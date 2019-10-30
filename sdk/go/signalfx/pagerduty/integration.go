// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pagerduty

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// SignalFx PagerDuty integrations
// 
// **Note:** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/pagerduty_integration.html.markdown.
type Integration struct {
	s *pulumi.ResourceState
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOpt) (*Integration, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["apiKey"] = nil
		inputs["enabled"] = nil
		inputs["name"] = nil
	} else {
		inputs["apiKey"] = args.ApiKey
		inputs["enabled"] = args.Enabled
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("signalfx:pagerduty/integration:Integration", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Integration{s: s}, nil
}

// GetIntegration gets an existing Integration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegration(ctx *pulumi.Context,
	name string, id pulumi.ID, state *IntegrationState, opts ...pulumi.ResourceOpt) (*Integration, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["apiKey"] = state.ApiKey
		inputs["enabled"] = state.Enabled
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("signalfx:pagerduty/integration:Integration", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Integration{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Integration) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Integration) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// PagerDuty API key.
func (r *Integration) ApiKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["apiKey"])
}

// Whether the integration is enabled.
func (r *Integration) Enabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enabled"])
}

// Name of the integration.
func (r *Integration) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering Integration resources.
type IntegrationState struct {
	// PagerDuty API key.
	ApiKey interface{}
	// Whether the integration is enabled.
	Enabled interface{}
	// Name of the integration.
	Name interface{}
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// PagerDuty API key.
	ApiKey interface{}
	// Whether the integration is enabled.
	Enabled interface{}
	// Name of the integration.
	Name interface{}
}
