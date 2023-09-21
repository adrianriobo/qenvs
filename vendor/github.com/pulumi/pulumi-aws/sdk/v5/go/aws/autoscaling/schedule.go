// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AutoScaling Schedule resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/autoscaling"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			foobarGroup, err := autoscaling.NewGroup(ctx, "foobarGroup", &autoscaling.GroupArgs{
//				AvailabilityZones: pulumi.StringArray{
//					pulumi.String("us-west-2a"),
//				},
//				MaxSize:                pulumi.Int(1),
//				MinSize:                pulumi.Int(1),
//				HealthCheckGracePeriod: pulumi.Int(300),
//				HealthCheckType:        pulumi.String("ELB"),
//				ForceDelete:            pulumi.Bool(true),
//				TerminationPolicies: pulumi.StringArray{
//					pulumi.String("OldestInstance"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = autoscaling.NewSchedule(ctx, "foobarSchedule", &autoscaling.ScheduleArgs{
//				ScheduledActionName:  pulumi.String("foobar"),
//				MinSize:              pulumi.Int(0),
//				MaxSize:              pulumi.Int(1),
//				DesiredCapacity:      pulumi.Int(0),
//				StartTime:            pulumi.String("2016-12-11T18:00:00Z"),
//				EndTime:              pulumi.String("2016-12-12T06:00:00Z"),
//				AutoscalingGroupName: foobarGroup.Name,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// AutoScaling ScheduledAction can be imported using the `auto-scaling-group-name` and `scheduled-action-name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:autoscaling/schedule:Schedule resource-name auto-scaling-group-name/scheduled-action-name
//
// ```
type Schedule struct {
	pulumi.CustomResourceState

	// ARN assigned by AWS to the autoscaling schedule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the Auto Scaling group.
	AutoscalingGroupName pulumi.StringOutput `pulumi:"autoscalingGroupName"`
	// The initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. Set to `-1` if you don't want to change the desired capacity at the scheduled time. Defaults to `0`.
	DesiredCapacity pulumi.IntOutput `pulumi:"desiredCapacity"`
	// The date and time for the recurring schedule to end, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The maximum size of the Auto Scaling group. Set to `-1` if you don't want to change the maximum size at the scheduled time. Defaults to `0`.
	MaxSize pulumi.IntOutput `pulumi:"maxSize"`
	// The minimum size of the Auto Scaling group. Set to `-1` if you don't want to change the minimum size at the scheduled time. Defaults to `0`.
	MinSize pulumi.IntOutput `pulumi:"minSize"`
	// The recurring schedule for this action specified using the Unix cron syntax format.
	Recurrence pulumi.StringOutput `pulumi:"recurrence"`
	// The name of this scaling action.
	//
	// The following arguments are optional:
	ScheduledActionName pulumi.StringOutput `pulumi:"scheduledActionName"`
	// The date and time for the recurring schedule to start, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Specifies the time zone for a cron expression. Valid values are the canonical names of the IANA time zones (such as `Etc/GMT+9` or `Pacific/Tahiti`).
	//
	// > **NOTE:** When `startTime` and `endTime` are specified with `recurrence` , they form the boundaries of when the recurring action will start and stop.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoscalingGroupName == nil {
		return nil, errors.New("invalid value for required argument 'AutoscalingGroupName'")
	}
	if args.ScheduledActionName == nil {
		return nil, errors.New("invalid value for required argument 'ScheduledActionName'")
	}
	var resource Schedule
	err := ctx.RegisterResource("aws:autoscaling/schedule:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchedule gets an existing Schedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("aws:autoscaling/schedule:Schedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schedule resources.
type scheduleState struct {
	// ARN assigned by AWS to the autoscaling schedule.
	Arn *string `pulumi:"arn"`
	// The name of the Auto Scaling group.
	AutoscalingGroupName *string `pulumi:"autoscalingGroupName"`
	// The initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. Set to `-1` if you don't want to change the desired capacity at the scheduled time. Defaults to `0`.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// The date and time for the recurring schedule to end, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	EndTime *string `pulumi:"endTime"`
	// The maximum size of the Auto Scaling group. Set to `-1` if you don't want to change the maximum size at the scheduled time. Defaults to `0`.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum size of the Auto Scaling group. Set to `-1` if you don't want to change the minimum size at the scheduled time. Defaults to `0`.
	MinSize *int `pulumi:"minSize"`
	// The recurring schedule for this action specified using the Unix cron syntax format.
	Recurrence *string `pulumi:"recurrence"`
	// The name of this scaling action.
	//
	// The following arguments are optional:
	ScheduledActionName *string `pulumi:"scheduledActionName"`
	// The date and time for the recurring schedule to start, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	StartTime *string `pulumi:"startTime"`
	// Specifies the time zone for a cron expression. Valid values are the canonical names of the IANA time zones (such as `Etc/GMT+9` or `Pacific/Tahiti`).
	//
	// > **NOTE:** When `startTime` and `endTime` are specified with `recurrence` , they form the boundaries of when the recurring action will start and stop.
	TimeZone *string `pulumi:"timeZone"`
}

type ScheduleState struct {
	// ARN assigned by AWS to the autoscaling schedule.
	Arn pulumi.StringPtrInput
	// The name of the Auto Scaling group.
	AutoscalingGroupName pulumi.StringPtrInput
	// The initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. Set to `-1` if you don't want to change the desired capacity at the scheduled time. Defaults to `0`.
	DesiredCapacity pulumi.IntPtrInput
	// The date and time for the recurring schedule to end, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	EndTime pulumi.StringPtrInput
	// The maximum size of the Auto Scaling group. Set to `-1` if you don't want to change the maximum size at the scheduled time. Defaults to `0`.
	MaxSize pulumi.IntPtrInput
	// The minimum size of the Auto Scaling group. Set to `-1` if you don't want to change the minimum size at the scheduled time. Defaults to `0`.
	MinSize pulumi.IntPtrInput
	// The recurring schedule for this action specified using the Unix cron syntax format.
	Recurrence pulumi.StringPtrInput
	// The name of this scaling action.
	//
	// The following arguments are optional:
	ScheduledActionName pulumi.StringPtrInput
	// The date and time for the recurring schedule to start, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	StartTime pulumi.StringPtrInput
	// Specifies the time zone for a cron expression. Valid values are the canonical names of the IANA time zones (such as `Etc/GMT+9` or `Pacific/Tahiti`).
	//
	// > **NOTE:** When `startTime` and `endTime` are specified with `recurrence` , they form the boundaries of when the recurring action will start and stop.
	TimeZone pulumi.StringPtrInput
}

func (ScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleState)(nil)).Elem()
}

