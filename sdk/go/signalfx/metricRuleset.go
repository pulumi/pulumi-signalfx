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

// Provides an Observability Cloud resource for managing metric rulesets.
//
// > **NOTE** When managing metric rulesets to drop data use a session token for an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
//
// ## Example
//
// ## Arguments
//
// The following arguments are supported in the resource block:
//
// * `metricName` - (Required) Name of the input metric
// * `aggregationRules` - (Optional) List of aggregation rules for the metric
//   - `enabled` - (Required) When false, this rule will not generate aggregated MTSs
//   - `name` - (Optional) name of the aggregation rule
//   - `matcher` - (Required) Matcher object
//   - `type` - (Required) Type of matcher. Must always be "dimension"
//   - `filters` - (Optional) List of filters to filter the set of input MTSs
//   - `property` - (Required) - Name of the dimension
//   - `propertyValue` - (Required) - Value of the dimension
//   - `not` - When true, this filter will match all values not matching the propertyValues
//   - `aggregator` - (Required) - Aggregator object
//   - `type` - (Required) Type of aggregator. Must always be "rollup"
//   - `dimensions` - (Required) List of dimensions to either be kept or dropped in the new aggregated MTSs
//   - `dropDimensions` - (Required) when true, the specified dimensions will be dropped from the aggregated MTSs
//   - `outputName` - (Required) name of the new aggregated metric
//
// * `routingRule` - (Required) Routing Rule object
//   - `destination` - (Required) - end destination of the input metric. Must be `RealTime` or `Drop`
type MetricRuleset struct {
	pulumi.CustomResourceState

	// Aggregation rules in the ruleset
	AggregationRules MetricRulesetAggregationRuleArrayOutput `pulumi:"aggregationRules"`
	// Timestamp of when the metric ruleset was created
	Created pulumi.StringOutput `pulumi:"created"`
	// ID of the creator of the metric ruleset
	Creator pulumi.StringOutput `pulumi:"creator"`
	// Timestamp of when the metric ruleset was last updated
	LastUpdated pulumi.StringOutput `pulumi:"lastUpdated"`
	// ID of user who last updated the metric ruleset
	LastUpdatedBy pulumi.StringOutput `pulumi:"lastUpdatedBy"`
	// Name of user who last updated this metric ruleset
	LastUpdatedByName pulumi.StringOutput `pulumi:"lastUpdatedByName"`
	// Name of the metric
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// Location to send the input metric
	RoutingRules MetricRulesetRoutingRuleArrayOutput `pulumi:"routingRules"`
	// Version of the ruleset
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewMetricRuleset registers a new resource with the given unique name, arguments, and options.
func NewMetricRuleset(ctx *pulumi.Context,
	name string, args *MetricRulesetArgs, opts ...pulumi.ResourceOption) (*MetricRuleset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	if args.RoutingRules == nil {
		return nil, errors.New("invalid value for required argument 'RoutingRules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MetricRuleset
	err := ctx.RegisterResource("signalfx:index/metricRuleset:MetricRuleset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricRuleset gets an existing MetricRuleset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricRuleset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricRulesetState, opts ...pulumi.ResourceOption) (*MetricRuleset, error) {
	var resource MetricRuleset
	err := ctx.ReadResource("signalfx:index/metricRuleset:MetricRuleset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricRuleset resources.
type metricRulesetState struct {
	// Aggregation rules in the ruleset
	AggregationRules []MetricRulesetAggregationRule `pulumi:"aggregationRules"`
	// Timestamp of when the metric ruleset was created
	Created *string `pulumi:"created"`
	// ID of the creator of the metric ruleset
	Creator *string `pulumi:"creator"`
	// Timestamp of when the metric ruleset was last updated
	LastUpdated *string `pulumi:"lastUpdated"`
	// ID of user who last updated the metric ruleset
	LastUpdatedBy *string `pulumi:"lastUpdatedBy"`
	// Name of user who last updated this metric ruleset
	LastUpdatedByName *string `pulumi:"lastUpdatedByName"`
	// Name of the metric
	MetricName *string `pulumi:"metricName"`
	// Location to send the input metric
	RoutingRules []MetricRulesetRoutingRule `pulumi:"routingRules"`
	// Version of the ruleset
	Version *string `pulumi:"version"`
}

type MetricRulesetState struct {
	// Aggregation rules in the ruleset
	AggregationRules MetricRulesetAggregationRuleArrayInput
	// Timestamp of when the metric ruleset was created
	Created pulumi.StringPtrInput
	// ID of the creator of the metric ruleset
	Creator pulumi.StringPtrInput
	// Timestamp of when the metric ruleset was last updated
	LastUpdated pulumi.StringPtrInput
	// ID of user who last updated the metric ruleset
	LastUpdatedBy pulumi.StringPtrInput
	// Name of user who last updated this metric ruleset
	LastUpdatedByName pulumi.StringPtrInput
	// Name of the metric
	MetricName pulumi.StringPtrInput
	// Location to send the input metric
	RoutingRules MetricRulesetRoutingRuleArrayInput
	// Version of the ruleset
	Version pulumi.StringPtrInput
}

func (MetricRulesetState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricRulesetState)(nil)).Elem()
}

type metricRulesetArgs struct {
	// Aggregation rules in the ruleset
	AggregationRules []MetricRulesetAggregationRule `pulumi:"aggregationRules"`
	// Name of the metric
	MetricName string `pulumi:"metricName"`
	// Location to send the input metric
	RoutingRules []MetricRulesetRoutingRule `pulumi:"routingRules"`
}

// The set of arguments for constructing a MetricRuleset resource.
type MetricRulesetArgs struct {
	// Aggregation rules in the ruleset
	AggregationRules MetricRulesetAggregationRuleArrayInput
	// Name of the metric
	MetricName pulumi.StringInput
	// Location to send the input metric
	RoutingRules MetricRulesetRoutingRuleArrayInput
}

func (MetricRulesetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricRulesetArgs)(nil)).Elem()
}

type MetricRulesetInput interface {
	pulumi.Input

	ToMetricRulesetOutput() MetricRulesetOutput
	ToMetricRulesetOutputWithContext(ctx context.Context) MetricRulesetOutput
}

func (*MetricRuleset) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricRuleset)(nil)).Elem()
}

