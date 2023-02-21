// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC resource.
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
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
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
// Basic usage with tags:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
//				CidrBlock:       pulumi.String("10.0.0.0/16"),
//				InstanceTenancy: pulumi.String("default"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("main"),
//				},
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
// VPC with CIDR from AWS IPAM:
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
//			testVpcIpam, err := ec2.NewVpcIpam(ctx, "testVpcIpam", &ec2.VpcIpamArgs{
//				OperatingRegions: ec2.VpcIpamOperatingRegionArray{
//					&ec2.VpcIpamOperatingRegionArgs{
//						RegionName: *pulumi.String(current.Name),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			testVpcIpamPool, err := ec2.NewVpcIpamPool(ctx, "testVpcIpamPool", &ec2.VpcIpamPoolArgs{
//				AddressFamily: pulumi.String("ipv4"),
//				IpamScopeId:   testVpcIpam.PrivateDefaultScopeId,
//				Locale:        *pulumi.String(current.Name),
//			})
//			if err != nil {
//				return err
//			}
//			testVpcIpamPoolCidr, err := ec2.NewVpcIpamPoolCidr(ctx, "testVpcIpamPoolCidr", &ec2.VpcIpamPoolCidrArgs{
//				IpamPoolId: testVpcIpamPool.ID(),
//				Cidr:       pulumi.String("172.2.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpc(ctx, "testVpc", &ec2.VpcArgs{
//				Ipv4IpamPoolId:    testVpcIpamPool.ID(),
//				Ipv4NetmaskLength: pulumi.Int(28),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				testVpcIpamPoolCidr,
//			}))
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
// VPCs can be imported using the `vpc id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/vpc:Vpc test_vpc vpc-a01106c2
//
// ```
type Vpc struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of VPC
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6IpamPoolId`
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrOutput `pulumi:"assignGeneratedIpv6CidrBlock"`
	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4NetmaskLength`.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId pulumi.StringOutput `pulumi:"defaultNetworkAclId"`
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId pulumi.StringOutput `pulumi:"defaultRouteTableId"`
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId pulumi.StringOutput `pulumi:"defaultSecurityGroupId"`
	DhcpOptionsId          pulumi.StringOutput `pulumi:"dhcpOptionsId"`
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	//
	// Deprecated: With the retirement of EC2-Classic the enable_classiclink attribute has been deprecated and will be removed in a future version.
	EnableClassiclink pulumi.BoolOutput `pulumi:"enableClassiclink"`
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	//
	// Deprecated: With the retirement of EC2-Classic the enable_classiclink_dns_support attribute has been deprecated and will be removed in a future version.
	EnableClassiclinkDnsSupport pulumi.BoolOutput `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolOutput `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults to true.
	EnableDnsSupport pulumi.BoolPtrOutput `pulumi:"enableDnsSupport"`
	// Indicates whether Network Address Usage metrics are enabled for your VPC. Defaults to false.
	EnableNetworkAddressUsageMetrics pulumi.BoolOutput `pulumi:"enableNetworkAddressUsageMetrics"`
	// A tenancy option for instances launched into the VPC. Default is `default`, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is `dedicated`, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
	InstanceTenancy pulumi.StringPtrOutput `pulumi:"instanceTenancy"`
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	Ipv4IpamPoolId pulumi.StringPtrOutput `pulumi:"ipv4IpamPoolId"`
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4IpamPoolId`.
	Ipv4NetmaskLength pulumi.IntPtrOutput `pulumi:"ipv4NetmaskLength"`
	// The association ID for the IPv6 CIDR block.
	Ipv6AssociationId pulumi.StringOutput `pulumi:"ipv6AssociationId"`
	// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6NetmaskLength`.
	Ipv6CidrBlock pulumi.StringOutput `pulumi:"ipv6CidrBlock"`
	// By default when an IPv6 CIDR is assigned to a VPC a default ipv6CidrBlockNetworkBorderGroup will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
	Ipv6CidrBlockNetworkBorderGroup pulumi.StringOutput `pulumi:"ipv6CidrBlockNetworkBorderGroup"`
	// IPAM Pool ID for a IPv6 pool. Conflicts with `assignGeneratedIpv6CidrBlock`.
	Ipv6IpamPoolId pulumi.StringPtrOutput `pulumi:"ipv6IpamPoolId"`
	// Netmask length to request from IPAM Pool. Conflicts with `ipv6CidrBlock`. This can be omitted if IPAM pool as a `allocationDefaultNetmaskLength` set. Valid values: `56`.
	Ipv6NetmaskLength pulumi.IntPtrOutput `pulumi:"ipv6NetmaskLength"`
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// `ec2.MainRouteTableAssociation`.
	MainRouteTableId pulumi.StringOutput `pulumi:"mainRouteTableId"`
	// The ID of the AWS account that owns the VPC.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVpc registers a new resource with the given unique name, arguments, and options.
