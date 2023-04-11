// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Service End point policy resource.
// API Version: 2020-11-01.
type ServiceEndpointPolicy struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Kind of service endpoint policy. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the service endpoint policy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the service endpoint policy resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// A collection of service endpoint policy definitions of the service endpoint policy.
	ServiceEndpointPolicyDefinitions ServiceEndpointPolicyDefinitionResponseArrayOutput `pulumi:"serviceEndpointPolicyDefinitions"`
	// A collection of references to subnets.
	Subnets SubnetResponseArrayOutput `pulumi:"subnets"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceEndpointPolicy registers a new resource with the given unique name, arguments, and options.
func NewServiceEndpointPolicy(ctx *pulumi.Context,
	name string, args *ServiceEndpointPolicyArgs, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180701:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ServiceEndpointPolicy"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ServiceEndpointPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceEndpointPolicy
	err := ctx.RegisterResource("azure-native:network:ServiceEndpointPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceEndpointPolicy gets an existing ServiceEndpointPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceEndpointPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceEndpointPolicyState, opts ...pulumi.ResourceOption) (*ServiceEndpointPolicy, error) {
	var resource ServiceEndpointPolicy
	err := ctx.ReadResource("azure-native:network:ServiceEndpointPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceEndpointPolicy resources.
type serviceEndpointPolicyState struct {
}

type ServiceEndpointPolicyState struct {
}

func (ServiceEndpointPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyState)(nil)).Elem()
}

type serviceEndpointPolicyArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A collection of service endpoint policy definitions of the service endpoint policy.
	ServiceEndpointPolicyDefinitions []ServiceEndpointPolicyDefinitionType `pulumi:"serviceEndpointPolicyDefinitions"`
	// The name of the service endpoint policy.
	ServiceEndpointPolicyName *string `pulumi:"serviceEndpointPolicyName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceEndpointPolicy resource.
type ServiceEndpointPolicyArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// A collection of service endpoint policy definitions of the service endpoint policy.
	ServiceEndpointPolicyDefinitions ServiceEndpointPolicyDefinitionTypeArrayInput
	// The name of the service endpoint policy.
	ServiceEndpointPolicyName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ServiceEndpointPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceEndpointPolicyArgs)(nil)).Elem()
}

type ServiceEndpointPolicyInput interface {
	pulumi.Input

	ToServiceEndpointPolicyOutput() ServiceEndpointPolicyOutput
	ToServiceEndpointPolicyOutputWithContext(ctx context.Context) ServiceEndpointPolicyOutput
}

func (*ServiceEndpointPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointPolicy)(nil)).Elem()
}

func (i *ServiceEndpointPolicy) ToServiceEndpointPolicyOutput() ServiceEndpointPolicyOutput {
	return i.ToServiceEndpointPolicyOutputWithContext(context.Background())
}

func (i *ServiceEndpointPolicy) ToServiceEndpointPolicyOutputWithContext(ctx context.Context) ServiceEndpointPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceEndpointPolicyOutput)
}

type ServiceEndpointPolicyOutput struct{ *pulumi.OutputState }

func (ServiceEndpointPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceEndpointPolicy)(nil)).Elem()
}

func (o ServiceEndpointPolicyOutput) ToServiceEndpointPolicyOutput() ServiceEndpointPolicyOutput {
	return o
}

func (o ServiceEndpointPolicyOutput) ToServiceEndpointPolicyOutputWithContext(ctx context.Context) ServiceEndpointPolicyOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o ServiceEndpointPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Kind of service endpoint policy. This is metadata used for the Azure portal experience.
func (o ServiceEndpointPolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o ServiceEndpointPolicyOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ServiceEndpointPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the service endpoint policy resource.
func (o ServiceEndpointPolicyOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the service endpoint policy resource.
func (o ServiceEndpointPolicyOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// A collection of service endpoint policy definitions of the service endpoint policy.
func (o ServiceEndpointPolicyOutput) ServiceEndpointPolicyDefinitions() ServiceEndpointPolicyDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) ServiceEndpointPolicyDefinitionResponseArrayOutput {
		return v.ServiceEndpointPolicyDefinitions
	}).(ServiceEndpointPolicyDefinitionResponseArrayOutput)
}

// A collection of references to subnets.
func (o ServiceEndpointPolicyOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) SubnetResponseArrayOutput { return v.Subnets }).(SubnetResponseArrayOutput)
}

// Resource tags.
func (o ServiceEndpointPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ServiceEndpointPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceEndpointPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceEndpointPolicyOutput{})
}
