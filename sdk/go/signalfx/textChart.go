// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This special type of chart doesn’t display any metric data. Rather, it lets you place a text note on the dashboard.
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
// 		_, err := signalfx.NewTextChart(ctx, "mynote0", &signalfx.TextChartArgs{
// 			Description: pulumi.String("Lorem ipsum dolor sit amet, laudem tibique iracundia at mea. Nam posse dolores ex, nec cu adhuc putent honestatis"),
// 			Markdown:    pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "    1. First ordered list item\n", "    2. Another item\n", "      * Unordered sub-list.\n", "    1. Actual numbers don't matter, just that it's a number\n", "      1. Ordered sub-list\n", "    4. And another item.\n", "\n", "       You can have properly indented paragraphs within list items. Notice the blank line above, and the leading spaces (at least one, but we'll use three here to also align the raw Markdown).\n", "\n", "       To have a line break without a paragraph, you will need to use two trailing spaces.⋅⋅\n", "       Note that this line is separate, but within the same paragraph.⋅⋅\n", "       (This is contrary to the typical GFM line break behaviour, where trailing spaces are not required.)\n", "\n", "    * Unordered list can use asterisks\n", "    - Or minuses\n", "    + Or pluses\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type TextChart struct {
	pulumi.CustomResourceState

	// Description of the text note.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Markdown text to display.
	Markdown pulumi.StringOutput `pulumi:"markdown"`
	// Name of the text note.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the chart.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewTextChart registers a new resource with the given unique name, arguments, and options.
func NewTextChart(ctx *pulumi.Context,
	name string, args *TextChartArgs, opts ...pulumi.ResourceOption) (*TextChart, error) {
	if args == nil || args.Markdown == nil {
		return nil, errors.New("missing required argument 'Markdown'")
	}
	if args == nil {
		args = &TextChartArgs{}
	}
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
}

// The set of arguments for constructing a TextChart resource.
type TextChartArgs struct {
	// Description of the text note.
	Description pulumi.StringPtrInput
	// Markdown text to display.
	Markdown pulumi.StringInput
	// Name of the text note.
	Name pulumi.StringPtrInput
}

func (TextChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*textChartArgs)(nil)).Elem()
}
