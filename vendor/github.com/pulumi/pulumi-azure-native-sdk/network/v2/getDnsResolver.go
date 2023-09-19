// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets properties of a DNS resolver.
// Azure REST API version: 2022-07-01.
func LookupDnsResolver(ctx *pulumi.Context, args *LookupDnsResolverArgs, opts ...pulumi.InvokeOption) (*LookupDnsResolverResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDnsResolverResult
	err := ctx.Invoke("azure-native:network:getDnsResolver", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDnsResolverArgs struct {
	// The name of the DNS resolver.
	DnsResolverName string `pulumi:"dnsResolverName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes a DNS resolver.
type LookupDnsResolverResult struct {
	// The current status of the DNS resolver. This is a read-only property and any attempt to set this value will be ignored.
	DnsResolverState string `pulumi:"dnsResolverState"`
	// ETag of the DNS resolver.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The current provisioning state of the DNS resolver. This is a read-only property and any attempt to set this value will be ignored.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resourceGuid property of the DNS resolver resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The reference to the virtual network. This cannot be changed after creation.
	VirtualNetwork SubResourceResponse `pulumi:"virtualNetwork"`
}

func LookupDnsResolverOutput(ctx *pulumi.Context, args LookupDnsResolverOutputArgs, opts ...pulumi.InvokeOption) LookupDnsResolverResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDnsResolverResult, error) {
			args := v.(LookupDnsResolverArgs)
			r, err := LookupDnsResolver(ctx, &args, opts...)
			var s LookupDnsResolverResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDnsResolverResultOutput)
}

type LookupDnsResolverOutputArgs struct {
	// The name of the DNS resolver.
	DnsResolverName pulumi.StringInput `pulumi:"dnsResolverName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDnsResolverOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsResolverArgs)(nil)).Elem()
}

// Describes a DNS resolver.
type LookupDnsResolverResultOutput struct{ *pulumi.OutputState }

func (LookupDnsResolverResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsResolverResult)(nil)).Elem()
}

func (o LookupDnsResolverResultOutput) ToLookupDnsResolverResultOutput() LookupDnsResolverResultOutput {
	return o
}

func (o LookupDnsResolverResultOutput) ToLookupDnsResolverResultOutputWithContext(ctx context.Context) LookupDnsResolverResultOutput {
	return o
}

// The current status of the DNS resolver. This is a read-only property and any attempt to set this value will be ignored.
func (o LookupDnsResolverResultOutput) DnsResolverState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.DnsResolverState }).(pulumi.StringOutput)
}

// ETag of the DNS resolver.
func (o LookupDnsResolverResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDnsResolverResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupDnsResolverResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDnsResolverResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current provisioning state of the DNS resolver. This is a read-only property and any attempt to set this value will be ignored.
func (o LookupDnsResolverResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resourceGuid property of the DNS resolver resource.
func (o LookupDnsResolverResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDnsResolverResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDnsResolverResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDnsResolverResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) string { return v.Type }).(pulumi.StringOutput)
}

// The reference to the virtual network. This cannot be changed after creation.
func (o LookupDnsResolverResultOutput) VirtualNetwork() SubResourceResponseOutput {
	return o.ApplyT(func(v LookupDnsResolverResult) SubResourceResponse { return v.VirtualNetwork }).(SubResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDnsResolverResultOutput{})
}