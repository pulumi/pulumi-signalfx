// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * SignalFx Webhook integration.
 *
 * > **NOTE** When managing integrations use a session token for an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const webhookMyteam = new signalfx.WebhookIntegration("webhook_myteam", {
 *     enabled: true,
 *     headers: [{
 *         headerKey: "some_header",
 *         headerValue: "value_for_that_header",
 *     }],
 *     sharedSecret: "abc1234",
 *     url: "https://www.example.com",
 * });
 * ```
 */
export class WebhookIntegration extends pulumi.CustomResource {
    /**
     * Get an existing WebhookIntegration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebhookIntegrationState, opts?: pulumi.CustomResourceOptions): WebhookIntegration {
        return new WebhookIntegration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:index/webhookIntegration:WebhookIntegration';

    /**
     * Returns true if the given object is an instance of WebhookIntegration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebhookIntegration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WebhookIntegration.__pulumiType;
    }

    /**
     * Whether the integration is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * A header to send with the request
     */
    public readonly headers!: pulumi.Output<outputs.WebhookIntegrationHeader[] | undefined>;
    /**
     * Name of the integration.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly sharedSecret!: pulumi.Output<string | undefined>;
    /**
     * The URL to request
     */
    public readonly url!: pulumi.Output<string | undefined>;

    /**
     * Create a WebhookIntegration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookIntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebhookIntegrationArgs | WebhookIntegrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WebhookIntegrationState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["headers"] = state ? state.headers : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["sharedSecret"] = state ? state.sharedSecret : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as WebhookIntegrationArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["headers"] = args ? args.headers : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["sharedSecret"] = args ? args.sharedSecret : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WebhookIntegration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebhookIntegration resources.
 */
export interface WebhookIntegrationState {
    /**
     * Whether the integration is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * A header to send with the request
     */
    headers?: pulumi.Input<pulumi.Input<inputs.WebhookIntegrationHeader>[]>;
    /**
     * Name of the integration.
     */
    name?: pulumi.Input<string>;
    sharedSecret?: pulumi.Input<string>;
    /**
     * The URL to request
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WebhookIntegration resource.
 */
export interface WebhookIntegrationArgs {
    /**
     * Whether the integration is enabled.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * A header to send with the request
     */
    headers?: pulumi.Input<pulumi.Input<inputs.WebhookIntegrationHeader>[]>;
    /**
     * Name of the integration.
     */
    name?: pulumi.Input<string>;
    sharedSecret?: pulumi.Input<string>;
    /**
     * The URL to request
     */
    url?: pulumi.Input<string>;
}
