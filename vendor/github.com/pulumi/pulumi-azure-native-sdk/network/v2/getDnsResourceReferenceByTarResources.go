// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the DNS records specified by the referencing targetResourceIds.
// Azure REST API version: 2018-05-01.
func GetDnsResourceReferenceByTarResources(ctx *pulumi.Context, args *GetDnsResourceReferenceByTarResourcesArgs, opts ...pulumi.InvokeOption) (*GetDnsResourceReferenceByTarResourcesResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetDnsResourceReferenceByTarResourcesResult
	err := ctx.Invoke("azure-native:network:getDnsResourceReferenceByTarResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDnsResourceReferenceByTarResourcesArgs struct {
	// A list of references to azure resources for which referencing dns records need to be queried.
	TargetResources []SubResource `pulumi:"targetResources"`
}

// Represents the properties of the Dns Resource Reference Result.
type GetDnsResourceReferenceByTarResourcesResult struct {
	// The result of dns resource reference request. A list of dns resource references for each of the azure resource in the request
	DnsResourceReferences []DnsResourceReferenceResponse `pulumi:"dnsResourceReferences"`
}

func GetDnsResourceReferenceByTarResourcesOutput(ctx *pulumi.Context, args GetDnsResourceReferenceByTarResourcesOutputArgs, opts ...pulumi.InvokeOption) GetDnsResourceReferenceByTarResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDnsResourceReferenceByTarResourcesResult, error) {
			args := v.(GetDnsResourceReferenceByTarResourcesArgs)
			r, err := GetDnsResourceReferenceByTarResources(ctx, &args, opts...)
			var s GetDnsResourceReferenceByTarResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDnsResourceReferenceByTarResourcesResultOutput)
}

type GetDnsResourceReferenceByTarResourcesOutputArgs struct {
	// A list of references to azure resources for which referencing dns records need to be queried.
	TargetResources SubResourceArrayInput `pulumi:"targetResources"`
}

func (GetDnsResourceReferenceByTarResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsResourceReferenceByTarResourcesArgs)(nil)).Elem()
}

// Represents the properties of the Dns Resource Reference Result.
type GetDnsResourceReferenceByTarResourcesResultOutput struct{ *pulumi.OutputState }

func (GetDnsResourceReferenceByTarResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsResourceReferenceByTarResourcesResult)(nil)).Elem()
}

func (o GetDnsResourceReferenceByTarResourcesResultOutput) ToGetDnsResourceReferenceByTarResourcesResultOutput() GetDnsResourceReferenceByTarResourcesResultOutput {
	return o
}

func (o GetDnsResourceReferenceByTarResourcesResultOutput) ToGetDnsResourceReferenceByTarResourcesResultOutputWithContext(ctx context.Context) GetDnsResourceReferenceByTarResourcesResultOutput {
	return o
}

func (o GetDnsResourceReferenceByTarResourcesResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDnsResourceReferenceByTarResourcesResult] {
	return pulumix.Output[GetDnsResourceReferenceByTarResourcesResult]{
		OutputState: o.OutputState,
	}
}

// The result of dns resource reference request. A list of dns resource references for each of the azure resource in the request
func (o GetDnsResourceReferenceByTarResourcesResultOutput) DnsResourceReferences() DnsResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v GetDnsResourceReferenceByTarResourcesResult) []DnsResourceReferenceResponse {
		return v.DnsResourceReferences
	}).(DnsResourceReferenceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDnsResourceReferenceByTarResourcesResultOutput{})
}
