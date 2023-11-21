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

// Public IP prefix resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2019-06-01, 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01.
type PublicIPPrefix struct {
	pulumi.CustomResourceState

	// The customIpPrefix that this prefix is associated with.
	CustomIPPrefix SubResourceResponsePtrOutput `pulumi:"customIPPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The extended location of the public ip address.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The allocated Prefix.
	IpPrefix pulumi.StringOutput `pulumi:"ipPrefix"`
	// The list of tags associated with the public IP prefix.
	IpTags IpTagResponseArrayOutput `pulumi:"ipTags"`
	// The reference to load balancer frontend IP configuration associated with the public IP prefix.
	LoadBalancerFrontendIpConfiguration SubResourceResponseOutput `pulumi:"loadBalancerFrontendIpConfiguration"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// NatGateway of Public IP Prefix.
	NatGateway NatGatewayResponsePtrOutput `pulumi:"natGateway"`
	// The Length of the Public IP Prefix.
	PrefixLength pulumi.IntPtrOutput `pulumi:"prefixLength"`
	// The provisioning state of the public IP prefix resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The public IP address version.
	PublicIPAddressVersion pulumi.StringPtrOutput `pulumi:"publicIPAddressVersion"`
	// The list of all referenced PublicIPAddresses.
	PublicIPAddresses ReferencedPublicIpAddressResponseArrayOutput `pulumi:"publicIPAddresses"`
	// The resource GUID property of the public IP prefix resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// The public IP prefix SKU.
	Sku PublicIPPrefixSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewPublicIPPrefix registers a new resource with the given unique name, arguments, and options.
func NewPublicIPPrefix(ctx *pulumi.Context,
	name string, args *PublicIPPrefixArgs, opts ...pulumi.ResourceOption) (*PublicIPPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:PublicIPPrefix"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:PublicIPPrefix"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PublicIPPrefix
	err := ctx.RegisterResource("azure-native:network:PublicIPPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIPPrefix gets an existing PublicIPPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIPPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIPPrefixState, opts ...pulumi.ResourceOption) (*PublicIPPrefix, error) {
	var resource PublicIPPrefix
	err := ctx.ReadResource("azure-native:network:PublicIPPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIPPrefix resources.
type publicIPPrefixState struct {
}

type PublicIPPrefixState struct {
}

func (PublicIPPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPPrefixState)(nil)).Elem()
}

type publicIPPrefixArgs struct {
	// The customIpPrefix that this prefix is associated with.
	CustomIPPrefix *SubResource `pulumi:"customIPPrefix"`
	// The extended location of the public ip address.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The list of tags associated with the public IP prefix.
	IpTags []IpTag `pulumi:"ipTags"`
	// Resource location.
	Location *string `pulumi:"location"`
	// NatGateway of Public IP Prefix.
	NatGateway *NatGatewayType `pulumi:"natGateway"`
	// The Length of the Public IP Prefix.
	PrefixLength *int `pulumi:"prefixLength"`
	// The public IP address version.
	PublicIPAddressVersion *string `pulumi:"publicIPAddressVersion"`
	// The name of the public IP prefix.
	PublicIpPrefixName *string `pulumi:"publicIpPrefixName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The public IP prefix SKU.
	Sku *PublicIPPrefixSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a PublicIPPrefix resource.
type PublicIPPrefixArgs struct {
	// The customIpPrefix that this prefix is associated with.
	CustomIPPrefix SubResourcePtrInput
	// The extended location of the public ip address.
	ExtendedLocation ExtendedLocationPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The list of tags associated with the public IP prefix.
	IpTags IpTagArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// NatGateway of Public IP Prefix.
	NatGateway NatGatewayTypePtrInput
	// The Length of the Public IP Prefix.
	PrefixLength pulumi.IntPtrInput
	// The public IP address version.
	PublicIPAddressVersion pulumi.StringPtrInput
	// The name of the public IP prefix.
	PublicIpPrefixName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The public IP prefix SKU.
	Sku PublicIPPrefixSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (PublicIPPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIPPrefixArgs)(nil)).Elem()
}

type PublicIPPrefixInput interface {
	pulumi.Input

	ToPublicIPPrefixOutput() PublicIPPrefixOutput
	ToPublicIPPrefixOutputWithContext(ctx context.Context) PublicIPPrefixOutput
}

func (*PublicIPPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPPrefix)(nil)).Elem()
}

func (i *PublicIPPrefix) ToPublicIPPrefixOutput() PublicIPPrefixOutput {
	return i.ToPublicIPPrefixOutputWithContext(context.Background())
}

func (i *PublicIPPrefix) ToPublicIPPrefixOutputWithContext(ctx context.Context) PublicIPPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPPrefixOutput)
}

func (i *PublicIPPrefix) ToOutput(ctx context.Context) pulumix.Output[*PublicIPPrefix] {
	return pulumix.Output[*PublicIPPrefix]{
		OutputState: i.ToPublicIPPrefixOutputWithContext(ctx).OutputState,
	}
}

type PublicIPPrefixOutput struct{ *pulumi.OutputState }

func (PublicIPPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIPPrefix)(nil)).Elem()
}

func (o PublicIPPrefixOutput) ToPublicIPPrefixOutput() PublicIPPrefixOutput {
	return o
}

func (o PublicIPPrefixOutput) ToPublicIPPrefixOutputWithContext(ctx context.Context) PublicIPPrefixOutput {
	return o
}

func (o PublicIPPrefixOutput) ToOutput(ctx context.Context) pulumix.Output[*PublicIPPrefix] {
	return pulumix.Output[*PublicIPPrefix]{
		OutputState: o.OutputState,
	}
}

// The customIpPrefix that this prefix is associated with.
func (o PublicIPPrefixOutput) CustomIPPrefix() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) SubResourceResponsePtrOutput { return v.CustomIPPrefix }).(SubResourceResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o PublicIPPrefixOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the public ip address.
func (o PublicIPPrefixOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The allocated Prefix.
func (o PublicIPPrefixOutput) IpPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.IpPrefix }).(pulumi.StringOutput)
}

// The list of tags associated with the public IP prefix.
func (o PublicIPPrefixOutput) IpTags() IpTagResponseArrayOutput {
	return o.ApplyT(func(v *PublicIPPrefix) IpTagResponseArrayOutput { return v.IpTags }).(IpTagResponseArrayOutput)
}

// The reference to load balancer frontend IP configuration associated with the public IP prefix.
func (o PublicIPPrefixOutput) LoadBalancerFrontendIpConfiguration() SubResourceResponseOutput {
	return o.ApplyT(func(v *PublicIPPrefix) SubResourceResponseOutput { return v.LoadBalancerFrontendIpConfiguration }).(SubResourceResponseOutput)
}

// Resource location.
func (o PublicIPPrefixOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o PublicIPPrefixOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// NatGateway of Public IP Prefix.
func (o PublicIPPrefixOutput) NatGateway() NatGatewayResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) NatGatewayResponsePtrOutput { return v.NatGateway }).(NatGatewayResponsePtrOutput)
}

// The Length of the Public IP Prefix.
func (o PublicIPPrefixOutput) PrefixLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.IntPtrOutput { return v.PrefixLength }).(pulumi.IntPtrOutput)
}

// The provisioning state of the public IP prefix resource.
func (o PublicIPPrefixOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The public IP address version.
func (o PublicIPPrefixOutput) PublicIPAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringPtrOutput { return v.PublicIPAddressVersion }).(pulumi.StringPtrOutput)
}

// The list of all referenced PublicIPAddresses.
func (o PublicIPPrefixOutput) PublicIPAddresses() ReferencedPublicIpAddressResponseArrayOutput {
	return o.ApplyT(func(v *PublicIPPrefix) ReferencedPublicIpAddressResponseArrayOutput { return v.PublicIPAddresses }).(ReferencedPublicIpAddressResponseArrayOutput)
}

// The resource GUID property of the public IP prefix resource.
func (o PublicIPPrefixOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The public IP prefix SKU.
func (o PublicIPPrefixOutput) Sku() PublicIPPrefixSkuResponsePtrOutput {
	return o.ApplyT(func(v *PublicIPPrefix) PublicIPPrefixSkuResponsePtrOutput { return v.Sku }).(PublicIPPrefixSkuResponsePtrOutput)
}

// Resource tags.
func (o PublicIPPrefixOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o PublicIPPrefixOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of availability zones denoting the IP allocated for the resource needs to come from.
func (o PublicIPPrefixOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PublicIPPrefix) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicIPPrefixOutput{})
}
