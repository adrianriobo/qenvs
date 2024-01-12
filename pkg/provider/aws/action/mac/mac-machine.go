package mac

import (
	_ "embed"
	"fmt"

	"github.com/adrianriobo/qenvs/pkg/manager"
	qenvsContext "github.com/adrianriobo/qenvs/pkg/manager/context"
	infra "github.com/adrianriobo/qenvs/pkg/provider"
	"github.com/adrianriobo/qenvs/pkg/provider/aws"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/data"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/modules/bastion"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/modules/network"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/services/ec2/ami"
	qEC2 "github.com/adrianriobo/qenvs/pkg/provider/aws/services/ec2/compute"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/services/ec2/keypair"
	securityGroup "github.com/adrianriobo/qenvs/pkg/provider/aws/services/ec2/security-group"
	"github.com/adrianriobo/qenvs/pkg/provider/util/command"
	"github.com/adrianriobo/qenvs/pkg/provider/util/output"
	"github.com/adrianriobo/qenvs/pkg/provider/util/security"
	"github.com/adrianriobo/qenvs/pkg/util"
	"github.com/adrianriobo/qenvs/pkg/util/file"
	"github.com/adrianriobo/qenvs/pkg/util/logging"
	resourcesUtil "github.com/adrianriobo/qenvs/pkg/util/resources"

	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-command/sdk/go/command/remote"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"golang.org/x/exp/slices"
)

//go:embed bootstrap.sh
var BootstrapScript []byte

type userDataValues struct {
	Username string
	Password string
}

type locked struct {
	pulumi.ResourceState
	Lock bool
}

func isMachineLocked(prefix string, h HostInformation) (bool, error) {
	s, err := manager.CheckStack(manager.Stack{
		StackName:   qenvsContext.GetStackInstanceName(stackMacMachine),
		ProjectName: qenvsContext.GetInstanceName(),
		BackedURL:   *h.BackedURL,
		ProviderCredentials: aws.GetClouProviderCredentials(
			map[string]string{
				aws.CONFIG_AWS_REGION: *h.Region}),
	})
	if err != nil {
		return false, err
	}
	outputs, err := manager.GetOutputs(s)
	if err != nil {
		return false, err
	}
	return outputs[fmt.Sprintf("%s-%s", prefix, outputLock)].Value.(bool), nil
}

// This function will use the information from the
// dedicated host holding the mac machine will check if stack exists
// if exists will get the lock value from it
func replaceMachine(prefix string, h HostInformation) error {
	machineLocked, err := isMachineLocked(prefix, h)
	if err != nil {
		return err
	}
	if !machineLocked {
		instances, err := data.GetInstanceByRegion(data.InstanceResquest{
			Tags: qenvsContext.GetTags(),
		}, *h.Region)
		if err != nil {
			return err
		}
		logging.Debugf("%v", instances)
		// TODO need to change the AMI ID to pick it based on labes? + copy across regions
		t, err := qEC2.ReplaceRootVolume(
			qEC2.ReplaceRootVolumeRequest{
				Region:     *h.Region,
				InstanceID: *instances[0].InstanceId,
				AMIID:      "ami-04aaa6dc45616cb76",
			})
		if err != nil {
			return err
		}
		logging.Debugf("%v", t)
	}
	return nil
}

// Release will set the lock as false
func (r *MacRequest) releaseLocked(h *HostInformation) error {
	// Check the stack
	r.Lock = false
	sr, _ := manager.UpStack(manager.Stack{
		StackName:   qenvsContext.GetStackInstanceName(stackMacMachine),
		ProjectName: qenvsContext.GetInstanceName(),
		BackedURL:   *h.BackedURL,
		ProviderCredentials: aws.GetClouProviderCredentials(
			map[string]string{
				aws.CONFIG_AWS_REGION: *h.Region}),
		DeployFunc: r.deployerMachine,
	})
	return r.manageResultsMachine(sr)
}

