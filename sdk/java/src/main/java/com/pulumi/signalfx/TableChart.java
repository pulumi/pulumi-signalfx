// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.TableChartArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.TableChartState;
import com.pulumi.signalfx.outputs.TableChartVizOption;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This special type of chart displays a Data Table. This Table can be grouped by a Dimension.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.TableChart;
 * import com.pulumi.signalfx.TableChartArgs;
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
 *         var table0 = new TableChart(&#34;table0&#34;, TableChartArgs.builder()        
 *             .description(&#34;beep&#34;)
 *             .disableSampling(false)
 *             .groupBies(&#34;ClusterName&#34;)
 *             .maxDelay(0)
 *             .programText(&#34;A = data(&#39;cpu.usage.total&#39;).publish(label=&#39;CPU Total&#39;)&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="signalfx:index/tableChart:TableChart")
public class TableChart extends com.pulumi.resources.CustomResource {
    /**
     * Description of the table chart.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the table chart.
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
     * Dimension to group by
     * 
     */
    @Export(name="groupBies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> groupBies;

    /**
     * @return Dimension to group by
     * 
     */
    public Output<Optional<List<String>>> groupBies() {
        return Codegen.optional(this.groupBies);
    }
    /**
     * (false by default) Whether to show the timestamp in the chart
     * 
     */
    @Export(name="hideTimestamp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hideTimestamp;

    /**
     * @return (false by default) Whether to show the timestamp in the chart
     * 
     */
    public Output<Optional<Boolean>> hideTimestamp() {
        return Codegen.optional(this.hideTimestamp);
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
     * The minimum resolution (in seconds) to use for computing the underlying program
     * 
     */
    @Export(name="minimumResolution", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minimumResolution;

    /**
     * @return The minimum resolution (in seconds) to use for computing the underlying program
     * 
     */
    public Output<Optional<Integer>> minimumResolution() {
        return Codegen.optional(this.minimumResolution);
    }
    /**
     * Name of the table chart.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the table chart.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The SignalFlow for your Data Table Chart
     * 
     */
    @Export(name="programText", refs={String.class}, tree="[0]")
    private Output<String> programText;

    /**
     * @return The SignalFlow for your Data Table Chart
     * 
     */
    public Output<String> programText() {
        return this.programText;
    }
    /**
     * How often (in seconds) to refresh the values of the Table
     * 
     */
    @Export(name="refreshInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> refreshInterval;

    /**
     * @return How often (in seconds) to refresh the values of the Table
     * 
     */
    public Output<Optional<Integer>> refreshInterval() {
        return Codegen.optional(this.refreshInterval);
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
     * Plot-level customization options, associated with a publish statement
     * 
     */
    @Export(name="vizOptions", refs={List.class,TableChartVizOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TableChartVizOption>> vizOptions;

    /**
     * @return Plot-level customization options, associated with a publish statement
     * 
     */
    public Output<Optional<List<TableChartVizOption>>> vizOptions() {
        return Codegen.optional(this.vizOptions);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TableChart(String name) {
        this(name, TableChartArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TableChart(String name, TableChartArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TableChart(String name, TableChartArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/tableChart:TableChart", name, args == null ? TableChartArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TableChart(String name, Output<String> id, @Nullable TableChartState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/tableChart:TableChart", name, state, makeResourceOptions(options, id));
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
    public static TableChart get(String name, Output<String> id, @Nullable TableChartState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TableChart(name, id, state, options);
    }
}
