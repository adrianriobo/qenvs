// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a resource group.
// Azure REST API version: 2022-09-01.
func LookupResourceGroup(ctx *pulumi.Context, args *LookupResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupResourceGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceGroupResult
	err := ctx.Invoke("azure-native:resources:getResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceGroupArgs struct {
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Resource group information.
type LookupResourceGroupResult struct {
	// The ID of the resource group.
	Id string `pulumi:"id"`
	// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
	Location string `pulumi:"location"`
	// The ID of the resource that manages this resource group.
	ManagedBy *string `pulumi:"managedBy"`
	// The name of the resource group.
	Name string `pulumi:"name"`
	// The resource group properties.
	Properties ResourceGroupPropertiesResponse `pulumi:"properties"`
	// The tags attached to the resource group.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource group.
	Type string `pulumi:"type"`
}

func LookupResourceGroupOutput(ctx *pulumi.Context, args LookupResourceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupResourceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceGroupResult, error) {
			args := v.(LookupResourceGroupArgs)
			r, err := LookupResourceGroup(ctx, &args, opts...)
			var s LookupResourceGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceGroupResultOutput)
}

type LookupResourceGroupOutputArgs struct {
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupResourceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGroupArgs)(nil)).Elem()
}

// Resource group information.
type LookupResourceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupResourceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGroupResult)(nil)).Elem()
}

func (o LookupResourceGroupResultOutput) ToLookupResourceGroupResultOutput() LookupResourceGroupResultOutput {
	return o
}

func (o LookupResourceGroupResultOutput) ToLookupResourceGroupResultOutputWithContext(ctx context.Context) LookupResourceGroupResultOutput {
	return o
}

func (o LookupResourceGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupResourceGroupResult] {
	return pulumix.Output[LookupResourceGroupResult]{
		OutputState: o.OutputState,
	}
}

// The ID of the resource group.
func (o LookupResourceGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
func (o LookupResourceGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// The ID of the resource that manages this resource group.
func (o LookupResourceGroupResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// The name of the resource group.
func (o LookupResourceGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource group properties.
func (o LookupResourceGroupResultOutput) Properties() ResourceGroupPropertiesResponseOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) ResourceGroupPropertiesResponse { return v.Properties }).(ResourceGroupPropertiesResponseOutput)
}

// The tags attached to the resource group.
func (o LookupResourceGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource group.
func (o LookupResourceGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceGroupResultOutput{})
}
