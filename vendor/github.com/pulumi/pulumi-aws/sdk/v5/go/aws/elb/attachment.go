// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attaches an EC2 instance to an Elastic Load Balancer (ELB). For attaching resources with Application Load Balancer (ALB) or Network Load Balancer (NLB), see the `lb.TargetGroupAttachment` resource.
//
// > **NOTE on ELB Instances and ELB Attachments:** This provider currently provides
// both a standalone ELB Attachment resource (describing an instance attached to
// an ELB), and an Elastic Load Balancer resource with
// `instances` defined in-line. At this time you cannot use an ELB with in-line
// instances in conjunction with an ELB Attachment resource. Doing so will cause a
// conflict and will overwrite attachments.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := elb.NewAttachment(ctx, "baz", &elb.AttachmentArgs{
//				Elb:      pulumi.Any(aws_elb.Bar.Id),
//				Instance: pulumi.Any(aws_instance.Foo.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Attachment struct {
	pulumi.CustomResourceState

	// The name of the ELB.
	Elb pulumi.StringOutput `pulumi:"elb"`
	// Instance ID to place in the ELB pool.
	Instance pulumi.StringOutput `pulumi:"instance"`
}

// NewAttachment registers a new resource with the given unique name, arguments, and options.
func NewAttachment(ctx *pulumi.Context,
	name string, args *AttachmentArgs, opts ...pulumi.ResourceOption) (*Attachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Elb == nil {
		return nil, errors.New("invalid value for required argument 'Elb'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:elasticloadbalancing/attachment:Attachment"),
		},
	})
	opts = append(opts, aliases)
	var resource Attachment
	err := ctx.RegisterResource("aws:elb/attachment:Attachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachment gets an existing Attachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachmentState, opts ...pulumi.ResourceOption) (*Attachment, error) {
	var resource Attachment
	err := ctx.ReadResource("aws:elb/attachment:Attachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Attachment resources.
type attachmentState struct {
	// The name of the ELB.
	Elb *string `pulumi:"elb"`
	// Instance ID to place in the ELB pool.
	Instance *string `pulumi:"instance"`
}

type AttachmentState struct {
	// The name of the ELB.
	Elb pulumi.StringPtrInput
	// Instance ID to place in the ELB pool.
	Instance pulumi.StringPtrInput
}

func (AttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentState)(nil)).Elem()
}

type attachmentArgs struct {
	// The name of the ELB.
	Elb string `pulumi:"elb"`
	// Instance ID to place in the ELB pool.
	Instance string `pulumi:"instance"`
}

// The set of arguments for constructing a Attachment resource.
type AttachmentArgs struct {
	// The name of the ELB.
	Elb pulumi.StringInput
	// Instance ID to place in the ELB pool.
	Instance pulumi.StringInput
}

func (AttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentArgs)(nil)).Elem()
}

type AttachmentInput interface {
	pulumi.Input

	ToAttachmentOutput() AttachmentOutput
	ToAttachmentOutputWithContext(ctx context.Context) AttachmentOutput
}

func (*Attachment) ElementType() reflect.Type {
	return reflect.TypeOf((**Attachment)(nil)).Elem()
}

func (i *Attachment) ToAttachmentOutput() AttachmentOutput {
	return i.ToAttachmentOutputWithContext(context.Background())
}

func (i *Attachment) ToAttachmentOutputWithContext(ctx context.Context) AttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentOutput)
}

// AttachmentArrayInput is an input type that accepts AttachmentArray and AttachmentArrayOutput values.
// You can construct a concrete instance of `AttachmentArrayInput` via:
//
//	AttachmentArray{ AttachmentArgs{...} }
type AttachmentArrayInput interface {
	pulumi.Input

	ToAttachmentArrayOutput() AttachmentArrayOutput
	ToAttachmentArrayOutputWithContext(context.Context) AttachmentArrayOutput
}

type AttachmentArray []AttachmentInput

func (AttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Attachment)(nil)).Elem()
}

func (i AttachmentArray) ToAttachmentArrayOutput() AttachmentArrayOutput {
	return i.ToAttachmentArrayOutputWithContext(context.Background())
}

func (i AttachmentArray) ToAttachmentArrayOutputWithContext(ctx context.Context) AttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentArrayOutput)
}

// AttachmentMapInput is an input type that accepts AttachmentMap and AttachmentMapOutput values.
// You can construct a concrete instance of `AttachmentMapInput` via:
//
//	AttachmentMap{ "key": AttachmentArgs{...} }
type AttachmentMapInput interface {
	pulumi.Input

	ToAttachmentMapOutput() AttachmentMapOutput
	ToAttachmentMapOutputWithContext(context.Context) AttachmentMapOutput
}

type AttachmentMap map[string]AttachmentInput

func (AttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Attachment)(nil)).Elem()
}

func (i AttachmentMap) ToAttachmentMapOutput() AttachmentMapOutput {
	return i.ToAttachmentMapOutputWithContext(context.Background())
}

func (i AttachmentMap) ToAttachmentMapOutputWithContext(ctx context.Context) AttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachmentMapOutput)
}

type AttachmentOutput struct{ *pulumi.OutputState }

func (AttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Attachment)(nil)).Elem()
}

func (o AttachmentOutput) ToAttachmentOutput() AttachmentOutput {
	return o
}

func (o AttachmentOutput) ToAttachmentOutputWithContext(ctx context.Context) AttachmentOutput {
	return o
}

// The name of the ELB.
func (o AttachmentOutput) Elb() pulumi.StringOutput {
	return o.ApplyT(func(v *Attachment) pulumi.StringOutput { return v.Elb }).(pulumi.StringOutput)
}

// Instance ID to place in the ELB pool.
func (o AttachmentOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *Attachment) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

type AttachmentArrayOutput struct{ *pulumi.OutputState }

func (AttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Attachment)(nil)).Elem()
}

func (o AttachmentArrayOutput) ToAttachmentArrayOutput() AttachmentArrayOutput {
	return o
}

func (o AttachmentArrayOutput) ToAttachmentArrayOutputWithContext(ctx context.Context) AttachmentArrayOutput {
	return o
}

func (o AttachmentArrayOutput) Index(i pulumi.IntInput) AttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Attachment {
		return vs[0].([]*Attachment)[vs[1].(int)]
	}).(AttachmentOutput)
}

type AttachmentMapOutput struct{ *pulumi.OutputState }

func (AttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Attachment)(nil)).Elem()
}

func (o AttachmentMapOutput) ToAttachmentMapOutput() AttachmentMapOutput {
	return o
}

func (o AttachmentMapOutput) ToAttachmentMapOutputWithContext(ctx context.Context) AttachmentMapOutput {
	return o
}

func (o AttachmentMapOutput) MapIndex(k pulumi.StringInput) AttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Attachment {
		return vs[0].(map[string]*Attachment)[vs[1].(string)]
	}).(AttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttachmentInput)(nil)).Elem(), &Attachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttachmentArrayInput)(nil)).Elem(), AttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttachmentMapInput)(nil)).Elem(), AttachmentMap{})
	pulumi.RegisterOutputType(AttachmentOutput{})
	pulumi.RegisterOutputType(AttachmentArrayOutput{})
	pulumi.RegisterOutputType(AttachmentMapOutput{})
}
