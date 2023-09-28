// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Specifies information about the SSH public key.
// Azure REST API version: 2023-03-01. Prior API version in Azure Native 1.x: 2020-12-01
type SshPublicKey struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// SSH public key used to authenticate to a virtual machine through ssh. If this property is not initially provided when the resource is created, the publicKey property will be populated when generateKeyPair is called. If the public key is provided upon resource creation, the provided public key needs to be at least 2048-bit and in ssh-rsa format.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSshPublicKey registers a new resource with the given unique name, arguments, and options.
func NewSshPublicKey(ctx *pulumi.Context,
	name string, args *SshPublicKeyArgs, opts ...pulumi.ResourceOption) (*SshPublicKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20191201:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:SshPublicKey"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230701:SshPublicKey"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SshPublicKey
	err := ctx.RegisterResource("azure-native:compute:SshPublicKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSshPublicKey gets an existing SshPublicKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSshPublicKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SshPublicKeyState, opts ...pulumi.ResourceOption) (*SshPublicKey, error) {
	var resource SshPublicKey
	err := ctx.ReadResource("azure-native:compute:SshPublicKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SshPublicKey resources.
type sshPublicKeyState struct {
}

type SshPublicKeyState struct {
}

func (SshPublicKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sshPublicKeyState)(nil)).Elem()
}

type sshPublicKeyArgs struct {
	// Resource location
	Location *string `pulumi:"location"`
	// SSH public key used to authenticate to a virtual machine through ssh. If this property is not initially provided when the resource is created, the publicKey property will be populated when generateKeyPair is called. If the public key is provided upon resource creation, the provided public key needs to be at least 2048-bit and in ssh-rsa format.
	PublicKey *string `pulumi:"publicKey"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SSH public key.
	SshPublicKeyName *string `pulumi:"sshPublicKeyName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SshPublicKey resource.
type SshPublicKeyArgs struct {
	// Resource location
	Location pulumi.StringPtrInput
	// SSH public key used to authenticate to a virtual machine through ssh. If this property is not initially provided when the resource is created, the publicKey property will be populated when generateKeyPair is called. If the public key is provided upon resource creation, the provided public key needs to be at least 2048-bit and in ssh-rsa format.
	PublicKey pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the SSH public key.
	SshPublicKeyName pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (SshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sshPublicKeyArgs)(nil)).Elem()
}

type SshPublicKeyInput interface {
	pulumi.Input

	ToSshPublicKeyOutput() SshPublicKeyOutput
	ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput
}

func (*SshPublicKey) ElementType() reflect.Type {
	return reflect.TypeOf((**SshPublicKey)(nil)).Elem()
}

func (i *SshPublicKey) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return i.ToSshPublicKeyOutputWithContext(context.Background())
}

func (i *SshPublicKey) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyOutput)
}

func (i *SshPublicKey) ToOutput(ctx context.Context) pulumix.Output[*SshPublicKey] {
	return pulumix.Output[*SshPublicKey]{
		OutputState: i.ToSshPublicKeyOutputWithContext(ctx).OutputState,
	}
}

type SshPublicKeyOutput struct{ *pulumi.OutputState }

func (SshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) ToOutput(ctx context.Context) pulumix.Output[*SshPublicKey] {
	return pulumix.Output[*SshPublicKey]{
		OutputState: o.OutputState,
	}
}

// Resource location
func (o SshPublicKeyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o SshPublicKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SSH public key used to authenticate to a virtual machine through ssh. If this property is not initially provided when the resource is created, the publicKey property will be populated when generateKeyPair is called. If the public key is provided upon resource creation, the provided public key needs to be at least 2048-bit and in ssh-rsa format.
func (o SshPublicKeyOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o SshPublicKeyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o SshPublicKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SshPublicKey) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SshPublicKeyOutput{})
}
