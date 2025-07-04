// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Splunk Observability Cloud time chart resource. This can be used to create and manage the different types of time charts.
//
// Time charts display data points over a period of time.
//
// ## Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := signalfx.NewTimeChart(ctx, "mychart0", &signalfx.TimeChartArgs{
//				Name:            pulumi.String("CPU Total Idle"),
//				ProgramText:     pulumi.String("data(\"cpu.total.idle\").publish(label=\"CPU Idle\")\n"),
//				TimeRange:       pulumi.Int(3600),
//				PlotType:        pulumi.String("LineChart"),
//				ShowDataMarkers: pulumi.Bool(true),
//				LegendOptionsFields: signalfx.TimeChartLegendOptionsFieldArray{
//					&signalfx.TimeChartLegendOptionsFieldArgs{
//						Property: pulumi.String("collector"),
//						Enabled:  pulumi.Bool(false),
//					},
//					&signalfx.TimeChartLegendOptionsFieldArgs{
//						Property: pulumi.String("hostname"),
//						Enabled:  pulumi.Bool(false),
//					},
//				},
//				VizOptions: signalfx.TimeChartVizOptionArray{
//					&signalfx.TimeChartVizOptionArgs{
//						Label: pulumi.String("CPU Idle"),
//						Axis:  pulumi.String("left"),
//						Color: pulumi.String("orange"),
//					},
//				},
//				AxisLeft: &signalfx.TimeChartAxisLeftArgs{
//					Label:        pulumi.String("CPU Total Idle"),
//					LowWatermark: pulumi.Float64(1000),
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
type TimeChart struct {
	pulumi.CustomResourceState

	// Force the chart to display zero on the y-axes, even if none of the data is near zero.
	AxesIncludeZero pulumi.BoolPtrOutput `pulumi:"axesIncludeZero"`
	// Specifies the digits Splunk Observability Cloud displays for values plotted on the chart. Defaults to `3`.
	AxesPrecision pulumi.IntPtrOutput `pulumi:"axesPrecision"`
	// Set of axis options.
	AxisLeft TimeChartAxisLeftPtrOutput `pulumi:"axisLeft"`
	// Set of axis options.
	AxisRight TimeChartAxisRightPtrOutput `pulumi:"axisRight"`
	// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy pulumi.StringPtrOutput `pulumi:"colorBy"`
	// Description of the chart.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
	DisableSampling pulumi.BoolPtrOutput `pulumi:"disableSampling"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrOutput `pulumi:"endTime"`
	// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
	EventOptions TimeChartEventOptionArrayOutput `pulumi:"eventOptions"`
	// Only used when `plotType` is `"Histogram"`. Histogram specific options.
	HistogramOptions TimeChartHistogramOptionArrayOutput `pulumi:"histogramOptions"`
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	//
	// Deprecated: Please use legend_options_fields
	LegendFieldsToHides pulumi.StringArrayOutput `pulumi:"legendFieldsToHides"`
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	LegendOptionsFields TimeChartLegendOptionsFieldArrayOutput `pulumi:"legendOptionsFields"`
	// How long (in seconds) to wait for late datapoints.
	MaxDelay pulumi.IntPtrOutput `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution pulumi.IntPtrOutput `pulumi:"minimumResolution"`
	// Name of the chart.
	Name pulumi.StringOutput `pulumi:"name"`
	// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
	OnChartLegendDimension pulumi.StringPtrOutput `pulumi:"onChartLegendDimension"`
	// The default plot display style for the visualization. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Default: `"LineChart"`.
	PlotType pulumi.StringPtrOutput `pulumi:"plotType"`
	// Signalflow program text for the chart. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
	ProgramText pulumi.StringOutput `pulumi:"programText"`
	// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
	ShowDataMarkers pulumi.BoolPtrOutput `pulumi:"showDataMarkers"`
	// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
	ShowEventLines pulumi.BoolPtrOutput `pulumi:"showEventLines"`
	// Whether area and bar charts in the visualization should be stacked. `false` by default.
	Stacked pulumi.BoolPtrOutput `pulumi:"stacked"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrOutput `pulumi:"startTime"`
	// Tags associated with the chart
	//
	// Deprecated: signalfx_time_chart.tags is being removed in the next release
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrOutput `pulumi:"timeRange"`
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://dev.splunk.com/observability/docs/signalflow/). `"UTC"` by default.
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix pulumi.StringPtrOutput `pulumi:"unitPrefix"`
	// The URL of the chart.
	Url pulumi.StringOutput `pulumi:"url"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions TimeChartVizOptionArrayOutput `pulumi:"vizOptions"`
}

