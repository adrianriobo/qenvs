// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines an Network Experiment Profile and lists of Experiments
// API Version: 2019-11-01.
type NetworkExperimentProfile struct {
	pulumi.CustomResourceState

	// The state of the Experiment
	EnabledState pulumi.StringPtrOutput `pulumi:"enabledState"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource status.
	ResourceState pulumi.StringOutput `pulumi:"resourceState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkExperimentProfile registers a new resource with the given unique name, arguments, and options.
func NewNetworkExperimentProfile(ctx *pulumi.Context,
	name string, args *NetworkExperimentProfileArgs, opts ...pulumi.ResourceOption) (*NetworkExperimentProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkExperimentProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource NetworkExperimentProfile
	err := ctx.RegisterResource("azure-native:network:NetworkExperimentProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkExperimentProfile gets an existing NetworkExperimentProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkExperimentProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkExperimentProfileState, opts ...pulumi.ResourceOption) (*NetworkExperimentProfile, error) {
	var resource NetworkExperimentProfile
	err := ctx.ReadResource("azure-native:network:NetworkExperimentProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkExperimentProfile resources.
type networkExperimentProfileState struct {
}

type NetworkExperimentProfileState struct {
}

func (NetworkExperimentProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkExperimentProfileState)(nil)).Elem()
}

type networkExperimentProfileArgs struct {
	// The state of the Experiment
	EnabledState *string `pulumi:"enabledState"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the Profile
	Name *string `pulumi:"name"`
	// The Profile identifier associated with the Tenant and Partner
	ProfileName *string `pulumi:"profileName"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkExperimentProfile resource.
type NetworkExperimentProfileArgs struct {
	// The state of the Experiment
	EnabledState pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the Profile
	Name pulumi.StringPtrInput
	// The Profile identifier associated with the Tenant and Partner
	ProfileName pulumi.StringPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkExperimentProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkExperimentProfileArgs)(nil)).Elem()
}

type NetworkExperimentProfileInput interface {
	pulumi.Input

	ToNetworkExperimentProfileOutput() NetworkExperimentProfileOutput
	ToNetworkExperimentProfileOutputWithContext(ctx context.Context) NetworkExperimentProfileOutput
}

func (*NetworkExperimentProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkExperimentProfile)(nil)).Elem()
}

func (i *NetworkExperimentProfile) ToNetworkExperimentProfileOutput() NetworkExperimentProfileOutput {
	return i.ToNetworkExperimentProfileOutputWithContext(context.Background())
}

func (i *NetworkExperimentProfile) ToNetworkExperimentProfileOutputWithContext(ctx context.Context) NetworkExperimentProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkExperimentProfileOutput)
}

type NetworkExperimentProfileOutput struct{ *pulumi.OutputState }

func (NetworkExperimentProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkExperimentProfile)(nil)).Elem()
}

func (o NetworkExperimentProfileOutput) ToNetworkExperimentProfileOutput() NetworkExperimentProfileOutput {
	return o
}

func (o NetworkExperimentProfileOutput) ToNetworkExperimentProfileOutputWithContext(ctx context.Context) NetworkExperimentProfileOutput {
	return o
}

// The state of the Experiment
func (o NetworkExperimentProfileOutput) EnabledState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringPtrOutput { return v.EnabledState }).(pulumi.StringPtrOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated.
func (o NetworkExperimentProfileOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o NetworkExperimentProfileOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NetworkExperimentProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource status.
func (o NetworkExperimentProfileOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

// Resource tags.
func (o NetworkExperimentProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o NetworkExperimentProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkExperimentProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkExperimentProfileOutput{})
}