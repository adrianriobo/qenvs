// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Rule Group resource.
// Azure REST API version: 2020-04-01. Prior API version in Azure Native 1.x: 2020-04-01
type FirewallPolicyRuleGroup struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Priority of the Firewall Policy Rule Group resource.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The provisioning state of the firewall policy rule group resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Group of Firewall Policy rules.
	Rules pulumi.ArrayOutput `pulumi:"rules"`
	// Rule Group type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallPolicyRuleGroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicyRuleGroup(ctx *pulumi.Context,
	name string, args *FirewallPolicyRuleGroupArgs, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallPolicyName == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicyName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20190601:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:FirewallPolicyRuleGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:FirewallPolicyRuleGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FirewallPolicyRuleGroup
	err := ctx.RegisterResource("azure-native:network:FirewallPolicyRuleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicyRuleGroup gets an existing FirewallPolicyRuleGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicyRuleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyRuleGroupState, opts ...pulumi.ResourceOption) (*FirewallPolicyRuleGroup, error) {
	var resource FirewallPolicyRuleGroup
	err := ctx.ReadResource("azure-native:network:FirewallPolicyRuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicyRuleGroup resources.
type firewallPolicyRuleGroupState struct {
}

type FirewallPolicyRuleGroupState struct {
}

func (FirewallPolicyRuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleGroupState)(nil)).Elem()
}

type firewallPolicyRuleGroupArgs struct {
	// The name of the Firewall Policy.
	FirewallPolicyName string `pulumi:"firewallPolicyName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Priority of the Firewall Policy Rule Group resource.
	Priority *int `pulumi:"priority"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the FirewallPolicyRuleGroup.
	RuleGroupName *string `pulumi:"ruleGroupName"`
	// Group of Firewall Policy rules.
	Rules []interface{} `pulumi:"rules"`
}

// The set of arguments for constructing a FirewallPolicyRuleGroup resource.
type FirewallPolicyRuleGroupArgs struct {
	// The name of the Firewall Policy.
	FirewallPolicyName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Priority of the Firewall Policy Rule Group resource.
	Priority pulumi.IntPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the FirewallPolicyRuleGroup.
	RuleGroupName pulumi.StringPtrInput
	// Group of Firewall Policy rules.
	Rules pulumi.ArrayInput
}

func (FirewallPolicyRuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyRuleGroupArgs)(nil)).Elem()
}

type FirewallPolicyRuleGroupInput interface {
	pulumi.Input

	ToFirewallPolicyRuleGroupOutput() FirewallPolicyRuleGroupOutput
	ToFirewallPolicyRuleGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleGroupOutput
}

func (*FirewallPolicyRuleGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleGroup)(nil)).Elem()
}

func (i *FirewallPolicyRuleGroup) ToFirewallPolicyRuleGroupOutput() FirewallPolicyRuleGroupOutput {
	return i.ToFirewallPolicyRuleGroupOutputWithContext(context.Background())
}

func (i *FirewallPolicyRuleGroup) ToFirewallPolicyRuleGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyRuleGroupOutput)
}

func (i *FirewallPolicyRuleGroup) ToOutput(ctx context.Context) pulumix.Output[*FirewallPolicyRuleGroup] {
	return pulumix.Output[*FirewallPolicyRuleGroup]{
		OutputState: i.ToFirewallPolicyRuleGroupOutputWithContext(ctx).OutputState,
	}
}

type FirewallPolicyRuleGroupOutput struct{ *pulumi.OutputState }

func (FirewallPolicyRuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyRuleGroup)(nil)).Elem()
}

func (o FirewallPolicyRuleGroupOutput) ToFirewallPolicyRuleGroupOutput() FirewallPolicyRuleGroupOutput {
	return o
}

func (o FirewallPolicyRuleGroupOutput) ToFirewallPolicyRuleGroupOutputWithContext(ctx context.Context) FirewallPolicyRuleGroupOutput {
	return o
}

func (o FirewallPolicyRuleGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*FirewallPolicyRuleGroup] {
	return pulumix.Output[*FirewallPolicyRuleGroup]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o FirewallPolicyRuleGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o FirewallPolicyRuleGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Priority of the Firewall Policy Rule Group resource.
func (o FirewallPolicyRuleGroupOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleGroup) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The provisioning state of the firewall policy rule group resource.
func (o FirewallPolicyRuleGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Group of Firewall Policy rules.
func (o FirewallPolicyRuleGroupOutput) Rules() pulumi.ArrayOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleGroup) pulumi.ArrayOutput { return v.Rules }).(pulumi.ArrayOutput)
}

// Rule Group type.
func (o FirewallPolicyRuleGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyRuleGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyRuleGroupOutput{})
}
