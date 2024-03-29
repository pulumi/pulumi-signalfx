// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class MetricRulesetAggregationRuleMatcherFilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final MetricRulesetAggregationRuleMatcherFilterArgs Empty = new MetricRulesetAggregationRuleMatcherFilterArgs();

    /**
     * Flag specifying equals or not equals
     * 
     */
    @Import(name="not", required=true)
    private Output<Boolean> not;

    /**
     * @return Flag specifying equals or not equals
     * 
     */
    public Output<Boolean> not() {
        return this.not;
    }

    /**
     * Name of dimension to match
     * 
     */
    @Import(name="property", required=true)
    private Output<String> property;

    /**
     * @return Name of dimension to match
     * 
     */
    public Output<String> property() {
        return this.property;
    }

    /**
     * List of property values to match
     * 
     */
    @Import(name="propertyValues", required=true)
    private Output<List<String>> propertyValues;

    /**
     * @return List of property values to match
     * 
     */
    public Output<List<String>> propertyValues() {
        return this.propertyValues;
    }

    private MetricRulesetAggregationRuleMatcherFilterArgs() {}

    private MetricRulesetAggregationRuleMatcherFilterArgs(MetricRulesetAggregationRuleMatcherFilterArgs $) {
        this.not = $.not;
        this.property = $.property;
        this.propertyValues = $.propertyValues;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MetricRulesetAggregationRuleMatcherFilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MetricRulesetAggregationRuleMatcherFilterArgs $;

        public Builder() {
            $ = new MetricRulesetAggregationRuleMatcherFilterArgs();
        }

        public Builder(MetricRulesetAggregationRuleMatcherFilterArgs defaults) {
            $ = new MetricRulesetAggregationRuleMatcherFilterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param not Flag specifying equals or not equals
         * 
         * @return builder
         * 
         */
        public Builder not(Output<Boolean> not) {
            $.not = not;
            return this;
        }

        /**
         * @param not Flag specifying equals or not equals
         * 
         * @return builder
         * 
         */
        public Builder not(Boolean not) {
            return not(Output.of(not));
        }

        /**
         * @param property Name of dimension to match
         * 
         * @return builder
         * 
         */
        public Builder property(Output<String> property) {
            $.property = property;
            return this;
        }

        /**
         * @param property Name of dimension to match
         * 
         * @return builder
         * 
         */
        public Builder property(String property) {
            return property(Output.of(property));
        }

        /**
         * @param propertyValues List of property values to match
         * 
         * @return builder
         * 
         */
        public Builder propertyValues(Output<List<String>> propertyValues) {
            $.propertyValues = propertyValues;
            return this;
        }

        /**
         * @param propertyValues List of property values to match
         * 
         * @return builder
         * 
         */
        public Builder propertyValues(List<String> propertyValues) {
            return propertyValues(Output.of(propertyValues));
        }

        /**
         * @param propertyValues List of property values to match
         * 
         * @return builder
         * 
         */
        public Builder propertyValues(String... propertyValues) {
            return propertyValues(List.of(propertyValues));
        }

        public MetricRulesetAggregationRuleMatcherFilterArgs build() {
            if ($.not == null) {
                throw new MissingRequiredPropertyException("MetricRulesetAggregationRuleMatcherFilterArgs", "not");
            }
            if ($.property == null) {
                throw new MissingRequiredPropertyException("MetricRulesetAggregationRuleMatcherFilterArgs", "property");
            }
            if ($.propertyValues == null) {
                throw new MissingRequiredPropertyException("MetricRulesetAggregationRuleMatcherFilterArgs", "propertyValues");
            }
            return $;
        }
    }

}
