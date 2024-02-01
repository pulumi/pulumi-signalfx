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
public final class HeatmapChartColorRange {
    /**
     * @return The color range to use. The starting hex color value for data values in a heatmap chart. Specify the value as a 6-character hexadecimal value preceded by the &#39;#&#39; character, for example &#34;#ea1849&#34; (grass green).
     * 
     */
    private String color;
    /**
     * @return The maximum value within the coloring range
     * 
     */
    private @Nullable Double maxValue;
    /**
     * @return The minimum value within the coloring range
     * 
     */
    private @Nullable Double minValue;

    private HeatmapChartColorRange() {}
    /**
     * @return The color range to use. The starting hex color value for data values in a heatmap chart. Specify the value as a 6-character hexadecimal value preceded by the &#39;#&#39; character, for example &#34;#ea1849&#34; (grass green).
     * 
     */
    public String color() {
        return this.color;
    }
    /**
     * @return The maximum value within the coloring range
     * 
     */
    public Optional<Double> maxValue() {
        return Optional.ofNullable(this.maxValue);
    }
    /**
     * @return The minimum value within the coloring range
     * 
     */
    public Optional<Double> minValue() {
        return Optional.ofNullable(this.minValue);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(HeatmapChartColorRange defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String color;
        private @Nullable Double maxValue;
        private @Nullable Double minValue;
        public Builder() {}
        public Builder(HeatmapChartColorRange defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.color = defaults.color;
    	      this.maxValue = defaults.maxValue;
    	      this.minValue = defaults.minValue;
        }

        @CustomType.Setter
        public Builder color(String color) {
            if (color == null) {
              throw new MissingRequiredPropertyException("HeatmapChartColorRange", "color");
            }
            this.color = color;
            return this;
        }
        @CustomType.Setter
        public Builder maxValue(@Nullable Double maxValue) {

            this.maxValue = maxValue;
            return this;
        }
        @CustomType.Setter
        public Builder minValue(@Nullable Double minValue) {

            this.minValue = minValue;
            return this;
        }
        public HeatmapChartColorRange build() {
            final var _resultValue = new HeatmapChartColorRange();
            _resultValue.color = color;
            _resultValue.maxValue = maxValue;
            _resultValue.minValue = minValue;
            return _resultValue;
        }
    }
}
