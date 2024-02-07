// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VirtualRouter Resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2022-01-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01.
type VirtualRouter struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The Gateway on which VirtualRouter is hosted.
	HostedGateway SubResourceResponsePtrOutput `pulumi:"hostedGateway"`
	// The Subnet on which VirtualRouter is hosted.
	HostedSubnet SubResourceResponsePtrOutput `pulumi:"hostedSubnet"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of references to VirtualRouterPeerings.
	Peerings SubResourceResponseArrayOutput `pulumi:"peerings"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// VirtualRouter ASN.
	VirtualRouterAsn pulumi.Float64PtrOutput `pulumi:"virtualRouterAsn"`
	// VirtualRouter IPs.
	VirtualRouterIps pulumi.StringArrayOutput `pulumi:"virtualRouterIps"`
}

// NewVirtualRouter registers a new resource with the given unique name, arguments, and options.
func NewVirtualRouter(ctx *pulumi.Context,
	name string, args *VirtualRouterArgs, opts ...pulumi.ResourceOption) (*VirtualRouter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:VirtualRouter"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230901:VirtualRouter"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualRouter
	err := ctx.RegisterResource("azure-native:network:VirtualRouter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualRouter gets an existing VirtualRouter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualRouterState, opts ...pulumi.ResourceOption) (*VirtualRouter, error) {
	var resource VirtualRouter
	err := ctx.ReadResource("azure-native:network:VirtualRouter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualRouter resources.
type virtualRouterState struct {
}

type VirtualRouterState struct {
}

func (VirtualRouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterState)(nil)).Elem()
}

type virtualRouterArgs struct {
	// The Gateway on which VirtualRouter is hosted.
	HostedGateway *SubResource `pulumi:"hostedGateway"`
	// The Subnet on which VirtualRouter is hosted.
	HostedSubnet *SubResource `pulumi:"hostedSubnet"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// VirtualRouter ASN.
	VirtualRouterAsn *float64 `pulumi:"virtualRouterAsn"`
	// VirtualRouter IPs.
	VirtualRouterIps []string `pulumi:"virtualRouterIps"`
	// The name of the Virtual Router.
	VirtualRouterName *string `pulumi:"virtualRouterName"`
}

// The set of arguments for constructing a VirtualRouter resource.
type VirtualRouterArgs struct {
	// The Gateway on which VirtualRouter is hosted.
	HostedGateway SubResourcePtrInput
	// The Subnet on which VirtualRouter is hosted.
	HostedSubnet SubResourcePtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// VirtualRouter ASN.
	VirtualRouterAsn pulumi.Float64PtrInput
	// VirtualRouter IPs.
	VirtualRouterIps pulumi.StringArrayInput
	// The name of the Virtual Router.
	VirtualRouterName pulumi.StringPtrInput
}

func (VirtualRouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterArgs)(nil)).Elem()
}

type VirtualRouterInput interface {
	pulumi.Input

	ToVirtualRouterOutput() VirtualRouterOutput
	ToVirtualRouterOutputWithContext(ctx context.Context) VirtualRouterOutput
}

func (*VirtualRouter) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRouter)(nil)).Elem()
}

func (i *VirtualRouter) ToVirtualRouterOutput() VirtualRouterOutput {
	return i.ToVirtualRouterOutputWithContext(context.Background())
}

func (i *VirtualRouter) ToVirtualRouterOutputWithContext(ctx context.Context) VirtualRouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRouterOutput)
}

type VirtualRouterOutput struct{ *pulumi.OutputState }

func (VirtualRouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRouter)(nil)).Elem()
}

func (o VirtualRouterOutput) ToVirtualRouterOutput() VirtualRouterOutput {
	return o
}

func (o VirtualRouterOutput) ToVirtualRouterOutputWithContext(ctx context.Context) VirtualRouterOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o VirtualRouterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The Gateway on which VirtualRouter is hosted.
func (o VirtualRouterOutput) HostedGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualRouter) SubResourceResponsePtrOutput { return v.HostedGateway }).(SubResourceResponsePtrOutput)
}

// The Subnet on which VirtualRouter is hosted.
func (o VirtualRouterOutput) HostedSubnet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualRouter) SubResourceResponsePtrOutput { return v.HostedSubnet }).(SubResourceResponsePtrOutput)
}

// Resource location.
func (o VirtualRouterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o VirtualRouterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of references to VirtualRouterPeerings.
func (o VirtualRouterOutput) Peerings() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualRouter) SubResourceResponseArrayOutput { return v.Peerings }).(SubResourceResponseArrayOutput)
}

// The provisioning state of the resource.
func (o VirtualRouterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o VirtualRouterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o VirtualRouterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// VirtualRouter ASN.
func (o VirtualRouterOutput) VirtualRouterAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.Float64PtrOutput { return v.VirtualRouterAsn }).(pulumi.Float64PtrOutput)
}

// VirtualRouter IPs.
func (o VirtualRouterOutput) VirtualRouterIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualRouter) pulumi.StringArrayOutput { return v.VirtualRouterIps }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualRouterOutput{})
}
