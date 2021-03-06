// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides a SignalFx time chart resource. This can be used to create and manage the different types of time charts.
 *
 * Time charts display data points over a period of time.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const mychart0 = new signalfx.TimeChart("mychart0", {
 *     axisLeft: {
 *         label: "CPU Total Idle",
 *         lowWatermark: 1000,
 *     },
 *     legendOptionsFields: [
 *         {
 *             enabled: false,
 *             property: "collector",
 *         },
 *         {
 *             enabled: false,
 *             property: "hostname",
 *         },
 *     ],
 *     plotType: "LineChart",
 *     programText: "data(\"cpu.total.idle\").publish(label=\"CPU Idle\")\n",
 *     showDataMarkers: true,
 *     timeRange: 3600,
 *     vizOptions: [{
 *         axis: "left",
 *         color: "orange",
 *         label: "CPU Idle",
 *     }],
 * });
 * ```
 */
export class TimeChart extends pulumi.CustomResource {
    /**
     * Get an existing TimeChart resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TimeChartState, opts?: pulumi.CustomResourceOptions): TimeChart {
        return new TimeChart(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/timeChart:TimeChart';

    /**
     * Returns true if the given object is an instance of TimeChart.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TimeChart {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TimeChart.__pulumiType;
    }

    /**
     * Force the chart to display zero on the y-axes, even if none of the data is near zero.
     */
    public readonly axesIncludeZero!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
     */
    public readonly axesPrecision!: pulumi.Output<number | undefined>;
    /**
     * Set of axis options.
     */
    public readonly axisLeft!: pulumi.Output<outputs.TimeChartAxisLeft | undefined>;
    /**
     * Set of axis options.
     */
    public readonly axisRight!: pulumi.Output<outputs.TimeChartAxisRight | undefined>;
    /**
     * Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
     */
    public readonly colorBy!: pulumi.Output<string | undefined>;
    /**
     * Description of the chart.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
     */
    public readonly disableSampling!: pulumi.Output<boolean | undefined>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    public readonly endTime!: pulumi.Output<number | undefined>;
    /**
     * Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
     */
    public readonly eventOptions!: pulumi.Output<outputs.TimeChartEventOption[] | undefined>;
    /**
     * Only used when `plotType` is `"Histogram"`. Histogram specific options.
     */
    public readonly histogramOptions!: pulumi.Output<outputs.TimeChartHistogramOption[] | undefined>;
    /**
     * List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
     *
     * @deprecated Please use legend_options_fields
     */
    public readonly legendFieldsToHides!: pulumi.Output<string[] | undefined>;
    /**
     * List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
     */
    public readonly legendOptionsFields!: pulumi.Output<outputs.TimeChartLegendOptionsField[] | undefined>;
    /**
     * How long (in seconds) to wait for late datapoints.
     */
    public readonly maxDelay!: pulumi.Output<number | undefined>;
    /**
     * The minimum resolution (in seconds) to use for computing the underlying program.
     */
    public readonly minimumResolution!: pulumi.Output<number | undefined>;
    /**
     * Name of the chart.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
     */
    public readonly onChartLegendDimension!: pulumi.Output<string | undefined>;
    /**
     * The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
     */
    public readonly plotType!: pulumi.Output<string | undefined>;
    /**
     * Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    public readonly programText!: pulumi.Output<string>;
    /**
     * Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
     */
    public readonly showDataMarkers!: pulumi.Output<boolean | undefined>;
    /**
     * Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
     */
    public readonly showEventLines!: pulumi.Output<boolean | undefined>;
    /**
     * Whether area and bar charts in the visualization should be stacked. `false` by default.
     */
    public readonly stacked!: pulumi.Output<boolean | undefined>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    public readonly startTime!: pulumi.Output<number | undefined>;
    /**
     * Tags associated with the chart
     *
     * @deprecated signalfx_time_chart.tags is being removed in the next release
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
     */
    public readonly timeRange!: pulumi.Output<number | undefined>;
    /**
     * Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
     */
    public readonly timezone!: pulumi.Output<string | undefined>;
    /**
     * Must be `"Metric"` or `"Binary`". `"Metric"` by default.
     */
    public readonly unitPrefix!: pulumi.Output<string | undefined>;
    /**
     * The URL of the chart.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    public readonly vizOptions!: pulumi.Output<outputs.TimeChartVizOption[] | undefined>;

    /**
     * Create a TimeChart resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TimeChartArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TimeChartArgs | TimeChartState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TimeChartState | undefined;
            inputs["axesIncludeZero"] = state ? state.axesIncludeZero : undefined;
            inputs["axesPrecision"] = state ? state.axesPrecision : undefined;
            inputs["axisLeft"] = state ? state.axisLeft : undefined;
            inputs["axisRight"] = state ? state.axisRight : undefined;
            inputs["colorBy"] = state ? state.colorBy : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disableSampling"] = state ? state.disableSampling : undefined;
            inputs["endTime"] = state ? state.endTime : undefined;
            inputs["eventOptions"] = state ? state.eventOptions : undefined;
            inputs["histogramOptions"] = state ? state.histogramOptions : undefined;
            inputs["legendFieldsToHides"] = state ? state.legendFieldsToHides : undefined;
            inputs["legendOptionsFields"] = state ? state.legendOptionsFields : undefined;
            inputs["maxDelay"] = state ? state.maxDelay : undefined;
            inputs["minimumResolution"] = state ? state.minimumResolution : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["onChartLegendDimension"] = state ? state.onChartLegendDimension : undefined;
            inputs["plotType"] = state ? state.plotType : undefined;
            inputs["programText"] = state ? state.programText : undefined;
            inputs["showDataMarkers"] = state ? state.showDataMarkers : undefined;
            inputs["showEventLines"] = state ? state.showEventLines : undefined;
            inputs["stacked"] = state ? state.stacked : undefined;
            inputs["startTime"] = state ? state.startTime : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["timeRange"] = state ? state.timeRange : undefined;
            inputs["timezone"] = state ? state.timezone : undefined;
            inputs["unitPrefix"] = state ? state.unitPrefix : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["vizOptions"] = state ? state.vizOptions : undefined;
        } else {
            const args = argsOrState as TimeChartArgs | undefined;
            if ((!args || args.programText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'programText'");
            }
            inputs["axesIncludeZero"] = args ? args.axesIncludeZero : undefined;
            inputs["axesPrecision"] = args ? args.axesPrecision : undefined;
            inputs["axisLeft"] = args ? args.axisLeft : undefined;
            inputs["axisRight"] = args ? args.axisRight : undefined;
            inputs["colorBy"] = args ? args.colorBy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disableSampling"] = args ? args.disableSampling : undefined;
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["eventOptions"] = args ? args.eventOptions : undefined;
            inputs["histogramOptions"] = args ? args.histogramOptions : undefined;
            inputs["legendFieldsToHides"] = args ? args.legendFieldsToHides : undefined;
            inputs["legendOptionsFields"] = args ? args.legendOptionsFields : undefined;
            inputs["maxDelay"] = args ? args.maxDelay : undefined;
            inputs["minimumResolution"] = args ? args.minimumResolution : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["onChartLegendDimension"] = args ? args.onChartLegendDimension : undefined;
            inputs["plotType"] = args ? args.plotType : undefined;
            inputs["programText"] = args ? args.programText : undefined;
            inputs["showDataMarkers"] = args ? args.showDataMarkers : undefined;
            inputs["showEventLines"] = args ? args.showEventLines : undefined;
            inputs["stacked"] = args ? args.stacked : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["timeRange"] = args ? args.timeRange : undefined;
            inputs["timezone"] = args ? args.timezone : undefined;
            inputs["unitPrefix"] = args ? args.unitPrefix : undefined;
            inputs["vizOptions"] = args ? args.vizOptions : undefined;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TimeChart.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TimeChart resources.
 */
