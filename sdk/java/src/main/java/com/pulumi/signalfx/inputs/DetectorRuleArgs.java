// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.inputs.DetectorRuleReminderNotificationArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DetectorRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final DetectorRuleArgs Empty = new DetectorRuleArgs();

    /**
     * Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * A detect label which matches a detect label within `program_text`.
     * 
     */
    @Import(name="detectLabel", required=true)
    private Output<String> detectLabel;

    /**
     * @return A detect label which matches a detect label within `program_text`.
     * 
     */
    public Output<String> detectLabel() {
        return this.detectLabel;
    }

    /**
     * When true, notifications and events will not be generated for the detect label. `false` by default.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return When true, notifications and events will not be generated for the detect label. `false` by default.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See [Create A Single Detector](https://dev.splunk.com/observability/reference/api/detectors/latest) for more info.
     * 
     */
    @Import(name="notifications")
    private @Nullable Output<List<String>> notifications;

    /**
     * @return List of strings specifying where notifications will be sent when an incident occurs. See [Create A Single Detector](https://dev.splunk.com/observability/reference/api/detectors/latest) for more info.
     * 
     */
    public Optional<Output<List<String>>> notifications() {
        return Optional.ofNullable(this.notifications);
    }

    /**
     * Custom notification message body when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html) for more info.
     * 
     */
    @Import(name="parameterizedBody")
    private @Nullable Output<String> parameterizedBody;

    /**
     * @return Custom notification message body when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html) for more info.
     * 
     */
    public Optional<Output<String>> parameterizedBody() {
        return Optional.ofNullable(this.parameterizedBody);
    }

    /**
     * Custom notification message subject when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html) for more info.
     * 
     */
    @Import(name="parameterizedSubject")
    private @Nullable Output<String> parameterizedSubject;

    /**
     * @return Custom notification message subject when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html) for more info.
     * 
     */
    public Optional<Output<String>> parameterizedSubject() {
        return Optional.ofNullable(this.parameterizedSubject);
    }

    /**
     * Reminder notification in a detector rule lets you send multiple notifications for active alerts over a defined period of time. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    @Import(name="reminderNotification")
    private @Nullable Output<DetectorRuleReminderNotificationArgs> reminderNotification;

    /**
     * @return Reminder notification in a detector rule lets you send multiple notifications for active alerts over a defined period of time. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    public Optional<Output<DetectorRuleReminderNotificationArgs>> reminderNotification() {
        return Optional.ofNullable(this.reminderNotification);
    }

    /**
     * URL of page to consult when an alert is triggered. This can be used with custom notification messages.
     * 
     */
    @Import(name="runbookUrl")
    private @Nullable Output<String> runbookUrl;

    /**
     * @return URL of page to consult when an alert is triggered. This can be used with custom notification messages.
     * 
     */
    public Optional<Output<String>> runbookUrl() {
        return Optional.ofNullable(this.runbookUrl);
    }

    /**
     * The severity of the rule, must be one of: `&#34;Critical&#34;`, `&#34;Major&#34;`, `&#34;Minor&#34;`, `&#34;Warning&#34;`, `&#34;Info&#34;`.
     * 
     */
    @Import(name="severity", required=true)
    private Output<String> severity;

    /**
     * @return The severity of the rule, must be one of: `&#34;Critical&#34;`, `&#34;Major&#34;`, `&#34;Minor&#34;`, `&#34;Warning&#34;`, `&#34;Info&#34;`.
     * 
     */
    public Output<String> severity() {
        return this.severity;
    }

    /**
     * Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
     * 
     */
    @Import(name="tip")
    private @Nullable Output<String> tip;

    /**
     * @return Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
     * 
     */
    public Optional<Output<String>> tip() {
        return Optional.ofNullable(this.tip);
    }

    private DetectorRuleArgs() {}

    private DetectorRuleArgs(DetectorRuleArgs $) {
        this.description = $.description;
        this.detectLabel = $.detectLabel;
        this.disabled = $.disabled;
        this.notifications = $.notifications;
        this.parameterizedBody = $.parameterizedBody;
        this.parameterizedSubject = $.parameterizedSubject;
        this.reminderNotification = $.reminderNotification;
        this.runbookUrl = $.runbookUrl;
        this.severity = $.severity;
        this.tip = $.tip;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DetectorRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DetectorRuleArgs $;

        public Builder() {
            $ = new DetectorRuleArgs();
        }

        public Builder(DetectorRuleArgs defaults) {
            $ = new DetectorRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param detectLabel A detect label which matches a detect label within `program_text`.
         * 
         * @return builder
         * 
         */
        public Builder detectLabel(Output<String> detectLabel) {
            $.detectLabel = detectLabel;
            return this;
        }

        /**
         * @param detectLabel A detect label which matches a detect label within `program_text`.
         * 
         * @return builder
         * 
         */
        public Builder detectLabel(String detectLabel) {
            return detectLabel(Output.of(detectLabel));
        }

        /**
         * @param disabled When true, notifications and events will not be generated for the detect label. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled When true, notifications and events will not be generated for the detect label. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param notifications List of strings specifying where notifications will be sent when an incident occurs. See [Create A Single Detector](https://dev.splunk.com/observability/reference/api/detectors/latest) for more info.
         * 
         * @return builder
         * 
         */
        public Builder notifications(@Nullable Output<List<String>> notifications) {
            $.notifications = notifications;
            return this;
        }

        /**
         * @param notifications List of strings specifying where notifications will be sent when an incident occurs. See [Create A Single Detector](https://dev.splunk.com/observability/reference/api/detectors/latest) for more info.
         * 
         * @return builder
         * 
         */
        public Builder notifications(List<String> notifications) {
            return notifications(Output.of(notifications));
        }

        /**
         * @param notifications List of strings specifying where notifications will be sent when an incident occurs. See [Create A Single Detector](https://dev.splunk.com/observability/reference/api/detectors/latest) for more info.
         * 
         * @return builder
         * 
         */
        public Builder notifications(String... notifications) {
            return notifications(List.of(notifications));
        }

        /**
         * @param parameterizedBody Custom notification message body when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html) for more info.
         * 
         * @return builder
         * 
         */
        public Builder parameterizedBody(@Nullable Output<String> parameterizedBody) {
            $.parameterizedBody = parameterizedBody;
            return this;
        }

        /**
         * @param parameterizedBody Custom notification message body when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html) for more info.
         * 
         * @return builder
         * 
         */
        public Builder parameterizedBody(String parameterizedBody) {
            return parameterizedBody(Output.of(parameterizedBody));
        }

        /**
         * @param parameterizedSubject Custom notification message subject when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html) for more info.
         * 
         * @return builder
         * 
         */
        public Builder parameterizedSubject(@Nullable Output<String> parameterizedSubject) {
            $.parameterizedSubject = parameterizedSubject;
            return this;
        }

        /**
         * @param parameterizedSubject Custom notification message subject when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html) for more info.
         * 
         * @return builder
         * 
         */
        public Builder parameterizedSubject(String parameterizedSubject) {
            return parameterizedSubject(Output.of(parameterizedSubject));
        }

        /**
         * @param reminderNotification Reminder notification in a detector rule lets you send multiple notifications for active alerts over a defined period of time. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
         * 
         * @return builder
         * 
         */
        public Builder reminderNotification(@Nullable Output<DetectorRuleReminderNotificationArgs> reminderNotification) {
            $.reminderNotification = reminderNotification;
            return this;
        }

        /**
         * @param reminderNotification Reminder notification in a detector rule lets you send multiple notifications for active alerts over a defined period of time. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
         * 
         * @return builder
         * 
         */
        public Builder reminderNotification(DetectorRuleReminderNotificationArgs reminderNotification) {
            return reminderNotification(Output.of(reminderNotification));
        }

        /**
         * @param runbookUrl URL of page to consult when an alert is triggered. This can be used with custom notification messages.
         * 
         * @return builder
         * 
         */
        public Builder runbookUrl(@Nullable Output<String> runbookUrl) {
            $.runbookUrl = runbookUrl;
            return this;
        }

        /**
         * @param runbookUrl URL of page to consult when an alert is triggered. This can be used with custom notification messages.
         * 
         * @return builder
         * 
         */
        public Builder runbookUrl(String runbookUrl) {
            return runbookUrl(Output.of(runbookUrl));
        }

        /**
         * @param severity The severity of the rule, must be one of: `&#34;Critical&#34;`, `&#34;Major&#34;`, `&#34;Minor&#34;`, `&#34;Warning&#34;`, `&#34;Info&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder severity(Output<String> severity) {
            $.severity = severity;
            return this;
        }

        /**
         * @param severity The severity of the rule, must be one of: `&#34;Critical&#34;`, `&#34;Major&#34;`, `&#34;Minor&#34;`, `&#34;Warning&#34;`, `&#34;Info&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder severity(String severity) {
            return severity(Output.of(severity));
        }

        /**
         * @param tip Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
         * 
         * @return builder
         * 
         */
        public Builder tip(@Nullable Output<String> tip) {
            $.tip = tip;
            return this;
        }

        /**
         * @param tip Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
         * 
         * @return builder
         * 
         */
        public Builder tip(String tip) {
            return tip(Output.of(tip));
        }

        public DetectorRuleArgs build() {
            if ($.detectLabel == null) {
                throw new MissingRequiredPropertyException("DetectorRuleArgs", "detectLabel");
            }
            if ($.severity == null) {
                throw new MissingRequiredPropertyException("DetectorRuleArgs", "severity");
            }
            return $;
        }
    }

}
