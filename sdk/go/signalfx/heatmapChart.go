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

// This chart type shows the specified plot in a heat map fashion. This format is similar to the [Infrastructure Navigator](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/built-in-content/infra-nav.html#infra), with squares representing each source for the selected metric, and the color of each square representing the value range of the metric.
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
//			_, err := signalfx.NewHeatmapChart(ctx, "myheatmapchart0", &signalfx.HeatmapChartArgs{
//				ColorRange: &signalfx.HeatmapChartColorRangeArgs{
//					Color:    pulumi.String("#ff0000"),
//					MaxValue: pulumi.Float64(100),
//					MinValue: pulumi.Float64(0),
//				},
//				ColorScales: signalfx.HeatmapChartColorScaleArray{
//					&signalfx.HeatmapChartColorScaleArgs{
//						Color: pulumi.String("green"),
//						Gte:   pulumi.Float64(99),
//					},
//					&signalfx.HeatmapChartColorScaleArgs{
//						Color: pulumi.String("yellow"),
//						Gte:   pulumi.Float64(95),
//						Lt:    pulumi.Float64(99),
//					},
//					&signalfx.HeatmapChartColorScaleArgs{
//						Color: pulumi.String("red"),
//						Lt:    pulumi.Float64(95),
//					},
//				},
//				Description:     pulumi.String("Very cool Heatmap"),
//				DisableSampling: pulumi.Bool(true),
//				GroupBies: pulumi.StringArray{
//					pulumi.String("hostname"),
//					pulumi.String("host"),
//				},
//				HideTimestamp: pulumi.Bool(true),
//				ProgramText:   pulumi.String("myfilters = filter(\"cluster_name\", \"prod\") and filter(\"role\", \"search\")\ndata(\"cpu.total.idle\", filter=myfilters).publish()\n\n"),
//				SortBy:        pulumi.String("+host"),
//				Timezone:      pulumi.String("Europe/Paris"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Arguments
//
// The following arguments are supported in the resource block:
//
// * `name` - (Required) Name of the chart.
// * `programText` - (Required) Signalflow program text for the chart. More info at <https://dev.splunk.com/observability/docs/signalflow/>.
// * `description` - (Optional) Description of the chart.
// * `unitPrefix` - (Optional) Must be `"Metric"` or `"Binary`". `"Metric"` by default.
// * `minimumResolution` - (Optional) The minimum resolution (in seconds) to use for computing the underlying program.
// * `maxDelay` - (Optional) How long (in seconds) to wait for late datapoints.
// * `timezone` - (Optional) The property value is a string that denotes the geographic region associated with the time zone, (default UTC).
// * `refreshInterval` - (Optional) How often (in seconds) to refresh the values of the heatmap.
// * `disableSampling` - (Optional) If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
// * `groupBy` - (Optional) Properties to group by in the heatmap (in nesting order).
// * `sortBy` - (Optional) The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
// * `hideTimestamp` - (Optional) Whether to show the timestamp in the chart. `false` by default.
// * `colorRange` - (Optional, Default) Values and color for the color range. Example: `colorRange : { min : 0, max : 100, color : "#0000ff" }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
//   - `minValue` - (Optional) The minimum value within the coloring range.
//   - `maxValue` - (Optional) The maximum value within the coloring range.
//   - `color` - (Required) The color range to use. The starting hex color value for data values in a heatmap chart. Specify the value as a 6-character hexadecimal value preceded by the '#' character, for example "#ea1849" (grass green).
//
// * `colorScale` - (Optional.  Conflicts with `colorRange`) One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `colorScale { gt = 60, color = "blue" } colorScale { lte = 60, color = "yellow" }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
//   - `gt` - (Optional) Indicates the lower threshold non-inclusive value for this range.
//   - `gte` - (Optional) Indicates the lower threshold inclusive value for this range.
//   - `lt` - (Optional) Indicates the upper threshold non-inclusive value for this range.
//   - `lte` - (Optional) Indicates the upper threshold inclusive value for this range.
//   - `color` - (Required) The color range to use. Hex values are not supported here. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
//
// ## Attributes
//
// In a addition to all arguments above, the following attributes are exported:
//
// * `id` - The ID of the chart.
// * `url` - The URL of the chart.
type HeatmapChart struct {
	pulumi.CustomResourceState

	// Values and color for the color range. Example: colorRange : { min : 0, max : 100, color : "#0000ff" }
	ColorRange HeatmapChartColorRangePtrOutput `pulumi:"colorRange"`
	// Single color range including both the color to display for that range and the borders of the range
	ColorScales HeatmapChartColorScaleArrayOutput `pulumi:"colorScales"`
	// Description of the chart (Optional)
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// (false by default) If false, samples a subset of the output MTS, which improves UI performance
	DisableSampling pulumi.BoolPtrOutput `pulumi:"disableSampling"`
	// Properties to group by in the heatmap (in nesting order)
	GroupBies pulumi.StringArrayOutput `pulumi:"groupBies"`
	// (false by default) Whether to show the timestamp in the chart
	HideTimestamp pulumi.BoolPtrOutput `pulumi:"hideTimestamp"`
	// How long (in seconds) to wait for late datapoints
	MaxDelay pulumi.IntPtrOutput `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program
	MinimumResolution pulumi.IntPtrOutput `pulumi:"minimumResolution"`
	// Name of the chart
	Name pulumi.StringOutput `pulumi:"name"`
	// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
	ProgramText pulumi.StringOutput `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the heatmap
	RefreshInterval pulumi.IntPtrOutput `pulumi:"refreshInterval"`
	// The property to use when sorting the elements. Must be prepended with + for ascending or - for descending (e.g. -foo)
	SortBy pulumi.StringPtrOutput `pulumi:"sortBy"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// (Metric by default) Must be "Metric" or "Binary"
	UnitPrefix pulumi.StringPtrOutput `pulumi:"unitPrefix"`
	// URL of the chart
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewHeatmapChart registers a new resource with the given unique name, arguments, and options.
func NewHeatmapChart(ctx *pulumi.Context,
	name string, args *HeatmapChartArgs, opts ...pulumi.ResourceOption) (*HeatmapChart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProgramText == nil {
		return nil, errors.New("invalid value for required argument 'ProgramText'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Values and color for the color range. Example: colorRange : { min : 0, max : 100, color : "#0000ff" }
	ColorRange *HeatmapChartColorRange `pulumi:"colorRange"`
	// Single color range including both the color to display for that range and the borders of the range
	ColorScales []HeatmapChartColorScale `pulumi:"colorScales"`
	// Description of the chart (Optional)
	Description *string `pulumi:"description"`
	// (false by default) If false, samples a subset of the output MTS, which improves UI performance
	DisableSampling *bool `pulumi:"disableSampling"`
	// Properties to group by in the heatmap (in nesting order)
	GroupBies []string `pulumi:"groupBies"`
	// (false by default) Whether to show the timestamp in the chart
	HideTimestamp *bool `pulumi:"hideTimestamp"`
	// How long (in seconds) to wait for late datapoints
	MaxDelay *int `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program
	MinimumResolution *int `pulumi:"minimumResolution"`
	// Name of the chart
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
	ProgramText *string `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the heatmap
	RefreshInterval *int `pulumi:"refreshInterval"`
	// The property to use when sorting the elements. Must be prepended with + for ascending or - for descending (e.g. -foo)
	SortBy *string `pulumi:"sortBy"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone *string `pulumi:"timezone"`
	// (Metric by default) Must be "Metric" or "Binary"
	UnitPrefix *string `pulumi:"unitPrefix"`
	// URL of the chart
	Url *string `pulumi:"url"`
}

