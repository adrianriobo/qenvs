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

// Describes a Virtual Machine run command.
// Azure REST API version: 2023-03-01. Prior API version in Azure Native 1.x: 2021-03-01.
//
// Other available API versions: 2023-07-01.
type VirtualMachineScaleSetVMRunCommand struct {
	pulumi.CustomResourceState

	// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
	AsyncExecution pulumi.BoolPtrOutput `pulumi:"asyncExecution"`
	// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	ErrorBlobManagedIdentity RunCommandManagedIdentityResponsePtrOutput `pulumi:"errorBlobManagedIdentity"`
	// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
	ErrorBlobUri pulumi.StringPtrOutput `pulumi:"errorBlobUri"`
	// The virtual machine run command instance view.
	InstanceView VirtualMachineRunCommandInstanceViewResponseOutput `pulumi:"instanceView"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	OutputBlobManagedIdentity RunCommandManagedIdentityResponsePtrOutput `pulumi:"outputBlobManagedIdentity"`
	// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
	OutputBlobUri pulumi.StringPtrOutput `pulumi:"outputBlobUri"`
	// The parameters used by the script.
	Parameters RunCommandInputParameterResponseArrayOutput `pulumi:"parameters"`
	// The parameters used by the script.
	ProtectedParameters RunCommandInputParameterResponseArrayOutput `pulumi:"protectedParameters"`
	// The provisioning state, which only appears in the response. If treatFailureAsDeploymentFailure set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If treatFailureAsDeploymentFailure set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Specifies the user account password on the VM when executing the run command.
	RunAsPassword pulumi.StringPtrOutput `pulumi:"runAsPassword"`
	// Specifies the user account on the VM when executing the run command.
	RunAsUser pulumi.StringPtrOutput `pulumi:"runAsUser"`
	// The source of the run command script.
	Source VirtualMachineRunCommandScriptSourceResponsePtrOutput `pulumi:"source"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds pulumi.IntPtrOutput `pulumi:"timeoutInSeconds"`
	// Optional. If set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
	TreatFailureAsDeploymentFailure pulumi.BoolPtrOutput `pulumi:"treatFailureAsDeploymentFailure"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualMachineScaleSetVMRunCommand registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSetVMRunCommand(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetVMRunCommandArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVMRunCommand, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmScaleSetName == nil {
		return nil, errors.New("invalid value for required argument 'VmScaleSetName'")
	}
	if args.AsyncExecution == nil {
		args.AsyncExecution = pulumi.BoolPtr(false)
	}
	if args.TreatFailureAsDeploymentFailure == nil {
		args.TreatFailureAsDeploymentFailure = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220801:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20221101:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230301:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230701:VirtualMachineScaleSetVMRunCommand"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualMachineScaleSetVMRunCommand
	err := ctx.RegisterResource("azure-native:compute:VirtualMachineScaleSetVMRunCommand", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSetVMRunCommand gets an existing VirtualMachineScaleSetVMRunCommand resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSetVMRunCommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetVMRunCommandState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVMRunCommand, error) {
	var resource VirtualMachineScaleSetVMRunCommand
	err := ctx.ReadResource("azure-native:compute:VirtualMachineScaleSetVMRunCommand", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSetVMRunCommand resources.
type virtualMachineScaleSetVMRunCommandState struct {
}

type VirtualMachineScaleSetVMRunCommandState struct {
}

func (VirtualMachineScaleSetVMRunCommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMRunCommandState)(nil)).Elem()
}

type virtualMachineScaleSetVMRunCommandArgs struct {
	// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
	AsyncExecution *bool `pulumi:"asyncExecution"`
	// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	ErrorBlobManagedIdentity *RunCommandManagedIdentity `pulumi:"errorBlobManagedIdentity"`
	// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
	ErrorBlobUri *string `pulumi:"errorBlobUri"`
	// The instance ID of the virtual machine.
	InstanceId string `pulumi:"instanceId"`
	// Resource location
	Location *string `pulumi:"location"`
	// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	OutputBlobManagedIdentity *RunCommandManagedIdentity `pulumi:"outputBlobManagedIdentity"`
	// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
	OutputBlobUri *string `pulumi:"outputBlobUri"`
	// The parameters used by the script.
	Parameters []RunCommandInputParameter `pulumi:"parameters"`
	// The parameters used by the script.
	ProtectedParameters []RunCommandInputParameter `pulumi:"protectedParameters"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the user account password on the VM when executing the run command.
	RunAsPassword *string `pulumi:"runAsPassword"`
	// Specifies the user account on the VM when executing the run command.
	RunAsUser *string `pulumi:"runAsUser"`
	// The name of the virtual machine run command.
	RunCommandName *string `pulumi:"runCommandName"`
	// The source of the run command script.
	Source *VirtualMachineRunCommandScriptSource `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
	// Optional. If set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
	TreatFailureAsDeploymentFailure *bool `pulumi:"treatFailureAsDeploymentFailure"`
	// The name of the VM scale set.
	VmScaleSetName string `pulumi:"vmScaleSetName"`
}

// The set of arguments for constructing a VirtualMachineScaleSetVMRunCommand resource.
type VirtualMachineScaleSetVMRunCommandArgs struct {
	// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
	AsyncExecution pulumi.BoolPtrInput
	// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	ErrorBlobManagedIdentity RunCommandManagedIdentityPtrInput
	// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
	ErrorBlobUri pulumi.StringPtrInput
	// The instance ID of the virtual machine.
	InstanceId pulumi.StringInput
	// Resource location
	Location pulumi.StringPtrInput
	// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	OutputBlobManagedIdentity RunCommandManagedIdentityPtrInput
	// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
	OutputBlobUri pulumi.StringPtrInput
	// The parameters used by the script.
	Parameters RunCommandInputParameterArrayInput
	// The parameters used by the script.
	ProtectedParameters RunCommandInputParameterArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the user account password on the VM when executing the run command.
	RunAsPassword pulumi.StringPtrInput
	// Specifies the user account on the VM when executing the run command.
	RunAsUser pulumi.StringPtrInput
	// The name of the virtual machine run command.
	RunCommandName pulumi.StringPtrInput
	// The source of the run command script.
	Source VirtualMachineRunCommandScriptSourcePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds pulumi.IntPtrInput
	// Optional. If set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
	TreatFailureAsDeploymentFailure pulumi.BoolPtrInput
	// The name of the VM scale set.
	VmScaleSetName pulumi.StringInput
}

func (VirtualMachineScaleSetVMRunCommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMRunCommandArgs)(nil)).Elem()
}

type VirtualMachineScaleSetVMRunCommandInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMRunCommandOutput() VirtualMachineScaleSetVMRunCommandOutput
	ToVirtualMachineScaleSetVMRunCommandOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMRunCommandOutput
}

func (*VirtualMachineScaleSetVMRunCommand) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMRunCommand)(nil)).Elem()
}

func (i *VirtualMachineScaleSetVMRunCommand) ToVirtualMachineScaleSetVMRunCommandOutput() VirtualMachineScaleSetVMRunCommandOutput {
	return i.ToVirtualMachineScaleSetVMRunCommandOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSetVMRunCommand) ToVirtualMachineScaleSetVMRunCommandOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMRunCommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMRunCommandOutput)
}

type VirtualMachineScaleSetVMRunCommandOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMRunCommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMRunCommand)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMRunCommandOutput) ToVirtualMachineScaleSetVMRunCommandOutput() VirtualMachineScaleSetVMRunCommandOutput {
	return o
}

func (o VirtualMachineScaleSetVMRunCommandOutput) ToVirtualMachineScaleSetVMRunCommandOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMRunCommandOutput {
	return o
}

// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
func (o VirtualMachineScaleSetVMRunCommandOutput) AsyncExecution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.BoolPtrOutput { return v.AsyncExecution }).(pulumi.BoolPtrOutput)
}

// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
func (o VirtualMachineScaleSetVMRunCommandOutput) ErrorBlobManagedIdentity() RunCommandManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) RunCommandManagedIdentityResponsePtrOutput {
		return v.ErrorBlobManagedIdentity
	}).(RunCommandManagedIdentityResponsePtrOutput)
}

// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
func (o VirtualMachineScaleSetVMRunCommandOutput) ErrorBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringPtrOutput { return v.ErrorBlobUri }).(pulumi.StringPtrOutput)
}

// The virtual machine run command instance view.
func (o VirtualMachineScaleSetVMRunCommandOutput) InstanceView() VirtualMachineRunCommandInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) VirtualMachineRunCommandInstanceViewResponseOutput {
		return v.InstanceView
	}).(VirtualMachineRunCommandInstanceViewResponseOutput)
}

// Resource location
func (o VirtualMachineScaleSetVMRunCommandOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o VirtualMachineScaleSetVMRunCommandOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
func (o VirtualMachineScaleSetVMRunCommandOutput) OutputBlobManagedIdentity() RunCommandManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) RunCommandManagedIdentityResponsePtrOutput {
		return v.OutputBlobManagedIdentity
	}).(RunCommandManagedIdentityResponsePtrOutput)
}

// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
func (o VirtualMachineScaleSetVMRunCommandOutput) OutputBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringPtrOutput { return v.OutputBlobUri }).(pulumi.StringPtrOutput)
}

// The parameters used by the script.
func (o VirtualMachineScaleSetVMRunCommandOutput) Parameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) RunCommandInputParameterResponseArrayOutput {
		return v.Parameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

// The parameters used by the script.
func (o VirtualMachineScaleSetVMRunCommandOutput) ProtectedParameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) RunCommandInputParameterResponseArrayOutput {
		return v.ProtectedParameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

// The provisioning state, which only appears in the response. If treatFailureAsDeploymentFailure set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If treatFailureAsDeploymentFailure set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
func (o VirtualMachineScaleSetVMRunCommandOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies the user account password on the VM when executing the run command.
func (o VirtualMachineScaleSetVMRunCommandOutput) RunAsPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringPtrOutput { return v.RunAsPassword }).(pulumi.StringPtrOutput)
}

// Specifies the user account on the VM when executing the run command.
func (o VirtualMachineScaleSetVMRunCommandOutput) RunAsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringPtrOutput { return v.RunAsUser }).(pulumi.StringPtrOutput)
}

// The source of the run command script.
func (o VirtualMachineScaleSetVMRunCommandOutput) Source() VirtualMachineRunCommandScriptSourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) VirtualMachineRunCommandScriptSourceResponsePtrOutput {
		return v.Source
	}).(VirtualMachineRunCommandScriptSourceResponsePtrOutput)
}

// Resource tags
func (o VirtualMachineScaleSetVMRunCommandOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The timeout in seconds to execute the run command.
func (o VirtualMachineScaleSetVMRunCommandOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.IntPtrOutput { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

// Optional. If set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
func (o VirtualMachineScaleSetVMRunCommandOutput) TreatFailureAsDeploymentFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.BoolPtrOutput {
		return v.TreatFailureAsDeploymentFailure
	}).(pulumi.BoolPtrOutput)
}

// Resource type
func (o VirtualMachineScaleSetVMRunCommandOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMRunCommandOutput{})
}
