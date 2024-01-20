// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TimeChartLegendOptionsField {
    private @Nullable Boolean enabled;
    private String property;

    private TimeChartLegendOptionsField() {}
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    public String property() {
        return this.property;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TimeChartLegendOptionsField defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enabled;
        private String property;
        public Builder() {}
        public Builder(TimeChartLegendOptionsField defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.property = defaults.property;
        }

        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {

            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder property(String property) {
            if (property == null) {
              throw new MissingRequiredPropertyException("TimeChartLegendOptionsField", "property");
            }
            this.property = property;
            return this;
        }
        public TimeChartLegendOptionsField build() {
            final var _resultValue = new TimeChartLegendOptionsField();
            _resultValue.enabled = enabled;
            _resultValue.property = property;
            return _resultValue;
        }
    }
}
