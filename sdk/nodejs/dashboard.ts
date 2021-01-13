// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class Dashboard extends pulumi.CustomResource {
    /**
     * Get an existing Dashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardState, opts?: pulumi.CustomResourceOptions): Dashboard {
        return new Dashboard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/dashboard:Dashboard';

    /**
     * Returns true if the given object is an instance of Dashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dashboard.__pulumiType;
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
     * Chart ID and layout information for the charts in the dashboard.
     */
    public readonly charts!: pulumi.Output<outputs.DashboardChart[] | undefined>;
    /**
     * Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
     */
    public readonly chartsResolution!: pulumi.Output<string | undefined>;
    /**
     * Column number for the layout.
     */
    public readonly columns!: pulumi.Output<outputs.DashboardColumn[] | undefined>;
    /**
     * The ID of the dashboard group that contains the dashboard.
     */
    public readonly dashboardGroup!: pulumi.Output<string>;
    /**
     * Variable description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly discoveryOptionsQuery!: pulumi.Output<string | undefined>;
    public readonly discoveryOptionsSelectors!: pulumi.Output<string[] | undefined>;
    /**
     * Seconds since epoch. Used for visualization.
     */
    public readonly endTime!: pulumi.Output<number | undefined>;
    /**
     * Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selectedEventOverlay` property instead.
     */
    public readonly eventOverlays!: pulumi.Output<outputs.DashboardEventOverlay[] | undefined>;
    /**
     * Filter to apply to the charts when displaying the dashboard.
     */
    public readonly filters!: pulumi.Output<outputs.DashboardFilter[] | undefined>;
    /**
     * Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
     */
    public readonly grids!: pulumi.Output<outputs.DashboardGrid[] | undefined>;
    /**
     * Name of the dashboard.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
     */
    public readonly selectedEventOverlays!: pulumi.Output<outputs.DashboardSelectedEventOverlay[] | undefined>;
    /**
     * Seconds since epoch. Used for visualization.
     */
    public readonly startTime!: pulumi.Output<number | undefined>;
    /**
     * The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
     */
    public readonly timeRange!: pulumi.Output<string | undefined>;
    /**
     * The URL of the dashboard.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Dashboard variable to apply to each chart in the dashboard.
     */
    public readonly variables!: pulumi.Output<outputs.DashboardVariable[] | undefined>;

    /**
     * Create a Dashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardArgs | DashboardState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DashboardState | undefined;
            inputs["authorizedWriterTeams"] = state ? state.authorizedWriterTeams : undefined;
            inputs["authorizedWriterUsers"] = state ? state.authorizedWriterUsers : undefined;
            inputs["charts"] = state ? state.charts : undefined;
            inputs["chartsResolution"] = state ? state.chartsResolution : undefined;
            inputs["columns"] = state ? state.columns : undefined;
            inputs["dashboardGroup"] = state ? state.dashboardGroup : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["discoveryOptionsQuery"] = state ? state.discoveryOptionsQuery : undefined;
            inputs["discoveryOptionsSelectors"] = state ? state.discoveryOptionsSelectors : undefined;
            inputs["endTime"] = state ? state.endTime : undefined;
            inputs["eventOverlays"] = state ? state.eventOverlays : undefined;
            inputs["filters"] = state ? state.filters : undefined;
            inputs["grids"] = state ? state.grids : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["selectedEventOverlays"] = state ? state.selectedEventOverlays : undefined;
            inputs["startTime"] = state ? state.startTime : undefined;
            inputs["timeRange"] = state ? state.timeRange : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as DashboardArgs | undefined;
            if ((!args || args.dashboardGroup === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'dashboardGroup'");
            }
            inputs["authorizedWriterTeams"] = args ? args.authorizedWriterTeams : undefined;
            inputs["authorizedWriterUsers"] = args ? args.authorizedWriterUsers : undefined;
            inputs["charts"] = args ? args.charts : undefined;
            inputs["chartsResolution"] = args ? args.chartsResolution : undefined;
            inputs["columns"] = args ? args.columns : undefined;
            inputs["dashboardGroup"] = args ? args.dashboardGroup : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["discoveryOptionsQuery"] = args ? args.discoveryOptionsQuery : undefined;
            inputs["discoveryOptionsSelectors"] = args ? args.discoveryOptionsSelectors : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["eventOverlays"] = args ? args.eventOverlays : undefined;
            inputs["filters"] = args ? args.filters : undefined;
            inputs["grids"] = args ? args.grids : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["selectedEventOverlays"] = args ? args.selectedEventOverlays : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["timeRange"] = args ? args.timeRange : undefined;
            inputs["variables"] = args ? args.variables : undefined;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Dashboard.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dashboard resources.
 */
