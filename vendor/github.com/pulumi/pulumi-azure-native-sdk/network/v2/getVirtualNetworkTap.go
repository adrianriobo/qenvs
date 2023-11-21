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

// Gets information about the specified virtual network tap.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01.
func LookupVirtualNetworkTap(ctx *pulumi.Context, args *LookupVirtualNetworkTapArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkTapResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualNetworkTapResult
	err := ctx.Invoke("azure-native:network:getVirtualNetworkTap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualNetworkTapArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of virtual network tap.
	TapName string `pulumi:"tapName"`
}

// Virtual Network Tap resource.
type LookupVirtualNetworkTapResult struct {
	// The reference to the private IP address on the internal Load Balancer that will receive the tap.
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfigurationResponse `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	// The reference to the private IP Address of the collector nic that will receive the tap.
	DestinationNetworkInterfaceIPConfiguration *NetworkInterfaceIPConfigurationResponse `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	// The VXLAN destination port that will receive the tapped traffic.
	DestinationPort *int `pulumi:"destinationPort"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Specifies the list of resource IDs for the network interface IP configuration that needs to be tapped.
	NetworkInterfaceTapConfigurations []NetworkInterfaceTapConfigurationResponse `pulumi:"networkInterfaceTapConfigurations"`
	// The provisioning state of the virtual network tap resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the virtual network tap resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupVirtualNetworkTapResult
func (val *LookupVirtualNetworkTapResult) Defaults() *LookupVirtualNetworkTapResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DestinationLoadBalancerFrontEndIPConfiguration = tmp.DestinationLoadBalancerFrontEndIPConfiguration.Defaults()

	tmp.DestinationNetworkInterfaceIPConfiguration = tmp.DestinationNetworkInterfaceIPConfiguration.Defaults()

	return &tmp
}

func LookupVirtualNetworkTapOutput(ctx *pulumi.Context, args LookupVirtualNetworkTapOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkTapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkTapResult, error) {
			args := v.(LookupVirtualNetworkTapArgs)
			r, err := LookupVirtualNetworkTap(ctx, &args, opts...)
			var s LookupVirtualNetworkTapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkTapResultOutput)
}

type LookupVirtualNetworkTapOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of virtual network tap.
	TapName pulumi.StringInput `pulumi:"tapName"`
}

func (LookupVirtualNetworkTapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkTapArgs)(nil)).Elem()
}

// Virtual Network Tap resource.
type LookupVirtualNetworkTapResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkTapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkTapResult)(nil)).Elem()
}

func (o LookupVirtualNetworkTapResultOutput) ToLookupVirtualNetworkTapResultOutput() LookupVirtualNetworkTapResultOutput {
	return o
}

func (o LookupVirtualNetworkTapResultOutput) ToLookupVirtualNetworkTapResultOutputWithContext(ctx context.Context) LookupVirtualNetworkTapResultOutput {
	return o
}

func (o LookupVirtualNetworkTapResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualNetworkTapResult] {
	return pulumix.Output[LookupVirtualNetworkTapResult]{
		OutputState: o.OutputState,
	}
}

// The reference to the private IP address on the internal Load Balancer that will receive the tap.
func (o LookupVirtualNetworkTapResultOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *FrontendIPConfigurationResponse {
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationResponsePtrOutput)
}

// The reference to the private IP Address of the collector nic that will receive the tap.
func (o LookupVirtualNetworkTapResultOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *NetworkInterfaceIPConfigurationResponse {
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponsePtrOutput)
}

// The VXLAN destination port that will receive the tapped traffic.
func (o LookupVirtualNetworkTapResultOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *int { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualNetworkTapResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupVirtualNetworkTapResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupVirtualNetworkTapResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupVirtualNetworkTapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the list of resource IDs for the network interface IP configuration that needs to be tapped.
func (o LookupVirtualNetworkTapResultOutput) NetworkInterfaceTapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) []NetworkInterfaceTapConfigurationResponse {
		return v.NetworkInterfaceTapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

// The provisioning state of the virtual network tap resource.
func (o LookupVirtualNetworkTapResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the virtual network tap resource.
func (o LookupVirtualNetworkTapResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupVirtualNetworkTapResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupVirtualNetworkTapResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkTapResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkTapResultOutput{})
}
