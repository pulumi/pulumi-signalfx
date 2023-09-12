// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an Observability Cloud resource for managing alert muting rules. See [Mute Notifications](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html) for more information.
//
// > **WARNING** Observability Cloud does not allow the start time of a **currently active** muting rule to be modified. As such, attempting to modify a currently active rule will destroy the existing rule and create a new rule. This may result in the emission of notifications.
//
// > **WARNING** Observability Cloud currently allows linking alert muting rule with only one detector ID. Specifying multiple detector IDs will make the muting rule obsolete.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := signalfx.NewAlertMutingRule(ctx, "roolMooterOne", &signalfx.AlertMutingRuleArgs{
//				Description: pulumi.String("mooted it NEW"),
//				StartTime:   pulumi.Int(1573063243),
//				StopTime:    pulumi.Int(0),
//				Detectors: pulumi.StringArray{
//					signalfx_detector.Some_detector_id,
//				},
//				Filters: signalfx.AlertMutingRuleFilterArray{
//					&signalfx.AlertMutingRuleFilterArgs{
//						Property:      pulumi.String("foo"),
//						PropertyValue: pulumi.String("bar"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AlertMutingRule struct {
	pulumi.CustomResourceState

	// The description for this muting rule
	Description pulumi.StringOutput `pulumi:"description"`
	// A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
	Detectors          pulumi.StringArrayOutput `pulumi:"detectors"`
	EffectiveStartTime pulumi.IntOutput         `pulumi:"effectiveStartTime"`
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
	Filters AlertMutingRuleFilterArrayOutput `pulumi:"filters"`
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime pulumi.IntOutput `pulumi:"startTime"`
	// Stop time of an alert muting rule as a Unix time stamp in seconds.
	StopTime pulumi.IntPtrOutput `pulumi:"stopTime"`
}

// NewAlertMutingRule registers a new resource with the given unique name, arguments, and options.
func NewAlertMutingRule(ctx *pulumi.Context,
	name string, args *AlertMutingRuleArgs, opts ...pulumi.ResourceOption) (*AlertMutingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.StartTime == nil {
		return nil, errors.New("invalid value for required argument 'StartTime'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AlertMutingRule
	err := ctx.RegisterResource("signalfx:index/alertMutingRule:AlertMutingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertMutingRule gets an existing AlertMutingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertMutingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertMutingRuleState, opts ...pulumi.ResourceOption) (*AlertMutingRule, error) {
	var resource AlertMutingRule
	err := ctx.ReadResource("signalfx:index/alertMutingRule:AlertMutingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertMutingRule resources.
type alertMutingRuleState struct {
	// The description for this muting rule
	Description *string `pulumi:"description"`
	// A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
	Detectors          []string `pulumi:"detectors"`
	EffectiveStartTime *int     `pulumi:"effectiveStartTime"`
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
	Filters []AlertMutingRuleFilter `pulumi:"filters"`
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime *int `pulumi:"startTime"`
	// Stop time of an alert muting rule as a Unix time stamp in seconds.
	StopTime *int `pulumi:"stopTime"`
}

type AlertMutingRuleState struct {
	// The description for this muting rule
	Description pulumi.StringPtrInput
	// A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
	Detectors          pulumi.StringArrayInput
	EffectiveStartTime pulumi.IntPtrInput
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
	Filters AlertMutingRuleFilterArrayInput
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime pulumi.IntPtrInput
	// Stop time of an alert muting rule as a Unix time stamp in seconds.
	StopTime pulumi.IntPtrInput
}

func (AlertMutingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertMutingRuleState)(nil)).Elem()
}

type alertMutingRuleArgs struct {
	// The description for this muting rule
	Description string `pulumi:"description"`
	// A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
	Detectors []string `pulumi:"detectors"`
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
	Filters []AlertMutingRuleFilter `pulumi:"filters"`
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime int `pulumi:"startTime"`
	// Stop time of an alert muting rule as a Unix time stamp in seconds.
	StopTime *int `pulumi:"stopTime"`
}

// The set of arguments for constructing a AlertMutingRule resource.
type AlertMutingRuleArgs struct {
	// The description for this muting rule
	Description pulumi.StringInput
	// A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
	Detectors pulumi.StringArrayInput
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
	Filters AlertMutingRuleFilterArrayInput
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime pulumi.IntInput
	// Stop time of an alert muting rule as a Unix time stamp in seconds.
	StopTime pulumi.IntPtrInput
}

func (AlertMutingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertMutingRuleArgs)(nil)).Elem()
}

type AlertMutingRuleInput interface {
	pulumi.Input

	ToAlertMutingRuleOutput() AlertMutingRuleOutput
	ToAlertMutingRuleOutputWithContext(ctx context.Context) AlertMutingRuleOutput
}

func (*AlertMutingRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertMutingRule)(nil)).Elem()
}