type scheduleArgs struct {
	// The name of the Auto Scaling group.
	AutoscalingGroupName string `pulumi:"autoscalingGroupName"`
	// The initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. Set to `-1` if you don't want to change the desired capacity at the scheduled time. Defaults to `0`.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// The date and time for the recurring schedule to end, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	EndTime *string `pulumi:"endTime"`
	// The maximum size of the Auto Scaling group. Set to `-1` if you don't want to change the maximum size at the scheduled time. Defaults to `0`.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum size of the Auto Scaling group. Set to `-1` if you don't want to change the minimum size at the scheduled time. Defaults to `0`.
	MinSize *int `pulumi:"minSize"`
	// The recurring schedule for this action specified using the Unix cron syntax format.
	Recurrence *string `pulumi:"recurrence"`
	// The name of this scaling action.
	//
	// The following arguments are optional:
	ScheduledActionName string `pulumi:"scheduledActionName"`
	// The date and time for the recurring schedule to start, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	StartTime *string `pulumi:"startTime"`
	// Specifies the time zone for a cron expression. Valid values are the canonical names of the IANA time zones (such as `Etc/GMT+9` or `Pacific/Tahiti`).
	//
	// > **NOTE:** When `startTime` and `endTime` are specified with `recurrence` , they form the boundaries of when the recurring action will start and stop.
	TimeZone *string `pulumi:"timeZone"`
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	// The name of the Auto Scaling group.
	AutoscalingGroupName pulumi.StringInput
	// The initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. Set to `-1` if you don't want to change the desired capacity at the scheduled time. Defaults to `0`.
	DesiredCapacity pulumi.IntPtrInput
	// The date and time for the recurring schedule to end, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	EndTime pulumi.StringPtrInput
	// The maximum size of the Auto Scaling group. Set to `-1` if you don't want to change the maximum size at the scheduled time. Defaults to `0`.
	MaxSize pulumi.IntPtrInput
	// The minimum size of the Auto Scaling group. Set to `-1` if you don't want to change the minimum size at the scheduled time. Defaults to `0`.
	MinSize pulumi.IntPtrInput
	// The recurring schedule for this action specified using the Unix cron syntax format.
	Recurrence pulumi.StringPtrInput
	// The name of this scaling action.
	//
	// The following arguments are optional:
	ScheduledActionName pulumi.StringInput
	// The date and time for the recurring schedule to start, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
	StartTime pulumi.StringPtrInput
	// Specifies the time zone for a cron expression. Valid values are the canonical names of the IANA time zones (such as `Etc/GMT+9` or `Pacific/Tahiti`).
	//
	// > **NOTE:** When `startTime` and `endTime` are specified with `recurrence` , they form the boundaries of when the recurring action will start and stop.
	TimeZone pulumi.StringPtrInput
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleArgs)(nil)).Elem()
}

type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput
}

func (*Schedule) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (i *Schedule) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i *Schedule) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

// ScheduleArrayInput is an input type that accepts ScheduleArray and ScheduleArrayOutput values.
// You can construct a concrete instance of `ScheduleArrayInput` via:
//
//	ScheduleArray{ ScheduleArgs{...} }
type ScheduleArrayInput interface {
	pulumi.Input

	ToScheduleArrayOutput() ScheduleArrayOutput
	ToScheduleArrayOutputWithContext(context.Context) ScheduleArrayOutput
}

