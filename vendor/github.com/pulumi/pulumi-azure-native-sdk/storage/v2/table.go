// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Properties of the table, including Id, resource name, resource type.
// Azure REST API version: 2022-09-01. Prior API version in Azure Native 1.x: 2021-02-01.
//
// Other available API versions: 2023-01-01.
type Table struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of stored access policies specified on the table.
	SignedIdentifiers TableSignedIdentifierResponseArrayOutput `pulumi:"signedIdentifiers"`
	// Table name under the specified account
	TableName pulumi.StringOutput `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage/v20190601:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220901:Table"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230101:Table"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Table
	err := ctx.RegisterResource("azure-native:storage:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("azure-native:storage:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
}

type TableState struct {
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// List of stored access policies specified on the table.
	SignedIdentifiers []TableSignedIdentifier `pulumi:"signedIdentifiers"`
	// A table name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of only alphanumeric characters and it cannot begin with a numeric character.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// List of stored access policies specified on the table.
	SignedIdentifiers TableSignedIdentifierArrayInput
	// A table name must be unique within a storage account and must be between 3 and 63 characters.The name must comprise of only alphanumeric characters and it cannot begin with a numeric character.
	TableName pulumi.StringPtrInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

func (i *Table) ToOutput(ctx context.Context) pulumix.Output[*Table] {
	return pulumix.Output[*Table]{
		OutputState: i.ToTableOutputWithContext(ctx).OutputState,
	}
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func (o TableOutput) ToOutput(ctx context.Context) pulumix.Output[*Table] {
	return pulumix.Output[*Table]{
		OutputState: o.OutputState,
	}
}

// The name of the resource
func (o TableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of stored access policies specified on the table.
func (o TableOutput) SignedIdentifiers() TableSignedIdentifierResponseArrayOutput {
	return o.ApplyT(func(v *Table) TableSignedIdentifierResponseArrayOutput { return v.SignedIdentifiers }).(TableSignedIdentifierResponseArrayOutput)
}

// Table name under the specified account
func (o TableOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TableOutput{})
}
