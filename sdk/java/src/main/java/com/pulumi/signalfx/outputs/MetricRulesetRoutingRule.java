// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class MetricRulesetRoutingRule {
    /**
     * @return end destination of the input metric. Must be `RealTime`, `Archived`, or `Drop`
     * 
     */
    private String destination;

    private MetricRulesetRoutingRule() {}
    /**
     * @return end destination of the input metric. Must be `RealTime`, `Archived`, or `Drop`
     * 
     */
    public String destination() {
        return this.destination;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MetricRulesetRoutingRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String destination;
        public Builder() {}
        public Builder(MetricRulesetRoutingRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.destination = defaults.destination;
        }

        @CustomType.Setter
        public Builder destination(String destination) {
            if (destination == null) {
              throw new MissingRequiredPropertyException("MetricRulesetRoutingRule", "destination");
            }
            this.destination = destination;
            return this;
        }
        public MetricRulesetRoutingRule build() {
            final var _resultValue = new MetricRulesetRoutingRule();
            _resultValue.destination = destination;
            return _resultValue;
        }
    }
}
