// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Specifies information about the gallery image definition that you want to create or update.
// Azure REST API version: 2022-03-03. Prior API version in Azure Native 1.x: 2020-09-30
type GalleryImage struct {
	pulumi.CustomResourceState

	// The architecture of the image. Applicable to OS disks only.
	Architecture pulumi.StringPtrOutput `pulumi:"architecture"`
	// The description of this gallery image definition resource. This property is updatable.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Describes the disallowed disk types.
	Disallowed DisallowedResponsePtrOutput `pulumi:"disallowed"`
	// The end of life date of the gallery image definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate pulumi.StringPtrOutput `pulumi:"endOfLifeDate"`
	// The Eula agreement for the gallery image definition.
	Eula pulumi.StringPtrOutput `pulumi:"eula"`
	// A list of gallery image features.
	Features GalleryImageFeatureResponseArrayOutput `pulumi:"features"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// This is the gallery image definition identifier.
	Identifier GalleryImageIdentifierResponseOutput `pulumi:"identifier"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
	OsState pulumi.StringOutput `pulumi:"osState"`
	// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	OsType pulumi.StringOutput `pulumi:"osType"`
	// The privacy statement uri.
	PrivacyStatementUri pulumi.StringPtrOutput `pulumi:"privacyStatementUri"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Describes the gallery image definition purchase plan. This is used by marketplace images.
	PurchasePlan ImagePurchasePlanResponsePtrOutput `pulumi:"purchasePlan"`
	// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
	Recommended RecommendedMachineConfigurationResponsePtrOutput `pulumi:"recommended"`
	// The release note uri.
	ReleaseNoteUri pulumi.StringPtrOutput `pulumi:"releaseNoteUri"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGalleryImage registers a new resource with the given unique name, arguments, and options.
func NewGalleryImage(ctx *pulumi.Context,
	name string, args *GalleryImageArgs, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GalleryName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryName'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.OsState == nil {
		return nil, errors.New("invalid value for required argument 'OsState'")
	}
	if args.OsType == nil {
		return nil, errors.New("invalid value for required argument 'OsType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20180601:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211001:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220103:GalleryImage"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220303:GalleryImage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GalleryImage
	err := ctx.RegisterResource("azure-native:compute:GalleryImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGalleryImage gets an existing GalleryImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGalleryImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageState, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	var resource GalleryImage
	err := ctx.ReadResource("azure-native:compute:GalleryImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GalleryImage resources.
type galleryImageState struct {
}

type GalleryImageState struct {
}

func (GalleryImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageState)(nil)).Elem()
}

type galleryImageArgs struct {
	// The architecture of the image. Applicable to OS disks only.
	Architecture *string `pulumi:"architecture"`
	// The description of this gallery image definition resource. This property is updatable.
	Description *string `pulumi:"description"`
	// Describes the disallowed disk types.
	Disallowed *Disallowed `pulumi:"disallowed"`
	// The end of life date of the gallery image definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate *string `pulumi:"endOfLifeDate"`
	// The Eula agreement for the gallery image definition.
	Eula *string `pulumi:"eula"`
	// A list of gallery image features.
	Features []GalleryImageFeature `pulumi:"features"`
	// The name of the gallery image definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
	GalleryImageName *string `pulumi:"galleryImageName"`
	// The name of the Shared Image Gallery in which the Image Definition is to be created.
	GalleryName string `pulumi:"galleryName"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// This is the gallery image definition identifier.
	Identifier GalleryImageIdentifier `pulumi:"identifier"`
	// Resource location
	Location *string `pulumi:"location"`
	// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
	OsState OperatingSystemStateTypes `pulumi:"osState"`
	// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	OsType OperatingSystemTypes `pulumi:"osType"`
	// The privacy statement uri.
	PrivacyStatementUri *string `pulumi:"privacyStatementUri"`
	// Describes the gallery image definition purchase plan. This is used by marketplace images.
	PurchasePlan *ImagePurchasePlan `pulumi:"purchasePlan"`
	// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
	Recommended *RecommendedMachineConfiguration `pulumi:"recommended"`
	// The release note uri.
	ReleaseNoteUri *string `pulumi:"releaseNoteUri"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GalleryImage resource.
type GalleryImageArgs struct {
	// The architecture of the image. Applicable to OS disks only.
	Architecture pulumi.StringPtrInput
	// The description of this gallery image definition resource. This property is updatable.
	Description pulumi.StringPtrInput
	// Describes the disallowed disk types.
	Disallowed DisallowedPtrInput
	// The end of life date of the gallery image definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate pulumi.StringPtrInput
	// The Eula agreement for the gallery image definition.
	Eula pulumi.StringPtrInput
	// A list of gallery image features.
	Features GalleryImageFeatureArrayInput
	// The name of the gallery image definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
	GalleryImageName pulumi.StringPtrInput
	// The name of the Shared Image Gallery in which the Image Definition is to be created.
	GalleryName pulumi.StringInput
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrInput
	// This is the gallery image definition identifier.
	Identifier GalleryImageIdentifierInput
	// Resource location
	Location pulumi.StringPtrInput
	// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
	OsState OperatingSystemStateTypesInput
	// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	OsType OperatingSystemTypesInput
	// The privacy statement uri.
	PrivacyStatementUri pulumi.StringPtrInput
	// Describes the gallery image definition purchase plan. This is used by marketplace images.
	PurchasePlan ImagePurchasePlanPtrInput
	// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
	Recommended RecommendedMachineConfigurationPtrInput
	// The release note uri.
	ReleaseNoteUri pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GalleryImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageArgs)(nil)).Elem()
}

type GalleryImageInput interface {
	pulumi.Input

	ToGalleryImageOutput() GalleryImageOutput
	ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput
}

func (*GalleryImage) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImage)(nil)).Elem()
}

