// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Splunk Observability Cloud resource for managing alert muting rules. See [Mute Notifications](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html) for more information.
//
// Splunk Observability Cloud currently allows linking an alert muting rule with only one detector ID. Specifying multiple detector IDs makes the muting rule obsolete.
//
// > **WARNING** Splunk Observability Cloud does not allow the start time of a **currently active** muting rule to be modified. Attempting to modify a currently active rule destroys the existing rule and creates a new rule. This might result in the emission of notifications.
//
// ## Example
//
// ## Arguments
//
// * `description` - (Required) The description for this muting rule
// * `startTime` - (Required) Starting time of an alert muting rule as a Unit time stamp in seconds.
// * `stopTime` - (Optional) Stop time of an alert muting rule as a Unix time stamp in seconds.
// * `detectors` - (Optional) A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
// * `filter` - (Optional) Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
//   - `property` - (Required) The property to filter.
//   - `propertyValue` - (Required) The property value to filter.
//   - `negated` - (Optional) Determines if this is a "not" filter. Defaults to `false`.
//
// ## Attributes
//
// In a addition to all arguments above, the following attributes are exported:
//
// * `id` - The ID of the alert muting rule.
// * `effectiveStartTime`
type AlertMutingRule struct {
	pulumi.CustomResourceState

	// description of the rule
	Description pulumi.StringOutput `pulumi:"description"`
	// detectors to which this muting rule applies
	Detectors          pulumi.StringArrayOutput `pulumi:"detectors"`
	EffectiveStartTime pulumi.IntOutput         `pulumi:"effectiveStartTime"`
	// list of alert muting filters for this rule
	Filters AlertMutingRuleFilterArrayOutput `pulumi:"filters"`
	// starting time of an alert muting rule as a Unix timestamp, in seconds
	StartTime pulumi.IntOutput `pulumi:"startTime"`
	// stop time of an alert muting rule as a Unix timestamp, in seconds
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
	// description of the rule
	Description *string `pulumi:"description"`
	// detectors to which this muting rule applies
	Detectors          []string `pulumi:"detectors"`
	EffectiveStartTime *int     `pulumi:"effectiveStartTime"`
	// list of alert muting filters for this rule
	Filters []AlertMutingRuleFilter `pulumi:"filters"`
	// starting time of an alert muting rule as a Unix timestamp, in seconds
	StartTime *int `pulumi:"startTime"`
	// stop time of an alert muting rule as a Unix timestamp, in seconds
	StopTime *int `pulumi:"stopTime"`
}

type AlertMutingRuleState struct {
	// description of the rule
	Description pulumi.StringPtrInput
	// detectors to which this muting rule applies
	Detectors          pulumi.StringArrayInput
	EffectiveStartTime pulumi.IntPtrInput
	// list of alert muting filters for this rule
	Filters AlertMutingRuleFilterArrayInput
	// starting time of an alert muting rule as a Unix timestamp, in seconds
	StartTime pulumi.IntPtrInput
	// stop time of an alert muting rule as a Unix timestamp, in seconds
	StopTime pulumi.IntPtrInput
}

func (AlertMutingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertMutingRuleState)(nil)).Elem()
}

type alertMutingRuleArgs struct {
	// description of the rule
	Description string `pulumi:"description"`
	// detectors to which this muting rule applies
	Detectors []string `pulumi:"detectors"`
	// list of alert muting filters for this rule
	Filters []AlertMutingRuleFilter `pulumi:"filters"`
	// starting time of an alert muting rule as a Unix timestamp, in seconds
	StartTime int `pulumi:"startTime"`
	// stop time of an alert muting rule as a Unix timestamp, in seconds
	StopTime *int `pulumi:"stopTime"`
}

// The set of arguments for constructing a AlertMutingRule resource.
type AlertMutingRuleArgs struct {
	// description of the rule
	Description pulumi.StringInput
	// detectors to which this muting rule applies
	Detectors pulumi.StringArrayInput
	// list of alert muting filters for this rule
	Filters AlertMutingRuleFilterArrayInput
	// starting time of an alert muting rule as a Unix timestamp, in seconds
	StartTime pulumi.IntInput
	// stop time of an alert muting rule as a Unix timestamp, in seconds
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

// description of the rule
func (o AlertMutingRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertMutingRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// detectors to which this muting rule applies
func (o AlertMutingRuleOutput) Detectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AlertMutingRule) pulumi.StringArrayOutput { return v.Detectors }).(pulumi.StringArrayOutput)
}

func (o AlertMutingRuleOutput) EffectiveStartTime() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertMutingRule) pulumi.IntOutput { return v.EffectiveStartTime }).(pulumi.IntOutput)
}

// list of alert muting filters for this rule
func (o AlertMutingRuleOutput) Filters() AlertMutingRuleFilterArrayOutput {
	return o.ApplyT(func(v *AlertMutingRule) AlertMutingRuleFilterArrayOutput { return v.Filters }).(AlertMutingRuleFilterArrayOutput)
}

// starting time of an alert muting rule as a Unix timestamp, in seconds
func (o AlertMutingRuleOutput) StartTime() pulumi.IntOutput {
	return o.ApplyT(func(v *AlertMutingRule) pulumi.IntOutput { return v.StartTime }).(pulumi.IntOutput)
}

// stop time of an alert muting rule as a Unix timestamp, in seconds
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
