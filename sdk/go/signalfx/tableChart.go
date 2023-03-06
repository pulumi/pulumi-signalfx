// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This special type of chart displays a Data Table. This Table can be grouped by a Dimension.
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
//			_, err := signalfx.NewTableChart(ctx, "table0", &signalfx.TableChartArgs{
//				Description:     pulumi.String("beep"),
//				DisableSampling: pulumi.Bool(false),
//				GroupBies: pulumi.StringArray{
//					pulumi.String("ClusterName"),
//				},
//				MaxDelay:    pulumi.Int(0),
//				ProgramText: pulumi.String("A = data('cpu.usage.total').publish(label='CPU Total')"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type TableChart struct {
	pulumi.CustomResourceState

	// Description of the table chart.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// (false by default) If false, samples a subset of the output MTS, which improves UI performance
	DisableSampling pulumi.BoolPtrOutput `pulumi:"disableSampling"`
	// Dimension to group by
	GroupBies pulumi.StringArrayOutput `pulumi:"groupBies"`
	// (false by default) Whether to show the timestamp in the chart
	HideTimestamp pulumi.BoolPtrOutput `pulumi:"hideTimestamp"`
	// How long (in seconds) to wait for late datapoints
	MaxDelay pulumi.IntPtrOutput `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program
	MinimumResolution pulumi.IntPtrOutput `pulumi:"minimumResolution"`
	// Name of the table chart.
	Name pulumi.StringOutput `pulumi:"name"`
	// The SignalFlow for your Data Table Chart
	ProgramText pulumi.StringOutput `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the Table
	RefreshInterval pulumi.IntPtrOutput `pulumi:"refreshInterval"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// (Metric by default) Must be "Metric" or "Binary"
	UnitPrefix pulumi.StringPtrOutput `pulumi:"unitPrefix"`
	// The URL of the chart.
	Url pulumi.StringOutput `pulumi:"url"`
	// Plot-level customization options, associated with a publish statement
	VizOptions TableChartVizOptionArrayOutput `pulumi:"vizOptions"`
}

// NewTableChart registers a new resource with the given unique name, arguments, and options.
func NewTableChart(ctx *pulumi.Context,
	name string, args *TableChartArgs, opts ...pulumi.ResourceOption) (*TableChart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProgramText == nil {
		return nil, errors.New("invalid value for required argument 'ProgramText'")
	}
	var resource TableChart
	err := ctx.RegisterResource("signalfx:index/tableChart:TableChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTableChart gets an existing TableChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTableChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableChartState, opts ...pulumi.ResourceOption) (*TableChart, error) {
	var resource TableChart
	err := ctx.ReadResource("signalfx:index/tableChart:TableChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TableChart resources.
type tableChartState struct {
	// Description of the table chart.
	Description *string `pulumi:"description"`
	// (false by default) If false, samples a subset of the output MTS, which improves UI performance
	DisableSampling *bool `pulumi:"disableSampling"`
	// Dimension to group by
	GroupBies []string `pulumi:"groupBies"`
	// (false by default) Whether to show the timestamp in the chart
	HideTimestamp *bool `pulumi:"hideTimestamp"`
	// How long (in seconds) to wait for late datapoints
	MaxDelay *int `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program
	MinimumResolution *int `pulumi:"minimumResolution"`
	// Name of the table chart.
	Name *string `pulumi:"name"`
	// The SignalFlow for your Data Table Chart
	ProgramText *string `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the Table
	RefreshInterval *int `pulumi:"refreshInterval"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone *string `pulumi:"timezone"`
	// (Metric by default) Must be "Metric" or "Binary"
	UnitPrefix *string `pulumi:"unitPrefix"`
	// The URL of the chart.
	Url *string `pulumi:"url"`
	// Plot-level customization options, associated with a publish statement
	VizOptions []TableChartVizOption `pulumi:"vizOptions"`
}

type TableChartState struct {
	// Description of the table chart.
	Description pulumi.StringPtrInput
	// (false by default) If false, samples a subset of the output MTS, which improves UI performance
	DisableSampling pulumi.BoolPtrInput
	// Dimension to group by
	GroupBies pulumi.StringArrayInput
	// (false by default) Whether to show the timestamp in the chart
	HideTimestamp pulumi.BoolPtrInput
	// How long (in seconds) to wait for late datapoints
	MaxDelay pulumi.IntPtrInput
	// The minimum resolution (in seconds) to use for computing the underlying program
	MinimumResolution pulumi.IntPtrInput
	// Name of the table chart.
	Name pulumi.StringPtrInput
	// The SignalFlow for your Data Table Chart
	ProgramText pulumi.StringPtrInput
	// How often (in seconds) to refresh the values of the Table
	RefreshInterval pulumi.IntPtrInput
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrInput
	// (Metric by default) Must be "Metric" or "Binary"
	UnitPrefix pulumi.StringPtrInput
	// The URL of the chart.
	Url pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement
	VizOptions TableChartVizOptionArrayInput
}

