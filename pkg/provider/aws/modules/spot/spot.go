package spot

import (
	"context"
	"reflect"

	"errors"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type BestSpotChoice struct {
	pulumi.CustomResourceState

	// Region for the best spot choice
	Region pulumi.StringOutput `pulumi:"region"`
	// AvailabilityZone for the best spot choice
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// Average price for spot
	AVGPrice pulumi.Float64PtrOutput `pulumi:"avgprice"`
	// Max price for spot
	MaxPrice pulumi.Float64PtrOutput `pulumi:"maxprice"`
	// Placement score
	Score pulumi.IntPtrOutput `pulumi:"score"`
}

// https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSpotPriceHistory.html
type bestSpotChoiceArgs struct {
	// Product description to filter requisites for the target machine
	ProductDescription *string `pulumi:"productDescription"`
	// Types of instances for the target machine
	InstanceTypes []string `pulumi:"instanceTypes"`
	// AMI name for the target machine
	AMIName *string `pulumi:"aminame"`
}

// The set of arguments for obtaining the best spot choice
type BestSpotChoiceArgs struct {
	// Product description to filter requisites for the target machine
	ProductDescription pulumi.StringPtrInput
	// Types of instances for the target machine
	InstanceTypes pulumi.StringArrayInput
	// AMI name for the target machine
	AMIName pulumi.StringPtrInput
}

// NewBestSpotChoice registers a new resource with the given unique name, arguments, and options.
func NewBestSpotChoice(ctx *pulumi.Context,
	name string, args *BestSpotChoiceArgs, opts ...pulumi.ResourceOption) (*BestSpotChoice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductDescription == nil {
		return nil, errors.New("invalid value for required argument 'ProductDescription'")
	}
	// if len(args.InstanceTypes) = 0 {
	// 	return nil, errors.New("invalid value for required argument 'ProductDescription'")
	// }
	var resource BestSpotChoice
	// here we will run the logic
	// choiceBestSpot(args)
	err := ctx.RegisterResource("rh:qe/aws:spot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBestSpotChoice gets an existing BestSpotChoice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBestSpotChoice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BestSpotChoiceState, opts ...pulumi.ResourceOption) (*BestSpotChoice, error) {
	var resource BestSpotChoice
	err := ctx.ReadResource("rh:qe/aws:spot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BestSpotChoice resources.
type bestSpotChoiceState struct {
	Region           *string  `pulumi:"region"`
	AvailabilityZone *string  `pulumi:"availabilityZone"`
	AVGPrice         *float64 `pulumi:"avgrice"`
	MaxPrice         *float64 `pulumi:"maxprice"`
	Score            *int64   `pulumi:"score"`
}

type BestSpotChoiceState struct {
	Region           pulumi.StringPtrInput
	AvailabilityZone pulumi.StringPtrInput
	AVGPrice         pulumi.Float64PtrInput
	MaxPrice         pulumi.Float64PtrInput
	Score            pulumi.IntPtrInput
}

func (BestSpotChoiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*bestSpotChoiceState)(nil)).Elem()
}

func (BestSpotChoiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bestSpotChoiceArgs)(nil)).Elem()
}

type BestSpotChoiceInput interface {
	pulumi.Input

	ToBestSpotChoiceOutput() BestSpotChoiceOutput
	ToBestSpotChoiceWithContext(ctx context.Context) BestSpotChoiceOutput
}

func (*BestSpotChoice) ElementType() reflect.Type {
	return reflect.TypeOf((**BestSpotChoice)(nil)).Elem()
}

func (i *BestSpotChoice) ToBestSpotChoiceOutput() BestSpotChoiceOutput {
	return i.ToBestSpotChoiceOutputWithContext(context.Background())
}

func (i *BestSpotChoice) ToBestSpotChoiceOutputWithContext(ctx context.Context) BestSpotChoiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BestSpotChoiceOutput)
}

func (i *BestSpotChoice) ToOutput(ctx context.Context) pulumix.Output[*BestSpotChoice] {
	return pulumix.Output[*BestSpotChoice]{
		OutputState: i.ToBestSpotChoiceOutputWithContext(ctx).OutputState,
	}
}

// BestSpotChoiceArrayInput is an input type that accepts BestSpotChoiceArray and BestSpotChoiceArrayOutput values.
// You can construct a concrete instance of `BestSpotChoiceArrayInput` via:
//
//	BestSpotChoiceArray{ BestSpotChoiceArgs{...} }
type BestSpotChoiceArrayInput interface {
	pulumi.Input

	ToBestSpotChoiceArrayOutput() BestSpotChoiceArrayOutput
	ToBestSpotChoiceArrayOutputWithContext(context.Context) BestSpotChoiceArrayOutput
}

type BestSpotChoiceArray []BestSpotChoiceInput

func (BestSpotChoiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BestSpotChoice)(nil)).Elem()
}

func (i BestSpotChoiceArray) ToBestSpotChoiceArrayOutput() BestSpotChoiceArrayOutput {
	return i.ToBestSpotChoiceArrayOutputWithContext(context.Background())
}

func (i BestSpotChoiceArray) ToBestSpotChoiceArrayOutputWithContext(ctx context.Context) BestSpotChoiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BestSpotChoiceArrayOutput)
}

