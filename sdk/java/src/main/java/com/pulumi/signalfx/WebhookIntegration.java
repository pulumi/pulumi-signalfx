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
 * SignalFx Webhook integration.
 * 
 * &gt; **NOTE** When managing integrations use a session token for an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you&#39;ll receive a 4xx error.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.WebhookIntegration;
 * import com.pulumi.signalfx.WebhookIntegrationArgs;
 * import com.pulumi.signalfx.inputs.WebhookIntegrationHeaderArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var webhookMyteam = new WebhookIntegration(&#34;webhookMyteam&#34;, WebhookIntegrationArgs.builder()        
 *             .enabled(true)
 *             .headers(WebhookIntegrationHeaderArgs.builder()
 *                 .headerKey(&#34;some_header&#34;)
 *                 .headerValue(&#34;value_for_that_header&#34;)
 *                 .build())
 *             .sharedSecret(&#34;abc1234&#34;)
 *             .url(&#34;https://www.example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="signalfx:index/webhookIntegration:WebhookIntegration")
public class WebhookIntegration extends com.pulumi.resources.CustomResource {
    /**
     * Whether the integration is enabled.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
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
    @Export(name="headers", type=List.class, parameters={WebhookIntegrationHeader.class})
    private Output</* @Nullable */ List<WebhookIntegrationHeader>> headers;

    /**
     * @return A header to send with the request
     * 
     */
    public Output<Optional<List<WebhookIntegrationHeader>>> headers() {
        return Codegen.optional(this.headers);
    }
    /**
     * Name of the integration.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the integration.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="sharedSecret", type=String.class, parameters={})
    private Output</* @Nullable */ String> sharedSecret;

    public Output<Optional<String>> sharedSecret() {
        return Codegen.optional(this.sharedSecret);
    }
    /**
     * The URL to request
     * 
     */
    @Export(name="url", type=String.class, parameters={})
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
    public WebhookIntegration(String name) {
        this(name, WebhookIntegrationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public WebhookIntegration(String name, WebhookIntegrationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public WebhookIntegration(String name, WebhookIntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/webhookIntegration:WebhookIntegration", name, args == null ? WebhookIntegrationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private WebhookIntegration(String name, Output<String> id, @Nullable WebhookIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/webhookIntegration:WebhookIntegration", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static WebhookIntegration get(String name, Output<String> id, @Nullable WebhookIntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new WebhookIntegration(name, id, state, options);
    }
}
