// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of AWS service names.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const awsServices = signalfx.aws.getServices({});
 * // Leaves out most of the integration bits, see the docs
 * // for signalfx_aws_integration for more
 * // …
 * const awsMyteam = new signalfx.aws.Integration("awsMyteam", {services: [awsServices.then(awsServices => awsServices.services)].map(__item => __item?.name)});
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
    return pulumi.runtime.invoke("signalfx:aws/getServices:getServices", {
        "services": args.services,
    }, opts);
}

/**
 * A collection of arguments for invoking getServices.
 */
export interface GetServicesArgs {
    readonly services?: inputs.aws.GetServicesService[];
}

/**
 * A collection of values returned by getServices.
 */
export interface GetServicesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly services?: outputs.aws.GetServicesService[];
}
