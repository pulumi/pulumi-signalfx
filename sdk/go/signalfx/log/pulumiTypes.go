// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ViewColumn struct {
	// Name of the log view.
	Name string `pulumi:"name"`
}

// ViewColumnInput is an input type that accepts ViewColumnArgs and ViewColumnOutput values.
// You can construct a concrete instance of `ViewColumnInput` via:
//
//	ViewColumnArgs{...}
type ViewColumnInput interface {
	pulumi.Input

	ToViewColumnOutput() ViewColumnOutput
	ToViewColumnOutputWithContext(context.Context) ViewColumnOutput
}

type ViewColumnArgs struct {
	// Name of the log view.
	Name pulumi.StringInput `pulumi:"name"`
}

func (ViewColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ViewColumn)(nil)).Elem()
}

func (i ViewColumnArgs) ToViewColumnOutput() ViewColumnOutput {
	return i.ToViewColumnOutputWithContext(context.Background())
}

func (i ViewColumnArgs) ToViewColumnOutputWithContext(ctx context.Context) ViewColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewColumnOutput)
}

// ViewColumnArrayInput is an input type that accepts ViewColumnArray and ViewColumnArrayOutput values.
// You can construct a concrete instance of `ViewColumnArrayInput` via:
//
//	ViewColumnArray{ ViewColumnArgs{...} }
type ViewColumnArrayInput interface {
	pulumi.Input

	ToViewColumnArrayOutput() ViewColumnArrayOutput
	ToViewColumnArrayOutputWithContext(context.Context) ViewColumnArrayOutput
}

type ViewColumnArray []ViewColumnInput

func (ViewColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ViewColumn)(nil)).Elem()
}

func (i ViewColumnArray) ToViewColumnArrayOutput() ViewColumnArrayOutput {
	return i.ToViewColumnArrayOutputWithContext(context.Background())
}

func (i ViewColumnArray) ToViewColumnArrayOutputWithContext(ctx context.Context) ViewColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewColumnArrayOutput)
}

type ViewColumnOutput struct{ *pulumi.OutputState }

func (ViewColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ViewColumn)(nil)).Elem()
}

func (o ViewColumnOutput) ToViewColumnOutput() ViewColumnOutput {
	return o
}

func (o ViewColumnOutput) ToViewColumnOutputWithContext(ctx context.Context) ViewColumnOutput {
	return o
}

// Name of the log view.
func (o ViewColumnOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ViewColumn) string { return v.Name }).(pulumi.StringOutput)
}

type ViewColumnArrayOutput struct{ *pulumi.OutputState }

func (ViewColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ViewColumn)(nil)).Elem()
}

func (o ViewColumnArrayOutput) ToViewColumnArrayOutput() ViewColumnArrayOutput {
	return o
}

func (o ViewColumnArrayOutput) ToViewColumnArrayOutputWithContext(ctx context.Context) ViewColumnArrayOutput {
	return o
}

func (o ViewColumnArrayOutput) Index(i pulumi.IntInput) ViewColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ViewColumn {
		return vs[0].([]ViewColumn)[vs[1].(int)]
	}).(ViewColumnOutput)
}

type ViewSortOption struct {
	// Name of the column
	Descending bool `pulumi:"descending"`
	// Name of the column
	Field string `pulumi:"field"`
}

// ViewSortOptionInput is an input type that accepts ViewSortOptionArgs and ViewSortOptionOutput values.
// You can construct a concrete instance of `ViewSortOptionInput` via:
//
//	ViewSortOptionArgs{...}
type ViewSortOptionInput interface {
	pulumi.Input

	ToViewSortOptionOutput() ViewSortOptionOutput
	ToViewSortOptionOutputWithContext(context.Context) ViewSortOptionOutput
}

type ViewSortOptionArgs struct {
	// Name of the column
	Descending pulumi.BoolInput `pulumi:"descending"`
	// Name of the column
	Field pulumi.StringInput `pulumi:"field"`
}

func (ViewSortOptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ViewSortOption)(nil)).Elem()
}

func (i ViewSortOptionArgs) ToViewSortOptionOutput() ViewSortOptionOutput {
	return i.ToViewSortOptionOutputWithContext(context.Background())
}

func (i ViewSortOptionArgs) ToViewSortOptionOutputWithContext(ctx context.Context) ViewSortOptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewSortOptionOutput)
}

// ViewSortOptionArrayInput is an input type that accepts ViewSortOptionArray and ViewSortOptionArrayOutput values.
// You can construct a concrete instance of `ViewSortOptionArrayInput` via:
//
//	ViewSortOptionArray{ ViewSortOptionArgs{...} }
type ViewSortOptionArrayInput interface {
	pulumi.Input

	ToViewSortOptionArrayOutput() ViewSortOptionArrayOutput
	ToViewSortOptionArrayOutputWithContext(context.Context) ViewSortOptionArrayOutput
}

type ViewSortOptionArray []ViewSortOptionInput

func (ViewSortOptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ViewSortOption)(nil)).Elem()
}

func (i ViewSortOptionArray) ToViewSortOptionArrayOutput() ViewSortOptionArrayOutput {
	return i.ToViewSortOptionArrayOutputWithContext(context.Background())
}

func (i ViewSortOptionArray) ToViewSortOptionArrayOutputWithContext(ctx context.Context) ViewSortOptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewSortOptionArrayOutput)
}

type ViewSortOptionOutput struct{ *pulumi.OutputState }

func (ViewSortOptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ViewSortOption)(nil)).Elem()
}

func (o ViewSortOptionOutput) ToViewSortOptionOutput() ViewSortOptionOutput {
	return o
}

func (o ViewSortOptionOutput) ToViewSortOptionOutputWithContext(ctx context.Context) ViewSortOptionOutput {
	return o
}

// Name of the column
func (o ViewSortOptionOutput) Descending() pulumi.BoolOutput {
	return o.ApplyT(func(v ViewSortOption) bool { return v.Descending }).(pulumi.BoolOutput)
}

// Name of the column
func (o ViewSortOptionOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v ViewSortOption) string { return v.Field }).(pulumi.StringOutput)
}

type ViewSortOptionArrayOutput struct{ *pulumi.OutputState }

func (ViewSortOptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ViewSortOption)(nil)).Elem()
}

func (o ViewSortOptionArrayOutput) ToViewSortOptionArrayOutput() ViewSortOptionArrayOutput {
	return o
}

func (o ViewSortOptionArrayOutput) ToViewSortOptionArrayOutputWithContext(ctx context.Context) ViewSortOptionArrayOutput {
	return o
}

func (o ViewSortOptionArrayOutput) Index(i pulumi.IntInput) ViewSortOptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ViewSortOption {
		return vs[0].([]ViewSortOption)[vs[1].(int)]
	}).(ViewSortOptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ViewColumnInput)(nil)).Elem(), ViewColumnArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ViewColumnArrayInput)(nil)).Elem(), ViewColumnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ViewSortOptionInput)(nil)).Elem(), ViewSortOptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ViewSortOptionArrayInput)(nil)).Elem(), ViewSortOptionArray{})
	pulumi.RegisterOutputType(ViewColumnOutput{})
	pulumi.RegisterOutputType(ViewColumnArrayOutput{})
	pulumi.RegisterOutputType(ViewSortOptionOutput{})
	pulumi.RegisterOutputType(ViewSortOptionArrayOutput{})
}
