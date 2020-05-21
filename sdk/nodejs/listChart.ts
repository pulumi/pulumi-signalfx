// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This chart type displays current data values in a list format.
 *
 * The name of each value in the chart reflects the name of the plot and any associated dimensions. We recommend you click the Pencil icon and give the plot a meaningful name, as in plot B below. Otherwise, just the raw metric name will be displayed on the chart, as in plot A below.
 *
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const mylistchart0 = new signalfx.ListChart("mylistchart0", {
 *     colorBy: "Metric",
 *     description: "Very cool List Chart",
 *     disableSampling: true,
 *     legendOptionsFields: [
 *         {
 *             enabled: false,
 *             property: "collector",
 *         },
 *         {
 *             enabled: true,
 *             property: "clusterName",
 *         },
 *         {
 *             enabled: true,
 *             property: "role",
 *         },
 *         {
 *             enabled: false,
 *             property: "collector",
 *         },
 *         {
 *             enabled: false,
 *             property: "host",
 *         },
 *     ],
 *     maxDelay: 2,
 *     maxPrecision: 2,
 *     programText: `myfilters = filter("clusterName", "prod") and filter("role", "search")
 * data("cpu.total.idle", filter=myfilters).publish()
 * `,
 *     refreshInterval: 1,
 *     sortBy: "-value",
 * });
 * ```
 */
export class ListChart extends pulumi.CustomResource {
    /**
     * Get an existing ListChart resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListChartState, opts?: pulumi.CustomResourceOptions): ListChart {
        return new ListChart(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/listChart:ListChart';

    /**
     * Returns true if the given object is an instance of ListChart.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ListChart {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ListChart.__pulumiType;
    }

    /**
     * Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
     */
    public readonly colorBy!: pulumi.Output<string | undefined>;
    /**
     * Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
     */
    public readonly colorScales!: pulumi.Output<outputs.ListChartColorScale[] | undefined>;
    /**
     * Description of the chart.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
     */
    public readonly disableSampling!: pulumi.Output<boolean | undefined>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    public readonly endTime!: pulumi.Output<number | undefined>;
    /**
     * List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
     */
    public readonly legendFieldsToHides!: pulumi.Output<string[] | undefined>;
    /**
     * List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
     */
    public readonly legendOptionsFields!: pulumi.Output<outputs.ListChartLegendOptionsField[] | undefined>;
    /**
     * How long (in seconds) to wait for late datapoints.
     */
    public readonly maxDelay!: pulumi.Output<number | undefined>;
    /**
     * Maximum number of digits to display when rounding values up or down.
     */
    public readonly maxPrecision!: pulumi.Output<number | undefined>;
    /**
     * Name of the chart.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    public readonly programText!: pulumi.Output<string>;
    /**
     * How often (in seconds) to refresh the values of the list.
     */
    public readonly refreshInterval!: pulumi.Output<number | undefined>;
    /**
     * The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
     */
    public readonly secondaryVisualization!: pulumi.Output<string | undefined>;
    /**
     * The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
     */
    public readonly sortBy!: pulumi.Output<string | undefined>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    public readonly startTime!: pulumi.Output<number | undefined>;
    /**
     * How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
     */
    public readonly timeRange!: pulumi.Output<number | undefined>;
    /**
     * Must be `"Metric"` or `"Binary`". `"Metric"` by default.
     */
    public readonly unitPrefix!: pulumi.Output<string | undefined>;
    /**
     * URL of the chart
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    public readonly vizOptions!: pulumi.Output<outputs.ListChartVizOption[] | undefined>;

    /**
     * Create a ListChart resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListChartArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListChartArgs | ListChartState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ListChartState | undefined;
            inputs["colorBy"] = state ? state.colorBy : undefined;
            inputs["colorScales"] = state ? state.colorScales : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disableSampling"] = state ? state.disableSampling : undefined;
            inputs["endTime"] = state ? state.endTime : undefined;
            inputs["legendFieldsToHides"] = state ? state.legendFieldsToHides : undefined;
            inputs["legendOptionsFields"] = state ? state.legendOptionsFields : undefined;
            inputs["maxDelay"] = state ? state.maxDelay : undefined;
            inputs["maxPrecision"] = state ? state.maxPrecision : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["programText"] = state ? state.programText : undefined;
            inputs["refreshInterval"] = state ? state.refreshInterval : undefined;
            inputs["secondaryVisualization"] = state ? state.secondaryVisualization : undefined;
            inputs["sortBy"] = state ? state.sortBy : undefined;
            inputs["startTime"] = state ? state.startTime : undefined;
            inputs["timeRange"] = state ? state.timeRange : undefined;
            inputs["unitPrefix"] = state ? state.unitPrefix : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["vizOptions"] = state ? state.vizOptions : undefined;
        } else {
            const args = argsOrState as ListChartArgs | undefined;
            if (!args || args.programText === undefined) {
                throw new Error("Missing required property 'programText'");
            }
            inputs["colorBy"] = args ? args.colorBy : undefined;
            inputs["colorScales"] = args ? args.colorScales : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disableSampling"] = args ? args.disableSampling : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["legendFieldsToHides"] = args ? args.legendFieldsToHides : undefined;
            inputs["legendOptionsFields"] = args ? args.legendOptionsFields : undefined;
            inputs["maxDelay"] = args ? args.maxDelay : undefined;
            inputs["maxPrecision"] = args ? args.maxPrecision : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["programText"] = args ? args.programText : undefined;
            inputs["refreshInterval"] = args ? args.refreshInterval : undefined;
            inputs["secondaryVisualization"] = args ? args.secondaryVisualization : undefined;
            inputs["sortBy"] = args ? args.sortBy : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["timeRange"] = args ? args.timeRange : undefined;
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
        super(ListChart.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ListChart resources.
 */
