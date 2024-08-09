// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.WebhookIntegrationArgs;
import com.pulumi.signalfx.inputs.WebhookIntegrationState;
import com.pulumi.signalfx.outputs.WebhookIntegrationHeader;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Splunk Observability Cloud webhook integration.
 * 
 * &gt; **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you&#39;ll receive a 4xx error.
 * 
 * ## Example
 * 
 */
@ResourceType(type="signalfx:index/webhookIntegration:WebhookIntegration")
public class WebhookIntegration extends com.pulumi.resources.CustomResource {
    /**
     * Whether the integration is enabled.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return Whether the integration is enabled.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * A header to send with the request
     * 
     */
    @Export(name="headers", refs={List.class,WebhookIntegrationHeader.class}, tree="[0,1]")
    private Output</* @Nullable */ List<WebhookIntegrationHeader>> headers;

    /**
     * @return A header to send with the request
     * 
     */
    public Output<Optional<List<WebhookIntegrationHeader>>> headers() {
        return Codegen.optional(this.headers);
    }
    /**
     * HTTP method used for the webhook request, such as &#39;GET&#39;, &#39;POST&#39; and &#39;PUT&#39;
     * 
     */
    @Export(name="method", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> method;

    /**
     * @return HTTP method used for the webhook request, such as &#39;GET&#39;, &#39;POST&#39; and &#39;PUT&#39;
     * 
     */
    public Output<Optional<String>> method() {
        return Codegen.optional(this.method);
    }
    /**
     * Name of the integration.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the integration.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Template for the payload to be sent with the webhook request in JSON format
     * 
     */
    @Export(name="payloadTemplate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> payloadTemplate;

    /**
     * @return Template for the payload to be sent with the webhook request in JSON format
     * 
     */
    public Output<Optional<String>> payloadTemplate() {
        return Codegen.optional(this.payloadTemplate);
    }
    @Export(name="sharedSecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sharedSecret;

    public Output<Optional<String>> sharedSecret() {
        return Codegen.optional(this.sharedSecret);
    }
    /**
     * The URL to request
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> url;

    /**
     * @return The URL to request
     * 
     */
    public Output<Optional<String>> url() {
        return Codegen.optional(this.url);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public WebhookIntegration(java.lang.String name) {
        this(name, WebhookIntegrationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public WebhookIntegration(java.lang.String name, WebhookIntegrationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public WebhookIntegration(java.lang.String name, WebhookIntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/webhookIntegration:WebhookIntegration", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private WebhookIntegration(java.lang.String name, Output<java.lang.String> id, @Nullable WebhookIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/webhookIntegration:WebhookIntegration", name, state, makeResourceOptions(options, id), false);
    }

    private static WebhookIntegrationArgs makeArgs(WebhookIntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? WebhookIntegrationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "headers",
                "sharedSecret"
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
    public static WebhookIntegration get(java.lang.String name, Output<java.lang.String> id, @Nullable WebhookIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new WebhookIntegration(name, id, state, options);
    }
}
