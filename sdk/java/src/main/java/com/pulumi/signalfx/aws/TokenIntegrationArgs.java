// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TokenIntegrationArgs extends com.pulumi.resources.ResourceArgs {

    public static final TokenIntegrationArgs Empty = new TokenIntegrationArgs();

    /**
     * The name of this integration
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of this integration
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private TokenIntegrationArgs() {}

    private TokenIntegrationArgs(TokenIntegrationArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TokenIntegrationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TokenIntegrationArgs $;

        public Builder() {
            $ = new TokenIntegrationArgs();
        }

        public Builder(TokenIntegrationArgs defaults) {
            $ = new TokenIntegrationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of this integration
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of this integration
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public TokenIntegrationArgs build() {
            return $;
        }
    }

}
