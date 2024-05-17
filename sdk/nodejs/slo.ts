// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a Splunk Observability Cloud slo resource. This can be used to create and manage SLOs.
 *
 * To learn more about this feature take a look on [documentation for SLO](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/slo-intro.html).
 *
 * ## Example
 *
 * ## Notification format
 *
 * As Splunk Observability Cloud supports different notification mechanisms, use a comma-delimited string to provide inputs. If you want to specify multiple notifications, each must be a member in the list, like so:
 *
 * See [Splunk Observability Cloud Docs](https://dev.splunk.com/observability/reference/api/detectors/latest) for more information.
 *
 * Here are some example of how to configure each notification type:
 *
 * ### Email
 *
 * ### Jira
 *
 * Note that the `credentialId` is the Splunk-provided ID shown after setting up your Jira integration. See also `signalfx.jira.Integration`.
 *
 * ### OpsGenie
 *
 * Note that the `credentialId` is the Splunk-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.
 *
 * ### PagerDuty
 *
 * ### Slack
 *
 * Exclude the `#` on the channel name:
 *
 * ### Team
 *
 * Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).
 *
 * ### TeamEmail
 *
 * Sends an email to every member of a team.
 *
 * ### Splunk On-Call (formerly VictorOps)
 *
 * ### Webhooks
 *
 * You need to include all the commas even if you only use a credential id.
 *
 * You can either configure a Webhook to use an existing integration's credential id:
 *
 * Or configure one inline:
 */
export class Slo extends pulumi.CustomResource {
    /**
     * Get an existing Slo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SloState, opts?: pulumi.CustomResourceOptions): Slo {
        return new Slo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/slo:Slo';

    /**
     * Returns true if the given object is an instance of Slo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Slo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Slo.__pulumiType;
    }

    /**
     * Description of the SLO.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Properties to configure an SLO object inputs
     */
    public readonly input!: pulumi.Output<outputs.SloInput>;
    /**
     * Name of the SLO. Each SLO name must be unique within an organization.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Define target value of the service level indicator in the appropriate time period.
     */
    public readonly target!: pulumi.Output<outputs.SloTarget>;
    /**
     * Type of the SLO. Currently just: `"RequestBased"` is supported.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Slo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SloArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SloArgs | SloState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SloState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["input"] = state ? state.input : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SloArgs | undefined;
            if ((!args || args.input === undefined) && !opts.urn) {
                throw new Error("Missing required property 'input'");
            }
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["input"] = args ? args.input : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Slo.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Slo resources.
 */
export interface SloState {
    /**
     * Description of the SLO.
     */
    description?: pulumi.Input<string>;
    /**
     * Properties to configure an SLO object inputs
     */
    input?: pulumi.Input<inputs.SloInput>;
    /**
     * Name of the SLO. Each SLO name must be unique within an organization.
     */
    name?: pulumi.Input<string>;
    /**
     * Define target value of the service level indicator in the appropriate time period.
     */
    target?: pulumi.Input<inputs.SloTarget>;
    /**
     * Type of the SLO. Currently just: `"RequestBased"` is supported.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Slo resource.
 */
export interface SloArgs {
    /**
     * Description of the SLO.
     */
    description?: pulumi.Input<string>;
    /**
     * Properties to configure an SLO object inputs
     */
    input: pulumi.Input<inputs.SloInput>;
    /**
     * Name of the SLO. Each SLO name must be unique within an organization.
     */
    name?: pulumi.Input<string>;
    /**
     * Define target value of the service level indicator in the appropriate time period.
     */
    target: pulumi.Input<inputs.SloTarget>;
    /**
     * Type of the SLO. Currently just: `"RequestBased"` is supported.
     */
    type: pulumi.Input<string>;
}
