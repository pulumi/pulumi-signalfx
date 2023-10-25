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

// SignalFx Webhook integration.
//
// > **NOTE** When managing integrations, use a session token of an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
type WebhookIntegration struct {
	pulumi.CustomResourceState

	// Whether the integration is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// A header to send with the request
	Headers WebhookIntegrationHeaderArrayOutput `pulumi:"headers"`
	// Name of the integration.
	Name         pulumi.StringOutput    `pulumi:"name"`
	SharedSecret pulumi.StringPtrOutput `pulumi:"sharedSecret"`
	// The URL to request
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewWebhookIntegration registers a new resource with the given unique name, arguments, and options.
func NewWebhookIntegration(ctx *pulumi.Context,
	name string, args *WebhookIntegrationArgs, opts ...pulumi.ResourceOption) (*WebhookIntegration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Headers != nil {
		args.Headers = pulumi.ToSecret(args.Headers).(WebhookIntegrationHeaderArrayInput)
	}
	if args.SharedSecret != nil {
		args.SharedSecret = pulumi.ToSecret(args.SharedSecret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"headers",
		"sharedSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebhookIntegration
	err := ctx.RegisterResource("signalfx:index/webhookIntegration:WebhookIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhookIntegration gets an existing WebhookIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhookIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookIntegrationState, opts ...pulumi.ResourceOption) (*WebhookIntegration, error) {
	var resource WebhookIntegration
	err := ctx.ReadResource("signalfx:index/webhookIntegration:WebhookIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebhookIntegration resources.
type webhookIntegrationState struct {
	// Whether the integration is enabled.
	Enabled *bool `pulumi:"enabled"`
	// A header to send with the request
	Headers []WebhookIntegrationHeader `pulumi:"headers"`
	// Name of the integration.
	Name         *string `pulumi:"name"`
	SharedSecret *string `pulumi:"sharedSecret"`
	// The URL to request
	Url *string `pulumi:"url"`
}

type WebhookIntegrationState struct {
	// Whether the integration is enabled.
	Enabled pulumi.BoolPtrInput
	// A header to send with the request
	Headers WebhookIntegrationHeaderArrayInput
	// Name of the integration.
	Name         pulumi.StringPtrInput
	SharedSecret pulumi.StringPtrInput
	// The URL to request
	Url pulumi.StringPtrInput
}

func (WebhookIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookIntegrationState)(nil)).Elem()
}

type webhookIntegrationArgs struct {
	// Whether the integration is enabled.
	Enabled bool `pulumi:"enabled"`
	// A header to send with the request
	Headers []WebhookIntegrationHeader `pulumi:"headers"`
	// Name of the integration.
	Name         *string `pulumi:"name"`
	SharedSecret *string `pulumi:"sharedSecret"`
	// The URL to request
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a WebhookIntegration resource.
type WebhookIntegrationArgs struct {
	// Whether the integration is enabled.
	Enabled pulumi.BoolInput
	// A header to send with the request
	Headers WebhookIntegrationHeaderArrayInput
	// Name of the integration.
	Name         pulumi.StringPtrInput
	SharedSecret pulumi.StringPtrInput
	// The URL to request
	Url pulumi.StringPtrInput
}

func (WebhookIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookIntegrationArgs)(nil)).Elem()
}

type WebhookIntegrationInput interface {
	pulumi.Input

	ToWebhookIntegrationOutput() WebhookIntegrationOutput
	ToWebhookIntegrationOutputWithContext(ctx context.Context) WebhookIntegrationOutput
}

func (*WebhookIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookIntegration)(nil)).Elem()
}

func (i *WebhookIntegration) ToWebhookIntegrationOutput() WebhookIntegrationOutput {
	return i.ToWebhookIntegrationOutputWithContext(context.Background())
}

func (i *WebhookIntegration) ToWebhookIntegrationOutputWithContext(ctx context.Context) WebhookIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookIntegrationOutput)
}

func (i *WebhookIntegration) ToOutput(ctx context.Context) pulumix.Output[*WebhookIntegration] {
	return pulumix.Output[*WebhookIntegration]{
		OutputState: i.ToWebhookIntegrationOutputWithContext(ctx).OutputState,
	}
}

// WebhookIntegrationArrayInput is an input type that accepts WebhookIntegrationArray and WebhookIntegrationArrayOutput values.
// You can construct a concrete instance of `WebhookIntegrationArrayInput` via:
//
//	WebhookIntegrationArray{ WebhookIntegrationArgs{...} }
type WebhookIntegrationArrayInput interface {
	pulumi.Input

	ToWebhookIntegrationArrayOutput() WebhookIntegrationArrayOutput
	ToWebhookIntegrationArrayOutputWithContext(context.Context) WebhookIntegrationArrayOutput
}

type WebhookIntegrationArray []WebhookIntegrationInput

func (WebhookIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebhookIntegration)(nil)).Elem()
}

func (i WebhookIntegrationArray) ToWebhookIntegrationArrayOutput() WebhookIntegrationArrayOutput {
	return i.ToWebhookIntegrationArrayOutputWithContext(context.Background())
}

func (i WebhookIntegrationArray) ToWebhookIntegrationArrayOutputWithContext(ctx context.Context) WebhookIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookIntegrationArrayOutput)
}

