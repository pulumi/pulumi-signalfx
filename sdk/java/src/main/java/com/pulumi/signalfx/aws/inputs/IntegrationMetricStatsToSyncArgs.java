// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.aws.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class IntegrationMetricStatsToSyncArgs extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationMetricStatsToSyncArgs Empty = new IntegrationMetricStatsToSyncArgs();

    /**
     * AWS metric that you want to pick statistics for
     * 
     */
    @Import(name="metric", required=true)
    private Output<String> metric;

    /**
     * @return AWS metric that you want to pick statistics for
     * 
     */
    public Output<String> metric() {
        return this.metric;
    }

    /**
     * An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
     * 
     */
    @Import(name="namespace", required=true)
    private Output<String> namespace;

    /**
     * @return An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
     * 
     */
    public Output<String> namespace() {
        return this.namespace;
    }

    /**
     * AWS statistics you want to collect
     * 
     */
    @Import(name="stats", required=true)
    private Output<List<String>> stats;

    /**
     * @return AWS statistics you want to collect
     * 
     */
    public Output<List<String>> stats() {
        return this.stats;
    }

    private IntegrationMetricStatsToSyncArgs() {}

    private IntegrationMetricStatsToSyncArgs(IntegrationMetricStatsToSyncArgs $) {
        this.metric = $.metric;
        this.namespace = $.namespace;
        this.stats = $.stats;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationMetricStatsToSyncArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationMetricStatsToSyncArgs $;

        public Builder() {
            $ = new IntegrationMetricStatsToSyncArgs();
        }

        public Builder(IntegrationMetricStatsToSyncArgs defaults) {
            $ = new IntegrationMetricStatsToSyncArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param metric AWS metric that you want to pick statistics for
         * 
         * @return builder
         * 
         */
        public Builder metric(Output<String> metric) {
            $.metric = metric;
            return this;
        }

        /**
         * @param metric AWS metric that you want to pick statistics for
         * 
         * @return builder
         * 
         */
        public Builder metric(String metric) {
            return metric(Output.of(metric));
        }

        /**
         * @param namespace An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
         * 
         * @return builder
         * 
         */
        public Builder namespace(Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param stats AWS statistics you want to collect
         * 
         * @return builder
         * 
         */
        public Builder stats(Output<List<String>> stats) {
            $.stats = stats;
            return this;
        }

        /**
         * @param stats AWS statistics you want to collect
         * 
         * @return builder
         * 
         */
        public Builder stats(List<String> stats) {
            return stats(Output.of(stats));
        }

        /**
         * @param stats AWS statistics you want to collect
         * 
         * @return builder
         * 
         */
        public Builder stats(String... stats) {
            return stats(List.of(stats));
        }

        public IntegrationMetricStatsToSyncArgs build() {
            $.metric = Objects.requireNonNull($.metric, "expected parameter 'metric' to be non-null");
            $.namespace = Objects.requireNonNull($.namespace, "expected parameter 'namespace' to be non-null");
            $.stats = Objects.requireNonNull($.stats, "expected parameter 'stats' to be non-null");
            return $;
        }
    }

}