// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified DDoS protection plan.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2018-02-01, 2022-05-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01.
func LookupDdosProtectionPlan(ctx *pulumi.Context, args *LookupDdosProtectionPlanArgs, opts ...pulumi.InvokeOption) (*LookupDdosProtectionPlanResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDdosProtectionPlanResult
	err := ctx.Invoke("azure-native:network:getDdosProtectionPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDdosProtectionPlanArgs struct {
	// The name of the DDoS protection plan.
	DdosProtectionPlanName string `pulumi:"ddosProtectionPlanName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A DDoS protection plan in a resource group.
type LookupDdosProtectionPlanResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the DDoS protection plan resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The list of public IPs associated with the DDoS protection plan resource. This list is read-only.
	PublicIPAddresses []SubResourceResponse `pulumi:"publicIPAddresses"`
	// The resource GUID property of the DDoS protection plan resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The list of virtual networks associated with the DDoS protection plan resource. This list is read-only.
	VirtualNetworks []SubResourceResponse `pulumi:"virtualNetworks"`
}

func LookupDdosProtectionPlanOutput(ctx *pulumi.Context, args LookupDdosProtectionPlanOutputArgs, opts ...pulumi.InvokeOption) LookupDdosProtectionPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDdosProtectionPlanResult, error) {
			args := v.(LookupDdosProtectionPlanArgs)
			r, err := LookupDdosProtectionPlan(ctx, &args, opts...)
			var s LookupDdosProtectionPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDdosProtectionPlanResultOutput)
}

type LookupDdosProtectionPlanOutputArgs struct {
	// The name of the DDoS protection plan.
	DdosProtectionPlanName pulumi.StringInput `pulumi:"ddosProtectionPlanName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDdosProtectionPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDdosProtectionPlanArgs)(nil)).Elem()
}

// A DDoS protection plan in a resource group.
type LookupDdosProtectionPlanResultOutput struct{ *pulumi.OutputState }

func (LookupDdosProtectionPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDdosProtectionPlanResult)(nil)).Elem()
}

func (o LookupDdosProtectionPlanResultOutput) ToLookupDdosProtectionPlanResultOutput() LookupDdosProtectionPlanResultOutput {
	return o
}

func (o LookupDdosProtectionPlanResultOutput) ToLookupDdosProtectionPlanResultOutputWithContext(ctx context.Context) LookupDdosProtectionPlanResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupDdosProtectionPlanResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupDdosProtectionPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupDdosProtectionPlanResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupDdosProtectionPlanResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the DDoS protection plan resource.
func (o LookupDdosProtectionPlanResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of public IPs associated with the DDoS protection plan resource. This list is read-only.
func (o LookupDdosProtectionPlanResultOutput) PublicIPAddresses() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) []SubResourceResponse { return v.PublicIPAddresses }).(SubResourceResponseArrayOutput)
}

// The resource GUID property of the DDoS protection plan resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
func (o LookupDdosProtectionPlanResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupDdosProtectionPlanResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupDdosProtectionPlanResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) string { return v.Type }).(pulumi.StringOutput)
}

// The list of virtual networks associated with the DDoS protection plan resource. This list is read-only.
func (o LookupDdosProtectionPlanResultOutput) VirtualNetworks() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupDdosProtectionPlanResult) []SubResourceResponse { return v.VirtualNetworks }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDdosProtectionPlanResultOutput{})
}
