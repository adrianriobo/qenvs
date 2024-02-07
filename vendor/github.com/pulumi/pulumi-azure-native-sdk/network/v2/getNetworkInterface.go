// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified network interface.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2015-05-01-preview, 2018-07-01, 2019-02-01, 2019-06-01, 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01.
func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:network:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNetworkInterfaceArgs struct {
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the network interface.
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A network interface in a resource group.
type LookupNetworkInterfaceResult struct {
	// Auxiliary mode of Network Interface resource.
	AuxiliaryMode *string `pulumi:"auxiliaryMode"`
	// Auxiliary sku of Network Interface resource.
	AuxiliarySku *string `pulumi:"auxiliarySku"`
	// Indicates whether to disable tcp state tracking.
	DisableTcpStateTracking *bool `pulumi:"disableTcpStateTracking"`
	// The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettingsResponse `pulumi:"dnsSettings"`
	// A reference to the dscp configuration to which the network interface is linked.
	DscpConfiguration SubResourceResponse `pulumi:"dscpConfiguration"`
	// If the network interface is configured for accelerated networking. Not applicable to VM sizes which require accelerated networking.
	EnableAcceleratedNetworking *bool `pulumi:"enableAcceleratedNetworking"`
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding *bool `pulumi:"enableIPForwarding"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The extended location of the network interface.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// A list of references to linked BareMetal resources.
	HostedWorkloads []string `pulumi:"hostedWorkloads"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterfaceIPConfigurationResponse `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The MAC address of the network interface.
	MacAddress string `pulumi:"macAddress"`
	// Migration phase of Network Interface resource.
	MigrationPhase *string `pulumi:"migrationPhase"`
	// Resource name.
	Name string `pulumi:"name"`
	// The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupResponse `pulumi:"networkSecurityGroup"`
	// Type of Network Interface resource.
	NicType *string `pulumi:"nicType"`
	// Whether this is a primary network interface on a virtual machine.
	Primary bool `pulumi:"primary"`
	// A reference to the private endpoint to which the network interface is linked.
	PrivateEndpoint PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// Privatelinkservice of the network interface resource.
	PrivateLinkService *PrivateLinkServiceResponse `pulumi:"privateLinkService"`
	// The provisioning state of the network interface resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the network interface resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of TapConfigurations of the network interface.
	TapConfigurations []NetworkInterfaceTapConfigurationResponse `pulumi:"tapConfigurations"`
	// Resource type.
	Type string `pulumi:"type"`
	// The reference to a virtual machine.
	VirtualMachine SubResourceResponse `pulumi:"virtualMachine"`
	// Whether the virtual machine this nic is attached to supports encryption.
	VnetEncryptionSupported bool `pulumi:"vnetEncryptionSupported"`
	// WorkloadType of the NetworkInterface for BareMetal resources
	WorkloadType *string `pulumi:"workloadType"`
}

// Defaults sets the appropriate defaults for LookupNetworkInterfaceResult
func (val *LookupNetworkInterfaceResult) Defaults() *LookupNetworkInterfaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PrivateEndpoint = *tmp.PrivateEndpoint.Defaults()

	return &tmp
}

func LookupNetworkInterfaceOutput(ctx *pulumi.Context, args LookupNetworkInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkInterfaceResult, error) {
			args := v.(LookupNetworkInterfaceArgs)
			r, err := LookupNetworkInterface(ctx, &args, opts...)
			var s LookupNetworkInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkInterfaceResultOutput)
}

type LookupNetworkInterfaceOutputArgs struct {
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the network interface.
	NetworkInterfaceName pulumi.StringInput `pulumi:"networkInterfaceName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceArgs)(nil)).Elem()
}

// A network interface in a resource group.
type LookupNetworkInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInterfaceResult)(nil)).Elem()
}

func (o LookupNetworkInterfaceResultOutput) ToLookupNetworkInterfaceResultOutput() LookupNetworkInterfaceResultOutput {
	return o
}

func (o LookupNetworkInterfaceResultOutput) ToLookupNetworkInterfaceResultOutputWithContext(ctx context.Context) LookupNetworkInterfaceResultOutput {
	return o
}

