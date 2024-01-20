// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.azure.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class IntegrationResourceFilterRule {
    private String filterSource;

    private IntegrationResourceFilterRule() {}
    public String filterSource() {
        return this.filterSource;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IntegrationResourceFilterRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String filterSource;
        public Builder() {}
        public Builder(IntegrationResourceFilterRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filterSource = defaults.filterSource;
        }

        @CustomType.Setter
        public Builder filterSource(String filterSource) {
            if (filterSource == null) {
              throw new MissingRequiredPropertyException("IntegrationResourceFilterRule", "filterSource");
            }
            this.filterSource = filterSource;
            return this;
        }
        public IntegrationResourceFilterRule build() {
            final var _resultValue = new IntegrationResourceFilterRule();
            _resultValue.filterSource = filterSource;
            return _resultValue;
        }
    }
}
