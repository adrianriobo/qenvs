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

// LoadBalancer resource.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2020-11-01.
//
// Other available API versions: 2015-05-01-preview, 2018-06-01, 2019-06-01, 2019-08-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01.
type LoadBalancer struct {
	pulumi.CustomResourceState

	// Collection of backend address pools used by a load balancer.
	BackendAddressPools BackendAddressPoolResponseArrayOutput `pulumi:"backendAddressPools"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The extended location of the load balancer.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// Object representing the frontend IPs to be used for the load balancer.
	FrontendIPConfigurations FrontendIPConfigurationResponseArrayOutput `pulumi:"frontendIPConfigurations"`
	// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound NAT rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools InboundNatPoolResponseArrayOutput `pulumi:"inboundNatPools"`
	// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatRules InboundNatRuleResponseArrayOutput `pulumi:"inboundNatRules"`
	// Object collection representing the load balancing rules Gets the provisioning.
	LoadBalancingRules LoadBalancingRuleResponseArrayOutput `pulumi:"loadBalancingRules"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The outbound rules.
	OutboundRules OutboundRuleResponseArrayOutput `pulumi:"outboundRules"`
	// Collection of probe objects used in the load balancer.
	Probes ProbeResponseArrayOutput `pulumi:"probes"`
	// The provisioning state of the load balancer resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the load balancer resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// The load balancer SKU.
	Sku LoadBalancerSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20150501preview:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:LoadBalancer"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230901:LoadBalancer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LoadBalancer
	err := ctx.RegisterResource("azure-native:network:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("azure-native:network:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
}

type LoadBalancerState struct {
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// Collection of backend address pools used by a load balancer.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	BackendAddressPools []BackendAddressPool `pulumi:"backendAddressPools"`
	// The extended location of the load balancer.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// Object representing the frontend IPs to be used for the load balancer.
	FrontendIPConfigurations []FrontendIPConfiguration `pulumi:"frontendIPConfigurations"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound NAT rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools []InboundNatPool `pulumi:"inboundNatPools"`
	// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	InboundNatRules []InboundNatRuleType `pulumi:"inboundNatRules"`
	// The name of the load balancer.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// Object collection representing the load balancing rules Gets the provisioning.
	LoadBalancingRules []LoadBalancingRule `pulumi:"loadBalancingRules"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The outbound rules.
	OutboundRules []OutboundRule `pulumi:"outboundRules"`
	// Collection of probe objects used in the load balancer.
	Probes []Probe `pulumi:"probes"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The load balancer SKU.
	Sku *LoadBalancerSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// Collection of backend address pools used by a load balancer.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	BackendAddressPools BackendAddressPoolArrayInput
	// The extended location of the load balancer.
	ExtendedLocation ExtendedLocationPtrInput
	// Object representing the frontend IPs to be used for the load balancer.
	FrontendIPConfigurations FrontendIPConfigurationArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound NAT rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools InboundNatPoolArrayInput
	// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	InboundNatRules InboundNatRuleTypeArrayInput
	// The name of the load balancer.
	LoadBalancerName pulumi.StringPtrInput
	// Object collection representing the load balancing rules Gets the provisioning.
	LoadBalancingRules LoadBalancingRuleArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The outbound rules.
	OutboundRules OutboundRuleArrayInput
	// Collection of probe objects used in the load balancer.
	Probes ProbeArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The load balancer SKU.
	Sku LoadBalancerSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}

type LoadBalancerInput interface {
	pulumi.Input

	ToLoadBalancerOutput() LoadBalancerOutput
	ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput
}

func (*LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (i *LoadBalancer) ToLoadBalancerOutput() LoadBalancerOutput {
	return i.ToLoadBalancerOutputWithContext(context.Background())
}

func (i *LoadBalancer) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerOutput)
}

type LoadBalancerOutput struct{ *pulumi.OutputState }

func (LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerOutput) ToLoadBalancerOutput() LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return o
}

// Collection of backend address pools used by a load balancer.
func (o LoadBalancerOutput) BackendAddressPools() BackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) BackendAddressPoolResponseArrayOutput { return v.BackendAddressPools }).(BackendAddressPoolResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LoadBalancerOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the load balancer.
func (o LoadBalancerOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *LoadBalancer) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Object representing the frontend IPs to be used for the load balancer.
func (o LoadBalancerOutput) FrontendIPConfigurations() FrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) FrontendIPConfigurationResponseArrayOutput { return v.FrontendIPConfigurations }).(FrontendIPConfigurationResponseArrayOutput)
}

// Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound NAT rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
func (o LoadBalancerOutput) InboundNatPools() InboundNatPoolResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) InboundNatPoolResponseArrayOutput { return v.InboundNatPools }).(InboundNatPoolResponseArrayOutput)
}

// Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
func (o LoadBalancerOutput) InboundNatRules() InboundNatRuleResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) InboundNatRuleResponseArrayOutput { return v.InboundNatRules }).(InboundNatRuleResponseArrayOutput)
}

// Object collection representing the load balancing rules Gets the provisioning.
func (o LoadBalancerOutput) LoadBalancingRules() LoadBalancingRuleResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancingRuleResponseArrayOutput { return v.LoadBalancingRules }).(LoadBalancingRuleResponseArrayOutput)
}

// Resource location.
func (o LoadBalancerOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LoadBalancerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The outbound rules.
func (o LoadBalancerOutput) OutboundRules() OutboundRuleResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) OutboundRuleResponseArrayOutput { return v.OutboundRules }).(OutboundRuleResponseArrayOutput)
}

// Collection of probe objects used in the load balancer.
func (o LoadBalancerOutput) Probes() ProbeResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) ProbeResponseArrayOutput { return v.Probes }).(ProbeResponseArrayOutput)
}

// The provisioning state of the load balancer resource.
func (o LoadBalancerOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the load balancer resource.
func (o LoadBalancerOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The load balancer SKU.
func (o LoadBalancerOutput) Sku() LoadBalancerSkuResponsePtrOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerSkuResponsePtrOutput { return v.Sku }).(LoadBalancerSkuResponsePtrOutput)
}

// Resource tags.
func (o LoadBalancerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LoadBalancerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerOutput{})
}
