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

// Gets the connection health of P2S clients of the virtual wan P2SVpnGateway in the specified resource group.
// Azure REST API version: 2023-02-01.
func GetP2sVpnGatewayP2sVpnConnectionHealth(ctx *pulumi.Context, args *GetP2sVpnGatewayP2sVpnConnectionHealthArgs, opts ...pulumi.InvokeOption) (*GetP2sVpnGatewayP2sVpnConnectionHealthResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetP2sVpnGatewayP2sVpnConnectionHealthResult
	err := ctx.Invoke("azure-native:network:getP2sVpnGatewayP2sVpnConnectionHealth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetP2sVpnGatewayP2sVpnConnectionHealthArgs struct {
	// The name of the P2SVpnGateway.
	GatewayName string `pulumi:"gatewayName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// P2SVpnGateway Resource.
type GetP2sVpnGatewayP2sVpnConnectionHealthResult struct {
	// List of all customer specified DNS servers IP addresses.
	CustomDnsServers []string `pulumi:"customDnsServers"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Enable Routing Preference property for the Public IP Interface of the P2SVpnGateway.
	IsRoutingPreferenceInternet *bool `pulumi:"isRoutingPreferenceInternet"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// List of all p2s connection configurations of the gateway.
	P2SConnectionConfigurations []P2SConnectionConfigurationResponse `pulumi:"p2SConnectionConfigurations"`
	// The provisioning state of the P2S VPN gateway resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The VirtualHub to which the gateway belongs.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
	// All P2S VPN clients' connection health status.
	VpnClientConnectionHealth VpnClientConnectionHealthResponse `pulumi:"vpnClientConnectionHealth"`
	// The scale unit for this p2s vpn gateway.
	VpnGatewayScaleUnit *int `pulumi:"vpnGatewayScaleUnit"`
	// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
	VpnServerConfiguration *SubResourceResponse `pulumi:"vpnServerConfiguration"`
}

func GetP2sVpnGatewayP2sVpnConnectionHealthOutput(ctx *pulumi.Context, args GetP2sVpnGatewayP2sVpnConnectionHealthOutputArgs, opts ...pulumi.InvokeOption) GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetP2sVpnGatewayP2sVpnConnectionHealthResult, error) {
			args := v.(GetP2sVpnGatewayP2sVpnConnectionHealthArgs)
			r, err := GetP2sVpnGatewayP2sVpnConnectionHealth(ctx, &args, opts...)
			var s GetP2sVpnGatewayP2sVpnConnectionHealthResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput)
}

type GetP2sVpnGatewayP2sVpnConnectionHealthOutputArgs struct {
	// The name of the P2SVpnGateway.
	GatewayName pulumi.StringInput `pulumi:"gatewayName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetP2sVpnGatewayP2sVpnConnectionHealthOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetP2sVpnGatewayP2sVpnConnectionHealthArgs)(nil)).Elem()
}

// P2SVpnGateway Resource.
type GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput struct{ *pulumi.OutputState }

func (GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetP2sVpnGatewayP2sVpnConnectionHealthResult)(nil)).Elem()
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) ToGetP2sVpnGatewayP2sVpnConnectionHealthResultOutput() GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput {
	return o
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) ToGetP2sVpnGatewayP2sVpnConnectionHealthResultOutputWithContext(ctx context.Context) GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput {
	return o
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetP2sVpnGatewayP2sVpnConnectionHealthResult] {
	return pulumix.Output[GetP2sVpnGatewayP2sVpnConnectionHealthResult]{
		OutputState: o.OutputState,
	}
}

// List of all customer specified DNS servers IP addresses.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) CustomDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) []string { return v.CustomDnsServers }).(pulumi.StringArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Enable Routing Preference property for the Public IP Interface of the P2SVpnGateway.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) IsRoutingPreferenceInternet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) *bool { return v.IsRoutingPreferenceInternet }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of all p2s connection configurations of the gateway.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) P2SConnectionConfigurations() P2SConnectionConfigurationResponseArrayOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) []P2SConnectionConfigurationResponse {
		return v.P2SConnectionConfigurations
	}).(P2SConnectionConfigurationResponseArrayOutput)
}

// The provisioning state of the P2S VPN gateway resource.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource tags.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) string { return v.Type }).(pulumi.StringOutput)
}

// The VirtualHub to which the gateway belongs.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

// All P2S VPN clients' connection health status.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) VpnClientConnectionHealth() VpnClientConnectionHealthResponseOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) VpnClientConnectionHealthResponse {
		return v.VpnClientConnectionHealth
	}).(VpnClientConnectionHealthResponseOutput)
}

// The scale unit for this p2s vpn gateway.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) VpnGatewayScaleUnit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) *int { return v.VpnGatewayScaleUnit }).(pulumi.IntPtrOutput)
}

// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
func (o GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput) VpnServerConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthResult) *SubResourceResponse {
		return v.VpnServerConfiguration
	}).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetP2sVpnGatewayP2sVpnConnectionHealthResultOutput{})
}
