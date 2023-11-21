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

// Lists active connectivity configurations in a network manager.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2021-05-01-preview, 2023-04-01, 2023-05-01, 2023-06-01.
func ListActiveConnectivityConfigurations(ctx *pulumi.Context, args *ListActiveConnectivityConfigurationsArgs, opts ...pulumi.InvokeOption) (*ListActiveConnectivityConfigurationsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListActiveConnectivityConfigurationsResult
	err := ctx.Invoke("azure-native:network:listActiveConnectivityConfigurations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveConnectivityConfigurationsArgs struct {
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// List of regions.
	Regions []string `pulumi:"regions"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// An optional query parameter which specifies the maximum number of records to be returned by the server.
	Top *int `pulumi:"top"`
}

// Result of the request to list active connectivity configurations. It contains a list of active connectivity configurations and a skiptoken to get the next set of results.
type ListActiveConnectivityConfigurationsResult struct {
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// Gets a page of active connectivity configurations.
	Value []ActiveConnectivityConfigurationResponse `pulumi:"value"`
}

func ListActiveConnectivityConfigurationsOutput(ctx *pulumi.Context, args ListActiveConnectivityConfigurationsOutputArgs, opts ...pulumi.InvokeOption) ListActiveConnectivityConfigurationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListActiveConnectivityConfigurationsResult, error) {
			args := v.(ListActiveConnectivityConfigurationsArgs)
			r, err := ListActiveConnectivityConfigurations(ctx, &args, opts...)
			var s ListActiveConnectivityConfigurationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListActiveConnectivityConfigurationsResultOutput)
}

type ListActiveConnectivityConfigurationsOutputArgs struct {
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// List of regions.
	Regions pulumi.StringArrayInput `pulumi:"regions"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
	// An optional query parameter which specifies the maximum number of records to be returned by the server.
	Top pulumi.IntPtrInput `pulumi:"top"`
}

func (ListActiveConnectivityConfigurationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveConnectivityConfigurationsArgs)(nil)).Elem()
}

// Result of the request to list active connectivity configurations. It contains a list of active connectivity configurations and a skiptoken to get the next set of results.
type ListActiveConnectivityConfigurationsResultOutput struct{ *pulumi.OutputState }

func (ListActiveConnectivityConfigurationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListActiveConnectivityConfigurationsResult)(nil)).Elem()
}

func (o ListActiveConnectivityConfigurationsResultOutput) ToListActiveConnectivityConfigurationsResultOutput() ListActiveConnectivityConfigurationsResultOutput {
	return o
}

func (o ListActiveConnectivityConfigurationsResultOutput) ToListActiveConnectivityConfigurationsResultOutputWithContext(ctx context.Context) ListActiveConnectivityConfigurationsResultOutput {
	return o
}

func (o ListActiveConnectivityConfigurationsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListActiveConnectivityConfigurationsResult] {
	return pulumix.Output[ListActiveConnectivityConfigurationsResult]{
		OutputState: o.OutputState,
	}
}

// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
func (o ListActiveConnectivityConfigurationsResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListActiveConnectivityConfigurationsResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

// Gets a page of active connectivity configurations.
func (o ListActiveConnectivityConfigurationsResultOutput) Value() ActiveConnectivityConfigurationResponseArrayOutput {
	return o.ApplyT(func(v ListActiveConnectivityConfigurationsResult) []ActiveConnectivityConfigurationResponse {
		return v.Value
	}).(ActiveConnectivityConfigurationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListActiveConnectivityConfigurationsResultOutput{})
}
