// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TimeChartVizOption {
    private @Nullable String axis;
    private @Nullable String color;
    private @Nullable String displayName;
    private String label;
    private @Nullable String plotType;
    private @Nullable String valuePrefix;
    private @Nullable String valueSuffix;
    private @Nullable String valueUnit;

    private TimeChartVizOption() {}
    public Optional<String> axis() {
        return Optional.ofNullable(this.axis);
    }
    public Optional<String> color() {
        return Optional.ofNullable(this.color);
    }
    public Optional<String> displayName() {
        return Optional.ofNullable(this.displayName);
    }
    public String label() {
        return this.label;
    }
    public Optional<String> plotType() {
        return Optional.ofNullable(this.plotType);
    }
    public Optional<String> valuePrefix() {
        return Optional.ofNullable(this.valuePrefix);
    }
    public Optional<String> valueSuffix() {
        return Optional.ofNullable(this.valueSuffix);
    }
    public Optional<String> valueUnit() {
        return Optional.ofNullable(this.valueUnit);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TimeChartVizOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String axis;
        private @Nullable String color;
        private @Nullable String displayName;
        private String label;
        private @Nullable String plotType;
        private @Nullable String valuePrefix;
        private @Nullable String valueSuffix;
        private @Nullable String valueUnit;
        public Builder() {}
        public Builder(TimeChartVizOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.axis = defaults.axis;
    	      this.color = defaults.color;
    	      this.displayName = defaults.displayName;
    	      this.label = defaults.label;
    	      this.plotType = defaults.plotType;
    	      this.valuePrefix = defaults.valuePrefix;
    	      this.valueSuffix = defaults.valueSuffix;
    	      this.valueUnit = defaults.valueUnit;
        }

        @CustomType.Setter
        public Builder axis(@Nullable String axis) {

            this.axis = axis;
            return this;
        }
        @CustomType.Setter
        public Builder color(@Nullable String color) {

            this.color = color;
            return this;
        }
        @CustomType.Setter
        public Builder displayName(@Nullable String displayName) {

            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder label(String label) {
            if (label == null) {
              throw new MissingRequiredPropertyException("TimeChartVizOption", "label");
            }
            this.label = label;
            return this;
        }
        @CustomType.Setter
        public Builder plotType(@Nullable String plotType) {

            this.plotType = plotType;
            return this;
        }
        @CustomType.Setter
        public Builder valuePrefix(@Nullable String valuePrefix) {

            this.valuePrefix = valuePrefix;
            return this;
        }
        @CustomType.Setter
        public Builder valueSuffix(@Nullable String valueSuffix) {

            this.valueSuffix = valueSuffix;
            return this;
        }
        @CustomType.Setter
        public Builder valueUnit(@Nullable String valueUnit) {

            this.valueUnit = valueUnit;
            return this;
        }
        public TimeChartVizOption build() {
            final var _resultValue = new TimeChartVizOption();
            _resultValue.axis = axis;
            _resultValue.color = color;
            _resultValue.displayName = displayName;
            _resultValue.label = label;
            _resultValue.plotType = plotType;
            _resultValue.valuePrefix = valuePrefix;
            _resultValue.valueSuffix = valueSuffix;
            _resultValue.valueUnit = valueUnit;
            return _resultValue;
        }
    }
}
