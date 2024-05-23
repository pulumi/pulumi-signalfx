// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.inputs.SloTargetAlertRuleArgs;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SloTargetArgs extends com.pulumi.resources.ResourceArgs {

    public static final SloTargetArgs Empty = new SloTargetArgs();

    /**
     * List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
     * 
     */
    @Import(name="alertRules", required=true)
    private Output<List<SloTargetAlertRuleArgs>> alertRules;

    /**
     * @return List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
     * 
     */
    public Output<List<SloTargetAlertRuleArgs>> alertRules() {
        return this.alertRules;
    }

    /**
     * Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
     * 
     */
    @Import(name="compliancePeriod")
    private @Nullable Output<String> compliancePeriod;

    /**
     * @return Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
     * 
     */
    public Optional<Output<String>> compliancePeriod() {
        return Optional.ofNullable(this.compliancePeriod);
    }

    /**
     * It can be used to change the cycle start time. For example, you can specify sunday as the start of the week (instead of the default monday)
     * 
     */
    @Import(name="cycleStart")
    private @Nullable Output<String> cycleStart;

    /**
     * @return It can be used to change the cycle start time. For example, you can specify sunday as the start of the week (instead of the default monday)
     * 
     */
    public Optional<Output<String>> cycleStart() {
        return Optional.ofNullable(this.cycleStart);
    }

    /**
     * The cycle type of the calendar window, e.g. week, month.
     * 
     */
    @Import(name="cycleType")
    private @Nullable Output<String> cycleType;

    /**
     * @return The cycle type of the calendar window, e.g. week, month.
     * 
     */
    public Optional<Output<String>> cycleType() {
        return Optional.ofNullable(this.cycleType);
    }

    /**
     * Target value in the form of a percentage
     * 
     */
    @Import(name="slo", required=true)
    private Output<Double> slo;

    /**
     * @return Target value in the form of a percentage
     * 
     */
    public Output<Double> slo() {
        return this.slo;
    }

    /**
     * SLO target type can be the following type: `&#34;RollingWindow&#34;`, `&#34;CalendarWindow&#34;`
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return SLO target type can be the following type: `&#34;RollingWindow&#34;`, `&#34;CalendarWindow&#34;`
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private SloTargetArgs() {}

    private SloTargetArgs(SloTargetArgs $) {
        this.alertRules = $.alertRules;
        this.compliancePeriod = $.compliancePeriod;
        this.cycleStart = $.cycleStart;
        this.cycleType = $.cycleType;
        this.slo = $.slo;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SloTargetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SloTargetArgs $;

        public Builder() {
            $ = new SloTargetArgs();
        }

        public Builder(SloTargetArgs defaults) {
            $ = new SloTargetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alertRules List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
         * 
         * @return builder
         * 
         */
        public Builder alertRules(Output<List<SloTargetAlertRuleArgs>> alertRules) {
            $.alertRules = alertRules;
            return this;
        }

        /**
         * @param alertRules List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
         * 
         * @return builder
         * 
         */
        public Builder alertRules(List<SloTargetAlertRuleArgs> alertRules) {
            return alertRules(Output.of(alertRules));
        }

        /**
         * @param alertRules List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
         * 
         * @return builder
         * 
         */
        public Builder alertRules(SloTargetAlertRuleArgs... alertRules) {
            return alertRules(List.of(alertRules));
        }

        /**
         * @param compliancePeriod Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
         * 
         * @return builder
         * 
         */
        public Builder compliancePeriod(@Nullable Output<String> compliancePeriod) {
            $.compliancePeriod = compliancePeriod;
            return this;
        }

        /**
         * @param compliancePeriod Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
         * 
         * @return builder
         * 
         */
        public Builder compliancePeriod(String compliancePeriod) {
            return compliancePeriod(Output.of(compliancePeriod));
        }

        /**
         * @param cycleStart It can be used to change the cycle start time. For example, you can specify sunday as the start of the week (instead of the default monday)
         * 
         * @return builder
         * 
         */
        public Builder cycleStart(@Nullable Output<String> cycleStart) {
            $.cycleStart = cycleStart;
            return this;
        }

        /**
         * @param cycleStart It can be used to change the cycle start time. For example, you can specify sunday as the start of the week (instead of the default monday)
         * 
         * @return builder
         * 
         */
        public Builder cycleStart(String cycleStart) {
            return cycleStart(Output.of(cycleStart));
        }

        /**
         * @param cycleType The cycle type of the calendar window, e.g. week, month.
         * 
         * @return builder
         * 
         */
        public Builder cycleType(@Nullable Output<String> cycleType) {
            $.cycleType = cycleType;
            return this;
        }

        /**
         * @param cycleType The cycle type of the calendar window, e.g. week, month.
         * 
         * @return builder
         * 
         */
        public Builder cycleType(String cycleType) {
            return cycleType(Output.of(cycleType));
        }

        /**
         * @param slo Target value in the form of a percentage
         * 
         * @return builder
         * 
         */
        public Builder slo(Output<Double> slo) {
            $.slo = slo;
            return this;
        }

        /**
         * @param slo Target value in the form of a percentage
         * 
         * @return builder
         * 
         */
        public Builder slo(Double slo) {
            return slo(Output.of(slo));
        }

        /**
         * @param type SLO target type can be the following type: `&#34;RollingWindow&#34;`, `&#34;CalendarWindow&#34;`
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type SLO target type can be the following type: `&#34;RollingWindow&#34;`, `&#34;CalendarWindow&#34;`
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public SloTargetArgs build() {
            if ($.alertRules == null) {
                throw new MissingRequiredPropertyException("SloTargetArgs", "alertRules");
            }
            if ($.slo == null) {
                throw new MissingRequiredPropertyException("SloTargetArgs", "slo");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("SloTargetArgs", "type");
            }
            return $;
        }
    }

}
