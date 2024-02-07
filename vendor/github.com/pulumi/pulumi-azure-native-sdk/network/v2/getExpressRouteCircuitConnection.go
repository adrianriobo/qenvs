// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Express Route Circuit Connection from the specified express route circuit.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01.
func LookupExpressRouteCircuitConnection(ctx *pulumi.Context, args *LookupExpressRouteCircuitConnectionArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExpressRouteCircuitConnectionResult
	err := ctx.Invoke("azure-native:network:getExpressRouteCircuitConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitConnectionArgs struct {
	// The name of the express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// The name of the express route circuit connection.
	ConnectionName string `pulumi:"connectionName"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.
type LookupExpressRouteCircuitConnectionResult struct {
	// /29 IP address space to carve out Customer addresses for tunnels.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// The authorization key.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// Express Route Circuit connection state.
	CircuitConnectionStatus string `pulumi:"circuitConnectionStatus"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
	ExpressRouteCircuitPeering *SubResourceResponse `pulumi:"expressRouteCircuitPeering"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IPv6 Address PrefixProperties of the express route circuit connection.
	Ipv6CircuitConnectionConfig *Ipv6CircuitConnectionConfigResponse `pulumi:"ipv6CircuitConnectionConfig"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
	PeerExpressRouteCircuitPeering *SubResourceResponse `pulumi:"peerExpressRouteCircuitPeering"`
	// The provisioning state of the express route circuit connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

func LookupExpressRouteCircuitConnectionOutput(ctx *pulumi.Context, args LookupExpressRouteCircuitConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteCircuitConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteCircuitConnectionResult, error) {
			args := v.(LookupExpressRouteCircuitConnectionArgs)
			r, err := LookupExpressRouteCircuitConnection(ctx, &args, opts...)
			var s LookupExpressRouteCircuitConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteCircuitConnectionResultOutput)
}

type LookupExpressRouteCircuitConnectionOutputArgs struct {
	// The name of the express route circuit.
	CircuitName pulumi.StringInput `pulumi:"circuitName"`
	// The name of the express route circuit connection.
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The name of the peering.
	PeeringName pulumi.StringInput `pulumi:"peeringName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteCircuitConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitConnectionArgs)(nil)).Elem()
}

// Express Route Circuit Connection in an ExpressRouteCircuitPeering resource.
type LookupExpressRouteCircuitConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteCircuitConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitConnectionResult)(nil)).Elem()
}

func (o LookupExpressRouteCircuitConnectionResultOutput) ToLookupExpressRouteCircuitConnectionResultOutput() LookupExpressRouteCircuitConnectionResultOutput {
	return o
}

func (o LookupExpressRouteCircuitConnectionResultOutput) ToLookupExpressRouteCircuitConnectionResultOutputWithContext(ctx context.Context) LookupExpressRouteCircuitConnectionResultOutput {
	return o
}

// /29 IP address space to carve out Customer addresses for tunnels.
func (o LookupExpressRouteCircuitConnectionResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

// The authorization key.
func (o LookupExpressRouteCircuitConnectionResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

// Express Route Circuit connection state.
func (o LookupExpressRouteCircuitConnectionResultOutput) CircuitConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) string { return v.CircuitConnectionStatus }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupExpressRouteCircuitConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Reference to Express Route Circuit Private Peering Resource of the circuit initiating connection.
func (o LookupExpressRouteCircuitConnectionResultOutput) ExpressRouteCircuitPeering() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *SubResourceResponse {
		return v.ExpressRouteCircuitPeering
	}).(SubResourceResponsePtrOutput)
}

// Resource ID.
func (o LookupExpressRouteCircuitConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// IPv6 Address PrefixProperties of the express route circuit connection.
func (o LookupExpressRouteCircuitConnectionResultOutput) Ipv6CircuitConnectionConfig() Ipv6CircuitConnectionConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *Ipv6CircuitConnectionConfigResponse {
		return v.Ipv6CircuitConnectionConfig
	}).(Ipv6CircuitConnectionConfigResponsePtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupExpressRouteCircuitConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Reference to Express Route Circuit Private Peering Resource of the peered circuit.
func (o LookupExpressRouteCircuitConnectionResultOutput) PeerExpressRouteCircuitPeering() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *SubResourceResponse {
		return v.PeerExpressRouteCircuitPeering
	}).(SubResourceResponsePtrOutput)
}

// The provisioning state of the express route circuit connection resource.
func (o LookupExpressRouteCircuitConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Type of the resource.
func (o LookupExpressRouteCircuitConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteCircuitConnectionResultOutput{})
}
