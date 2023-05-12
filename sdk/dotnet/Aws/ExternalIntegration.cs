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
    /// SignalFx AWS CloudWatch integrations using Role ARNs. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
    /// 
    /// &gt; **NOTE** When managing integrations use a session token for an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
    /// 
    /// &gt; **WARNING** This resource implements a part of a workflow. You must use it with `signalfx.aws.Integration`. Check with SignalFx support for your realm's AWS account id.
    /// 
    /// ## Example Usage
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
    ///     var awsMyteamExtern = new SignalFx.Aws.ExternalIntegration("awsMyteamExtern");
    /// 
    ///     var signalfxAssumePolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
    ///     {
    ///         Statements = new[]
    ///         {
    ///             new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
    ///             {
    ///                 Actions = new[]
    ///                 {
    ///                     "sts:AssumeRole",
    ///                 },
    ///                 Principals = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
    ///                     {
    ///                         Type = "AWS",
    ///                         Identifiers = new[]
    ///                         {
    ///                             awsMyteamExtern.SignalfxAwsAccount,
    ///                         },
    ///                     },
    ///                 },
    ///                 Conditions = new[]
    ///                 {
    ///                     new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
    ///                     {
    ///                         Test = "StringEquals",
    ///                         Variable = "sts:ExternalId",
    ///                         Values = new[]
    ///                         {
    ///                             awsMyteamExtern.ExternalId,
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var awsSplunkRole = new Aws.Iam.Role("awsSplunkRole", new()
    ///     {
    ///         Description = "signalfx integration to read out data and send it to signalfxs aws account",
    ///         AssumeRolePolicy = signalfxAssumePolicy.Apply(getPolicyDocumentResult =&gt; getPolicyDocumentResult.Json),
    ///     });
    /// 
    ///     var awsSplunkPolicy = new Aws.Iam.Policy("awsSplunkPolicy", new()
    ///     {
    ///         Description = "AWS permissions required by the Splunk Observability Cloud",
    ///         PolicyDocument = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Effect"": ""Allow"",
    ///       ""Action"": [
    ///         ""apigateway:GET"",
    ///         ""autoscaling:DescribeAutoScalingGroups"",
    ///         ""cloudfront:GetDistributionConfig"",
    ///         ""cloudfront:ListDistributions"",
    ///         ""cloudfront:ListTagsForResource"",
    ///         ""cloudwatch:DescribeAlarms"",
    ///         ""cloudwatch:GetMetricData"",
    ///         ""cloudwatch:GetMetricStatistics"",
    ///         ""cloudwatch:ListMetrics"",
    ///         ""directconnect:DescribeConnections"",
    ///         ""dynamodb:DescribeTable"",
    ///         ""dynamodb:ListTables"",
    ///         ""dynamodb:ListTagsOfResource"",
    ///         ""ec2:DescribeInstances"",
    ///         ""ec2:DescribeInstanceStatus"",
    ///         ""ec2:DescribeRegions"",
    ///         ""ec2:DescribeReservedInstances"",
    ///         ""ec2:DescribeReservedInstancesModifications"",
    ///         ""ec2:DescribeTags"",
    ///         ""ec2:DescribeVolumes"",
    ///         ""ecs:DescribeClusters"",
    ///         ""ecs:DescribeServices"",
    ///         ""ecs:DescribeTasks"",
    ///         ""ecs:ListClusters"",
    ///         ""ecs:ListServices"",
    ///         ""ecs:ListTagsForResource"",
    ///         ""ecs:ListTaskDefinitions"",
    ///         ""ecs:ListTasks"",
    ///         ""elasticache:DescribeCacheClusters"",
    ///         ""elasticloadbalancing:DescribeLoadBalancerAttributes"",
    ///         ""elasticloadbalancing:DescribeLoadBalancers"",
    ///         ""elasticloadbalancing:DescribeTags"",
    ///         ""elasticloadbalancing:DescribeTargetGroups"",
    ///         ""elasticmapreduce:DescribeCluster"",
    ///         ""elasticmapreduce:ListClusters"",
    ///         ""es:DescribeElasticsearchDomain"",
    ///         ""es:ListDomainNames"",
    ///         ""kinesis:DescribeStream"",
    ///         ""kinesis:ListShards"",
    ///         ""kinesis:ListStreams"",
    ///         ""kinesis:ListTagsForStream"",
    ///         ""lambda:GetAlias"",
    ///         ""lambda:ListFunctions"",
    ///         ""lambda:ListTags"",
    ///         ""logs:DeleteSubscriptionFilter"",
    ///         ""logs:DescribeLogGroups"",
    ///         ""logs:DescribeSubscriptionFilters"",
    ///         ""logs:PutSubscriptionFilter"",
    ///         ""organizations:DescribeOrganization"",
    ///         ""rds:DescribeDBClusters"",
    ///         ""rds:DescribeDBInstances"",
    ///         ""rds:ListTagsForResource"",
    ///         ""redshift:DescribeClusters"",
    ///         ""redshift:DescribeLoggingStatus"",
    ///         ""s3:GetBucketLocation"",
    ///         ""s3:GetBucketLogging"",
    ///         ""s3:GetBucketNotification"",
    ///         ""s3:GetBucketTagging"",
    ///         ""s3:ListAllMyBuckets"",
    ///         ""s3:ListBucket"",
    ///         ""s3:PutBucketNotification"",
    ///         ""sqs:GetQueueAttributes"",
    ///         ""sqs:ListQueues"",
    ///         ""sqs:ListQueueTags"",
    ///         ""states:ListStateMachines"",
    ///         ""tag:GetResources"",
    ///         ""workspaces:DescribeWorkspaces""
    ///       ],
    ///       ""Resource"": ""*""
    ///     }
    ///   ]
    /// }
    /// ",
    ///     });
    /// 
    ///     var splunkRolePolicyAttach = new Aws.Iam.RolePolicyAttachment("splunkRolePolicyAttach", new()
    ///     {
    ///         Role = awsSplunkRole.Name,
    ///         PolicyArn = awsSplunkPolicy.Arn,
    ///     });
    /// 
    ///     var awsMyteam = new SignalFx.Aws.Integration("awsMyteam", new()
    ///     {
    ///         Enabled = true,
    ///         IntegrationId = awsMyteamExtern.Id,
    ///         ExternalId = awsMyteamExtern.ExternalId,
    ///         RoleArn = awsSplunkRole.Arn,
    ///         Regions = new[]
    ///         {
    ///             "us-east-1",
    ///         },
    ///         PollRate = 300,
    ///         ImportCloudWatch = true,
    ///         EnableAwsUsage = true,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:aws/externalIntegration:ExternalIntegration")]
    public partial class ExternalIntegration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The external ID to use with your IAM role and with `signalfx.aws.Integration`.
        /// </summary>
        [Output("externalId")]
        public Output<string> ExternalId { get; private set; } = null!;

        /// <summary>
        /// The name of this integration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
        /// </summary>
        [Output("signalfxAwsAccount")]
        public Output<string> SignalfxAwsAccount { get; private set; } = null!;


        /// <summary>
        /// Create a ExternalIntegration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExternalIntegration(string name, ExternalIntegrationArgs? args = null, CustomResourceOptions? options = null)
            : base("signalfx:aws/externalIntegration:ExternalIntegration", name, args ?? new ExternalIntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExternalIntegration(string name, Input<string> id, ExternalIntegrationState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:aws/externalIntegration:ExternalIntegration", name, state, MakeResourceOptions(options, id))
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
                    "signalfxAwsAccount",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExternalIntegration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExternalIntegration Get(string name, Input<string> id, ExternalIntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new ExternalIntegration(name, id, state, options);
        }
    }

    public sealed class ExternalIntegrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of this integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ExternalIntegrationArgs()
        {
        }
        public static new ExternalIntegrationArgs Empty => new ExternalIntegrationArgs();
    }

    public sealed class ExternalIntegrationState : global::Pulumi.ResourceArgs
    {
        [Input("externalId")]
        private Input<string>? _externalId;

        /// <summary>
        /// The external ID to use with your IAM role and with `signalfx.aws.Integration`.
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
        /// The name of this integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("signalfxAwsAccount")]
        private Input<string>? _signalfxAwsAccount;

        /// <summary>
        /// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
        /// </summary>
        public Input<string>? SignalfxAwsAccount
        {
            get => _signalfxAwsAccount;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _signalfxAwsAccount = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ExternalIntegrationState()
        {
        }
        public static new ExternalIntegrationState Empty => new ExternalIntegrationState();
    }
}
