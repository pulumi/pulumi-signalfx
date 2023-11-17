// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type IntegrationCustomNamespaceSyncRule struct {
	// Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filterAction` and `filterSource` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
	DefaultAction *string `pulumi:"defaultAction"`
	// Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
	FilterAction *string `pulumi:"filterAction"`
	// Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
	FilterSource *string `pulumi:"filterSource"`
	// An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
	Namespace string `pulumi:"namespace"`
}

// IntegrationCustomNamespaceSyncRuleInput is an input type that accepts IntegrationCustomNamespaceSyncRuleArgs and IntegrationCustomNamespaceSyncRuleOutput values.
// You can construct a concrete instance of `IntegrationCustomNamespaceSyncRuleInput` via:
//
//	IntegrationCustomNamespaceSyncRuleArgs{...}
type IntegrationCustomNamespaceSyncRuleInput interface {
	pulumi.Input

	ToIntegrationCustomNamespaceSyncRuleOutput() IntegrationCustomNamespaceSyncRuleOutput
	ToIntegrationCustomNamespaceSyncRuleOutputWithContext(context.Context) IntegrationCustomNamespaceSyncRuleOutput
}

type IntegrationCustomNamespaceSyncRuleArgs struct {
	// Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filterAction` and `filterSource` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
	DefaultAction pulumi.StringPtrInput `pulumi:"defaultAction"`
	// Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
	FilterAction pulumi.StringPtrInput `pulumi:"filterAction"`
	// Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
	FilterSource pulumi.StringPtrInput `pulumi:"filterSource"`
	// An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
	Namespace pulumi.StringInput `pulumi:"namespace"`
}

func (IntegrationCustomNamespaceSyncRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationCustomNamespaceSyncRule)(nil)).Elem()
}

func (i IntegrationCustomNamespaceSyncRuleArgs) ToIntegrationCustomNamespaceSyncRuleOutput() IntegrationCustomNamespaceSyncRuleOutput {
	return i.ToIntegrationCustomNamespaceSyncRuleOutputWithContext(context.Background())
}

func (i IntegrationCustomNamespaceSyncRuleArgs) ToIntegrationCustomNamespaceSyncRuleOutputWithContext(ctx context.Context) IntegrationCustomNamespaceSyncRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationCustomNamespaceSyncRuleOutput)
}

// IntegrationCustomNamespaceSyncRuleArrayInput is an input type that accepts IntegrationCustomNamespaceSyncRuleArray and IntegrationCustomNamespaceSyncRuleArrayOutput values.
// You can construct a concrete instance of `IntegrationCustomNamespaceSyncRuleArrayInput` via:
//
//	IntegrationCustomNamespaceSyncRuleArray{ IntegrationCustomNamespaceSyncRuleArgs{...} }
type IntegrationCustomNamespaceSyncRuleArrayInput interface {
	pulumi.Input

	ToIntegrationCustomNamespaceSyncRuleArrayOutput() IntegrationCustomNamespaceSyncRuleArrayOutput
	ToIntegrationCustomNamespaceSyncRuleArrayOutputWithContext(context.Context) IntegrationCustomNamespaceSyncRuleArrayOutput
}

type IntegrationCustomNamespaceSyncRuleArray []IntegrationCustomNamespaceSyncRuleInput

func (IntegrationCustomNamespaceSyncRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationCustomNamespaceSyncRule)(nil)).Elem()
}

func (i IntegrationCustomNamespaceSyncRuleArray) ToIntegrationCustomNamespaceSyncRuleArrayOutput() IntegrationCustomNamespaceSyncRuleArrayOutput {
	return i.ToIntegrationCustomNamespaceSyncRuleArrayOutputWithContext(context.Background())
}

func (i IntegrationCustomNamespaceSyncRuleArray) ToIntegrationCustomNamespaceSyncRuleArrayOutputWithContext(ctx context.Context) IntegrationCustomNamespaceSyncRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationCustomNamespaceSyncRuleArrayOutput)
}

type IntegrationCustomNamespaceSyncRuleOutput struct{ *pulumi.OutputState }

func (IntegrationCustomNamespaceSyncRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationCustomNamespaceSyncRule)(nil)).Elem()
}

func (o IntegrationCustomNamespaceSyncRuleOutput) ToIntegrationCustomNamespaceSyncRuleOutput() IntegrationCustomNamespaceSyncRuleOutput {
	return o
}

func (o IntegrationCustomNamespaceSyncRuleOutput) ToIntegrationCustomNamespaceSyncRuleOutputWithContext(ctx context.Context) IntegrationCustomNamespaceSyncRuleOutput {
	return o
}

// Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filterAction` and `filterSource` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
func (o IntegrationCustomNamespaceSyncRuleOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationCustomNamespaceSyncRule) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

// Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
func (o IntegrationCustomNamespaceSyncRuleOutput) FilterAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationCustomNamespaceSyncRule) *string { return v.FilterAction }).(pulumi.StringPtrOutput)
}

// Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
func (o IntegrationCustomNamespaceSyncRuleOutput) FilterSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationCustomNamespaceSyncRule) *string { return v.FilterSource }).(pulumi.StringPtrOutput)
}

// An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
func (o IntegrationCustomNamespaceSyncRuleOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationCustomNamespaceSyncRule) string { return v.Namespace }).(pulumi.StringOutput)
}

type IntegrationCustomNamespaceSyncRuleArrayOutput struct{ *pulumi.OutputState }

func (IntegrationCustomNamespaceSyncRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationCustomNamespaceSyncRule)(nil)).Elem()
}

func (o IntegrationCustomNamespaceSyncRuleArrayOutput) ToIntegrationCustomNamespaceSyncRuleArrayOutput() IntegrationCustomNamespaceSyncRuleArrayOutput {
	return o
}

func (o IntegrationCustomNamespaceSyncRuleArrayOutput) ToIntegrationCustomNamespaceSyncRuleArrayOutputWithContext(ctx context.Context) IntegrationCustomNamespaceSyncRuleArrayOutput {
	return o
}

func (o IntegrationCustomNamespaceSyncRuleArrayOutput) Index(i pulumi.IntInput) IntegrationCustomNamespaceSyncRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationCustomNamespaceSyncRule {
		return vs[0].([]IntegrationCustomNamespaceSyncRule)[vs[1].(int)]
	}).(IntegrationCustomNamespaceSyncRuleOutput)
}

type IntegrationMetricStatsToSync struct {
	// AWS metric that you want to pick statistics for
	Metric string `pulumi:"metric"`
	// An AWS namespace having AWS metric that you want to pick statistics for
	Namespace string `pulumi:"namespace"`
	// AWS statistics you want to collect
	Stats []string `pulumi:"stats"`
}

// IntegrationMetricStatsToSyncInput is an input type that accepts IntegrationMetricStatsToSyncArgs and IntegrationMetricStatsToSyncOutput values.
// You can construct a concrete instance of `IntegrationMetricStatsToSyncInput` via:
//
//	IntegrationMetricStatsToSyncArgs{...}
type IntegrationMetricStatsToSyncInput interface {
	pulumi.Input

	ToIntegrationMetricStatsToSyncOutput() IntegrationMetricStatsToSyncOutput
	ToIntegrationMetricStatsToSyncOutputWithContext(context.Context) IntegrationMetricStatsToSyncOutput
}

type IntegrationMetricStatsToSyncArgs struct {
	// AWS metric that you want to pick statistics for
	Metric pulumi.StringInput `pulumi:"metric"`
	// An AWS namespace having AWS metric that you want to pick statistics for
	Namespace pulumi.StringInput `pulumi:"namespace"`
	// AWS statistics you want to collect
	Stats pulumi.StringArrayInput `pulumi:"stats"`
}

func (IntegrationMetricStatsToSyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationMetricStatsToSync)(nil)).Elem()
}

func (i IntegrationMetricStatsToSyncArgs) ToIntegrationMetricStatsToSyncOutput() IntegrationMetricStatsToSyncOutput {
	return i.ToIntegrationMetricStatsToSyncOutputWithContext(context.Background())
}

func (i IntegrationMetricStatsToSyncArgs) ToIntegrationMetricStatsToSyncOutputWithContext(ctx context.Context) IntegrationMetricStatsToSyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMetricStatsToSyncOutput)
}

// IntegrationMetricStatsToSyncArrayInput is an input type that accepts IntegrationMetricStatsToSyncArray and IntegrationMetricStatsToSyncArrayOutput values.
// You can construct a concrete instance of `IntegrationMetricStatsToSyncArrayInput` via:
//
//	IntegrationMetricStatsToSyncArray{ IntegrationMetricStatsToSyncArgs{...} }
type IntegrationMetricStatsToSyncArrayInput interface {
	pulumi.Input

	ToIntegrationMetricStatsToSyncArrayOutput() IntegrationMetricStatsToSyncArrayOutput
	ToIntegrationMetricStatsToSyncArrayOutputWithContext(context.Context) IntegrationMetricStatsToSyncArrayOutput
}

type IntegrationMetricStatsToSyncArray []IntegrationMetricStatsToSyncInput

func (IntegrationMetricStatsToSyncArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationMetricStatsToSync)(nil)).Elem()
}

func (i IntegrationMetricStatsToSyncArray) ToIntegrationMetricStatsToSyncArrayOutput() IntegrationMetricStatsToSyncArrayOutput {
	return i.ToIntegrationMetricStatsToSyncArrayOutputWithContext(context.Background())
}

