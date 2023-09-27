// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The operation that retrieves information about the capacity reservation.
// Azure REST API version: 2023-03-01.
func LookupCapacityReservation(ctx *pulumi.Context, args *LookupCapacityReservationArgs, opts ...pulumi.InvokeOption) (*LookupCapacityReservationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCapacityReservationResult
	err := ctx.Invoke("azure-native:compute:getCapacityReservation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCapacityReservationArgs struct {
	// The name of the capacity reservation group.
	CapacityReservationGroupName string `pulumi:"capacityReservationGroupName"`
	// The name of the capacity reservation.
	CapacityReservationName string `pulumi:"capacityReservationName"`
	// The expand expression to apply on the operation. 'InstanceView' retrieves a snapshot of the runtime properties of the capacity reservation that is managed by the platform and can change outside of control plane operations.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Specifies information about the capacity reservation.
type LookupCapacityReservationResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// The Capacity reservation instance view.
	InstanceView CapacityReservationInstanceViewResponse `pulumi:"instanceView"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Specifies the value of fault domain count that Capacity Reservation supports for requested VM size. **Note:** The fault domain count specified for a resource (like virtual machines scale set) must be less than or equal to this value if it deploys using capacity reservation. Minimum api-version: 2022-08-01.
	PlatformFaultDomainCount int `pulumi:"platformFaultDomainCount"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// The date time when the capacity reservation was last updated.
	ProvisioningTime string `pulumi:"provisioningTime"`
	// A unique id generated and assigned to the capacity reservation by the platform which does not change throughout the lifetime of the resource.
	ReservationId string `pulumi:"reservationId"`
	// SKU of the resource for which capacity needs be reserved. The SKU name and capacity is required to be set. Currently VM Skus with the capability called 'CapacityReservationSupported' set to true are supported. Refer to List Microsoft.Compute SKUs in a region (https://docs.microsoft.com/rest/api/compute/resourceskus/list) for supported values.
	Sku SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Specifies the time at which the Capacity Reservation resource was created. Minimum api-version: 2021-11-01.
	TimeCreated string `pulumi:"timeCreated"`
	// Resource type
	Type string `pulumi:"type"`
	// A list of all virtual machine resource ids that are associated with the capacity reservation.
	VirtualMachinesAssociated []SubResourceReadOnlyResponse `pulumi:"virtualMachinesAssociated"`
	// Availability Zone to use for this capacity reservation. The zone has to be single value and also should be part for the list of zones specified during the capacity reservation group creation. The zone can be assigned only during creation. If not provided, the reservation supports only non-zonal deployments. If provided, enforces VM/VMSS using this capacity reservation to be in same zone.
	Zones []string `pulumi:"zones"`
}

func LookupCapacityReservationOutput(ctx *pulumi.Context, args LookupCapacityReservationOutputArgs, opts ...pulumi.InvokeOption) LookupCapacityReservationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapacityReservationResult, error) {
			args := v.(LookupCapacityReservationArgs)
			r, err := LookupCapacityReservation(ctx, &args, opts...)
			var s LookupCapacityReservationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapacityReservationResultOutput)
}

type LookupCapacityReservationOutputArgs struct {
	// The name of the capacity reservation group.
	CapacityReservationGroupName pulumi.StringInput `pulumi:"capacityReservationGroupName"`
	// The name of the capacity reservation.
	CapacityReservationName pulumi.StringInput `pulumi:"capacityReservationName"`
	// The expand expression to apply on the operation. 'InstanceView' retrieves a snapshot of the runtime properties of the capacity reservation that is managed by the platform and can change outside of control plane operations.
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCapacityReservationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationArgs)(nil)).Elem()
}

// Specifies information about the capacity reservation.
type LookupCapacityReservationResultOutput struct{ *pulumi.OutputState }

func (LookupCapacityReservationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityReservationResult)(nil)).Elem()
}

func (o LookupCapacityReservationResultOutput) ToLookupCapacityReservationResultOutput() LookupCapacityReservationResultOutput {
	return o
}

func (o LookupCapacityReservationResultOutput) ToLookupCapacityReservationResultOutputWithContext(ctx context.Context) LookupCapacityReservationResultOutput {
	return o
}

func (o LookupCapacityReservationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupCapacityReservationResult] {
	return pulumix.Output[LookupCapacityReservationResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id
func (o LookupCapacityReservationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Capacity reservation instance view.
func (o LookupCapacityReservationResultOutput) InstanceView() CapacityReservationInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) CapacityReservationInstanceViewResponse { return v.InstanceView }).(CapacityReservationInstanceViewResponseOutput)
}

// Resource location
func (o LookupCapacityReservationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o LookupCapacityReservationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the value of fault domain count that Capacity Reservation supports for requested VM size. **Note:** The fault domain count specified for a resource (like virtual machines scale set) must be less than or equal to this value if it deploys using capacity reservation. Minimum api-version: 2022-08-01.
func (o LookupCapacityReservationResultOutput) PlatformFaultDomainCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) int { return v.PlatformFaultDomainCount }).(pulumi.IntOutput)
}

// The provisioning state, which only appears in the response.
func (o LookupCapacityReservationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The date time when the capacity reservation was last updated.
func (o LookupCapacityReservationResultOutput) ProvisioningTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.ProvisioningTime }).(pulumi.StringOutput)
}

// A unique id generated and assigned to the capacity reservation by the platform which does not change throughout the lifetime of the resource.
func (o LookupCapacityReservationResultOutput) ReservationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.ReservationId }).(pulumi.StringOutput)
}

// SKU of the resource for which capacity needs be reserved. The SKU name and capacity is required to be set. Currently VM Skus with the capability called 'CapacityReservationSupported' set to true are supported. Refer to List Microsoft.Compute SKUs in a region (https://docs.microsoft.com/rest/api/compute/resourceskus/list) for supported values.
func (o LookupCapacityReservationResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// Resource tags
func (o LookupCapacityReservationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the time at which the Capacity Reservation resource was created. Minimum api-version: 2021-11-01.
func (o LookupCapacityReservationResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

// Resource type
func (o LookupCapacityReservationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) string { return v.Type }).(pulumi.StringOutput)
}

// A list of all virtual machine resource ids that are associated with the capacity reservation.
func (o LookupCapacityReservationResultOutput) VirtualMachinesAssociated() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) []SubResourceReadOnlyResponse {
		return v.VirtualMachinesAssociated
	}).(SubResourceReadOnlyResponseArrayOutput)
}

// Availability Zone to use for this capacity reservation. The zone has to be single value and also should be part for the list of zones specified during the capacity reservation group creation. The zone can be assigned only during creation. If not provided, the reservation supports only non-zonal deployments. If provided, enforces VM/VMSS using this capacity reservation to be in same zone.
func (o LookupCapacityReservationResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCapacityReservationResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityReservationResultOutput{})
}
