// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IP address pool resource for IPAM.
//
// ## Example Usage
//
// Basic usage:
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
//			current, err := aws.GetRegion(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			exampleVpcIpam, err := ec2.NewVpcIpam(ctx, "exampleVpcIpam", &ec2.VpcIpamArgs{
//				OperatingRegions: ec2.VpcIpamOperatingRegionArray{
//					&ec2.VpcIpamOperatingRegionArgs{
//						RegionName: pulumi.String(current.Name),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcIpamPool(ctx, "exampleVpcIpamPool", &ec2.VpcIpamPoolArgs{
//				AddressFamily: pulumi.String("ipv4"),
//				IpamScopeId:   exampleVpcIpam.PrivateDefaultScopeId,
//				Locale:        pulumi.String(current.Name),
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
// Nested Pools:
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
//			current, err := aws.GetRegion(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			example, err := ec2.NewVpcIpam(ctx, "example", &ec2.VpcIpamArgs{
//				OperatingRegions: ec2.VpcIpamOperatingRegionArray{
//					&ec2.VpcIpamOperatingRegionArgs{
//						RegionName: pulumi.String(current.Name),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			parent, err := ec2.NewVpcIpamPool(ctx, "parent", &ec2.VpcIpamPoolArgs{
//				AddressFamily: pulumi.String("ipv4"),
//				IpamScopeId:   example.PrivateDefaultScopeId,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcIpamPoolCidr(ctx, "parentTest", &ec2.VpcIpamPoolCidrArgs{
//				IpamPoolId: parent.ID(),
//				Cidr:       pulumi.String("172.2.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			child, err := ec2.NewVpcIpamPool(ctx, "child", &ec2.VpcIpamPoolArgs{
//				AddressFamily:    pulumi.String("ipv4"),
//				IpamScopeId:      example.PrivateDefaultScopeId,
//				Locale:           pulumi.String(current.Name),
//				SourceIpamPoolId: parent.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcIpamPoolCidr(ctx, "childTest", &ec2.VpcIpamPoolCidrArgs{
//				IpamPoolId: child.ID(),
//				Cidr:       pulumi.String("172.2.0.0/24"),
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
// IPAMs can be imported using the `ipam pool id`, e.g.
//
// ```sh
//
//	$ pulumi import aws:ec2/vpcIpamPool:VpcIpamPool example ipam-pool-0958f95207d978e1e
//
// ```
type VpcIpamPool struct {
	pulumi.CustomResourceState

	// The IP protocol assigned to this pool. You must choose either IPv4 or IPv6 protocol for a pool.
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16 (unless you provide a different netmask value when you create the new allocation).
	AllocationDefaultNetmaskLength pulumi.IntPtrOutput `pulumi:"allocationDefaultNetmaskLength"`
	// The maximum netmask length that will be required for CIDR allocations in this pool.
	AllocationMaxNetmaskLength pulumi.IntPtrOutput `pulumi:"allocationMaxNetmaskLength"`
	// The minimum netmask length that will be required for CIDR allocations in this pool.
	AllocationMinNetmaskLength pulumi.IntPtrOutput `pulumi:"allocationMinNetmaskLength"`
	// Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	AllocationResourceTags pulumi.StringMapOutput `pulumi:"allocationResourceTags"`
	// Amazon Resource Name (ARN) of IPAM
	Arn pulumi.StringOutput `pulumi:"arn"`
	// If you include this argument, IPAM automatically imports any VPCs you have in your scope that fall
	// within the CIDR range in the pool.
	AutoImport pulumi.BoolPtrOutput `pulumi:"autoImport"`
	// Limits which AWS service the pool can be used in. Only useable on public scopes. Valid Values: `ec2`.
	AwsService pulumi.StringPtrOutput `pulumi:"awsService"`
	// A description for the IPAM pool.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the scope in which you would like to create the IPAM pool.
	IpamScopeId   pulumi.StringOutput `pulumi:"ipamScopeId"`
	IpamScopeType pulumi.StringOutput `pulumi:"ipamScopeType"`
	// The locale in which you would like to create the IPAM pool. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. Possible values: Any AWS region, such as `us-east-1`.
	Locale    pulumi.StringPtrOutput `pulumi:"locale"`
	PoolDepth pulumi.IntOutput       `pulumi:"poolDepth"`
	// Defines whether or not IPv6 pool space is publicly advertisable over the internet. This option is not available for IPv4 pool space.
	PubliclyAdvertisable pulumi.BoolPtrOutput `pulumi:"publiclyAdvertisable"`
	// The ID of the source IPAM pool. Use this argument to create a child pool within an existing pool.
	SourceIpamPoolId pulumi.StringPtrOutput `pulumi:"sourceIpamPoolId"`
	// The ID of the IPAM
	State pulumi.StringOutput `pulumi:"state"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVpcIpamPool registers a new resource with the given unique name, arguments, and options.
func NewVpcIpamPool(ctx *pulumi.Context,
	name string, args *VpcIpamPoolArgs, opts ...pulumi.ResourceOption) (*VpcIpamPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'AddressFamily'")
	}
	if args.IpamScopeId == nil {
		return nil, errors.New("invalid value for required argument 'IpamScopeId'")
	}
	var resource VpcIpamPool
	err := ctx.RegisterResource("aws:ec2/vpcIpamPool:VpcIpamPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIpamPool gets an existing VpcIpamPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIpamPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIpamPoolState, opts ...pulumi.ResourceOption) (*VpcIpamPool, error) {
	var resource VpcIpamPool
	err := ctx.ReadResource("aws:ec2/vpcIpamPool:VpcIpamPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIpamPool resources.
type vpcIpamPoolState struct {
	// The IP protocol assigned to this pool. You must choose either IPv4 or IPv6 protocol for a pool.
	AddressFamily *string `pulumi:"addressFamily"`
	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16 (unless you provide a different netmask value when you create the new allocation).
	AllocationDefaultNetmaskLength *int `pulumi:"allocationDefaultNetmaskLength"`
	// The maximum netmask length that will be required for CIDR allocations in this pool.
	AllocationMaxNetmaskLength *int `pulumi:"allocationMaxNetmaskLength"`
	// The minimum netmask length that will be required for CIDR allocations in this pool.
	AllocationMinNetmaskLength *int `pulumi:"allocationMinNetmaskLength"`
	// Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	AllocationResourceTags map[string]string `pulumi:"allocationResourceTags"`
	// Amazon Resource Name (ARN) of IPAM
	Arn *string `pulumi:"arn"`
	// If you include this argument, IPAM automatically imports any VPCs you have in your scope that fall
	// within the CIDR range in the pool.
	AutoImport *bool `pulumi:"autoImport"`
	// Limits which AWS service the pool can be used in. Only useable on public scopes. Valid Values: `ec2`.
	AwsService *string `pulumi:"awsService"`
	// A description for the IPAM pool.
	Description *string `pulumi:"description"`
	// The ID of the scope in which you would like to create the IPAM pool.
	IpamScopeId   *string `pulumi:"ipamScopeId"`
	IpamScopeType *string `pulumi:"ipamScopeType"`
	// The locale in which you would like to create the IPAM pool. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. Possible values: Any AWS region, such as `us-east-1`.
	Locale    *string `pulumi:"locale"`
	PoolDepth *int    `pulumi:"poolDepth"`
	// Defines whether or not IPv6 pool space is publicly advertisable over the internet. This option is not available for IPv4 pool space.
	PubliclyAdvertisable *bool `pulumi:"publiclyAdvertisable"`
	// The ID of the source IPAM pool. Use this argument to create a child pool within an existing pool.
	SourceIpamPoolId *string `pulumi:"sourceIpamPoolId"`
	// The ID of the IPAM
	State *string `pulumi:"state"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VpcIpamPoolState struct {
	// The IP protocol assigned to this pool. You must choose either IPv4 or IPv6 protocol for a pool.
	AddressFamily pulumi.StringPtrInput
	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16 (unless you provide a different netmask value when you create the new allocation).
	AllocationDefaultNetmaskLength pulumi.IntPtrInput
	// The maximum netmask length that will be required for CIDR allocations in this pool.
	AllocationMaxNetmaskLength pulumi.IntPtrInput
	// The minimum netmask length that will be required for CIDR allocations in this pool.
	AllocationMinNetmaskLength pulumi.IntPtrInput
	// Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	AllocationResourceTags pulumi.StringMapInput
	// Amazon Resource Name (ARN) of IPAM
	Arn pulumi.StringPtrInput
	// If you include this argument, IPAM automatically imports any VPCs you have in your scope that fall
	// within the CIDR range in the pool.
	AutoImport pulumi.BoolPtrInput
	// Limits which AWS service the pool can be used in. Only useable on public scopes. Valid Values: `ec2`.
	AwsService pulumi.StringPtrInput
	// A description for the IPAM pool.
	Description pulumi.StringPtrInput
	// The ID of the scope in which you would like to create the IPAM pool.
	IpamScopeId   pulumi.StringPtrInput
	IpamScopeType pulumi.StringPtrInput
	// The locale in which you would like to create the IPAM pool. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. Possible values: Any AWS region, such as `us-east-1`.
	Locale    pulumi.StringPtrInput
	PoolDepth pulumi.IntPtrInput
	// Defines whether or not IPv6 pool space is publicly advertisable over the internet. This option is not available for IPv4 pool space.
	PubliclyAdvertisable pulumi.BoolPtrInput
	// The ID of the source IPAM pool. Use this argument to create a child pool within an existing pool.
	SourceIpamPoolId pulumi.StringPtrInput
	// The ID of the IPAM
	State pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (VpcIpamPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamPoolState)(nil)).Elem()
}

