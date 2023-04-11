// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource group information.
// API Version: 2019-05-01.
type ResourceGroup struct {
	pulumi.CustomResourceState

	// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
	Location pulumi.StringOutput `pulumi:"location"`
	// The ID of the resource that manages this resource group.
	ManagedBy pulumi.StringPtrOutput `pulumi:"managedBy"`
	// The name of the resource group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource group properties.
	Properties ResourceGroupPropertiesResponseOutput `pulumi:"properties"`
	// The tags attached to the resource group.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource group.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewResourceGroup(ctx *pulumi.Context,
	name string, args *ResourceGroupArgs, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	if args == nil {
		args = &ResourceGroupArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20151101:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160201:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160701:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20160901:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20170510:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180201:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20180501:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190301:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190501:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190510:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190701:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20190801:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20191001:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:ResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220901:ResourceGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ResourceGroup
	err := ctx.RegisterResource("azure-native:resources:ResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGroup gets an existing ResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupState, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	var resource ResourceGroup
	err := ctx.ReadResource("azure-native:resources:ResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGroup resources.
type resourceGroupState struct {
}

type ResourceGroupState struct {
}

func (ResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupState)(nil)).Elem()
}

type resourceGroupArgs struct {
	// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
	Location *string `pulumi:"location"`
	// The ID of the resource that manages this resource group.
	ManagedBy *string `pulumi:"managedBy"`
	// The name of the resource group to create or update. Can include alphanumeric, underscore, parentheses, hyphen, period (except at end), and Unicode characters that match the allowed characters.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The tags attached to the resource group.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ResourceGroup resource.
type ResourceGroupArgs struct {
	// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
	Location pulumi.StringPtrInput
	// The ID of the resource that manages this resource group.
	ManagedBy pulumi.StringPtrInput
	// The name of the resource group to create or update. Can include alphanumeric, underscore, parentheses, hyphen, period (except at end), and Unicode characters that match the allowed characters.
	ResourceGroupName pulumi.StringPtrInput
	// The tags attached to the resource group.
	Tags pulumi.StringMapInput
}

func (ResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupArgs)(nil)).Elem()
}

type ResourceGroupInput interface {
	pulumi.Input

	ToResourceGroupOutput() ResourceGroupOutput
	ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput
}

func (*ResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroup)(nil)).Elem()
}

func (i *ResourceGroup) ToResourceGroupOutput() ResourceGroupOutput {
	return i.ToResourceGroupOutputWithContext(context.Background())
}

func (i *ResourceGroup) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupOutput)
}

type ResourceGroupOutput struct{ *pulumi.OutputState }

func (ResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGroup)(nil)).Elem()
}

func (o ResourceGroupOutput) ToResourceGroupOutput() ResourceGroupOutput {
	return o
}

func (o ResourceGroupOutput) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return o
}

// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
func (o ResourceGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The ID of the resource that manages this resource group.
func (o ResourceGroupOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The name of the resource group.
func (o ResourceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource group properties.
func (o ResourceGroupOutput) Properties() ResourceGroupPropertiesResponseOutput {
	return o.ApplyT(func(v *ResourceGroup) ResourceGroupPropertiesResponseOutput { return v.Properties }).(ResourceGroupPropertiesResponseOutput)
}

// The tags attached to the resource group.
func (o ResourceGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource group.
func (o ResourceGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceGroupOutput{})
}
