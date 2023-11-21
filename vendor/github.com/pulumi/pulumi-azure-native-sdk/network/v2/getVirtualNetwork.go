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

// Gets the specified virtual network by resource group.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2015-05-01-preview, 2019-06-01, 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01.
func LookupVirtualNetwork(ctx *pulumi.Context, args *LookupVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualNetworkResult
	err := ctx.Invoke("azure-native:network:getVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualNetworkArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual network.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// Virtual Network resource.
type LookupVirtualNetworkResult struct {
	// The AddressSpace that contains an array of IP address ranges that can be used by subnets.
	AddressSpace *AddressSpaceResponse `pulumi:"addressSpace"`
	// Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
	BgpCommunities *VirtualNetworkBgpCommunitiesResponse `pulumi:"bgpCommunities"`
	// The DDoS protection plan associated with the virtual network.
	DdosProtectionPlan *SubResourceResponse `pulumi:"ddosProtectionPlan"`
	// The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
	DhcpOptions *DhcpOptionsResponse `pulumi:"dhcpOptions"`
	// Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
	EnableDdosProtection *bool `pulumi:"enableDdosProtection"`
	// Indicates if VM protection is enabled for all the subnets in the virtual network.
	EnableVmProtection *bool `pulumi:"enableVmProtection"`
	// Indicates if encryption is enabled on virtual network and if VM without encryption is allowed in encrypted VNet.
	Encryption *VirtualNetworkEncryptionResponse `pulumi:"encryption"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The extended location of the virtual network.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// A collection of references to flow log resources.
	FlowLogs []FlowLogResponse `pulumi:"flowLogs"`
	// The FlowTimeout value (in minutes) for the Virtual Network
	FlowTimeoutInMinutes *int `pulumi:"flowTimeoutInMinutes"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Array of IpAllocation which reference this VNET.
	IpAllocations []SubResourceResponse `pulumi:"ipAllocations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the virtual network resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resourceGuid property of the Virtual Network resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// A list of subnets in a Virtual Network.
	Subnets []SubnetResponse `pulumi:"subnets"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// A list of peerings in a Virtual Network.
	VirtualNetworkPeerings []VirtualNetworkPeeringResponse `pulumi:"virtualNetworkPeerings"`
}

// Defaults sets the appropriate defaults for LookupVirtualNetworkResult
func (val *LookupVirtualNetworkResult) Defaults() *LookupVirtualNetworkResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.EnableDdosProtection == nil {
		enableDdosProtection_ := false
		tmp.EnableDdosProtection = &enableDdosProtection_
	}
	if tmp.EnableVmProtection == nil {
		enableVmProtection_ := false
		tmp.EnableVmProtection = &enableVmProtection_
	}
	return &tmp
}

func LookupVirtualNetworkOutput(ctx *pulumi.Context, args LookupVirtualNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkResult, error) {
			args := v.(LookupVirtualNetworkArgs)
			r, err := LookupVirtualNetwork(ctx, &args, opts...)
			var s LookupVirtualNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkResultOutput)
}

type LookupVirtualNetworkOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the virtual network.
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (LookupVirtualNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkArgs)(nil)).Elem()
}

// Virtual Network resource.
type LookupVirtualNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkResult)(nil)).Elem()
}

func (o LookupVirtualNetworkResultOutput) ToLookupVirtualNetworkResultOutput() LookupVirtualNetworkResultOutput {
	return o
}

func (o LookupVirtualNetworkResultOutput) ToLookupVirtualNetworkResultOutputWithContext(ctx context.Context) LookupVirtualNetworkResultOutput {
	return o
}

func (o LookupVirtualNetworkResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualNetworkResult] {
	return pulumix.Output[LookupVirtualNetworkResult]{
		OutputState: o.OutputState,
	}
}

// The AddressSpace that contains an array of IP address ranges that can be used by subnets.
func (o LookupVirtualNetworkResultOutput) AddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *AddressSpaceResponse { return v.AddressSpace }).(AddressSpaceResponsePtrOutput)
}

// Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
func (o LookupVirtualNetworkResultOutput) BgpCommunities() VirtualNetworkBgpCommunitiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *VirtualNetworkBgpCommunitiesResponse { return v.BgpCommunities }).(VirtualNetworkBgpCommunitiesResponsePtrOutput)
}

// The DDoS protection plan associated with the virtual network.
func (o LookupVirtualNetworkResultOutput) DdosProtectionPlan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *SubResourceResponse { return v.DdosProtectionPlan }).(SubResourceResponsePtrOutput)
}

// The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
func (o LookupVirtualNetworkResultOutput) DhcpOptions() DhcpOptionsResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *DhcpOptionsResponse { return v.DhcpOptions }).(DhcpOptionsResponsePtrOutput)
}

// Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It requires a DDoS protection plan associated with the resource.
func (o LookupVirtualNetworkResultOutput) EnableDdosProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *bool { return v.EnableDdosProtection }).(pulumi.BoolPtrOutput)
}

// Indicates if VM protection is enabled for all the subnets in the virtual network.
func (o LookupVirtualNetworkResultOutput) EnableVmProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *bool { return v.EnableVmProtection }).(pulumi.BoolPtrOutput)
}

// Indicates if encryption is enabled on virtual network and if VM without encryption is allowed in encrypted VNet.
func (o LookupVirtualNetworkResultOutput) Encryption() VirtualNetworkEncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *VirtualNetworkEncryptionResponse { return v.Encryption }).(VirtualNetworkEncryptionResponsePtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupVirtualNetworkResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the virtual network.
func (o LookupVirtualNetworkResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// A collection of references to flow log resources.
func (o LookupVirtualNetworkResultOutput) FlowLogs() FlowLogResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []FlowLogResponse { return v.FlowLogs }).(FlowLogResponseArrayOutput)
}

// The FlowTimeout value (in minutes) for the Virtual Network
func (o LookupVirtualNetworkResultOutput) FlowTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *int { return v.FlowTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

// Resource ID.
func (o LookupVirtualNetworkResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Array of IpAllocation which reference this VNET.
func (o LookupVirtualNetworkResultOutput) IpAllocations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []SubResourceResponse { return v.IpAllocations }).(SubResourceResponseArrayOutput)
}

// Resource location.
func (o LookupVirtualNetworkResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupVirtualNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the virtual network resource.
func (o LookupVirtualNetworkResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resourceGuid property of the Virtual Network resource.
func (o LookupVirtualNetworkResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// A list of subnets in a Virtual Network.
func (o LookupVirtualNetworkResultOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

// Resource tags.
func (o LookupVirtualNetworkResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupVirtualNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of peerings in a Virtual Network.
func (o LookupVirtualNetworkResultOutput) VirtualNetworkPeerings() VirtualNetworkPeeringResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkResult) []VirtualNetworkPeeringResponse { return v.VirtualNetworkPeerings }).(VirtualNetworkPeeringResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkResultOutput{})
}
