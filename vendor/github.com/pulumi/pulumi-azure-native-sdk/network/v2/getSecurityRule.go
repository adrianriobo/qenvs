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

// Get the specified network security rule.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2017-03-01, 2019-06-01, 2022-07-01, 2023-04-01, 2023-05-01, 2023-06-01.
func LookupSecurityRule(ctx *pulumi.Context, args *LookupSecurityRuleArgs, opts ...pulumi.InvokeOption) (*LookupSecurityRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityRuleResult
	err := ctx.Invoke("azure-native:network:getSecurityRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityRuleArgs struct {
	// The name of the network security group.
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the security rule.
	SecurityRuleName string `pulumi:"securityRuleName"`
}

// Network security rule.
type LookupSecurityRuleResult struct {
	// The network traffic is allowed or denied.
	Access string `pulumi:"access"`
	// A description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	// The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes []string `pulumi:"destinationAddressPrefixes"`
	// The application security group specified as destination.
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupResponse `pulumi:"destinationApplicationSecurityGroups"`
	// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// The destination port ranges.
	DestinationPortRanges []string `pulumi:"destinationPortRanges"`
	// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction string `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority int `pulumi:"priority"`
	// Network protocol this rule applies to.
	Protocol string `pulumi:"protocol"`
	// The provisioning state of the security rule resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix *string `pulumi:"sourceAddressPrefix"`
	// The CIDR or source IP ranges.
	SourceAddressPrefixes []string `pulumi:"sourceAddressPrefixes"`
	// The application security group specified as source.
	SourceApplicationSecurityGroups []ApplicationSecurityGroupResponse `pulumi:"sourceApplicationSecurityGroups"`
	// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// The source port ranges.
	SourcePortRanges []string `pulumi:"sourcePortRanges"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

func LookupSecurityRuleOutput(ctx *pulumi.Context, args LookupSecurityRuleOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityRuleResult, error) {
			args := v.(LookupSecurityRuleArgs)
			r, err := LookupSecurityRule(ctx, &args, opts...)
			var s LookupSecurityRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityRuleResultOutput)
}

type LookupSecurityRuleOutputArgs struct {
	// The name of the network security group.
	NetworkSecurityGroupName pulumi.StringInput `pulumi:"networkSecurityGroupName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the security rule.
	SecurityRuleName pulumi.StringInput `pulumi:"securityRuleName"`
}

func (LookupSecurityRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityRuleArgs)(nil)).Elem()
}

// Network security rule.
type LookupSecurityRuleResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityRuleResult)(nil)).Elem()
}

func (o LookupSecurityRuleResultOutput) ToLookupSecurityRuleResultOutput() LookupSecurityRuleResultOutput {
	return o
}

func (o LookupSecurityRuleResultOutput) ToLookupSecurityRuleResultOutputWithContext(ctx context.Context) LookupSecurityRuleResultOutput {
	return o
}

func (o LookupSecurityRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSecurityRuleResult] {
	return pulumix.Output[LookupSecurityRuleResult]{
		OutputState: o.OutputState,
	}
}

// The network traffic is allowed or denied.
func (o LookupSecurityRuleResultOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.Access }).(pulumi.StringOutput)
}

// A description for this rule. Restricted to 140 chars.
func (o LookupSecurityRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
func (o LookupSecurityRuleResultOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

// The destination address prefixes. CIDR or destination IP ranges.
func (o LookupSecurityRuleResultOutput) DestinationAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []string { return v.DestinationAddressPrefixes }).(pulumi.StringArrayOutput)
}

// The application security group specified as destination.
func (o LookupSecurityRuleResultOutput) DestinationApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []ApplicationSecurityGroupResponse {
		return v.DestinationApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

// The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
func (o LookupSecurityRuleResultOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

// The destination port ranges.
func (o LookupSecurityRuleResultOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

// The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
func (o LookupSecurityRuleResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.Direction }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupSecurityRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupSecurityRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupSecurityRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
func (o LookupSecurityRuleResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) int { return v.Priority }).(pulumi.IntOutput)
}

// Network protocol this rule applies to.
func (o LookupSecurityRuleResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// The provisioning state of the security rule resource.
func (o LookupSecurityRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
func (o LookupSecurityRuleResultOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

// The CIDR or source IP ranges.
func (o LookupSecurityRuleResultOutput) SourceAddressPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []string { return v.SourceAddressPrefixes }).(pulumi.StringArrayOutput)
}

// The application security group specified as source.
func (o LookupSecurityRuleResultOutput) SourceApplicationSecurityGroups() ApplicationSecurityGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []ApplicationSecurityGroupResponse {
		return v.SourceApplicationSecurityGroups
	}).(ApplicationSecurityGroupResponseArrayOutput)
}

// The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
func (o LookupSecurityRuleResultOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

// The source port ranges.
func (o LookupSecurityRuleResultOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

// The type of the resource.
func (o LookupSecurityRuleResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityRuleResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityRuleResultOutput{})
}
