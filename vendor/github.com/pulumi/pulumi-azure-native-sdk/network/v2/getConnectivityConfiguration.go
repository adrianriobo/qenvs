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

// Gets a Network Connectivity Configuration, specified by the resource group, network manager name, and connectivity Configuration name
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2021-02-01-preview, 2021-05-01-preview, 2023-04-01, 2023-05-01, 2023-06-01.
func LookupConnectivityConfiguration(ctx *pulumi.Context, args *LookupConnectivityConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupConnectivityConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectivityConfigurationResult
	err := ctx.Invoke("azure-native:network:getConnectivityConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectivityConfigurationArgs struct {
	// The name of the network manager connectivity configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The network manager connectivity configuration resource
type LookupConnectivityConfigurationResult struct {
	// Groups for configuration
	AppliesToGroups []ConnectivityGroupItemResponse `pulumi:"appliesToGroups"`
	// Connectivity topology type.
	ConnectivityTopology string `pulumi:"connectivityTopology"`
	// Flag if need to remove current existing peerings.
	DeleteExistingPeering *string `pulumi:"deleteExistingPeering"`
	// A description of the connectivity configuration.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// List of hubItems
	Hubs []HubResponse `pulumi:"hubs"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Flag if global mesh is supported.
	IsGlobal *string `pulumi:"isGlobal"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the connectivity configuration resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupConnectivityConfigurationOutput(ctx *pulumi.Context, args LookupConnectivityConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupConnectivityConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectivityConfigurationResult, error) {
			args := v.(LookupConnectivityConfigurationArgs)
			r, err := LookupConnectivityConfiguration(ctx, &args, opts...)
			var s LookupConnectivityConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectivityConfigurationResultOutput)
}

type LookupConnectivityConfigurationOutputArgs struct {
	// The name of the network manager connectivity configuration.
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectivityConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectivityConfigurationArgs)(nil)).Elem()
}

// The network manager connectivity configuration resource
type LookupConnectivityConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupConnectivityConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectivityConfigurationResult)(nil)).Elem()
}

func (o LookupConnectivityConfigurationResultOutput) ToLookupConnectivityConfigurationResultOutput() LookupConnectivityConfigurationResultOutput {
	return o
}

func (o LookupConnectivityConfigurationResultOutput) ToLookupConnectivityConfigurationResultOutputWithContext(ctx context.Context) LookupConnectivityConfigurationResultOutput {
	return o
}

func (o LookupConnectivityConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConnectivityConfigurationResult] {
	return pulumix.Output[LookupConnectivityConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// Groups for configuration
func (o LookupConnectivityConfigurationResultOutput) AppliesToGroups() ConnectivityGroupItemResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) []ConnectivityGroupItemResponse {
		return v.AppliesToGroups
	}).(ConnectivityGroupItemResponseArrayOutput)
}

// Connectivity topology type.
func (o LookupConnectivityConfigurationResultOutput) ConnectivityTopology() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.ConnectivityTopology }).(pulumi.StringOutput)
}

// Flag if need to remove current existing peerings.
func (o LookupConnectivityConfigurationResultOutput) DeleteExistingPeering() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) *string { return v.DeleteExistingPeering }).(pulumi.StringPtrOutput)
}

// A description of the connectivity configuration.
func (o LookupConnectivityConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupConnectivityConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// List of hubItems
func (o LookupConnectivityConfigurationResultOutput) Hubs() HubResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) []HubResponse { return v.Hubs }).(HubResponseArrayOutput)
}

// Resource ID.
func (o LookupConnectivityConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Flag if global mesh is supported.
func (o LookupConnectivityConfigurationResultOutput) IsGlobal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) *string { return v.IsGlobal }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupConnectivityConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the connectivity configuration resource.
func (o LookupConnectivityConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o LookupConnectivityConfigurationResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o LookupConnectivityConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupConnectivityConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectivityConfigurationResultOutput{})
}
