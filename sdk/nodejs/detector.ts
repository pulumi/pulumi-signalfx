// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides a SignalFx detector resource. This can be used to create and manage detectors.
 *
 * > **NOTE** If you're interested in using SignalFx detector features such as Historical Anomaly, Resource Running Out, or others then consider building them in the UI first then using the "Show SignalFlow" feature to extract the value for `programText`. You may also consult the [documentation for detector functions in signalflow-library](https://github.com/signalfx/signalflow-library/tree/master/library/signalfx/detectors).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const config = new pulumi.Config();
 * const clusters = config.getObject("clusters") || [
 *     "clusterA",
 *     "clusterB",
 * ];
 * const applicationDelay: signalfx.Detector[];
 * for (const range = {value: 0}; range.value < clusters.length; range.value++) {
 *     applicationDelay.push(new signalfx.Detector(`applicationDelay-${range.value}`, {
 *         description: `your application is slow - ${clusters[range.value]}`,
 *         maxDelay: 30,
 *         tags: [
 *             "app-backend",
 *             "staging",
 *         ],
 *         authorizedWriterTeams: [signalfx_team.mycoolteam.id],
 *         authorizedWriterUsers: ["abc123"],
 *         programText: `signal = data('app.delay', filter('cluster','${clusters[range.value]}'), extrapolation='last_value', maxExtrapolations=5).max()
 * detect(when(signal > 60, '5m')).publish('Processing old messages 5m')
 * detect(when(signal > 60, '30m')).publish('Processing old messages 30m')
 * `,
 *         rules: [
 *             {
 *                 description: "maximum > 60 for 5m",
 *                 severity: "Warning",
 *                 detectLabel: "Processing old messages 5m",
 *                 notifications: ["Email,foo-alerts@bar.com"],
 *             },
 *             {
 *                 description: "maximum > 60 for 30m",
 *                 severity: "Critical",
 *                 detectLabel: "Processing old messages 30m",
 *                 notifications: ["Email,foo-alerts@bar.com"],
 *             },
 *         ],
 *     }));
 * }
 * ```
 * ## Notification Format
 *
 * As SignalFx supports different notification mechanisms a comma-delimited string is used to provide inputs. If you'd like to specify multiple notifications, then each should be a member in the list, like so:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * This will likely be changed in a future iteration of the provider. See [SignalFx Docs](https://developers.signalfx.com/detectors_reference.html#operation/Create%20Single%20Detector) for more information. For now, here are some example of how to configure each notification type:
 *
 * ### Email
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ### Jira
 *
 * Note that the `credentialId` is the SignalFx-provided ID shown after setting up your Jira integration. (See also `signalfx.jira.Integration`.)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ### Opsgenie
 *
 * Note that the `credentialId` is the SignalFx-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ### PagerDuty
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ### Slack
 *
 * Exclude the `#` on the channel name!
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ### Team
 *
 * Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ### TeamEmail
 *
 * Sends an email to every member of a team.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ### VictorOps
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ### Webhook
 *
 * > **NOTE** You need to include all the commas even if you only use a credential id below.
 *
 * You can either configure a Webhook to use an existing integration's credential id:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * or configure one inline:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ## Import
 *
 * Detectors can be imported using their string ID (recoverable from URL`/#/detector/v2/abc123/edit`, e.g.
 *
 * ```sh
 *  $ pulumi import signalfx:index/detector:Detector application_delay abc123
 * ```
 */
export class Detector extends pulumi.CustomResource {
    /**
     * Get an existing Detector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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
     * Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorizedWriterUsers`).
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
     * How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
     */
    public readonly maxDelay!: pulumi.Output<number | undefined>;
    /**
     * How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
     */
    public readonly minDelay!: pulumi.Output<number | undefined>;
    /**
     * Name of the detector.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    public readonly programText!: pulumi.Output<string>;
    /**
     * Set of rules used for alerting.
     */
    public readonly rules!: pulumi.Output<outputs.DetectorRule[]>;
    /**
     * When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
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
     * Tags associated with the detector.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Team IDs to associate the detector to.
     */
    public readonly teams!: pulumi.Output<string[] | undefined>;
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
     */
    public readonly timeRange!: pulumi.Output<number | undefined>;
    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     */
    public readonly timezone!: pulumi.Output<string | undefined>;
    /**
     * The URL of the detector.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    public readonly vizOptions!: pulumi.Output<outputs.DetectorVizOption[] | undefined>;

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
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DetectorState | undefined;
            inputs["authorizedWriterTeams"] = state ? state.authorizedWriterTeams : undefined;
            inputs["authorizedWriterUsers"] = state ? state.authorizedWriterUsers : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disableSampling"] = state ? state.disableSampling : undefined;
            inputs["endTime"] = state ? state.endTime : undefined;
            inputs["maxDelay"] = state ? state.maxDelay : undefined;
            inputs["minDelay"] = state ? state.minDelay : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["programText"] = state ? state.programText : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["showDataMarkers"] = state ? state.showDataMarkers : undefined;
            inputs["showEventLines"] = state ? state.showEventLines : undefined;
            inputs["startTime"] = state ? state.startTime : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["teams"] = state ? state.teams : undefined;
            inputs["timeRange"] = state ? state.timeRange : undefined;
            inputs["timezone"] = state ? state.timezone : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["vizOptions"] = state ? state.vizOptions : undefined;
        } else {
            const args = argsOrState as DetectorArgs | undefined;
            if ((!args || args.programText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'programText'");
            }
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            inputs["authorizedWriterTeams"] = args ? args.authorizedWriterTeams : undefined;
            inputs["authorizedWriterUsers"] = args ? args.authorizedWriterUsers : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disableSampling"] = args ? args.disableSampling : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["maxDelay"] = args ? args.maxDelay : undefined;
            inputs["minDelay"] = args ? args.minDelay : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["programText"] = args ? args.programText : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["showDataMarkers"] = args ? args.showDataMarkers : undefined;
            inputs["showEventLines"] = args ? args.showEventLines : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["teams"] = args ? args.teams : undefined;
            inputs["timeRange"] = args ? args.timeRange : undefined;
            inputs["timezone"] = args ? args.timezone : undefined;
            inputs["vizOptions"] = args ? args.vizOptions : undefined;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Detector.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Detector resources.
 */
export interface DetectorState {
    /**
     * Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorizedWriterUsers`).
     */
    authorizedWriterTeams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
     */
    authorizedWriterUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
     */
    description?: pulumi.Input<string>;
    /**
     * When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
     */
    disableSampling?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    endTime?: pulumi.Input<number>;
    /**
     * How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
     */
    maxDelay?: pulumi.Input<number>;
    /**
     * How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
     */
    minDelay?: pulumi.Input<number>;
    /**
     * Name of the detector.
     */
    name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    programText?: pulumi.Input<string>;
    /**
     * Set of rules used for alerting.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.DetectorRule>[]>;
    /**
     * When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
     */
    showDataMarkers?: pulumi.Input<boolean>;
    /**
     * When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
     */
    showEventLines?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    startTime?: pulumi.Input<number>;
    /**
     * Tags associated with the detector.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Team IDs to associate the detector to.
     */
    teams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
     */
    timeRange?: pulumi.Input<number>;
    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     */
    timezone?: pulumi.Input<string>;
    /**
     * The URL of the detector.
     */
    url?: pulumi.Input<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    vizOptions?: pulumi.Input<pulumi.Input<inputs.DetectorVizOption>[]>;
}

/**
 * The set of arguments for constructing a Detector resource.
 */
export interface DetectorArgs {
    /**
     * Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorizedWriterUsers`).
     */
    authorizedWriterTeams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
     */
    authorizedWriterUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
     */
    description?: pulumi.Input<string>;
    /**
     * When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
     */
    disableSampling?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    endTime?: pulumi.Input<number>;
    /**
     * How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
     */
    maxDelay?: pulumi.Input<number>;
    /**
     * How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
     */
    minDelay?: pulumi.Input<number>;
    /**
     * Name of the detector.
     */
    name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    programText: pulumi.Input<string>;
    /**
     * Set of rules used for alerting.
     */
    rules: pulumi.Input<pulumi.Input<inputs.DetectorRule>[]>;
    /**
     * When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
     */
    showDataMarkers?: pulumi.Input<boolean>;
    /**
     * When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
     */
    showEventLines?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    startTime?: pulumi.Input<number>;
    /**
     * Tags associated with the detector.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Team IDs to associate the detector to.
     */
    teams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
     */
    timeRange?: pulumi.Input<number>;
    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     */
    timezone?: pulumi.Input<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    vizOptions?: pulumi.Input<pulumi.Input<inputs.DetectorVizOption>[]>;
}
