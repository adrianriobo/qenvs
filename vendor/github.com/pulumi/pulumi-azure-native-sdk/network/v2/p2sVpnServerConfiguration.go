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

// P2SVpnServerConfiguration Resource.
// Azure REST API version: 2019-07-01. Prior API version in Azure Native 1.x: 2019-07-01
type P2sVpnServerConfiguration struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Paren VirtualWan resource name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// List of references to P2SVpnGateways.
	P2SVpnGateways SubResourceResponseArrayOutput `pulumi:"p2SVpnGateways"`
	// Radius client root certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigRadiusClientRootCertificates P2SVpnServerConfigRadiusClientRootCertificateResponseArrayOutput `pulumi:"p2SVpnServerConfigRadiusClientRootCertificates"`
	// Radius Server root certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigRadiusServerRootCertificates P2SVpnServerConfigRadiusServerRootCertificateResponseArrayOutput `pulumi:"p2SVpnServerConfigRadiusServerRootCertificates"`
	// VPN client revoked certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigVpnClientRevokedCertificates P2SVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput `pulumi:"p2SVpnServerConfigVpnClientRevokedCertificates"`
	// VPN client root certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigVpnClientRootCertificates P2SVpnServerConfigVpnClientRootCertificateResponseArrayOutput `pulumi:"p2SVpnServerConfigVpnClientRootCertificates"`
	// The provisioning state of the P2S VPN server configuration resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress pulumi.StringPtrOutput `pulumi:"radiusServerAddress"`
	// The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret pulumi.StringPtrOutput `pulumi:"radiusServerSecret"`
	// VpnClientIpsecPolicies for P2SVpnServerConfiguration.
	VpnClientIpsecPolicies IpsecPolicyResponseArrayOutput `pulumi:"vpnClientIpsecPolicies"`
	// VPN protocols for the P2SVpnServerConfiguration.
	VpnProtocols pulumi.StringArrayOutput `pulumi:"vpnProtocols"`
}

// NewP2sVpnServerConfiguration registers a new resource with the given unique name, arguments, and options.
func NewP2sVpnServerConfiguration(ctx *pulumi.Context,
	name string, args *P2sVpnServerConfigurationArgs, opts ...pulumi.ResourceOption) (*P2sVpnServerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualWanName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualWanName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180801:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:P2sVpnServerConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:P2sVpnServerConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource P2sVpnServerConfiguration
	err := ctx.RegisterResource("azure-native:network:P2sVpnServerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetP2sVpnServerConfiguration gets an existing P2sVpnServerConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetP2sVpnServerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *P2sVpnServerConfigurationState, opts ...pulumi.ResourceOption) (*P2sVpnServerConfiguration, error) {
	var resource P2sVpnServerConfiguration
	err := ctx.ReadResource("azure-native:network:P2sVpnServerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering P2sVpnServerConfiguration resources.
type p2sVpnServerConfigurationState struct {
}

type P2sVpnServerConfigurationState struct {
}

func (P2sVpnServerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*p2sVpnServerConfigurationState)(nil)).Elem()
}

type p2sVpnServerConfigurationArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Paren VirtualWan resource name.
	Name *string `pulumi:"name"`
	// Radius client root certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigRadiusClientRootCertificates []P2SVpnServerConfigRadiusClientRootCertificate `pulumi:"p2SVpnServerConfigRadiusClientRootCertificates"`
	// Radius Server root certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigRadiusServerRootCertificates []P2SVpnServerConfigRadiusServerRootCertificate `pulumi:"p2SVpnServerConfigRadiusServerRootCertificates"`
	// VPN client revoked certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigVpnClientRevokedCertificates []P2SVpnServerConfigVpnClientRevokedCertificate `pulumi:"p2SVpnServerConfigVpnClientRevokedCertificates"`
	// VPN client root certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigVpnClientRootCertificates []P2SVpnServerConfigVpnClientRootCertificate `pulumi:"p2SVpnServerConfigVpnClientRootCertificates"`
	// The name of the P2SVpnServerConfiguration.
	P2SVpnServerConfigurationName *string `pulumi:"p2SVpnServerConfigurationName"`
	// The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress *string `pulumi:"radiusServerAddress"`
	// The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret *string `pulumi:"radiusServerSecret"`
	// The resource group name of the VirtualWan.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualWan.
	VirtualWanName string `pulumi:"virtualWanName"`
	// VpnClientIpsecPolicies for P2SVpnServerConfiguration.
	VpnClientIpsecPolicies []IpsecPolicy `pulumi:"vpnClientIpsecPolicies"`
	// VPN protocols for the P2SVpnServerConfiguration.
	VpnProtocols []string `pulumi:"vpnProtocols"`
}

// The set of arguments for constructing a P2sVpnServerConfiguration resource.
type P2sVpnServerConfigurationArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Paren VirtualWan resource name.
	Name pulumi.StringPtrInput
	// Radius client root certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigRadiusClientRootCertificates P2SVpnServerConfigRadiusClientRootCertificateArrayInput
	// Radius Server root certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigRadiusServerRootCertificates P2SVpnServerConfigRadiusServerRootCertificateArrayInput
	// VPN client revoked certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigVpnClientRevokedCertificates P2SVpnServerConfigVpnClientRevokedCertificateArrayInput
	// VPN client root certificate of P2SVpnServerConfiguration.
	P2SVpnServerConfigVpnClientRootCertificates P2SVpnServerConfigVpnClientRootCertificateArrayInput
	// The name of the P2SVpnServerConfiguration.
	P2SVpnServerConfigurationName pulumi.StringPtrInput
	// The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
	RadiusServerAddress pulumi.StringPtrInput
	// The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
	RadiusServerSecret pulumi.StringPtrInput
	// The resource group name of the VirtualWan.
	ResourceGroupName pulumi.StringInput
	// The name of the VirtualWan.
	VirtualWanName pulumi.StringInput
	// VpnClientIpsecPolicies for P2SVpnServerConfiguration.
	VpnClientIpsecPolicies IpsecPolicyArrayInput
	// VPN protocols for the P2SVpnServerConfiguration.
	VpnProtocols pulumi.StringArrayInput
}

