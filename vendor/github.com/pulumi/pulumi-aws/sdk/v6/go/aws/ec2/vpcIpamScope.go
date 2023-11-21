// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a scope for AWS IPAM.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := aws.GetRegion(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			exampleVpcIpam, err := ec2.NewVpcIpam(ctx, "exampleVpcIpam", &ec2.VpcIpamArgs{
//				OperatingRegions: ec2.VpcIpamOperatingRegionArray{
//					&ec2.VpcIpamOperatingRegionArgs{
//						RegionName: *pulumi.String(current.Name),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcIpamScope(ctx, "exampleVpcIpamScope", &ec2.VpcIpamScopeArgs{
//				IpamId:      exampleVpcIpam.ID(),
//				Description: pulumi.String("Another Scope"),
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
// Using `pulumi import`, import IPAMs using the `scope_id`. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2/vpcIpamScope:VpcIpamScope example ipam-scope-0513c69f283d11dfb
//
// ```
type VpcIpamScope struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the scope.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description for the scope you're creating.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ARN of the IPAM for which you're creating this scope.
	IpamArn pulumi.StringOutput `pulumi:"ipamArn"`
	// The ID of the IPAM for which you're creating this scope.
	IpamId        pulumi.StringOutput `pulumi:"ipamId"`
	IpamScopeType pulumi.StringOutput `pulumi:"ipamScopeType"`
	// Defines if the scope is the default scope or not.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// The number of pools in the scope.
	PoolCount pulumi.IntOutput `pulumi:"poolCount"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVpcIpamScope registers a new resource with the given unique name, arguments, and options.
func NewVpcIpamScope(ctx *pulumi.Context,
	name string, args *VpcIpamScopeArgs, opts ...pulumi.ResourceOption) (*VpcIpamScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpamId == nil {
		return nil, errors.New("invalid value for required argument 'IpamId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcIpamScope
	err := ctx.RegisterResource("aws:ec2/vpcIpamScope:VpcIpamScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIpamScope gets an existing VpcIpamScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIpamScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIpamScopeState, opts ...pulumi.ResourceOption) (*VpcIpamScope, error) {
	var resource VpcIpamScope
	err := ctx.ReadResource("aws:ec2/vpcIpamScope:VpcIpamScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIpamScope resources.
type vpcIpamScopeState struct {
	// The Amazon Resource Name (ARN) of the scope.
	Arn *string `pulumi:"arn"`
	// A description for the scope you're creating.
	Description *string `pulumi:"description"`
	// The ARN of the IPAM for which you're creating this scope.
	IpamArn *string `pulumi:"ipamArn"`
	// The ID of the IPAM for which you're creating this scope.
	IpamId        *string `pulumi:"ipamId"`
	IpamScopeType *string `pulumi:"ipamScopeType"`
	// Defines if the scope is the default scope or not.
	IsDefault *bool `pulumi:"isDefault"`
	// The number of pools in the scope.
	PoolCount *int `pulumi:"poolCount"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VpcIpamScopeState struct {
	// The Amazon Resource Name (ARN) of the scope.
	Arn pulumi.StringPtrInput
	// A description for the scope you're creating.
	Description pulumi.StringPtrInput
	// The ARN of the IPAM for which you're creating this scope.
	IpamArn pulumi.StringPtrInput
	// The ID of the IPAM for which you're creating this scope.
	IpamId        pulumi.StringPtrInput
	IpamScopeType pulumi.StringPtrInput
	// Defines if the scope is the default scope or not.
	IsDefault pulumi.BoolPtrInput
	// The number of pools in the scope.
	PoolCount pulumi.IntPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (VpcIpamScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamScopeState)(nil)).Elem()
}

