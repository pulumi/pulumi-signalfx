// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationCustomNamespaceSyncRule struct {
	// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
	DefaultAction *string `pulumi:"defaultAction"`
	// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
	FilterAction *string `pulumi:"filterAction"`
	// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
	FilterSource *string `pulumi:"filterSource"`
	// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
	Namespace string `pulumi:"namespace"`
}

// IntegrationCustomNamespaceSyncRuleInput is an input type that accepts IntegrationCustomNamespaceSyncRuleArgs and IntegrationCustomNamespaceSyncRuleOutput values.
// You can construct a concrete instance of `IntegrationCustomNamespaceSyncRuleInput` via:
//
//          IntegrationCustomNamespaceSyncRuleArgs{...}
type IntegrationCustomNamespaceSyncRuleInput interface {
	pulumi.Input

	ToIntegrationCustomNamespaceSyncRuleOutput() IntegrationCustomNamespaceSyncRuleOutput
	ToIntegrationCustomNamespaceSyncRuleOutputWithContext(context.Context) IntegrationCustomNamespaceSyncRuleOutput
}

type IntegrationCustomNamespaceSyncRuleArgs struct {
	// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
	DefaultAction pulumi.StringPtrInput `pulumi:"defaultAction"`
	// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
	FilterAction pulumi.StringPtrInput `pulumi:"filterAction"`
	// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
	FilterSource pulumi.StringPtrInput `pulumi:"filterSource"`
	// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
	Namespace pulumi.StringInput `pulumi:"namespace"`
}

func (IntegrationCustomNamespaceSyncRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationCustomNamespaceSyncRule)(nil)).Elem()
}

func (i IntegrationCustomNamespaceSyncRuleArgs) ToIntegrationCustomNamespaceSyncRuleOutput() IntegrationCustomNamespaceSyncRuleOutput {
	return i.ToIntegrationCustomNamespaceSyncRuleOutputWithContext(context.Background())
}

func (i IntegrationCustomNamespaceSyncRuleArgs) ToIntegrationCustomNamespaceSyncRuleOutputWithContext(ctx context.Context) IntegrationCustomNamespaceSyncRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationCustomNamespaceSyncRuleOutput)
}

// IntegrationCustomNamespaceSyncRuleArrayInput is an input type that accepts IntegrationCustomNamespaceSyncRuleArray and IntegrationCustomNamespaceSyncRuleArrayOutput values.
// You can construct a concrete instance of `IntegrationCustomNamespaceSyncRuleArrayInput` via:
//
//          IntegrationCustomNamespaceSyncRuleArray{ IntegrationCustomNamespaceSyncRuleArgs{...} }
type IntegrationCustomNamespaceSyncRuleArrayInput interface {
	pulumi.Input

	ToIntegrationCustomNamespaceSyncRuleArrayOutput() IntegrationCustomNamespaceSyncRuleArrayOutput
	ToIntegrationCustomNamespaceSyncRuleArrayOutputWithContext(context.Context) IntegrationCustomNamespaceSyncRuleArrayOutput
}

type IntegrationCustomNamespaceSyncRuleArray []IntegrationCustomNamespaceSyncRuleInput

func (IntegrationCustomNamespaceSyncRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationCustomNamespaceSyncRule)(nil)).Elem()
}

func (i IntegrationCustomNamespaceSyncRuleArray) ToIntegrationCustomNamespaceSyncRuleArrayOutput() IntegrationCustomNamespaceSyncRuleArrayOutput {
	return i.ToIntegrationCustomNamespaceSyncRuleArrayOutputWithContext(context.Background())
}

func (i IntegrationCustomNamespaceSyncRuleArray) ToIntegrationCustomNamespaceSyncRuleArrayOutputWithContext(ctx context.Context) IntegrationCustomNamespaceSyncRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationCustomNamespaceSyncRuleArrayOutput)
}

type IntegrationCustomNamespaceSyncRuleOutput struct{ *pulumi.OutputState }

func (IntegrationCustomNamespaceSyncRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationCustomNamespaceSyncRule)(nil)).Elem()
}

func (o IntegrationCustomNamespaceSyncRuleOutput) ToIntegrationCustomNamespaceSyncRuleOutput() IntegrationCustomNamespaceSyncRuleOutput {
	return o
}

