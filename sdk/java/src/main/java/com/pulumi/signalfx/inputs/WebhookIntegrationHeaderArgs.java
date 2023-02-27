// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class WebhookIntegrationHeaderArgs extends com.pulumi.resources.ResourceArgs {

    public static final WebhookIntegrationHeaderArgs Empty = new WebhookIntegrationHeaderArgs();

    /**
     * The key of the header to send
     * 
     */
    @Import(name="headerKey", required=true)
    private Output<String> headerKey;

    /**
     * @return The key of the header to send
     * 
     */
    public Output<String> headerKey() {
        return this.headerKey;
    }

    /**
     * The value of the header to send
     * 
     */
    @Import(name="headerValue", required=true)
    private Output<String> headerValue;

    /**
     * @return The value of the header to send
     * 
     */
    public Output<String> headerValue() {
        return this.headerValue;
    }

    private WebhookIntegrationHeaderArgs() {}

    private WebhookIntegrationHeaderArgs(WebhookIntegrationHeaderArgs $) {
        this.headerKey = $.headerKey;
        this.headerValue = $.headerValue;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WebhookIntegrationHeaderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WebhookIntegrationHeaderArgs $;

        public Builder() {
            $ = new WebhookIntegrationHeaderArgs();
        }

        public Builder(WebhookIntegrationHeaderArgs defaults) {
            $ = new WebhookIntegrationHeaderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param headerKey The key of the header to send
         * 
         * @return builder
         * 
         */
        public Builder headerKey(Output<String> headerKey) {
            $.headerKey = headerKey;
            return this;
        }

        /**
         * @param headerKey The key of the header to send
         * 
         * @return builder
         * 
         */
        public Builder headerKey(String headerKey) {
            return headerKey(Output.of(headerKey));
        }

        /**
         * @param headerValue The value of the header to send
         * 
         * @return builder
         * 
         */
        public Builder headerValue(Output<String> headerValue) {
            $.headerValue = headerValue;
            return this;
        }

        /**
         * @param headerValue The value of the header to send
         * 
         * @return builder
         * 
         */
        public Builder headerValue(String headerValue) {
            return headerValue(Output.of(headerValue));
        }

        public WebhookIntegrationHeaderArgs build() {
            $.headerKey = Objects.requireNonNull($.headerKey, "expected parameter 'headerKey' to be non-null");
            $.headerValue = Objects.requireNonNull($.headerValue, "expected parameter 'headerValue' to be non-null");
            return $;
        }
    }

}