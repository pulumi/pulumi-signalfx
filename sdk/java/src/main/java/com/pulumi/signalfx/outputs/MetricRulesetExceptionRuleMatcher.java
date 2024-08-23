// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.outputs.MetricRulesetExceptionRuleMatcherFilter;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class MetricRulesetExceptionRuleMatcher {
    /**
     * @return List of filters to filter the set of input MTSs
     * 
     */
    private @Nullable List<MetricRulesetExceptionRuleMatcherFilter> filters;
    /**
     * @return Type of matcher. Must always be &#34;dimension&#34;
     * 
     */
    private String type;

    private MetricRulesetExceptionRuleMatcher() {}
    /**
     * @return List of filters to filter the set of input MTSs
     * 
     */
    public List<MetricRulesetExceptionRuleMatcherFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return Type of matcher. Must always be &#34;dimension&#34;
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MetricRulesetExceptionRuleMatcher defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<MetricRulesetExceptionRuleMatcherFilter> filters;
        private String type;
        public Builder() {}
        public Builder(MetricRulesetExceptionRuleMatcher defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filters = defaults.filters;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder filters(@Nullable List<MetricRulesetExceptionRuleMatcherFilter> filters) {

            this.filters = filters;
            return this;
        }
        public Builder filters(MetricRulesetExceptionRuleMatcherFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("MetricRulesetExceptionRuleMatcher", "type");
            }
            this.type = type;
            return this;
        }
        public MetricRulesetExceptionRuleMatcher build() {
            final var _resultValue = new MetricRulesetExceptionRuleMatcher();
            _resultValue.filters = filters;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
