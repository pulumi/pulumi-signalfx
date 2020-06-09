// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get a list of dimension values matching the provided query.
 *
 * > **NOTE** This data source only allows 1000 values, as it's kinda nuts to make anything with `forEach` that big in SignalFx. This is negotiable.
 */
export function getDimensionValues(args: GetDimensionValuesArgs, opts?: pulumi.InvokeOptions): Promise<GetDimensionValuesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("signalfx:index/getDimensionValues:getDimensionValues", {
        "query": args.query,
    }, opts);
}

/**
 * A collection of arguments for invoking getDimensionValues.
 */
export interface GetDimensionValuesArgs {
    readonly query: string;
}

/**
 * A collection of values returned by getDimensionValues.
 */
export interface GetDimensionValuesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly query: string;
    readonly values: string[];
}