func (i IntegrationMetricStatsToSyncArray) ToIntegrationMetricStatsToSyncArrayOutputWithContext(ctx context.Context) IntegrationMetricStatsToSyncArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMetricStatsToSyncArrayOutput)
}

type IntegrationMetricStatsToSyncOutput struct{ *pulumi.OutputState }

func (IntegrationMetricStatsToSyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationMetricStatsToSync)(nil)).Elem()
}

func (o IntegrationMetricStatsToSyncOutput) ToIntegrationMetricStatsToSyncOutput() IntegrationMetricStatsToSyncOutput {
	return o
}

func (o IntegrationMetricStatsToSyncOutput) ToIntegrationMetricStatsToSyncOutputWithContext(ctx context.Context) IntegrationMetricStatsToSyncOutput {
	return o
}

// AWS metric that you want to pick statistics for
func (o IntegrationMetricStatsToSyncOutput) Metric() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationMetricStatsToSync) string { return v.Metric }).(pulumi.StringOutput)
}

// An AWS namespace having AWS metric that you want to pick statistics for
func (o IntegrationMetricStatsToSyncOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationMetricStatsToSync) string { return v.Namespace }).(pulumi.StringOutput)
}

// AWS statistics you want to collect
func (o IntegrationMetricStatsToSyncOutput) Stats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IntegrationMetricStatsToSync) []string { return v.Stats }).(pulumi.StringArrayOutput)
}

type IntegrationMetricStatsToSyncArrayOutput struct{ *pulumi.OutputState }

func (IntegrationMetricStatsToSyncArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationMetricStatsToSync)(nil)).Elem()
}

func (o IntegrationMetricStatsToSyncArrayOutput) ToIntegrationMetricStatsToSyncArrayOutput() IntegrationMetricStatsToSyncArrayOutput {
	return o
}

func (o IntegrationMetricStatsToSyncArrayOutput) ToIntegrationMetricStatsToSyncArrayOutputWithContext(ctx context.Context) IntegrationMetricStatsToSyncArrayOutput {
	return o
}

func (o IntegrationMetricStatsToSyncArrayOutput) Index(i pulumi.IntInput) IntegrationMetricStatsToSyncOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationMetricStatsToSync {
		return vs[0].([]IntegrationMetricStatsToSync)[vs[1].(int)]
	}).(IntegrationMetricStatsToSyncOutput)
}

type IntegrationNamespaceSyncRule struct {
	// Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filterAction` and `filterSource` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
	DefaultAction *string `pulumi:"defaultAction"`
	// Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
	FilterAction *string `pulumi:"filterAction"`
	// Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
	FilterSource *string `pulumi:"filterSource"`
	// An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See `services` field description below for additional information.
	Namespace string `pulumi:"namespace"`
}

// IntegrationNamespaceSyncRuleInput is an input type that accepts IntegrationNamespaceSyncRuleArgs and IntegrationNamespaceSyncRuleOutput values.
// You can construct a concrete instance of `IntegrationNamespaceSyncRuleInput` via:
//
//	IntegrationNamespaceSyncRuleArgs{...}
type IntegrationNamespaceSyncRuleInput interface {
	pulumi.Input

	ToIntegrationNamespaceSyncRuleOutput() IntegrationNamespaceSyncRuleOutput
	ToIntegrationNamespaceSyncRuleOutputWithContext(context.Context) IntegrationNamespaceSyncRuleOutput
}

type IntegrationNamespaceSyncRuleArgs struct {
	// Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filterAction` and `filterSource` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
	DefaultAction pulumi.StringPtrInput `pulumi:"defaultAction"`
	// Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
	FilterAction pulumi.StringPtrInput `pulumi:"filterAction"`
	// Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
	FilterSource pulumi.StringPtrInput `pulumi:"filterSource"`
	// An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See `services` field description below for additional information.
	Namespace pulumi.StringInput `pulumi:"namespace"`
}

func (IntegrationNamespaceSyncRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationNamespaceSyncRule)(nil)).Elem()
}

func (i IntegrationNamespaceSyncRuleArgs) ToIntegrationNamespaceSyncRuleOutput() IntegrationNamespaceSyncRuleOutput {
	return i.ToIntegrationNamespaceSyncRuleOutputWithContext(context.Background())
}

func (i IntegrationNamespaceSyncRuleArgs) ToIntegrationNamespaceSyncRuleOutputWithContext(ctx context.Context) IntegrationNamespaceSyncRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationNamespaceSyncRuleOutput)
}

