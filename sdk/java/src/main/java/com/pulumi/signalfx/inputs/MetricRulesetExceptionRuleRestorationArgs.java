// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MetricRulesetExceptionRuleRestorationArgs extends com.pulumi.resources.ResourceArgs {

    public static final MetricRulesetExceptionRuleRestorationArgs Empty = new MetricRulesetExceptionRuleRestorationArgs();

    /**
     * ID of the restoration job.
     * 
     */
    @Import(name="restorationId")
    private @Nullable Output<String> restorationId;

    /**
     * @return ID of the restoration job.
     * 
     */
    public Optional<Output<String>> restorationId() {
        return Optional.ofNullable(this.restorationId);
    }

    /**
     * Time from which the restoration job will restore archived data, in the form of *nix time in milliseconds
     * 
     */
    @Import(name="startTime", required=true)
    private Output<String> startTime;

    /**
     * @return Time from which the restoration job will restore archived data, in the form of *nix time in milliseconds
     * 
     */
    public Output<String> startTime() {
        return this.startTime;
    }

    /**
     * Time to which the restoration job will restore archived data, in the form of *nix time in milliseconds
     * 
     */
    @Import(name="stopTime")
    private @Nullable Output<String> stopTime;

    /**
     * @return Time to which the restoration job will restore archived data, in the form of *nix time in milliseconds
     * 
     */
    public Optional<Output<String>> stopTime() {
        return Optional.ofNullable(this.stopTime);
    }

    private MetricRulesetExceptionRuleRestorationArgs() {}

    private MetricRulesetExceptionRuleRestorationArgs(MetricRulesetExceptionRuleRestorationArgs $) {
        this.restorationId = $.restorationId;
        this.startTime = $.startTime;
        this.stopTime = $.stopTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MetricRulesetExceptionRuleRestorationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MetricRulesetExceptionRuleRestorationArgs $;

        public Builder() {
            $ = new MetricRulesetExceptionRuleRestorationArgs();
        }

        public Builder(MetricRulesetExceptionRuleRestorationArgs defaults) {
            $ = new MetricRulesetExceptionRuleRestorationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param restorationId ID of the restoration job.
         * 
         * @return builder
         * 
         */
        public Builder restorationId(@Nullable Output<String> restorationId) {
            $.restorationId = restorationId;
            return this;
        }

        /**
         * @param restorationId ID of the restoration job.
         * 
         * @return builder
         * 
         */
        public Builder restorationId(String restorationId) {
            return restorationId(Output.of(restorationId));
        }

        /**
         * @param startTime Time from which the restoration job will restore archived data, in the form of *nix time in milliseconds
         * 
         * @return builder
         * 
         */
        public Builder startTime(Output<String> startTime) {
            $.startTime = startTime;
            return this;
        }

        /**
         * @param startTime Time from which the restoration job will restore archived data, in the form of *nix time in milliseconds
         * 
         * @return builder
         * 
         */
        public Builder startTime(String startTime) {
            return startTime(Output.of(startTime));
        }

        /**
         * @param stopTime Time to which the restoration job will restore archived data, in the form of *nix time in milliseconds
         * 
         * @return builder
         * 
         */
        public Builder stopTime(@Nullable Output<String> stopTime) {
            $.stopTime = stopTime;
            return this;
        }

        /**
         * @param stopTime Time to which the restoration job will restore archived data, in the form of *nix time in milliseconds
         * 
         * @return builder
         * 
         */
        public Builder stopTime(String stopTime) {
            return stopTime(Output.of(stopTime));
        }

        public MetricRulesetExceptionRuleRestorationArgs build() {
            if ($.startTime == null) {
                throw new MissingRequiredPropertyException("MetricRulesetExceptionRuleRestorationArgs", "startTime");
            }
            return $;
        }
    }

}
