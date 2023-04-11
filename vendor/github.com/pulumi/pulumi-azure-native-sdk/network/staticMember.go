// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// StaticMember Item.
// API Version: 2022-02-01-preview.
type StaticMember struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource Id.
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStaticMember registers a new resource with the given unique name, arguments, and options.
func NewStaticMember(ctx *pulumi.Context,
	name string, args *StaticMemberArgs, opts ...pulumi.ResourceOption) (*StaticMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkGroupName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkGroupName'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210501preview:StaticMember"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:StaticMember"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:StaticMember"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:StaticMember"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:StaticMember"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:StaticMember"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:StaticMember"),
		},
	})
	opts = append(opts, aliases)
	var resource StaticMember
	err := ctx.RegisterResource("azure-native:network:StaticMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticMember gets an existing StaticMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticMemberState, opts ...pulumi.ResourceOption) (*StaticMember, error) {
	var resource StaticMember
	err := ctx.ReadResource("azure-native:network:StaticMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticMember resources.
type staticMemberState struct {
}

type StaticMemberState struct {
}

func (StaticMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticMemberState)(nil)).Elem()
}

type staticMemberArgs struct {
	// The name of the network group.
	NetworkGroupName string `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource Id.
	ResourceId *string `pulumi:"resourceId"`
	// The name of the static member.
	StaticMemberName *string `pulumi:"staticMemberName"`
}

// The set of arguments for constructing a StaticMember resource.
type StaticMemberArgs struct {
	// The name of the network group.
	NetworkGroupName pulumi.StringInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource Id.
	ResourceId pulumi.StringPtrInput
	// The name of the static member.
	StaticMemberName pulumi.StringPtrInput
}

func (StaticMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticMemberArgs)(nil)).Elem()
}

type StaticMemberInput interface {
	pulumi.Input

	ToStaticMemberOutput() StaticMemberOutput
	ToStaticMemberOutputWithContext(ctx context.Context) StaticMemberOutput
}

func (*StaticMember) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticMember)(nil)).Elem()
}

func (i *StaticMember) ToStaticMemberOutput() StaticMemberOutput {
	return i.ToStaticMemberOutputWithContext(context.Background())
}

func (i *StaticMember) ToStaticMemberOutputWithContext(ctx context.Context) StaticMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticMemberOutput)
}

type StaticMemberOutput struct{ *pulumi.OutputState }

func (StaticMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticMember)(nil)).Elem()
}

func (o StaticMemberOutput) ToStaticMemberOutput() StaticMemberOutput {
	return o
}

func (o StaticMemberOutput) ToStaticMemberOutputWithContext(ctx context.Context) StaticMemberOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o StaticMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource name.
func (o StaticMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource Id.
func (o StaticMemberOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticMember) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The system metadata related to this resource.
func (o StaticMemberOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StaticMember) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o StaticMemberOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticMember) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticMemberOutput{})
}
