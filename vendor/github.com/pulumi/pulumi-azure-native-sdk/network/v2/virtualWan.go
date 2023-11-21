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

// VirtualWAN Resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2019-07-01, 2023-04-01, 2023-05-01, 2023-06-01.
type VirtualWan struct {
	pulumi.CustomResourceState

	// True if branch to branch traffic is allowed.
	AllowBranchToBranchTraffic pulumi.BoolPtrOutput `pulumi:"allowBranchToBranchTraffic"`
	// True if Vnet to Vnet traffic is allowed.
	AllowVnetToVnetTraffic pulumi.BoolPtrOutput `pulumi:"allowVnetToVnetTraffic"`
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption pulumi.BoolPtrOutput `pulumi:"disableVpnEncryption"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The office local breakout category.
	Office365LocalBreakoutCategory pulumi.StringOutput `pulumi:"office365LocalBreakoutCategory"`
	// The provisioning state of the virtual WAN resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// List of VirtualHubs in the VirtualWAN.
	VirtualHubs SubResourceResponseArrayOutput `pulumi:"virtualHubs"`
	// List of VpnSites in the VirtualWAN.
	VpnSites SubResourceResponseArrayOutput `pulumi:"vpnSites"`
}

// NewVirtualWan registers a new resource with the given unique name, arguments, and options.
func NewVirtualWan(ctx *pulumi.Context,
	name string, args *VirtualWanArgs, opts ...pulumi.ResourceOption) (*VirtualWan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:VirtualWan"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:VirtualWan"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualWan
	err := ctx.RegisterResource("azure-native:network:VirtualWan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualWan gets an existing VirtualWan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualWan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualWanState, opts ...pulumi.ResourceOption) (*VirtualWan, error) {
	var resource VirtualWan
	err := ctx.ReadResource("azure-native:network:VirtualWan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualWan resources.
type virtualWanState struct {
}

type VirtualWanState struct {
}

func (VirtualWanState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWanState)(nil)).Elem()
}

type virtualWanArgs struct {
	// True if branch to branch traffic is allowed.
	AllowBranchToBranchTraffic *bool `pulumi:"allowBranchToBranchTraffic"`
	// True if Vnet to Vnet traffic is allowed.
	AllowVnetToVnetTraffic *bool `pulumi:"allowVnetToVnetTraffic"`
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption *bool `pulumi:"disableVpnEncryption"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The resource group name of the VirtualWan.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the VirtualWAN.
	Type *string `pulumi:"type"`
	// The name of the VirtualWAN being created or updated.
	VirtualWANName *string `pulumi:"virtualWANName"`
}

// The set of arguments for constructing a VirtualWan resource.
type VirtualWanArgs struct {
	// True if branch to branch traffic is allowed.
	AllowBranchToBranchTraffic pulumi.BoolPtrInput
	// True if Vnet to Vnet traffic is allowed.
	AllowVnetToVnetTraffic pulumi.BoolPtrInput
	// Vpn encryption to be disabled or not.
	DisableVpnEncryption pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The resource group name of the VirtualWan.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the VirtualWAN.
	Type pulumi.StringPtrInput
	// The name of the VirtualWAN being created or updated.
	VirtualWANName pulumi.StringPtrInput
}

func (VirtualWanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualWanArgs)(nil)).Elem()
}

type VirtualWanInput interface {
	pulumi.Input

	ToVirtualWanOutput() VirtualWanOutput
	ToVirtualWanOutputWithContext(ctx context.Context) VirtualWanOutput
}

func (*VirtualWan) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualWan)(nil)).Elem()
}

func (i *VirtualWan) ToVirtualWanOutput() VirtualWanOutput {
	return i.ToVirtualWanOutputWithContext(context.Background())
}

func (i *VirtualWan) ToVirtualWanOutputWithContext(ctx context.Context) VirtualWanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualWanOutput)
}

func (i *VirtualWan) ToOutput(ctx context.Context) pulumix.Output[*VirtualWan] {
	return pulumix.Output[*VirtualWan]{
		OutputState: i.ToVirtualWanOutputWithContext(ctx).OutputState,
	}
}

type VirtualWanOutput struct{ *pulumi.OutputState }

func (VirtualWanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualWan)(nil)).Elem()
}

func (o VirtualWanOutput) ToVirtualWanOutput() VirtualWanOutput {
	return o
}

func (o VirtualWanOutput) ToVirtualWanOutputWithContext(ctx context.Context) VirtualWanOutput {
	return o
}

func (o VirtualWanOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualWan] {
	return pulumix.Output[*VirtualWan]{
		OutputState: o.OutputState,
	}
}

// True if branch to branch traffic is allowed.
func (o VirtualWanOutput) AllowBranchToBranchTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.BoolPtrOutput { return v.AllowBranchToBranchTraffic }).(pulumi.BoolPtrOutput)
}

// True if Vnet to Vnet traffic is allowed.
func (o VirtualWanOutput) AllowVnetToVnetTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.BoolPtrOutput { return v.AllowVnetToVnetTraffic }).(pulumi.BoolPtrOutput)
}

// Vpn encryption to be disabled or not.
func (o VirtualWanOutput) DisableVpnEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.BoolPtrOutput { return v.DisableVpnEncryption }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o VirtualWanOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource location.
func (o VirtualWanOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o VirtualWanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The office local breakout category.
func (o VirtualWanOutput) Office365LocalBreakoutCategory() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.Office365LocalBreakoutCategory }).(pulumi.StringOutput)
}

// The provisioning state of the virtual WAN resource.
func (o VirtualWanOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o VirtualWanOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o VirtualWanOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualWan) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// List of VirtualHubs in the VirtualWAN.
func (o VirtualWanOutput) VirtualHubs() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualWan) SubResourceResponseArrayOutput { return v.VirtualHubs }).(SubResourceResponseArrayOutput)
}

// List of VpnSites in the VirtualWAN.
func (o VirtualWanOutput) VpnSites() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualWan) SubResourceResponseArrayOutput { return v.VpnSites }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualWanOutput{})
}
