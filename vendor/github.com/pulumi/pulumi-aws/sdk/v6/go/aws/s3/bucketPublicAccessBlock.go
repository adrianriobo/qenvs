// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages S3 bucket-level Public Access Block configuration. For more information about these settings, see the [AWS S3 Block Public Access documentation](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketPublicAccessBlock(ctx, "exampleBucketPublicAccessBlock", &s3.BucketPublicAccessBlockArgs{
//				Bucket:                exampleBucketV2.ID(),
//				BlockPublicAcls:       pulumi.Bool(true),
//				BlockPublicPolicy:     pulumi.Bool(true),
//				IgnorePublicAcls:      pulumi.Bool(true),
//				RestrictPublicBuckets: pulumi.Bool(true),
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
// Using `pulumi import`, import `aws_s3_bucket_public_access_block` using the bucket name. For example:
//
// ```sh
//
//	$ pulumi import aws:s3/bucketPublicAccessBlock:BucketPublicAccessBlock example my-bucket
//
// ```
type BucketPublicAccessBlock struct {
	pulumi.CustomResourceState

	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls pulumi.BoolPtrOutput `pulumi:"blockPublicAcls"`
	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy pulumi.BoolPtrOutput `pulumi:"blockPublicPolicy"`
	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore public ACLs on this bucket and any objects that it contains.
	IgnorePublicAcls pulumi.BoolPtrOutput `pulumi:"ignorePublicAcls"`
	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
	RestrictPublicBuckets pulumi.BoolPtrOutput `pulumi:"restrictPublicBuckets"`
}

// NewBucketPublicAccessBlock registers a new resource with the given unique name, arguments, and options.
func NewBucketPublicAccessBlock(ctx *pulumi.Context,
	name string, args *BucketPublicAccessBlockArgs, opts ...pulumi.ResourceOption) (*BucketPublicAccessBlock, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BucketPublicAccessBlock
	err := ctx.RegisterResource("aws:s3/bucketPublicAccessBlock:BucketPublicAccessBlock", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketPublicAccessBlock gets an existing BucketPublicAccessBlock resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketPublicAccessBlock(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketPublicAccessBlockState, opts ...pulumi.ResourceOption) (*BucketPublicAccessBlock, error) {
	var resource BucketPublicAccessBlock
	err := ctx.ReadResource("aws:s3/bucketPublicAccessBlock:BucketPublicAccessBlock", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketPublicAccessBlock resources.
type bucketPublicAccessBlockState struct {
	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls *bool `pulumi:"blockPublicAcls"`
	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy *bool `pulumi:"blockPublicPolicy"`
	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket *string `pulumi:"bucket"`
	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore public ACLs on this bucket and any objects that it contains.
	IgnorePublicAcls *bool `pulumi:"ignorePublicAcls"`
	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
	RestrictPublicBuckets *bool `pulumi:"restrictPublicBuckets"`
}

type BucketPublicAccessBlockState struct {
	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy pulumi.BoolPtrInput
	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket pulumi.StringPtrInput
	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore public ACLs on this bucket and any objects that it contains.
	IgnorePublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
	RestrictPublicBuckets pulumi.BoolPtrInput
}

func (BucketPublicAccessBlockState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPublicAccessBlockState)(nil)).Elem()
}

type bucketPublicAccessBlockArgs struct {
	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls *bool `pulumi:"blockPublicAcls"`
	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy *bool `pulumi:"blockPublicPolicy"`
	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket string `pulumi:"bucket"`
	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore public ACLs on this bucket and any objects that it contains.
	IgnorePublicAcls *bool `pulumi:"ignorePublicAcls"`
	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
	RestrictPublicBuckets *bool `pulumi:"restrictPublicBuckets"`
}

// The set of arguments for constructing a BucketPublicAccessBlock resource.
type BucketPublicAccessBlockArgs struct {
	// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
	// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
	// * PUT Object calls will fail if the request includes an object ACL.
	BlockPublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
	// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	BlockPublicPolicy pulumi.BoolPtrInput
	// S3 Bucket to which this Public Access Block configuration should be applied.
	Bucket pulumi.StringInput
	// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
	// * Ignore public ACLs on this bucket and any objects that it contains.
	IgnorePublicAcls pulumi.BoolPtrInput
	// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
	// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
	RestrictPublicBuckets pulumi.BoolPtrInput
}

func (BucketPublicAccessBlockArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketPublicAccessBlockArgs)(nil)).Elem()
}

