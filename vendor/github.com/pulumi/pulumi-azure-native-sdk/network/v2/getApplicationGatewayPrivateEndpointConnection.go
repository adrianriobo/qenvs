// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified private endpoint connection on application gateway.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01.
func LookupApplicationGatewayPrivateEndpointConnection(ctx *pulumi.Context, args *LookupApplicationGatewayPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationGatewayPrivateEndpointConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationGatewayPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:network:getApplicationGatewayPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApplicationGatewayPrivateEndpointConnectionArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName string `pulumi:"applicationGatewayName"`
	// The name of the application gateway private endpoint connection.
	ConnectionName string `pulumi:"connectionName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Private Endpoint connection on an application gateway.
type LookupApplicationGatewayPrivateEndpointConnectionResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The consumer link id.
	LinkIdentifier string `pulumi:"linkIdentifier"`
	// Name of the private endpoint connection on an application gateway.
	Name *string `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the application gateway private endpoint connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupApplicationGatewayPrivateEndpointConnectionResult
func (val *LookupApplicationGatewayPrivateEndpointConnectionResult) Defaults() *LookupApplicationGatewayPrivateEndpointConnectionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PrivateEndpoint = *tmp.PrivateEndpoint.Defaults()

	return &tmp
}

func LookupApplicationGatewayPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupApplicationGatewayPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationGatewayPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationGatewayPrivateEndpointConnectionResult, error) {
			args := v.(LookupApplicationGatewayPrivateEndpointConnectionArgs)
			r, err := LookupApplicationGatewayPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupApplicationGatewayPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationGatewayPrivateEndpointConnectionResultOutput)
}

type LookupApplicationGatewayPrivateEndpointConnectionOutputArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName pulumi.StringInput `pulumi:"applicationGatewayName"`
	// The name of the application gateway private endpoint connection.
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationGatewayPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGatewayPrivateEndpointConnectionArgs)(nil)).Elem()
}

// Private Endpoint connection on an application gateway.
type LookupApplicationGatewayPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationGatewayPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationGatewayPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) ToLookupApplicationGatewayPrivateEndpointConnectionResultOutput() LookupApplicationGatewayPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) ToLookupApplicationGatewayPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupApplicationGatewayPrivateEndpointConnectionResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The consumer link id.
func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) LinkIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) string { return v.LinkIdentifier }).(pulumi.StringOutput)
}

// Name of the private endpoint connection on an application gateway.
func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The resource of private end point.
func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponseOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

// The provisioning state of the application gateway private endpoint connection resource.
func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Type of the resource.
func (o LookupApplicationGatewayPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationGatewayPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationGatewayPrivateEndpointConnectionResultOutput{})
}