// NewTimeChart registers a new resource with the given unique name, arguments, and options.
func NewTimeChart(ctx *pulumi.Context,
	name string, args *TimeChartArgs, opts ...pulumi.ResourceOption) (*TimeChart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProgramText == nil {
		return nil, errors.New("invalid value for required argument 'ProgramText'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TimeChart
	err := ctx.RegisterResource("signalfx:index/timeChart:TimeChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTimeChart gets an existing TimeChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTimeChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TimeChartState, opts ...pulumi.ResourceOption) (*TimeChart, error) {
	var resource TimeChart
	err := ctx.ReadResource("signalfx:index/timeChart:TimeChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TimeChart resources.
type timeChartState struct {
	// Force the chart to display zero on the y-axes, even if none of the data is near zero.
	AxesIncludeZero *bool `pulumi:"axesIncludeZero"`
	// Specifies the digits Splunk Observability Cloud displays for values plotted on the chart. Defaults to `3`.
	AxesPrecision *int `pulumi:"axesPrecision"`
	// Set of axis options.
	AxisLeft *TimeChartAxisLeft `pulumi:"axisLeft"`
	// Set of axis options.
	AxisRight *TimeChartAxisRight `pulumi:"axisRight"`
	// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy *string `pulumi:"colorBy"`
	// Description of the chart.
	Description *string `pulumi:"description"`
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
	DisableSampling *bool `pulumi:"disableSampling"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime *int `pulumi:"endTime"`
	// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
	EventOptions []TimeChartEventOption `pulumi:"eventOptions"`
	// Only used when `plotType` is `"Histogram"`. Histogram specific options.
	HistogramOptions []TimeChartHistogramOption `pulumi:"histogramOptions"`
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	//
	// Deprecated: Please use legend_options_fields
	LegendFieldsToHides []string `pulumi:"legendFieldsToHides"`
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	LegendOptionsFields []TimeChartLegendOptionsField `pulumi:"legendOptionsFields"`
	// How long (in seconds) to wait for late datapoints.
	MaxDelay *int `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution *int `pulumi:"minimumResolution"`
	// Name of the chart.
	Name *string `pulumi:"name"`
	// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
	OnChartLegendDimension *string `pulumi:"onChartLegendDimension"`
	// The default plot display style for the visualization. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Default: `"LineChart"`.
	PlotType *string `pulumi:"plotType"`
	// Signalflow program text for the chart. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
	ProgramText *string `pulumi:"programText"`
	// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
	ShowDataMarkers *bool `pulumi:"showDataMarkers"`
	// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
	ShowEventLines *bool `pulumi:"showEventLines"`
	// Whether area and bar charts in the visualization should be stacked. `false` by default.
	Stacked *bool `pulumi:"stacked"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime *int `pulumi:"startTime"`
	// Tags associated with the chart
	//
	// Deprecated: signalfx_time_chart.tags is being removed in the next release
	Tags []string `pulumi:"tags"`
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange *int `pulumi:"timeRange"`
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://dev.splunk.com/observability/docs/signalflow/). `"UTC"` by default.
	Timezone *string `pulumi:"timezone"`
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix *string `pulumi:"unitPrefix"`
	// The URL of the chart.
	Url *string `pulumi:"url"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions []TimeChartVizOption `pulumi:"vizOptions"`
}

type TimeChartState struct {
	// Force the chart to display zero on the y-axes, even if none of the data is near zero.
	AxesIncludeZero pulumi.BoolPtrInput
	// Specifies the digits Splunk Observability Cloud displays for values plotted on the chart. Defaults to `3`.
	AxesPrecision pulumi.IntPtrInput
	// Set of axis options.
	AxisLeft TimeChartAxisLeftPtrInput
	// Set of axis options.
	AxisRight TimeChartAxisRightPtrInput
	// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy pulumi.StringPtrInput
	// Description of the chart.
	Description pulumi.StringPtrInput
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
	DisableSampling pulumi.BoolPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrInput
	// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
	EventOptions TimeChartEventOptionArrayInput
	// Only used when `plotType` is `"Histogram"`. Histogram specific options.
	HistogramOptions TimeChartHistogramOptionArrayInput
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	//
	// Deprecated: Please use legend_options_fields
	LegendFieldsToHides pulumi.StringArrayInput
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	LegendOptionsFields TimeChartLegendOptionsFieldArrayInput
	// How long (in seconds) to wait for late datapoints.
	MaxDelay pulumi.IntPtrInput
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution pulumi.IntPtrInput
	// Name of the chart.
	Name pulumi.StringPtrInput
	// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
	OnChartLegendDimension pulumi.StringPtrInput
	// The default plot display style for the visualization. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Default: `"LineChart"`.
	PlotType pulumi.StringPtrInput
	// Signalflow program text for the chart. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
	ProgramText pulumi.StringPtrInput
	// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
	ShowDataMarkers pulumi.BoolPtrInput
	// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
	ShowEventLines pulumi.BoolPtrInput
	// Whether area and bar charts in the visualization should be stacked. `false` by default.
	Stacked pulumi.BoolPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrInput
	// Tags associated with the chart
	//
	// Deprecated: signalfx_time_chart.tags is being removed in the next release
	Tags pulumi.StringArrayInput
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrInput
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://dev.splunk.com/observability/docs/signalflow/). `"UTC"` by default.
	Timezone pulumi.StringPtrInput
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix pulumi.StringPtrInput
	// The URL of the chart.
	Url pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement.
	VizOptions TimeChartVizOptionArrayInput
}

func (TimeChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*timeChartState)(nil)).Elem()
}

type timeChartArgs struct {
	// Force the chart to display zero on the y-axes, even if none of the data is near zero.
	AxesIncludeZero *bool `pulumi:"axesIncludeZero"`
	// Specifies the digits Splunk Observability Cloud displays for values plotted on the chart. Defaults to `3`.
	AxesPrecision *int `pulumi:"axesPrecision"`
	// Set of axis options.
	AxisLeft *TimeChartAxisLeft `pulumi:"axisLeft"`
	// Set of axis options.
	AxisRight *TimeChartAxisRight `pulumi:"axisRight"`
	// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy *string `pulumi:"colorBy"`
	// Description of the chart.
	Description *string `pulumi:"description"`
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
	DisableSampling *bool `pulumi:"disableSampling"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime *int `pulumi:"endTime"`
	// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
	EventOptions []TimeChartEventOption `pulumi:"eventOptions"`
	// Only used when `plotType` is `"Histogram"`. Histogram specific options.
	HistogramOptions []TimeChartHistogramOption `pulumi:"histogramOptions"`
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	//
	// Deprecated: Please use legend_options_fields
	LegendFieldsToHides []string `pulumi:"legendFieldsToHides"`
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	LegendOptionsFields []TimeChartLegendOptionsField `pulumi:"legendOptionsFields"`
	// How long (in seconds) to wait for late datapoints.
	MaxDelay *int `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution *int `pulumi:"minimumResolution"`
	// Name of the chart.
	Name *string `pulumi:"name"`
	// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
	OnChartLegendDimension *string `pulumi:"onChartLegendDimension"`
	// The default plot display style for the visualization. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Default: `"LineChart"`.
	PlotType *string `pulumi:"plotType"`
	// Signalflow program text for the chart. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
	ProgramText string `pulumi:"programText"`
	// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
	ShowDataMarkers *bool `pulumi:"showDataMarkers"`
	// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
	ShowEventLines *bool `pulumi:"showEventLines"`
	// Whether area and bar charts in the visualization should be stacked. `false` by default.
	Stacked *bool `pulumi:"stacked"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime *int `pulumi:"startTime"`
	// Tags associated with the chart
	//
	// Deprecated: signalfx_time_chart.tags is being removed in the next release
	Tags []string `pulumi:"tags"`
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange *int `pulumi:"timeRange"`
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://dev.splunk.com/observability/docs/signalflow/). `"UTC"` by default.
	Timezone *string `pulumi:"timezone"`
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix *string `pulumi:"unitPrefix"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions []TimeChartVizOption `pulumi:"vizOptions"`
}

// The set of arguments for constructing a TimeChart resource.
type TimeChartArgs struct {
	// Force the chart to display zero on the y-axes, even if none of the data is near zero.
	AxesIncludeZero pulumi.BoolPtrInput
	// Specifies the digits Splunk Observability Cloud displays for values plotted on the chart. Defaults to `3`.
	AxesPrecision pulumi.IntPtrInput
	// Set of axis options.
	AxisLeft TimeChartAxisLeftPtrInput
	// Set of axis options.
	AxisRight TimeChartAxisRightPtrInput
	// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy pulumi.StringPtrInput
	// Description of the chart.
	Description pulumi.StringPtrInput
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
	DisableSampling pulumi.BoolPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrInput
	// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
	EventOptions TimeChartEventOptionArrayInput
	// Only used when `plotType` is `"Histogram"`. Histogram specific options.
	HistogramOptions TimeChartHistogramOptionArrayInput
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	//
	// Deprecated: Please use legend_options_fields
	LegendFieldsToHides pulumi.StringArrayInput
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	LegendOptionsFields TimeChartLegendOptionsFieldArrayInput
	// How long (in seconds) to wait for late datapoints.
	MaxDelay pulumi.IntPtrInput
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution pulumi.IntPtrInput
	// Name of the chart.
	Name pulumi.StringPtrInput
	// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
	OnChartLegendDimension pulumi.StringPtrInput
	// The default plot display style for the visualization. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Default: `"LineChart"`.
	PlotType pulumi.StringPtrInput
	// Signalflow program text for the chart. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
	ProgramText pulumi.StringInput
	// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
	ShowDataMarkers pulumi.BoolPtrInput
	// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
	ShowEventLines pulumi.BoolPtrInput
	// Whether area and bar charts in the visualization should be stacked. `false` by default.
	Stacked pulumi.BoolPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrInput
	// Tags associated with the chart
	//
	// Deprecated: signalfx_time_chart.tags is being removed in the next release
	Tags pulumi.StringArrayInput
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrInput
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://dev.splunk.com/observability/docs/signalflow/). `"UTC"` by default.
	Timezone pulumi.StringPtrInput
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement.
	VizOptions TimeChartVizOptionArrayInput
}

func (TimeChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*timeChartArgs)(nil)).Elem()
}

type TimeChartInput interface {
	pulumi.Input

	ToTimeChartOutput() TimeChartOutput
	ToTimeChartOutputWithContext(ctx context.Context) TimeChartOutput
}

func (*TimeChart) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeChart)(nil)).Elem()
}

