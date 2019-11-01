// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Signalfx.Aws
{
    /// <summary>
    /// SignalFx AWS CloudWatch integrations. For help with this integration see [Monitoring Amazon Web Services](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#monitor-amazon-web-services).
    /// 
    /// **Note:** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.
    /// 
    /// &gt; **WARNING** This resource implements a part of a workflow. You must use it with one of either `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
    /// 
    /// ## Service Names
    /// 
    /// Fields that expect an AWS service/namespace will work with one of: "AWS/ApiGateway" "AWS/AppStream" "AWS/AutoScaling" "AWS/Billing" "AWS/CloudFront" "AWS/CloudSearch" "AWS/Events" "AWS/Logs" "AWS/Connect" "AWS/DMS" "AWS/DX" "AWS/DynamoDB" "AWS/EC2" "AWS/EC2Spot" "AWS/ECS" "AWS/ElasticBeanstalk" "AWS/EBS" "AWS/EFS" "AWS/ELB" "AWS/ApplicationELB" "AWS/NetworkELB" "AWS/ElasticTranscoder" "AWS/ElastiCache" "AWS/ES" "AWS/ElasticMapReduce" "AWS/GameLift" "AWS/Inspector" "AWS/IoT" "AWS/KMS" "AWS/KinesisAnalytics" "AWS/Firehose" "AWS/Kinesis" "AWS/KinesisVideo" "AWS/Lambda" "AWS/Lex" "AWS/ML" "AWS/OpsWorks" "AWS/Polly" "AWS/Redshift" "AWS/RDS" "AWS/Route53" "AWS/SageMaker" "AWS/DDoSProtection" "AWS/SES" "AWS/SNS" "AWS/SQS" "AWS/S3" "AWS/SWF" "AWS/States" "AWS/StorageGateway" "AWS/Translate" "AWS/NATGateway" "AWS/VPN (VPN)" "WAF" "AWS/WorkSpaces".
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/aws_integration.html.markdown.
    /// </summary>
    public partial class Integration : Pulumi.CustomResource
    {
        /// <summary>
        /// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
        /// </summary>
        [Output("customCloudwatchNamespaces")]
        public Output<ImmutableArray<string>> CustomCloudwatchNamespaces { get; private set; } = null!;

        /// <summary>
        /// Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `custom_cloudwatch_namespaces` property.
        /// </summary>
        [Output("customNamespaceSyncRules")]
        public Output<ImmutableArray<Outputs.IntegrationCustomNamespaceSyncRules>> CustomNamespaceSyncRules { get; private set; } = null!;

        /// <summary>
        /// Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
        /// </summary>
        [Output("enableAwsUsage")]
        public Output<bool?> EnableAwsUsage { get; private set; } = null!;

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
        /// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
        /// </summary>
        [Output("importCloudWatch")]
        public Output<bool?> ImportCloudWatch { get; private set; } = null!;

        /// <summary>
        /// The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
        /// </summary>
        [Output("integrationId")]
        public Output<string> IntegrationId { get; private set; } = null!;

        /// <summary>
        /// If you specify `auth_method = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
        /// </summary>
        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

        /// <summary>
        /// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
        /// </summary>
        [Output("namespaceSyncRules")]
        public Output<ImmutableArray<Outputs.IntegrationNamespaceSyncRules>> NamespaceSyncRules { get; private set; } = null!;

        /// <summary>
        /// AWS poll rate (in seconds). One of `60` or `300`.
        /// </summary>
        [Output("pollRate")]
        public Output<int?> PollRate { get; private set; } = null!;

        /// <summary>
        /// List of AWS regions that SignalFx should monitor.
        /// </summary>
        [Output("regions")]
        public Output<ImmutableArray<string>> Regions { get; private set; } = null!;

        /// <summary>
        /// Role ARN that you add to an existing AWS integration object.
        /// </summary>
        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        /// <summary>
        /// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespace_sync_rule`.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<string>> Services { get; private set; } = null!;

        /// <summary>
        /// Used with `signalfx_aws_token_integration`. Use this property to specify the token.
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("signalfx:aws/integration:Integration", name, args, MakeResourceOptions(options, ""))
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

    public sealed class IntegrationArgs : Pulumi.ResourceArgs
    {
        [Input("customCloudwatchNamespaces")]
        private InputList<string>? _customCloudwatchNamespaces;

        /// <summary>
        /// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
        /// </summary>
        public InputList<string> CustomCloudwatchNamespaces
        {
            get => _customCloudwatchNamespaces ?? (_customCloudwatchNamespaces = new InputList<string>());
            set => _customCloudwatchNamespaces = value;
        }

        [Input("customNamespaceSyncRules")]
        private InputList<Inputs.IntegrationCustomNamespaceSyncRulesArgs>? _customNamespaceSyncRules;

        /// <summary>
        /// Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `custom_cloudwatch_namespaces` property.
        /// </summary>
        public InputList<Inputs.IntegrationCustomNamespaceSyncRulesArgs> CustomNamespaceSyncRules
        {
            get => _customNamespaceSyncRules ?? (_customNamespaceSyncRules = new InputList<Inputs.IntegrationCustomNamespaceSyncRulesArgs>());
            set => _customNamespaceSyncRules = value;
        }

        /// <summary>
        /// Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
        /// </summary>
        [Input("enableAwsUsage")]
        public Input<bool>? EnableAwsUsage { get; set; }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The `external_id` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
        /// </summary>
        [Input("importCloudWatch")]
        public Input<bool>? ImportCloudWatch { get; set; }

        /// <summary>
        /// The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
        /// </summary>
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        /// <summary>
        /// If you specify `auth_method = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("namespaceSyncRules")]
        private InputList<Inputs.IntegrationNamespaceSyncRulesArgs>? _namespaceSyncRules;

        /// <summary>
        /// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
        /// </summary>
        public InputList<Inputs.IntegrationNamespaceSyncRulesArgs> NamespaceSyncRules
        {
            get => _namespaceSyncRules ?? (_namespaceSyncRules = new InputList<Inputs.IntegrationNamespaceSyncRulesArgs>());
            set => _namespaceSyncRules = value;
        }

        /// <summary>
        /// AWS poll rate (in seconds). One of `60` or `300`.
        /// </summary>
        [Input("pollRate")]
        public Input<int>? PollRate { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// List of AWS regions that SignalFx should monitor.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// Role ARN that you add to an existing AWS integration object.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespace_sync_rule`.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        /// <summary>
        /// Used with `signalfx_aws_token_integration`. Use this property to specify the token.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public IntegrationArgs()
        {
        }
    }

    public sealed class IntegrationState : Pulumi.ResourceArgs
    {
        [Input("customCloudwatchNamespaces")]
        private InputList<string>? _customCloudwatchNamespaces;

        /// <summary>
        /// List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
        /// </summary>
        public InputList<string> CustomCloudwatchNamespaces
        {
            get => _customCloudwatchNamespaces ?? (_customCloudwatchNamespaces = new InputList<string>());
            set => _customCloudwatchNamespaces = value;
        }

        [Input("customNamespaceSyncRules")]
        private InputList<Inputs.IntegrationCustomNamespaceSyncRulesGetArgs>? _customNamespaceSyncRules;

        /// <summary>
        /// Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `custom_cloudwatch_namespaces` property.
        /// </summary>
        public InputList<Inputs.IntegrationCustomNamespaceSyncRulesGetArgs> CustomNamespaceSyncRules
        {
            get => _customNamespaceSyncRules ?? (_customNamespaceSyncRules = new InputList<Inputs.IntegrationCustomNamespaceSyncRulesGetArgs>());
            set => _customNamespaceSyncRules = value;
        }

        /// <summary>
        /// Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
        /// </summary>
        [Input("enableAwsUsage")]
        public Input<bool>? EnableAwsUsage { get; set; }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The `external_id` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
        /// </summary>
        [Input("importCloudWatch")]
        public Input<bool>? ImportCloudWatch { get; set; }

        /// <summary>
        /// The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
        /// </summary>
        [Input("integrationId")]
        public Input<string>? IntegrationId { get; set; }

        /// <summary>
        /// If you specify `auth_method = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("namespaceSyncRules")]
        private InputList<Inputs.IntegrationNamespaceSyncRulesGetArgs>? _namespaceSyncRules;

        /// <summary>
        /// Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
        /// </summary>
        public InputList<Inputs.IntegrationNamespaceSyncRulesGetArgs> NamespaceSyncRules
        {
            get => _namespaceSyncRules ?? (_namespaceSyncRules = new InputList<Inputs.IntegrationNamespaceSyncRulesGetArgs>());
            set => _namespaceSyncRules = value;
        }

        /// <summary>
        /// AWS poll rate (in seconds). One of `60` or `300`.
        /// </summary>
        [Input("pollRate")]
        public Input<int>? PollRate { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// List of AWS regions that SignalFx should monitor.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        /// <summary>
        /// Role ARN that you add to an existing AWS integration object.
        /// </summary>
        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespace_sync_rule`.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        /// <summary>
        /// Used with `signalfx_aws_token_integration`. Use this property to specify the token.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public IntegrationState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class IntegrationCustomNamespaceSyncRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<string> DefaultAction { get; set; } = null!;

        /// <summary>
        /// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        [Input("filterAction", required: true)]
        public Input<string> FilterAction { get; set; } = null!;

        /// <summary>
        /// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        /// </summary>
        [Input("filterSource", required: true)]
        public Input<string> FilterSource { get; set; } = null!;

        /// <summary>
        /// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public IntegrationCustomNamespaceSyncRulesArgs()
        {
        }
    }

    public sealed class IntegrationCustomNamespaceSyncRulesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<string> DefaultAction { get; set; } = null!;

        /// <summary>
        /// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        [Input("filterAction", required: true)]
        public Input<string> FilterAction { get; set; } = null!;

        /// <summary>
        /// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        /// </summary>
        [Input("filterSource", required: true)]
        public Input<string> FilterSource { get; set; } = null!;

        /// <summary>
        /// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public IntegrationCustomNamespaceSyncRulesGetArgs()
        {
        }
    }

    public sealed class IntegrationNamespaceSyncRulesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<string> DefaultAction { get; set; } = null!;

        /// <summary>
        /// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        [Input("filterAction", required: true)]
        public Input<string> FilterAction { get; set; } = null!;

        /// <summary>
        /// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        /// </summary>
        [Input("filterSource", required: true)]
        public Input<string> FilterSource { get; set; } = null!;

        /// <summary>
        /// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public IntegrationNamespaceSyncRulesArgs()
        {
        }
    }

    public sealed class IntegrationNamespaceSyncRulesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<string> DefaultAction { get; set; } = null!;

        /// <summary>
        /// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        [Input("filterAction", required: true)]
        public Input<string> FilterAction { get; set; } = null!;

        /// <summary>
        /// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        /// </summary>
        [Input("filterSource", required: true)]
        public Input<string> FilterSource { get; set; } = null!;

        /// <summary>
        /// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public IntegrationNamespaceSyncRulesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class IntegrationCustomNamespaceSyncRules
    {
        /// <summary>
        /// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        public readonly string DefaultAction;
        /// <summary>
        /// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        public readonly string FilterAction;
        /// <summary>
        /// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        /// </summary>
        public readonly string FilterSource;
        /// <summary>
        /// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private IntegrationCustomNamespaceSyncRules(
            string defaultAction,
            string filterAction,
            string filterSource,
            string @namespace)
        {
            DefaultAction = defaultAction;
            FilterAction = filterAction;
            FilterSource = filterSource;
            Namespace = @namespace;
        }
    }

    [OutputType]
    public sealed class IntegrationNamespaceSyncRules
    {
        /// <summary>
        /// Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        public readonly string DefaultAction;
        /// <summary>
        /// Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        public readonly string FilterAction;
        /// <summary>
        /// Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        /// </summary>
        public readonly string FilterSource;
        /// <summary>
        /// An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private IntegrationNamespaceSyncRules(
            string defaultAction,
            string filterAction,
            string filterSource,
            string @namespace)
        {
            DefaultAction = defaultAction;
            FilterAction = filterAction;
            FilterSource = filterSource;
            Namespace = @namespace;
        }
    }
    }
}
