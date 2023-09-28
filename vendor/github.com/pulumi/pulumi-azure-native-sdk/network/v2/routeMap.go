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

// The RouteMap child resource of a Virtual hub.
// Azure REST API version: 2023-02-01.
type RouteMap struct {
	pulumi.CustomResourceState

	// List of connections which have this RoutMap associated for inbound traffic.
	AssociatedInboundConnections pulumi.StringArrayOutput `pulumi:"associatedInboundConnections"`
	// List of connections which have this RoutMap associated for outbound traffic.
	AssociatedOutboundConnections pulumi.StringArrayOutput `pulumi:"associatedOutboundConnections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the RouteMap resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// List of RouteMap rules to be applied.
	Rules RouteMapRuleResponseArrayOutput `pulumi:"rules"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRouteMap registers a new resource with the given unique name, arguments, and options.
func NewRouteMap(ctx *pulumi.Context,
	name string, args *RouteMapArgs, opts ...pulumi.ResourceOption) (*RouteMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHubName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20220501:RouteMap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:RouteMap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:RouteMap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:RouteMap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:RouteMap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:RouteMap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:RouteMap"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RouteMap
	err := ctx.RegisterResource("azure-native:network:RouteMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteMap gets an existing RouteMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteMapState, opts ...pulumi.ResourceOption) (*RouteMap, error) {
	var resource RouteMap
	err := ctx.ReadResource("azure-native:network:RouteMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteMap resources.
type routeMapState struct {
}

type RouteMapState struct {
}

func (RouteMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeMapState)(nil)).Elem()
}

type routeMapArgs struct {
	// List of connections which have this RoutMap associated for inbound traffic.
	AssociatedInboundConnections []string `pulumi:"associatedInboundConnections"`
	// List of connections which have this RoutMap associated for outbound traffic.
	AssociatedOutboundConnections []string `pulumi:"associatedOutboundConnections"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The resource group name of the RouteMap's resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the RouteMap.
	RouteMapName *string `pulumi:"routeMapName"`
	// List of RouteMap rules to be applied.
	Rules []RouteMapRule `pulumi:"rules"`
	// The name of the VirtualHub containing the RouteMap.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// The set of arguments for constructing a RouteMap resource.
type RouteMapArgs struct {
	// List of connections which have this RoutMap associated for inbound traffic.
	AssociatedInboundConnections pulumi.StringArrayInput
	// List of connections which have this RoutMap associated for outbound traffic.
	AssociatedOutboundConnections pulumi.StringArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The resource group name of the RouteMap's resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the RouteMap.
	RouteMapName pulumi.StringPtrInput
	// List of RouteMap rules to be applied.
	Rules RouteMapRuleArrayInput
	// The name of the VirtualHub containing the RouteMap.
	VirtualHubName pulumi.StringInput
}

func (RouteMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeMapArgs)(nil)).Elem()
}

type RouteMapInput interface {
	pulumi.Input

	ToRouteMapOutput() RouteMapOutput
	ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput
}

func (*RouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteMap)(nil)).Elem()
}

func (i *RouteMap) ToRouteMapOutput() RouteMapOutput {
	return i.ToRouteMapOutputWithContext(context.Background())
}

func (i *RouteMap) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapOutput)
}

func (i *RouteMap) ToOutput(ctx context.Context) pulumix.Output[*RouteMap] {
	return pulumix.Output[*RouteMap]{
		OutputState: i.ToRouteMapOutputWithContext(ctx).OutputState,
	}
}

type RouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteMap)(nil)).Elem()
}

func (o RouteMapOutput) ToRouteMapOutput() RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToOutput(ctx context.Context) pulumix.Output[*RouteMap] {
	return pulumix.Output[*RouteMap]{
		OutputState: o.OutputState,
	}
}

// List of connections which have this RoutMap associated for inbound traffic.
func (o RouteMapOutput) AssociatedInboundConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringArrayOutput { return v.AssociatedInboundConnections }).(pulumi.StringArrayOutput)
}

// List of connections which have this RoutMap associated for outbound traffic.
func (o RouteMapOutput) AssociatedOutboundConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringArrayOutput { return v.AssociatedOutboundConnections }).(pulumi.StringArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o RouteMapOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o RouteMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the RouteMap resource.
func (o RouteMapOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of RouteMap rules to be applied.
func (o RouteMapOutput) Rules() RouteMapRuleResponseArrayOutput {
	return o.ApplyT(func(v *RouteMap) RouteMapRuleResponseArrayOutput { return v.Rules }).(RouteMapRuleResponseArrayOutput)
}

// Resource type.
func (o RouteMapOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteMap) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteMapOutput{})
}
