// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.aws.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class IntegrationMetricStatsToSyncArgs extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationMetricStatsToSyncArgs Empty = new IntegrationMetricStatsToSyncArgs();

    @Import(name="metric", required=true)
    private Output<String> metric;

    public Output<String> metric() {
        return this.metric;
    }

    @Import(name="namespace", required=true)
    private Output<String> namespace;

    public Output<String> namespace() {
        return this.namespace;
    }

    @Import(name="stats", required=true)
    private Output<List<String>> stats;

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

        public Builder metric(Output<String> metric) {
            $.metric = metric;
            return this;
        }

        public Builder metric(String metric) {
            return metric(Output.of(metric));
        }

        public Builder namespace(Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public Builder stats(Output<List<String>> stats) {
            $.stats = stats;
            return this;
        }

        public Builder stats(List<String> stats) {
            return stats(Output.of(stats));
        }

        public Builder stats(String... stats) {
            return stats(List.of(stats));
        }

        public IntegrationMetricStatsToSyncArgs build() {
            if ($.metric == null) {
                throw new MissingRequiredPropertyException("IntegrationMetricStatsToSyncArgs", "metric");
            }
            if ($.namespace == null) {
                throw new MissingRequiredPropertyException("IntegrationMetricStatsToSyncArgs", "namespace");
            }
            if ($.stats == null) {
                throw new MissingRequiredPropertyException("IntegrationMetricStatsToSyncArgs", "stats");
            }
            return $;
        }
    }

}