type HeatmapChartState struct {
	// Values and color for the color range. Example: colorRange : { min : 0, max : 100, color : "#0000ff" }
	ColorRange HeatmapChartColorRangePtrInput
	// Single color range including both the color to display for that range and the borders of the range
	ColorScales HeatmapChartColorScaleArrayInput
	// Description of the chart (Optional)
	Description pulumi.StringPtrInput
	// (false by default) If false, samples a subset of the output MTS, which improves UI performance
	DisableSampling pulumi.BoolPtrInput
	// Properties to group by in the heatmap (in nesting order)
	GroupBies pulumi.StringArrayInput
	// (false by default) Whether to show the timestamp in the chart
	HideTimestamp pulumi.BoolPtrInput
	// How long (in seconds) to wait for late datapoints
	MaxDelay pulumi.IntPtrInput
	// The minimum resolution (in seconds) to use for computing the underlying program
	MinimumResolution pulumi.IntPtrInput
	// Name of the chart
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
	ProgramText pulumi.StringPtrInput
	// How often (in seconds) to refresh the values of the heatmap
	RefreshInterval pulumi.IntPtrInput
	// The property to use when sorting the elements. Must be prepended with + for ascending or - for descending (e.g. -foo)
	SortBy pulumi.StringPtrInput
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrInput
	// (Metric by default) Must be "Metric" or "Binary"
	UnitPrefix pulumi.StringPtrInput
	// URL of the chart
	Url pulumi.StringPtrInput
}

