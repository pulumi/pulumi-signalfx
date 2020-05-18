// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// SignalFx AWS CloudWatch integrations. For help with this integration see [Monitoring Amazon Web Services](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#monitor-amazon-web-services).
//
// > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.
//
// > **WARNING** This resource implements a part of a workflow. You must use it with one of either `aws.ExternalIntegration` or `aws.TokenIntegration`.
//
//
// ## Service Names
//
// > **NOTE** You can use the data source "aws.getServices" to specify all services.
type Integration struct {
	pulumi.CustomResourceState

	// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
	CustomCloudwatchNamespaces pulumi.StringArrayOutput `pulumi:"customCloudwatchNamespaces"`
	// Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
	CustomNamespaceSyncRules IntegrationCustomNamespaceSyncRuleArrayOutput `pulumi:"customNamespaceSyncRules"`
	// Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
	EnableAwsUsage pulumi.BoolPtrOutput `pulumi:"enableAwsUsage"`
	// Whether the integration is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The `externalId` property from one of a `aws.ExternalIntegration` or `aws.TokenIntegration`
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
	ImportCloudWatch pulumi.BoolPtrOutput `pulumi:"importCloudWatch"`
	// The id of one of a `aws.ExternalIntegration` or `aws.TokenIntegration`.
	IntegrationId pulumi.StringOutput `pulumi:"integrationId"`
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
	NamespaceSyncRules IntegrationNamespaceSyncRuleArrayOutput `pulumi:"namespaceSyncRules"`
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate pulumi.IntPtrOutput `pulumi:"pollRate"`
	// List of AWS regions that SignalFx should monitor.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services pulumi.StringArrayOutput `pulumi:"services"`
	// Used with `signalfx_aws_token_integration`. Use this property to specify the token.
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
	UseGetMetricDataMethod pulumi.BoolPtrOutput `pulumi:"useGetMetricDataMethod"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.IntegrationId == nil {
		return nil, errors.New("missing required argument 'IntegrationId'")
	}
	if args == nil {
		args = &IntegrationArgs{}
	}
	var resource Integration
	err := ctx.RegisterResource("signalfx:aws/integration:Integration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("signalfx:aws/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
	CustomCloudwatchNamespaces []string `pulumi:"customCloudwatchNamespaces"`
	// Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
	CustomNamespaceSyncRules []IntegrationCustomNamespaceSyncRule `pulumi:"customNamespaceSyncRules"`
	// Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
	EnableAwsUsage *bool `pulumi:"enableAwsUsage"`
	// Whether the integration is enabled.
	Enabled *bool `pulumi:"enabled"`
	// The `externalId` property from one of a `aws.ExternalIntegration` or `aws.TokenIntegration`
	ExternalId *string `pulumi:"externalId"`
	// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
	ImportCloudWatch *bool `pulumi:"importCloudWatch"`
	// The id of one of a `aws.ExternalIntegration` or `aws.TokenIntegration`.
	IntegrationId *string `pulumi:"integrationId"`
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
	Key *string `pulumi:"key"`
	// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
	NamespaceSyncRules []IntegrationNamespaceSyncRule `pulumi:"namespaceSyncRules"`
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate *int `pulumi:"pollRate"`
	// List of AWS regions that SignalFx should monitor.
	Regions []string `pulumi:"regions"`
	// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
	RoleArn *string `pulumi:"roleArn"`
	// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services []string `pulumi:"services"`
	// Used with `signalfx_aws_token_integration`. Use this property to specify the token.
	Token *string `pulumi:"token"`
	// Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
	UseGetMetricDataMethod *bool `pulumi:"useGetMetricDataMethod"`
}

type IntegrationState struct {
	// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
	CustomCloudwatchNamespaces pulumi.StringArrayInput
	// Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
	CustomNamespaceSyncRules IntegrationCustomNamespaceSyncRuleArrayInput
	// Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
	EnableAwsUsage pulumi.BoolPtrInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolPtrInput
	// The `externalId` property from one of a `aws.ExternalIntegration` or `aws.TokenIntegration`
	ExternalId pulumi.StringPtrInput
	// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
	ImportCloudWatch pulumi.BoolPtrInput
	// The id of one of a `aws.ExternalIntegration` or `aws.TokenIntegration`.
	IntegrationId pulumi.StringPtrInput
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
	Key pulumi.StringPtrInput
	// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
	NamespaceSyncRules IntegrationNamespaceSyncRuleArrayInput
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate pulumi.IntPtrInput
	// List of AWS regions that SignalFx should monitor.
	Regions pulumi.StringArrayInput
	// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
	RoleArn pulumi.StringPtrInput
	// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services pulumi.StringArrayInput
	// Used with `signalfx_aws_token_integration`. Use this property to specify the token.
	Token pulumi.StringPtrInput
	// Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
	UseGetMetricDataMethod pulumi.BoolPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
	CustomCloudwatchNamespaces []string `pulumi:"customCloudwatchNamespaces"`
	// Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
	CustomNamespaceSyncRules []IntegrationCustomNamespaceSyncRule `pulumi:"customNamespaceSyncRules"`
	// Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
	EnableAwsUsage *bool `pulumi:"enableAwsUsage"`
	// Whether the integration is enabled.
	Enabled bool `pulumi:"enabled"`
	// The `externalId` property from one of a `aws.ExternalIntegration` or `aws.TokenIntegration`
	ExternalId *string `pulumi:"externalId"`
	// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
	ImportCloudWatch *bool `pulumi:"importCloudWatch"`
	// The id of one of a `aws.ExternalIntegration` or `aws.TokenIntegration`.
	IntegrationId string `pulumi:"integrationId"`
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
	Key *string `pulumi:"key"`
	// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
	NamespaceSyncRules []IntegrationNamespaceSyncRule `pulumi:"namespaceSyncRules"`
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate *int `pulumi:"pollRate"`
	// List of AWS regions that SignalFx should monitor.
	Regions []string `pulumi:"regions"`
	// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
	RoleArn *string `pulumi:"roleArn"`
	// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services []string `pulumi:"services"`
	// Used with `signalfx_aws_token_integration`. Use this property to specify the token.
	Token *string `pulumi:"token"`
	// Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
	UseGetMetricDataMethod *bool `pulumi:"useGetMetricDataMethod"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
	CustomCloudwatchNamespaces pulumi.StringArrayInput
	// Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
	CustomNamespaceSyncRules IntegrationCustomNamespaceSyncRuleArrayInput
	// Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
	EnableAwsUsage pulumi.BoolPtrInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolInput
	// The `externalId` property from one of a `aws.ExternalIntegration` or `aws.TokenIntegration`
	ExternalId pulumi.StringPtrInput
	// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
	ImportCloudWatch pulumi.BoolPtrInput
	// The id of one of a `aws.ExternalIntegration` or `aws.TokenIntegration`.
	IntegrationId pulumi.StringInput
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
	Key pulumi.StringPtrInput
	// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
	NamespaceSyncRules IntegrationNamespaceSyncRuleArrayInput
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate pulumi.IntPtrInput
	// List of AWS regions that SignalFx should monitor.
	Regions pulumi.StringArrayInput
	// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
	RoleArn pulumi.StringPtrInput
	// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services pulumi.StringArrayInput
	// Used with `signalfx_aws_token_integration`. Use this property to specify the token.
	Token pulumi.StringPtrInput
	// Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
	UseGetMetricDataMethod pulumi.BoolPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}
