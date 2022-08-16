// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * SignalFx GCP Integration
 *
 * > **NOTE** When managing integrations use a session token for an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const gcpMyteam = new signalfx.gcp.Integration("gcp_myteam", {
 *     enabled: true,
 *     pollRate: 300,
 *     projectServiceKeys: [
 *         {
 *             projectId: "gcp_project_id_1",
 *             projectKey: fs.readFileSync("/path/to/gcp_credentials_1.json", "utf-8"),
 *         },
 *         {
 *             projectId: "gcp_project_id_2",
 *             projectKey: fs.readFileSync("/path/to/gcp_credentials_2.json", "utf-8"),
 *         },
 *     ],
 *     services: ["compute"],
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
    public static readonly __pulumiType = 'signalfx:gcp/integration:Integration';

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
     * Whether the integration is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Name of the integration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     */
    public readonly namedToken!: pulumi.Output<string | undefined>;
    /**
     * GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     */
    public readonly pollRate!: pulumi.Output<number | undefined>;
    /**
     * GCP projects to add.
     */
    public readonly projectServiceKeys!: pulumi.Output<outputs.gcp.IntegrationProjectServiceKey[] | undefined>;
    /**
     * GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See the documentation for [Creating Integrations](https://dev.splunk.com/observability/reference/api/integrations/latest#endpoint-create-integration) for valid values.
     */
    public readonly services!: pulumi.Output<string[] | undefined>;
    /**
     * When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
     */
    public readonly useMetricSourceProjectForQuota!: pulumi.Output<boolean | undefined>;
    /**
     * [Compute Metadata Whitelist](https://docs.splunk.com/Observability/infrastructure/navigators/gcp.html#compute-engine-instance).
     */
    public readonly whitelists!: pulumi.Output<string[] | undefined>;

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
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namedToken"] = state ? state.namedToken : undefined;
            resourceInputs["pollRate"] = state ? state.pollRate : undefined;
            resourceInputs["projectServiceKeys"] = state ? state.projectServiceKeys : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["useMetricSourceProjectForQuota"] = state ? state.useMetricSourceProjectForQuota : undefined;
            resourceInputs["whitelists"] = state ? state.whitelists : undefined;
        } else {
            const args = argsOrState as IntegrationArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namedToken"] = args ? args.namedToken : undefined;
            resourceInputs["pollRate"] = args ? args.pollRate : undefined;
            resourceInputs["projectServiceKeys"] = args ? args.projectServiceKeys : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["useMetricSourceProjectForQuota"] = args ? args.useMetricSourceProjectForQuota : undefined;
            resourceInputs["whitelists"] = args ? args.whitelists : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Integration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Integration resources.
 */
export interface IntegrationState {
    /**
     * Whether the integration is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Name of the integration.
     */
    name?: pulumi.Input<string>;
    /**
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     */
    namedToken?: pulumi.Input<string>;
    /**
     * GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     */
    pollRate?: pulumi.Input<number>;
    /**
     * GCP projects to add.
     */
    projectServiceKeys?: pulumi.Input<pulumi.Input<inputs.gcp.IntegrationProjectServiceKey>[]>;
    /**
     * GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See the documentation for [Creating Integrations](https://dev.splunk.com/observability/reference/api/integrations/latest#endpoint-create-integration) for valid values.
     */
    services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
     */
    useMetricSourceProjectForQuota?: pulumi.Input<boolean>;
    /**
     * [Compute Metadata Whitelist](https://docs.splunk.com/Observability/infrastructure/navigators/gcp.html#compute-engine-instance).
     */
    whitelists?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Integration resource.
 */
export interface IntegrationArgs {
    /**
     * Whether the integration is enabled.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Name of the integration.
     */
    name?: pulumi.Input<string>;
    /**
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     */
    namedToken?: pulumi.Input<string>;
    /**
     * GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     */
    pollRate?: pulumi.Input<number>;
    /**
     * GCP projects to add.
     */
    projectServiceKeys?: pulumi.Input<pulumi.Input<inputs.gcp.IntegrationProjectServiceKey>[]>;
    /**
     * GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See the documentation for [Creating Integrations](https://dev.splunk.com/observability/reference/api/integrations/latest#endpoint-create-integration) for valid values.
     */
    services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
     */
    useMetricSourceProjectForQuota?: pulumi.Input<boolean>;
    /**
     * [Compute Metadata Whitelist](https://docs.splunk.com/Observability/infrastructure/navigators/gcp.html#compute-engine-instance).
     */
    whitelists?: pulumi.Input<pulumi.Input<string>[]>;
}
