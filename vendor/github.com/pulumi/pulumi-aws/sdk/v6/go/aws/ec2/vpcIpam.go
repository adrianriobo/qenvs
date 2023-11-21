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

// Provides an IPAM resource.
//
// ## Import
//
// Using `pulumi import`, import IPAMs using the IPAM `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2/vpcIpam:VpcIpam example ipam-0178368ad2146a492
//
// ```
type VpcIpam struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of IPAM
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.
	Cascade pulumi.BoolPtrOutput `pulumi:"cascade"`
	// The IPAM's default resource discovery association ID.
	DefaultResourceDiscoveryAssociationId pulumi.StringOutput `pulumi:"defaultResourceDiscoveryAssociationId"`
	// The IPAM's default resource discovery ID.
	DefaultResourceDiscoveryId pulumi.StringOutput `pulumi:"defaultResourceDiscoveryId"`
	// A description for the IPAM.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. You **must** set your provider block region as an operating_region.
	OperatingRegions VpcIpamOperatingRegionArrayOutput `pulumi:"operatingRegions"`
	// The ID of the IPAM's private scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private IP space. The public scope is intended for all internet-routable IP space.
	PrivateDefaultScopeId pulumi.StringOutput `pulumi:"privateDefaultScopeId"`
	// The ID of the IPAM's public scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private
	// IP space. The public scope is intended for all internet-routable IP space.
	PublicDefaultScopeId pulumi.StringOutput `pulumi:"publicDefaultScopeId"`
	// The number of scopes in the IPAM.
	ScopeCount pulumi.IntOutput `pulumi:"scopeCount"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVpcIpam registers a new resource with the given unique name, arguments, and options.
func NewVpcIpam(ctx *pulumi.Context,
	name string, args *VpcIpamArgs, opts ...pulumi.ResourceOption) (*VpcIpam, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OperatingRegions == nil {
		return nil, errors.New("invalid value for required argument 'OperatingRegions'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcIpam
	err := ctx.RegisterResource("aws:ec2/vpcIpam:VpcIpam", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIpam gets an existing VpcIpam resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIpam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIpamState, opts ...pulumi.ResourceOption) (*VpcIpam, error) {
	var resource VpcIpam
	err := ctx.ReadResource("aws:ec2/vpcIpam:VpcIpam", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIpam resources.
type vpcIpamState struct {
	// Amazon Resource Name (ARN) of IPAM
	Arn *string `pulumi:"arn"`
	// Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.
	Cascade *bool `pulumi:"cascade"`
	// The IPAM's default resource discovery association ID.
	DefaultResourceDiscoveryAssociationId *string `pulumi:"defaultResourceDiscoveryAssociationId"`
	// The IPAM's default resource discovery ID.
	DefaultResourceDiscoveryId *string `pulumi:"defaultResourceDiscoveryId"`
	// A description for the IPAM.
	Description *string `pulumi:"description"`
	// Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. You **must** set your provider block region as an operating_region.
	OperatingRegions []VpcIpamOperatingRegion `pulumi:"operatingRegions"`
	// The ID of the IPAM's private scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private IP space. The public scope is intended for all internet-routable IP space.
	PrivateDefaultScopeId *string `pulumi:"privateDefaultScopeId"`
	// The ID of the IPAM's public scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private
	// IP space. The public scope is intended for all internet-routable IP space.
	PublicDefaultScopeId *string `pulumi:"publicDefaultScopeId"`
	// The number of scopes in the IPAM.
	ScopeCount *int `pulumi:"scopeCount"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VpcIpamState struct {
	// Amazon Resource Name (ARN) of IPAM
	Arn pulumi.StringPtrInput
	// Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.
	Cascade pulumi.BoolPtrInput
	// The IPAM's default resource discovery association ID.
	DefaultResourceDiscoveryAssociationId pulumi.StringPtrInput
	// The IPAM's default resource discovery ID.
	DefaultResourceDiscoveryId pulumi.StringPtrInput
	// A description for the IPAM.
	Description pulumi.StringPtrInput
	// Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. You **must** set your provider block region as an operating_region.
	OperatingRegions VpcIpamOperatingRegionArrayInput
	// The ID of the IPAM's private scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private IP space. The public scope is intended for all internet-routable IP space.
	PrivateDefaultScopeId pulumi.StringPtrInput
	// The ID of the IPAM's public scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private
	// IP space. The public scope is intended for all internet-routable IP space.
	PublicDefaultScopeId pulumi.StringPtrInput
	// The number of scopes in the IPAM.
	ScopeCount pulumi.IntPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (VpcIpamState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamState)(nil)).Elem()
}

