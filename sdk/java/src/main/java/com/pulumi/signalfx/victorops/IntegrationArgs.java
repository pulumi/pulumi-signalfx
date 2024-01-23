// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.victorops;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IntegrationArgs extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationArgs Empty = new IntegrationArgs();

    /**
     * Whether the integration is enabled or not
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Whether the integration is enabled or not
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * Name of the integration
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the integration
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Opsgenie API URL for integration
     * 
     */
    @Import(name="postUrl")
    private @Nullable Output<String> postUrl;

    /**
     * @return Opsgenie API URL for integration
     * 
     */
    public Optional<Output<String>> postUrl() {
        return Optional.ofNullable(this.postUrl);
    }

    private IntegrationArgs() {}

    private IntegrationArgs(IntegrationArgs $) {
        this.enabled = $.enabled;
        this.name = $.name;
        this.postUrl = $.postUrl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationArgs $;

        public Builder() {
            $ = new IntegrationArgs();
        }

        public Builder(IntegrationArgs defaults) {
            $ = new IntegrationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Whether the integration is enabled or not
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether the integration is enabled or not
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param name Name of the integration
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the integration
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param postUrl Opsgenie API URL for integration
         * 
         * @return builder
         * 
         */
        public Builder postUrl(@Nullable Output<String> postUrl) {
            $.postUrl = postUrl;
            return this;
        }

        /**
         * @param postUrl Opsgenie API URL for integration
         * 
         * @return builder
         * 
         */
        public Builder postUrl(String postUrl) {
            return postUrl(Output.of(postUrl));
        }

        public IntegrationArgs build() {
            if ($.enabled == null) {
                throw new MissingRequiredPropertyException("IntegrationArgs", "enabled");
            }
            return $;
        }
    }

}
