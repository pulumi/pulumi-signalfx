// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.inputs.MetricRulesetAggregationRuleMatcherFilterArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MetricRulesetAggregationRuleMatcherArgs extends com.pulumi.resources.ResourceArgs {

    public static final MetricRulesetAggregationRuleMatcherArgs Empty = new MetricRulesetAggregationRuleMatcherArgs();

    /**
     * List of filters to filter the set of input MTSs
     * 
     */
    @Import(name="filters")
    private @Nullable Output<List<MetricRulesetAggregationRuleMatcherFilterArgs>> filters;

    /**
     * @return List of filters to filter the set of input MTSs
     * 
     */
    public Optional<Output<List<MetricRulesetAggregationRuleMatcherFilterArgs>>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * Type of matcher. Must always be &#34;dimension&#34;
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of matcher. Must always be &#34;dimension&#34;
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private MetricRulesetAggregationRuleMatcherArgs() {}

    private MetricRulesetAggregationRuleMatcherArgs(MetricRulesetAggregationRuleMatcherArgs $) {
        this.filters = $.filters;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MetricRulesetAggregationRuleMatcherArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MetricRulesetAggregationRuleMatcherArgs $;

        public Builder() {
            $ = new MetricRulesetAggregationRuleMatcherArgs();
        }

        public Builder(MetricRulesetAggregationRuleMatcherArgs defaults) {
            $ = new MetricRulesetAggregationRuleMatcherArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters List of filters to filter the set of input MTSs
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable Output<List<MetricRulesetAggregationRuleMatcherFilterArgs>> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters List of filters to filter the set of input MTSs
         * 
         * @return builder
         * 
         */
        public Builder filters(List<MetricRulesetAggregationRuleMatcherFilterArgs> filters) {
            return filters(Output.of(filters));
        }

        /**
         * @param filters List of filters to filter the set of input MTSs
         * 
         * @return builder
         * 
         */
        public Builder filters(MetricRulesetAggregationRuleMatcherFilterArgs... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param type Type of matcher. Must always be &#34;dimension&#34;
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of matcher. Must always be &#34;dimension&#34;
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public MetricRulesetAggregationRuleMatcherArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
