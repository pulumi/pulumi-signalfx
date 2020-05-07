// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the signalfx package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.AuthToken == nil {
		args.AuthToken = pulumi.StringPtr(getEnvOrDefault("", nil, "SFX_AUTH_TOKEN").(string))
	}
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
	// Application URL for your SignalFx org, often customzied for organizations using SSO
	CustomAppUrl *string `pulumi:"customAppUrl"`
	// Timeout duration for a single HTTP call in seconds. Defaults to 120
	TimeoutSeconds *int `pulumi:"timeoutSeconds"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// API URL for your SignalFx org, may include a realm
	ApiUrl pulumi.StringPtrInput
	// SignalFx auth token
	AuthToken pulumi.StringPtrInput
	// Application URL for your SignalFx org, often customzied for organizations using SSO
	CustomAppUrl pulumi.StringPtrInput
	// Timeout duration for a single HTTP call in seconds. Defaults to 120
	TimeoutSeconds pulumi.IntPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}