func NewVpc(ctx *pulumi.Context,
	name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil {
		args = &VpcArgs{}
	}

	var resource Vpc
	err := ctx.RegisterResource("aws:ec2/vpc:Vpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpc gets an existing Vpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcState, opts ...pulumi.ResourceOption) (*Vpc, error) {
	var resource Vpc
	err := ctx.ReadResource("aws:ec2/vpc:Vpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vpc resources.
type vpcState struct {
	// Amazon Resource Name (ARN) of VPC
	Arn *string `pulumi:"arn"`
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6IpamPoolId`
	AssignGeneratedIpv6CidrBlock *bool `pulumi:"assignGeneratedIpv6CidrBlock"`
	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4NetmaskLength`.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId *string `pulumi:"defaultNetworkAclId"`
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId *string `pulumi:"defaultRouteTableId"`
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId *string `pulumi:"defaultSecurityGroupId"`
	DhcpOptionsId          *string `pulumi:"dhcpOptionsId"`
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	//
	// Deprecated: With the retirement of EC2-Classic the enable_classiclink attribute has been deprecated and will be removed in a future version.
	EnableClassiclink *bool `pulumi:"enableClassiclink"`
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	//
	// Deprecated: With the retirement of EC2-Classic the enable_classiclink_dns_support attribute has been deprecated and will be removed in a future version.
	EnableClassiclinkDnsSupport *bool `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames *bool `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults to true.
	EnableDnsSupport *bool `pulumi:"enableDnsSupport"`
	// Indicates whether Network Address Usage metrics are enabled for your VPC. Defaults to false.
	EnableNetworkAddressUsageMetrics *bool `pulumi:"enableNetworkAddressUsageMetrics"`
	// A tenancy option for instances launched into the VPC. Default is `default`, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is `dedicated`, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
	InstanceTenancy *string `pulumi:"instanceTenancy"`
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	Ipv4IpamPoolId *string `pulumi:"ipv4IpamPoolId"`
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4IpamPoolId`.
	Ipv4NetmaskLength *int `pulumi:"ipv4NetmaskLength"`
	// The association ID for the IPv6 CIDR block.
	Ipv6AssociationId *string `pulumi:"ipv6AssociationId"`
	// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6NetmaskLength`.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// By default when an IPv6 CIDR is assigned to a VPC a default ipv6CidrBlockNetworkBorderGroup will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
	Ipv6CidrBlockNetworkBorderGroup *string `pulumi:"ipv6CidrBlockNetworkBorderGroup"`
	// IPAM Pool ID for a IPv6 pool. Conflicts with `assignGeneratedIpv6CidrBlock`.
	Ipv6IpamPoolId *string `pulumi:"ipv6IpamPoolId"`
	// Netmask length to request from IPAM Pool. Conflicts with `ipv6CidrBlock`. This can be omitted if IPAM pool as a `allocationDefaultNetmaskLength` set. Valid values: `56`.
	Ipv6NetmaskLength *int `pulumi:"ipv6NetmaskLength"`
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// `ec2.MainRouteTableAssociation`.
	MainRouteTableId *string `pulumi:"mainRouteTableId"`
	// The ID of the AWS account that owns the VPC.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VpcState struct {
	// Amazon Resource Name (ARN) of VPC
	Arn pulumi.StringPtrInput
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6IpamPoolId`
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrInput
	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4NetmaskLength`.
	CidrBlock pulumi.StringPtrInput
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId pulumi.StringPtrInput
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId pulumi.StringPtrInput
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId pulumi.StringPtrInput
	DhcpOptionsId          pulumi.StringPtrInput
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	//
	// Deprecated: With the retirement of EC2-Classic the enable_classiclink attribute has been deprecated and will be removed in a future version.
	EnableClassiclink pulumi.BoolPtrInput
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	//
	// Deprecated: With the retirement of EC2-Classic the enable_classiclink_dns_support attribute has been deprecated and will be removed in a future version.
	EnableClassiclinkDnsSupport pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS support in the VPC. Defaults to true.
	EnableDnsSupport pulumi.BoolPtrInput
	// Indicates whether Network Address Usage metrics are enabled for your VPC. Defaults to false.
	EnableNetworkAddressUsageMetrics pulumi.BoolPtrInput
	// A tenancy option for instances launched into the VPC. Default is `default`, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is `dedicated`, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
	InstanceTenancy pulumi.StringPtrInput
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	Ipv4IpamPoolId pulumi.StringPtrInput
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4IpamPoolId`.
	Ipv4NetmaskLength pulumi.IntPtrInput
	// The association ID for the IPv6 CIDR block.
	Ipv6AssociationId pulumi.StringPtrInput
	// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6NetmaskLength`.
	Ipv6CidrBlock pulumi.StringPtrInput
	// By default when an IPv6 CIDR is assigned to a VPC a default ipv6CidrBlockNetworkBorderGroup will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
	Ipv6CidrBlockNetworkBorderGroup pulumi.StringPtrInput
	// IPAM Pool ID for a IPv6 pool. Conflicts with `assignGeneratedIpv6CidrBlock`.
	Ipv6IpamPoolId pulumi.StringPtrInput
	// Netmask length to request from IPAM Pool. Conflicts with `ipv6CidrBlock`. This can be omitted if IPAM pool as a `allocationDefaultNetmaskLength` set. Valid values: `56`.
	Ipv6NetmaskLength pulumi.IntPtrInput
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// `ec2.MainRouteTableAssociation`.
	MainRouteTableId pulumi.StringPtrInput
	// The ID of the AWS account that owns the VPC.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (VpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcState)(nil)).Elem()
}

type vpcArgs struct {
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6IpamPoolId`
	AssignGeneratedIpv6CidrBlock *bool `pulumi:"assignGeneratedIpv6CidrBlock"`
	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4NetmaskLength`.
	CidrBlock *string `pulumi:"cidrBlock"`
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	//
	// Deprecated: With the retirement of EC2-Classic the enable_classiclink attribute has been deprecated and will be removed in a future version.
	EnableClassiclink *bool `pulumi:"enableClassiclink"`
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	//
	// Deprecated: With the retirement of EC2-Classic the enable_classiclink_dns_support attribute has been deprecated and will be removed in a future version.
	EnableClassiclinkDnsSupport *bool `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames *bool `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults to true.
	EnableDnsSupport *bool `pulumi:"enableDnsSupport"`
	// Indicates whether Network Address Usage metrics are enabled for your VPC. Defaults to false.
	EnableNetworkAddressUsageMetrics *bool `pulumi:"enableNetworkAddressUsageMetrics"`
	// A tenancy option for instances launched into the VPC. Default is `default`, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is `dedicated`, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
	InstanceTenancy *string `pulumi:"instanceTenancy"`
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	Ipv4IpamPoolId *string `pulumi:"ipv4IpamPoolId"`
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4IpamPoolId`.
	Ipv4NetmaskLength *int `pulumi:"ipv4NetmaskLength"`
	// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6NetmaskLength`.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// By default when an IPv6 CIDR is assigned to a VPC a default ipv6CidrBlockNetworkBorderGroup will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
	Ipv6CidrBlockNetworkBorderGroup *string `pulumi:"ipv6CidrBlockNetworkBorderGroup"`
	// IPAM Pool ID for a IPv6 pool. Conflicts with `assignGeneratedIpv6CidrBlock`.
	Ipv6IpamPoolId *string `pulumi:"ipv6IpamPoolId"`
	// Netmask length to request from IPAM Pool. Conflicts with `ipv6CidrBlock`. This can be omitted if IPAM pool as a `allocationDefaultNetmaskLength` set. Valid values: `56`.
	Ipv6NetmaskLength *int `pulumi:"ipv6NetmaskLength"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Vpc resource.
type VpcArgs struct {
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6IpamPoolId`
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrInput
	// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4NetmaskLength`.
	CidrBlock pulumi.StringPtrInput
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	//
	// Deprecated: With the retirement of EC2-Classic the enable_classiclink attribute has been deprecated and will be removed in a future version.
	EnableClassiclink pulumi.BoolPtrInput
	// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
	// Only valid in regions and accounts that support EC2 Classic.
	//
	// Deprecated: With the retirement of EC2-Classic the enable_classiclink_dns_support attribute has been deprecated and will be removed in a future version.
	EnableClassiclinkDnsSupport pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS support in the VPC. Defaults to true.
	EnableDnsSupport pulumi.BoolPtrInput
	// Indicates whether Network Address Usage metrics are enabled for your VPC. Defaults to false.
	EnableNetworkAddressUsageMetrics pulumi.BoolPtrInput
	// A tenancy option for instances launched into the VPC. Default is `default`, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is `dedicated`, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
	InstanceTenancy pulumi.StringPtrInput
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
	Ipv4IpamPoolId pulumi.StringPtrInput
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4IpamPoolId`.
	Ipv4NetmaskLength pulumi.IntPtrInput
	// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6NetmaskLength`.
	Ipv6CidrBlock pulumi.StringPtrInput
	// By default when an IPv6 CIDR is assigned to a VPC a default ipv6CidrBlockNetworkBorderGroup will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
	Ipv6CidrBlockNetworkBorderGroup pulumi.StringPtrInput
	// IPAM Pool ID for a IPv6 pool. Conflicts with `assignGeneratedIpv6CidrBlock`.
	Ipv6IpamPoolId pulumi.StringPtrInput
	// Netmask length to request from IPAM Pool. Conflicts with `ipv6CidrBlock`. This can be omitted if IPAM pool as a `allocationDefaultNetmaskLength` set. Valid values: `56`.
	Ipv6NetmaskLength pulumi.IntPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (VpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcArgs)(nil)).Elem()
}

type VpcInput interface {
	pulumi.Input

	ToVpcOutput() VpcOutput
	ToVpcOutputWithContext(ctx context.Context) VpcOutput
}

func (*Vpc) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil)).Elem()
}

