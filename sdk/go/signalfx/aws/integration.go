// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SignalFx AWS CloudWatch integrations. For help with this integration see [Monitoring Amazon Web Services](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#monitor-amazon-web-services).
//
// > **NOTE** When managing integrations use a session token for an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
//
// > **WARNING** This resource implements a part of a workflow. You must use it with one of either `aws.ExternalIntegration` or `aws.TokenIntegration`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi-signalfx/sdk/v5/go/signalfx/aws"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		awsMyteamExternal, err := aws.NewExternalIntegration(ctx, "awsMyteamExternal", nil)
// 		if err != nil {
// 			return err
// 		}
// 		awsSfxRole, err := iam.NewRole(ctx, "awsSfxRole", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = aws.NewIntegration(ctx, "awsMyteam", &aws.IntegrationArgs{
// 			Enabled:       pulumi.Bool(true),
// 			IntegrationId: awsMyteamExternal.ID(),
// 			ExternalId:    awsMyteamExternal.ExternalId,
// 			RoleArn:       awsSfxRole.Arn,
// 			Regions: pulumi.StringArray{
// 				pulumi.String("us-east-1"),
// 			},
// 			PollRate:         pulumi.Int(300),
// 			ImportCloudWatch: pulumi.Bool(true),
// 			EnableAwsUsage:   pulumi.Bool(true),
// 			CustomNamespaceSyncRules: aws.IntegrationCustomNamespaceSyncRuleArray{
// 				&aws.IntegrationCustomNamespaceSyncRuleArgs{
// 					DefaultAction: pulumi.String("Exclude"),
// 					FilterAction:  pulumi.String("Include"),
// 					FilterSource:  pulumi.String("filter('code', '200')"),
// 					Namespace:     pulumi.String("my-custom-namespace"),
// 				},
// 			},
// 			NamespaceSyncRules: aws.IntegrationNamespaceSyncRuleArray{
// 				&aws.IntegrationNamespaceSyncRuleArgs{
// 					DefaultAction: pulumi.String("Exclude"),
// 					FilterAction:  pulumi.String("Include"),
// 					FilterSource:  pulumi.String("filter('code', '200')"),
// 					Namespace:     pulumi.String("AWS/EC2"),
// 				},
// 			},
// 			MetricStatsToSyncs: aws.IntegrationMetricStatsToSyncArray{
// 				&aws.IntegrationMetricStatsToSyncArgs{
// 					Namespace: pulumi.String("AWS/EC2"),
// 					Metric:    pulumi.String("NetworkPacketsIn"),
// 					Stats: pulumi.StringArray{
// 						pulumi.String("upper"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
	// Controls how SignalFx checks for large amounts of data for this AWS integration. If `true`, SignalFx monitors the amount of data coming in from the integration.
	EnableCheckLargeVolume pulumi.BoolPtrOutput `pulumi:"enableCheckLargeVolume"`
	// Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
	EnableLogsSync pulumi.BoolOutput `pulumi:"enableLogsSync"`
	// Whether the integration is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The `externalId` property from one of a `aws.ExternalIntegration` or `aws.TokenIntegration`
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
	ImportCloudWatch pulumi.BoolPtrOutput `pulumi:"importCloudWatch"`
	// The id of one of a `aws.ExternalIntegration` or `aws.TokenIntegration`.
	IntegrationId pulumi.StringOutput `pulumi:"integrationId"`
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that SignalFx collects for this metric. If you specify this property, SignalFx retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, SignalFx retrieves the AWS standard set of statistics.
	MetricStatsToSyncs IntegrationMetricStatsToSyncArrayOutput `pulumi:"metricStatsToSyncs"`
	// Name of the org token to be used for data ingestion. If not specified then default access token is used.
	NamedToken pulumi.StringPtrOutput `pulumi:"namedToken"`
	// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
	NamespaceSyncRules IntegrationNamespaceSyncRuleArrayOutput `pulumi:"namespaceSyncRules"`
	// AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
	PollRate pulumi.IntPtrOutput `pulumi:"pollRate"`
	// List of AWS regions that SignalFx should monitor.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services pulumi.StringArrayOutput `pulumi:"services"`
	// Indicates that SignalFx should sync metrics and metadata from custom AWS namespaces only (see the `customNamespaceSyncRule` above). Defaults to `false`.
	SyncCustomNamespacesOnly pulumi.BoolPtrOutput `pulumi:"syncCustomNamespacesOnly"`
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
	UseGetMetricDataMethod pulumi.BoolPtrOutput `pulumi:"useGetMetricDataMethod"`
	// Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics. Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.
	UseMetricStreamsSync pulumi.BoolOutput `pulumi:"useMetricStreamsSync"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.IntegrationId == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationId'")
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
	// Controls how SignalFx checks for large amounts of data for this AWS integration. If `true`, SignalFx monitors the amount of data coming in from the integration.
	EnableCheckLargeVolume *bool `pulumi:"enableCheckLargeVolume"`
	// Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
	EnableLogsSync *bool `pulumi:"enableLogsSync"`
	// Whether the integration is enabled.
	Enabled *bool `pulumi:"enabled"`
	// The `externalId` property from one of a `aws.ExternalIntegration` or `aws.TokenIntegration`
	ExternalId *string `pulumi:"externalId"`
	// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
	ImportCloudWatch *bool `pulumi:"importCloudWatch"`
	// The id of one of a `aws.ExternalIntegration` or `aws.TokenIntegration`.
	IntegrationId *string `pulumi:"integrationId"`
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
	Key *string `pulumi:"key"`
	// Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that SignalFx collects for this metric. If you specify this property, SignalFx retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, SignalFx retrieves the AWS standard set of statistics.
	MetricStatsToSyncs []IntegrationMetricStatsToSync `pulumi:"metricStatsToSyncs"`
	// Name of the org token to be used for data ingestion. If not specified then default access token is used.
	NamedToken *string `pulumi:"namedToken"`
	// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
	NamespaceSyncRules []IntegrationNamespaceSyncRule `pulumi:"namespaceSyncRules"`
	// AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
	PollRate *int `pulumi:"pollRate"`
	// List of AWS regions that SignalFx should monitor.
	Regions []string `pulumi:"regions"`
	// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
	RoleArn *string `pulumi:"roleArn"`
	// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services []string `pulumi:"services"`
	// Indicates that SignalFx should sync metrics and metadata from custom AWS namespaces only (see the `customNamespaceSyncRule` above). Defaults to `false`.
	SyncCustomNamespacesOnly *bool `pulumi:"syncCustomNamespacesOnly"`
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
	Token *string `pulumi:"token"`
	// Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
	UseGetMetricDataMethod *bool `pulumi:"useGetMetricDataMethod"`
	// Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics. Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.
	UseMetricStreamsSync *bool `pulumi:"useMetricStreamsSync"`
}

type IntegrationState struct {
	// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
	CustomCloudwatchNamespaces pulumi.StringArrayInput
	// Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
	CustomNamespaceSyncRules IntegrationCustomNamespaceSyncRuleArrayInput
	// Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
	EnableAwsUsage pulumi.BoolPtrInput
	// Controls how SignalFx checks for large amounts of data for this AWS integration. If `true`, SignalFx monitors the amount of data coming in from the integration.
	EnableCheckLargeVolume pulumi.BoolPtrInput
	// Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
	EnableLogsSync pulumi.BoolPtrInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolPtrInput
	// The `externalId` property from one of a `aws.ExternalIntegration` or `aws.TokenIntegration`
	ExternalId pulumi.StringPtrInput
	// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
	ImportCloudWatch pulumi.BoolPtrInput
	// The id of one of a `aws.ExternalIntegration` or `aws.TokenIntegration`.
	IntegrationId pulumi.StringPtrInput
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
	Key pulumi.StringPtrInput
	// Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that SignalFx collects for this metric. If you specify this property, SignalFx retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, SignalFx retrieves the AWS standard set of statistics.
	MetricStatsToSyncs IntegrationMetricStatsToSyncArrayInput
	// Name of the org token to be used for data ingestion. If not specified then default access token is used.
	NamedToken pulumi.StringPtrInput
	// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
	NamespaceSyncRules IntegrationNamespaceSyncRuleArrayInput
	// AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
	PollRate pulumi.IntPtrInput
	// List of AWS regions that SignalFx should monitor.
	Regions pulumi.StringArrayInput
	// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
	RoleArn pulumi.StringPtrInput
	// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services pulumi.StringArrayInput
	// Indicates that SignalFx should sync metrics and metadata from custom AWS namespaces only (see the `customNamespaceSyncRule` above). Defaults to `false`.
	SyncCustomNamespacesOnly pulumi.BoolPtrInput
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
	Token pulumi.StringPtrInput
	// Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
	UseGetMetricDataMethod pulumi.BoolPtrInput
	// Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics. Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.
	UseMetricStreamsSync pulumi.BoolPtrInput
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
	// Controls how SignalFx checks for large amounts of data for this AWS integration. If `true`, SignalFx monitors the amount of data coming in from the integration.
	EnableCheckLargeVolume *bool `pulumi:"enableCheckLargeVolume"`
	// Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
	EnableLogsSync *bool `pulumi:"enableLogsSync"`
	// Whether the integration is enabled.
	Enabled bool `pulumi:"enabled"`
	// The `externalId` property from one of a `aws.ExternalIntegration` or `aws.TokenIntegration`
	ExternalId *string `pulumi:"externalId"`
	// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
	ImportCloudWatch *bool `pulumi:"importCloudWatch"`
	// The id of one of a `aws.ExternalIntegration` or `aws.TokenIntegration`.
	IntegrationId string `pulumi:"integrationId"`
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
	Key *string `pulumi:"key"`
	// Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that SignalFx collects for this metric. If you specify this property, SignalFx retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, SignalFx retrieves the AWS standard set of statistics.
	MetricStatsToSyncs []IntegrationMetricStatsToSync `pulumi:"metricStatsToSyncs"`
	// Name of the org token to be used for data ingestion. If not specified then default access token is used.
	NamedToken *string `pulumi:"namedToken"`
	// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
	NamespaceSyncRules []IntegrationNamespaceSyncRule `pulumi:"namespaceSyncRules"`
	// AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
	PollRate *int `pulumi:"pollRate"`
	// List of AWS regions that SignalFx should monitor.
	Regions []string `pulumi:"regions"`
	// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
	RoleArn *string `pulumi:"roleArn"`
	// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services []string `pulumi:"services"`
	// Indicates that SignalFx should sync metrics and metadata from custom AWS namespaces only (see the `customNamespaceSyncRule` above). Defaults to `false`.
	SyncCustomNamespacesOnly *bool `pulumi:"syncCustomNamespacesOnly"`
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
	Token *string `pulumi:"token"`
	// Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
	UseGetMetricDataMethod *bool `pulumi:"useGetMetricDataMethod"`
	// Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics. Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.
	UseMetricStreamsSync *bool `pulumi:"useMetricStreamsSync"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
	CustomCloudwatchNamespaces pulumi.StringArrayInput
	// Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
	CustomNamespaceSyncRules IntegrationCustomNamespaceSyncRuleArrayInput
	// Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
	EnableAwsUsage pulumi.BoolPtrInput
	// Controls how SignalFx checks for large amounts of data for this AWS integration. If `true`, SignalFx monitors the amount of data coming in from the integration.
	EnableCheckLargeVolume pulumi.BoolPtrInput
	// Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
	EnableLogsSync pulumi.BoolPtrInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolInput
	// The `externalId` property from one of a `aws.ExternalIntegration` or `aws.TokenIntegration`
	ExternalId pulumi.StringPtrInput
	// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
	ImportCloudWatch pulumi.BoolPtrInput
	// The id of one of a `aws.ExternalIntegration` or `aws.TokenIntegration`.
	IntegrationId pulumi.StringInput
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
	Key pulumi.StringPtrInput
	// Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that SignalFx collects for this metric. If you specify this property, SignalFx retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, SignalFx retrieves the AWS standard set of statistics.
	MetricStatsToSyncs IntegrationMetricStatsToSyncArrayInput
	// Name of the org token to be used for data ingestion. If not specified then default access token is used.
	NamedToken pulumi.StringPtrInput
	// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
	NamespaceSyncRules IntegrationNamespaceSyncRuleArrayInput
	// AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
	PollRate pulumi.IntPtrInput
	// List of AWS regions that SignalFx should monitor.
	Regions pulumi.StringArrayInput
	// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
	RoleArn pulumi.StringPtrInput
	// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services pulumi.StringArrayInput
	// Indicates that SignalFx should sync metrics and metadata from custom AWS namespaces only (see the `customNamespaceSyncRule` above). Defaults to `false`.
	SyncCustomNamespacesOnly pulumi.BoolPtrInput
	// If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
	Token pulumi.StringPtrInput
	// Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
	UseGetMetricDataMethod pulumi.BoolPtrInput
	// Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics. Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.
	UseMetricStreamsSync pulumi.BoolPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}

