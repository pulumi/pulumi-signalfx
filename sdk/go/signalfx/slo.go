// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Splunk Observability Cloud slo resource. This can be used to create and manage SLOs.
//
// To learn more about this feature take a look on [documentation for SLO](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/slo-intro.html).
//
// ## Example
//
// ## Notification format
//
// As Splunk Observability Cloud supports different notification mechanisms, use a comma-delimited string to provide inputs. If you want to specify multiple notifications, each must be a member in the list, like so:
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// See [Splunk Observability Cloud Docs](https://dev.splunk.com/observability/reference/api/detectors/latest) for more information.
//
// Here are some example of how to configure each notification type:
//
// ### Email
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Jira
//
// Note that the `credentialId` is the Splunk-provided ID shown after setting up your Jira integration. See also `jira.Integration`.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### OpsGenie
//
// Note that the `credentialId` is the Splunk-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### PagerDuty
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Slack
//
// Exclude the `#` on the channel name:
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Team
//
// Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### TeamEmail
//
// Sends an email to every member of a team.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Splunk On-Call (formerly VictorOps)
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Webhooks
//
// You need to include all the commas even if you only use a credential id.
//
// You can either configure a Webhook to use an existing integration's credential id:
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// Or configure one inline:
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Arguments
//
// * `name` - (Required) Name of the SLO. Each SLO name must be unique within an organization.
// * `description` - (Optional) Description of the SLO.
// * `type` - (Required) Type of the SLO. Currently just: `"RequestBased"` is supported.
// * `input` - (Required) Properties to configure an SLO object inputs
//   - `programText` - (Required) SignalFlow program and arguments text strings that define the streams used as successful event count and total event count
//   - `goodEventsLabel` - (Required) Label used in `"programText"` that refers to the data block which contains the stream of successful events
//   - `totalEventsLabel` - (Required) Label used in `"programText"` that refers to the data block which contains the stream of total events
//
// * `target` - (Required) Define target value of the service level indicator in the appropriate time period.
//   - `type` - (Required) SLO target type can be the following type: `"RollingWindow"`
//   - `compliancePeriod` - (Required for `"RollingWindow"` type) Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
//   - `slo` - (Required) Target value in the form of a percentage
//   - `alertRule` - (Required) List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
//   - `type` - (Required) SLO alert rule can be one of the following types: BREACH, ERROR_BUDGET_LEFT, BURN_RATE. Within an SLO object, you can only specify one SLO alertRule per type. For example, you can't specify two alertRule of type BREACH. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
//   - `rule` - (Required) Set of rules used for alerting.
//   - `severity` - (Required) The severity of the rule, must be one of: `"Critical"`, `"Major"`, `"Minor"`, `"Warning"`, `"Info"`.
//   - `description` - (Optional) Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
//   - `disabled` - (Optional) When true, notifications and events will not be generated for the detect label. `false` by default.
//   - `notifications` - (Optional) List of strings specifying where notifications will be sent when an incident occurs. See [Create SLO](https://dev.splunk.com/observability/reference/api/slo/latest#endpoint-create-new-slo) for more info.
//   - `parameterizedBody` - (Optional) Custom notification message body when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
//   - `parameterizedSubject` - (Optional) Custom notification message subject when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
//   - `runbookUrl` - (Optional) URL of page to consult when an alert is triggered. This can be used with custom notification messages.
//   - `tip` - (Optional) Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
//   - `parameters` - (Optional) Parameters for the SLO alert rule. Each SLO alert rule type accepts different parameters. If not specified, default parameters are used.
//   - `fireLasting` - (Optional) Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: `"BREACH"` and `"ERROR_BUDGET_LEFT"` alert rules use the fireLasting parameter. Default: `"5m"`
//   - `percentOfLasting` - (Optional) Percentage of the `"fireLasting"` duration that the alert condition is met before the alert is triggered. Note: `"BREACH"` and `"ERROR_BUDGET_LEFT"` alert rules use the `"percentOfLasting"` parameter. Default: `100`
//   - `percentErrorBudgetLeft` - (Optional) Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: `"ERROR_BUDGET_LEFT"` alert rules use the `"percentErrorBudgetLeft"` parameter. Default: `100`
//   - `shortWindow1` - (Optional) Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of `"longWindow1"`. Note: `"BURN_RATE"` alert rules use the `"shortWindow1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
//   - `shortWindow2` - (Optional) Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of `"longWindow2"`. Note: `"BURN_RATE"` alert rules use the `"shortWindow2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
//   - `longWindow1` - (Optional) Long window 1 used in burn rate alert calculation. This value must be longer than `"shortWindow1"` and shorter than 90 days. Note: `"BURN_RATE"` alert rules use the `"longWindow1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
//   - `longWindow2` - (Optional) Long window 2 used in burn rate alert calculation. This value must be longer than `"shortWindow2"` and shorter than 90 days. Note: `"BURN_RATE"` alert rules use the `"longWindow2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
//   - `burnRateThreshold1` - (Optional) Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: `"BURN_RATE"` alert rules use the `"burnRateThreshold1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
//   - `burnRateThreshold2` - (Optional) Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: `"BURN_RATE"` alert rules use the `"burnRateThreshold2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
type Slo struct {
	pulumi.CustomResourceState

	// Description of the SLO
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// SignalFlow program and arguments text strings that define the streams used as successful event count and total event
	// count
	Input SloInputTypeOutput `pulumi:"input"`
	// Name of the SLO
	Name pulumi.StringOutput `pulumi:"name"`
	// Define target value of the service level indicator in the appropriate time period.
	Target SloTargetOutput `pulumi:"target"`
	// Type of the SLO. Currently only RequestBased SLO is supported
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSlo registers a new resource with the given unique name, arguments, and options.
func NewSlo(ctx *pulumi.Context,
	name string, args *SloArgs, opts ...pulumi.ResourceOption) (*Slo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Input == nil {
		return nil, errors.New("invalid value for required argument 'Input'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Slo
	err := ctx.RegisterResource("signalfx:index/slo:Slo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlo gets an existing Slo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SloState, opts ...pulumi.ResourceOption) (*Slo, error) {
	var resource Slo
	err := ctx.ReadResource("signalfx:index/slo:Slo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Slo resources.
type sloState struct {
	// Description of the SLO
	Description *string `pulumi:"description"`
	// SignalFlow program and arguments text strings that define the streams used as successful event count and total event
	// count
	Input *SloInputType `pulumi:"input"`
	// Name of the SLO
	Name *string `pulumi:"name"`
	// Define target value of the service level indicator in the appropriate time period.
	Target *SloTarget `pulumi:"target"`
	// Type of the SLO. Currently only RequestBased SLO is supported
	Type *string `pulumi:"type"`
}

type SloState struct {
	// Description of the SLO
	Description pulumi.StringPtrInput
	// SignalFlow program and arguments text strings that define the streams used as successful event count and total event
	// count
	Input SloInputTypePtrInput
	// Name of the SLO
	Name pulumi.StringPtrInput
	// Define target value of the service level indicator in the appropriate time period.
	Target SloTargetPtrInput
	// Type of the SLO. Currently only RequestBased SLO is supported
	Type pulumi.StringPtrInput
}

func (SloState) ElementType() reflect.Type {
	return reflect.TypeOf((*sloState)(nil)).Elem()
}

type sloArgs struct {
	// Description of the SLO
	Description *string `pulumi:"description"`
	// SignalFlow program and arguments text strings that define the streams used as successful event count and total event
	// count
	Input SloInputType `pulumi:"input"`
	// Name of the SLO
	Name *string `pulumi:"name"`
	// Define target value of the service level indicator in the appropriate time period.
	Target SloTarget `pulumi:"target"`
	// Type of the SLO. Currently only RequestBased SLO is supported
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Slo resource.
type SloArgs struct {
	// Description of the SLO
	Description pulumi.StringPtrInput
	// SignalFlow program and arguments text strings that define the streams used as successful event count and total event
	// count
	Input SloInputTypeInput
	// Name of the SLO
	Name pulumi.StringPtrInput
	// Define target value of the service level indicator in the appropriate time period.
	Target SloTargetInput
	// Type of the SLO. Currently only RequestBased SLO is supported
	Type pulumi.StringInput
}

func (SloArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sloArgs)(nil)).Elem()
}

type SloInput interface {
	pulumi.Input

	ToSloOutput() SloOutput
	ToSloOutputWithContext(ctx context.Context) SloOutput
}

func (*Slo) ElementType() reflect.Type {
	return reflect.TypeOf((**Slo)(nil)).Elem()
}

func (i *Slo) ToSloOutput() SloOutput {
	return i.ToSloOutputWithContext(context.Background())
}

func (i *Slo) ToSloOutputWithContext(ctx context.Context) SloOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SloOutput)
}

// SloArrayInput is an input type that accepts SloArray and SloArrayOutput values.
// You can construct a concrete instance of `SloArrayInput` via:
//
//	SloArray{ SloArgs{...} }
type SloArrayInput interface {
	pulumi.Input

	ToSloArrayOutput() SloArrayOutput
	ToSloArrayOutputWithContext(context.Context) SloArrayOutput
}

type SloArray []SloInput

func (SloArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Slo)(nil)).Elem()
}

func (i SloArray) ToSloArrayOutput() SloArrayOutput {
	return i.ToSloArrayOutputWithContext(context.Background())
}

func (i SloArray) ToSloArrayOutputWithContext(ctx context.Context) SloArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SloArrayOutput)
}

