// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Availability sets overview](https://docs.microsoft.com/azure/virtual-machines/availability-set-overview). For more information on Azure planned maintenance, see [Maintenance and updates for Virtual Machines in Azure](https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates). Currently, a VM can only be added to an availability set at creation time. An existing VM cannot be added to an availability set.
// Azure REST API version: 2023-03-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2016-04-30-preview, 2023-07-01.
type AvailabilitySet struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Fault Domain count.
	PlatformFaultDomainCount pulumi.IntPtrOutput `pulumi:"platformFaultDomainCount"`
	// Update Domain count.
	PlatformUpdateDomainCount pulumi.IntPtrOutput `pulumi:"platformUpdateDomainCount"`
	// Specifies information about the proximity placement group that the availability set should be assigned to. Minimum api-version: 2018-04-01.
	ProximityPlacementGroup SubResourceResponsePtrOutput `pulumi:"proximityPlacementGroup"`
	// Sku of the availability set, only name is required to be set. See AvailabilitySetSkuTypes for possible set of values. Use 'Aligned' for virtual machines with managed disks and 'Classic' for virtual machines with unmanaged disks. Default value is 'Classic'.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The resource status information.
	Statuses InstanceViewStatusResponseArrayOutput `pulumi:"statuses"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of references to all virtual machines in the availability set.
	VirtualMachines SubResourceResponseArrayOutput `pulumi:"virtualMachines"`
}

// NewAvailabilitySet registers a new resource with the given unique name, arguments, and options.
func NewAvailabilitySet(ctx *pulumi.Context,
	name string, args *AvailabilitySetArgs, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20150615:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160330:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230701:AvailabilitySet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AvailabilitySet
	err := ctx.RegisterResource("azure-native:compute:AvailabilitySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAvailabilitySet gets an existing AvailabilitySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAvailabilitySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilitySetState, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	var resource AvailabilitySet
	err := ctx.ReadResource("azure-native:compute:AvailabilitySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AvailabilitySet resources.
type availabilitySetState struct {
}

type AvailabilitySetState struct {
}

func (AvailabilitySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetState)(nil)).Elem()
}

type availabilitySetArgs struct {
	// The name of the availability set.
	AvailabilitySetName *string `pulumi:"availabilitySetName"`
	// Resource location
	Location *string `pulumi:"location"`
	// Fault Domain count.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Update Domain count.
	PlatformUpdateDomainCount *int `pulumi:"platformUpdateDomainCount"`
	// Specifies information about the proximity placement group that the availability set should be assigned to. Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResource `pulumi:"proximityPlacementGroup"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sku of the availability set, only name is required to be set. See AvailabilitySetSkuTypes for possible set of values. Use 'Aligned' for virtual machines with managed disks and 'Classic' for virtual machines with unmanaged disks. Default value is 'Classic'.
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// A list of references to all virtual machines in the availability set.
	VirtualMachines []SubResource `pulumi:"virtualMachines"`
}

// The set of arguments for constructing a AvailabilitySet resource.
type AvailabilitySetArgs struct {
	// The name of the availability set.
	AvailabilitySetName pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Fault Domain count.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Update Domain count.
	PlatformUpdateDomainCount pulumi.IntPtrInput
	// Specifies information about the proximity placement group that the availability set should be assigned to. Minimum api-version: 2018-04-01.
	ProximityPlacementGroup SubResourcePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Sku of the availability set, only name is required to be set. See AvailabilitySetSkuTypes for possible set of values. Use 'Aligned' for virtual machines with managed disks and 'Classic' for virtual machines with unmanaged disks. Default value is 'Classic'.
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// A list of references to all virtual machines in the availability set.
	VirtualMachines SubResourceArrayInput
}

func (AvailabilitySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetArgs)(nil)).Elem()
}

type AvailabilitySetInput interface {
	pulumi.Input

	ToAvailabilitySetOutput() AvailabilitySetOutput
	ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput
}

func (*AvailabilitySet) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilitySet)(nil)).Elem()
}

func (i *AvailabilitySet) ToAvailabilitySetOutput() AvailabilitySetOutput {
	return i.ToAvailabilitySetOutputWithContext(context.Background())
}

func (i *AvailabilitySet) ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetOutput)
}

func (i *AvailabilitySet) ToOutput(ctx context.Context) pulumix.Output[*AvailabilitySet] {
	return pulumix.Output[*AvailabilitySet]{
		OutputState: i.ToAvailabilitySetOutputWithContext(ctx).OutputState,
	}
}

type AvailabilitySetOutput struct{ *pulumi.OutputState }

func (AvailabilitySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilitySet)(nil)).Elem()
}

func (o AvailabilitySetOutput) ToAvailabilitySetOutput() AvailabilitySetOutput {
	return o
}

func (o AvailabilitySetOutput) ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput {
	return o
}

func (o AvailabilitySetOutput) ToOutput(ctx context.Context) pulumix.Output[*AvailabilitySet] {
	return pulumix.Output[*AvailabilitySet]{
		OutputState: o.OutputState,
	}
}

// Resource location
func (o AvailabilitySetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o AvailabilitySetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Fault Domain count.
func (o AvailabilitySetOutput) PlatformFaultDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.IntPtrOutput { return v.PlatformFaultDomainCount }).(pulumi.IntPtrOutput)
}

// Update Domain count.
func (o AvailabilitySetOutput) PlatformUpdateDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.IntPtrOutput { return v.PlatformUpdateDomainCount }).(pulumi.IntPtrOutput)
}

// Specifies information about the proximity placement group that the availability set should be assigned to. Minimum api-version: 2018-04-01.
func (o AvailabilitySetOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) SubResourceResponsePtrOutput { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

// Sku of the availability set, only name is required to be set. See AvailabilitySetSkuTypes for possible set of values. Use 'Aligned' for virtual machines with managed disks and 'Classic' for virtual machines with unmanaged disks. Default value is 'Classic'.
func (o AvailabilitySetOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The resource status information.
func (o AvailabilitySetOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *AvailabilitySet) InstanceViewStatusResponseArrayOutput { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

// Resource tags
func (o AvailabilitySetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o AvailabilitySetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of references to all virtual machines in the availability set.
func (o AvailabilitySetOutput) VirtualMachines() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *AvailabilitySet) SubResourceResponseArrayOutput { return v.VirtualMachines }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AvailabilitySetOutput{})
}
