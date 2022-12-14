package compute

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi-tls/sdk/v4/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	securityGroup "github.com/adrianriobo/qenvs/pkg/infra/aws/services/ec2/security-group"
	supportMatrix "github.com/adrianriobo/qenvs/pkg/infra/aws/support-matrix"
)

// Composable resources and information provided
// to create a compute asset
type Request struct {
	ProjecName string
	Specs      *supportMatrix.SupportedHost
	Public     bool
	BastionSG  *ec2.SecurityGroup
	KeyPair    *ec2.KeyPair
	// contains value if key is created within this module
	PublicKeyOpenssh  pulumi.StringOutput
	VPC               *ec2.Vpc
	AvailabilityZones []string
	Region            string
	Subnets           []*ec2.Subnet
	SpotPrice         string
}

// Methods to be implemented by any specific type of compute supported
// these methods implement specifics on each different type
type ComputeRequest interface {
	// Get Request information
	GetRequest() *Request
	// Get ami value for the compute
	GetAMI(ctx *pulumi.Context) (*ec2.LookupAmiResult, error)
	// Get userdata if any
	GetUserdata(ctx *pulumi.Context) (pulumi.StringPtrInput, error)
	// Create dedicated host if compute requires it
	GetDedicatedHost(ctx *pulumi.Context) (*ec2.DedicatedHost, error)
	// In case a host has any specific ingress rule, this ingress rule will take effect on a SG
	// with egress to all in case of specific SG it is required to define them within CustomSecurityGroups
	CustomIngressRules() []securityGroup.IngressRules
	CustomSecurityGroups(ctx *pulumi.Context) ([]*ec2.SecurityGroup, error)
	// Get script to be executed after initalization (not suited for userdata)
	GetPostScript(ctx *pulumi.Context) (pulumi.StringPtrInput, error)
	// Command to be executed to ensure the provisioned is finish with fully setup
	ReadinessCommand() string
	// Create function to get a compute based on request
	Create(ctx *pulumi.Context, computeRequested ComputeRequest) (*Compute, error)
}

// Related resources created within the compute asset
type Compute struct {
	Specs      *supportMatrix.SupportedHost
	Name       string
	Username   string
	Instance   *ec2.Instance
	InstanceIP pulumi.StringOutput
	AWSKeyPair *ec2.KeyPair
	SG         []*ec2.SecurityGroup
	// contains value if key is created within this module
	PrivateKey *tls.PrivateKey
}
