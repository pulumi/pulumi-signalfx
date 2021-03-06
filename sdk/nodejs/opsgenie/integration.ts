// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * SignalFx Opsgenie integration.
 *
 * > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const opgenieMyteam = new signalfx.opsgenie.Integration("opgenie_myteam", {
 *     apiKey: "farts",
 *     apiUrl: "https://api.opsgenie.com",
 *     enabled: true,
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
    public static readonly __pulumiType = 'signalfx:opsgenie/integration:Integration';

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
     * The API key
     */
    public readonly apiKey!: pulumi.Output<string>;
    /**
     * Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
     */
    public readonly apiUrl!: pulumi.Output<string | undefined>;
    /**
     * Whether the integration is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Name of the integration.
     */
    public readonly name!: pulumi.Output<string>;

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
            inputs["apiKey"] = state ? state.apiKey : undefined;
            inputs["apiUrl"] = state ? state.apiUrl : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as IntegrationArgs | undefined;
            if ((!args || args.apiKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiKey'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            inputs["apiKey"] = args ? args.apiKey : undefined;
            inputs["apiUrl"] = args ? args.apiUrl : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["name"] = args ? args.name : undefined;
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
     * The API key
     */
    readonly apiKey?: pulumi.Input<string>;
    /**
     * Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
     */
    readonly apiUrl?: pulumi.Input<string>;
    /**
     * Whether the integration is enabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Name of the integration.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Integration resource.
 */
export interface IntegrationArgs {
    /**
     * The API key
     */
    readonly apiKey: pulumi.Input<string>;
    /**
     * Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
     */
    readonly apiUrl?: pulumi.Input<string>;
    /**
     * Whether the integration is enabled.
     */
    readonly enabled: pulumi.Input<boolean>;
    /**
     * Name of the integration.
     */
    readonly name?: pulumi.Input<string>;
}
