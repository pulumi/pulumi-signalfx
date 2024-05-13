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
    public MetricRuleset(String name) {
        this(name, MetricRulesetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MetricRuleset(String name, MetricRulesetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MetricRuleset(String name, MetricRulesetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/metricRuleset:MetricRuleset", name, args == null ? MetricRulesetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MetricRuleset(String name, Output<String> id, @Nullable MetricRulesetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/metricRuleset:MetricRuleset", name, state, makeResourceOptions(options, id));
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
    public static MetricRuleset get(String name, Output<String> id, @Nullable MetricRulesetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MetricRuleset(name, id, state, options);
    }
}