// this creates the stack for the mac machine
func (r *MacRequest) createMacMachine(dhAZ *string, host *ec2Types.Host) error {
	r.AvailabilityZone = *dhAZ
	r.dedicatedHost = host
	// TODO we need to check wich host we will use
	// some logic based on states for dh + locks on machines on the dh
	r.Lock = true
	// If request does not set onlyHost we will create the mac machine
	if len(r.AvailabilityZone) == 0 {
		r.AvailabilityZone = *r.dedicatedHost.AvailabilityZone
	}
	region := r.AvailabilityZone[:len(r.AvailabilityZone)-1]
	i := slices.IndexFunc(
		r.dedicatedHost.Tags,
		func(t ec2Types.Tag) bool {
			return *t.Key == backedURLTagName
		})
	dhBackedURL := r.dedicatedHost.Tags[i].Value
	cs := manager.Stack{
		StackName:   qenvsContext.GetStackInstanceName(stackMacMachine),
		ProjectName: qenvsContext.GetInstanceName(),
		BackedURL:   *dhBackedURL,
		ProviderCredentials: aws.GetClouProviderCredentials(
			map[string]string{
				aws.CONFIG_AWS_REGION: region}),
		DeployFunc: r.deployerMachine,
	}
	sr, _ := manager.UpStack(cs)
	return r.manageResultsMachine(sr)
}

// this creates the stack for the mac machine
func (r *MacRequest) createAirgapMacMachine(dhAZ *string, host *ec2Types.Host) error {
	r.airgapPhaseConnectivity = network.ON
	err := r.createMacMachine(dhAZ, host)
	if err != nil {
		return nil
	}
	r.airgapPhaseConnectivity = network.OFF
	return r.createMacMachine(dhAZ, host)
}

// Main function to deploy all requried resources to azure
func (r *MacRequest) deployerMachine(ctx *pulumi.Context) error {
	// Export the region for delete
	ctx.Export(fmt.Sprintf("%s-%s", r.Prefix, outputRegion), pulumi.String(r.Region))
	// Lookup AMI
	ami, err := ami.GetAMIByName(ctx,
		fmt.Sprintf(amiRegex, r.Version),
		amiOwner,
		map[string]string{
			"architecture": awsArchIDbyArch[r.Architecture]})
	if err != nil {
		return err
	}
	nr := network.NetworkRequest{
		Prefix:                  r.Prefix,
		ID:                      awsMacMachineID,
		Region:                  r.Region,
		AZ:                      r.AvailabilityZone,
		Airgap:                  r.Airgap,
		AirgapPhaseConnectivity: r.airgapPhaseConnectivity,
	}
	vpc, targetSubnet, targetRouteTableAssociation, bastion, _, err := nr.Network(ctx)
	if err != nil {
		return err
	}
	// Create Keypair
	kpr := keypair.KeyPairRequest{
		Name: resourcesUtil.GetResourceName(
			r.Prefix, awsMacMachineID, "pk")}
	keyResources, err := kpr.Create(ctx)
	if err != nil {
		return err
	}
	ctx.Export(fmt.Sprintf("%s-%s", r.Prefix, outputUserPrivateKey),
		keyResources.PrivateKey.PrivateKeyPem)
	// Security groups
	securityGroups, err := r.securityGroups(ctx, vpc)
	if err != nil {
		return err
	}
	// Create instance
	i, err := r.instance(ctx, targetSubnet, ami, keyResources, securityGroups)
	if err != nil {
		return err
	}
	ctx.Export(fmt.Sprintf("%s-%s", r.Prefix, outputUsername),
		pulumi.String(defaultUsername))
	if r.Airgap {
		ctx.Export(fmt.Sprintf("%s-%s", r.Prefix, outputHost),
			i.PrivateIp)
	} else {
		ctx.Export(fmt.Sprintf("%s-%s", r.Prefix, outputHost),
			i.PublicIp)
	}
	// Bootstrap script
	bSDependecies := []pulumi.Resource{i}
	if bastion != nil {
		bSDependecies = append(bSDependecies,
			[]pulumi.Resource{bastion.Instance, targetRouteTableAssociation}...)
	}
	bc, userPassword, err := r.bootstrapscript(
		ctx, i, keyResources.PrivateKey, bastion, bSDependecies)
	if err != nil {
		return err
	}
	ctx.Export(fmt.Sprintf("%s-%s", r.Prefix, outputUserPassword), userPassword.Result)
	// Create a lock on the machine
	if err := newMachineLock(ctx,
		resourcesUtil.GetResourceName(
			r.Prefix, awsMacMachineID, "mac-lock"), r.Lock); err != nil {
		return err
	}
	ctx.Export(fmt.Sprintf("%s-%s", r.Prefix, outputLock), pulumi.Bool(r.Lock))
	// Check if replace is needed
	if r.Replace {
		// use the replace
		archID := awsArchIDbyArch[r.Architecture]
		macAMIKey := fmt.Sprintf("%s-%s", archID, r.Version)
		macAMIID := macAMIs[macAMIKey]
		logging.Debugf("%s", macAMIID)
	}
	return r.readiness(ctx, i, keyResources.PrivateKey, bastion, []pulumi.Resource{bc})
}