type IntegrationInput interface {
	pulumi.Input

	ToIntegrationOutput() IntegrationOutput
	ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput
}

func (*Integration) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil)).Elem()
}

func (i *Integration) ToIntegrationOutput() IntegrationOutput {
	return i.ToIntegrationOutputWithContext(context.Background())
}

func (i *Integration) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationOutput)
}

// IntegrationArrayInput is an input type that accepts IntegrationArray and IntegrationArrayOutput values.
// You can construct a concrete instance of `IntegrationArrayInput` via:
//
//          IntegrationArray{ IntegrationArgs{...} }
type IntegrationArrayInput interface {
	pulumi.Input

	ToIntegrationArrayOutput() IntegrationArrayOutput
	ToIntegrationArrayOutputWithContext(context.Context) IntegrationArrayOutput
}

type IntegrationArray []IntegrationInput

func (IntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Integration)(nil)).Elem()
}

func (i IntegrationArray) ToIntegrationArrayOutput() IntegrationArrayOutput {
	return i.ToIntegrationArrayOutputWithContext(context.Background())
}

func (i IntegrationArray) ToIntegrationArrayOutputWithContext(ctx context.Context) IntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationArrayOutput)
}

// IntegrationMapInput is an input type that accepts IntegrationMap and IntegrationMapOutput values.
// You can construct a concrete instance of `IntegrationMapInput` via:
//
//          IntegrationMap{ "key": IntegrationArgs{...} }
type IntegrationMapInput interface {
	pulumi.Input

	ToIntegrationMapOutput() IntegrationMapOutput
	ToIntegrationMapOutputWithContext(context.Context) IntegrationMapOutput
}

