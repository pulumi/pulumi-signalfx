// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.inputs.MetricRulesetAggregationRuleArgs;
import com.pulumi.signalfx.inputs.MetricRulesetExceptionRuleArgs;
import com.pulumi.signalfx.inputs.MetricRulesetRoutingRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MetricRulesetState extends com.pulumi.resources.ResourceArgs {

    public static final MetricRulesetState Empty = new MetricRulesetState();

    /**
     * List of aggregation rules for the metric
     * 
     */
    @Import(name="aggregationRules")
    private @Nullable Output<List<MetricRulesetAggregationRuleArgs>> aggregationRules;

    /**
     * @return List of aggregation rules for the metric
     * 
     */
    public Optional<Output<List<MetricRulesetAggregationRuleArgs>>> aggregationRules() {
        return Optional.ofNullable(this.aggregationRules);
    }

    /**
     * Timestamp of when the metric ruleset was created
     * 
     */
    @Import(name="created")
    private @Nullable Output<String> created;

    /**
     * @return Timestamp of when the metric ruleset was created
     * 
     */
    public Optional<Output<String>> created() {
        return Optional.ofNullable(this.created);
    }

    /**
     * ID of the creator of the metric ruleset
     * 
     */
    @Import(name="creator")
    private @Nullable Output<String> creator;

    /**
     * @return ID of the creator of the metric ruleset
     * 
     */
    public Optional<Output<String>> creator() {
        return Optional.ofNullable(this.creator);
    }

    /**
     * Information about the metric ruleset
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Information about the metric ruleset
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * List of exception rules for the metric
     * 
     */
    @Import(name="exceptionRules")
    private @Nullable Output<List<MetricRulesetExceptionRuleArgs>> exceptionRules;

    /**
     * @return List of exception rules for the metric
     * 
     */
    public Optional<Output<List<MetricRulesetExceptionRuleArgs>>> exceptionRules() {
        return Optional.ofNullable(this.exceptionRules);
    }

    /**
     * Timestamp of when the metric ruleset was last updated
     * 
     */
    @Import(name="lastUpdated")
    private @Nullable Output<String> lastUpdated;

    /**
     * @return Timestamp of when the metric ruleset was last updated
     * 
     */
    public Optional<Output<String>> lastUpdated() {
        return Optional.ofNullable(this.lastUpdated);
    }

    /**
     * ID of user who last updated the metric ruleset
     * 
     */
    @Import(name="lastUpdatedBy")
    private @Nullable Output<String> lastUpdatedBy;

    /**
     * @return ID of user who last updated the metric ruleset
     * 
     */
    public Optional<Output<String>> lastUpdatedBy() {
        return Optional.ofNullable(this.lastUpdatedBy);
    }

    /**
     * Name of user who last updated this metric ruleset
     * 
     */
    @Import(name="lastUpdatedByName")
    private @Nullable Output<String> lastUpdatedByName;

    /**
     * @return Name of user who last updated this metric ruleset
     * 
     */
    public Optional<Output<String>> lastUpdatedByName() {
        return Optional.ofNullable(this.lastUpdatedByName);
    }

    /**
     * Name of the input metric
     * 
     */
    @Import(name="metricName")
    private @Nullable Output<String> metricName;

    /**
     * @return Name of the input metric
     * 
     */
    public Optional<Output<String>> metricName() {
        return Optional.ofNullable(this.metricName);
    }

    /**
     * Routing Rule object
     * 
     */
    @Import(name="routingRules")
    private @Nullable Output<List<MetricRulesetRoutingRuleArgs>> routingRules;

    /**
     * @return Routing Rule object
     * 
     */
    public Optional<Output<List<MetricRulesetRoutingRuleArgs>>> routingRules() {
        return Optional.ofNullable(this.routingRules);
    }

    /**
     * Version of the ruleset
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Version of the ruleset
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private MetricRulesetState() {}

    private MetricRulesetState(MetricRulesetState $) {
        this.aggregationRules = $.aggregationRules;
        this.created = $.created;
        this.creator = $.creator;
        this.description = $.description;
        this.exceptionRules = $.exceptionRules;
        this.lastUpdated = $.lastUpdated;
        this.lastUpdatedBy = $.lastUpdatedBy;
        this.lastUpdatedByName = $.lastUpdatedByName;
        this.metricName = $.metricName;
        this.routingRules = $.routingRules;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MetricRulesetState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MetricRulesetState $;

        public Builder() {
            $ = new MetricRulesetState();
        }

        public Builder(MetricRulesetState defaults) {
            $ = new MetricRulesetState(Objects.requireNonNull(defaults));
        }

        /**
         * @param aggregationRules List of aggregation rules for the metric
         * 
         * @return builder
         * 
         */
        public Builder aggregationRules(@Nullable Output<List<MetricRulesetAggregationRuleArgs>> aggregationRules) {
            $.aggregationRules = aggregationRules;
            return this;
        }

        /**
         * @param aggregationRules List of aggregation rules for the metric
         * 
         * @return builder
         * 
         */
        public Builder aggregationRules(List<MetricRulesetAggregationRuleArgs> aggregationRules) {
            return aggregationRules(Output.of(aggregationRules));
        }

        /**
         * @param aggregationRules List of aggregation rules for the metric
         * 
         * @return builder
         * 
         */
        public Builder aggregationRules(MetricRulesetAggregationRuleArgs... aggregationRules) {
            return aggregationRules(List.of(aggregationRules));
        }

        /**
         * @param created Timestamp of when the metric ruleset was created
         * 
         * @return builder
         * 
         */
        public Builder created(@Nullable Output<String> created) {
            $.created = created;
            return this;
        }

        /**
         * @param created Timestamp of when the metric ruleset was created
         * 
         * @return builder
         * 
         */
        public Builder created(String created) {
            return created(Output.of(created));
        }

        /**
         * @param creator ID of the creator of the metric ruleset
         * 
         * @return builder
         * 
         */
        public Builder creator(@Nullable Output<String> creator) {
            $.creator = creator;
            return this;
        }

        /**
         * @param creator ID of the creator of the metric ruleset
         * 
         * @return builder
         * 
         */
        public Builder creator(String creator) {
            return creator(Output.of(creator));
        }

        /**
         * @param description Information about the metric ruleset
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Information about the metric ruleset
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param exceptionRules List of exception rules for the metric
         * 
         * @return builder
         * 
         */
        public Builder exceptionRules(@Nullable Output<List<MetricRulesetExceptionRuleArgs>> exceptionRules) {
            $.exceptionRules = exceptionRules;
            return this;
        }

        /**
         * @param exceptionRules List of exception rules for the metric
         * 
         * @return builder
         * 
         */
        public Builder exceptionRules(List<MetricRulesetExceptionRuleArgs> exceptionRules) {
            return exceptionRules(Output.of(exceptionRules));
        }

        /**
         * @param exceptionRules List of exception rules for the metric
         * 
         * @return builder
         * 
         */
        public Builder exceptionRules(MetricRulesetExceptionRuleArgs... exceptionRules) {
            return exceptionRules(List.of(exceptionRules));
        }

        /**
         * @param lastUpdated Timestamp of when the metric ruleset was last updated
         * 
         * @return builder
         * 
         */
        public Builder lastUpdated(@Nullable Output<String> lastUpdated) {
            $.lastUpdated = lastUpdated;
            return this;
        }

        /**
         * @param lastUpdated Timestamp of when the metric ruleset was last updated
         * 
         * @return builder
         * 
         */
        public Builder lastUpdated(String lastUpdated) {
            return lastUpdated(Output.of(lastUpdated));
        }

        /**
         * @param lastUpdatedBy ID of user who last updated the metric ruleset
         * 
         * @return builder
         * 
         */
        public Builder lastUpdatedBy(@Nullable Output<String> lastUpdatedBy) {
            $.lastUpdatedBy = lastUpdatedBy;
            return this;
        }

        /**
         * @param lastUpdatedBy ID of user who last updated the metric ruleset
         * 
         * @return builder
         * 
         */
        public Builder lastUpdatedBy(String lastUpdatedBy) {
            return lastUpdatedBy(Output.of(lastUpdatedBy));
        }

        /**
         * @param lastUpdatedByName Name of user who last updated this metric ruleset
         * 
         * @return builder
         * 
         */
        public Builder lastUpdatedByName(@Nullable Output<String> lastUpdatedByName) {
            $.lastUpdatedByName = lastUpdatedByName;
            return this;
        }

        /**
         * @param lastUpdatedByName Name of user who last updated this metric ruleset
         * 
         * @return builder
         * 
         */
        public Builder lastUpdatedByName(String lastUpdatedByName) {
            return lastUpdatedByName(Output.of(lastUpdatedByName));
        }

        /**
         * @param metricName Name of the input metric
         * 
         * @return builder
         * 
         */
        public Builder metricName(@Nullable Output<String> metricName) {
            $.metricName = metricName;
            return this;
        }

        /**
         * @param metricName Name of the input metric
         * 
         * @return builder
         * 
         */
        public Builder metricName(String metricName) {
            return metricName(Output.of(metricName));
        }

        /**
         * @param routingRules Routing Rule object
         * 
         * @return builder
         * 
         */
        public Builder routingRules(@Nullable Output<List<MetricRulesetRoutingRuleArgs>> routingRules) {
            $.routingRules = routingRules;
            return this;
        }

        /**
         * @param routingRules Routing Rule object
         * 
         * @return builder
         * 
         */
        public Builder routingRules(List<MetricRulesetRoutingRuleArgs> routingRules) {
            return routingRules(Output.of(routingRules));
        }

        /**
         * @param routingRules Routing Rule object
         * 
         * @return builder
         * 
         */
        public Builder routingRules(MetricRulesetRoutingRuleArgs... routingRules) {
            return routingRules(List.of(routingRules));
        }

        /**
         * @param version Version of the ruleset
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Version of the ruleset
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public MetricRulesetState build() {
            return $;
        }
    }

}
