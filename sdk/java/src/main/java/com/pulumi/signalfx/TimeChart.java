// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.TimeChartArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.TimeChartState;
import com.pulumi.signalfx.outputs.TimeChartAxisLeft;
import com.pulumi.signalfx.outputs.TimeChartAxisRight;
import com.pulumi.signalfx.outputs.TimeChartEventOption;
import com.pulumi.signalfx.outputs.TimeChartHistogramOption;
import com.pulumi.signalfx.outputs.TimeChartLegendOptionsField;
import com.pulumi.signalfx.outputs.TimeChartVizOption;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Splunk Observability Cloud time chart resource. This can be used to create and manage the different types of time charts.
 * 
 * Time charts display data points over a period of time.
 * 
 * ## Example
 * 
 */
@ResourceType(type="signalfx:index/timeChart:TimeChart")
public class TimeChart extends com.pulumi.resources.CustomResource {
    /**
     * Force the chart to display zero on the y-axes, even if none of the data is near zero.
     * 
     */
    @Export(name="axesIncludeZero", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> axesIncludeZero;

    /**
     * @return Force the chart to display zero on the y-axes, even if none of the data is near zero.
     * 
     */
    public Output<Optional<Boolean>> axesIncludeZero() {
        return Codegen.optional(this.axesIncludeZero);
    }
    /**
     * Specifies the digits Splunk Observability Cloud displays for values plotted on the chart. Defaults to `3`.
     * 
     */
    @Export(name="axesPrecision", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> axesPrecision;

    /**
     * @return Specifies the digits Splunk Observability Cloud displays for values plotted on the chart. Defaults to `3`.
     * 
     */
    public Output<Optional<Integer>> axesPrecision() {
        return Codegen.optional(this.axesPrecision);
    }
    /**
     * Set of axis options.
     * 
     */
    @Export(name="axisLeft", refs={TimeChartAxisLeft.class}, tree="[0]")
    private Output</* @Nullable */ TimeChartAxisLeft> axisLeft;

    /**
     * @return Set of axis options.
     * 
     */
    public Output<Optional<TimeChartAxisLeft>> axisLeft() {
        return Codegen.optional(this.axisLeft);
    }
    /**
     * Set of axis options.
     * 
     */
    @Export(name="axisRight", refs={TimeChartAxisRight.class}, tree="[0]")
    private Output</* @Nullable */ TimeChartAxisRight> axisRight;

    /**
     * @return Set of axis options.
     * 
     */
    public Output<Optional<TimeChartAxisRight>> axisRight() {
        return Codegen.optional(this.axisRight);
    }
    /**
     * Must be `&#34;Dimension&#34;` or `&#34;Metric&#34;`. `&#34;Dimension&#34;` by default.
     * 
     */
    @Export(name="colorBy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> colorBy;

    /**
     * @return Must be `&#34;Dimension&#34;` or `&#34;Metric&#34;`. `&#34;Dimension&#34;` by default.
     * 
     */
    public Output<Optional<String>> colorBy() {
        return Codegen.optional(this.colorBy);
    }
    /**
     * Description of the chart.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the chart.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
     * 
     */
    @Export(name="disableSampling", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableSampling;

    /**
     * @return If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
     * 
     */
    public Output<Optional<Boolean>> disableSampling() {
        return Codegen.optional(this.disableSampling);
    }
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    @Export(name="endTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> endTime;

    /**
     * @return Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    public Output<Optional<Integer>> endTime() {
        return Codegen.optional(this.endTime);
    }
    /**
     * Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
     * 
     */
    @Export(name="eventOptions", refs={List.class,TimeChartEventOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TimeChartEventOption>> eventOptions;

    /**
     * @return Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
     * 
     */
    public Output<Optional<List<TimeChartEventOption>>> eventOptions() {
        return Codegen.optional(this.eventOptions);
    }
    /**
     * Only used when `plot_type` is `&#34;Histogram&#34;`. Histogram specific options.
     * 
     */
    @Export(name="histogramOptions", refs={List.class,TimeChartHistogramOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TimeChartHistogramOption>> histogramOptions;

    /**
     * @return Only used when `plot_type` is `&#34;Histogram&#34;`. Histogram specific options.
     * 
     */
    public Output<Optional<List<TimeChartHistogramOption>>> histogramOptions() {
        return Codegen.optional(this.histogramOptions);
    }
    /**
     * List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
     * 
     * @deprecated
     * Please use legend_options_fields
     * 
     */
    @Deprecated /* Please use legend_options_fields */
    @Export(name="legendFieldsToHides", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> legendFieldsToHides;

    /**
     * @return List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
     * 
     */
    public Output<Optional<List<String>>> legendFieldsToHides() {
        return Codegen.optional(this.legendFieldsToHides);
    }
    /**
     * List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
     * 
     */
    @Export(name="legendOptionsFields", refs={List.class,TimeChartLegendOptionsField.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TimeChartLegendOptionsField>> legendOptionsFields;

    /**
     * @return List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
     * 
     */
    public Output<Optional<List<TimeChartLegendOptionsField>>> legendOptionsFields() {
        return Codegen.optional(this.legendOptionsFields);
    }
    /**
     * How long (in seconds) to wait for late datapoints.
     * 
     */
    @Export(name="maxDelay", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxDelay;

    /**
     * @return How long (in seconds) to wait for late datapoints.
     * 
     */
    public Output<Optional<Integer>> maxDelay() {
        return Codegen.optional(this.maxDelay);
    }
    /**
     * The minimum resolution (in seconds) to use for computing the underlying program.
     * 
     */
    @Export(name="minimumResolution", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minimumResolution;

    /**
     * @return The minimum resolution (in seconds) to use for computing the underlying program.
     * 
     */
    public Output<Optional<Integer>> minimumResolution() {
        return Codegen.optional(this.minimumResolution);
    }
    /**
     * Name of the chart.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the chart.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `&#34;metric&#34;`, `&#34;plot_label&#34;` and any dimension.
     * 
     */
    @Export(name="onChartLegendDimension", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> onChartLegendDimension;

    /**
     * @return Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `&#34;metric&#34;`, `&#34;plot_label&#34;` and any dimension.
     * 
     */
    public Output<Optional<String>> onChartLegendDimension() {
        return Codegen.optional(this.onChartLegendDimension);
    }
    /**
     * The default plot display style for the visualization. Must be `&#34;LineChart&#34;`, `&#34;AreaChart&#34;`, `&#34;ColumnChart&#34;`, or `&#34;Histogram&#34;`. Default: `&#34;LineChart&#34;`.
     * 
     */
    @Export(name="plotType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> plotType;

    /**
     * @return The default plot display style for the visualization. Must be `&#34;LineChart&#34;`, `&#34;AreaChart&#34;`, `&#34;ColumnChart&#34;`, or `&#34;Histogram&#34;`. Default: `&#34;LineChart&#34;`.
     * 
     */
    public Output<Optional<String>> plotType() {
        return Codegen.optional(this.plotType);
    }
    /**
     * Signalflow program text for the chart. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
     * 
     */
    @Export(name="programText", refs={String.class}, tree="[0]")
    private Output<String> programText;

    /**
     * @return Signalflow program text for the chart. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
     * 
     */
    public Output<String> programText() {
        return this.programText;
    }
    /**
     * Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
     * 
     */
    @Export(name="showDataMarkers", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> showDataMarkers;

    /**
     * @return Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
     * 
     */
    public Output<Optional<Boolean>> showDataMarkers() {
        return Codegen.optional(this.showDataMarkers);
    }
    /**
     * Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
     * 
     */
    @Export(name="showEventLines", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> showEventLines;

    /**
     * @return Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
     * 
     */
    public Output<Optional<Boolean>> showEventLines() {
        return Codegen.optional(this.showEventLines);
    }
    /**
     * Whether area and bar charts in the visualization should be stacked. `false` by default.
     * 
     */
    @Export(name="stacked", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> stacked;

    /**
     * @return Whether area and bar charts in the visualization should be stacked. `false` by default.
     * 
     */
    public Output<Optional<Boolean>> stacked() {
        return Codegen.optional(this.stacked);
    }
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    @Export(name="startTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> startTime;

    /**
     * @return Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    public Output<Optional<Integer>> startTime() {
        return Codegen.optional(this.startTime);
    }
    /**
     * Tags associated with the chart
     * 
     * @deprecated
     * signalfx_time_chart.tags is being removed in the next release
     * 
     */
    @Deprecated /* signalfx_time_chart.tags is being removed in the next release */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return Tags associated with the chart
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
     * 
     */
    @Export(name="timeRange", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeRange;

    /**
     * @return How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
     * 
     */
    public Output<Optional<Integer>> timeRange() {
        return Codegen.optional(this.timeRange);
    }
    /**
     * Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set &#34;timezone&#34;: &#34;Europe/Paris&#34; and then use the transformation sum(cycle=&#34;week&#34;, cycle_start=&#34;Monday&#34;) in your chart&#39;s SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://dev.splunk.com/observability/docs/signalflow/). `&#34;UTC&#34;` by default.
     * 
     */
    @Export(name="timezone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> timezone;

    /**
     * @return Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set &#34;timezone&#34;: &#34;Europe/Paris&#34; and then use the transformation sum(cycle=&#34;week&#34;, cycle_start=&#34;Monday&#34;) in your chart&#39;s SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://dev.splunk.com/observability/docs/signalflow/). `&#34;UTC&#34;` by default.
     * 
     */
    public Output<Optional<String>> timezone() {
        return Codegen.optional(this.timezone);
    }
    /**
     * Must be `&#34;Metric&#34;` or `&#34;Binary`&#34;. `&#34;Metric&#34;` by default.
     * 
     */
    @Export(name="unitPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> unitPrefix;

    /**
     * @return Must be `&#34;Metric&#34;` or `&#34;Binary`&#34;. `&#34;Metric&#34;` by default.
     * 
     */
    public Output<Optional<String>> unitPrefix() {
        return Codegen.optional(this.unitPrefix);
    }
    /**
     * The URL of the chart.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL of the chart.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * Plot-level customization options, associated with a publish statement.
     * 
     */
    @Export(name="vizOptions", refs={List.class,TimeChartVizOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TimeChartVizOption>> vizOptions;

    /**
     * @return Plot-level customization options, associated with a publish statement.
     * 
     */
    public Output<Optional<List<TimeChartVizOption>>> vizOptions() {
        return Codegen.optional(this.vizOptions);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TimeChart(java.lang.String name) {
        this(name, TimeChartArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TimeChart(java.lang.String name, TimeChartArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TimeChart(java.lang.String name, TimeChartArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/timeChart:TimeChart", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TimeChart(java.lang.String name, Output<java.lang.String> id, @Nullable TimeChartState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/timeChart:TimeChart", name, state, makeResourceOptions(options, id), false);
    }

    private static TimeChartArgs makeArgs(TimeChartArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TimeChartArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static TimeChart get(java.lang.String name, Output<java.lang.String> id, @Nullable TimeChartState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TimeChart(name, id, state, options);
    }
}
