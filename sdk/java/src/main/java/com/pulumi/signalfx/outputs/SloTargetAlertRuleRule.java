// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.outputs.SloTargetAlertRuleRuleParameters;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SloTargetAlertRuleRule {
    /**
     * @return Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
     * 
     */
    private @Nullable String description;
    /**
     * @return When true, notifications and events will not be generated for the detect label. `false` by default.
     * 
     */
    private @Nullable Boolean disabled;
    /**
     * @return List of strings specifying where notifications will be sent when an incident occurs. See [Create SLO](https://dev.splunk.com/observability/reference/api/slo/latest#endpoint-create-new-slo) for more info.
     * 
     */
    private @Nullable List<String> notifications;
    /**
     * @return Custom notification message body when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
     * 
     */
    private @Nullable String parameterizedBody;
    /**
     * @return Custom notification message subject when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
     * 
     */
    private @Nullable String parameterizedSubject;
    /**
     * @return Parameters for the SLO alert rule. Each SLO alert rule type accepts different parameters. If not specified, default parameters are used.
     * 
     */
    private @Nullable SloTargetAlertRuleRuleParameters parameters;
    /**
     * @return URL of page to consult when an alert is triggered. This can be used with custom notification messages.
     * 
     */
    private @Nullable String runbookUrl;
    /**
     * @return The severity of the rule, must be one of: `&#34;Critical&#34;`, `&#34;Major&#34;`, `&#34;Minor&#34;`, `&#34;Warning&#34;`, `&#34;Info&#34;`.
     * 
     */
    private String severity;
    /**
     * @return Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
     * 
     */
    private @Nullable String tip;

    private SloTargetAlertRuleRule() {}
    /**
     * @return Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return When true, notifications and events will not be generated for the detect label. `false` by default.
     * 
     */
    public Optional<Boolean> disabled() {
        return Optional.ofNullable(this.disabled);
    }
    /**
     * @return List of strings specifying where notifications will be sent when an incident occurs. See [Create SLO](https://dev.splunk.com/observability/reference/api/slo/latest#endpoint-create-new-slo) for more info.
     * 
     */
    public List<String> notifications() {
        return this.notifications == null ? List.of() : this.notifications;
    }
    /**
     * @return Custom notification message body when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
     * 
     */
    public Optional<String> parameterizedBody() {
        return Optional.ofNullable(this.parameterizedBody);
    }
    /**
     * @return Custom notification message subject when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
     * 
     */
    public Optional<String> parameterizedSubject() {
        return Optional.ofNullable(this.parameterizedSubject);
    }
    /**
     * @return Parameters for the SLO alert rule. Each SLO alert rule type accepts different parameters. If not specified, default parameters are used.
     * 
     */
    public Optional<SloTargetAlertRuleRuleParameters> parameters() {
        return Optional.ofNullable(this.parameters);
    }
    /**
     * @return URL of page to consult when an alert is triggered. This can be used with custom notification messages.
     * 
     */
    public Optional<String> runbookUrl() {
        return Optional.ofNullable(this.runbookUrl);
    }
    /**
     * @return The severity of the rule, must be one of: `&#34;Critical&#34;`, `&#34;Major&#34;`, `&#34;Minor&#34;`, `&#34;Warning&#34;`, `&#34;Info&#34;`.
     * 
     */
    public String severity() {
        return this.severity;
    }
    /**
     * @return Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
     * 
     */
    public Optional<String> tip() {
        return Optional.ofNullable(this.tip);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SloTargetAlertRuleRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable Boolean disabled;
        private @Nullable List<String> notifications;
        private @Nullable String parameterizedBody;
        private @Nullable String parameterizedSubject;
        private @Nullable SloTargetAlertRuleRuleParameters parameters;
        private @Nullable String runbookUrl;
        private String severity;
        private @Nullable String tip;
        public Builder() {}
        public Builder(SloTargetAlertRuleRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.disabled = defaults.disabled;
    	      this.notifications = defaults.notifications;
    	      this.parameterizedBody = defaults.parameterizedBody;
    	      this.parameterizedSubject = defaults.parameterizedSubject;
    	      this.parameters = defaults.parameters;
    	      this.runbookUrl = defaults.runbookUrl;
    	      this.severity = defaults.severity;
    	      this.tip = defaults.tip;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder disabled(@Nullable Boolean disabled) {

            this.disabled = disabled;
            return this;
        }
        @CustomType.Setter
        public Builder notifications(@Nullable List<String> notifications) {

            this.notifications = notifications;
            return this;
        }
        public Builder notifications(String... notifications) {
            return notifications(List.of(notifications));
        }
        @CustomType.Setter
        public Builder parameterizedBody(@Nullable String parameterizedBody) {

            this.parameterizedBody = parameterizedBody;
            return this;
        }
        @CustomType.Setter
        public Builder parameterizedSubject(@Nullable String parameterizedSubject) {

            this.parameterizedSubject = parameterizedSubject;
            return this;
        }
        @CustomType.Setter
        public Builder parameters(@Nullable SloTargetAlertRuleRuleParameters parameters) {

            this.parameters = parameters;
            return this;
        }
        @CustomType.Setter
        public Builder runbookUrl(@Nullable String runbookUrl) {

            this.runbookUrl = runbookUrl;
            return this;
        }
        @CustomType.Setter
        public Builder severity(String severity) {
            if (severity == null) {
              throw new MissingRequiredPropertyException("SloTargetAlertRuleRule", "severity");
            }
            this.severity = severity;
            return this;
        }
        @CustomType.Setter
        public Builder tip(@Nullable String tip) {

            this.tip = tip;
            return this;
        }
        public SloTargetAlertRuleRule build() {
            final var _resultValue = new SloTargetAlertRuleRule();
            _resultValue.description = description;
            _resultValue.disabled = disabled;
            _resultValue.notifications = notifications;
            _resultValue.parameterizedBody = parameterizedBody;
            _resultValue.parameterizedSubject = parameterizedSubject;
            _resultValue.parameters = parameters;
            _resultValue.runbookUrl = runbookUrl;
            _resultValue.severity = severity;
            _resultValue.tip = tip;
            return _resultValue;
        }
    }
}
