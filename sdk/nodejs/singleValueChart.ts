// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This chart type displays a single number in a large font, representing the current value of a single metric on a plot line.
 * 
 * If the time period is in the past, the number represents the value of the metric near the end of the time period.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 * 
 * const mysvchart0 = new signalfx.SingleValueChart("mysvchart0", {
 *     colorBy: "Dimension",
 *     description: "Very cool Single Value Chart",
 *     isTimestampHidden: true,
 *     maxDelay: 2,
 *     maxPrecision: 2,
 *     programText: `myfilters = filter("clusterName", "prod") and filter("role", "search")
 * data("cpu.total.idle", filter=myfilters).publish()
 * `,
 *     refreshInterval: 1,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/single_value_chart.html.markdown.
 */
export class SingleValueChart extends pulumi.CustomResource {
    /**
     * Get an existing SingleValueChart resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SingleValueChartState, opts?: pulumi.CustomResourceOptions): SingleValueChart {
        return new SingleValueChart(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/singleValueChart:SingleValueChart';

    /**
     * Returns true if the given object is an instance of SingleValueChart.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SingleValueChart {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SingleValueChart.__pulumiType;
    }

    /**
     * Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
     */
    public readonly colorBy!: pulumi.Output<string | undefined>;
    /**
     * Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
     */
    public readonly colorScales!: pulumi.Output<outputs.SingleValueChartColorScale[] | undefined>;
    /**
     * Description of the chart.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether to hide the timestamp in the chart. `false` by default.
     */
    public readonly isTimestampHidden!: pulumi.Output<boolean | undefined>;
    /**
     * How long (in seconds) to wait for late datapoints
     */
    public readonly maxDelay!: pulumi.Output<number | undefined>;
    /**
     * The maximum precision to for value displayed.
     */
    public readonly maxPrecision!: pulumi.Output<number | undefined>;
    /**
     * Name of the chart.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    public readonly programText!: pulumi.Output<string>;
    /**
     * How often (in seconds) to refresh the value.
     */
    public readonly refreshInterval!: pulumi.Output<number | undefined>;
    /**
     * The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
     */
    public readonly secondaryVisualization!: pulumi.Output<string | undefined>;
    /**
     * Whether to show a trend line below the current value. `false` by default.
     */
    public readonly showSparkLine!: pulumi.Output<boolean | undefined>;
    /**
     * Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
     */
    public readonly unitPrefix!: pulumi.Output<string | undefined>;
    /**
     * URL of the chart
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    public readonly vizOptions!: pulumi.Output<outputs.SingleValueChartVizOption[] | undefined>;

    /**
     * Create a SingleValueChart resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SingleValueChartArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SingleValueChartArgs | SingleValueChartState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SingleValueChartState | undefined;
            inputs["colorBy"] = state ? state.colorBy : undefined;
            inputs["colorScales"] = state ? state.colorScales : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["isTimestampHidden"] = state ? state.isTimestampHidden : undefined;
            inputs["maxDelay"] = state ? state.maxDelay : undefined;
            inputs["maxPrecision"] = state ? state.maxPrecision : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["programText"] = state ? state.programText : undefined;
            inputs["refreshInterval"] = state ? state.refreshInterval : undefined;
            inputs["secondaryVisualization"] = state ? state.secondaryVisualization : undefined;
            inputs["showSparkLine"] = state ? state.showSparkLine : undefined;
            inputs["unitPrefix"] = state ? state.unitPrefix : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["vizOptions"] = state ? state.vizOptions : undefined;
        } else {
            const args = argsOrState as SingleValueChartArgs | undefined;
            if (!args || args.programText === undefined) {
                throw new Error("Missing required property 'programText'");
            }
            inputs["colorBy"] = args ? args.colorBy : undefined;
            inputs["colorScales"] = args ? args.colorScales : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["isTimestampHidden"] = args ? args.isTimestampHidden : undefined;
            inputs["maxDelay"] = args ? args.maxDelay : undefined;
            inputs["maxPrecision"] = args ? args.maxPrecision : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["programText"] = args ? args.programText : undefined;
            inputs["refreshInterval"] = args ? args.refreshInterval : undefined;
            inputs["secondaryVisualization"] = args ? args.secondaryVisualization : undefined;
            inputs["showSparkLine"] = args ? args.showSparkLine : undefined;
            inputs["unitPrefix"] = args ? args.unitPrefix : undefined;
            inputs["vizOptions"] = args ? args.vizOptions : undefined;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SingleValueChart.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SingleValueChart resources.
 */
export interface SingleValueChartState {
    /**
     * Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
     */
    readonly colorBy?: pulumi.Input<string>;
    /**
     * Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
     */
    readonly colorScales?: pulumi.Input<pulumi.Input<inputs.SingleValueChartColorScale>[]>;
    /**
     * Description of the chart.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether to hide the timestamp in the chart. `false` by default.
     */
    readonly isTimestampHidden?: pulumi.Input<boolean>;
    /**
     * How long (in seconds) to wait for late datapoints
     */
    readonly maxDelay?: pulumi.Input<number>;
    /**
     * The maximum precision to for value displayed.
     */
    readonly maxPrecision?: pulumi.Input<number>;
    /**
     * Name of the chart.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    readonly programText?: pulumi.Input<string>;
    /**
     * How often (in seconds) to refresh the value.
     */
    readonly refreshInterval?: pulumi.Input<number>;
    /**
     * The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
     */
    readonly secondaryVisualization?: pulumi.Input<string>;
    /**
     * Whether to show a trend line below the current value. `false` by default.
     */
    readonly showSparkLine?: pulumi.Input<boolean>;
    /**
     * Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
     */
    readonly unitPrefix?: pulumi.Input<string>;
    /**
     * URL of the chart
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    readonly vizOptions?: pulumi.Input<pulumi.Input<inputs.SingleValueChartVizOption>[]>;
}

/**
 * The set of arguments for constructing a SingleValueChart resource.
 */
export interface SingleValueChartArgs {
    /**
     * Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
     */
    readonly colorBy?: pulumi.Input<string>;
    /**
     * Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
     */
    readonly colorScales?: pulumi.Input<pulumi.Input<inputs.SingleValueChartColorScale>[]>;
    /**
     * Description of the chart.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether to hide the timestamp in the chart. `false` by default.
     */
    readonly isTimestampHidden?: pulumi.Input<boolean>;
    /**
     * How long (in seconds) to wait for late datapoints
     */
    readonly maxDelay?: pulumi.Input<number>;
    /**
     * The maximum precision to for value displayed.
     */
    readonly maxPrecision?: pulumi.Input<number>;
    /**
     * Name of the chart.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    readonly programText: pulumi.Input<string>;
    /**
     * How often (in seconds) to refresh the value.
     */
    readonly refreshInterval?: pulumi.Input<number>;
    /**
     * The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
     */
    readonly secondaryVisualization?: pulumi.Input<string>;
    /**
     * Whether to show a trend line below the current value. `false` by default.
     */
    readonly showSparkLine?: pulumi.Input<boolean>;
    /**
     * Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
     */
    readonly unitPrefix?: pulumi.Input<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    readonly vizOptions?: pulumi.Input<pulumi.Input<inputs.SingleValueChartVizOption>[]>;
}
