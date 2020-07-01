// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This chart type displays current data values in a list format.
//
// The name of each value in the chart reflects the name of the plot and any associated dimensions. We recommend you click the Pencil icon and give the plot a meaningful name, as in plot B below. Otherwise, just the raw metric name will be displayed on the chart, as in plot A below.
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
// 		_, err := signalfx.NewListChart(ctx, "mylistchart0", &signalfx.ListChartArgs{
// 			ColorBy:         pulumi.String("Metric"),
// 			Description:     pulumi.String("Very cool List Chart"),
// 			DisableSampling: pulumi.Bool(true),
// 			LegendOptionsFields: signalfx.ListChartLegendOptionsFieldArray{
// 				&signalfx.ListChartLegendOptionsFieldArgs{
// 					Enabled:  pulumi.Bool(false),
// 					Property: pulumi.String("collector"),
// 				},
// 				&signalfx.ListChartLegendOptionsFieldArgs{
// 					Enabled:  pulumi.Bool(true),
// 					Property: pulumi.String("cluster_name"),
// 				},
// 				&signalfx.ListChartLegendOptionsFieldArgs{
// 					Enabled:  pulumi.Bool(true),
// 					Property: pulumi.String("role"),
// 				},
// 				&signalfx.ListChartLegendOptionsFieldArgs{
// 					Enabled:  pulumi.Bool(false),
// 					Property: pulumi.String("collector"),
// 				},
// 				&signalfx.ListChartLegendOptionsFieldArgs{
// 					Enabled:  pulumi.Bool(false),
// 					Property: pulumi.String("host"),
// 				},
// 			},
// 			MaxDelay:        pulumi.Int(2),
// 			MaxPrecision:    pulumi.Int(2),
// 			ProgramText:     pulumi.String(fmt.Sprintf("%v%v%v", "myfilters = filter(\"cluster_name\", \"prod\") and filter(\"role\", \"search\")\n", "data(\"cpu.total.idle\", filter=myfilters).publish()\n", "\n")),
// 			RefreshInterval: pulumi.Int(1),
// 			SortBy:          pulumi.String("-value"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ListChart struct {
	pulumi.CustomResourceState

	// Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy pulumi.StringPtrOutput `pulumi:"colorBy"`
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales ListChartColorScaleArrayOutput `pulumi:"colorScales"`
	// Description of the chart.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling pulumi.BoolPtrOutput `pulumi:"disableSampling"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrOutput `pulumi:"endTime"`
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	//
	// Deprecated: Please use legend_options_fields
	LegendFieldsToHides pulumi.StringArrayOutput `pulumi:"legendFieldsToHides"`
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	LegendOptionsFields ListChartLegendOptionsFieldArrayOutput `pulumi:"legendOptionsFields"`
	// How long (in seconds) to wait for late datapoints.
	MaxDelay pulumi.IntPtrOutput `pulumi:"maxDelay"`
	// Maximum number of digits to display when rounding values up or down.
	MaxPrecision pulumi.IntPtrOutput `pulumi:"maxPrecision"`
	// Name of the chart.
	Name pulumi.StringOutput `pulumi:"name"`
	// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringOutput `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the list.
	RefreshInterval pulumi.IntPtrOutput `pulumi:"refreshInterval"`
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
	SecondaryVisualization pulumi.StringPtrOutput `pulumi:"secondaryVisualization"`
	// The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
	SortBy pulumi.StringPtrOutput `pulumi:"sortBy"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrOutput `pulumi:"startTime"`
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrOutput `pulumi:"timeRange"`
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix pulumi.StringPtrOutput `pulumi:"unitPrefix"`
	// URL of the chart
	Url pulumi.StringOutput `pulumi:"url"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions ListChartVizOptionArrayOutput `pulumi:"vizOptions"`
}

