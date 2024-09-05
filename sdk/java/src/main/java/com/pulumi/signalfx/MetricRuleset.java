// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.MetricRulesetArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.MetricRulesetState;
import com.pulumi.signalfx.outputs.MetricRulesetAggregationRule;
import com.pulumi.signalfx.outputs.MetricRulesetExceptionRule;
import com.pulumi.signalfx.outputs.MetricRulesetRoutingRule;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Observability Cloud resource for managing metric rulesets.
 * 
 * &gt; **NOTE** When managing metric rulesets to drop data use a session token for an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you&#39;ll receive a 4xx error.
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
 * import com.pulumi.signalfx.MetricRuleset;
 * import com.pulumi.signalfx.MetricRulesetArgs;
 * import com.pulumi.signalfx.inputs.MetricRulesetAggregationRuleArgs;
 * import com.pulumi.signalfx.inputs.MetricRulesetExceptionRuleArgs;
 * import com.pulumi.signalfx.inputs.MetricRulesetRoutingRuleArgs;
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
 *         var cpuUtilizationMetricRuleset = new MetricRuleset("cpuUtilizationMetricRuleset", MetricRulesetArgs.builder()
 *             .metricName("cpu.utilization")
 *             .description("Routing ruleset for cpu.utilization")
 *             .aggregationRules(MetricRulesetAggregationRuleArgs.builder()
 *                 .name("cpu.utilization by service rule")
 *                 .description("Aggregates cpu.utilization data by service")
 *                 .enabled(true)
 *                 .matchers(MetricRulesetAggregationRuleMatcherArgs.builder()
 *                     .type("dimension")
 *                     .filters(MetricRulesetAggregationRuleMatcherFilterArgs.builder()
 *                         .property("realm")
 *                         .propertyValues("us-east-1")
 *                         .not(false)
 *                         .build())
 *                     .build())
 *                 .aggregators(MetricRulesetAggregationRuleAggregatorArgs.builder()
 *                     .type("rollup")
 *                     .dimensions("service")
 *                     .dropDimensions(false)
 *                     .outputName("cpu.utilization.by.service.agg")
 *                     .build())
 *                 .build())
 *             .exceptionRules(MetricRulesetExceptionRuleArgs.builder()
 *                 .name("Exception rule us-east-2")
 *                 .description("Routes us-east-2 data to real-time")
 *                 .enabled(true)
 *                 .matchers(MetricRulesetExceptionRuleMatcherArgs.builder()
 *                     .type("dimension")
 *                     .filters(MetricRulesetExceptionRuleMatcherFilterArgs.builder()
 *                         .property("realm")
 *                         .propertyValues("us-east-2")
 *                         .not(false)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .routingRules(MetricRulesetRoutingRuleArgs.builder()
 *                 .destination("Archived")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="signalfx:index/metricRuleset:MetricRuleset")
public class MetricRuleset extends com.pulumi.resources.CustomResource {
    /**
     * List of aggregation rules for the metric
     * 
     */
    @Export(name="aggregationRules", refs={List.class,MetricRulesetAggregationRule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<MetricRulesetAggregationRule>> aggregationRules;

    /**
     * @return List of aggregation rules for the metric
     * 
     */
    public Output<Optional<List<MetricRulesetAggregationRule>>> aggregationRules() {
        return Codegen.optional(this.aggregationRules);
    }
    /**
     * Timestamp of when the metric ruleset was created
     * 
     */
    @Export(name="created", refs={String.class}, tree="[0]")
    private Output<String> created;

    /**
     * @return Timestamp of when the metric ruleset was created
     * 
     */
    public Output<String> created() {
        return this.created;
    }
    /**
     * ID of the creator of the metric ruleset
     * 
     */
    @Export(name="creator", refs={String.class}, tree="[0]")
    private Output<String> creator;

    /**
     * @return ID of the creator of the metric ruleset
     * 
     */
    public Output<String> creator() {
        return this.creator;
    }
    /**
     * Information about the metric ruleset
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Information about the metric ruleset
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * List of exception rules for the metric
     * 
     */
    @Export(name="exceptionRules", refs={List.class,MetricRulesetExceptionRule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<MetricRulesetExceptionRule>> exceptionRules;

    /**
     * @return List of exception rules for the metric
     * 
     */
    public Output<Optional<List<MetricRulesetExceptionRule>>> exceptionRules() {
        return Codegen.optional(this.exceptionRules);
    }
    /**
     * Timestamp of when the metric ruleset was last updated
     * 
     */
    @Export(name="lastUpdated", refs={String.class}, tree="[0]")
    private Output<String> lastUpdated;

    /**
     * @return Timestamp of when the metric ruleset was last updated
     * 
     */
    public Output<String> lastUpdated() {
        return this.lastUpdated;
    }
    /**
     * ID of user who last updated the metric ruleset
     * 
     */
    @Export(name="lastUpdatedBy", refs={String.class}, tree="[0]")
    private Output<String> lastUpdatedBy;

    /**
     * @return ID of user who last updated the metric ruleset
     * 
     */
    public Output<String> lastUpdatedBy() {
        return this.lastUpdatedBy;
    }
    /**
     * Name of user who last updated this metric ruleset
     * 
     */
    @Export(name="lastUpdatedByName", refs={String.class}, tree="[0]")
    private Output<String> lastUpdatedByName;

    /**
     * @return Name of user who last updated this metric ruleset
     * 
     */
    public Output<String> lastUpdatedByName() {
        return this.lastUpdatedByName;
    }
    /**
     * Name of the input metric
     * 
     */
    @Export(name="metricName", refs={String.class}, tree="[0]")
    private Output<String> metricName;

    /**
     * @return Name of the input metric
     * 
     */
    public Output<String> metricName() {
        return this.metricName;
    }
    /**
     * Routing Rule object
     * 
     */
    @Export(name="routingRules", refs={List.class,MetricRulesetRoutingRule.class}, tree="[0,1]")
    private Output<List<MetricRulesetRoutingRule>> routingRules;

    /**
     * @return Routing Rule object
     * 
     */
    public Output<List<MetricRulesetRoutingRule>> routingRules() {
        return this.routingRules;
    }
    /**
     * Version of the ruleset
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return Version of the ruleset
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MetricRuleset(java.lang.String name) {
        this(name, MetricRulesetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MetricRuleset(java.lang.String name, MetricRulesetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MetricRuleset(java.lang.String name, MetricRulesetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/metricRuleset:MetricRuleset", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private MetricRuleset(java.lang.String name, Output<java.lang.String> id, @Nullable MetricRulesetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/metricRuleset:MetricRuleset", name, state, makeResourceOptions(options, id), false);
    }

    private static MetricRulesetArgs makeArgs(MetricRulesetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MetricRulesetArgs.Empty : args;
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
    public static MetricRuleset get(java.lang.String name, Output<java.lang.String> id, @Nullable MetricRulesetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MetricRuleset(name, id, state, options);
    }
}
