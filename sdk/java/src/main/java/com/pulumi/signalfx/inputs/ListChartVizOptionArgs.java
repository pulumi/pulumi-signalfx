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


public final class ListChartVizOptionArgs extends com.pulumi.resources.ResourceArgs {

    public static final ListChartVizOptionArgs Empty = new ListChartVizOptionArgs();

    /**
     * Color to use
     * 
     */
    @Import(name="color")
    private @Nullable Output<String> color;

    /**
     * @return Color to use
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
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     * 
     */
    @Import(name="label", required=true)
    private Output<String> label;

    /**
     * @return The label used in the publish statement that displays the plot (metric time series data) you want to customize
     * 
     */
    public Output<String> label() {
        return this.label;
    }

    /**
     * An arbitrary prefix to display with the value of this plot
     * 
     */
    @Import(name="valuePrefix")
    private @Nullable Output<String> valuePrefix;

    /**
     * @return An arbitrary prefix to display with the value of this plot
     * 
     */
    public Optional<Output<String>> valuePrefix() {
        return Optional.ofNullable(this.valuePrefix);
    }

    /**
     * An arbitrary suffix to display with the value of this plot
     * 
     */
    @Import(name="valueSuffix")
    private @Nullable Output<String> valueSuffix;

    /**
     * @return An arbitrary suffix to display with the value of this plot
     * 
     */
    public Optional<Output<String>> valueSuffix() {
        return Optional.ofNullable(this.valueSuffix);
    }

    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     * 
     */
    @Import(name="valueUnit")
    private @Nullable Output<String> valueUnit;

    /**
     * @return A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     * 
     */
    public Optional<Output<String>> valueUnit() {
        return Optional.ofNullable(this.valueUnit);
    }

    private ListChartVizOptionArgs() {}

    private ListChartVizOptionArgs(ListChartVizOptionArgs $) {
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
    public static Builder builder(ListChartVizOptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListChartVizOptionArgs $;

        public Builder() {
            $ = new ListChartVizOptionArgs();
        }

        public Builder(ListChartVizOptionArgs defaults) {
            $ = new ListChartVizOptionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param color Color to use
         * 
         * @return builder
         * 
         */
        public Builder color(@Nullable Output<String> color) {
            $.color = color;
            return this;
        }

        /**
         * @param color Color to use
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
         * @param label The label used in the publish statement that displays the plot (metric time series data) you want to customize
         * 
         * @return builder
         * 
         */
        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label The label used in the publish statement that displays the plot (metric time series data) you want to customize
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param valuePrefix An arbitrary prefix to display with the value of this plot
         * 
         * @return builder
         * 
         */
        public Builder valuePrefix(@Nullable Output<String> valuePrefix) {
            $.valuePrefix = valuePrefix;
            return this;
        }

        /**
         * @param valuePrefix An arbitrary prefix to display with the value of this plot
         * 
         * @return builder
         * 
         */
        public Builder valuePrefix(String valuePrefix) {
            return valuePrefix(Output.of(valuePrefix));
        }

        /**
         * @param valueSuffix An arbitrary suffix to display with the value of this plot
         * 
         * @return builder
         * 
         */
        public Builder valueSuffix(@Nullable Output<String> valueSuffix) {
            $.valueSuffix = valueSuffix;
            return this;
        }

        /**
         * @param valueSuffix An arbitrary suffix to display with the value of this plot
         * 
         * @return builder
         * 
         */
        public Builder valueSuffix(String valueSuffix) {
            return valueSuffix(Output.of(valueSuffix));
        }

        /**
         * @param valueUnit A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
         * 
         * @return builder
         * 
         */
        public Builder valueUnit(@Nullable Output<String> valueUnit) {
            $.valueUnit = valueUnit;
            return this;
        }

        /**
         * @param valueUnit A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
         * 
         * @return builder
         * 
         */
        public Builder valueUnit(String valueUnit) {
            return valueUnit(Output.of(valueUnit));
        }

        public ListChartVizOptionArgs build() {
            if ($.label == null) {
                throw new MissingRequiredPropertyException("ListChartVizOptionArgs", "label");
            }
            return $;
        }
    }

}