func (i *MetricRuleset) ToMetricRulesetOutput() MetricRulesetOutput {
	return i.ToMetricRulesetOutputWithContext(context.Background())
}

func (i *MetricRuleset) ToMetricRulesetOutputWithContext(ctx context.Context) MetricRulesetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricRulesetOutput)
}

// MetricRulesetArrayInput is an input type that accepts MetricRulesetArray and MetricRulesetArrayOutput values.
// You can construct a concrete instance of `MetricRulesetArrayInput` via:
//
//	MetricRulesetArray{ MetricRulesetArgs{...} }
type MetricRulesetArrayInput interface {
	pulumi.Input

	ToMetricRulesetArrayOutput() MetricRulesetArrayOutput
	ToMetricRulesetArrayOutputWithContext(context.Context) MetricRulesetArrayOutput
}

type MetricRulesetArray []MetricRulesetInput

func (MetricRulesetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetricRuleset)(nil)).Elem()
}

func (i MetricRulesetArray) ToMetricRulesetArrayOutput() MetricRulesetArrayOutput {
	return i.ToMetricRulesetArrayOutputWithContext(context.Background())
}

func (i MetricRulesetArray) ToMetricRulesetArrayOutputWithContext(ctx context.Context) MetricRulesetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricRulesetArrayOutput)
}

// MetricRulesetMapInput is an input type that accepts MetricRulesetMap and MetricRulesetMapOutput values.
// You can construct a concrete instance of `MetricRulesetMapInput` via:
//
//	MetricRulesetMap{ "key": MetricRulesetArgs{...} }
type MetricRulesetMapInput interface {
	pulumi.Input

	ToMetricRulesetMapOutput() MetricRulesetMapOutput
	ToMetricRulesetMapOutputWithContext(context.Context) MetricRulesetMapOutput
}

type MetricRulesetMap map[string]MetricRulesetInput

func (MetricRulesetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetricRuleset)(nil)).Elem()
}

func (i MetricRulesetMap) ToMetricRulesetMapOutput() MetricRulesetMapOutput {
	return i.ToMetricRulesetMapOutputWithContext(context.Background())
}

