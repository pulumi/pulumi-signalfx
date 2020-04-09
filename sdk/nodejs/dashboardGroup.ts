// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * In the SignalFx web UI, a [dashboard group](https://developers.signalfx.com/dashboard_groups_reference.html) is a collection of dashboards.
 * 
 * > **NOTE** Dashboard groups cannot be accessed directly, but just via a dashboard contained in them. This is the reason why make show won't show any of yours dashboard groups.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/dashboard_group.html.markdown.
 */
export class DashboardGroup extends pulumi.CustomResource {
    /**
     * Get an existing DashboardGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardGroupState, opts?: pulumi.CustomResourceOptions): DashboardGroup {
        return new DashboardGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/dashboardGroup:DashboardGroup';

    /**
     * Returns true if the given object is an instance of DashboardGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DashboardGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DashboardGroup.__pulumiType;
    }

    /**
     * Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
     */
    public readonly authorizedWriterTeams!: pulumi.Output<string[] | undefined>;
    /**
     * User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
     */
    public readonly authorizedWriterUsers!: pulumi.Output<string[] | undefined>;
    /**
     * [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     */
    public readonly dashboards!: pulumi.Output<outputs.DashboardGroupDashboard[] | undefined>;
    /**
     * Description of the dashboard group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly importQualifiers!: pulumi.Output<outputs.DashboardGroupImportQualifier[] | undefined>;
    /**
     * Name of the dashboard group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Team IDs to associate the dashboard group to.
     */
    public readonly teams!: pulumi.Output<string[] | undefined>;

    /**
     * Create a DashboardGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DashboardGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardGroupArgs | DashboardGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DashboardGroupState | undefined;
            inputs["authorizedWriterTeams"] = state ? state.authorizedWriterTeams : undefined;
            inputs["authorizedWriterUsers"] = state ? state.authorizedWriterUsers : undefined;
            inputs["dashboards"] = state ? state.dashboards : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["importQualifiers"] = state ? state.importQualifiers : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["teams"] = state ? state.teams : undefined;
        } else {
            const args = argsOrState as DashboardGroupArgs | undefined;
            inputs["authorizedWriterTeams"] = args ? args.authorizedWriterTeams : undefined;
            inputs["authorizedWriterUsers"] = args ? args.authorizedWriterUsers : undefined;
            inputs["dashboards"] = args ? args.dashboards : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["importQualifiers"] = args ? args.importQualifiers : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["teams"] = args ? args.teams : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DashboardGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DashboardGroup resources.
 */
export interface DashboardGroupState {
    /**
     * Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterTeams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     */
    readonly dashboards?: pulumi.Input<pulumi.Input<inputs.DashboardGroupDashboard>[]>;
    /**
     * Description of the dashboard group.
     */
    readonly description?: pulumi.Input<string>;
    readonly importQualifiers?: pulumi.Input<pulumi.Input<inputs.DashboardGroupImportQualifier>[]>;
    /**
     * Name of the dashboard group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Team IDs to associate the dashboard group to.
     */
    readonly teams?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DashboardGroup resource.
 */
export interface DashboardGroupArgs {
    /**
     * Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterTeams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     */
    readonly dashboards?: pulumi.Input<pulumi.Input<inputs.DashboardGroupDashboard>[]>;
    /**
     * Description of the dashboard group.
     */
    readonly description?: pulumi.Input<string>;
    readonly importQualifiers?: pulumi.Input<pulumi.Input<inputs.DashboardGroupImportQualifier>[]>;
    /**
     * Name of the dashboard group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Team IDs to associate the dashboard group to.
     */
    readonly teams?: pulumi.Input<pulumi.Input<string>[]>;
}
