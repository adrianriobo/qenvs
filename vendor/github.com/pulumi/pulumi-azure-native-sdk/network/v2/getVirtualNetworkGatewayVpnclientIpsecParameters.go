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

// The Get VpnclientIpsecParameters operation retrieves information about the vpnclient ipsec policy for P2S client of virtual network gateway in the specified resource group through Network resource provider.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01.
func GetVirtualNetworkGatewayVpnclientIpsecParameters(ctx *pulumi.Context, args *GetVirtualNetworkGatewayVpnclientIpsecParametersArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayVpnclientIpsecParametersResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetVirtualNetworkGatewayVpnclientIpsecParametersResult
	err := ctx.Invoke("azure-native:network:getVirtualNetworkGatewayVpnclientIpsecParameters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayVpnclientIpsecParametersArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The virtual network gateway name.
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}

// An IPSec parameters for a virtual network gateway P2S connection.
type GetVirtualNetworkGatewayVpnclientIpsecParametersResult struct {
	// The DH Group used in IKE Phase 1 for initial SA.
	DhGroup string `pulumi:"dhGroup"`
	// The IKE encryption algorithm (IKE phase 2).
	IkeEncryption string `pulumi:"ikeEncryption"`
	// The IKE integrity algorithm (IKE phase 2).
	IkeIntegrity string `pulumi:"ikeIntegrity"`
	// The IPSec encryption algorithm (IKE phase 1).
	IpsecEncryption string `pulumi:"ipsecEncryption"`
	// The IPSec integrity algorithm (IKE phase 1).
	IpsecIntegrity string `pulumi:"ipsecIntegrity"`
	// The Pfs Group used in IKE Phase 2 for new child SA.
	PfsGroup string `pulumi:"pfsGroup"`
	// The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for P2S client..
	SaDataSizeKilobytes int `pulumi:"saDataSizeKilobytes"`
	// The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for P2S client.
	SaLifeTimeSeconds int `pulumi:"saLifeTimeSeconds"`
}

func GetVirtualNetworkGatewayVpnclientIpsecParametersOutput(ctx *pulumi.Context, args GetVirtualNetworkGatewayVpnclientIpsecParametersOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualNetworkGatewayVpnclientIpsecParametersResult, error) {
			args := v.(GetVirtualNetworkGatewayVpnclientIpsecParametersArgs)
			r, err := GetVirtualNetworkGatewayVpnclientIpsecParameters(ctx, &args, opts...)
			var s GetVirtualNetworkGatewayVpnclientIpsecParametersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput)
}

type GetVirtualNetworkGatewayVpnclientIpsecParametersOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The virtual network gateway name.
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (GetVirtualNetworkGatewayVpnclientIpsecParametersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayVpnclientIpsecParametersArgs)(nil)).Elem()
}

// An IPSec parameters for a virtual network gateway P2S connection.
type GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayVpnclientIpsecParametersResult)(nil)).Elem()
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) ToGetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput() GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) ToGetVirtualNetworkGatewayVpnclientIpsecParametersResultOutputWithContext(ctx context.Context) GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetVirtualNetworkGatewayVpnclientIpsecParametersResult] {
	return pulumix.Output[GetVirtualNetworkGatewayVpnclientIpsecParametersResult]{
		OutputState: o.OutputState,
	}
}

// The DH Group used in IKE Phase 1 for initial SA.
func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) DhGroup() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.DhGroup }).(pulumi.StringOutput)
}

// The IKE encryption algorithm (IKE phase 2).
func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) IkeEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.IkeEncryption }).(pulumi.StringOutput)
}

// The IKE integrity algorithm (IKE phase 2).
func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) IkeIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.IkeIntegrity }).(pulumi.StringOutput)
}

// The IPSec encryption algorithm (IKE phase 1).
func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) IpsecEncryption() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.IpsecEncryption }).(pulumi.StringOutput)
}

// The IPSec integrity algorithm (IKE phase 1).
func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) IpsecIntegrity() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.IpsecIntegrity }).(pulumi.StringOutput)
}

// The Pfs Group used in IKE Phase 2 for new child SA.
func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) PfsGroup() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) string { return v.PfsGroup }).(pulumi.StringOutput)
}

// The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for P2S client..
func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) SaDataSizeKilobytes() pulumi.IntOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) int { return v.SaDataSizeKilobytes }).(pulumi.IntOutput)
}

// The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for P2S client.
func (o GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput) SaLifeTimeSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayVpnclientIpsecParametersResult) int { return v.SaLifeTimeSeconds }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNetworkGatewayVpnclientIpsecParametersResultOutput{})
}
