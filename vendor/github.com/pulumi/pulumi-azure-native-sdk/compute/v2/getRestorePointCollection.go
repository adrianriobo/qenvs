// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The operation to get the restore point collection.
// Azure REST API version: 2023-03-01.
//
// Other available API versions: 2023-07-01, 2023-09-01.
func LookupRestorePointCollection(ctx *pulumi.Context, args *LookupRestorePointCollectionArgs, opts ...pulumi.InvokeOption) (*LookupRestorePointCollectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRestorePointCollectionResult
	err := ctx.Invoke("azure-native:compute:getRestorePointCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRestorePointCollectionArgs struct {
	// The expand expression to apply on the operation. If expand=restorePoints, server will return all contained restore points in the restorePointCollection.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the restore point collection.
	RestorePointCollectionName string `pulumi:"restorePointCollectionName"`
}

// Create or update Restore Point collection parameters.
type LookupRestorePointCollectionResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// The provisioning state of the restore point collection.
	ProvisioningState string `pulumi:"provisioningState"`
	// The unique id of the restore point collection.
	RestorePointCollectionId string `pulumi:"restorePointCollectionId"`
	// A list containing all restore points created under this restore point collection.
	RestorePoints []RestorePointResponse `pulumi:"restorePoints"`
	// The properties of the source resource that this restore point collection is created from.
	Source *RestorePointCollectionSourcePropertiesResponse `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupRestorePointCollectionOutput(ctx *pulumi.Context, args LookupRestorePointCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupRestorePointCollectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRestorePointCollectionResult, error) {
			args := v.(LookupRestorePointCollectionArgs)
			r, err := LookupRestorePointCollection(ctx, &args, opts...)
			var s LookupRestorePointCollectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRestorePointCollectionResultOutput)
}

type LookupRestorePointCollectionOutputArgs struct {
	// The expand expression to apply on the operation. If expand=restorePoints, server will return all contained restore points in the restorePointCollection.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the restore point collection.
	RestorePointCollectionName pulumi.StringInput `pulumi:"restorePointCollectionName"`
}

func (LookupRestorePointCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestorePointCollectionArgs)(nil)).Elem()
}

// Create or update Restore Point collection parameters.
type LookupRestorePointCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupRestorePointCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestorePointCollectionResult)(nil)).Elem()
}

func (o LookupRestorePointCollectionResultOutput) ToLookupRestorePointCollectionResultOutput() LookupRestorePointCollectionResultOutput {
	return o
}

func (o LookupRestorePointCollectionResultOutput) ToLookupRestorePointCollectionResultOutputWithContext(ctx context.Context) LookupRestorePointCollectionResultOutput {
	return o
}

// Resource Id
func (o LookupRestorePointCollectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupRestorePointCollectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupRestorePointCollectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the restore point collection.
func (o LookupRestorePointCollectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The unique id of the restore point collection.
func (o LookupRestorePointCollectionResultOutput) RestorePointCollectionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.RestorePointCollectionId }).(pulumi.StringOutput)
}

// A list containing all restore points created under this restore point collection.
func (o LookupRestorePointCollectionResultOutput) RestorePoints() RestorePointResponseArrayOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) []RestorePointResponse { return v.RestorePoints }).(RestorePointResponseArrayOutput)
}

// The properties of the source resource that this restore point collection is created from.
func (o LookupRestorePointCollectionResultOutput) Source() RestorePointCollectionSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) *RestorePointCollectionSourcePropertiesResponse {
		return v.Source
	}).(RestorePointCollectionSourcePropertiesResponsePtrOutput)
}

// Resource tags
func (o LookupRestorePointCollectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupRestorePointCollectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestorePointCollectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRestorePointCollectionResultOutput{})
}
