// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SloTargetAlertRuleRuleParametersArgs extends com.pulumi.resources.ResourceArgs {

    public static final SloTargetAlertRuleRuleParametersArgs Empty = new SloTargetAlertRuleRuleParametersArgs();

    /**
     * Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burn_rate_threshold_1 parameter.
     * 
     */
    @Import(name="burnRateThreshold1")
    private @Nullable Output<Double> burnRateThreshold1;

    /**
     * @return Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burn_rate_threshold_1 parameter.
     * 
     */
    public Optional<Output<Double>> burnRateThreshold1() {
        return Optional.ofNullable(this.burnRateThreshold1);
    }

    /**
     * Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burn_rate_threshold_2 parameter.
     * 
     */
    @Import(name="burnRateThreshold2")
    private @Nullable Output<Double> burnRateThreshold2;

    /**
     * @return Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burn_rate_threshold_2 parameter.
     * 
     */
    public Optional<Output<Double>> burnRateThreshold2() {
        return Optional.ofNullable(this.burnRateThreshold2);
    }

    /**
     * Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the fire_lasting parameter
     * 
     */
    @Import(name="fireLasting")
    private @Nullable Output<String> fireLasting;

    /**
     * @return Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the fire_lasting parameter
     * 
     */
    public Optional<Output<String>> fireLasting() {
        return Optional.ofNullable(this.fireLasting);
    }

    /**
     * Long window 1 used in burn rate alert calculation. This value must be longer than short_window_1` and shorter than 90 days. Note: BURN_RATE alert rules use the long_window_1 parameter.
     * 
     */
    @Import(name="longWindow1")
    private @Nullable Output<String> longWindow1;

    /**
     * @return Long window 1 used in burn rate alert calculation. This value must be longer than short_window_1` and shorter than 90 days. Note: BURN_RATE alert rules use the long_window_1 parameter.
     * 
     */
    public Optional<Output<String>> longWindow1() {
        return Optional.ofNullable(this.longWindow1);
    }

    /**
     * Long window 2 used in burn rate alert calculation. This value must be longer than short_window_2` and shorter than 90 days. Note: BURN_RATE alert rules use the long_window_2 parameter.
     * 
     */
    @Import(name="longWindow2")
    private @Nullable Output<String> longWindow2;

    /**
     * @return Long window 2 used in burn rate alert calculation. This value must be longer than short_window_2` and shorter than 90 days. Note: BURN_RATE alert rules use the long_window_2 parameter.
     * 
     */
    public Optional<Output<String>> longWindow2() {
        return Optional.ofNullable(this.longWindow2);
    }

    /**
     * Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: ERROR_BUDGET_LEFT alert rules use the percent_error_budget_left parameter.
     * 
     */
    @Import(name="percentErrorBudgetLeft")
    private @Nullable Output<Double> percentErrorBudgetLeft;

    /**
     * @return Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: ERROR_BUDGET_LEFT alert rules use the percent_error_budget_left parameter.
     * 
     */
    public Optional<Output<Double>> percentErrorBudgetLeft() {
        return Optional.ofNullable(this.percentErrorBudgetLeft);
    }

    /**
     * Percentage of the fire_lasting duration that the alert condition is met before the alert is triggered. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the percent_of_lasting parameter
     * 
     */
    @Import(name="percentOfLasting")
    private @Nullable Output<Double> percentOfLasting;

    /**
     * @return Percentage of the fire_lasting duration that the alert condition is met before the alert is triggered. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the percent_of_lasting parameter
     * 
     */
    public Optional<Output<Double>> percentOfLasting() {
        return Optional.ofNullable(this.percentOfLasting);
    }

    /**
     * Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_1. Note: BURN_RATE alert rules use the short_window_1 parameter.
     * 
     */
    @Import(name="shortWindow1")
    private @Nullable Output<String> shortWindow1;

    /**
     * @return Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_1. Note: BURN_RATE alert rules use the short_window_1 parameter.
     * 
     */
    public Optional<Output<String>> shortWindow1() {
        return Optional.ofNullable(this.shortWindow1);
    }

    /**
     * Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_2. Note: BURN_RATE alert rules use the short_window_2 parameter.
     * 
     */
    @Import(name="shortWindow2")
    private @Nullable Output<String> shortWindow2;

    /**
     * @return Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_2. Note: BURN_RATE alert rules use the short_window_2 parameter.
     * 
     */
    public Optional<Output<String>> shortWindow2() {
        return Optional.ofNullable(this.shortWindow2);
    }

    private SloTargetAlertRuleRuleParametersArgs() {}

    private SloTargetAlertRuleRuleParametersArgs(SloTargetAlertRuleRuleParametersArgs $) {
        this.burnRateThreshold1 = $.burnRateThreshold1;
        this.burnRateThreshold2 = $.burnRateThreshold2;
        this.fireLasting = $.fireLasting;
        this.longWindow1 = $.longWindow1;
        this.longWindow2 = $.longWindow2;
        this.percentErrorBudgetLeft = $.percentErrorBudgetLeft;
        this.percentOfLasting = $.percentOfLasting;
        this.shortWindow1 = $.shortWindow1;
        this.shortWindow2 = $.shortWindow2;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SloTargetAlertRuleRuleParametersArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SloTargetAlertRuleRuleParametersArgs $;

        public Builder() {
            $ = new SloTargetAlertRuleRuleParametersArgs();
        }

        public Builder(SloTargetAlertRuleRuleParametersArgs defaults) {
            $ = new SloTargetAlertRuleRuleParametersArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param burnRateThreshold1 Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burn_rate_threshold_1 parameter.
         * 
         * @return builder
         * 
         */
        public Builder burnRateThreshold1(@Nullable Output<Double> burnRateThreshold1) {
            $.burnRateThreshold1 = burnRateThreshold1;
            return this;
        }

        /**
         * @param burnRateThreshold1 Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burn_rate_threshold_1 parameter.
         * 
         * @return builder
         * 
         */
        public Builder burnRateThreshold1(Double burnRateThreshold1) {
            return burnRateThreshold1(Output.of(burnRateThreshold1));
        }

        /**
         * @param burnRateThreshold2 Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burn_rate_threshold_2 parameter.
         * 
         * @return builder
         * 
         */
        public Builder burnRateThreshold2(@Nullable Output<Double> burnRateThreshold2) {
            $.burnRateThreshold2 = burnRateThreshold2;
            return this;
        }

        /**
         * @param burnRateThreshold2 Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burn_rate_threshold_2 parameter.
         * 
         * @return builder
         * 
         */
        public Builder burnRateThreshold2(Double burnRateThreshold2) {
            return burnRateThreshold2(Output.of(burnRateThreshold2));
        }

        /**
         * @param fireLasting Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the fire_lasting parameter
         * 
         * @return builder
         * 
         */
        public Builder fireLasting(@Nullable Output<String> fireLasting) {
            $.fireLasting = fireLasting;
            return this;
        }

        /**
         * @param fireLasting Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the fire_lasting parameter
         * 
         * @return builder
         * 
         */
        public Builder fireLasting(String fireLasting) {
            return fireLasting(Output.of(fireLasting));
        }

        /**
         * @param longWindow1 Long window 1 used in burn rate alert calculation. This value must be longer than short_window_1` and shorter than 90 days. Note: BURN_RATE alert rules use the long_window_1 parameter.
         * 
         * @return builder
         * 
         */
        public Builder longWindow1(@Nullable Output<String> longWindow1) {
            $.longWindow1 = longWindow1;
            return this;
        }

        /**
         * @param longWindow1 Long window 1 used in burn rate alert calculation. This value must be longer than short_window_1` and shorter than 90 days. Note: BURN_RATE alert rules use the long_window_1 parameter.
         * 
         * @return builder
         * 
         */
        public Builder longWindow1(String longWindow1) {
            return longWindow1(Output.of(longWindow1));
        }

        /**
         * @param longWindow2 Long window 2 used in burn rate alert calculation. This value must be longer than short_window_2` and shorter than 90 days. Note: BURN_RATE alert rules use the long_window_2 parameter.
         * 
         * @return builder
         * 
         */
        public Builder longWindow2(@Nullable Output<String> longWindow2) {
            $.longWindow2 = longWindow2;
            return this;
        }

        /**
         * @param longWindow2 Long window 2 used in burn rate alert calculation. This value must be longer than short_window_2` and shorter than 90 days. Note: BURN_RATE alert rules use the long_window_2 parameter.
         * 
         * @return builder
         * 
         */
        public Builder longWindow2(String longWindow2) {
            return longWindow2(Output.of(longWindow2));
        }

        /**
         * @param percentErrorBudgetLeft Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: ERROR_BUDGET_LEFT alert rules use the percent_error_budget_left parameter.
         * 
         * @return builder
         * 
         */
        public Builder percentErrorBudgetLeft(@Nullable Output<Double> percentErrorBudgetLeft) {
            $.percentErrorBudgetLeft = percentErrorBudgetLeft;
            return this;
        }

        /**
         * @param percentErrorBudgetLeft Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: ERROR_BUDGET_LEFT alert rules use the percent_error_budget_left parameter.
         * 
         * @return builder
         * 
         */
        public Builder percentErrorBudgetLeft(Double percentErrorBudgetLeft) {
            return percentErrorBudgetLeft(Output.of(percentErrorBudgetLeft));
        }

        /**
         * @param percentOfLasting Percentage of the fire_lasting duration that the alert condition is met before the alert is triggered. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the percent_of_lasting parameter
         * 
         * @return builder
         * 
         */
        public Builder percentOfLasting(@Nullable Output<Double> percentOfLasting) {
            $.percentOfLasting = percentOfLasting;
            return this;
        }

        /**
         * @param percentOfLasting Percentage of the fire_lasting duration that the alert condition is met before the alert is triggered. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the percent_of_lasting parameter
         * 
         * @return builder
         * 
         */
        public Builder percentOfLasting(Double percentOfLasting) {
            return percentOfLasting(Output.of(percentOfLasting));
        }

        /**
         * @param shortWindow1 Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_1. Note: BURN_RATE alert rules use the short_window_1 parameter.
         * 
         * @return builder
         * 
         */
        public Builder shortWindow1(@Nullable Output<String> shortWindow1) {
            $.shortWindow1 = shortWindow1;
            return this;
        }

        /**
         * @param shortWindow1 Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_1. Note: BURN_RATE alert rules use the short_window_1 parameter.
         * 
         * @return builder
         * 
         */
        public Builder shortWindow1(String shortWindow1) {
            return shortWindow1(Output.of(shortWindow1));
        }

        /**
         * @param shortWindow2 Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_2. Note: BURN_RATE alert rules use the short_window_2 parameter.
         * 
         * @return builder
         * 
         */
        public Builder shortWindow2(@Nullable Output<String> shortWindow2) {
            $.shortWindow2 = shortWindow2;
            return this;
        }

        /**
         * @param shortWindow2 Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_2. Note: BURN_RATE alert rules use the short_window_2 parameter.
         * 
         * @return builder
         * 
         */
        public Builder shortWindow2(String shortWindow2) {
            return shortWindow2(Output.of(shortWindow2));
        }

        public SloTargetAlertRuleRuleParametersArgs build() {
            return $;
        }
    }

}
