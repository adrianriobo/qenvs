// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a network manager security configuration admin rule.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2021-02-01-preview, 2021-05-01-preview, 2023-04-01, 2023-05-01, 2023-06-01.
func LookupAdminRule(ctx *pulumi.Context, args *LookupAdminRuleArgs, opts ...pulumi.InvokeOption) (*LookupAdminRuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAdminRuleResult
	err := ctx.Invoke("azure-native:network:getAdminRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdminRuleArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	// The name of the rule.
	RuleName string `pulumi:"ruleName"`
}

// Network admin rule.
type LookupAdminRuleResult struct {
	// Indicates the access allowed for this particular rule
	Access string `pulumi:"access"`
	// A description for this rule. Restricted to 140 chars.
	Description *string `pulumi:"description"`
	// The destination port ranges.
	DestinationPortRanges []string `pulumi:"destinationPortRanges"`
	// The destination address prefixes. CIDR or destination IP ranges.
	Destinations []AddressPrefixItemResponse `pulumi:"destinations"`
	// Indicates if the traffic matched against the rule in inbound or outbound.
	Direction string `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Whether the rule is custom or default.
	// Expected value is 'Custom'.
	Kind string `pulumi:"kind"`
	// Resource name.
	Name string `pulumi:"name"`
	// The priority of the rule. The value can be between 1 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority int `pulumi:"priority"`
	// Network protocol this rule applies to.
	Protocol string `pulumi:"protocol"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The source port ranges.
	SourcePortRanges []string `pulumi:"sourcePortRanges"`
	// The CIDR or source IP ranges.
	Sources []AddressPrefixItemResponse `pulumi:"sources"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupAdminRuleOutput(ctx *pulumi.Context, args LookupAdminRuleOutputArgs, opts ...pulumi.InvokeOption) LookupAdminRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAdminRuleResult, error) {
			args := v.(LookupAdminRuleArgs)
			r, err := LookupAdminRule(ctx, &args, opts...)
			var s LookupAdminRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAdminRuleResultOutput)
}

type LookupAdminRuleOutputArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName pulumi.StringInput `pulumi:"ruleCollectionName"`
	// The name of the rule.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupAdminRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdminRuleArgs)(nil)).Elem()
}

// Network admin rule.
type LookupAdminRuleResultOutput struct{ *pulumi.OutputState }

func (LookupAdminRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdminRuleResult)(nil)).Elem()
}

func (o LookupAdminRuleResultOutput) ToLookupAdminRuleResultOutput() LookupAdminRuleResultOutput {
	return o
}

func (o LookupAdminRuleResultOutput) ToLookupAdminRuleResultOutputWithContext(ctx context.Context) LookupAdminRuleResultOutput {
	return o
}

// Indicates the access allowed for this particular rule
func (o LookupAdminRuleResultOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Access }).(pulumi.StringOutput)
}

// A description for this rule. Restricted to 140 chars.
func (o LookupAdminRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination port ranges.
func (o LookupAdminRuleResultOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) []string { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

// The destination address prefixes. CIDR or destination IP ranges.
func (o LookupAdminRuleResultOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) []AddressPrefixItemResponse { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

// Indicates if the traffic matched against the rule in inbound or outbound.
func (o LookupAdminRuleResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Direction }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupAdminRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupAdminRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether the rule is custom or default.
// Expected value is 'Custom'.
func (o LookupAdminRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupAdminRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The priority of the rule. The value can be between 1 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
func (o LookupAdminRuleResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) int { return v.Priority }).(pulumi.IntOutput)
}

// Network protocol this rule applies to.
func (o LookupAdminRuleResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupAdminRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o LookupAdminRuleResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The source port ranges.
func (o LookupAdminRuleResultOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) []string { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

// The CIDR or source IP ranges.
func (o LookupAdminRuleResultOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) []AddressPrefixItemResponse { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
}

// The system metadata related to this resource.
func (o LookupAdminRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupAdminRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdminRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdminRuleResultOutput{})
}
