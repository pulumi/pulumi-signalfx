// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Displays a listing of events as a widget in a dashboard.
type EventFeedChart struct {
	pulumi.CustomResourceState

	// Description of the text note.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrOutput `pulumi:"endTime"`
	// Name of the text note.
	Name pulumi.StringOutput `pulumi:"name"`
	// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringOutput `pulumi:"programText"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrOutput `pulumi:"startTime"`
	// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrOutput `pulumi:"timeRange"`
	// The URL of the chart.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewEventFeedChart registers a new resource with the given unique name, arguments, and options.
func NewEventFeedChart(ctx *pulumi.Context,
	name string, args *EventFeedChartArgs, opts ...pulumi.ResourceOption) (*EventFeedChart, error) {
	if args == nil || args.ProgramText == nil {
		return nil, errors.New("missing required argument 'ProgramText'")
	}
	if args == nil {
		args = &EventFeedChartArgs{}
	}
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
	// Description of the text note.
	Description *string `pulumi:"description"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime *int `pulumi:"endTime"`
	// Name of the text note.
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText *string `pulumi:"programText"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime *int `pulumi:"startTime"`
	// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
	TimeRange *int `pulumi:"timeRange"`
	// The URL of the chart.
	Url *string `pulumi:"url"`
}

type EventFeedChartState struct {
	// Description of the text note.
	Description pulumi.StringPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrInput
	// Name of the text note.
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrInput
	// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrInput
	// The URL of the chart.
	Url pulumi.StringPtrInput
}

func (EventFeedChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventFeedChartState)(nil)).Elem()
}

type eventFeedChartArgs struct {
	// Description of the text note.
	Description *string `pulumi:"description"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime *int `pulumi:"endTime"`
	// Name of the text note.
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText string `pulumi:"programText"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime *int `pulumi:"startTime"`
	// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
	TimeRange *int `pulumi:"timeRange"`
}

// The set of arguments for constructing a EventFeedChart resource.
type EventFeedChartArgs struct {
	// Description of the text note.
	Description pulumi.StringPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrInput
	// Name of the text note.
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrInput
	// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
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

func (EventFeedChart) ElementType() reflect.Type {
	return reflect.TypeOf((*EventFeedChart)(nil)).Elem()
}

func (i EventFeedChart) ToEventFeedChartOutput() EventFeedChartOutput {
	return i.ToEventFeedChartOutputWithContext(context.Background())
}

func (i EventFeedChart) ToEventFeedChartOutputWithContext(ctx context.Context) EventFeedChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventFeedChartOutput)
}

type EventFeedChartOutput struct {
	*pulumi.OutputState
}

func (EventFeedChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventFeedChartOutput)(nil)).Elem()
}

func (o EventFeedChartOutput) ToEventFeedChartOutput() EventFeedChartOutput {
	return o
}

func (o EventFeedChartOutput) ToEventFeedChartOutputWithContext(ctx context.Context) EventFeedChartOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventFeedChartOutput{})
}
