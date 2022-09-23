// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewNetworkAcl(ctx, "main", &ec2.NetworkAclArgs{
//				VpcId: pulumi.Any(aws_vpc.Main.Id),
//				Egress: ec2.NetworkAclEgressArray{
//					&ec2.NetworkAclEgressArgs{
//						Protocol:  pulumi.String("tcp"),
//						RuleNo:    pulumi.Int(200),
//						Action:    pulumi.String("allow"),
//						CidrBlock: pulumi.String("10.3.0.0/18"),
//						FromPort:  pulumi.Int(443),
//						ToPort:    pulumi.Int(443),
//					},
//				},
//				Ingress: ec2.NetworkAclIngressArray{
//					&ec2.NetworkAclIngressArgs{
//						Protocol:  pulumi.String("tcp"),
//						RuleNo:    pulumi.Int(100),
//						Action:    pulumi.String("allow"),
//						CidrBlock: pulumi.String("10.3.0.0/18"),
//						FromPort:  pulumi.Int(80),
//						ToPort:    pulumi.Int(80),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("main"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Network ACLs can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/networkAcl:NetworkAcl main acl-7aaabd18
//
// ```
type NetworkAcl struct {
	pulumi.CustomResourceState

	// The ARN of the network ACL
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies an egress rule. Parameters defined below.
	Egress NetworkAclEgressArrayOutput `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress NetworkAclIngressArrayOutput `pulumi:"ingress"`
	// The ID of the AWS account that owns the network ACL.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A list of Subnet IDs to apply the ACL to
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The ID of the associated VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewNetworkAcl registers a new resource with the given unique name, arguments, and options.
func NewNetworkAcl(ctx *pulumi.Context,
	name string, args *NetworkAclArgs, opts ...pulumi.ResourceOption) (*NetworkAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource NetworkAcl
	err := ctx.RegisterResource("aws:ec2/networkAcl:NetworkAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkAcl gets an existing NetworkAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkAclState, opts ...pulumi.ResourceOption) (*NetworkAcl, error) {
	var resource NetworkAcl
	err := ctx.ReadResource("aws:ec2/networkAcl:NetworkAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkAcl resources.
type networkAclState struct {
	// The ARN of the network ACL
	Arn *string `pulumi:"arn"`
	// Specifies an egress rule. Parameters defined below.
	Egress []NetworkAclEgress `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress []NetworkAclIngress `pulumi:"ingress"`
	// The ID of the AWS account that owns the network ACL.
	OwnerId *string `pulumi:"ownerId"`
	// A list of Subnet IDs to apply the ACL to
	SubnetIds []string `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The ID of the associated VPC.
	VpcId *string `pulumi:"vpcId"`
}

type NetworkAclState struct {
	// The ARN of the network ACL
	Arn pulumi.StringPtrInput
	// Specifies an egress rule. Parameters defined below.
	Egress NetworkAclEgressArrayInput
	// Specifies an ingress rule. Parameters defined below.
	Ingress NetworkAclIngressArrayInput
	// The ID of the AWS account that owns the network ACL.
	OwnerId pulumi.StringPtrInput
	// A list of Subnet IDs to apply the ACL to
	SubnetIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The ID of the associated VPC.
	VpcId pulumi.StringPtrInput
}

func (NetworkAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclState)(nil)).Elem()
}

type networkAclArgs struct {
	// Specifies an egress rule. Parameters defined below.
	Egress []NetworkAclEgress `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress []NetworkAclIngress `pulumi:"ingress"`
	// A list of Subnet IDs to apply the ACL to
	SubnetIds []string `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the associated VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a NetworkAcl resource.
type NetworkAclArgs struct {
	// Specifies an egress rule. Parameters defined below.
	Egress NetworkAclEgressArrayInput
	// Specifies an ingress rule. Parameters defined below.
	Ingress NetworkAclIngressArrayInput
	// A list of Subnet IDs to apply the ACL to
	SubnetIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The ID of the associated VPC.
	VpcId pulumi.StringInput
}

func (NetworkAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclArgs)(nil)).Elem()
}

type NetworkAclInput interface {
	pulumi.Input

	ToNetworkAclOutput() NetworkAclOutput
	ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput
}

func (*NetworkAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAcl)(nil)).Elem()
}

func (i *NetworkAcl) ToNetworkAclOutput() NetworkAclOutput {
	return i.ToNetworkAclOutputWithContext(context.Background())
}

func (i *NetworkAcl) ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclOutput)
}

// NetworkAclArrayInput is an input type that accepts NetworkAclArray and NetworkAclArrayOutput values.
// You can construct a concrete instance of `NetworkAclArrayInput` via:
//
//	NetworkAclArray{ NetworkAclArgs{...} }
type NetworkAclArrayInput interface {
	pulumi.Input

	ToNetworkAclArrayOutput() NetworkAclArrayOutput
	ToNetworkAclArrayOutputWithContext(context.Context) NetworkAclArrayOutput
}

type NetworkAclArray []NetworkAclInput

func (NetworkAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkAcl)(nil)).Elem()
}

func (i NetworkAclArray) ToNetworkAclArrayOutput() NetworkAclArrayOutput {
	return i.ToNetworkAclArrayOutputWithContext(context.Background())
}

func (i NetworkAclArray) ToNetworkAclArrayOutputWithContext(ctx context.Context) NetworkAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclArrayOutput)
}

// NetworkAclMapInput is an input type that accepts NetworkAclMap and NetworkAclMapOutput values.
// You can construct a concrete instance of `NetworkAclMapInput` via:
//
//	NetworkAclMap{ "key": NetworkAclArgs{...} }
type NetworkAclMapInput interface {
	pulumi.Input

	ToNetworkAclMapOutput() NetworkAclMapOutput
	ToNetworkAclMapOutputWithContext(context.Context) NetworkAclMapOutput
}

type NetworkAclMap map[string]NetworkAclInput

func (NetworkAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkAcl)(nil)).Elem()
}

func (i NetworkAclMap) ToNetworkAclMapOutput() NetworkAclMapOutput {
	return i.ToNetworkAclMapOutputWithContext(context.Background())
}

func (i NetworkAclMap) ToNetworkAclMapOutputWithContext(ctx context.Context) NetworkAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclMapOutput)
}

type NetworkAclOutput struct{ *pulumi.OutputState }

func (NetworkAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAcl)(nil)).Elem()
}

func (o NetworkAclOutput) ToNetworkAclOutput() NetworkAclOutput {
	return o
}

func (o NetworkAclOutput) ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput {
	return o
}

// The ARN of the network ACL
func (o NetworkAclOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies an egress rule. Parameters defined below.
func (o NetworkAclOutput) Egress() NetworkAclEgressArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) NetworkAclEgressArrayOutput { return v.Egress }).(NetworkAclEgressArrayOutput)
}

// Specifies an ingress rule. Parameters defined below.
func (o NetworkAclOutput) Ingress() NetworkAclIngressArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) NetworkAclIngressArrayOutput { return v.Ingress }).(NetworkAclIngressArrayOutput)
}

// The ID of the AWS account that owns the network ACL.
func (o NetworkAclOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// A list of Subnet IDs to apply the ACL to
func (o NetworkAclOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o NetworkAclOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o NetworkAclOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The ID of the associated VPC.
func (o NetworkAclOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type NetworkAclArrayOutput struct{ *pulumi.OutputState }

func (NetworkAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkAcl)(nil)).Elem()
}

func (o NetworkAclArrayOutput) ToNetworkAclArrayOutput() NetworkAclArrayOutput {
	return o
}

func (o NetworkAclArrayOutput) ToNetworkAclArrayOutputWithContext(ctx context.Context) NetworkAclArrayOutput {
	return o
}

func (o NetworkAclArrayOutput) Index(i pulumi.IntInput) NetworkAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkAcl {
		return vs[0].([]*NetworkAcl)[vs[1].(int)]
	}).(NetworkAclOutput)
}

type NetworkAclMapOutput struct{ *pulumi.OutputState }

func (NetworkAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkAcl)(nil)).Elem()
}

func (o NetworkAclMapOutput) ToNetworkAclMapOutput() NetworkAclMapOutput {
	return o
}

func (o NetworkAclMapOutput) ToNetworkAclMapOutputWithContext(ctx context.Context) NetworkAclMapOutput {
	return o
}

func (o NetworkAclMapOutput) MapIndex(k pulumi.StringInput) NetworkAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkAcl {
		return vs[0].(map[string]*NetworkAcl)[vs[1].(string)]
	}).(NetworkAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclInput)(nil)).Elem(), &NetworkAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclArrayInput)(nil)).Elem(), NetworkAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclMapInput)(nil)).Elem(), NetworkAclMap{})
	pulumi.RegisterOutputType(NetworkAclOutput{})
	pulumi.RegisterOutputType(NetworkAclArrayOutput{})
	pulumi.RegisterOutputType(NetworkAclMapOutput{})
}
