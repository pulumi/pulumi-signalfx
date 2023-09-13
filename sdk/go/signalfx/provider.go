// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The provider type for the signalfx package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// API URL for your SignalFx org, may include a realm
	ApiUrl pulumi.StringPtrOutput `pulumi:"apiUrl"`
	// SignalFx auth token
	AuthToken pulumi.StringPtrOutput `pulumi:"authToken"`
	// Application URL for your SignalFx org, often customized for organizations using SSO
	CustomAppUrl pulumi.StringPtrOutput `pulumi:"customAppUrl"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:signalfx", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// API URL for your SignalFx org, may include a realm
	ApiUrl *string `pulumi:"apiUrl"`
	// SignalFx auth token
	AuthToken *string `pulumi:"authToken"`
	// Application URL for your SignalFx org, often customized for organizations using SSO
	CustomAppUrl *string `pulumi:"customAppUrl"`
	// Max retries for a single HTTP call. Defaults to 4
	RetryMaxAttempts *int `pulumi:"retryMaxAttempts"`
	// Maximum retry wait for a single HTTP call in seconds. Defaults to 30
	RetryWaitMaxSeconds *int `pulumi:"retryWaitMaxSeconds"`
	// Minimum retry wait for a single HTTP call in seconds. Defaults to 1
	RetryWaitMinSeconds *int `pulumi:"retryWaitMinSeconds"`
	// Timeout duration for a single HTTP call in seconds. Defaults to 120
	TimeoutSeconds *int `pulumi:"timeoutSeconds"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// API URL for your SignalFx org, may include a realm
	ApiUrl pulumi.StringPtrInput
	// SignalFx auth token
	AuthToken pulumi.StringPtrInput
	// Application URL for your SignalFx org, often customized for organizations using SSO
	CustomAppUrl pulumi.StringPtrInput
	// Max retries for a single HTTP call. Defaults to 4
	RetryMaxAttempts pulumi.IntPtrInput
	// Maximum retry wait for a single HTTP call in seconds. Defaults to 30
	RetryWaitMaxSeconds pulumi.IntPtrInput
	// Minimum retry wait for a single HTTP call in seconds. Defaults to 1
	RetryWaitMinSeconds pulumi.IntPtrInput
	// Timeout duration for a single HTTP call in seconds. Defaults to 120
	TimeoutSeconds pulumi.IntPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

func (i *Provider) ToOutput(ctx context.Context) pulumix.Output[*Provider] {
	return pulumix.Output[*Provider]{
		OutputState: i.ToProviderOutputWithContext(ctx).OutputState,
	}
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) ToOutput(ctx context.Context) pulumix.Output[*Provider] {
	return pulumix.Output[*Provider]{
		OutputState: o.OutputState,
	}
}

// API URL for your SignalFx org, may include a realm
func (o ProviderOutput) ApiUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApiUrl }).(pulumi.StringPtrOutput)
}

// SignalFx auth token
func (o ProviderOutput) AuthToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.AuthToken }).(pulumi.StringPtrOutput)
}

// Application URL for your SignalFx org, often customized for organizations using SSO
func (o ProviderOutput) CustomAppUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.CustomAppUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
