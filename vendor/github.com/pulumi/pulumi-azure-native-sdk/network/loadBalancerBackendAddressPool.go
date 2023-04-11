// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Pool of backend IP addresses.
// API Version: 2020-11-01.
type LoadBalancerBackendAddressPool struct {
	pulumi.CustomResourceState

	// An array of references to IP addresses defined in network interfaces.
	BackendIPConfigurations NetworkInterfaceIPConfigurationResponseArrayOutput `pulumi:"backendIPConfigurations"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// An array of backend addresses.
	LoadBalancerBackendAddresses LoadBalancerBackendAddressResponseArrayOutput `pulumi:"loadBalancerBackendAddresses"`
	// An array of references to load balancing rules that use this backend address pool.
	LoadBalancingRules SubResourceResponseArrayOutput `pulumi:"loadBalancingRules"`
	// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// A reference to an outbound rule that uses this backend address pool.
	OutboundRule SubResourceResponseOutput `pulumi:"outboundRule"`
	// An array of references to outbound rules that use this backend address pool.
	OutboundRules SubResourceResponseArrayOutput `pulumi:"outboundRules"`
	// The provisioning state of the backend address pool resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLoadBalancerBackendAddressPool registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerBackendAddressPool(ctx *pulumi.Context,
	name string, args *LoadBalancerBackendAddressPoolArgs, opts ...pulumi.ResourceOption) (*LoadBalancerBackendAddressPool, error) {
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
			Type: pulumi.String("azure-native:network/v20200401:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:LoadBalancerBackendAddressPool"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:LoadBalancerBackendAddressPool"),
		},
	})
	opts = append(opts, aliases)
	var resource LoadBalancerBackendAddressPool
	err := ctx.RegisterResource("azure-native:network:LoadBalancerBackendAddressPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerBackendAddressPool gets an existing LoadBalancerBackendAddressPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerBackendAddressPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerBackendAddressPoolState, opts ...pulumi.ResourceOption) (*LoadBalancerBackendAddressPool, error) {
	var resource LoadBalancerBackendAddressPool
	err := ctx.ReadResource("azure-native:network:LoadBalancerBackendAddressPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancerBackendAddressPool resources.
type loadBalancerBackendAddressPoolState struct {
}

type LoadBalancerBackendAddressPoolState struct {
}

func (LoadBalancerBackendAddressPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendAddressPoolState)(nil)).Elem()
}

type loadBalancerBackendAddressPoolArgs struct {
	// The name of the backend address pool.
	BackendAddressPoolName *string `pulumi:"backendAddressPoolName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// An array of backend addresses.
	LoadBalancerBackendAddresses []LoadBalancerBackendAddress `pulumi:"loadBalancerBackendAddresses"`
	// The name of the load balancer.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LoadBalancerBackendAddressPool resource.
type LoadBalancerBackendAddressPoolArgs struct {
	// The name of the backend address pool.
	BackendAddressPoolName pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// An array of backend addresses.
	LoadBalancerBackendAddresses LoadBalancerBackendAddressArrayInput
	// The name of the load balancer.
	LoadBalancerName pulumi.StringInput
	// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (LoadBalancerBackendAddressPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendAddressPoolArgs)(nil)).Elem()
}

type LoadBalancerBackendAddressPoolInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput
	ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput
}

func (*LoadBalancerBackendAddressPool) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerBackendAddressPool)(nil)).Elem()
}

func (i *LoadBalancerBackendAddressPool) ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput {
	return i.ToLoadBalancerBackendAddressPoolOutputWithContext(context.Background())
}

func (i *LoadBalancerBackendAddressPool) ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolOutput)
}

type LoadBalancerBackendAddressPoolOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerBackendAddressPool)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolOutput) ToLoadBalancerBackendAddressPoolOutput() LoadBalancerBackendAddressPoolOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolOutput) ToLoadBalancerBackendAddressPoolOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolOutput {
	return o
}

// An array of references to IP addresses defined in network interfaces.
func (o LoadBalancerBackendAddressPoolOutput) BackendIPConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) NetworkInterfaceIPConfigurationResponseArrayOutput {
		return v.BackendIPConfigurations
	}).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LoadBalancerBackendAddressPoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// An array of backend addresses.
func (o LoadBalancerBackendAddressPoolOutput) LoadBalancerBackendAddresses() LoadBalancerBackendAddressResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) LoadBalancerBackendAddressResponseArrayOutput {
		return v.LoadBalancerBackendAddresses
	}).(LoadBalancerBackendAddressResponseArrayOutput)
}

// An array of references to load balancing rules that use this backend address pool.
func (o LoadBalancerBackendAddressPoolOutput) LoadBalancingRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponseArrayOutput { return v.LoadBalancingRules }).(SubResourceResponseArrayOutput)
}

// The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
func (o LoadBalancerBackendAddressPoolOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// A reference to an outbound rule that uses this backend address pool.
func (o LoadBalancerBackendAddressPoolOutput) OutboundRule() SubResourceResponseOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponseOutput { return v.OutboundRule }).(SubResourceResponseOutput)
}

// An array of references to outbound rules that use this backend address pool.
func (o LoadBalancerBackendAddressPoolOutput) OutboundRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) SubResourceResponseArrayOutput { return v.OutboundRules }).(SubResourceResponseArrayOutput)
}

// The provisioning state of the backend address pool resource.
func (o LoadBalancerBackendAddressPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Type of the resource.
func (o LoadBalancerBackendAddressPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancerBackendAddressPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolOutput{})
}
