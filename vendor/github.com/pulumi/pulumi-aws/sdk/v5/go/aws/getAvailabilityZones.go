// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Availability Zones data source allows access to the list of AWS
// Availability Zones which can be accessed by an AWS account within the region
// configured in the provider.
//
// This is different from the `getAvailabilityZone` (singular) data source,
// which provides some details about a specific availability zone.
//
// > When [Local Zones](https://aws.amazon.com/about-aws/global-infrastructure/localzones/) are enabled in a region, by default the API and this data source include both Local Zones and Availability Zones. To return only Availability Zones, see the example section below.
//
// ## Example Usage
// ### By State
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			available, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
//				State: pulumi.StringRef("available"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewSubnet(ctx, "primary", &ec2.SubnetArgs{
//				AvailabilityZone: *pulumi.String(available.Names[0]),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewSubnet(ctx, "secondary", &ec2.SubnetArgs{
//				AvailabilityZone: *pulumi.String(available.Names[1]),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### By Filter
//
// All Local Zones (regardless of opt-in status):
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
//				AllAvailabilityZones: pulumi.BoolRef(true),
//				Filters: []aws.GetAvailabilityZonesFilter{
//					{
//						Name: "opt-in-status",
//						Values: []string{
//							"not-opted-in",
//							"opted-in",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Only Availability Zones (no Local Zones):
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
//				Filters: []aws.GetAvailabilityZonesFilter{
//					{
//						Name: "opt-in-status",
//						Values: []string{
//							"opt-in-not-required",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAvailabilityZones(ctx *pulumi.Context, args *GetAvailabilityZonesArgs, opts ...pulumi.InvokeOption) (*GetAvailabilityZonesResult, error) {
	var rv GetAvailabilityZonesResult
	err := ctx.Invoke("aws:index/getAvailabilityZones:getAvailabilityZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesArgs struct {
	// Set to `true` to include all Availability Zones and Local Zones regardless of your opt in status.
	AllAvailabilityZones *bool `pulumi:"allAvailabilityZones"`
	// List of Availability Zone names to exclude.
	ExcludeNames []string `pulumi:"excludeNames"`
	// List of Availability Zone IDs to exclude.
	ExcludeZoneIds []string `pulumi:"excludeZoneIds"`
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetAvailabilityZonesFilter `pulumi:"filters"`
	// Allows to filter list of Availability Zones based on their
	// current state. Can be either `"available"`, `"information"`, `"impaired"` or
	// `"unavailable"`. By default the list includes a complete set of Availability Zones
	// to which the underlying AWS account has access, regardless of their state.
	State *string `pulumi:"state"`
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResult struct {
	AllAvailabilityZones *bool                        `pulumi:"allAvailabilityZones"`
	ExcludeNames         []string                     `pulumi:"excludeNames"`
	ExcludeZoneIds       []string                     `pulumi:"excludeZoneIds"`
	Filters              []GetAvailabilityZonesFilter `pulumi:"filters"`
	// A set of the Availability Zone Group names. For Availability Zones, this is the same value as the Region name. For Local Zones, the name of the associated group, for example `us-west-2-lax-1`.
	GroupNames []string `pulumi:"groupNames"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of the Availability Zone names available to the account.
	Names []string `pulumi:"names"`
	State *string  `pulumi:"state"`
	// List of the Availability Zone IDs available to the account.
	ZoneIds []string `pulumi:"zoneIds"`
}

func GetAvailabilityZonesOutput(ctx *pulumi.Context, args GetAvailabilityZonesOutputArgs, opts ...pulumi.InvokeOption) GetAvailabilityZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAvailabilityZonesResult, error) {
			args := v.(GetAvailabilityZonesArgs)
			r, err := GetAvailabilityZones(ctx, &args, opts...)
			var s GetAvailabilityZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAvailabilityZonesResultOutput)
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesOutputArgs struct {
	// Set to `true` to include all Availability Zones and Local Zones regardless of your opt in status.
	AllAvailabilityZones pulumi.BoolPtrInput `pulumi:"allAvailabilityZones"`
	// List of Availability Zone names to exclude.
	ExcludeNames pulumi.StringArrayInput `pulumi:"excludeNames"`
	// List of Availability Zone IDs to exclude.
	ExcludeZoneIds pulumi.StringArrayInput `pulumi:"excludeZoneIds"`
	// Configuration block(s) for filtering. Detailed below.
	Filters GetAvailabilityZonesFilterArrayInput `pulumi:"filters"`
	// Allows to filter list of Availability Zones based on their
	// current state. Can be either `"available"`, `"information"`, `"impaired"` or
	// `"unavailable"`. By default the list includes a complete set of Availability Zones
	// to which the underlying AWS account has access, regardless of their state.
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (GetAvailabilityZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZonesArgs)(nil)).Elem()
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResultOutput struct{ *pulumi.OutputState }

func (GetAvailabilityZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZonesResult)(nil)).Elem()
}

func (o GetAvailabilityZonesResultOutput) ToGetAvailabilityZonesResultOutput() GetAvailabilityZonesResultOutput {
	return o
}

func (o GetAvailabilityZonesResultOutput) ToGetAvailabilityZonesResultOutputWithContext(ctx context.Context) GetAvailabilityZonesResultOutput {
	return o
}

func (o GetAvailabilityZonesResultOutput) AllAvailabilityZones() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) *bool { return v.AllAvailabilityZones }).(pulumi.BoolPtrOutput)
}

func (o GetAvailabilityZonesResultOutput) ExcludeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) []string { return v.ExcludeNames }).(pulumi.StringArrayOutput)
}

func (o GetAvailabilityZonesResultOutput) ExcludeZoneIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) []string { return v.ExcludeZoneIds }).(pulumi.StringArrayOutput)
}

func (o GetAvailabilityZonesResultOutput) Filters() GetAvailabilityZonesFilterArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) []GetAvailabilityZonesFilter { return v.Filters }).(GetAvailabilityZonesFilterArrayOutput)
}

// A set of the Availability Zone Group names. For Availability Zones, this is the same value as the Region name. For Local Zones, the name of the associated group, for example `us-west-2-lax-1`.
func (o GetAvailabilityZonesResultOutput) GroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) []string { return v.GroupNames }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAvailabilityZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of the Availability Zone names available to the account.
func (o GetAvailabilityZonesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetAvailabilityZonesResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// List of the Availability Zone IDs available to the account.
func (o GetAvailabilityZonesResultOutput) ZoneIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) []string { return v.ZoneIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAvailabilityZonesResultOutput{})
}
