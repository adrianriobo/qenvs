// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IAM instance profile.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"ec2.amazonaws.com",
//								},
//							},
//						},
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
//				Path:             pulumi.String("/"),
//				AssumeRolePolicy: *pulumi.String(assumeRole.Json),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewInstanceProfile(ctx, "testProfile", &iam.InstanceProfileArgs{
//				Role: role.Name,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import Instance Profiles using the `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:iam/instanceProfile:InstanceProfile test_profile app-instance-profile-1
//
// ```
type InstanceProfile struct {
	pulumi.CustomResourceState

	// ARN assigned by AWS to the instance profile.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Creation timestamp of the instance profile.
	CreateDate pulumi.StringOutput `pulumi:"createDate"`
	// Name of the instance profile. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`. Can be a string of characters consisting of upper and lowercase alphanumeric characters and these special characters: `_`, `+`, `=`, `,`, `.`, `@`, `-`. Spaces are not allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// Path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. Can be a string of characters consisting of either a forward slash (`/`) by itself or a string that must begin and end with forward slashes. Can include any ASCII character from the ! (\u0021) through the DEL character (\u007F), including most punctuation characters, digits, and upper and lowercase letters.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Name of the role to add to the profile.
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// Map of resource tags for the IAM Instance Profile. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// [Unique ID][1] assigned by AWS.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewInstanceProfile registers a new resource with the given unique name, arguments, and options.
func NewInstanceProfile(ctx *pulumi.Context,
	name string, args *InstanceProfileArgs, opts ...pulumi.ResourceOption) (*InstanceProfile, error) {
	if args == nil {
		args = &InstanceProfileArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceProfile
	err := ctx.RegisterResource("aws:iam/instanceProfile:InstanceProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceProfile gets an existing InstanceProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceProfileState, opts ...pulumi.ResourceOption) (*InstanceProfile, error) {
	var resource InstanceProfile
	err := ctx.ReadResource("aws:iam/instanceProfile:InstanceProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceProfile resources.
type instanceProfileState struct {
	// ARN assigned by AWS to the instance profile.
	Arn *string `pulumi:"arn"`
	// Creation timestamp of the instance profile.
	CreateDate *string `pulumi:"createDate"`
	// Name of the instance profile. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`. Can be a string of characters consisting of upper and lowercase alphanumeric characters and these special characters: `_`, `+`, `=`, `,`, `.`, `@`, `-`. Spaces are not allowed.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. Can be a string of characters consisting of either a forward slash (`/`) by itself or a string that must begin and end with forward slashes. Can include any ASCII character from the ! (\u0021) through the DEL character (\u007F), including most punctuation characters, digits, and upper and lowercase letters.
	Path *string `pulumi:"path"`
	// Name of the role to add to the profile.
	Role interface{} `pulumi:"role"`
	// Map of resource tags for the IAM Instance Profile. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// [Unique ID][1] assigned by AWS.
	UniqueId *string `pulumi:"uniqueId"`
}

type InstanceProfileState struct {
	// ARN assigned by AWS to the instance profile.
	Arn pulumi.StringPtrInput
	// Creation timestamp of the instance profile.
	CreateDate pulumi.StringPtrInput
	// Name of the instance profile. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`. Can be a string of characters consisting of upper and lowercase alphanumeric characters and these special characters: `_`, `+`, `=`, `,`, `.`, `@`, `-`. Spaces are not allowed.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. Can be a string of characters consisting of either a forward slash (`/`) by itself or a string that must begin and end with forward slashes. Can include any ASCII character from the ! (\u0021) through the DEL character (\u007F), including most punctuation characters, digits, and upper and lowercase letters.
	Path pulumi.StringPtrInput
	// Name of the role to add to the profile.
	Role pulumi.Input
	// Map of resource tags for the IAM Instance Profile. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// [Unique ID][1] assigned by AWS.
	UniqueId pulumi.StringPtrInput
}

func (InstanceProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceProfileState)(nil)).Elem()
}

type instanceProfileArgs struct {
	// Name of the instance profile. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`. Can be a string of characters consisting of upper and lowercase alphanumeric characters and these special characters: `_`, `+`, `=`, `,`, `.`, `@`, `-`. Spaces are not allowed.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. Can be a string of characters consisting of either a forward slash (`/`) by itself or a string that must begin and end with forward slashes. Can include any ASCII character from the ! (\u0021) through the DEL character (\u007F), including most punctuation characters, digits, and upper and lowercase letters.
	Path *string `pulumi:"path"`
	// Name of the role to add to the profile.
	Role interface{} `pulumi:"role"`
	// Map of resource tags for the IAM Instance Profile. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a InstanceProfile resource.
type InstanceProfileArgs struct {
	// Name of the instance profile. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`. Can be a string of characters consisting of upper and lowercase alphanumeric characters and these special characters: `_`, `+`, `=`, `,`, `.`, `@`, `-`. Spaces are not allowed.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. Can be a string of characters consisting of either a forward slash (`/`) by itself or a string that must begin and end with forward slashes. Can include any ASCII character from the ! (\u0021) through the DEL character (\u007F), including most punctuation characters, digits, and upper and lowercase letters.
	Path pulumi.StringPtrInput
	// Name of the role to add to the profile.
	Role pulumi.Input
	// Map of resource tags for the IAM Instance Profile. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (InstanceProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceProfileArgs)(nil)).Elem()
}

type InstanceProfileInput interface {
	pulumi.Input

	ToInstanceProfileOutput() InstanceProfileOutput
	ToInstanceProfileOutputWithContext(ctx context.Context) InstanceProfileOutput
}

func (*InstanceProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceProfile)(nil)).Elem()
}

func (i *InstanceProfile) ToInstanceProfileOutput() InstanceProfileOutput {
	return i.ToInstanceProfileOutputWithContext(context.Background())
}

func (i *InstanceProfile) ToInstanceProfileOutputWithContext(ctx context.Context) InstanceProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceProfileOutput)
}

