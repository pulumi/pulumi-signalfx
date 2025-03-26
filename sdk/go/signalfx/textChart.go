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

// This special type of chart doesn’t display any metric data. Rather, it lets you place a text note on the dashboard.
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
//			_, err := signalfx.NewTextChart(ctx, "mynote0", &signalfx.TextChartArgs{
//				Name:        pulumi.String("Important Dashboard Note"),
//				Description: pulumi.String("Lorem ipsum dolor sit amet, laudem tibique iracundia at mea. Nam posse dolores ex, nec cu adhuc putent honestatis"),
//				Markdown: pulumi.String(`1. First ordered list item
//
// 2. Another item
//   - Unordered sub-list.
//
// 1. Actual numbers don't matter, just that it's a number
//  1. Ordered sub-list
//
// 4. And another item.
//
//	You can have properly indented paragraphs within list items. Notice the blank line above, and the leading spaces (at least one, but we'll use three here to also align the raw Markdown).
//
//	To have a line break without a paragraph, you will need to use two trailing spaces.⋅⋅
//	Note that this line is separate, but within the same paragraph.⋅⋅
//	(This is contrary to the typical GFM line break behaviour, where trailing spaces are not required.)
//
// * Unordered list can use asterisks
// - Or minuses
// + Or pluses
// `),
//
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type TextChart struct {
	pulumi.CustomResourceState

	// Description of the text note.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Markdown text to display.
	Markdown pulumi.StringOutput `pulumi:"markdown"`
	// Name of the text note.
	Name pulumi.StringOutput `pulumi:"name"`
	// Tags associated with the resource
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The URL of the chart.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewTextChart registers a new resource with the given unique name, arguments, and options.
func NewTextChart(ctx *pulumi.Context,
	name string, args *TextChartArgs, opts ...pulumi.ResourceOption) (*TextChart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Markdown == nil {
		return nil, errors.New("invalid value for required argument 'Markdown'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TextChart
	err := ctx.RegisterResource("signalfx:index/textChart:TextChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTextChart gets an existing TextChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTextChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TextChartState, opts ...pulumi.ResourceOption) (*TextChart, error) {
	var resource TextChart
	err := ctx.ReadResource("signalfx:index/textChart:TextChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TextChart resources.
type textChartState struct {
	// Description of the text note.
	Description *string `pulumi:"description"`
	// Markdown text to display.
	Markdown *string `pulumi:"markdown"`
	// Name of the text note.
	Name *string `pulumi:"name"`
	// Tags associated with the resource
	Tags []string `pulumi:"tags"`
	// The URL of the chart.
	Url *string `pulumi:"url"`
}

type TextChartState struct {
	// Description of the text note.
	Description pulumi.StringPtrInput
	// Markdown text to display.
	Markdown pulumi.StringPtrInput
	// Name of the text note.
	Name pulumi.StringPtrInput
	// Tags associated with the resource
	Tags pulumi.StringArrayInput
	// The URL of the chart.
	Url pulumi.StringPtrInput
}

func (TextChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*textChartState)(nil)).Elem()
}

type textChartArgs struct {
	// Description of the text note.
	Description *string `pulumi:"description"`
	// Markdown text to display.
	Markdown string `pulumi:"markdown"`
	// Name of the text note.
	Name *string `pulumi:"name"`
	// Tags associated with the resource
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a TextChart resource.
type TextChartArgs struct {
	// Description of the text note.
	Description pulumi.StringPtrInput
	// Markdown text to display.
	Markdown pulumi.StringInput
	// Name of the text note.
	Name pulumi.StringPtrInput
	// Tags associated with the resource
	Tags pulumi.StringArrayInput
}

func (TextChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*textChartArgs)(nil)).Elem()
}

type TextChartInput interface {
	pulumi.Input

	ToTextChartOutput() TextChartOutput
	ToTextChartOutputWithContext(ctx context.Context) TextChartOutput
}

func (*TextChart) ElementType() reflect.Type {
	return reflect.TypeOf((**TextChart)(nil)).Elem()
}

func (i *TextChart) ToTextChartOutput() TextChartOutput {
	return i.ToTextChartOutputWithContext(context.Background())
}

func (i *TextChart) ToTextChartOutputWithContext(ctx context.Context) TextChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TextChartOutput)
}

// TextChartArrayInput is an input type that accepts TextChartArray and TextChartArrayOutput values.
// You can construct a concrete instance of `TextChartArrayInput` via:
//
//	TextChartArray{ TextChartArgs{...} }
type TextChartArrayInput interface {
	pulumi.Input

	ToTextChartArrayOutput() TextChartArrayOutput
	ToTextChartArrayOutputWithContext(context.Context) TextChartArrayOutput
}

type TextChartArray []TextChartInput

func (TextChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TextChart)(nil)).Elem()
}

