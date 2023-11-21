// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Traffic Manager endpoint.
// Azure REST API version: 2022-04-01.
//
// Other available API versions: 2017-03-01, 2018-02-01, 2022-04-01-preview.
func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEndpointResult
	err := ctx.Invoke("azure-native:network:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	// The name of the Traffic Manager endpoint.
	EndpointName string `pulumi:"endpointName"`
	// The type of the Traffic Manager endpoint.
	EndpointType string `pulumi:"endpointType"`
	// The name of the Traffic Manager profile.
	ProfileName string `pulumi:"profileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a Traffic Manager endpoint.
type LookupEndpointResult struct {
	// If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in the traffic routing method.
	AlwaysServe *string `pulumi:"alwaysServe"`
	// List of custom headers.
	CustomHeaders []EndpointPropertiesResponseCustomHeaders `pulumi:"customHeaders"`
	// Specifies the location of the external or nested endpoints when using the 'Performance' traffic routing method.
	EndpointLocation *string `pulumi:"endpointLocation"`
	// The monitoring status of the endpoint.
	EndpointMonitorStatus *string `pulumi:"endpointMonitorStatus"`
	// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
	EndpointStatus *string `pulumi:"endpointStatus"`
	// The list of countries/regions mapped to this endpoint when using the 'Geographic' traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
	GeoMapping []string `pulumi:"geoMapping"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id *string `pulumi:"id"`
	// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints *float64 `pulumi:"minChildEndpoints"`
	// The minimum number of IPv4 (DNS record type A) endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpointsIPv4 *float64 `pulumi:"minChildEndpointsIPv4"`
	// The minimum number of IPv6 (DNS record type AAAA) endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpointsIPv6 *float64 `pulumi:"minChildEndpointsIPv6"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The priority of this endpoint when using the 'Priority' traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
	Priority *float64 `pulumi:"priority"`
	// The list of subnets, IP addresses, and/or address ranges mapped to this endpoint when using the 'Subnet' traffic routing method. An empty list will match all ranges not covered by other endpoints.
	Subnets []EndpointPropertiesResponseSubnets `pulumi:"subnets"`
	// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target *string `pulumi:"target"`
	// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `pulumi:"type"`
	// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight *float64 `pulumi:"weight"`
}

func LookupEndpointOutput(ctx *pulumi.Context, args LookupEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointResult, error) {
			args := v.(LookupEndpointArgs)
			r, err := LookupEndpoint(ctx, &args, opts...)
			var s LookupEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointResultOutput)
}

type LookupEndpointOutputArgs struct {
	// The name of the Traffic Manager endpoint.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The type of the Traffic Manager endpoint.
	EndpointType pulumi.StringInput `pulumi:"endpointType"`
	// The name of the Traffic Manager profile.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointArgs)(nil)).Elem()
}

// Class representing a Traffic Manager endpoint.
type LookupEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointResult)(nil)).Elem()
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutput() LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutputWithContext(ctx context.Context) LookupEndpointResultOutput {
	return o
}

// If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in the traffic routing method.
func (o LookupEndpointResultOutput) AlwaysServe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.AlwaysServe }).(pulumi.StringPtrOutput)
}

// List of custom headers.
func (o LookupEndpointResultOutput) CustomHeaders() EndpointPropertiesResponseCustomHeadersArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []EndpointPropertiesResponseCustomHeaders { return v.CustomHeaders }).(EndpointPropertiesResponseCustomHeadersArrayOutput)
}

// Specifies the location of the external or nested endpoints when using the 'Performance' traffic routing method.
func (o LookupEndpointResultOutput) EndpointLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EndpointLocation }).(pulumi.StringPtrOutput)
}

// The monitoring status of the endpoint.
func (o LookupEndpointResultOutput) EndpointMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EndpointMonitorStatus }).(pulumi.StringPtrOutput)
}

// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
func (o LookupEndpointResultOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

// The list of countries/regions mapped to this endpoint when using the 'Geographic' traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
func (o LookupEndpointResultOutput) GeoMapping() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []string { return v.GeoMapping }).(pulumi.StringArrayOutput)
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
func (o LookupEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
func (o LookupEndpointResultOutput) MinChildEndpoints() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *float64 { return v.MinChildEndpoints }).(pulumi.Float64PtrOutput)
}

// The minimum number of IPv4 (DNS record type A) endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
func (o LookupEndpointResultOutput) MinChildEndpointsIPv4() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *float64 { return v.MinChildEndpointsIPv4 }).(pulumi.Float64PtrOutput)
}

// The minimum number of IPv6 (DNS record type AAAA) endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
func (o LookupEndpointResultOutput) MinChildEndpointsIPv6() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *float64 { return v.MinChildEndpointsIPv6 }).(pulumi.Float64PtrOutput)
}

// The name of the resource
func (o LookupEndpointResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The priority of this endpoint when using the 'Priority' traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
func (o LookupEndpointResultOutput) Priority() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *float64 { return v.Priority }).(pulumi.Float64PtrOutput)
}

// The list of subnets, IP addresses, and/or address ranges mapped to this endpoint when using the 'Subnet' traffic routing method. An empty list will match all ranges not covered by other endpoints.
func (o LookupEndpointResultOutput) Subnets() EndpointPropertiesResponseSubnetsArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []EndpointPropertiesResponseSubnets { return v.Subnets }).(EndpointPropertiesResponseSubnetsArrayOutput)
}

// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
func (o LookupEndpointResultOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Target }).(pulumi.StringPtrOutput)
}

// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
func (o LookupEndpointResultOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
func (o LookupEndpointResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
func (o LookupEndpointResultOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *float64 { return v.Weight }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointResultOutput{})
}