// InstanceProfileArrayInput is an input type that accepts InstanceProfileArray and InstanceProfileArrayOutput values.
// You can construct a concrete instance of `InstanceProfileArrayInput` via:
//
//	InstanceProfileArray{ InstanceProfileArgs{...} }
type InstanceProfileArrayInput interface {
	pulumi.Input

	ToInstanceProfileArrayOutput() InstanceProfileArrayOutput
	ToInstanceProfileArrayOutputWithContext(context.Context) InstanceProfileArrayOutput
}

type InstanceProfileArray []InstanceProfileInput

func (InstanceProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceProfile)(nil)).Elem()
}

func (i InstanceProfileArray) ToInstanceProfileArrayOutput() InstanceProfileArrayOutput {
	return i.ToInstanceProfileArrayOutputWithContext(context.Background())
}

func (i InstanceProfileArray) ToInstanceProfileArrayOutputWithContext(ctx context.Context) InstanceProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceProfileArrayOutput)
}

// InstanceProfileMapInput is an input type that accepts InstanceProfileMap and InstanceProfileMapOutput values.
// You can construct a concrete instance of `InstanceProfileMapInput` via:
//
//	InstanceProfileMap{ "key": InstanceProfileArgs{...} }
type InstanceProfileMapInput interface {
	pulumi.Input

	ToInstanceProfileMapOutput() InstanceProfileMapOutput
	ToInstanceProfileMapOutputWithContext(context.Context) InstanceProfileMapOutput
}

type InstanceProfileMap map[string]InstanceProfileInput

func (InstanceProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceProfile)(nil)).Elem()
}

func (i InstanceProfileMap) ToInstanceProfileMapOutput() InstanceProfileMapOutput {
	return i.ToInstanceProfileMapOutputWithContext(context.Background())
}

func (i InstanceProfileMap) ToInstanceProfileMapOutputWithContext(ctx context.Context) InstanceProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceProfileMapOutput)
}

type InstanceProfileOutput struct{ *pulumi.OutputState }

func (InstanceProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceProfile)(nil)).Elem()
}

func (o InstanceProfileOutput) ToInstanceProfileOutput() InstanceProfileOutput {
	return o
}

func (o InstanceProfileOutput) ToInstanceProfileOutputWithContext(ctx context.Context) InstanceProfileOutput {
	return o
}

// ARN assigned by AWS to the instance profile.
func (o InstanceProfileOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Creation timestamp of the instance profile.
func (o InstanceProfileOutput) CreateDate() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringOutput { return v.CreateDate }).(pulumi.StringOutput)
}

// Name of the instance profile. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`. Can be a string of characters consisting of upper and lowercase alphanumeric characters and these special characters: `_`, `+`, `=`, `,`, `.`, `@`, `-`. Spaces are not allowed.
func (o InstanceProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o InstanceProfileOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// Path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. Can be a string of characters consisting of either a forward slash (`/`) by itself or a string that must begin and end with forward slashes. Can include any ASCII character from the ! (\u0021) through the DEL character (\u007F), including most punctuation characters, digits, and upper and lowercase letters.
func (o InstanceProfileOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// Name of the role to add to the profile.
func (o InstanceProfileOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

// Map of resource tags for the IAM Instance Profile. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o InstanceProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o InstanceProfileOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// [Unique ID][1] assigned by AWS.
func (o InstanceProfileOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceProfile) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

type InstanceProfileArrayOutput struct{ *pulumi.OutputState }

func (InstanceProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceProfile)(nil)).Elem()
}

func (o InstanceProfileArrayOutput) ToInstanceProfileArrayOutput() InstanceProfileArrayOutput {
	return o
}

func (o InstanceProfileArrayOutput) ToInstanceProfileArrayOutputWithContext(ctx context.Context) InstanceProfileArrayOutput {
	return o
}

func (o InstanceProfileArrayOutput) Index(i pulumi.IntInput) InstanceProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceProfile {
		return vs[0].([]*InstanceProfile)[vs[1].(int)]
	}).(InstanceProfileOutput)
}

type InstanceProfileMapOutput struct{ *pulumi.OutputState }

func (InstanceProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceProfile)(nil)).Elem()
}

func (o InstanceProfileMapOutput) ToInstanceProfileMapOutput() InstanceProfileMapOutput {
	return o
}

func (o InstanceProfileMapOutput) ToInstanceProfileMapOutputWithContext(ctx context.Context) InstanceProfileMapOutput {
	return o
}

func (o InstanceProfileMapOutput) MapIndex(k pulumi.StringInput) InstanceProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceProfile {
		return vs[0].(map[string]*InstanceProfile)[vs[1].(string)]
	}).(InstanceProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceProfileInput)(nil)).Elem(), &InstanceProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceProfileArrayInput)(nil)).Elem(), InstanceProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceProfileMapInput)(nil)).Elem(), InstanceProfileMap{})
	pulumi.RegisterOutputType(InstanceProfileOutput{})
	pulumi.RegisterOutputType(InstanceProfileArrayOutput{})
	pulumi.RegisterOutputType(InstanceProfileMapOutput{})
}
