// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a managed prefix list resource.
//
// > **NOTE on Managed Prefix Lists and Managed Prefix List Entries:** The provider
// currently provides both a standalone Managed Prefix List Entry resource (a single entry),
// and a Managed Prefix List resource with entries defined in-line. At this time you
// cannot use a Managed Prefix List with in-line rules in conjunction with any Managed
// Prefix List Entry resources. Doing so will cause a conflict of entries and will overwrite entries.
//
// > **NOTE on `maxEntries`:** When you reference a Prefix List in a resource,
// the maximum number of entries for the prefix lists counts as the same number of rules
// or entries for the resource. For example, if you create a prefix list with a maximum
// of 20 entries and you reference that prefix list in a security group rule, this counts
// as 20 rules for the security group.
//
// ## Example Usage
//
// # Basic usage
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
//			_, err := ec2.NewManagedPrefixList(ctx, "example", &ec2.ManagedPrefixListArgs{
//				AddressFamily: pulumi.String("IPv4"),
//				MaxEntries:    pulumi.Int(5),
//				Entries: ec2.ManagedPrefixListEntryTypeArray{
//					&ec2.ManagedPrefixListEntryTypeArgs{
//						Cidr:        pulumi.Any(aws_vpc.Example.Cidr_block),
//						Description: pulumi.String("Primary"),
//					},
//					&ec2.ManagedPrefixListEntryTypeArgs{
//						Cidr:        pulumi.Any(aws_vpc_ipv4_cidr_block_association.Example.Cidr_block),
//						Description: pulumi.String("Secondary"),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"Env": pulumi.String("live"),
//				},
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
// Prefix Lists can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/managedPrefixList:ManagedPrefixList default pl-0570a1d2d725c16be
//
// ```
type ManagedPrefixList struct {
	pulumi.CustomResourceState

	// Address family (`IPv4` or `IPv6`) of this prefix list.
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// ARN of the prefix list.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries ManagedPrefixListEntryTypeArrayOutput `pulumi:"entries"`
	// Maximum number of entries that this prefix list can contain.
	MaxEntries pulumi.IntOutput `pulumi:"maxEntries"`
	// Name of this resource. The name must not start with `com.amazonaws`.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the AWS account that owns this prefix list.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Latest version of this prefix list.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewManagedPrefixList registers a new resource with the given unique name, arguments, and options.
func NewManagedPrefixList(ctx *pulumi.Context,
	name string, args *ManagedPrefixListArgs, opts ...pulumi.ResourceOption) (*ManagedPrefixList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'AddressFamily'")
	}
	if args.MaxEntries == nil {
		return nil, errors.New("invalid value for required argument 'MaxEntries'")
	}
	var resource ManagedPrefixList
	err := ctx.RegisterResource("aws:ec2/managedPrefixList:ManagedPrefixList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPrefixList gets an existing ManagedPrefixList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPrefixList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPrefixListState, opts ...pulumi.ResourceOption) (*ManagedPrefixList, error) {
	var resource ManagedPrefixList
	err := ctx.ReadResource("aws:ec2/managedPrefixList:ManagedPrefixList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPrefixList resources.
type managedPrefixListState struct {
	// Address family (`IPv4` or `IPv6`) of this prefix list.
	AddressFamily *string `pulumi:"addressFamily"`
	// ARN of the prefix list.
	Arn *string `pulumi:"arn"`
	// Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries []ManagedPrefixListEntryType `pulumi:"entries"`
	// Maximum number of entries that this prefix list can contain.
	MaxEntries *int `pulumi:"maxEntries"`
	// Name of this resource. The name must not start with `com.amazonaws`.
	Name *string `pulumi:"name"`
	// ID of the AWS account that owns this prefix list.
	OwnerId *string `pulumi:"ownerId"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Latest version of this prefix list.
	Version *int `pulumi:"version"`
}

type ManagedPrefixListState struct {
	// Address family (`IPv4` or `IPv6`) of this prefix list.
	AddressFamily pulumi.StringPtrInput
	// ARN of the prefix list.
	Arn pulumi.StringPtrInput
	// Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries ManagedPrefixListEntryTypeArrayInput
	// Maximum number of entries that this prefix list can contain.
	MaxEntries pulumi.IntPtrInput
	// Name of this resource. The name must not start with `com.amazonaws`.
	Name pulumi.StringPtrInput
	// ID of the AWS account that owns this prefix list.
	OwnerId pulumi.StringPtrInput
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Latest version of this prefix list.
	Version pulumi.IntPtrInput
}

func (ManagedPrefixListState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrefixListState)(nil)).Elem()
}

type managedPrefixListArgs struct {
	// Address family (`IPv4` or `IPv6`) of this prefix list.
	AddressFamily string `pulumi:"addressFamily"`
	// Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries []ManagedPrefixListEntryType `pulumi:"entries"`
	// Maximum number of entries that this prefix list can contain.
	MaxEntries int `pulumi:"maxEntries"`
	// Name of this resource. The name must not start with `com.amazonaws`.
	Name *string `pulumi:"name"`
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedPrefixList resource.
type ManagedPrefixListArgs struct {
	// Address family (`IPv4` or `IPv6`) of this prefix list.
	AddressFamily pulumi.StringInput
	// Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entries ManagedPrefixListEntryTypeArrayInput
	// Maximum number of entries that this prefix list can contain.
	MaxEntries pulumi.IntInput
	// Name of this resource. The name must not start with `com.amazonaws`.
	Name pulumi.StringPtrInput
	// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ManagedPrefixListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrefixListArgs)(nil)).Elem()
}

type ManagedPrefixListInput interface {
	pulumi.Input

	ToManagedPrefixListOutput() ManagedPrefixListOutput
	ToManagedPrefixListOutputWithContext(ctx context.Context) ManagedPrefixListOutput
}

func (*ManagedPrefixList) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrefixList)(nil)).Elem()
}

func (i *ManagedPrefixList) ToManagedPrefixListOutput() ManagedPrefixListOutput {
	return i.ToManagedPrefixListOutputWithContext(context.Background())
}

func (i *ManagedPrefixList) ToManagedPrefixListOutputWithContext(ctx context.Context) ManagedPrefixListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrefixListOutput)
}

// ManagedPrefixListArrayInput is an input type that accepts ManagedPrefixListArray and ManagedPrefixListArrayOutput values.
// You can construct a concrete instance of `ManagedPrefixListArrayInput` via:
//
//	ManagedPrefixListArray{ ManagedPrefixListArgs{...} }
type ManagedPrefixListArrayInput interface {
	pulumi.Input

	ToManagedPrefixListArrayOutput() ManagedPrefixListArrayOutput
	ToManagedPrefixListArrayOutputWithContext(context.Context) ManagedPrefixListArrayOutput
}

type ManagedPrefixListArray []ManagedPrefixListInput

func (ManagedPrefixListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedPrefixList)(nil)).Elem()
}

func (i ManagedPrefixListArray) ToManagedPrefixListArrayOutput() ManagedPrefixListArrayOutput {
	return i.ToManagedPrefixListArrayOutputWithContext(context.Background())
}

func (i ManagedPrefixListArray) ToManagedPrefixListArrayOutputWithContext(ctx context.Context) ManagedPrefixListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrefixListArrayOutput)
}

// ManagedPrefixListMapInput is an input type that accepts ManagedPrefixListMap and ManagedPrefixListMapOutput values.
// You can construct a concrete instance of `ManagedPrefixListMapInput` via:
//
//	ManagedPrefixListMap{ "key": ManagedPrefixListArgs{...} }
type ManagedPrefixListMapInput interface {
	pulumi.Input

	ToManagedPrefixListMapOutput() ManagedPrefixListMapOutput
	ToManagedPrefixListMapOutputWithContext(context.Context) ManagedPrefixListMapOutput
}

type ManagedPrefixListMap map[string]ManagedPrefixListInput

func (ManagedPrefixListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedPrefixList)(nil)).Elem()
}

func (i ManagedPrefixListMap) ToManagedPrefixListMapOutput() ManagedPrefixListMapOutput {
	return i.ToManagedPrefixListMapOutputWithContext(context.Background())
}

func (i ManagedPrefixListMap) ToManagedPrefixListMapOutputWithContext(ctx context.Context) ManagedPrefixListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrefixListMapOutput)
}

type ManagedPrefixListOutput struct{ *pulumi.OutputState }

func (ManagedPrefixListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrefixList)(nil)).Elem()
}

func (o ManagedPrefixListOutput) ToManagedPrefixListOutput() ManagedPrefixListOutput {
	return o
}

func (o ManagedPrefixListOutput) ToManagedPrefixListOutputWithContext(ctx context.Context) ManagedPrefixListOutput {
	return o
}

// Address family (`IPv4` or `IPv6`) of this prefix list.
func (o ManagedPrefixListOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrefixList) pulumi.StringOutput { return v.AddressFamily }).(pulumi.StringOutput)
}

// ARN of the prefix list.
func (o ManagedPrefixListOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrefixList) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
func (o ManagedPrefixListOutput) Entries() ManagedPrefixListEntryTypeArrayOutput {
	return o.ApplyT(func(v *ManagedPrefixList) ManagedPrefixListEntryTypeArrayOutput { return v.Entries }).(ManagedPrefixListEntryTypeArrayOutput)
}

// Maximum number of entries that this prefix list can contain.
func (o ManagedPrefixListOutput) MaxEntries() pulumi.IntOutput {
	return o.ApplyT(func(v *ManagedPrefixList) pulumi.IntOutput { return v.MaxEntries }).(pulumi.IntOutput)
}

// Name of this resource. The name must not start with `com.amazonaws`.
func (o ManagedPrefixListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrefixList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the AWS account that owns this prefix list.
func (o ManagedPrefixListOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrefixList) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// Map of tags to assign to this resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ManagedPrefixListOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedPrefixList) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ManagedPrefixListOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedPrefixList) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Latest version of this prefix list.
func (o ManagedPrefixListOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *ManagedPrefixList) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type ManagedPrefixListArrayOutput struct{ *pulumi.OutputState }

func (ManagedPrefixListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ManagedPrefixList)(nil)).Elem()
}

func (o ManagedPrefixListArrayOutput) ToManagedPrefixListArrayOutput() ManagedPrefixListArrayOutput {
	return o
}

func (o ManagedPrefixListArrayOutput) ToManagedPrefixListArrayOutputWithContext(ctx context.Context) ManagedPrefixListArrayOutput {
	return o
}

func (o ManagedPrefixListArrayOutput) Index(i pulumi.IntInput) ManagedPrefixListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ManagedPrefixList {
		return vs[0].([]*ManagedPrefixList)[vs[1].(int)]
	}).(ManagedPrefixListOutput)
}

type ManagedPrefixListMapOutput struct{ *pulumi.OutputState }

func (ManagedPrefixListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ManagedPrefixList)(nil)).Elem()
}

func (o ManagedPrefixListMapOutput) ToManagedPrefixListMapOutput() ManagedPrefixListMapOutput {
	return o
}

func (o ManagedPrefixListMapOutput) ToManagedPrefixListMapOutputWithContext(ctx context.Context) ManagedPrefixListMapOutput {
	return o
}

func (o ManagedPrefixListMapOutput) MapIndex(k pulumi.StringInput) ManagedPrefixListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ManagedPrefixList {
		return vs[0].(map[string]*ManagedPrefixList)[vs[1].(string)]
	}).(ManagedPrefixListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPrefixListInput)(nil)).Elem(), &ManagedPrefixList{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPrefixListArrayInput)(nil)).Elem(), ManagedPrefixListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManagedPrefixListMapInput)(nil)).Elem(), ManagedPrefixListMap{})
	pulumi.RegisterOutputType(ManagedPrefixListOutput{})
	pulumi.RegisterOutputType(ManagedPrefixListArrayOutput{})
	pulumi.RegisterOutputType(ManagedPrefixListMapOutput{})
}