type ScheduleArray []ScheduleInput

func (ScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Schedule)(nil)).Elem()
}

func (i ScheduleArray) ToScheduleArrayOutput() ScheduleArrayOutput {
	return i.ToScheduleArrayOutputWithContext(context.Background())
}

func (i ScheduleArray) ToScheduleArrayOutputWithContext(ctx context.Context) ScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleArrayOutput)
}

// ScheduleMapInput is an input type that accepts ScheduleMap and ScheduleMapOutput values.
// You can construct a concrete instance of `ScheduleMapInput` via:
//
//	ScheduleMap{ "key": ScheduleArgs{...} }
type ScheduleMapInput interface {
	pulumi.Input

	ToScheduleMapOutput() ScheduleMapOutput
	ToScheduleMapOutputWithContext(context.Context) ScheduleMapOutput
}

type ScheduleMap map[string]ScheduleInput

func (ScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Schedule)(nil)).Elem()
}

func (i ScheduleMap) ToScheduleMapOutput() ScheduleMapOutput {
	return i.ToScheduleMapOutputWithContext(context.Background())
}

func (i ScheduleMap) ToScheduleMapOutputWithContext(ctx context.Context) ScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleMapOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

// ARN assigned by AWS to the autoscaling schedule.
func (o ScheduleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The name of the Auto Scaling group.
func (o ScheduleOutput) AutoscalingGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.AutoscalingGroupName }).(pulumi.StringOutput)
}

// The initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain. Set to `-1` if you don't want to change the desired capacity at the scheduled time. Defaults to `0`.
func (o ScheduleOutput) DesiredCapacity() pulumi.IntOutput {
	return o.ApplyT(func(v *Schedule) pulumi.IntOutput { return v.DesiredCapacity }).(pulumi.IntOutput)
}

// The date and time for the recurring schedule to end, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
func (o ScheduleOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// The maximum size of the Auto Scaling group. Set to `-1` if you don't want to change the maximum size at the scheduled time. Defaults to `0`.
func (o ScheduleOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Schedule) pulumi.IntOutput { return v.MaxSize }).(pulumi.IntOutput)
}

// The minimum size of the Auto Scaling group. Set to `-1` if you don't want to change the minimum size at the scheduled time. Defaults to `0`.
func (o ScheduleOutput) MinSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Schedule) pulumi.IntOutput { return v.MinSize }).(pulumi.IntOutput)
}

// The recurring schedule for this action specified using the Unix cron syntax format.
func (o ScheduleOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Recurrence }).(pulumi.StringOutput)
}

// The name of this scaling action.
//
// The following arguments are optional:
func (o ScheduleOutput) ScheduledActionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.ScheduledActionName }).(pulumi.StringOutput)
}

// The date and time for the recurring schedule to start, in UTC with the format `"YYYY-MM-DDThh:mm:ssZ"` (e.g. `"2021-06-01T00:00:00Z"`).
func (o ScheduleOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// Specifies the time zone for a cron expression. Valid values are the canonical names of the IANA time zones (such as `Etc/GMT+9` or `Pacific/Tahiti`).
//
// > **NOTE:** When `startTime` and `endTime` are specified with `recurrence` , they form the boundaries of when the recurring action will start and stop.
func (o ScheduleOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

type ScheduleArrayOutput struct{ *pulumi.OutputState }

func (ScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Schedule)(nil)).Elem()
}

func (o ScheduleArrayOutput) ToScheduleArrayOutput() ScheduleArrayOutput {
	return o
}

func (o ScheduleArrayOutput) ToScheduleArrayOutputWithContext(ctx context.Context) ScheduleArrayOutput {
	return o
}

func (o ScheduleArrayOutput) Index(i pulumi.IntInput) ScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Schedule {
		return vs[0].([]*Schedule)[vs[1].(int)]
	}).(ScheduleOutput)
}

type ScheduleMapOutput struct{ *pulumi.OutputState }

func (ScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Schedule)(nil)).Elem()
}

func (o ScheduleMapOutput) ToScheduleMapOutput() ScheduleMapOutput {
	return o
}

func (o ScheduleMapOutput) ToScheduleMapOutputWithContext(ctx context.Context) ScheduleMapOutput {
	return o
}

func (o ScheduleMapOutput) MapIndex(k pulumi.StringInput) ScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Schedule {
		return vs[0].(map[string]*Schedule)[vs[1].(string)]
	}).(ScheduleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleInput)(nil)).Elem(), &Schedule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleArrayInput)(nil)).Elem(), ScheduleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleMapInput)(nil)).Elem(), ScheduleMap{})
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(ScheduleArrayOutput{})
	pulumi.RegisterOutputType(ScheduleMapOutput{})
}