func (i TextChartArray) ToTextChartArrayOutput() TextChartArrayOutput {
	return i.ToTextChartArrayOutputWithContext(context.Background())
}

func (i TextChartArray) ToTextChartArrayOutputWithContext(ctx context.Context) TextChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TextChartArrayOutput)
}

// TextChartMapInput is an input type that accepts TextChartMap and TextChartMapOutput values.
// You can construct a concrete instance of `TextChartMapInput` via:
//
//	TextChartMap{ "key": TextChartArgs{...} }
type TextChartMapInput interface {
	pulumi.Input

	ToTextChartMapOutput() TextChartMapOutput
	ToTextChartMapOutputWithContext(context.Context) TextChartMapOutput
}

type TextChartMap map[string]TextChartInput

func (TextChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TextChart)(nil)).Elem()
}

func (i TextChartMap) ToTextChartMapOutput() TextChartMapOutput {
	return i.ToTextChartMapOutputWithContext(context.Background())
}

func (i TextChartMap) ToTextChartMapOutputWithContext(ctx context.Context) TextChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TextChartMapOutput)
}

type TextChartOutput struct{ *pulumi.OutputState }

func (TextChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TextChart)(nil)).Elem()
}

func (o TextChartOutput) ToTextChartOutput() TextChartOutput {
	return o
}

func (o TextChartOutput) ToTextChartOutputWithContext(ctx context.Context) TextChartOutput {
	return o
}

// Description of the text note.
func (o TextChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TextChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Markdown text to display.
func (o TextChartOutput) Markdown() pulumi.StringOutput {
	return o.ApplyT(func(v *TextChart) pulumi.StringOutput { return v.Markdown }).(pulumi.StringOutput)
}

// Name of the text note.
func (o TextChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TextChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Tags associated with the resource
func (o TextChartOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TextChart) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The URL of the chart.
func (o TextChartOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *TextChart) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type TextChartArrayOutput struct{ *pulumi.OutputState }

func (TextChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TextChart)(nil)).Elem()
}

func (o TextChartArrayOutput) ToTextChartArrayOutput() TextChartArrayOutput {
	return o
}

func (o TextChartArrayOutput) ToTextChartArrayOutputWithContext(ctx context.Context) TextChartArrayOutput {
	return o
}

func (o TextChartArrayOutput) Index(i pulumi.IntInput) TextChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TextChart {
		return vs[0].([]*TextChart)[vs[1].(int)]
	}).(TextChartOutput)
}

type TextChartMapOutput struct{ *pulumi.OutputState }

func (TextChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TextChart)(nil)).Elem()
}

func (o TextChartMapOutput) ToTextChartMapOutput() TextChartMapOutput {
	return o
}

func (o TextChartMapOutput) ToTextChartMapOutputWithContext(ctx context.Context) TextChartMapOutput {
	return o
}

func (o TextChartMapOutput) MapIndex(k pulumi.StringInput) TextChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TextChart {
		return vs[0].(map[string]*TextChart)[vs[1].(string)]
	}).(TextChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TextChartInput)(nil)).Elem(), &TextChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*TextChartArrayInput)(nil)).Elem(), TextChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TextChartMapInput)(nil)).Elem(), TextChartMap{})
	pulumi.RegisterOutputType(TextChartOutput{})
	pulumi.RegisterOutputType(TextChartArrayOutput{})
	pulumi.RegisterOutputType(TextChartMapOutput{})
}
