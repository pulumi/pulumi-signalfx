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

// Displays a listing of events as a widget in a dashboard.
//
// ## Arguments
//
// The following arguments are supported in the resource block:
//
// * `name` - (Required) Name of the text note.
// * `programText` - (Required) Signalflow program text for the chart. More info[in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
// * `description` - (Optional) Description of the text note.
// * `timeRange` - (Optional) From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
// * `startTime` - (Optional) Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
// * `endTime` - (Optional) Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
//
// ## Attributes
//
// In a addition to all arguments above, the following attributes are exported:
//
// * `id` - The ID of the chart.
// * `url` - The URL of the chart.
type EventFeedChart struct {
	pulumi.CustomResourceState

	// Description of the chart (Optional)
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Seconds since epoch to end the visualization
	EndTime pulumi.IntPtrOutput `pulumi:"endTime"`
	// Name of the chart
	Name pulumi.StringOutput `pulumi:"name"`
	// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
	ProgramText pulumi.StringOutput `pulumi:"programText"`
	// Seconds since epoch to start the visualization
	StartTime pulumi.IntPtrOutput `pulumi:"startTime"`
	// Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
	TimeRange pulumi.IntPtrOutput `pulumi:"timeRange"`
	// URL of the chart
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewEventFeedChart registers a new resource with the given unique name, arguments, and options.
func NewEventFeedChart(ctx *pulumi.Context,
	name string, args *EventFeedChartArgs, opts ...pulumi.ResourceOption) (*EventFeedChart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProgramText == nil {
		return nil, errors.New("invalid value for required argument 'ProgramText'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EventFeedChart
	err := ctx.RegisterResource("signalfx:index/eventFeedChart:EventFeedChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventFeedChart gets an existing EventFeedChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventFeedChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventFeedChartState, opts ...pulumi.ResourceOption) (*EventFeedChart, error) {
	var resource EventFeedChart
	err := ctx.ReadResource("signalfx:index/eventFeedChart:EventFeedChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventFeedChart resources.
type eventFeedChartState struct {
	// Description of the chart (Optional)
	Description *string `pulumi:"description"`
	// Seconds since epoch to end the visualization
	EndTime *int `pulumi:"endTime"`
	// Name of the chart
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
	ProgramText *string `pulumi:"programText"`
	// Seconds since epoch to start the visualization
	StartTime *int `pulumi:"startTime"`
	// Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
	TimeRange *int `pulumi:"timeRange"`
	// URL of the chart
	Url *string `pulumi:"url"`
}

type EventFeedChartState struct {
	// Description of the chart (Optional)
	Description pulumi.StringPtrInput
	// Seconds since epoch to end the visualization
	EndTime pulumi.IntPtrInput
	// Name of the chart
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
	ProgramText pulumi.StringPtrInput
	// Seconds since epoch to start the visualization
	StartTime pulumi.IntPtrInput
	// Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
	TimeRange pulumi.IntPtrInput
	// URL of the chart
	Url pulumi.StringPtrInput
}

func (EventFeedChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventFeedChartState)(nil)).Elem()
}

type eventFeedChartArgs struct {
	// Description of the chart (Optional)
	Description *string `pulumi:"description"`
	// Seconds since epoch to end the visualization
	EndTime *int `pulumi:"endTime"`
	// Name of the chart
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
	ProgramText string `pulumi:"programText"`
	// Seconds since epoch to start the visualization
	StartTime *int `pulumi:"startTime"`
	// Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
	TimeRange *int `pulumi:"timeRange"`
}

// The set of arguments for constructing a EventFeedChart resource.
type EventFeedChartArgs struct {
	// Description of the chart (Optional)
	Description pulumi.StringPtrInput
	// Seconds since epoch to end the visualization
	EndTime pulumi.IntPtrInput
	// Name of the chart
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
	ProgramText pulumi.StringInput
	// Seconds since epoch to start the visualization
	StartTime pulumi.IntPtrInput
	// Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
	TimeRange pulumi.IntPtrInput
}

func (EventFeedChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventFeedChartArgs)(nil)).Elem()
}

type EventFeedChartInput interface {
	pulumi.Input

	ToEventFeedChartOutput() EventFeedChartOutput
	ToEventFeedChartOutputWithContext(ctx context.Context) EventFeedChartOutput
}

func (*EventFeedChart) ElementType() reflect.Type {
	return reflect.TypeOf((**EventFeedChart)(nil)).Elem()
}

func (i *EventFeedChart) ToEventFeedChartOutput() EventFeedChartOutput {
	return i.ToEventFeedChartOutputWithContext(context.Background())
}

func (i *EventFeedChart) ToEventFeedChartOutputWithContext(ctx context.Context) EventFeedChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventFeedChartOutput)
}

