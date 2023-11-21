// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified peering for the express route circuit.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2017-09-01, 2019-02-01, 2019-06-01, 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01.
func LookupExpressRouteCircuitPeering(ctx *pulumi.Context, args *LookupExpressRouteCircuitPeeringArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitPeeringResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExpressRouteCircuitPeeringResult
	err := ctx.Invoke("azure-native:network:getExpressRouteCircuitPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitPeeringArgs struct {
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Peering in an ExpressRouteCircuit resource.
type LookupExpressRouteCircuitPeeringResult struct {
	// The Azure ASN.
	AzureASN *int `pulumi:"azureASN"`
	// The list of circuit connections associated with Azure Private Peering for this circuit.
	Connections []ExpressRouteCircuitConnectionResponse `pulumi:"connections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The ExpressRoute connection.
	ExpressRouteConnection *ExpressRouteConnectionIdResponse `pulumi:"expressRouteConnection"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The IPv6 peering configuration.
	Ipv6PeeringConfig *Ipv6ExpressRouteCircuitPeeringConfigResponse `pulumi:"ipv6PeeringConfig"`
	// Who was the last to modify the peering.
	LastModifiedBy string `pulumi:"lastModifiedBy"`
	// The Microsoft peering configuration.
	MicrosoftPeeringConfig *ExpressRouteCircuitPeeringConfigResponse `pulumi:"microsoftPeeringConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The peer ASN.
	PeerASN *float64 `pulumi:"peerASN"`
	// The list of peered circuit connections associated with Azure Private Peering for this circuit.
	PeeredConnections []PeerExpressRouteCircuitConnectionResponse `pulumi:"peeredConnections"`
	// The peering type.
	PeeringType *string `pulumi:"peeringType"`
	// The primary port.
	PrimaryAzurePort *string `pulumi:"primaryAzurePort"`
	// The primary address prefix.
	PrimaryPeerAddressPrefix *string `pulumi:"primaryPeerAddressPrefix"`
	// The provisioning state of the express route circuit peering resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The reference to the RouteFilter resource.
	RouteFilter *SubResourceResponse `pulumi:"routeFilter"`
	// The secondary port.
	SecondaryAzurePort *string `pulumi:"secondaryAzurePort"`
	// The secondary address prefix.
	SecondaryPeerAddressPrefix *string `pulumi:"secondaryPeerAddressPrefix"`
	// The shared key.
	SharedKey *string `pulumi:"sharedKey"`
	// The peering state.
	State *string `pulumi:"state"`
	// The peering stats of express route circuit.
	Stats *ExpressRouteCircuitStatsResponse `pulumi:"stats"`
	// Type of the resource.
	Type string `pulumi:"type"`
	// The VLAN ID.
	VlanId *int `pulumi:"vlanId"`
}

func LookupExpressRouteCircuitPeeringOutput(ctx *pulumi.Context, args LookupExpressRouteCircuitPeeringOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteCircuitPeeringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteCircuitPeeringResult, error) {
			args := v.(LookupExpressRouteCircuitPeeringArgs)
			r, err := LookupExpressRouteCircuitPeering(ctx, &args, opts...)
			var s LookupExpressRouteCircuitPeeringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteCircuitPeeringResultOutput)
}

type LookupExpressRouteCircuitPeeringOutputArgs struct {
	// The name of the express route circuit.
	CircuitName pulumi.StringInput `pulumi:"circuitName"`
	// The name of the peering.
	PeeringName pulumi.StringInput `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteCircuitPeeringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitPeeringArgs)(nil)).Elem()
}

// Peering in an ExpressRouteCircuit resource.
type LookupExpressRouteCircuitPeeringResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteCircuitPeeringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitPeeringResult)(nil)).Elem()
}

func (o LookupExpressRouteCircuitPeeringResultOutput) ToLookupExpressRouteCircuitPeeringResultOutput() LookupExpressRouteCircuitPeeringResultOutput {
	return o
}