func (i WebhookIntegrationArray) ToOutput(ctx context.Context) pulumix.Output[[]*WebhookIntegration] {
	return pulumix.Output[[]*WebhookIntegration]{
		OutputState: i.ToWebhookIntegrationArrayOutputWithContext(ctx).OutputState,
	}
}

// WebhookIntegrationMapInput is an input type that accepts WebhookIntegrationMap and WebhookIntegrationMapOutput values.
// You can construct a concrete instance of `WebhookIntegrationMapInput` via:
//
//	WebhookIntegrationMap{ "key": WebhookIntegrationArgs{...} }
type WebhookIntegrationMapInput interface {
	pulumi.Input

	ToWebhookIntegrationMapOutput() WebhookIntegrationMapOutput
	ToWebhookIntegrationMapOutputWithContext(context.Context) WebhookIntegrationMapOutput
}

type WebhookIntegrationMap map[string]WebhookIntegrationInput

func (WebhookIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebhookIntegration)(nil)).Elem()
}

func (i WebhookIntegrationMap) ToWebhookIntegrationMapOutput() WebhookIntegrationMapOutput {
	return i.ToWebhookIntegrationMapOutputWithContext(context.Background())
}

func (i WebhookIntegrationMap) ToWebhookIntegrationMapOutputWithContext(ctx context.Context) WebhookIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookIntegrationMapOutput)
}

func (i WebhookIntegrationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WebhookIntegration] {
	return pulumix.Output[map[string]*WebhookIntegration]{
		OutputState: i.ToWebhookIntegrationMapOutputWithContext(ctx).OutputState,
	}
}

type WebhookIntegrationOutput struct{ *pulumi.OutputState }

func (WebhookIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookIntegration)(nil)).Elem()
}

func (o WebhookIntegrationOutput) ToWebhookIntegrationOutput() WebhookIntegrationOutput {
	return o
}

func (o WebhookIntegrationOutput) ToWebhookIntegrationOutputWithContext(ctx context.Context) WebhookIntegrationOutput {
	return o
}

func (o WebhookIntegrationOutput) ToOutput(ctx context.Context) pulumix.Output[*WebhookIntegration] {
	return pulumix.Output[*WebhookIntegration]{
		OutputState: o.OutputState,
	}
}

// Whether the integration is enabled.
func (o WebhookIntegrationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *WebhookIntegration) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// A header to send with the request
func (o WebhookIntegrationOutput) Headers() WebhookIntegrationHeaderArrayOutput {
	return o.ApplyT(func(v *WebhookIntegration) WebhookIntegrationHeaderArrayOutput { return v.Headers }).(WebhookIntegrationHeaderArrayOutput)
}

// Name of the integration.
func (o WebhookIntegrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebhookIntegration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebhookIntegrationOutput) SharedSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookIntegration) pulumi.StringPtrOutput { return v.SharedSecret }).(pulumi.StringPtrOutput)
}

// The URL to request
func (o WebhookIntegrationOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookIntegration) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

type WebhookIntegrationArrayOutput struct{ *pulumi.OutputState }

func (WebhookIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebhookIntegration)(nil)).Elem()
}

func (o WebhookIntegrationArrayOutput) ToWebhookIntegrationArrayOutput() WebhookIntegrationArrayOutput {
	return o
}

func (o WebhookIntegrationArrayOutput) ToWebhookIntegrationArrayOutputWithContext(ctx context.Context) WebhookIntegrationArrayOutput {
	return o
}

func (o WebhookIntegrationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WebhookIntegration] {
	return pulumix.Output[[]*WebhookIntegration]{
		OutputState: o.OutputState,
	}
}

func (o WebhookIntegrationArrayOutput) Index(i pulumi.IntInput) WebhookIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebhookIntegration {
		return vs[0].([]*WebhookIntegration)[vs[1].(int)]
	}).(WebhookIntegrationOutput)
}

type WebhookIntegrationMapOutput struct{ *pulumi.OutputState }

func (WebhookIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebhookIntegration)(nil)).Elem()
}

func (o WebhookIntegrationMapOutput) ToWebhookIntegrationMapOutput() WebhookIntegrationMapOutput {
	return o
}

func (o WebhookIntegrationMapOutput) ToWebhookIntegrationMapOutputWithContext(ctx context.Context) WebhookIntegrationMapOutput {
	return o
}

func (o WebhookIntegrationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WebhookIntegration] {
	return pulumix.Output[map[string]*WebhookIntegration]{
		OutputState: o.OutputState,
	}
}

func (o WebhookIntegrationMapOutput) MapIndex(k pulumi.StringInput) WebhookIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebhookIntegration {
		return vs[0].(map[string]*WebhookIntegration)[vs[1].(string)]
	}).(WebhookIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookIntegrationInput)(nil)).Elem(), &WebhookIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookIntegrationArrayInput)(nil)).Elem(), WebhookIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookIntegrationMapInput)(nil)).Elem(), WebhookIntegrationMap{})
	pulumi.RegisterOutputType(WebhookIntegrationOutput{})
	pulumi.RegisterOutputType(WebhookIntegrationArrayOutput{})
	pulumi.RegisterOutputType(WebhookIntegrationMapOutput{})
}