func (i *Vpc) ToVpcOutput() VpcOutput {
	return i.ToVpcOutputWithContext(context.Background())
}

func (i *Vpc) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcOutput)
}

// VpcArrayInput is an input type that accepts VpcArray and VpcArrayOutput values.
// You can construct a concrete instance of `VpcArrayInput` via:
//
//	VpcArray{ VpcArgs{...} }
type VpcArrayInput interface {
	pulumi.Input

	ToVpcArrayOutput() VpcArrayOutput
	ToVpcArrayOutputWithContext(context.Context) VpcArrayOutput
}

type VpcArray []VpcInput

func (VpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (i VpcArray) ToVpcArrayOutput() VpcArrayOutput {
	return i.ToVpcArrayOutputWithContext(context.Background())
}

func (i VpcArray) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcArrayOutput)
}

// VpcMapInput is an input type that accepts VpcMap and VpcMapOutput values.
// You can construct a concrete instance of `VpcMapInput` via:
//
//	VpcMap{ "key": VpcArgs{...} }
type VpcMapInput interface {
	pulumi.Input

	ToVpcMapOutput() VpcMapOutput
	ToVpcMapOutputWithContext(context.Context) VpcMapOutput
}

type VpcMap map[string]VpcInput

func (VpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (i VpcMap) ToVpcMapOutput() VpcMapOutput {
	return i.ToVpcMapOutputWithContext(context.Background())
}

func (i VpcMap) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcMapOutput)
}

