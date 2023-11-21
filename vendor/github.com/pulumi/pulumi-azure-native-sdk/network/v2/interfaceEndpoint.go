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

// Interface endpoint resource.
// Azure REST API version: 2019-02-01.
type InterfaceEndpoint struct {
	pulumi.CustomResourceState

	// A reference to the service being brought into the virtual network.
	EndpointService EndpointServiceResponsePtrOutput `pulumi:"endpointService"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// A first-party service's FQDN that is mapped to the private IP allocated via this interface endpoint.
	Fqdn pulumi.StringPtrOutput `pulumi:"fqdn"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets an array of references to the network interfaces created for this interface endpoint.
	NetworkInterfaces NetworkInterfaceResponseArrayOutput `pulumi:"networkInterfaces"`
	// A read-only property that identifies who created this interface endpoint.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The provisioning state of the interface endpoint. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The ID of the subnet from which the private IP will be allocated.
	Subnet SubnetResponsePtrOutput `pulumi:"subnet"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewInterfaceEndpoint registers a new resource with the given unique name, arguments, and options.
func NewInterfaceEndpoint(ctx *pulumi.Context,
	name string, args *InterfaceEndpointArgs, opts ...pulumi.ResourceOption) (*InterfaceEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Subnet != nil {
		args.Subnet = args.Subnet.ToSubnetTypePtrOutput().ApplyT(func(v *SubnetType) *SubnetType { return v.Defaults() }).(SubnetTypePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180801:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:InterfaceEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:InterfaceEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource InterfaceEndpoint
	err := ctx.RegisterResource("azure-native:network:InterfaceEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterfaceEndpoint gets an existing InterfaceEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterfaceEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterfaceEndpointState, opts ...pulumi.ResourceOption) (*InterfaceEndpoint, error) {
	var resource InterfaceEndpoint
	err := ctx.ReadResource("azure-native:network:InterfaceEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InterfaceEndpoint resources.
type interfaceEndpointState struct {
}

type InterfaceEndpointState struct {
}

func (InterfaceEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceEndpointState)(nil)).Elem()
}

type interfaceEndpointArgs struct {
	// A reference to the service being brought into the virtual network.
	EndpointService *EndpointService `pulumi:"endpointService"`
	// A first-party service's FQDN that is mapped to the private IP allocated via this interface endpoint.
	Fqdn *string `pulumi:"fqdn"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the interface endpoint.
	InterfaceEndpointName *string `pulumi:"interfaceEndpointName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the subnet from which the private IP will be allocated.
	Subnet *SubnetType `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a InterfaceEndpoint resource.
type InterfaceEndpointArgs struct {
	// A reference to the service being brought into the virtual network.
	EndpointService EndpointServicePtrInput
	// A first-party service's FQDN that is mapped to the private IP allocated via this interface endpoint.
	Fqdn pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the interface endpoint.
	InterfaceEndpointName pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The ID of the subnet from which the private IP will be allocated.
	Subnet SubnetTypePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (InterfaceEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceEndpointArgs)(nil)).Elem()
}

type InterfaceEndpointInput interface {
	pulumi.Input

	ToInterfaceEndpointOutput() InterfaceEndpointOutput
	ToInterfaceEndpointOutputWithContext(ctx context.Context) InterfaceEndpointOutput
}

func (*InterfaceEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceEndpoint)(nil)).Elem()
}

func (i *InterfaceEndpoint) ToInterfaceEndpointOutput() InterfaceEndpointOutput {
	return i.ToInterfaceEndpointOutputWithContext(context.Background())
}

func (i *InterfaceEndpoint) ToInterfaceEndpointOutputWithContext(ctx context.Context) InterfaceEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterfaceEndpointOutput)
}

func (i *InterfaceEndpoint) ToOutput(ctx context.Context) pulumix.Output[*InterfaceEndpoint] {
	return pulumix.Output[*InterfaceEndpoint]{
		OutputState: i.ToInterfaceEndpointOutputWithContext(ctx).OutputState,
	}
}

type InterfaceEndpointOutput struct{ *pulumi.OutputState }

func (InterfaceEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterfaceEndpoint)(nil)).Elem()
}

func (o InterfaceEndpointOutput) ToInterfaceEndpointOutput() InterfaceEndpointOutput {
	return o
}

func (o InterfaceEndpointOutput) ToInterfaceEndpointOutputWithContext(ctx context.Context) InterfaceEndpointOutput {
	return o
}

func (o InterfaceEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*InterfaceEndpoint] {
	return pulumix.Output[*InterfaceEndpoint]{
		OutputState: o.OutputState,
	}
}

// A reference to the service being brought into the virtual network.
func (o InterfaceEndpointOutput) EndpointService() EndpointServiceResponsePtrOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) EndpointServiceResponsePtrOutput { return v.EndpointService }).(EndpointServiceResponsePtrOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated.
func (o InterfaceEndpointOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// A first-party service's FQDN that is mapped to the private IP allocated via this interface endpoint.
func (o InterfaceEndpointOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringPtrOutput { return v.Fqdn }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o InterfaceEndpointOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o InterfaceEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets an array of references to the network interfaces created for this interface endpoint.
func (o InterfaceEndpointOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) NetworkInterfaceResponseArrayOutput { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

// A read-only property that identifies who created this interface endpoint.
func (o InterfaceEndpointOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The provisioning state of the interface endpoint. Possible values are: 'Updating', 'Deleting', and 'Failed'.
func (o InterfaceEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The ID of the subnet from which the private IP will be allocated.
func (o InterfaceEndpointOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) SubnetResponsePtrOutput { return v.Subnet }).(SubnetResponsePtrOutput)
}

// Resource tags.
func (o InterfaceEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o InterfaceEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InterfaceEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InterfaceEndpointOutput{})
}
