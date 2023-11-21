// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a specific AWS EC2 Public IPv4 Pool.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.GetPublicIpv4Pool(ctx, &ec2.GetPublicIpv4PoolArgs{
//				PoolId: "ipv4pool-ec2-000df99cff0c1ec10",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPublicIpv4Pool(ctx *pulumi.Context, args *GetPublicIpv4PoolArgs, opts ...pulumi.InvokeOption) (*GetPublicIpv4PoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPublicIpv4PoolResult
	err := ctx.Invoke("aws:ec2/getPublicIpv4Pool:getPublicIpv4Pool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPublicIpv4Pool.
type GetPublicIpv4PoolArgs struct {
	// AWS resource IDs of a public IPv4 pool (as a string) for which this data source will fetch detailed information.
	PoolId string `pulumi:"poolId"`
	// Any tags for the address pool.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getPublicIpv4Pool.
type GetPublicIpv4PoolResult struct {
	// Description of the pool, if any.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the location from which the address pool is advertised.
	// * poolAddressRanges` - List of Address Ranges in the Pool; each address range record contains:
	NetworkBorderGroup string                              `pulumi:"networkBorderGroup"`
	PoolAddressRanges  []GetPublicIpv4PoolPoolAddressRange `pulumi:"poolAddressRanges"`
	PoolId             string                              `pulumi:"poolId"`
	// Any tags for the address pool.
	Tags map[string]string `pulumi:"tags"`
	// Total number of addresses in the pool.
	TotalAddressCount int `pulumi:"totalAddressCount"`
	// Total number of available addresses in the pool.
	TotalAvailableAddressCount int `pulumi:"totalAvailableAddressCount"`
}

func GetPublicIpv4PoolOutput(ctx *pulumi.Context, args GetPublicIpv4PoolOutputArgs, opts ...pulumi.InvokeOption) GetPublicIpv4PoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPublicIpv4PoolResult, error) {
			args := v.(GetPublicIpv4PoolArgs)
			r, err := GetPublicIpv4Pool(ctx, &args, opts...)
			var s GetPublicIpv4PoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPublicIpv4PoolResultOutput)
}

// A collection of arguments for invoking getPublicIpv4Pool.
type GetPublicIpv4PoolOutputArgs struct {
	// AWS resource IDs of a public IPv4 pool (as a string) for which this data source will fetch detailed information.
	PoolId pulumi.StringInput `pulumi:"poolId"`
	// Any tags for the address pool.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetPublicIpv4PoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicIpv4PoolArgs)(nil)).Elem()
}

// A collection of values returned by getPublicIpv4Pool.
type GetPublicIpv4PoolResultOutput struct{ *pulumi.OutputState }

func (GetPublicIpv4PoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPublicIpv4PoolResult)(nil)).Elem()
}

func (o GetPublicIpv4PoolResultOutput) ToGetPublicIpv4PoolResultOutput() GetPublicIpv4PoolResultOutput {
	return o
}

func (o GetPublicIpv4PoolResultOutput) ToGetPublicIpv4PoolResultOutputWithContext(ctx context.Context) GetPublicIpv4PoolResultOutput {
	return o
}

// Description of the pool, if any.
func (o GetPublicIpv4PoolResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPublicIpv4PoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the location from which the address pool is advertised.
// * poolAddressRanges` - List of Address Ranges in the Pool; each address range record contains:
func (o GetPublicIpv4PoolResultOutput) NetworkBorderGroup() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolResult) string { return v.NetworkBorderGroup }).(pulumi.StringOutput)
}

func (o GetPublicIpv4PoolResultOutput) PoolAddressRanges() GetPublicIpv4PoolPoolAddressRangeArrayOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolResult) []GetPublicIpv4PoolPoolAddressRange { return v.PoolAddressRanges }).(GetPublicIpv4PoolPoolAddressRangeArrayOutput)
}

func (o GetPublicIpv4PoolResultOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolResult) string { return v.PoolId }).(pulumi.StringOutput)
}

// Any tags for the address pool.
func (o GetPublicIpv4PoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Total number of addresses in the pool.
func (o GetPublicIpv4PoolResultOutput) TotalAddressCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolResult) int { return v.TotalAddressCount }).(pulumi.IntOutput)
}

// Total number of available addresses in the pool.
func (o GetPublicIpv4PoolResultOutput) TotalAvailableAddressCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetPublicIpv4PoolResult) int { return v.TotalAvailableAddressCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPublicIpv4PoolResultOutput{})
}
