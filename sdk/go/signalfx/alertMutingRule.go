// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a SignalFx resource for managing alert muting rules. See [Mute Notifications](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html) for more information.
//
// > **WARNING** SignalFx does not allow the start time of a **currently active** muting rule to be modified. As such, attempting to modify a currently active rule will destroy the existing rule and create a new rule. This may result in the emission of notifications.
type AlertMutingRule struct {
	pulumi.CustomResourceState

	// The description for this muting rule
	Description pulumi.StringOutput `pulumi:"description"`
	// A convenience attribute that associated this muting rule with specific detector ids.
	Detectors          pulumi.StringArrayOutput `pulumi:"detectors"`
	EffectiveStartTime pulumi.IntOutput         `pulumi:"effectiveStartTime"`
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.
	Filters AlertMutingRuleFilterArrayOutput `pulumi:"filters"`
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime pulumi.IntOutput `pulumi:"startTime"`
	// Starting time of an alert muting rule as a Unix time stamp in seconds.
	StopTime pulumi.IntPtrOutput `pulumi:"stopTime"`
}

// NewAlertMutingRule registers a new resource with the given unique name, arguments, and options.
func NewAlertMutingRule(ctx *pulumi.Context,
	name string, args *AlertMutingRuleArgs, opts ...pulumi.ResourceOption) (*AlertMutingRule, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.Filters == nil {
		return nil, errors.New("missing required argument 'Filters'")
	}
	if args == nil || args.StartTime == nil {
		return nil, errors.New("missing required argument 'StartTime'")
	}
	if args == nil {
		args = &AlertMutingRuleArgs{}
	}
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
	// A convenience attribute that associated this muting rule with specific detector ids.
	Detectors          []string `pulumi:"detectors"`
	EffectiveStartTime *int     `pulumi:"effectiveStartTime"`
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.
	Filters []AlertMutingRuleFilter `pulumi:"filters"`
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime *int `pulumi:"startTime"`
	// Starting time of an alert muting rule as a Unix time stamp in seconds.
	StopTime *int `pulumi:"stopTime"`
}

type AlertMutingRuleState struct {
	// The description for this muting rule
	Description pulumi.StringPtrInput
	// A convenience attribute that associated this muting rule with specific detector ids.
	Detectors          pulumi.StringArrayInput
	EffectiveStartTime pulumi.IntPtrInput
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.
	Filters AlertMutingRuleFilterArrayInput
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime pulumi.IntPtrInput
	// Starting time of an alert muting rule as a Unix time stamp in seconds.
	StopTime pulumi.IntPtrInput
}

func (AlertMutingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertMutingRuleState)(nil)).Elem()
}

type alertMutingRuleArgs struct {
	// The description for this muting rule
	Description string `pulumi:"description"`
	// A convenience attribute that associated this muting rule with specific detector ids.
	Detectors []string `pulumi:"detectors"`
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.
	Filters []AlertMutingRuleFilter `pulumi:"filters"`
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime int `pulumi:"startTime"`
	// Starting time of an alert muting rule as a Unix time stamp in seconds.
	StopTime *int `pulumi:"stopTime"`
}

// The set of arguments for constructing a AlertMutingRule resource.
type AlertMutingRuleArgs struct {
	// The description for this muting rule
	Description pulumi.StringInput
	// A convenience attribute that associated this muting rule with specific detector ids.
	Detectors pulumi.StringArrayInput
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.
	Filters AlertMutingRuleFilterArrayInput
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime pulumi.IntInput
	// Starting time of an alert muting rule as a Unix time stamp in seconds.
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

func (AlertMutingRule) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertMutingRule)(nil)).Elem()
}

func (i AlertMutingRule) ToAlertMutingRuleOutput() AlertMutingRuleOutput {
	return i.ToAlertMutingRuleOutputWithContext(context.Background())
}

func (i AlertMutingRule) ToAlertMutingRuleOutputWithContext(ctx context.Context) AlertMutingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertMutingRuleOutput)
}

type AlertMutingRuleOutput struct {
	*pulumi.OutputState
}

func (AlertMutingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertMutingRuleOutput)(nil)).Elem()
}

func (o AlertMutingRuleOutput) ToAlertMutingRuleOutput() AlertMutingRuleOutput {
	return o
}

func (o AlertMutingRuleOutput) ToAlertMutingRuleOutputWithContext(ctx context.Context) AlertMutingRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AlertMutingRuleOutput{})
}
