// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This special type of chart doesn’t display any metric data. Rather, it lets you place a text note on the dashboard.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const mynote0 = new signalfx.TextChart("mynote0", {
 *     description: "Lorem ipsum dolor sit amet, laudem tibique iracundia at mea. Nam posse dolores ex, nec cu adhuc putent honestatis",
 *     markdown: `    1. First ordered list item
 *     2. Another item
 *       * Unordered sub-list.
 *     1. Actual numbers don't matter, just that it's a number
 *       1. Ordered sub-list
 *     4. And another item.
 *
 *        You can have properly indented paragraphs within list items. Notice the blank line above, and the leading spaces (at least one, but we'll use three here to also align the raw Markdown).
 *
 *        To have a line break without a paragraph, you will need to use two trailing spaces.⋅⋅
 *        Note that this line is separate, but within the same paragraph.⋅⋅
 *        (This is contrary to the typical GFM line break behaviour, where trailing spaces are not required.)
 *
 *     * Unordered list can use asterisks
 *     - Or minuses
 *     + Or pluses
 * `,
 * });
 * ```
 */
export class TextChart extends pulumi.CustomResource {
    /**
     * Get an existing TextChart resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TextChartState, opts?: pulumi.CustomResourceOptions): TextChart {
        return new TextChart(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/textChart:TextChart';

    /**
     * Returns true if the given object is an instance of TextChart.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TextChart {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TextChart.__pulumiType;
    }

    /**
     * Description of the text note.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Markdown text to display.
     */
    public readonly markdown!: pulumi.Output<string>;
    /**
     * Name of the text note.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL of the chart.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a TextChart resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TextChartArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TextChartArgs | TextChartState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TextChartState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["markdown"] = state ? state.markdown : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as TextChartArgs | undefined;
            if ((!args || args.markdown === undefined) && !opts.urn) {
                throw new Error("Missing required property 'markdown'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["markdown"] = args ? args.markdown : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TextChart.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TextChart resources.
 */
export interface TextChartState {
    /**
     * Description of the text note.
     */
    description?: pulumi.Input<string>;
    /**
     * Markdown text to display.
     */
    markdown?: pulumi.Input<string>;
    /**
     * Name of the text note.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL of the chart.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TextChart resource.
 */
export interface TextChartArgs {
    /**
     * Description of the text note.
     */
    description?: pulumi.Input<string>;
    /**
     * Markdown text to display.
     */
    markdown: pulumi.Input<string>;
    /**
     * Name of the text note.
     */
    name?: pulumi.Input<string>;
}