type VpcOutput struct{ *pulumi.OutputState }

func (VpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vpc)(nil)).Elem()
}

func (o VpcOutput) ToVpcOutput() VpcOutput {
	return o
}

func (o VpcOutput) ToVpcOutputWithContext(ctx context.Context) VpcOutput {
	return o
}

// Amazon Resource Name (ARN) of VPC
func (o VpcOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block. Default is `false`. Conflicts with `ipv6IpamPoolId`
func (o VpcOutput) AssignGeneratedIpv6CidrBlock() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Vpc) pulumi.BoolPtrOutput { return v.AssignGeneratedIpv6CidrBlock }).(pulumi.BoolPtrOutput)
}

// The IPv4 CIDR block for the VPC. CIDR can be explicitly set or it can be derived from IPAM using `ipv4NetmaskLength`.
func (o VpcOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// The ID of the network ACL created by default on VPC creation
func (o VpcOutput) DefaultNetworkAclId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.DefaultNetworkAclId }).(pulumi.StringOutput)
}

// The ID of the route table created by default on VPC creation
func (o VpcOutput) DefaultRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.DefaultRouteTableId }).(pulumi.StringOutput)
}

// The ID of the security group created by default on VPC creation
func (o VpcOutput) DefaultSecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.DefaultSecurityGroupId }).(pulumi.StringOutput)
}

