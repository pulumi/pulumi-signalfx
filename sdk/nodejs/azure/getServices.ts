// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of Azure service names.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const azureServices = signalfx.azure.getServices({});
 * // Leaves out most of the integration bits, see the docs
 * // for signalfx.azure.Integration for more
 * const azureMyteam = new signalfx.azure.Integration("azureMyteam", {services: [azureServices.then(azureServices => azureServices.services)].map(__item => __item?.name)});
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
    return pulumi.runtime.invoke("signalfx:azure/getServices:getServices", {
        "services": args.services,
    }, opts);
}

/**
 * A collection of arguments for invoking getServices.
 */
export interface GetServicesArgs {
    readonly services?: inputs.azure.GetServicesService[];
}

/**
 * A collection of values returned by getServices.
 */
export interface GetServicesResult {
    readonly services?: outputs.azure.GetServicesService[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
