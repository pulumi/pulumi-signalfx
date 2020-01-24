// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/detector.html.markdown.
 */
export class Detector extends pulumi.CustomResource {
    /**
     * Get an existing Detector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DetectorState, opts?: pulumi.CustomResourceOptions): Detector {
        return new Detector(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/detector:Detector';

    /**
     * Returns true if the given object is an instance of Detector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Detector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Detector.__pulumiType;
    }

    /**
     * Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
     */
    public readonly authorizedWriterTeams!: pulumi.Output<string[] | undefined>;
    /**
     * User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
     */
    public readonly authorizedWriterUsers!: pulumi.Output<string[] | undefined>;
    /**
     * Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI. 
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
     */
    public readonly disableSampling!: pulumi.Output<boolean | undefined>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    public readonly endTime!: pulumi.Output<number | undefined>;
    /**
     * How long (in seconds) to wait for late datapoints. See <https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints> for more info. Max value is `900` seconds (15 minutes).
     */
    public readonly maxDelay!: pulumi.Output<number | undefined>;
    /**
     * Name of the detector.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Signalflow program text for the detector. More info at <https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html>.
     */
    public readonly programText!: pulumi.Output<string>;
    /**
     * Set of rules used for alerting.
     */
    public readonly rules!: pulumi.Output<outputs.DetectorRule[]>;
    /**
     * When `true`, markers will be drawn for each datapoint within the visualization. `false` by default.
     */
    public readonly showDataMarkers!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
     */
    public readonly showEventLines!: pulumi.Output<boolean | undefined>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    public readonly startTime!: pulumi.Output<number | undefined>;
    /**
     * Team IDs to associate the detector to.
     */
    public readonly teams!: pulumi.Output<string[] | undefined>;
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`. Defaults to 3600.
     */
    public readonly timeRange!: pulumi.Output<number | undefined>;
    /**
     * URL of the detector
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Detector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DetectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DetectorArgs | DetectorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DetectorState | undefined;
            inputs["authorizedWriterTeams"] = state ? state.authorizedWriterTeams : undefined;
            inputs["authorizedWriterUsers"] = state ? state.authorizedWriterUsers : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disableSampling"] = state ? state.disableSampling : undefined;
            inputs["endTime"] = state ? state.endTime : undefined;
            inputs["maxDelay"] = state ? state.maxDelay : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["programText"] = state ? state.programText : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["showDataMarkers"] = state ? state.showDataMarkers : undefined;
            inputs["showEventLines"] = state ? state.showEventLines : undefined;
            inputs["startTime"] = state ? state.startTime : undefined;
            inputs["teams"] = state ? state.teams : undefined;
            inputs["timeRange"] = state ? state.timeRange : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as DetectorArgs | undefined;
            if (!args || args.programText === undefined) {
                throw new Error("Missing required property 'programText'");
            }
            if (!args || args.rules === undefined) {
                throw new Error("Missing required property 'rules'");
            }
            inputs["authorizedWriterTeams"] = args ? args.authorizedWriterTeams : undefined;
            inputs["authorizedWriterUsers"] = args ? args.authorizedWriterUsers : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disableSampling"] = args ? args.disableSampling : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["maxDelay"] = args ? args.maxDelay : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["programText"] = args ? args.programText : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["showDataMarkers"] = args ? args.showDataMarkers : undefined;
            inputs["showEventLines"] = args ? args.showEventLines : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["teams"] = args ? args.teams : undefined;
            inputs["timeRange"] = args ? args.timeRange : undefined;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Detector.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Detector resources.
 */
export interface DetectorState {
    /**
     * Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterTeams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI. 
     */
    readonly description?: pulumi.Input<string>;
    /**
     * When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
     */
    readonly disableSampling?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly endTime?: pulumi.Input<number>;
    /**
     * How long (in seconds) to wait for late datapoints. See <https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints> for more info. Max value is `900` seconds (15 minutes).
     */
    readonly maxDelay?: pulumi.Input<number>;
    /**
     * Name of the detector.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the detector. More info at <https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html>.
     */
    readonly programText?: pulumi.Input<string>;
    /**
     * Set of rules used for alerting.
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.DetectorRule>[]>;
    /**
     * When `true`, markers will be drawn for each datapoint within the visualization. `false` by default.
     */
    readonly showDataMarkers?: pulumi.Input<boolean>;
    /**
     * When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
     */
    readonly showEventLines?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly startTime?: pulumi.Input<number>;
    /**
     * Team IDs to associate the detector to.
     */
    readonly teams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`. Defaults to 3600.
     */
    readonly timeRange?: pulumi.Input<number>;
    /**
     * URL of the detector
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Detector resource.
 */
export interface DetectorArgs {
    /**
     * Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterTeams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
     */
    readonly authorizedWriterUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI. 
     */
    readonly description?: pulumi.Input<string>;
    /**
     * When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
     */
    readonly disableSampling?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly endTime?: pulumi.Input<number>;
    /**
     * How long (in seconds) to wait for late datapoints. See <https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints> for more info. Max value is `900` seconds (15 minutes).
     */
    readonly maxDelay?: pulumi.Input<number>;
    /**
     * Name of the detector.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the detector. More info at <https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html>.
     */
    readonly programText: pulumi.Input<string>;
    /**
     * Set of rules used for alerting.
     */
    readonly rules: pulumi.Input<pulumi.Input<inputs.DetectorRule>[]>;
    /**
     * When `true`, markers will be drawn for each datapoint within the visualization. `false` by default.
     */
    readonly showDataMarkers?: pulumi.Input<boolean>;
    /**
     * When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
     */
    readonly showEventLines?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly startTime?: pulumi.Input<number>;
    /**
     * Team IDs to associate the detector to.
     */
    readonly teams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`. Defaults to 3600.
     */
    readonly timeRange?: pulumi.Input<number>;
}
