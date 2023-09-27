// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Object model for the Azure PowerShell script.
// Azure REST API version: 2020-10-01. Prior API version in Azure Native 1.x: 2020-10-01
type AzurePowerShellScript struct {
	pulumi.CustomResourceState

	// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
	Arguments pulumi.StringPtrOutput `pulumi:"arguments"`
	// Azure PowerShell module version to be used.
	AzPowerShellVersion pulumi.StringOutput `pulumi:"azPowerShellVersion"`
	// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
	CleanupPreference pulumi.StringPtrOutput `pulumi:"cleanupPreference"`
	// Container settings.
	ContainerSettings ContainerConfigurationResponsePtrOutput `pulumi:"containerSettings"`
	// The environment variables to pass over to the script.
	EnvironmentVariables EnvironmentVariableResponseArrayOutput `pulumi:"environmentVariables"`
	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
	ForceUpdateTag pulumi.StringPtrOutput `pulumi:"forceUpdateTag"`
	// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Type of the script.
	// Expected value is 'AzurePowerShell'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The location of the ACI and the storage account for the deployment script.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of script outputs.
	Outputs pulumi.MapOutput `pulumi:"outputs"`
	// Uri for the script. This is the entry point for the external script.
	PrimaryScriptUri pulumi.StringPtrOutput `pulumi:"primaryScriptUri"`
	// State of the script execution. This only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
	RetentionInterval pulumi.StringOutput `pulumi:"retentionInterval"`
	// Script body.
	ScriptContent pulumi.StringPtrOutput `pulumi:"scriptContent"`
	// Contains the results of script execution.
	Status ScriptStatusResponseOutput `pulumi:"status"`
	// Storage Account settings.
	StorageAccountSettings StorageAccountConfigurationResponsePtrOutput `pulumi:"storageAccountSettings"`
	// Supporting files for the external script.
	SupportingScriptUris pulumi.StringArrayOutput `pulumi:"supportingScriptUris"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
	Timeout pulumi.StringPtrOutput `pulumi:"timeout"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAzurePowerShellScript registers a new resource with the given unique name, arguments, and options.
func NewAzurePowerShellScript(ctx *pulumi.Context,
	name string, args *AzurePowerShellScriptArgs, opts ...pulumi.ResourceOption) (*AzurePowerShellScript, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AzPowerShellVersion == nil {
		return nil, errors.New("invalid value for required argument 'AzPowerShellVersion'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RetentionInterval == nil {
		return nil, errors.New("invalid value for required argument 'RetentionInterval'")
	}
	if args.CleanupPreference == nil {
		args.CleanupPreference = pulumi.StringPtr("Always")
	}
	args.Kind = pulumi.String("AzurePowerShell")
	if args.Timeout == nil {
		args.Timeout = pulumi.StringPtr("P1D")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20191001preview:AzurePowerShellScript"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:AzurePowerShellScript"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AzurePowerShellScript
	err := ctx.RegisterResource("azure-native:resources:AzurePowerShellScript", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzurePowerShellScript gets an existing AzurePowerShellScript resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzurePowerShellScript(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzurePowerShellScriptState, opts ...pulumi.ResourceOption) (*AzurePowerShellScript, error) {
	var resource AzurePowerShellScript
	err := ctx.ReadResource("azure-native:resources:AzurePowerShellScript", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzurePowerShellScript resources.
type azurePowerShellScriptState struct {
}

type AzurePowerShellScriptState struct {
}

func (AzurePowerShellScriptState) ElementType() reflect.Type {
	return reflect.TypeOf((*azurePowerShellScriptState)(nil)).Elem()
}

type azurePowerShellScriptArgs struct {
	// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
	Arguments *string `pulumi:"arguments"`
	// Azure PowerShell module version to be used.
	AzPowerShellVersion string `pulumi:"azPowerShellVersion"`
	// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
	CleanupPreference *string `pulumi:"cleanupPreference"`
	// Container settings.
	ContainerSettings *ContainerConfiguration `pulumi:"containerSettings"`
	// The environment variables to pass over to the script.
	EnvironmentVariables []EnvironmentVariable `pulumi:"environmentVariables"`
	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Type of the script.
	// Expected value is 'AzurePowerShell'.
	Kind string `pulumi:"kind"`
	// The location of the ACI and the storage account for the deployment script.
	Location *string `pulumi:"location"`
	// Uri for the script. This is the entry point for the external script.
	PrimaryScriptUri *string `pulumi:"primaryScriptUri"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
	RetentionInterval string `pulumi:"retentionInterval"`
	// Script body.
	ScriptContent *string `pulumi:"scriptContent"`
	// Name of the deployment script.
	ScriptName *string `pulumi:"scriptName"`
	// Storage Account settings.
	StorageAccountSettings *StorageAccountConfiguration `pulumi:"storageAccountSettings"`
	// Supporting files for the external script.
	SupportingScriptUris []string `pulumi:"supportingScriptUris"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
	Timeout *string `pulumi:"timeout"`
}

// The set of arguments for constructing a AzurePowerShellScript resource.
type AzurePowerShellScriptArgs struct {
	// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
	Arguments pulumi.StringPtrInput
	// Azure PowerShell module version to be used.
	AzPowerShellVersion pulumi.StringInput
	// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
	CleanupPreference pulumi.StringPtrInput
	// Container settings.
	ContainerSettings ContainerConfigurationPtrInput
	// The environment variables to pass over to the script.
	EnvironmentVariables EnvironmentVariableArrayInput
	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
	ForceUpdateTag pulumi.StringPtrInput
	// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
	Identity ManagedServiceIdentityPtrInput
	// Type of the script.
	// Expected value is 'AzurePowerShell'.
	Kind pulumi.StringInput
	// The location of the ACI and the storage account for the deployment script.
	Location pulumi.StringPtrInput
	// Uri for the script. This is the entry point for the external script.
	PrimaryScriptUri pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
	RetentionInterval pulumi.StringInput
	// Script body.
	ScriptContent pulumi.StringPtrInput
	// Name of the deployment script.
	ScriptName pulumi.StringPtrInput
	// Storage Account settings.
	StorageAccountSettings StorageAccountConfigurationPtrInput
	// Supporting files for the external script.
	SupportingScriptUris pulumi.StringArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
	Timeout pulumi.StringPtrInput
}

func (AzurePowerShellScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azurePowerShellScriptArgs)(nil)).Elem()
}

type AzurePowerShellScriptInput interface {
	pulumi.Input

	ToAzurePowerShellScriptOutput() AzurePowerShellScriptOutput
	ToAzurePowerShellScriptOutputWithContext(ctx context.Context) AzurePowerShellScriptOutput
}

func (*AzurePowerShellScript) ElementType() reflect.Type {
	return reflect.TypeOf((**AzurePowerShellScript)(nil)).Elem()
}

func (i *AzurePowerShellScript) ToAzurePowerShellScriptOutput() AzurePowerShellScriptOutput {
	return i.ToAzurePowerShellScriptOutputWithContext(context.Background())
}

func (i *AzurePowerShellScript) ToAzurePowerShellScriptOutputWithContext(ctx context.Context) AzurePowerShellScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzurePowerShellScriptOutput)
}

func (i *AzurePowerShellScript) ToOutput(ctx context.Context) pulumix.Output[*AzurePowerShellScript] {
	return pulumix.Output[*AzurePowerShellScript]{
		OutputState: i.ToAzurePowerShellScriptOutputWithContext(ctx).OutputState,
	}
}

type AzurePowerShellScriptOutput struct{ *pulumi.OutputState }

func (AzurePowerShellScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzurePowerShellScript)(nil)).Elem()
}

func (o AzurePowerShellScriptOutput) ToAzurePowerShellScriptOutput() AzurePowerShellScriptOutput {
	return o
}

func (o AzurePowerShellScriptOutput) ToAzurePowerShellScriptOutputWithContext(ctx context.Context) AzurePowerShellScriptOutput {
	return o
}

func (o AzurePowerShellScriptOutput) ToOutput(ctx context.Context) pulumix.Output[*AzurePowerShellScript] {
	return pulumix.Output[*AzurePowerShellScript]{
		OutputState: o.OutputState,
	}
}

// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
func (o AzurePowerShellScriptOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.Arguments }).(pulumi.StringPtrOutput)
}

