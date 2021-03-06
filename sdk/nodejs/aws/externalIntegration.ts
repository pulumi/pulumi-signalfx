// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * SignalFx AWS CloudWatch integrations using Role ARNs. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
 *
 * > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.
 *
 * > **WARNING** This resource implements a part of a workflow. You must use it with `signalfx.aws.Integration`. Check with SignalFx support for your realm's AWS account id.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const awsMyteamExtern = new signalfx.aws.ExternalIntegration("awsMyteamExtern", {});
 * const signalfxAssumePolicy = pulumi.all([awsMyteamExtern.signalfxAwsAccount, awsMyteamExtern.externalId]).apply(([signalfxAwsAccount, externalId]) => aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [{
 *             type: "AWS",
 *             identifiers: [signalfxAwsAccount],
 *         }],
 *         conditions: [{
 *             test: "StringEquals",
 *             variable: "sts:ExternalId",
 *             values: [externalId],
 *         }],
 *     }],
 * }));
 * const awsSfxRole = new aws.iam.Role("awsSfxRole", {
 *     description: "signalfx integration to read out data and send it to signalfxs aws account",
 *     assumeRolePolicy: signalfxAssumePolicy.json,
 * });
 * const awsReadPermissions = new aws.iam.Policy("awsReadPermissions", {
 *     description: "farts",
 *     policy: `{
 * 	"Version": "2012-10-17",
 * 	"Statement": [
 * 		{
 * 			"Action": [
 * 				"dynamodb:ListTables",
 * 		    "dynamodb:DescribeTable",
 * 		    "dynamodb:ListTagsOfResource",
 * 		    "ec2:DescribeInstances",
 * 		    "ec2:DescribeInstanceStatus",
 * 		    "ec2:DescribeVolumes",
 * 		    "ec2:DescribeReservedInstances",
 * 		    "ec2:DescribeReservedInstancesModifications",
 * 		    "ec2:DescribeTags",
 * 		    "organizations:DescribeOrganization",
 * 		    "cloudwatch:ListMetrics",
 * 		    "cloudwatch:GetMetricData",
 * 		    "cloudwatch:GetMetricStatistics",
 * 		    "cloudwatch:DescribeAlarms",
 * 		    "sqs:ListQueues",
 * 		    "sqs:GetQueueAttributes",
 * 		    "sqs:ListQueueTags",
 * 		    "elasticmapreduce:ListClusters",
 * 		    "elasticmapreduce:DescribeCluster",
 * 		    "kinesis:ListShards",
 * 		    "kinesis:ListStreams",
 * 		    "kinesis:DescribeStream",
 * 		    "kinesis:ListTagsForStream",
 * 		    "rds:DescribeDBInstances",
 * 		    "rds:ListTagsForResource",
 * 		    "elasticloadbalancing:DescribeLoadBalancers",
 * 		    "elasticloadbalancing:DescribeTags",
 * 		    "elasticache:describeCacheClusters",
 * 		    "redshift:DescribeClusters",
 * 		    "lambda:GetAlias",
 * 		    "lambda:ListFunctions",
 * 		    "lambda:ListTags",
 * 		    "autoscaling:DescribeAutoScalingGroups",
 * 		    "s3:ListAllMyBuckets",
 * 		    "s3:ListBucket",
 * 		    "s3:GetBucketLocation",
 * 		    "s3:GetBucketTagging",
 * 		    "ecs:ListServices",
 * 		    "ecs:ListTasks",
 * 		    "ecs:DescribeTasks",
 * 		    "ecs:DescribeServices",
 * 		    "ecs:ListClusters",
 * 		    "ecs:DescribeClusters",
 * 		    "ecs:ListTaskDefinitions",
 * 		    "ecs:ListTagsForResource",
 * 		    "apigateway:GET",
 * 		    "cloudfront:ListDistributions",
 * 		    "cloudfront:ListTagsForResource",
 * 		    "tag:GetResources",
 * 		    "es:ListDomainNames",
 * 		    "es:DescribeElasticsearchDomain"
 * 			],
 * 			"Effect": "Allow",
 * 			"Resource": "*"
 * 		}
 * 	]
 * }
 * `,
 * });
 * const sfx_read_attach = new aws.iam.RolePolicyAttachment("sfx-read-attach", {
 *     role: awsSfxRole.name,
 *     policyArn: awsReadPermissions.arn,
 * });
 * const awsMyteam = new signalfx.aws.Integration("awsMyteam", {
 *     enabled: true,
 *     integrationId: awsMyteamExtern.id,
 *     externalId: awsMyteamExtern.externalId,
 *     roleArn: awsSfxRole.arn,
 *     regions: ["us-east-1"],
 *     pollRate: 300,
 *     importCloudWatch: true,
 *     enableAwsUsage: true,
 * });
 * ```
 */
export class ExternalIntegration extends pulumi.CustomResource {
    /**
     * Get an existing ExternalIntegration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExternalIntegrationState, opts?: pulumi.CustomResourceOptions): ExternalIntegration {
        return new ExternalIntegration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:aws/externalIntegration:ExternalIntegration';

    /**
     * Returns true if the given object is an instance of ExternalIntegration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExternalIntegration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExternalIntegration.__pulumiType;
    }

    /**
     * The external ID to use with your IAM role and with `signalfx.aws.Integration`.
     */
    public /*out*/ readonly externalId!: pulumi.Output<string>;
    /**
     * The name of this integration
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The AWS Account ARN to use with your policies/roles, provided by SignalFx.
     */
    public /*out*/ readonly signalfxAwsAccount!: pulumi.Output<string>;

    /**
     * Create a ExternalIntegration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExternalIntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExternalIntegrationArgs | ExternalIntegrationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExternalIntegrationState | undefined;
            inputs["externalId"] = state ? state.externalId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["signalfxAwsAccount"] = state ? state.signalfxAwsAccount : undefined;
        } else {
            const args = argsOrState as ExternalIntegrationArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["externalId"] = undefined /*out*/;
            inputs["signalfxAwsAccount"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ExternalIntegration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExternalIntegration resources.
 */
export interface ExternalIntegrationState {
    /**
     * The external ID to use with your IAM role and with `signalfx.aws.Integration`.
     */
    readonly externalId?: pulumi.Input<string>;
    /**
     * The name of this integration
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AWS Account ARN to use with your policies/roles, provided by SignalFx.
     */
    readonly signalfxAwsAccount?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExternalIntegration resource.
 */
export interface ExternalIntegrationArgs {
    /**
     * The name of this integration
     */
    readonly name?: pulumi.Input<string>;
}