type vpcIpamPoolArgs struct {
	// The IP protocol assigned to this pool. You must choose either IPv4 or IPv6 protocol for a pool.
	AddressFamily string `pulumi:"addressFamily"`
	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16 (unless you provide a different netmask value when you create the new allocation).
	AllocationDefaultNetmaskLength *int `pulumi:"allocationDefaultNetmaskLength"`
	// The maximum netmask length that will be required for CIDR allocations in this pool.
	AllocationMaxNetmaskLength *int `pulumi:"allocationMaxNetmaskLength"`
	// The minimum netmask length that will be required for CIDR allocations in this pool.
	AllocationMinNetmaskLength *int `pulumi:"allocationMinNetmaskLength"`
	// Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	AllocationResourceTags map[string]string `pulumi:"allocationResourceTags"`
	// If you include this argument, IPAM automatically imports any VPCs you have in your scope that fall
	// within the CIDR range in the pool.
	AutoImport *bool `pulumi:"autoImport"`
	// Limits which AWS service the pool can be used in. Only useable on public scopes. Valid Values: `ec2`.
	AwsService *string `pulumi:"awsService"`
	// A description for the IPAM pool.
	Description *string `pulumi:"description"`
	// The ID of the scope in which you would like to create the IPAM pool.
	IpamScopeId string `pulumi:"ipamScopeId"`
	// The locale in which you would like to create the IPAM pool. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. Possible values: Any AWS region, such as `us-east-1`.
	Locale *string `pulumi:"locale"`
	// Defines whether or not IPv6 pool space is publicly advertisable over the internet. This option is not available for IPv4 pool space.
	PubliclyAdvertisable *bool `pulumi:"publiclyAdvertisable"`
	// The ID of the source IPAM pool. Use this argument to create a child pool within an existing pool.
	SourceIpamPoolId *string `pulumi:"sourceIpamPoolId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VpcIpamPool resource.
type VpcIpamPoolArgs struct {
	// The IP protocol assigned to this pool. You must choose either IPv4 or IPv6 protocol for a pool.
	AddressFamily pulumi.StringInput
	// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16 (unless you provide a different netmask value when you create the new allocation).
	AllocationDefaultNetmaskLength pulumi.IntPtrInput
	// The maximum netmask length that will be required for CIDR allocations in this pool.
	AllocationMaxNetmaskLength pulumi.IntPtrInput
	// The minimum netmask length that will be required for CIDR allocations in this pool.
	AllocationMinNetmaskLength pulumi.IntPtrInput
	// Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
	AllocationResourceTags pulumi.StringMapInput
	// If you include this argument, IPAM automatically imports any VPCs you have in your scope that fall
	// within the CIDR range in the pool.
	AutoImport pulumi.BoolPtrInput
	// Limits which AWS service the pool can be used in. Only useable on public scopes. Valid Values: `ec2`.
	AwsService pulumi.StringPtrInput
	// A description for the IPAM pool.
	Description pulumi.StringPtrInput
	// The ID of the scope in which you would like to create the IPAM pool.
	IpamScopeId pulumi.StringInput
	// The locale in which you would like to create the IPAM pool. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. Possible values: Any AWS region, such as `us-east-1`.
	Locale pulumi.StringPtrInput
	// Defines whether or not IPv6 pool space is publicly advertisable over the internet. This option is not available for IPv4 pool space.
	PubliclyAdvertisable pulumi.BoolPtrInput
	// The ID of the source IPAM pool. Use this argument to create a child pool within an existing pool.
	SourceIpamPoolId pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (VpcIpamPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamPoolArgs)(nil)).Elem()
}

type VpcIpamPoolInput interface {
	pulumi.Input

	ToVpcIpamPoolOutput() VpcIpamPoolOutput
	ToVpcIpamPoolOutputWithContext(ctx context.Context) VpcIpamPoolOutput
}

func (*VpcIpamPool) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamPool)(nil)).Elem()
}

func (i *VpcIpamPool) ToVpcIpamPoolOutput() VpcIpamPoolOutput {
	return i.ToVpcIpamPoolOutputWithContext(context.Background())
}

func (i *VpcIpamPool) ToVpcIpamPoolOutputWithContext(ctx context.Context) VpcIpamPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamPoolOutput)
}

// VpcIpamPoolArrayInput is an input type that accepts VpcIpamPoolArray and VpcIpamPoolArrayOutput values.
// You can construct a concrete instance of `VpcIpamPoolArrayInput` via:
//
//	VpcIpamPoolArray{ VpcIpamPoolArgs{...} }
type VpcIpamPoolArrayInput interface {
	pulumi.Input

	ToVpcIpamPoolArrayOutput() VpcIpamPoolArrayOutput
	ToVpcIpamPoolArrayOutputWithContext(context.Context) VpcIpamPoolArrayOutput
}

type VpcIpamPoolArray []VpcIpamPoolInput

func (VpcIpamPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamPool)(nil)).Elem()
}

func (i VpcIpamPoolArray) ToVpcIpamPoolArrayOutput() VpcIpamPoolArrayOutput {
	return i.ToVpcIpamPoolArrayOutputWithContext(context.Background())
}

func (i VpcIpamPoolArray) ToVpcIpamPoolArrayOutputWithContext(ctx context.Context) VpcIpamPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamPoolArrayOutput)
}

// VpcIpamPoolMapInput is an input type that accepts VpcIpamPoolMap and VpcIpamPoolMapOutput values.
// You can construct a concrete instance of `VpcIpamPoolMapInput` via:
//
//	VpcIpamPoolMap{ "key": VpcIpamPoolArgs{...} }
type VpcIpamPoolMapInput interface {
	pulumi.Input

	ToVpcIpamPoolMapOutput() VpcIpamPoolMapOutput
	ToVpcIpamPoolMapOutputWithContext(context.Context) VpcIpamPoolMapOutput
}

type VpcIpamPoolMap map[string]VpcIpamPoolInput

func (VpcIpamPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamPool)(nil)).Elem()
}

func (i VpcIpamPoolMap) ToVpcIpamPoolMapOutput() VpcIpamPoolMapOutput {
	return i.ToVpcIpamPoolMapOutputWithContext(context.Background())
}

func (i VpcIpamPoolMap) ToVpcIpamPoolMapOutputWithContext(ctx context.Context) VpcIpamPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamPoolMapOutput)
}

type VpcIpamPoolOutput struct{ *pulumi.OutputState }

func (VpcIpamPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamPool)(nil)).Elem()
}

func (o VpcIpamPoolOutput) ToVpcIpamPoolOutput() VpcIpamPoolOutput {
	return o
}

func (o VpcIpamPoolOutput) ToVpcIpamPoolOutputWithContext(ctx context.Context) VpcIpamPoolOutput {
	return o
}

// The IP protocol assigned to this pool. You must choose either IPv4 or IPv6 protocol for a pool.
func (o VpcIpamPoolOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringOutput { return v.AddressFamily }).(pulumi.StringOutput)
}

// A default netmask length for allocations added to this pool. If, for example, the CIDR assigned to this pool is 10.0.0.0/8 and you enter 16 here, new allocations will default to 10.0.0.0/16 (unless you provide a different netmask value when you create the new allocation).
func (o VpcIpamPoolOutput) AllocationDefaultNetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.IntPtrOutput { return v.AllocationDefaultNetmaskLength }).(pulumi.IntPtrOutput)
}

// The maximum netmask length that will be required for CIDR allocations in this pool.
func (o VpcIpamPoolOutput) AllocationMaxNetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.IntPtrOutput { return v.AllocationMaxNetmaskLength }).(pulumi.IntPtrOutput)
}

// The minimum netmask length that will be required for CIDR allocations in this pool.
func (o VpcIpamPoolOutput) AllocationMinNetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.IntPtrOutput { return v.AllocationMinNetmaskLength }).(pulumi.IntPtrOutput)
}

// Tags that are required for resources that use CIDRs from this IPAM pool. Resources that do not have these tags will not be allowed to allocate space from the pool. If the resources have their tags changed after they have allocated space or if the allocation tagging requirements are changed on the pool, the resource may be marked as noncompliant.
func (o VpcIpamPoolOutput) AllocationResourceTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringMapOutput { return v.AllocationResourceTags }).(pulumi.StringMapOutput)
}

// Amazon Resource Name (ARN) of IPAM
func (o VpcIpamPoolOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// If you include this argument, IPAM automatically imports any VPCs you have in your scope that fall
// within the CIDR range in the pool.
func (o VpcIpamPoolOutput) AutoImport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.BoolPtrOutput { return v.AutoImport }).(pulumi.BoolPtrOutput)
}

// Limits which AWS service the pool can be used in. Only useable on public scopes. Valid Values: `ec2`.
func (o VpcIpamPoolOutput) AwsService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringPtrOutput { return v.AwsService }).(pulumi.StringPtrOutput)
}

// A description for the IPAM pool.
func (o VpcIpamPoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the scope in which you would like to create the IPAM pool.
func (o VpcIpamPoolOutput) IpamScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringOutput { return v.IpamScopeId }).(pulumi.StringOutput)
}

func (o VpcIpamPoolOutput) IpamScopeType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringOutput { return v.IpamScopeType }).(pulumi.StringOutput)
}

// The locale in which you would like to create the IPAM pool. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. Possible values: Any AWS region, such as `us-east-1`.
func (o VpcIpamPoolOutput) Locale() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringPtrOutput { return v.Locale }).(pulumi.StringPtrOutput)
}

func (o VpcIpamPoolOutput) PoolDepth() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.IntOutput { return v.PoolDepth }).(pulumi.IntOutput)
}

// Defines whether or not IPv6 pool space is publicly advertisable over the internet. This option is not available for IPv4 pool space.
func (o VpcIpamPoolOutput) PubliclyAdvertisable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.BoolPtrOutput { return v.PubliclyAdvertisable }).(pulumi.BoolPtrOutput)
}

// The ID of the source IPAM pool. Use this argument to create a child pool within an existing pool.
func (o VpcIpamPoolOutput) SourceIpamPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringPtrOutput { return v.SourceIpamPoolId }).(pulumi.StringPtrOutput)
}

// The ID of the IPAM
func (o VpcIpamPoolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VpcIpamPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o VpcIpamPoolOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpamPool) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type VpcIpamPoolArrayOutput struct{ *pulumi.OutputState }

func (VpcIpamPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamPool)(nil)).Elem()
}

func (o VpcIpamPoolArrayOutput) ToVpcIpamPoolArrayOutput() VpcIpamPoolArrayOutput {
	return o
}

func (o VpcIpamPoolArrayOutput) ToVpcIpamPoolArrayOutputWithContext(ctx context.Context) VpcIpamPoolArrayOutput {
	return o
}

func (o VpcIpamPoolArrayOutput) Index(i pulumi.IntInput) VpcIpamPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcIpamPool {
		return vs[0].([]*VpcIpamPool)[vs[1].(int)]
	}).(VpcIpamPoolOutput)
}

type VpcIpamPoolMapOutput struct{ *pulumi.OutputState }

func (VpcIpamPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamPool)(nil)).Elem()
}

func (o VpcIpamPoolMapOutput) ToVpcIpamPoolMapOutput() VpcIpamPoolMapOutput {
	return o
}

func (o VpcIpamPoolMapOutput) ToVpcIpamPoolMapOutputWithContext(ctx context.Context) VpcIpamPoolMapOutput {
	return o
}

func (o VpcIpamPoolMapOutput) MapIndex(k pulumi.StringInput) VpcIpamPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcIpamPool {
		return vs[0].(map[string]*VpcIpamPool)[vs[1].(string)]
	}).(VpcIpamPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamPoolInput)(nil)).Elem(), &VpcIpamPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamPoolArrayInput)(nil)).Elem(), VpcIpamPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamPoolMapInput)(nil)).Elem(), VpcIpamPoolMap{})
	pulumi.RegisterOutputType(VpcIpamPoolOutput{})
	pulumi.RegisterOutputType(VpcIpamPoolArrayOutput{})
	pulumi.RegisterOutputType(VpcIpamPoolMapOutput{})
}