// EventFeedChartArrayInput is an input type that accepts EventFeedChartArray and EventFeedChartArrayOutput values.
// You can construct a concrete instance of `EventFeedChartArrayInput` via:
//
//	EventFeedChartArray{ EventFeedChartArgs{...} }
type EventFeedChartArrayInput interface {
	pulumi.Input

	ToEventFeedChartArrayOutput() EventFeedChartArrayOutput
	ToEventFeedChartArrayOutputWithContext(context.Context) EventFeedChartArrayOutput
}

type EventFeedChartArray []EventFeedChartInput

func (EventFeedChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventFeedChart)(nil)).Elem()
}

func (i EventFeedChartArray) ToEventFeedChartArrayOutput() EventFeedChartArrayOutput {
	return i.ToEventFeedChartArrayOutputWithContext(context.Background())
}

func (i EventFeedChartArray) ToEventFeedChartArrayOutputWithContext(ctx context.Context) EventFeedChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventFeedChartArrayOutput)
}

// EventFeedChartMapInput is an input type that accepts EventFeedChartMap and EventFeedChartMapOutput values.
// You can construct a concrete instance of `EventFeedChartMapInput` via:
//
//	EventFeedChartMap{ "key": EventFeedChartArgs{...} }
type EventFeedChartMapInput interface {
	pulumi.Input

	ToEventFeedChartMapOutput() EventFeedChartMapOutput
	ToEventFeedChartMapOutputWithContext(context.Context) EventFeedChartMapOutput
}

type EventFeedChartMap map[string]EventFeedChartInput

func (EventFeedChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventFeedChart)(nil)).Elem()
}

func (i EventFeedChartMap) ToEventFeedChartMapOutput() EventFeedChartMapOutput {
	return i.ToEventFeedChartMapOutputWithContext(context.Background())
}

func (i EventFeedChartMap) ToEventFeedChartMapOutputWithContext(ctx context.Context) EventFeedChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventFeedChartMapOutput)
}

type EventFeedChartOutput struct{ *pulumi.OutputState }

func (EventFeedChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventFeedChart)(nil)).Elem()
}

func (o EventFeedChartOutput) ToEventFeedChartOutput() EventFeedChartOutput {
	return o
}

func (o EventFeedChartOutput) ToEventFeedChartOutputWithContext(ctx context.Context) EventFeedChartOutput {
	return o
}

// Description of the chart (Optional)
func (o EventFeedChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventFeedChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Seconds since epoch to end the visualization
func (o EventFeedChartOutput) EndTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EventFeedChart) pulumi.IntPtrOutput { return v.EndTime }).(pulumi.IntPtrOutput)
}

// Name of the chart
func (o EventFeedChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventFeedChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
func (o EventFeedChartOutput) ProgramText() pulumi.StringOutput {
	return o.ApplyT(func(v *EventFeedChart) pulumi.StringOutput { return v.ProgramText }).(pulumi.StringOutput)
}

// Seconds since epoch to start the visualization
func (o EventFeedChartOutput) StartTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EventFeedChart) pulumi.IntPtrOutput { return v.StartTime }).(pulumi.IntPtrOutput)
}

// Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
func (o EventFeedChartOutput) TimeRange() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *EventFeedChart) pulumi.IntPtrOutput { return v.TimeRange }).(pulumi.IntPtrOutput)
}

// URL of the chart
func (o EventFeedChartOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *EventFeedChart) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type EventFeedChartArrayOutput struct{ *pulumi.OutputState }

func (EventFeedChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventFeedChart)(nil)).Elem()
}

func (o EventFeedChartArrayOutput) ToEventFeedChartArrayOutput() EventFeedChartArrayOutput {
	return o
}

func (o EventFeedChartArrayOutput) ToEventFeedChartArrayOutputWithContext(ctx context.Context) EventFeedChartArrayOutput {
	return o
}

func (o EventFeedChartArrayOutput) Index(i pulumi.IntInput) EventFeedChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventFeedChart {
		return vs[0].([]*EventFeedChart)[vs[1].(int)]
	}).(EventFeedChartOutput)
}

type EventFeedChartMapOutput struct{ *pulumi.OutputState }

func (EventFeedChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventFeedChart)(nil)).Elem()
}

func (o EventFeedChartMapOutput) ToEventFeedChartMapOutput() EventFeedChartMapOutput {
	return o
}

func (o EventFeedChartMapOutput) ToEventFeedChartMapOutputWithContext(ctx context.Context) EventFeedChartMapOutput {
	return o
}

func (o EventFeedChartMapOutput) MapIndex(k pulumi.StringInput) EventFeedChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventFeedChart {
		return vs[0].(map[string]*EventFeedChart)[vs[1].(string)]
	}).(EventFeedChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventFeedChartInput)(nil)).Elem(), &EventFeedChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventFeedChartArrayInput)(nil)).Elem(), EventFeedChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventFeedChartMapInput)(nil)).Elem(), EventFeedChartMap{})
	pulumi.RegisterOutputType(EventFeedChartOutput{})
	pulumi.RegisterOutputType(EventFeedChartArrayOutput{})
	pulumi.RegisterOutputType(EventFeedChartMapOutput{})
}
