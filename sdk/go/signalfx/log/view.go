// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package log

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-signalfx/sdk/v6/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// You can add logs data to your Observability Cloud dashboards without turning your logs into metrics first. A log view displays log lines in a table form in a dashboard and shows you in detail what is happening and why.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-signalfx/sdk/v6/go/signalfx/log"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := log.NewView(ctx, "myLogView", &log.ViewArgs{
//				Columns: log.ViewColumnArray{
//					&log.ViewColumnArgs{
//						Name: pulumi.String("severity"),
//					},
//					&log.ViewColumnArgs{
//						Name: pulumi.String("time"),
//					},
//					&log.ViewColumnArgs{
//						Name: pulumi.String("amount.currency_code"),
//					},
//					&log.ViewColumnArgs{
//						Name: pulumi.String("amount.nanos"),
//					},
//					&log.ViewColumnArgs{
//						Name: pulumi.String("amount.units"),
//					},
//					&log.ViewColumnArgs{
//						Name: pulumi.String("message"),
//					},
//				},
//				Description: pulumi.String("Lorem ipsum dolor sit amet, laudem tibique iracundia at mea. Nam posse dolores ex, nec cu adhuc putent honestatis"),
//				ProgramText: pulumi.String("logs(filter=field('message') == 'Transaction processed' and field('service.name') == 'paymentservice').publish()\n\n"),
//				SortOptions: log.ViewSortOptionArray{
//					&log.ViewSortOptionArgs{
//						Descending: pulumi.Bool(false),
//						Field:      pulumi.String("severity"),
//					},
//				},
//				TimeRange: pulumi.Int(900),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type View struct {
	pulumi.CustomResourceState

	// The column headers to show on the log view.
	Columns ViewColumnArrayOutput `pulumi:"columns"`
	// The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
	DefaultConnection pulumi.StringPtrOutput `pulumi:"defaultConnection"`
	// Description of the log view.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrOutput `pulumi:"endTime"`
	// Name of the log view.
	Name pulumi.StringOutput `pulumi:"name"`
	// Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
	ProgramText pulumi.StringOutput `pulumi:"programText"`
	// The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
	SortOptions ViewSortOptionArrayOutput `pulumi:"sortOptions"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrOutput `pulumi:"startTime"`
	// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrOutput `pulumi:"timeRange"`
	// The URL of the log view.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewView registers a new resource with the given unique name, arguments, and options.
func NewView(ctx *pulumi.Context,
	name string, args *ViewArgs, opts ...pulumi.ResourceOption) (*View, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProgramText == nil {
		return nil, errors.New("invalid value for required argument 'ProgramText'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("signalfx:logs/view:View"),
		},
	})
	opts = append(opts, aliases)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource View
	err := ctx.RegisterResource("signalfx:log/view:View", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetView gets an existing View resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewState, opts ...pulumi.ResourceOption) (*View, error) {
	var resource View
	err := ctx.ReadResource("signalfx:log/view:View", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering View resources.
type viewState struct {
	// The column headers to show on the log view.
	Columns []ViewColumn `pulumi:"columns"`
	// The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
	DefaultConnection *string `pulumi:"defaultConnection"`
	// Description of the log view.
	Description *string `pulumi:"description"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime *int `pulumi:"endTime"`
	// Name of the log view.
	Name *string `pulumi:"name"`
	// Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
	ProgramText *string `pulumi:"programText"`
	// The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
	SortOptions []ViewSortOption `pulumi:"sortOptions"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime *int `pulumi:"startTime"`
	// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
	TimeRange *int `pulumi:"timeRange"`
	// The URL of the log view.
	Url *string `pulumi:"url"`
}

type ViewState struct {
	// The column headers to show on the log view.
	Columns ViewColumnArrayInput
	// The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
	DefaultConnection pulumi.StringPtrInput
	// Description of the log view.
	Description pulumi.StringPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrInput
	// Name of the log view.
	Name pulumi.StringPtrInput
	// Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
	ProgramText pulumi.StringPtrInput
	// The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
	SortOptions ViewSortOptionArrayInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrInput
	// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrInput
	// The URL of the log view.
	Url pulumi.StringPtrInput
}

func (ViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewState)(nil)).Elem()
}

type viewArgs struct {
	// The column headers to show on the log view.
	Columns []ViewColumn `pulumi:"columns"`
	// The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
	DefaultConnection *string `pulumi:"defaultConnection"`
	// Description of the log view.
	Description *string `pulumi:"description"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime *int `pulumi:"endTime"`
	// Name of the log view.
	Name *string `pulumi:"name"`
	// Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
	ProgramText string `pulumi:"programText"`
	// The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
	SortOptions []ViewSortOption `pulumi:"sortOptions"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime *int `pulumi:"startTime"`
	// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
	TimeRange *int `pulumi:"timeRange"`
}

// The set of arguments for constructing a View resource.
type ViewArgs struct {
	// The column headers to show on the log view.
	Columns ViewColumnArrayInput
	// The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
	DefaultConnection pulumi.StringPtrInput
	// Description of the log view.
	Description pulumi.StringPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrInput
	// Name of the log view.
	Name pulumi.StringPtrInput
	// Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
	ProgramText pulumi.StringInput
	// The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
	SortOptions ViewSortOptionArrayInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrInput
	// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
	TimeRange pulumi.IntPtrInput
}

func (ViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewArgs)(nil)).Elem()
}

type ViewInput interface {
	pulumi.Input

	ToViewOutput() ViewOutput
	ToViewOutputWithContext(ctx context.Context) ViewOutput
}

func (*View) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (i *View) ToViewOutput() ViewOutput {
	return i.ToViewOutputWithContext(context.Background())
}

func (i *View) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewOutput)
}

// ViewArrayInput is an input type that accepts ViewArray and ViewArrayOutput values.
// You can construct a concrete instance of `ViewArrayInput` via:
//
//	ViewArray{ ViewArgs{...} }
type ViewArrayInput interface {
	pulumi.Input

	ToViewArrayOutput() ViewArrayOutput
	ToViewArrayOutputWithContext(context.Context) ViewArrayOutput
}

type ViewArray []ViewInput

func (ViewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*View)(nil)).Elem()
}

func (i ViewArray) ToViewArrayOutput() ViewArrayOutput {
	return i.ToViewArrayOutputWithContext(context.Background())
}

func (i ViewArray) ToViewArrayOutputWithContext(ctx context.Context) ViewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewArrayOutput)
}

// ViewMapInput is an input type that accepts ViewMap and ViewMapOutput values.
// You can construct a concrete instance of `ViewMapInput` via:
//
//	ViewMap{ "key": ViewArgs{...} }
type ViewMapInput interface {
	pulumi.Input

	ToViewMapOutput() ViewMapOutput
	ToViewMapOutputWithContext(context.Context) ViewMapOutput
}

type ViewMap map[string]ViewInput

func (ViewMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*View)(nil)).Elem()
}

func (i ViewMap) ToViewMapOutput() ViewMapOutput {
	return i.ToViewMapOutputWithContext(context.Background())
}

func (i ViewMap) ToViewMapOutputWithContext(ctx context.Context) ViewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewMapOutput)
}

type ViewOutput struct{ *pulumi.OutputState }

func (ViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (o ViewOutput) ToViewOutput() ViewOutput {
	return o
}

func (o ViewOutput) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return o
}

// The column headers to show on the log view.
func (o ViewOutput) Columns() ViewColumnArrayOutput {
	return o.ApplyT(func(v *View) ViewColumnArrayOutput { return v.Columns }).(ViewColumnArrayOutput)
}

// The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
func (o ViewOutput) DefaultConnection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *View) pulumi.StringPtrOutput { return v.DefaultConnection }).(pulumi.StringPtrOutput)
}

// Description of the log view.
func (o ViewOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *View) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
func (o ViewOutput) EndTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *View) pulumi.IntPtrOutput { return v.EndTime }).(pulumi.IntPtrOutput)
}

// Name of the log view.
func (o ViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
func (o ViewOutput) ProgramText() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.ProgramText }).(pulumi.StringOutput)
}

// The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
func (o ViewOutput) SortOptions() ViewSortOptionArrayOutput {
	return o.ApplyT(func(v *View) ViewSortOptionArrayOutput { return v.SortOptions }).(ViewSortOptionArrayOutput)
}

// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
func (o ViewOutput) StartTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *View) pulumi.IntPtrOutput { return v.StartTime }).(pulumi.IntPtrOutput)
}

// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
func (o ViewOutput) TimeRange() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *View) pulumi.IntPtrOutput { return v.TimeRange }).(pulumi.IntPtrOutput)
}

// The URL of the log view.
func (o ViewOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type ViewArrayOutput struct{ *pulumi.OutputState }

func (ViewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*View)(nil)).Elem()
}

func (o ViewArrayOutput) ToViewArrayOutput() ViewArrayOutput {
	return o
}

func (o ViewArrayOutput) ToViewArrayOutputWithContext(ctx context.Context) ViewArrayOutput {
	return o
}

func (o ViewArrayOutput) Index(i pulumi.IntInput) ViewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *View {
		return vs[0].([]*View)[vs[1].(int)]
	}).(ViewOutput)
}

type ViewMapOutput struct{ *pulumi.OutputState }

func (ViewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*View)(nil)).Elem()
}

func (o ViewMapOutput) ToViewMapOutput() ViewMapOutput {
	return o
}

func (o ViewMapOutput) ToViewMapOutputWithContext(ctx context.Context) ViewMapOutput {
	return o
}

func (o ViewMapOutput) MapIndex(k pulumi.StringInput) ViewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *View {
		return vs[0].(map[string]*View)[vs[1].(string)]
	}).(ViewOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ViewInput)(nil)).Elem(), &View{})
	pulumi.RegisterInputType(reflect.TypeOf((*ViewArrayInput)(nil)).Elem(), ViewArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ViewMapInput)(nil)).Elem(), ViewMap{})
	pulumi.RegisterOutputType(ViewOutput{})
	pulumi.RegisterOutputType(ViewArrayOutput{})
	pulumi.RegisterOutputType(ViewMapOutput{})
}
