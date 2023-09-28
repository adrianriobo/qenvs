// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified private link service by resource group.
// Azure REST API version: 2023-02-01.
func LookupPrivateLinkService(ctx *pulumi.Context, args *LookupPrivateLinkServiceArgs, opts ...pulumi.InvokeOption) (*LookupPrivateLinkServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateLinkServiceResult
	err := ctx.Invoke("azure-native:network:getPrivateLinkService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateLinkServiceArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the private link service.
	ServiceName string `pulumi:"serviceName"`
}

// Private link service resource.
type LookupPrivateLinkServiceResult struct {
	// The alias of the private link service.
	Alias string `pulumi:"alias"`
	// The auto-approval list of the private link service.
	AutoApproval *PrivateLinkServicePropertiesResponseAutoApproval `pulumi:"autoApproval"`
	// Whether the private link service is enabled for proxy protocol or not.
	EnableProxyProtocol *bool `pulumi:"enableProxyProtocol"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The extended location of the load balancer.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// The list of Fqdn.
	Fqdns []string `pulumi:"fqdns"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// An array of private link service IP configurations.
	IpConfigurations []PrivateLinkServiceIpConfigurationResponse `pulumi:"ipConfigurations"`
	// An array of references to the load balancer IP configurations.
	LoadBalancerFrontendIpConfigurations []FrontendIPConfigurationResponse `pulumi:"loadBalancerFrontendIpConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// An array of references to the network interfaces created for this private link service.
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	// An array of list about connections to the private endpoint.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provisioning state of the private link service resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The visibility list of the private link service.
	Visibility *PrivateLinkServicePropertiesResponseVisibility `pulumi:"visibility"`
}

func LookupPrivateLinkServiceOutput(ctx *pulumi.Context, args LookupPrivateLinkServiceOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateLinkServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateLinkServiceResult, error) {
			args := v.(LookupPrivateLinkServiceArgs)
			r, err := LookupPrivateLinkService(ctx, &args, opts...)
			var s LookupPrivateLinkServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateLinkServiceResultOutput)
}

type LookupPrivateLinkServiceOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the private link service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupPrivateLinkServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServiceArgs)(nil)).Elem()
}

// Private link service resource.
type LookupPrivateLinkServiceResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateLinkServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateLinkServiceResult)(nil)).Elem()
}

func (o LookupPrivateLinkServiceResultOutput) ToLookupPrivateLinkServiceResultOutput() LookupPrivateLinkServiceResultOutput {
	return o
}

func (o LookupPrivateLinkServiceResultOutput) ToLookupPrivateLinkServiceResultOutputWithContext(ctx context.Context) LookupPrivateLinkServiceResultOutput {
	return o
}

func (o LookupPrivateLinkServiceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateLinkServiceResult] {
	return pulumix.Output[LookupPrivateLinkServiceResult]{
		OutputState: o.OutputState,
	}
}

// The alias of the private link service.
func (o LookupPrivateLinkServiceResultOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) string { return v.Alias }).(pulumi.StringOutput)
}

// The auto-approval list of the private link service.
func (o LookupPrivateLinkServiceResultOutput) AutoApproval() PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *PrivateLinkServicePropertiesResponseAutoApproval {
		return v.AutoApproval
	}).(PrivateLinkServicePropertiesResponseAutoApprovalPtrOutput)
}

// Whether the private link service is enabled for proxy protocol or not.
func (o LookupPrivateLinkServiceResultOutput) EnableProxyProtocol() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *bool { return v.EnableProxyProtocol }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupPrivateLinkServiceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the load balancer.
func (o LookupPrivateLinkServiceResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The list of Fqdn.
func (o LookupPrivateLinkServiceResultOutput) Fqdns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) []string { return v.Fqdns }).(pulumi.StringArrayOutput)
}

// Resource ID.
func (o LookupPrivateLinkServiceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// An array of private link service IP configurations.
func (o LookupPrivateLinkServiceResultOutput) IpConfigurations() PrivateLinkServiceIpConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) []PrivateLinkServiceIpConfigurationResponse {
		return v.IpConfigurations
	}).(PrivateLinkServiceIpConfigurationResponseArrayOutput)
}

// An array of references to the load balancer IP configurations.
func (o LookupPrivateLinkServiceResultOutput) LoadBalancerFrontendIpConfigurations() FrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) []FrontendIPConfigurationResponse {
		return v.LoadBalancerFrontendIpConfigurations
	}).(FrontendIPConfigurationResponseArrayOutput)
}

// Resource location.
func (o LookupPrivateLinkServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupPrivateLinkServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// An array of references to the network interfaces created for this private link service.
func (o LookupPrivateLinkServiceResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

// An array of list about connections to the private endpoint.
func (o LookupPrivateLinkServiceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state of the private link service resource.
func (o LookupPrivateLinkServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupPrivateLinkServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupPrivateLinkServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The visibility list of the private link service.
func (o LookupPrivateLinkServiceResultOutput) Visibility() PrivateLinkServicePropertiesResponseVisibilityPtrOutput {
	return o.ApplyT(func(v LookupPrivateLinkServiceResult) *PrivateLinkServicePropertiesResponseVisibility {
		return v.Visibility
	}).(PrivateLinkServicePropertiesResponseVisibilityPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateLinkServiceResultOutput{})
}
