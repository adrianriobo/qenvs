// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Peering in an ExpressRoute Cross Connection resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01.
type ExpressRouteCrossConnectionPeering struct {
	pulumi.CustomResourceState

	// The Azure ASN.
	AzureASN pulumi.IntOutput `pulumi:"azureASN"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrOutput `pulumi:"gatewayManagerEtag"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput `pulumi:"ipv6PeeringConfig"`
	// Who was the last to modify the peering.
	LastModifiedBy pulumi.StringOutput `pulumi:"lastModifiedBy"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringConfigResponsePtrOutput `pulumi:"microsoftPeeringConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The peer ASN.
	PeerASN pulumi.Float64PtrOutput `pulumi:"peerASN"`
	// The peering type.
	PeeringType pulumi.StringPtrOutput `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort pulumi.StringOutput `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix pulumi.StringPtrOutput `pulumi:"primaryPeerAddressPrefix"`
	// The provisioning state of the express route cross connection peering resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The secondary port.
	SecondaryAzurePort pulumi.StringOutput `pulumi:"secondaryAzurePort"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix pulumi.StringPtrOutput `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey pulumi.StringPtrOutput `pulumi:"sharedKey"`
	// The peering state.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The VLAN ID.
	VlanId pulumi.IntPtrOutput `pulumi:"vlanId"`
}

// NewExpressRouteCrossConnectionPeering registers a new resource with the given unique name, arguments, and options.
func NewExpressRouteCrossConnectionPeering(ctx *pulumi.Context,
	name string, args *ExpressRouteCrossConnectionPeeringArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCrossConnectionPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CrossConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'CrossConnectionName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230901:ExpressRouteCrossConnectionPeering"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ExpressRouteCrossConnectionPeering
	err := ctx.RegisterResource("azure-native:network:ExpressRouteCrossConnectionPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRouteCrossConnectionPeering gets an existing ExpressRouteCrossConnectionPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRouteCrossConnectionPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCrossConnectionPeeringState, opts ...pulumi.ResourceOption) (*ExpressRouteCrossConnectionPeering, error) {
	var resource ExpressRouteCrossConnectionPeering
	err := ctx.ReadResource("azure-native:network:ExpressRouteCrossConnectionPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRouteCrossConnectionPeering resources.
type expressRouteCrossConnectionPeeringState struct {
}

type ExpressRouteCrossConnectionPeeringState struct {
}

func (ExpressRouteCrossConnectionPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCrossConnectionPeeringState)(nil)).Elem()
}

type expressRouteCrossConnectionPeeringArgs struct {
	// The name of the ExpressRouteCrossConnection.
	CrossConnectionName string `pulumi:"crossConnectionName"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig *Ipv6ExpressRouteCircuitPeeringConfig `pulumi:"ipv6PeeringConfig"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringConfig `pulumi:"microsoftPeeringConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The peer ASN.
	PeerASN *float64 `pulumi:"peerASN"`
	// The name of the peering.
	PeeringName *string `pulumi:"peeringName"`
	// The peering type.
	PeeringType *string `pulumi:"peeringType"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix *string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// The peering state.
	State *string `pulumi:"state"`
	// The VLAN ID.
	VlanId *int `pulumi:"vlanId"`
}

// The set of arguments for constructing a ExpressRouteCrossConnectionPeering resource.
type ExpressRouteCrossConnectionPeeringArgs struct {
	// The name of the ExpressRouteCrossConnection.
	CrossConnectionName pulumi.StringInput
	// The GatewayManager Etag.
	GatewayManagerEtag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The IPv6 peering configuration.
	Ipv6PeeringConfig Ipv6ExpressRouteCircuitPeeringConfigPtrInput
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig ExpressRouteCircuitPeeringConfigPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The peer ASN.
	PeerASN pulumi.Float64PtrInput
	// The name of the peering.
	PeeringName pulumi.StringPtrInput
	// The peering type.
	PeeringType pulumi.StringPtrInput
	// The primary address prefix.
	PrimaryPeerAddressPrefix pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The secondary address prefix.
	SecondaryPeerAddressPrefix pulumi.StringPtrInput
	// The shared key.
	SharedKey pulumi.StringPtrInput
	// The peering state.
	State pulumi.StringPtrInput
	// The VLAN ID.
	VlanId pulumi.IntPtrInput
}

func (ExpressRouteCrossConnectionPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCrossConnectionPeeringArgs)(nil)).Elem()
}

