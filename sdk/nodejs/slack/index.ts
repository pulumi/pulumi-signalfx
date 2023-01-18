// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { IntegrationArgs, IntegrationState } from "./integration";
export type Integration = import("./integration").Integration;
export const Integration: typeof import("./integration").Integration = null as any;
utilities.lazyLoad(exports, ["Integration"], () => require("./integration"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "signalfx:slack/integration:Integration":
                return new Integration(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("signalfx", "slack/integration", _module)
