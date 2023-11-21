// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// `ec2.Vpc` provides details about a specific VPC.
//
// This resource can prove useful when a module accepts a vpc id as
// an input variable and needs to, for example, determine the CIDR block of that
// VPC.
func LookupVpc(ctx *pulumi.Context, args *LookupVpcArgs, opts ...pulumi.InvokeOption) (*LookupVpcResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcResult
	err := ctx.Invoke("aws:ec2/getVpc:getVpc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpc.
type LookupVpcArgs struct {
	// Cidr block of the desired VPC.
	CidrBlock *string `pulumi:"cidrBlock"`
	// Boolean constraint on whether the desired VPC is
	// the default VPC for the region.
	Default *bool `pulumi:"default"`
	// DHCP options id of the desired VPC.
	DhcpOptionsId *string `pulumi:"dhcpOptionsId"`
	// Custom filter block as described below.
	Filters []GetVpcFilter `pulumi:"filters"`
	// ID of the specific VPC to retrieve.
	Id *string `pulumi:"id"`
	// Current state of the desired VPC.
	// Can be either `"pending"` or `"available"`.
	State *string `pulumi:"state"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired VPC.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getVpc.
type LookupVpcResult struct {
	// ARN of VPC
	Arn string `pulumi:"arn"`
	// CIDR block for the association.
	CidrBlock             string                       `pulumi:"cidrBlock"`
	CidrBlockAssociations []GetVpcCidrBlockAssociation `pulumi:"cidrBlockAssociations"`
	Default               bool                         `pulumi:"default"`
	DhcpOptionsId         string                       `pulumi:"dhcpOptionsId"`
	// Whether or not the VPC has DNS hostname support
	EnableDnsHostnames bool `pulumi:"enableDnsHostnames"`
	// Whether or not the VPC has DNS support
	EnableDnsSupport bool `pulumi:"enableDnsSupport"`
	// Whether Network Address Usage metrics are enabled for your VPC
	EnableNetworkAddressUsageMetrics bool           `pulumi:"enableNetworkAddressUsageMetrics"`
	Filters                          []GetVpcFilter `pulumi:"filters"`
	Id                               string         `pulumi:"id"`
	// Allowed tenancy of instances launched into the
	// selected VPC. May be any of `"default"`, `"dedicated"`, or `"host"`.
	InstanceTenancy string `pulumi:"instanceTenancy"`
	// Association ID for the IPv6 CIDR block.
	Ipv6AssociationId string `pulumi:"ipv6AssociationId"`
	// IPv6 CIDR block.
	Ipv6CidrBlock string `pulumi:"ipv6CidrBlock"`
	// ID of the main route table associated with this VPC.
	MainRouteTableId string `pulumi:"mainRouteTableId"`
	// ID of the AWS account that owns the VPC.
	OwnerId string `pulumi:"ownerId"`
	// State of the association.
	State string            `pulumi:"state"`
	Tags  map[string]string `pulumi:"tags"`
}

func LookupVpcOutput(ctx *pulumi.Context, args LookupVpcOutputArgs, opts ...pulumi.InvokeOption) LookupVpcResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcResult, error) {
			args := v.(LookupVpcArgs)
			r, err := LookupVpc(ctx, &args, opts...)
			var s LookupVpcResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcResultOutput)
}

// A collection of arguments for invoking getVpc.
type LookupVpcOutputArgs struct {
	// Cidr block of the desired VPC.
	CidrBlock pulumi.StringPtrInput `pulumi:"cidrBlock"`
	// Boolean constraint on whether the desired VPC is
	// the default VPC for the region.
	Default pulumi.BoolPtrInput `pulumi:"default"`
	// DHCP options id of the desired VPC.
	DhcpOptionsId pulumi.StringPtrInput `pulumi:"dhcpOptionsId"`
	// Custom filter block as described below.
	Filters GetVpcFilterArrayInput `pulumi:"filters"`
	// ID of the specific VPC to retrieve.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Current state of the desired VPC.
	// Can be either `"pending"` or `"available"`.
	State pulumi.StringPtrInput `pulumi:"state"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired VPC.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupVpcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcArgs)(nil)).Elem()
}

// A collection of values returned by getVpc.
type LookupVpcResultOutput struct{ *pulumi.OutputState }

func (LookupVpcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcResult)(nil)).Elem()
}

func (o LookupVpcResultOutput) ToLookupVpcResultOutput() LookupVpcResultOutput {
	return o
}

func (o LookupVpcResultOutput) ToLookupVpcResultOutputWithContext(ctx context.Context) LookupVpcResultOutput {
	return o
}

func (o LookupVpcResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVpcResult] {
	return pulumix.Output[LookupVpcResult]{
		OutputState: o.OutputState,
	}
}

// ARN of VPC
func (o LookupVpcResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.Arn }).(pulumi.StringOutput)
}

// CIDR block for the association.
func (o LookupVpcResultOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.CidrBlock }).(pulumi.StringOutput)
}

func (o LookupVpcResultOutput) CidrBlockAssociations() GetVpcCidrBlockAssociationArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []GetVpcCidrBlockAssociation { return v.CidrBlockAssociations }).(GetVpcCidrBlockAssociationArrayOutput)
}

func (o LookupVpcResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcResult) bool { return v.Default }).(pulumi.BoolOutput)
}

func (o LookupVpcResultOutput) DhcpOptionsId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.DhcpOptionsId }).(pulumi.StringOutput)
}

// Whether or not the VPC has DNS hostname support
func (o LookupVpcResultOutput) EnableDnsHostnames() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcResult) bool { return v.EnableDnsHostnames }).(pulumi.BoolOutput)
}

// Whether or not the VPC has DNS support
func (o LookupVpcResultOutput) EnableDnsSupport() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcResult) bool { return v.EnableDnsSupport }).(pulumi.BoolOutput)
}

// Whether Network Address Usage metrics are enabled for your VPC
func (o LookupVpcResultOutput) EnableNetworkAddressUsageMetrics() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcResult) bool { return v.EnableNetworkAddressUsageMetrics }).(pulumi.BoolOutput)
}

func (o LookupVpcResultOutput) Filters() GetVpcFilterArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []GetVpcFilter { return v.Filters }).(GetVpcFilterArrayOutput)
}

func (o LookupVpcResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.Id }).(pulumi.StringOutput)
}

// Allowed tenancy of instances launched into the
// selected VPC. May be any of `"default"`, `"dedicated"`, or `"host"`.
func (o LookupVpcResultOutput) InstanceTenancy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.InstanceTenancy }).(pulumi.StringOutput)
}

// Association ID for the IPv6 CIDR block.
func (o LookupVpcResultOutput) Ipv6AssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.Ipv6AssociationId }).(pulumi.StringOutput)
}

// IPv6 CIDR block.
func (o LookupVpcResultOutput) Ipv6CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.Ipv6CidrBlock }).(pulumi.StringOutput)
}

// ID of the main route table associated with this VPC.
func (o LookupVpcResultOutput) MainRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.MainRouteTableId }).(pulumi.StringOutput)
}

// ID of the AWS account that owns the VPC.
func (o LookupVpcResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// State of the association.
func (o LookupVpcResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupVpcResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpcResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcResultOutput{})
}