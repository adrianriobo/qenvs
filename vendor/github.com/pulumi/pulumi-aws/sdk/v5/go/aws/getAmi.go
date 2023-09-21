// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of a registered AMI for use in other
// resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
//				ExecutableUsers: []string{
//					"self",
//				},
//				Filters: []ec2.GetAmiFilter{
//					{
//						Name: "name",
//						Values: []string{
//							"myami-*",
//						},
//					},
//					{
//						Name: "root-device-type",
//						Values: []string{
//							"ebs",
//						},
//					},
//					{
//						Name: "virtualization-type",
//						Values: []string{
//							"hvm",
//						},
//					},
//				},
//				MostRecent: pulumi.BoolRef(true),
//				NameRegex:  pulumi.StringRef("^myami-\\d{3}"),
//				Owners: []string{
//					"self",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: aws.getAmi has been deprecated in favor of aws.ec2.getAmi
func GetAmi(ctx *pulumi.Context, args *GetAmiArgs, opts ...pulumi.InvokeOption) (*GetAmiResult, error) {
	var rv GetAmiResult
	err := ctx.Invoke("aws:index/getAmi:getAmi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAmi.
type GetAmiArgs struct {
	// Limit search to users with *explicit* launch permission on
	// the image. Valid items are the numeric account ID or `self`.
	ExecutableUsers []string `pulumi:"executableUsers"`
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-images in the AWS CLI reference][1].
	Filters []GetAmiFilter `pulumi:"filters"`
	// If true, all deprecated AMIs are included in the response. If false, no deprecated AMIs are included in the response. If no value is specified, the default value is false.
	IncludeDeprecated *bool `pulumi:"includeDeprecated"`
	// If more than one result is returned, use the most
	// recent AMI.
	MostRecent *bool `pulumi:"mostRecent"`
	// Regex string to apply to the AMI list returned
	// by AWS. This allows more advanced filtering not supported from the AWS API. This
	// filtering is done locally on what AWS returns, and could have a performance
	// impact if the result is large. Combine this with other
	// options to narrow down the list AWS returns.
	//
	// > **NOTE:** If more or less than a single match is returned by the search,
	// this call will fail. Ensure that your search is specific enough to return
	// a single AMI ID only, or use `mostRecent` to choose the most recent one. If
	// you want to match multiple AMIs, use the `ec2.getAmiIds` data source instead.
	NameRegex *string `pulumi:"nameRegex"`
	// List of AMI owners to limit search. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g., `amazon`, `aws-marketplace`, `microsoft`).
	Owners []string `pulumi:"owners"`
	// Any tags assigned to the image.
	// * `tags.#.key` - Key name of the tag.
	// * `tags.#.value` - Value of the tag.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getAmi.
type GetAmiResult struct {
	// OS architecture of the AMI (ie: `i386` or `x8664`).
	Architecture string `pulumi:"architecture"`
	// ARN of the AMI.
	Arn string `pulumi:"arn"`
	// Set of objects with block device mappings of the AMI.
	BlockDeviceMappings []GetAmiBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// Boot mode of the image.
	BootMode string `pulumi:"bootMode"`
	// Date and time the image was created.
	CreationDate string `pulumi:"creationDate"`
	// Date and time when the image will be deprecated.
	DeprecationTime string `pulumi:"deprecationTime"`
	// Description of the AMI that was provided during image
	// creation.
	Description string `pulumi:"description"`
	// Whether enhanced networking with ENA is enabled.
	EnaSupport      bool           `pulumi:"enaSupport"`
	ExecutableUsers []string       `pulumi:"executableUsers"`
	Filters         []GetAmiFilter `pulumi:"filters"`
	// Hypervisor type of the image.
	Hypervisor string `pulumi:"hypervisor"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the AMI. Should be the same as the resource `id`.
	ImageId string `pulumi:"imageId"`
	// Location of the AMI.
	ImageLocation string `pulumi:"imageLocation"`
	// AWS account alias (for example, `amazon`, `self`) or
	// the AWS account ID of the AMI owner.
	ImageOwnerAlias string `pulumi:"imageOwnerAlias"`
	// Type of image.
	ImageType string `pulumi:"imageType"`
	// Instance Metadata Service (IMDS) support mode for the image. Set to `v2.0` if instances ran from this image enforce IMDSv2.
	ImdsSupport       string `pulumi:"imdsSupport"`
	IncludeDeprecated *bool  `pulumi:"includeDeprecated"`
	// Kernel associated with the image, if any. Only applicable
	// for machine images.
	KernelId   string `pulumi:"kernelId"`
	MostRecent *bool  `pulumi:"mostRecent"`
	// Name of the AMI that was provided during image creation.
	Name      string  `pulumi:"name"`
	NameRegex *string `pulumi:"nameRegex"`
	// AWS account ID of the image owner.
	OwnerId string   `pulumi:"ownerId"`
	Owners  []string `pulumi:"owners"`
	// Value is Windows for `Windows` AMIs; otherwise blank.
	Platform string `pulumi:"platform"`
	// Platform details associated with the billing code of the AMI.
	PlatformDetails string `pulumi:"platformDetails"`
	// Any product codes associated with the AMI.
	// * `product_codes.#.product_code_id` - The product code.
	// * `product_codes.#.product_code_type` - The type of product code.
	ProductCodes []GetAmiProductCode `pulumi:"productCodes"`
	// `true` if the image has public launch permissions.
	Public bool `pulumi:"public"`
	// RAM disk associated with the image, if any. Only applicable
	// for machine images.
	RamdiskId string `pulumi:"ramdiskId"`
	// Device name of the root device.
	RootDeviceName string `pulumi:"rootDeviceName"`
	// Type of root device (ie: `ebs` or `instance-store`).
	RootDeviceType string `pulumi:"rootDeviceType"`
	// Snapshot id associated with the root device, if any
	// (only applies to `ebs` root devices).
	RootSnapshotId string `pulumi:"rootSnapshotId"`
	// Whether enhanced networking is enabled.
	SriovNetSupport string `pulumi:"sriovNetSupport"`
	// Current state of the AMI. If the state is `available`, the image
	// is successfully registered and can be used to launch an instance.
	State string `pulumi:"state"`
	// Describes a state change. Fields are `UNSET` if not available.
	// * `state_reason.code` - The reason code for the state change.
	// * `state_reason.message` - The message for the state change.
	StateReason map[string]string `pulumi:"stateReason"`
	// Any tags assigned to the image.
	// * `tags.#.key` - Key name of the tag.
	// * `tags.#.value` - Value of the tag.
	Tags map[string]string `pulumi:"tags"`
	// If the image is configured for NitroTPM support, the value is `v2.0`.
	TpmSupport string `pulumi:"tpmSupport"`
	// Operation of the Amazon EC2 instance and the billing code that is associated with the AMI.
	UsageOperation string `pulumi:"usageOperation"`
	// Type of virtualization of the AMI (ie: `hvm` or
	// `paravirtual`).
	VirtualizationType string `pulumi:"virtualizationType"`
}

func GetAmiOutput(ctx *pulumi.Context, args GetAmiOutputArgs, opts ...pulumi.InvokeOption) GetAmiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAmiResult, error) {
			args := v.(GetAmiArgs)
			r, err := GetAmi(ctx, &args, opts...)
			var s GetAmiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAmiResultOutput)
}