// Azure PowerShell module version to be used.
func (o AzurePowerShellScriptOutput) AzPowerShellVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.AzPowerShellVersion }).(pulumi.StringOutput)
}

// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
func (o AzurePowerShellScriptOutput) CleanupPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.CleanupPreference }).(pulumi.StringPtrOutput)
}

// Container settings.
func (o AzurePowerShellScriptOutput) ContainerSettings() ContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) ContainerConfigurationResponsePtrOutput { return v.ContainerSettings }).(ContainerConfigurationResponsePtrOutput)
}

// The environment variables to pass over to the script.
func (o AzurePowerShellScriptOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) EnvironmentVariableResponseArrayOutput { return v.EnvironmentVariables }).(EnvironmentVariableResponseArrayOutput)
}

// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
func (o AzurePowerShellScriptOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
func (o AzurePowerShellScriptOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Type of the script.
// Expected value is 'AzurePowerShell'.
func (o AzurePowerShellScriptOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The location of the ACI and the storage account for the deployment script.
func (o AzurePowerShellScriptOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of this resource.
func (o AzurePowerShellScriptOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of script outputs.
func (o AzurePowerShellScriptOutput) Outputs() pulumi.MapOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.MapOutput { return v.Outputs }).(pulumi.MapOutput)
}

// Uri for the script. This is the entry point for the external script.
func (o AzurePowerShellScriptOutput) PrimaryScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.PrimaryScriptUri }).(pulumi.StringPtrOutput)
}

// State of the script execution. This only appears in the response.
func (o AzurePowerShellScriptOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
func (o AzurePowerShellScriptOutput) RetentionInterval() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.RetentionInterval }).(pulumi.StringOutput)
}

// Script body.
func (o AzurePowerShellScriptOutput) ScriptContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.ScriptContent }).(pulumi.StringPtrOutput)
}

// Contains the results of script execution.
func (o AzurePowerShellScriptOutput) Status() ScriptStatusResponseOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) ScriptStatusResponseOutput { return v.Status }).(ScriptStatusResponseOutput)
}

// Storage Account settings.
func (o AzurePowerShellScriptOutput) StorageAccountSettings() StorageAccountConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) StorageAccountConfigurationResponsePtrOutput {
		return v.StorageAccountSettings
	}).(StorageAccountConfigurationResponsePtrOutput)
}

// Supporting files for the external script.
func (o AzurePowerShellScriptOutput) SupportingScriptUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringArrayOutput { return v.SupportingScriptUris }).(pulumi.StringArrayOutput)
}

// The system metadata related to this resource.
func (o AzurePowerShellScriptOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AzurePowerShellScriptOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
func (o AzurePowerShellScriptOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringPtrOutput { return v.Timeout }).(pulumi.StringPtrOutput)
}

// Type of this resource.
func (o AzurePowerShellScriptOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzurePowerShellScript) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzurePowerShellScriptOutput{})
}
