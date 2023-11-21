// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Specifies information about the Dedicated host.
// Azure REST API version: 2023-03-01. Prior API version in Azure Native 1.x: 2020-12-01.
//
// Other available API versions: 2023-07-01.
type DedicatedHost struct {
	pulumi.CustomResourceState

	// Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
	AutoReplaceOnFailure pulumi.BoolPtrOutput `pulumi:"autoReplaceOnFailure"`
	// A unique id generated and assigned to the dedicated host by the platform. Does not change throughout the lifetime of the host.
	HostId pulumi.StringOutput `pulumi:"hostId"`
	// The dedicated host instance view.
	InstanceView DedicatedHostInstanceViewResponseOutput `pulumi:"instanceView"`
	// Specifies the software license type that will be applied to the VMs deployed on the dedicated host. Possible values are: **None,** **Windows_Server_Hybrid,** **Windows_Server_Perpetual.** The default value is: **None.**
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Fault domain of the dedicated host within a dedicated host group.
	PlatformFaultDomain pulumi.IntPtrOutput `pulumi:"platformFaultDomain"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The date when the host was first provisioned.
	ProvisioningTime pulumi.StringOutput `pulumi:"provisioningTime"`
	// SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the time at which the Dedicated Host resource was created. Minimum api-version: 2021-11-01.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of references to all virtual machines in the Dedicated Host.
	VirtualMachines SubResourceReadOnlyResponseArrayOutput `pulumi:"virtualMachines"`
}

// NewDedicatedHost registers a new resource with the given unique name, arguments, and options.
func NewDedicatedHost(ctx *pulumi.Context,
	name string, args *DedicatedHostArgs, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostGroupName == nil {
		return nil, errors.New("invalid value for required argument 'HostGroupName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20190301:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:DedicatedHost"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230701:DedicatedHost"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DedicatedHost
	err := ctx.RegisterResource("azure-native:compute:DedicatedHost", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedHost gets an existing DedicatedHost resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHostState, opts ...pulumi.ResourceOption) (*DedicatedHost, error) {
	var resource DedicatedHost
	err := ctx.ReadResource("azure-native:compute:DedicatedHost", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedHost resources.
type dedicatedHostState struct {
}

type DedicatedHostState struct {
}

func (DedicatedHostState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostState)(nil)).Elem()
}

type dedicatedHostArgs struct {
	// Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
	AutoReplaceOnFailure *bool `pulumi:"autoReplaceOnFailure"`
	// The name of the dedicated host group.
	HostGroupName string `pulumi:"hostGroupName"`
	// The name of the dedicated host .
	HostName *string `pulumi:"hostName"`
	// Specifies the software license type that will be applied to the VMs deployed on the dedicated host. Possible values are: **None,** **Windows_Server_Hybrid,** **Windows_Server_Perpetual.** The default value is: **None.**
	LicenseType *DedicatedHostLicenseTypes `pulumi:"licenseType"`
	// Resource location
	Location *string `pulumi:"location"`
	// Fault domain of the dedicated host within a dedicated host group.
	PlatformFaultDomain *int `pulumi:"platformFaultDomain"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
	Sku Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DedicatedHost resource.
type DedicatedHostArgs struct {
	// Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
	AutoReplaceOnFailure pulumi.BoolPtrInput
	// The name of the dedicated host group.
	HostGroupName pulumi.StringInput
	// The name of the dedicated host .
	HostName pulumi.StringPtrInput
	// Specifies the software license type that will be applied to the VMs deployed on the dedicated host. Possible values are: **None,** **Windows_Server_Hybrid,** **Windows_Server_Perpetual.** The default value is: **None.**
	LicenseType DedicatedHostLicenseTypesPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Fault domain of the dedicated host within a dedicated host group.
	PlatformFaultDomain pulumi.IntPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
	Sku SkuInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (DedicatedHostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHostArgs)(nil)).Elem()
}

type DedicatedHostInput interface {
	pulumi.Input

	ToDedicatedHostOutput() DedicatedHostOutput
	ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput
}

func (*DedicatedHost) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHost)(nil)).Elem()
}

func (i *DedicatedHost) ToDedicatedHostOutput() DedicatedHostOutput {
	return i.ToDedicatedHostOutputWithContext(context.Background())
}

func (i *DedicatedHost) ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedHostOutput)
}

func (i *DedicatedHost) ToOutput(ctx context.Context) pulumix.Output[*DedicatedHost] {
	return pulumix.Output[*DedicatedHost]{
		OutputState: i.ToDedicatedHostOutputWithContext(ctx).OutputState,
	}
}

type DedicatedHostOutput struct{ *pulumi.OutputState }

func (DedicatedHostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedHost)(nil)).Elem()
}

func (o DedicatedHostOutput) ToDedicatedHostOutput() DedicatedHostOutput {
	return o
}

func (o DedicatedHostOutput) ToDedicatedHostOutputWithContext(ctx context.Context) DedicatedHostOutput {
	return o
}

func (o DedicatedHostOutput) ToOutput(ctx context.Context) pulumix.Output[*DedicatedHost] {
	return pulumix.Output[*DedicatedHost]{
		OutputState: o.OutputState,
	}
}

// Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
func (o DedicatedHostOutput) AutoReplaceOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.BoolPtrOutput { return v.AutoReplaceOnFailure }).(pulumi.BoolPtrOutput)
}

// A unique id generated and assigned to the dedicated host by the platform. Does not change throughout the lifetime of the host.
func (o DedicatedHostOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.HostId }).(pulumi.StringOutput)
}

// The dedicated host instance view.
func (o DedicatedHostOutput) InstanceView() DedicatedHostInstanceViewResponseOutput {
	return o.ApplyT(func(v *DedicatedHost) DedicatedHostInstanceViewResponseOutput { return v.InstanceView }).(DedicatedHostInstanceViewResponseOutput)
}

// Specifies the software license type that will be applied to the VMs deployed on the dedicated host. Possible values are: **None,** **Windows_Server_Hybrid,** **Windows_Server_Perpetual.** The default value is: **None.**
func (o DedicatedHostOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location
func (o DedicatedHostOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o DedicatedHostOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Fault domain of the dedicated host within a dedicated host group.
func (o DedicatedHostOutput) PlatformFaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.IntPtrOutput { return v.PlatformFaultDomain }).(pulumi.IntPtrOutput)
}

// The provisioning state, which only appears in the response.
func (o DedicatedHostOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The date when the host was first provisioned.
func (o DedicatedHostOutput) ProvisioningTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.ProvisioningTime }).(pulumi.StringOutput)
}

// SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
func (o DedicatedHostOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *DedicatedHost) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Resource tags
func (o DedicatedHostOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the time at which the Dedicated Host resource was created. Minimum api-version: 2021-11-01.
func (o DedicatedHostOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

// Resource type
func (o DedicatedHostOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedHost) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A list of references to all virtual machines in the Dedicated Host.
func (o DedicatedHostOutput) VirtualMachines() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v *DedicatedHost) SubResourceReadOnlyResponseArrayOutput { return v.VirtualMachines }).(SubResourceReadOnlyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DedicatedHostOutput{})
}
