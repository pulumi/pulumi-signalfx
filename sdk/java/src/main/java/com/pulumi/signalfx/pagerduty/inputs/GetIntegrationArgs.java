// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.pagerduty.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetIntegrationArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIntegrationArgs Empty = new GetIntegrationArgs();

    /**
     * Specify the exact name of the desired PagerDuty integration
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Specify the exact name of the desired PagerDuty integration
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetIntegrationArgs() {}

    private GetIntegrationArgs(GetIntegrationArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIntegrationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIntegrationArgs $;

        public Builder() {
            $ = new GetIntegrationArgs();
        }

        public Builder(GetIntegrationArgs defaults) {
            $ = new GetIntegrationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Specify the exact name of the desired PagerDuty integration
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Specify the exact name of the desired PagerDuty integration
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetIntegrationArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