func (o IntegrationCustomNamespaceSyncRuleOutput) ToIntegrationCustomNamespaceSyncRuleOutputWithContext(ctx context.Context) IntegrationCustomNamespaceSyncRuleOutput {
	return o
}

// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
func (o IntegrationCustomNamespaceSyncRuleOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationCustomNamespaceSyncRule) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
func (o IntegrationCustomNamespaceSyncRuleOutput) FilterAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationCustomNamespaceSyncRule) *string { return v.FilterAction }).(pulumi.StringPtrOutput)
}

// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
func (o IntegrationCustomNamespaceSyncRuleOutput) FilterSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationCustomNamespaceSyncRule) *string { return v.FilterSource }).(pulumi.StringPtrOutput)
}

// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
func (o IntegrationCustomNamespaceSyncRuleOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationCustomNamespaceSyncRule) string { return v.Namespace }).(pulumi.StringOutput)
}

type IntegrationCustomNamespaceSyncRuleArrayOutput struct{ *pulumi.OutputState }

func (IntegrationCustomNamespaceSyncRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationCustomNamespaceSyncRule)(nil)).Elem()
}

func (o IntegrationCustomNamespaceSyncRuleArrayOutput) ToIntegrationCustomNamespaceSyncRuleArrayOutput() IntegrationCustomNamespaceSyncRuleArrayOutput {
	return o
}

func (o IntegrationCustomNamespaceSyncRuleArrayOutput) ToIntegrationCustomNamespaceSyncRuleArrayOutputWithContext(ctx context.Context) IntegrationCustomNamespaceSyncRuleArrayOutput {
	return o
}

func (o IntegrationCustomNamespaceSyncRuleArrayOutput) Index(i pulumi.IntInput) IntegrationCustomNamespaceSyncRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationCustomNamespaceSyncRule {
		return vs[0].([]IntegrationCustomNamespaceSyncRule)[vs[1].(int)]
	}).(IntegrationCustomNamespaceSyncRuleOutput)
}

type IntegrationNamespaceSyncRule struct {
	// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
	DefaultAction *string `pulumi:"defaultAction"`
	// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
	FilterAction *string `pulumi:"filterAction"`
	// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
	FilterSource *string `pulumi:"filterSource"`
	// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
	Namespace string `pulumi:"namespace"`
}

// IntegrationNamespaceSyncRuleInput is an input type that accepts IntegrationNamespaceSyncRuleArgs and IntegrationNamespaceSyncRuleOutput values.
// You can construct a concrete instance of `IntegrationNamespaceSyncRuleInput` via:
//
//          IntegrationNamespaceSyncRuleArgs{...}
type IntegrationNamespaceSyncRuleInput interface {
	pulumi.Input

	ToIntegrationNamespaceSyncRuleOutput() IntegrationNamespaceSyncRuleOutput
	ToIntegrationNamespaceSyncRuleOutputWithContext(context.Context) IntegrationNamespaceSyncRuleOutput
}

type IntegrationNamespaceSyncRuleArgs struct {
	// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
	DefaultAction pulumi.StringPtrInput `pulumi:"defaultAction"`
	// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
	FilterAction pulumi.StringPtrInput `pulumi:"filterAction"`
	// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
	FilterSource pulumi.StringPtrInput `pulumi:"filterSource"`
	// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
	Namespace pulumi.StringInput `pulumi:"namespace"`
}

func (IntegrationNamespaceSyncRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationNamespaceSyncRule)(nil)).Elem()
}

func (i IntegrationNamespaceSyncRuleArgs) ToIntegrationNamespaceSyncRuleOutput() IntegrationNamespaceSyncRuleOutput {
	return i.ToIntegrationNamespaceSyncRuleOutputWithContext(context.Background())
}

func (i IntegrationNamespaceSyncRuleArgs) ToIntegrationNamespaceSyncRuleOutputWithContext(ctx context.Context) IntegrationNamespaceSyncRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationNamespaceSyncRuleOutput)
}

