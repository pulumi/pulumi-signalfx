// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type IntegrationCustomNamespacesPerService struct {
	Namespaces []string `pulumi:"namespaces"`
	Service    string   `pulumi:"service"`
}

// IntegrationCustomNamespacesPerServiceInput is an input type that accepts IntegrationCustomNamespacesPerServiceArgs and IntegrationCustomNamespacesPerServiceOutput values.
// You can construct a concrete instance of `IntegrationCustomNamespacesPerServiceInput` via:
//
//	IntegrationCustomNamespacesPerServiceArgs{...}
type IntegrationCustomNamespacesPerServiceInput interface {
	pulumi.Input

	ToIntegrationCustomNamespacesPerServiceOutput() IntegrationCustomNamespacesPerServiceOutput
	ToIntegrationCustomNamespacesPerServiceOutputWithContext(context.Context) IntegrationCustomNamespacesPerServiceOutput
}

type IntegrationCustomNamespacesPerServiceArgs struct {
	Namespaces pulumi.StringArrayInput `pulumi:"namespaces"`
	Service    pulumi.StringInput      `pulumi:"service"`
}

func (IntegrationCustomNamespacesPerServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationCustomNamespacesPerService)(nil)).Elem()
}

func (i IntegrationCustomNamespacesPerServiceArgs) ToIntegrationCustomNamespacesPerServiceOutput() IntegrationCustomNamespacesPerServiceOutput {
	return i.ToIntegrationCustomNamespacesPerServiceOutputWithContext(context.Background())
}

func (i IntegrationCustomNamespacesPerServiceArgs) ToIntegrationCustomNamespacesPerServiceOutputWithContext(ctx context.Context) IntegrationCustomNamespacesPerServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationCustomNamespacesPerServiceOutput)
}

// IntegrationCustomNamespacesPerServiceArrayInput is an input type that accepts IntegrationCustomNamespacesPerServiceArray and IntegrationCustomNamespacesPerServiceArrayOutput values.
// You can construct a concrete instance of `IntegrationCustomNamespacesPerServiceArrayInput` via:
//
//	IntegrationCustomNamespacesPerServiceArray{ IntegrationCustomNamespacesPerServiceArgs{...} }
type IntegrationCustomNamespacesPerServiceArrayInput interface {
	pulumi.Input

	ToIntegrationCustomNamespacesPerServiceArrayOutput() IntegrationCustomNamespacesPerServiceArrayOutput
	ToIntegrationCustomNamespacesPerServiceArrayOutputWithContext(context.Context) IntegrationCustomNamespacesPerServiceArrayOutput
}

type IntegrationCustomNamespacesPerServiceArray []IntegrationCustomNamespacesPerServiceInput

func (IntegrationCustomNamespacesPerServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationCustomNamespacesPerService)(nil)).Elem()
}

func (i IntegrationCustomNamespacesPerServiceArray) ToIntegrationCustomNamespacesPerServiceArrayOutput() IntegrationCustomNamespacesPerServiceArrayOutput {
	return i.ToIntegrationCustomNamespacesPerServiceArrayOutputWithContext(context.Background())
}

func (i IntegrationCustomNamespacesPerServiceArray) ToIntegrationCustomNamespacesPerServiceArrayOutputWithContext(ctx context.Context) IntegrationCustomNamespacesPerServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationCustomNamespacesPerServiceArrayOutput)
}

type IntegrationCustomNamespacesPerServiceOutput struct{ *pulumi.OutputState }

func (IntegrationCustomNamespacesPerServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationCustomNamespacesPerService)(nil)).Elem()
}

func (o IntegrationCustomNamespacesPerServiceOutput) ToIntegrationCustomNamespacesPerServiceOutput() IntegrationCustomNamespacesPerServiceOutput {
	return o
}

func (o IntegrationCustomNamespacesPerServiceOutput) ToIntegrationCustomNamespacesPerServiceOutputWithContext(ctx context.Context) IntegrationCustomNamespacesPerServiceOutput {
	return o
}

func (o IntegrationCustomNamespacesPerServiceOutput) Namespaces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IntegrationCustomNamespacesPerService) []string { return v.Namespaces }).(pulumi.StringArrayOutput)
}

func (o IntegrationCustomNamespacesPerServiceOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationCustomNamespacesPerService) string { return v.Service }).(pulumi.StringOutput)
}

type IntegrationCustomNamespacesPerServiceArrayOutput struct{ *pulumi.OutputState }

func (IntegrationCustomNamespacesPerServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationCustomNamespacesPerService)(nil)).Elem()
}

func (o IntegrationCustomNamespacesPerServiceArrayOutput) ToIntegrationCustomNamespacesPerServiceArrayOutput() IntegrationCustomNamespacesPerServiceArrayOutput {
	return o
}

func (o IntegrationCustomNamespacesPerServiceArrayOutput) ToIntegrationCustomNamespacesPerServiceArrayOutputWithContext(ctx context.Context) IntegrationCustomNamespacesPerServiceArrayOutput {
	return o
}

func (o IntegrationCustomNamespacesPerServiceArrayOutput) Index(i pulumi.IntInput) IntegrationCustomNamespacesPerServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationCustomNamespacesPerService {
		return vs[0].([]IntegrationCustomNamespacesPerService)[vs[1].(int)]
	}).(IntegrationCustomNamespacesPerServiceOutput)
}

