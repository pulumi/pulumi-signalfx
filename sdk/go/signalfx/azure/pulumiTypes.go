// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationCustomNamespacesPerService struct {
	// The additional namespaces.
	Namespaces []string `pulumi:"namespaces"`
	// The name of the service.
	Service string `pulumi:"service"`
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
	// The additional namespaces.
	Namespaces pulumi.StringArrayInput `pulumi:"namespaces"`
	// The name of the service.
	Service pulumi.StringInput `pulumi:"service"`
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

// The additional namespaces.
func (o IntegrationCustomNamespacesPerServiceOutput) Namespaces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IntegrationCustomNamespacesPerService) []string { return v.Namespaces }).(pulumi.StringArrayOutput)
}

// The name of the service.
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
	Filter IntegrationResourceFilterRuleFilter `pulumi:"filter"`
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
	Filter IntegrationResourceFilterRuleFilterInput `pulumi:"filter"`
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

func (o IntegrationResourceFilterRuleOutput) Filter() IntegrationResourceFilterRuleFilterOutput {
	return o.ApplyT(func(v IntegrationResourceFilterRule) IntegrationResourceFilterRuleFilter { return v.Filter }).(IntegrationResourceFilterRuleFilterOutput)
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

type IntegrationResourceFilterRuleFilter struct {
	Source string `pulumi:"source"`
}

// IntegrationResourceFilterRuleFilterInput is an input type that accepts IntegrationResourceFilterRuleFilterArgs and IntegrationResourceFilterRuleFilterOutput values.
// You can construct a concrete instance of `IntegrationResourceFilterRuleFilterInput` via:
//
//	IntegrationResourceFilterRuleFilterArgs{...}
type IntegrationResourceFilterRuleFilterInput interface {
	pulumi.Input

	ToIntegrationResourceFilterRuleFilterOutput() IntegrationResourceFilterRuleFilterOutput
	ToIntegrationResourceFilterRuleFilterOutputWithContext(context.Context) IntegrationResourceFilterRuleFilterOutput
}

type IntegrationResourceFilterRuleFilterArgs struct {
	Source pulumi.StringInput `pulumi:"source"`
}

func (IntegrationResourceFilterRuleFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationResourceFilterRuleFilter)(nil)).Elem()
}

func (i IntegrationResourceFilterRuleFilterArgs) ToIntegrationResourceFilterRuleFilterOutput() IntegrationResourceFilterRuleFilterOutput {
	return i.ToIntegrationResourceFilterRuleFilterOutputWithContext(context.Background())
}

func (i IntegrationResourceFilterRuleFilterArgs) ToIntegrationResourceFilterRuleFilterOutputWithContext(ctx context.Context) IntegrationResourceFilterRuleFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationResourceFilterRuleFilterOutput)
}

type IntegrationResourceFilterRuleFilterOutput struct{ *pulumi.OutputState }

func (IntegrationResourceFilterRuleFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationResourceFilterRuleFilter)(nil)).Elem()
}

func (o IntegrationResourceFilterRuleFilterOutput) ToIntegrationResourceFilterRuleFilterOutput() IntegrationResourceFilterRuleFilterOutput {
	return o
}

func (o IntegrationResourceFilterRuleFilterOutput) ToIntegrationResourceFilterRuleFilterOutputWithContext(ctx context.Context) IntegrationResourceFilterRuleFilterOutput {
	return o
}

func (o IntegrationResourceFilterRuleFilterOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationResourceFilterRuleFilter) string { return v.Source }).(pulumi.StringOutput)
}

type GetServicesService struct {
	Name string `pulumi:"name"`
}

// GetServicesServiceInput is an input type that accepts GetServicesServiceArgs and GetServicesServiceOutput values.
// You can construct a concrete instance of `GetServicesServiceInput` via:
//
//	GetServicesServiceArgs{...}
type GetServicesServiceInput interface {
	pulumi.Input

	ToGetServicesServiceOutput() GetServicesServiceOutput
	ToGetServicesServiceOutputWithContext(context.Context) GetServicesServiceOutput
}

type GetServicesServiceArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetServicesServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServicesService)(nil)).Elem()
}

func (i GetServicesServiceArgs) ToGetServicesServiceOutput() GetServicesServiceOutput {
	return i.ToGetServicesServiceOutputWithContext(context.Background())
}

func (i GetServicesServiceArgs) ToGetServicesServiceOutputWithContext(ctx context.Context) GetServicesServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServicesServiceOutput)
}

// GetServicesServiceArrayInput is an input type that accepts GetServicesServiceArray and GetServicesServiceArrayOutput values.
// You can construct a concrete instance of `GetServicesServiceArrayInput` via:
//
//	GetServicesServiceArray{ GetServicesServiceArgs{...} }
type GetServicesServiceArrayInput interface {
	pulumi.Input

	ToGetServicesServiceArrayOutput() GetServicesServiceArrayOutput
	ToGetServicesServiceArrayOutputWithContext(context.Context) GetServicesServiceArrayOutput
}

type GetServicesServiceArray []GetServicesServiceInput

func (GetServicesServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServicesService)(nil)).Elem()
}

func (i GetServicesServiceArray) ToGetServicesServiceArrayOutput() GetServicesServiceArrayOutput {
	return i.ToGetServicesServiceArrayOutputWithContext(context.Background())
}

func (i GetServicesServiceArray) ToGetServicesServiceArrayOutputWithContext(ctx context.Context) GetServicesServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServicesServiceArrayOutput)
}

type GetServicesServiceOutput struct{ *pulumi.OutputState }

func (GetServicesServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServicesService)(nil)).Elem()
}

func (o GetServicesServiceOutput) ToGetServicesServiceOutput() GetServicesServiceOutput {
	return o
}

func (o GetServicesServiceOutput) ToGetServicesServiceOutputWithContext(ctx context.Context) GetServicesServiceOutput {
	return o
}

func (o GetServicesServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetServicesService) string { return v.Name }).(pulumi.StringOutput)
}

type GetServicesServiceArrayOutput struct{ *pulumi.OutputState }

func (GetServicesServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServicesService)(nil)).Elem()
}

func (o GetServicesServiceArrayOutput) ToGetServicesServiceArrayOutput() GetServicesServiceArrayOutput {
	return o
}

func (o GetServicesServiceArrayOutput) ToGetServicesServiceArrayOutputWithContext(ctx context.Context) GetServicesServiceArrayOutput {
	return o
}

func (o GetServicesServiceArrayOutput) Index(i pulumi.IntInput) GetServicesServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetServicesService {
		return vs[0].([]GetServicesService)[vs[1].(int)]
	}).(GetServicesServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationCustomNamespacesPerServiceInput)(nil)).Elem(), IntegrationCustomNamespacesPerServiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationCustomNamespacesPerServiceArrayInput)(nil)).Elem(), IntegrationCustomNamespacesPerServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationResourceFilterRuleInput)(nil)).Elem(), IntegrationResourceFilterRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationResourceFilterRuleArrayInput)(nil)).Elem(), IntegrationResourceFilterRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationResourceFilterRuleFilterInput)(nil)).Elem(), IntegrationResourceFilterRuleFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServicesServiceInput)(nil)).Elem(), GetServicesServiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServicesServiceArrayInput)(nil)).Elem(), GetServicesServiceArray{})
	pulumi.RegisterOutputType(IntegrationCustomNamespacesPerServiceOutput{})
	pulumi.RegisterOutputType(IntegrationCustomNamespacesPerServiceArrayOutput{})
	pulumi.RegisterOutputType(IntegrationResourceFilterRuleOutput{})
	pulumi.RegisterOutputType(IntegrationResourceFilterRuleArrayOutput{})
	pulumi.RegisterOutputType(IntegrationResourceFilterRuleFilterOutput{})
	pulumi.RegisterOutputType(GetServicesServiceOutput{})
	pulumi.RegisterOutputType(GetServicesServiceArrayOutput{})
}
