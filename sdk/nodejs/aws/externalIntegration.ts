// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Splunk Observability AWS CloudWatch integrations using Role ARNs. For help with this integration see [Connect to AWS CloudWatch](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/aws/aws-apiconfig.html).
 *
 * > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
 *
 * > **WARNING** This resource implements part of a workflow. Use it with `signalfx.aws.Integration`. Check with Splunk support for your realm's AWS account id.
 */
export class ExternalIntegration extends pulumi.CustomResource {
    /**
     * Get an existing ExternalIntegration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExternalIntegrationState, opts?: pulumi.CustomResourceOptions): ExternalIntegration {
        return new ExternalIntegration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:aws/externalIntegration:ExternalIntegration';

    /**
     * Returns true if the given object is an instance of ExternalIntegration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExternalIntegration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExternalIntegration.__pulumiType;
    }

    /**
     * The external ID to use with your IAM role and with `signalfx.aws.Integration`.
     */
    public /*out*/ readonly externalId!: pulumi.Output<string>;
    /**
     * The name of this integration
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The AWS Account ARN to use with your policies/roles, provided by Splunk Observability Cloud.
     */
    public /*out*/ readonly signalfxAwsAccount!: pulumi.Output<string>;

    /**
     * Create a ExternalIntegration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExternalIntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExternalIntegrationArgs | ExternalIntegrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ExternalIntegrationState | undefined;
            resourceInputs["externalId"] = state ? state.externalId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["signalfxAwsAccount"] = state ? state.signalfxAwsAccount : undefined;
        } else {
            const args = argsOrState as ExternalIntegrationArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["externalId"] = undefined /*out*/;
            resourceInputs["signalfxAwsAccount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["externalId", "signalfxAwsAccount"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ExternalIntegration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ExternalIntegration resources.
 */
export interface ExternalIntegrationState {
    /**
     * The external ID to use with your IAM role and with `signalfx.aws.Integration`.
     */
    externalId?: pulumi.Input<string>;
    /**
     * The name of this integration
     */
    name?: pulumi.Input<string>;
    /**
     * The AWS Account ARN to use with your policies/roles, provided by Splunk Observability Cloud.
     */
    signalfxAwsAccount?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ExternalIntegration resource.
 */
export interface ExternalIntegrationArgs {
    /**
     * The name of this integration
     */
    name?: pulumi.Input<string>;
}
