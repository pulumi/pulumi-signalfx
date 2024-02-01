// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.inputs.MetricRulesetAggregationRuleAggregatorArgs;
import com.pulumi.signalfx.inputs.MetricRulesetAggregationRuleMatcherArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MetricRulesetAggregationRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final MetricRulesetAggregationRuleArgs Empty = new MetricRulesetAggregationRuleArgs();

    /**
     * The aggregator for this rule
     * 
     */
    @Import(name="aggregators", required=true)
    private Output<List<MetricRulesetAggregationRuleAggregatorArgs>> aggregators;

    /**
     * @return The aggregator for this rule
     * 
     */
    public Output<List<MetricRulesetAggregationRuleAggregatorArgs>> aggregators() {
        return this.aggregators;
    }

    /**
     * Status of this aggregation rule
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Status of this aggregation rule
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * The matcher for this rule
     * 
     */
    @Import(name="matchers", required=true)
    private Output<List<MetricRulesetAggregationRuleMatcherArgs>> matchers;

    /**
     * @return The matcher for this rule
     * 
     */
    public Output<List<MetricRulesetAggregationRuleMatcherArgs>> matchers() {
        return this.matchers;
    }

    /**
     * Name of this aggregation rule
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of this aggregation rule
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private MetricRulesetAggregationRuleArgs() {}

    private MetricRulesetAggregationRuleArgs(MetricRulesetAggregationRuleArgs $) {
        this.aggregators = $.aggregators;
        this.enabled = $.enabled;
        this.matchers = $.matchers;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MetricRulesetAggregationRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MetricRulesetAggregationRuleArgs $;

        public Builder() {
            $ = new MetricRulesetAggregationRuleArgs();
        }

        public Builder(MetricRulesetAggregationRuleArgs defaults) {
            $ = new MetricRulesetAggregationRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param aggregators The aggregator for this rule
         * 
         * @return builder
         * 
         */
        public Builder aggregators(Output<List<MetricRulesetAggregationRuleAggregatorArgs>> aggregators) {
            $.aggregators = aggregators;
            return this;
        }

        /**
         * @param aggregators The aggregator for this rule
         * 
         * @return builder
         * 
         */
        public Builder aggregators(List<MetricRulesetAggregationRuleAggregatorArgs> aggregators) {
            return aggregators(Output.of(aggregators));
        }

        /**
         * @param aggregators The aggregator for this rule
         * 
         * @return builder
         * 
         */
        public Builder aggregators(MetricRulesetAggregationRuleAggregatorArgs... aggregators) {
            return aggregators(List.of(aggregators));
        }

        /**
         * @param enabled Status of this aggregation rule
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Status of this aggregation rule
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param matchers The matcher for this rule
         * 
         * @return builder
         * 
         */
        public Builder matchers(Output<List<MetricRulesetAggregationRuleMatcherArgs>> matchers) {
            $.matchers = matchers;
            return this;
        }

        /**
         * @param matchers The matcher for this rule
         * 
         * @return builder
         * 
         */
        public Builder matchers(List<MetricRulesetAggregationRuleMatcherArgs> matchers) {
            return matchers(Output.of(matchers));
        }

        /**
         * @param matchers The matcher for this rule
         * 
         * @return builder
         * 
         */
        public Builder matchers(MetricRulesetAggregationRuleMatcherArgs... matchers) {
            return matchers(List.of(matchers));
        }

        /**
         * @param name Name of this aggregation rule
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of this aggregation rule
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public MetricRulesetAggregationRuleArgs build() {
            if ($.aggregators == null) {
                throw new MissingRequiredPropertyException("MetricRulesetAggregationRuleArgs", "aggregators");
            }
            if ($.enabled == null) {
                throw new MissingRequiredPropertyException("MetricRulesetAggregationRuleArgs", "enabled");
            }
            if ($.matchers == null) {
                throw new MissingRequiredPropertyException("MetricRulesetAggregationRuleArgs", "matchers");
            }
            return $;
        }
    }

}
