// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.outputs.MetricRulesetAggregationRuleAggregator;
import com.pulumi.signalfx.outputs.MetricRulesetAggregationRuleMatcher;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class MetricRulesetAggregationRule {
    private List<MetricRulesetAggregationRuleAggregator> aggregators;
    private Boolean enabled;
    private List<MetricRulesetAggregationRuleMatcher> matchers;
    private @Nullable String name;

    private MetricRulesetAggregationRule() {}
    public List<MetricRulesetAggregationRuleAggregator> aggregators() {
        return this.aggregators;
    }
    public Boolean enabled() {
        return this.enabled;
    }
    public List<MetricRulesetAggregationRuleMatcher> matchers() {
        return this.matchers;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MetricRulesetAggregationRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<MetricRulesetAggregationRuleAggregator> aggregators;
        private Boolean enabled;
        private List<MetricRulesetAggregationRuleMatcher> matchers;
        private @Nullable String name;
        public Builder() {}
        public Builder(MetricRulesetAggregationRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aggregators = defaults.aggregators;
    	      this.enabled = defaults.enabled;
    	      this.matchers = defaults.matchers;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder aggregators(List<MetricRulesetAggregationRuleAggregator> aggregators) {
            if (aggregators == null) {
              throw new MissingRequiredPropertyException("MetricRulesetAggregationRule", "aggregators");
            }
            this.aggregators = aggregators;
            return this;
        }
        public Builder aggregators(MetricRulesetAggregationRuleAggregator... aggregators) {
            return aggregators(List.of(aggregators));
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            if (enabled == null) {
              throw new MissingRequiredPropertyException("MetricRulesetAggregationRule", "enabled");
            }
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder matchers(List<MetricRulesetAggregationRuleMatcher> matchers) {
            if (matchers == null) {
              throw new MissingRequiredPropertyException("MetricRulesetAggregationRule", "matchers");
            }
            this.matchers = matchers;
            return this;
        }
        public Builder matchers(MetricRulesetAggregationRuleMatcher... matchers) {
            return matchers(List.of(matchers));
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        public MetricRulesetAggregationRule build() {
            final var _resultValue = new MetricRulesetAggregationRule();
            _resultValue.aggregators = aggregators;
            _resultValue.enabled = enabled;
            _resultValue.matchers = matchers;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}
