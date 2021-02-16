// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * SignalFx AWS CloudWatch integrations. For help with this integration see [Monitoring Amazon Web Services](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#monitor-amazon-web-services).
 *
 * > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.
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
 *         namespace: "fart",
 *     }],
 *     namespaceSyncRules: [{
 *         defaultAction: "Exclude",
 *         filterAction: "Include",
 *         filterSource: "filter('code', '200')",
 *         namespace: "AWS/EC2",
 *     }],
 * });
 * ```
 * ## Service Names
 *
 * > **NOTE** You can use the data source "signalfx.aws.getServices" to specify all services.
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
     * List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
     */
    public readonly customCloudwatchNamespaces!: pulumi.Output<string[] | undefined>;
    /**
     * Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
     */
    public readonly customNamespaceSyncRules!: pulumi.Output<outputs.aws.IntegrationCustomNamespaceSyncRule[] | undefined>;
    /**
     * Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
     */
    public readonly enableAwsUsage!: pulumi.Output<boolean | undefined>;
    /**
     * Controls how SignalFx checks for large amounts of data for this AWS integration. If `true`, SignalFx monitors the amount of data coming in from the integration.
     */
    public readonly enableCheckLargeVolume!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the integration is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The `externalId` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
     */
    public readonly externalId!: pulumi.Output<string | undefined>;
    /**
     * Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
     */
    public readonly importCloudWatch!: pulumi.Output<boolean | undefined>;
    /**
     * The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
     */
    public readonly integrationId!: pulumi.Output<string>;
    /**
     * If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * A named token to use for ingest
     */
    public readonly namedToken!: pulumi.Output<string | undefined>;
    /**
     * Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
     */
    public readonly namespaceSyncRules!: pulumi.Output<outputs.aws.IntegrationNamespaceSyncRule[] | undefined>;
    /**
     * AWS poll rate (in seconds). Value between `60` and `300`.
     */
    public readonly pollRate!: pulumi.Output<number | undefined>;
    /**
     * List of AWS regions that SignalFx should monitor.
     */
    public readonly regions!: pulumi.Output<string[] | undefined>;
    /**
     * Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
     */
    public readonly services!: pulumi.Output<string[] | undefined>;
    /**
     * Used with `signalfx_aws_token_integration`. Use this property to specify the token.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
     */
    public readonly useGetMetricDataMethod!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Integration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationArgs | IntegrationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationState | undefined;
            inputs["customCloudwatchNamespaces"] = state ? state.customCloudwatchNamespaces : undefined;
            inputs["customNamespaceSyncRules"] = state ? state.customNamespaceSyncRules : undefined;
            inputs["enableAwsUsage"] = state ? state.enableAwsUsage : undefined;
            inputs["enableCheckLargeVolume"] = state ? state.enableCheckLargeVolume : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["externalId"] = state ? state.externalId : undefined;
            inputs["importCloudWatch"] = state ? state.importCloudWatch : undefined;
            inputs["integrationId"] = state ? state.integrationId : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["namedToken"] = state ? state.namedToken : undefined;
            inputs["namespaceSyncRules"] = state ? state.namespaceSyncRules : undefined;
            inputs["pollRate"] = state ? state.pollRate : undefined;
            inputs["regions"] = state ? state.regions : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["services"] = state ? state.services : undefined;
            inputs["token"] = state ? state.token : undefined;
            inputs["useGetMetricDataMethod"] = state ? state.useGetMetricDataMethod : undefined;
        } else {
            const args = argsOrState as IntegrationArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.integrationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationId'");
            }
            inputs["customCloudwatchNamespaces"] = args ? args.customCloudwatchNamespaces : undefined;
            inputs["customNamespaceSyncRules"] = args ? args.customNamespaceSyncRules : undefined;
            inputs["enableAwsUsage"] = args ? args.enableAwsUsage : undefined;
            inputs["enableCheckLargeVolume"] = args ? args.enableCheckLargeVolume : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["externalId"] = args ? args.externalId : undefined;
            inputs["importCloudWatch"] = args ? args.importCloudWatch : undefined;
            inputs["integrationId"] = args ? args.integrationId : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["namedToken"] = args ? args.namedToken : undefined;
            inputs["namespaceSyncRules"] = args ? args.namespaceSyncRules : undefined;
            inputs["pollRate"] = args ? args.pollRate : undefined;
            inputs["regions"] = args ? args.regions : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["services"] = args ? args.services : undefined;
            inputs["token"] = args ? args.token : undefined;
            inputs["useGetMetricDataMethod"] = args ? args.useGetMetricDataMethod : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Integration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Integration resources.
 */
export interface IntegrationState {
    /**
     * List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
     */
    readonly customCloudwatchNamespaces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
     */
    readonly customNamespaceSyncRules?: pulumi.Input<pulumi.Input<inputs.aws.IntegrationCustomNamespaceSyncRule>[]>;
    /**
     * Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
     */
    readonly enableAwsUsage?: pulumi.Input<boolean>;
    /**
     * Controls how SignalFx checks for large amounts of data for this AWS integration. If `true`, SignalFx monitors the amount of data coming in from the integration.
     */
    readonly enableCheckLargeVolume?: pulumi.Input<boolean>;
    /**
     * Whether the integration is enabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The `externalId` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
     */
    readonly externalId?: pulumi.Input<string>;
    /**
     * Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
     */
    readonly importCloudWatch?: pulumi.Input<boolean>;
    /**
     * The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
     */
    readonly integrationId?: pulumi.Input<string>;
    /**
     * If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * A named token to use for ingest
     */
    readonly namedToken?: pulumi.Input<string>;
    /**
     * Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
     */
    readonly namespaceSyncRules?: pulumi.Input<pulumi.Input<inputs.aws.IntegrationNamespaceSyncRule>[]>;
    /**
     * AWS poll rate (in seconds). Value between `60` and `300`.
     */
    readonly pollRate?: pulumi.Input<number>;
    /**
     * List of AWS regions that SignalFx should monitor.
     */
    readonly regions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
     */
    readonly services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used with `signalfx_aws_token_integration`. Use this property to specify the token.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
     */
    readonly useGetMetricDataMethod?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Integration resource.
 */
export interface IntegrationArgs {
    /**
     * List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS; SignalFx imports the metrics so you can monitor them.
     */
    readonly customCloudwatchNamespaces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Each element controls the data collected by SignalFx for the specified namespace. Conflicts with the `customCloudwatchNamespaces` property.
     */
    readonly customNamespaceSyncRules?: pulumi.Input<pulumi.Input<inputs.aws.IntegrationCustomNamespaceSyncRule>[]>;
    /**
     * Flag that controls how SignalFx imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`, SignalFx imports the metrics.
     */
    readonly enableAwsUsage?: pulumi.Input<boolean>;
    /**
     * Controls how SignalFx checks for large amounts of data for this AWS integration. If `true`, SignalFx monitors the amount of data coming in from the integration.
     */
    readonly enableCheckLargeVolume?: pulumi.Input<boolean>;
    /**
     * Whether the integration is enabled.
     */
    readonly enabled: pulumi.Input<boolean>;
    /**
     * The `externalId` property from one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`
     */
    readonly externalId?: pulumi.Input<string>;
    /**
     * Flag that controls how SignalFx imports Cloud Watch metrics. If true, SignalFx imports Cloud Watch metrics from AWS.
     */
    readonly importCloudWatch?: pulumi.Input<boolean>;
    /**
     * The id of one of a `signalfx.aws.ExternalIntegration` or `signalfx.aws.TokenIntegration`.
     */
    readonly integrationId: pulumi.Input<string>;
    /**
     * If you specify `authMethod = \"SecurityToken\"` in your request to create an AWS integration object, use this property to specify the key.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * A named token to use for ingest
     */
    readonly namedToken?: pulumi.Input<string>;
    /**
     * Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that SignalFx collects for the namespace. Conflicts with the `services` property. If you don't specify either property, SignalFx syncs all data in all AWS namespaces.
     */
    readonly namespaceSyncRules?: pulumi.Input<pulumi.Input<inputs.aws.IntegrationNamespaceSyncRule>[]>;
    /**
     * AWS poll rate (in seconds). Value between `60` and `300`.
     */
    readonly pollRate?: pulumi.Input<number>;
    /**
     * List of AWS regions that SignalFx should monitor.
     */
    readonly regions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Role ARN that you add to an existing AWS integration object. **Note**: Ensure you use the `arn` property of your role, not the id!
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * List of AWS services that you want SignalFx to monitor. Each element is a string designating an AWS service. Conflicts with `namespaceSyncRule`. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
     */
    readonly services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used with `signalfx_aws_token_integration`. Use this property to specify the token.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * Enable the use of Amazon's `GetMetricData` for collecting metrics. Note that this requires the inclusion of the `"cloudwatch:GetMetricData"` permission.
     */
    readonly useGetMetricDataMethod?: pulumi.Input<boolean>;
}
