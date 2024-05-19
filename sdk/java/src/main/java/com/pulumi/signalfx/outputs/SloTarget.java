// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.outputs.SloTargetAlertRule;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SloTarget {
    /**
     * @return SLO alert rules
     * 
     */
    private List<SloTargetAlertRule> alertRules;
    /**
     * @return (Required for `RollingWindow` type) Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
     * 
     */
    private @Nullable String compliancePeriod;
    /**
     * @return (Optional for `CalendarWindow` type)  It can be used to change the cycle start time. For example, you can specify sunday as the start of the week (instead of the default monday)
     * 
     */
    private @Nullable String cycleStart;
    /**
     * @return (Required for `CalendarWindow` type) The cycle type of the calendar window, e.g. week, month.
     * 
     */
    private @Nullable String cycleType;
    /**
     * @return Target value in the form of a percentage
     * 
     */
    private Double slo;
    /**
     * @return SLO target type can be the following type: `RollingWindow`
     * 
     */
    private String type;

    private SloTarget() {}
    /**
     * @return SLO alert rules
     * 
     */
    public List<SloTargetAlertRule> alertRules() {
        return this.alertRules;
    }
    /**
     * @return (Required for `RollingWindow` type) Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
     * 
     */
    public Optional<String> compliancePeriod() {
        return Optional.ofNullable(this.compliancePeriod);
    }
    /**
     * @return (Optional for `CalendarWindow` type)  It can be used to change the cycle start time. For example, you can specify sunday as the start of the week (instead of the default monday)
     * 
     */
    public Optional<String> cycleStart() {
        return Optional.ofNullable(this.cycleStart);
    }
    /**
     * @return (Required for `CalendarWindow` type) The cycle type of the calendar window, e.g. week, month.
     * 
     */
    public Optional<String> cycleType() {
        return Optional.ofNullable(this.cycleType);
    }
    /**
     * @return Target value in the form of a percentage
     * 
     */
    public Double slo() {
        return this.slo;
    }
    /**
     * @return SLO target type can be the following type: `RollingWindow`
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SloTarget defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<SloTargetAlertRule> alertRules;
        private @Nullable String compliancePeriod;
        private @Nullable String cycleStart;
        private @Nullable String cycleType;
        private Double slo;
        private String type;
        public Builder() {}
        public Builder(SloTarget defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.alertRules = defaults.alertRules;
    	      this.compliancePeriod = defaults.compliancePeriod;
    	      this.cycleStart = defaults.cycleStart;
    	      this.cycleType = defaults.cycleType;
    	      this.slo = defaults.slo;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder alertRules(List<SloTargetAlertRule> alertRules) {
            if (alertRules == null) {
              throw new MissingRequiredPropertyException("SloTarget", "alertRules");
            }
            this.alertRules = alertRules;
            return this;
        }
        public Builder alertRules(SloTargetAlertRule... alertRules) {
            return alertRules(List.of(alertRules));
        }
        @CustomType.Setter
        public Builder compliancePeriod(@Nullable String compliancePeriod) {

            this.compliancePeriod = compliancePeriod;
            return this;
        }
        @CustomType.Setter
        public Builder cycleStart(@Nullable String cycleStart) {

            this.cycleStart = cycleStart;
            return this;
        }
        @CustomType.Setter
        public Builder cycleType(@Nullable String cycleType) {

            this.cycleType = cycleType;
            return this;
        }
        @CustomType.Setter
        public Builder slo(Double slo) {
            if (slo == null) {
              throw new MissingRequiredPropertyException("SloTarget", "slo");
            }
            this.slo = slo;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("SloTarget", "type");
            }
            this.type = type;
            return this;
        }
        public SloTarget build() {
            final var _resultValue = new SloTarget();
            _resultValue.alertRules = alertRules;
            _resultValue.compliancePeriod = compliancePeriod;
            _resultValue.cycleStart = cycleStart;
            _resultValue.cycleType = cycleType;
            _resultValue.slo = slo;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
