// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about an SSH public key.
// Azure REST API version: 2023-03-01.
//
// Other available API versions: 2023-07-01, 2023-09-01.
func LookupSshPublicKey(ctx *pulumi.Context, args *LookupSshPublicKeyArgs, opts ...pulumi.InvokeOption) (*LookupSshPublicKeyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSshPublicKeyResult
	err := ctx.Invoke("azure-native:compute:getSshPublicKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSshPublicKeyArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SSH public key.
	SshPublicKeyName string `pulumi:"sshPublicKeyName"`
}

// Specifies information about the SSH public key.
type LookupSshPublicKeyResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// SSH public key used to authenticate to a virtual machine through ssh. If this property is not initially provided when the resource is created, the publicKey property will be populated when generateKeyPair is called. If the public key is provided upon resource creation, the provided public key needs to be at least 2048-bit and in ssh-rsa format.
	PublicKey *string `pulumi:"publicKey"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupSshPublicKeyOutput(ctx *pulumi.Context, args LookupSshPublicKeyOutputArgs, opts ...pulumi.InvokeOption) LookupSshPublicKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSshPublicKeyResult, error) {
			args := v.(LookupSshPublicKeyArgs)
			r, err := LookupSshPublicKey(ctx, &args, opts...)
			var s LookupSshPublicKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSshPublicKeyResultOutput)
}

type LookupSshPublicKeyOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the SSH public key.
	SshPublicKeyName pulumi.StringInput `pulumi:"sshPublicKeyName"`
}

func (LookupSshPublicKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSshPublicKeyArgs)(nil)).Elem()
}

// Specifies information about the SSH public key.
type LookupSshPublicKeyResultOutput struct{ *pulumi.OutputState }

func (LookupSshPublicKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSshPublicKeyResult)(nil)).Elem()
}

func (o LookupSshPublicKeyResultOutput) ToLookupSshPublicKeyResultOutput() LookupSshPublicKeyResultOutput {
	return o
}

func (o LookupSshPublicKeyResultOutput) ToLookupSshPublicKeyResultOutputWithContext(ctx context.Context) LookupSshPublicKeyResultOutput {
	return o
}

// Resource Id
func (o LookupSshPublicKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupSshPublicKeyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupSshPublicKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

// SSH public key used to authenticate to a virtual machine through ssh. If this property is not initially provided when the resource is created, the publicKey property will be populated when generateKeyPair is called. If the public key is provided upon resource creation, the provided public key needs to be at least 2048-bit and in ssh-rsa format.
func (o LookupSshPublicKeyResultOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o LookupSshPublicKeyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupSshPublicKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSshPublicKeyResultOutput{})
}
