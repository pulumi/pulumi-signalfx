// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TimeChartAxisRightWatermark {
    /**
     * @return Label of the left axis.
     * 
     */
    private @Nullable String label;
    private Double value;

    private TimeChartAxisRightWatermark() {}
    /**
     * @return Label of the left axis.
     * 
     */
    public Optional<String> label() {
        return Optional.ofNullable(this.label);
    }
    public Double value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TimeChartAxisRightWatermark defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String label;
        private Double value;
        public Builder() {}
        public Builder(TimeChartAxisRightWatermark defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.label = defaults.label;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder label(@Nullable String label) {
            this.label = label;
            return this;
        }
        @CustomType.Setter
        public Builder value(Double value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public TimeChartAxisRightWatermark build() {
            final var o = new TimeChartAxisRightWatermark();
            o.label = label;
            o.value = value;
            return o;
        }
    }
}
