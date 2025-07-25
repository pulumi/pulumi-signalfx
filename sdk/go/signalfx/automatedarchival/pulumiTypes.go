// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automatedarchival

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type ExemptMetricExemptMetric struct {
	// Timestamp of when the automated archival setting was created
	Created *int `pulumi:"created"`
	// ID of the creator of the automated archival setting
	Creator *string `pulumi:"creator"`
	// Timestamp of when the automated archival setting was last updated
	LastUpdated *int `pulumi:"lastUpdated"`
	// ID of user who last updated the automated archival setting
	LastUpdatedBy *string `pulumi:"lastUpdatedBy"`
	// Name of the metric to be exempted from automated archival
	Name string `pulumi:"name"`
}

// ExemptMetricExemptMetricInput is an input type that accepts ExemptMetricExemptMetricArgs and ExemptMetricExemptMetricOutput values.
// You can construct a concrete instance of `ExemptMetricExemptMetricInput` via:
//
//	ExemptMetricExemptMetricArgs{...}
type ExemptMetricExemptMetricInput interface {
	pulumi.Input

	ToExemptMetricExemptMetricOutput() ExemptMetricExemptMetricOutput
	ToExemptMetricExemptMetricOutputWithContext(context.Context) ExemptMetricExemptMetricOutput
}

type ExemptMetricExemptMetricArgs struct {
	// Timestamp of when the automated archival setting was created
	Created pulumi.IntPtrInput `pulumi:"created"`
	// ID of the creator of the automated archival setting
	Creator pulumi.StringPtrInput `pulumi:"creator"`
	// Timestamp of when the automated archival setting was last updated
	LastUpdated pulumi.IntPtrInput `pulumi:"lastUpdated"`
	// ID of user who last updated the automated archival setting
	LastUpdatedBy pulumi.StringPtrInput `pulumi:"lastUpdatedBy"`
	// Name of the metric to be exempted from automated archival
	Name pulumi.StringInput `pulumi:"name"`
}

func (ExemptMetricExemptMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExemptMetricExemptMetric)(nil)).Elem()
}

func (i ExemptMetricExemptMetricArgs) ToExemptMetricExemptMetricOutput() ExemptMetricExemptMetricOutput {
	return i.ToExemptMetricExemptMetricOutputWithContext(context.Background())
}

func (i ExemptMetricExemptMetricArgs) ToExemptMetricExemptMetricOutputWithContext(ctx context.Context) ExemptMetricExemptMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExemptMetricExemptMetricOutput)
}

// ExemptMetricExemptMetricArrayInput is an input type that accepts ExemptMetricExemptMetricArray and ExemptMetricExemptMetricArrayOutput values.
// You can construct a concrete instance of `ExemptMetricExemptMetricArrayInput` via:
//
//	ExemptMetricExemptMetricArray{ ExemptMetricExemptMetricArgs{...} }
type ExemptMetricExemptMetricArrayInput interface {
	pulumi.Input

	ToExemptMetricExemptMetricArrayOutput() ExemptMetricExemptMetricArrayOutput
	ToExemptMetricExemptMetricArrayOutputWithContext(context.Context) ExemptMetricExemptMetricArrayOutput
}

type ExemptMetricExemptMetricArray []ExemptMetricExemptMetricInput

func (ExemptMetricExemptMetricArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExemptMetricExemptMetric)(nil)).Elem()
}

func (i ExemptMetricExemptMetricArray) ToExemptMetricExemptMetricArrayOutput() ExemptMetricExemptMetricArrayOutput {
	return i.ToExemptMetricExemptMetricArrayOutputWithContext(context.Background())
}

func (i ExemptMetricExemptMetricArray) ToExemptMetricExemptMetricArrayOutputWithContext(ctx context.Context) ExemptMetricExemptMetricArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExemptMetricExemptMetricArrayOutput)
}

type ExemptMetricExemptMetricOutput struct{ *pulumi.OutputState }

func (ExemptMetricExemptMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExemptMetricExemptMetric)(nil)).Elem()
}

func (o ExemptMetricExemptMetricOutput) ToExemptMetricExemptMetricOutput() ExemptMetricExemptMetricOutput {
	return o
}

func (o ExemptMetricExemptMetricOutput) ToExemptMetricExemptMetricOutputWithContext(ctx context.Context) ExemptMetricExemptMetricOutput {
	return o
}

// Timestamp of when the automated archival setting was created
func (o ExemptMetricExemptMetricOutput) Created() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExemptMetricExemptMetric) *int { return v.Created }).(pulumi.IntPtrOutput)
}

// ID of the creator of the automated archival setting
func (o ExemptMetricExemptMetricOutput) Creator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExemptMetricExemptMetric) *string { return v.Creator }).(pulumi.StringPtrOutput)
}

// Timestamp of when the automated archival setting was last updated
func (o ExemptMetricExemptMetricOutput) LastUpdated() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ExemptMetricExemptMetric) *int { return v.LastUpdated }).(pulumi.IntPtrOutput)
}

// ID of user who last updated the automated archival setting
func (o ExemptMetricExemptMetricOutput) LastUpdatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExemptMetricExemptMetric) *string { return v.LastUpdatedBy }).(pulumi.StringPtrOutput)
}

// Name of the metric to be exempted from automated archival
func (o ExemptMetricExemptMetricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ExemptMetricExemptMetric) string { return v.Name }).(pulumi.StringOutput)
}

type ExemptMetricExemptMetricArrayOutput struct{ *pulumi.OutputState }

func (ExemptMetricExemptMetricArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExemptMetricExemptMetric)(nil)).Elem()
}

func (o ExemptMetricExemptMetricArrayOutput) ToExemptMetricExemptMetricArrayOutput() ExemptMetricExemptMetricArrayOutput {
	return o
}

func (o ExemptMetricExemptMetricArrayOutput) ToExemptMetricExemptMetricArrayOutputWithContext(ctx context.Context) ExemptMetricExemptMetricArrayOutput {
	return o
}

func (o ExemptMetricExemptMetricArrayOutput) Index(i pulumi.IntInput) ExemptMetricExemptMetricOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExemptMetricExemptMetric {
		return vs[0].([]ExemptMetricExemptMetric)[vs[1].(int)]
	}).(ExemptMetricExemptMetricOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExemptMetricExemptMetricInput)(nil)).Elem(), ExemptMetricExemptMetricArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExemptMetricExemptMetricArrayInput)(nil)).Elem(), ExemptMetricExemptMetricArray{})
	pulumi.RegisterOutputType(ExemptMetricExemptMetricOutput{})
	pulumi.RegisterOutputType(ExemptMetricExemptMetricArrayOutput{})
}