func (i *AlertMutingRule) ToAlertMutingRuleOutput() AlertMutingRuleOutput {
	return i.ToAlertMutingRuleOutputWithContext(context.Background())
}

func (i *AlertMutingRule) ToAlertMutingRuleOutputWithContext(ctx context.Context) AlertMutingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertMutingRuleOutput)
}

func (i *AlertMutingRule) ToOutput(ctx context.Context) pulumix.Output[*AlertMutingRule] {
	return pulumix.Output[*AlertMutingRule]{
		OutputState: i.ToAlertMutingRuleOutputWithContext(ctx).OutputState,
	}
}

// AlertMutingRuleArrayInput is an input type that accepts AlertMutingRuleArray and AlertMutingRuleArrayOutput values.
// You can construct a concrete instance of `AlertMutingRuleArrayInput` via:
//
//	AlertMutingRuleArray{ AlertMutingRuleArgs{...} }
type AlertMutingRuleArrayInput interface {
	pulumi.Input

	ToAlertMutingRuleArrayOutput() AlertMutingRuleArrayOutput
	ToAlertMutingRuleArrayOutputWithContext(context.Context) AlertMutingRuleArrayOutput
}

type AlertMutingRuleArray []AlertMutingRuleInput

func (AlertMutingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertMutingRule)(nil)).Elem()
}

func (i AlertMutingRuleArray) ToAlertMutingRuleArrayOutput() AlertMutingRuleArrayOutput {
	return i.ToAlertMutingRuleArrayOutputWithContext(context.Background())
}

func (i AlertMutingRuleArray) ToAlertMutingRuleArrayOutputWithContext(ctx context.Context) AlertMutingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertMutingRuleArrayOutput)
}

func (i AlertMutingRuleArray) ToOutput(ctx context.Context) pulumix.Output[[]*AlertMutingRule] {
	return pulumix.Output[[]*AlertMutingRule]{
		OutputState: i.ToAlertMutingRuleArrayOutputWithContext(ctx).OutputState,
	}
}

// AlertMutingRuleMapInput is an input type that accepts AlertMutingRuleMap and AlertMutingRuleMapOutput values.
// You can construct a concrete instance of `AlertMutingRuleMapInput` via:
//
//	AlertMutingRuleMap{ "key": AlertMutingRuleArgs{...} }
type AlertMutingRuleMapInput interface {
	pulumi.Input

	ToAlertMutingRuleMapOutput() AlertMutingRuleMapOutput
	ToAlertMutingRuleMapOutputWithContext(context.Context) AlertMutingRuleMapOutput
}

type AlertMutingRuleMap map[string]AlertMutingRuleInput

func (AlertMutingRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertMutingRule)(nil)).Elem()
}

func (i AlertMutingRuleMap) ToAlertMutingRuleMapOutput() AlertMutingRuleMapOutput {
	return i.ToAlertMutingRuleMapOutputWithContext(context.Background())
}

func (i AlertMutingRuleMap) ToAlertMutingRuleMapOutputWithContext(ctx context.Context) AlertMutingRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertMutingRuleMapOutput)
}

func (i AlertMutingRuleMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AlertMutingRule] {
	return pulumix.Output[map[string]*AlertMutingRule]{
		OutputState: i.ToAlertMutingRuleMapOutputWithContext(ctx).OutputState,
	}
}

