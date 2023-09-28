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

// Class representing a Traffic Manager endpoint.
// Azure REST API version: 2022-04-01. Prior API version in Azure Native 1.x: 2018-08-01
type Endpoint struct {
	pulumi.CustomResourceState

	// If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in the traffic routing method.
	AlwaysServe pulumi.StringPtrOutput `pulumi:"alwaysServe"`
	// List of custom headers.
	CustomHeaders EndpointPropertiesResponseCustomHeadersArrayOutput `pulumi:"customHeaders"`
	// Specifies the location of the external or nested endpoints when using the 'Performance' traffic routing method.
	EndpointLocation pulumi.StringPtrOutput `pulumi:"endpointLocation"`
	// The monitoring status of the endpoint.
	EndpointMonitorStatus pulumi.StringPtrOutput `pulumi:"endpointMonitorStatus"`
	// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
	EndpointStatus pulumi.StringPtrOutput `pulumi:"endpointStatus"`
	// The list of countries/regions mapped to this endpoint when using the 'Geographic' traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
	GeoMapping pulumi.StringArrayOutput `pulumi:"geoMapping"`
	// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints pulumi.Float64PtrOutput `pulumi:"minChildEndpoints"`
	// The minimum number of IPv4 (DNS record type A) endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpointsIPv4 pulumi.Float64PtrOutput `pulumi:"minChildEndpointsIPv4"`
	// The minimum number of IPv6 (DNS record type AAAA) endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpointsIPv6 pulumi.Float64PtrOutput `pulumi:"minChildEndpointsIPv6"`
	// The name of the resource
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The priority of this endpoint when using the 'Priority' traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
	Priority pulumi.Float64PtrOutput `pulumi:"priority"`
	// The list of subnets, IP addresses, and/or address ranges mapped to this endpoint when using the 'Subnet' traffic routing method. An empty list will match all ranges not covered by other endpoints.
	Subnets EndpointPropertiesResponseSubnetsArrayOutput `pulumi:"subnets"`
	// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target pulumi.StringPtrOutput `pulumi:"target"`
	// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceId pulumi.StringPtrOutput `pulumi:"targetResourceId"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight pulumi.Float64PtrOutput `pulumi:"weight"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointType == nil {
		return nil, errors.New("invalid value for required argument 'EndpointType'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20151101:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170501:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401:Endpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:Endpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Endpoint
	err := ctx.RegisterResource("azure-native:network:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("azure-native:network:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
}

type EndpointState struct {
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in the traffic routing method.
	AlwaysServe *string `pulumi:"alwaysServe"`
	// List of custom headers.
	CustomHeaders []EndpointPropertiesCustomHeaders `pulumi:"customHeaders"`
	// Specifies the location of the external or nested endpoints when using the 'Performance' traffic routing method.
	EndpointLocation *string `pulumi:"endpointLocation"`
	// The monitoring status of the endpoint.
	EndpointMonitorStatus *string `pulumi:"endpointMonitorStatus"`
	// The name of the Traffic Manager endpoint to be created or updated.
	EndpointName *string `pulumi:"endpointName"`
	// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
	EndpointStatus *string `pulumi:"endpointStatus"`
	// The type of the Traffic Manager endpoint to be created or updated.
	EndpointType string `pulumi:"endpointType"`
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
	// The name of the Traffic Manager profile.
	ProfileName string `pulumi:"profileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of subnets, IP addresses, and/or address ranges mapped to this endpoint when using the 'Subnet' traffic routing method. An empty list will match all ranges not covered by other endpoints.
	Subnets []EndpointPropertiesSubnets `pulumi:"subnets"`
	// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target *string `pulumi:"target"`
	// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type *string `pulumi:"type"`
	// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight *float64 `pulumi:"weight"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in the traffic routing method.
	AlwaysServe pulumi.StringPtrInput
	// List of custom headers.
	CustomHeaders EndpointPropertiesCustomHeadersArrayInput
	// Specifies the location of the external or nested endpoints when using the 'Performance' traffic routing method.
	EndpointLocation pulumi.StringPtrInput
	// The monitoring status of the endpoint.
	EndpointMonitorStatus pulumi.StringPtrInput
	// The name of the Traffic Manager endpoint to be created or updated.
	EndpointName pulumi.StringPtrInput
	// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
	EndpointStatus pulumi.StringPtrInput
	// The type of the Traffic Manager endpoint to be created or updated.
	EndpointType pulumi.StringInput
	// The list of countries/regions mapped to this endpoint when using the 'Geographic' traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
	GeoMapping pulumi.StringArrayInput
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
	Id pulumi.StringPtrInput
	// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpoints pulumi.Float64PtrInput
	// The minimum number of IPv4 (DNS record type A) endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpointsIPv4 pulumi.Float64PtrInput
	// The minimum number of IPv6 (DNS record type AAAA) endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
	MinChildEndpointsIPv6 pulumi.Float64PtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The priority of this endpoint when using the 'Priority' traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
	Priority pulumi.Float64PtrInput
	// The name of the Traffic Manager profile.
	ProfileName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The list of subnets, IP addresses, and/or address ranges mapped to this endpoint when using the 'Subnet' traffic routing method. An empty list will match all ranges not covered by other endpoints.
	Subnets EndpointPropertiesSubnetsArrayInput
	// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
	Target pulumi.StringPtrInput
	// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
	TargetResourceId pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
	Type pulumi.StringPtrInput
	// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
	Weight pulumi.Float64PtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

func (i *Endpoint) ToOutput(ctx context.Context) pulumix.Output[*Endpoint] {
	return pulumix.Output[*Endpoint]{
		OutputState: i.ToEndpointOutputWithContext(ctx).OutputState,
	}
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func (o EndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*Endpoint] {
	return pulumix.Output[*Endpoint]{
		OutputState: o.OutputState,
	}
}

// If Always Serve is enabled, probing for endpoint health will be disabled and endpoints will be included in the traffic routing method.
func (o EndpointOutput) AlwaysServe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.AlwaysServe }).(pulumi.StringPtrOutput)
}

// List of custom headers.
func (o EndpointOutput) CustomHeaders() EndpointPropertiesResponseCustomHeadersArrayOutput {
	return o.ApplyT(func(v *Endpoint) EndpointPropertiesResponseCustomHeadersArrayOutput { return v.CustomHeaders }).(EndpointPropertiesResponseCustomHeadersArrayOutput)
}

// Specifies the location of the external or nested endpoints when using the 'Performance' traffic routing method.
func (o EndpointOutput) EndpointLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.EndpointLocation }).(pulumi.StringPtrOutput)
}

