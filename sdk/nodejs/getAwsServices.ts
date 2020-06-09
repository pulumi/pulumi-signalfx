// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/** @deprecated signalfx.getAwsServices has been deprecated in favor of signalfx.aws.getServices */
export function getAwsServices(args?: GetAwsServicesArgs, opts?: pulumi.InvokeOptions): Promise<GetAwsServicesResult> {
    pulumi.log.warn("getAwsServices is deprecated: signalfx.getAwsServices has been deprecated in favor of signalfx.aws.getServices")
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("signalfx:index/getAwsServices:getAwsServices", {
        "services": args.services,
    }, opts);
}

/**
 * A collection of arguments for invoking getAwsServices.
 */
export interface GetAwsServicesArgs {
    readonly services?: inputs.GetAwsServicesService[];
}

/**
 * A collection of values returned by getAwsServices.
 */
export interface GetAwsServicesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly services?: outputs.GetAwsServicesService[];
}
