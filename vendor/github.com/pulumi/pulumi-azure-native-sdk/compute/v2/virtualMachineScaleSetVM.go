// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a virtual machine scale set virtual machine.
// Azure REST API version: 2023-03-01. Prior API version in Azure Native 1.x: 2021-03-01.
//
// Other available API versions: 2023-07-01.
type VirtualMachineScaleSetVM struct {
	pulumi.CustomResourceState

	// Specifies additional capabilities enabled or disabled on the virtual machine in the scale set. For instance: whether the virtual machine has the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
	AdditionalCapabilities AdditionalCapabilitiesResponsePtrOutput `pulumi:"additionalCapabilities"`
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Availability sets overview](https://docs.microsoft.com/azure/virtual-machines/availability-set-overview). For more information on Azure planned maintenance, see [Maintenance and updates for Virtual Machines in Azure](https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates). Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
	AvailabilitySet SubResourceResponsePtrOutput `pulumi:"availabilitySet"`
	// Specifies the boot diagnostic settings state. Minimum api-version: 2015-06-15.
	DiagnosticsProfile DiagnosticsProfileResponsePtrOutput `pulumi:"diagnosticsProfile"`
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile HardwareProfileResponsePtrOutput `pulumi:"hardwareProfile"`
	// The identity of the virtual machine, if configured.
	Identity VirtualMachineIdentityResponsePtrOutput `pulumi:"identity"`
	// The virtual machine instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The virtual machine instance view.
	InstanceView VirtualMachineScaleSetVMInstanceViewResponseOutput `pulumi:"instanceView"`
	// Specifies whether the latest model has been applied to the virtual machine.
	LatestModelApplied pulumi.BoolOutput `pulumi:"latestModelApplied"`
	// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies whether the model applied to the virtual machine is the model of the virtual machine scale set or the customized model for the virtual machine.
	ModelDefinitionApplied pulumi.StringOutput `pulumi:"modelDefinitionApplied"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// Specifies the network profile configuration of the virtual machine.
	NetworkProfileConfiguration VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput `pulumi:"networkProfileConfiguration"`
	// Specifies the operating system settings for the virtual machine.
	OsProfile OSProfileResponsePtrOutput `pulumi:"osProfile"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanResponsePtrOutput `pulumi:"plan"`
	// Specifies the protection policy of the virtual machine.
	ProtectionPolicy VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput `pulumi:"protectionPolicy"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The virtual machine child extension resources.
	Resources VirtualMachineExtensionResponseArrayOutput `pulumi:"resources"`
	// Specifies the Security related profile settings for the virtual machine.
	SecurityProfile SecurityProfileResponsePtrOutput `pulumi:"securityProfile"`
	// The virtual machine SKU.
	Sku SkuResponseOutput `pulumi:"sku"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// UserData for the VM, which must be base-64 encoded. Customer should not pass any secrets in here. <br><br>Minimum api-version: 2021-03-01
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// Azure VM unique ID.
	VmId pulumi.StringOutput `pulumi:"vmId"`
	// The virtual machine zones.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewVirtualMachineScaleSetVM registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSetVM(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetVMArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVM, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmScaleSetName == nil {
		return nil, errors.New("invalid value for required argument 'VmScaleSetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230701:VirtualMachineScaleSetVM"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualMachineScaleSetVM
	err := ctx.RegisterResource("azure-native:compute:VirtualMachineScaleSetVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSetVM gets an existing VirtualMachineScaleSetVM resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSetVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetVMState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVM, error) {
	var resource VirtualMachineScaleSetVM
	err := ctx.ReadResource("azure-native:compute:VirtualMachineScaleSetVM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSetVM resources.
type virtualMachineScaleSetVMState struct {
}

type VirtualMachineScaleSetVMState struct {
}

func (VirtualMachineScaleSetVMState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMState)(nil)).Elem()
}

type virtualMachineScaleSetVMArgs struct {
	// Specifies additional capabilities enabled or disabled on the virtual machine in the scale set. For instance: whether the virtual machine has the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
	AdditionalCapabilities *AdditionalCapabilities `pulumi:"additionalCapabilities"`
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Availability sets overview](https://docs.microsoft.com/azure/virtual-machines/availability-set-overview). For more information on Azure planned maintenance, see [Maintenance and updates for Virtual Machines in Azure](https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates). Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
	AvailabilitySet *SubResource `pulumi:"availabilitySet"`
	// Specifies the boot diagnostic settings state. Minimum api-version: 2015-06-15.
	DiagnosticsProfile *DiagnosticsProfile `pulumi:"diagnosticsProfile"`
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile *HardwareProfile `pulumi:"hardwareProfile"`
	// The identity of the virtual machine, if configured.
	Identity *VirtualMachineIdentity `pulumi:"identity"`
	// The instance ID of the virtual machine.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
	LicenseType *string `pulumi:"licenseType"`
	// Resource location
	Location *string `pulumi:"location"`
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
	// Specifies the network profile configuration of the virtual machine.
	NetworkProfileConfiguration *VirtualMachineScaleSetVMNetworkProfileConfiguration `pulumi:"networkProfileConfiguration"`
	// Specifies the operating system settings for the virtual machine.
	OsProfile *OSProfile `pulumi:"osProfile"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan *Plan `pulumi:"plan"`
	// Specifies the protection policy of the virtual machine.
	ProtectionPolicy *VirtualMachineScaleSetVMProtectionPolicy `pulumi:"protectionPolicy"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the Security related profile settings for the virtual machine.
	SecurityProfile *SecurityProfile `pulumi:"securityProfile"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile *StorageProfile `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// UserData for the VM, which must be base-64 encoded. Customer should not pass any secrets in here. <br><br>Minimum api-version: 2021-03-01
	UserData *string `pulumi:"userData"`
	// The name of the VM scale set where the extension should be create or updated.
	VmScaleSetName string `pulumi:"vmScaleSetName"`
}

// The set of arguments for constructing a VirtualMachineScaleSetVM resource.
type VirtualMachineScaleSetVMArgs struct {
	// Specifies additional capabilities enabled or disabled on the virtual machine in the scale set. For instance: whether the virtual machine has the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
	AdditionalCapabilities AdditionalCapabilitiesPtrInput
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Availability sets overview](https://docs.microsoft.com/azure/virtual-machines/availability-set-overview). For more information on Azure planned maintenance, see [Maintenance and updates for Virtual Machines in Azure](https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates). Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
	AvailabilitySet SubResourcePtrInput
	// Specifies the boot diagnostic settings state. Minimum api-version: 2015-06-15.
	DiagnosticsProfile DiagnosticsProfilePtrInput
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile HardwareProfilePtrInput
	// The identity of the virtual machine, if configured.
	Identity VirtualMachineIdentityPtrInput
	// The instance ID of the virtual machine.
	InstanceId pulumi.StringPtrInput
	// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
	LicenseType pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile NetworkProfilePtrInput
	// Specifies the network profile configuration of the virtual machine.
	NetworkProfileConfiguration VirtualMachineScaleSetVMNetworkProfileConfigurationPtrInput
	// Specifies the operating system settings for the virtual machine.
	OsProfile OSProfilePtrInput
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanPtrInput
	// Specifies the protection policy of the virtual machine.
	ProtectionPolicy VirtualMachineScaleSetVMProtectionPolicyPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the Security related profile settings for the virtual machine.
	SecurityProfile SecurityProfilePtrInput
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile StorageProfilePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// UserData for the VM, which must be base-64 encoded. Customer should not pass any secrets in here. <br><br>Minimum api-version: 2021-03-01
	UserData pulumi.StringPtrInput
	// The name of the VM scale set where the extension should be create or updated.
	VmScaleSetName pulumi.StringInput
}

func (VirtualMachineScaleSetVMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMArgs)(nil)).Elem()
}

type VirtualMachineScaleSetVMInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMOutput() VirtualMachineScaleSetVMOutput
	ToVirtualMachineScaleSetVMOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMOutput
}

func (*VirtualMachineScaleSetVM) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVM)(nil)).Elem()
}

func (i *VirtualMachineScaleSetVM) ToVirtualMachineScaleSetVMOutput() VirtualMachineScaleSetVMOutput {
	return i.ToVirtualMachineScaleSetVMOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSetVM) ToVirtualMachineScaleSetVMOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMOutput)
}

type VirtualMachineScaleSetVMOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVM)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMOutput) ToVirtualMachineScaleSetVMOutput() VirtualMachineScaleSetVMOutput {
	return o
}

func (o VirtualMachineScaleSetVMOutput) ToVirtualMachineScaleSetVMOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMOutput {
	return o
}

// Specifies additional capabilities enabled or disabled on the virtual machine in the scale set. For instance: whether the virtual machine has the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
func (o VirtualMachineScaleSetVMOutput) AdditionalCapabilities() AdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) AdditionalCapabilitiesResponsePtrOutput {
		return v.AdditionalCapabilities
	}).(AdditionalCapabilitiesResponsePtrOutput)
}

// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Availability sets overview](https://docs.microsoft.com/azure/virtual-machines/availability-set-overview). For more information on Azure planned maintenance, see [Maintenance and updates for Virtual Machines in Azure](https://docs.microsoft.com/azure/virtual-machines/maintenance-and-updates). Currently, a VM can only be added to availability set at creation time. An existing VM cannot be added to an availability set.
func (o VirtualMachineScaleSetVMOutput) AvailabilitySet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) SubResourceResponsePtrOutput { return v.AvailabilitySet }).(SubResourceResponsePtrOutput)
}