func (i MetricRulesetMap) ToMetricRulesetMapOutputWithContext(ctx context.Context) MetricRulesetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricRulesetMapOutput)
}

type MetricRulesetOutput struct{ *pulumi.OutputState }

func (MetricRulesetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricRuleset)(nil)).Elem()
}

func (o MetricRulesetOutput) ToMetricRulesetOutput() MetricRulesetOutput {
	return o
}

func (o MetricRulesetOutput) ToMetricRulesetOutputWithContext(ctx context.Context) MetricRulesetOutput {
	return o
}

// Aggregation rules in the ruleset
func (o MetricRulesetOutput) AggregationRules() MetricRulesetAggregationRuleArrayOutput {
	return o.ApplyT(func(v *MetricRuleset) MetricRulesetAggregationRuleArrayOutput { return v.AggregationRules }).(MetricRulesetAggregationRuleArrayOutput)
}

// Timestamp of when the metric ruleset was created
func (o MetricRulesetOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricRuleset) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// ID of the creator of the metric ruleset
func (o MetricRulesetOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricRuleset) pulumi.StringOutput { return v.Creator }).(pulumi.StringOutput)
}

// Timestamp of when the metric ruleset was last updated
func (o MetricRulesetOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricRuleset) pulumi.StringOutput { return v.LastUpdated }).(pulumi.StringOutput)
}

// ID of user who last updated the metric ruleset
func (o MetricRulesetOutput) LastUpdatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricRuleset) pulumi.StringOutput { return v.LastUpdatedBy }).(pulumi.StringOutput)
}

// Name of user who last updated this metric ruleset
func (o MetricRulesetOutput) LastUpdatedByName() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricRuleset) pulumi.StringOutput { return v.LastUpdatedByName }).(pulumi.StringOutput)
}

// Name of the metric
func (o MetricRulesetOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricRuleset) pulumi.StringOutput { return v.MetricName }).(pulumi.StringOutput)
}

// Location to send the input metric
func (o MetricRulesetOutput) RoutingRules() MetricRulesetRoutingRuleArrayOutput {
	return o.ApplyT(func(v *MetricRuleset) MetricRulesetRoutingRuleArrayOutput { return v.RoutingRules }).(MetricRulesetRoutingRuleArrayOutput)
}

// Version of the ruleset
func (o MetricRulesetOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricRuleset) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type MetricRulesetArrayOutput struct{ *pulumi.OutputState }

func (MetricRulesetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MetricRuleset)(nil)).Elem()
}

func (o MetricRulesetArrayOutput) ToMetricRulesetArrayOutput() MetricRulesetArrayOutput {
	return o
}

func (o MetricRulesetArrayOutput) ToMetricRulesetArrayOutputWithContext(ctx context.Context) MetricRulesetArrayOutput {
	return o
}

func (o MetricRulesetArrayOutput) Index(i pulumi.IntInput) MetricRulesetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MetricRuleset {
		return vs[0].([]*MetricRuleset)[vs[1].(int)]
	}).(MetricRulesetOutput)
}

type MetricRulesetMapOutput struct{ *pulumi.OutputState }

func (MetricRulesetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MetricRuleset)(nil)).Elem()
}

func (o MetricRulesetMapOutput) ToMetricRulesetMapOutput() MetricRulesetMapOutput {
	return o
}

func (o MetricRulesetMapOutput) ToMetricRulesetMapOutputWithContext(ctx context.Context) MetricRulesetMapOutput {
	return o
}

func (o MetricRulesetMapOutput) MapIndex(k pulumi.StringInput) MetricRulesetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MetricRuleset {
		return vs[0].(map[string]*MetricRuleset)[vs[1].(string)]
	}).(MetricRulesetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetricRulesetInput)(nil)).Elem(), &MetricRuleset{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricRulesetArrayInput)(nil)).Elem(), MetricRulesetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricRulesetMapInput)(nil)).Elem(), MetricRulesetMap{})
	pulumi.RegisterOutputType(MetricRulesetOutput{})
	pulumi.RegisterOutputType(MetricRulesetArrayOutput{})
	pulumi.RegisterOutputType(MetricRulesetMapOutput{})
}
