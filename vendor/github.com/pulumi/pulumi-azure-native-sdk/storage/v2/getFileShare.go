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

// Gets properties of a specified share.
// Azure REST API version: 2022-09-01.
//
// Other available API versions: 2023-01-01.
func LookupFileShare(ctx *pulumi.Context, args *LookupFileShareArgs, opts ...pulumi.InvokeOption) (*LookupFileShareResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFileShareResult
	err := ctx.Invoke("azure-native:storage:getFileShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileShareArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// Optional, used to expand the properties within share's properties. Valid values are: stats. Should be passed as a string with delimiter ','.
	Expand *string `pulumi:"expand"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ShareName string `pulumi:"shareName"`
}

// Properties of the file share, including Id, resource name, resource type, Etag.
type LookupFileShareResult struct {
	// Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
	AccessTier *string `pulumi:"accessTier"`
	// Indicates the last modification time for share access tier.
	AccessTierChangeTime string `pulumi:"accessTierChangeTime"`
	// Indicates if there is a pending transition for access tier.
	AccessTierStatus string `pulumi:"accessTierStatus"`
	// Indicates whether the share was deleted.
	Deleted bool `pulumi:"deleted"`
	// The deleted time if the share was deleted.
	DeletedTime string `pulumi:"deletedTime"`
	// The authentication protocol that is used for the file share. Can only be specified when creating a share.
	EnabledProtocols *string `pulumi:"enabledProtocols"`
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Returns the date and time the share was last modified.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	// Specifies whether the lease on a share is of infinite or fixed duration, only when the share is leased.
	LeaseDuration string `pulumi:"leaseDuration"`
	// Lease state of the share.
	LeaseState string `pulumi:"leaseState"`
	// The lease status of the share.
	LeaseStatus string `pulumi:"leaseStatus"`
	// A name-value pair to associate with the share as metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Remaining retention days for share that was soft deleted.
	RemainingRetentionDays int `pulumi:"remainingRetentionDays"`
	// The property is for NFS share only. The default is NoRootSquash.
	RootSquash *string `pulumi:"rootSquash"`
	// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400.
	ShareQuota *int `pulumi:"shareQuota"`
	// The approximate size of the data stored on the share. Note that this value may not include all recently created or recently resized files.
	ShareUsageBytes float64 `pulumi:"shareUsageBytes"`
	// List of stored access policies specified on the share.
	SignedIdentifiers []SignedIdentifierResponse `pulumi:"signedIdentifiers"`
	// Creation time of share snapshot returned in the response of list shares with expand param "snapshots".
	SnapshotTime string `pulumi:"snapshotTime"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The version of the share.
	Version string `pulumi:"version"`
}

func LookupFileShareOutput(ctx *pulumi.Context, args LookupFileShareOutputArgs, opts ...pulumi.InvokeOption) LookupFileShareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileShareResult, error) {
			args := v.(LookupFileShareArgs)
			r, err := LookupFileShare(ctx, &args, opts...)
			var s LookupFileShareResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileShareResultOutput)
}

type LookupFileShareOutputArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// Optional, used to expand the properties within share's properties. Valid values are: stats. Should be passed as a string with delimiter ','.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ShareName pulumi.StringInput `pulumi:"shareName"`
}

func (LookupFileShareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileShareArgs)(nil)).Elem()
}

// Properties of the file share, including Id, resource name, resource type, Etag.
type LookupFileShareResultOutput struct{ *pulumi.OutputState }

func (LookupFileShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileShareResult)(nil)).Elem()
}

func (o LookupFileShareResultOutput) ToLookupFileShareResultOutput() LookupFileShareResultOutput {
	return o
}

func (o LookupFileShareResultOutput) ToLookupFileShareResultOutputWithContext(ctx context.Context) LookupFileShareResultOutput {
	return o
}

func (o LookupFileShareResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFileShareResult] {
	return pulumix.Output[LookupFileShareResult]{
		OutputState: o.OutputState,
	}
}

// Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
func (o LookupFileShareResultOutput) AccessTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileShareResult) *string { return v.AccessTier }).(pulumi.StringPtrOutput)
}

// Indicates the last modification time for share access tier.
func (o LookupFileShareResultOutput) AccessTierChangeTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.AccessTierChangeTime }).(pulumi.StringOutput)
}

// Indicates if there is a pending transition for access tier.
func (o LookupFileShareResultOutput) AccessTierStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.AccessTierStatus }).(pulumi.StringOutput)
}

// Indicates whether the share was deleted.
func (o LookupFileShareResultOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFileShareResult) bool { return v.Deleted }).(pulumi.BoolOutput)
}

// The deleted time if the share was deleted.
func (o LookupFileShareResultOutput) DeletedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.DeletedTime }).(pulumi.StringOutput)
}

// The authentication protocol that is used for the file share. Can only be specified when creating a share.
func (o LookupFileShareResultOutput) EnabledProtocols() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileShareResult) *string { return v.EnabledProtocols }).(pulumi.StringPtrOutput)
}

// Resource Etag.
func (o LookupFileShareResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFileShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Id }).(pulumi.StringOutput)
}

// Returns the date and time the share was last modified.
func (o LookupFileShareResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// Specifies whether the lease on a share is of infinite or fixed duration, only when the share is leased.
func (o LookupFileShareResultOutput) LeaseDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.LeaseDuration }).(pulumi.StringOutput)
}

// Lease state of the share.
func (o LookupFileShareResultOutput) LeaseState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.LeaseState }).(pulumi.StringOutput)
}

// The lease status of the share.
func (o LookupFileShareResultOutput) LeaseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.LeaseStatus }).(pulumi.StringOutput)
}

// A name-value pair to associate with the share as metadata.
func (o LookupFileShareResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFileShareResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the resource
func (o LookupFileShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Name }).(pulumi.StringOutput)
}

// Remaining retention days for share that was soft deleted.
func (o LookupFileShareResultOutput) RemainingRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFileShareResult) int { return v.RemainingRetentionDays }).(pulumi.IntOutput)
}

// The property is for NFS share only. The default is NoRootSquash.
func (o LookupFileShareResultOutput) RootSquash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileShareResult) *string { return v.RootSquash }).(pulumi.StringPtrOutput)
}

// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400.
func (o LookupFileShareResultOutput) ShareQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFileShareResult) *int { return v.ShareQuota }).(pulumi.IntPtrOutput)
}

// The approximate size of the data stored on the share. Note that this value may not include all recently created or recently resized files.
func (o LookupFileShareResultOutput) ShareUsageBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupFileShareResult) float64 { return v.ShareUsageBytes }).(pulumi.Float64Output)
}

// List of stored access policies specified on the share.
func (o LookupFileShareResultOutput) SignedIdentifiers() SignedIdentifierResponseArrayOutput {
	return o.ApplyT(func(v LookupFileShareResult) []SignedIdentifierResponse { return v.SignedIdentifiers }).(SignedIdentifierResponseArrayOutput)
}

// Creation time of share snapshot returned in the response of list shares with expand param "snapshots".
func (o LookupFileShareResultOutput) SnapshotTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.SnapshotTime }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFileShareResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Type }).(pulumi.StringOutput)
}

// The version of the share.
func (o LookupFileShareResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileShareResultOutput{})
}