export interface DashboardState {
    /**
     * Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterTeams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Chart ID and layout information for the charts in the dashboard.
     */
    readonly charts?: pulumi.Input<pulumi.Input<inputs.DashboardChart>[]>;
    /**
     * Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
     */
    readonly chartsResolution?: pulumi.Input<string>;
    /**
     * Column number for the layout.
     */
    readonly columns?: pulumi.Input<pulumi.Input<inputs.DashboardColumn>[]>;
    /**
     * The ID of the dashboard group that contains the dashboard.
     */
    readonly dashboardGroup?: pulumi.Input<string>;
    /**
     * Variable description.
     */
    readonly description?: pulumi.Input<string>;
    readonly discoveryOptionsQuery?: pulumi.Input<string>;
    readonly discoveryOptionsSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Seconds since epoch. Used for visualization.
     */
    readonly endTime?: pulumi.Input<number>;
    /**
     * Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selectedEventOverlay` property instead.
     */
    readonly eventOverlays?: pulumi.Input<pulumi.Input<inputs.DashboardEventOverlay>[]>;
    /**
     * Filter to apply to the charts when displaying the dashboard.
     */
    readonly filters?: pulumi.Input<pulumi.Input<inputs.DashboardFilter>[]>;
    /**
     * Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
     */
    readonly grids?: pulumi.Input<pulumi.Input<inputs.DashboardGrid>[]>;
    /**
     * Name of the dashboard.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
     */
    readonly selectedEventOverlays?: pulumi.Input<pulumi.Input<inputs.DashboardSelectedEventOverlay>[]>;
    /**
     * Seconds since epoch. Used for visualization.
     */
    readonly startTime?: pulumi.Input<number>;
    /**
     * The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
     */
    readonly timeRange?: pulumi.Input<string>;
    /**
     * The URL of the dashboard.
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Dashboard variable to apply to each chart in the dashboard.
     */
    readonly variables?: pulumi.Input<pulumi.Input<inputs.DashboardVariable>[]>;
}

/**
 * The set of arguments for constructing a Dashboard resource.
 */
export interface DashboardArgs {
    /**
     * Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterTeams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Chart ID and layout information for the charts in the dashboard.
     */
    readonly charts?: pulumi.Input<pulumi.Input<inputs.DashboardChart>[]>;
    /**
     * Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
     */
    readonly chartsResolution?: pulumi.Input<string>;
    /**
     * Column number for the layout.
     */
    readonly columns?: pulumi.Input<pulumi.Input<inputs.DashboardColumn>[]>;
    /**
     * The ID of the dashboard group that contains the dashboard.
     */
    readonly dashboardGroup: pulumi.Input<string>;
    /**
     * Variable description.
     */
    readonly description?: pulumi.Input<string>;
    readonly discoveryOptionsQuery?: pulumi.Input<string>;
    readonly discoveryOptionsSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Seconds since epoch. Used for visualization.
     */
    readonly endTime?: pulumi.Input<number>;
    /**
     * Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selectedEventOverlay` property instead.
     */
    readonly eventOverlays?: pulumi.Input<pulumi.Input<inputs.DashboardEventOverlay>[]>;
    /**
     * Filter to apply to the charts when displaying the dashboard.
     */
    readonly filters?: pulumi.Input<pulumi.Input<inputs.DashboardFilter>[]>;
    /**
     * Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
     */
    readonly grids?: pulumi.Input<pulumi.Input<inputs.DashboardGrid>[]>;
    /**
     * Name of the dashboard.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
     */
    readonly selectedEventOverlays?: pulumi.Input<pulumi.Input<inputs.DashboardSelectedEventOverlay>[]>;
    /**
     * Seconds since epoch. Used for visualization.
     */
    readonly startTime?: pulumi.Input<number>;
    /**
     * The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
     */
    readonly timeRange?: pulumi.Input<string>;
    /**
     * Dashboard variable to apply to each chart in the dashboard.
     */
    readonly variables?: pulumi.Input<pulumi.Input<inputs.DashboardVariable>[]>;
}