// SloMapInput is an input type that accepts SloMap and SloMapOutput values.
// You can construct a concrete instance of `SloMapInput` via:
//
//	SloMap{ "key": SloArgs{...} }
type SloMapInput interface {
	pulumi.Input

	ToSloMapOutput() SloMapOutput
	ToSloMapOutputWithContext(context.Context) SloMapOutput
}

type SloMap map[string]SloInput

func (SloMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Slo)(nil)).Elem()
}

func (i SloMap) ToSloMapOutput() SloMapOutput {
	return i.ToSloMapOutputWithContext(context.Background())
}

func (i SloMap) ToSloMapOutputWithContext(ctx context.Context) SloMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SloMapOutput)
}

type SloOutput struct{ *pulumi.OutputState }

func (SloOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Slo)(nil)).Elem()
}

func (o SloOutput) ToSloOutput() SloOutput {
	return o
}

func (o SloOutput) ToSloOutputWithContext(ctx context.Context) SloOutput {
	return o
}

// Description of the SLO
func (o SloOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Slo) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// SignalFlow program and arguments text strings that define the streams used as successful event count and total event
// count
func (o SloOutput) Input() SloInputTypeOutput {
	return o.ApplyT(func(v *Slo) SloInputTypeOutput { return v.Input }).(SloInputTypeOutput)
}

