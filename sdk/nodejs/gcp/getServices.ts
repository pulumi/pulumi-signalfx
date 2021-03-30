// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of GCP service names.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const gcpServices = signalfx.gcp.getServices({});
 * // Leaves out most of the integration bits, see the docs
 * // for signalfx_gcp_integration for more
 * // …
 * const gcpMyteam = new signalfx.gcp.Integration("gcpMyteam", {services: [gcpServices.then(gcpServices => gcpServices.services)].map(__item => __item?.name)});
 * ```
 */
export function getServices(args?: GetServicesArgs, opts?: pulumi.InvokeOptions): Promise<GetServicesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("signalfx:gcp/getServices:getServices", {
        "services": args.services,
    }, opts);
}

/**
 * A collection of arguments for invoking getServices.
 */
export interface GetServicesArgs {
    readonly services?: inputs.gcp.GetServicesService[];
}

/**
 * A collection of values returned by getServices.
 */
export interface GetServicesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly services?: outputs.gcp.GetServicesService[];
}