// A collection of arguments for invoking getAmi.
type GetAmiOutputArgs struct {
	// Limit search to users with *explicit* launch permission on
	// the image. Valid items are the numeric account ID or `self`.
	ExecutableUsers pulumi.StringArrayInput `pulumi:"executableUsers"`
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-images in the AWS CLI reference][1].
	Filters GetAmiFilterArrayInput `pulumi:"filters"`
	// If true, all deprecated AMIs are included in the response. If false, no deprecated AMIs are included in the response. If no value is specified, the default value is false.
	IncludeDeprecated pulumi.BoolPtrInput `pulumi:"includeDeprecated"`
	// If more than one result is returned, use the most
	// recent AMI.
	MostRecent pulumi.BoolPtrInput `pulumi:"mostRecent"`
	// Regex string to apply to the AMI list returned
	// by AWS. This allows more advanced filtering not supported from the AWS API. This
	// filtering is done locally on what AWS returns, and could have a performance
	// impact if the result is large. Combine this with other
	// options to narrow down the list AWS returns.
	//
	// > **NOTE:** If more or less than a single match is returned by the search,
	// this call will fail. Ensure that your search is specific enough to return
	// a single AMI ID only, or use `mostRecent` to choose the most recent one. If
	// you want to match multiple AMIs, use the `ec2.getAmiIds` data source instead.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// List of AMI owners to limit search. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g., `amazon`, `aws-marketplace`, `microsoft`).
	Owners pulumi.StringArrayInput `pulumi:"owners"`
	// Any tags assigned to the image.
	// * `tags.#.key` - Key name of the tag.
	// * `tags.#.value` - Value of the tag.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetAmiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiArgs)(nil)).Elem()
}

