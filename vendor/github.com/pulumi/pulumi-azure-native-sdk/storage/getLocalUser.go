// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the local user of the storage account by username.
// API Version: 2021-08-01.
func LookupLocalUser(ctx *pulumi.Context, args *LookupLocalUserArgs, opts ...pulumi.InvokeOption) (*LookupLocalUserResult, error) {
	var rv LookupLocalUserResult
	err := ctx.Invoke("azure-native:storage:getLocalUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLocalUserArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of local user. The username must contain lowercase letters and numbers only. It must be unique only within the storage account.
	Username string `pulumi:"username"`
}

// The local user associated with the storage accounts.
type LookupLocalUserResult struct {
	// Indicates whether shared key exists. Set it to false to remove existing shared key.
	HasSharedKey *bool `pulumi:"hasSharedKey"`
	// Indicates whether ssh key exists. Set it to false to remove existing SSH key.
	HasSshKey *bool `pulumi:"hasSshKey"`
	// Indicates whether ssh password exists. Set it to false to remove existing SSH password.
	HasSshPassword *bool `pulumi:"hasSshPassword"`
	// Optional, local user home directory.
	HomeDirectory *string `pulumi:"homeDirectory"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The permission scopes of the local user.
	PermissionScopes []PermissionScopeResponse `pulumi:"permissionScopes"`
	// A unique Security Identifier that is generated by the server.
	Sid string `pulumi:"sid"`
	// Optional, local user ssh authorized keys for SFTP.
	SshAuthorizedKeys []SshPublicKeyResponse `pulumi:"sshAuthorizedKeys"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupLocalUserOutput(ctx *pulumi.Context, args LookupLocalUserOutputArgs, opts ...pulumi.InvokeOption) LookupLocalUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalUserResult, error) {
			args := v.(LookupLocalUserArgs)
			r, err := LookupLocalUser(ctx, &args, opts...)
			var s LookupLocalUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalUserResultOutput)
}

type LookupLocalUserOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of local user. The username must contain lowercase letters and numbers only. It must be unique only within the storage account.
	Username pulumi.StringInput `pulumi:"username"`
}

func (LookupLocalUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalUserArgs)(nil)).Elem()
}

// The local user associated with the storage accounts.
type LookupLocalUserResultOutput struct{ *pulumi.OutputState }

func (LookupLocalUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalUserResult)(nil)).Elem()
}

func (o LookupLocalUserResultOutput) ToLookupLocalUserResultOutput() LookupLocalUserResultOutput {
	return o
}

func (o LookupLocalUserResultOutput) ToLookupLocalUserResultOutputWithContext(ctx context.Context) LookupLocalUserResultOutput {
	return o
}

// Indicates whether shared key exists. Set it to false to remove existing shared key.
func (o LookupLocalUserResultOutput) HasSharedKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalUserResult) *bool { return v.HasSharedKey }).(pulumi.BoolPtrOutput)
}

// Indicates whether ssh key exists. Set it to false to remove existing SSH key.
func (o LookupLocalUserResultOutput) HasSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalUserResult) *bool { return v.HasSshKey }).(pulumi.BoolPtrOutput)
}

// Indicates whether ssh password exists. Set it to false to remove existing SSH password.
func (o LookupLocalUserResultOutput) HasSshPassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalUserResult) *bool { return v.HasSshPassword }).(pulumi.BoolPtrOutput)
}

// Optional, local user home directory.
func (o LookupLocalUserResultOutput) HomeDirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalUserResult) *string { return v.HomeDirectory }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLocalUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLocalUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// The permission scopes of the local user.
func (o LookupLocalUserResultOutput) PermissionScopes() PermissionScopeResponseArrayOutput {
	return o.ApplyT(func(v LookupLocalUserResult) []PermissionScopeResponse { return v.PermissionScopes }).(PermissionScopeResponseArrayOutput)
}

// A unique Security Identifier that is generated by the server.
func (o LookupLocalUserResultOutput) Sid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalUserResult) string { return v.Sid }).(pulumi.StringOutput)
}

// Optional, local user ssh authorized keys for SFTP.
func (o LookupLocalUserResultOutput) SshAuthorizedKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v LookupLocalUserResult) []SshPublicKeyResponse { return v.SshAuthorizedKeys }).(SshPublicKeyResponseArrayOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupLocalUserResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLocalUserResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLocalUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalUserResultOutput{})
}
