// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to get a list of dimension values matching the provided query.
 *
 * > **NOTE** The maximum number of values for this data source is 1,000. If you need more, reach out to Splunk support.
 */
export function getDimensionValues(args: GetDimensionValuesArgs, opts?: pulumi.InvokeOptions): Promise<GetDimensionValuesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("signalfx:index/getDimensionValues:getDimensionValues", {
        "limit": args.limit,
        "orderBy": args.orderBy,
        "query": args.query,
    }, opts);
}

/**
 * A collection of arguments for invoking getDimensionValues.
 */
export interface GetDimensionValuesArgs {
    limit?: number;
    orderBy?: string;
    query: string;
}

/**
 * A collection of values returned by getDimensionValues.
 */
export interface GetDimensionValuesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly limit?: number;
    readonly orderBy?: string;
    readonly query: string;
    readonly values: string[];
}
/**
 * Use this data source to get a list of dimension values matching the provided query.
 *
 * > **NOTE** The maximum number of values for this data source is 1,000. If you need more, reach out to Splunk support.
 */
export function getDimensionValuesOutput(args: GetDimensionValuesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetDimensionValuesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("signalfx:index/getDimensionValues:getDimensionValues", {
        "limit": args.limit,
        "orderBy": args.orderBy,
        "query": args.query,
    }, opts);
}

/**
 * A collection of arguments for invoking getDimensionValues.
 */
export interface GetDimensionValuesOutputArgs {
    limit?: pulumi.Input<number>;
    orderBy?: pulumi.Input<string>;
    query: pulumi.Input<string>;
}