// Auxiliary mode of Network Interface resource.
func (o LookupNetworkInterfaceResultOutput) AuxiliaryMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.AuxiliaryMode }).(pulumi.StringPtrOutput)
}

// Auxiliary sku of Network Interface resource.
func (o LookupNetworkInterfaceResultOutput) AuxiliarySku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.AuxiliarySku }).(pulumi.StringPtrOutput)
}

// Indicates whether to disable tcp state tracking.
func (o LookupNetworkInterfaceResultOutput) DisableTcpStateTracking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.DisableTcpStateTracking }).(pulumi.BoolPtrOutput)
}

// The DNS settings in network interface.
func (o LookupNetworkInterfaceResultOutput) DnsSettings() NetworkInterfaceDnsSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *NetworkInterfaceDnsSettingsResponse { return v.DnsSettings }).(NetworkInterfaceDnsSettingsResponsePtrOutput)
}

// A reference to the dscp configuration to which the network interface is linked.
func (o LookupNetworkInterfaceResultOutput) DscpConfiguration() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) SubResourceResponse { return v.DscpConfiguration }).(SubResourceResponseOutput)
}

// If the network interface is configured for accelerated networking. Not applicable to VM sizes which require accelerated networking.
func (o LookupNetworkInterfaceResultOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

// Indicates whether IP forwarding is enabled on this network interface.
func (o LookupNetworkInterfaceResultOutput) EnableIPForwarding() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *bool { return v.EnableIPForwarding }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkInterfaceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the network interface.
func (o LookupNetworkInterfaceResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// A list of references to linked BareMetal resources.
func (o LookupNetworkInterfaceResultOutput) HostedWorkloads() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []string { return v.HostedWorkloads }).(pulumi.StringArrayOutput)
}

// Resource ID.
func (o LookupNetworkInterfaceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// A list of IPConfigurations of the network interface.
func (o LookupNetworkInterfaceResultOutput) IpConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []NetworkInterfaceIPConfigurationResponse {
		return v.IpConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

// Resource location.
func (o LookupNetworkInterfaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The MAC address of the network interface.
func (o LookupNetworkInterfaceResultOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.MacAddress }).(pulumi.StringOutput)
}

// Migration phase of Network Interface resource.
func (o LookupNetworkInterfaceResultOutput) MigrationPhase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.MigrationPhase }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNetworkInterfaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The reference to the NetworkSecurityGroup resource.
func (o LookupNetworkInterfaceResultOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *NetworkSecurityGroupResponse { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

// Type of Network Interface resource.
func (o LookupNetworkInterfaceResultOutput) NicType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.NicType }).(pulumi.StringPtrOutput)
}

// Whether this is a primary network interface on a virtual machine.
func (o LookupNetworkInterfaceResultOutput) Primary() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) bool { return v.Primary }).(pulumi.BoolOutput)
}

// A reference to the private endpoint to which the network interface is linked.
func (o LookupNetworkInterfaceResultOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponseOutput)
}

// Privatelinkservice of the network interface resource.
func (o LookupNetworkInterfaceResultOutput) PrivateLinkService() PrivateLinkServiceResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *PrivateLinkServiceResponse { return v.PrivateLinkService }).(PrivateLinkServiceResponsePtrOutput)
}

// The provisioning state of the network interface resource.
func (o LookupNetworkInterfaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the network interface resource.
func (o LookupNetworkInterfaceResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupNetworkInterfaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A list of TapConfigurations of the network interface.
func (o LookupNetworkInterfaceResultOutput) TapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) []NetworkInterfaceTapConfigurationResponse {
		return v.TapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

// Resource type.
func (o LookupNetworkInterfaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The reference to a virtual machine.
func (o LookupNetworkInterfaceResultOutput) VirtualMachine() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) SubResourceResponse { return v.VirtualMachine }).(SubResourceResponseOutput)
}

// Whether the virtual machine this nic is attached to supports encryption.
func (o LookupNetworkInterfaceResultOutput) VnetEncryptionSupported() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) bool { return v.VnetEncryptionSupported }).(pulumi.BoolOutput)
}

// WorkloadType of the NetworkInterface for BareMetal resources
func (o LookupNetworkInterfaceResultOutput) WorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkInterfaceResult) *string { return v.WorkloadType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInterfaceResultOutput{})
}
