// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the backend health for given combination of backend pool and http setting of the specified application gateway in a resource group.
// API Version: 2020-11-01.
func GetApplicationGatewayBackendHealthOnDemand(ctx *pulumi.Context, args *GetApplicationGatewayBackendHealthOnDemandArgs, opts ...pulumi.InvokeOption) (*GetApplicationGatewayBackendHealthOnDemandResult, error) {
	var rv GetApplicationGatewayBackendHealthOnDemandResult
	err := ctx.Invoke("azure-native:network:getApplicationGatewayBackendHealthOnDemand", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetApplicationGatewayBackendHealthOnDemandArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName string `pulumi:"applicationGatewayName"`
	// Reference to backend pool of application gateway to which probe request will be sent.
	BackendAddressPool *SubResource `pulumi:"backendAddressPool"`
	// Reference to backend http setting of application gateway to be used for test probe.
	BackendHttpSettings *SubResource `pulumi:"backendHttpSettings"`
	// Expands BackendAddressPool and BackendHttpSettings referenced in backend health.
	Expand *string `pulumi:"expand"`
	// Host name to send the probe to.
	Host *string `pulumi:"host"`
	// Criterion for classifying a healthy probe response.
	Match *ApplicationGatewayProbeHealthResponseMatch `pulumi:"match"`
	// Relative path of probe. Valid path starts from '/'. Probe is sent to <Protocol>://<host>:<port><path>.
	Path *string `pulumi:"path"`
	// Whether the host header should be picked from the backend http settings. Default value is false.
	PickHostNameFromBackendHttpSettings *bool `pulumi:"pickHostNameFromBackendHttpSettings"`
	// The protocol used for the probe.
	Protocol *string `pulumi:"protocol"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The probe timeout in seconds. Probe marked as failed if valid response is not received with this timeout period. Acceptable values are from 1 second to 86400 seconds.
	Timeout *int `pulumi:"timeout"`
}

// Result of on demand test probe.
type GetApplicationGatewayBackendHealthOnDemandResult struct {
	// Reference to an ApplicationGatewayBackendAddressPool resource.
	BackendAddressPool *ApplicationGatewayBackendAddressPoolResponse `pulumi:"backendAddressPool"`
	// Application gateway BackendHealthHttp settings.
	BackendHealthHttpSettings *ApplicationGatewayBackendHealthHttpSettingsResponse `pulumi:"backendHealthHttpSettings"`
}

func GetApplicationGatewayBackendHealthOnDemandOutput(ctx *pulumi.Context, args GetApplicationGatewayBackendHealthOnDemandOutputArgs, opts ...pulumi.InvokeOption) GetApplicationGatewayBackendHealthOnDemandResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetApplicationGatewayBackendHealthOnDemandResult, error) {
			args := v.(GetApplicationGatewayBackendHealthOnDemandArgs)
			r, err := GetApplicationGatewayBackendHealthOnDemand(ctx, &args, opts...)
			var s GetApplicationGatewayBackendHealthOnDemandResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetApplicationGatewayBackendHealthOnDemandResultOutput)
}

type GetApplicationGatewayBackendHealthOnDemandOutputArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName pulumi.StringInput `pulumi:"applicationGatewayName"`
	// Reference to backend pool of application gateway to which probe request will be sent.
	BackendAddressPool SubResourcePtrInput `pulumi:"backendAddressPool"`
	// Reference to backend http setting of application gateway to be used for test probe.
	BackendHttpSettings SubResourcePtrInput `pulumi:"backendHttpSettings"`
	// Expands BackendAddressPool and BackendHttpSettings referenced in backend health.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// Host name to send the probe to.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// Criterion for classifying a healthy probe response.
	Match ApplicationGatewayProbeHealthResponseMatchPtrInput `pulumi:"match"`
	// Relative path of probe. Valid path starts from '/'. Probe is sent to <Protocol>://<host>:<port><path>.
	Path pulumi.StringPtrInput `pulumi:"path"`
	// Whether the host header should be picked from the backend http settings. Default value is false.
	PickHostNameFromBackendHttpSettings pulumi.BoolPtrInput `pulumi:"pickHostNameFromBackendHttpSettings"`
	// The protocol used for the probe.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The probe timeout in seconds. Probe marked as failed if valid response is not received with this timeout period. Acceptable values are from 1 second to 86400 seconds.
	Timeout pulumi.IntPtrInput `pulumi:"timeout"`
}

func (GetApplicationGatewayBackendHealthOnDemandOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationGatewayBackendHealthOnDemandArgs)(nil)).Elem()
}

// Result of on demand test probe.
type GetApplicationGatewayBackendHealthOnDemandResultOutput struct{ *pulumi.OutputState }

func (GetApplicationGatewayBackendHealthOnDemandResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationGatewayBackendHealthOnDemandResult)(nil)).Elem()
}

func (o GetApplicationGatewayBackendHealthOnDemandResultOutput) ToGetApplicationGatewayBackendHealthOnDemandResultOutput() GetApplicationGatewayBackendHealthOnDemandResultOutput {
	return o
}

func (o GetApplicationGatewayBackendHealthOnDemandResultOutput) ToGetApplicationGatewayBackendHealthOnDemandResultOutputWithContext(ctx context.Context) GetApplicationGatewayBackendHealthOnDemandResultOutput {
	return o
}

// Reference to an ApplicationGatewayBackendAddressPool resource.
func (o GetApplicationGatewayBackendHealthOnDemandResultOutput) BackendAddressPool() ApplicationGatewayBackendAddressPoolResponsePtrOutput {
	return o.ApplyT(func(v GetApplicationGatewayBackendHealthOnDemandResult) *ApplicationGatewayBackendAddressPoolResponse {
		return v.BackendAddressPool
	}).(ApplicationGatewayBackendAddressPoolResponsePtrOutput)
}

// Application gateway BackendHealthHttp settings.
func (o GetApplicationGatewayBackendHealthOnDemandResultOutput) BackendHealthHttpSettings() ApplicationGatewayBackendHealthHttpSettingsResponsePtrOutput {
	return o.ApplyT(func(v GetApplicationGatewayBackendHealthOnDemandResult) *ApplicationGatewayBackendHealthHttpSettingsResponse {
		return v.BackendHealthHttpSettings
	}).(ApplicationGatewayBackendHealthHttpSettingsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApplicationGatewayBackendHealthOnDemandResultOutput{})
}
