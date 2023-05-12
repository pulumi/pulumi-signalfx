// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This chart type displays a single number in a large font, representing the current value of a single metric on a plot line.
//
// If the time period is in the past, the number represents the value of the metric near the end of the time period.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-signalfx/sdk/v5/go/signalfx"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := signalfx.NewSingleValueChart(ctx, "mysvchart0", &signalfx.SingleValueChartArgs{
//				ColorBy:           pulumi.String("Dimension"),
//				Description:       pulumi.String("Very cool Single Value Chart"),
//				IsTimestampHidden: pulumi.Bool(true),
//				MaxDelay:          pulumi.Int(2),
//				MaxPrecision:      pulumi.Int(2),
//				ProgramText:       pulumi.String("myfilters = filter(\"cluster_name\", \"prod\") and filter(\"role\", \"search\")\ndata(\"cpu.total.idle\", filter=myfilters).publish()\n\n"),
//				RefreshInterval:   pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type SingleValueChart struct {
	pulumi.CustomResourceState

	// Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
	ColorBy pulumi.StringPtrOutput `pulumi:"colorBy"`
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales SingleValueChartColorScaleArrayOutput `pulumi:"colorScales"`
	// Description of the chart.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to hide the timestamp in the chart. `false` by default.
	IsTimestampHidden pulumi.BoolPtrOutput `pulumi:"isTimestampHidden"`
	// How long (in seconds) to wait for late datapoints
	MaxDelay pulumi.IntPtrOutput `pulumi:"maxDelay"`
	// The maximum precision to for value displayed.
	MaxPrecision pulumi.IntPtrOutput `pulumi:"maxPrecision"`
	// Name of the chart.
	Name pulumi.StringOutput `pulumi:"name"`
	// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringOutput `pulumi:"programText"`
	// How often (in seconds) to refresh the value.
	RefreshInterval pulumi.IntPtrOutput `pulumi:"refreshInterval"`
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
	SecondaryVisualization pulumi.StringPtrOutput `pulumi:"secondaryVisualization"`
	// Whether to show a trend line below the current value. `false` by default.
	ShowSparkLine pulumi.BoolPtrOutput `pulumi:"showSparkLine"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
	UnitPrefix pulumi.StringPtrOutput `pulumi:"unitPrefix"`
	// The URL of the chart.
	Url pulumi.StringOutput `pulumi:"url"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions SingleValueChartVizOptionArrayOutput `pulumi:"vizOptions"`
}

