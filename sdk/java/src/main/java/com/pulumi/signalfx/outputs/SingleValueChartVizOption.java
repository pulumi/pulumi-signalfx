// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
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
    private final @Nullable String color;
    /**
     * @return Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     * 
     */
    private final @Nullable String displayName;
    /**
     * @return Label used in the publish statement that displays the plot (metric time series data) you want to customize.
     * 
     */
    private final String label;
    private final @Nullable String valuePrefix;
    private final @Nullable String valueSuffix;
    /**
     * @return A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
     * * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
     * 
     */
    private final @Nullable String valueUnit;

    @CustomType.Constructor
    private SingleValueChartVizOption(
        @CustomType.Parameter("color") @Nullable String color,
        @CustomType.Parameter("displayName") @Nullable String displayName,
        @CustomType.Parameter("label") String label,
        @CustomType.Parameter("valuePrefix") @Nullable String valuePrefix,
        @CustomType.Parameter("valueSuffix") @Nullable String valueSuffix,
        @CustomType.Parameter("valueUnit") @Nullable String valueUnit) {
        this.color = color;
        this.displayName = displayName;
        this.label = label;
        this.valuePrefix = valuePrefix;
        this.valueSuffix = valueSuffix;
        this.valueUnit = valueUnit;
    }

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
    public Optional<String> valuePrefix() {
        return Optional.ofNullable(this.valuePrefix);
    }
    public Optional<String> valueSuffix() {
        return Optional.ofNullable(this.valueSuffix);
    }
    /**
     * @return A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
     * * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
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

    public static final class Builder {
        private @Nullable String color;
        private @Nullable String displayName;
        private String label;
        private @Nullable String valuePrefix;
        private @Nullable String valueSuffix;
        private @Nullable String valueUnit;

        public Builder() {
    	      // Empty
        }

        public Builder(SingleValueChartVizOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.color = defaults.color;
    	      this.displayName = defaults.displayName;
    	      this.label = defaults.label;
    	      this.valuePrefix = defaults.valuePrefix;
    	      this.valueSuffix = defaults.valueSuffix;
    	      this.valueUnit = defaults.valueUnit;
        }

        public Builder color(@Nullable String color) {
            this.color = color;
            return this;
        }
        public Builder displayName(@Nullable String displayName) {
            this.displayName = displayName;
            return this;
        }
        public Builder label(String label) {
            this.label = Objects.requireNonNull(label);
            return this;
        }
        public Builder valuePrefix(@Nullable String valuePrefix) {
            this.valuePrefix = valuePrefix;
            return this;
        }
        public Builder valueSuffix(@Nullable String valueSuffix) {
            this.valueSuffix = valueSuffix;
            return this;
        }
        public Builder valueUnit(@Nullable String valueUnit) {
            this.valueUnit = valueUnit;
            return this;
        }        public SingleValueChartVizOption build() {
            return new SingleValueChartVizOption(color, displayName, label, valuePrefix, valueSuffix, valueUnit);
        }
    }
}
