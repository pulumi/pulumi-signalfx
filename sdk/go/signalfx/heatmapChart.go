// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This chart type displays the specified plot in a heatmap fashion. This format is similar to the [Infrastructure Navigator](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/built-in-content/infra-nav.html#infra), with squares representing each source for the selected metric, and the color of each square representing the value range of the metric.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-signalfx/sdk/v3/go/signalfx"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := signalfx.NewHeatmapChart(ctx, "myheatmapchart0", &signalfx.HeatmapChartArgs{
// 			ColorRange: &signalfx.HeatmapChartColorRangeArgs{
// 				Color:    pulumi.String("#ff0000"),
// 				MaxValue: pulumi.Float64(100),
// 				MinValue: pulumi.Float64(0),
// 			},
// 			ColorScales: signalfx.HeatmapChartColorScaleArray{
// 				&signalfx.HeatmapChartColorScaleArgs{
// 					Color: pulumi.String("green"),
// 					Gte:   pulumi.Float64(99),
// 				},
// 				&signalfx.HeatmapChartColorScaleArgs{
// 					Color: pulumi.String("yellow"),
// 					Gte:   pulumi.Float64(95),
// 					Lt:    pulumi.Float64(99),
// 				},
// 				&signalfx.HeatmapChartColorScaleArgs{
// 					Color: pulumi.String("red"),
// 					Lt:    pulumi.Float64(95),
// 				},
// 			},
// 			Description:     pulumi.String("Very cool Heatmap"),
// 			DisableSampling: pulumi.Bool(true),
// 			GroupBies: pulumi.StringArray{
// 				pulumi.String("hostname"),
// 				pulumi.String("host"),
// 			},
// 			HideTimestamp: pulumi.Bool(true),
// 			ProgramText:   pulumi.String(fmt.Sprintf("%v%v%v", "myfilters = filter(\"cluster_name\", \"prod\") and filter(\"role\", \"search\")\n", "data(\"cpu.total.idle\", filter=myfilters).publish()\n", "\n")),
// 			SortBy:        pulumi.String("+host"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type HeatmapChart struct {
	pulumi.CustomResourceState

	// Values and color for the color range. Example: `colorRange : { min : 0, max : 100, color : "#0000ff" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorRange HeatmapChartColorRangePtrOutput `pulumi:"colorRange"`
	// One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `colorScale { gt = 60, color = "blue" } colorScale { lte = 60, color = "yellow" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales HeatmapChartColorScaleArrayOutput `pulumi:"colorScales"`
	// Description of the chart.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling pulumi.BoolPtrOutput `pulumi:"disableSampling"`
	// Properties to group by in the heatmap (in nesting order).
	GroupBies pulumi.StringArrayOutput `pulumi:"groupBies"`
	// Whether to show the timestamp in the chart. `false` by default.
	HideTimestamp pulumi.BoolPtrOutput `pulumi:"hideTimestamp"`
	// How long (in seconds) to wait for late datapoints.
	MaxDelay pulumi.IntPtrOutput `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution pulumi.IntPtrOutput `pulumi:"minimumResolution"`
	// Name of the chart.
	Name pulumi.StringOutput `pulumi:"name"`
	// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
	ProgramText pulumi.StringOutput `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the heatmap.
	RefreshInterval pulumi.IntPtrOutput `pulumi:"refreshInterval"`
	// The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
	SortBy pulumi.StringPtrOutput `pulumi:"sortBy"`
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix pulumi.StringPtrOutput `pulumi:"unitPrefix"`
	// The URL of the chart.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewHeatmapChart registers a new resource with the given unique name, arguments, and options.