// NewSingleValueChart registers a new resource with the given unique name, arguments, and options.
func NewSingleValueChart(ctx *pulumi.Context,
	name string, args *SingleValueChartArgs, opts ...pulumi.ResourceOption) (*SingleValueChart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProgramText == nil {
		return nil, errors.New("invalid value for required argument 'ProgramText'")
	}
	var resource SingleValueChart
	err := ctx.RegisterResource("signalfx:index/singleValueChart:SingleValueChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSingleValueChart gets an existing SingleValueChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSingleValueChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SingleValueChartState, opts ...pulumi.ResourceOption) (*SingleValueChart, error) {
	var resource SingleValueChart
	err := ctx.ReadResource("signalfx:index/singleValueChart:SingleValueChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SingleValueChart resources.
type singleValueChartState struct {
	// Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
	ColorBy *string `pulumi:"colorBy"`
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales []SingleValueChartColorScale `pulumi:"colorScales"`
	// Description of the chart.
	Description *string `pulumi:"description"`
	// Whether to hide the timestamp in the chart. `false` by default.
	IsTimestampHidden *bool `pulumi:"isTimestampHidden"`
	// How long (in seconds) to wait for late datapoints
	MaxDelay *int `pulumi:"maxDelay"`
	// The maximum precision to for value displayed.
	MaxPrecision *int `pulumi:"maxPrecision"`
	// Name of the chart.
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText *string `pulumi:"programText"`
	// How often (in seconds) to refresh the value.
	RefreshInterval *int `pulumi:"refreshInterval"`
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
	SecondaryVisualization *string `pulumi:"secondaryVisualization"`
	// Whether to show a trend line below the current value. `false` by default.
	ShowSparkLine *bool `pulumi:"showSparkLine"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone *string `pulumi:"timezone"`
	// Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
	UnitPrefix *string `pulumi:"unitPrefix"`
	// The URL of the chart.
	Url *string `pulumi:"url"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions []SingleValueChartVizOption `pulumi:"vizOptions"`
}

type SingleValueChartState struct {
	// Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
	ColorBy pulumi.StringPtrInput
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales SingleValueChartColorScaleArrayInput
	// Description of the chart.
	Description pulumi.StringPtrInput
	// Whether to hide the timestamp in the chart. `false` by default.
	IsTimestampHidden pulumi.BoolPtrInput
	// How long (in seconds) to wait for late datapoints
	MaxDelay pulumi.IntPtrInput
	// The maximum precision to for value displayed.
	MaxPrecision pulumi.IntPtrInput
	// Name of the chart.
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringPtrInput
	// How often (in seconds) to refresh the value.
	RefreshInterval pulumi.IntPtrInput
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
	SecondaryVisualization pulumi.StringPtrInput
	// Whether to show a trend line below the current value. `false` by default.
	ShowSparkLine pulumi.BoolPtrInput
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrInput
	// Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
	UnitPrefix pulumi.StringPtrInput
	// The URL of the chart.
	Url pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement.
	VizOptions SingleValueChartVizOptionArrayInput
}

func (SingleValueChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*singleValueChartState)(nil)).Elem()
}

type singleValueChartArgs struct {
	// Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
	ColorBy *string `pulumi:"colorBy"`
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales []SingleValueChartColorScale `pulumi:"colorScales"`
	// Description of the chart.
	Description *string `pulumi:"description"`
	// Whether to hide the timestamp in the chart. `false` by default.
	IsTimestampHidden *bool `pulumi:"isTimestampHidden"`
	// How long (in seconds) to wait for late datapoints
	MaxDelay *int `pulumi:"maxDelay"`
	// The maximum precision to for value displayed.
	MaxPrecision *int `pulumi:"maxPrecision"`
	// Name of the chart.
	Name *string `pulumi:"name"`
	// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText string `pulumi:"programText"`
	// How often (in seconds) to refresh the value.
	RefreshInterval *int `pulumi:"refreshInterval"`
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
	SecondaryVisualization *string `pulumi:"secondaryVisualization"`
	// Whether to show a trend line below the current value. `false` by default.
	ShowSparkLine *bool `pulumi:"showSparkLine"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone *string `pulumi:"timezone"`
	// Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
	UnitPrefix *string `pulumi:"unitPrefix"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions []SingleValueChartVizOption `pulumi:"vizOptions"`
}

// The set of arguments for constructing a SingleValueChart resource.
type SingleValueChartArgs struct {
	// Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
	ColorBy pulumi.StringPtrInput
	// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
	ColorScales SingleValueChartColorScaleArrayInput
	// Description of the chart.
	Description pulumi.StringPtrInput
	// Whether to hide the timestamp in the chart. `false` by default.
	IsTimestampHidden pulumi.BoolPtrInput
	// How long (in seconds) to wait for late datapoints
	MaxDelay pulumi.IntPtrInput
	// The maximum precision to for value displayed.
	MaxPrecision pulumi.IntPtrInput
	// Name of the chart.
	Name pulumi.StringPtrInput
	// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringInput
	// How often (in seconds) to refresh the value.
	RefreshInterval pulumi.IntPtrInput
	// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
	SecondaryVisualization pulumi.StringPtrInput
	// Whether to show a trend line below the current value. `false` by default.
	ShowSparkLine pulumi.BoolPtrInput
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrInput
	// Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
	UnitPrefix pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement.
	VizOptions SingleValueChartVizOptionArrayInput
}

func (SingleValueChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*singleValueChartArgs)(nil)).Elem()
}

type SingleValueChartInput interface {
	pulumi.Input

	ToSingleValueChartOutput() SingleValueChartOutput
	ToSingleValueChartOutputWithContext(ctx context.Context) SingleValueChartOutput
}

func (*SingleValueChart) ElementType() reflect.Type {
	return reflect.TypeOf((**SingleValueChart)(nil)).Elem()
}

func (i *SingleValueChart) ToSingleValueChartOutput() SingleValueChartOutput {
	return i.ToSingleValueChartOutputWithContext(context.Background())
}

func (i *SingleValueChart) ToSingleValueChartOutputWithContext(ctx context.Context) SingleValueChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SingleValueChartOutput)
}

// SingleValueChartArrayInput is an input type that accepts SingleValueChartArray and SingleValueChartArrayOutput values.
// You can construct a concrete instance of `SingleValueChartArrayInput` via:
//
//	SingleValueChartArray{ SingleValueChartArgs{...} }
type SingleValueChartArrayInput interface {
	pulumi.Input

	ToSingleValueChartArrayOutput() SingleValueChartArrayOutput
	ToSingleValueChartArrayOutputWithContext(context.Context) SingleValueChartArrayOutput
}

type SingleValueChartArray []SingleValueChartInput

func (SingleValueChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SingleValueChart)(nil)).Elem()
}

func (i SingleValueChartArray) ToSingleValueChartArrayOutput() SingleValueChartArrayOutput {
	return i.ToSingleValueChartArrayOutputWithContext(context.Background())
}

func (i SingleValueChartArray) ToSingleValueChartArrayOutputWithContext(ctx context.Context) SingleValueChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SingleValueChartArrayOutput)
}

// SingleValueChartMapInput is an input type that accepts SingleValueChartMap and SingleValueChartMapOutput values.
// You can construct a concrete instance of `SingleValueChartMapInput` via:
//
//	SingleValueChartMap{ "key": SingleValueChartArgs{...} }
type SingleValueChartMapInput interface {
	pulumi.Input

	ToSingleValueChartMapOutput() SingleValueChartMapOutput
	ToSingleValueChartMapOutputWithContext(context.Context) SingleValueChartMapOutput
}

type SingleValueChartMap map[string]SingleValueChartInput

