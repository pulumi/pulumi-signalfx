// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manage Splunk Observability Cloud org tokens.
 *
 * > **NOTE** When managing Org tokens, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
 *
 * ## Example
 *
 * ## Arguments
 *
 * The following arguments are supported in the resource block:
 *
 * * `name` - (Required) Name of the token.
 * * `description` - (Optional) Description of the token.
 * * `disabled` - (Optional) Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
 * * `secret` - The secret token created by the API. You cannot set this value.
 * * `notifications` - (Optional) Where to send notifications about this token's limits. See the Notification Format laid out in detectors.
 * * `hostOrUsageLimits` - (Optional) Specify Usage-based limits for this token.
 *   * `hostLimit` - (Optional) Max number of hosts that can use this token
 *   * `hostNotificationThreshold` - (Optional) Notification threshold for hosts
 *   * `containerLimit` - (Optional) Max number of Docker containers that can use this token
 *   * `containerNotificationThreshold` - (Optional) Notification threshold for Docker containers
 *   * `customMetricsLimit` - (Optional) Max number of custom metrics that can be sent with this token
 *   * `customMetricsNotificationThreshold` - (Optional) Notification threshold for custom metrics
 *   * `highResMetricsLimit` - (Optional) Max number of hi-res metrics that can be sent with this toke
 *   * `highResMetricsNotificationThreshold` - (Optional) Notification threshold for hi-res metrics
 * * `dpmLimits` (Optional) Specify DPM-based limits for this token.
 *   * `dpmNotificationThreshold` - (Optional) DPM level at which Splunk Observability Cloud sends the notification for this token. If you don't specify a notification, Splunk Observability Cloud sends the generic notification.
 *   * `dpmLimit` - (Required) The datapoints per minute (dpm) limit for this token. If you exceed this limit, Splunk Observability Cloud sends out an alert.
 *
 * ## Attributes
 *
 * In a addition to all arguments above, the following attributes are exported:
 *
 * * `id` - The ID of the token.
 * * `secret` - The assigned token.
 */
export class OrgToken extends pulumi.CustomResource {
    /**
     * Get an existing OrgToken resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgTokenState, opts?: pulumi.CustomResourceOptions): OrgToken {
        return new OrgToken(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/orgToken:OrgToken';

    /**
     * Returns true if the given object is an instance of OrgToken.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgToken {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgToken.__pulumiType;
    }

    /**
     * Authentication scope, ex: INGEST, API, RUM ... (Optional)
     */
    public readonly authScopes!: pulumi.Output<string[]>;
    /**
     * Description of the token (Optional)
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication.
     * Defaults to `false`
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    public readonly dpmLimits!: pulumi.Output<outputs.OrgTokenDpmLimits | undefined>;
    public readonly hostOrUsageLimits!: pulumi.Output<outputs.OrgTokenHostOrUsageLimits | undefined>;
    /**
     * Name of the token
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See
     * https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     */
    public readonly notifications!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly secret!: pulumi.Output<string>;

    /**
     * Create a OrgToken resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OrgTokenArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgTokenArgs | OrgTokenState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgTokenState | undefined;
            resourceInputs["authScopes"] = state ? state.authScopes : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["dpmLimits"] = state ? state.dpmLimits : undefined;
            resourceInputs["hostOrUsageLimits"] = state ? state.hostOrUsageLimits : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notifications"] = state ? state.notifications : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
        } else {
            const args = argsOrState as OrgTokenArgs | undefined;
            resourceInputs["authScopes"] = args ? args.authScopes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["dpmLimits"] = args ? args.dpmLimits : undefined;
            resourceInputs["hostOrUsageLimits"] = args ? args.hostOrUsageLimits : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notifications"] = args ? args.notifications : undefined;
            resourceInputs["secret"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["secret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(OrgToken.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgToken resources.
 */
export interface OrgTokenState {
    /**
     * Authentication scope, ex: INGEST, API, RUM ... (Optional)
     */
    authScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the token (Optional)
     */
    description?: pulumi.Input<string>;
    /**
     * Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication.
     * Defaults to `false`
     */
    disabled?: pulumi.Input<boolean>;
    dpmLimits?: pulumi.Input<inputs.OrgTokenDpmLimits>;
    hostOrUsageLimits?: pulumi.Input<inputs.OrgTokenHostOrUsageLimits>;
    /**
     * Name of the token
     */
    name?: pulumi.Input<string>;
    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See
     * https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     */
    notifications?: pulumi.Input<pulumi.Input<string>[]>;
    secret?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrgToken resource.
 */
export interface OrgTokenArgs {
    /**
     * Authentication scope, ex: INGEST, API, RUM ... (Optional)
     */
    authScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the token (Optional)
     */
    description?: pulumi.Input<string>;
    /**
     * Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication.
     * Defaults to `false`
     */
    disabled?: pulumi.Input<boolean>;
    dpmLimits?: pulumi.Input<inputs.OrgTokenDpmLimits>;
    hostOrUsageLimits?: pulumi.Input<inputs.OrgTokenHostOrUsageLimits>;
    /**
     * Name of the token
     */
    name?: pulumi.Input<string>;
    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See
     * https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     */
    notifications?: pulumi.Input<pulumi.Input<string>[]>;
}
