// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists active connectivity configurations in a network manager.
// API Version: 2021-02-01-preview.
func ListActiveConnectivityConfiguration(ctx *pulumi.Context, args *ListActiveConnectivityConfigurationArgs, opts ...pulumi.InvokeOption) (*ListActiveConnectivityConfigurationResult, error) {
	var rv ListActiveConnectivityConfigurationResult
	err := ctx.Invoke("azure-native:network:listActiveConnectivityConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveConnectivityConfigurationArgs struct {
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// List of regions.
	Regions []string `pulumi:"regions"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
}

// Result of the request to list active connectivity configurations. It contains a list of active connectivity configurations and a skiptoken to get the next set of results.
type ListActiveConnectivityConfigurationResult struct {
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// Gets a page of active connectivity configurations.
	Value []ActiveConnectivityConfigurationResponse `pulumi:"value"`
}

func ListActiveConnectivityConfigurationOutput(ctx *pulumi.Context, args ListActiveConnectivityConfigurationOutputArgs, opts ...pulumi.InvokeOption) ListActiveConnectivityConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActiveConnectivityConfigurationResult, error) {
			args := v.(ListActiveConnectivityConfigurationArgs)
			r, err := ListActiveConnectivityConfiguration(ctx, &args, opts...)
			var s ListActiveConnectivityConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActiveConnectivityConfigurationResultOutput)
}

type ListActiveConnectivityConfigurationOutputArgs struct {
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// List of regions.
	Regions pulumi.StringArrayInput `pulumi:"regions"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListActiveConnectivityConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveConnectivityConfigurationArgs)(nil)).Elem()
}

// Result of the request to list active connectivity configurations. It contains a list of active connectivity configurations and a skiptoken to get the next set of results.
type ListActiveConnectivityConfigurationResultOutput struct{ *pulumi.OutputState }

func (ListActiveConnectivityConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveConnectivityConfigurationResult)(nil)).Elem()
}

func (o ListActiveConnectivityConfigurationResultOutput) ToListActiveConnectivityConfigurationResultOutput() ListActiveConnectivityConfigurationResultOutput {
	return o
}

func (o ListActiveConnectivityConfigurationResultOutput) ToListActiveConnectivityConfigurationResultOutputWithContext(ctx context.Context) ListActiveConnectivityConfigurationResultOutput {
	return o
}

// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
func (o ListActiveConnectivityConfigurationResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListActiveConnectivityConfigurationResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

// Gets a page of active connectivity configurations.
func (o ListActiveConnectivityConfigurationResultOutput) Value() ActiveConnectivityConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListActiveConnectivityConfigurationResult) []ActiveConnectivityConfigurationResponse {
		return v.Value
	}).(ActiveConnectivityConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActiveConnectivityConfigurationResultOutput{})
}
