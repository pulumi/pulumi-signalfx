// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * You can add logs data to your Observability Cloud dashboards without turning your logs into metrics first.
 *
 * A log timeline chart displays timeline visualization in a dashboard and shows you in detail what is happening and why.
 *
 * ## Example
 *
 * ## Arguments
 *
 * The following arguments are supported in the resource block:
 *
 * * `name` - (Required) Name of the log timeline.
 * * `programText` - (Required) Signalflow program text for the log timeline. More info at https://dev.splunk.com/observability/docs/.
 * * `description` - (Optional) Description of the log timeline.
 * * `timeRange` - (Optional) From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
 * * `startTime` - (Optional) Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
 * * `endTime` - (Optional) Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
 * * `defaultConnection` - (Optional) The connection that the log timeline uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
 *
 * ## Attributes
 *
 * In a addition to all arguments above, the following attributes are exported:
 *
 * * `id` - The ID of the log timeline.
 * * `url` - The URL of the log timeline.
 */
export class Timeline extends pulumi.CustomResource {
    /**
     * Get an existing Timeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TimelineState, opts?: pulumi.CustomResourceOptions): Timeline {
        return new Timeline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:log/timeline:Timeline';

    /**
     * Returns true if the given object is an instance of Timeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Timeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Timeline.__pulumiType;
    }

    /**
     * default connection that the dashboard uses
     */
    public readonly defaultConnection!: pulumi.Output<string | undefined>;
    /**
     * Description of the chart (Optional)
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Seconds since epoch to end the visualization
     */
    public readonly endTime!: pulumi.Output<number | undefined>;
    /**
     * Name of the chart
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
     */
    public readonly programText!: pulumi.Output<string>;
    /**
     * Seconds since epoch to start the visualization
     */
    public readonly startTime!: pulumi.Output<number | undefined>;
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
     */
    public readonly timeRange!: pulumi.Output<number | undefined>;
    /**
     * URL of the chart
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Timeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TimelineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TimelineArgs | TimelineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TimelineState | undefined;
            resourceInputs["defaultConnection"] = state ? state.defaultConnection : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["programText"] = state ? state.programText : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["timeRange"] = state ? state.timeRange : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as TimelineArgs | undefined;
            if ((!args || args.programText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'programText'");
            }
            resourceInputs["defaultConnection"] = args ? args.defaultConnection : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["programText"] = args ? args.programText : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["timeRange"] = args ? args.timeRange : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Timeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Timeline resources.
 */
export interface TimelineState {
    /**
     * default connection that the dashboard uses
     */
    defaultConnection?: pulumi.Input<string>;
    /**
     * Description of the chart (Optional)
     */
    description?: pulumi.Input<string>;
    /**
     * Seconds since epoch to end the visualization
     */
    endTime?: pulumi.Input<number>;
    /**
     * Name of the chart
     */
    name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
     */
    programText?: pulumi.Input<string>;
    /**
     * Seconds since epoch to start the visualization
     */
    startTime?: pulumi.Input<number>;
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
     */
    timeRange?: pulumi.Input<number>;
    /**
     * URL of the chart
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Timeline resource.
 */
export interface TimelineArgs {
    /**
     * default connection that the dashboard uses
     */
    defaultConnection?: pulumi.Input<string>;
    /**
     * Description of the chart (Optional)
     */
    description?: pulumi.Input<string>;
    /**
     * Seconds since epoch to end the visualization
     */
    endTime?: pulumi.Input<number>;
    /**
     * Name of the chart
     */
    name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
     */
    programText: pulumi.Input<string>;
    /**
     * Seconds since epoch to start the visualization
     */
    startTime?: pulumi.Input<number>;
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
     */
    timeRange?: pulumi.Input<number>;
}
