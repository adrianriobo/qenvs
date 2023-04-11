// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The operation to get the restore point.
// API Version: 2021-03-01.
func LookupRestorePoint(ctx *pulumi.Context, args *LookupRestorePointArgs, opts ...pulumi.InvokeOption) (*LookupRestorePointResult, error) {
	var rv LookupRestorePointResult
	err := ctx.Invoke("azure-native:compute:getRestorePoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRestorePointArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the restore point collection.
	RestorePointCollectionName string `pulumi:"restorePointCollectionName"`
	// The name of the restore point.
	RestorePointName string `pulumi:"restorePointName"`
}

// Restore Point details.
type LookupRestorePointResult struct {
	// Gets the consistency mode for the restore point. Please refer to https://aka.ms/RestorePoints for more details.
	ConsistencyMode string `pulumi:"consistencyMode"`
	// List of disk resource ids that the customer wishes to exclude from the restore point. If no disks are specified, all disks will be included.
	ExcludeDisks []ApiEntityReferenceResponse `pulumi:"excludeDisks"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name string `pulumi:"name"`
	// Gets the provisioning state of the restore point.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets the details of the VM captured at the time of the restore point creation.
	SourceMetadata RestorePointSourceMetadataResponse `pulumi:"sourceMetadata"`
	// Gets the creation time of the restore point.
	TimeCreated *string `pulumi:"timeCreated"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupRestorePointOutput(ctx *pulumi.Context, args LookupRestorePointOutputArgs, opts ...pulumi.InvokeOption) LookupRestorePointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRestorePointResult, error) {
			args := v.(LookupRestorePointArgs)
			r, err := LookupRestorePoint(ctx, &args, opts...)
			var s LookupRestorePointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRestorePointResultOutput)
}

type LookupRestorePointOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the restore point collection.
	RestorePointCollectionName pulumi.StringInput `pulumi:"restorePointCollectionName"`
	// The name of the restore point.
	RestorePointName pulumi.StringInput `pulumi:"restorePointName"`
}

func (LookupRestorePointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestorePointArgs)(nil)).Elem()
}

// Restore Point details.
type LookupRestorePointResultOutput struct{ *pulumi.OutputState }

func (LookupRestorePointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestorePointResult)(nil)).Elem()
}

func (o LookupRestorePointResultOutput) ToLookupRestorePointResultOutput() LookupRestorePointResultOutput {
	return o
}

func (o LookupRestorePointResultOutput) ToLookupRestorePointResultOutputWithContext(ctx context.Context) LookupRestorePointResultOutput {
	return o
}

// Gets the consistency mode for the restore point. Please refer to https://aka.ms/RestorePoints for more details.
func (o LookupRestorePointResultOutput) ConsistencyMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.ConsistencyMode }).(pulumi.StringOutput)
}

// List of disk resource ids that the customer wishes to exclude from the restore point. If no disks are specified, all disks will be included.
func (o LookupRestorePointResultOutput) ExcludeDisks() ApiEntityReferenceResponseArrayOutput {
	return o.ApplyT(func(v LookupRestorePointResult) []ApiEntityReferenceResponse { return v.ExcludeDisks }).(ApiEntityReferenceResponseArrayOutput)
}

// Resource Id
func (o LookupRestorePointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupRestorePointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the restore point.
func (o LookupRestorePointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the details of the VM captured at the time of the restore point creation.
func (o LookupRestorePointResultOutput) SourceMetadata() RestorePointSourceMetadataResponseOutput {
	return o.ApplyT(func(v LookupRestorePointResult) RestorePointSourceMetadataResponse { return v.SourceMetadata }).(RestorePointSourceMetadataResponseOutput)
}

// Gets the creation time of the restore point.
func (o LookupRestorePointResultOutput) TimeCreated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRestorePointResult) *string { return v.TimeCreated }).(pulumi.StringPtrOutput)
}

// Resource type
func (o LookupRestorePointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRestorePointResultOutput{})
}
