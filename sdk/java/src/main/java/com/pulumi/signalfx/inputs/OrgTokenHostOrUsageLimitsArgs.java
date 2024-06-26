// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrgTokenHostOrUsageLimitsArgs extends com.pulumi.resources.ResourceArgs {

    public static final OrgTokenHostOrUsageLimitsArgs Empty = new OrgTokenHostOrUsageLimitsArgs();

    /**
     * Max number of Docker containers that can use this token
     * 
     */
    @Import(name="containerLimit")
    private @Nullable Output<Integer> containerLimit;

    /**
     * @return Max number of Docker containers that can use this token
     * 
     */
    public Optional<Output<Integer>> containerLimit() {
        return Optional.ofNullable(this.containerLimit);
    }

    /**
     * Notification threshold for Docker containers
     * 
     */
    @Import(name="containerNotificationThreshold")
    private @Nullable Output<Integer> containerNotificationThreshold;

    /**
     * @return Notification threshold for Docker containers
     * 
     */
    public Optional<Output<Integer>> containerNotificationThreshold() {
        return Optional.ofNullable(this.containerNotificationThreshold);
    }

    /**
     * Max number of custom metrics that can be sent with this token
     * 
     */
    @Import(name="customMetricsLimit")
    private @Nullable Output<Integer> customMetricsLimit;

    /**
     * @return Max number of custom metrics that can be sent with this token
     * 
     */
    public Optional<Output<Integer>> customMetricsLimit() {
        return Optional.ofNullable(this.customMetricsLimit);
    }

    /**
     * Notification threshold for custom metrics
     * 
     */
    @Import(name="customMetricsNotificationThreshold")
    private @Nullable Output<Integer> customMetricsNotificationThreshold;

    /**
     * @return Notification threshold for custom metrics
     * 
     */
    public Optional<Output<Integer>> customMetricsNotificationThreshold() {
        return Optional.ofNullable(this.customMetricsNotificationThreshold);
    }

    /**
     * Max number of hi-res metrics that can be sent with this toke
     * 
     */
    @Import(name="highResMetricsLimit")
    private @Nullable Output<Integer> highResMetricsLimit;

    /**
     * @return Max number of hi-res metrics that can be sent with this toke
     * 
     */
    public Optional<Output<Integer>> highResMetricsLimit() {
        return Optional.ofNullable(this.highResMetricsLimit);
    }

    /**
     * Notification threshold for hi-res metrics
     * 
     */
    @Import(name="highResMetricsNotificationThreshold")
    private @Nullable Output<Integer> highResMetricsNotificationThreshold;

    /**
     * @return Notification threshold for hi-res metrics
     * 
     */
    public Optional<Output<Integer>> highResMetricsNotificationThreshold() {
        return Optional.ofNullable(this.highResMetricsNotificationThreshold);
    }

    /**
     * Max number of hosts that can use this token
     * 
     */
    @Import(name="hostLimit")
    private @Nullable Output<Integer> hostLimit;

    /**
     * @return Max number of hosts that can use this token
     * 
     */
    public Optional<Output<Integer>> hostLimit() {
        return Optional.ofNullable(this.hostLimit);
    }

    /**
     * Notification threshold for hosts
     * 
     */
    @Import(name="hostNotificationThreshold")
    private @Nullable Output<Integer> hostNotificationThreshold;

    /**
     * @return Notification threshold for hosts
     * 
     */
    public Optional<Output<Integer>> hostNotificationThreshold() {
        return Optional.ofNullable(this.hostNotificationThreshold);
    }

    private OrgTokenHostOrUsageLimitsArgs() {}

    private OrgTokenHostOrUsageLimitsArgs(OrgTokenHostOrUsageLimitsArgs $) {
        this.containerLimit = $.containerLimit;
        this.containerNotificationThreshold = $.containerNotificationThreshold;
        this.customMetricsLimit = $.customMetricsLimit;
        this.customMetricsNotificationThreshold = $.customMetricsNotificationThreshold;
        this.highResMetricsLimit = $.highResMetricsLimit;
        this.highResMetricsNotificationThreshold = $.highResMetricsNotificationThreshold;
        this.hostLimit = $.hostLimit;
        this.hostNotificationThreshold = $.hostNotificationThreshold;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrgTokenHostOrUsageLimitsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrgTokenHostOrUsageLimitsArgs $;

        public Builder() {
            $ = new OrgTokenHostOrUsageLimitsArgs();
        }

        public Builder(OrgTokenHostOrUsageLimitsArgs defaults) {
            $ = new OrgTokenHostOrUsageLimitsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param containerLimit Max number of Docker containers that can use this token
         * 
         * @return builder
         * 
         */
        public Builder containerLimit(@Nullable Output<Integer> containerLimit) {
            $.containerLimit = containerLimit;
            return this;
        }

        /**
         * @param containerLimit Max number of Docker containers that can use this token
         * 
         * @return builder
         * 
         */
        public Builder containerLimit(Integer containerLimit) {
            return containerLimit(Output.of(containerLimit));
        }

        /**
         * @param containerNotificationThreshold Notification threshold for Docker containers
         * 
         * @return builder
         * 
         */
        public Builder containerNotificationThreshold(@Nullable Output<Integer> containerNotificationThreshold) {
            $.containerNotificationThreshold = containerNotificationThreshold;
            return this;
        }

        /**
         * @param containerNotificationThreshold Notification threshold for Docker containers
         * 
         * @return builder
         * 
         */
        public Builder containerNotificationThreshold(Integer containerNotificationThreshold) {
            return containerNotificationThreshold(Output.of(containerNotificationThreshold));
        }

        /**
         * @param customMetricsLimit Max number of custom metrics that can be sent with this token
         * 
         * @return builder
         * 
         */
        public Builder customMetricsLimit(@Nullable Output<Integer> customMetricsLimit) {
            $.customMetricsLimit = customMetricsLimit;
            return this;
        }

        /**
         * @param customMetricsLimit Max number of custom metrics that can be sent with this token
         * 
         * @return builder
         * 
         */
        public Builder customMetricsLimit(Integer customMetricsLimit) {
            return customMetricsLimit(Output.of(customMetricsLimit));
        }

        /**
         * @param customMetricsNotificationThreshold Notification threshold for custom metrics
         * 
         * @return builder
         * 
         */
        public Builder customMetricsNotificationThreshold(@Nullable Output<Integer> customMetricsNotificationThreshold) {
            $.customMetricsNotificationThreshold = customMetricsNotificationThreshold;
            return this;
        }

        /**
         * @param customMetricsNotificationThreshold Notification threshold for custom metrics
         * 
         * @return builder
         * 
         */
        public Builder customMetricsNotificationThreshold(Integer customMetricsNotificationThreshold) {
            return customMetricsNotificationThreshold(Output.of(customMetricsNotificationThreshold));
        }

        /**
         * @param highResMetricsLimit Max number of hi-res metrics that can be sent with this toke
         * 
         * @return builder
         * 
         */
        public Builder highResMetricsLimit(@Nullable Output<Integer> highResMetricsLimit) {
            $.highResMetricsLimit = highResMetricsLimit;
            return this;
        }

        /**
         * @param highResMetricsLimit Max number of hi-res metrics that can be sent with this toke
         * 
         * @return builder
         * 
         */
        public Builder highResMetricsLimit(Integer highResMetricsLimit) {
            return highResMetricsLimit(Output.of(highResMetricsLimit));
        }

        /**
         * @param highResMetricsNotificationThreshold Notification threshold for hi-res metrics
         * 
         * @return builder
         * 
         */
        public Builder highResMetricsNotificationThreshold(@Nullable Output<Integer> highResMetricsNotificationThreshold) {
            $.highResMetricsNotificationThreshold = highResMetricsNotificationThreshold;
            return this;
        }

        /**
         * @param highResMetricsNotificationThreshold Notification threshold for hi-res metrics
         * 
         * @return builder
         * 
         */
        public Builder highResMetricsNotificationThreshold(Integer highResMetricsNotificationThreshold) {
            return highResMetricsNotificationThreshold(Output.of(highResMetricsNotificationThreshold));
        }

        /**
         * @param hostLimit Max number of hosts that can use this token
         * 
         * @return builder
         * 
         */
        public Builder hostLimit(@Nullable Output<Integer> hostLimit) {
            $.hostLimit = hostLimit;
            return this;
        }

        /**
         * @param hostLimit Max number of hosts that can use this token
         * 
         * @return builder
         * 
         */
        public Builder hostLimit(Integer hostLimit) {
            return hostLimit(Output.of(hostLimit));
        }

        /**
         * @param hostNotificationThreshold Notification threshold for hosts
         * 
         * @return builder
         * 
         */
        public Builder hostNotificationThreshold(@Nullable Output<Integer> hostNotificationThreshold) {
            $.hostNotificationThreshold = hostNotificationThreshold;
            return this;
        }

        /**
         * @param hostNotificationThreshold Notification threshold for hosts
         * 
         * @return builder
         * 
         */
        public Builder hostNotificationThreshold(Integer hostNotificationThreshold) {
            return hostNotificationThreshold(Output.of(hostNotificationThreshold));
        }

        public OrgTokenHostOrUsageLimitsArgs build() {
            return $;
        }
    }

}
