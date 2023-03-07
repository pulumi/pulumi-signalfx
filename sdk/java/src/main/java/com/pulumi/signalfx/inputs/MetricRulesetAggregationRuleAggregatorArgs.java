// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class MetricRulesetAggregationRuleAggregatorArgs extends com.pulumi.resources.ResourceArgs {

    public static final MetricRulesetAggregationRuleAggregatorArgs Empty = new MetricRulesetAggregationRuleAggregatorArgs();

    /**
     * List of dimensions to either be kept or dropped in the new aggregated MTSs
     * 
     */
    @Import(name="dimensions", required=true)
    private Output<List<String>> dimensions;

    /**
     * @return List of dimensions to either be kept or dropped in the new aggregated MTSs
     * 
     */
    public Output<List<String>> dimensions() {
        return this.dimensions;
    }

    /**
     * when true, the specified dimensions will be dropped from the aggregated MTSs
     * 
     */
    @Import(name="dropDimensions", required=true)
    private Output<Boolean> dropDimensions;

    /**
     * @return when true, the specified dimensions will be dropped from the aggregated MTSs
     * 
     */
    public Output<Boolean> dropDimensions() {
        return this.dropDimensions;
    }

    /**
     * name of the new aggregated metric
     * 
     */
    @Import(name="outputName", required=true)
    private Output<String> outputName;

    /**
     * @return name of the new aggregated metric
     * 
     */
    public Output<String> outputName() {
        return this.outputName;
    }

    /**
     * Type of aggregator. Must always be &#34;rollup&#34;
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of aggregator. Must always be &#34;rollup&#34;
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private MetricRulesetAggregationRuleAggregatorArgs() {}

    private MetricRulesetAggregationRuleAggregatorArgs(MetricRulesetAggregationRuleAggregatorArgs $) {
        this.dimensions = $.dimensions;
        this.dropDimensions = $.dropDimensions;
        this.outputName = $.outputName;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MetricRulesetAggregationRuleAggregatorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MetricRulesetAggregationRuleAggregatorArgs $;

        public Builder() {
            $ = new MetricRulesetAggregationRuleAggregatorArgs();
        }

        public Builder(MetricRulesetAggregationRuleAggregatorArgs defaults) {
            $ = new MetricRulesetAggregationRuleAggregatorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dimensions List of dimensions to either be kept or dropped in the new aggregated MTSs
         * 
         * @return builder
         * 
         */
        public Builder dimensions(Output<List<String>> dimensions) {
            $.dimensions = dimensions;
            return this;
        }

        /**
         * @param dimensions List of dimensions to either be kept or dropped in the new aggregated MTSs
         * 
         * @return builder
         * 
         */
        public Builder dimensions(List<String> dimensions) {
            return dimensions(Output.of(dimensions));
        }

        /**
         * @param dimensions List of dimensions to either be kept or dropped in the new aggregated MTSs
         * 
         * @return builder
         * 
         */
        public Builder dimensions(String... dimensions) {
            return dimensions(List.of(dimensions));
        }

        /**
         * @param dropDimensions when true, the specified dimensions will be dropped from the aggregated MTSs
         * 
         * @return builder
         * 
         */
        public Builder dropDimensions(Output<Boolean> dropDimensions) {
            $.dropDimensions = dropDimensions;
            return this;
        }

        /**
         * @param dropDimensions when true, the specified dimensions will be dropped from the aggregated MTSs
         * 
         * @return builder
         * 
         */
        public Builder dropDimensions(Boolean dropDimensions) {
            return dropDimensions(Output.of(dropDimensions));
        }

        /**
         * @param outputName name of the new aggregated metric
         * 
         * @return builder
         * 
         */
        public Builder outputName(Output<String> outputName) {
            $.outputName = outputName;
            return this;
        }

        /**
         * @param outputName name of the new aggregated metric
         * 
         * @return builder
         * 
         */
        public Builder outputName(String outputName) {
            return outputName(Output.of(outputName));
        }

        /**
         * @param type Type of aggregator. Must always be &#34;rollup&#34;
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of aggregator. Must always be &#34;rollup&#34;
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public MetricRulesetAggregationRuleAggregatorArgs build() {
            $.dimensions = Objects.requireNonNull($.dimensions, "expected parameter 'dimensions' to be non-null");
            $.dropDimensions = Objects.requireNonNull($.dropDimensions, "expected parameter 'dropDimensions' to be non-null");
            $.outputName = Objects.requireNonNull($.outputName, "expected parameter 'outputName' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