type IntegrationMap map[string]IntegrationInput

func (IntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Integration)(nil)).Elem()
}

func (i IntegrationMap) ToIntegrationMapOutput() IntegrationMapOutput {
	return i.ToIntegrationMapOutputWithContext(context.Background())
}

func (i IntegrationMap) ToIntegrationMapOutputWithContext(ctx context.Context) IntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMapOutput)
}

type IntegrationOutput struct{ *pulumi.OutputState }

func (IntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil)).Elem()
}

func (o IntegrationOutput) ToIntegrationOutput() IntegrationOutput {
	return o
}

func (o IntegrationOutput) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return o
}

type IntegrationArrayOutput struct{ *pulumi.OutputState }

func (IntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Integration)(nil)).Elem()
}

func (o IntegrationArrayOutput) ToIntegrationArrayOutput() IntegrationArrayOutput {
	return o
}

func (o IntegrationArrayOutput) ToIntegrationArrayOutputWithContext(ctx context.Context) IntegrationArrayOutput {
	return o
}

func (o IntegrationArrayOutput) Index(i pulumi.IntInput) IntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Integration {
		return vs[0].([]*Integration)[vs[1].(int)]
	}).(IntegrationOutput)
}

type IntegrationMapOutput struct{ *pulumi.OutputState }

func (IntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Integration)(nil)).Elem()
}

func (o IntegrationMapOutput) ToIntegrationMapOutput() IntegrationMapOutput {
	return o
}

func (o IntegrationMapOutput) ToIntegrationMapOutputWithContext(ctx context.Context) IntegrationMapOutput {
	return o
}

func (o IntegrationMapOutput) MapIndex(k pulumi.StringInput) IntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Integration {
		return vs[0].(map[string]*Integration)[vs[1].(string)]
	}).(IntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationInput)(nil)).Elem(), &Integration{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationArrayInput)(nil)).Elem(), IntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMapInput)(nil)).Elem(), IntegrationMap{})
	pulumi.RegisterOutputType(IntegrationOutput{})
	pulumi.RegisterOutputType(IntegrationArrayOutput{})
	pulumi.RegisterOutputType(IntegrationMapOutput{})
}