export interface TimeChartState {
    /**
     * Force the chart to display zero on the y-axes, even if none of the data is near zero.
     */
    readonly axesIncludeZero?: pulumi.Input<boolean>;
    /**
     * Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
     */
    readonly axesPrecision?: pulumi.Input<number>;
    /**
     * Set of axis options.
     */
    readonly axisLeft?: pulumi.Input<inputs.TimeChartAxisLeft>;
    /**
     * Set of axis options.
     */
    readonly axisRight?: pulumi.Input<inputs.TimeChartAxisRight>;
    /**
     * Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
     */
    readonly colorBy?: pulumi.Input<string>;
    /**
     * Description of the chart.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
     */
    readonly disableSampling?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly endTime?: pulumi.Input<number>;
    /**
     * Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
     */
    readonly eventOptions?: pulumi.Input<pulumi.Input<inputs.TimeChartEventOption>[]>;
    /**
     * Only used when `plotType` is `"Histogram"`. Histogram specific options.
     */
    readonly histogramOptions?: pulumi.Input<pulumi.Input<inputs.TimeChartHistogramOption>[]>;
    /**
     * List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
     *
     * @deprecated Please use legend_options_fields
     */
    readonly legendFieldsToHides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
     */
    readonly legendOptionsFields?: pulumi.Input<pulumi.Input<inputs.TimeChartLegendOptionsField>[]>;
    /**
     * How long (in seconds) to wait for late datapoints.
     */
    readonly maxDelay?: pulumi.Input<number>;
    /**
     * The minimum resolution (in seconds) to use for computing the underlying program.
     */
    readonly minimumResolution?: pulumi.Input<number>;
    /**
     * Name of the chart.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
     */
    readonly onChartLegendDimension?: pulumi.Input<string>;
    /**
     * The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
     */
    readonly plotType?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    readonly programText?: pulumi.Input<string>;
    /**
     * Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
     */
    readonly showDataMarkers?: pulumi.Input<boolean>;
    /**
     * Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
     */
    readonly showEventLines?: pulumi.Input<boolean>;
    /**
     * Whether area and bar charts in the visualization should be stacked. `false` by default.
     */
    readonly stacked?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly startTime?: pulumi.Input<number>;
    /**
     * Tags associated with the chart
     *
     * @deprecated signalfx_time_chart.tags is being removed in the next release
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
     */
    readonly timeRange?: pulumi.Input<number>;
    /**
     * Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
     */
    readonly timezone?: pulumi.Input<string>;
    /**
     * Must be `"Metric"` or `"Binary`". `"Metric"` by default.
     */
    readonly unitPrefix?: pulumi.Input<string>;
    /**
     * The URL of the chart.
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    readonly vizOptions?: pulumi.Input<pulumi.Input<inputs.TimeChartVizOption>[]>;
}

/**
 * The set of arguments for constructing a TimeChart resource.
 */
export interface TimeChartArgs {
    /**
     * Force the chart to display zero on the y-axes, even if none of the data is near zero.
     */
    readonly axesIncludeZero?: pulumi.Input<boolean>;
    /**
     * Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
     */
    readonly axesPrecision?: pulumi.Input<number>;
    /**
     * Set of axis options.
     */
    readonly axisLeft?: pulumi.Input<inputs.TimeChartAxisLeft>;
    /**
     * Set of axis options.
     */
    readonly axisRight?: pulumi.Input<inputs.TimeChartAxisRight>;
    /**
     * Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
     */
    readonly colorBy?: pulumi.Input<string>;
    /**
     * Description of the chart.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
     */
    readonly disableSampling?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly endTime?: pulumi.Input<number>;
    /**
     * Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
     */
    readonly eventOptions?: pulumi.Input<pulumi.Input<inputs.TimeChartEventOption>[]>;
    /**
     * Only used when `plotType` is `"Histogram"`. Histogram specific options.
     */
    readonly histogramOptions?: pulumi.Input<pulumi.Input<inputs.TimeChartHistogramOption>[]>;
    /**
     * List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legendOptionsFields`.
     *
     * @deprecated Please use legend_options_fields
     */
    readonly legendFieldsToHides?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legendFieldsToHide`.
     */
    readonly legendOptionsFields?: pulumi.Input<pulumi.Input<inputs.TimeChartLegendOptionsField>[]>;
    /**
     * How long (in seconds) to wait for late datapoints.
     */
    readonly maxDelay?: pulumi.Input<number>;
    /**
     * The minimum resolution (in seconds) to use for computing the underlying program.
     */
    readonly minimumResolution?: pulumi.Input<number>;
    /**
     * Name of the chart.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plotLabel"` and any dimension.
     */
    readonly onChartLegendDimension?: pulumi.Input<string>;
    /**
     * The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
     */
    readonly plotType?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
     */
    readonly programText: pulumi.Input<string>;
    /**
     * Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
     */
    readonly showDataMarkers?: pulumi.Input<boolean>;
    /**
     * Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
     */
    readonly showEventLines?: pulumi.Input<boolean>;
    /**
     * Whether area and bar charts in the visualization should be stacked. `false` by default.
     */
    readonly stacked?: pulumi.Input<boolean>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    readonly startTime?: pulumi.Input<number>;
    /**
     * Tags associated with the chart
     *
     * @deprecated signalfx_time_chart.tags is being removed in the next release
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `startTime` and `endTime`.
     */
    readonly timeRange?: pulumi.Input<number>;
    /**
     * Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
     */
    readonly timezone?: pulumi.Input<string>;
    /**
     * Must be `"Metric"` or `"Binary`". `"Metric"` by default.
     */
    readonly unitPrefix?: pulumi.Input<string>;
    /**
     * Plot-level customization options, associated with a publish statement.
     */
    readonly vizOptions?: pulumi.Input<pulumi.Input<inputs.TimeChartVizOption>[]>;
}
