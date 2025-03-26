// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.SingleValueChartArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.SingleValueChartState;
import com.pulumi.signalfx.outputs.SingleValueChartColorScale;
import com.pulumi.signalfx.outputs.SingleValueChartVizOption;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This chart type displays a single number in a large font, representing the current value of a single metric on a plot line.
 * 
 * If the time period is in the past, the number represents the value of the metric near the end of the time period.
 * 
 * ## Example
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.SingleValueChart;
 * import com.pulumi.signalfx.SingleValueChartArgs;
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
 *         var mysvchart0 = new SingleValueChart("mysvchart0", SingleValueChartArgs.builder()
 *             .name("CPU Total Idle - Single Value")
 *             .programText("""
 * myfilters = filter("cluster_name", "prod") and filter("role", "search")
 * data("cpu.total.idle", filter=myfilters).publish()
 *             """)
 *             .description("Very cool Single Value Chart")
 *             .colorBy("Dimension")
 *             .maxDelay(2)
 *             .refreshInterval(1)
 *             .maxPrecision(2)
 *             .isTimestampHidden(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="signalfx:index/singleValueChart:SingleValueChart")
public class SingleValueChart extends com.pulumi.resources.CustomResource {
    /**
     * Must be `&#34;Dimension&#34;`, `&#34;Scale&#34;` or `&#34;Metric&#34;`. `&#34;Dimension&#34;` by default.
     * 
     */
    @Export(name="colorBy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> colorBy;

    /**
     * @return Must be `&#34;Dimension&#34;`, `&#34;Scale&#34;` or `&#34;Metric&#34;`. `&#34;Dimension&#34;` by default.
     * 
     */
    public Output<Optional<String>> colorBy() {
        return Codegen.optional(this.colorBy);
    }
    /**
     * Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = &#34;blue&#34; }, { lte = 60, color = &#34;yellow&#34; }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
     * 
     */
    @Export(name="colorScales", refs={List.class,SingleValueChartColorScale.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SingleValueChartColorScale>> colorScales;

    /**
     * @return Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = &#34;blue&#34; }, { lte = 60, color = &#34;yellow&#34; }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
     * 
     */
    public Output<Optional<List<SingleValueChartColorScale>>> colorScales() {
        return Codegen.optional(this.colorScales);
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
     * Whether to hide the timestamp in the chart. `false` by default.
     * 
     */
    @Export(name="isTimestampHidden", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isTimestampHidden;

    /**
     * @return Whether to hide the timestamp in the chart. `false` by default.
     * 
     */
    public Output<Optional<Boolean>> isTimestampHidden() {
        return Codegen.optional(this.isTimestampHidden);
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
     * The maximum precision to for value displayed.
     * 
     */
    @Export(name="maxPrecision", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxPrecision;

    /**
     * @return The maximum precision to for value displayed.
     * 
     */
    public Output<Optional<Integer>> maxPrecision() {
        return Codegen.optional(this.maxPrecision);
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
     * How often (in seconds) to refresh the value.
     * 
     */
    @Export(name="refreshInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> refreshInterval;

    /**
     * @return How often (in seconds) to refresh the value.
     * 
     */
    public Output<Optional<Integer>> refreshInterval() {
        return Codegen.optional(this.refreshInterval);
    }
    /**
     * The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the Splunk Observability Cloud default is used (`None`).
     * 
     */
    @Export(name="secondaryVisualization", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> secondaryVisualization;

    /**
     * @return The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the Splunk Observability Cloud default is used (`None`).
     * 
     */
    public Output<Optional<String>> secondaryVisualization() {
        return Codegen.optional(this.secondaryVisualization);
    }
    /**
     * Whether to show a trend line below the current value. `false` by default.
     * 
     */
    @Export(name="showSparkLine", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> showSparkLine;

    /**
     * @return Whether to show a trend line below the current value. `false` by default.
     * 
     */
    public Output<Optional<Boolean>> showSparkLine() {
        return Codegen.optional(this.showSparkLine);
    }
    /**
     * Tags associated with the resource
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return Tags associated with the resource
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
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
     * Must be `&#34;Metric&#34;` or `&#34;Binary&#34;`. `&#34;Metric&#34;` by default.
     * 
     */
    @Export(name="unitPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> unitPrefix;

    /**
     * @return Must be `&#34;Metric&#34;` or `&#34;Binary&#34;`. `&#34;Metric&#34;` by default.
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
    @Export(name="vizOptions", refs={List.class,SingleValueChartVizOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<SingleValueChartVizOption>> vizOptions;

    /**
     * @return Plot-level customization options, associated with a publish statement.
     * 
     */
    public Output<Optional<List<SingleValueChartVizOption>>> vizOptions() {
        return Codegen.optional(this.vizOptions);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SingleValueChart(java.lang.String name) {
        this(name, SingleValueChartArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SingleValueChart(java.lang.String name, SingleValueChartArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SingleValueChart(java.lang.String name, SingleValueChartArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/singleValueChart:SingleValueChart", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SingleValueChart(java.lang.String name, Output<java.lang.String> id, @Nullable SingleValueChartState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/singleValueChart:SingleValueChart", name, state, makeResourceOptions(options, id), false);
    }

    private static SingleValueChartArgs makeArgs(SingleValueChartArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SingleValueChartArgs.Empty : args;
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
    public static SingleValueChart get(java.lang.String name, Output<java.lang.String> id, @Nullable SingleValueChartState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SingleValueChart(name, id, state, options);
    }
}