// A collection of values returned by getAmi.
type GetAmiResultOutput struct{ *pulumi.OutputState }

func (GetAmiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAmiResult)(nil)).Elem()
}

func (o GetAmiResultOutput) ToGetAmiResultOutput() GetAmiResultOutput {
	return o
}

func (o GetAmiResultOutput) ToGetAmiResultOutputWithContext(ctx context.Context) GetAmiResultOutput {
	return o
}

// OS architecture of the AMI (ie: `i386` or `x8664`).
func (o GetAmiResultOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.Architecture }).(pulumi.StringOutput)
}

// ARN of the AMI.
func (o GetAmiResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Set of objects with block device mappings of the AMI.
func (o GetAmiResultOutput) BlockDeviceMappings() GetAmiBlockDeviceMappingArrayOutput {
	return o.ApplyT(func(v GetAmiResult) []GetAmiBlockDeviceMapping { return v.BlockDeviceMappings }).(GetAmiBlockDeviceMappingArrayOutput)
}

// Boot mode of the image.
func (o GetAmiResultOutput) BootMode() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.BootMode }).(pulumi.StringOutput)
}

// Date and time the image was created.
func (o GetAmiResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Date and time when the image will be deprecated.
func (o GetAmiResultOutput) DeprecationTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.DeprecationTime }).(pulumi.StringOutput)
}

// Description of the AMI that was provided during image
// creation.
func (o GetAmiResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.Description }).(pulumi.StringOutput)
}

// Whether enhanced networking with ENA is enabled.
func (o GetAmiResultOutput) EnaSupport() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAmiResult) bool { return v.EnaSupport }).(pulumi.BoolOutput)
}

func (o GetAmiResultOutput) ExecutableUsers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAmiResult) []string { return v.ExecutableUsers }).(pulumi.StringArrayOutput)
}

func (o GetAmiResultOutput) Filters() GetAmiFilterArrayOutput {
	return o.ApplyT(func(v GetAmiResult) []GetAmiFilter { return v.Filters }).(GetAmiFilterArrayOutput)
}

// Hypervisor type of the image.
func (o GetAmiResultOutput) Hypervisor() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.Hypervisor }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAmiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the AMI. Should be the same as the resource `id`.
func (o GetAmiResultOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.ImageId }).(pulumi.StringOutput)
}

// Location of the AMI.
func (o GetAmiResultOutput) ImageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.ImageLocation }).(pulumi.StringOutput)
}

// AWS account alias (for example, `amazon`, `self`) or
// the AWS account ID of the AMI owner.
func (o GetAmiResultOutput) ImageOwnerAlias() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.ImageOwnerAlias }).(pulumi.StringOutput)
}

// Type of image.
func (o GetAmiResultOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.ImageType }).(pulumi.StringOutput)
}

