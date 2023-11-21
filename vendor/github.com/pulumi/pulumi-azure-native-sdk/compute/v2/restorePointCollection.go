// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create or update Restore Point collection parameters.
// Azure REST API version: 2023-03-01. Prior API version in Azure Native 1.x: 2021-03-01.
//
// Other available API versions: 2023-07-01.
type RestorePointCollection struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the restore point collection.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The unique id of the restore point collection.
	RestorePointCollectionId pulumi.StringOutput `pulumi:"restorePointCollectionId"`
	// A list containing all restore points created under this restore point collection.
	RestorePoints RestorePointResponseArrayOutput `pulumi:"restorePoints"`
	// The properties of the source resource that this restore point collection is created from.
	Source RestorePointCollectionSourcePropertiesResponsePtrOutput `pulumi:"source"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRestorePointCollection registers a new resource with the given unique name, arguments, and options.
func NewRestorePointCollection(ctx *pulumi.Context,
	name string, args *RestorePointCollectionArgs, opts ...pulumi.ResourceOption) (*RestorePointCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20210301:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:RestorePointCollection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230701:RestorePointCollection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RestorePointCollection
	err := ctx.RegisterResource("azure-native:compute:RestorePointCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestorePointCollection gets an existing RestorePointCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestorePointCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestorePointCollectionState, opts ...pulumi.ResourceOption) (*RestorePointCollection, error) {
	var resource RestorePointCollection
	err := ctx.ReadResource("azure-native:compute:RestorePointCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RestorePointCollection resources.
type restorePointCollectionState struct {
}

type RestorePointCollectionState struct {
}

func (RestorePointCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*restorePointCollectionState)(nil)).Elem()
}

type restorePointCollectionArgs struct {
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the restore point collection.
	RestorePointCollectionName *string `pulumi:"restorePointCollectionName"`
	// The properties of the source resource that this restore point collection is created from.
	Source *RestorePointCollectionSourceProperties `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RestorePointCollection resource.
type RestorePointCollectionArgs struct {
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the restore point collection.
	RestorePointCollectionName pulumi.StringPtrInput
	// The properties of the source resource that this restore point collection is created from.
	Source RestorePointCollectionSourcePropertiesPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (RestorePointCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restorePointCollectionArgs)(nil)).Elem()
}

type RestorePointCollectionInput interface {
	pulumi.Input

	ToRestorePointCollectionOutput() RestorePointCollectionOutput
	ToRestorePointCollectionOutputWithContext(ctx context.Context) RestorePointCollectionOutput
}

func (*RestorePointCollection) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePointCollection)(nil)).Elem()
}

func (i *RestorePointCollection) ToRestorePointCollectionOutput() RestorePointCollectionOutput {
	return i.ToRestorePointCollectionOutputWithContext(context.Background())
}

func (i *RestorePointCollection) ToRestorePointCollectionOutputWithContext(ctx context.Context) RestorePointCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestorePointCollectionOutput)
}

type RestorePointCollectionOutput struct{ *pulumi.OutputState }

func (RestorePointCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestorePointCollection)(nil)).Elem()
}

func (o RestorePointCollectionOutput) ToRestorePointCollectionOutput() RestorePointCollectionOutput {
	return o
}

func (o RestorePointCollectionOutput) ToRestorePointCollectionOutputWithContext(ctx context.Context) RestorePointCollectionOutput {
	return o
}

// Resource location
func (o RestorePointCollectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o RestorePointCollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the restore point collection.
func (o RestorePointCollectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The unique id of the restore point collection.
func (o RestorePointCollectionOutput) RestorePointCollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringOutput { return v.RestorePointCollectionId }).(pulumi.StringOutput)
}

// A list containing all restore points created under this restore point collection.
func (o RestorePointCollectionOutput) RestorePoints() RestorePointResponseArrayOutput {
	return o.ApplyT(func(v *RestorePointCollection) RestorePointResponseArrayOutput { return v.RestorePoints }).(RestorePointResponseArrayOutput)
}

// The properties of the source resource that this restore point collection is created from.
func (o RestorePointCollectionOutput) Source() RestorePointCollectionSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *RestorePointCollection) RestorePointCollectionSourcePropertiesResponsePtrOutput {
		return v.Source
	}).(RestorePointCollectionSourcePropertiesResponsePtrOutput)
}

// Resource tags
func (o RestorePointCollectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o RestorePointCollectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RestorePointCollection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RestorePointCollectionOutput{})
}
