// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a private endpoint connection under a disk access resource.
// API Version: 2020-12-01.
func LookupDiskAccessAPrivateEndpointConnection(ctx *pulumi.Context, args *LookupDiskAccessAPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupDiskAccessAPrivateEndpointConnectionResult, error) {
	var rv LookupDiskAccessAPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:compute:getDiskAccessAPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskAccessAPrivateEndpointConnectionArgs struct {
	// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	DiskAccessName string `pulumi:"diskAccessName"`
	// The name of the private endpoint connection
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The Private Endpoint Connection resource.
type LookupDiskAccessAPrivateEndpointConnectionResult struct {
	// private endpoint connection Id
	Id string `pulumi:"id"`
	// private endpoint connection name
	Name string `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between DiskAccess and Virtual Network.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// private endpoint connection type
	Type string `pulumi:"type"`
}

func LookupDiskAccessAPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupDiskAccessAPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupDiskAccessAPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskAccessAPrivateEndpointConnectionResult, error) {
			args := v.(LookupDiskAccessAPrivateEndpointConnectionArgs)
			r, err := LookupDiskAccessAPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupDiskAccessAPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskAccessAPrivateEndpointConnectionResultOutput)
}

type LookupDiskAccessAPrivateEndpointConnectionOutputArgs struct {
	// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	DiskAccessName pulumi.StringInput `pulumi:"diskAccessName"`
	// The name of the private endpoint connection
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDiskAccessAPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskAccessAPrivateEndpointConnectionArgs)(nil)).Elem()
}

// The Private Endpoint Connection resource.
type LookupDiskAccessAPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupDiskAccessAPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskAccessAPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) ToLookupDiskAccessAPrivateEndpointConnectionResultOutput() LookupDiskAccessAPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) ToLookupDiskAccessAPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupDiskAccessAPrivateEndpointConnectionResultOutput {
	return o
}

// private endpoint connection Id
func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// private endpoint connection name
func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) PrivateEndpointResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponseOutput)
}

// A collection of information about the state of the connection between DiskAccess and Virtual Network.
func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// private endpoint connection type
func (o LookupDiskAccessAPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskAccessAPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskAccessAPrivateEndpointConnectionResultOutput{})
}