func (i *TimeChart) ToTimeChartOutput() TimeChartOutput {
	return i.ToTimeChartOutputWithContext(context.Background())
}

func (i *TimeChart) ToTimeChartOutputWithContext(ctx context.Context) TimeChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeChartOutput)
}

// TimeChartArrayInput is an input type that accepts TimeChartArray and TimeChartArrayOutput values.
// You can construct a concrete instance of `TimeChartArrayInput` via:
//
//	TimeChartArray{ TimeChartArgs{...} }
type TimeChartArrayInput interface {
	pulumi.Input

	ToTimeChartArrayOutput() TimeChartArrayOutput
	ToTimeChartArrayOutputWithContext(context.Context) TimeChartArrayOutput
}

type TimeChartArray []TimeChartInput

func (TimeChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TimeChart)(nil)).Elem()
}

func (i TimeChartArray) ToTimeChartArrayOutput() TimeChartArrayOutput {
	return i.ToTimeChartArrayOutputWithContext(context.Background())
}

func (i TimeChartArray) ToTimeChartArrayOutputWithContext(ctx context.Context) TimeChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeChartArrayOutput)
}

// TimeChartMapInput is an input type that accepts TimeChartMap and TimeChartMapOutput values.
// You can construct a concrete instance of `TimeChartMapInput` via:
//
//	TimeChartMap{ "key": TimeChartArgs{...} }
type TimeChartMapInput interface {
	pulumi.Input

	ToTimeChartMapOutput() TimeChartMapOutput
	ToTimeChartMapOutputWithContext(context.Context) TimeChartMapOutput
}

