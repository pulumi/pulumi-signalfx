// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This chart type displays an overview of your SLO and can give more specific insights into your SLI performance using different filter and customized time ranges.
 *
 * ## Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const myslochart0 = new signalfx.SloChart("myslochart0", {sloId: "GbOHXbSAEAA"});
 * ```
 */
export class SloChart extends pulumi.CustomResource {
    /**
     * Get an existing SloChart resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SloChartState, opts?: pulumi.CustomResourceOptions): SloChart {
        return new SloChart(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/sloChart:SloChart';

    /**
     * Returns true if the given object is an instance of SloChart.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SloChart {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SloChart.__pulumiType;
    }

    /**
     * ID of SLO object.
     */
    public readonly sloId!: pulumi.Output<string>;
    /**
     * The URL of the chart.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a SloChart resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SloChartArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SloChartArgs | SloChartState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SloChartState | undefined;
            resourceInputs["sloId"] = state ? state.sloId : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as SloChartArgs | undefined;
            if ((!args || args.sloId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sloId'");
            }
            resourceInputs["sloId"] = args ? args.sloId : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "signalfx:slo/chart:Chart" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(SloChart.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SloChart resources.
 */
export interface SloChartState {
    /**
     * ID of SLO object.
     */
    sloId?: pulumi.Input<string>;
    /**
     * The URL of the chart.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SloChart resource.
 */
export interface SloChartArgs {
    /**
     * ID of SLO object.
     */
    sloId: pulumi.Input<string>;
}
