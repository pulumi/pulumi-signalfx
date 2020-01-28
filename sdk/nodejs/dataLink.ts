// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manage SignalFx [Data Links](https://docs.signalfx.com/en/latest/managing/data-links.html).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 * 
 * // A global link to SignalFx dashboard.
 * const myDataLink = new signalfx.DataLink("myDataLink", {
 *     propertyName: "pname",
 *     propertyValue: "pvalue",
 *     targetSignalfxDashboards: [{
 *         dashboardGroupId: signalfx_dashboard_group_mydashboardgroup0.id,
 *         dashboardId: signalfx_dashboard_mydashboard0.id,
 *         isDefault: true,
 *         name: "sfxDash",
 *     }],
 * });
 * // A dashboard-specific link to an external URL
 * const myDataLinkDash = new signalfx.DataLink("myDataLinkDash", {
 *     contextDashboardId: signalfx_dashboard_mydashboard0.id,
 *     propertyName: "pname2",
 *     propertyValue: "pvalue",
 *     targetExternalUrls: [{
 *         isDefault: false,
 *         name: "exUrl",
 *         propertyKeyMapping: {
 *             foo: "bar",
 *         },
 *         timeFormat: "ISO8601",
 *         url: "https://www.example.com",
 *     }],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/data_link.html.markdown.
 */
export class DataLink extends pulumi.CustomResource {
    /**
     * Get an existing DataLink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataLinkState, opts?: pulumi.CustomResourceOptions): DataLink {
        return new DataLink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/dataLink:DataLink';

    /**
     * Returns true if the given object is an instance of DataLink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataLink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataLink.__pulumiType;
    }

    /**
     * If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
     */
    public readonly contextDashboardId!: pulumi.Output<string | undefined>;
    /**
     * Name (key) of the metadata that's the trigger of a data link. If you specify `propertyValue`, you must specify `propertyName`.
     */
    public readonly propertyName!: pulumi.Output<string | undefined>;
    /**
     * Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `propertyName`.
     */
    public readonly propertyValue!: pulumi.Output<string | undefined>;
    /**
     * Link to an external URL
     */
    public readonly targetExternalUrls!: pulumi.Output<outputs.DataLinkTargetExternalUrl[] | undefined>;
    /**
     * Link to a SignalFx dashboard
     */
    public readonly targetSignalfxDashboards!: pulumi.Output<outputs.DataLinkTargetSignalfxDashboard[] | undefined>;
    /**
     * Link to an external URL
     */
    public readonly targetSplunks!: pulumi.Output<outputs.DataLinkTargetSplunk[] | undefined>;

    /**
     * Create a DataLink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DataLinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataLinkArgs | DataLinkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DataLinkState | undefined;
            inputs["contextDashboardId"] = state ? state.contextDashboardId : undefined;
            inputs["propertyName"] = state ? state.propertyName : undefined;
            inputs["propertyValue"] = state ? state.propertyValue : undefined;
            inputs["targetExternalUrls"] = state ? state.targetExternalUrls : undefined;
            inputs["targetSignalfxDashboards"] = state ? state.targetSignalfxDashboards : undefined;
            inputs["targetSplunks"] = state ? state.targetSplunks : undefined;
        } else {
            const args = argsOrState as DataLinkArgs | undefined;
            inputs["contextDashboardId"] = args ? args.contextDashboardId : undefined;
            inputs["propertyName"] = args ? args.propertyName : undefined;
            inputs["propertyValue"] = args ? args.propertyValue : undefined;
            inputs["targetExternalUrls"] = args ? args.targetExternalUrls : undefined;
            inputs["targetSignalfxDashboards"] = args ? args.targetSignalfxDashboards : undefined;
            inputs["targetSplunks"] = args ? args.targetSplunks : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DataLink.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataLink resources.
 */
export interface DataLinkState {
    /**
     * If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
     */
    readonly contextDashboardId?: pulumi.Input<string>;
    /**
     * Name (key) of the metadata that's the trigger of a data link. If you specify `propertyValue`, you must specify `propertyName`.
     */
    readonly propertyName?: pulumi.Input<string>;
    /**
     * Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `propertyName`.
     */
    readonly propertyValue?: pulumi.Input<string>;
    /**
     * Link to an external URL
     */
    readonly targetExternalUrls?: pulumi.Input<pulumi.Input<inputs.DataLinkTargetExternalUrl>[]>;
    /**
     * Link to a SignalFx dashboard
     */
    readonly targetSignalfxDashboards?: pulumi.Input<pulumi.Input<inputs.DataLinkTargetSignalfxDashboard>[]>;
    /**
     * Link to an external URL
     */
    readonly targetSplunks?: pulumi.Input<pulumi.Input<inputs.DataLinkTargetSplunk>[]>;
}

/**
 * The set of arguments for constructing a DataLink resource.
 */
export interface DataLinkArgs {
    /**
     * If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
     */
    readonly contextDashboardId?: pulumi.Input<string>;
    /**
     * Name (key) of the metadata that's the trigger of a data link. If you specify `propertyValue`, you must specify `propertyName`.
     */
    readonly propertyName?: pulumi.Input<string>;
    /**
     * Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `propertyName`.
     */
    readonly propertyValue?: pulumi.Input<string>;
    /**
     * Link to an external URL
     */
    readonly targetExternalUrls?: pulumi.Input<pulumi.Input<inputs.DataLinkTargetExternalUrl>[]>;
    /**
     * Link to a SignalFx dashboard
     */
    readonly targetSignalfxDashboards?: pulumi.Input<pulumi.Input<inputs.DataLinkTargetSignalfxDashboard>[]>;
    /**
     * Link to an external URL
     */
    readonly targetSplunks?: pulumi.Input<pulumi.Input<inputs.DataLinkTargetSplunk>[]>;
}