// Specifies the boot diagnostic settings state. Minimum api-version: 2015-06-15.
func (o VirtualMachineScaleSetVMOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) DiagnosticsProfileResponsePtrOutput { return v.DiagnosticsProfile }).(DiagnosticsProfileResponsePtrOutput)
}

// Specifies the hardware settings for the virtual machine.
func (o VirtualMachineScaleSetVMOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) HardwareProfileResponsePtrOutput { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

// The identity of the virtual machine, if configured.
func (o VirtualMachineScaleSetVMOutput) Identity() VirtualMachineIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) VirtualMachineIdentityResponsePtrOutput { return v.Identity }).(VirtualMachineIdentityResponsePtrOutput)
}

// The virtual machine instance ID.
func (o VirtualMachineScaleSetVMOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The virtual machine instance view.
func (o VirtualMachineScaleSetVMOutput) InstanceView() VirtualMachineScaleSetVMInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) VirtualMachineScaleSetVMInstanceViewResponseOutput {
		return v.InstanceView
	}).(VirtualMachineScaleSetVMInstanceViewResponseOutput)
}

// Specifies whether the latest model has been applied to the virtual machine.
func (o VirtualMachineScaleSetVMOutput) LatestModelApplied() pulumi.BoolOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.BoolOutput { return v.LatestModelApplied }).(pulumi.BoolOutput)
}

// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
func (o VirtualMachineScaleSetVMOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Resource location
func (o VirtualMachineScaleSetVMOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Specifies whether the model applied to the virtual machine is the model of the virtual machine scale set or the customized model for the virtual machine.
func (o VirtualMachineScaleSetVMOutput) ModelDefinitionApplied() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.ModelDefinitionApplied }).(pulumi.StringOutput)
}

// Resource name
func (o VirtualMachineScaleSetVMOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the network interfaces of the virtual machine.
func (o VirtualMachineScaleSetVMOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

// Specifies the network profile configuration of the virtual machine.
func (o VirtualMachineScaleSetVMOutput) NetworkProfileConfiguration() VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput {
		return v.NetworkProfileConfiguration
	}).(VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput)
}

// Specifies the operating system settings for the virtual machine.
func (o VirtualMachineScaleSetVMOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) OSProfileResponsePtrOutput { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
func (o VirtualMachineScaleSetVMOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

// Specifies the protection policy of the virtual machine.
func (o VirtualMachineScaleSetVMOutput) ProtectionPolicy() VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput {
		return v.ProtectionPolicy
	}).(VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput)
}

// The provisioning state, which only appears in the response.
func (o VirtualMachineScaleSetVMOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The virtual machine child extension resources.
func (o VirtualMachineScaleSetVMOutput) Resources() VirtualMachineExtensionResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) VirtualMachineExtensionResponseArrayOutput { return v.Resources }).(VirtualMachineExtensionResponseArrayOutput)
}

// Specifies the Security related profile settings for the virtual machine.
func (o VirtualMachineScaleSetVMOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) SecurityProfileResponsePtrOutput { return v.SecurityProfile }).(SecurityProfileResponsePtrOutput)
}

// The virtual machine SKU.
func (o VirtualMachineScaleSetVMOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

// Specifies the storage settings for the virtual machine disks.
func (o VirtualMachineScaleSetVMOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

// Resource tags
func (o VirtualMachineScaleSetVMOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o VirtualMachineScaleSetVMOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// UserData for the VM, which must be base-64 encoded. Customer should not pass any secrets in here. <br><br>Minimum api-version: 2021-03-01
func (o VirtualMachineScaleSetVMOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

// Azure VM unique ID.
func (o VirtualMachineScaleSetVMOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

// The virtual machine zones.
func (o VirtualMachineScaleSetVMOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMOutput{})
}
