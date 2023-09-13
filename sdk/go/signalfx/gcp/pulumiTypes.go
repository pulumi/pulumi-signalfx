// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

var _ = internal.GetEnvOrDefault

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

func (i IntegrationProjectServiceKeyArgs) ToOutput(ctx context.Context) pulumix.Output[IntegrationProjectServiceKey] {
	return pulumix.Output[IntegrationProjectServiceKey]{
		OutputState: i.ToIntegrationProjectServiceKeyOutputWithContext(ctx).OutputState,
	}
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

func (i IntegrationProjectServiceKeyArray) ToOutput(ctx context.Context) pulumix.Output[[]IntegrationProjectServiceKey] {
	return pulumix.Output[[]IntegrationProjectServiceKey]{
		OutputState: i.ToIntegrationProjectServiceKeyArrayOutputWithContext(ctx).OutputState,
	}
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

func (o IntegrationProjectServiceKeyOutput) ToOutput(ctx context.Context) pulumix.Output[IntegrationProjectServiceKey] {
	return pulumix.Output[IntegrationProjectServiceKey]{
		OutputState: o.OutputState,
	}
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

func (o IntegrationProjectServiceKeyArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]IntegrationProjectServiceKey] {
	return pulumix.Output[[]IntegrationProjectServiceKey]{
		OutputState: o.OutputState,
	}
}

func (o IntegrationProjectServiceKeyArrayOutput) Index(i pulumi.IntInput) IntegrationProjectServiceKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IntegrationProjectServiceKey {
		return vs[0].([]IntegrationProjectServiceKey)[vs[1].(int)]
	}).(IntegrationProjectServiceKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationProjectServiceKeyInput)(nil)).Elem(), IntegrationProjectServiceKeyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationProjectServiceKeyArrayInput)(nil)).Elem(), IntegrationProjectServiceKeyArray{})
	pulumi.RegisterOutputType(IntegrationProjectServiceKeyOutput{})
	pulumi.RegisterOutputType(IntegrationProjectServiceKeyArrayOutput{})
}
