// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Aws
{
    /// <summary>
    /// AWS CloudWatch integrations for Splunk Observability Cloud. For help with this integration see [Monitoring Amazon Web Services](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/aws/get-awstoc.html).
    /// 
    /// This resource implements a part of a workflow. Use it with one of either `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
    /// 
    /// &gt; **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
    /// 
    /// ## Example
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// using SignalFx = Pulumi.SignalFx;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // This resource returns an account id in `external_id`…
    ///     var awsMyteamExternal = new SignalFx.Aws.ExternalIntegration("aws_myteam_external", new()
    ///     {
    ///         Name = "My AWS integration",
    ///     });
    /// 
    ///     // Make yourself an AWS IAM role here, use `signalfx_aws_external_integration.aws_myteam_external.external_id`
    ///     var awsSfxRole = new Aws.Index.IamRole("aws_sfx_role");
    /// 
    ///     var awsMyteam = new SignalFx.Aws.Integration("aws_myteam", new()
    ///     {
    ///         Enabled = true,
    ///         IntegrationId = awsMyteamExternal.Id,
    ///         ExternalId = awsMyteamExternal.ExternalId,
    ///         RoleArn = awsSfxRole.Arn,
    ///         Regions = new[]
    ///         {
    ///             "us-east-1",
    ///         },
    ///         PollRate = 300,
    ///         ImportCloudWatch = true,
    ///         EnableAwsUsage = true,
    ///         CustomNamespaceSyncRules = new[]
    ///         {
    ///             new SignalFx.Aws.Inputs.IntegrationCustomNamespaceSyncRuleArgs
    ///             {
    ///                 DefaultAction = "Exclude",
    ///                 FilterAction = "Include",
    ///                 FilterSource = "filter('code', '200')",
    ///                 Namespace = "my-custom-namespace",
    ///             },
    ///         },
    ///         NamespaceSyncRules = new[]
    ///         {
    ///             new SignalFx.Aws.Inputs.IntegrationNamespaceSyncRuleArgs
    ///             {
    ///                 DefaultAction = "Exclude",
    ///                 FilterAction = "Include",
    ///                 FilterSource = "filter('code', '200')",
    ///                 Namespace = "AWS/EC2",
    ///             },
    ///         },
    ///         MetricStatsToSyncs = new[]
    ///         {
    ///             new SignalFx.Aws.Inputs.IntegrationMetricStatsToSyncArgs
    ///             {
    ///                 Namespace = "AWS/EC2",
    ///                 Metric = "NetworkPacketsIn",
    ///                 Stats = new[]
    ///                 {
    ///                     "upper",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:aws/integration:Integration")]
    public partial class Integration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The mechanism used to authenticate with AWS. Use one of `signalfx.aws.ExternalIntegration` or
        /// `signalfx.aws.TokenIntegration` to define this
        /// </summary>
        [Output("authMethod")]
        public Output<string> AuthMethod { get; private set; } = null!;

        /// <summary>
        /// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; Splunk Observability Cloud imports the metrics so you can monitor them.
        /// </summary>
        [Output("customCloudwatchNamespaces")]
        public Output<ImmutableArray<string>> CustomCloudwatchNamespaces { get; private set; } = null!;

        /// <summary>
        /// Each element controls the data collected by Splunk Observability Cloud for the specified namespace. Conflicts with the `custom_cloudwatch_namespaces` property.
        /// </summary>
        [Output("customNamespaceSyncRules")]
        public Output<ImmutableArray<Outputs.IntegrationCustomNamespaceSyncRule>> CustomNamespaceSyncRules { get; private set; } = null!;

        /// <summary>
        /// Flag that controls how Splunk Observability Cloud imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, Splunk Observability Cloud imports the metrics.
        /// </summary>
        [Output("enableAwsUsage")]
        public Output<bool?> EnableAwsUsage { get; private set; } = null!;

        /// <summary>
        /// Controls how Splunk Observability Cloud checks for large amounts of data for this AWS integration. If `true`, Splunk Observability Cloud monitors the amount of data coming in from the integration.
        /// </summary>
        [Output("enableCheckLargeVolume")]
        public Output<bool?> EnableCheckLargeVolume { get; private set; } = null!;

        /// <summary>
        /// Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
        /// </summary>
        [Output("enableLogsSync")]
        public Output<bool> EnableLogsSync { get; private set; } = null!;

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The `external_id` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
        /// </summary>
        [Output("externalId")]
        public Output<string?> ExternalId { get; private set; } = null!;

        /// <summary>
        /// Flag that controls how Splunk Observability Cloud imports Cloud Watch metrics. If true, Splunk Observability Cloud imports Cloud Watch metrics from AWS.
        /// </summary>
        [Output("importCloudWatch")]
        public Output<bool?> ImportCloudWatch { get; private set; } = null!;

        /// <summary>
        /// The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
        /// </summary>
        [Output("integrationId")]
        public Output<string> IntegrationId { get; private set; } = null!;

        /// <summary>
        /// If you specify `auth_method = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
        /// </summary>
        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        /// <summary>
        /// Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that Splunk Observability Cloud collects for this metric. If you specify this property, Splunk Observability Cloud retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, Splunk Observability Cloud retrieves the AWS standard set of statistics.
        /// </summary>
        [Output("metricStatsToSyncs")]
        public Output<ImmutableArray<Outputs.IntegrationMetricStatsToSync>> MetricStatsToSyncs { get; private set; } = null!;

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name of the org token to be used for data ingestion. If not specified then default access token is used.
        /// </summary>
        [Output("namedToken")]
        public Output<string?> NamedToken { get; private set; } = null!;

        /// <summary>
        /// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that Splunk Observability Cloud collects for the namespace. Conflicts with the `services` property. If you don't specify either property, Splunk Observability Cloud syncs all data in all AWS namespaces.
        /// </summary>
        [Output("namespaceSyncRules")]
        public Output<ImmutableArray<Outputs.IntegrationNamespaceSyncRule>> NamespaceSyncRules { get; private set; } = null!;

        /// <summary>
        /// AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
        /// </summary>
        [Output("pollRate")]
        public Output<int?> PollRate { get; private set; } = null!;

        /// <summary>
        /// List of AWS regions that Splunk Observability Cloud should monitor. It cannot be empty.
        /// </summary>
        [Output("regions")]
        public Output<ImmutableArray<string>> Regions { get; private set; } = null!;

        /// <summary>
        /// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// List of AWS services that you want Splunk Observability Cloud to monitor. Each element is a string designating an AWS service. Can be an empty list to import data for all supported services. Conflicts with `namespace_sync_rule`. See [Amazon Web Services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#amazon-web-services) for a list of valid values.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<string>> Services { get; private set; } = null!;

        /// <summary>
        /// Indicates that Splunk Observability Cloud should sync metrics and metadata from custom AWS namespaces only (see the `custom_namespace_sync_rule` above). Defaults to `false`.
        /// </summary>
        [Output("syncCustomNamespacesOnly")]
        public Output<bool?> SyncCustomNamespacesOnly { get; private set; } = null!;

        /// <summary>
        /// If you specify `auth_method = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;

        /// <summary>
        /// Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics.&lt;br&gt;
        /// Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.&lt;br&gt;
        /// Note you need to deploy additional resources on your AWS account to enable CloudWatch metrics streaming. Select one of the [CloudFormation templates](https://docs.splunk.com/Observability/gdi/get-data-in/connect/aws/aws-cloudformation.html) to deploy all the required resources.
        /// </summary>
        [Output("useMetricStreamsSync")]
        public Output<bool> UseMetricStreamsSync { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("signalfx:aws/integration:Integration", name, args ?? new IntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Integration(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:aws/integration:Integration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "externalId",
                    "key",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Integration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Integration Get(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new Integration(name, id, state, options);
        }
    }

    public sealed class IntegrationArgs : global::Pulumi.ResourceArgs
    {
        [Input("customCloudwatchNamespaces")]
        private InputList<string>? _customCloudwatchNamespaces;

        /// <summary>
        /// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; Splunk Observability Cloud imports the metrics so you can monitor them.
        /// </summary>
        public InputList<string> CustomCloudwatchNamespaces
        {
            get => _customCloudwatchNamespaces ?? (_customCloudwatchNamespaces = new InputList<string>());
            set => _customCloudwatchNamespaces = value;
        }

        [Input("customNamespaceSyncRules")]
        private InputList<Inputs.IntegrationCustomNamespaceSyncRuleArgs>? _customNamespaceSyncRules;

        /// <summary>
        /// Each element controls the data collected by Splunk Observability Cloud for the specified namespace. Conflicts with the `custom_cloudwatch_namespaces` property.
        /// </summary>
        public InputList<Inputs.IntegrationCustomNamespaceSyncRuleArgs> CustomNamespaceSyncRules
        {
            get => _customNamespaceSyncRules ?? (_customNamespaceSyncRules = new InputList<Inputs.IntegrationCustomNamespaceSyncRuleArgs>());
            set => _customNamespaceSyncRules = value;
        }

        /// <summary>
        /// Flag that controls how Splunk Observability Cloud imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, Splunk Observability Cloud imports the metrics.
        /// </summary>
        [Input("enableAwsUsage")]
        public Input<bool>? EnableAwsUsage { get; set; }

        /// <summary>
        /// Controls how Splunk Observability Cloud checks for large amounts of data for this AWS integration. If `true`, Splunk Observability Cloud monitors the amount of data coming in from the integration.
        /// </summary>
        [Input("enableCheckLargeVolume")]
        public Input<bool>? EnableCheckLargeVolume { get; set; }

        /// <summary>
        /// Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
        /// </summary>
        [Input("enableLogsSync")]
        public Input<bool>? EnableLogsSync { get; set; }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("externalId")]
        private Input<string>? _externalId;

        /// <summary>
        /// The `external_id` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
        /// </summary>
        public Input<string>? ExternalId
        {
            get => _externalId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _externalId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Flag that controls how Splunk Observability Cloud imports Cloud Watch metrics. If true, Splunk Observability Cloud imports Cloud Watch metrics from AWS.
        /// </summary>
        [Input("importCloudWatch")]
        public Input<bool>? ImportCloudWatch { get; set; }

        /// <summary>
        /// The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
        /// </summary>
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// If you specify `auth_method = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("metricStatsToSyncs")]
        private InputList<Inputs.IntegrationMetricStatsToSyncArgs>? _metricStatsToSyncs;

        /// <summary>
        /// Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that Splunk Observability Cloud collects for this metric. If you specify this property, Splunk Observability Cloud retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, Splunk Observability Cloud retrieves the AWS standard set of statistics.
        /// </summary>
        public InputList<Inputs.IntegrationMetricStatsToSyncArgs> MetricStatsToSyncs
        {
            get => _metricStatsToSyncs ?? (_metricStatsToSyncs = new InputList<Inputs.IntegrationMetricStatsToSyncArgs>());
            set => _metricStatsToSyncs = value;
        }

        /// <summary>
        /// Name of the org token to be used for data ingestion. If not specified then default access token is used.
        /// </summary>
        [Input("namedToken")]
        public Input<string>? NamedToken { get; set; }

        [Input("namespaceSyncRules")]
        private InputList<Inputs.IntegrationNamespaceSyncRuleArgs>? _namespaceSyncRules;

        /// <summary>
        /// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that Splunk Observability Cloud collects for the namespace. Conflicts with the `services` property. If you don't specify either property, Splunk Observability Cloud syncs all data in all AWS namespaces.
        /// </summary>
        public InputList<Inputs.IntegrationNamespaceSyncRuleArgs> NamespaceSyncRules
        {
            get => _namespaceSyncRules ?? (_namespaceSyncRules = new InputList<Inputs.IntegrationNamespaceSyncRuleArgs>());
            set => _namespaceSyncRules = value;
        }

        /// <summary>
        /// AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
        /// </summary>
        [Input("pollRate")]
        public Input<int>? PollRate { get; set; }

        [Input("regions", required: true)]
        private InputList<string>? _regions;

        /// <summary>
        /// List of AWS regions that Splunk Observability Cloud should monitor. It cannot be empty.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// List of AWS services that you want Splunk Observability Cloud to monitor. Each element is a string designating an AWS service. Can be an empty list to import data for all supported services. Conflicts with `namespace_sync_rule`. See [Amazon Web Services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#amazon-web-services) for a list of valid values.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        /// <summary>
        /// Indicates that Splunk Observability Cloud should sync metrics and metadata from custom AWS namespaces only (see the `custom_namespace_sync_rule` above). Defaults to `false`.
        /// </summary>
        [Input("syncCustomNamespacesOnly")]
        public Input<bool>? SyncCustomNamespacesOnly { get; set; }

        /// <summary>
        /// If you specify `auth_method = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics.&lt;br&gt;
        /// Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.&lt;br&gt;
        /// Note you need to deploy additional resources on your AWS account to enable CloudWatch metrics streaming. Select one of the [CloudFormation templates](https://docs.splunk.com/Observability/gdi/get-data-in/connect/aws/aws-cloudformation.html) to deploy all the required resources.
        /// </summary>
        [Input("useMetricStreamsSync")]
        public Input<bool>? UseMetricStreamsSync { get; set; }

        public IntegrationArgs()
        {
        }
        public static new IntegrationArgs Empty => new IntegrationArgs();
    }

    public sealed class IntegrationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The mechanism used to authenticate with AWS. Use one of `signalfx.aws.ExternalIntegration` or
        /// `signalfx.aws.TokenIntegration` to define this
        /// </summary>
        [Input("authMethod")]
        public Input<string>? AuthMethod { get; set; }

        [Input("customCloudwatchNamespaces")]
        private InputList<string>? _customCloudwatchNamespaces;

        /// <summary>
        /// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; Splunk Observability Cloud imports the metrics so you can monitor them.
        /// </summary>
        public InputList<string> CustomCloudwatchNamespaces
        {
            get => _customCloudwatchNamespaces ?? (_customCloudwatchNamespaces = new InputList<string>());
            set => _customCloudwatchNamespaces = value;
        }

        [Input("customNamespaceSyncRules")]
        private InputList<Inputs.IntegrationCustomNamespaceSyncRuleGetArgs>? _customNamespaceSyncRules;

        /// <summary>
        /// Each element controls the data collected by Splunk Observability Cloud for the specified namespace. Conflicts with the `custom_cloudwatch_namespaces` property.
        /// </summary>
        public InputList<Inputs.IntegrationCustomNamespaceSyncRuleGetArgs> CustomNamespaceSyncRules
        {
            get => _customNamespaceSyncRules ?? (_customNamespaceSyncRules = new InputList<Inputs.IntegrationCustomNamespaceSyncRuleGetArgs>());
            set => _customNamespaceSyncRules = value;
        }

        /// <summary>
        /// Flag that controls how Splunk Observability Cloud imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, Splunk Observability Cloud imports the metrics.
        /// </summary>
        [Input("enableAwsUsage")]
        public Input<bool>? EnableAwsUsage { get; set; }

        /// <summary>
        /// Controls how Splunk Observability Cloud checks for large amounts of data for this AWS integration. If `true`, Splunk Observability Cloud monitors the amount of data coming in from the integration.
        /// </summary>
        [Input("enableCheckLargeVolume")]
        public Input<bool>? EnableCheckLargeVolume { get; set; }

        /// <summary>
        /// Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
        /// </summary>
        [Input("enableLogsSync")]
        public Input<bool>? EnableLogsSync { get; set; }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("externalId")]
        private Input<string>? _externalId;

        /// <summary>
        /// The `external_id` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
        /// </summary>
        public Input<string>? ExternalId
        {
            get => _externalId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _externalId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Flag that controls how Splunk Observability Cloud imports Cloud Watch metrics. If true, Splunk Observability Cloud imports Cloud Watch metrics from AWS.
        /// </summary>
        [Input("importCloudWatch")]
        public Input<bool>? ImportCloudWatch { get; set; }

        /// <summary>
        /// The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
        /// </summary>
        [Input("integrationId")]
        public Input<string>? IntegrationId { get; set; }

        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// If you specify `auth_method = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("metricStatsToSyncs")]
        private InputList<Inputs.IntegrationMetricStatsToSyncGetArgs>? _metricStatsToSyncs;

        /// <summary>
        /// Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that Splunk Observability Cloud collects for this metric. If you specify this property, Splunk Observability Cloud retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, Splunk Observability Cloud retrieves the AWS standard set of statistics.
        /// </summary>
        public InputList<Inputs.IntegrationMetricStatsToSyncGetArgs> MetricStatsToSyncs
        {
            get => _metricStatsToSyncs ?? (_metricStatsToSyncs = new InputList<Inputs.IntegrationMetricStatsToSyncGetArgs>());
            set => _metricStatsToSyncs = value;
        }

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the org token to be used for data ingestion. If not specified then default access token is used.
        /// </summary>
        [Input("namedToken")]
        public Input<string>? NamedToken { get; set; }

        [Input("namespaceSyncRules")]
        private InputList<Inputs.IntegrationNamespaceSyncRuleGetArgs>? _namespaceSyncRules;

        /// <summary>
        /// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that Splunk Observability Cloud collects for the namespace. Conflicts with the `services` property. If you don't specify either property, Splunk Observability Cloud syncs all data in all AWS namespaces.
        /// </summary>
        public InputList<Inputs.IntegrationNamespaceSyncRuleGetArgs> NamespaceSyncRules
        {
            get => _namespaceSyncRules ?? (_namespaceSyncRules = new InputList<Inputs.IntegrationNamespaceSyncRuleGetArgs>());
            set => _namespaceSyncRules = value;
        }

        /// <summary>
        /// AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
        /// </summary>
        [Input("pollRate")]
        public Input<int>? PollRate { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// List of AWS regions that Splunk Observability Cloud should monitor. It cannot be empty.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// List of AWS services that you want Splunk Observability Cloud to monitor. Each element is a string designating an AWS service. Can be an empty list to import data for all supported services. Conflicts with `namespace_sync_rule`. See [Amazon Web Services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#amazon-web-services) for a list of valid values.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        /// <summary>
        /// Indicates that Splunk Observability Cloud should sync metrics and metadata from custom AWS namespaces only (see the `custom_namespace_sync_rule` above). Defaults to `false`.
        /// </summary>
        [Input("syncCustomNamespacesOnly")]
        public Input<bool>? SyncCustomNamespacesOnly { get; set; }

        /// <summary>
        /// If you specify `auth_method = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics.&lt;br&gt;
        /// Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.&lt;br&gt;
        /// Note you need to deploy additional resources on your AWS account to enable CloudWatch metrics streaming. Select one of the [CloudFormation templates](https://docs.splunk.com/Observability/gdi/get-data-in/connect/aws/aws-cloudformation.html) to deploy all the required resources.
        /// </summary>
        [Input("useMetricStreamsSync")]
        public Input<bool>? UseMetricStreamsSync { get; set; }

        public IntegrationState()
        {
        }
        public static new IntegrationState Empty => new IntegrationState();
    }
}
