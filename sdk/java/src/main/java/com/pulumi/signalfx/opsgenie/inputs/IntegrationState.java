// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.opsgenie.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IntegrationState extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationState Empty = new IntegrationState();

    /**
     * The API key
     * 
     */
    @Import(name="apiKey")
    private @Nullable Output<String> apiKey;

    /**
     * @return The API key
     * 
     */
    public Optional<Output<String>> apiKey() {
        return Optional.ofNullable(this.apiKey);
    }

    /**
     * Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
     * 
     */
    @Import(name="apiUrl")
    private @Nullable Output<String> apiUrl;

    /**
     * @return Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
     * 
     */
    public Optional<Output<String>> apiUrl() {
        return Optional.ofNullable(this.apiUrl);
    }

    /**
     * Whether the integration is enabled.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether the integration is enabled.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Name of the integration.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the integration.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private IntegrationState() {}

    private IntegrationState(IntegrationState $) {
        this.apiKey = $.apiKey;
        this.apiUrl = $.apiUrl;
        this.enabled = $.enabled;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationState $;

        public Builder() {
            $ = new IntegrationState();
        }

        public Builder(IntegrationState defaults) {
            $ = new IntegrationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiKey The API key
         * 
         * @return builder
         * 
         */
        public Builder apiKey(@Nullable Output<String> apiKey) {
            $.apiKey = apiKey;
            return this;
        }

        /**
         * @param apiKey The API key
         * 
         * @return builder
         * 
         */
        public Builder apiKey(String apiKey) {
            return apiKey(Output.of(apiKey));
        }

        /**
         * @param apiUrl Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
         * 
         * @return builder
         * 
         */
        public Builder apiUrl(@Nullable Output<String> apiUrl) {
            $.apiUrl = apiUrl;
            return this;
        }

        /**
         * @param apiUrl Opsgenie API URL. Will default to `https://api.opsgenie.com`. You might also want `https://api.eu.opsgenie.com`.
         * 
         * @return builder
         * 
         */
        public Builder apiUrl(String apiUrl) {
            return apiUrl(Output.of(apiUrl));
        }

        /**
         * @param enabled Whether the integration is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether the integration is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param name Name of the integration.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the integration.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public IntegrationState build() {
            return $;
        }
    }

}
