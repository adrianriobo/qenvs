// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified interface endpoint by resource group.
// Azure REST API version: 2019-02-01.
func LookupInterfaceEndpoint(ctx *pulumi.Context, args *LookupInterfaceEndpointArgs, opts ...pulumi.InvokeOption) (*LookupInterfaceEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupInterfaceEndpointResult
	err := ctx.Invoke("azure-native:network:getInterfaceEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupInterfaceEndpointArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the interface endpoint.
	InterfaceEndpointName string `pulumi:"interfaceEndpointName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Interface endpoint resource.
type LookupInterfaceEndpointResult struct {
	// A reference to the service being brought into the virtual network.
	EndpointService *EndpointServiceResponse `pulumi:"endpointService"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// A first-party service's FQDN that is mapped to the private IP allocated via this interface endpoint.
	Fqdn *string `pulumi:"fqdn"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Gets an array of references to the network interfaces created for this interface endpoint.
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	// A read-only property that identifies who created this interface endpoint.
	Owner string `pulumi:"owner"`
	// The provisioning state of the interface endpoint. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState string `pulumi:"provisioningState"`
	// The ID of the subnet from which the private IP will be allocated.
	Subnet *SubnetResponse `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupInterfaceEndpointResult
func (val *LookupInterfaceEndpointResult) Defaults() *LookupInterfaceEndpointResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Subnet = tmp.Subnet.Defaults()

	return &tmp
}

func LookupInterfaceEndpointOutput(ctx *pulumi.Context, args LookupInterfaceEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupInterfaceEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInterfaceEndpointResult, error) {
			args := v.(LookupInterfaceEndpointArgs)
			r, err := LookupInterfaceEndpoint(ctx, &args, opts...)
			var s LookupInterfaceEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInterfaceEndpointResultOutput)
}

type LookupInterfaceEndpointOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the interface endpoint.
	InterfaceEndpointName pulumi.StringInput `pulumi:"interfaceEndpointName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInterfaceEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInterfaceEndpointArgs)(nil)).Elem()
}

// Interface endpoint resource.
type LookupInterfaceEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupInterfaceEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInterfaceEndpointResult)(nil)).Elem()
}

func (o LookupInterfaceEndpointResultOutput) ToLookupInterfaceEndpointResultOutput() LookupInterfaceEndpointResultOutput {
	return o
}

func (o LookupInterfaceEndpointResultOutput) ToLookupInterfaceEndpointResultOutputWithContext(ctx context.Context) LookupInterfaceEndpointResultOutput {
	return o
}

func (o LookupInterfaceEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupInterfaceEndpointResult] {
	return pulumix.Output[LookupInterfaceEndpointResult]{
		OutputState: o.OutputState,
	}
}

// A reference to the service being brought into the virtual network.
func (o LookupInterfaceEndpointResultOutput) EndpointService() EndpointServiceResponsePtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *EndpointServiceResponse { return v.EndpointService }).(EndpointServiceResponsePtrOutput)
}

// Gets a unique read-only string that changes whenever the resource is updated.
func (o LookupInterfaceEndpointResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// A first-party service's FQDN that is mapped to the private IP allocated via this interface endpoint.
func (o LookupInterfaceEndpointResultOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupInterfaceEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupInterfaceEndpointResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupInterfaceEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets an array of references to the network interfaces created for this interface endpoint.
func (o LookupInterfaceEndpointResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

// A read-only property that identifies who created this interface endpoint.
func (o LookupInterfaceEndpointResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) string { return v.Owner }).(pulumi.StringOutput)
}

// The provisioning state of the interface endpoint. Possible values are: 'Updating', 'Deleting', and 'Failed'.
func (o LookupInterfaceEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The ID of the subnet from which the private IP will be allocated.
func (o LookupInterfaceEndpointResultOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

// Resource tags.
func (o LookupInterfaceEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupInterfaceEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInterfaceEndpointResultOutput{})
}