type TimeChartMap map[string]TimeChartInput

func (TimeChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TimeChart)(nil)).Elem()
}

func (i TimeChartMap) ToTimeChartMapOutput() TimeChartMapOutput {
	return i.ToTimeChartMapOutputWithContext(context.Background())
}

func (i TimeChartMap) ToTimeChartMapOutputWithContext(ctx context.Context) TimeChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeChartMapOutput)
}

type TimeChartOutput struct{ *pulumi.OutputState }

func (TimeChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeChart)(nil)).Elem()
}

func (o TimeChartOutput) ToTimeChartOutput() TimeChartOutput {
	return o
}

func (o TimeChartOutput) ToTimeChartOutputWithContext(ctx context.Context) TimeChartOutput {
	return o
}

// Force the chart to display zero on the y-axes, even if none of the data is near zero.
func (o TimeChartOutput) AxesIncludeZero() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.BoolPtrOutput { return v.AxesIncludeZero }).(pulumi.BoolPtrOutput)
}

// Specifies the digits Splunk Observability Cloud displays for values plotted on the chart. Defaults to `3`.
func (o TimeChartOutput) AxesPrecision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.IntPtrOutput { return v.AxesPrecision }).(pulumi.IntPtrOutput)
}

// Set of axis options.
func (o TimeChartOutput) AxisLeft() TimeChartAxisLeftPtrOutput {
	return o.ApplyT(func(v *TimeChart) TimeChartAxisLeftPtrOutput { return v.AxisLeft }).(TimeChartAxisLeftPtrOutput)
}

