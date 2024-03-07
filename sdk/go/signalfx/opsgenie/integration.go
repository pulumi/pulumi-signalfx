// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opsgenie

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Splunk Observability Cloud Opsgenie integration.
//
// > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
//
// ## Example
//
// ## Arguments
//
// * `name` - (Required) Name of the integration.
// * `enabled` - (Required) Whether the integration is enabled.
// * `apiKey` - (Required) The API key
// * `apiUrl` - (Optional) Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
//
// ## Attributes
//
// In a addition to all arguments above, the following attributes are exported:
//
// * `id` - The ID of the integration.
type Integration struct {
	pulumi.CustomResourceState

	// Opsgenie API key
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// Opsgenie API URL for integration
	ApiUrl pulumi.StringPtrOutput `pulumi:"apiUrl"`
	// Whether the integration is enabled or not
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Name of the integration
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ApiKey != nil {
		args.ApiKey = pulumi.ToSecret(args.ApiKey).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Integration
	err := ctx.RegisterResource("signalfx:opsgenie/integration:Integration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegration gets an existing Integration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationState, opts ...pulumi.ResourceOption) (*Integration, error) {
	var resource Integration
	err := ctx.ReadResource("signalfx:opsgenie/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// Opsgenie API key
	ApiKey *string `pulumi:"apiKey"`
	// Opsgenie API URL for integration
	ApiUrl *string `pulumi:"apiUrl"`
	// Whether the integration is enabled or not
	Enabled *bool `pulumi:"enabled"`
	// Name of the integration
	Name *string `pulumi:"name"`
}

type IntegrationState struct {
	// Opsgenie API key
	ApiKey pulumi.StringPtrInput
	// Opsgenie API URL for integration
	ApiUrl pulumi.StringPtrInput
	// Whether the integration is enabled or not
	Enabled pulumi.BoolPtrInput
	// Name of the integration
	Name pulumi.StringPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// Opsgenie API key
	ApiKey string `pulumi:"apiKey"`
	// Opsgenie API URL for integration
	ApiUrl *string `pulumi:"apiUrl"`
	// Whether the integration is enabled or not
	Enabled bool `pulumi:"enabled"`
	// Name of the integration
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// Opsgenie API key
	ApiKey pulumi.StringInput
	// Opsgenie API URL for integration
	ApiUrl pulumi.StringPtrInput
	// Whether the integration is enabled or not
	Enabled pulumi.BoolInput
	// Name of the integration
	Name pulumi.StringPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}

type IntegrationInput interface {
	pulumi.Input

	ToIntegrationOutput() IntegrationOutput
	ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput
}

func (*Integration) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil)).Elem()
}

func (i *Integration) ToIntegrationOutput() IntegrationOutput {
	return i.ToIntegrationOutputWithContext(context.Background())
}

func (i *Integration) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationOutput)
}

// IntegrationArrayInput is an input type that accepts IntegrationArray and IntegrationArrayOutput values.
// You can construct a concrete instance of `IntegrationArrayInput` via:
//
//	IntegrationArray{ IntegrationArgs{...} }
type IntegrationArrayInput interface {
	pulumi.Input

	ToIntegrationArrayOutput() IntegrationArrayOutput
	ToIntegrationArrayOutputWithContext(context.Context) IntegrationArrayOutput
}

type IntegrationArray []IntegrationInput

func (IntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Integration)(nil)).Elem()
}

func (i IntegrationArray) ToIntegrationArrayOutput() IntegrationArrayOutput {
	return i.ToIntegrationArrayOutputWithContext(context.Background())
}

func (i IntegrationArray) ToIntegrationArrayOutputWithContext(ctx context.Context) IntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationArrayOutput)
}

// IntegrationMapInput is an input type that accepts IntegrationMap and IntegrationMapOutput values.
// You can construct a concrete instance of `IntegrationMapInput` via:
//
//	IntegrationMap{ "key": IntegrationArgs{...} }
type IntegrationMapInput interface {
	pulumi.Input

	ToIntegrationMapOutput() IntegrationMapOutput
	ToIntegrationMapOutputWithContext(context.Context) IntegrationMapOutput
}

type IntegrationMap map[string]IntegrationInput

func (IntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Integration)(nil)).Elem()
}

func (i IntegrationMap) ToIntegrationMapOutput() IntegrationMapOutput {
	return i.ToIntegrationMapOutputWithContext(context.Background())
}

func (i IntegrationMap) ToIntegrationMapOutputWithContext(ctx context.Context) IntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMapOutput)
}

type IntegrationOutput struct{ *pulumi.OutputState }

func (IntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil)).Elem()
}

func (o IntegrationOutput) ToIntegrationOutput() IntegrationOutput {
	return o
}

func (o IntegrationOutput) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return o
}

// Opsgenie API key
func (o IntegrationOutput) ApiKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.ApiKey }).(pulumi.StringOutput)
}

// Opsgenie API URL for integration
func (o IntegrationOutput) ApiUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringPtrOutput { return v.ApiUrl }).(pulumi.StringPtrOutput)
}

// Whether the integration is enabled or not
func (o IntegrationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Integration) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Name of the integration
func (o IntegrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type IntegrationArrayOutput struct{ *pulumi.OutputState }

func (IntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Integration)(nil)).Elem()
}

func (o IntegrationArrayOutput) ToIntegrationArrayOutput() IntegrationArrayOutput {
	return o
}

func (o IntegrationArrayOutput) ToIntegrationArrayOutputWithContext(ctx context.Context) IntegrationArrayOutput {
	return o
}

func (o IntegrationArrayOutput) Index(i pulumi.IntInput) IntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Integration {
		return vs[0].([]*Integration)[vs[1].(int)]
	}).(IntegrationOutput)
}

type IntegrationMapOutput struct{ *pulumi.OutputState }

func (IntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Integration)(nil)).Elem()
}

func (o IntegrationMapOutput) ToIntegrationMapOutput() IntegrationMapOutput {
	return o
}

func (o IntegrationMapOutput) ToIntegrationMapOutputWithContext(ctx context.Context) IntegrationMapOutput {
	return o
}

func (o IntegrationMapOutput) MapIndex(k pulumi.StringInput) IntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Integration {
		return vs[0].(map[string]*Integration)[vs[1].(string)]
	}).(IntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationInput)(nil)).Elem(), &Integration{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationArrayInput)(nil)).Elem(), IntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMapInput)(nil)).Elem(), IntegrationMap{})
	pulumi.RegisterOutputType(IntegrationOutput{})
	pulumi.RegisterOutputType(IntegrationArrayOutput{})
	pulumi.RegisterOutputType(IntegrationMapOutput{})
}
