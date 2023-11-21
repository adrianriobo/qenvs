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

// NetworkSecurityGroup resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2015-05-01-preview, 2016-03-30, 2019-06-01, 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01.
type NetworkSecurityGroup struct {
	pulumi.CustomResourceState

	// The default security rules of network security group.
	DefaultSecurityRules SecurityRuleResponseArrayOutput `pulumi:"defaultSecurityRules"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A collection of references to flow log resources.
	FlowLogs FlowLogResponseArrayOutput `pulumi:"flowLogs"`
	// When enabled, flows created from Network Security Group connections will be re-evaluated when rules are updates. Initial enablement will trigger re-evaluation.
	FlushConnection pulumi.BoolPtrOutput `pulumi:"flushConnection"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// A collection of references to network interfaces.
	NetworkInterfaces NetworkInterfaceResponseArrayOutput `pulumi:"networkInterfaces"`
	// The provisioning state of the network security group resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the network security group resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// A collection of security rules of the network security group.
	SecurityRules SecurityRuleResponseArrayOutput `pulumi:"securityRules"`
	// A collection of references to subnets.
	Subnets SubnetResponseArrayOutput `pulumi:"subnets"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityGroup(ctx *pulumi.Context,
	name string, args *NetworkSecurityGroupArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20150501preview:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:NetworkSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:NetworkSecurityGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkSecurityGroup
	err := ctx.RegisterResource("azure-native:network:NetworkSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityGroup gets an existing NetworkSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityGroupState, opts ...pulumi.ResourceOption) (*NetworkSecurityGroup, error) {
	var resource NetworkSecurityGroup
	err := ctx.ReadResource("azure-native:network:NetworkSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityGroup resources.
type networkSecurityGroupState struct {
}

type NetworkSecurityGroupState struct {
}

func (NetworkSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityGroupState)(nil)).Elem()
}

type networkSecurityGroupArgs struct {
	// When enabled, flows created from Network Security Group connections will be re-evaluated when rules are updates. Initial enablement will trigger re-evaluation.
	FlushConnection *bool `pulumi:"flushConnection"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the network security group.
	NetworkSecurityGroupName *string `pulumi:"networkSecurityGroupName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A collection of security rules of the network security group.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	SecurityRules []SecurityRuleType `pulumi:"securityRules"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkSecurityGroup resource.
type NetworkSecurityGroupArgs struct {
	// When enabled, flows created from Network Security Group connections will be re-evaluated when rules are updates. Initial enablement will trigger re-evaluation.
	FlushConnection pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the network security group.
	NetworkSecurityGroupName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// A collection of security rules of the network security group.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	SecurityRules SecurityRuleTypeArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityGroupArgs)(nil)).Elem()
}

type NetworkSecurityGroupInput interface {
	pulumi.Input

	ToNetworkSecurityGroupOutput() NetworkSecurityGroupOutput
	ToNetworkSecurityGroupOutputWithContext(ctx context.Context) NetworkSecurityGroupOutput
}

func (*NetworkSecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroup)(nil)).Elem()
}

func (i *NetworkSecurityGroup) ToNetworkSecurityGroupOutput() NetworkSecurityGroupOutput {
	return i.ToNetworkSecurityGroupOutputWithContext(context.Background())
}

func (i *NetworkSecurityGroup) ToNetworkSecurityGroupOutputWithContext(ctx context.Context) NetworkSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupOutput)
}

func (i *NetworkSecurityGroup) ToOutput(ctx context.Context) pulumix.Output[*NetworkSecurityGroup] {
	return pulumix.Output[*NetworkSecurityGroup]{
		OutputState: i.ToNetworkSecurityGroupOutputWithContext(ctx).OutputState,
	}
}

type NetworkSecurityGroupOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityGroup)(nil)).Elem()
}

func (o NetworkSecurityGroupOutput) ToNetworkSecurityGroupOutput() NetworkSecurityGroupOutput {
	return o
}

func (o NetworkSecurityGroupOutput) ToNetworkSecurityGroupOutputWithContext(ctx context.Context) NetworkSecurityGroupOutput {
	return o
}

func (o NetworkSecurityGroupOutput) ToOutput(ctx context.Context) pulumix.Output[*NetworkSecurityGroup] {
	return pulumix.Output[*NetworkSecurityGroup]{
		OutputState: o.OutputState,
	}
}

// The default security rules of network security group.
func (o NetworkSecurityGroupOutput) DefaultSecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) SecurityRuleResponseArrayOutput { return v.DefaultSecurityRules }).(SecurityRuleResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o NetworkSecurityGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A collection of references to flow log resources.
func (o NetworkSecurityGroupOutput) FlowLogs() FlowLogResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) FlowLogResponseArrayOutput { return v.FlowLogs }).(FlowLogResponseArrayOutput)
}

// When enabled, flows created from Network Security Group connections will be re-evaluated when rules are updates. Initial enablement will trigger re-evaluation.
func (o NetworkSecurityGroupOutput) FlushConnection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.BoolPtrOutput { return v.FlushConnection }).(pulumi.BoolPtrOutput)
}

// Resource location.
func (o NetworkSecurityGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o NetworkSecurityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A collection of references to network interfaces.
func (o NetworkSecurityGroupOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) NetworkInterfaceResponseArrayOutput { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

// The provisioning state of the network security group resource.
func (o NetworkSecurityGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the network security group resource.
func (o NetworkSecurityGroupOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// A collection of security rules of the network security group.
func (o NetworkSecurityGroupOutput) SecurityRules() SecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) SecurityRuleResponseArrayOutput { return v.SecurityRules }).(SecurityRuleResponseArrayOutput)
}

// A collection of references to subnets.
func (o NetworkSecurityGroupOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) SubnetResponseArrayOutput { return v.Subnets }).(SubnetResponseArrayOutput)
}

// Resource tags.
func (o NetworkSecurityGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o NetworkSecurityGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkSecurityGroupOutput{})
}