type AlertMutingRuleOutput struct{ *pulumi.OutputState }

func (AlertMutingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertMutingRule)(nil)).Elem()
}

func (o AlertMutingRuleOutput) ToAlertMutingRuleOutput() AlertMutingRuleOutput {
	return o
}

func (o AlertMutingRuleOutput) ToAlertMutingRuleOutputWithContext(ctx context.Context) AlertMutingRuleOutput {
	return o
}

func (o AlertMutingRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*AlertMutingRule] {
	return pulumix.Output[*AlertMutingRule]{
		OutputState: o.OutputState,
	}
}

// The description for this muting rule
func (o AlertMutingRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertMutingRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
func (o AlertMutingRuleOutput) Detectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AlertMutingRule) pulumi.StringArrayOutput { return v.Detectors }).(pulumi.StringArrayOutput)
}

func (o AlertMutingRuleOutput) EffectiveStartTime() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertMutingRule) pulumi.IntOutput { return v.EffectiveStartTime }).(pulumi.IntOutput)
}

// Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
func (o AlertMutingRuleOutput) Filters() AlertMutingRuleFilterArrayOutput {
	return o.ApplyT(func(v *AlertMutingRule) AlertMutingRuleFilterArrayOutput { return v.Filters }).(AlertMutingRuleFilterArrayOutput)
}

// Starting time of an alert muting rule as a Unit time stamp in seconds.
func (o AlertMutingRuleOutput) StartTime() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertMutingRule) pulumi.IntOutput { return v.StartTime }).(pulumi.IntOutput)
}

// Stop time of an alert muting rule as a Unix time stamp in seconds.
func (o AlertMutingRuleOutput) StopTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AlertMutingRule) pulumi.IntPtrOutput { return v.StopTime }).(pulumi.IntPtrOutput)
}

type AlertMutingRuleArrayOutput struct{ *pulumi.OutputState }

func (AlertMutingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AlertMutingRule)(nil)).Elem()
}

func (o AlertMutingRuleArrayOutput) ToAlertMutingRuleArrayOutput() AlertMutingRuleArrayOutput {
	return o
}

func (o AlertMutingRuleArrayOutput) ToAlertMutingRuleArrayOutputWithContext(ctx context.Context) AlertMutingRuleArrayOutput {
	return o
}

func (o AlertMutingRuleArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AlertMutingRule] {
	return pulumix.Output[[]*AlertMutingRule]{
		OutputState: o.OutputState,
	}
}

func (o AlertMutingRuleArrayOutput) Index(i pulumi.IntInput) AlertMutingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AlertMutingRule {
		return vs[0].([]*AlertMutingRule)[vs[1].(int)]
	}).(AlertMutingRuleOutput)
}

type AlertMutingRuleMapOutput struct{ *pulumi.OutputState }

func (AlertMutingRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AlertMutingRule)(nil)).Elem()
}

func (o AlertMutingRuleMapOutput) ToAlertMutingRuleMapOutput() AlertMutingRuleMapOutput {
	return o
}

func (o AlertMutingRuleMapOutput) ToAlertMutingRuleMapOutputWithContext(ctx context.Context) AlertMutingRuleMapOutput {
	return o
}

func (o AlertMutingRuleMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AlertMutingRule] {
	return pulumix.Output[map[string]*AlertMutingRule]{
		OutputState: o.OutputState,
	}
}

func (o AlertMutingRuleMapOutput) MapIndex(k pulumi.StringInput) AlertMutingRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AlertMutingRule {
		return vs[0].(map[string]*AlertMutingRule)[vs[1].(string)]
	}).(AlertMutingRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertMutingRuleInput)(nil)).Elem(), &AlertMutingRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertMutingRuleArrayInput)(nil)).Elem(), AlertMutingRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertMutingRuleMapInput)(nil)).Elem(), AlertMutingRuleMap{})
	pulumi.RegisterOutputType(AlertMutingRuleOutput{})
	pulumi.RegisterOutputType(AlertMutingRuleArrayOutput{})
	pulumi.RegisterOutputType(AlertMutingRuleMapOutput{})
}