export interface ListChartState {
    /**
     * Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
     */
    readonly colorBy?: pulumi.Input<string>;
    /**
     * Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
     */
    readonly colorScales?: pulumi.Input<pulumi.Input<inputs.ListChartColorScale>[]>;
    /**
     * Description of the chart.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
     */
    readonly disableSampling?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly endTime?: pulumi.Input<number>;
    /**
     * List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
     * @deprecated Please use legend_options_fields
     */
    readonly legendFieldsToHides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
     */
    readonly legendOptionsFields?: pulumi.Input<pulumi.Input<inputs.ListChartLegendOptionsField>[]>;
    /**
     * How long (in seconds) to wait for late datapoints.
     */
    readonly maxDelay?: pulumi.Input<number>;
    /**
     * Maximum number of digits to display when rounding values up or down.
     */
    readonly maxPrecision?: pulumi.Input<number>;
    /**
     * Name of the chart.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    readonly programText?: pulumi.Input<string>;
    /**
     * How often (in seconds) to refresh the values of the list.
     */
    readonly refreshInterval?: pulumi.Input<number>;
    /**
     * The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
     */
    readonly secondaryVisualization?: pulumi.Input<string>;
    /**
     * The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
     */
    readonly sortBy?: pulumi.Input<string>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly startTime?: pulumi.Input<number>;
    /**
     * How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
     */
    readonly timeRange?: pulumi.Input<number>;
    /**
     * Must be `"Metric"` or `"Binary`". `"Metric"` by default.
     */
    readonly unitPrefix?: pulumi.Input<string>;
    /**
     * URL of the chart
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    readonly vizOptions?: pulumi.Input<pulumi.Input<inputs.ListChartVizOption>[]>;
}

/**
 * The set of arguments for constructing a ListChart resource.
 */
export interface ListChartArgs {
    /**
     * Must be one of `"Scale"`, `"Dimension"` or `"Metric"`. `"Dimension"` by default.
     */
    readonly colorBy?: pulumi.Input<string>;
    /**
     * Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
     */
    readonly colorScales?: pulumi.Input<pulumi.Input<inputs.ListChartColorScale>[]>;
    /**
     * Description of the chart.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
     */
    readonly disableSampling?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly endTime?: pulumi.Input<number>;
    /**
     * List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
     * @deprecated Please use legend_options_fields
     */
    readonly legendFieldsToHides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
     */
    readonly legendOptionsFields?: pulumi.Input<pulumi.Input<inputs.ListChartLegendOptionsField>[]>;
    /**
     * How long (in seconds) to wait for late datapoints.
     */
    readonly maxDelay?: pulumi.Input<number>;
    /**
     * Maximum number of digits to display when rounding values up or down.
     */
    readonly maxPrecision?: pulumi.Input<number>;
    /**
     * Name of the chart.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    readonly programText: pulumi.Input<string>;
    /**
     * How often (in seconds) to refresh the values of the list.
     */
    readonly refreshInterval?: pulumi.Input<number>;
    /**
     * The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`Sparkline`).
     */
    readonly secondaryVisualization?: pulumi.Input<string>;
    /**
     * The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `"value"` for Value, `"sf_originatingMetric"` for Metric, and `"sfMetric"` for plot.
     */
    readonly sortBy?: pulumi.Input<string>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly startTime?: pulumi.Input<number>;
    /**
     * How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
     */
    readonly timeRange?: pulumi.Input<number>;
    /**
     * Must be `"Metric"` or `"Binary`". `"Metric"` by default.
     */
    readonly unitPrefix?: pulumi.Input<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    readonly vizOptions?: pulumi.Input<pulumi.Input<inputs.ListChartVizOption>[]>;
}
