// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a DNS zone.
// API Version: 2018-05-01.
type Zone struct {
	pulumi.CustomResourceState

	// The etag of the zone.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets pulumi.Float64Output `pulumi:"maxNumberOfRecordSets"`
	// The maximum number of records per record set that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordsPerRecordSet pulumi.Float64Output `pulumi:"maxNumberOfRecordsPerRecordSet"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
	NumberOfRecordSets pulumi.Float64Output `pulumi:"numberOfRecordSets"`
	// A list of references to virtual networks that register hostnames in this DNS zone. This is a only when ZoneType is Private.
	RegistrationVirtualNetworks SubResourceResponseArrayOutput `pulumi:"registrationVirtualNetworks"`
	// A list of references to virtual networks that resolve records in this DNS zone. This is a only when ZoneType is Private.
	ResolutionVirtualNetworks SubResourceResponseArrayOutput `pulumi:"resolutionVirtualNetworks"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The type of this DNS zone (Public or Private).
	ZoneType pulumi.StringPtrOutput `pulumi:"zoneType"`
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.ZoneType) {
		args.ZoneType = ZoneType("Public")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20150504preview:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160401:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301preview:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180501:Zone"),
		},
	})
	opts = append(opts, aliases)
	var resource Zone
	err := ctx.RegisterResource("azure-native:network:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("azure-native:network:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Zone resources.
type zoneState struct {
}

type ZoneState struct {
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	// Resource location.
	Location *string `pulumi:"location"`
	// A list of references to virtual networks that register hostnames in this DNS zone. This is a only when ZoneType is Private.
	RegistrationVirtualNetworks []SubResource `pulumi:"registrationVirtualNetworks"`
	// A list of references to virtual networks that resolve records in this DNS zone. This is a only when ZoneType is Private.
	ResolutionVirtualNetworks []SubResource `pulumi:"resolutionVirtualNetworks"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the DNS zone (without a terminating dot).
	ZoneName *string `pulumi:"zoneName"`
	// The type of this DNS zone (Public or Private).
	ZoneType *ZoneType `pulumi:"zoneType"`
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// Resource location.
	Location pulumi.StringPtrInput
	// A list of references to virtual networks that register hostnames in this DNS zone. This is a only when ZoneType is Private.
	RegistrationVirtualNetworks SubResourceArrayInput
	// A list of references to virtual networks that resolve records in this DNS zone. This is a only when ZoneType is Private.
	ResolutionVirtualNetworks SubResourceArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the DNS zone (without a terminating dot).
	ZoneName pulumi.StringPtrInput
	// The type of this DNS zone (Public or Private).
	ZoneType ZoneTypePtrInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

type ZoneOutput struct{ *pulumi.OutputState }

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Zone)(nil)).Elem()
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

// The etag of the zone.
func (o ZoneOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o ZoneOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The maximum number of record sets that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
func (o ZoneOutput) MaxNumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v *Zone) pulumi.Float64Output { return v.MaxNumberOfRecordSets }).(pulumi.Float64Output)
}

// The maximum number of records per record set that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
func (o ZoneOutput) MaxNumberOfRecordsPerRecordSet() pulumi.Float64Output {
	return o.ApplyT(func(v *Zone) pulumi.Float64Output { return v.MaxNumberOfRecordsPerRecordSet }).(pulumi.Float64Output)
}

// Resource name.
func (o ZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name servers for this DNS zone. This is a read-only property and any attempt to set this value will be ignored.
func (o ZoneOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringArrayOutput { return v.NameServers }).(pulumi.StringArrayOutput)
}

// The current number of record sets in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
func (o ZoneOutput) NumberOfRecordSets() pulumi.Float64Output {
	return o.ApplyT(func(v *Zone) pulumi.Float64Output { return v.NumberOfRecordSets }).(pulumi.Float64Output)
}

// A list of references to virtual networks that register hostnames in this DNS zone. This is a only when ZoneType is Private.
func (o ZoneOutput) RegistrationVirtualNetworks() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *Zone) SubResourceResponseArrayOutput { return v.RegistrationVirtualNetworks }).(SubResourceResponseArrayOutput)
}

// A list of references to virtual networks that resolve records in this DNS zone. This is a only when ZoneType is Private.
func (o ZoneOutput) ResolutionVirtualNetworks() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *Zone) SubResourceResponseArrayOutput { return v.ResolutionVirtualNetworks }).(SubResourceResponseArrayOutput)
}

// Resource tags.
func (o ZoneOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ZoneOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The type of this DNS zone (Public or Private).
func (o ZoneOutput) ZoneType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Zone) pulumi.StringPtrOutput { return v.ZoneType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ZoneOutput{})
}
