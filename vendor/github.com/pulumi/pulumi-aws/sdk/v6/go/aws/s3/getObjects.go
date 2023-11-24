// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// > **NOTE on `maxKeys`:** Retrieving very large numbers of keys can adversely affect the provider's performance.
//
// The objects data source returns keys (i.e., file names) and other metadata about objects in an S3 bucket.
func GetObjects(ctx *pulumi.Context, args *GetObjectsArgs, opts ...pulumi.InvokeOption) (*GetObjectsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetObjectsResult
	err := ctx.Invoke("aws:s3/getObjects:getObjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getObjects.
type GetObjectsArgs struct {
	// Lists object keys in this S3 bucket. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
	Bucket string `pulumi:"bucket"`
	// Character used to group keys (Default: none)
	Delimiter *string `pulumi:"delimiter"`
	// Encodes keys using this method (Default: none; besides none, only "url" can be used)
	EncodingType *string `pulumi:"encodingType"`
	// Boolean specifying whether to populate the owner list (Default: false)
	FetchOwner *bool `pulumi:"fetchOwner"`
	// Maximum object keys to return (Default: 1000)
	MaxKeys *int `pulumi:"maxKeys"`
	// Limits results to object keys with this prefix (Default: none)
	Prefix *string `pulumi:"prefix"`
	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests. If included, the only valid value is `requester`.
	RequestPayer *string `pulumi:"requestPayer"`
	// Returns key names lexicographically after a specific object key in your bucket (Default: none; S3 lists object keys in UTF-8 character encoding in lexicographical order)
	StartAfter *string `pulumi:"startAfter"`
}

// A collection of values returned by getObjects.
type GetObjectsResult struct {
	Bucket string `pulumi:"bucket"`
	// List of any keys between `prefix` and the next occurrence of `delimiter` (i.e., similar to subdirectories of the `prefix` "directory"); the list is only returned when you specify `delimiter`
	CommonPrefixes []string `pulumi:"commonPrefixes"`
	Delimiter      *string  `pulumi:"delimiter"`
	EncodingType   *string  `pulumi:"encodingType"`
	FetchOwner     *bool    `pulumi:"fetchOwner"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of strings representing object keys
	Keys    []string `pulumi:"keys"`
	MaxKeys *int     `pulumi:"maxKeys"`
	// List of strings representing object owner IDs (see `fetchOwner` above)
	Owners []string `pulumi:"owners"`
	Prefix *string  `pulumi:"prefix"`
	// If present, indicates that the requester was successfully charged for the request.
	RequestCharged string  `pulumi:"requestCharged"`
	RequestPayer   *string `pulumi:"requestPayer"`
	StartAfter     *string `pulumi:"startAfter"`
}

func GetObjectsOutput(ctx *pulumi.Context, args GetObjectsOutputArgs, opts ...pulumi.InvokeOption) GetObjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetObjectsResult, error) {
			args := v.(GetObjectsArgs)
			r, err := GetObjects(ctx, &args, opts...)
			var s GetObjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetObjectsResultOutput)
}

// A collection of arguments for invoking getObjects.
type GetObjectsOutputArgs struct {
	// Lists object keys in this S3 bucket. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// Character used to group keys (Default: none)
	Delimiter pulumi.StringPtrInput `pulumi:"delimiter"`
	// Encodes keys using this method (Default: none; besides none, only "url" can be used)
	EncodingType pulumi.StringPtrInput `pulumi:"encodingType"`
	// Boolean specifying whether to populate the owner list (Default: false)
	FetchOwner pulumi.BoolPtrInput `pulumi:"fetchOwner"`
	// Maximum object keys to return (Default: 1000)
	MaxKeys pulumi.IntPtrInput `pulumi:"maxKeys"`
	// Limits results to object keys with this prefix (Default: none)
	Prefix pulumi.StringPtrInput `pulumi:"prefix"`
	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests. If included, the only valid value is `requester`.
	RequestPayer pulumi.StringPtrInput `pulumi:"requestPayer"`
	// Returns key names lexicographically after a specific object key in your bucket (Default: none; S3 lists object keys in UTF-8 character encoding in lexicographical order)
	StartAfter pulumi.StringPtrInput `pulumi:"startAfter"`
}

func (GetObjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetObjectsArgs)(nil)).Elem()
}

// A collection of values returned by getObjects.
type GetObjectsResultOutput struct{ *pulumi.OutputState }

func (GetObjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetObjectsResult)(nil)).Elem()
}

func (o GetObjectsResultOutput) ToGetObjectsResultOutput() GetObjectsResultOutput {
	return o
}

func (o GetObjectsResultOutput) ToGetObjectsResultOutputWithContext(ctx context.Context) GetObjectsResultOutput {
	return o
}

func (o GetObjectsResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectsResult) string { return v.Bucket }).(pulumi.StringOutput)
}

// List of any keys between `prefix` and the next occurrence of `delimiter` (i.e., similar to subdirectories of the `prefix` "directory"); the list is only returned when you specify `delimiter`
func (o GetObjectsResultOutput) CommonPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetObjectsResult) []string { return v.CommonPrefixes }).(pulumi.StringArrayOutput)
}

func (o GetObjectsResultOutput) Delimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetObjectsResult) *string { return v.Delimiter }).(pulumi.StringPtrOutput)
}

func (o GetObjectsResultOutput) EncodingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetObjectsResult) *string { return v.EncodingType }).(pulumi.StringPtrOutput)
}

func (o GetObjectsResultOutput) FetchOwner() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetObjectsResult) *bool { return v.FetchOwner }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetObjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of strings representing object keys
func (o GetObjectsResultOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetObjectsResult) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

func (o GetObjectsResultOutput) MaxKeys() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetObjectsResult) *int { return v.MaxKeys }).(pulumi.IntPtrOutput)
}

// List of strings representing object owner IDs (see `fetchOwner` above)
func (o GetObjectsResultOutput) Owners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetObjectsResult) []string { return v.Owners }).(pulumi.StringArrayOutput)
}

func (o GetObjectsResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetObjectsResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

// If present, indicates that the requester was successfully charged for the request.
func (o GetObjectsResultOutput) RequestCharged() pulumi.StringOutput {
	return o.ApplyT(func(v GetObjectsResult) string { return v.RequestCharged }).(pulumi.StringOutput)
}

func (o GetObjectsResultOutput) RequestPayer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetObjectsResult) *string { return v.RequestPayer }).(pulumi.StringPtrOutput)
}

func (o GetObjectsResultOutput) StartAfter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetObjectsResult) *string { return v.StartAfter }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetObjectsResultOutput{})
}
