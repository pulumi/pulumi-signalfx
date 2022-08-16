// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage SignalFx org tokens.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-signalfx/sdk/v5/go/signalfx"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := signalfx.NewOrgToken(ctx, "myteamkey0", &signalfx.OrgTokenArgs{
//				Description: pulumi.String("My team's rad key"),
//				HostOrUsageLimits: &OrgTokenHostOrUsageLimitsArgs{
//					ContainerLimit:                      pulumi.Int(200),
//					ContainerNotificationThreshold:      pulumi.Int(180),
//					CustomMetricsLimit:                  pulumi.Int(1000),
//					CustomMetricsNotificationThreshold:  pulumi.Int(900),
//					HighResMetricsLimit:                 pulumi.Int(1000),
//					HighResMetricsNotificationThreshold: pulumi.Int(900),
//					HostLimit:                           pulumi.Int(100),
//					HostNotificationThreshold:           pulumi.Int(90),
//				},
//				Notifications: pulumi.StringArray{
//					pulumi.String("Email,foo-alerts@bar.com"),
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
type OrgToken struct {
	pulumi.CustomResourceState

	// Authentication scope, ex: INGEST, API, RUM ... (Optional)
	AuthScopes pulumi.StringArrayOutput `pulumi:"authScopes"`
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
	// Authentication scope, ex: INGEST, API, RUM ... (Optional)
	AuthScopes []string `pulumi:"authScopes"`
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
	// Authentication scope, ex: INGEST, API, RUM ... (Optional)
	AuthScopes pulumi.StringArrayInput
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
	// Authentication scope, ex: INGEST, API, RUM ... (Optional)
	AuthScopes []string `pulumi:"authScopes"`
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
	// Authentication scope, ex: INGEST, API, RUM ... (Optional)
	AuthScopes pulumi.StringArrayInput
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

func (*OrgToken) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgToken)(nil)).Elem()
}

func (i *OrgToken) ToOrgTokenOutput() OrgTokenOutput {
	return i.ToOrgTokenOutputWithContext(context.Background())
}

func (i *OrgToken) ToOrgTokenOutputWithContext(ctx context.Context) OrgTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgTokenOutput)
}

// OrgTokenArrayInput is an input type that accepts OrgTokenArray and OrgTokenArrayOutput values.
// You can construct a concrete instance of `OrgTokenArrayInput` via:
//
//	OrgTokenArray{ OrgTokenArgs{...} }
type OrgTokenArrayInput interface {
	pulumi.Input

	ToOrgTokenArrayOutput() OrgTokenArrayOutput
	ToOrgTokenArrayOutputWithContext(context.Context) OrgTokenArrayOutput
}

type OrgTokenArray []OrgTokenInput

func (OrgTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgToken)(nil)).Elem()
}

func (i OrgTokenArray) ToOrgTokenArrayOutput() OrgTokenArrayOutput {
	return i.ToOrgTokenArrayOutputWithContext(context.Background())
}

func (i OrgTokenArray) ToOrgTokenArrayOutputWithContext(ctx context.Context) OrgTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgTokenArrayOutput)
}

// OrgTokenMapInput is an input type that accepts OrgTokenMap and OrgTokenMapOutput values.
// You can construct a concrete instance of `OrgTokenMapInput` via:
//
//	OrgTokenMap{ "key": OrgTokenArgs{...} }
type OrgTokenMapInput interface {
	pulumi.Input

	ToOrgTokenMapOutput() OrgTokenMapOutput
	ToOrgTokenMapOutputWithContext(context.Context) OrgTokenMapOutput
}

type OrgTokenMap map[string]OrgTokenInput

func (OrgTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgToken)(nil)).Elem()
}

func (i OrgTokenMap) ToOrgTokenMapOutput() OrgTokenMapOutput {
	return i.ToOrgTokenMapOutputWithContext(context.Background())
}

func (i OrgTokenMap) ToOrgTokenMapOutputWithContext(ctx context.Context) OrgTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgTokenMapOutput)
}

type OrgTokenOutput struct{ *pulumi.OutputState }

func (OrgTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgToken)(nil)).Elem()
}

func (o OrgTokenOutput) ToOrgTokenOutput() OrgTokenOutput {
	return o
}

func (o OrgTokenOutput) ToOrgTokenOutputWithContext(ctx context.Context) OrgTokenOutput {
	return o
}

// Authentication scope, ex: INGEST, API, RUM ... (Optional)
func (o OrgTokenOutput) AuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrgToken) pulumi.StringArrayOutput { return v.AuthScopes }).(pulumi.StringArrayOutput)
}

// Description of the token.
func (o OrgTokenOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgToken) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
func (o OrgTokenOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OrgToken) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Specify DPM-based limits for this token.
func (o OrgTokenOutput) DpmLimits() OrgTokenDpmLimitsPtrOutput {
	return o.ApplyT(func(v *OrgToken) OrgTokenDpmLimitsPtrOutput { return v.DpmLimits }).(OrgTokenDpmLimitsPtrOutput)
}

// Specify Usage-based limits for this token.
func (o OrgTokenOutput) HostOrUsageLimits() OrgTokenHostOrUsageLimitsPtrOutput {
	return o.ApplyT(func(v *OrgToken) OrgTokenHostOrUsageLimitsPtrOutput { return v.HostOrUsageLimits }).(OrgTokenHostOrUsageLimitsPtrOutput)
}

// Name of the token.
func (o OrgTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of strings specifying where notifications will be sent when an incident occurs. See
// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
func (o OrgTokenOutput) Notifications() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrgToken) pulumi.StringArrayOutput { return v.Notifications }).(pulumi.StringArrayOutput)
}

// The secret token created by the API. You cannot set this value.
func (o OrgTokenOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgToken) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

type OrgTokenArrayOutput struct{ *pulumi.OutputState }

func (OrgTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgToken)(nil)).Elem()
}

func (o OrgTokenArrayOutput) ToOrgTokenArrayOutput() OrgTokenArrayOutput {
	return o
}

func (o OrgTokenArrayOutput) ToOrgTokenArrayOutputWithContext(ctx context.Context) OrgTokenArrayOutput {
	return o
}

func (o OrgTokenArrayOutput) Index(i pulumi.IntInput) OrgTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgToken {
		return vs[0].([]*OrgToken)[vs[1].(int)]
	}).(OrgTokenOutput)
}

type OrgTokenMapOutput struct{ *pulumi.OutputState }

func (OrgTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgToken)(nil)).Elem()
}

func (o OrgTokenMapOutput) ToOrgTokenMapOutput() OrgTokenMapOutput {
	return o
}

func (o OrgTokenMapOutput) ToOrgTokenMapOutputWithContext(ctx context.Context) OrgTokenMapOutput {
	return o
}

func (o OrgTokenMapOutput) MapIndex(k pulumi.StringInput) OrgTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgToken {
		return vs[0].(map[string]*OrgToken)[vs[1].(string)]
	}).(OrgTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgTokenInput)(nil)).Elem(), &OrgToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgTokenArrayInput)(nil)).Elem(), OrgTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgTokenMapInput)(nil)).Elem(), OrgTokenMap{})
	pulumi.RegisterOutputType(OrgTokenOutput{})
	pulumi.RegisterOutputType(OrgTokenArrayOutput{})
	pulumi.RegisterOutputType(OrgTokenMapOutput{})
}
