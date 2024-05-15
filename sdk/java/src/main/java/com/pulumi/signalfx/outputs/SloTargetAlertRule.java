// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.outputs.SloTargetAlertRuleRule;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class SloTargetAlertRule {
    /**
     * @return Set of rules used for alerting
     * 
     */
    private List<SloTargetAlertRuleRule> rules;
    /**
     * @return Type of the SLO. Currently just: `&#34;RequestBased&#34;` is supported.
     * 
     */
    private String type;

    private SloTargetAlertRule() {}
    /**
     * @return Set of rules used for alerting
     * 
     */
    public List<SloTargetAlertRuleRule> rules() {
        return this.rules;
    }
    /**
     * @return Type of the SLO. Currently just: `&#34;RequestBased&#34;` is supported.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SloTargetAlertRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<SloTargetAlertRuleRule> rules;
        private String type;
        public Builder() {}
        public Builder(SloTargetAlertRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.rules = defaults.rules;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder rules(List<SloTargetAlertRuleRule> rules) {
            if (rules == null) {
              throw new MissingRequiredPropertyException("SloTargetAlertRule", "rules");
            }
            this.rules = rules;
            return this;
        }
        public Builder rules(SloTargetAlertRuleRule... rules) {
            return rules(List.of(rules));
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("SloTargetAlertRule", "type");
            }
            this.type = type;
            return this;
        }
        public SloTargetAlertRule build() {
            final var _resultValue = new SloTargetAlertRule();
            _resultValue.rules = rules;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
