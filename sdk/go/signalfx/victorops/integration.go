// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package victorops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// SignalFx VictorOps integration.
//
// > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.
type Integration struct {
	pulumi.CustomResourceState

	// Whether the integration is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Name of the integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Victor Ops REST API URL.
	PostUrl pulumi.StringPtrOutput `pulumi:"postUrl"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil {
		args = &IntegrationArgs{}
	}
	var resource Integration
	err := ctx.RegisterResource("signalfx:victorops/integration:Integration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("signalfx:victorops/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// Whether the integration is enabled.
	Enabled *bool `pulumi:"enabled"`
	// Name of the integration.
	Name *string `pulumi:"name"`
	// Victor Ops REST API URL.
	PostUrl *string `pulumi:"postUrl"`
}

type IntegrationState struct {
	// Whether the integration is enabled.
	Enabled pulumi.BoolPtrInput
	// Name of the integration.
	Name pulumi.StringPtrInput
	// Victor Ops REST API URL.
	PostUrl pulumi.StringPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// Whether the integration is enabled.
	Enabled bool `pulumi:"enabled"`
	// Name of the integration.
	Name *string `pulumi:"name"`
	// Victor Ops REST API URL.
	PostUrl *string `pulumi:"postUrl"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// Whether the integration is enabled.
	Enabled pulumi.BoolInput
	// Name of the integration.
	Name pulumi.StringPtrInput
	// Victor Ops REST API URL.
	PostUrl pulumi.StringPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}

type IntegrationInput interface {
	pulumi.Input

	ToIntegrationOutput() IntegrationOutput
	ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput
}

func (Integration) ElementType() reflect.Type {
	return reflect.TypeOf((*Integration)(nil)).Elem()
}

func (i Integration) ToIntegrationOutput() IntegrationOutput {
	return i.ToIntegrationOutputWithContext(context.Background())
}

func (i Integration) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationOutput)
}

type IntegrationOutput struct {
	*pulumi.OutputState
}

func (IntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationOutput)(nil)).Elem()
}

func (o IntegrationOutput) ToIntegrationOutput() IntegrationOutput {
	return o
}

func (o IntegrationOutput) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationOutput{})
}