func (i BestSpotChoiceArray) ToOutput(ctx context.Context) pulumix.Output[[]*BestSpotChoice] {
	return pulumix.Output[[]*BestSpotChoice]{
		OutputState: i.ToBestSpotChoiceArrayOutputWithContext(ctx).OutputState,
	}
}

// BestSpotChoiceMapInput is an input type that accepts BestSpotChoiceMap and BestSpotChoiceMapOutput values.
// You can construct a concrete instance of `BestSpotChoiceMapInput` via:
//
//	BestSpotChoiceMap{ "key": BestSpotChoiceArgs{...} }
type BestSpotChoiceMapInput interface {
	pulumi.Input

	ToBestSpotChoiceMapOutput() BestSpotChoiceMapOutput
	ToBestSpotChoiceMapOutputWithContext(context.Context) BestSpotChoiceMapOutput
}

type BestSpotChoiceMap map[string]BestSpotChoiceInput

func (BestSpotChoiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BestSpotChoice)(nil)).Elem()
}

func (i BestSpotChoiceMap) ToBestSpotChoiceMapOutput() BestSpotChoiceMapOutput {
	return i.ToBestSpotChoiceMapOutputWithContext(context.Background())
}

func (i BestSpotChoiceMap) ToBestSpotChoiceMapOutputWithContext(ctx context.Context) BestSpotChoiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BestSpotChoiceMapOutput)
}

func (i BestSpotChoiceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*BestSpotChoice] {
	return pulumix.Output[map[string]*BestSpotChoice]{
		OutputState: i.ToBestSpotChoiceMapOutputWithContext(ctx).OutputState,
	}
}

type BestSpotChoiceOutput struct{ *pulumi.OutputState }

func (BestSpotChoiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BestSpotChoice)(nil)).Elem()
}

func (o BestSpotChoiceOutput) ToBestSpotChoiceOutput() BestSpotChoiceOutput {
	return o
}

func (o BestSpotChoiceOutput) ToBestSpotChoiceOutputWithContext(ctx context.Context) BestSpotChoiceOutput {
	return o
}

func (o BestSpotChoiceOutput) ToOutput(ctx context.Context) pulumix.Output[*BestSpotChoice] {
	return pulumix.Output[*BestSpotChoice]{
		OutputState: o.OutputState,
	}
}

func (o BestSpotChoiceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *BestSpotChoice) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o BestSpotChoiceOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BestSpotChoice) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o BestSpotChoiceOutput) AVGPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BestSpotChoice) pulumi.Float64PtrOutput { return v.AVGPrice }).(pulumi.Float64PtrOutput)
}

func (o BestSpotChoiceOutput) MaxPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BestSpotChoice) pulumi.Float64PtrOutput { return v.MaxPrice }).(pulumi.Float64PtrOutput)
}

func (o BestSpotChoiceOutput) Score() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BestSpotChoice) pulumi.IntPtrOutput { return v.Score }).(pulumi.IntPtrOutput)
}

type BestSpotChoiceArrayOutput struct{ *pulumi.OutputState }

func (BestSpotChoiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BestSpotChoice)(nil)).Elem()
}

func (o BestSpotChoiceArrayOutput) ToBestSpotChoiceArrayOutput() BestSpotChoiceArrayOutput {
	return o
}

func (o BestSpotChoiceArrayOutput) ToBestSpotChoiceArrayOutputWithContext(ctx context.Context) BestSpotChoiceArrayOutput {
	return o
}

func (o BestSpotChoiceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*BestSpotChoice] {
	return pulumix.Output[[]*BestSpotChoice]{
		OutputState: o.OutputState,
	}
}

func (o BestSpotChoiceArrayOutput) Index(i pulumi.IntInput) BestSpotChoiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BestSpotChoice {
		return vs[0].([]*BestSpotChoice)[vs[1].(int)]
	}).(BestSpotChoiceOutput)
}

type BestSpotChoiceMapOutput struct{ *pulumi.OutputState }

func (BestSpotChoiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BestSpotChoice)(nil)).Elem()
}

func (o BestSpotChoiceMapOutput) ToBestSpotChoiceMapOutput() BestSpotChoiceMapOutput {
	return o
}

func (o BestSpotChoiceMapOutput) ToBestSpotChoiceMapOutputWithContext(ctx context.Context) BestSpotChoiceMapOutput {
	return o
}

func (o BestSpotChoiceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*BestSpotChoice] {
	return pulumix.Output[map[string]*BestSpotChoice]{
		OutputState: o.OutputState,
	}
}

func (o BestSpotChoiceMapOutput) MapIndex(k pulumi.StringInput) BestSpotChoiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BestSpotChoice {
		return vs[0].(map[string]*BestSpotChoice)[vs[1].(string)]
	}).(BestSpotChoiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BestSpotChoiceInput)(nil)).Elem(), &BestSpotChoice{})
	pulumi.RegisterInputType(reflect.TypeOf((*BestSpotChoiceArrayInput)(nil)).Elem(), BestSpotChoiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BestSpotChoiceMapInput)(nil)).Elem(), BestSpotChoiceMap{})
	pulumi.RegisterOutputType(BestSpotChoiceOutput{})
	pulumi.RegisterOutputType(BestSpotChoiceArrayOutput{})
	pulumi.RegisterOutputType(BestSpotChoiceMapOutput{})
}
