// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the specified NSP profile.
// Azure REST API version: 2021-02-01-preview.
func LookupNspProfile(ctx *pulumi.Context, args *LookupNspProfileArgs, opts ...pulumi.InvokeOption) (*LookupNspProfileResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNspProfileResult
	err := ctx.Invoke("azure-native:network:getNspProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNspProfileArgs struct {
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName string `pulumi:"networkSecurityPerimeterName"`
	// The name of the NSP profile.
	ProfileName string `pulumi:"profileName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The network security perimeter profile resource
type LookupNspProfileResult struct {
	// Version number that increases with every update to access rules within the profile.
	AccessRulesVersion string `pulumi:"accessRulesVersion"`
	// Version number that increases with every update to diagnostic settings within the profile.
	DiagnosticSettingsVersion string `pulumi:"diagnosticSettingsVersion"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupNspProfileOutput(ctx *pulumi.Context, args LookupNspProfileOutputArgs, opts ...pulumi.InvokeOption) LookupNspProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNspProfileResult, error) {
			args := v.(LookupNspProfileArgs)
			r, err := LookupNspProfile(ctx, &args, opts...)
			var s LookupNspProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNspProfileResultOutput)
}

type LookupNspProfileOutputArgs struct {
	// The name of the network security perimeter.
	NetworkSecurityPerimeterName pulumi.StringInput `pulumi:"networkSecurityPerimeterName"`
	// The name of the NSP profile.
	ProfileName pulumi.StringInput `pulumi:"profileName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNspProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspProfileArgs)(nil)).Elem()
}

// The network security perimeter profile resource
type LookupNspProfileResultOutput struct{ *pulumi.OutputState }

func (LookupNspProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNspProfileResult)(nil)).Elem()
}

func (o LookupNspProfileResultOutput) ToLookupNspProfileResultOutput() LookupNspProfileResultOutput {
	return o
}

func (o LookupNspProfileResultOutput) ToLookupNspProfileResultOutputWithContext(ctx context.Context) LookupNspProfileResultOutput {
	return o
}

func (o LookupNspProfileResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupNspProfileResult] {
	return pulumix.Output[LookupNspProfileResult]{
		OutputState: o.OutputState,
	}
}

// Version number that increases with every update to access rules within the profile.
func (o LookupNspProfileResultOutput) AccessRulesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspProfileResult) string { return v.AccessRulesVersion }).(pulumi.StringOutput)
}

// Version number that increases with every update to diagnostic settings within the profile.
func (o LookupNspProfileResultOutput) DiagnosticSettingsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspProfileResult) string { return v.DiagnosticSettingsVersion }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupNspProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupNspProfileResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNspProfileResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupNspProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupNspProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNspProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupNspProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNspProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNspProfileResultOutput{})
}
