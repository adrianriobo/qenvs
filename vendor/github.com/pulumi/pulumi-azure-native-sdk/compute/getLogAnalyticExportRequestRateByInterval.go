// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Export logs that show Api requests made by this subscription in the given time window to show throttling activities.
// API Version: 2020-12-01.
func GetLogAnalyticExportRequestRateByInterval(ctx *pulumi.Context, args *GetLogAnalyticExportRequestRateByIntervalArgs, opts ...pulumi.InvokeOption) (*GetLogAnalyticExportRequestRateByIntervalResult, error) {
	var rv GetLogAnalyticExportRequestRateByIntervalResult
	err := ctx.Invoke("azure-native:compute:getLogAnalyticExportRequestRateByInterval", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLogAnalyticExportRequestRateByIntervalArgs struct {
	// SAS Uri of the logging blob container to which LogAnalytics Api writes output logs to.
	BlobContainerSasUri string `pulumi:"blobContainerSasUri"`
	// From time of the query
	FromTime string `pulumi:"fromTime"`
	// Group query result by Client Application ID.
	GroupByClientApplicationId *bool `pulumi:"groupByClientApplicationId"`
	// Group query result by Operation Name.
	GroupByOperationName *bool `pulumi:"groupByOperationName"`
	// Group query result by Resource Name.
	GroupByResourceName *bool `pulumi:"groupByResourceName"`
	// Group query result by Throttle Policy applied.
	GroupByThrottlePolicy *bool `pulumi:"groupByThrottlePolicy"`
	// Group query result by User Agent.
	GroupByUserAgent *bool `pulumi:"groupByUserAgent"`
	// Interval value in minutes used to create LogAnalytics call rate logs.
	IntervalLength IntervalInMins `pulumi:"intervalLength"`
	// The location upon which virtual-machine-sizes is queried.
	Location string `pulumi:"location"`
	// To time of the query
	ToTime string `pulumi:"toTime"`
}

// LogAnalytics operation status response
type GetLogAnalyticExportRequestRateByIntervalResult struct {
	// LogAnalyticsOutput
	Properties LogAnalyticsOutputResponse `pulumi:"properties"`
}

func GetLogAnalyticExportRequestRateByIntervalOutput(ctx *pulumi.Context, args GetLogAnalyticExportRequestRateByIntervalOutputArgs, opts ...pulumi.InvokeOption) GetLogAnalyticExportRequestRateByIntervalResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogAnalyticExportRequestRateByIntervalResult, error) {
			args := v.(GetLogAnalyticExportRequestRateByIntervalArgs)
			r, err := GetLogAnalyticExportRequestRateByInterval(ctx, &args, opts...)
			var s GetLogAnalyticExportRequestRateByIntervalResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogAnalyticExportRequestRateByIntervalResultOutput)
}

type GetLogAnalyticExportRequestRateByIntervalOutputArgs struct {
	// SAS Uri of the logging blob container to which LogAnalytics Api writes output logs to.
	BlobContainerSasUri pulumi.StringInput `pulumi:"blobContainerSasUri"`
	// From time of the query
	FromTime pulumi.StringInput `pulumi:"fromTime"`
	// Group query result by Client Application ID.
	GroupByClientApplicationId pulumi.BoolPtrInput `pulumi:"groupByClientApplicationId"`
	// Group query result by Operation Name.
	GroupByOperationName pulumi.BoolPtrInput `pulumi:"groupByOperationName"`
	// Group query result by Resource Name.
	GroupByResourceName pulumi.BoolPtrInput `pulumi:"groupByResourceName"`
	// Group query result by Throttle Policy applied.
	GroupByThrottlePolicy pulumi.BoolPtrInput `pulumi:"groupByThrottlePolicy"`
	// Group query result by User Agent.
	GroupByUserAgent pulumi.BoolPtrInput `pulumi:"groupByUserAgent"`
	// Interval value in minutes used to create LogAnalytics call rate logs.
	IntervalLength IntervalInMinsInput `pulumi:"intervalLength"`
	// The location upon which virtual-machine-sizes is queried.
	Location pulumi.StringInput `pulumi:"location"`
	// To time of the query
	ToTime pulumi.StringInput `pulumi:"toTime"`
}

func (GetLogAnalyticExportRequestRateByIntervalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogAnalyticExportRequestRateByIntervalArgs)(nil)).Elem()
}

// LogAnalytics operation status response
type GetLogAnalyticExportRequestRateByIntervalResultOutput struct{ *pulumi.OutputState }

func (GetLogAnalyticExportRequestRateByIntervalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogAnalyticExportRequestRateByIntervalResult)(nil)).Elem()
}

func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) ToGetLogAnalyticExportRequestRateByIntervalResultOutput() GetLogAnalyticExportRequestRateByIntervalResultOutput {
	return o
}

func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) ToGetLogAnalyticExportRequestRateByIntervalResultOutputWithContext(ctx context.Context) GetLogAnalyticExportRequestRateByIntervalResultOutput {
	return o
}

// LogAnalyticsOutput
func (o GetLogAnalyticExportRequestRateByIntervalResultOutput) Properties() LogAnalyticsOutputResponseOutput {
	return o.ApplyT(func(v GetLogAnalyticExportRequestRateByIntervalResult) LogAnalyticsOutputResponse {
		return v.Properties
	}).(LogAnalyticsOutputResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogAnalyticExportRequestRateByIntervalResultOutput{})
}