// IntegrationNamespaceSyncRuleArrayInput is an input type that accepts IntegrationNamespaceSyncRuleArray and IntegrationNamespaceSyncRuleArrayOutput values.
// You can construct a concrete instance of `IntegrationNamespaceSyncRuleArrayInput` via:
//
//	IntegrationNamespaceSyncRuleArray{ IntegrationNamespaceSyncRuleArgs{...} }
type IntegrationNamespaceSyncRuleArrayInput interface {
	pulumi.Input

	ToIntegrationNamespaceSyncRuleArrayOutput() IntegrationNamespaceSyncRuleArrayOutput
	ToIntegrationNamespaceSyncRuleArrayOutputWithContext(context.Context) IntegrationNamespaceSyncRuleArrayOutput
}

type IntegrationNamespaceSyncRuleArray []IntegrationNamespaceSyncRuleInput

func (IntegrationNamespaceSyncRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationNamespaceSyncRule)(nil)).Elem()
}

func (i IntegrationNamespaceSyncRuleArray) ToIntegrationNamespaceSyncRuleArrayOutput() IntegrationNamespaceSyncRuleArrayOutput {
	return i.ToIntegrationNamespaceSyncRuleArrayOutputWithContext(context.Background())
}

func (i IntegrationNamespaceSyncRuleArray) ToIntegrationNamespaceSyncRuleArrayOutputWithContext(ctx context.Context) IntegrationNamespaceSyncRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationNamespaceSyncRuleArrayOutput)
}

type IntegrationNamespaceSyncRuleOutput struct{ *pulumi.OutputState }

func (IntegrationNamespaceSyncRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationNamespaceSyncRule)(nil)).Elem()
}

func (o IntegrationNamespaceSyncRuleOutput) ToIntegrationNamespaceSyncRuleOutput() IntegrationNamespaceSyncRuleOutput {
	return o
}

func (o IntegrationNamespaceSyncRuleOutput) ToIntegrationNamespaceSyncRuleOutputWithContext(ctx context.Context) IntegrationNamespaceSyncRuleOutput {
	return o
}

// Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filterAction` and `filterSource` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
func (o IntegrationNamespaceSyncRuleOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationNamespaceSyncRule) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

// Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
func (o IntegrationNamespaceSyncRuleOutput) FilterAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationNamespaceSyncRule) *string { return v.FilterAction }).(pulumi.StringPtrOutput)
}

// Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
func (o IntegrationNamespaceSyncRuleOutput) FilterSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationNamespaceSyncRule) *string { return v.FilterSource }).(pulumi.StringPtrOutput)
}

// An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See `services` field description below for additional information.
func (o IntegrationNamespaceSyncRuleOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationNamespaceSyncRule) string { return v.Namespace }).(pulumi.StringOutput)
}

type IntegrationNamespaceSyncRuleArrayOutput struct{ *pulumi.OutputState }

func (IntegrationNamespaceSyncRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationNamespaceSyncRule)(nil)).Elem()
}

func (o IntegrationNamespaceSyncRuleArrayOutput) ToIntegrationNamespaceSyncRuleArrayOutput() IntegrationNamespaceSyncRuleArrayOutput {
	return o
}

func (o IntegrationNamespaceSyncRuleArrayOutput) ToIntegrationNamespaceSyncRuleArrayOutputWithContext(ctx context.Context) IntegrationNamespaceSyncRuleArrayOutput {
	return o
}

func (o IntegrationNamespaceSyncRuleArrayOutput) Index(i pulumi.IntInput) IntegrationNamespaceSyncRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationNamespaceSyncRule {
		return vs[0].([]IntegrationNamespaceSyncRule)[vs[1].(int)]
	}).(IntegrationNamespaceSyncRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationCustomNamespaceSyncRuleInput)(nil)).Elem(), IntegrationCustomNamespaceSyncRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationCustomNamespaceSyncRuleArrayInput)(nil)).Elem(), IntegrationCustomNamespaceSyncRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMetricStatsToSyncInput)(nil)).Elem(), IntegrationMetricStatsToSyncArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMetricStatsToSyncArrayInput)(nil)).Elem(), IntegrationMetricStatsToSyncArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationNamespaceSyncRuleInput)(nil)).Elem(), IntegrationNamespaceSyncRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationNamespaceSyncRuleArrayInput)(nil)).Elem(), IntegrationNamespaceSyncRuleArray{})
	pulumi.RegisterOutputType(IntegrationCustomNamespaceSyncRuleOutput{})
	pulumi.RegisterOutputType(IntegrationCustomNamespaceSyncRuleArrayOutput{})
	pulumi.RegisterOutputType(IntegrationMetricStatsToSyncOutput{})
	pulumi.RegisterOutputType(IntegrationMetricStatsToSyncArrayOutput{})
	pulumi.RegisterOutputType(IntegrationNamespaceSyncRuleOutput{})
	pulumi.RegisterOutputType(IntegrationNamespaceSyncRuleArrayOutput{})
}
