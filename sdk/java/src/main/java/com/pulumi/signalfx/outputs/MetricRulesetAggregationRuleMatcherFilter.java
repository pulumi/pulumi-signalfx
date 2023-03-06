// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class MetricRulesetAggregationRuleMatcherFilter {
    /**
     * @return When true, this filter will match all values not matching the property_values
     * 
     */
    private Boolean not;
    /**
     * @return Name of the dimension
     * 
     */
    private String property;
    /**
     * @return Value of the dimension
     * 
     */
    private List<String> propertyValues;

    private MetricRulesetAggregationRuleMatcherFilter() {}
    /**
     * @return When true, this filter will match all values not matching the property_values
     * 
     */
    public Boolean not() {
        return this.not;
    }
    /**
     * @return Name of the dimension
     * 
     */
    public String property() {
        return this.property;
    }
    /**
     * @return Value of the dimension
     * 
     */
    public List<String> propertyValues() {
        return this.propertyValues;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MetricRulesetAggregationRuleMatcherFilter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean not;
        private String property;
        private List<String> propertyValues;
        public Builder() {}
        public Builder(MetricRulesetAggregationRuleMatcherFilter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.not = defaults.not;
    	      this.property = defaults.property;
    	      this.propertyValues = defaults.propertyValues;
        }

        @CustomType.Setter
        public Builder not(Boolean not) {
            this.not = Objects.requireNonNull(not);
            return this;
        }
        @CustomType.Setter
        public Builder property(String property) {
            this.property = Objects.requireNonNull(property);
            return this;
        }
        @CustomType.Setter
        public Builder propertyValues(List<String> propertyValues) {
            this.propertyValues = Objects.requireNonNull(propertyValues);
            return this;
        }
        public Builder propertyValues(String... propertyValues) {
            return propertyValues(List.of(propertyValues));
        }
        public MetricRulesetAggregationRuleMatcherFilter build() {
            final var o = new MetricRulesetAggregationRuleMatcherFilter();
            o.not = not;
            o.property = property;
            o.propertyValues = propertyValues;
            return o;
        }
    }
}
