// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.azure.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class IntegrationResourceFilterRuleFilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationResourceFilterRuleFilterArgs Empty = new IntegrationResourceFilterRuleFilterArgs();

    @Import(name="source", required=true)
    private Output<String> source;

    public Output<String> source() {
        return this.source;
    }

    private IntegrationResourceFilterRuleFilterArgs() {}

    private IntegrationResourceFilterRuleFilterArgs(IntegrationResourceFilterRuleFilterArgs $) {
        this.source = $.source;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationResourceFilterRuleFilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationResourceFilterRuleFilterArgs $;

        public Builder() {
            $ = new IntegrationResourceFilterRuleFilterArgs();
        }

        public Builder(IntegrationResourceFilterRuleFilterArgs defaults) {
            $ = new IntegrationResourceFilterRuleFilterArgs(Objects.requireNonNull(defaults));
        }

        public Builder source(Output<String> source) {
            $.source = source;
            return this;
        }

        public Builder source(String source) {
            return source(Output.of(source));
        }

        public IntegrationResourceFilterRuleFilterArgs build() {
            $.source = Objects.requireNonNull($.source, "expected parameter 'source' to be non-null");
            return $;
        }
    }

}
