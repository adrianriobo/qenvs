// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// IpAllocation resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01.
type IpAllocation struct {
	pulumi.CustomResourceState

	// IpAllocation tags.
	AllocationTags pulumi.StringMapOutput `pulumi:"allocationTags"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The IPAM allocation ID.
	IpamAllocationId pulumi.StringPtrOutput `pulumi:"ipamAllocationId"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The address prefix for the IpAllocation.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The address prefix length for the IpAllocation.
	PrefixLength pulumi.IntPtrOutput `pulumi:"prefixLength"`
	// The address prefix Type for the IpAllocation.
	PrefixType pulumi.StringPtrOutput `pulumi:"prefixType"`
	// The Subnet that using the prefix of this IpAllocation resource.
	Subnet SubResourceResponseOutput `pulumi:"subnet"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The VirtualNetwork that using the prefix of this IpAllocation resource.
	VirtualNetwork SubResourceResponseOutput `pulumi:"virtualNetwork"`
}

// NewIpAllocation registers a new resource with the given unique name, arguments, and options.
func NewIpAllocation(ctx *pulumi.Context,
	name string, args *IpAllocationArgs, opts ...pulumi.ResourceOption) (*IpAllocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.PrefixLength == nil {
		args.PrefixLength = pulumi.IntPtr(0)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20200301:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:IpAllocation"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:IpAllocation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IpAllocation
	err := ctx.RegisterResource("azure-native:network:IpAllocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpAllocation gets an existing IpAllocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpAllocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpAllocationState, opts ...pulumi.ResourceOption) (*IpAllocation, error) {
	var resource IpAllocation
	err := ctx.ReadResource("azure-native:network:IpAllocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpAllocation resources.
type ipAllocationState struct {
}

type IpAllocationState struct {
}

func (IpAllocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAllocationState)(nil)).Elem()
}

type ipAllocationArgs struct {
	// IpAllocation tags.
	AllocationTags map[string]string `pulumi:"allocationTags"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the IpAllocation.
	IpAllocationName *string `pulumi:"ipAllocationName"`
	// The IPAM allocation ID.
	IpamAllocationId *string `pulumi:"ipamAllocationId"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The address prefix for the IpAllocation.
	Prefix *string `pulumi:"prefix"`
	// The address prefix length for the IpAllocation.
	PrefixLength *int `pulumi:"prefixLength"`
	// The address prefix Type for the IpAllocation.
	PrefixType *string `pulumi:"prefixType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type for the IpAllocation.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a IpAllocation resource.
type IpAllocationArgs struct {
	// IpAllocation tags.
	AllocationTags pulumi.StringMapInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the IpAllocation.
	IpAllocationName pulumi.StringPtrInput
	// The IPAM allocation ID.
	IpamAllocationId pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The address prefix for the IpAllocation.
	Prefix pulumi.StringPtrInput
	// The address prefix length for the IpAllocation.
	PrefixLength pulumi.IntPtrInput
	// The address prefix Type for the IpAllocation.
	PrefixType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type for the IpAllocation.
	Type pulumi.StringPtrInput
}

func (IpAllocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAllocationArgs)(nil)).Elem()
}

type IpAllocationInput interface {
	pulumi.Input

	ToIpAllocationOutput() IpAllocationOutput
	ToIpAllocationOutputWithContext(ctx context.Context) IpAllocationOutput
}

func (*IpAllocation) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAllocation)(nil)).Elem()
}

func (i *IpAllocation) ToIpAllocationOutput() IpAllocationOutput {
	return i.ToIpAllocationOutputWithContext(context.Background())
}

func (i *IpAllocation) ToIpAllocationOutputWithContext(ctx context.Context) IpAllocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAllocationOutput)
}

func (i *IpAllocation) ToOutput(ctx context.Context) pulumix.Output[*IpAllocation] {
	return pulumix.Output[*IpAllocation]{
		OutputState: i.ToIpAllocationOutputWithContext(ctx).OutputState,
	}
}

type IpAllocationOutput struct{ *pulumi.OutputState }

func (IpAllocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAllocation)(nil)).Elem()
}

func (o IpAllocationOutput) ToIpAllocationOutput() IpAllocationOutput {
	return o
}

func (o IpAllocationOutput) ToIpAllocationOutputWithContext(ctx context.Context) IpAllocationOutput {
	return o
}

func (o IpAllocationOutput) ToOutput(ctx context.Context) pulumix.Output[*IpAllocation] {
	return pulumix.Output[*IpAllocation]{
		OutputState: o.OutputState,
	}
}

// IpAllocation tags.
func (o IpAllocationOutput) AllocationTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringMapOutput { return v.AllocationTags }).(pulumi.StringMapOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o IpAllocationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The IPAM allocation ID.
func (o IpAllocationOutput) IpamAllocationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringPtrOutput { return v.IpamAllocationId }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o IpAllocationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o IpAllocationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The address prefix for the IpAllocation.
func (o IpAllocationOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringPtrOutput { return v.Prefix }).(pulumi.StringPtrOutput)
}

// The address prefix length for the IpAllocation.
func (o IpAllocationOutput) PrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.IntPtrOutput { return v.PrefixLength }).(pulumi.IntPtrOutput)
}

// The address prefix Type for the IpAllocation.
func (o IpAllocationOutput) PrefixType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringPtrOutput { return v.PrefixType }).(pulumi.StringPtrOutput)
}

// The Subnet that using the prefix of this IpAllocation resource.
func (o IpAllocationOutput) Subnet() SubResourceResponseOutput {
	return o.ApplyT(func(v *IpAllocation) SubResourceResponseOutput { return v.Subnet }).(SubResourceResponseOutput)
}

// Resource tags.
func (o IpAllocationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o IpAllocationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAllocation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The VirtualNetwork that using the prefix of this IpAllocation resource.
func (o IpAllocationOutput) VirtualNetwork() SubResourceResponseOutput {
	return o.ApplyT(func(v *IpAllocation) SubResourceResponseOutput { return v.VirtualNetwork }).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(IpAllocationOutput{})
}
