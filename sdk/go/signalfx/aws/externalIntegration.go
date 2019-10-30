// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// SignalFx AWS CloudWatch integrations using Role ARNs. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
// 
// **Note:** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.
// 
// > **WARNING** This resource implements a part of a workflow. You must use it with one of either `aws.Integration`. Check with SignalFx support for your realm's AWS account id.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/aws_external_integration.html.markdown.
type ExternalIntegration struct {
	s *pulumi.ResourceState
}

// NewExternalIntegration registers a new resource with the given unique name, arguments, and options.
func NewExternalIntegration(ctx *pulumi.Context,
	name string, args *ExternalIntegrationArgs, opts ...pulumi.ResourceOpt) (*ExternalIntegration, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
	} else {
		inputs["name"] = args.Name
	}
	inputs["externalId"] = nil
	s, err := ctx.RegisterResource("signalfx:aws/externalIntegration:ExternalIntegration", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ExternalIntegration{s: s}, nil
}

// GetExternalIntegration gets an existing ExternalIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalIntegration(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ExternalIntegrationState, opts ...pulumi.ResourceOpt) (*ExternalIntegration, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["externalId"] = state.ExternalId
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("signalfx:aws/externalIntegration:ExternalIntegration", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ExternalIntegration{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ExternalIntegration) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ExternalIntegration) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The external ID to use with your IAM role and with `aws.Integration`.
func (r *ExternalIntegration) ExternalId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["externalId"])
}

// The name of this integration
func (r *ExternalIntegration) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering ExternalIntegration resources.
type ExternalIntegrationState struct {
	// The external ID to use with your IAM role and with `aws.Integration`.
	ExternalId interface{}
	// The name of this integration
	Name interface{}
}

// The set of arguments for constructing a ExternalIntegration resource.
type ExternalIntegrationArgs struct {
	// The name of this integration
	Name interface{}
}