type IntegrationResourceFilterRule struct {
	FilterSource string `pulumi:"filterSource"`
}

// IntegrationResourceFilterRuleInput is an input type that accepts IntegrationResourceFilterRuleArgs and IntegrationResourceFilterRuleOutput values.
// You can construct a concrete instance of `IntegrationResourceFilterRuleInput` via:
//
//	IntegrationResourceFilterRuleArgs{...}
type IntegrationResourceFilterRuleInput interface {
	pulumi.Input

	ToIntegrationResourceFilterRuleOutput() IntegrationResourceFilterRuleOutput
	ToIntegrationResourceFilterRuleOutputWithContext(context.Context) IntegrationResourceFilterRuleOutput
}

type IntegrationResourceFilterRuleArgs struct {
	FilterSource pulumi.StringInput `pulumi:"filterSource"`
}

func (IntegrationResourceFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationResourceFilterRule)(nil)).Elem()
}

func (i IntegrationResourceFilterRuleArgs) ToIntegrationResourceFilterRuleOutput() IntegrationResourceFilterRuleOutput {
	return i.ToIntegrationResourceFilterRuleOutputWithContext(context.Background())
}

func (i IntegrationResourceFilterRuleArgs) ToIntegrationResourceFilterRuleOutputWithContext(ctx context.Context) IntegrationResourceFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationResourceFilterRuleOutput)
}

// IntegrationResourceFilterRuleArrayInput is an input type that accepts IntegrationResourceFilterRuleArray and IntegrationResourceFilterRuleArrayOutput values.
// You can construct a concrete instance of `IntegrationResourceFilterRuleArrayInput` via:
//
//	IntegrationResourceFilterRuleArray{ IntegrationResourceFilterRuleArgs{...} }
type IntegrationResourceFilterRuleArrayInput interface {
	pulumi.Input

	ToIntegrationResourceFilterRuleArrayOutput() IntegrationResourceFilterRuleArrayOutput
	ToIntegrationResourceFilterRuleArrayOutputWithContext(context.Context) IntegrationResourceFilterRuleArrayOutput
}

type IntegrationResourceFilterRuleArray []IntegrationResourceFilterRuleInput

func (IntegrationResourceFilterRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationResourceFilterRule)(nil)).Elem()
}

func (i IntegrationResourceFilterRuleArray) ToIntegrationResourceFilterRuleArrayOutput() IntegrationResourceFilterRuleArrayOutput {
	return i.ToIntegrationResourceFilterRuleArrayOutputWithContext(context.Background())
}

func (i IntegrationResourceFilterRuleArray) ToIntegrationResourceFilterRuleArrayOutputWithContext(ctx context.Context) IntegrationResourceFilterRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationResourceFilterRuleArrayOutput)
}

type IntegrationResourceFilterRuleOutput struct{ *pulumi.OutputState }

func (IntegrationResourceFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationResourceFilterRule)(nil)).Elem()
}

func (o IntegrationResourceFilterRuleOutput) ToIntegrationResourceFilterRuleOutput() IntegrationResourceFilterRuleOutput {
	return o
}

func (o IntegrationResourceFilterRuleOutput) ToIntegrationResourceFilterRuleOutputWithContext(ctx context.Context) IntegrationResourceFilterRuleOutput {
	return o
}

func (o IntegrationResourceFilterRuleOutput) FilterSource() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationResourceFilterRule) string { return v.FilterSource }).(pulumi.StringOutput)
}

type IntegrationResourceFilterRuleArrayOutput struct{ *pulumi.OutputState }

func (IntegrationResourceFilterRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationResourceFilterRule)(nil)).Elem()
}

func (o IntegrationResourceFilterRuleArrayOutput) ToIntegrationResourceFilterRuleArrayOutput() IntegrationResourceFilterRuleArrayOutput {
	return o
}

func (o IntegrationResourceFilterRuleArrayOutput) ToIntegrationResourceFilterRuleArrayOutputWithContext(ctx context.Context) IntegrationResourceFilterRuleArrayOutput {
	return o
}

func (o IntegrationResourceFilterRuleArrayOutput) Index(i pulumi.IntInput) IntegrationResourceFilterRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationResourceFilterRule {
		return vs[0].([]IntegrationResourceFilterRule)[vs[1].(int)]
	}).(IntegrationResourceFilterRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationCustomNamespacesPerServiceInput)(nil)).Elem(), IntegrationCustomNamespacesPerServiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationCustomNamespacesPerServiceArrayInput)(nil)).Elem(), IntegrationCustomNamespacesPerServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationResourceFilterRuleInput)(nil)).Elem(), IntegrationResourceFilterRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationResourceFilterRuleArrayInput)(nil)).Elem(), IntegrationResourceFilterRuleArray{})
	pulumi.RegisterOutputType(IntegrationCustomNamespacesPerServiceOutput{})
	pulumi.RegisterOutputType(IntegrationCustomNamespacesPerServiceArrayOutput{})
	pulumi.RegisterOutputType(IntegrationResourceFilterRuleOutput{})
	pulumi.RegisterOutputType(IntegrationResourceFilterRuleArrayOutput{})
}
