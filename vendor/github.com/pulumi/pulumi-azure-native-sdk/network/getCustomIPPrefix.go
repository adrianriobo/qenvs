// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified custom IP prefix in a specified resource group.
// API Version: 2020-11-01.
func LookupCustomIPPrefix(ctx *pulumi.Context, args *LookupCustomIPPrefixArgs, opts ...pulumi.InvokeOption) (*LookupCustomIPPrefixResult, error) {
	var rv LookupCustomIPPrefixResult
	err := ctx.Invoke("azure-native:network:getCustomIPPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomIPPrefixArgs struct {
	// The name of the custom IP prefix.
	CustomIpPrefixName string `pulumi:"customIpPrefixName"`
	// Expands referenced resources.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Custom IP prefix resource.
type LookupCustomIPPrefixResult struct {
	// The prefix range in CIDR notation. Should include the start address and the prefix length.
	Cidr *string `pulumi:"cidr"`
	// The commissioned state of the Custom IP Prefix.
	CommissionedState *string `pulumi:"commissionedState"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The extended location of the custom IP prefix.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the custom IP prefix resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The list of all referenced PublicIpPrefixes.
	PublicIpPrefixes []SubResourceResponse `pulumi:"publicIpPrefixes"`
	// The resource GUID property of the custom IP prefix resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

func LookupCustomIPPrefixOutput(ctx *pulumi.Context, args LookupCustomIPPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupCustomIPPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomIPPrefixResult, error) {
			args := v.(LookupCustomIPPrefixArgs)
			r, err := LookupCustomIPPrefix(ctx, &args, opts...)
			var s LookupCustomIPPrefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomIPPrefixResultOutput)
}

type LookupCustomIPPrefixOutputArgs struct {
	// The name of the custom IP prefix.
	CustomIpPrefixName pulumi.StringInput `pulumi:"customIpPrefixName"`
	// Expands referenced resources.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCustomIPPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomIPPrefixArgs)(nil)).Elem()
}

// Custom IP prefix resource.
type LookupCustomIPPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupCustomIPPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomIPPrefixResult)(nil)).Elem()
}

func (o LookupCustomIPPrefixResultOutput) ToLookupCustomIPPrefixResultOutput() LookupCustomIPPrefixResultOutput {
	return o
}

func (o LookupCustomIPPrefixResultOutput) ToLookupCustomIPPrefixResultOutputWithContext(ctx context.Context) LookupCustomIPPrefixResultOutput {
	return o
}

// The prefix range in CIDR notation. Should include the start address and the prefix length.
func (o LookupCustomIPPrefixResultOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Cidr }).(pulumi.StringPtrOutput)
}

// The commissioned state of the Custom IP Prefix.
func (o LookupCustomIPPrefixResultOutput) CommissionedState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.CommissionedState }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupCustomIPPrefixResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the custom IP prefix.
func (o LookupCustomIPPrefixResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Resource ID.
func (o LookupCustomIPPrefixResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupCustomIPPrefixResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupCustomIPPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the custom IP prefix resource.
func (o LookupCustomIPPrefixResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of all referenced PublicIpPrefixes.
func (o LookupCustomIPPrefixResultOutput) PublicIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) []SubResourceResponse { return v.PublicIpPrefixes }).(SubResourceResponseArrayOutput)
}

// The resource GUID property of the custom IP prefix resource.
func (o LookupCustomIPPrefixResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupCustomIPPrefixResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupCustomIPPrefixResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of availability zones denoting the IP allocated for the resource needs to come from.
func (o LookupCustomIPPrefixResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCustomIPPrefixResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomIPPrefixResultOutput{})
}