func (i *GalleryImage) ToGalleryImageOutput() GalleryImageOutput {
	return i.ToGalleryImageOutputWithContext(context.Background())
}

func (i *GalleryImage) ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageOutput)
}

func (i *GalleryImage) ToOutput(ctx context.Context) pulumix.Output[*GalleryImage] {
	return pulumix.Output[*GalleryImage]{
		OutputState: i.ToGalleryImageOutputWithContext(ctx).OutputState,
	}
}

type GalleryImageOutput struct{ *pulumi.OutputState }

func (GalleryImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImage)(nil)).Elem()
}

func (o GalleryImageOutput) ToGalleryImageOutput() GalleryImageOutput {
	return o
}

func (o GalleryImageOutput) ToGalleryImageOutputWithContext(ctx context.Context) GalleryImageOutput {
	return o
}

func (o GalleryImageOutput) ToOutput(ctx context.Context) pulumix.Output[*GalleryImage] {
	return pulumix.Output[*GalleryImage]{
		OutputState: o.OutputState,
	}
}

// The architecture of the image. Applicable to OS disks only.
func (o GalleryImageOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.Architecture }).(pulumi.StringPtrOutput)
}

// The description of this gallery image definition resource. This property is updatable.
func (o GalleryImageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Describes the disallowed disk types.
func (o GalleryImageOutput) Disallowed() DisallowedResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImage) DisallowedResponsePtrOutput { return v.Disallowed }).(DisallowedResponsePtrOutput)
}

// The end of life date of the gallery image definition. This property can be used for decommissioning purposes. This property is updatable.
func (o GalleryImageOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

// The Eula agreement for the gallery image definition.
func (o GalleryImageOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.Eula }).(pulumi.StringPtrOutput)
}

// A list of gallery image features.
func (o GalleryImageOutput) Features() GalleryImageFeatureResponseArrayOutput {
	return o.ApplyT(func(v *GalleryImage) GalleryImageFeatureResponseArrayOutput { return v.Features }).(GalleryImageFeatureResponseArrayOutput)
}

// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
func (o GalleryImageOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// This is the gallery image definition identifier.
func (o GalleryImageOutput) Identifier() GalleryImageIdentifierResponseOutput {
	return o.ApplyT(func(v *GalleryImage) GalleryImageIdentifierResponseOutput { return v.Identifier }).(GalleryImageIdentifierResponseOutput)
}

// Resource location
func (o GalleryImageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o GalleryImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
func (o GalleryImageOutput) OsState() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.OsState }).(pulumi.StringOutput)
}

// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
func (o GalleryImageOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

// The privacy statement uri.
func (o GalleryImageOutput) PrivacyStatementUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.PrivacyStatementUri }).(pulumi.StringPtrOutput)
}

// The provisioning state, which only appears in the response.
func (o GalleryImageOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Describes the gallery image definition purchase plan. This is used by marketplace images.
func (o GalleryImageOutput) PurchasePlan() ImagePurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImage) ImagePurchasePlanResponsePtrOutput { return v.PurchasePlan }).(ImagePurchasePlanResponsePtrOutput)
}

// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
func (o GalleryImageOutput) Recommended() RecommendedMachineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GalleryImage) RecommendedMachineConfigurationResponsePtrOutput { return v.Recommended }).(RecommendedMachineConfigurationResponsePtrOutput)
}

// The release note uri.
func (o GalleryImageOutput) ReleaseNoteUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringPtrOutput { return v.ReleaseNoteUri }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o GalleryImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o GalleryImageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GalleryImage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GalleryImageOutput{})
}