type ExpressRouteCrossConnectionPeeringInput interface {
	pulumi.Input

	ToExpressRouteCrossConnectionPeeringOutput() ExpressRouteCrossConnectionPeeringOutput
	ToExpressRouteCrossConnectionPeeringOutputWithContext(ctx context.Context) ExpressRouteCrossConnectionPeeringOutput
}

func (*ExpressRouteCrossConnectionPeering) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCrossConnectionPeering)(nil)).Elem()
}

func (i *ExpressRouteCrossConnectionPeering) ToExpressRouteCrossConnectionPeeringOutput() ExpressRouteCrossConnectionPeeringOutput {
	return i.ToExpressRouteCrossConnectionPeeringOutputWithContext(context.Background())
}

func (i *ExpressRouteCrossConnectionPeering) ToExpressRouteCrossConnectionPeeringOutputWithContext(ctx context.Context) ExpressRouteCrossConnectionPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCrossConnectionPeeringOutput)
}

type ExpressRouteCrossConnectionPeeringOutput struct{ *pulumi.OutputState }

func (ExpressRouteCrossConnectionPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCrossConnectionPeering)(nil)).Elem()
}

func (o ExpressRouteCrossConnectionPeeringOutput) ToExpressRouteCrossConnectionPeeringOutput() ExpressRouteCrossConnectionPeeringOutput {
	return o
}

func (o ExpressRouteCrossConnectionPeeringOutput) ToExpressRouteCrossConnectionPeeringOutputWithContext(ctx context.Context) ExpressRouteCrossConnectionPeeringOutput {
	return o
}

// The Azure ASN.
func (o ExpressRouteCrossConnectionPeeringOutput) AzureASN() pulumi.IntOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.IntOutput { return v.AzureASN }).(pulumi.IntOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ExpressRouteCrossConnectionPeeringOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The GatewayManager Etag.
func (o ExpressRouteCrossConnectionPeeringOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

// The IPv6 peering configuration.
func (o ExpressRouteCrossConnectionPeeringOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

// Who was the last to modify the peering.
func (o ExpressRouteCrossConnectionPeeringOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringOutput { return v.LastModifiedBy }).(pulumi.StringOutput)
}

// The Microsoft peering configuration.
func (o ExpressRouteCrossConnectionPeeringOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) ExpressRouteCircuitPeeringConfigResponsePtrOutput {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o ExpressRouteCrossConnectionPeeringOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The peer ASN.
func (o ExpressRouteCrossConnectionPeeringOutput) PeerASN() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.Float64PtrOutput { return v.PeerASN }).(pulumi.Float64PtrOutput)
}

// The peering type.
func (o ExpressRouteCrossConnectionPeeringOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.PeeringType }).(pulumi.StringPtrOutput)
}

// The primary port.
func (o ExpressRouteCrossConnectionPeeringOutput) PrimaryAzurePort() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringOutput { return v.PrimaryAzurePort }).(pulumi.StringOutput)
}

// The primary address prefix.
func (o ExpressRouteCrossConnectionPeeringOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

// The provisioning state of the express route cross connection peering resource.
func (o ExpressRouteCrossConnectionPeeringOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The secondary port.
func (o ExpressRouteCrossConnectionPeeringOutput) SecondaryAzurePort() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringOutput { return v.SecondaryAzurePort }).(pulumi.StringOutput)
}

// The secondary address prefix.
func (o ExpressRouteCrossConnectionPeeringOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput {
		return v.SecondaryPeerAddressPrefix
	}).(pulumi.StringPtrOutput)
}

// The shared key.
func (o ExpressRouteCrossConnectionPeeringOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.SharedKey }).(pulumi.StringPtrOutput)
}

// The peering state.
func (o ExpressRouteCrossConnectionPeeringOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// The VLAN ID.
func (o ExpressRouteCrossConnectionPeeringOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.IntPtrOutput { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCrossConnectionPeeringOutput{})
}
