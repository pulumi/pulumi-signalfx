// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Handles management of Splunk Observability Cloud teams.
 *
 * You can configure [team notification policies](https://docs.splunk.com/observability/en/admin/user-management/teams/team-notifications.html) using this resource and the various `notifications_*` properties.
 *
 * > **NOTE** When managing teams, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
 *
 * ## Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const myteam0 = new signalfx.Team("myteam0", {
 *     name: "Best Team Ever",
 *     description: "Super great team no jerks definitely",
 *     members: [
 *         "userid1",
 *         "userid2",
 *     ],
 *     notificationsCriticals: ["PagerDuty,credentialId"],
 *     notificationsInfos: ["Email,notify@example.com"],
 * });
 * ```
 */
export class Team extends pulumi.CustomResource {
    /**
     * Get an existing Team resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamState, opts?: pulumi.CustomResourceOptions): Team {
        return new Team(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/team:Team';

    /**
     * Returns true if the given object is an instance of Team.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Team {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Team.__pulumiType;
    }

    /**
     * Description of the team.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of user IDs to include in the team.
     */
    public readonly members!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the team.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Where to send notifications for critical alerts
     */
    public readonly notificationsCriticals!: pulumi.Output<string[] | undefined>;
    /**
     * Where to send notifications for default alerts
     */
    public readonly notificationsDefaults!: pulumi.Output<string[] | undefined>;
    /**
     * Where to send notifications for info alerts
     */
    public readonly notificationsInfos!: pulumi.Output<string[] | undefined>;
    /**
     * Where to send notifications for major alerts
     */
    public readonly notificationsMajors!: pulumi.Output<string[] | undefined>;
    /**
     * Where to send notifications for minor alerts
     */
    public readonly notificationsMinors!: pulumi.Output<string[] | undefined>;
    /**
     * Where to send notifications for warning alerts
     */
    public readonly notificationsWarnings!: pulumi.Output<string[] | undefined>;
    /**
     * The URL of the team.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Team resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TeamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamArgs | TeamState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TeamState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationsCriticals"] = state ? state.notificationsCriticals : undefined;
            resourceInputs["notificationsDefaults"] = state ? state.notificationsDefaults : undefined;
            resourceInputs["notificationsInfos"] = state ? state.notificationsInfos : undefined;
            resourceInputs["notificationsMajors"] = state ? state.notificationsMajors : undefined;
            resourceInputs["notificationsMinors"] = state ? state.notificationsMinors : undefined;
            resourceInputs["notificationsWarnings"] = state ? state.notificationsWarnings : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as TeamArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationsCriticals"] = args ? args.notificationsCriticals : undefined;
            resourceInputs["notificationsDefaults"] = args ? args.notificationsDefaults : undefined;
            resourceInputs["notificationsInfos"] = args ? args.notificationsInfos : undefined;
            resourceInputs["notificationsMajors"] = args ? args.notificationsMajors : undefined;
            resourceInputs["notificationsMinors"] = args ? args.notificationsMinors : undefined;
            resourceInputs["notificationsWarnings"] = args ? args.notificationsWarnings : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Team.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Team resources.
 */
export interface TeamState {
    /**
     * Description of the team.
     */
    description?: pulumi.Input<string>;
    /**
     * List of user IDs to include in the team.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the team.
     */
    name?: pulumi.Input<string>;
    /**
     * Where to send notifications for critical alerts
     */
    notificationsCriticals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for default alerts
     */
    notificationsDefaults?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for info alerts
     */
    notificationsInfos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for major alerts
     */
    notificationsMajors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for minor alerts
     */
    notificationsMinors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for warning alerts
     */
    notificationsWarnings?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL of the team.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Team resource.
 */
export interface TeamArgs {
    /**
     * Description of the team.
     */
    description?: pulumi.Input<string>;
    /**
     * List of user IDs to include in the team.
     */
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the team.
     */
    name?: pulumi.Input<string>;
    /**
     * Where to send notifications for critical alerts
     */
    notificationsCriticals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for default alerts
     */
    notificationsDefaults?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for info alerts
     */
    notificationsInfos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for major alerts
     */
    notificationsMajors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for minor alerts
     */
    notificationsMinors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for warning alerts
     */
    notificationsWarnings?: pulumi.Input<pulumi.Input<string>[]>;
}