// Set of axis options.
func (o TimeChartOutput) AxisRight() TimeChartAxisRightPtrOutput {
	return o.ApplyT(func(v *TimeChart) TimeChartAxisRightPtrOutput { return v.AxisRight }).(TimeChartAxisRightPtrOutput)
}

// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
func (o TimeChartOutput) ColorBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringPtrOutput { return v.ColorBy }).(pulumi.StringPtrOutput)
}

// Description of the chart.
func (o TimeChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
func (o TimeChartOutput) DisableSampling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.BoolPtrOutput { return v.DisableSampling }).(pulumi.BoolPtrOutput)
}

// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
func (o TimeChartOutput) EndTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.IntPtrOutput { return v.EndTime }).(pulumi.IntPtrOutput)
}

// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
func (o TimeChartOutput) EventOptions() TimeChartEventOptionArrayOutput {
	return o.ApplyT(func(v *TimeChart) TimeChartEventOptionArrayOutput { return v.EventOptions }).(TimeChartEventOptionArrayOutput)
}

// Only used when `plotType` is `"Histogram"`. Histogram specific options.
func (o TimeChartOutput) HistogramOptions() TimeChartHistogramOptionArrayOutput {
	return o.ApplyT(func(v *TimeChart) TimeChartHistogramOptionArrayOutput { return v.HistogramOptions }).(TimeChartHistogramOptionArrayOutput)
}

// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
//
// Deprecated: Please use legend_options_fields
func (o TimeChartOutput) LegendFieldsToHides() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringArrayOutput { return v.LegendFieldsToHides }).(pulumi.StringArrayOutput)
}

// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
func (o TimeChartOutput) LegendOptionsFields() TimeChartLegendOptionsFieldArrayOutput {
	return o.ApplyT(func(v *TimeChart) TimeChartLegendOptionsFieldArrayOutput { return v.LegendOptionsFields }).(TimeChartLegendOptionsFieldArrayOutput)
}

// How long (in seconds) to wait for late datapoints.
func (o TimeChartOutput) MaxDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.IntPtrOutput { return v.MaxDelay }).(pulumi.IntPtrOutput)
}