func (SingleValueChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SingleValueChart)(nil)).Elem()
}

func (i SingleValueChartMap) ToSingleValueChartMapOutput() SingleValueChartMapOutput {
	return i.ToSingleValueChartMapOutputWithContext(context.Background())
}

func (i SingleValueChartMap) ToSingleValueChartMapOutputWithContext(ctx context.Context) SingleValueChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SingleValueChartMapOutput)
}

type SingleValueChartOutput struct{ *pulumi.OutputState }

func (SingleValueChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SingleValueChart)(nil)).Elem()
}

func (o SingleValueChartOutput) ToSingleValueChartOutput() SingleValueChartOutput {
	return o
}

func (o SingleValueChartOutput) ToSingleValueChartOutputWithContext(ctx context.Context) SingleValueChartOutput {
	return o
}

// Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
func (o SingleValueChartOutput) ColorBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.StringPtrOutput { return v.ColorBy }).(pulumi.StringPtrOutput)
}

// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
func (o SingleValueChartOutput) ColorScales() SingleValueChartColorScaleArrayOutput {
	return o.ApplyT(func(v *SingleValueChart) SingleValueChartColorScaleArrayOutput { return v.ColorScales }).(SingleValueChartColorScaleArrayOutput)
}

// Description of the chart.
func (o SingleValueChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether to hide the timestamp in the chart. `false` by default.
func (o SingleValueChartOutput) IsTimestampHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.BoolPtrOutput { return v.IsTimestampHidden }).(pulumi.BoolPtrOutput)
}

// How long (in seconds) to wait for late datapoints
func (o SingleValueChartOutput) MaxDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.IntPtrOutput { return v.MaxDelay }).(pulumi.IntPtrOutput)
}

// The maximum precision to for value displayed.
func (o SingleValueChartOutput) MaxPrecision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.IntPtrOutput { return v.MaxPrecision }).(pulumi.IntPtrOutput)
}

// Name of the chart.
func (o SingleValueChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
func (o SingleValueChartOutput) ProgramText() pulumi.StringOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.StringOutput { return v.ProgramText }).(pulumi.StringOutput)
}

// How often (in seconds) to refresh the value.
func (o SingleValueChartOutput) RefreshInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.IntPtrOutput { return v.RefreshInterval }).(pulumi.IntPtrOutput)
}

// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
func (o SingleValueChartOutput) SecondaryVisualization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.StringPtrOutput { return v.SecondaryVisualization }).(pulumi.StringPtrOutput)
}

// Whether to show a trend line below the current value. `false` by default.
func (o SingleValueChartOutput) ShowSparkLine() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.BoolPtrOutput { return v.ShowSparkLine }).(pulumi.BoolPtrOutput)
}

// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
func (o SingleValueChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
func (o SingleValueChartOutput) UnitPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.StringPtrOutput { return v.UnitPrefix }).(pulumi.StringPtrOutput)
}

// The URL of the chart.
func (o SingleValueChartOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *SingleValueChart) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Plot-level customization options, associated with a publish statement.
func (o SingleValueChartOutput) VizOptions() SingleValueChartVizOptionArrayOutput {
	return o.ApplyT(func(v *SingleValueChart) SingleValueChartVizOptionArrayOutput { return v.VizOptions }).(SingleValueChartVizOptionArrayOutput)
}

type SingleValueChartArrayOutput struct{ *pulumi.OutputState }

func (SingleValueChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SingleValueChart)(nil)).Elem()
}

func (o SingleValueChartArrayOutput) ToSingleValueChartArrayOutput() SingleValueChartArrayOutput {
	return o
}

func (o SingleValueChartArrayOutput) ToSingleValueChartArrayOutputWithContext(ctx context.Context) SingleValueChartArrayOutput {
	return o
}

func (o SingleValueChartArrayOutput) Index(i pulumi.IntInput) SingleValueChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SingleValueChart {
		return vs[0].([]*SingleValueChart)[vs[1].(int)]
	}).(SingleValueChartOutput)
}

type SingleValueChartMapOutput struct{ *pulumi.OutputState }

func (SingleValueChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SingleValueChart)(nil)).Elem()
}

func (o SingleValueChartMapOutput) ToSingleValueChartMapOutput() SingleValueChartMapOutput {
	return o
}

func (o SingleValueChartMapOutput) ToSingleValueChartMapOutputWithContext(ctx context.Context) SingleValueChartMapOutput {
	return o
}

func (o SingleValueChartMapOutput) MapIndex(k pulumi.StringInput) SingleValueChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SingleValueChart {
		return vs[0].(map[string]*SingleValueChart)[vs[1].(string)]
	}).(SingleValueChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SingleValueChartInput)(nil)).Elem(), &SingleValueChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*SingleValueChartArrayInput)(nil)).Elem(), SingleValueChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SingleValueChartMapInput)(nil)).Elem(), SingleValueChartMap{})
	pulumi.RegisterOutputType(SingleValueChartOutput{})
	pulumi.RegisterOutputType(SingleValueChartArrayOutput{})
	pulumi.RegisterOutputType(SingleValueChartMapOutput{})
}