// NewListChart registers a new resource with the given unique name, arguments, and options.
func NewListChart(ctx *pulumi.Context,
	name string, args *ListChartArgs, opts ...pulumi.ResourceOption) (*ListChart, error) {
	if args == nil || args.ProgramText == nil {
		return nil, errors.New("missing required argument 'ProgramText'")
	}
	if args == nil {
		args = &ListChartArgs{}
	}
	var resource ListChart
	err := ctx.RegisterResource("signalfx:index/listChart:ListChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListChart gets an existing ListChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListChartState, opts ...pulumi.ResourceOption) (*ListChart, error) {
	var resource ListChart
	err := ctx.ReadResource("signalfx:index/listChart:ListChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListChart resources.
type listChartState struct {
	// Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy *string `pulumi:"colorBy"`
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales []ListChartColorScale `pulumi:"colorScales"`
	// Description of the chart.
	Description *string `pulumi:"description"`
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling *bool `pulumi:"disableSampling"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime *int `pulumi:"endTime"`
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	//
	// Deprecated: Please use legend_options_fields
	LegendFieldsToHides []string `pulumi:"legendFieldsToHides"`
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	LegendOptionsFields []ListChartLegendOptionsField `pulumi:"legendOptionsFields"`
	// How long (in seconds) to wait for late datapoints.
	MaxDelay *int `pulumi:"maxDelay"`
	// Maximum number of digits to display when rounding values up or down.
	MaxPrecision *int `pulumi:"maxPrecision"`
	// Name of the chart.
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText *string `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the list.
	RefreshInterval *int `pulumi:"refreshInterval"`
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
	SecondaryVisualization *string `pulumi:"secondaryVisualization"`
	// The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
	SortBy *string `pulumi:"sortBy"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime *int `pulumi:"startTime"`
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange *int `pulumi:"timeRange"`
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix *string `pulumi:"unitPrefix"`
	// URL of the chart
	Url *string `pulumi:"url"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions []ListChartVizOption `pulumi:"vizOptions"`
}

type ListChartState struct {
	// Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy pulumi.StringPtrInput
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales ListChartColorScaleArrayInput
	// Description of the chart.
	Description pulumi.StringPtrInput
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling pulumi.BoolPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrInput
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	//
	// Deprecated: Please use legend_options_fields
	LegendFieldsToHides pulumi.StringArrayInput
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	LegendOptionsFields ListChartLegendOptionsFieldArrayInput
	// How long (in seconds) to wait for late datapoints.
	MaxDelay pulumi.IntPtrInput
	// Maximum number of digits to display when rounding values up or down.
	MaxPrecision pulumi.IntPtrInput
	// Name of the chart.
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringPtrInput
	// How often (in seconds) to refresh the values of the list.
	RefreshInterval pulumi.IntPtrInput
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
	SecondaryVisualization pulumi.StringPtrInput
	// The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
	SortBy pulumi.StringPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrInput
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrInput
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix pulumi.StringPtrInput
	// URL of the chart
	Url pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement.
	VizOptions ListChartVizOptionArrayInput
}

func (ListChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*listChartState)(nil)).Elem()
}

type listChartArgs struct {
	// Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy *string `pulumi:"colorBy"`
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales []ListChartColorScale `pulumi:"colorScales"`
	// Description of the chart.
	Description *string `pulumi:"description"`
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling *bool `pulumi:"disableSampling"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime *int `pulumi:"endTime"`
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	//
	// Deprecated: Please use legend_options_fields
	LegendFieldsToHides []string `pulumi:"legendFieldsToHides"`
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	LegendOptionsFields []ListChartLegendOptionsField `pulumi:"legendOptionsFields"`
	// How long (in seconds) to wait for late datapoints.
	MaxDelay *int `pulumi:"maxDelay"`
	// Maximum number of digits to display when rounding values up or down.
	MaxPrecision *int `pulumi:"maxPrecision"`
	// Name of the chart.
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText string `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the list.
	RefreshInterval *int `pulumi:"refreshInterval"`
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
	SecondaryVisualization *string `pulumi:"secondaryVisualization"`
	// The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
	SortBy *string `pulumi:"sortBy"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime *int `pulumi:"startTime"`
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange *int `pulumi:"timeRange"`
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix *string `pulumi:"unitPrefix"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions []ListChartVizOption `pulumi:"vizOptions"`
}

// The set of arguments for constructing a ListChart resource.
type ListChartArgs struct {
	// Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
	ColorBy pulumi.StringPtrInput
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales ListChartColorScaleArrayInput
	// Description of the chart.
	Description pulumi.StringPtrInput
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling pulumi.BoolPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrInput
	// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
	//
	// Deprecated: Please use legend_options_fields
	LegendFieldsToHides pulumi.StringArrayInput
	// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
	LegendOptionsFields ListChartLegendOptionsFieldArrayInput
	// How long (in seconds) to wait for late datapoints.
	MaxDelay pulumi.IntPtrInput
	// Maximum number of digits to display when rounding values up or down.
	MaxPrecision pulumi.IntPtrInput
	// Name of the chart.
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringInput
	// How often (in seconds) to refresh the values of the list.
	RefreshInterval pulumi.IntPtrInput
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
	SecondaryVisualization pulumi.StringPtrInput
	// The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
	SortBy pulumi.StringPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrInput
	// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrInput
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement.
	VizOptions ListChartVizOptionArrayInput
}

func (ListChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listChartArgs)(nil)).Elem()
}
