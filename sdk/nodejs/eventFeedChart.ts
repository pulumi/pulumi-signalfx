// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Displays a listing of events as a widget in a dashboard.
 */
export class EventFeedChart extends pulumi.CustomResource {
    /**
     * Get an existing EventFeedChart resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventFeedChartState, opts?: pulumi.CustomResourceOptions): EventFeedChart {
        return new EventFeedChart(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/eventFeedChart:EventFeedChart';

    /**
     * Returns true if the given object is an instance of EventFeedChart.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventFeedChart {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventFeedChart.__pulumiType;
    }

    /**
     * Description of the text note.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    public readonly endTime!: pulumi.Output<number | undefined>;
    /**
     * Name of the text note.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    public readonly programText!: pulumi.Output<string>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    public readonly startTime!: pulumi.Output<number | undefined>;
    /**
     * From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
     */
    public readonly timeRange!: pulumi.Output<number | undefined>;
    /**
     * The URL of the chart.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a EventFeedChart resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventFeedChartArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventFeedChartArgs | EventFeedChartState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventFeedChartState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["programText"] = state ? state.programText : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["timeRange"] = state ? state.timeRange : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as EventFeedChartArgs | undefined;
            if ((!args || args.programText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'programText'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["programText"] = args ? args.programText : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["timeRange"] = args ? args.timeRange : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventFeedChart.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventFeedChart resources.
 */
export interface EventFeedChartState {
    /**
     * Description of the text note.
     */
    description?: pulumi.Input<string>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    endTime?: pulumi.Input<number>;
    /**
     * Name of the text note.
     */
    name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    programText?: pulumi.Input<string>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    startTime?: pulumi.Input<number>;
    /**
     * From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
     */
    timeRange?: pulumi.Input<number>;
    /**
     * The URL of the chart.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventFeedChart resource.
 */
export interface EventFeedChartArgs {
    /**
     * Description of the text note.
     */
    description?: pulumi.Input<string>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    endTime?: pulumi.Input<number>;
    /**
     * Name of the text note.
     */
    name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    programText: pulumi.Input<string>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    startTime?: pulumi.Input<number>;
    /**
     * From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
     */
    timeRange?: pulumi.Input<number>;
}
