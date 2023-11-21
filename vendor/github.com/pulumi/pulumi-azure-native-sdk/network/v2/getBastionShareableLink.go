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

// Return the Bastion Shareable Links for all the VMs specified in the request.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01.
func GetBastionShareableLink(ctx *pulumi.Context, args *GetBastionShareableLinkArgs, opts ...pulumi.InvokeOption) (*GetBastionShareableLinkResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetBastionShareableLinkResult
	err := ctx.Invoke("azure-native:network:getBastionShareableLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBastionShareableLinkArgs struct {
	// The name of the Bastion Host.
	BastionHostName string `pulumi:"bastionHostName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// List of VM references.
	Vms []BastionShareableLink `pulumi:"vms"`
}

// Response for all the Bastion Shareable Link endpoints.
type GetBastionShareableLinkResult struct {
	// The URL to get the next set of results.
	NextLink *string `pulumi:"nextLink"`
	// List of Bastion Shareable Links for the request.
	Value []BastionShareableLinkResponse `pulumi:"value"`
}

func GetBastionShareableLinkOutput(ctx *pulumi.Context, args GetBastionShareableLinkOutputArgs, opts ...pulumi.InvokeOption) GetBastionShareableLinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBastionShareableLinkResult, error) {
			args := v.(GetBastionShareableLinkArgs)
			r, err := GetBastionShareableLink(ctx, &args, opts...)
			var s GetBastionShareableLinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBastionShareableLinkResultOutput)
}

type GetBastionShareableLinkOutputArgs struct {
	// The name of the Bastion Host.
	BastionHostName pulumi.StringInput `pulumi:"bastionHostName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// List of VM references.
	Vms BastionShareableLinkArrayInput `pulumi:"vms"`
}

func (GetBastionShareableLinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBastionShareableLinkArgs)(nil)).Elem()
}

// Response for all the Bastion Shareable Link endpoints.
type GetBastionShareableLinkResultOutput struct{ *pulumi.OutputState }

func (GetBastionShareableLinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBastionShareableLinkResult)(nil)).Elem()
}

func (o GetBastionShareableLinkResultOutput) ToGetBastionShareableLinkResultOutput() GetBastionShareableLinkResultOutput {
	return o
}

func (o GetBastionShareableLinkResultOutput) ToGetBastionShareableLinkResultOutputWithContext(ctx context.Context) GetBastionShareableLinkResultOutput {
	return o
}

func (o GetBastionShareableLinkResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetBastionShareableLinkResult] {
	return pulumix.Output[GetBastionShareableLinkResult]{
		OutputState: o.OutputState,
	}
}

// The URL to get the next set of results.
func (o GetBastionShareableLinkResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBastionShareableLinkResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// List of Bastion Shareable Links for the request.
func (o GetBastionShareableLinkResultOutput) Value() BastionShareableLinkResponseArrayOutput {
	return o.ApplyT(func(v GetBastionShareableLinkResult) []BastionShareableLinkResponse { return v.Value }).(BastionShareableLinkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBastionShareableLinkResultOutput{})
}
