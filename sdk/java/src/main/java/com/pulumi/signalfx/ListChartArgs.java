// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.inputs.ListChartColorScaleArgs;
import com.pulumi.signalfx.inputs.ListChartLegendOptionsFieldArgs;
import com.pulumi.signalfx.inputs.ListChartVizOptionArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ListChartArgs extends com.pulumi.resources.ResourceArgs {

    public static final ListChartArgs Empty = new ListChartArgs();

    /**
     * (Metric by default) Must be &#34;Scale&#34;, &#34;Metric&#34; or &#34;Dimension&#34;
     * 
     */
    @Import(name="colorBy")
    private @Nullable Output<String> colorBy;

    /**
     * @return (Metric by default) Must be &#34;Scale&#34;, &#34;Metric&#34; or &#34;Dimension&#34;
     * 
     */
    public Optional<Output<String>> colorBy() {
        return Optional.ofNullable(this.colorBy);
    }

    /**
     * Single color range including both the color to display for that range and the borders of the range
     * 
     */
    @Import(name="colorScales")
    private @Nullable Output<List<ListChartColorScaleArgs>> colorScales;

    /**
     * @return Single color range including both the color to display for that range and the borders of the range
     * 
     */
    public Optional<Output<List<ListChartColorScaleArgs>>> colorScales() {
        return Optional.ofNullable(this.colorScales);
    }

    /**
     * Description of the chart (Optional)
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the chart (Optional)
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * (false by default) If false, samples a subset of the output MTS, which improves UI performance
     * 
     */
    @Import(name="disableSampling")
    private @Nullable Output<Boolean> disableSampling;

    /**
     * @return (false by default) If false, samples a subset of the output MTS, which improves UI performance
     * 
     */
    public Optional<Output<Boolean>> disableSampling() {
        return Optional.ofNullable(this.disableSampling);
    }

    /**
     * Seconds since epoch to end the visualization
     * 
     */
    @Import(name="endTime")
    private @Nullable Output<Integer> endTime;

    /**
     * @return Seconds since epoch to end the visualization
     * 
     */
    public Optional<Output<Integer>> endTime() {
        return Optional.ofNullable(this.endTime);
    }

    /**
     * (false by default) If `true`, missing data points in the chart would be hidden
     * 
     */
    @Import(name="hideMissingValues")
    private @Nullable Output<Boolean> hideMissingValues;

    /**
     * @return (false by default) If `true`, missing data points in the chart would be hidden
     * 
     */
    public Optional<Output<Boolean>> hideMissingValues() {
        return Optional.ofNullable(this.hideMissingValues);
    }

    /**
     * List of properties that shouldn&#39;t be displayed in the chart legend (i.e. dimension names)
     * 
     * @deprecated
     * Please use legend_options_fields
     * 
     */
    @Deprecated /* Please use legend_options_fields */
    @Import(name="legendFieldsToHides")
    private @Nullable Output<List<String>> legendFieldsToHides;

    /**
     * @return List of properties that shouldn&#39;t be displayed in the chart legend (i.e. dimension names)
     * 
     * @deprecated
     * Please use legend_options_fields
     * 
     */
    @Deprecated /* Please use legend_options_fields */
    public Optional<Output<List<String>>> legendFieldsToHides() {
        return Optional.ofNullable(this.legendFieldsToHides);
    }

    /**
     * List of property and enabled flags to control the order and presence of datatable labels in a chart.
     * 
     */
    @Import(name="legendOptionsFields")
    private @Nullable Output<List<ListChartLegendOptionsFieldArgs>> legendOptionsFields;

    /**
     * @return List of property and enabled flags to control the order and presence of datatable labels in a chart.
     * 
     */
    public Optional<Output<List<ListChartLegendOptionsFieldArgs>>> legendOptionsFields() {
        return Optional.ofNullable(this.legendOptionsFields);
    }

    /**
     * How long (in seconds) to wait for late datapoints
     * 
     */
    @Import(name="maxDelay")
    private @Nullable Output<Integer> maxDelay;

    /**
     * @return How long (in seconds) to wait for late datapoints
     * 
     */
    public Optional<Output<Integer>> maxDelay() {
        return Optional.ofNullable(this.maxDelay);
    }

    /**
     * Maximum number of digits to display when rounding values up or down
     * 
     */
    @Import(name="maxPrecision")
    private @Nullable Output<Integer> maxPrecision;

    /**
     * @return Maximum number of digits to display when rounding values up or down
     * 
     */
    public Optional<Output<Integer>> maxPrecision() {
        return Optional.ofNullable(this.maxPrecision);
    }

    /**
     * Name of the chart
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the chart
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Signalflow program text for the chart. More info at &#34;https://developers.signalfx.com/docs/signalflow-overview&#34;
     * 
     */
    @Import(name="programText", required=true)
    private Output<String> programText;

    /**
     * @return Signalflow program text for the chart. More info at &#34;https://developers.signalfx.com/docs/signalflow-overview&#34;
     * 
     */
    public Output<String> programText() {
        return this.programText;
    }

    /**
     * How often (in seconds) to refresh the values of the list
     * 
     */
    @Import(name="refreshInterval")
    private @Nullable Output<Integer> refreshInterval;

    /**
     * @return How often (in seconds) to refresh the values of the list
     * 
     */
    public Optional<Output<Integer>> refreshInterval() {
        return Optional.ofNullable(this.refreshInterval);
    }

    /**
     * (false by default) What kind of secondary visualization to show (None, Radial, Linear, Sparkline)
     * 
     */
    @Import(name="secondaryVisualization")
    private @Nullable Output<String> secondaryVisualization;

    /**
     * @return (false by default) What kind of secondary visualization to show (None, Radial, Linear, Sparkline)
     * 
     */
    public Optional<Output<String>> secondaryVisualization() {
        return Optional.ofNullable(this.secondaryVisualization);
    }

    /**
     * The property to use when sorting the elements. Use &#39;value&#39; if you want to sort by value. Must be prepended with + for
     * ascending or - for descending (e.g. -foo)
     * 
     */
    @Import(name="sortBy")
    private @Nullable Output<String> sortBy;

    /**
     * @return The property to use when sorting the elements. Use &#39;value&#39; if you want to sort by value. Must be prepended with + for
     * ascending or - for descending (e.g. -foo)
     * 
     */
    public Optional<Output<String>> sortBy() {
        return Optional.ofNullable(this.sortBy);
    }

    /**
     * Seconds since epoch to start the visualization
     * 
     */
    @Import(name="startTime")
    private @Nullable Output<Integer> startTime;

    /**
     * @return Seconds since epoch to start the visualization
     * 
     */
    public Optional<Output<Integer>> startTime() {
        return Optional.ofNullable(this.startTime);
    }

    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
     * 
     */
    @Import(name="timeRange")
    private @Nullable Output<Integer> timeRange;

    /**
     * @return Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
     * 
     */
    public Optional<Output<Integer>> timeRange() {
        return Optional.ofNullable(this.timeRange);
    }

    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     * 
     */
    @Import(name="timezone")
    private @Nullable Output<String> timezone;

    /**
     * @return The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     * 
     */
    public Optional<Output<String>> timezone() {
        return Optional.ofNullable(this.timezone);
    }

    /**
     * (Metric by default) Must be &#34;Metric&#34; or &#34;Binary&#34;
     * 
     */
    @Import(name="unitPrefix")
    private @Nullable Output<String> unitPrefix;

    /**
     * @return (Metric by default) Must be &#34;Metric&#34; or &#34;Binary&#34;
     * 
     */
    public Optional<Output<String>> unitPrefix() {
        return Optional.ofNullable(this.unitPrefix);
    }

    /**
     * Plot-level customization options, associated with a publish statement
     * 
     */
    @Import(name="vizOptions")
    private @Nullable Output<List<ListChartVizOptionArgs>> vizOptions;

    /**
     * @return Plot-level customization options, associated with a publish statement
     * 
     */
    public Optional<Output<List<ListChartVizOptionArgs>>> vizOptions() {
        return Optional.ofNullable(this.vizOptions);
    }

    private ListChartArgs() {}

    private ListChartArgs(ListChartArgs $) {
        this.colorBy = $.colorBy;
        this.colorScales = $.colorScales;
        this.description = $.description;
        this.disableSampling = $.disableSampling;
        this.endTime = $.endTime;
        this.hideMissingValues = $.hideMissingValues;
        this.legendFieldsToHides = $.legendFieldsToHides;
        this.legendOptionsFields = $.legendOptionsFields;
        this.maxDelay = $.maxDelay;
        this.maxPrecision = $.maxPrecision;
        this.name = $.name;
        this.programText = $.programText;
        this.refreshInterval = $.refreshInterval;
        this.secondaryVisualization = $.secondaryVisualization;
        this.sortBy = $.sortBy;
        this.startTime = $.startTime;
        this.timeRange = $.timeRange;
        this.timezone = $.timezone;
        this.unitPrefix = $.unitPrefix;
        this.vizOptions = $.vizOptions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ListChartArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListChartArgs $;

        public Builder() {
            $ = new ListChartArgs();
        }

        public Builder(ListChartArgs defaults) {
            $ = new ListChartArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param colorBy (Metric by default) Must be &#34;Scale&#34;, &#34;Metric&#34; or &#34;Dimension&#34;
         * 
         * @return builder
         * 
         */
        public Builder colorBy(@Nullable Output<String> colorBy) {
            $.colorBy = colorBy;
            return this;
        }

        /**
         * @param colorBy (Metric by default) Must be &#34;Scale&#34;, &#34;Metric&#34; or &#34;Dimension&#34;
         * 
         * @return builder
         * 
         */
        public Builder colorBy(String colorBy) {
            return colorBy(Output.of(colorBy));
        }

        /**
         * @param colorScales Single color range including both the color to display for that range and the borders of the range
         * 
         * @return builder
         * 
         */
        public Builder colorScales(@Nullable Output<List<ListChartColorScaleArgs>> colorScales) {
            $.colorScales = colorScales;
            return this;
        }

        /**
         * @param colorScales Single color range including both the color to display for that range and the borders of the range
         * 
         * @return builder
         * 
         */
        public Builder colorScales(List<ListChartColorScaleArgs> colorScales) {
            return colorScales(Output.of(colorScales));
        }

        /**
         * @param colorScales Single color range including both the color to display for that range and the borders of the range
         * 
         * @return builder
         * 
         */
        public Builder colorScales(ListChartColorScaleArgs... colorScales) {
            return colorScales(List.of(colorScales));
        }

        /**
         * @param description Description of the chart (Optional)
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the chart (Optional)
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disableSampling (false by default) If false, samples a subset of the output MTS, which improves UI performance
         * 
         * @return builder
         * 
         */
        public Builder disableSampling(@Nullable Output<Boolean> disableSampling) {
            $.disableSampling = disableSampling;
            return this;
        }

        /**
         * @param disableSampling (false by default) If false, samples a subset of the output MTS, which improves UI performance
         * 
         * @return builder
         * 
         */
        public Builder disableSampling(Boolean disableSampling) {
            return disableSampling(Output.of(disableSampling));
        }

        /**
         * @param endTime Seconds since epoch to end the visualization
         * 
         * @return builder
         * 
         */
        public Builder endTime(@Nullable Output<Integer> endTime) {
            $.endTime = endTime;
            return this;
        }

        /**
         * @param endTime Seconds since epoch to end the visualization
         * 
         * @return builder
         * 
         */
        public Builder endTime(Integer endTime) {
            return endTime(Output.of(endTime));
        }

        /**
         * @param hideMissingValues (false by default) If `true`, missing data points in the chart would be hidden
         * 
         * @return builder
         * 
         */
        public Builder hideMissingValues(@Nullable Output<Boolean> hideMissingValues) {
            $.hideMissingValues = hideMissingValues;
            return this;
        }

        /**
         * @param hideMissingValues (false by default) If `true`, missing data points in the chart would be hidden
         * 
         * @return builder
         * 
         */
        public Builder hideMissingValues(Boolean hideMissingValues) {
            return hideMissingValues(Output.of(hideMissingValues));
        }

        /**
         * @param legendFieldsToHides List of properties that shouldn&#39;t be displayed in the chart legend (i.e. dimension names)
         * 
         * @return builder
         * 
         * @deprecated
         * Please use legend_options_fields
         * 
         */
        @Deprecated /* Please use legend_options_fields */
        public Builder legendFieldsToHides(@Nullable Output<List<String>> legendFieldsToHides) {
            $.legendFieldsToHides = legendFieldsToHides;
            return this;
        }

        /**
         * @param legendFieldsToHides List of properties that shouldn&#39;t be displayed in the chart legend (i.e. dimension names)
         * 
         * @return builder
         * 
         * @deprecated
         * Please use legend_options_fields
         * 
         */
        @Deprecated /* Please use legend_options_fields */
        public Builder legendFieldsToHides(List<String> legendFieldsToHides) {
            return legendFieldsToHides(Output.of(legendFieldsToHides));
        }

        /**
         * @param legendFieldsToHides List of properties that shouldn&#39;t be displayed in the chart legend (i.e. dimension names)
         * 
         * @return builder
         * 
         * @deprecated
         * Please use legend_options_fields
         * 
         */
        @Deprecated /* Please use legend_options_fields */
        public Builder legendFieldsToHides(String... legendFieldsToHides) {
            return legendFieldsToHides(List.of(legendFieldsToHides));
        }

        /**
         * @param legendOptionsFields List of property and enabled flags to control the order and presence of datatable labels in a chart.
         * 
         * @return builder
         * 
         */
        public Builder legendOptionsFields(@Nullable Output<List<ListChartLegendOptionsFieldArgs>> legendOptionsFields) {
            $.legendOptionsFields = legendOptionsFields;
            return this;
        }

        /**
         * @param legendOptionsFields List of property and enabled flags to control the order and presence of datatable labels in a chart.
         * 
         * @return builder
         * 
         */
        public Builder legendOptionsFields(List<ListChartLegendOptionsFieldArgs> legendOptionsFields) {
            return legendOptionsFields(Output.of(legendOptionsFields));
        }

        /**
         * @param legendOptionsFields List of property and enabled flags to control the order and presence of datatable labels in a chart.
         * 
         * @return builder
         * 
         */
        public Builder legendOptionsFields(ListChartLegendOptionsFieldArgs... legendOptionsFields) {
            return legendOptionsFields(List.of(legendOptionsFields));
        }

        /**
         * @param maxDelay How long (in seconds) to wait for late datapoints
         * 
         * @return builder
         * 
         */
        public Builder maxDelay(@Nullable Output<Integer> maxDelay) {
            $.maxDelay = maxDelay;
            return this;
        }

        /**
         * @param maxDelay How long (in seconds) to wait for late datapoints
         * 
         * @return builder
         * 
         */
        public Builder maxDelay(Integer maxDelay) {
            return maxDelay(Output.of(maxDelay));
        }

        /**
         * @param maxPrecision Maximum number of digits to display when rounding values up or down
         * 
         * @return builder
         * 
         */
        public Builder maxPrecision(@Nullable Output<Integer> maxPrecision) {
            $.maxPrecision = maxPrecision;
            return this;
        }

        /**
         * @param maxPrecision Maximum number of digits to display when rounding values up or down
         * 
         * @return builder
         * 
         */
        public Builder maxPrecision(Integer maxPrecision) {
            return maxPrecision(Output.of(maxPrecision));
        }

        /**
         * @param name Name of the chart
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the chart
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param programText Signalflow program text for the chart. More info at &#34;https://developers.signalfx.com/docs/signalflow-overview&#34;
         * 
         * @return builder
         * 
         */
        public Builder programText(Output<String> programText) {
            $.programText = programText;
            return this;
        }

        /**
         * @param programText Signalflow program text for the chart. More info at &#34;https://developers.signalfx.com/docs/signalflow-overview&#34;
         * 
         * @return builder
         * 
         */
        public Builder programText(String programText) {
            return programText(Output.of(programText));
        }

        /**
         * @param refreshInterval How often (in seconds) to refresh the values of the list
         * 
         * @return builder
         * 
         */
        public Builder refreshInterval(@Nullable Output<Integer> refreshInterval) {
            $.refreshInterval = refreshInterval;
            return this;
        }

        /**
         * @param refreshInterval How often (in seconds) to refresh the values of the list
         * 
         * @return builder
         * 
         */
        public Builder refreshInterval(Integer refreshInterval) {
            return refreshInterval(Output.of(refreshInterval));
        }

        /**
         * @param secondaryVisualization (false by default) What kind of secondary visualization to show (None, Radial, Linear, Sparkline)
         * 
         * @return builder
         * 
         */
        public Builder secondaryVisualization(@Nullable Output<String> secondaryVisualization) {
            $.secondaryVisualization = secondaryVisualization;
            return this;
        }

        /**
         * @param secondaryVisualization (false by default) What kind of secondary visualization to show (None, Radial, Linear, Sparkline)
         * 
         * @return builder
         * 
         */
        public Builder secondaryVisualization(String secondaryVisualization) {
            return secondaryVisualization(Output.of(secondaryVisualization));
        }

        /**
         * @param sortBy The property to use when sorting the elements. Use &#39;value&#39; if you want to sort by value. Must be prepended with + for
         * ascending or - for descending (e.g. -foo)
         * 
         * @return builder
         * 
         */
        public Builder sortBy(@Nullable Output<String> sortBy) {
            $.sortBy = sortBy;
            return this;
        }

        /**
         * @param sortBy The property to use when sorting the elements. Use &#39;value&#39; if you want to sort by value. Must be prepended with + for
         * ascending or - for descending (e.g. -foo)
         * 
         * @return builder
         * 
         */
        public Builder sortBy(String sortBy) {
            return sortBy(Output.of(sortBy));
        }

        /**
         * @param startTime Seconds since epoch to start the visualization
         * 
         * @return builder
         * 
         */
        public Builder startTime(@Nullable Output<Integer> startTime) {
            $.startTime = startTime;
            return this;
        }

        /**
         * @param startTime Seconds since epoch to start the visualization
         * 
         * @return builder
         * 
         */
        public Builder startTime(Integer startTime) {
            return startTime(Output.of(startTime));
        }

        /**
         * @param timeRange Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
         * 
         * @return builder
         * 
         */
        public Builder timeRange(@Nullable Output<Integer> timeRange) {
            $.timeRange = timeRange;
            return this;
        }

        /**
         * @param timeRange Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
         * 
         * @return builder
         * 
         */
        public Builder timeRange(Integer timeRange) {
            return timeRange(Output.of(timeRange));
        }

        /**
         * @param timezone The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
         * 
         * @return builder
         * 
         */
        public Builder timezone(@Nullable Output<String> timezone) {
            $.timezone = timezone;
            return this;
        }

        /**
         * @param timezone The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
         * 
         * @return builder
         * 
         */
        public Builder timezone(String timezone) {
            return timezone(Output.of(timezone));
        }

        /**
         * @param unitPrefix (Metric by default) Must be &#34;Metric&#34; or &#34;Binary&#34;
         * 
         * @return builder
         * 
         */
        public Builder unitPrefix(@Nullable Output<String> unitPrefix) {
            $.unitPrefix = unitPrefix;
            return this;
        }

        /**
         * @param unitPrefix (Metric by default) Must be &#34;Metric&#34; or &#34;Binary&#34;
         * 
         * @return builder
         * 
         */
        public Builder unitPrefix(String unitPrefix) {
            return unitPrefix(Output.of(unitPrefix));
        }

        /**
         * @param vizOptions Plot-level customization options, associated with a publish statement
         * 
         * @return builder
         * 
         */
        public Builder vizOptions(@Nullable Output<List<ListChartVizOptionArgs>> vizOptions) {
            $.vizOptions = vizOptions;
            return this;
        }

        /**
         * @param vizOptions Plot-level customization options, associated with a publish statement
         * 
         * @return builder
         * 
         */
        public Builder vizOptions(List<ListChartVizOptionArgs> vizOptions) {
            return vizOptions(Output.of(vizOptions));
        }

        /**
         * @param vizOptions Plot-level customization options, associated with a publish statement
         * 
         * @return builder
         * 
         */
        public Builder vizOptions(ListChartVizOptionArgs... vizOptions) {
            return vizOptions(List.of(vizOptions));
        }

        public ListChartArgs build() {
            if ($.programText == null) {
                throw new MissingRequiredPropertyException("ListChartArgs", "programText");
            }
            return $;
        }
    }

}
