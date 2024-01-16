package mac

import (
	"fmt"

	"github.com/adrianriobo/qenvs/pkg/manager"
	qenvsContext "github.com/adrianriobo/qenvs/pkg/manager/context"
	"github.com/adrianriobo/qenvs/pkg/provider/aws"
	"github.com/adrianriobo/qenvs/pkg/provider/aws/data"
	"github.com/adrianriobo/qenvs/pkg/provider/util/output"
	"github.com/adrianriobo/qenvs/pkg/util/logging"
	resourcesUtil "github.com/adrianriobo/qenvs/pkg/util/resources"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// this creates the stack for the dedicated host
func (r *MacRequest) createDedicatedHost() (*HostInformation, error) {
	logging.Debugf("creating dedicated host for mac machine %s", r.Architecture)
	backedURL := fmt.Sprintf("%s/%s", qenvsContext.GetBackedURL(), qenvsContext.GetID())
	cs := manager.Stack{
		StackName:   qenvsContext.GetStackInstanceName(stackDedicatedHost),
		ProjectName: qenvsContext.GetInstanceName(),
		BackedURL:   backedURL,
		ProviderCredentials: aws.GetClouProviderCredentials(
			map[string]string{
				aws.CONFIG_AWS_REGION: *r.Region}),
		DeployFunc: r.deployerDedicatedHost,
	}
	sr, _ := manager.UpStack(cs)
	dhID, _, err := r.manageResultsDedicatedHost(sr)
	if err != nil {
		return nil, err
	}
	logging.Debugf("dedicated host with host id %s has been created successfully", *dhID)
	host, err := data.GetDedicatedHost(*dhID)
	if err != nil {
		return nil, err
	}
	return &HostInformation{
		BackedURL: &backedURL,
		Region:    r.Region,
		Host:      host,
	}, nil
}

// this function will create the dedicated host resource
func (r *MacRequest) deployerDedicatedHost(ctx *pulumi.Context) (err error) {
	ctx.Export(fmt.Sprintf("%s-%s", r.Prefix, outputRegion), pulumi.String(*r.Region))
	dh, err := ec2.NewDedicatedHost(ctx,
		resourcesUtil.GetResourceName(r.Prefix, awsMacMachineID, "dh"),
		&ec2.DedicatedHostArgs{
			AutoPlacement:    pulumi.String("off"),
			AvailabilityZone: pulumi.String(*r.AvailabilityZone),
			InstanceType:     pulumi.String(macTypesByArch[r.Architecture]),
			Tags: qenvsContext.ResourceTagsWithCustom(
				map[string]string{
					backedURLTagName: fmt.Sprintf("%s/%s", qenvsContext.GetBackedURL(), qenvsContext.GetID()),
					archTagName:      r.Architecture,
				}),
		})
	ctx.Export(fmt.Sprintf("%s-%s", r.Prefix, outputDedicatedHostID), dh.ID())
	ctx.Export(fmt.Sprintf("%s-%s", r.Prefix, outputDedicatedHostAZ), pulumi.String(*r.AvailabilityZone))
	if err != nil {
		return err
	}
	return nil
}

// results for dedicated host it will return dedicatedhost ID and dedicatedhost AZ
// also write results to files on the target folder
func (r *MacRequest) manageResultsDedicatedHost(stackResult auto.UpResult) (*string, *string, error) {
	if err := output.Write(stackResult, qenvsContext.GetResultsOutputPath(), map[string]string{
		fmt.Sprintf("%s-%s", r.Prefix, outputDedicatedHostID): "dedicatedHostID",
	}); err != nil {
		return nil, nil, err
	}
	dhID, ok := stackResult.Outputs[fmt.Sprintf("%s-%s", r.Prefix, outputDedicatedHostID)].Value.(string)
	if !ok {
		return nil, nil, fmt.Errorf("error getting dedicated host ID")
	}
	dhAZ, ok := stackResult.Outputs[fmt.Sprintf("%s-%s", r.Prefix, outputDedicatedHostAZ)].Value.(string)
	if !ok {
		return nil, nil, fmt.Errorf("error getting dedicated host AZ")
	}
	return &dhID, &dhAZ, nil
}