type vpcIpamArgs struct {
	// Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.
	Cascade *bool `pulumi:"cascade"`
	// A description for the IPAM.
	Description *string `pulumi:"description"`
	// Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. You **must** set your provider block region as an operating_region.
	OperatingRegions []VpcIpamOperatingRegion `pulumi:"operatingRegions"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VpcIpam resource.
type VpcIpamArgs struct {
	// Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.
	Cascade pulumi.BoolPtrInput
	// A description for the IPAM.
	Description pulumi.StringPtrInput
	// Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. You **must** set your provider block region as an operating_region.
	OperatingRegions VpcIpamOperatingRegionArrayInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (VpcIpamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamArgs)(nil)).Elem()
}

type VpcIpamInput interface {
	pulumi.Input

	ToVpcIpamOutput() VpcIpamOutput
	ToVpcIpamOutputWithContext(ctx context.Context) VpcIpamOutput
}

func (*VpcIpam) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpam)(nil)).Elem()
}

func (i *VpcIpam) ToVpcIpamOutput() VpcIpamOutput {
	return i.ToVpcIpamOutputWithContext(context.Background())
}

func (i *VpcIpam) ToVpcIpamOutputWithContext(ctx context.Context) VpcIpamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamOutput)
}

// VpcIpamArrayInput is an input type that accepts VpcIpamArray and VpcIpamArrayOutput values.
// You can construct a concrete instance of `VpcIpamArrayInput` via:
//
//	VpcIpamArray{ VpcIpamArgs{...} }
type VpcIpamArrayInput interface {
	pulumi.Input

	ToVpcIpamArrayOutput() VpcIpamArrayOutput
	ToVpcIpamArrayOutputWithContext(context.Context) VpcIpamArrayOutput
}

type VpcIpamArray []VpcIpamInput

func (VpcIpamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpam)(nil)).Elem()
}

func (i VpcIpamArray) ToVpcIpamArrayOutput() VpcIpamArrayOutput {
	return i.ToVpcIpamArrayOutputWithContext(context.Background())
}

func (i VpcIpamArray) ToVpcIpamArrayOutputWithContext(ctx context.Context) VpcIpamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamArrayOutput)
}

// VpcIpamMapInput is an input type that accepts VpcIpamMap and VpcIpamMapOutput values.
// You can construct a concrete instance of `VpcIpamMapInput` via:
//
//	VpcIpamMap{ "key": VpcIpamArgs{...} }
type VpcIpamMapInput interface {
	pulumi.Input

	ToVpcIpamMapOutput() VpcIpamMapOutput
	ToVpcIpamMapOutputWithContext(context.Context) VpcIpamMapOutput
}

type VpcIpamMap map[string]VpcIpamInput

func (VpcIpamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpam)(nil)).Elem()
}

func (i VpcIpamMap) ToVpcIpamMapOutput() VpcIpamMapOutput {
	return i.ToVpcIpamMapOutputWithContext(context.Background())
}

func (i VpcIpamMap) ToVpcIpamMapOutputWithContext(ctx context.Context) VpcIpamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamMapOutput)
}

type VpcIpamOutput struct{ *pulumi.OutputState }

func (VpcIpamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpam)(nil)).Elem()
}

func (o VpcIpamOutput) ToVpcIpamOutput() VpcIpamOutput {
	return o
}

func (o VpcIpamOutput) ToVpcIpamOutputWithContext(ctx context.Context) VpcIpamOutput {
	return o
}

// Amazon Resource Name (ARN) of IPAM
func (o VpcIpamOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpam) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and any allocations in the pools in private scopes.
func (o VpcIpamOutput) Cascade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcIpam) pulumi.BoolPtrOutput { return v.Cascade }).(pulumi.BoolPtrOutput)
}

// The IPAM's default resource discovery association ID.
func (o VpcIpamOutput) DefaultResourceDiscoveryAssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpam) pulumi.StringOutput { return v.DefaultResourceDiscoveryAssociationId }).(pulumi.StringOutput)
}

// The IPAM's default resource discovery ID.
func (o VpcIpamOutput) DefaultResourceDiscoveryId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpam) pulumi.StringOutput { return v.DefaultResourceDiscoveryId }).(pulumi.StringOutput)
}

// A description for the IPAM.
func (o VpcIpamOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcIpam) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. You **must** set your provider block region as an operating_region.
func (o VpcIpamOutput) OperatingRegions() VpcIpamOperatingRegionArrayOutput {
	return o.ApplyT(func(v *VpcIpam) VpcIpamOperatingRegionArrayOutput { return v.OperatingRegions }).(VpcIpamOperatingRegionArrayOutput)
}

// The ID of the IPAM's private scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private IP space. The public scope is intended for all internet-routable IP space.
func (o VpcIpamOutput) PrivateDefaultScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpam) pulumi.StringOutput { return v.PrivateDefaultScopeId }).(pulumi.StringOutput)
}

// The ID of the IPAM's public scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private
// IP space. The public scope is intended for all internet-routable IP space.
func (o VpcIpamOutput) PublicDefaultScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpam) pulumi.StringOutput { return v.PublicDefaultScopeId }).(pulumi.StringOutput)
}

// The number of scopes in the IPAM.
func (o VpcIpamOutput) ScopeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcIpam) pulumi.IntOutput { return v.ScopeCount }).(pulumi.IntOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VpcIpamOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpam) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o VpcIpamOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpam) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type VpcIpamArrayOutput struct{ *pulumi.OutputState }

func (VpcIpamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpam)(nil)).Elem()
}

func (o VpcIpamArrayOutput) ToVpcIpamArrayOutput() VpcIpamArrayOutput {
	return o
}

func (o VpcIpamArrayOutput) ToVpcIpamArrayOutputWithContext(ctx context.Context) VpcIpamArrayOutput {
	return o
}

func (o VpcIpamArrayOutput) Index(i pulumi.IntInput) VpcIpamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcIpam {
		return vs[0].([]*VpcIpam)[vs[1].(int)]
	}).(VpcIpamOutput)
}

type VpcIpamMapOutput struct{ *pulumi.OutputState }

func (VpcIpamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpam)(nil)).Elem()
}

func (o VpcIpamMapOutput) ToVpcIpamMapOutput() VpcIpamMapOutput {
	return o
}

func (o VpcIpamMapOutput) ToVpcIpamMapOutputWithContext(ctx context.Context) VpcIpamMapOutput {
	return o
}

func (o VpcIpamMapOutput) MapIndex(k pulumi.StringInput) VpcIpamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcIpam {
		return vs[0].(map[string]*VpcIpam)[vs[1].(string)]
	}).(VpcIpamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamInput)(nil)).Elem(), &VpcIpam{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamArrayInput)(nil)).Elem(), VpcIpamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamMapInput)(nil)).Elem(), VpcIpamMap{})
	pulumi.RegisterOutputType(VpcIpamOutput{})
	pulumi.RegisterOutputType(VpcIpamArrayOutput{})
	pulumi.RegisterOutputType(VpcIpamMapOutput{})
}
