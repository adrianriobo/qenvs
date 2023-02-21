// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to accept a pending VPC Endpoint Connection accept request to VPC Endpoint Service.
//
// ## Example Usage
// ### Accept cross-account request
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
//			exampleVpcEndpointService, err := ec2.NewVpcEndpointService(ctx, "exampleVpcEndpointService", &ec2.VpcEndpointServiceArgs{
//				AcceptanceRequired: pulumi.Bool(false),
//				NetworkLoadBalancerArns: pulumi.StringArray{
//					aws_lb.Example.Arn,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleVpcEndpoint, err := ec2.NewVpcEndpoint(ctx, "exampleVpcEndpoint", &ec2.VpcEndpointArgs{
//				VpcId:             pulumi.Any(aws_vpc.Test_alternate.Id),
//				ServiceName:       pulumi.Any(aws_vpc_endpoint_service.Test.Service_name),
//				VpcEndpointType:   pulumi.String("Interface"),
//				PrivateDnsEnabled: pulumi.Bool(false),
//				SecurityGroupIds: pulumi.StringArray{
//					aws_security_group.Test.Id,
//				},
//			}, pulumi.Provider("aws.alternate"))
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcEndpointConnectionAccepter(ctx, "exampleVpcEndpointConnectionAccepter", &ec2.VpcEndpointConnectionAccepterArgs{
//				VpcEndpointServiceId: exampleVpcEndpointService.ID(),
//				VpcEndpointId:        exampleVpcEndpoint.ID(),
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
// VPC Endpoint Services can be imported using ID of the connection, which is the `VPC Endpoint Service ID` and `VPC Endpoint ID` separated by underscore (`_`). e.g.
//
// ```sh
//
//	$ pulumi import aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter foo vpce-svc-0f97a19d3fa8220bc_vpce-010601a6db371e263
//
// ```
type VpcEndpointConnectionAccepter struct {
	pulumi.CustomResourceState

	// AWS VPC Endpoint ID.
	VpcEndpointId pulumi.StringOutput `pulumi:"vpcEndpointId"`
	// AWS VPC Endpoint Service ID.
	VpcEndpointServiceId pulumi.StringOutput `pulumi:"vpcEndpointServiceId"`
	// State of the VPC Endpoint.
	VpcEndpointState pulumi.StringOutput `pulumi:"vpcEndpointState"`
}

// NewVpcEndpointConnectionAccepter registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointConnectionAccepter(ctx *pulumi.Context,
	name string, args *VpcEndpointConnectionAccepterArgs, opts ...pulumi.ResourceOption) (*VpcEndpointConnectionAccepter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'VpcEndpointId'")
	}
	if args.VpcEndpointServiceId == nil {
		return nil, errors.New("invalid value for required argument 'VpcEndpointServiceId'")
	}
	var resource VpcEndpointConnectionAccepter
	err := ctx.RegisterResource("aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointConnectionAccepter gets an existing VpcEndpointConnectionAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointConnectionAccepter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointConnectionAccepterState, opts ...pulumi.ResourceOption) (*VpcEndpointConnectionAccepter, error) {
	var resource VpcEndpointConnectionAccepter
	err := ctx.ReadResource("aws:ec2/vpcEndpointConnectionAccepter:VpcEndpointConnectionAccepter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointConnectionAccepter resources.
type vpcEndpointConnectionAccepterState struct {
	// AWS VPC Endpoint ID.
	VpcEndpointId *string `pulumi:"vpcEndpointId"`
	// AWS VPC Endpoint Service ID.
	VpcEndpointServiceId *string `pulumi:"vpcEndpointServiceId"`
	// State of the VPC Endpoint.
	VpcEndpointState *string `pulumi:"vpcEndpointState"`
}

type VpcEndpointConnectionAccepterState struct {
	// AWS VPC Endpoint ID.
	VpcEndpointId pulumi.StringPtrInput
	// AWS VPC Endpoint Service ID.
	VpcEndpointServiceId pulumi.StringPtrInput
	// State of the VPC Endpoint.
	VpcEndpointState pulumi.StringPtrInput
}

func (VpcEndpointConnectionAccepterState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointConnectionAccepterState)(nil)).Elem()
}

type vpcEndpointConnectionAccepterArgs struct {
	// AWS VPC Endpoint ID.
	VpcEndpointId string `pulumi:"vpcEndpointId"`
	// AWS VPC Endpoint Service ID.
	VpcEndpointServiceId string `pulumi:"vpcEndpointServiceId"`
}

// The set of arguments for constructing a VpcEndpointConnectionAccepter resource.
type VpcEndpointConnectionAccepterArgs struct {
	// AWS VPC Endpoint ID.
	VpcEndpointId pulumi.StringInput
	// AWS VPC Endpoint Service ID.
	VpcEndpointServiceId pulumi.StringInput
}

func (VpcEndpointConnectionAccepterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointConnectionAccepterArgs)(nil)).Elem()
}