func (HeatmapChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*heatmapChartState)(nil)).Elem()
}

type heatmapChartArgs struct {
	// Values and color for the color range. Example: colorRange : { min : 0, max : 100, color : "#0000ff" }
	ColorRange *HeatmapChartColorRange `pulumi:"colorRange"`
	// Single color range including both the color to display for that range and the borders of the range
	ColorScales []HeatmapChartColorScale `pulumi:"colorScales"`
	// Description of the chart (Optional)
	Description *string `pulumi:"description"`
	// (false by default) If false, samples a subset of the output MTS, which improves UI performance
	DisableSampling *bool `pulumi:"disableSampling"`
	// Properties to group by in the heatmap (in nesting order)
	GroupBies []string `pulumi:"groupBies"`
	// (false by default) Whether to show the timestamp in the chart
	HideTimestamp *bool `pulumi:"hideTimestamp"`
	// How long (in seconds) to wait for late datapoints
	MaxDelay *int `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program
	MinimumResolution *int `pulumi:"minimumResolution"`
	// Name of the chart
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
	ProgramText string `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the heatmap
	RefreshInterval *int `pulumi:"refreshInterval"`
	// The property to use when sorting the elements. Must be prepended with + for ascending or - for descending (e.g. -foo)
	SortBy *string `pulumi:"sortBy"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone *string `pulumi:"timezone"`
	// (Metric by default) Must be "Metric" or "Binary"
	UnitPrefix *string `pulumi:"unitPrefix"`
}

// The set of arguments for constructing a HeatmapChart resource.
type HeatmapChartArgs struct {
	// Values and color for the color range. Example: colorRange : { min : 0, max : 100, color : "#0000ff" }
	ColorRange HeatmapChartColorRangePtrInput
	// Single color range including both the color to display for that range and the borders of the range
	ColorScales HeatmapChartColorScaleArrayInput
	// Description of the chart (Optional)
	Description pulumi.StringPtrInput
	// (false by default) If false, samples a subset of the output MTS, which improves UI performance
	DisableSampling pulumi.BoolPtrInput
	// Properties to group by in the heatmap (in nesting order)
	GroupBies pulumi.StringArrayInput
	// (false by default) Whether to show the timestamp in the chart
	HideTimestamp pulumi.BoolPtrInput
	// How long (in seconds) to wait for late datapoints
	MaxDelay pulumi.IntPtrInput
	// The minimum resolution (in seconds) to use for computing the underlying program
	MinimumResolution pulumi.IntPtrInput
	// Name of the chart
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
	ProgramText pulumi.StringInput
	// How often (in seconds) to refresh the values of the heatmap
	RefreshInterval pulumi.IntPtrInput
	// The property to use when sorting the elements. Must be prepended with + for ascending or - for descending (e.g. -foo)
	SortBy pulumi.StringPtrInput
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrInput
	// (Metric by default) Must be "Metric" or "Binary"
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

func (*HeatmapChart) ElementType() reflect.Type {
	return reflect.TypeOf((**HeatmapChart)(nil)).Elem()
}

func (i *HeatmapChart) ToHeatmapChartOutput() HeatmapChartOutput {
	return i.ToHeatmapChartOutputWithContext(context.Background())
}

func (i *HeatmapChart) ToHeatmapChartOutputWithContext(ctx context.Context) HeatmapChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeatmapChartOutput)
}

// HeatmapChartArrayInput is an input type that accepts HeatmapChartArray and HeatmapChartArrayOutput values.
// You can construct a concrete instance of `HeatmapChartArrayInput` via:
//
//	HeatmapChartArray{ HeatmapChartArgs{...} }
type HeatmapChartArrayInput interface {
	pulumi.Input

	ToHeatmapChartArrayOutput() HeatmapChartArrayOutput
	ToHeatmapChartArrayOutputWithContext(context.Context) HeatmapChartArrayOutput
}

type HeatmapChartArray []HeatmapChartInput

func (HeatmapChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HeatmapChart)(nil)).Elem()
}

func (i HeatmapChartArray) ToHeatmapChartArrayOutput() HeatmapChartArrayOutput {
	return i.ToHeatmapChartArrayOutputWithContext(context.Background())
}

func (i HeatmapChartArray) ToHeatmapChartArrayOutputWithContext(ctx context.Context) HeatmapChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeatmapChartArrayOutput)
}

// HeatmapChartMapInput is an input type that accepts HeatmapChartMap and HeatmapChartMapOutput values.
// You can construct a concrete instance of `HeatmapChartMapInput` via:
//
//	HeatmapChartMap{ "key": HeatmapChartArgs{...} }
type HeatmapChartMapInput interface {
	pulumi.Input

	ToHeatmapChartMapOutput() HeatmapChartMapOutput
	ToHeatmapChartMapOutputWithContext(context.Context) HeatmapChartMapOutput
}

type HeatmapChartMap map[string]HeatmapChartInput

func (HeatmapChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HeatmapChart)(nil)).Elem()
}

func (i HeatmapChartMap) ToHeatmapChartMapOutput() HeatmapChartMapOutput {
	return i.ToHeatmapChartMapOutputWithContext(context.Background())
}

func (i HeatmapChartMap) ToHeatmapChartMapOutputWithContext(ctx context.Context) HeatmapChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HeatmapChartMapOutput)
}

type HeatmapChartOutput struct{ *pulumi.OutputState }

func (HeatmapChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HeatmapChart)(nil)).Elem()
}

func (o HeatmapChartOutput) ToHeatmapChartOutput() HeatmapChartOutput {
	return o
}

func (o HeatmapChartOutput) ToHeatmapChartOutputWithContext(ctx context.Context) HeatmapChartOutput {
	return o
}

// Values and color for the color range. Example: colorRange : { min : 0, max : 100, color : "#0000ff" }
func (o HeatmapChartOutput) ColorRange() HeatmapChartColorRangePtrOutput {
	return o.ApplyT(func(v *HeatmapChart) HeatmapChartColorRangePtrOutput { return v.ColorRange }).(HeatmapChartColorRangePtrOutput)
}

// Single color range including both the color to display for that range and the borders of the range
func (o HeatmapChartOutput) ColorScales() HeatmapChartColorScaleArrayOutput {
	return o.ApplyT(func(v *HeatmapChart) HeatmapChartColorScaleArrayOutput { return v.ColorScales }).(HeatmapChartColorScaleArrayOutput)
}

// Description of the chart (Optional)
func (o HeatmapChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// (false by default) If false, samples a subset of the output MTS, which improves UI performance
func (o HeatmapChartOutput) DisableSampling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.BoolPtrOutput { return v.DisableSampling }).(pulumi.BoolPtrOutput)
}

// Properties to group by in the heatmap (in nesting order)
func (o HeatmapChartOutput) GroupBies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.StringArrayOutput { return v.GroupBies }).(pulumi.StringArrayOutput)
}

// (false by default) Whether to show the timestamp in the chart
func (o HeatmapChartOutput) HideTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.BoolPtrOutput { return v.HideTimestamp }).(pulumi.BoolPtrOutput)
}

// How long (in seconds) to wait for late datapoints
func (o HeatmapChartOutput) MaxDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.IntPtrOutput { return v.MaxDelay }).(pulumi.IntPtrOutput)
}

// The minimum resolution (in seconds) to use for computing the underlying program
func (o HeatmapChartOutput) MinimumResolution() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.IntPtrOutput { return v.MinimumResolution }).(pulumi.IntPtrOutput)
}

// Name of the chart
func (o HeatmapChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
func (o HeatmapChartOutput) ProgramText() pulumi.StringOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.StringOutput { return v.ProgramText }).(pulumi.StringOutput)
}

// How often (in seconds) to refresh the values of the heatmap
func (o HeatmapChartOutput) RefreshInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.IntPtrOutput { return v.RefreshInterval }).(pulumi.IntPtrOutput)
}

// The property to use when sorting the elements. Must be prepended with + for ascending or - for descending (e.g. -foo)
func (o HeatmapChartOutput) SortBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.StringPtrOutput { return v.SortBy }).(pulumi.StringPtrOutput)
}

// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
func (o HeatmapChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// (Metric by default) Must be "Metric" or "Binary"
func (o HeatmapChartOutput) UnitPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.StringPtrOutput { return v.UnitPrefix }).(pulumi.StringPtrOutput)
}

// URL of the chart
func (o HeatmapChartOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *HeatmapChart) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type HeatmapChartArrayOutput struct{ *pulumi.OutputState }

func (HeatmapChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HeatmapChart)(nil)).Elem()
}

func (o HeatmapChartArrayOutput) ToHeatmapChartArrayOutput() HeatmapChartArrayOutput {
	return o
}

func (o HeatmapChartArrayOutput) ToHeatmapChartArrayOutputWithContext(ctx context.Context) HeatmapChartArrayOutput {
	return o
}

func (o HeatmapChartArrayOutput) Index(i pulumi.IntInput) HeatmapChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HeatmapChart {
		return vs[0].([]*HeatmapChart)[vs[1].(int)]
	}).(HeatmapChartOutput)
}

type HeatmapChartMapOutput struct{ *pulumi.OutputState }

func (HeatmapChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HeatmapChart)(nil)).Elem()
}

func (o HeatmapChartMapOutput) ToHeatmapChartMapOutput() HeatmapChartMapOutput {
	return o
}

func (o HeatmapChartMapOutput) ToHeatmapChartMapOutputWithContext(ctx context.Context) HeatmapChartMapOutput {
	return o
}

func (o HeatmapChartMapOutput) MapIndex(k pulumi.StringInput) HeatmapChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HeatmapChart {
		return vs[0].(map[string]*HeatmapChart)[vs[1].(string)]
	}).(HeatmapChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HeatmapChartInput)(nil)).Elem(), &HeatmapChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*HeatmapChartArrayInput)(nil)).Elem(), HeatmapChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HeatmapChartMapInput)(nil)).Elem(), HeatmapChartMap{})
	pulumi.RegisterOutputType(HeatmapChartOutput{})
	pulumi.RegisterOutputType(HeatmapChartArrayOutput{})
	pulumi.RegisterOutputType(HeatmapChartMapOutput{})
}
