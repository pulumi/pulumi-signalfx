// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.ProviderArgs;
import com.pulumi.signalfx.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the signalfx package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:signalfx")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * API URL for your Splunk Observability Cloud org, may include a realm
     * 
     */
    @Export(name="apiUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> apiUrl;

    /**
     * @return API URL for your Splunk Observability Cloud org, may include a realm
     * 
     */
    public Output<Optional<String>> apiUrl() {
        return Codegen.optional(this.apiUrl);
    }
    /**
     * Splunk Observability Cloud auth token
     * 
     */
    @Export(name="authToken", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authToken;

    /**
     * @return Splunk Observability Cloud auth token
     * 
     */
    public Output<Optional<String>> authToken() {
        return Codegen.optional(this.authToken);
    }
    /**
     * Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
     * 
     */
    @Export(name="customAppUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> customAppUrl;

    /**
     * @return Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
     * 
     */
    public Output<Optional<String>> customAppUrl() {
        return Codegen.optional(this.customAppUrl);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, @Nullable ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, @Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
