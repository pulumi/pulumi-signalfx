// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Splunk Observability AWS CloudWatch integrations. For help with this integration see [Monitoring Amazon Web Services](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#monitor-amazon-web-services).
 *
 * > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
 *
 * > **WARNING** This resource implements a part of a workflow. You must use it with one of either `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * // This resource returns an account id in `external_id`…
 * const awsMyteamExternal = new signalfx.aws.ExternalIntegration("awsMyteamExternal", {});
 * // Make yourself an AWS IAM role here, use `signalfx_aws_external_integration.aws_myteam_external.external_id`
 * const awsSfxRole = new aws.iam.Role("awsSfxRole", {});
 * // Stuff here that uses the external and account ID
 * const awsMyteam = new signalfx.aws.Integration("awsMyteam", {
 *     enabled: true,
 *     integrationId: awsMyteamExternal.id,
 *     externalId: awsMyteamExternal.externalId,
 *     roleArn: awsSfxRole.arn,
 *     regions: ["us-east-1"],
 *     pollRate: 300,
 *     importCloudWatch: true,
 *     enableAwsUsage: true,
 *     customNamespaceSyncRules: [{
 *         defaultAction: "Exclude",
 *         filterAction: "Include",
 *         filterSource: "filter('code', '200')",
 *         namespace: "my-custom-namespace",
 *     }],
 *     namespaceSyncRules: [{
 *         defaultAction: "Exclude",
 *         filterAction: "Include",
 *         filterSource: "filter('code', '200')",
 *         namespace: "AWS/EC2",
 *     }],
 *     metricStatsToSyncs: [{
 *         namespace: "AWS/EC2",
 *         metric: "NetworkPacketsIn",
 *         stats: ["upper"],
 *     }],
 * });
 * ```
 */
export class Integration extends pulumi.CustomResource {
    /**
     * Get an existing Integration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationState, opts?: pulumi.CustomResourceOptions): Integration {
        return new Integration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:aws/integration:Integration';

    /**
     * Returns true if the given object is an instance of Integration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Integration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Integration.__pulumiType;
    }

    /**
     * The mechanism used to authenticate with AWS. Use one of `signalfx_aws_external_integration` or
     * `signalfx_aws_token_integration` to define this
     */
    public /*out*/ readonly authMethod!: pulumi.Output<string>;
    /**
     * List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; Splunk Observability imports the metrics so you can monitor them.
     */
    public readonly customCloudwatchNamespaces!: pulumi.Output<string[] | undefined>;
    /**
     * Each element controls the data collected by Splunk Observability for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
     */
    public readonly customNamespaceSyncRules!: pulumi.Output<outputs.aws.IntegrationCustomNamespaceSyncRule[] | undefined>;
    /**
     * Flag that controls how Splunk Observability imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, Splunk Observability imports the metrics.
     */
    public readonly enableAwsUsage!: pulumi.Output<boolean | undefined>;
    /**
     * Controls how Splunk Observability checks for large amounts of data for this AWS integration. If `true`, Splunk Observability monitors the amount of data coming in from the integration.
     */
    public readonly enableCheckLargeVolume!: pulumi.Output<boolean | undefined>;
    /**
     * Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
     */
    public readonly enableLogsSync!: pulumi.Output<boolean>;
    /**
     * Whether the integration is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The `externalId` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
     */
    public readonly externalId!: pulumi.Output<string | undefined>;
    /**
     * Flag that controls how Splunk Observability imports Cloud Watch metrics. If true, Splunk Observability imports Cloud Watch metrics from AWS.
     */
    public readonly importCloudWatch!: pulumi.Output<boolean | undefined>;
    /**
     * The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
     */
    public readonly integrationId!: pulumi.Output<string>;
    /**
     * If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that Splunk Observability collects for this metric. If you specify this property, Splunk Observability retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, Splunk Observability retrieves the AWS standard set of statistics.
     */
    public readonly metricStatsToSyncs!: pulumi.Output<outputs.aws.IntegrationMetricStatsToSync[] | undefined>;
    /**
     * Name of the integration.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     */
    public readonly namedToken!: pulumi.Output<string | undefined>;
    /**
     * Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that Splunk Observability collects for the namespace. Conflicts with the `services` property. If you don't specify either property, Splunk Observability syncs all data in all AWS namespaces.
     */
    public readonly namespaceSyncRules!: pulumi.Output<outputs.aws.IntegrationNamespaceSyncRule[] | undefined>;
    /**
     * AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     */
    public readonly pollRate!: pulumi.Output<number | undefined>;
    /**
     * List of AWS regions that Splunk Observability should monitor.
     */
    public readonly regions!: pulumi.Output<string[] | undefined>;
    /**
     * Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * List of AWS services that you want Splunk Observability to monitor. Each element is a string designating an AWS service. Can be an empty list to import data for all supported services. Conflicts with `namespaceSyncRule`. See [Amazon Web Services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#amazon-web-services) for a list of valid values.
     */
    public readonly services!: pulumi.Output<string[] | undefined>;
    /**
     * Indicates that Splunk Observability should sync metrics and metadata from custom AWS namespaces only (see the `customNamespaceSyncRule` above). Defaults to `false`.
     */
    public readonly syncCustomNamespacesOnly!: pulumi.Output<boolean | undefined>;
    /**
     * If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics.<br>
     * Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.<br>
     * Note you need to deploy additional resources on your AWS account to enable CloudWatch metrics streaming. Select one of the [CloudFormation templates](https://docs.splunk.com/Observability/gdi/get-data-in/connect/aws/aws-cloudformation.html) to deploy all the required resources.
     */
    public readonly useMetricStreamsSync!: pulumi.Output<boolean>;

