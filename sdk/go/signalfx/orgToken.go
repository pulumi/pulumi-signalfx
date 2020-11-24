// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manage SignalFx org tokens.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-signalfx/sdk/v3/go/signalfx"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := signalfx.NewOrgToken(ctx, "myteamkey0", &signalfx.OrgTokenArgs{
// 			Description: pulumi.String("My team's rad key"),
// 			HostOrUsageLimits: &signalfx.OrgTokenHostOrUsageLimitsArgs{
// 				ContainerLimit:                      pulumi.Int(200),
// 				ContainerNotificationThreshold:      pulumi.Int(180),
// 				CustomMetricsLimit:                  pulumi.Int(1000),
// 				CustomMetricsNotificationThreshold:  pulumi.Int(900),
// 				HighResMetricsLimit:                 pulumi.Int(1000),
// 				HighResMetricsNotificationThreshold: pulumi.Int(900),
// 				HostLimit:                           pulumi.Int(100),
// 				HostNotificationThreshold:           pulumi.Int(90),
// 			},
// 			Notifications: pulumi.StringArray{
// 				pulumi.String("Email,foo-alerts@bar.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
	// The secret token created by the API. You cannot set this value.
	Secret pulumi.StringOutput `pulumi:"secret"`
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
	// The secret token created by the API. You cannot set this value.
	Secret *string `pulumi:"secret"`
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
	// The secret token created by the API. You cannot set this value.
	Secret pulumi.StringPtrInput
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

type OrgTokenInput interface {
	pulumi.Input

	ToOrgTokenOutput() OrgTokenOutput
	ToOrgTokenOutputWithContext(ctx context.Context) OrgTokenOutput
}

func (OrgToken) ElementType() reflect.Type {
	return reflect.TypeOf((*OrgToken)(nil)).Elem()
}

func (i OrgToken) ToOrgTokenOutput() OrgTokenOutput {
	return i.ToOrgTokenOutputWithContext(context.Background())
}

func (i OrgToken) ToOrgTokenOutputWithContext(ctx context.Context) OrgTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgTokenOutput)
}

type OrgTokenOutput struct {
	*pulumi.OutputState
}

func (OrgTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrgTokenOutput)(nil)).Elem()
}

func (o OrgTokenOutput) ToOrgTokenOutput() OrgTokenOutput {
	return o
}

func (o OrgTokenOutput) ToOrgTokenOutputWithContext(ctx context.Context) OrgTokenOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrgTokenOutput{})
}
