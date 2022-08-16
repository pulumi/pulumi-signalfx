// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationProjectServiceKey struct {
	ProjectId  string `pulumi:"projectId"`
	ProjectKey string `pulumi:"projectKey"`
}

// IntegrationProjectServiceKeyInput is an input type that accepts IntegrationProjectServiceKeyArgs and IntegrationProjectServiceKeyOutput values.
// You can construct a concrete instance of `IntegrationProjectServiceKeyInput` via:
//
//	IntegrationProjectServiceKeyArgs{...}
type IntegrationProjectServiceKeyInput interface {
	pulumi.Input

	ToIntegrationProjectServiceKeyOutput() IntegrationProjectServiceKeyOutput
	ToIntegrationProjectServiceKeyOutputWithContext(context.Context) IntegrationProjectServiceKeyOutput
}

type IntegrationProjectServiceKeyArgs struct {
	ProjectId  pulumi.StringInput `pulumi:"projectId"`
	ProjectKey pulumi.StringInput `pulumi:"projectKey"`
}

func (IntegrationProjectServiceKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationProjectServiceKey)(nil)).Elem()
}

func (i IntegrationProjectServiceKeyArgs) ToIntegrationProjectServiceKeyOutput() IntegrationProjectServiceKeyOutput {
	return i.ToIntegrationProjectServiceKeyOutputWithContext(context.Background())
}

func (i IntegrationProjectServiceKeyArgs) ToIntegrationProjectServiceKeyOutputWithContext(ctx context.Context) IntegrationProjectServiceKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationProjectServiceKeyOutput)
}

// IntegrationProjectServiceKeyArrayInput is an input type that accepts IntegrationProjectServiceKeyArray and IntegrationProjectServiceKeyArrayOutput values.
// You can construct a concrete instance of `IntegrationProjectServiceKeyArrayInput` via:
//
//	IntegrationProjectServiceKeyArray{ IntegrationProjectServiceKeyArgs{...} }
type IntegrationProjectServiceKeyArrayInput interface {
	pulumi.Input

	ToIntegrationProjectServiceKeyArrayOutput() IntegrationProjectServiceKeyArrayOutput
	ToIntegrationProjectServiceKeyArrayOutputWithContext(context.Context) IntegrationProjectServiceKeyArrayOutput
}

type IntegrationProjectServiceKeyArray []IntegrationProjectServiceKeyInput

func (IntegrationProjectServiceKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationProjectServiceKey)(nil)).Elem()
}

func (i IntegrationProjectServiceKeyArray) ToIntegrationProjectServiceKeyArrayOutput() IntegrationProjectServiceKeyArrayOutput {
	return i.ToIntegrationProjectServiceKeyArrayOutputWithContext(context.Background())
}

func (i IntegrationProjectServiceKeyArray) ToIntegrationProjectServiceKeyArrayOutputWithContext(ctx context.Context) IntegrationProjectServiceKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationProjectServiceKeyArrayOutput)
}

type IntegrationProjectServiceKeyOutput struct{ *pulumi.OutputState }

func (IntegrationProjectServiceKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationProjectServiceKey)(nil)).Elem()
}

func (o IntegrationProjectServiceKeyOutput) ToIntegrationProjectServiceKeyOutput() IntegrationProjectServiceKeyOutput {
	return o
}

func (o IntegrationProjectServiceKeyOutput) ToIntegrationProjectServiceKeyOutputWithContext(ctx context.Context) IntegrationProjectServiceKeyOutput {
	return o
}

func (o IntegrationProjectServiceKeyOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationProjectServiceKey) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o IntegrationProjectServiceKeyOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v IntegrationProjectServiceKey) string { return v.ProjectKey }).(pulumi.StringOutput)
}

type IntegrationProjectServiceKeyArrayOutput struct{ *pulumi.OutputState }

func (IntegrationProjectServiceKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IntegrationProjectServiceKey)(nil)).Elem()
}

func (o IntegrationProjectServiceKeyArrayOutput) ToIntegrationProjectServiceKeyArrayOutput() IntegrationProjectServiceKeyArrayOutput {
	return o
}

func (o IntegrationProjectServiceKeyArrayOutput) ToIntegrationProjectServiceKeyArrayOutputWithContext(ctx context.Context) IntegrationProjectServiceKeyArrayOutput {
	return o
}

func (o IntegrationProjectServiceKeyArrayOutput) Index(i pulumi.IntInput) IntegrationProjectServiceKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationProjectServiceKey {
		return vs[0].([]IntegrationProjectServiceKey)[vs[1].(int)]
	}).(IntegrationProjectServiceKeyOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationProjectServiceKeyInput)(nil)).Elem(), IntegrationProjectServiceKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationProjectServiceKeyArrayInput)(nil)).Elem(), IntegrationProjectServiceKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServicesServiceInput)(nil)).Elem(), GetServicesServiceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetServicesServiceArrayInput)(nil)).Elem(), GetServicesServiceArray{})
	pulumi.RegisterOutputType(IntegrationProjectServiceKeyOutput{})
	pulumi.RegisterOutputType(IntegrationProjectServiceKeyArrayOutput{})
	pulumi.RegisterOutputType(GetServicesServiceOutput{})
	pulumi.RegisterOutputType(GetServicesServiceArrayOutput{})
}