type VpcEndpointConnectionAccepterInput interface {
	pulumi.Input

	ToVpcEndpointConnectionAccepterOutput() VpcEndpointConnectionAccepterOutput
	ToVpcEndpointConnectionAccepterOutputWithContext(ctx context.Context) VpcEndpointConnectionAccepterOutput
}

func (*VpcEndpointConnectionAccepter) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointConnectionAccepter)(nil)).Elem()
}

func (i *VpcEndpointConnectionAccepter) ToVpcEndpointConnectionAccepterOutput() VpcEndpointConnectionAccepterOutput {
	return i.ToVpcEndpointConnectionAccepterOutputWithContext(context.Background())
}

func (i *VpcEndpointConnectionAccepter) ToVpcEndpointConnectionAccepterOutputWithContext(ctx context.Context) VpcEndpointConnectionAccepterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointConnectionAccepterOutput)
}

// VpcEndpointConnectionAccepterArrayInput is an input type that accepts VpcEndpointConnectionAccepterArray and VpcEndpointConnectionAccepterArrayOutput values.
// You can construct a concrete instance of `VpcEndpointConnectionAccepterArrayInput` via:
//
//	VpcEndpointConnectionAccepterArray{ VpcEndpointConnectionAccepterArgs{...} }
type VpcEndpointConnectionAccepterArrayInput interface {
	pulumi.Input

	ToVpcEndpointConnectionAccepterArrayOutput() VpcEndpointConnectionAccepterArrayOutput
	ToVpcEndpointConnectionAccepterArrayOutputWithContext(context.Context) VpcEndpointConnectionAccepterArrayOutput
}

type VpcEndpointConnectionAccepterArray []VpcEndpointConnectionAccepterInput

func (VpcEndpointConnectionAccepterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointConnectionAccepter)(nil)).Elem()
}

func (i VpcEndpointConnectionAccepterArray) ToVpcEndpointConnectionAccepterArrayOutput() VpcEndpointConnectionAccepterArrayOutput {
	return i.ToVpcEndpointConnectionAccepterArrayOutputWithContext(context.Background())
}

func (i VpcEndpointConnectionAccepterArray) ToVpcEndpointConnectionAccepterArrayOutputWithContext(ctx context.Context) VpcEndpointConnectionAccepterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointConnectionAccepterArrayOutput)
}

// VpcEndpointConnectionAccepterMapInput is an input type that accepts VpcEndpointConnectionAccepterMap and VpcEndpointConnectionAccepterMapOutput values.
// You can construct a concrete instance of `VpcEndpointConnectionAccepterMapInput` via:
//
//	VpcEndpointConnectionAccepterMap{ "key": VpcEndpointConnectionAccepterArgs{...} }
type VpcEndpointConnectionAccepterMapInput interface {
	pulumi.Input

	ToVpcEndpointConnectionAccepterMapOutput() VpcEndpointConnectionAccepterMapOutput
	ToVpcEndpointConnectionAccepterMapOutputWithContext(context.Context) VpcEndpointConnectionAccepterMapOutput
}

type VpcEndpointConnectionAccepterMap map[string]VpcEndpointConnectionAccepterInput

func (VpcEndpointConnectionAccepterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointConnectionAccepter)(nil)).Elem()
}

func (i VpcEndpointConnectionAccepterMap) ToVpcEndpointConnectionAccepterMapOutput() VpcEndpointConnectionAccepterMapOutput {
	return i.ToVpcEndpointConnectionAccepterMapOutputWithContext(context.Background())
}

func (i VpcEndpointConnectionAccepterMap) ToVpcEndpointConnectionAccepterMapOutputWithContext(ctx context.Context) VpcEndpointConnectionAccepterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointConnectionAccepterMapOutput)
}

type VpcEndpointConnectionAccepterOutput struct{ *pulumi.OutputState }

func (VpcEndpointConnectionAccepterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointConnectionAccepter)(nil)).Elem()
}

func (o VpcEndpointConnectionAccepterOutput) ToVpcEndpointConnectionAccepterOutput() VpcEndpointConnectionAccepterOutput {
	return o
}

func (o VpcEndpointConnectionAccepterOutput) ToVpcEndpointConnectionAccepterOutputWithContext(ctx context.Context) VpcEndpointConnectionAccepterOutput {
	return o
}

// AWS VPC Endpoint ID.
func (o VpcEndpointConnectionAccepterOutput) VpcEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnectionAccepter) pulumi.StringOutput { return v.VpcEndpointId }).(pulumi.StringOutput)
}

// AWS VPC Endpoint Service ID.
func (o VpcEndpointConnectionAccepterOutput) VpcEndpointServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnectionAccepter) pulumi.StringOutput { return v.VpcEndpointServiceId }).(pulumi.StringOutput)
}

// State of the VPC Endpoint.
func (o VpcEndpointConnectionAccepterOutput) VpcEndpointState() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnectionAccepter) pulumi.StringOutput { return v.VpcEndpointState }).(pulumi.StringOutput)
}

type VpcEndpointConnectionAccepterArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointConnectionAccepterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointConnectionAccepter)(nil)).Elem()
}

func (o VpcEndpointConnectionAccepterArrayOutput) ToVpcEndpointConnectionAccepterArrayOutput() VpcEndpointConnectionAccepterArrayOutput {
	return o
}

func (o VpcEndpointConnectionAccepterArrayOutput) ToVpcEndpointConnectionAccepterArrayOutputWithContext(ctx context.Context) VpcEndpointConnectionAccepterArrayOutput {
	return o
}

func (o VpcEndpointConnectionAccepterArrayOutput) Index(i pulumi.IntInput) VpcEndpointConnectionAccepterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointConnectionAccepter {
		return vs[0].([]*VpcEndpointConnectionAccepter)[vs[1].(int)]
	}).(VpcEndpointConnectionAccepterOutput)
}

type VpcEndpointConnectionAccepterMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointConnectionAccepterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointConnectionAccepter)(nil)).Elem()
}

func (o VpcEndpointConnectionAccepterMapOutput) ToVpcEndpointConnectionAccepterMapOutput() VpcEndpointConnectionAccepterMapOutput {
	return o
}

func (o VpcEndpointConnectionAccepterMapOutput) ToVpcEndpointConnectionAccepterMapOutputWithContext(ctx context.Context) VpcEndpointConnectionAccepterMapOutput {
	return o
}

func (o VpcEndpointConnectionAccepterMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointConnectionAccepterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointConnectionAccepter {
		return vs[0].(map[string]*VpcEndpointConnectionAccepter)[vs[1].(string)]
	}).(VpcEndpointConnectionAccepterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointConnectionAccepterInput)(nil)).Elem(), &VpcEndpointConnectionAccepter{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointConnectionAccepterArrayInput)(nil)).Elem(), VpcEndpointConnectionAccepterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointConnectionAccepterMapInput)(nil)).Elem(), VpcEndpointConnectionAccepterMap{})
	pulumi.RegisterOutputType(VpcEndpointConnectionAccepterOutput{})
	pulumi.RegisterOutputType(VpcEndpointConnectionAccepterArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointConnectionAccepterMapOutput{})
}