// Write exported values in context to files o a selected target folder
func (r *MacRequest) manageResultsMachine(stackResult auto.UpResult) error {
	results := map[string]string{
		fmt.Sprintf("%s-%s", r.Prefix, outputUsername):       "username",
		fmt.Sprintf("%s-%s", r.Prefix, outputUserPassword):   "userpassword",
		fmt.Sprintf("%s-%s", r.Prefix, outputUserPrivateKey): "id_rsa",
		fmt.Sprintf("%s-%s", r.Prefix, outputHost):           "host",
	}
	if r.Airgap {
		err := bastion.WriteOutputs(stackResult, r.Prefix, qenvsContext.GetResultsOutputPath())
		if err != nil {
			return err
		}
	}
	return output.Write(stackResult, qenvsContext.GetResultsOutputPath(), results)
}

// security group for mac machine with ingress rules for ssh and vnc
func (r *MacRequest) securityGroups(ctx *pulumi.Context,
	vpc *ec2.Vpc) (pulumi.StringArray, error) {
	// ingress for ssh access from 0.0.0.0
	sshIngressRule := securityGroup.SSH_TCP
	sshIngressRule.CidrBlocks = infra.NETWORKING_CIDR_ANY_IPV4
	// ingress for vnc access from 0.0.0.0
	vncIngressRule := securityGroup.IngressRules{
		Description: fmt.Sprintf("VNC port for %s", awsMacMachineID),
		FromPort:    vncDefaultPort,
		ToPort:      vncDefaultPort,
		Protocol:    "tcp",
		CidrBlocks:  infra.NETWORKING_CIDR_ANY_IPV4,
	}
	// Create SG with ingress rules
	sg, err := securityGroup.SGRequest{
		Name:        resourcesUtil.GetResourceName(r.Prefix, awsMacMachineID, "sg"),
		VPC:         vpc,
		Description: fmt.Sprintf("sg for %s", awsMacMachineID),
		IngressRules: []securityGroup.IngressRules{
			sshIngressRule, vncIngressRule},
	}.Create(ctx)
	if err != nil {
		return nil, err
	}
	// Convert to an array of IDs
	sgs := util.ArrayConvert([]*ec2.SecurityGroup{sg.SG},
		func(sg *ec2.SecurityGroup) pulumi.StringInput {
			return sg.ID()
		})
	return pulumi.StringArray(sgs[:]), nil
}

// Create the mac instance
func (r *MacRequest) instance(ctx *pulumi.Context,
	subnet *ec2.Subnet,
	ami *ec2.LookupAmiResult,
	keyResources *keypair.KeyPairResources,
	securityGroups pulumi.StringArray,
) (*ec2.Instance, error) {
	instanceArgs := ec2.InstanceArgs{
		HostId:                   pulumi.String(*r.dedicatedHost.HostId),
		SubnetId:                 subnet.ID(),
		Ami:                      pulumi.String(ami.Id),
		InstanceType:             pulumi.String(macTypesByArch[r.Architecture]),
		KeyName:                  keyResources.AWSKeyPair.KeyName,
		AssociatePublicIpAddress: pulumi.Bool(true),
		VpcSecurityGroupIds:      securityGroups,
		RootBlockDevice: ec2.InstanceRootBlockDeviceArgs{
			VolumeSize: pulumi.Int(diskSize),
		},
		Tags: qenvsContext.ResourceTags(),
	}
	if r.Airgap {
		instanceArgs.AssociatePublicIpAddress = pulumi.Bool(false)
	}
	return ec2.NewInstance(ctx,
		resourcesUtil.GetResourceName(r.Prefix, awsMacMachineID, "instance"),
		&instanceArgs)
}

