// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * You can add logs data to your Observability Cloud dashboards without turning your logs into metrics first. A log view displays log lines in a table form in a dashboard and shows you in detail what is happening and why.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const myLogView = new signalfx.logs.View("myLogView", {
 *     columns: [
 *         {
 *             name: "severity",
 *         },
 *         {
 *             name: "time",
 *         },
 *         {
 *             name: "amount.currency_code",
 *         },
 *         {
 *             name: "amount.nanos",
 *         },
 *         {
 *             name: "amount.units",
 *         },
 *         {
 *             name: "message",
 *         },
 *     ],
 *     description: "Lorem ipsum dolor sit amet, laudem tibique iracundia at mea. Nam posse dolores ex, nec cu adhuc putent honestatis",
 *     programText: `logs(filter=field('message') == 'Transaction processed' and field('service.name') == 'paymentservice').publish()
 *
 * `,
 *     sortOptions: [{
 *         descending: false,
 *         field: "severity",
 *     }],
 *     timeRange: 900,
 * });
 * ```
 */
export class View extends pulumi.CustomResource {
    /**
     * Get an existing View resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ViewState, opts?: pulumi.CustomResourceOptions): View {
        return new View(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:logs/view:View';

    /**
     * Returns true if the given object is an instance of View.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is View {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === View.__pulumiType;
    }

    /**
     * The column headers to show on the log view.
     */
    public readonly columns!: pulumi.Output<outputs.logs.ViewColumn[] | undefined>;
    /**
     * The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
     */
    public readonly defaultConnection!: pulumi.Output<string | undefined>;
    /**
     * Description of the log view.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    public readonly endTime!: pulumi.Output<number | undefined>;
    /**
     * Name of the log view.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
     */
    public readonly programText!: pulumi.Output<string>;
    /**
     * The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
     */
    public readonly sortOptions!: pulumi.Output<outputs.logs.ViewSortOption[] | undefined>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    public readonly startTime!: pulumi.Output<number | undefined>;
    /**
     * From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
     */
    public readonly timeRange!: pulumi.Output<number | undefined>;
    /**
     * The URL of the log view.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a View resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ViewArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ViewArgs | ViewState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ViewState | undefined;
            resourceInputs["columns"] = state ? state.columns : undefined;
            resourceInputs["defaultConnection"] = state ? state.defaultConnection : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endTime"] = state ? state.endTime : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["programText"] = state ? state.programText : undefined;
            resourceInputs["sortOptions"] = state ? state.sortOptions : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
            resourceInputs["timeRange"] = state ? state.timeRange : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as ViewArgs | undefined;
            if ((!args || args.programText === undefined) && !opts.urn) {
                throw new Error("Missing required property 'programText'");
            }
            resourceInputs["columns"] = args ? args.columns : undefined;
            resourceInputs["defaultConnection"] = args ? args.defaultConnection : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endTime"] = args ? args.endTime : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["programText"] = args ? args.programText : undefined;
            resourceInputs["sortOptions"] = args ? args.sortOptions : undefined;
            resourceInputs["startTime"] = args ? args.startTime : undefined;
            resourceInputs["timeRange"] = args ? args.timeRange : undefined;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(View.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering View resources.
 */
export interface ViewState {
    /**
     * The column headers to show on the log view.
     */
    columns?: pulumi.Input<pulumi.Input<inputs.logs.ViewColumn>[]>;
    /**
     * The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
     */
    defaultConnection?: pulumi.Input<string>;
    /**
     * Description of the log view.
     */
    description?: pulumi.Input<string>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    endTime?: pulumi.Input<number>;
    /**
     * Name of the log view.
     */
    name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
     */
    programText?: pulumi.Input<string>;
    /**
     * The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
     */
    sortOptions?: pulumi.Input<pulumi.Input<inputs.logs.ViewSortOption>[]>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    startTime?: pulumi.Input<number>;
    /**
     * From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
     */
    timeRange?: pulumi.Input<number>;
    /**
     * The URL of the log view.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a View resource.
 */
export interface ViewArgs {
    /**
     * The column headers to show on the log view.
     */
    columns?: pulumi.Input<pulumi.Input<inputs.logs.ViewColumn>[]>;
    /**
     * The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
     */
    defaultConnection?: pulumi.Input<string>;
    /**
     * Description of the log view.
     */
    description?: pulumi.Input<string>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    endTime?: pulumi.Input<number>;
    /**
     * Name of the log view.
     */
    name?: pulumi.Input<string>;
    /**
     * Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
     */
    programText: pulumi.Input<string>;
    /**
     * The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
     */
    sortOptions?: pulumi.Input<pulumi.Input<inputs.logs.ViewSortOption>[]>;
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
     */
    startTime?: pulumi.Input<number>;
    /**
     * From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `startTime` and `endTime`.
     */
    timeRange?: pulumi.Input<number>;
}