    /**
     * Create a Integration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationArgs | IntegrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationState | undefined;
            resourceInputs["authMethod"] = state ? state.authMethod : undefined;
            resourceInputs["customCloudwatchNamespaces"] = state ? state.customCloudwatchNamespaces : undefined;
            resourceInputs["customNamespaceSyncRules"] = state ? state.customNamespaceSyncRules : undefined;
            resourceInputs["enableAwsUsage"] = state ? state.enableAwsUsage : undefined;
            resourceInputs["enableCheckLargeVolume"] = state ? state.enableCheckLargeVolume : undefined;
            resourceInputs["enableLogsSync"] = state ? state.enableLogsSync : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["externalId"] = state ? state.externalId : undefined;
            resourceInputs["importCloudWatch"] = state ? state.importCloudWatch : undefined;
            resourceInputs["integrationId"] = state ? state.integrationId : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["metricStatsToSyncs"] = state ? state.metricStatsToSyncs : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namedToken"] = state ? state.namedToken : undefined;
            resourceInputs["namespaceSyncRules"] = state ? state.namespaceSyncRules : undefined;
            resourceInputs["pollRate"] = state ? state.pollRate : undefined;
            resourceInputs["regions"] = state ? state.regions : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["syncCustomNamespacesOnly"] = state ? state.syncCustomNamespacesOnly : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["useMetricStreamsSync"] = state ? state.useMetricStreamsSync : undefined;
        } else {
            const args = argsOrState as IntegrationArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.integrationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationId'");
            }
            resourceInputs["customCloudwatchNamespaces"] = args ? args.customCloudwatchNamespaces : undefined;
            resourceInputs["customNamespaceSyncRules"] = args ? args.customNamespaceSyncRules : undefined;
            resourceInputs["enableAwsUsage"] = args ? args.enableAwsUsage : undefined;
            resourceInputs["enableCheckLargeVolume"] = args ? args.enableCheckLargeVolume : undefined;
            resourceInputs["enableLogsSync"] = args ? args.enableLogsSync : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["externalId"] = args?.externalId ? pulumi.secret(args.externalId) : undefined;
            resourceInputs["importCloudWatch"] = args ? args.importCloudWatch : undefined;
            resourceInputs["integrationId"] = args ? args.integrationId : undefined;
            resourceInputs["key"] = args?.key ? pulumi.secret(args.key) : undefined;
            resourceInputs["metricStatsToSyncs"] = args ? args.metricStatsToSyncs : undefined;
            resourceInputs["namedToken"] = args ? args.namedToken : undefined;
            resourceInputs["namespaceSyncRules"] = args ? args.namespaceSyncRules : undefined;
            resourceInputs["pollRate"] = args ? args.pollRate : undefined;
            resourceInputs["regions"] = args ? args.regions : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["syncCustomNamespacesOnly"] = args ? args.syncCustomNamespacesOnly : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["useMetricStreamsSync"] = args ? args.useMetricStreamsSync : undefined;
            resourceInputs["authMethod"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["externalId", "key"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Integration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Integration resources.
 */
export interface IntegrationState {
    /**
     * The mechanism used to authenticate with AWS. Use one of `signalfx_aws_external_integration` or
     * `signalfx_aws_token_integration` to define this
     */
    authMethod?: pulumi.Input<string>;
    /**
     * List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; Splunk Observability imports the metrics so you can monitor them.
     */
    customCloudwatchNamespaces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Each element controls the data collected by Splunk Observability for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
     */
    customNamespaceSyncRules?: pulumi.Input<pulumi.Input<inputs.aws.IntegrationCustomNamespaceSyncRule>[]>;
    /**
     * Flag that controls how Splunk Observability imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, Splunk Observability imports the metrics.
     */
    enableAwsUsage?: pulumi.Input<boolean>;
    /**
     * Controls how Splunk Observability checks for large amounts of data for this AWS integration. If `true`, Splunk Observability monitors the amount of data coming in from the integration.
     */
    enableCheckLargeVolume?: pulumi.Input<boolean>;
    /**
     * Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
     */
    enableLogsSync?: pulumi.Input<boolean>;
    /**
     * Whether the integration is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The `externalId` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
     */
    externalId?: pulumi.Input<string>;
    /**
     * Flag that controls how Splunk Observability imports Cloud Watch metrics. If true, Splunk Observability imports Cloud Watch metrics from AWS.
     */
    importCloudWatch?: pulumi.Input<boolean>;
    /**
     * The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
     */
    integrationId?: pulumi.Input<string>;
    /**
     * If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
     */
    key?: pulumi.Input<string>;
    /**
     * Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that Splunk Observability collects for this metric. If you specify this property, Splunk Observability retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, Splunk Observability retrieves the AWS standard set of statistics.
     */
    metricStatsToSyncs?: pulumi.Input<pulumi.Input<inputs.aws.IntegrationMetricStatsToSync>[]>;
    /**
     * Name of the integration.
     */
    name?: pulumi.Input<string>;
    /**
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     */
    namedToken?: pulumi.Input<string>;
    /**
     * Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that Splunk Observability collects for the namespace. Conflicts with the `services` property. If you don't specify either property, Splunk Observability syncs all data in all AWS namespaces.
     */
    namespaceSyncRules?: pulumi.Input<pulumi.Input<inputs.aws.IntegrationNamespaceSyncRule>[]>;
    /**
     * AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     */
    pollRate?: pulumi.Input<number>;
    /**
     * List of AWS regions that Splunk Observability should monitor.
     */
    regions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
     */
    roleArn?: pulumi.Input<string>;
    /**
     * List of AWS services that you want Splunk Observability to monitor. Each element is a string designating an AWS service. Can be an empty list to import data for all supported services. Conflicts with `namespaceSyncRule`. See [Amazon Web Services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#amazon-web-services) for a list of valid values.
     */
    services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates that Splunk Observability should sync metrics and metadata from custom AWS namespaces only (see the `customNamespaceSyncRule` above). Defaults to `false`.
     */
    syncCustomNamespacesOnly?: pulumi.Input<boolean>;
    /**
     * If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
     */
    token?: pulumi.Input<string>;
    /**
     * Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics.<br>
     * Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.<br>
     * Note you need to deploy additional resources on your AWS account to enable CloudWatch metrics streaming. Select one of the [CloudFormation templates](https://docs.splunk.com/Observability/gdi/get-data-in/connect/aws/aws-cloudformation.html) to deploy all the required resources.
     */
    useMetricStreamsSync?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Integration resource.
 */
export interface IntegrationArgs {
    /**
     * List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; Splunk Observability imports the metrics so you can monitor them.
     */
    customCloudwatchNamespaces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Each element controls the data collected by Splunk Observability for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
     */
    customNamespaceSyncRules?: pulumi.Input<pulumi.Input<inputs.aws.IntegrationCustomNamespaceSyncRule>[]>;
    /**
     * Flag that controls how Splunk Observability imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, Splunk Observability imports the metrics.
     */
    enableAwsUsage?: pulumi.Input<boolean>;
    /**
     * Controls how Splunk Observability checks for large amounts of data for this AWS integration. If `true`, Splunk Observability monitors the amount of data coming in from the integration.
     */
    enableCheckLargeVolume?: pulumi.Input<boolean>;
    /**
     * Enable the AWS logs synchronization. Note that this requires the inclusion of `"logs:DescribeLogGroups"`,  `"logs:DeleteSubscriptionFilter"`, `"logs:DescribeSubscriptionFilters"`, `"logs:PutSubscriptionFilter"`, and `"s3:GetBucketLogging"`,  `"s3:GetBucketNotification"`, `"s3:PutBucketNotification"` permissions. Additional permissions may be required to capture logs from specific AWS services.
     */
    enableLogsSync?: pulumi.Input<boolean>;
    /**
     * Whether the integration is enabled.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The `externalId` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
     */
    externalId?: pulumi.Input<string>;
    /**
     * Flag that controls how Splunk Observability imports Cloud Watch metrics. If true, Splunk Observability imports Cloud Watch metrics from AWS.
     */
    importCloudWatch?: pulumi.Input<boolean>;
    /**
     * The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
     */
    integrationId: pulumi.Input<string>;
    /**
     * If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key (this is typically equivalent to the `AWS_SECRET_ACCESS_KEY` environment variable).
     */
    key?: pulumi.Input<string>;
    /**
     * Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics that Splunk Observability collects for this metric. If you specify this property, Splunk Observability retrieves only specified AWS statistics when AWS metric streams are not used. When AWS metric streams are used this property specifies additional extended statistics to collect (please note that AWS metric streams API supports percentile stats only; other stats are ignored). If you don't specify this property, Splunk Observability retrieves the AWS standard set of statistics.
     */
    metricStatsToSyncs?: pulumi.Input<pulumi.Input<inputs.aws.IntegrationMetricStatsToSync>[]>;
    /**
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     */
    namedToken?: pulumi.Input<string>;
    /**
     * Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that Splunk Observability collects for the namespace. Conflicts with the `services` property. If you don't specify either property, Splunk Observability syncs all data in all AWS namespaces.
     */
    namespaceSyncRules?: pulumi.Input<pulumi.Input<inputs.aws.IntegrationNamespaceSyncRule>[]>;
    /**
     * AWS poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     */
    pollRate?: pulumi.Input<number>;
    /**
     * List of AWS regions that Splunk Observability should monitor.
     */
    regions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
     */
    roleArn?: pulumi.Input<string>;
    /**
     * List of AWS services that you want Splunk Observability to monitor. Each element is a string designating an AWS service. Can be an empty list to import data for all supported services. Conflicts with `namespaceSyncRule`. See [Amazon Web Services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#amazon-web-services) for a list of valid values.
     */
    services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates that Splunk Observability should sync metrics and metadata from custom AWS namespaces only (see the `customNamespaceSyncRule` above). Defaults to `false`.
     */
    syncCustomNamespacesOnly?: pulumi.Input<boolean>;
    /**
     * If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the token (this is typically equivalent to the `AWS_ACCESS_KEY_ID` environment variable).
     */
    token?: pulumi.Input<string>;
    /**
     * Enable the use of Amazon Cloudwatch Metric Streams for ingesting metrics.<br>
     * Note that this requires the inclusion of `"cloudwatch:ListMetricStreams"`,`"cloudwatch:GetMetricStream"`, `"cloudwatch:PutMetricStream"`, `"cloudwatch:DeleteMetricStream"`, `"cloudwatch:StartMetricStreams"`, `"cloudwatch:StopMetricStreams"` and `"iam:PassRole"` permissions.<br>
     * Note you need to deploy additional resources on your AWS account to enable CloudWatch metrics streaming. Select one of the [CloudFormation templates](https://docs.splunk.com/Observability/gdi/get-data-in/connect/aws/aws-cloudformation.html) to deploy all the required resources.
     */
    useMetricStreamsSync?: pulumi.Input<boolean>;
}
