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

// Get a specified connection created by this management group.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-04-01, 2023-05-01, 2023-06-01.
func LookupManagementGroupNetworkManagerConnection(ctx *pulumi.Context, args *LookupManagementGroupNetworkManagerConnectionArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupNetworkManagerConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagementGroupNetworkManagerConnectionResult
	err := ctx.Invoke("azure-native:network:getManagementGroupNetworkManagerConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupNetworkManagerConnectionArgs struct {
	// The management group Id which uniquely identify the Microsoft Azure management group.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// Name for the network manager connection.
	NetworkManagerConnectionName string `pulumi:"networkManagerConnectionName"`
}

// The Network Manager Connection resource
type LookupManagementGroupNetworkManagerConnectionResult struct {
	// A description of the network manager connection.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Network Manager Id.
	NetworkManagerId *string `pulumi:"networkManagerId"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupManagementGroupNetworkManagerConnectionOutput(ctx *pulumi.Context, args LookupManagementGroupNetworkManagerConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupManagementGroupNetworkManagerConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementGroupNetworkManagerConnectionResult, error) {
			args := v.(LookupManagementGroupNetworkManagerConnectionArgs)
			r, err := LookupManagementGroupNetworkManagerConnection(ctx, &args, opts...)
			var s LookupManagementGroupNetworkManagerConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementGroupNetworkManagerConnectionResultOutput)
}

type LookupManagementGroupNetworkManagerConnectionOutputArgs struct {
	// The management group Id which uniquely identify the Microsoft Azure management group.
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	// Name for the network manager connection.
	NetworkManagerConnectionName pulumi.StringInput `pulumi:"networkManagerConnectionName"`
}

func (LookupManagementGroupNetworkManagerConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupNetworkManagerConnectionArgs)(nil)).Elem()
}

// The Network Manager Connection resource
type LookupManagementGroupNetworkManagerConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupManagementGroupNetworkManagerConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupNetworkManagerConnectionResult)(nil)).Elem()
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) ToLookupManagementGroupNetworkManagerConnectionResultOutput() LookupManagementGroupNetworkManagerConnectionResultOutput {
	return o
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) ToLookupManagementGroupNetworkManagerConnectionResultOutputWithContext(ctx context.Context) LookupManagementGroupNetworkManagerConnectionResultOutput {
	return o
}

func (o LookupManagementGroupNetworkManagerConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupManagementGroupNetworkManagerConnectionResult] {
	return pulumix.Output[LookupManagementGroupNetworkManagerConnectionResult]{
		OutputState: o.OutputState,
	}
}

// A description of the network manager connection.
func (o LookupManagementGroupNetworkManagerConnectionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupManagementGroupNetworkManagerConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupManagementGroupNetworkManagerConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupManagementGroupNetworkManagerConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network Manager Id.
func (o LookupManagementGroupNetworkManagerConnectionResultOutput) NetworkManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) *string { return v.NetworkManagerId }).(pulumi.StringPtrOutput)
}

// The system metadata related to this resource.
func (o LookupManagementGroupNetworkManagerConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupManagementGroupNetworkManagerConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupNetworkManagerConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementGroupNetworkManagerConnectionResultOutput{})
}
