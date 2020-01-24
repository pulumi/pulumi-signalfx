// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This chart type displays current data values in a list format.
// 
// The name of each value in the chart reflects the name of the plot and any associated dimensions. We recommend you click the Pencil icon and give the plot a meaningful name, as in plot B below. Otherwise, just the raw metric name will be displayed on the chart, as in plot A below.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/list_chart.html.markdown.
type ListChart struct {
	s *pulumi.ResourceState
}

// NewListChart registers a new resource with the given unique name, arguments, and options.
func NewListChart(ctx *pulumi.Context,
	name string, args *ListChartArgs, opts ...pulumi.ResourceOpt) (*ListChart, error) {
	if args == nil || args.ProgramText == nil {
		return nil, errors.New("missing required argument 'ProgramText'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["colorBy"] = nil
		inputs["colorScales"] = nil
		inputs["description"] = nil
		inputs["disableSampling"] = nil
		inputs["endTime"] = nil
		inputs["legendFieldsToHides"] = nil
		inputs["legendOptionsFields"] = nil
		inputs["maxDelay"] = nil
		inputs["maxPrecision"] = nil
		inputs["name"] = nil
		inputs["programText"] = nil
		inputs["refreshInterval"] = nil
		inputs["secondaryVisualization"] = nil
		inputs["sortBy"] = nil
		inputs["startTime"] = nil
		inputs["timeRange"] = nil
		inputs["unitPrefix"] = nil
		inputs["vizOptions"] = nil
	} else {
		inputs["colorBy"] = args.ColorBy
		inputs["colorScales"] = args.ColorScales
		inputs["description"] = args.Description
		inputs["disableSampling"] = args.DisableSampling
		inputs["endTime"] = args.EndTime
		inputs["legendFieldsToHides"] = args.LegendFieldsToHides
		inputs["legendOptionsFields"] = args.LegendOptionsFields
		inputs["maxDelay"] = args.MaxDelay
		inputs["maxPrecision"] = args.MaxPrecision
		inputs["name"] = args.Name
		inputs["programText"] = args.ProgramText
		inputs["refreshInterval"] = args.RefreshInterval
		inputs["secondaryVisualization"] = args.SecondaryVisualization
		inputs["sortBy"] = args.SortBy
		inputs["startTime"] = args.StartTime
		inputs["timeRange"] = args.TimeRange
		inputs["unitPrefix"] = args.UnitPrefix
		inputs["vizOptions"] = args.VizOptions
	}
	inputs["url"] = nil
	s, err := ctx.RegisterResource("signalfx:index/listChart:ListChart", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ListChart{s: s}, nil
}

// GetListChart gets an existing ListChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListChart(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ListChartState, opts ...pulumi.ResourceOpt) (*ListChart, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["colorBy"] = state.ColorBy
		inputs["colorScales"] = state.ColorScales
		inputs["description"] = state.Description
		inputs["disableSampling"] = state.DisableSampling
		inputs["endTime"] = state.EndTime
		inputs["legendFieldsToHides"] = state.LegendFieldsToHides
		inputs["legendOptionsFields"] = state.LegendOptionsFields
		inputs["maxDelay"] = state.MaxDelay
		inputs["maxPrecision"] = state.MaxPrecision
		inputs["name"] = state.Name
		inputs["programText"] = state.ProgramText
		inputs["refreshInterval"] = state.RefreshInterval
		inputs["secondaryVisualization"] = state.SecondaryVisualization
		inputs["sortBy"] = state.SortBy
		inputs["startTime"] = state.StartTime
		inputs["timeRange"] = state.TimeRange
		inputs["unitPrefix"] = state.UnitPrefix
		inputs["url"] = state.Url
		inputs["vizOptions"] = state.VizOptions
	}
	s, err := ctx.ReadResource("signalfx:index/listChart:ListChart", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ListChart{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ListChart) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ListChart) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
func (r *ListChart) ColorBy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["colorBy"])
}

// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
func (r *ListChart) ColorScales() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["colorScales"])
}

// Description of the chart.
func (r *ListChart) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
func (r *ListChart) DisableSampling() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["disableSampling"])
}

// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
func (r *ListChart) EndTime() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["endTime"])
}

// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
func (r *ListChart) LegendFieldsToHides() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["legendFieldsToHides"])
}

// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
// * `property` The name of the property to display. Note the special values of `plotLabel` (corresponding with the API's `sfMetric`) which shows the label of the time series `publish()` and `metric` (corresponding with the API's `sf_originatingMetric`) that shows the name of the metric for the time series being displayed.
// * `enabled` True or False depending on if you want the property to be shown or hidden.
func (r *ListChart) LegendOptionsFields() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["legendOptionsFields"])
}

// How long (in seconds) to wait for late datapoints.
func (r *ListChart) MaxDelay() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxDelay"])
}

// Maximum number of digits to display when rounding values up or down.
func (r *ListChart) MaxPrecision() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["maxPrecision"])
}

// Name of the chart.
func (r *ListChart) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
func (r *ListChart) ProgramText() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["programText"])
}

// How often (in seconds) to refresh the values of the list.
func (r *ListChart) RefreshInterval() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["refreshInterval"])
}

// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
func (r *ListChart) SecondaryVisualization() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["secondaryVisualization"])
}

// The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
func (r *ListChart) SortBy() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sortBy"])
}

// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
func (r *ListChart) StartTime() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["startTime"])
}

// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
func (r *ListChart) TimeRange() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["timeRange"])
}

// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
func (r *ListChart) UnitPrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["unitPrefix"])
}

// URL of the chart
func (r *ListChart) Url() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["url"])
}

// Plot-level customization options, associated with a publish statement.
func (r *ListChart) VizOptions() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["vizOptions"])
}

// Input properties used for looking up and filtering ListChart resources.
type ListChartState struct {
	// Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy interface{}
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales interface{}
	// Description of the chart.
	Description interface{}
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling interface{}
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime interface{}
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	LegendFieldsToHides interface{}
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	// * `property` The name of the property to display. Note the special values of `plotLabel` (corresponding with the API's `sfMetric`) which shows the label of the time series `publish()` and `metric` (corresponding with the API's `sf_originatingMetric`) that shows the name of the metric for the time series being displayed.
	// * `enabled` True or False depending on if you want the property to be shown or hidden.
	LegendOptionsFields interface{}
	// How long (in seconds) to wait for late datapoints.
	MaxDelay interface{}
	// Maximum number of digits to display when rounding values up or down.
	MaxPrecision interface{}
	// Name of the chart.
	Name interface{}
	// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
	ProgramText interface{}
	// How often (in seconds) to refresh the values of the list.
	RefreshInterval interface{}
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
	SecondaryVisualization interface{}
	// The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
	SortBy interface{}
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime interface{}
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange interface{}
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix interface{}
	// URL of the chart
	Url interface{}
	// Plot-level customization options, associated with a publish statement.
	VizOptions interface{}
}

// The set of arguments for constructing a ListChart resource.
type ListChartArgs struct {
	// Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy interface{}
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales interface{}
	// Description of the chart.
	Description interface{}
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling interface{}
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime interface{}
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	LegendFieldsToHides interface{}
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	// * `property` The name of the property to display. Note the special values of `plotLabel` (corresponding with the API's `sfMetric`) which shows the label of the time series `publish()` and `metric` (corresponding with the API's `sf_originatingMetric`) that shows the name of the metric for the time series being displayed.
	// * `enabled` True or False depending on if you want the property to be shown or hidden.
	LegendOptionsFields interface{}
	// How long (in seconds) to wait for late datapoints.
	MaxDelay interface{}
	// Maximum number of digits to display when rounding values up or down.
	MaxPrecision interface{}
	// Name of the chart.
	Name interface{}
	// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
	ProgramText interface{}
	// How often (in seconds) to refresh the values of the list.
	RefreshInterval interface{}
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
	SecondaryVisualization interface{}
	// The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
	SortBy interface{}
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime interface{}
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange interface{}
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix interface{}
	// Plot-level customization options, associated with a publish statement.
	VizOptions interface{}
}
