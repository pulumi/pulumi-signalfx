// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data sources allows for obtaining a list of dimension values by on query provided.
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
    /**
     * This allows you to define how many dimensions are returned as the values output.
     */
    limit?: number;
    orderBy?: string;
    /**
     * Acepts a query string that allows for defining a key value deintion, wild card matching on values, or where the dimension value exists. Refer to https://dev.splunk.com/observability/reference/api/metrics_metadata/latest#endpoint-retrieve-dimensions-query for more details
     */
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
    /**
     * This allows you to define how many dimensions are returned as the values output.
     */
    readonly limit?: number;
    readonly orderBy?: string;
    /**
     * Acepts a query string that allows for defining a key value deintion, wild card matching on values, or where the dimension value exists. Refer to https://dev.splunk.com/observability/reference/api/metrics_metadata/latest#endpoint-retrieve-dimensions-query for more details
     */
    readonly query: string;
    /**
     * List of all the match dimension values that the provided query, ordered by orderBy field
     */
    readonly values: string[];
}
/**
 * This data sources allows for obtaining a list of dimension values by on query provided.
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
    /**
     * This allows you to define how many dimensions are returned as the values output.
     */
    limit?: pulumi.Input<number>;
    orderBy?: pulumi.Input<string>;
    /**
     * Acepts a query string that allows for defining a key value deintion, wild card matching on values, or where the dimension value exists. Refer to https://dev.splunk.com/observability/reference/api/metrics_metadata/latest#endpoint-retrieve-dimensions-query for more details
     */
    query: pulumi.Input<string>;
}
