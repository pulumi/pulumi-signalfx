// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.HeatmapChartArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.HeatmapChartState;
import com.pulumi.signalfx.outputs.HeatmapChartColorRange;
import com.pulumi.signalfx.outputs.HeatmapChartColorScale;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This chart type shows the specified plot in a heat map fashion. This format is similar to the [Infrastructure Navigator](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/built-in-content/infra-nav.html#infra), with squares representing each source for the selected metric, and the color of each square representing the value range of the metric.
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
 * import com.pulumi.signalfx.HeatmapChart;
 * import com.pulumi.signalfx.HeatmapChartArgs;
 * import com.pulumi.signalfx.inputs.HeatmapChartColorRangeArgs;
 * import com.pulumi.signalfx.inputs.HeatmapChartColorScaleArgs;
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
 *         var myheatmapchart0 = new HeatmapChart("myheatmapchart0", HeatmapChartArgs.builder()
 *             .name("CPU Total Idle - Heatmap")
 *             .programText("""
 * myfilters = filter("cluster_name", "prod") and filter("role", "search")
 * data("cpu.total.idle", filter=myfilters).publish()
 *             """)
 *             .description("Very cool Heatmap")
 *             .disableSampling(true)
 *             .sortBy("+host")
 *             .groupBies(            
 *                 "hostname",
 *                 "host")
 *             .hideTimestamp(true)
 *             .timezone("Europe/Paris")
 *             .colorRange(HeatmapChartColorRangeArgs.builder()
 *                 .minValue(0.0)
 *                 .maxValue(100.0)
 *                 .color("#ff0000")
 *                 .build())
 *             .colorScales(            
 *                 HeatmapChartColorScaleArgs.builder()
 *                     .gte(99.0)
 *                     .color("green")
 *                     .build(),
 *                 HeatmapChartColorScaleArgs.builder()
 *                     .lt(99.0)
 *                     .gte(95.0)
 *                     .color("yellow")
 *                     .build(),
 *                 HeatmapChartColorScaleArgs.builder()
 *                     .lt(95.0)
 *                     .color("red")
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="signalfx:index/heatmapChart:HeatmapChart")
public class HeatmapChart extends com.pulumi.resources.CustomResource {
    /**
     * Values and color for the color range. Example: `color_range : { min : 0, max : 100, color : &#34;#0000ff&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
     * 
     */
    @Export(name="colorRange", refs={HeatmapChartColorRange.class}, tree="[0]")
    private Output</* @Nullable */ HeatmapChartColorRange> colorRange;

    /**
     * @return Values and color for the color range. Example: `color_range : { min : 0, max : 100, color : &#34;#0000ff&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
     * 
     */
    public Output<Optional<HeatmapChartColorRange>> colorRange() {
        return Codegen.optional(this.colorRange);
    }
    /**
     * One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = &#34;blue&#34; } color_scale { lte = 60, color = &#34;yellow&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
     * 
     */
    @Export(name="colorScales", refs={List.class,HeatmapChartColorScale.class}, tree="[0,1]")
    private Output</* @Nullable */ List<HeatmapChartColorScale>> colorScales;

    /**
     * @return One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = &#34;blue&#34; } color_scale { lte = 60, color = &#34;yellow&#34; }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
     * 
     */
    public Output<Optional<List<HeatmapChartColorScale>>> colorScales() {
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
     * If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
     * 
     */
    @Export(name="disableSampling", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableSampling;

    /**
     * @return If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
     * 
     */
    public Output<Optional<Boolean>> disableSampling() {
        return Codegen.optional(this.disableSampling);
    }
    /**
     * Properties to group by in the heatmap (in nesting order).
     * 
     */
    @Export(name="groupBies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> groupBies;

    /**
     * @return Properties to group by in the heatmap (in nesting order).
     * 
     */
    public Output<Optional<List<String>>> groupBies() {
        return Codegen.optional(this.groupBies);
    }
    /**
     * Whether to show the timestamp in the chart. `false` by default.
     * 
     */
    @Export(name="hideTimestamp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hideTimestamp;

    /**
     * @return Whether to show the timestamp in the chart. `false` by default.
     * 
     */
    public Output<Optional<Boolean>> hideTimestamp() {
        return Codegen.optional(this.hideTimestamp);
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
     * Signalflow program text for the chart. More info at &lt;https://dev.splunk.com/observability/docs/signalflow/&gt;.
     * 
     */
    @Export(name="programText", refs={String.class}, tree="[0]")
    private Output<String> programText;

    /**
     * @return Signalflow program text for the chart. More info at &lt;https://dev.splunk.com/observability/docs/signalflow/&gt;.
     * 
     */
    public Output<String> programText() {
        return this.programText;
    }
    /**
     * How often (in seconds) to refresh the values of the heatmap.
     * 
     */
    @Export(name="refreshInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> refreshInterval;

    /**
     * @return How often (in seconds) to refresh the values of the heatmap.
     * 
     */
    public Output<Optional<Integer>> refreshInterval() {
        return Codegen.optional(this.refreshInterval);
    }
    /**
     * The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
     * 
     */
    @Export(name="sortBy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sortBy;

    /**
     * @return The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
     * 
     */
    public Output<Optional<String>> sortBy() {
        return Codegen.optional(this.sortBy);
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
     * The property value is a string that denotes the geographic region associated with the time zone, (default UTC).
     * 
     */
    @Export(name="timezone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> timezone;

    /**
     * @return The property value is a string that denotes the geographic region associated with the time zone, (default UTC).
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HeatmapChart(java.lang.String name) {
        this(name, HeatmapChartArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HeatmapChart(java.lang.String name, HeatmapChartArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HeatmapChart(java.lang.String name, HeatmapChartArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/heatmapChart:HeatmapChart", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HeatmapChart(java.lang.String name, Output<java.lang.String> id, @Nullable HeatmapChartState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/heatmapChart:HeatmapChart", name, state, makeResourceOptions(options, id), false);
    }

    private static HeatmapChartArgs makeArgs(HeatmapChartArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HeatmapChartArgs.Empty : args;
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
    public static HeatmapChart get(java.lang.String name, Output<java.lang.String> id, @Nullable HeatmapChartState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HeatmapChart(name, id, state, options);
    }
}
