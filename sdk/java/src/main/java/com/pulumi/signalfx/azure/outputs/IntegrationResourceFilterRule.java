// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.azure.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.signalfx.azure.outputs.IntegrationResourceFilterRuleFilter;
import java.util.Objects;

@CustomType
public final class IntegrationResourceFilterRule {
    private IntegrationResourceFilterRuleFilter filter;

    private IntegrationResourceFilterRule() {}
    public IntegrationResourceFilterRuleFilter filter() {
        return this.filter;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IntegrationResourceFilterRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private IntegrationResourceFilterRuleFilter filter;
        public Builder() {}
        public Builder(IntegrationResourceFilterRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filter = defaults.filter;
        }

        @CustomType.Setter
        public Builder filter(IntegrationResourceFilterRuleFilter filter) {
            this.filter = Objects.requireNonNull(filter);
            return this;
        }
        public IntegrationResourceFilterRule build() {
            final var o = new IntegrationResourceFilterRule();
            o.filter = filter;
            return o;
        }
    }
}