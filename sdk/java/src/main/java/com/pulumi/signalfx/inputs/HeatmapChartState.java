// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.inputs.HeatmapChartColorRangeArgs;
import com.pulumi.signalfx.inputs.HeatmapChartColorScaleArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HeatmapChartState extends com.pulumi.resources.ResourceArgs {

    public static final HeatmapChartState Empty = new HeatmapChartState();

    /**
     * Values and color for the color range. Example: `color_range : { min : 0, max : 100, color : &#34;#0000ff&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
     * 
     */
    @Import(name="colorRange")
    private @Nullable Output<HeatmapChartColorRangeArgs> colorRange;

    /**
     * @return Values and color for the color range. Example: `color_range : { min : 0, max : 100, color : &#34;#0000ff&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
     * 
     */
    public Optional<Output<HeatmapChartColorRangeArgs>> colorRange() {
        return Optional.ofNullable(this.colorRange);
    }

    /**
     * One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = &#34;blue&#34; } color_scale { lte = 60, color = &#34;yellow&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
     * 
     */
    @Import(name="colorScales")
    private @Nullable Output<List<HeatmapChartColorScaleArgs>> colorScales;

    /**
     * @return One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = &#34;blue&#34; } color_scale { lte = 60, color = &#34;yellow&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
     * 
     */
    public Optional<Output<List<HeatmapChartColorScaleArgs>>> colorScales() {
        return Optional.ofNullable(this.colorScales);
    }

    /**
     * Description of the chart.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the chart.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
     * 
     */
    @Import(name="disableSampling")
    private @Nullable Output<Boolean> disableSampling;

    /**
     * @return If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
     * 
     */
    public Optional<Output<Boolean>> disableSampling() {
        return Optional.ofNullable(this.disableSampling);
    }

    /**
     * Properties to group by in the heatmap (in nesting order).
     * 
     */
    @Import(name="groupBies")
    private @Nullable Output<List<String>> groupBies;

    /**
     * @return Properties to group by in the heatmap (in nesting order).
     * 
     */
    public Optional<Output<List<String>>> groupBies() {
        return Optional.ofNullable(this.groupBies);
    }

    /**
     * Whether to show the timestamp in the chart. `false` by default.
     * 
     */
    @Import(name="hideTimestamp")
    private @Nullable Output<Boolean> hideTimestamp;

    /**
     * @return Whether to show the timestamp in the chart. `false` by default.
     * 
     */
    public Optional<Output<Boolean>> hideTimestamp() {
        return Optional.ofNullable(this.hideTimestamp);
    }

    /**
     * How long (in seconds) to wait for late datapoints.
     * 
     */
    @Import(name="maxDelay")
    private @Nullable Output<Integer> maxDelay;

    /**
     * @return How long (in seconds) to wait for late datapoints.
     * 
     */
    public Optional<Output<Integer>> maxDelay() {
        return Optional.ofNullable(this.maxDelay);
    }

    /**
     * The minimum resolution (in seconds) to use for computing the underlying program.
     * 
     */
    @Import(name="minimumResolution")
    private @Nullable Output<Integer> minimumResolution;

    /**
     * @return The minimum resolution (in seconds) to use for computing the underlying program.
     * 
     */
    public Optional<Output<Integer>> minimumResolution() {
        return Optional.ofNullable(this.minimumResolution);
    }

    /**
     * Name of the chart.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the chart.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Signalflow program text for the chart. More info at &lt;https://dev.splunk.com/observability/docs/signalflow/&gt;.
     * 
     */
    @Import(name="programText")
    private @Nullable Output<String> programText;

    /**
     * @return Signalflow program text for the chart. More info at &lt;https://dev.splunk.com/observability/docs/signalflow/&gt;.
     * 
     */
    public Optional<Output<String>> programText() {
        return Optional.ofNullable(this.programText);
    }

    /**
     * How often (in seconds) to refresh the values of the heatmap.
     * 
     */
    @Import(name="refreshInterval")
    private @Nullable Output<Integer> refreshInterval;

    /**
     * @return How often (in seconds) to refresh the values of the heatmap.
     * 
     */
    public Optional<Output<Integer>> refreshInterval() {
        return Optional.ofNullable(this.refreshInterval);
    }

    /**
     * The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
     * 
     */
    @Import(name="sortBy")
    private @Nullable Output<String> sortBy;

    /**
     * @return The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
     * 
     */
    public Optional<Output<String>> sortBy() {
        return Optional.ofNullable(this.sortBy);
    }

    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (default UTC).
     * 
     */
    @Import(name="timezone")
    private @Nullable Output<String> timezone;

    /**
     * @return The property value is a string that denotes the geographic region associated with the time zone, (default UTC).
     * 
     */
    public Optional<Output<String>> timezone() {
        return Optional.ofNullable(this.timezone);
    }

    /**
     * Must be `&#34;Metric&#34;` or `&#34;Binary`&#34;. `&#34;Metric&#34;` by default.
     * 
     */
    @Import(name="unitPrefix")
    private @Nullable Output<String> unitPrefix;

    /**
     * @return Must be `&#34;Metric&#34;` or `&#34;Binary`&#34;. `&#34;Metric&#34;` by default.
     * 
     */
    public Optional<Output<String>> unitPrefix() {
        return Optional.ofNullable(this.unitPrefix);
    }

    /**
     * The URL of the chart.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The URL of the chart.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private HeatmapChartState() {}

    private HeatmapChartState(HeatmapChartState $) {
        this.colorRange = $.colorRange;
        this.colorScales = $.colorScales;
        this.description = $.description;
        this.disableSampling = $.disableSampling;
        this.groupBies = $.groupBies;
        this.hideTimestamp = $.hideTimestamp;
        this.maxDelay = $.maxDelay;
        this.minimumResolution = $.minimumResolution;
        this.name = $.name;
        this.programText = $.programText;
        this.refreshInterval = $.refreshInterval;
        this.sortBy = $.sortBy;
        this.timezone = $.timezone;
        this.unitPrefix = $.unitPrefix;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HeatmapChartState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HeatmapChartState $;

        public Builder() {
            $ = new HeatmapChartState();
        }

        public Builder(HeatmapChartState defaults) {
            $ = new HeatmapChartState(Objects.requireNonNull(defaults));
        }

        /**
         * @param colorRange Values and color for the color range. Example: `color_range : { min : 0, max : 100, color : &#34;#0000ff&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
         * 
         * @return builder
         * 
         */
        public Builder colorRange(@Nullable Output<HeatmapChartColorRangeArgs> colorRange) {
            $.colorRange = colorRange;
            return this;
        }

        /**
         * @param colorRange Values and color for the color range. Example: `color_range : { min : 0, max : 100, color : &#34;#0000ff&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
         * 
         * @return builder
         * 
         */
        public Builder colorRange(HeatmapChartColorRangeArgs colorRange) {
            return colorRange(Output.of(colorRange));
        }

        /**
         * @param colorScales One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = &#34;blue&#34; } color_scale { lte = 60, color = &#34;yellow&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
         * 
         * @return builder
         * 
         */
        public Builder colorScales(@Nullable Output<List<HeatmapChartColorScaleArgs>> colorScales) {
            $.colorScales = colorScales;
            return this;
        }

        /**
         * @param colorScales One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = &#34;blue&#34; } color_scale { lte = 60, color = &#34;yellow&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
         * 
         * @return builder
         * 
         */
        public Builder colorScales(List<HeatmapChartColorScaleArgs> colorScales) {
            return colorScales(Output.of(colorScales));
        }

        /**
         * @param colorScales One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = &#34;blue&#34; } color_scale { lte = 60, color = &#34;yellow&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
         * 
         * @return builder
         * 
         */
        public Builder colorScales(HeatmapChartColorScaleArgs... colorScales) {
            return colorScales(List.of(colorScales));
        }

        /**
         * @param description Description of the chart.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the chart.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disableSampling If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder disableSampling(@Nullable Output<Boolean> disableSampling) {
            $.disableSampling = disableSampling;
            return this;
        }

        /**
         * @param disableSampling If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder disableSampling(Boolean disableSampling) {
            return disableSampling(Output.of(disableSampling));
        }

        /**
         * @param groupBies Properties to group by in the heatmap (in nesting order).
         * 
         * @return builder
         * 
         */
        public Builder groupBies(@Nullable Output<List<String>> groupBies) {
            $.groupBies = groupBies;
            return this;
        }

        /**
         * @param groupBies Properties to group by in the heatmap (in nesting order).
         * 
         * @return builder
         * 
         */
        public Builder groupBies(List<String> groupBies) {
            return groupBies(Output.of(groupBies));
        }

        /**
         * @param groupBies Properties to group by in the heatmap (in nesting order).
         * 
         * @return builder
         * 
         */
        public Builder groupBies(String... groupBies) {
            return groupBies(List.of(groupBies));
        }

        /**
         * @param hideTimestamp Whether to show the timestamp in the chart. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder hideTimestamp(@Nullable Output<Boolean> hideTimestamp) {
            $.hideTimestamp = hideTimestamp;
            return this;
        }

        /**
         * @param hideTimestamp Whether to show the timestamp in the chart. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder hideTimestamp(Boolean hideTimestamp) {
            return hideTimestamp(Output.of(hideTimestamp));
        }

        /**
         * @param maxDelay How long (in seconds) to wait for late datapoints.
         * 
         * @return builder
         * 
         */
        public Builder maxDelay(@Nullable Output<Integer> maxDelay) {
            $.maxDelay = maxDelay;
            return this;
        }

        /**
         * @param maxDelay How long (in seconds) to wait for late datapoints.
         * 
         * @return builder
         * 
         */
        public Builder maxDelay(Integer maxDelay) {
            return maxDelay(Output.of(maxDelay));
        }

        /**
         * @param minimumResolution The minimum resolution (in seconds) to use for computing the underlying program.
         * 
         * @return builder
         * 
         */
        public Builder minimumResolution(@Nullable Output<Integer> minimumResolution) {
            $.minimumResolution = minimumResolution;
            return this;
        }

        /**
         * @param minimumResolution The minimum resolution (in seconds) to use for computing the underlying program.
         * 
         * @return builder
         * 
         */
        public Builder minimumResolution(Integer minimumResolution) {
            return minimumResolution(Output.of(minimumResolution));
        }

        /**
         * @param name Name of the chart.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the chart.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param programText Signalflow program text for the chart. More info at &lt;https://dev.splunk.com/observability/docs/signalflow/&gt;.
         * 
         * @return builder
         * 
         */
        public Builder programText(@Nullable Output<String> programText) {
            $.programText = programText;
            return this;
        }

        /**
         * @param programText Signalflow program text for the chart. More info at &lt;https://dev.splunk.com/observability/docs/signalflow/&gt;.
         * 
         * @return builder
         * 
         */
        public Builder programText(String programText) {
            return programText(Output.of(programText));
        }

        /**
         * @param refreshInterval How often (in seconds) to refresh the values of the heatmap.
         * 
         * @return builder
         * 
         */
        public Builder refreshInterval(@Nullable Output<Integer> refreshInterval) {
            $.refreshInterval = refreshInterval;
            return this;
        }

        /**
         * @param refreshInterval How often (in seconds) to refresh the values of the heatmap.
         * 
         * @return builder
         * 
         */
        public Builder refreshInterval(Integer refreshInterval) {
            return refreshInterval(Output.of(refreshInterval));
        }

        /**
         * @param sortBy The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
         * 
         * @return builder
         * 
         */
        public Builder sortBy(@Nullable Output<String> sortBy) {
            $.sortBy = sortBy;
            return this;
        }

        /**
         * @param sortBy The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
         * 
         * @return builder
         * 
         */
        public Builder sortBy(String sortBy) {
            return sortBy(Output.of(sortBy));
        }

        /**
         * @param timezone The property value is a string that denotes the geographic region associated with the time zone, (default UTC).
         * 
         * @return builder
         * 
         */
        public Builder timezone(@Nullable Output<String> timezone) {
            $.timezone = timezone;
            return this;
        }

        /**
         * @param timezone The property value is a string that denotes the geographic region associated with the time zone, (default UTC).
         * 
         * @return builder
         * 
         */
        public Builder timezone(String timezone) {
            return timezone(Output.of(timezone));
        }

        /**
         * @param unitPrefix Must be `&#34;Metric&#34;` or `&#34;Binary`&#34;. `&#34;Metric&#34;` by default.
         * 
         * @return builder
         * 
         */
        public Builder unitPrefix(@Nullable Output<String> unitPrefix) {
            $.unitPrefix = unitPrefix;
            return this;
        }

        /**
         * @param unitPrefix Must be `&#34;Metric&#34;` or `&#34;Binary`&#34;. `&#34;Metric&#34;` by default.
         * 
         * @return builder
         * 
         */
        public Builder unitPrefix(String unitPrefix) {
            return unitPrefix(Output.of(unitPrefix));
        }

        /**
         * @param url The URL of the chart.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL of the chart.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public HeatmapChartState build() {
            return $;
        }
    }

}
