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

// Retrieves a network manager security user configuration.
// Azure REST API version: 2022-04-01-preview.
//
// Other available API versions: 2021-05-01-preview.
func LookupSecurityUserConfiguration(ctx *pulumi.Context, args *LookupSecurityUserConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSecurityUserConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityUserConfigurationResult
	err := ctx.Invoke("azure-native:network:getSecurityUserConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityUserConfigurationArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines the security user configuration
type LookupSecurityUserConfigurationResult struct {
	// Flag if need to delete existing network security groups.
	DeleteExistingNSGs *string `pulumi:"deleteExistingNSGs"`
	// A description of the security user configuration.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupSecurityUserConfigurationOutput(ctx *pulumi.Context, args LookupSecurityUserConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityUserConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityUserConfigurationResult, error) {
			args := v.(LookupSecurityUserConfigurationArgs)
			r, err := LookupSecurityUserConfiguration(ctx, &args, opts...)
			var s LookupSecurityUserConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityUserConfigurationResultOutput)
}

type LookupSecurityUserConfigurationOutputArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSecurityUserConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityUserConfigurationArgs)(nil)).Elem()
}

// Defines the security user configuration
type LookupSecurityUserConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityUserConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityUserConfigurationResult)(nil)).Elem()
}

func (o LookupSecurityUserConfigurationResultOutput) ToLookupSecurityUserConfigurationResultOutput() LookupSecurityUserConfigurationResultOutput {
	return o
}

func (o LookupSecurityUserConfigurationResultOutput) ToLookupSecurityUserConfigurationResultOutputWithContext(ctx context.Context) LookupSecurityUserConfigurationResultOutput {
	return o
}

func (o LookupSecurityUserConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSecurityUserConfigurationResult] {
	return pulumix.Output[LookupSecurityUserConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// Flag if need to delete existing network security groups.
func (o LookupSecurityUserConfigurationResultOutput) DeleteExistingNSGs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) *string { return v.DeleteExistingNSGs }).(pulumi.StringPtrOutput)
}

// A description of the security user configuration.
func (o LookupSecurityUserConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupSecurityUserConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupSecurityUserConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupSecurityUserConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupSecurityUserConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o LookupSecurityUserConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupSecurityUserConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityUserConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityUserConfigurationResultOutput{})
}
