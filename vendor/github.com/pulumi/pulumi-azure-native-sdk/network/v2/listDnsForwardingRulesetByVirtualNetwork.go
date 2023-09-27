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

// Lists DNS forwarding ruleset resource IDs attached to a virtual network.
// Azure REST API version: 2022-07-01.
func ListDnsForwardingRulesetByVirtualNetwork(ctx *pulumi.Context, args *ListDnsForwardingRulesetByVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*ListDnsForwardingRulesetByVirtualNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDnsForwardingRulesetByVirtualNetworkResult
	err := ctx.Invoke("azure-native:network:listDnsForwardingRulesetByVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDnsForwardingRulesetByVirtualNetworkArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top *int `pulumi:"top"`
	// The name of the virtual network.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// The response to an enumeration operation on Virtual Network DNS Forwarding Ruleset.
type ListDnsForwardingRulesetByVirtualNetworkResult struct {
	// The continuation token for the next page of results.
	NextLink string `pulumi:"nextLink"`
	// Enumeration of the Virtual Network DNS Forwarding Ruleset.
	Value []VirtualNetworkDnsForwardingRulesetResponse `pulumi:"value"`
}

func ListDnsForwardingRulesetByVirtualNetworkOutput(ctx *pulumi.Context, args ListDnsForwardingRulesetByVirtualNetworkOutputArgs, opts ...pulumi.InvokeOption) ListDnsForwardingRulesetByVirtualNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDnsForwardingRulesetByVirtualNetworkResult, error) {
			args := v.(ListDnsForwardingRulesetByVirtualNetworkArgs)
			r, err := ListDnsForwardingRulesetByVirtualNetwork(ctx, &args, opts...)
			var s ListDnsForwardingRulesetByVirtualNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDnsForwardingRulesetByVirtualNetworkResultOutput)
}

type ListDnsForwardingRulesetByVirtualNetworkOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The maximum number of results to return. If not specified, returns up to 100 results.
	Top pulumi.IntPtrInput `pulumi:"top"`
	// The name of the virtual network.
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (ListDnsForwardingRulesetByVirtualNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDnsForwardingRulesetByVirtualNetworkArgs)(nil)).Elem()
}

// The response to an enumeration operation on Virtual Network DNS Forwarding Ruleset.
type ListDnsForwardingRulesetByVirtualNetworkResultOutput struct{ *pulumi.OutputState }

func (ListDnsForwardingRulesetByVirtualNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDnsForwardingRulesetByVirtualNetworkResult)(nil)).Elem()
}

func (o ListDnsForwardingRulesetByVirtualNetworkResultOutput) ToListDnsForwardingRulesetByVirtualNetworkResultOutput() ListDnsForwardingRulesetByVirtualNetworkResultOutput {
	return o
}

func (o ListDnsForwardingRulesetByVirtualNetworkResultOutput) ToListDnsForwardingRulesetByVirtualNetworkResultOutputWithContext(ctx context.Context) ListDnsForwardingRulesetByVirtualNetworkResultOutput {
	return o
}

func (o ListDnsForwardingRulesetByVirtualNetworkResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListDnsForwardingRulesetByVirtualNetworkResult] {
	return pulumix.Output[ListDnsForwardingRulesetByVirtualNetworkResult]{
		OutputState: o.OutputState,
	}
}

// The continuation token for the next page of results.
func (o ListDnsForwardingRulesetByVirtualNetworkResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListDnsForwardingRulesetByVirtualNetworkResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Enumeration of the Virtual Network DNS Forwarding Ruleset.
func (o ListDnsForwardingRulesetByVirtualNetworkResultOutput) Value() VirtualNetworkDnsForwardingRulesetResponseArrayOutput {
	return o.ApplyT(func(v ListDnsForwardingRulesetByVirtualNetworkResult) []VirtualNetworkDnsForwardingRulesetResponse {
		return v.Value
	}).(VirtualNetworkDnsForwardingRulesetResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDnsForwardingRulesetByVirtualNetworkResultOutput{})
}
