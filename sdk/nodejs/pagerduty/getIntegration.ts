// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get information on an existing PagerDuty integration.
 *
 * ## Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const pdIntegration = signalfx.pagerduty.getIntegration({
 *     name: "PD-Integration",
 * });
 * ```
 */
export function getIntegration(args: GetIntegrationArgs, opts?: pulumi.InvokeOptions): Promise<GetIntegrationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("signalfx:pagerduty/getIntegration:getIntegration", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIntegration.
 */
export interface GetIntegrationArgs {
    /**
     * Specify the exact name of the desired PagerDuty integration
     */
    name: string;
}

/**
 * A collection of values returned by getIntegration.
 */
export interface GetIntegrationResult {
    /**
     * Whether the integration is enabled.
     */
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the integration.
     */
    readonly name: string;
}
/**
 * Use this data source to get information on an existing PagerDuty integration.
 *
 * ## Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const pdIntegration = signalfx.pagerduty.getIntegration({
 *     name: "PD-Integration",
 * });
 * ```
 */
export function getIntegrationOutput(args: GetIntegrationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIntegrationResult> {
    return pulumi.output(args).apply((a: any) => getIntegration(a, opts))
}

/**
 * A collection of arguments for invoking getIntegration.
 */
export interface GetIntegrationOutputArgs {
    /**
     * Specify the exact name of the desired PagerDuty integration
     */
    name: pulumi.Input<string>;
}
