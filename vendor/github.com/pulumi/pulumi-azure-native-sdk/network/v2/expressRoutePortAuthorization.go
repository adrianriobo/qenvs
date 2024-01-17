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

// ExpressRoutePort Authorization resource definition.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2022-01-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01.
type ExpressRoutePortAuthorization struct {
	pulumi.CustomResourceState

	// The authorization key.
	AuthorizationKey pulumi.StringOutput `pulumi:"authorizationKey"`
	// The authorization use status.
	AuthorizationUseStatus pulumi.StringOutput `pulumi:"authorizationUseStatus"`
	// The reference to the ExpressRoute circuit resource using the authorization.
	CircuitResourceUri pulumi.StringOutput `pulumi:"circuitResourceUri"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The provisioning state of the authorization resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewExpressRoutePortAuthorization registers a new resource with the given unique name, arguments, and options.
func NewExpressRoutePortAuthorization(ctx *pulumi.Context,
	name string, args *ExpressRoutePortAuthorizationArgs, opts ...pulumi.ResourceOption) (*ExpressRoutePortAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExpressRoutePortName == nil {
		return nil, errors.New("invalid value for required argument 'ExpressRoutePortName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:ExpressRoutePortAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230901:ExpressRoutePortAuthorization"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ExpressRoutePortAuthorization
	err := ctx.RegisterResource("azure-native:network:ExpressRoutePortAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExpressRoutePortAuthorization gets an existing ExpressRoutePortAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExpressRoutePortAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRoutePortAuthorizationState, opts ...pulumi.ResourceOption) (*ExpressRoutePortAuthorization, error) {
	var resource ExpressRoutePortAuthorization
	err := ctx.ReadResource("azure-native:network:ExpressRoutePortAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExpressRoutePortAuthorization resources.
type expressRoutePortAuthorizationState struct {
}

type ExpressRoutePortAuthorizationState struct {
}

func (ExpressRoutePortAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRoutePortAuthorizationState)(nil)).Elem()
}

type expressRoutePortAuthorizationArgs struct {
	// The name of the authorization.
	AuthorizationName *string `pulumi:"authorizationName"`
	// The name of the express route port.
	ExpressRoutePortName string `pulumi:"expressRoutePortName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ExpressRoutePortAuthorization resource.
type ExpressRoutePortAuthorizationArgs struct {
	// The name of the authorization.
	AuthorizationName pulumi.StringPtrInput
	// The name of the express route port.
	ExpressRoutePortName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ExpressRoutePortAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRoutePortAuthorizationArgs)(nil)).Elem()
}

type ExpressRoutePortAuthorizationInput interface {
	pulumi.Input

	ToExpressRoutePortAuthorizationOutput() ExpressRoutePortAuthorizationOutput
	ToExpressRoutePortAuthorizationOutputWithContext(ctx context.Context) ExpressRoutePortAuthorizationOutput
}

func (*ExpressRoutePortAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePortAuthorization)(nil)).Elem()
}

func (i *ExpressRoutePortAuthorization) ToExpressRoutePortAuthorizationOutput() ExpressRoutePortAuthorizationOutput {
	return i.ToExpressRoutePortAuthorizationOutputWithContext(context.Background())
}

func (i *ExpressRoutePortAuthorization) ToExpressRoutePortAuthorizationOutputWithContext(ctx context.Context) ExpressRoutePortAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRoutePortAuthorizationOutput)
}

type ExpressRoutePortAuthorizationOutput struct{ *pulumi.OutputState }

func (ExpressRoutePortAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRoutePortAuthorization)(nil)).Elem()
}

func (o ExpressRoutePortAuthorizationOutput) ToExpressRoutePortAuthorizationOutput() ExpressRoutePortAuthorizationOutput {
	return o
}

func (o ExpressRoutePortAuthorizationOutput) ToExpressRoutePortAuthorizationOutputWithContext(ctx context.Context) ExpressRoutePortAuthorizationOutput {
	return o
}

// The authorization key.
func (o ExpressRoutePortAuthorizationOutput) AuthorizationKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.AuthorizationKey }).(pulumi.StringOutput)
}

// The authorization use status.
func (o ExpressRoutePortAuthorizationOutput) AuthorizationUseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.AuthorizationUseStatus }).(pulumi.StringOutput)
}

// The reference to the ExpressRoute circuit resource using the authorization.
func (o ExpressRoutePortAuthorizationOutput) CircuitResourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.CircuitResourceUri }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ExpressRoutePortAuthorizationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o ExpressRoutePortAuthorizationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the authorization resource.
func (o ExpressRoutePortAuthorizationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Type of the resource.
func (o ExpressRoutePortAuthorizationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRoutePortAuthorization) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRoutePortAuthorizationOutput{})
}
