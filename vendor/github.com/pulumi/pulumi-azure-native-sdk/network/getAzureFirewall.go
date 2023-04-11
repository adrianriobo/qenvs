// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified Azure Firewall.
// API Version: 2020-11-01.
func LookupAzureFirewall(ctx *pulumi.Context, args *LookupAzureFirewallArgs, opts ...pulumi.InvokeOption) (*LookupAzureFirewallResult, error) {
	var rv LookupAzureFirewallResult
	err := ctx.Invoke("azure-native:network:getAzureFirewall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAzureFirewallArgs struct {
	// The name of the Azure Firewall.
	AzureFirewallName string `pulumi:"azureFirewallName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure Firewall resource.
type LookupAzureFirewallResult struct {
	// The additional properties used to further config this azure firewall.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections []AzureFirewallApplicationRuleCollectionResponse `pulumi:"applicationRuleCollections"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The firewallPolicy associated with this azure firewall.
	FirewallPolicy *SubResourceResponse `pulumi:"firewallPolicy"`
	// IP addresses associated with AzureFirewall.
	HubIPAddresses *HubIPAddressesResponse `pulumi:"hubIPAddresses"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// IP configuration of the Azure Firewall resource.
	IpConfigurations []AzureFirewallIPConfigurationResponse `pulumi:"ipConfigurations"`
	// IpGroups associated with AzureFirewall.
	IpGroups []AzureFirewallIpGroupsResponse `pulumi:"ipGroups"`
	// Resource location.
	Location *string `pulumi:"location"`
	// IP configuration of the Azure Firewall used for management traffic.
	ManagementIpConfiguration *AzureFirewallIPConfigurationResponse `pulumi:"managementIpConfiguration"`
	// Resource name.
	Name string `pulumi:"name"`
	// Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections []AzureFirewallNatRuleCollectionResponse `pulumi:"natRuleCollections"`
	// Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections []AzureFirewallNetworkRuleCollectionResponse `pulumi:"networkRuleCollections"`
	// The provisioning state of the Azure firewall resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The Azure Firewall Resource SKU.
	Sku *AzureFirewallSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The operation mode for Threat Intelligence.
	ThreatIntelMode *string `pulumi:"threatIntelMode"`
	// Resource type.
	Type string `pulumi:"type"`
	// The virtualHub to which the firewall belongs.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

func LookupAzureFirewallOutput(ctx *pulumi.Context, args LookupAzureFirewallOutputArgs, opts ...pulumi.InvokeOption) LookupAzureFirewallResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureFirewallResult, error) {
			args := v.(LookupAzureFirewallArgs)
			r, err := LookupAzureFirewall(ctx, &args, opts...)
			var s LookupAzureFirewallResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureFirewallResultOutput)
}

type LookupAzureFirewallOutputArgs struct {
	// The name of the Azure Firewall.
	AzureFirewallName pulumi.StringInput `pulumi:"azureFirewallName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAzureFirewallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureFirewallArgs)(nil)).Elem()
}

// Azure Firewall resource.
type LookupAzureFirewallResultOutput struct{ *pulumi.OutputState }

func (LookupAzureFirewallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureFirewallResult)(nil)).Elem()
}

func (o LookupAzureFirewallResultOutput) ToLookupAzureFirewallResultOutput() LookupAzureFirewallResultOutput {
	return o
}

func (o LookupAzureFirewallResultOutput) ToLookupAzureFirewallResultOutputWithContext(ctx context.Context) LookupAzureFirewallResultOutput {
	return o
}

// The additional properties used to further config this azure firewall.
func (o LookupAzureFirewallResultOutput) AdditionalProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) map[string]string { return v.AdditionalProperties }).(pulumi.StringMapOutput)
}

// Collection of application rule collections used by Azure Firewall.
func (o LookupAzureFirewallResultOutput) ApplicationRuleCollections() AzureFirewallApplicationRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []AzureFirewallApplicationRuleCollectionResponse {
		return v.ApplicationRuleCollections
	}).(AzureFirewallApplicationRuleCollectionResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupAzureFirewallResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The firewallPolicy associated with this azure firewall.
func (o LookupAzureFirewallResultOutput) FirewallPolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *SubResourceResponse { return v.FirewallPolicy }).(SubResourceResponsePtrOutput)
}

// IP addresses associated with AzureFirewall.
func (o LookupAzureFirewallResultOutput) HubIPAddresses() HubIPAddressesResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *HubIPAddressesResponse { return v.HubIPAddresses }).(HubIPAddressesResponsePtrOutput)
}

// Resource ID.
func (o LookupAzureFirewallResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// IP configuration of the Azure Firewall resource.
func (o LookupAzureFirewallResultOutput) IpConfigurations() AzureFirewallIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []AzureFirewallIPConfigurationResponse { return v.IpConfigurations }).(AzureFirewallIPConfigurationResponseArrayOutput)
}

// IpGroups associated with AzureFirewall.
func (o LookupAzureFirewallResultOutput) IpGroups() AzureFirewallIpGroupsResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []AzureFirewallIpGroupsResponse { return v.IpGroups }).(AzureFirewallIpGroupsResponseArrayOutput)
}

// Resource location.
func (o LookupAzureFirewallResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// IP configuration of the Azure Firewall used for management traffic.
func (o LookupAzureFirewallResultOutput) ManagementIpConfiguration() AzureFirewallIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *AzureFirewallIPConfigurationResponse {
		return v.ManagementIpConfiguration
	}).(AzureFirewallIPConfigurationResponsePtrOutput)
}

// Resource name.
func (o LookupAzureFirewallResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) string { return v.Name }).(pulumi.StringOutput)
}

// Collection of NAT rule collections used by Azure Firewall.
func (o LookupAzureFirewallResultOutput) NatRuleCollections() AzureFirewallNatRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []AzureFirewallNatRuleCollectionResponse {
		return v.NatRuleCollections
	}).(AzureFirewallNatRuleCollectionResponseArrayOutput)
}

// Collection of network rule collections used by Azure Firewall.
func (o LookupAzureFirewallResultOutput) NetworkRuleCollections() AzureFirewallNetworkRuleCollectionResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []AzureFirewallNetworkRuleCollectionResponse {
		return v.NetworkRuleCollections
	}).(AzureFirewallNetworkRuleCollectionResponseArrayOutput)
}

// The provisioning state of the Azure firewall resource.
func (o LookupAzureFirewallResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The Azure Firewall Resource SKU.
func (o LookupAzureFirewallResultOutput) Sku() AzureFirewallSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *AzureFirewallSkuResponse { return v.Sku }).(AzureFirewallSkuResponsePtrOutput)
}

// Resource tags.
func (o LookupAzureFirewallResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The operation mode for Threat Intelligence.
func (o LookupAzureFirewallResultOutput) ThreatIntelMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *string { return v.ThreatIntelMode }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupAzureFirewallResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) string { return v.Type }).(pulumi.StringOutput)
}

// The virtualHub to which the firewall belongs.
func (o LookupAzureFirewallResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

// A list of availability zones denoting where the resource needs to come from.
func (o LookupAzureFirewallResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAzureFirewallResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureFirewallResultOutput{})
}
