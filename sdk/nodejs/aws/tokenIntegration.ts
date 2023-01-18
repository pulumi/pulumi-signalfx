// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * SignalFx AWS CloudWatch integrations using security tokens. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
 *
 * > **NOTE** When managing integrations use a session token for an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
 *
 * > **WARNING** This resource implements a part of a workflow. You must use it with `signalfx.aws.Integration`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const awsMyteamToken = new signalfx.aws.TokenIntegration("awsMyteamToken", {});
 * // Make yourself an AWS IAM role here
 * const awsSfxRole = new aws.iam.Role("awsSfxRole", {});
 * // Stuff here that uses the external and account ID
 * const awsMyteam = new signalfx.aws.Integration("awsMyteam", {
 *     enabled: true,
 *     integrationId: awsMyteamToken.id,
 *     token: "put_your_token_here",
 *     key: "put_your_key_here",
 *     regions: ["us-east-1"],
 *     pollRate: 300,
 *     importCloudWatch: true,
 *     enableAwsUsage: true,
 *     customNamespaceSyncRules: [{
 *         defaultAction: "Exclude",
 *         filterAction: "Include",
 *         filterSource: "filter('code', '200')",
 *         namespace: "my-custom-namespace",
 *     }],
 *     namespaceSyncRules: [{
 *         defaultAction: "Exclude",
 *         filterAction: "Include",
 *         filterSource: "filter('code', '200')",
 *         namespace: "AWS/EC2",
 *     }],
 * });
 * ```
 */
export class TokenIntegration extends pulumi.CustomResource {
    /**
     * Get an existing TokenIntegration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TokenIntegrationState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["signalfxAwsAccount"] = state ? state.signalfxAwsAccount : undefined;
            resourceInputs["tokenId"] = state ? state.tokenId : undefined;
        } else {
            const args = argsOrState as TokenIntegrationArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["signalfxAwsAccount"] = undefined /*out*/;
            resourceInputs["tokenId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["signalfxAwsAccount", "tokenId"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(TokenIntegration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TokenIntegration resources.
 */
export interface TokenIntegrationState {
    /**
     * The name of this integration
     */
    name?: pulumi.Input<string>;
    /**
     * The AWS Account ARN to use with your policies/roles, provided by SignalFx.
     */
    signalfxAwsAccount?: pulumi.Input<string>;
    /**
     * The SignalFx-generated AWS token to use with an AWS integration.
     */
    tokenId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TokenIntegration resource.
 */
export interface TokenIntegrationArgs {
    /**
     * The name of this integration
     */
    name?: pulumi.Input<string>;
}