func (TableChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableChartState)(nil)).Elem()
}

type tableChartArgs struct {
	// Description of the table chart.
	Description *string `pulumi:"description"`
	// (false by default) If false, samples a subset of the output MTS, which improves UI performance
	DisableSampling *bool `pulumi:"disableSampling"`
	// Dimension to group by
	GroupBies []string `pulumi:"groupBies"`
	// (false by default) Whether to show the timestamp in the chart
	HideTimestamp *bool `pulumi:"hideTimestamp"`
	// How long (in seconds) to wait for late datapoints
	MaxDelay *int `pulumi:"maxDelay"`
	// The minimum resolution (in seconds) to use for computing the underlying program
	MinimumResolution *int `pulumi:"minimumResolution"`
	// Name of the table chart.
	Name *string `pulumi:"name"`
	// The SignalFlow for your Data Table Chart
	ProgramText string `pulumi:"programText"`
	// How often (in seconds) to refresh the values of the Table
	RefreshInterval *int `pulumi:"refreshInterval"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone *string `pulumi:"timezone"`
	// (Metric by default) Must be "Metric" or "Binary"
	UnitPrefix *string `pulumi:"unitPrefix"`
	// Plot-level customization options, associated with a publish statement
	VizOptions []TableChartVizOption `pulumi:"vizOptions"`
}

// The set of arguments for constructing a TableChart resource.
type TableChartArgs struct {
	// Description of the table chart.
	Description pulumi.StringPtrInput
	// (false by default) If false, samples a subset of the output MTS, which improves UI performance
	DisableSampling pulumi.BoolPtrInput
	// Dimension to group by
	GroupBies pulumi.StringArrayInput
	// (false by default) Whether to show the timestamp in the chart
	HideTimestamp pulumi.BoolPtrInput
	// How long (in seconds) to wait for late datapoints
	MaxDelay pulumi.IntPtrInput
	// The minimum resolution (in seconds) to use for computing the underlying program
	MinimumResolution pulumi.IntPtrInput
	// Name of the table chart.
	Name pulumi.StringPtrInput
	// The SignalFlow for your Data Table Chart
	ProgramText pulumi.StringInput
	// How often (in seconds) to refresh the values of the Table
	RefreshInterval pulumi.IntPtrInput
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrInput
	// (Metric by default) Must be "Metric" or "Binary"
	UnitPrefix pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement
	VizOptions TableChartVizOptionArrayInput
}

func (TableChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableChartArgs)(nil)).Elem()
}

type TableChartInput interface {
	pulumi.Input

	ToTableChartOutput() TableChartOutput
	ToTableChartOutputWithContext(ctx context.Context) TableChartOutput
}

func (*TableChart) ElementType() reflect.Type {
	return reflect.TypeOf((**TableChart)(nil)).Elem()
}

func (i *TableChart) ToTableChartOutput() TableChartOutput {
	return i.ToTableChartOutputWithContext(context.Background())
}

func (i *TableChart) ToTableChartOutputWithContext(ctx context.Context) TableChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableChartOutput)
}

// TableChartArrayInput is an input type that accepts TableChartArray and TableChartArrayOutput values.
// You can construct a concrete instance of `TableChartArrayInput` via:
//
//	TableChartArray{ TableChartArgs{...} }
type TableChartArrayInput interface {
	pulumi.Input

	ToTableChartArrayOutput() TableChartArrayOutput
	ToTableChartArrayOutputWithContext(context.Context) TableChartArrayOutput
}

type TableChartArray []TableChartInput

func (TableChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TableChart)(nil)).Elem()
}

func (i TableChartArray) ToTableChartArrayOutput() TableChartArrayOutput {
	return i.ToTableChartArrayOutputWithContext(context.Background())
}

func (i TableChartArray) ToTableChartArrayOutputWithContext(ctx context.Context) TableChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableChartArrayOutput)
}

// TableChartMapInput is an input type that accepts TableChartMap and TableChartMapOutput values.
// You can construct a concrete instance of `TableChartMapInput` via:
//
//	TableChartMap{ "key": TableChartArgs{...} }
type TableChartMapInput interface {
	pulumi.Input

	ToTableChartMapOutput() TableChartMapOutput
	ToTableChartMapOutputWithContext(context.Context) TableChartMapOutput
}

type TableChartMap map[string]TableChartInput

func (TableChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TableChart)(nil)).Elem()
}

func (i TableChartMap) ToTableChartMapOutput() TableChartMapOutput {
	return i.ToTableChartMapOutputWithContext(context.Background())
}

func (i TableChartMap) ToTableChartMapOutputWithContext(ctx context.Context) TableChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableChartMapOutput)
}

type TableChartOutput struct{ *pulumi.OutputState }

func (TableChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableChart)(nil)).Elem()
}

func (o TableChartOutput) ToTableChartOutput() TableChartOutput {
	return o
}

func (o TableChartOutput) ToTableChartOutputWithContext(ctx context.Context) TableChartOutput {
	return o
}

// Description of the table chart.
func (o TableChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// (false by default) If false, samples a subset of the output MTS, which improves UI performance
func (o TableChartOutput) DisableSampling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TableChart) pulumi.BoolPtrOutput { return v.DisableSampling }).(pulumi.BoolPtrOutput)
}

// Dimension to group by
func (o TableChartOutput) GroupBies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableChart) pulumi.StringArrayOutput { return v.GroupBies }).(pulumi.StringArrayOutput)
}

// (false by default) Whether to show the timestamp in the chart
func (o TableChartOutput) HideTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TableChart) pulumi.BoolPtrOutput { return v.HideTimestamp }).(pulumi.BoolPtrOutput)
}

// How long (in seconds) to wait for late datapoints
func (o TableChartOutput) MaxDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TableChart) pulumi.IntPtrOutput { return v.MaxDelay }).(pulumi.IntPtrOutput)
}

// The minimum resolution (in seconds) to use for computing the underlying program
func (o TableChartOutput) MinimumResolution() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TableChart) pulumi.IntPtrOutput { return v.MinimumResolution }).(pulumi.IntPtrOutput)
}

// Name of the table chart.
func (o TableChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TableChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The SignalFlow for your Data Table Chart
func (o TableChartOutput) ProgramText() pulumi.StringOutput {
	return o.ApplyT(func(v *TableChart) pulumi.StringOutput { return v.ProgramText }).(pulumi.StringOutput)
}

// How often (in seconds) to refresh the values of the Table
func (o TableChartOutput) RefreshInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TableChart) pulumi.IntPtrOutput { return v.RefreshInterval }).(pulumi.IntPtrOutput)
}

// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
func (o TableChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// (Metric by default) Must be "Metric" or "Binary"
func (o TableChartOutput) UnitPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableChart) pulumi.StringPtrOutput { return v.UnitPrefix }).(pulumi.StringPtrOutput)
}

// The URL of the chart.
func (o TableChartOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *TableChart) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Plot-level customization options, associated with a publish statement
func (o TableChartOutput) VizOptions() TableChartVizOptionArrayOutput {
	return o.ApplyT(func(v *TableChart) TableChartVizOptionArrayOutput { return v.VizOptions }).(TableChartVizOptionArrayOutput)
}

type TableChartArrayOutput struct{ *pulumi.OutputState }

func (TableChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TableChart)(nil)).Elem()
}

func (o TableChartArrayOutput) ToTableChartArrayOutput() TableChartArrayOutput {
	return o
}

func (o TableChartArrayOutput) ToTableChartArrayOutputWithContext(ctx context.Context) TableChartArrayOutput {
	return o
}

func (o TableChartArrayOutput) Index(i pulumi.IntInput) TableChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TableChart {
		return vs[0].([]*TableChart)[vs[1].(int)]
	}).(TableChartOutput)
}

type TableChartMapOutput struct{ *pulumi.OutputState }

func (TableChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TableChart)(nil)).Elem()
}

func (o TableChartMapOutput) ToTableChartMapOutput() TableChartMapOutput {
	return o
}

func (o TableChartMapOutput) ToTableChartMapOutputWithContext(ctx context.Context) TableChartMapOutput {
	return o
}

func (o TableChartMapOutput) MapIndex(k pulumi.StringInput) TableChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TableChart {
		return vs[0].(map[string]*TableChart)[vs[1].(string)]
	}).(TableChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableChartInput)(nil)).Elem(), &TableChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableChartArrayInput)(nil)).Elem(), TableChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableChartMapInput)(nil)).Elem(), TableChartMap{})
	pulumi.RegisterOutputType(TableChartOutput{})
	pulumi.RegisterOutputType(TableChartArrayOutput{})
	pulumi.RegisterOutputType(TableChartMapOutput{})
}
