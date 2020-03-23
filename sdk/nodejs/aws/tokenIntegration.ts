// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * SignalFx AWS CloudWatch integrations using security tokens. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
 * 
 * > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.
 * 
 * > **WARNING** This resource implements a part of a workflow. You must use it with `signalfx.aws.Integration`.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/aws_token_integration.html.markdown.
 */
export class TokenIntegration extends pulumi.CustomResource {
    /**
     * Get an existing TokenIntegration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TokenIntegrationState, opts?: pulumi.CustomResourceOptions): TokenIntegration {
        return new TokenIntegration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:aws/tokenIntegration:TokenIntegration';

    /**
     * Returns true if the given object is an instance of TokenIntegration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TokenIntegration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TokenIntegration.__pulumiType;
    }

    /**
     * The name of this integration
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The AWS Account ARN to use with your policies/roles, provided by SignalFx.
     */
    public /*out*/ readonly signalfxAwsAccount!: pulumi.Output<string>;
    /**
     * The SignalFx-generated AWS token to use with an AWS integration.
     */
    public /*out*/ readonly tokenId!: pulumi.Output<string>;

    /**
     * Create a TokenIntegration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TokenIntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TokenIntegrationArgs | TokenIntegrationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TokenIntegrationState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["signalfxAwsAccount"] = state ? state.signalfxAwsAccount : undefined;
            inputs["tokenId"] = state ? state.tokenId : undefined;
        } else {
            const args = argsOrState as TokenIntegrationArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["signalfxAwsAccount"] = undefined /*out*/;
            inputs["tokenId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TokenIntegration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TokenIntegration resources.
 */
export interface TokenIntegrationState {
    /**
     * The name of this integration
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The AWS Account ARN to use with your policies/roles, provided by SignalFx.
     */
    readonly signalfxAwsAccount?: pulumi.Input<string>;
    /**
     * The SignalFx-generated AWS token to use with an AWS integration.
     */
    readonly tokenId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TokenIntegration resource.
 */
export interface TokenIntegrationArgs {
    /**
     * The name of this integration
     */
    readonly name?: pulumi.Input<string>;
}
