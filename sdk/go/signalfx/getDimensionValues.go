// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of dimension values matching the provided query.
//
// > **NOTE** This data source only allows 1000 values, as it's kinda nuts to make anything with `forEach` that big in SignalFx. This is negotiable.
func GetDimensionValues(ctx *pulumi.Context, args *GetDimensionValuesArgs, opts ...pulumi.InvokeOption) (*GetDimensionValuesResult, error) {
	var rv GetDimensionValuesResult
	err := ctx.Invoke("signalfx:index/getDimensionValues:getDimensionValues", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDimensionValues.
type GetDimensionValuesArgs struct {
	Query string `pulumi:"query"`
}

// A collection of values returned by getDimensionValues.
type GetDimensionValuesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id     string   `pulumi:"id"`
	Query  string   `pulumi:"query"`
	Values []string `pulumi:"values"`
}

func GetDimensionValuesOutput(ctx *pulumi.Context, args GetDimensionValuesOutputArgs, opts ...pulumi.InvokeOption) GetDimensionValuesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDimensionValuesResult, error) {
			args := v.(GetDimensionValuesArgs)
			r, err := GetDimensionValues(ctx, &args, opts...)
			return *r, err
		}).(GetDimensionValuesResultOutput)
}

// A collection of arguments for invoking getDimensionValues.
type GetDimensionValuesOutputArgs struct {
	Query pulumi.StringInput `pulumi:"query"`
}

func (GetDimensionValuesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDimensionValuesArgs)(nil)).Elem()
}

// A collection of values returned by getDimensionValues.
type GetDimensionValuesResultOutput struct{ *pulumi.OutputState }

func (GetDimensionValuesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDimensionValuesResult)(nil)).Elem()
}

func (o GetDimensionValuesResultOutput) ToGetDimensionValuesResultOutput() GetDimensionValuesResultOutput {
	return o
}

func (o GetDimensionValuesResultOutput) ToGetDimensionValuesResultOutputWithContext(ctx context.Context) GetDimensionValuesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetDimensionValuesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDimensionValuesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDimensionValuesResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v GetDimensionValuesResult) string { return v.Query }).(pulumi.StringOutput)
}

func (o GetDimensionValuesResultOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDimensionValuesResult) []string { return v.Values }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDimensionValuesResultOutput{})
}
