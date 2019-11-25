// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a SignalFx time chart resource. This can be used to create and manage the different types of time charts.
// 
// Time charts display data points over a period of time.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/time_chart.html.markdown.
type TimeChart struct {
	s *pulumi.ResourceState
}

// NewTimeChart registers a new resource with the given unique name, arguments, and options.
func NewTimeChart(ctx *pulumi.Context,
	name string, args *TimeChartArgs, opts ...pulumi.ResourceOpt) (*TimeChart, error) {
	if args == nil || args.ProgramText == nil {
		return nil, errors.New("missing required argument 'ProgramText'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["axesIncludeZero"] = nil
		inputs["axesPrecision"] = nil
		inputs["axisLeft"] = nil
		inputs["axisRight"] = nil
		inputs["colorBy"] = nil
		inputs["description"] = nil
		inputs["disableSampling"] = nil
		inputs["endTime"] = nil
		inputs["eventOptions"] = nil
		inputs["histogramOptions"] = nil
		inputs["legendFieldsToHides"] = nil
		inputs["legendOptionsFields"] = nil
		inputs["maxDelay"] = nil
		inputs["minimumResolution"] = nil
		inputs["name"] = nil
		inputs["onChartLegendDimension"] = nil
		inputs["plotType"] = nil
		inputs["programText"] = nil
		inputs["showDataMarkers"] = nil
		inputs["showEventLines"] = nil
		inputs["stacked"] = nil
		inputs["startTime"] = nil
		inputs["tags"] = nil
		inputs["timeRange"] = nil
		inputs["timezone"] = nil
		inputs["unitPrefix"] = nil
		inputs["vizOptions"] = nil
	} else {
		inputs["axesIncludeZero"] = args.AxesIncludeZero
		inputs["axesPrecision"] = args.AxesPrecision
		inputs["axisLeft"] = args.AxisLeft
		inputs["axisRight"] = args.AxisRight
		inputs["colorBy"] = args.ColorBy
		inputs["description"] = args.Description
		inputs["disableSampling"] = args.DisableSampling
		inputs["endTime"] = args.EndTime
		inputs["eventOptions"] = args.EventOptions
		inputs["histogramOptions"] = args.HistogramOptions
		inputs["legendFieldsToHides"] = args.LegendFieldsToHides
		inputs["legendOptionsFields"] = args.LegendOptionsFields
		inputs["maxDelay"] = args.MaxDelay
		inputs["minimumResolution"] = args.MinimumResolution
		inputs["name"] = args.Name
		inputs["onChartLegendDimension"] = args.OnChartLegendDimension
		inputs["plotType"] = args.PlotType
		inputs["programText"] = args.ProgramText
		inputs["showDataMarkers"] = args.ShowDataMarkers
		inputs["showEventLines"] = args.ShowEventLines
		inputs["stacked"] = args.Stacked
		inputs["startTime"] = args.StartTime
		inputs["tags"] = args.Tags
		inputs["timeRange"] = args.TimeRange
		inputs["timezone"] = args.Timezone
		inputs["unitPrefix"] = args.UnitPrefix
		inputs["vizOptions"] = args.VizOptions
	}
	inputs["url"] = nil
	s, err := ctx.RegisterResource("signalfx:index/timeChart:TimeChart", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TimeChart{s: s}, nil
}

// GetTimeChart gets an existing TimeChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTimeChart(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TimeChartState, opts ...pulumi.ResourceOpt) (*TimeChart, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["axesIncludeZero"] = state.AxesIncludeZero
		inputs["axesPrecision"] = state.AxesPrecision
		inputs["axisLeft"] = state.AxisLeft
		inputs["axisRight"] = state.AxisRight
		inputs["colorBy"] = state.ColorBy
		inputs["description"] = state.Description
		inputs["disableSampling"] = state.DisableSampling
		inputs["endTime"] = state.EndTime
		inputs["eventOptions"] = state.EventOptions
		inputs["histogramOptions"] = state.HistogramOptions
		inputs["legendFieldsToHides"] = state.LegendFieldsToHides
		inputs["legendOptionsFields"] = state.LegendOptionsFields
		inputs["maxDelay"] = state.MaxDelay
		inputs["minimumResolution"] = state.MinimumResolution
		inputs["name"] = state.Name
		inputs["onChartLegendDimension"] = state.OnChartLegendDimension
		inputs["plotType"] = state.PlotType
		inputs["programText"] = state.ProgramText
		inputs["showDataMarkers"] = state.ShowDataMarkers
		inputs["showEventLines"] = state.ShowEventLines
		inputs["stacked"] = state.Stacked
		inputs["startTime"] = state.StartTime
		inputs["tags"] = state.Tags
		inputs["timeRange"] = state.TimeRange
		inputs["timezone"] = state.Timezone
		inputs["unitPrefix"] = state.UnitPrefix
		inputs["url"] = state.Url
		inputs["vizOptions"] = state.VizOptions
	}
	s, err := ctx.ReadResource("signalfx:index/timeChart:TimeChart", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TimeChart{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TimeChart) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TimeChart) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Force the chart to display zero on the y-axes, even if none of the data is near zero.
func (r *TimeChart) AxesIncludeZero() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["axesIncludeZero"])
}

// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
func (r *TimeChart) AxesPrecision() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["axesPrecision"])
}

// Set of axis options.
func (r *TimeChart) AxisLeft() pulumi.Output {
	return r.s.State["axisLeft"]
}

// Set of axis options.
func (r *TimeChart) AxisRight() pulumi.Output {
	return r.s.State["axisRight"]
}

// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
func (r *TimeChart) ColorBy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["colorBy"])
}

// Description of the chart.
func (r *TimeChart) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
func (r *TimeChart) DisableSampling() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["disableSampling"])
}

// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
func (r *TimeChart) EndTime() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["endTime"])
}

// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
func (r *TimeChart) EventOptions() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["eventOptions"])
}

// Only used when `plotType` is `"Histogram"`. Histogram specific options.
func (r *TimeChart) HistogramOptions() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["histogramOptions"])
}

// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
func (r *TimeChart) LegendFieldsToHides() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["legendFieldsToHides"])
}

// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
// * `property` The name of the property to display. Note the special values of `plotLabel` (corresponding with the API's `sfMetric`) which shows the label of the time series `publish()` and `metric` (corresponding with the API's `sf_originatingMetric`) that shows the name of the metric for the time series being displayed.
// * `enabled` True or False depending on if you want the property to be shown or hidden.
func (r *TimeChart) LegendOptionsFields() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["legendOptionsFields"])
}

// How long (in seconds) to wait for late datapoints.
func (r *TimeChart) MaxDelay() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxDelay"])
}

// The minimum resolution (in seconds) to use for computing the underlying program.
func (r *TimeChart) MinimumResolution() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["minimumResolution"])
}

// Name of the chart.
func (r *TimeChart) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
func (r *TimeChart) OnChartLegendDimension() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["onChartLegendDimension"])
}

// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
func (r *TimeChart) PlotType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["plotType"])
}

// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
func (r *TimeChart) ProgramText() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["programText"])
}

// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
func (r *TimeChart) ShowDataMarkers() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["showDataMarkers"])
}

// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
func (r *TimeChart) ShowEventLines() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["showEventLines"])
}

// Whether area and bar charts in the visualization should be stacked. `false` by default.
func (r *TimeChart) Stacked() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["stacked"])
}

// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
func (r *TimeChart) StartTime() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["startTime"])
}

// Tags associated with the chart
func (r *TimeChart) Tags() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["tags"])
}

// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
func (r *TimeChart) TimeRange() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["timeRange"])
}

// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
func (r *TimeChart) Timezone() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["timezone"])
}

// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
func (r *TimeChart) UnitPrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["unitPrefix"])
}

// URL of the chart
func (r *TimeChart) Url() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["url"])
}

// Plot-level customization options, associated with a publish statement.
func (r *TimeChart) VizOptions() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["vizOptions"])
}

// Input properties used for looking up and filtering TimeChart resources.
type TimeChartState struct {
	// Force the chart to display zero on the y-axes, even if none of the data is near zero.
	AxesIncludeZero interface{}
	// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
	AxesPrecision interface{}
	// Set of axis options.
	AxisLeft interface{}
	// Set of axis options.
	AxisRight interface{}
	// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy interface{}
	// Description of the chart.
	Description interface{}
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
	DisableSampling interface{}
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime interface{}
	// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
	EventOptions interface{}
	// Only used when `plotType` is `"Histogram"`. Histogram specific options.
	HistogramOptions interface{}
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	LegendFieldsToHides interface{}
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	// * `property` The name of the property to display. Note the special values of `plotLabel` (corresponding with the API's `sfMetric`) which shows the label of the time series `publish()` and `metric` (corresponding with the API's `sf_originatingMetric`) that shows the name of the metric for the time series being displayed.
	// * `enabled` True or False depending on if you want the property to be shown or hidden.
	LegendOptionsFields interface{}
	// How long (in seconds) to wait for late datapoints.
	MaxDelay interface{}
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution interface{}
	// Name of the chart.
	Name interface{}
	// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
	OnChartLegendDimension interface{}
	// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
	PlotType interface{}
	// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
	ProgramText interface{}
	// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
	ShowDataMarkers interface{}
	// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
	ShowEventLines interface{}
	// Whether area and bar charts in the visualization should be stacked. `false` by default.
	Stacked interface{}
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime interface{}
	// Tags associated with the chart
	Tags interface{}
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange interface{}
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
	Timezone interface{}
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix interface{}
	// URL of the chart
	Url interface{}
	// Plot-level customization options, associated with a publish statement.
	VizOptions interface{}
}

// The set of arguments for constructing a TimeChart resource.
type TimeChartArgs struct {
	// Force the chart to display zero on the y-axes, even if none of the data is near zero.
	AxesIncludeZero interface{}
	// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
	AxesPrecision interface{}
	// Set of axis options.
	AxisLeft interface{}
	// Set of axis options.
	AxisRight interface{}
	// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy interface{}
	// Description of the chart.
	Description interface{}
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
	DisableSampling interface{}
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime interface{}
	// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
	EventOptions interface{}
	// Only used when `plotType` is `"Histogram"`. Histogram specific options.
	HistogramOptions interface{}
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	LegendFieldsToHides interface{}
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	// * `property` The name of the property to display. Note the special values of `plotLabel` (corresponding with the API's `sfMetric`) which shows the label of the time series `publish()` and `metric` (corresponding with the API's `sf_originatingMetric`) that shows the name of the metric for the time series being displayed.
	// * `enabled` True or False depending on if you want the property to be shown or hidden.
	LegendOptionsFields interface{}
	// How long (in seconds) to wait for late datapoints.
	MaxDelay interface{}
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution interface{}
	// Name of the chart.
	Name interface{}
	// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
	OnChartLegendDimension interface{}
	// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
	PlotType interface{}
	// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
	ProgramText interface{}
	// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
	ShowDataMarkers interface{}
	// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
	ShowEventLines interface{}
	// Whether area and bar charts in the visualization should be stacked. `false` by default.
	Stacked interface{}
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime interface{}
	// Tags associated with the chart
	Tags interface{}
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange interface{}
	// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
	Timezone interface{}
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix interface{}
	// Plot-level customization options, associated with a publish statement.
	VizOptions interface{}
}