// The minimum resolution (in seconds) to use for computing the underlying program.
func (o TimeChartOutput) MinimumResolution() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.IntPtrOutput { return v.MinimumResolution }).(pulumi.IntPtrOutput)
}

// Name of the chart.
func (o TimeChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
func (o TimeChartOutput) OnChartLegendDimension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringPtrOutput { return v.OnChartLegendDimension }).(pulumi.StringPtrOutput)
}

// The default plot display style for the visualization. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Default: `"LineChart"`.
func (o TimeChartOutput) PlotType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringPtrOutput { return v.PlotType }).(pulumi.StringPtrOutput)
}

// Signalflow program text for the chart. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
func (o TimeChartOutput) ProgramText() pulumi.StringOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringOutput { return v.ProgramText }).(pulumi.StringOutput)
}

// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
func (o TimeChartOutput) ShowDataMarkers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.BoolPtrOutput { return v.ShowDataMarkers }).(pulumi.BoolPtrOutput)
}

// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
func (o TimeChartOutput) ShowEventLines() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.BoolPtrOutput { return v.ShowEventLines }).(pulumi.BoolPtrOutput)
}

// Whether area and bar charts in the visualization should be stacked. `false` by default.
func (o TimeChartOutput) Stacked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.BoolPtrOutput { return v.Stacked }).(pulumi.BoolPtrOutput)
}

// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
func (o TimeChartOutput) StartTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.IntPtrOutput { return v.StartTime }).(pulumi.IntPtrOutput)
}

// Tags associated with the chart
//
// Deprecated: signalfx_time_chart.tags is being removed in the next release
func (o TimeChartOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
func (o TimeChartOutput) TimeRange() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.IntPtrOutput { return v.TimeRange }).(pulumi.IntPtrOutput)
}

// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://dev.splunk.com/observability/docs/signalflow/). `"UTC"` by default.
func (o TimeChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
func (o TimeChartOutput) UnitPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringPtrOutput { return v.UnitPrefix }).(pulumi.StringPtrOutput)
}

// The URL of the chart.
func (o TimeChartOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *TimeChart) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Plot-level customization options, associated with a publish statement.
func (o TimeChartOutput) VizOptions() TimeChartVizOptionArrayOutput {
	return o.ApplyT(func(v *TimeChart) TimeChartVizOptionArrayOutput { return v.VizOptions }).(TimeChartVizOptionArrayOutput)
}

type TimeChartArrayOutput struct{ *pulumi.OutputState }

func (TimeChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TimeChart)(nil)).Elem()
}

func (o TimeChartArrayOutput) ToTimeChartArrayOutput() TimeChartArrayOutput {
	return o
}

func (o TimeChartArrayOutput) ToTimeChartArrayOutputWithContext(ctx context.Context) TimeChartArrayOutput {
	return o
}

func (o TimeChartArrayOutput) Index(i pulumi.IntInput) TimeChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TimeChart {
		return vs[0].([]*TimeChart)[vs[1].(int)]
	}).(TimeChartOutput)
}

type TimeChartMapOutput struct{ *pulumi.OutputState }

func (TimeChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TimeChart)(nil)).Elem()
}

func (o TimeChartMapOutput) ToTimeChartMapOutput() TimeChartMapOutput {
	return o
}

func (o TimeChartMapOutput) ToTimeChartMapOutputWithContext(ctx context.Context) TimeChartMapOutput {
	return o
}

func (o TimeChartMapOutput) MapIndex(k pulumi.StringInput) TimeChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TimeChart {
		return vs[0].(map[string]*TimeChart)[vs[1].(string)]
	}).(TimeChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TimeChartInput)(nil)).Elem(), &TimeChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimeChartArrayInput)(nil)).Elem(), TimeChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimeChartMapInput)(nil)).Elem(), TimeChartMap{})
	pulumi.RegisterOutputType(TimeChartOutput{})
	pulumi.RegisterOutputType(TimeChartArrayOutput{})
	pulumi.RegisterOutputType(TimeChartMapOutput{})
}
