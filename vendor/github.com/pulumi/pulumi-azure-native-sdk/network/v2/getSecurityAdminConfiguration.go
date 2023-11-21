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

// Retrieves a network manager security admin configuration.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2021-05-01-preview, 2023-04-01, 2023-05-01, 2023-06-01.
func LookupSecurityAdminConfiguration(ctx *pulumi.Context, args *LookupSecurityAdminConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSecurityAdminConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityAdminConfigurationResult
	err := ctx.Invoke("azure-native:network:getSecurityAdminConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityAdminConfigurationArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Defines the security admin configuration
type LookupSecurityAdminConfigurationResult struct {
	// Enum list of network intent policy based services.
	ApplyOnNetworkIntentPolicyBasedServices []string `pulumi:"applyOnNetworkIntentPolicyBasedServices"`
	// A description of the security configuration.
	Description *string `pulumi:"description"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupSecurityAdminConfigurationOutput(ctx *pulumi.Context, args LookupSecurityAdminConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityAdminConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityAdminConfigurationResult, error) {
			args := v.(LookupSecurityAdminConfigurationArgs)
			r, err := LookupSecurityAdminConfiguration(ctx, &args, opts...)
			var s LookupSecurityAdminConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityAdminConfigurationResultOutput)
}

type LookupSecurityAdminConfigurationOutputArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringInput `pulumi:"configurationName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSecurityAdminConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityAdminConfigurationArgs)(nil)).Elem()
}

// Defines the security admin configuration
type LookupSecurityAdminConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityAdminConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityAdminConfigurationResult)(nil)).Elem()
}

func (o LookupSecurityAdminConfigurationResultOutput) ToLookupSecurityAdminConfigurationResultOutput() LookupSecurityAdminConfigurationResultOutput {
	return o
}

func (o LookupSecurityAdminConfigurationResultOutput) ToLookupSecurityAdminConfigurationResultOutputWithContext(ctx context.Context) LookupSecurityAdminConfigurationResultOutput {
	return o
}

func (o LookupSecurityAdminConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSecurityAdminConfigurationResult] {
	return pulumix.Output[LookupSecurityAdminConfigurationResult]{
		OutputState: o.OutputState,
	}
}

// Enum list of network intent policy based services.
func (o LookupSecurityAdminConfigurationResultOutput) ApplyOnNetworkIntentPolicyBasedServices() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) []string {
		return v.ApplyOnNetworkIntentPolicyBasedServices
	}).(pulumi.StringArrayOutput)
}

// A description of the security configuration.
func (o LookupSecurityAdminConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupSecurityAdminConfigurationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupSecurityAdminConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupSecurityAdminConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupSecurityAdminConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o LookupSecurityAdminConfigurationResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The system metadata related to this resource.
func (o LookupSecurityAdminConfigurationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o LookupSecurityAdminConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityAdminConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityAdminConfigurationResultOutput{})
}
