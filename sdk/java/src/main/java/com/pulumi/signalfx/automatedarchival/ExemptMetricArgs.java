// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.automatedarchival;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.automatedarchival.inputs.ExemptMetricExemptMetricArgs;
import java.util.List;
import java.util.Objects;


public final class ExemptMetricArgs extends com.pulumi.resources.ResourceArgs {

    public static final ExemptMetricArgs Empty = new ExemptMetricArgs();

    /**
     * List of metrics to be exempted from automated archival
     * 
     */
    @Import(name="exemptMetrics", required=true)
    private Output<List<ExemptMetricExemptMetricArgs>> exemptMetrics;

    /**
     * @return List of metrics to be exempted from automated archival
     * 
     */
    public Output<List<ExemptMetricExemptMetricArgs>> exemptMetrics() {
        return this.exemptMetrics;
    }

    private ExemptMetricArgs() {}

    private ExemptMetricArgs(ExemptMetricArgs $) {
        this.exemptMetrics = $.exemptMetrics;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ExemptMetricArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ExemptMetricArgs $;

        public Builder() {
            $ = new ExemptMetricArgs();
        }

        public Builder(ExemptMetricArgs defaults) {
            $ = new ExemptMetricArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param exemptMetrics List of metrics to be exempted from automated archival
         * 
         * @return builder
         * 
         */
        public Builder exemptMetrics(Output<List<ExemptMetricExemptMetricArgs>> exemptMetrics) {
            $.exemptMetrics = exemptMetrics;
            return this;
        }

        /**
         * @param exemptMetrics List of metrics to be exempted from automated archival
         * 
         * @return builder
         * 
         */
        public Builder exemptMetrics(List<ExemptMetricExemptMetricArgs> exemptMetrics) {
            return exemptMetrics(Output.of(exemptMetrics));
        }

        /**
         * @param exemptMetrics List of metrics to be exempted from automated archival
         * 
         * @return builder
         * 
         */
        public Builder exemptMetrics(ExemptMetricExemptMetricArgs... exemptMetrics) {
            return exemptMetrics(List.of(exemptMetrics));
        }

        public ExemptMetricArgs build() {
            if ($.exemptMetrics == null) {
                throw new MissingRequiredPropertyException("ExemptMetricArgs", "exemptMetrics");
            }
            return $;
        }
    }

}
