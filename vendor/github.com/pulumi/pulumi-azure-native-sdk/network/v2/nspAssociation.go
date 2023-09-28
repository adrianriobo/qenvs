// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The NSP resource association resource
// Azure REST API version: 2021-02-01-preview. Prior API version in Azure Native 1.x: 2021-02-01-preview
type NspAssociation struct {
	pulumi.CustomResourceState

	// Access mode on the association.
	AccessMode pulumi.StringPtrOutput `pulumi:"accessMode"`
	// Specifies if there are provisioning issues
	HasProvisioningIssues pulumi.StringOutput `pulumi:"hasProvisioningIssues"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The PaaS resource to be associated.
	PrivateLinkResource SubResourceResponsePtrOutput `pulumi:"privateLinkResource"`
	// Profile id to which the PaaS resource is associated.
	Profile SubResourceResponsePtrOutput `pulumi:"profile"`
	// The provisioning state of the resource  association resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNspAssociation registers a new resource with the given unique name, arguments, and options.
func NewNspAssociation(ctx *pulumi.Context,
	name string, args *NspAssociationArgs, opts ...pulumi.ResourceOption) (*NspAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkSecurityPerimeterName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityPerimeterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210201preview:NspAssociation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NspAssociation
	err := ctx.RegisterResource("azure-native:network:NspAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNspAssociation gets an existing NspAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNspAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NspAssociationState, opts ...pulumi.ResourceOption) (*NspAssociation, error) {
	var resource NspAssociation
	err := ctx.ReadResource("azure-native:network:NspAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NspAssociation resources.
type nspAssociationState struct {
}

type NspAssociationState struct {
}

func (NspAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*nspAssociationState)(nil)).Elem()
}

type nspAssociationArgs struct {
	// Access mode on the association.
	AccessMode *string `pulumi:"accessMode"`
	// The name of the NSP association.
	AssociationName *string `pulumi:"associationName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	// The PaaS resource to be associated.
	PrivateLinkResource *SubResource `pulumi:"privateLinkResource"`
	// Profile id to which the PaaS resource is associated.
	Profile *SubResource `pulumi:"profile"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NspAssociation resource.
type NspAssociationArgs struct {
	// Access mode on the association.
	AccessMode pulumi.StringPtrInput
	// The name of the NSP association.
	AssociationName pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName pulumi.StringInput
	// The PaaS resource to be associated.
	PrivateLinkResource SubResourcePtrInput
	// Profile id to which the PaaS resource is associated.
	Profile SubResourcePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NspAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nspAssociationArgs)(nil)).Elem()
}

type NspAssociationInput interface {
	pulumi.Input

	ToNspAssociationOutput() NspAssociationOutput
	ToNspAssociationOutputWithContext(ctx context.Context) NspAssociationOutput
}

func (*NspAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**NspAssociation)(nil)).Elem()
}

func (i *NspAssociation) ToNspAssociationOutput() NspAssociationOutput {
	return i.ToNspAssociationOutputWithContext(context.Background())
}

func (i *NspAssociation) ToNspAssociationOutputWithContext(ctx context.Context) NspAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NspAssociationOutput)
}

func (i *NspAssociation) ToOutput(ctx context.Context) pulumix.Output[*NspAssociation] {
	return pulumix.Output[*NspAssociation]{
		OutputState: i.ToNspAssociationOutputWithContext(ctx).OutputState,
	}
}

type NspAssociationOutput struct{ *pulumi.OutputState }

func (NspAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NspAssociation)(nil)).Elem()
}

func (o NspAssociationOutput) ToNspAssociationOutput() NspAssociationOutput {
	return o
}

func (o NspAssociationOutput) ToNspAssociationOutputWithContext(ctx context.Context) NspAssociationOutput {
	return o
}

func (o NspAssociationOutput) ToOutput(ctx context.Context) pulumix.Output[*NspAssociation] {
	return pulumix.Output[*NspAssociation]{
		OutputState: o.OutputState,
	}
}

// Access mode on the association.
func (o NspAssociationOutput) AccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringPtrOutput { return v.AccessMode }).(pulumi.StringPtrOutput)
}

// Specifies if there are provisioning issues
func (o NspAssociationOutput) HasProvisioningIssues() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringOutput { return v.HasProvisioningIssues }).(pulumi.StringOutput)
}

// Resource location.
func (o NspAssociationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NspAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The PaaS resource to be associated.
func (o NspAssociationOutput) PrivateLinkResource() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NspAssociation) SubResourceResponsePtrOutput { return v.PrivateLinkResource }).(SubResourceResponsePtrOutput)
}

// Profile id to which the PaaS resource is associated.
func (o NspAssociationOutput) Profile() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *NspAssociation) SubResourceResponsePtrOutput { return v.Profile }).(SubResourceResponsePtrOutput)
}

// The provisioning state of the resource  association resource.
func (o NspAssociationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o NspAssociationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o NspAssociationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NspAssociation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NspAssociationOutput{})
}
