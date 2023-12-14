// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about a gallery image definition.
// Azure REST API version: 2022-03-03.
//
// Other available API versions: 2022-08-03.
func LookupGalleryImage(ctx *pulumi.Context, args *LookupGalleryImageArgs, opts ...pulumi.InvokeOption) (*LookupGalleryImageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGalleryImageResult
	err := ctx.Invoke("azure-native:compute:getGalleryImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryImageArgs struct {
	// The name of the gallery image definition to be retrieved.
	GalleryImageName string `pulumi:"galleryImageName"`
	// The name of the Shared Image Gallery from which the Image Definitions are to be retrieved.
	GalleryName string `pulumi:"galleryName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Specifies information about the gallery image definition that you want to create or update.
type LookupGalleryImageResult struct {
	// The architecture of the image. Applicable to OS disks only.
	Architecture *string `pulumi:"architecture"`
	// The description of this gallery image definition resource. This property is updatable.
	Description *string `pulumi:"description"`
	// Describes the disallowed disk types.
	Disallowed *DisallowedResponse `pulumi:"disallowed"`
	// The end of life date of the gallery image definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate *string `pulumi:"endOfLifeDate"`
	// The Eula agreement for the gallery image definition.
	Eula *string `pulumi:"eula"`
	// A list of gallery image features.
	Features []GalleryImageFeatureResponse `pulumi:"features"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Resource Id
	Id string `pulumi:"id"`
	// This is the gallery image definition identifier.
	Identifier GalleryImageIdentifierResponse `pulumi:"identifier"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
	OsState string `pulumi:"osState"`
	// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	OsType string `pulumi:"osType"`
	// The privacy statement uri.
	PrivacyStatementUri *string `pulumi:"privacyStatementUri"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Describes the gallery image definition purchase plan. This is used by marketplace images.
	PurchasePlan *ImagePurchasePlanResponse `pulumi:"purchasePlan"`
	// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
	Recommended *RecommendedMachineConfigurationResponse `pulumi:"recommended"`
	// The release note uri.
	ReleaseNoteUri *string `pulumi:"releaseNoteUri"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupGalleryImageOutput(ctx *pulumi.Context, args LookupGalleryImageOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGalleryImageResult, error) {
			args := v.(LookupGalleryImageArgs)
			r, err := LookupGalleryImage(ctx, &args, opts...)
			var s LookupGalleryImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGalleryImageResultOutput)
}

type LookupGalleryImageOutputArgs struct {
	// The name of the gallery image definition to be retrieved.
	GalleryImageName pulumi.StringInput `pulumi:"galleryImageName"`
	// The name of the Shared Image Gallery from which the Image Definitions are to be retrieved.
	GalleryName pulumi.StringInput `pulumi:"galleryName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGalleryImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryImageArgs)(nil)).Elem()
}

// Specifies information about the gallery image definition that you want to create or update.
type LookupGalleryImageResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryImageResult)(nil)).Elem()
}

func (o LookupGalleryImageResultOutput) ToLookupGalleryImageResultOutput() LookupGalleryImageResultOutput {
	return o
}

func (o LookupGalleryImageResultOutput) ToLookupGalleryImageResultOutputWithContext(ctx context.Context) LookupGalleryImageResultOutput {
	return o
}

// The architecture of the image. Applicable to OS disks only.
func (o LookupGalleryImageResultOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.Architecture }).(pulumi.StringPtrOutput)
}

// The description of this gallery image definition resource. This property is updatable.
func (o LookupGalleryImageResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Describes the disallowed disk types.
func (o LookupGalleryImageResultOutput) Disallowed() DisallowedResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *DisallowedResponse { return v.Disallowed }).(DisallowedResponsePtrOutput)
}

// The end of life date of the gallery image definition. This property can be used for decommissioning purposes. This property is updatable.
func (o LookupGalleryImageResultOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

// The Eula agreement for the gallery image definition.
func (o LookupGalleryImageResultOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.Eula }).(pulumi.StringPtrOutput)
}

// A list of gallery image features.
func (o LookupGalleryImageResultOutput) Features() GalleryImageFeatureResponseArrayOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) []GalleryImageFeatureResponse { return v.Features }).(GalleryImageFeatureResponseArrayOutput)
}

// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
func (o LookupGalleryImageResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupGalleryImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// This is the gallery image definition identifier.
func (o LookupGalleryImageResultOutput) Identifier() GalleryImageIdentifierResponseOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) GalleryImageIdentifierResponse { return v.Identifier }).(GalleryImageIdentifierResponseOutput)
}

// Resource location
func (o LookupGalleryImageResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupGalleryImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
func (o LookupGalleryImageResultOutput) OsState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.OsState }).(pulumi.StringOutput)
}

// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
func (o LookupGalleryImageResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.OsType }).(pulumi.StringOutput)
}

// The privacy statement uri.
func (o LookupGalleryImageResultOutput) PrivacyStatementUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.PrivacyStatementUri }).(pulumi.StringPtrOutput)
}

// The provisioning state, which only appears in the response.
func (o LookupGalleryImageResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Describes the gallery image definition purchase plan. This is used by marketplace images.
func (o LookupGalleryImageResultOutput) PurchasePlan() ImagePurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *ImagePurchasePlanResponse { return v.PurchasePlan }).(ImagePurchasePlanResponsePtrOutput)
}

// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
func (o LookupGalleryImageResultOutput) Recommended() RecommendedMachineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *RecommendedMachineConfigurationResponse { return v.Recommended }).(RecommendedMachineConfigurationResponsePtrOutput)
}

// The release note uri.
func (o LookupGalleryImageResultOutput) ReleaseNoteUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.ReleaseNoteUri }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o LookupGalleryImageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupGalleryImageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryImageResultOutput{})
}
