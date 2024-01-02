// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DetectorVizOptionArgs extends com.pulumi.resources.ResourceArgs {

    public static final DetectorVizOptionArgs Empty = new DetectorVizOptionArgs();

    /**
     * Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
     * 
     */
    @Import(name="color")
    private @Nullable Output<String> color;

    /**
     * @return Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
     * 
     */
    public Optional<Output<String>> color() {
        return Optional.ofNullable(this.color);
    }

    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Label used in the publish statement that displays the plot (metric time series data) you want to customize.
     * 
     */
    @Import(name="label", required=true)
    private Output<String> label;

    /**
     * @return Label used in the publish statement that displays the plot (metric time series data) you want to customize.
     * 
     */
    public Output<String> label() {
        return this.label;
    }

    /**
     * , `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
     * 
     * **Notes**
     * 
     * It is highly recommended that you use both `max_delay` in your detector configuration and an `extrapolation` policy in your program text to reduce false positives/negatives.
     * 
     * `max_delay` allows SignalFx to continue with computation if there is a lag in receiving data points.
     * 
     * `extrapolation` allows you to specify how to handle missing data. An extrapolation policy can be added to individual signals by updating the data block in your `program_text`.
     * 
     * See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info.
     * 
     */
    @Import(name="valuePrefix")
    private @Nullable Output<String> valuePrefix;

    /**
     * @return , `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
     * 
     * **Notes**
     * 
     * It is highly recommended that you use both `max_delay` in your detector configuration and an `extrapolation` policy in your program text to reduce false positives/negatives.
     * 
     * `max_delay` allows SignalFx to continue with computation if there is a lag in receiving data points.
     * 
     * `extrapolation` allows you to specify how to handle missing data. An extrapolation policy can be added to individual signals by updating the data block in your `program_text`.
     * 
     * See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info.
     * 
     */
    public Optional<Output<String>> valuePrefix() {
        return Optional.ofNullable(this.valuePrefix);
    }

    @Import(name="valueSuffix")
    private @Nullable Output<String> valueSuffix;

    public Optional<Output<String>> valueSuffix() {
        return Optional.ofNullable(this.valueSuffix);
    }

    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
     * 
     */
    @Import(name="valueUnit")
    private @Nullable Output<String> valueUnit;

    /**
     * @return A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
     * 
     */
    public Optional<Output<String>> valueUnit() {
        return Optional.ofNullable(this.valueUnit);
    }

    private DetectorVizOptionArgs() {}

    private DetectorVizOptionArgs(DetectorVizOptionArgs $) {
        this.color = $.color;
        this.displayName = $.displayName;
        this.label = $.label;
        this.valuePrefix = $.valuePrefix;
        this.valueSuffix = $.valueSuffix;
        this.valueUnit = $.valueUnit;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DetectorVizOptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DetectorVizOptionArgs $;

        public Builder() {
            $ = new DetectorVizOptionArgs();
        }

        public Builder(DetectorVizOptionArgs defaults) {
            $ = new DetectorVizOptionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param color Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
         * 
         * @return builder
         * 
         */
        public Builder color(@Nullable Output<String> color) {
            $.color = color;
            return this;
        }

        /**
         * @param color Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
         * 
         * @return builder
         * 
         */
        public Builder color(String color) {
            return color(Output.of(color));
        }

        /**
         * @param displayName Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param label Label used in the publish statement that displays the plot (metric time series data) you want to customize.
         * 
         * @return builder
         * 
         */
        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label Label used in the publish statement that displays the plot (metric time series data) you want to customize.
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param valuePrefix , `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
         * 
         * **Notes**
         * 
         * It is highly recommended that you use both `max_delay` in your detector configuration and an `extrapolation` policy in your program text to reduce false positives/negatives.
         * 
         * `max_delay` allows SignalFx to continue with computation if there is a lag in receiving data points.
         * 
         * `extrapolation` allows you to specify how to handle missing data. An extrapolation policy can be added to individual signals by updating the data block in your `program_text`.
         * 
         * See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info.
         * 
         * @return builder
         * 
         */
        public Builder valuePrefix(@Nullable Output<String> valuePrefix) {
            $.valuePrefix = valuePrefix;
            return this;
        }

        /**
         * @param valuePrefix , `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
         * 
         * **Notes**
         * 
         * It is highly recommended that you use both `max_delay` in your detector configuration and an `extrapolation` policy in your program text to reduce false positives/negatives.
         * 
         * `max_delay` allows SignalFx to continue with computation if there is a lag in receiving data points.
         * 
         * `extrapolation` allows you to specify how to handle missing data. An extrapolation policy can be added to individual signals by updating the data block in your `program_text`.
         * 
         * See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info.
         * 
         * @return builder
         * 
         */
        public Builder valuePrefix(String valuePrefix) {
            return valuePrefix(Output.of(valuePrefix));
        }

        public Builder valueSuffix(@Nullable Output<String> valueSuffix) {
            $.valueSuffix = valueSuffix;
            return this;
        }

        public Builder valueSuffix(String valueSuffix) {
            return valueSuffix(Output.of(valueSuffix));
        }

        /**
         * @param valueUnit A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
         * 
         * @return builder
         * 
         */
        public Builder valueUnit(@Nullable Output<String> valueUnit) {
            $.valueUnit = valueUnit;
            return this;
        }

        /**
         * @param valueUnit A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
         * 
         * @return builder
         * 
         */
        public Builder valueUnit(String valueUnit) {
            return valueUnit(Output.of(valueUnit));
        }

        public DetectorVizOptionArgs build() {
            if ($.label == null) {
                throw new MissingRequiredPropertyException("DetectorVizOptionArgs", "label");
            }
            return $;
        }
    }

}