type BucketPublicAccessBlockInput interface {
	pulumi.Input

	ToBucketPublicAccessBlockOutput() BucketPublicAccessBlockOutput
	ToBucketPublicAccessBlockOutputWithContext(ctx context.Context) BucketPublicAccessBlockOutput
}

func (*BucketPublicAccessBlock) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketPublicAccessBlock)(nil)).Elem()
}

func (i *BucketPublicAccessBlock) ToBucketPublicAccessBlockOutput() BucketPublicAccessBlockOutput {
	return i.ToBucketPublicAccessBlockOutputWithContext(context.Background())
}

func (i *BucketPublicAccessBlock) ToBucketPublicAccessBlockOutputWithContext(ctx context.Context) BucketPublicAccessBlockOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPublicAccessBlockOutput)
}

// BucketPublicAccessBlockArrayInput is an input type that accepts BucketPublicAccessBlockArray and BucketPublicAccessBlockArrayOutput values.
// You can construct a concrete instance of `BucketPublicAccessBlockArrayInput` via:
//
//	BucketPublicAccessBlockArray{ BucketPublicAccessBlockArgs{...} }
type BucketPublicAccessBlockArrayInput interface {
	pulumi.Input

	ToBucketPublicAccessBlockArrayOutput() BucketPublicAccessBlockArrayOutput
	ToBucketPublicAccessBlockArrayOutputWithContext(context.Context) BucketPublicAccessBlockArrayOutput
}

type BucketPublicAccessBlockArray []BucketPublicAccessBlockInput

func (BucketPublicAccessBlockArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketPublicAccessBlock)(nil)).Elem()
}

func (i BucketPublicAccessBlockArray) ToBucketPublicAccessBlockArrayOutput() BucketPublicAccessBlockArrayOutput {
	return i.ToBucketPublicAccessBlockArrayOutputWithContext(context.Background())
}

func (i BucketPublicAccessBlockArray) ToBucketPublicAccessBlockArrayOutputWithContext(ctx context.Context) BucketPublicAccessBlockArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPublicAccessBlockArrayOutput)
}

// BucketPublicAccessBlockMapInput is an input type that accepts BucketPublicAccessBlockMap and BucketPublicAccessBlockMapOutput values.
// You can construct a concrete instance of `BucketPublicAccessBlockMapInput` via:
//
//	BucketPublicAccessBlockMap{ "key": BucketPublicAccessBlockArgs{...} }
type BucketPublicAccessBlockMapInput interface {
	pulumi.Input

	ToBucketPublicAccessBlockMapOutput() BucketPublicAccessBlockMapOutput
	ToBucketPublicAccessBlockMapOutputWithContext(context.Context) BucketPublicAccessBlockMapOutput
}

type BucketPublicAccessBlockMap map[string]BucketPublicAccessBlockInput

func (BucketPublicAccessBlockMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketPublicAccessBlock)(nil)).Elem()
}

func (i BucketPublicAccessBlockMap) ToBucketPublicAccessBlockMapOutput() BucketPublicAccessBlockMapOutput {
	return i.ToBucketPublicAccessBlockMapOutputWithContext(context.Background())
}

func (i BucketPublicAccessBlockMap) ToBucketPublicAccessBlockMapOutputWithContext(ctx context.Context) BucketPublicAccessBlockMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketPublicAccessBlockMapOutput)
}

type BucketPublicAccessBlockOutput struct{ *pulumi.OutputState }

func (BucketPublicAccessBlockOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketPublicAccessBlock)(nil)).Elem()
}

func (o BucketPublicAccessBlockOutput) ToBucketPublicAccessBlockOutput() BucketPublicAccessBlockOutput {
	return o
}

func (o BucketPublicAccessBlockOutput) ToBucketPublicAccessBlockOutputWithContext(ctx context.Context) BucketPublicAccessBlockOutput {
	return o
}

// Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
// * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
// * PUT Object calls will fail if the request includes an object ACL.
func (o BucketPublicAccessBlockOutput) BlockPublicAcls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketPublicAccessBlock) pulumi.BoolPtrOutput { return v.BlockPublicAcls }).(pulumi.BoolPtrOutput)
}

// Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
// * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
func (o BucketPublicAccessBlockOutput) BlockPublicPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketPublicAccessBlock) pulumi.BoolPtrOutput { return v.BlockPublicPolicy }).(pulumi.BoolPtrOutput)
}

// S3 Bucket to which this Public Access Block configuration should be applied.
func (o BucketPublicAccessBlockOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketPublicAccessBlock) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
// * Ignore public ACLs on this bucket and any objects that it contains.
func (o BucketPublicAccessBlockOutput) IgnorePublicAcls() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketPublicAccessBlock) pulumi.BoolPtrOutput { return v.IgnorePublicAcls }).(pulumi.BoolPtrOutput)
}

// Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
// * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
func (o BucketPublicAccessBlockOutput) RestrictPublicBuckets() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BucketPublicAccessBlock) pulumi.BoolPtrOutput { return v.RestrictPublicBuckets }).(pulumi.BoolPtrOutput)
}

type BucketPublicAccessBlockArrayOutput struct{ *pulumi.OutputState }

func (BucketPublicAccessBlockArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketPublicAccessBlock)(nil)).Elem()
}

func (o BucketPublicAccessBlockArrayOutput) ToBucketPublicAccessBlockArrayOutput() BucketPublicAccessBlockArrayOutput {
	return o
}

func (o BucketPublicAccessBlockArrayOutput) ToBucketPublicAccessBlockArrayOutputWithContext(ctx context.Context) BucketPublicAccessBlockArrayOutput {
	return o
}

func (o BucketPublicAccessBlockArrayOutput) Index(i pulumi.IntInput) BucketPublicAccessBlockOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketPublicAccessBlock {
		return vs[0].([]*BucketPublicAccessBlock)[vs[1].(int)]
	}).(BucketPublicAccessBlockOutput)
}

type BucketPublicAccessBlockMapOutput struct{ *pulumi.OutputState }

func (BucketPublicAccessBlockMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketPublicAccessBlock)(nil)).Elem()
}

func (o BucketPublicAccessBlockMapOutput) ToBucketPublicAccessBlockMapOutput() BucketPublicAccessBlockMapOutput {
	return o
}

func (o BucketPublicAccessBlockMapOutput) ToBucketPublicAccessBlockMapOutputWithContext(ctx context.Context) BucketPublicAccessBlockMapOutput {
	return o
}

func (o BucketPublicAccessBlockMapOutput) MapIndex(k pulumi.StringInput) BucketPublicAccessBlockOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketPublicAccessBlock {
		return vs[0].(map[string]*BucketPublicAccessBlock)[vs[1].(string)]
	}).(BucketPublicAccessBlockOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPublicAccessBlockInput)(nil)).Elem(), &BucketPublicAccessBlock{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPublicAccessBlockArrayInput)(nil)).Elem(), BucketPublicAccessBlockArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketPublicAccessBlockMapInput)(nil)).Elem(), BucketPublicAccessBlockMap{})
	pulumi.RegisterOutputType(BucketPublicAccessBlockOutput{})
	pulumi.RegisterOutputType(BucketPublicAccessBlockArrayOutput{})
	pulumi.RegisterOutputType(BucketPublicAccessBlockMapOutput{})
}
