// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Inbound NAT rule of the load balancer.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2019-06-01, 2023-04-01, 2023-05-01, 2023-06-01.
type InboundNatRule struct {
	pulumi.CustomResourceState

	// A reference to backendAddressPool resource.
	BackendAddressPool SubResourceResponsePtrOutput `pulumi:"backendAddressPool"`
	// A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
	BackendIPConfiguration NetworkInterfaceIPConfigurationResponseOutput `pulumi:"backendIPConfiguration"`
	// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
	BackendPort pulumi.IntPtrOutput `pulumi:"backendPort"`
	// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
	EnableFloatingIP pulumi.BoolPtrOutput `pulumi:"enableFloatingIP"`
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset pulumi.BoolPtrOutput `pulumi:"enableTcpReset"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A reference to frontend IP addresses.
	FrontendIPConfiguration SubResourceResponsePtrOutput `pulumi:"frontendIPConfiguration"`
	// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
	FrontendPort pulumi.IntPtrOutput `pulumi:"frontendPort"`
	// The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
	FrontendPortRangeEnd pulumi.IntPtrOutput `pulumi:"frontendPortRangeEnd"`
	// The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
	FrontendPortRangeStart pulumi.IntPtrOutput `pulumi:"frontendPortRangeStart"`
	// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes pulumi.IntPtrOutput `pulumi:"idleTimeoutInMinutes"`
	// The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The reference to the transport protocol used by the load balancing rule.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The provisioning state of the inbound NAT rule resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewInboundNatRule registers a new resource with the given unique name, arguments, and options.
func NewInboundNatRule(ctx *pulumi.Context,
	name string, args *InboundNatRuleArgs, opts ...pulumi.ResourceOption) (*InboundNatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerName == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20170601:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:InboundNatRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:InboundNatRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource InboundNatRule
	err := ctx.RegisterResource("azure-native:network:InboundNatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInboundNatRule gets an existing InboundNatRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInboundNatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InboundNatRuleState, opts ...pulumi.ResourceOption) (*InboundNatRule, error) {
	var resource InboundNatRule
	err := ctx.ReadResource("azure-native:network:InboundNatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InboundNatRule resources.
type inboundNatRuleState struct {
}

type InboundNatRuleState struct {
}

func (InboundNatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundNatRuleState)(nil)).Elem()
}

type inboundNatRuleArgs struct {
	// A reference to backendAddressPool resource.
	BackendAddressPool *SubResource `pulumi:"backendAddressPool"`
	// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
	BackendPort *int `pulumi:"backendPort"`
	// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
	EnableFloatingIP *bool `pulumi:"enableFloatingIP"`
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `pulumi:"enableTcpReset"`
	// A reference to frontend IP addresses.
	FrontendIPConfiguration *SubResource `pulumi:"frontendIPConfiguration"`
	// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
	FrontendPort *int `pulumi:"frontendPort"`
	// The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
	FrontendPortRangeEnd *int `pulumi:"frontendPortRangeEnd"`
	// The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
	FrontendPortRangeStart *int `pulumi:"frontendPortRangeStart"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// The name of the inbound NAT rule.
	InboundNatRuleName *string `pulumi:"inboundNatRuleName"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The reference to the transport protocol used by the load balancing rule.
	Protocol *string `pulumi:"protocol"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a InboundNatRule resource.
type InboundNatRuleArgs struct {
	// A reference to backendAddressPool resource.
	BackendAddressPool SubResourcePtrInput
	// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
	BackendPort pulumi.IntPtrInput
	// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
	EnableFloatingIP pulumi.BoolPtrInput
	// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
	EnableTcpReset pulumi.BoolPtrInput
	// A reference to frontend IP addresses.
	FrontendIPConfiguration SubResourcePtrInput
	// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
	FrontendPort pulumi.IntPtrInput
	// The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
	FrontendPortRangeEnd pulumi.IntPtrInput
	// The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
	FrontendPortRangeStart pulumi.IntPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// The name of the inbound NAT rule.
	InboundNatRuleName pulumi.StringPtrInput
	// The name of the load balancer.
	LoadBalancerName pulumi.StringInput
	// The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The reference to the transport protocol used by the load balancing rule.
	Protocol pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (InboundNatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inboundNatRuleArgs)(nil)).Elem()
}