func (P2sVpnServerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*p2sVpnServerConfigurationArgs)(nil)).Elem()
}

type P2sVpnServerConfigurationInput interface {
	pulumi.Input

	ToP2sVpnServerConfigurationOutput() P2sVpnServerConfigurationOutput
	ToP2sVpnServerConfigurationOutputWithContext(ctx context.Context) P2sVpnServerConfigurationOutput
}

func (*P2sVpnServerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**P2sVpnServerConfiguration)(nil)).Elem()
}

func (i *P2sVpnServerConfiguration) ToP2sVpnServerConfigurationOutput() P2sVpnServerConfigurationOutput {
	return i.ToP2sVpnServerConfigurationOutputWithContext(context.Background())
}

func (i *P2sVpnServerConfiguration) ToP2sVpnServerConfigurationOutputWithContext(ctx context.Context) P2sVpnServerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(P2sVpnServerConfigurationOutput)
}

func (i *P2sVpnServerConfiguration) ToOutput(ctx context.Context) pulumix.Output[*P2sVpnServerConfiguration] {
	return pulumix.Output[*P2sVpnServerConfiguration]{
		OutputState: i.ToP2sVpnServerConfigurationOutputWithContext(ctx).OutputState,
	}
}

type P2sVpnServerConfigurationOutput struct{ *pulumi.OutputState }

func (P2sVpnServerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**P2sVpnServerConfiguration)(nil)).Elem()
}

func (o P2sVpnServerConfigurationOutput) ToP2sVpnServerConfigurationOutput() P2sVpnServerConfigurationOutput {
	return o
}

func (o P2sVpnServerConfigurationOutput) ToP2sVpnServerConfigurationOutputWithContext(ctx context.Context) P2sVpnServerConfigurationOutput {
	return o
}

func (o P2sVpnServerConfigurationOutput) ToOutput(ctx context.Context) pulumix.Output[*P2sVpnServerConfiguration] {
	return pulumix.Output[*P2sVpnServerConfiguration]{
		OutputState: o.OutputState,
	}
}

// A unique read-only string that changes whenever the resource is updated.
func (o P2sVpnServerConfigurationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Paren VirtualWan resource name.
func (o P2sVpnServerConfigurationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// List of references to P2SVpnGateways.
func (o P2sVpnServerConfigurationOutput) P2SVpnGateways() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) SubResourceResponseArrayOutput { return v.P2SVpnGateways }).(SubResourceResponseArrayOutput)
}

// Radius client root certificate of P2SVpnServerConfiguration.
func (o P2sVpnServerConfigurationOutput) P2SVpnServerConfigRadiusClientRootCertificates() P2SVpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) P2SVpnServerConfigRadiusClientRootCertificateResponseArrayOutput {
		return v.P2SVpnServerConfigRadiusClientRootCertificates
	}).(P2SVpnServerConfigRadiusClientRootCertificateResponseArrayOutput)
}

// Radius Server root certificate of P2SVpnServerConfiguration.
func (o P2sVpnServerConfigurationOutput) P2SVpnServerConfigRadiusServerRootCertificates() P2SVpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) P2SVpnServerConfigRadiusServerRootCertificateResponseArrayOutput {
		return v.P2SVpnServerConfigRadiusServerRootCertificates
	}).(P2SVpnServerConfigRadiusServerRootCertificateResponseArrayOutput)
}

// VPN client revoked certificate of P2SVpnServerConfiguration.
func (o P2sVpnServerConfigurationOutput) P2SVpnServerConfigVpnClientRevokedCertificates() P2SVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) P2SVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput {
		return v.P2SVpnServerConfigVpnClientRevokedCertificates
	}).(P2SVpnServerConfigVpnClientRevokedCertificateResponseArrayOutput)
}

// VPN client root certificate of P2SVpnServerConfiguration.
func (o P2sVpnServerConfigurationOutput) P2SVpnServerConfigVpnClientRootCertificates() P2SVpnServerConfigVpnClientRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) P2SVpnServerConfigVpnClientRootCertificateResponseArrayOutput {
		return v.P2SVpnServerConfigVpnClientRootCertificates
	}).(P2SVpnServerConfigVpnClientRootCertificateResponseArrayOutput)
}

// The provisioning state of the P2S VPN server configuration resource.
func (o P2sVpnServerConfigurationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
func (o P2sVpnServerConfigurationOutput) RadiusServerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringPtrOutput { return v.RadiusServerAddress }).(pulumi.StringPtrOutput)
}

// The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
func (o P2sVpnServerConfigurationOutput) RadiusServerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringPtrOutput { return v.RadiusServerSecret }).(pulumi.StringPtrOutput)
}

// VpnClientIpsecPolicies for P2SVpnServerConfiguration.
func (o P2sVpnServerConfigurationOutput) VpnClientIpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) IpsecPolicyResponseArrayOutput { return v.VpnClientIpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

// VPN protocols for the P2SVpnServerConfiguration.
func (o P2sVpnServerConfigurationOutput) VpnProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *P2sVpnServerConfiguration) pulumi.StringArrayOutput { return v.VpnProtocols }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(P2sVpnServerConfigurationOutput{})
}