// IntegrationNamespaceSyncRuleArrayInput is an input type that accepts IntegrationNamespaceSyncRuleArray and IntegrationNamespaceSyncRuleArrayOutput values.
// You can construct a concrete instance of `IntegrationNamespaceSyncRuleArrayInput` via:
//
//          IntegrationNamespaceSyncRuleArray{ IntegrationNamespaceSyncRuleArgs{...} }
type IntegrationNamespaceSyncRuleArrayInput interface {
	pulumi.Input

	ToIntegrationNamespaceSyncRuleArrayOutput() IntegrationNamespaceSyncRuleArrayOutput
	ToIntegrationNamespaceSyncRuleArrayOutputWithContext(context.Context) IntegrationNamespaceSyncRuleArrayOutput
}

type IntegrationNamespaceSyncRuleArray []IntegrationNamespaceSyncRuleInput

func (IntegrationNamespaceSyncRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationNamespaceSyncRule)(nil)).Elem()
}

func (i IntegrationNamespaceSyncRuleArray) ToIntegrationNamespaceSyncRuleArrayOutput() IntegrationNamespaceSyncRuleArrayOutput {
	return i.ToIntegrationNamespaceSyncRuleArrayOutputWithContext(context.Background())
}

func (i IntegrationNamespaceSyncRuleArray) ToIntegrationNamespaceSyncRuleArrayOutputWithContext(ctx context.Context) IntegrationNamespaceSyncRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationNamespaceSyncRuleArrayOutput)
}

type IntegrationNamespaceSyncRuleOutput struct{ *pulumi.OutputState }

func (IntegrationNamespaceSyncRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationNamespaceSyncRule)(nil)).Elem()
}

func (o IntegrationNamespaceSyncRuleOutput) ToIntegrationNamespaceSyncRuleOutput() IntegrationNamespaceSyncRuleOutput {
	return o
}

func (o IntegrationNamespaceSyncRuleOutput) ToIntegrationNamespaceSyncRuleOutputWithContext(ctx context.Context) IntegrationNamespaceSyncRuleOutput {
	return o
}

// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
func (o IntegrationNamespaceSyncRuleOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationNamespaceSyncRule) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
func (o IntegrationNamespaceSyncRuleOutput) FilterAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationNamespaceSyncRule) *string { return v.FilterAction }).(pulumi.StringPtrOutput)
}

// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
func (o IntegrationNamespaceSyncRuleOutput) FilterSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IntegrationNamespaceSyncRule) *string { return v.FilterSource }).(pulumi.StringPtrOutput)
}

// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
func (o IntegrationNamespaceSyncRuleOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationNamespaceSyncRule) string { return v.Namespace }).(pulumi.StringOutput)
}

type IntegrationNamespaceSyncRuleArrayOutput struct{ *pulumi.OutputState }

func (IntegrationNamespaceSyncRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationNamespaceSyncRule)(nil)).Elem()
}

func (o IntegrationNamespaceSyncRuleArrayOutput) ToIntegrationNamespaceSyncRuleArrayOutput() IntegrationNamespaceSyncRuleArrayOutput {
	return o
}

func (o IntegrationNamespaceSyncRuleArrayOutput) ToIntegrationNamespaceSyncRuleArrayOutputWithContext(ctx context.Context) IntegrationNamespaceSyncRuleArrayOutput {
	return o
}

func (o IntegrationNamespaceSyncRuleArrayOutput) Index(i pulumi.IntInput) IntegrationNamespaceSyncRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationNamespaceSyncRule {
		return vs[0].([]IntegrationNamespaceSyncRule)[vs[1].(int)]
	}).(IntegrationNamespaceSyncRuleOutput)
}

type GetServicesService struct {
	Name string `pulumi:"name"`
}

// GetServicesServiceInput is an input type that accepts GetServicesServiceArgs and GetServicesServiceOutput values.
// You can construct a concrete instance of `GetServicesServiceInput` via:
//
//          GetServicesServiceArgs{...}
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
//          GetServicesServiceArray{ GetServicesServiceArgs{...} }
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
	pulumi.RegisterOutputType(IntegrationCustomNamespaceSyncRuleOutput{})
	pulumi.RegisterOutputType(IntegrationCustomNamespaceSyncRuleArrayOutput{})
	pulumi.RegisterOutputType(IntegrationNamespaceSyncRuleOutput{})
	pulumi.RegisterOutputType(IntegrationNamespaceSyncRuleArrayOutput{})
	pulumi.RegisterOutputType(GetServicesServiceOutput{})
	pulumi.RegisterOutputType(GetServicesServiceArrayOutput{})
}
