// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This special type of chart displays a data table. This table can be grouped by a dimension.
 *
 * ## Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * // signalfx_list_chart.Logs-Exec_0:
 * const table0 = new signalfx.TableChart("table0", {
 *     description: "beep",
 *     disableSampling: false,
 *     groupBies: ["ClusterName"],
 *     maxDelay: 0,
 *     programText: "A = data('cpu.usage.total').publish(label='CPU Total')",
 * });
 * ```
 *
 * ## Arguments
 *
 * The following arguments are supported in the resource block:
 *
 * * `name` - (Required) Name of the table chart.
 * * `programText` - (Required) The SignalFlow for your Data Table Chart
 * * `description` - (Optional) Description of the table chart.
 * * `groupBy` - (Optional) Dimension to group by
 *
 * ## Attributes
 *
 * In a addition to all arguments above, the following attributes are exported:
 *
 * * `id` - The ID of the chart.
 * * `url` - The URL of the chart.
 */
export class TableChart extends pulumi.CustomResource {
    /**
     * Get an existing TableChart resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableChartState, opts?: pulumi.CustomResourceOptions): TableChart {
        return new TableChart(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/tableChart:TableChart';

    /**
     * Returns true if the given object is an instance of TableChart.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TableChart {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TableChart.__pulumiType;
    }

    /**
     * Description of the chart (Optional)
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * (false by default) If false, samples a subset of the output MTS, which improves UI performance
     */
    public readonly disableSampling!: pulumi.Output<boolean | undefined>;
    /**
     * Properties to group by in the Table (in nesting order)
     */
    public readonly groupBies!: pulumi.Output<string[] | undefined>;
    /**
     * (false by default) Whether to show the timestamp in the chart
     */
    public readonly hideTimestamp!: pulumi.Output<boolean | undefined>;
    /**
     * How long (in seconds) to wait for late datapoints
     */
    public readonly maxDelay!: pulumi.Output<number | undefined>;
    /**
     * The minimum resolution (in seconds) to use for computing the underlying program
     */
    public readonly minimumResolution!: pulumi.Output<number | undefined>;
    /**
     * Name of the chart
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
     */
    public readonly programText!: pulumi.Output<string>;
    /**
     * How often (in seconds) to refresh the values of the Table
     */
    public readonly refreshInterval!: pulumi.Output<number | undefined>;
    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     */
    public readonly timezone!: pulumi.Output<string | undefined>;
    /**
     * (Metric by default) Must be "Metric" or "Binary"
     */
    public readonly unitPrefix!: pulumi.Output<string | undefined>;
    /**
     * URL of the chart
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Plot-level customization options, associated with a publish statement
     */
    public readonly vizOptions!: pulumi.Output<outputs.TableChartVizOption[] | undefined>;

    /**
     * Create a TableChart resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableChartArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableChartArgs | TableChartState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TableChartState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disableSampling"] = state ? state.disableSampling : undefined;
            resourceInputs["groupBies"] = state ? state.groupBies : undefined;
            resourceInputs["hideTimestamp"] = state ? state.hideTimestamp : undefined;
            resourceInputs["maxDelay"] = state ? state.maxDelay : undefined;
            resourceInputs["minimumResolution"] = state ? state.minimumResolution : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["programText"] = state ? state.programText : undefined;
            resourceInputs["refreshInterval"] = state ? state.refreshInterval : undefined;
            resourceInputs["timezone"] = state ? state.timezone : undefined;
            resourceInputs["unitPrefix"] = state ? state.unitPrefix : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["vizOptions"] = state ? state.vizOptions : undefined;
        } else {
            const args = argsOrState as TableChartArgs | undefined;
            if ((!args || args.programText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'programText'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableSampling"] = args ? args.disableSampling : undefined;
            resourceInputs["groupBies"] = args ? args.groupBies : undefined;
            resourceInputs["hideTimestamp"] = args ? args.hideTimestamp : undefined;
            resourceInputs["maxDelay"] = args ? args.maxDelay : undefined;
            resourceInputs["minimumResolution"] = args ? args.minimumResolution : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["programText"] = args ? args.programText : undefined;
            resourceInputs["refreshInterval"] = args ? args.refreshInterval : undefined;
            resourceInputs["timezone"] = args ? args.timezone : undefined;
            resourceInputs["unitPrefix"] = args ? args.unitPrefix : undefined;
            resourceInputs["vizOptions"] = args ? args.vizOptions : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TableChart.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TableChart resources.
 */
export interface TableChartState {
    /**
     * Description of the chart (Optional)
     */
    description?: pulumi.Input<string>;
    /**
     * (false by default) If false, samples a subset of the output MTS, which improves UI performance
     */
    disableSampling?: pulumi.Input<boolean>;
    /**
     * Properties to group by in the Table (in nesting order)
     */
    groupBies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (false by default) Whether to show the timestamp in the chart
     */
    hideTimestamp?: pulumi.Input<boolean>;
    /**
     * How long (in seconds) to wait for late datapoints
     */
    maxDelay?: pulumi.Input<number>;
    /**
     * The minimum resolution (in seconds) to use for computing the underlying program
     */
    minimumResolution?: pulumi.Input<number>;
    /**
     * Name of the chart
     */
    name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
     */
    programText?: pulumi.Input<string>;
    /**
     * How often (in seconds) to refresh the values of the Table
     */
    refreshInterval?: pulumi.Input<number>;
    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     */
    timezone?: pulumi.Input<string>;
    /**
     * (Metric by default) Must be "Metric" or "Binary"
     */
    unitPrefix?: pulumi.Input<string>;
    /**
     * URL of the chart
     */
    url?: pulumi.Input<string>;
    /**
     * Plot-level customization options, associated with a publish statement
     */
    vizOptions?: pulumi.Input<pulumi.Input<inputs.TableChartVizOption>[]>;
}

/**
 * The set of arguments for constructing a TableChart resource.
 */
export interface TableChartArgs {
    /**
     * Description of the chart (Optional)
     */
    description?: pulumi.Input<string>;
    /**
     * (false by default) If false, samples a subset of the output MTS, which improves UI performance
     */
    disableSampling?: pulumi.Input<boolean>;
    /**
     * Properties to group by in the Table (in nesting order)
     */
    groupBies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (false by default) Whether to show the timestamp in the chart
     */
    hideTimestamp?: pulumi.Input<boolean>;
    /**
     * How long (in seconds) to wait for late datapoints
     */
    maxDelay?: pulumi.Input<number>;
    /**
     * The minimum resolution (in seconds) to use for computing the underlying program
     */
    minimumResolution?: pulumi.Input<number>;
    /**
     * Name of the chart
     */
    name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
     */
    programText: pulumi.Input<string>;
    /**
     * How often (in seconds) to refresh the values of the Table
     */
    refreshInterval?: pulumi.Input<number>;
    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     */
    timezone?: pulumi.Input<string>;
    /**
     * (Metric by default) Must be "Metric" or "Binary"
     */
    unitPrefix?: pulumi.Input<string>;
    /**
     * Plot-level customization options, associated with a publish statement
     */
    vizOptions?: pulumi.Input<pulumi.Input<inputs.TableChartVizOption>[]>;
}
