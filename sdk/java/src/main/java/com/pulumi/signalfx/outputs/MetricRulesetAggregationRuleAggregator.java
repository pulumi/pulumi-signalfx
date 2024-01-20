// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class MetricRulesetAggregationRuleAggregator {
    private List<String> dimensions;
    private Boolean dropDimensions;
    private String outputName;
    private String type;

    private MetricRulesetAggregationRuleAggregator() {}
    public List<String> dimensions() {
        return this.dimensions;
    }
    public Boolean dropDimensions() {
        return this.dropDimensions;
    }
    public String outputName() {
        return this.outputName;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MetricRulesetAggregationRuleAggregator defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> dimensions;
        private Boolean dropDimensions;
        private String outputName;
        private String type;
        public Builder() {}
        public Builder(MetricRulesetAggregationRuleAggregator defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dimensions = defaults.dimensions;
    	      this.dropDimensions = defaults.dropDimensions;
    	      this.outputName = defaults.outputName;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder dimensions(List<String> dimensions) {
            if (dimensions == null) {
              throw new MissingRequiredPropertyException("MetricRulesetAggregationRuleAggregator", "dimensions");
            }
            this.dimensions = dimensions;
            return this;
        }
        public Builder dimensions(String... dimensions) {
            return dimensions(List.of(dimensions));
        }
        @CustomType.Setter
        public Builder dropDimensions(Boolean dropDimensions) {
            if (dropDimensions == null) {
              throw new MissingRequiredPropertyException("MetricRulesetAggregationRuleAggregator", "dropDimensions");
            }
            this.dropDimensions = dropDimensions;
            return this;
        }
        @CustomType.Setter
        public Builder outputName(String outputName) {
            if (outputName == null) {
              throw new MissingRequiredPropertyException("MetricRulesetAggregationRuleAggregator", "outputName");
            }
            this.outputName = outputName;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("MetricRulesetAggregationRuleAggregator", "type");
            }
            this.type = type;
            return this;
        }
        public MetricRulesetAggregationRuleAggregator build() {
            final var _resultValue = new MetricRulesetAggregationRuleAggregator();
            _resultValue.dimensions = dimensions;
            _resultValue.dropDimensions = dropDimensions;
            _resultValue.outputName = outputName;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
