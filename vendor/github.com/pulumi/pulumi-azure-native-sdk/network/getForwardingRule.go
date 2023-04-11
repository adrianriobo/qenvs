// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets properties of a forwarding rule in a DNS forwarding ruleset.
// API Version: 2020-04-01-preview.
func LookupForwardingRule(ctx *pulumi.Context, args *LookupForwardingRuleArgs, opts ...pulumi.InvokeOption) (*LookupForwardingRuleResult, error) {
	var rv LookupForwardingRuleResult
	err := ctx.Invoke("azure-native:network:getForwardingRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupForwardingRuleArgs struct {
	// The name of the DNS forwarding ruleset.
	DnsForwardingRulesetName string `pulumi:"dnsForwardingRulesetName"`
	// The name of the forwarding rule.
	ForwardingRuleName string `pulumi:"forwardingRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes a forwarding rule within a DNS forwarding ruleset.
type LookupForwardingRuleResult struct {
	// The domain name for the forwarding rule.
	DomainName string `pulumi:"domainName"`
	// ETag of the forwarding rule.
	Etag string `pulumi:"etag"`
	// The state of forwarding rule.
	ForwardingRuleState *string `pulumi:"forwardingRuleState"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Metadata attached to the forwarding rule.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The current provisioning state of the forwarding rule. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState string `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// DNS servers to forward the DNS query to.
	TargetDnsServers []TargetDnsServerResponse `pulumi:"targetDnsServers"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupForwardingRuleOutput(ctx *pulumi.Context, args LookupForwardingRuleOutputArgs, opts ...pulumi.InvokeOption) LookupForwardingRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupForwardingRuleResult, error) {
			args := v.(LookupForwardingRuleArgs)
			r, err := LookupForwardingRule(ctx, &args, opts...)
			var s LookupForwardingRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupForwardingRuleResultOutput)
}

type LookupForwardingRuleOutputArgs struct {
	// The name of the DNS forwarding ruleset.
	DnsForwardingRulesetName pulumi.StringInput `pulumi:"dnsForwardingRulesetName"`
	// The name of the forwarding rule.
	ForwardingRuleName pulumi.StringInput `pulumi:"forwardingRuleName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupForwardingRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupForwardingRuleArgs)(nil)).Elem()
}

// Describes a forwarding rule within a DNS forwarding ruleset.
type LookupForwardingRuleResultOutput struct{ *pulumi.OutputState }

func (LookupForwardingRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupForwardingRuleResult)(nil)).Elem()
}

func (o LookupForwardingRuleResultOutput) ToLookupForwardingRuleResultOutput() LookupForwardingRuleResultOutput {
	return o
}

func (o LookupForwardingRuleResultOutput) ToLookupForwardingRuleResultOutputWithContext(ctx context.Context) LookupForwardingRuleResultOutput {
	return o
}

// The domain name for the forwarding rule.
func (o LookupForwardingRuleResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.DomainName }).(pulumi.StringOutput)
}

// ETag of the forwarding rule.
func (o LookupForwardingRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The state of forwarding rule.
func (o LookupForwardingRuleResultOutput) ForwardingRuleState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) *string { return v.ForwardingRuleState }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupForwardingRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Metadata attached to the forwarding rule.
func (o LookupForwardingRuleResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the resource
func (o LookupForwardingRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the forwarding rule. This is a read-only property and any attempt to set this value will be ignored.
func (o LookupForwardingRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupForwardingRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// DNS servers to forward the DNS query to.
func (o LookupForwardingRuleResultOutput) TargetDnsServers() TargetDnsServerResponseArrayOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) []TargetDnsServerResponse { return v.TargetDnsServers }).(TargetDnsServerResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupForwardingRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupForwardingRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupForwardingRuleResultOutput{})
}
