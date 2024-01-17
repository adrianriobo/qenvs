// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the details of a RoutingIntent.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01.
func LookupRoutingIntent(ctx *pulumi.Context, args *LookupRoutingIntentArgs, opts ...pulumi.InvokeOption) (*LookupRoutingIntentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRoutingIntentResult
	err := ctx.Invoke("azure-native:network:getRoutingIntent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoutingIntentArgs struct {
	// The resource group name of the RoutingIntent.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the RoutingIntent.
	RoutingIntentName string `pulumi:"routingIntentName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// The routing intent child resource of a Virtual hub.
type LookupRoutingIntentResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the RoutingIntent resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// List of routing policies.
	RoutingPolicies []RoutingPolicyResponse `pulumi:"routingPolicies"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupRoutingIntentOutput(ctx *pulumi.Context, args LookupRoutingIntentOutputArgs, opts ...pulumi.InvokeOption) LookupRoutingIntentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoutingIntentResult, error) {
			args := v.(LookupRoutingIntentArgs)
			r, err := LookupRoutingIntent(ctx, &args, opts...)
			var s LookupRoutingIntentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoutingIntentResultOutput)
}

type LookupRoutingIntentOutputArgs struct {
	// The resource group name of the RoutingIntent.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the RoutingIntent.
	RoutingIntentName pulumi.StringInput `pulumi:"routingIntentName"`
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupRoutingIntentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutingIntentArgs)(nil)).Elem()
}

// The routing intent child resource of a Virtual hub.
type LookupRoutingIntentResultOutput struct{ *pulumi.OutputState }

func (LookupRoutingIntentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoutingIntentResult)(nil)).Elem()
}

func (o LookupRoutingIntentResultOutput) ToLookupRoutingIntentResultOutput() LookupRoutingIntentResultOutput {
	return o
}

func (o LookupRoutingIntentResultOutput) ToLookupRoutingIntentResultOutputWithContext(ctx context.Context) LookupRoutingIntentResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupRoutingIntentResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupRoutingIntentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the resource that is unique within a resource group. This name can be used to access the resource.
func (o LookupRoutingIntentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The provisioning state of the RoutingIntent resource.
func (o LookupRoutingIntentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of routing policies.
func (o LookupRoutingIntentResultOutput) RoutingPolicies() RoutingPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) []RoutingPolicyResponse { return v.RoutingPolicies }).(RoutingPolicyResponseArrayOutput)
}

// Resource type.
func (o LookupRoutingIntentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoutingIntentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoutingIntentResultOutput{})
}
