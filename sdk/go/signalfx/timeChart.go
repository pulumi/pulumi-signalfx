// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a SignalFx time chart resource. This can be used to create and manage the different types of time charts.
//
// Time charts display data points over a period of time.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-signalfx/sdk/v2/go/signalfx"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := signalfx.NewTimeChart(ctx, "mychart0", &signalfx.TimeChartArgs{
// 			AxisLeft: &signalfx.TimeChartAxisLeftArgs{
// 				Label:        pulumi.String("CPU Total Idle"),
// 				LowWatermark: pulumi.Float64(1000),
// 			},
// 			LegendOptionsFields: signalfx.TimeChartLegendOptionsFieldArray{
// 				&signalfx.TimeChartLegendOptionsFieldArgs{
// 					Enabled:  pulumi.Bool(false),
// 					Property: pulumi.String("collector"),
// 				},
// 				&signalfx.TimeChartLegendOptionsFieldArgs{
// 					Enabled:  pulumi.Bool(false),
// 					Property: pulumi.String("hostname"),
// 				},
// 			},
// 			PlotType:        pulumi.String("LineChart"),
// 			ProgramText:     pulumi.String(fmt.Sprintf("%v%v", "data(\"cpu.total.idle\").publish(label=\"CPU Idle\")\n", "\n")),
// 			ShowDataMarkers: pulumi.Bool(true),
// 			TimeRange:       pulumi.Int(3600),
// 			VizOptions: signalfx.TimeChartVizOptionArray{
// 				&signalfx.TimeChartVizOptionArgs{
// 					Axis:  pulumi.String("left"),
// 					Color: pulumi.String("orange"),
// 					Label: pulumi.String("CPU Idle"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type TimeChart struct {
	pulumi.CustomResourceState

	// Force the chart to display zero on the y-axes, even if none of the data is near zero.
	AxesIncludeZero pulumi.BoolPtrOutput `pulumi:"axesIncludeZero"`
	// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
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
	// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
	PlotType pulumi.StringPtrOutput `pulumi:"plotType"`
	// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
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
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
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
	if args == nil || args.ProgramText == nil {
		return nil, errors.New("missing required argument 'ProgramText'")
	}
	if args == nil {
		args = &TimeChartArgs{}
	}
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
	// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
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
	// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
	PlotType *string `pulumi:"plotType"`
	// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
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
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
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
	// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
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
	// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
	PlotType pulumi.StringPtrInput
	// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
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
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
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
	// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
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
	// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
	PlotType *string `pulumi:"plotType"`
	// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
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
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
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
	// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
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
	// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
	PlotType pulumi.StringPtrInput
	// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
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
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
	Timezone pulumi.StringPtrInput
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement.
	VizOptions TimeChartVizOptionArrayInput
}

func (TimeChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*timeChartArgs)(nil)).Elem()
}
