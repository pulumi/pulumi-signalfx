// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.aws.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class IntegrationMetricStatsToSync {
    /**
     * @return AWS metric that you want to pick statistics for
     * 
     */
    private String metric;
    /**
     * @return An AWS namespace having AWS metric that you want to pick statistics for
     * 
     */
    private String namespace;
    /**
     * @return AWS statistics you want to collect
     * 
     */
    private List<String> stats;

    private IntegrationMetricStatsToSync() {}
    /**
     * @return AWS metric that you want to pick statistics for
     * 
     */
    public String metric() {
        return this.metric;
    }
    /**
     * @return An AWS namespace having AWS metric that you want to pick statistics for
     * 
     */
    public String namespace() {
        return this.namespace;
    }
    /**
     * @return AWS statistics you want to collect
     * 
     */
    public List<String> stats() {
        return this.stats;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IntegrationMetricStatsToSync defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String metric;
        private String namespace;
        private List<String> stats;
        public Builder() {}
        public Builder(IntegrationMetricStatsToSync defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.metric = defaults.metric;
    	      this.namespace = defaults.namespace;
    	      this.stats = defaults.stats;
        }

        @CustomType.Setter
        public Builder metric(String metric) {
            if (metric == null) {
              throw new MissingRequiredPropertyException("IntegrationMetricStatsToSync", "metric");
            }
            this.metric = metric;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(String namespace) {
            if (namespace == null) {
              throw new MissingRequiredPropertyException("IntegrationMetricStatsToSync", "namespace");
            }
            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder stats(List<String> stats) {
            if (stats == null) {
              throw new MissingRequiredPropertyException("IntegrationMetricStatsToSync", "stats");
            }
            this.stats = stats;
            return this;
        }
        public Builder stats(String... stats) {
            return stats(List.of(stats));
        }
        public IntegrationMetricStatsToSync build() {
            final var _resultValue = new IntegrationMetricStatsToSync();
            _resultValue.metric = metric;
            _resultValue.namespace = namespace;
            _resultValue.stats = stats;
            return _resultValue;
        }
    }
}
