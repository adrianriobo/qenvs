// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the existing immutability policy along with the corresponding ETag in response headers and body.
// Azure REST API version: 2022-09-01.
func LookupBlobContainerImmutabilityPolicy(ctx *pulumi.Context, args *LookupBlobContainerImmutabilityPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBlobContainerImmutabilityPolicyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBlobContainerImmutabilityPolicyResult
	err := ctx.Invoke("azure-native:storage:getBlobContainerImmutabilityPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobContainerImmutabilityPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ContainerName string `pulumi:"containerName"`
	// The name of the blob container immutabilityPolicy within the specified storage account. ImmutabilityPolicy Name must be 'default'
	ImmutabilityPolicyName string `pulumi:"immutabilityPolicyName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
type LookupBlobContainerImmutabilityPolicyResult struct {
	// This property can only be changed for unlocked time-based retention policies. When enabled, new blocks can be written to an append blob while maintaining immutability protection and compliance. Only new blocks can be added and any existing blocks cannot be modified or deleted. This property cannot be changed with ExtendImmutabilityPolicy API.
	AllowProtectedAppendWrites *bool `pulumi:"allowProtectedAppendWrites"`
	// This property can only be changed for unlocked time-based retention policies. When enabled, new blocks can be written to both 'Append and Bock Blobs' while maintaining immutability protection and compliance. Only new blocks can be added and any existing blocks cannot be modified or deleted. This property cannot be changed with ExtendImmutabilityPolicy API. The 'allowProtectedAppendWrites' and 'allowProtectedAppendWritesAll' properties are mutually exclusive.
	AllowProtectedAppendWritesAll *bool `pulumi:"allowProtectedAppendWritesAll"`
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The immutability period for the blobs in the container since the policy creation, in days.
	ImmutabilityPeriodSinceCreationInDays *int `pulumi:"immutabilityPeriodSinceCreationInDays"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
	State string `pulumi:"state"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupBlobContainerImmutabilityPolicyOutput(ctx *pulumi.Context, args LookupBlobContainerImmutabilityPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBlobContainerImmutabilityPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBlobContainerImmutabilityPolicyResult, error) {
			args := v.(LookupBlobContainerImmutabilityPolicyArgs)
			r, err := LookupBlobContainerImmutabilityPolicy(ctx, &args, opts...)
			var s LookupBlobContainerImmutabilityPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBlobContainerImmutabilityPolicyResultOutput)
}

type LookupBlobContainerImmutabilityPolicyOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ContainerName pulumi.StringInput `pulumi:"containerName"`
	// The name of the blob container immutabilityPolicy within the specified storage account. ImmutabilityPolicy Name must be 'default'
	ImmutabilityPolicyName pulumi.StringInput `pulumi:"immutabilityPolicyName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBlobContainerImmutabilityPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerImmutabilityPolicyArgs)(nil)).Elem()
}

// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
type LookupBlobContainerImmutabilityPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBlobContainerImmutabilityPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobContainerImmutabilityPolicyResult)(nil)).Elem()
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) ToLookupBlobContainerImmutabilityPolicyResultOutput() LookupBlobContainerImmutabilityPolicyResultOutput {
	return o
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) ToLookupBlobContainerImmutabilityPolicyResultOutputWithContext(ctx context.Context) LookupBlobContainerImmutabilityPolicyResultOutput {
	return o
}

func (o LookupBlobContainerImmutabilityPolicyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBlobContainerImmutabilityPolicyResult] {
	return pulumix.Output[LookupBlobContainerImmutabilityPolicyResult]{
		OutputState: o.OutputState,
	}
}

// This property can only be changed for unlocked time-based retention policies. When enabled, new blocks can be written to an append blob while maintaining immutability protection and compliance. Only new blocks can be added and any existing blocks cannot be modified or deleted. This property cannot be changed with ExtendImmutabilityPolicy API.
func (o LookupBlobContainerImmutabilityPolicyResultOutput) AllowProtectedAppendWrites() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) *bool { return v.AllowProtectedAppendWrites }).(pulumi.BoolPtrOutput)
}

// This property can only be changed for unlocked time-based retention policies. When enabled, new blocks can be written to both 'Append and Bock Blobs' while maintaining immutability protection and compliance. Only new blocks can be added and any existing blocks cannot be modified or deleted. This property cannot be changed with ExtendImmutabilityPolicy API. The 'allowProtectedAppendWrites' and 'allowProtectedAppendWritesAll' properties are mutually exclusive.
func (o LookupBlobContainerImmutabilityPolicyResultOutput) AllowProtectedAppendWritesAll() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) *bool { return v.AllowProtectedAppendWritesAll }).(pulumi.BoolPtrOutput)
}

// Resource Etag.
func (o LookupBlobContainerImmutabilityPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupBlobContainerImmutabilityPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// The immutability period for the blobs in the container since the policy creation, in days.
func (o LookupBlobContainerImmutabilityPolicyResultOutput) ImmutabilityPeriodSinceCreationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) *int {
		return v.ImmutabilityPeriodSinceCreationInDays
	}).(pulumi.IntPtrOutput)
}

// The name of the resource
func (o LookupBlobContainerImmutabilityPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
func (o LookupBlobContainerImmutabilityPolicyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) string { return v.State }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBlobContainerImmutabilityPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobContainerImmutabilityPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobContainerImmutabilityPolicyResultOutput{})
}
