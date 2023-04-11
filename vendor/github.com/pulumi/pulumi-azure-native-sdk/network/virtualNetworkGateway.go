// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A common class for general resource information.
// API Version: 2020-11-01.
type VirtualNetworkGateway struct {
	pulumi.CustomResourceState

	// ActiveActive flag.
	ActiveActive pulumi.BoolPtrOutput `pulumi:"activeActive"`
	// Virtual network gateway's BGP speaker settings.
	BgpSettings BgpSettingsResponsePtrOutput `pulumi:"bgpSettings"`
	// The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
	CustomRoutes AddressSpaceResponsePtrOutput `pulumi:"customRoutes"`
	// Whether BGP is enabled for this virtual network gateway or not.
	EnableBgp pulumi.BoolPtrOutput `pulumi:"enableBgp"`
	// Whether dns forwarding is enabled or not.
	EnableDnsForwarding pulumi.BoolPtrOutput `pulumi:"enableDnsForwarding"`
	// Whether private IP needs to be enabled on this gateway for connections or not.
	EnablePrivateIpAddress pulumi.BoolPtrOutput `pulumi:"enablePrivateIpAddress"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The extended location of type local virtual network gateway.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The reference to the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
	GatewayDefaultSite SubResourceResponsePtrOutput `pulumi:"gatewayDefaultSite"`
	// The type of this virtual network gateway.
	GatewayType pulumi.StringPtrOutput `pulumi:"gatewayType"`
	// The IP address allocated by the gateway to which dns requests can be sent.
	InboundDnsForwardingEndpoint pulumi.StringOutput `pulumi:"inboundDnsForwardingEndpoint"`
	// IP configurations for virtual network gateway.
	IpConfigurations VirtualNetworkGatewayIPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the virtual network gateway resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the virtual network gateway resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
	Sku VirtualNetworkGatewaySkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Customer vnet resource id. VirtualNetworkGateway of type local gateway is associated with the customer vnet.
	VNetExtendedLocationResourceId pulumi.StringPtrOutput `pulumi:"vNetExtendedLocationResourceId"`
	// The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
	VpnClientConfiguration VpnClientConfigurationResponsePtrOutput `pulumi:"vpnClientConfiguration"`
	// The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
	VpnGatewayGeneration pulumi.StringPtrOutput `pulumi:"vpnGatewayGeneration"`
	// The type of this virtual network gateway.
	VpnType pulumi.StringPtrOutput `pulumi:"vpnType"`
}

// NewVirtualNetworkGateway registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetworkGateway(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20150615:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualNetworkGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VirtualNetworkGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkGateway
	err := ctx.RegisterResource("azure-native:network:VirtualNetworkGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetworkGateway gets an existing VirtualNetworkGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetworkGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayState, opts ...pulumi.ResourceOption) (*VirtualNetworkGateway, error) {
	var resource VirtualNetworkGateway
	err := ctx.ReadResource("azure-native:network:VirtualNetworkGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetworkGateway resources.
type virtualNetworkGatewayState struct {
}

type VirtualNetworkGatewayState struct {
}

func (VirtualNetworkGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayState)(nil)).Elem()
}

type virtualNetworkGatewayArgs struct {
	// ActiveActive flag.
	ActiveActive *bool `pulumi:"activeActive"`
	// Virtual network gateway's BGP speaker settings.
	BgpSettings *BgpSettings `pulumi:"bgpSettings"`
	// The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
	CustomRoutes *AddressSpace `pulumi:"customRoutes"`
	// Whether BGP is enabled for this virtual network gateway or not.
	EnableBgp *bool `pulumi:"enableBgp"`
	// Whether dns forwarding is enabled or not.
	EnableDnsForwarding *bool `pulumi:"enableDnsForwarding"`
	// Whether private IP needs to be enabled on this gateway for connections or not.
	EnablePrivateIpAddress *bool `pulumi:"enablePrivateIpAddress"`
	// The extended location of type local virtual network gateway.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The reference to the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
	GatewayDefaultSite *SubResource `pulumi:"gatewayDefaultSite"`
	// The type of this virtual network gateway.
	GatewayType *string `pulumi:"gatewayType"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IP configurations for virtual network gateway.
	IpConfigurations []VirtualNetworkGatewayIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
	Sku *VirtualNetworkGatewaySku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Customer vnet resource id. VirtualNetworkGateway of type local gateway is associated with the customer vnet.
	VNetExtendedLocationResourceId *string `pulumi:"vNetExtendedLocationResourceId"`
	// The name of the virtual network gateway.
	VirtualNetworkGatewayName *string `pulumi:"virtualNetworkGatewayName"`
	// The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
	VpnClientConfiguration *VpnClientConfiguration `pulumi:"vpnClientConfiguration"`
	// The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
	VpnGatewayGeneration *string `pulumi:"vpnGatewayGeneration"`
	// The type of this virtual network gateway.
	VpnType *string `pulumi:"vpnType"`
}

