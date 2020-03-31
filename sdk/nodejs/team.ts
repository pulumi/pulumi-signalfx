// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Handles management of SignalFx teams.
 * 
 * You can configure [team notification policies](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html) using this resource and the various `notifications_*` properties.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/team.html.markdown.
 */
export class Team extends pulumi.CustomResource {
    /**
     * Get an existing Team resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * URL of the team
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TeamState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["members"] = state ? state.members : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationsCriticals"] = state ? state.notificationsCriticals : undefined;
            inputs["notificationsDefaults"] = state ? state.notificationsDefaults : undefined;
            inputs["notificationsInfos"] = state ? state.notificationsInfos : undefined;
            inputs["notificationsMajors"] = state ? state.notificationsMajors : undefined;
            inputs["notificationsMinors"] = state ? state.notificationsMinors : undefined;
            inputs["notificationsWarnings"] = state ? state.notificationsWarnings : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as TeamArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["members"] = args ? args.members : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationsCriticals"] = args ? args.notificationsCriticals : undefined;
            inputs["notificationsDefaults"] = args ? args.notificationsDefaults : undefined;
            inputs["notificationsInfos"] = args ? args.notificationsInfos : undefined;
            inputs["notificationsMajors"] = args ? args.notificationsMajors : undefined;
            inputs["notificationsMinors"] = args ? args.notificationsMinors : undefined;
            inputs["notificationsWarnings"] = args ? args.notificationsWarnings : undefined;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Team.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Team resources.
 */
export interface TeamState {
    /**
     * Description of the team.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of user IDs to include in the team.
     */
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the team.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Where to send notifications for critical alerts
     */
    readonly notificationsCriticals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for default alerts
     */
    readonly notificationsDefaults?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for info alerts
     */
    readonly notificationsInfos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for major alerts
     */
    readonly notificationsMajors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for minor alerts
     */
    readonly notificationsMinors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for warning alerts
     */
    readonly notificationsWarnings?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URL of the team
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Team resource.
 */
export interface TeamArgs {
    /**
     * Description of the team.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of user IDs to include in the team.
     */
    readonly members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the team.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Where to send notifications for critical alerts
     */
    readonly notificationsCriticals?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for default alerts
     */
    readonly notificationsDefaults?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for info alerts
     */
    readonly notificationsInfos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for major alerts
     */
    readonly notificationsMajors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for minor alerts
     */
    readonly notificationsMinors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where to send notifications for warning alerts
     */
    readonly notificationsWarnings?: pulumi.Input<pulumi.Input<string>[]>;
}