type vpcIpamScopeArgs struct {
	// A description for the scope you're creating.
	Description *string `pulumi:"description"`
	// The ID of the IPAM for which you're creating this scope.
	IpamId string `pulumi:"ipamId"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VpcIpamScope resource.
type VpcIpamScopeArgs struct {
	// A description for the scope you're creating.
	Description pulumi.StringPtrInput
	// The ID of the IPAM for which you're creating this scope.
	IpamId pulumi.StringInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (VpcIpamScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamScopeArgs)(nil)).Elem()
}

type VpcIpamScopeInput interface {
	pulumi.Input

	ToVpcIpamScopeOutput() VpcIpamScopeOutput
	ToVpcIpamScopeOutputWithContext(ctx context.Context) VpcIpamScopeOutput
}

func (*VpcIpamScope) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamScope)(nil)).Elem()
}

func (i *VpcIpamScope) ToVpcIpamScopeOutput() VpcIpamScopeOutput {
	return i.ToVpcIpamScopeOutputWithContext(context.Background())
}

func (i *VpcIpamScope) ToVpcIpamScopeOutputWithContext(ctx context.Context) VpcIpamScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamScopeOutput)
}

// VpcIpamScopeArrayInput is an input type that accepts VpcIpamScopeArray and VpcIpamScopeArrayOutput values.
// You can construct a concrete instance of `VpcIpamScopeArrayInput` via:
//
//	VpcIpamScopeArray{ VpcIpamScopeArgs{...} }
type VpcIpamScopeArrayInput interface {
	pulumi.Input

	ToVpcIpamScopeArrayOutput() VpcIpamScopeArrayOutput
	ToVpcIpamScopeArrayOutputWithContext(context.Context) VpcIpamScopeArrayOutput
}

type VpcIpamScopeArray []VpcIpamScopeInput

func (VpcIpamScopeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamScope)(nil)).Elem()
}

func (i VpcIpamScopeArray) ToVpcIpamScopeArrayOutput() VpcIpamScopeArrayOutput {
	return i.ToVpcIpamScopeArrayOutputWithContext(context.Background())
}

func (i VpcIpamScopeArray) ToVpcIpamScopeArrayOutputWithContext(ctx context.Context) VpcIpamScopeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamScopeArrayOutput)
}

// VpcIpamScopeMapInput is an input type that accepts VpcIpamScopeMap and VpcIpamScopeMapOutput values.
// You can construct a concrete instance of `VpcIpamScopeMapInput` via:
//
//	VpcIpamScopeMap{ "key": VpcIpamScopeArgs{...} }
type VpcIpamScopeMapInput interface {
	pulumi.Input

	ToVpcIpamScopeMapOutput() VpcIpamScopeMapOutput
	ToVpcIpamScopeMapOutputWithContext(context.Context) VpcIpamScopeMapOutput
}

type VpcIpamScopeMap map[string]VpcIpamScopeInput

func (VpcIpamScopeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamScope)(nil)).Elem()
}

func (i VpcIpamScopeMap) ToVpcIpamScopeMapOutput() VpcIpamScopeMapOutput {
	return i.ToVpcIpamScopeMapOutputWithContext(context.Background())
}

func (i VpcIpamScopeMap) ToVpcIpamScopeMapOutputWithContext(ctx context.Context) VpcIpamScopeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamScopeMapOutput)
}

type VpcIpamScopeOutput struct{ *pulumi.OutputState }

func (VpcIpamScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamScope)(nil)).Elem()
}

func (o VpcIpamScopeOutput) ToVpcIpamScopeOutput() VpcIpamScopeOutput {
	return o
}

func (o VpcIpamScopeOutput) ToVpcIpamScopeOutputWithContext(ctx context.Context) VpcIpamScopeOutput {
	return o
}

// The Amazon Resource Name (ARN) of the scope.
func (o VpcIpamScopeOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamScope) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A description for the scope you're creating.
func (o VpcIpamScopeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcIpamScope) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ARN of the IPAM for which you're creating this scope.
func (o VpcIpamScopeOutput) IpamArn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamScope) pulumi.StringOutput { return v.IpamArn }).(pulumi.StringOutput)
}

// The ID of the IPAM for which you're creating this scope.
func (o VpcIpamScopeOutput) IpamId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamScope) pulumi.StringOutput { return v.IpamId }).(pulumi.StringOutput)
}

func (o VpcIpamScopeOutput) IpamScopeType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamScope) pulumi.StringOutput { return v.IpamScopeType }).(pulumi.StringOutput)
}

// Defines if the scope is the default scope or not.
func (o VpcIpamScopeOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpcIpamScope) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// The number of pools in the scope.
func (o VpcIpamScopeOutput) PoolCount() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcIpamScope) pulumi.IntOutput { return v.PoolCount }).(pulumi.IntOutput)
}

// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VpcIpamScopeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpamScope) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Deprecated: Please use `tags` instead.
func (o VpcIpamScopeOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpamScope) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type VpcIpamScopeArrayOutput struct{ *pulumi.OutputState }

func (VpcIpamScopeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamScope)(nil)).Elem()
}

func (o VpcIpamScopeArrayOutput) ToVpcIpamScopeArrayOutput() VpcIpamScopeArrayOutput {
	return o
}

func (o VpcIpamScopeArrayOutput) ToVpcIpamScopeArrayOutputWithContext(ctx context.Context) VpcIpamScopeArrayOutput {
	return o
}

func (o VpcIpamScopeArrayOutput) Index(i pulumi.IntInput) VpcIpamScopeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcIpamScope {
		return vs[0].([]*VpcIpamScope)[vs[1].(int)]
	}).(VpcIpamScopeOutput)
}

type VpcIpamScopeMapOutput struct{ *pulumi.OutputState }

func (VpcIpamScopeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamScope)(nil)).Elem()
}

func (o VpcIpamScopeMapOutput) ToVpcIpamScopeMapOutput() VpcIpamScopeMapOutput {
	return o
}

func (o VpcIpamScopeMapOutput) ToVpcIpamScopeMapOutputWithContext(ctx context.Context) VpcIpamScopeMapOutput {
	return o
}

func (o VpcIpamScopeMapOutput) MapIndex(k pulumi.StringInput) VpcIpamScopeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcIpamScope {
		return vs[0].(map[string]*VpcIpamScope)[vs[1].(string)]
	}).(VpcIpamScopeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamScopeInput)(nil)).Elem(), &VpcIpamScope{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamScopeArrayInput)(nil)).Elem(), VpcIpamScopeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamScopeMapInput)(nil)).Elem(), VpcIpamScopeMap{})
	pulumi.RegisterOutputType(VpcIpamScopeOutput{})
	pulumi.RegisterOutputType(VpcIpamScopeArrayOutput{})
	pulumi.RegisterOutputType(VpcIpamScopeMapOutput{})
}
