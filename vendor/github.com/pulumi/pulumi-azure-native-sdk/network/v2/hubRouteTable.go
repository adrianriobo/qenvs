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

// RouteTable resource in a virtual hub.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01.
type HubRouteTable struct {
	pulumi.CustomResourceState

	// List of all connections associated with this route table.
	AssociatedConnections pulumi.StringArrayOutput `pulumi:"associatedConnections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// List of labels associated with this route table.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// List of all connections that advertise to this route table.
	PropagatingConnections pulumi.StringArrayOutput `pulumi:"propagatingConnections"`
	// The provisioning state of the RouteTable resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// List of all routes.
	Routes HubRouteResponseArrayOutput `pulumi:"routes"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHubRouteTable registers a new resource with the given unique name, arguments, and options.
func NewHubRouteTable(ctx *pulumi.Context,
	name string, args *HubRouteTableArgs, opts ...pulumi.ResourceOption) (*HubRouteTable, error) {
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
			Type: pulumi.String("azure-native:network/v20200401:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:HubRouteTable"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:HubRouteTable"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HubRouteTable
	err := ctx.RegisterResource("azure-native:network:HubRouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHubRouteTable gets an existing HubRouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHubRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubRouteTableState, opts ...pulumi.ResourceOption) (*HubRouteTable, error) {
	var resource HubRouteTable
	err := ctx.ReadResource("azure-native:network:HubRouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HubRouteTable resources.
type hubRouteTableState struct {
}

type HubRouteTableState struct {
}

func (HubRouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*hubRouteTableState)(nil)).Elem()
}

type hubRouteTableArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// List of labels associated with this route table.
	Labels []string `pulumi:"labels"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the RouteTable.
	RouteTableName *string `pulumi:"routeTableName"`
	// List of all routes.
	Routes []HubRoute `pulumi:"routes"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// The set of arguments for constructing a HubRouteTable resource.
type HubRouteTableArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// List of labels associated with this route table.
	Labels pulumi.StringArrayInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput
	// The name of the RouteTable.
	RouteTableName pulumi.StringPtrInput
	// List of all routes.
	Routes HubRouteArrayInput
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput
}

func (HubRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hubRouteTableArgs)(nil)).Elem()
}

type HubRouteTableInput interface {
	pulumi.Input

	ToHubRouteTableOutput() HubRouteTableOutput
	ToHubRouteTableOutputWithContext(ctx context.Context) HubRouteTableOutput
}

func (*HubRouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((**HubRouteTable)(nil)).Elem()
}

func (i *HubRouteTable) ToHubRouteTableOutput() HubRouteTableOutput {
	return i.ToHubRouteTableOutputWithContext(context.Background())
}

func (i *HubRouteTable) ToHubRouteTableOutputWithContext(ctx context.Context) HubRouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubRouteTableOutput)
}

type HubRouteTableOutput struct{ *pulumi.OutputState }

func (HubRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HubRouteTable)(nil)).Elem()
}

func (o HubRouteTableOutput) ToHubRouteTableOutput() HubRouteTableOutput {
	return o
}

func (o HubRouteTableOutput) ToHubRouteTableOutputWithContext(ctx context.Context) HubRouteTableOutput {
	return o
}

// List of all connections associated with this route table.
func (o HubRouteTableOutput) AssociatedConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HubRouteTable) pulumi.StringArrayOutput { return v.AssociatedConnections }).(pulumi.StringArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o HubRouteTableOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *HubRouteTable) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// List of labels associated with this route table.
func (o HubRouteTableOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HubRouteTable) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o HubRouteTableOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HubRouteTable) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// List of all connections that advertise to this route table.
func (o HubRouteTableOutput) PropagatingConnections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HubRouteTable) pulumi.StringArrayOutput { return v.PropagatingConnections }).(pulumi.StringArrayOutput)
}

// The provisioning state of the RouteTable resource.
func (o HubRouteTableOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HubRouteTable) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of all routes.
func (o HubRouteTableOutput) Routes() HubRouteResponseArrayOutput {
	return o.ApplyT(func(v *HubRouteTable) HubRouteResponseArrayOutput { return v.Routes }).(HubRouteResponseArrayOutput)
}

// Resource type.
func (o HubRouteTableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HubRouteTable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HubRouteTableOutput{})
}