// Instance Metadata Service (IMDS) support mode for the image. Set to `v2.0` if instances ran from this image enforce IMDSv2.
func (o GetAmiResultOutput) ImdsSupport() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.ImdsSupport }).(pulumi.StringOutput)
}

func (o GetAmiResultOutput) IncludeDeprecated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAmiResult) *bool { return v.IncludeDeprecated }).(pulumi.BoolPtrOutput)
}

// Kernel associated with the image, if any. Only applicable
// for machine images.
func (o GetAmiResultOutput) KernelId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.KernelId }).(pulumi.StringOutput)
}

func (o GetAmiResultOutput) MostRecent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAmiResult) *bool { return v.MostRecent }).(pulumi.BoolPtrOutput)
}

// Name of the AMI that was provided during image creation.
func (o GetAmiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetAmiResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAmiResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// AWS account ID of the image owner.
func (o GetAmiResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o GetAmiResultOutput) Owners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAmiResult) []string { return v.Owners }).(pulumi.StringArrayOutput)
}

// Value is Windows for `Windows` AMIs; otherwise blank.
func (o GetAmiResultOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.Platform }).(pulumi.StringOutput)
}

// Platform details associated with the billing code of the AMI.
func (o GetAmiResultOutput) PlatformDetails() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.PlatformDetails }).(pulumi.StringOutput)
}

// Any product codes associated with the AMI.
// * `product_codes.#.product_code_id` - The product code.
// * `product_codes.#.product_code_type` - The type of product code.
func (o GetAmiResultOutput) ProductCodes() GetAmiProductCodeArrayOutput {
	return o.ApplyT(func(v GetAmiResult) []GetAmiProductCode { return v.ProductCodes }).(GetAmiProductCodeArrayOutput)
}

// `true` if the image has public launch permissions.
func (o GetAmiResultOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v GetAmiResult) bool { return v.Public }).(pulumi.BoolOutput)
}

// RAM disk associated with the image, if any. Only applicable
// for machine images.
func (o GetAmiResultOutput) RamdiskId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.RamdiskId }).(pulumi.StringOutput)
}

// Device name of the root device.
func (o GetAmiResultOutput) RootDeviceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.RootDeviceName }).(pulumi.StringOutput)
}

// Type of root device (ie: `ebs` or `instance-store`).
func (o GetAmiResultOutput) RootDeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.RootDeviceType }).(pulumi.StringOutput)
}

// Snapshot id associated with the root device, if any
// (only applies to `ebs` root devices).
func (o GetAmiResultOutput) RootSnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.RootSnapshotId }).(pulumi.StringOutput)
}

// Whether enhanced networking is enabled.
func (o GetAmiResultOutput) SriovNetSupport() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.SriovNetSupport }).(pulumi.StringOutput)
}

// Current state of the AMI. If the state is `available`, the image
// is successfully registered and can be used to launch an instance.
func (o GetAmiResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.State }).(pulumi.StringOutput)
}

// Describes a state change. Fields are `UNSET` if not available.
// * `state_reason.code` - The reason code for the state change.
// * `state_reason.message` - The message for the state change.
func (o GetAmiResultOutput) StateReason() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetAmiResult) map[string]string { return v.StateReason }).(pulumi.StringMapOutput)
}

// Any tags assigned to the image.
// * `tags.#.key` - Key name of the tag.
// * `tags.#.value` - Value of the tag.
func (o GetAmiResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetAmiResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// If the image is configured for NitroTPM support, the value is `v2.0`.
func (o GetAmiResultOutput) TpmSupport() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.TpmSupport }).(pulumi.StringOutput)
}

// Operation of the Amazon EC2 instance and the billing code that is associated with the AMI.
func (o GetAmiResultOutput) UsageOperation() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.UsageOperation }).(pulumi.StringOutput)
}

// Type of virtualization of the AMI (ie: `hvm` or
// `paravirtual`).
func (o GetAmiResultOutput) VirtualizationType() pulumi.StringOutput {
	return o.ApplyT(func(v GetAmiResult) string { return v.VirtualizationType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAmiResultOutput{})
}
