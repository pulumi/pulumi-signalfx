// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package signalfx

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manage SignalFx org tokens.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/org_token.html.markdown.
type OrgToken struct {
	pulumi.CustomResourceState

	// Description of the token.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Specify DPM-based limits for this token.
	DpmLimits OrgTokenDpmLimitsPtrOutput `pulumi:"dpmLimits"`
	// Specify Usage-based limits for this token.
	HostOrUsageLimits OrgTokenHostOrUsageLimitsPtrOutput `pulumi:"hostOrUsageLimits"`
	// Name of the token.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of strings specifying where notifications will be sent when an incident occurs. See
	// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
	Notifications pulumi.StringArrayOutput `pulumi:"notifications"`
}

// NewOrgToken registers a new resource with the given unique name, arguments, and options.
func NewOrgToken(ctx *pulumi.Context,
	name string, args *OrgTokenArgs, opts ...pulumi.ResourceOption) (*OrgToken, error) {
	if args == nil {
		args = &OrgTokenArgs{}
	}
	var resource OrgToken
	err := ctx.RegisterResource("signalfx:index/orgToken:OrgToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgToken gets an existing OrgToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgTokenState, opts ...pulumi.ResourceOption) (*OrgToken, error) {
	var resource OrgToken
	err := ctx.ReadResource("signalfx:index/orgToken:OrgToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgToken resources.
type orgTokenState struct {
	// Description of the token.
	Description *string `pulumi:"description"`
	// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
	Disabled *bool `pulumi:"disabled"`
	// Specify DPM-based limits for this token.
	DpmLimits *OrgTokenDpmLimits `pulumi:"dpmLimits"`
	// Specify Usage-based limits for this token.
	HostOrUsageLimits *OrgTokenHostOrUsageLimits `pulumi:"hostOrUsageLimits"`
	// Name of the token.
	Name *string `pulumi:"name"`
	// List of strings specifying where notifications will be sent when an incident occurs. See
	// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
	Notifications []string `pulumi:"notifications"`
}

type OrgTokenState struct {
	// Description of the token.
	Description pulumi.StringPtrInput
	// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
	Disabled pulumi.BoolPtrInput
	// Specify DPM-based limits for this token.
	DpmLimits OrgTokenDpmLimitsPtrInput
	// Specify Usage-based limits for this token.
	HostOrUsageLimits OrgTokenHostOrUsageLimitsPtrInput
	// Name of the token.
	Name pulumi.StringPtrInput
	// List of strings specifying where notifications will be sent when an incident occurs. See
	// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
	Notifications pulumi.StringArrayInput
}

func (OrgTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgTokenState)(nil)).Elem()
}

type orgTokenArgs struct {
	// Description of the token.
	Description *string `pulumi:"description"`
	// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
	Disabled *bool `pulumi:"disabled"`
	// Specify DPM-based limits for this token.
	DpmLimits *OrgTokenDpmLimits `pulumi:"dpmLimits"`
	// Specify Usage-based limits for this token.
	HostOrUsageLimits *OrgTokenHostOrUsageLimits `pulumi:"hostOrUsageLimits"`
	// Name of the token.
	Name *string `pulumi:"name"`
	// List of strings specifying where notifications will be sent when an incident occurs. See
	// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
	Notifications []string `pulumi:"notifications"`
}

// The set of arguments for constructing a OrgToken resource.
type OrgTokenArgs struct {
	// Description of the token.
	Description pulumi.StringPtrInput
	// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
	Disabled pulumi.BoolPtrInput
	// Specify DPM-based limits for this token.
	DpmLimits OrgTokenDpmLimitsPtrInput
	// Specify Usage-based limits for this token.
	HostOrUsageLimits OrgTokenHostOrUsageLimitsPtrInput
	// Name of the token.
	Name pulumi.StringPtrInput
	// List of strings specifying where notifications will be sent when an incident occurs. See
	// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
	Notifications pulumi.StringArrayInput
}

func (OrgTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgTokenArgs)(nil)).Elem()
}
