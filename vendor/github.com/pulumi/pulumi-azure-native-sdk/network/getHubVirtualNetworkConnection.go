// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the details of a HubVirtualNetworkConnection.
// API Version: 2020-11-01.
func LookupHubVirtualNetworkConnection(ctx *pulumi.Context, args *LookupHubVirtualNetworkConnectionArgs, opts ...pulumi.InvokeOption) (*LookupHubVirtualNetworkConnectionResult, error) {
	var rv LookupHubVirtualNetworkConnectionResult
	err := ctx.Invoke("azure-native:network:getHubVirtualNetworkConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubVirtualNetworkConnectionArgs struct {
	// The name of the vpn connection.
	ConnectionName string `pulumi:"connectionName"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// HubVirtualNetworkConnection Resource.
type LookupHubVirtualNetworkConnectionResult struct {
	// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
	AllowHubToRemoteVnetTransit *bool `pulumi:"allowHubToRemoteVnetTransit"`
	// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
	AllowRemoteVnetToUseHubVnetGateways *bool `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the hub virtual network connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Reference to the remote virtual network.
	RemoteVirtualNetwork *SubResourceResponse `pulumi:"remoteVirtualNetwork"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration *RoutingConfigurationResponse `pulumi:"routingConfiguration"`
}

func LookupHubVirtualNetworkConnectionOutput(ctx *pulumi.Context, args LookupHubVirtualNetworkConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupHubVirtualNetworkConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHubVirtualNetworkConnectionResult, error) {
			args := v.(LookupHubVirtualNetworkConnectionArgs)
			r, err := LookupHubVirtualNetworkConnection(ctx, &args, opts...)
			var s LookupHubVirtualNetworkConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHubVirtualNetworkConnectionResultOutput)
}

type LookupHubVirtualNetworkConnectionOutputArgs struct {
	// The name of the vpn connection.
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupHubVirtualNetworkConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubVirtualNetworkConnectionArgs)(nil)).Elem()
}

// HubVirtualNetworkConnection Resource.
type LookupHubVirtualNetworkConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupHubVirtualNetworkConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHubVirtualNetworkConnectionResult)(nil)).Elem()
}

func (o LookupHubVirtualNetworkConnectionResultOutput) ToLookupHubVirtualNetworkConnectionResultOutput() LookupHubVirtualNetworkConnectionResultOutput {
	return o
}

func (o LookupHubVirtualNetworkConnectionResultOutput) ToLookupHubVirtualNetworkConnectionResultOutputWithContext(ctx context.Context) LookupHubVirtualNetworkConnectionResultOutput {
	return o
}

// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
func (o LookupHubVirtualNetworkConnectionResultOutput) AllowHubToRemoteVnetTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *bool { return v.AllowHubToRemoteVnetTransit }).(pulumi.BoolPtrOutput)
}

// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
func (o LookupHubVirtualNetworkConnectionResultOutput) AllowRemoteVnetToUseHubVnetGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *bool { return v.AllowRemoteVnetToUseHubVnetGateways }).(pulumi.BoolPtrOutput)
}

// Enable internet security.
func (o LookupHubVirtualNetworkConnectionResultOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *bool { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupHubVirtualNetworkConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupHubVirtualNetworkConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupHubVirtualNetworkConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the hub virtual network connection resource.
func (o LookupHubVirtualNetworkConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Reference to the remote virtual network.
func (o LookupHubVirtualNetworkConnectionResultOutput) RemoteVirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *SubResourceResponse { return v.RemoteVirtualNetwork }).(SubResourceResponsePtrOutput)
}

// The Routing Configuration indicating the associated and propagated route tables on this connection.
func (o LookupHubVirtualNetworkConnectionResultOutput) RoutingConfiguration() RoutingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupHubVirtualNetworkConnectionResult) *RoutingConfigurationResponse {
		return v.RoutingConfiguration
	}).(RoutingConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHubVirtualNetworkConnectionResultOutput{})
}