// The set of arguments for constructing a VirtualNetworkGateway resource.
type VirtualNetworkGatewayArgs struct {
	// ActiveActive flag.
	ActiveActive pulumi.BoolPtrInput
	// Virtual network gateway's BGP speaker settings.
	BgpSettings BgpSettingsPtrInput
	// The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
	CustomRoutes AddressSpacePtrInput
	// Whether BGP is enabled for this virtual network gateway or not.
	EnableBgp pulumi.BoolPtrInput
	// Whether dns forwarding is enabled or not.
	EnableDnsForwarding pulumi.BoolPtrInput
	// Whether private IP needs to be enabled on this gateway for connections or not.
	EnablePrivateIpAddress pulumi.BoolPtrInput
	// The extended location of type local virtual network gateway.
	ExtendedLocation ExtendedLocationPtrInput
	// The reference to the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
	GatewayDefaultSite SubResourcePtrInput
	// The type of this virtual network gateway.
	GatewayType pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// IP configurations for virtual network gateway.
	IpConfigurations VirtualNetworkGatewayIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
	Sku VirtualNetworkGatewaySkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Customer vnet resource id. VirtualNetworkGateway of type local gateway is associated with the customer vnet.
	VNetExtendedLocationResourceId pulumi.StringPtrInput
	// The name of the virtual network gateway.
	VirtualNetworkGatewayName pulumi.StringPtrInput
	// The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
	VpnClientConfiguration VpnClientConfigurationPtrInput
	// The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
	VpnGatewayGeneration pulumi.StringPtrInput
	// The type of this virtual network gateway.
	VpnType pulumi.StringPtrInput
}

func (VirtualNetworkGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayArgs)(nil)).Elem()
}

type VirtualNetworkGatewayInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayOutput() VirtualNetworkGatewayOutput
	ToVirtualNetworkGatewayOutputWithContext(ctx context.Context) VirtualNetworkGatewayOutput
}

func (*VirtualNetworkGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGateway)(nil)).Elem()
}

func (i *VirtualNetworkGateway) ToVirtualNetworkGatewayOutput() VirtualNetworkGatewayOutput {
	return i.ToVirtualNetworkGatewayOutputWithContext(context.Background())
}

func (i *VirtualNetworkGateway) ToVirtualNetworkGatewayOutputWithContext(ctx context.Context) VirtualNetworkGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayOutput)
}

type VirtualNetworkGatewayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGateway)(nil)).Elem()
}

func (o VirtualNetworkGatewayOutput) ToVirtualNetworkGatewayOutput() VirtualNetworkGatewayOutput {
	return o
}

func (o VirtualNetworkGatewayOutput) ToVirtualNetworkGatewayOutputWithContext(ctx context.Context) VirtualNetworkGatewayOutput {
	return o
}

// ActiveActive flag.
func (o VirtualNetworkGatewayOutput) ActiveActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.BoolPtrOutput { return v.ActiveActive }).(pulumi.BoolPtrOutput)
}

// Virtual network gateway's BGP speaker settings.
func (o VirtualNetworkGatewayOutput) BgpSettings() BgpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) BgpSettingsResponsePtrOutput { return v.BgpSettings }).(BgpSettingsResponsePtrOutput)
}

// The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
func (o VirtualNetworkGatewayOutput) CustomRoutes() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) AddressSpaceResponsePtrOutput { return v.CustomRoutes }).(AddressSpaceResponsePtrOutput)
}

// Whether BGP is enabled for this virtual network gateway or not.
func (o VirtualNetworkGatewayOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.BoolPtrOutput { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

// Whether dns forwarding is enabled or not.
func (o VirtualNetworkGatewayOutput) EnableDnsForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.BoolPtrOutput { return v.EnableDnsForwarding }).(pulumi.BoolPtrOutput)
}

// Whether private IP needs to be enabled on this gateway for connections or not.
func (o VirtualNetworkGatewayOutput) EnablePrivateIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.BoolPtrOutput { return v.EnablePrivateIpAddress }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o VirtualNetworkGatewayOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of type local virtual network gateway.
func (o VirtualNetworkGatewayOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The reference to the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
func (o VirtualNetworkGatewayOutput) GatewayDefaultSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) SubResourceResponsePtrOutput { return v.GatewayDefaultSite }).(SubResourceResponsePtrOutput)
}

// The type of this virtual network gateway.
func (o VirtualNetworkGatewayOutput) GatewayType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringPtrOutput { return v.GatewayType }).(pulumi.StringPtrOutput)
}

// The IP address allocated by the gateway to which dns requests can be sent.
func (o VirtualNetworkGatewayOutput) InboundDnsForwardingEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringOutput { return v.InboundDnsForwardingEndpoint }).(pulumi.StringOutput)
}

// IP configurations for virtual network gateway.
func (o VirtualNetworkGatewayOutput) IpConfigurations() VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) VirtualNetworkGatewayIPConfigurationResponseArrayOutput {
		return v.IpConfigurations
	}).(VirtualNetworkGatewayIPConfigurationResponseArrayOutput)
}

// Resource location.
func (o VirtualNetworkGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o VirtualNetworkGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the virtual network gateway resource.
func (o VirtualNetworkGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the virtual network gateway resource.
func (o VirtualNetworkGatewayOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
func (o VirtualNetworkGatewayOutput) Sku() VirtualNetworkGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) VirtualNetworkGatewaySkuResponsePtrOutput { return v.Sku }).(VirtualNetworkGatewaySkuResponsePtrOutput)
}

// Resource tags.
func (o VirtualNetworkGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o VirtualNetworkGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Customer vnet resource id. VirtualNetworkGateway of type local gateway is associated with the customer vnet.
func (o VirtualNetworkGatewayOutput) VNetExtendedLocationResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringPtrOutput { return v.VNetExtendedLocationResourceId }).(pulumi.StringPtrOutput)
}

// The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
func (o VirtualNetworkGatewayOutput) VpnClientConfiguration() VpnClientConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) VpnClientConfigurationResponsePtrOutput {
		return v.VpnClientConfiguration
	}).(VpnClientConfigurationResponsePtrOutput)
}

// The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
func (o VirtualNetworkGatewayOutput) VpnGatewayGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringPtrOutput { return v.VpnGatewayGeneration }).(pulumi.StringPtrOutput)
}

// The type of this virtual network gateway.
func (o VirtualNetworkGatewayOutput) VpnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGateway) pulumi.StringPtrOutput { return v.VpnType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkGatewayOutput{})
}
