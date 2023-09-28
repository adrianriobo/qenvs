// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Describes a DNS record set (a collection of DNS records with the same name and type).
// Azure REST API version: 2018-05-01. Prior API version in Azure Native 1.x: 2018-05-01
type RecordSet struct {
	pulumi.CustomResourceState

	// The list of A records in the record set.
	ARecords ARecordResponseArrayOutput `pulumi:"aRecords"`
	// The list of AAAA records in the record set.
	AaaaRecords AaaaRecordResponseArrayOutput `pulumi:"aaaaRecords"`
	// The list of CAA records in the record set.
	CaaRecords CaaRecordResponseArrayOutput `pulumi:"caaRecords"`
	// The CNAME record in the  record set.
	CnameRecord CnameRecordResponsePtrOutput `pulumi:"cnameRecord"`
	// The etag of the record set.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Fully qualified domain name of the record set.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The metadata attached to the record set.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The list of MX records in the record set.
	MxRecords MxRecordResponseArrayOutput `pulumi:"mxRecords"`
	// The name of the record set.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of NS records in the record set.
	NsRecords NsRecordResponseArrayOutput `pulumi:"nsRecords"`
	// provisioning State of the record set.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The list of PTR records in the record set.
	PtrRecords PtrRecordResponseArrayOutput `pulumi:"ptrRecords"`
	// The SOA record in the record set.
	SoaRecord SoaRecordResponsePtrOutput `pulumi:"soaRecord"`
	// The list of SRV records in the record set.
	SrvRecords SrvRecordResponseArrayOutput `pulumi:"srvRecords"`
	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource SubResourceResponsePtrOutput `pulumi:"targetResource"`
	// The TTL (time-to-live) of the records in the record set.
	Ttl pulumi.Float64PtrOutput `pulumi:"ttl"`
	// The list of TXT records in the record set.
	TxtRecords TxtRecordResponseArrayOutput `pulumi:"txtRecords"`
	// The type of the record set.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRecordSet registers a new resource with the given unique name, arguments, and options.
func NewRecordSet(ctx *pulumi.Context,
	name string, args *RecordSetArgs, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RecordType == nil {
		return nil, errors.New("invalid value for required argument 'RecordType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ZoneName == nil {
		return nil, errors.New("invalid value for required argument 'ZoneName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20150504preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160401:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180501:RecordSet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230701preview:RecordSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RecordSet
	err := ctx.RegisterResource("azure-native:network:RecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecordSet gets an existing RecordSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetState, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	var resource RecordSet
	err := ctx.ReadResource("azure-native:network:RecordSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RecordSet resources.
type recordSetState struct {
}

type RecordSetState struct {
}

func (RecordSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetState)(nil)).Elem()
}

type recordSetArgs struct {
	// The list of A records in the record set.
	ARecords []ARecord `pulumi:"aRecords"`
	// The list of AAAA records in the record set.
	AaaaRecords []AaaaRecord `pulumi:"aaaaRecords"`
	// The list of CAA records in the record set.
	CaaRecords []CaaRecord `pulumi:"caaRecords"`
	// The CNAME record in the  record set.
	CnameRecord *CnameRecord `pulumi:"cnameRecord"`
	// The metadata attached to the record set.
	Metadata map[string]string `pulumi:"metadata"`
	// The list of MX records in the record set.
	MxRecords []MxRecord `pulumi:"mxRecords"`
	// The list of NS records in the record set.
	NsRecords []NsRecord `pulumi:"nsRecords"`
	// The list of PTR records in the record set.
	PtrRecords []PtrRecord `pulumi:"ptrRecords"`
	// The type of DNS record in this record set. Record sets of type SOA can be updated but not created (they are created when the DNS zone is created).
	RecordType string `pulumi:"recordType"`
	// The name of the record set, relative to the name of the zone.
	RelativeRecordSetName *string `pulumi:"relativeRecordSetName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SOA record in the record set.
	SoaRecord *SoaRecord `pulumi:"soaRecord"`
	// The list of SRV records in the record set.
	SrvRecords []SrvRecord `pulumi:"srvRecords"`
	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource *SubResource `pulumi:"targetResource"`
	// The TTL (time-to-live) of the records in the record set.
	Ttl *float64 `pulumi:"ttl"`
	// The list of TXT records in the record set.
	TxtRecords []TxtRecord `pulumi:"txtRecords"`
	// The name of the DNS zone (without a terminating dot).
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a RecordSet resource.
type RecordSetArgs struct {
	// The list of A records in the record set.
	ARecords ARecordArrayInput
	// The list of AAAA records in the record set.
	AaaaRecords AaaaRecordArrayInput
	// The list of CAA records in the record set.
	CaaRecords CaaRecordArrayInput
	// The CNAME record in the  record set.
	CnameRecord CnameRecordPtrInput
	// The metadata attached to the record set.
	Metadata pulumi.StringMapInput
	// The list of MX records in the record set.
	MxRecords MxRecordArrayInput
	// The list of NS records in the record set.
	NsRecords NsRecordArrayInput
	// The list of PTR records in the record set.
	PtrRecords PtrRecordArrayInput
	// The type of DNS record in this record set. Record sets of type SOA can be updated but not created (they are created when the DNS zone is created).
	RecordType pulumi.StringInput
	// The name of the record set, relative to the name of the zone.
	RelativeRecordSetName pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The SOA record in the record set.
	SoaRecord SoaRecordPtrInput
	// The list of SRV records in the record set.
	SrvRecords SrvRecordArrayInput
	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource SubResourcePtrInput
	// The TTL (time-to-live) of the records in the record set.
	Ttl pulumi.Float64PtrInput
	// The list of TXT records in the record set.
	TxtRecords TxtRecordArrayInput
	// The name of the DNS zone (without a terminating dot).
	ZoneName pulumi.StringInput
}

func (RecordSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetArgs)(nil)).Elem()
}

type RecordSetInput interface {
	pulumi.Input

	ToRecordSetOutput() RecordSetOutput
	ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput
}

func (*RecordSet) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordSet)(nil)).Elem()
}

func (i *RecordSet) ToRecordSetOutput() RecordSetOutput {
	return i.ToRecordSetOutputWithContext(context.Background())
}

func (i *RecordSet) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecordSetOutput)
}

func (i *RecordSet) ToOutput(ctx context.Context) pulumix.Output[*RecordSet] {
	return pulumix.Output[*RecordSet]{
		OutputState: i.ToRecordSetOutputWithContext(ctx).OutputState,
	}
}

type RecordSetOutput struct{ *pulumi.OutputState }

func (RecordSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecordSet)(nil)).Elem()
}

func (o RecordSetOutput) ToRecordSetOutput() RecordSetOutput {
	return o
}

func (o RecordSetOutput) ToRecordSetOutputWithContext(ctx context.Context) RecordSetOutput {
	return o
}

func (o RecordSetOutput) ToOutput(ctx context.Context) pulumix.Output[*RecordSet] {
	return pulumix.Output[*RecordSet]{
		OutputState: o.OutputState,
	}
}

// The list of A records in the record set.
func (o RecordSetOutput) ARecords() ARecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) ARecordResponseArrayOutput { return v.ARecords }).(ARecordResponseArrayOutput)
}

// The list of AAAA records in the record set.
func (o RecordSetOutput) AaaaRecords() AaaaRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) AaaaRecordResponseArrayOutput { return v.AaaaRecords }).(AaaaRecordResponseArrayOutput)
}

// The list of CAA records in the record set.
func (o RecordSetOutput) CaaRecords() CaaRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) CaaRecordResponseArrayOutput { return v.CaaRecords }).(CaaRecordResponseArrayOutput)
}

// The CNAME record in the  record set.
func (o RecordSetOutput) CnameRecord() CnameRecordResponsePtrOutput {
	return o.ApplyT(func(v *RecordSet) CnameRecordResponsePtrOutput { return v.CnameRecord }).(CnameRecordResponsePtrOutput)
}

// The etag of the record set.
func (o RecordSetOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified domain name of the record set.
func (o RecordSetOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// The metadata attached to the record set.
func (o RecordSetOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The list of MX records in the record set.
func (o RecordSetOutput) MxRecords() MxRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) MxRecordResponseArrayOutput { return v.MxRecords }).(MxRecordResponseArrayOutput)
}

// The name of the record set.
func (o RecordSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of NS records in the record set.
func (o RecordSetOutput) NsRecords() NsRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) NsRecordResponseArrayOutput { return v.NsRecords }).(NsRecordResponseArrayOutput)
}

// provisioning State of the record set.
func (o RecordSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of PTR records in the record set.
func (o RecordSetOutput) PtrRecords() PtrRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) PtrRecordResponseArrayOutput { return v.PtrRecords }).(PtrRecordResponseArrayOutput)
}

// The SOA record in the record set.
func (o RecordSetOutput) SoaRecord() SoaRecordResponsePtrOutput {
	return o.ApplyT(func(v *RecordSet) SoaRecordResponsePtrOutput { return v.SoaRecord }).(SoaRecordResponsePtrOutput)
}

// The list of SRV records in the record set.
func (o RecordSetOutput) SrvRecords() SrvRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) SrvRecordResponseArrayOutput { return v.SrvRecords }).(SrvRecordResponseArrayOutput)
}

// A reference to an azure resource from where the dns resource value is taken.
func (o RecordSetOutput) TargetResource() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *RecordSet) SubResourceResponsePtrOutput { return v.TargetResource }).(SubResourceResponsePtrOutput)
}

// The TTL (time-to-live) of the records in the record set.
func (o RecordSetOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.Float64PtrOutput { return v.Ttl }).(pulumi.Float64PtrOutput)
}

// The list of TXT records in the record set.
func (o RecordSetOutput) TxtRecords() TxtRecordResponseArrayOutput {
	return o.ApplyT(func(v *RecordSet) TxtRecordResponseArrayOutput { return v.TxtRecords }).(TxtRecordResponseArrayOutput)
}

// The type of the record set.
func (o RecordSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RecordSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RecordSetOutput{})
}