func (o VpcOutput) DhcpOptionsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.DhcpOptionsId }).(pulumi.StringOutput)
}

// A boolean flag to enable/disable ClassicLink
// for the VPC. Only valid in regions and accounts that support EC2 Classic.
// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
//
// Deprecated: With the retirement of EC2-Classic the enable_classiclink attribute has been deprecated and will be removed in a future version.
func (o VpcOutput) EnableClassiclink() pulumi.BoolOutput {
	return o.ApplyT(func(v *Vpc) pulumi.BoolOutput { return v.EnableClassiclink }).(pulumi.BoolOutput)
}

// A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
// Only valid in regions and accounts that support EC2 Classic.
//
// Deprecated: With the retirement of EC2-Classic the enable_classiclink_dns_support attribute has been deprecated and will be removed in a future version.
func (o VpcOutput) EnableClassiclinkDnsSupport() pulumi.BoolOutput {
	return o.ApplyT(func(v *Vpc) pulumi.BoolOutput { return v.EnableClassiclinkDnsSupport }).(pulumi.BoolOutput)
}

// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
func (o VpcOutput) EnableDnsHostnames() pulumi.BoolOutput {
	return o.ApplyT(func(v *Vpc) pulumi.BoolOutput { return v.EnableDnsHostnames }).(pulumi.BoolOutput)
}

// A boolean flag to enable/disable DNS support in the VPC. Defaults to true.
func (o VpcOutput) EnableDnsSupport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Vpc) pulumi.BoolPtrOutput { return v.EnableDnsSupport }).(pulumi.BoolPtrOutput)
}

// Indicates whether Network Address Usage metrics are enabled for your VPC. Defaults to false.
func (o VpcOutput) EnableNetworkAddressUsageMetrics() pulumi.BoolOutput {
	return o.ApplyT(func(v *Vpc) pulumi.BoolOutput { return v.EnableNetworkAddressUsageMetrics }).(pulumi.BoolOutput)
}

// A tenancy option for instances launched into the VPC. Default is `default`, which ensures that EC2 instances launched in this VPC use the EC2 instance tenancy attribute specified when the EC2 instance is launched. The only other option is `dedicated`, which ensures that EC2 instances launched in this VPC are run on dedicated tenancy instances regardless of the tenancy attribute specified at launch. This has a dedicated per region fee of $2 per hour, plus an hourly per instance usage fee.
func (o VpcOutput) InstanceTenancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringPtrOutput { return v.InstanceTenancy }).(pulumi.StringPtrOutput)
}

// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across AWS Regions and accounts. Using IPAM you can monitor IP address usage throughout your AWS Organization.
func (o VpcOutput) Ipv4IpamPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringPtrOutput { return v.Ipv4IpamPoolId }).(pulumi.StringPtrOutput)
}

// The netmask length of the IPv4 CIDR you want to allocate to this VPC. Requires specifying a `ipv4IpamPoolId`.
func (o VpcOutput) Ipv4NetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Vpc) pulumi.IntPtrOutput { return v.Ipv4NetmaskLength }).(pulumi.IntPtrOutput)
}

// The association ID for the IPv6 CIDR block.
func (o VpcOutput) Ipv6AssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.Ipv6AssociationId }).(pulumi.StringOutput)
}

// IPv6 CIDR block to request from an IPAM Pool. Can be set explicitly or derived from IPAM using `ipv6NetmaskLength`.
func (o VpcOutput) Ipv6CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.Ipv6CidrBlock }).(pulumi.StringOutput)
}

// By default when an IPv6 CIDR is assigned to a VPC a default ipv6CidrBlockNetworkBorderGroup will be set to the region of the VPC. This can be changed to restrict advertisement of public addresses to specific Network Border Groups such as LocalZones.
func (o VpcOutput) Ipv6CidrBlockNetworkBorderGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.Ipv6CidrBlockNetworkBorderGroup }).(pulumi.StringOutput)
}

// IPAM Pool ID for a IPv6 pool. Conflicts with `assignGeneratedIpv6CidrBlock`.
func (o VpcOutput) Ipv6IpamPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringPtrOutput { return v.Ipv6IpamPoolId }).(pulumi.StringPtrOutput)
}

// Netmask length to request from IPAM Pool. Conflicts with `ipv6CidrBlock`. This can be omitted if IPAM pool as a `allocationDefaultNetmaskLength` set. Valid values: `56`.
func (o VpcOutput) Ipv6NetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Vpc) pulumi.IntPtrOutput { return v.Ipv6NetmaskLength }).(pulumi.IntPtrOutput)
}

// The ID of the main route table associated with
// this VPC. Note that you can change a VPC's main route table by using an
// `ec2.MainRouteTableAssociation`.
func (o VpcOutput) MainRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.MainRouteTableId }).(pulumi.StringOutput)
}

// The ID of the AWS account that owns the VPC.
func (o VpcOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VpcOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o VpcOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Vpc) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type VpcArrayOutput struct{ *pulumi.OutputState }

func (VpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vpc)(nil)).Elem()
}

func (o VpcArrayOutput) ToVpcArrayOutput() VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) ToVpcArrayOutputWithContext(ctx context.Context) VpcArrayOutput {
	return o
}

func (o VpcArrayOutput) Index(i pulumi.IntInput) VpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vpc {
		return vs[0].([]*Vpc)[vs[1].(int)]
	}).(VpcOutput)
}

type VpcMapOutput struct{ *pulumi.OutputState }

func (VpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vpc)(nil)).Elem()
}

func (o VpcMapOutput) ToVpcMapOutput() VpcMapOutput {
	return o
}

func (o VpcMapOutput) ToVpcMapOutputWithContext(ctx context.Context) VpcMapOutput {
	return o
}

func (o VpcMapOutput) MapIndex(k pulumi.StringInput) VpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vpc {
		return vs[0].(map[string]*Vpc)[vs[1].(string)]
	}).(VpcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcInput)(nil)).Elem(), &Vpc{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcArrayInput)(nil)).Elem(), VpcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcMapInput)(nil)).Elem(), VpcMap{})
	pulumi.RegisterOutputType(VpcOutput{})
	pulumi.RegisterOutputType(VpcArrayOutput{})
	pulumi.RegisterOutputType(VpcMapOutput{})
}