// Name of the SLO
func (o SloOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Slo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Define target value of the service level indicator in the appropriate time period.
func (o SloOutput) Target() SloTargetOutput {
	return o.ApplyT(func(v *Slo) SloTargetOutput { return v.Target }).(SloTargetOutput)
}

// Type of the SLO. Currently only RequestBased SLO is supported
func (o SloOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Slo) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type SloArrayOutput struct{ *pulumi.OutputState }

func (SloArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Slo)(nil)).Elem()
}

func (o SloArrayOutput) ToSloArrayOutput() SloArrayOutput {
	return o
}

func (o SloArrayOutput) ToSloArrayOutputWithContext(ctx context.Context) SloArrayOutput {
	return o
}

func (o SloArrayOutput) Index(i pulumi.IntInput) SloOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Slo {
		return vs[0].([]*Slo)[vs[1].(int)]
	}).(SloOutput)
}

type SloMapOutput struct{ *pulumi.OutputState }

func (SloMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Slo)(nil)).Elem()
}

func (o SloMapOutput) ToSloMapOutput() SloMapOutput {
	return o
}

func (o SloMapOutput) ToSloMapOutputWithContext(ctx context.Context) SloMapOutput {
	return o
}

func (o SloMapOutput) MapIndex(k pulumi.StringInput) SloOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Slo {
		return vs[0].(map[string]*Slo)[vs[1].(string)]
	}).(SloOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SloInput)(nil)).Elem(), &Slo{})
	pulumi.RegisterInputType(reflect.TypeOf((*SloArrayInput)(nil)).Elem(), SloArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SloMapInput)(nil)).Elem(), SloMap{})
	pulumi.RegisterOutputType(SloOutput{})
	pulumi.RegisterOutputType(SloArrayOutput{})
	pulumi.RegisterOutputType(SloMapOutput{})
}
