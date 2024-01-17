// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The operation to get the run command.
// Azure REST API version: 2023-03-01.
//
// Other available API versions: 2023-07-01, 2023-09-01.
func LookupVirtualMachineRunCommandByVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineRunCommandByVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineRunCommandByVirtualMachineResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualMachineRunCommandByVirtualMachineResult
	err := ctx.Invoke("azure-native:compute:getVirtualMachineRunCommandByVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualMachineRunCommandByVirtualMachineArgs struct {
	// The expand expression to apply on the operation.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine run command.
	RunCommandName string `pulumi:"runCommandName"`
	// The name of the virtual machine containing the run command.
	VmName string `pulumi:"vmName"`
}

// Describes a Virtual Machine run command.
type LookupVirtualMachineRunCommandByVirtualMachineResult struct {
	// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
	AsyncExecution *bool `pulumi:"asyncExecution"`
	// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	ErrorBlobManagedIdentity *RunCommandManagedIdentityResponse `pulumi:"errorBlobManagedIdentity"`
	// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
	ErrorBlobUri *string `pulumi:"errorBlobUri"`
	// Resource Id
	Id string `pulumi:"id"`
	// The virtual machine run command instance view.
	InstanceView VirtualMachineRunCommandInstanceViewResponse `pulumi:"instanceView"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
	OutputBlobManagedIdentity *RunCommandManagedIdentityResponse `pulumi:"outputBlobManagedIdentity"`
	// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
	OutputBlobUri *string `pulumi:"outputBlobUri"`
	// The parameters used by the script.
	Parameters []RunCommandInputParameterResponse `pulumi:"parameters"`
	// The parameters used by the script.
	ProtectedParameters []RunCommandInputParameterResponse `pulumi:"protectedParameters"`
	// The provisioning state, which only appears in the response. If treatFailureAsDeploymentFailure set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If treatFailureAsDeploymentFailure set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
	ProvisioningState string `pulumi:"provisioningState"`
	// Specifies the user account password on the VM when executing the run command.
	RunAsPassword *string `pulumi:"runAsPassword"`
	// Specifies the user account on the VM when executing the run command.
	RunAsUser *string `pulumi:"runAsUser"`
	// The source of the run command script.
	Source *VirtualMachineRunCommandScriptSourceResponse `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
	// Optional. If set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
	TreatFailureAsDeploymentFailure *bool `pulumi:"treatFailureAsDeploymentFailure"`
	// Resource type
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupVirtualMachineRunCommandByVirtualMachineResult
func (val *LookupVirtualMachineRunCommandByVirtualMachineResult) Defaults() *LookupVirtualMachineRunCommandByVirtualMachineResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AsyncExecution == nil {
		asyncExecution_ := false
		tmp.AsyncExecution = &asyncExecution_
	}
	if tmp.TreatFailureAsDeploymentFailure == nil {
		treatFailureAsDeploymentFailure_ := false
		tmp.TreatFailureAsDeploymentFailure = &treatFailureAsDeploymentFailure_
	}
	return &tmp
}

func LookupVirtualMachineRunCommandByVirtualMachineOutput(ctx *pulumi.Context, args LookupVirtualMachineRunCommandByVirtualMachineOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineRunCommandByVirtualMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineRunCommandByVirtualMachineResult, error) {
			args := v.(LookupVirtualMachineRunCommandByVirtualMachineArgs)
			r, err := LookupVirtualMachineRunCommandByVirtualMachine(ctx, &args, opts...)
			var s LookupVirtualMachineRunCommandByVirtualMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineRunCommandByVirtualMachineResultOutput)
}

type LookupVirtualMachineRunCommandByVirtualMachineOutputArgs struct {
	// The expand expression to apply on the operation.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the virtual machine run command.
	RunCommandName pulumi.StringInput `pulumi:"runCommandName"`
	// The name of the virtual machine containing the run command.
	VmName pulumi.StringInput `pulumi:"vmName"`
}

func (LookupVirtualMachineRunCommandByVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineRunCommandByVirtualMachineArgs)(nil)).Elem()
}

// Describes a Virtual Machine run command.
type LookupVirtualMachineRunCommandByVirtualMachineResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineRunCommandByVirtualMachineResult)(nil)).Elem()
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ToLookupVirtualMachineRunCommandByVirtualMachineResultOutput() LookupVirtualMachineRunCommandByVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ToLookupVirtualMachineRunCommandByVirtualMachineResultOutputWithContext(ctx context.Context) LookupVirtualMachineRunCommandByVirtualMachineResultOutput {
	return o
}

// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) AsyncExecution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *bool { return v.AsyncExecution }).(pulumi.BoolPtrOutput)
}

// User-assigned managed identity that has access to errorBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ErrorBlobManagedIdentity() RunCommandManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *RunCommandManagedIdentityResponse {
		return v.ErrorBlobManagedIdentity
	}).(RunCommandManagedIdentityResponsePtrOutput)
}

// Specifies the Azure storage blob where script error stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer errorBlobManagedIdentity parameter.
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ErrorBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *string { return v.ErrorBlobUri }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

// The virtual machine run command instance view.
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) InstanceView() VirtualMachineRunCommandInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) VirtualMachineRunCommandInstanceViewResponse {
		return v.InstanceView
	}).(VirtualMachineRunCommandInstanceViewResponseOutput)
}

// Resource location
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

// User-assigned managed identity that has access to outputBlobUri storage blob. Use an empty object in case of system-assigned identity. Make sure managed identity has been given access to blob's container with 'Storage Blob Data Contributor' role assignment. In case of user-assigned identity, make sure you add it under VM's identity. For more info on managed identity and Run Command, refer https://aka.ms/ManagedIdentity and https://aka.ms/RunCommandManaged
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) OutputBlobManagedIdentity() RunCommandManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *RunCommandManagedIdentityResponse {
		return v.OutputBlobManagedIdentity
	}).(RunCommandManagedIdentityResponsePtrOutput)
}

// Specifies the Azure storage blob where script output stream will be uploaded. Use a SAS URI with read, append, create, write access OR use managed identity to provide the VM access to the blob. Refer outputBlobManagedIdentity parameter.
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) OutputBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *string { return v.OutputBlobUri }).(pulumi.StringPtrOutput)
}

// The parameters used by the script.
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Parameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) []RunCommandInputParameterResponse {
		return v.Parameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

// The parameters used by the script.
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ProtectedParameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) []RunCommandInputParameterResponse {
		return v.ProtectedParameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

// The provisioning state, which only appears in the response. If treatFailureAsDeploymentFailure set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If treatFailureAsDeploymentFailure set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies the user account password on the VM when executing the run command.
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) RunAsPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *string { return v.RunAsPassword }).(pulumi.StringPtrOutput)
}

// Specifies the user account on the VM when executing the run command.
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) RunAsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *string { return v.RunAsUser }).(pulumi.StringPtrOutput)
}

// The source of the run command script.
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Source() VirtualMachineRunCommandScriptSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *VirtualMachineRunCommandScriptSourceResponse {
		return v.Source
	}).(VirtualMachineRunCommandScriptSourceResponsePtrOutput)
}

// Resource tags
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The timeout in seconds to execute the run command.
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *int { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

// Optional. If set to true, any failure in the script will fail the deployment and ProvisioningState will be marked as Failed. If set to false, ProvisioningState would only reflect whether the run command was run or not by the extensions platform, it would not indicate whether script failed in case of script failures. See instance view of run command in case of script failures to see executionMessage, output, error: https://aka.ms/runcommandmanaged#get-execution-status-and-results
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) TreatFailureAsDeploymentFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) *bool {
		return v.TreatFailureAsDeploymentFailure
	}).(pulumi.BoolPtrOutput)
}

// Resource type
func (o LookupVirtualMachineRunCommandByVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineRunCommandByVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineRunCommandByVirtualMachineResultOutput{})
}
