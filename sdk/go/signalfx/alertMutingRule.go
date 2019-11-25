// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a SignalFx resource for managing alert muting rules. See [Mute Notifications](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html) for more information.
// 
// > **WARNING** SignalFx does not allow the start time of a **currently active** muting rule to be modified. As such, attempting to modify a currently active rule will destroy the existing rule and create a new rule. This may result in the emission of notifications.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/alert_muting_rule.html.markdown.
type AlertMutingRule struct {
	s *pulumi.ResourceState
}

// NewAlertMutingRule registers a new resource with the given unique name, arguments, and options.
func NewAlertMutingRule(ctx *pulumi.Context,
	name string, args *AlertMutingRuleArgs, opts ...pulumi.ResourceOpt) (*AlertMutingRule, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.Filters == nil {
		return nil, errors.New("missing required argument 'Filters'")
	}
	if args == nil || args.StartTime == nil {
		return nil, errors.New("missing required argument 'StartTime'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["detectors"] = nil
		inputs["filters"] = nil
		inputs["startTime"] = nil
		inputs["stopTime"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["detectors"] = args.Detectors
		inputs["filters"] = args.Filters
		inputs["startTime"] = args.StartTime
		inputs["stopTime"] = args.StopTime
	}
	inputs["effectiveStartTime"] = nil
	s, err := ctx.RegisterResource("signalfx:index/alertMutingRule:AlertMutingRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AlertMutingRule{s: s}, nil
}

// GetAlertMutingRule gets an existing AlertMutingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertMutingRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AlertMutingRuleState, opts ...pulumi.ResourceOpt) (*AlertMutingRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["detectors"] = state.Detectors
		inputs["effectiveStartTime"] = state.EffectiveStartTime
		inputs["filters"] = state.Filters
		inputs["startTime"] = state.StartTime
		inputs["stopTime"] = state.StopTime
	}
	s, err := ctx.ReadResource("signalfx:index/alertMutingRule:AlertMutingRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AlertMutingRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AlertMutingRule) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AlertMutingRule) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The description for this muting rule
func (r *AlertMutingRule) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// A convenience attribute that associated this muting rule with specific detector ids.
func (r *AlertMutingRule) Detectors() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["detectors"])
}

func (r *AlertMutingRule) EffectiveStartTime() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["effectiveStartTime"])
}

// Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.
func (r *AlertMutingRule) Filters() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["filters"])
}

// Starting time of an alert muting rule as a Unit time stamp in seconds.
func (r *AlertMutingRule) StartTime() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["startTime"])
}

// Starting time of an alert muting rule as a Unix time stamp in seconds.
func (r *AlertMutingRule) StopTime() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["stopTime"])
}

// Input properties used for looking up and filtering AlertMutingRule resources.
type AlertMutingRuleState struct {
	// The description for this muting rule
	Description interface{}
	// A convenience attribute that associated this muting rule with specific detector ids.
	Detectors interface{}
	EffectiveStartTime interface{}
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.
	Filters interface{}
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime interface{}
	// Starting time of an alert muting rule as a Unix time stamp in seconds.
	StopTime interface{}
}

// The set of arguments for constructing a AlertMutingRule resource.
type AlertMutingRuleArgs struct {
	// The description for this muting rule
	Description interface{}
	// A convenience attribute that associated this muting rule with specific detector ids.
	Detectors interface{}
	// Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.
	Filters interface{}
	// Starting time of an alert muting rule as a Unit time stamp in seconds.
	StartTime interface{}
	// Starting time of an alert muting rule as a Unix time stamp in seconds.
	StopTime interface{}
}
