// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.ListChartArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.ListChartState;
import com.pulumi.signalfx.outputs.ListChartColorScale;
import com.pulumi.signalfx.outputs.ListChartLegendOptionsField;
import com.pulumi.signalfx.outputs.ListChartVizOption;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This chart type displays current data values in a list format.
 * 
 * The name of each value in the chart reflects the name of the plot and any associated dimensions. We recommend you click the Pencil icon and give the plot a meaningful name, as in plot B from the example. Otherwise, just the raw metric name will be displayed on the chart, as in plot A from the example.
 * 
 * ## Example
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.ListChart;
 * import com.pulumi.signalfx.ListChartArgs;
 * import com.pulumi.signalfx.inputs.ListChartLegendOptionsFieldArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var mylistchart0 = new ListChart(&#34;mylistchart0&#34;, ListChartArgs.builder()        
 *             .colorBy(&#34;Metric&#34;)
 *             .description(&#34;Very cool List Chart&#34;)
 *             .disableSampling(true)
 *             .hideMissingValues(true)
 *             .legendOptionsFields(            
 *                 ListChartLegendOptionsFieldArgs.builder()
 *                     .enabled(false)
 *                     .property(&#34;collector&#34;)
 *                     .build(),
 *                 ListChartLegendOptionsFieldArgs.builder()
 *                     .enabled(true)
 *                     .property(&#34;cluster_name&#34;)
 *                     .build(),
 *                 ListChartLegendOptionsFieldArgs.builder()
 *                     .enabled(true)
 *                     .property(&#34;role&#34;)
 *                     .build(),
 *                 ListChartLegendOptionsFieldArgs.builder()
 *                     .enabled(false)
 *                     .property(&#34;collector&#34;)
 *                     .build(),
 *                 ListChartLegendOptionsFieldArgs.builder()
 *                     .enabled(false)
 *                     .property(&#34;host&#34;)
 *                     .build())
 *             .maxDelay(2)
 *             .maxPrecision(2)
 *             .programText(&#34;&#34;&#34;
 * myfilters = filter(&#34;cluster_name&#34;, &#34;prod&#34;) and filter(&#34;role&#34;, &#34;search&#34;)
 * data(&#34;cpu.total.idle&#34;, filter=myfilters).publish()
 * 
 *             &#34;&#34;&#34;)
 *             .refreshInterval(1)
 *             .sortBy(&#34;-value&#34;)
 *             .timezone(&#34;Europe/Paris&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Arguments
 * 
 * The following arguments are supported in the resource block:
 * 
 * * `name` - (Required) Name of the chart.
 * * `program_text` - (Required) Signalflow program text for the chart. More info[in the Splunk Observability Cloud docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
 * * `description` - (Optional) Description of the chart.
 * * `unit_prefix` - (Optional) Must be `&#34;Metric&#34;` or `&#34;Binary`&#34;. `&#34;Metric&#34;` by default.
 * * `color_by` - (Optional) Must be one of `&#34;Scale&#34;`, `&#34;Dimension&#34;` or `&#34;Metric&#34;`. `&#34;Dimension&#34;` by default.
 * * `max_delay` - (Optional) How long (in seconds) to wait for late datapoints.
 * * `timezone` - (Optional) The property value is a string that denotes the geographic region associated with the time zone, (default UTC).
 * * `disable_sampling` - (Optional) If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
 * * `refresh_interval` - (Optional) How often (in seconds) to refresh the values of the list.
 * * `hide_missing_values` - (Optional) Determines whether to hide missing data points in the chart. If `true`, missing data points in the chart would be hidden. `false` by default.
 * * `viz_options` - (Optional) Plot-level customization options, associated with a publish statement.
 *     * `label` - (Required) Label used in the publish statement that displays the plot (metric time series data) you want to customize.
 *     * `display_name` - (Optional) Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
 *     * `color` - (Optional) The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
 *     * `value_unit` - (Optional) A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
 *     * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
 * * `legend_fields_to_hide` - (Optional) List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
 * * `legend_options_fields` - (Optional) List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
 *     * `property` The name of the property to display. Note the special values of `sf_metric` (corresponding with the API&#39;s `Plot Name`) which shows the label of the time series `publish()` and `sf_originatingMetric` (corresponding with the API&#39;s `metric (sf metric)`) that shows the [name of the metric](https://dev.splunk.com/observability/docs/signalflow/functions/data_function/) for the time series being displayed.
 *     * `enabled` True or False depending on if you want the property to be shown or hidden.
 * * `max_precision` - (Optional) Maximum number of digits to display when rounding values up or down.
 * * `secondary_visualization` - (Optional) The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the Splunk Observability Cloud default is used (`Sparkline`).
 * * `color_scale` - (Optional. `color_by` must be `&#34;Scale&#34;`) Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = &#34;blue&#34; }, { lte = 60, color = &#34;yellow&#34; }]`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
 *     * `gt` - (Optional) Indicates the lower threshold non-inclusive value for this range.
 *     * `gte` - (Optional) Indicates the lower threshold inclusive value for this range.
 *     * `lt` - (Optional) Indicates the upper threshold non-inculsive value for this range.
 *     * `lte` - (Optional) Indicates the upper threshold inclusive value for this range.
 *     * `color` - (Required) The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
 * * `sort_by` - (Optional) The property to use when sorting the elements. Use `value` if you want to sort by value. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`). Note there are some special values for some of the options provided in the UX: `&#34;value&#34;` for Value, `&#34;sf_originatingMetric&#34;` for Metric, and `&#34;sf_metric&#34;` for plot.
 * * `time_range` - (Optional) How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
 * * `start_time` - (Optional) Seconds since epoch. Used for visualization. Conflicts with `time_range`.
 * * `end_time` - (Optional) Seconds since epoch. Used for visualization. Conflicts with `time_range`.
 * 
 * ## Attributes
 * 
 * In a addition to all arguments above, the following attributes are exported:
 * 
 * * `id` - The ID of the chart.
 * * `url` - The URL of the chart.
 * 
 */
@ResourceType(type="signalfx:index/listChart:ListChart")
public class ListChart extends com.pulumi.resources.CustomResource {
    /**
     * (Metric by default) Must be &#34;Scale&#34;, &#34;Metric&#34; or &#34;Dimension&#34;
     * 
     */
    @Export(name="colorBy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> colorBy;

    /**
     * @return (Metric by default) Must be &#34;Scale&#34;, &#34;Metric&#34; or &#34;Dimension&#34;
     * 
     */
    public Output<Optional<String>> colorBy() {
        return Codegen.optional(this.colorBy);
    }
    /**
     * Single color range including both the color to display for that range and the borders of the range
     * 
     */
    @Export(name="colorScales", refs={List.class,ListChartColorScale.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ListChartColorScale>> colorScales;

    /**
     * @return Single color range including both the color to display for that range and the borders of the range
     * 
     */
    public Output<Optional<List<ListChartColorScale>>> colorScales() {
        return Codegen.optional(this.colorScales);
    }
    /**
     * Description of the chart (Optional)
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the chart (Optional)
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * (false by default) If false, samples a subset of the output MTS, which improves UI performance
     * 
     */
    @Export(name="disableSampling", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableSampling;

    /**
     * @return (false by default) If false, samples a subset of the output MTS, which improves UI performance
     * 
     */
    public Output<Optional<Boolean>> disableSampling() {
        return Codegen.optional(this.disableSampling);
    }
    /**
     * Seconds since epoch to end the visualization
     * 
     */
    @Export(name="endTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> endTime;

    /**
     * @return Seconds since epoch to end the visualization
     * 
     */
    public Output<Optional<Integer>> endTime() {
        return Codegen.optional(this.endTime);
    }
    /**
     * (false by default) If `true`, missing data points in the chart would be hidden
     * 
     */
    @Export(name="hideMissingValues", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hideMissingValues;

    /**
     * @return (false by default) If `true`, missing data points in the chart would be hidden
     * 
     */
    public Output<Optional<Boolean>> hideMissingValues() {
        return Codegen.optional(this.hideMissingValues);
    }
    /**
     * List of properties that shouldn&#39;t be displayed in the chart legend (i.e. dimension names)
     * 
     * @deprecated
     * Please use legend_options_fields
     * 
     */
    @Deprecated /* Please use legend_options_fields */
    @Export(name="legendFieldsToHides", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> legendFieldsToHides;

    /**
     * @return List of properties that shouldn&#39;t be displayed in the chart legend (i.e. dimension names)
     * 
     */
    public Output<Optional<List<String>>> legendFieldsToHides() {
        return Codegen.optional(this.legendFieldsToHides);
    }
    /**
     * List of property and enabled flags to control the order and presence of datatable labels in a chart.
     * 
     */
    @Export(name="legendOptionsFields", refs={List.class,ListChartLegendOptionsField.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ListChartLegendOptionsField>> legendOptionsFields;

    /**
     * @return List of property and enabled flags to control the order and presence of datatable labels in a chart.
     * 
     */
    public Output<Optional<List<ListChartLegendOptionsField>>> legendOptionsFields() {
        return Codegen.optional(this.legendOptionsFields);
    }
    /**
     * How long (in seconds) to wait for late datapoints
     * 
     */
    @Export(name="maxDelay", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxDelay;

    /**
     * @return How long (in seconds) to wait for late datapoints
     * 
     */
    public Output<Optional<Integer>> maxDelay() {
        return Codegen.optional(this.maxDelay);
    }
    /**
     * Maximum number of digits to display when rounding values up or down
     * 
     */
    @Export(name="maxPrecision", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxPrecision;

    /**
     * @return Maximum number of digits to display when rounding values up or down
     * 
     */
    public Output<Optional<Integer>> maxPrecision() {
        return Codegen.optional(this.maxPrecision);
    }
    /**
     * Name of the chart
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the chart
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Signalflow program text for the chart. More info at &#34;https://developers.signalfx.com/docs/signalflow-overview&#34;
     * 
     */
    @Export(name="programText", refs={String.class}, tree="[0]")
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
    @Export(name="refreshInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> refreshInterval;

    /**
     * @return How often (in seconds) to refresh the values of the list
     * 
     */
    public Output<Optional<Integer>> refreshInterval() {
        return Codegen.optional(this.refreshInterval);
    }
    /**
     * (false by default) What kind of secondary visualization to show (None, Radial, Linear, Sparkline)
     * 
     */
    @Export(name="secondaryVisualization", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secondaryVisualization;

    /**
     * @return (false by default) What kind of secondary visualization to show (None, Radial, Linear, Sparkline)
     * 
     */
    public Output<Optional<String>> secondaryVisualization() {
        return Codegen.optional(this.secondaryVisualization);
    }
    /**
     * The property to use when sorting the elements. Use &#39;value&#39; if you want to sort by value. Must be prepended with + for
     * ascending or - for descending (e.g. -foo)
     * 
     */
    @Export(name="sortBy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sortBy;

    /**
     * @return The property to use when sorting the elements. Use &#39;value&#39; if you want to sort by value. Must be prepended with + for
     * ascending or - for descending (e.g. -foo)
     * 
     */
    public Output<Optional<String>> sortBy() {
        return Codegen.optional(this.sortBy);
    }
    /**
     * Seconds since epoch to start the visualization
     * 
     */
    @Export(name="startTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> startTime;

    /**
     * @return Seconds since epoch to start the visualization
     * 
     */
    public Output<Optional<Integer>> startTime() {
        return Codegen.optional(this.startTime);
    }
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
     * 
     */
    @Export(name="timeRange", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeRange;

    /**
     * @return Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`
     * 
     */
    public Output<Optional<Integer>> timeRange() {
        return Codegen.optional(this.timeRange);
    }
    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     * 
     */
    @Export(name="timezone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> timezone;

    /**
     * @return The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     * 
     */
    public Output<Optional<String>> timezone() {
        return Codegen.optional(this.timezone);
    }
    /**
     * (Metric by default) Must be &#34;Metric&#34; or &#34;Binary&#34;
     * 
     */
    @Export(name="unitPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> unitPrefix;

    /**
     * @return (Metric by default) Must be &#34;Metric&#34; or &#34;Binary&#34;
     * 
     */
    public Output<Optional<String>> unitPrefix() {
        return Codegen.optional(this.unitPrefix);
    }
    /**
     * URL of the chart
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return URL of the chart
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * Plot-level customization options, associated with a publish statement
     * 
     */
    @Export(name="vizOptions", refs={List.class,ListChartVizOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ListChartVizOption>> vizOptions;

    /**
     * @return Plot-level customization options, associated with a publish statement
     * 
     */
    public Output<Optional<List<ListChartVizOption>>> vizOptions() {
        return Codegen.optional(this.vizOptions);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ListChart(String name) {
        this(name, ListChartArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ListChart(String name, ListChartArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ListChart(String name, ListChartArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/listChart:ListChart", name, args == null ? ListChartArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ListChart(String name, Output<String> id, @Nullable ListChartState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/listChart:ListChart", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static ListChart get(String name, Output<String> id, @Nullable ListChartState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ListChart(name, id, state, options);
    }
}