type InboundNatRuleInput interface {
	pulumi.Input

	ToInboundNatRuleOutput() InboundNatRuleOutput
	ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput
}

func (*InboundNatRule) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundNatRule)(nil)).Elem()
}

func (i *InboundNatRule) ToInboundNatRuleOutput() InboundNatRuleOutput {
	return i.ToInboundNatRuleOutputWithContext(context.Background())
}

func (i *InboundNatRule) ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundNatRuleOutput)
}

type InboundNatRuleOutput struct{ *pulumi.OutputState }

func (InboundNatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InboundNatRule)(nil)).Elem()
}

func (o InboundNatRuleOutput) ToInboundNatRuleOutput() InboundNatRuleOutput {
	return o
}

func (o InboundNatRuleOutput) ToInboundNatRuleOutputWithContext(ctx context.Context) InboundNatRuleOutput {
	return o
}

// A reference to backendAddressPool resource.
func (o InboundNatRuleOutput) BackendAddressPool() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *InboundNatRule) SubResourceResponsePtrOutput { return v.BackendAddressPool }).(SubResourceResponsePtrOutput)
}

// A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
func (o InboundNatRuleOutput) BackendIPConfiguration() NetworkInterfaceIPConfigurationResponseOutput {
	return o.ApplyT(func(v *InboundNatRule) NetworkInterfaceIPConfigurationResponseOutput { return v.BackendIPConfiguration }).(NetworkInterfaceIPConfigurationResponseOutput)
}

// The port used for the internal endpoint. Acceptable values range from 1 to 65535.
func (o InboundNatRuleOutput) BackendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.IntPtrOutput { return v.BackendPort }).(pulumi.IntPtrOutput)
}

// Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
func (o InboundNatRuleOutput) EnableFloatingIP() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.BoolPtrOutput { return v.EnableFloatingIP }).(pulumi.BoolPtrOutput)
}

// Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
func (o InboundNatRuleOutput) EnableTcpReset() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.BoolPtrOutput { return v.EnableTcpReset }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o InboundNatRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A reference to frontend IP addresses.
func (o InboundNatRuleOutput) FrontendIPConfiguration() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *InboundNatRule) SubResourceResponsePtrOutput { return v.FrontendIPConfiguration }).(SubResourceResponsePtrOutput)
}

// The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
func (o InboundNatRuleOutput) FrontendPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.IntPtrOutput { return v.FrontendPort }).(pulumi.IntPtrOutput)
}

// The port range end for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeStart. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
func (o InboundNatRuleOutput) FrontendPortRangeEnd() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.IntPtrOutput { return v.FrontendPortRangeEnd }).(pulumi.IntPtrOutput)
}

// The port range start for the external endpoint. This property is used together with BackendAddressPool and FrontendPortRangeEnd. Individual inbound NAT rule port mappings will be created for each backend address from BackendAddressPool. Acceptable values range from 1 to 65534.
func (o InboundNatRuleOutput) FrontendPortRangeStart() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.IntPtrOutput { return v.FrontendPortRangeStart }).(pulumi.IntPtrOutput)
}

// The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
func (o InboundNatRuleOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.IntPtrOutput { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

// The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
func (o InboundNatRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The reference to the transport protocol used by the load balancing rule.
func (o InboundNatRuleOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The provisioning state of the inbound NAT rule resource.
func (o InboundNatRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Type of the resource.
func (o InboundNatRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InboundNatRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InboundNatRuleOutput{})
}
