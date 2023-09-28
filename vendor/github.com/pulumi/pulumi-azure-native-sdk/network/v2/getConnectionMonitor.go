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

// Gets a connection monitor by name.
// Azure REST API version: 2023-02-01.
func LookupConnectionMonitor(ctx *pulumi.Context, args *LookupConnectionMonitorArgs, opts ...pulumi.InvokeOption) (*LookupConnectionMonitorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectionMonitorResult
	err := ctx.Invoke("azure-native:network:getConnectionMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupConnectionMonitorArgs struct {
	// The name of the connection monitor.
	ConnectionMonitorName string `pulumi:"connectionMonitorName"`
	// The name of the Network Watcher resource.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the resource group containing Network Watcher.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Information about the connection monitor.
type LookupConnectionMonitorResult struct {
	// Determines if the connection monitor will start automatically once created.
	AutoStart *bool `pulumi:"autoStart"`
	// Type of connection monitor.
	ConnectionMonitorType string `pulumi:"connectionMonitorType"`
	// Describes the destination of connection monitor.
	Destination *ConnectionMonitorDestinationResponse `pulumi:"destination"`
	// List of connection monitor endpoints.
	Endpoints []ConnectionMonitorEndpointResponse `pulumi:"endpoints"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// ID of the connection monitor.
	Id string `pulumi:"id"`
	// Connection monitor location.
	Location *string `pulumi:"location"`
	// Monitoring interval in seconds.
	MonitoringIntervalInSeconds *int `pulumi:"monitoringIntervalInSeconds"`
	// The monitoring status of the connection monitor.
	MonitoringStatus string `pulumi:"monitoringStatus"`
	// Name of the connection monitor.
	Name string `pulumi:"name"`
	// Optional notes to be associated with the connection monitor.
	Notes *string `pulumi:"notes"`
	// List of connection monitor outputs.
	Outputs []ConnectionMonitorOutputResponse `pulumi:"outputs"`
	// The provisioning state of the connection monitor.
	ProvisioningState string `pulumi:"provisioningState"`
	// Describes the source of connection monitor.
	Source *ConnectionMonitorSourceResponse `pulumi:"source"`
	// The date and time when the connection monitor was started.
	StartTime string `pulumi:"startTime"`
	// Connection monitor tags.
	Tags map[string]string `pulumi:"tags"`
	// List of connection monitor test configurations.
	TestConfigurations []ConnectionMonitorTestConfigurationResponse `pulumi:"testConfigurations"`
	// List of connection monitor test groups.
	TestGroups []ConnectionMonitorTestGroupResponse `pulumi:"testGroups"`
	// Connection monitor type.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupConnectionMonitorResult
func (val *LookupConnectionMonitorResult) Defaults() *LookupConnectionMonitorResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AutoStart == nil {
		autoStart_ := true
		tmp.AutoStart = &autoStart_
	}
	if tmp.MonitoringIntervalInSeconds == nil {
		monitoringIntervalInSeconds_ := 60
		tmp.MonitoringIntervalInSeconds = &monitoringIntervalInSeconds_
	}
	return &tmp
}

func LookupConnectionMonitorOutput(ctx *pulumi.Context, args LookupConnectionMonitorOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionMonitorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionMonitorResult, error) {
			args := v.(LookupConnectionMonitorArgs)
			r, err := LookupConnectionMonitor(ctx, &args, opts...)
			var s LookupConnectionMonitorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionMonitorResultOutput)
}

type LookupConnectionMonitorOutputArgs struct {
	// The name of the connection monitor.
	ConnectionMonitorName pulumi.StringInput `pulumi:"connectionMonitorName"`
	// The name of the Network Watcher resource.
	NetworkWatcherName pulumi.StringInput `pulumi:"networkWatcherName"`
	// The name of the resource group containing Network Watcher.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupConnectionMonitorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionMonitorArgs)(nil)).Elem()
}

// Information about the connection monitor.
type LookupConnectionMonitorResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionMonitorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionMonitorResult)(nil)).Elem()
}

func (o LookupConnectionMonitorResultOutput) ToLookupConnectionMonitorResultOutput() LookupConnectionMonitorResultOutput {
	return o
}

func (o LookupConnectionMonitorResultOutput) ToLookupConnectionMonitorResultOutputWithContext(ctx context.Context) LookupConnectionMonitorResultOutput {
	return o
}

func (o LookupConnectionMonitorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupConnectionMonitorResult] {
	return pulumix.Output[LookupConnectionMonitorResult]{
		OutputState: o.OutputState,
	}
}

// Determines if the connection monitor will start automatically once created.
func (o LookupConnectionMonitorResultOutput) AutoStart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *bool { return v.AutoStart }).(pulumi.BoolPtrOutput)
}

// Type of connection monitor.
func (o LookupConnectionMonitorResultOutput) ConnectionMonitorType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.ConnectionMonitorType }).(pulumi.StringOutput)
}

// Describes the destination of connection monitor.
func (o LookupConnectionMonitorResultOutput) Destination() ConnectionMonitorDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *ConnectionMonitorDestinationResponse { return v.Destination }).(ConnectionMonitorDestinationResponsePtrOutput)
}

// List of connection monitor endpoints.
func (o LookupConnectionMonitorResultOutput) Endpoints() ConnectionMonitorEndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) []ConnectionMonitorEndpointResponse { return v.Endpoints }).(ConnectionMonitorEndpointResponseArrayOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupConnectionMonitorResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.Etag }).(pulumi.StringOutput)
}

// ID of the connection monitor.
func (o LookupConnectionMonitorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Connection monitor location.
func (o LookupConnectionMonitorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Monitoring interval in seconds.
func (o LookupConnectionMonitorResultOutput) MonitoringIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *int { return v.MonitoringIntervalInSeconds }).(pulumi.IntPtrOutput)
}

// The monitoring status of the connection monitor.
func (o LookupConnectionMonitorResultOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.MonitoringStatus }).(pulumi.StringOutput)
}

// Name of the connection monitor.
func (o LookupConnectionMonitorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional notes to be associated with the connection monitor.
func (o LookupConnectionMonitorResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

// List of connection monitor outputs.
func (o LookupConnectionMonitorResultOutput) Outputs() ConnectionMonitorOutputResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) []ConnectionMonitorOutputResponse { return v.Outputs }).(ConnectionMonitorOutputResponseArrayOutput)
}

// The provisioning state of the connection monitor.
func (o LookupConnectionMonitorResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Describes the source of connection monitor.
func (o LookupConnectionMonitorResultOutput) Source() ConnectionMonitorSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) *ConnectionMonitorSourceResponse { return v.Source }).(ConnectionMonitorSourceResponsePtrOutput)
}

// The date and time when the connection monitor was started.
func (o LookupConnectionMonitorResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// Connection monitor tags.
func (o LookupConnectionMonitorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// List of connection monitor test configurations.
func (o LookupConnectionMonitorResultOutput) TestConfigurations() ConnectionMonitorTestConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) []ConnectionMonitorTestConfigurationResponse {
		return v.TestConfigurations
	}).(ConnectionMonitorTestConfigurationResponseArrayOutput)
}

// List of connection monitor test groups.
func (o LookupConnectionMonitorResultOutput) TestGroups() ConnectionMonitorTestGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) []ConnectionMonitorTestGroupResponse { return v.TestGroups }).(ConnectionMonitorTestGroupResponseArrayOutput)
}

// Connection monitor type.
func (o LookupConnectionMonitorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionMonitorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionMonitorResultOutput{})
}
