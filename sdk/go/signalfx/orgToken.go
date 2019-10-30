// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manage SignalFx org tokens.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/org_token.html.markdown.
type OrgToken struct {
	s *pulumi.ResourceState
}

// NewOrgToken registers a new resource with the given unique name, arguments, and options.
func NewOrgToken(ctx *pulumi.Context,
	name string, args *OrgTokenArgs, opts ...pulumi.ResourceOpt) (*OrgToken, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["disabled"] = nil
		inputs["dpmLimits"] = nil
		inputs["hostOrUsageLimits"] = nil
		inputs["name"] = nil
		inputs["notifications"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["disabled"] = args.Disabled
		inputs["dpmLimits"] = args.DpmLimits
		inputs["hostOrUsageLimits"] = args.HostOrUsageLimits
		inputs["name"] = args.Name
		inputs["notifications"] = args.Notifications
	}
	s, err := ctx.RegisterResource("signalfx:index/orgToken:OrgToken", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OrgToken{s: s}, nil
}

// GetOrgToken gets an existing OrgToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgToken(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OrgTokenState, opts ...pulumi.ResourceOpt) (*OrgToken, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["disabled"] = state.Disabled
		inputs["dpmLimits"] = state.DpmLimits
		inputs["hostOrUsageLimits"] = state.HostOrUsageLimits
		inputs["name"] = state.Name
		inputs["notifications"] = state.Notifications
	}
	s, err := ctx.ReadResource("signalfx:index/orgToken:OrgToken", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OrgToken{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OrgToken) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OrgToken) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Description of the token.
func (r *OrgToken) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
func (r *OrgToken) Disabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disabled"])
}

// Specify DPM-based limits for this token.
func (r *OrgToken) DpmLimits() *pulumi.Output {
	return r.s.State["dpmLimits"]
}

// Specify Usage-based limits for this token.
func (r *OrgToken) HostOrUsageLimits() *pulumi.Output {
	return r.s.State["hostOrUsageLimits"]
}

// Name of the token.
func (r *OrgToken) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// List of strings specifying where notifications will be sent when an incident occurs. See
// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
func (r *OrgToken) Notifications() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["notifications"])
}

// Input properties used for looking up and filtering OrgToken resources.
type OrgTokenState struct {
	// Description of the token.
	Description interface{}
	// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
	Disabled interface{}
	// Specify DPM-based limits for this token.
	DpmLimits interface{}
	// Specify Usage-based limits for this token.
	HostOrUsageLimits interface{}
	// Name of the token.
	Name interface{}
	// List of strings specifying where notifications will be sent when an incident occurs. See
	// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
	Notifications interface{}
}

// The set of arguments for constructing a OrgToken resource.
type OrgTokenArgs struct {
	// Description of the token.
	Description interface{}
	// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
	Disabled interface{}
	// Specify DPM-based limits for this token.
	DpmLimits interface{}
	// Specify Usage-based limits for this token.
	HostOrUsageLimits interface{}
	// Name of the token.
	Name interface{}
	// List of strings specifying where notifications will be sent when an incident occurs. See
	// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
	Notifications interface{}
}