// The monitoring status of the endpoint.
func (o EndpointOutput) EndpointMonitorStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.EndpointMonitorStatus }).(pulumi.StringPtrOutput)
}

// The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
func (o EndpointOutput) EndpointStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.EndpointStatus }).(pulumi.StringPtrOutput)
}

// The list of countries/regions mapped to this endpoint when using the 'Geographic' traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
func (o EndpointOutput) GeoMapping() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringArrayOutput { return v.GeoMapping }).(pulumi.StringArrayOutput)
}

// The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
func (o EndpointOutput) MinChildEndpoints() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.Float64PtrOutput { return v.MinChildEndpoints }).(pulumi.Float64PtrOutput)
}

// The minimum number of IPv4 (DNS record type A) endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
func (o EndpointOutput) MinChildEndpointsIPv4() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.Float64PtrOutput { return v.MinChildEndpointsIPv4 }).(pulumi.Float64PtrOutput)
}

// The minimum number of IPv6 (DNS record type AAAA) endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
func (o EndpointOutput) MinChildEndpointsIPv6() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.Float64PtrOutput { return v.MinChildEndpointsIPv6 }).(pulumi.Float64PtrOutput)
}

// The name of the resource
func (o EndpointOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The priority of this endpoint when using the 'Priority' traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
func (o EndpointOutput) Priority() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.Float64PtrOutput { return v.Priority }).(pulumi.Float64PtrOutput)
}

// The list of subnets, IP addresses, and/or address ranges mapped to this endpoint when using the 'Subnet' traffic routing method. An empty list will match all ranges not covered by other endpoints.
func (o EndpointOutput) Subnets() EndpointPropertiesResponseSubnetsArrayOutput {
	return o.ApplyT(func(v *Endpoint) EndpointPropertiesResponseSubnetsArrayOutput { return v.Subnets }).(EndpointPropertiesResponseSubnetsArrayOutput)
}

// The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
func (o EndpointOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.Target }).(pulumi.StringPtrOutput)
}

// The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
func (o EndpointOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

// The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
func (o EndpointOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.
func (o EndpointOutput) Weight() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.Float64PtrOutput { return v.Weight }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointOutput{})
}
