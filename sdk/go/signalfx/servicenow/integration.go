// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicenow

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Observability Cloud ServiceNow integrations. For help with this integration see [Integration with ServiceNow](https://docs.splunk.com/Observability/admin/notif-services/servicenow.html).
//
// > **NOTE** When managing integrations use a session token for an administrator to authenticate the Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-signalfx/sdk/v5/go/signalfx/servicenow"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicenow.NewIntegration(ctx, "serviceNowMyteam", &servicenow.IntegrationArgs{
//				AlertResolvedPayloadTemplate:  pulumi.String("{\"close_notes\": \"{{{messageTitle}}} (customized close msg)\"}"),
//				AlertTriggeredPayloadTemplate: pulumi.String("{\"short_description\": \"{{{messageTitle}}} (customized)\"}"),
//				Enabled:                       pulumi.Bool(false),
//				InstanceName:                  pulumi.String("myinst.service-now.com"),
//				IssueType:                     pulumi.String("Incident"),
//				Password:                      pulumi.String("youd0ntsee1t"),
//				Username:                      pulumi.String("thisis_me"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Integration struct {
	pulumi.CustomResourceState

	// A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
	AlertResolvedPayloadTemplate pulumi.StringPtrOutput `pulumi:"alertResolvedPayloadTemplate"`
	// A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
	AlertTriggeredPayloadTemplate pulumi.StringPtrOutput `pulumi:"alertTriggeredPayloadTemplate"`
	// Whether the integration is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Name of the ServiceNow instance, for example `myinst.service-now.com`.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
	IssueType pulumi.StringOutput `pulumi:"issueType"`
	// Name of the integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password used to authenticate the ServiceNow integration.
	Password pulumi.StringOutput `pulumi:"password"`
	// User name used to authenticate the ServiceNow integration.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.IssueType == nil {
		return nil, errors.New("invalid value for required argument 'IssueType'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource Integration
	err := ctx.RegisterResource("signalfx:servicenow/integration:Integration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("signalfx:servicenow/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
	AlertResolvedPayloadTemplate *string `pulumi:"alertResolvedPayloadTemplate"`
	// A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
	AlertTriggeredPayloadTemplate *string `pulumi:"alertTriggeredPayloadTemplate"`
	// Whether the integration is enabled.
	Enabled *bool `pulumi:"enabled"`
	// Name of the ServiceNow instance, for example `myinst.service-now.com`.
	InstanceName *string `pulumi:"instanceName"`
	// The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
	IssueType *string `pulumi:"issueType"`
	// Name of the integration.
	Name *string `pulumi:"name"`
	// Password used to authenticate the ServiceNow integration.
	Password *string `pulumi:"password"`
	// User name used to authenticate the ServiceNow integration.
	Username *string `pulumi:"username"`
}

type IntegrationState struct {
	// A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
	AlertResolvedPayloadTemplate pulumi.StringPtrInput
	// A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
	AlertTriggeredPayloadTemplate pulumi.StringPtrInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolPtrInput
	// Name of the ServiceNow instance, for example `myinst.service-now.com`.
	InstanceName pulumi.StringPtrInput
	// The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
	IssueType pulumi.StringPtrInput
	// Name of the integration.
	Name pulumi.StringPtrInput
	// Password used to authenticate the ServiceNow integration.
	Password pulumi.StringPtrInput
	// User name used to authenticate the ServiceNow integration.
	Username pulumi.StringPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
	AlertResolvedPayloadTemplate *string `pulumi:"alertResolvedPayloadTemplate"`
	// A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
	AlertTriggeredPayloadTemplate *string `pulumi:"alertTriggeredPayloadTemplate"`
	// Whether the integration is enabled.
	Enabled bool `pulumi:"enabled"`
	// Name of the ServiceNow instance, for example `myinst.service-now.com`.
	InstanceName string `pulumi:"instanceName"`
	// The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
	IssueType string `pulumi:"issueType"`
	// Name of the integration.
	Name *string `pulumi:"name"`
	// Password used to authenticate the ServiceNow integration.
	Password string `pulumi:"password"`
	// User name used to authenticate the ServiceNow integration.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
	AlertResolvedPayloadTemplate pulumi.StringPtrInput
	// A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
	AlertTriggeredPayloadTemplate pulumi.StringPtrInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolInput
	// Name of the ServiceNow instance, for example `myinst.service-now.com`.
	InstanceName pulumi.StringInput
	// The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
	IssueType pulumi.StringInput
	// Name of the integration.
	Name pulumi.StringPtrInput
	// Password used to authenticate the ServiceNow integration.
	Password pulumi.StringInput
	// User name used to authenticate the ServiceNow integration.
	Username pulumi.StringInput
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

// A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
func (o IntegrationOutput) AlertResolvedPayloadTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringPtrOutput { return v.AlertResolvedPayloadTemplate }).(pulumi.StringPtrOutput)
}

// A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
func (o IntegrationOutput) AlertTriggeredPayloadTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringPtrOutput { return v.AlertTriggeredPayloadTemplate }).(pulumi.StringPtrOutput)
}

// Whether the integration is enabled.
func (o IntegrationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Integration) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Name of the ServiceNow instance, for example `myinst.service-now.com`.
func (o IntegrationOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
func (o IntegrationOutput) IssueType() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.IssueType }).(pulumi.StringOutput)
}

// Name of the integration.
func (o IntegrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password used to authenticate the ServiceNow integration.
func (o IntegrationOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// User name used to authenticate the ServiceNow integration.
func (o IntegrationOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
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
