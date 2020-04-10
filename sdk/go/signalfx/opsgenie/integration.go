// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsgenie

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// SignalFx Opsgenie integration.
//
// > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.
type Integration struct {
	pulumi.CustomResourceState

	// The API key
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
	ApiUrl pulumi.StringPtrOutput `pulumi:"apiUrl"`
	// Whether the integration is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Name of the integration.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil || args.ApiKey == nil {
		return nil, errors.New("missing required argument 'ApiKey'")
	}
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil {
		args = &IntegrationArgs{}
	}
	var resource Integration
	err := ctx.RegisterResource("signalfx:opsgenie/integration:Integration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegration gets an existing Integration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationState, opts ...pulumi.ResourceOption) (*Integration, error) {
	var resource Integration
	err := ctx.ReadResource("signalfx:opsgenie/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// The API key
	ApiKey *string `pulumi:"apiKey"`
	// Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
	ApiUrl *string `pulumi:"apiUrl"`
	// Whether the integration is enabled.
	Enabled *bool `pulumi:"enabled"`
	// Name of the integration.
	Name *string `pulumi:"name"`
}

type IntegrationState struct {
	// The API key
	ApiKey pulumi.StringPtrInput
	// Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
	ApiUrl pulumi.StringPtrInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolPtrInput
	// Name of the integration.
	Name pulumi.StringPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// The API key
	ApiKey string `pulumi:"apiKey"`
	// Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
	ApiUrl *string `pulumi:"apiUrl"`
	// Whether the integration is enabled.
	Enabled bool `pulumi:"enabled"`
	// Name of the integration.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// The API key
	ApiKey pulumi.StringInput
	// Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
	ApiUrl pulumi.StringPtrInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolInput
	// Name of the integration.
	Name pulumi.StringPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}