func NewHeatmapChart(ctx *pulumi.Context,
	name string, args *HeatmapChartArgs, opts ...pulumi.ResourceOption) (*HeatmapChart, error) {
	if args == nil || args.ProgramText == nil {
		return nil, errors.New("missing required argument 'ProgramText'")
	}
	if args == nil {
		args = &HeatmapChartArgs{}
	}
	var resource HeatmapChart
	err := ctx.RegisterResource("signalfx:index/heatmapChart:HeatmapChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHeatmapChart gets an existing HeatmapChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHeatmapChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HeatmapChartState, opts ...pulumi.ResourceOption) (*HeatmapChart, error) {
	var resource HeatmapChart
	err := ctx.ReadResource("signalfx:index/heatmapChart:HeatmapChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HeatmapChart resources.
type heatmapChartState struct {
	// Values and color for the color range. Example: `colorRange : { min : 0, max : 100, color : "#0000ff" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorRange *HeatmapChartColorRange `pulumi:"colorRange"`
	// One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `colorScale { gt = 60, color = "blue" } colorScale { lte = 60, color = "yellow" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales []HeatmapChartColorScale `pulumi:"colorScales"`
	// Description of the chart.
	Description *string `pulumi:"description"`
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling *bool `pulumi:"disableSampling"`
	// Properties to group by in the heatmap (in nesting order).
	GroupBies []string `pulumi:"groupBies"`
	// Whether to show the timestamp in the chart. `false` by default.
	HideTimestamp *bool `pulumi:"hideTimestamp"`
	// How long (in seconds) to wait for late datapoints.
	MaxDelay *int `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution *int `pulumi:"minimumResolution"`
	// Name of the chart.
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
	ProgramText *string `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the heatmap.
	RefreshInterval *int `pulumi:"refreshInterval"`
	// The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
	SortBy *string `pulumi:"sortBy"`
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix *string `pulumi:"unitPrefix"`
	// The URL of the chart.
	Url *string `pulumi:"url"`
}

type HeatmapChartState struct {
	// Values and color for the color range. Example: `colorRange : { min : 0, max : 100, color : "#0000ff" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorRange HeatmapChartColorRangePtrInput
	// One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `colorScale { gt = 60, color = "blue" } colorScale { lte = 60, color = "yellow" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales HeatmapChartColorScaleArrayInput
	// Description of the chart.
	Description pulumi.StringPtrInput
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling pulumi.BoolPtrInput
	// Properties to group by in the heatmap (in nesting order).
	GroupBies pulumi.StringArrayInput
	// Whether to show the timestamp in the chart. `false` by default.
	HideTimestamp pulumi.BoolPtrInput
	// How long (in seconds) to wait for late datapoints.
	MaxDelay pulumi.IntPtrInput
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution pulumi.IntPtrInput
	// Name of the chart.
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
	ProgramText pulumi.StringPtrInput
	// How often (in seconds) to refresh the values of the heatmap.
	RefreshInterval pulumi.IntPtrInput
	// The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
	SortBy pulumi.StringPtrInput
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix pulumi.StringPtrInput
	// The URL of the chart.
	Url pulumi.StringPtrInput
}

func (HeatmapChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*heatmapChartState)(nil)).Elem()
}

type heatmapChartArgs struct {
	// Values and color for the color range. Example: `colorRange : { min : 0, max : 100, color : "#0000ff" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorRange *HeatmapChartColorRange `pulumi:"colorRange"`
	// One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `colorScale { gt = 60, color = "blue" } colorScale { lte = 60, color = "yellow" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales []HeatmapChartColorScale `pulumi:"colorScales"`
	// Description of the chart.
	Description *string `pulumi:"description"`
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling *bool `pulumi:"disableSampling"`
	// Properties to group by in the heatmap (in nesting order).
	GroupBies []string `pulumi:"groupBies"`
	// Whether to show the timestamp in the chart. `false` by default.
	HideTimestamp *bool `pulumi:"hideTimestamp"`
	// How long (in seconds) to wait for late datapoints.
	MaxDelay *int `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution *int `pulumi:"minimumResolution"`
	// Name of the chart.
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
	ProgramText string `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the heatmap.
	RefreshInterval *int `pulumi:"refreshInterval"`
	// The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
	SortBy *string `pulumi:"sortBy"`
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix *string `pulumi:"unitPrefix"`
}

// The set of arguments for constructing a HeatmapChart resource.
type HeatmapChartArgs struct {
	// Values and color for the color range. Example: `colorRange : { min : 0, max : 100, color : "#0000ff" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorRange HeatmapChartColorRangePtrInput
	// One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `colorScale { gt = 60, color = "blue" } colorScale { lte = 60, color = "yellow" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales HeatmapChartColorScaleArrayInput
	// Description of the chart.
	Description pulumi.StringPtrInput
	// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
	DisableSampling pulumi.BoolPtrInput
	// Properties to group by in the heatmap (in nesting order).
	GroupBies pulumi.StringArrayInput
	// Whether to show the timestamp in the chart. `false` by default.
	HideTimestamp pulumi.BoolPtrInput
	// How long (in seconds) to wait for late datapoints.
	MaxDelay pulumi.IntPtrInput
	// The minimum resolution (in seconds) to use for computing the underlying program.
	MinimumResolution pulumi.IntPtrInput
	// Name of the chart.
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
	ProgramText pulumi.StringInput
	// How often (in seconds) to refresh the values of the heatmap.
	RefreshInterval pulumi.IntPtrInput
	// The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
	SortBy pulumi.StringPtrInput
	// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
	UnitPrefix pulumi.StringPtrInput
}

func (HeatmapChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*heatmapChartArgs)(nil)).Elem()
}

type HeatmapChartInput interface {
	pulumi.Input

	ToHeatmapChartOutput() HeatmapChartOutput
	ToHeatmapChartOutputWithContext(ctx context.Context) HeatmapChartOutput
}

func (HeatmapChart) ElementType() reflect.Type {
	return reflect.TypeOf((*HeatmapChart)(nil)).Elem()
}

func (i HeatmapChart) ToHeatmapChartOutput() HeatmapChartOutput {
	return i.ToHeatmapChartOutputWithContext(context.Background())
}

func (i HeatmapChart) ToHeatmapChartOutputWithContext(ctx context.Context) HeatmapChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeatmapChartOutput)
}

type HeatmapChartOutput struct {
	*pulumi.OutputState
}

func (HeatmapChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HeatmapChartOutput)(nil)).Elem()
}

func (o HeatmapChartOutput) ToHeatmapChartOutput() HeatmapChartOutput {
	return o
}

func (o HeatmapChartOutput) ToHeatmapChartOutputWithContext(ctx context.Context) HeatmapChartOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HeatmapChartOutput{})
}
