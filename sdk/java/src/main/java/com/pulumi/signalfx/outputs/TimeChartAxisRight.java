// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.signalfx.outputs.TimeChartAxisRightWatermark;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TimeChartAxisRight {
    /**
     * @return A line to draw as a high watermark.
     * 
     */
    private @Nullable Double highWatermark;
    /**
     * @return A label to attach to the high watermark line.
     * 
     */
    private @Nullable String highWatermarkLabel;
    /**
     * @return Label used in the publish statement that displays the event query you want to customize.
     * 
     */
    private @Nullable String label;
    /**
     * @return A line to draw as a low watermark.
     * 
     */
    private @Nullable Double lowWatermark;
    /**
     * @return A label to attach to the low watermark line.
     * 
     */
    private @Nullable String lowWatermarkLabel;
    /**
     * @return The maximum value for the right axis.
     * 
     */
    private @Nullable Double maxValue;
    /**
     * @return The minimum value for the right axis.
     * 
     */
    private @Nullable Double minValue;
    private @Nullable List<TimeChartAxisRightWatermark> watermarks;

    private TimeChartAxisRight() {}
    /**
     * @return A line to draw as a high watermark.
     * 
     */
    public Optional<Double> highWatermark() {
        return Optional.ofNullable(this.highWatermark);
    }
    /**
     * @return A label to attach to the high watermark line.
     * 
     */
    public Optional<String> highWatermarkLabel() {
        return Optional.ofNullable(this.highWatermarkLabel);
    }
    /**
     * @return Label used in the publish statement that displays the event query you want to customize.
     * 
     */
    public Optional<String> label() {
        return Optional.ofNullable(this.label);
    }
    /**
     * @return A line to draw as a low watermark.
     * 
     */
    public Optional<Double> lowWatermark() {
        return Optional.ofNullable(this.lowWatermark);
    }
    /**
     * @return A label to attach to the low watermark line.
     * 
     */
    public Optional<String> lowWatermarkLabel() {
        return Optional.ofNullable(this.lowWatermarkLabel);
    }
    /**
     * @return The maximum value for the right axis.
     * 
     */
    public Optional<Double> maxValue() {
        return Optional.ofNullable(this.maxValue);
    }
    /**
     * @return The minimum value for the right axis.
     * 
     */
    public Optional<Double> minValue() {
        return Optional.ofNullable(this.minValue);
    }
    public List<TimeChartAxisRightWatermark> watermarks() {
        return this.watermarks == null ? List.of() : this.watermarks;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TimeChartAxisRight defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Double highWatermark;
        private @Nullable String highWatermarkLabel;
        private @Nullable String label;
        private @Nullable Double lowWatermark;
        private @Nullable String lowWatermarkLabel;
        private @Nullable Double maxValue;
        private @Nullable Double minValue;
        private @Nullable List<TimeChartAxisRightWatermark> watermarks;
        public Builder() {}
        public Builder(TimeChartAxisRight defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.highWatermark = defaults.highWatermark;
    	      this.highWatermarkLabel = defaults.highWatermarkLabel;
    	      this.label = defaults.label;
    	      this.lowWatermark = defaults.lowWatermark;
    	      this.lowWatermarkLabel = defaults.lowWatermarkLabel;
    	      this.maxValue = defaults.maxValue;
    	      this.minValue = defaults.minValue;
    	      this.watermarks = defaults.watermarks;
        }

        @CustomType.Setter
        public Builder highWatermark(@Nullable Double highWatermark) {
            this.highWatermark = highWatermark;
            return this;
        }
        @CustomType.Setter
        public Builder highWatermarkLabel(@Nullable String highWatermarkLabel) {
            this.highWatermarkLabel = highWatermarkLabel;
            return this;
        }
        @CustomType.Setter
        public Builder label(@Nullable String label) {
            this.label = label;
            return this;
        }
        @CustomType.Setter
        public Builder lowWatermark(@Nullable Double lowWatermark) {
            this.lowWatermark = lowWatermark;
            return this;
        }
        @CustomType.Setter
        public Builder lowWatermarkLabel(@Nullable String lowWatermarkLabel) {
            this.lowWatermarkLabel = lowWatermarkLabel;
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
        @CustomType.Setter
        public Builder watermarks(@Nullable List<TimeChartAxisRightWatermark> watermarks) {
            this.watermarks = watermarks;
            return this;
        }
        public Builder watermarks(TimeChartAxisRightWatermark... watermarks) {
            return watermarks(List.of(watermarks));
        }
        public TimeChartAxisRight build() {
            final var _resultValue = new TimeChartAxisRight();
            _resultValue.highWatermark = highWatermark;
            _resultValue.highWatermarkLabel = highWatermarkLabel;
            _resultValue.label = label;
            _resultValue.lowWatermark = lowWatermark;
            _resultValue.lowWatermarkLabel = lowWatermarkLabel;
            _resultValue.maxValue = maxValue;
            _resultValue.minValue = minValue;
            _resultValue.watermarks = watermarks;
            return _resultValue;
        }
    }
}
