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

// Gets the specified network profile in a specified resource group.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01.
func LookupNetworkProfile(ctx *pulumi.Context, args *LookupNetworkProfileArgs, opts ...pulumi.InvokeOption) (*LookupNetworkProfileResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkProfileResult
	err := ctx.Invoke("azure-native:network:getNetworkProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkProfileArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the public IP prefix.
	NetworkProfileName string `pulumi:"networkProfileName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Network profile resource.
type LookupNetworkProfileResult struct {
	// List of chid container network interface configurations.
	ContainerNetworkInterfaceConfigurations []ContainerNetworkInterfaceConfigurationResponse `pulumi:"containerNetworkInterfaceConfigurations"`
	// List of child container network interfaces.
	ContainerNetworkInterfaces []ContainerNetworkInterfaceResponse `pulumi:"containerNetworkInterfaces"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the network profile resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the network profile resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupNetworkProfileOutput(ctx *pulumi.Context, args LookupNetworkProfileOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkProfileResult, error) {
			args := v.(LookupNetworkProfileArgs)
			r, err := LookupNetworkProfile(ctx, &args, opts...)
			var s LookupNetworkProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkProfileResultOutput)
}

type LookupNetworkProfileOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the public IP prefix.
	NetworkProfileName pulumi.StringInput `pulumi:"networkProfileName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkProfileArgs)(nil)).Elem()
}

// Network profile resource.
type LookupNetworkProfileResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkProfileResult)(nil)).Elem()
}

func (o LookupNetworkProfileResultOutput) ToLookupNetworkProfileResultOutput() LookupNetworkProfileResultOutput {
	return o
}

func (o LookupNetworkProfileResultOutput) ToLookupNetworkProfileResultOutputWithContext(ctx context.Context) LookupNetworkProfileResultOutput {
	return o
}

func (o LookupNetworkProfileResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNetworkProfileResult] {
	return pulumix.Output[LookupNetworkProfileResult]{
		OutputState: o.OutputState,
	}
}

// List of chid container network interface configurations.
func (o LookupNetworkProfileResultOutput) ContainerNetworkInterfaceConfigurations() ContainerNetworkInterfaceConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) []ContainerNetworkInterfaceConfigurationResponse {
		return v.ContainerNetworkInterfaceConfigurations
	}).(ContainerNetworkInterfaceConfigurationResponseArrayOutput)
}

// List of child container network interfaces.
func (o LookupNetworkProfileResultOutput) ContainerNetworkInterfaces() ContainerNetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) []ContainerNetworkInterfaceResponse {
		return v.ContainerNetworkInterfaces
	}).(ContainerNetworkInterfaceResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkProfileResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupNetworkProfileResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupNetworkProfileResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNetworkProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the network profile resource.
func (o LookupNetworkProfileResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the network profile resource.
func (o LookupNetworkProfileResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupNetworkProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupNetworkProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkProfileResultOutput{})
}
