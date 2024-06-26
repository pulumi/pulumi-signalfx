// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.servicenow.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IntegrationState extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationState Empty = new IntegrationState();

    /**
     * A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     * 
     */
    @Import(name="alertResolvedPayloadTemplate")
    private @Nullable Output<String> alertResolvedPayloadTemplate;

    /**
     * @return A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     * 
     */
    public Optional<Output<String>> alertResolvedPayloadTemplate() {
        return Optional.ofNullable(this.alertResolvedPayloadTemplate);
    }

    /**
     * A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     * 
     */
    @Import(name="alertTriggeredPayloadTemplate")
    private @Nullable Output<String> alertTriggeredPayloadTemplate;

    /**
     * @return A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     * 
     */
    public Optional<Output<String>> alertTriggeredPayloadTemplate() {
        return Optional.ofNullable(this.alertTriggeredPayloadTemplate);
    }

    /**
     * Whether the integration is enabled.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether the integration is enabled.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Name of the ServiceNow instance, for example `myinst.service-now.com`.
     * 
     */
    @Import(name="instanceName")
    private @Nullable Output<String> instanceName;

    /**
     * @return Name of the ServiceNow instance, for example `myinst.service-now.com`.
     * 
     */
    public Optional<Output<String>> instanceName() {
        return Optional.ofNullable(this.instanceName);
    }

    /**
     * The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
     * 
     */
    @Import(name="issueType")
    private @Nullable Output<String> issueType;

    /**
     * @return The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
     * 
     */
    public Optional<Output<String>> issueType() {
        return Optional.ofNullable(this.issueType);
    }

    /**
     * Name of the integration.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the integration.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Password used to authenticate the ServiceNow integration.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Password used to authenticate the ServiceNow integration.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * User name used to authenticate the ServiceNow integration.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return User name used to authenticate the ServiceNow integration.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private IntegrationState() {}

    private IntegrationState(IntegrationState $) {
        this.alertResolvedPayloadTemplate = $.alertResolvedPayloadTemplate;
        this.alertTriggeredPayloadTemplate = $.alertTriggeredPayloadTemplate;
        this.enabled = $.enabled;
        this.instanceName = $.instanceName;
        this.issueType = $.issueType;
        this.name = $.name;
        this.password = $.password;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationState $;

        public Builder() {
            $ = new IntegrationState();
        }

        public Builder(IntegrationState defaults) {
            $ = new IntegrationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alertResolvedPayloadTemplate A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
         * 
         * @return builder
         * 
         */
        public Builder alertResolvedPayloadTemplate(@Nullable Output<String> alertResolvedPayloadTemplate) {
            $.alertResolvedPayloadTemplate = alertResolvedPayloadTemplate;
            return this;
        }

        /**
         * @param alertResolvedPayloadTemplate A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
         * 
         * @return builder
         * 
         */
        public Builder alertResolvedPayloadTemplate(String alertResolvedPayloadTemplate) {
            return alertResolvedPayloadTemplate(Output.of(alertResolvedPayloadTemplate));
        }

        /**
         * @param alertTriggeredPayloadTemplate A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
         * 
         * @return builder
         * 
         */
        public Builder alertTriggeredPayloadTemplate(@Nullable Output<String> alertTriggeredPayloadTemplate) {
            $.alertTriggeredPayloadTemplate = alertTriggeredPayloadTemplate;
            return this;
        }

        /**
         * @param alertTriggeredPayloadTemplate A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
         * 
         * @return builder
         * 
         */
        public Builder alertTriggeredPayloadTemplate(String alertTriggeredPayloadTemplate) {
            return alertTriggeredPayloadTemplate(Output.of(alertTriggeredPayloadTemplate));
        }

        /**
         * @param enabled Whether the integration is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether the integration is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param instanceName Name of the ServiceNow instance, for example `myinst.service-now.com`.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(@Nullable Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName Name of the ServiceNow instance, for example `myinst.service-now.com`.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        /**
         * @param issueType The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
         * 
         * @return builder
         * 
         */
        public Builder issueType(@Nullable Output<String> issueType) {
            $.issueType = issueType;
            return this;
        }

        /**
         * @param issueType The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
         * 
         * @return builder
         * 
         */
        public Builder issueType(String issueType) {
            return issueType(Output.of(issueType));
        }

        /**
         * @param name Name of the integration.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the integration.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password Password used to authenticate the ServiceNow integration.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Password used to authenticate the ServiceNow integration.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param username User name used to authenticate the ServiceNow integration.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username User name used to authenticate the ServiceNow integration.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public IntegrationState build() {
            return $;
        }
    }

}
