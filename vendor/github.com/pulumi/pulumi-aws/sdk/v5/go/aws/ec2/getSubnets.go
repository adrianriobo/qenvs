// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can be useful for getting back a set of subnet IDs.
func GetSubnets(ctx *pulumi.Context, args *GetSubnetsArgs, opts ...pulumi.InvokeOption) (*GetSubnetsResult, error) {
	var rv GetSubnetsResult
	err := ctx.Invoke("aws:ec2/getSubnets:getSubnets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnets.
type GetSubnetsArgs struct {
	// Custom filter block as described below.
	Filters []GetSubnetsFilter `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired subnets.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getSubnets.
type GetSubnetsResult struct {
	Filters []GetSubnetsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of all the subnet ids found.
	Ids  []string          `pulumi:"ids"`
	Tags map[string]string `pulumi:"tags"`
}

func GetSubnetsOutput(ctx *pulumi.Context, args GetSubnetsOutputArgs, opts ...pulumi.InvokeOption) GetSubnetsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSubnetsResult, error) {
			args := v.(GetSubnetsArgs)
			r, err := GetSubnets(ctx, &args, opts...)
			var s GetSubnetsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSubnetsResultOutput)
}

// A collection of arguments for invoking getSubnets.
type GetSubnetsOutputArgs struct {
	// Custom filter block as described below.
	Filters GetSubnetsFilterArrayInput `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired subnets.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetSubnetsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetsArgs)(nil)).Elem()
}

// A collection of values returned by getSubnets.
type GetSubnetsResultOutput struct{ *pulumi.OutputState }

func (GetSubnetsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetsResult)(nil)).Elem()
}

func (o GetSubnetsResultOutput) ToGetSubnetsResultOutput() GetSubnetsResultOutput {
	return o
}

func (o GetSubnetsResultOutput) ToGetSubnetsResultOutputWithContext(ctx context.Context) GetSubnetsResultOutput {
	return o
}

func (o GetSubnetsResultOutput) Filters() GetSubnetsFilterArrayOutput {
	return o.ApplyT(func(v GetSubnetsResult) []GetSubnetsFilter { return v.Filters }).(GetSubnetsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSubnetsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of all the subnet ids found.
func (o GetSubnetsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSubnetsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetSubnetsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSubnetsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSubnetsResultOutput{})
}
