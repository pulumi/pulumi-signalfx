// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TimeChartAxisLeftWatermark {
    /**
     * @return Label used in the publish statement that displays the event query you want to customize.
     * 
     */
    private @Nullable String label;
    private Double value;

    private TimeChartAxisLeftWatermark() {}
    /**
     * @return Label used in the publish statement that displays the event query you want to customize.
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

    public static Builder builder(TimeChartAxisLeftWatermark defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String label;
        private Double value;
        public Builder() {}
        public Builder(TimeChartAxisLeftWatermark defaults) {
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
            if (value == null) {
              throw new MissingRequiredPropertyException("TimeChartAxisLeftWatermark", "value");
            }
            this.value = value;
            return this;
        }
        public TimeChartAxisLeftWatermark build() {
            final var _resultValue = new TimeChartAxisLeftWatermark();
            _resultValue.label = label;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}