func (o LookupExpressRouteCircuitPeeringResultOutput) ToLookupExpressRouteCircuitPeeringResultOutputWithContext(ctx context.Context) LookupExpressRouteCircuitPeeringResultOutput {
	return o
}

// The Azure ASN.
func (o LookupExpressRouteCircuitPeeringResultOutput) AzureASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *int { return v.AzureASN }).(pulumi.IntPtrOutput)
}

// The list of circuit connections associated with Azure Private Peering for this circuit.
func (o LookupExpressRouteCircuitPeeringResultOutput) Connections() ExpressRouteCircuitConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) []ExpressRouteCircuitConnectionResponse {
		return v.Connections
	}).(ExpressRouteCircuitConnectionResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupExpressRouteCircuitPeeringResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The ExpressRoute connection.
func (o LookupExpressRouteCircuitPeeringResultOutput) ExpressRouteConnection() ExpressRouteConnectionIdResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *ExpressRouteConnectionIdResponse {
		return v.ExpressRouteConnection
	}).(ExpressRouteConnectionIdResponsePtrOutput)
}

// The GatewayManager Etag.
func (o LookupExpressRouteCircuitPeeringResultOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupExpressRouteCircuitPeeringResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The IPv6 peering configuration.
func (o LookupExpressRouteCircuitPeeringResultOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *Ipv6ExpressRouteCircuitPeeringConfigResponse {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

// Who was the last to modify the peering.
func (o LookupExpressRouteCircuitPeeringResultOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

// The Microsoft peering configuration.
func (o LookupExpressRouteCircuitPeeringResultOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *ExpressRouteCircuitPeeringConfigResponse {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupExpressRouteCircuitPeeringResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The peer ASN.
func (o LookupExpressRouteCircuitPeeringResultOutput) PeerASN() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *float64 { return v.PeerASN }).(pulumi.Float64PtrOutput)
}

// The list of peered circuit connections associated with Azure Private Peering for this circuit.
func (o LookupExpressRouteCircuitPeeringResultOutput) PeeredConnections() PeerExpressRouteCircuitConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) []PeerExpressRouteCircuitConnectionResponse {
		return v.PeeredConnections
	}).(PeerExpressRouteCircuitConnectionResponseArrayOutput)
}

// The peering type.
func (o LookupExpressRouteCircuitPeeringResultOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.PeeringType }).(pulumi.StringPtrOutput)
}

// The primary port.
func (o LookupExpressRouteCircuitPeeringResultOutput) PrimaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.PrimaryAzurePort }).(pulumi.StringPtrOutput)
}

// The primary address prefix.
func (o LookupExpressRouteCircuitPeeringResultOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

// The provisioning state of the express route circuit peering resource.
func (o LookupExpressRouteCircuitPeeringResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The reference to the RouteFilter resource.
func (o LookupExpressRouteCircuitPeeringResultOutput) RouteFilter() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *SubResourceResponse { return v.RouteFilter }).(SubResourceResponsePtrOutput)
}

// The secondary port.
func (o LookupExpressRouteCircuitPeeringResultOutput) SecondaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.SecondaryAzurePort }).(pulumi.StringPtrOutput)
}

// The secondary address prefix.
func (o LookupExpressRouteCircuitPeeringResultOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

// The shared key.
func (o LookupExpressRouteCircuitPeeringResultOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

// The peering state.
func (o LookupExpressRouteCircuitPeeringResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The peering stats of express route circuit.
func (o LookupExpressRouteCircuitPeeringResultOutput) Stats() ExpressRouteCircuitStatsResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *ExpressRouteCircuitStatsResponse { return v.Stats }).(ExpressRouteCircuitStatsResponsePtrOutput)
}

// Type of the resource.
func (o LookupExpressRouteCircuitPeeringResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) string { return v.Type }).(pulumi.StringOutput)
}

// The VLAN ID.
func (o LookupExpressRouteCircuitPeeringResultOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteCircuitPeeringResultOutput{})
}