func (r *MacRequest) bootstrapscript(ctx *pulumi.Context,
	m *ec2.Instance,
	mk *tls.PrivateKey,
	b *bastion.Bastion,
	dependecies []pulumi.Resource) (
	*remote.Command,
	*random.RandomPassword,
	error) {
	// Bootstrap script
	remoteCommand, userPassword, err := r.getBootstrapScript(ctx)
	if err != nil {
		return nil, nil, err
	}
	rc, err := remote.NewCommand(ctx,
		resourcesUtil.GetResourceName(r.Prefix, awsMacMachineID, "bootstrap-cmd"),
		&remote.CommandArgs{
			Connection: remoteCommandArgs(m, mk, b),
			Create:     remoteCommand,
			Update:     remoteCommand,
		}, pulumi.Timeouts(
			&pulumi.CustomTimeouts{
				Create: remoteTimeout,
				Update: remoteTimeout}),
		pulumi.DependsOn(dependecies))
	return rc, userPassword, err
}

// fuction will return the bootstrap script which will be execute on the mac machine
// during the start of the machine
func (r *MacRequest) getBootstrapScript(ctx *pulumi.Context) (
	pulumi.StringPtrInput,
	*random.RandomPassword,
	error) {
	password, err := security.CreatePassword(ctx,
		resourcesUtil.GetResourceName(r.Prefix, awsMacMachineID, "passwd"))
	if err != nil {
		return nil, nil, err
	}

	postscript := password.Result.ApplyT(func(password string) (string, error) {
		return file.Template(
			userDataValues{
				defaultUsername,
				password},
			resourcesUtil.GetResourceName(r.Prefix, awsMacMachineID, "mac-bootstrap"),
			string(BootstrapScript[:]))

	}).(pulumi.StringOutput)
	return postscript, password, nil
}

func (r *MacRequest) readiness(ctx *pulumi.Context,
	m *ec2.Instance,
	mk *tls.PrivateKey,
	b *bastion.Bastion,
	dependecies []pulumi.Resource) error {
	_, err := remote.NewCommand(ctx,
		resourcesUtil.GetResourceName(r.Prefix, awsMacMachineID, "readiness-cmd"),
		&remote.CommandArgs{
			Connection: remoteCommandArgs(m, mk, b),
			Create:     pulumi.String(command.CommandPing),
			Update:     pulumi.String(command.CommandPing),
		}, pulumi.Timeouts(
			&pulumi.CustomTimeouts{
				Create: remoteTimeout,
				Update: remoteTimeout}),
		pulumi.DependsOn(dependecies))
	return err
}

// helper function to set the connection args
// based on bastion or direct connection to target host
func remoteCommandArgs(
	m *ec2.Instance,
	mk *tls.PrivateKey,
	b *bastion.Bastion) remote.ConnectionArgs {
	ca := remote.ConnectionArgs{
		Host:           m.PublicIp,
		PrivateKey:     mk.PrivateKeyOpenssh,
		User:           pulumi.String(defaultUsername),
		Port:           pulumi.Float64(defaultSSHPort),
		DialErrorLimit: pulumi.Int(-1)}
	if b != nil {
		// If airgap set the private IP for host
		// And bastion details
		ca.Host = m.PrivateIp
		ca.Proxy = remote.ProxyConnectionArgs{
			Host:           b.Instance.PublicIp,
			PrivateKey:     b.PrivateKey.PrivateKeyOpenssh,
			User:           pulumi.String(b.Usarname),
			Port:           pulumi.Float64(b.Port),
			DialErrorLimit: pulumi.Int(-1)}

	}
	return ca
}

// This will create mark to see if machine is lock or free
func newMachineLock(ctx *pulumi.Context, name string, lockedValue bool, opts ...pulumi.ResourceOption) error {
	return ctx.RegisterComponentResource(
		urnLock,
		name,
		&locked{
			Lock: lockedValue,
		},
		opts...)
}
