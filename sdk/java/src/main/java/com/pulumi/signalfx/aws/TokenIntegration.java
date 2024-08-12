// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.aws.TokenIntegrationArgs;
import com.pulumi.signalfx.aws.inputs.TokenIntegrationState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Splunk Observability AWS CloudWatch integrations using security tokens. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
 * 
 * &gt; **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observabilit Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
 * 
 * &gt; **WARNING** This resource implements a part of a workflow. You must use it with `signalfx.aws.Integration`.
 * 
 * ## Example
 * 
 */
@ResourceType(type="signalfx:aws/tokenIntegration:TokenIntegration")
public class TokenIntegration extends com.pulumi.resources.CustomResource {
    /**
     * The name of this integration
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of this integration
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The AWS Account ARN to use with your policies/roles, provided by Splunk Observability Cloud.
     * 
     */
    @Export(name="signalfxAwsAccount", refs={String.class}, tree="[0]")
    private Output<String> signalfxAwsAccount;

    /**
     * @return The AWS Account ARN to use with your policies/roles, provided by Splunk Observability Cloud.
     * 
     */
    public Output<String> signalfxAwsAccount() {
        return this.signalfxAwsAccount;
    }
    /**
     * The SignalFx-generated AWS token to use with an AWS integration.
     * 
     */
    @Export(name="tokenId", refs={String.class}, tree="[0]")
    private Output<String> tokenId;

    /**
     * @return The SignalFx-generated AWS token to use with an AWS integration.
     * 
     */
    public Output<String> tokenId() {
        return this.tokenId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TokenIntegration(java.lang.String name) {
        this(name, TokenIntegrationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TokenIntegration(java.lang.String name, @Nullable TokenIntegrationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TokenIntegration(java.lang.String name, @Nullable TokenIntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:aws/tokenIntegration:TokenIntegration", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TokenIntegration(java.lang.String name, Output<java.lang.String> id, @Nullable TokenIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:aws/tokenIntegration:TokenIntegration", name, state, makeResourceOptions(options, id), false);
    }

    private static TokenIntegrationArgs makeArgs(@Nullable TokenIntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TokenIntegrationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "signalfxAwsAccount",
                "tokenId"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static TokenIntegration get(java.lang.String name, Output<java.lang.String> id, @Nullable TokenIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TokenIntegration(name, id, state, options);
    }
}
