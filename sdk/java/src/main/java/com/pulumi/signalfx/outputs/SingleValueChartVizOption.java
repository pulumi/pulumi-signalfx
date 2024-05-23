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
public final class SingleValueChartVizOption {
    /**
     * @return The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     * 
     */
    private @Nullable String color;
    /**
     * @return Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     * 
     */
    private @Nullable String displayName;
    /**
     * @return Label used in the publish statement that displays the plot (metric time series data) you want to customize.
     * 
     */
    private String label;
    /**
     * @return , `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
     * 
     */
    private @Nullable String valuePrefix;
    /**
     * @return An arbitrary suffix to display with the value of this plot
     * 
     */
    private @Nullable String valueSuffix;
    /**
     * @return A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
     * 
     */
    private @Nullable String valueUnit;

    private SingleValueChartVizOption() {}
    /**
     * @return The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     * 
     */
    public Optional<String> color() {
        return Optional.ofNullable(this.color);
    }
    /**
     * @return Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     * 
     */
    public Optional<String> displayName() {
        return Optional.ofNullable(this.displayName);
    }
    /**
     * @return Label used in the publish statement that displays the plot (metric time series data) you want to customize.
     * 
     */
    public String label() {
        return this.label;
    }
    /**
     * @return , `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
     * 
     */
    public Optional<String> valuePrefix() {
        return Optional.ofNullable(this.valuePrefix);
    }
    /**
     * @return An arbitrary suffix to display with the value of this plot
     * 
     */
    public Optional<String> valueSuffix() {
        return Optional.ofNullable(this.valueSuffix);
    }
    /**
     * @return A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
     * 
     */
    public Optional<String> valueUnit() {
        return Optional.ofNullable(this.valueUnit);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SingleValueChartVizOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String color;
        private @Nullable String displayName;
        private String label;
        private @Nullable String valuePrefix;
        private @Nullable String valueSuffix;
        private @Nullable String valueUnit;
        public Builder() {}
        public Builder(SingleValueChartVizOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.color = defaults.color;
    	      this.displayName = defaults.displayName;
    	      this.label = defaults.label;
    	      this.valuePrefix = defaults.valuePrefix;
    	      this.valueSuffix = defaults.valueSuffix;
    	      this.valueUnit = defaults.valueUnit;
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
              throw new MissingRequiredPropertyException("SingleValueChartVizOption", "label");
            }
            this.label = label;
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
        public SingleValueChartVizOption build() {
            final var _resultValue = new SingleValueChartVizOption();
            _resultValue.color = color;
            _resultValue.displayName = displayName;
            _resultValue.label = label;
            _resultValue.valuePrefix = valuePrefix;
            _resultValue.valueSuffix = valueSuffix;
            _resultValue.valueUnit = valueUnit;
            return _resultValue;
        }
    }
}
