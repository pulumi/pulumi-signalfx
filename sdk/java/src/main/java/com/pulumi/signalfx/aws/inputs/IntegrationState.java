// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.aws.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.aws.inputs.IntegrationCustomNamespaceSyncRuleArgs;
import com.pulumi.signalfx.aws.inputs.IntegrationMetricStatsToSyncArgs;
import com.pulumi.signalfx.aws.inputs.IntegrationNamespaceSyncRuleArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IntegrationState extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationState Empty = new IntegrationState();

    /**
     * The mechanism used to authenticate with AWS. Use one of `signalfx.aws.ExternalIntegration` or
     * `signalfx.aws.TokenIntegration` to define this
     * 
     */
    @Import(name="authMethod")
    private @Nullable Output<String> authMethod;

    /**
     * @return The mechanism used to authenticate with AWS. Use one of `signalfx.aws.ExternalIntegration` or
     * `signalfx.aws.TokenIntegration` to define this
     * 
     */
    public Optional<Output<String>> authMethod() {
        return Optional.ofNullable(this.authMethod);
    }

    /**
     * List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS;
     * Splunk Observability imports the metrics so you can monitor them.
     * 
     */
    @Import(name="customCloudwatchNamespaces")
    private @Nullable Output<List<String>> customCloudwatchNamespaces;

    /**
     * @return List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS;
     * Splunk Observability imports the metrics so you can monitor them.
     * 
     */
    public Optional<Output<List<String>>> customCloudwatchNamespaces() {
        return Optional.ofNullable(this.customCloudwatchNamespaces);
    }

    /**
     * Each element controls the data collected by Splunk Observability for the specified namespace. If you specify this
     * property, Splunk Observability ignores values in the &#34;custom_cloudwatch_namespaces&#34; property.
     * 
     */
    @Import(name="customNamespaceSyncRules")
    private @Nullable Output<List<IntegrationCustomNamespaceSyncRuleArgs>> customNamespaceSyncRules;

    /**
     * @return Each element controls the data collected by Splunk Observability for the specified namespace. If you specify this
     * property, Splunk Observability ignores values in the &#34;custom_cloudwatch_namespaces&#34; property.
     * 
     */
    public Optional<Output<List<IntegrationCustomNamespaceSyncRuleArgs>>> customNamespaceSyncRules() {
        return Optional.ofNullable(this.customNamespaceSyncRules);
    }

    /**
     * Flag that controls how Splunk Observability imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`,
     * Splunk Observability imports the metrics.
     * 
     */
    @Import(name="enableAwsUsage")
    private @Nullable Output<Boolean> enableAwsUsage;

    /**
     * @return Flag that controls how Splunk Observability imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`,
     * Splunk Observability imports the metrics.
     * 
     */
    public Optional<Output<Boolean>> enableAwsUsage() {
        return Optional.ofNullable(this.enableAwsUsage);
    }

    /**
     * Controls how Splunk Observability checks for large amounts of data for this AWS integration. If true, Splunk
     * Observability monitors the amount of data coming in from the integration.
     * 
     */
    @Import(name="enableCheckLargeVolume")
    private @Nullable Output<Boolean> enableCheckLargeVolume;

    /**
     * @return Controls how Splunk Observability checks for large amounts of data for this AWS integration. If true, Splunk
     * Observability monitors the amount of data coming in from the integration.
     * 
     */
    public Optional<Output<Boolean>> enableCheckLargeVolume() {
        return Optional.ofNullable(this.enableCheckLargeVolume);
    }

    /**
     * Enables AWS logs synchronization.
     * 
     */
    @Import(name="enableLogsSync")
    private @Nullable Output<Boolean> enableLogsSync;

    /**
     * @return Enables AWS logs synchronization.
     * 
     */
    public Optional<Output<Boolean>> enableLogsSync() {
        return Optional.ofNullable(this.enableLogsSync);
    }

    /**
     * Whether the integration is enabled or not
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether the integration is enabled or not
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * Used with `signalfx.aws.ExternalIntegration`. Use this property to specify the external id.
     * 
     */
    @Import(name="externalId")
    private @Nullable Output<String> externalId;

    /**
     * @return Used with `signalfx.aws.ExternalIntegration`. Use this property to specify the external id.
     * 
     */
    public Optional<Output<String>> externalId() {
        return Optional.ofNullable(this.externalId);
    }

    /**
     * Flag that controls how Splunk Observability imports Cloud Watch metrics. If true, Splunk Observability imports Cloud
     * Watch metrics from AWS.
     * 
     */
    @Import(name="importCloudWatch")
    private @Nullable Output<Boolean> importCloudWatch;

    /**
     * @return Flag that controls how Splunk Observability imports Cloud Watch metrics. If true, Splunk Observability imports Cloud
     * Watch metrics from AWS.
     * 
     */
    public Optional<Output<Boolean>> importCloudWatch() {
        return Optional.ofNullable(this.importCloudWatch);
    }

    /**
     * The ID of this integration
     * 
     */
    @Import(name="integrationId")
    private @Nullable Output<String> integrationId;

    /**
     * @return The ID of this integration
     * 
     */
    public Optional<Output<String>> integrationId() {
        return Optional.ofNullable(this.integrationId);
    }

    /**
     * Used with `signalfx.aws.TokenIntegration`. Use this property to specify the token.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return Used with `signalfx.aws.TokenIntegration`. Use this property to specify the token.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics
     * that Splunk Observability collects for this metric. If you specify this property, Splunk Observability retrieves only
     * specified AWS statistics. If you don&#39;t specify this property, Splunk Observability retrieves the AWS standard set of
     * statistics.
     * 
     */
    @Import(name="metricStatsToSyncs")
    private @Nullable Output<List<IntegrationMetricStatsToSyncArgs>> metricStatsToSyncs;

    /**
     * @return Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics
     * that Splunk Observability collects for this metric. If you specify this property, Splunk Observability retrieves only
     * specified AWS statistics. If you don&#39;t specify this property, Splunk Observability retrieves the AWS standard set of
     * statistics.
     * 
     */
    public Optional<Output<List<IntegrationMetricStatsToSyncArgs>>> metricStatsToSyncs() {
        return Optional.ofNullable(this.metricStatsToSyncs);
    }

    /**
     * Name of the integration. Please specify the name in `signalfx.aws.ExternalIntegration` or
     * `signalfx_aws_integration_token`
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the integration. Please specify the name in `signalfx.aws.ExternalIntegration` or
     * `signalfx_aws_integration_token`
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A named token to use for ingest
     * 
     */
    @Import(name="namedToken")
    private @Nullable Output<String> namedToken;

    /**
     * @return A named token to use for ingest
     * 
     */
    public Optional<Output<String>> namedToken() {
        return Optional.ofNullable(this.namedToken);
    }

    /**
     * Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that
     * Splunk Observability collects for the namespace. If you specify this property, Splunk Observability ignores the values
     * in the AWS CloudWatch Integration Model &#34;services&#34; property. If you don&#39;t specify either property, Splunk Observability
     * syncs all data in all AWS namespaces.
     * 
     */
    @Import(name="namespaceSyncRules")
    private @Nullable Output<List<IntegrationNamespaceSyncRuleArgs>> namespaceSyncRules;

    /**
     * @return Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that
     * Splunk Observability collects for the namespace. If you specify this property, Splunk Observability ignores the values
     * in the AWS CloudWatch Integration Model &#34;services&#34; property. If you don&#39;t specify either property, Splunk Observability
     * syncs all data in all AWS namespaces.
     * 
     */
    public Optional<Output<List<IntegrationNamespaceSyncRuleArgs>>> namespaceSyncRules() {
        return Optional.ofNullable(this.namespaceSyncRules);
    }

    /**
     * AWS poll rate (in seconds). Between `60` and `600`.
     * 
     */
    @Import(name="pollRate")
    private @Nullable Output<Integer> pollRate;

    /**
     * @return AWS poll rate (in seconds). Between `60` and `600`.
     * 
     */
    public Optional<Output<Integer>> pollRate() {
        return Optional.ofNullable(this.pollRate);
    }

    /**
     * List of AWS regions that Splunk Observability should monitor.
     * 
     */
    @Import(name="regions")
    private @Nullable Output<List<String>> regions;

    /**
     * @return List of AWS regions that Splunk Observability should monitor.
     * 
     */
    public Optional<Output<List<String>>> regions() {
        return Optional.ofNullable(this.regions);
    }

    /**
     * Used with `signalfx.aws.ExternalIntegration`. Use this property to specify the AIM role ARN.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return Used with `signalfx.aws.ExternalIntegration`. Use this property to specify the AIM role ARN.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * List of AWS services that you want Splunk Observability to monitor. Each element is a string designating an AWS service.
     * 
     */
    @Import(name="services")
    private @Nullable Output<List<String>> services;

    /**
     * @return List of AWS services that you want Splunk Observability to monitor. Each element is a string designating an AWS service.
     * 
     */
    public Optional<Output<List<String>>> services() {
        return Optional.ofNullable(this.services);
    }

    /**
     * Indicates that Splunk Observability should sync metrics and metadata from custom AWS namespaces only (see the
     * `custom_namespace_sync_rule` field for details). Defaults to `false`.
     * 
     */
    @Import(name="syncCustomNamespacesOnly")
    private @Nullable Output<Boolean> syncCustomNamespacesOnly;

    /**
     * @return Indicates that Splunk Observability should sync metrics and metadata from custom AWS namespaces only (see the
     * `custom_namespace_sync_rule` field for details). Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> syncCustomNamespacesOnly() {
        return Optional.ofNullable(this.syncCustomNamespacesOnly);
    }

    /**
     * Used with `signalfx.aws.TokenIntegration`. Use this property to specify the token.
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return Used with `signalfx.aws.TokenIntegration`. Use this property to specify the token.
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    /**
     * Enables the use of Cloudwatch Metric Streams for metrics synchronization.
     * 
     */
    @Import(name="useMetricStreamsSync")
    private @Nullable Output<Boolean> useMetricStreamsSync;

    /**
     * @return Enables the use of Cloudwatch Metric Streams for metrics synchronization.
     * 
     */
    public Optional<Output<Boolean>> useMetricStreamsSync() {
        return Optional.ofNullable(this.useMetricStreamsSync);
    }

    private IntegrationState() {}

    private IntegrationState(IntegrationState $) {
        this.authMethod = $.authMethod;
        this.customCloudwatchNamespaces = $.customCloudwatchNamespaces;
        this.customNamespaceSyncRules = $.customNamespaceSyncRules;
        this.enableAwsUsage = $.enableAwsUsage;
        this.enableCheckLargeVolume = $.enableCheckLargeVolume;
        this.enableLogsSync = $.enableLogsSync;
        this.enabled = $.enabled;
        this.externalId = $.externalId;
        this.importCloudWatch = $.importCloudWatch;
        this.integrationId = $.integrationId;
        this.key = $.key;
        this.metricStatsToSyncs = $.metricStatsToSyncs;
        this.name = $.name;
        this.namedToken = $.namedToken;
        this.namespaceSyncRules = $.namespaceSyncRules;
        this.pollRate = $.pollRate;
        this.regions = $.regions;
        this.roleArn = $.roleArn;
        this.services = $.services;
        this.syncCustomNamespacesOnly = $.syncCustomNamespacesOnly;
        this.token = $.token;
        this.useMetricStreamsSync = $.useMetricStreamsSync;
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
         * @param authMethod The mechanism used to authenticate with AWS. Use one of `signalfx.aws.ExternalIntegration` or
         * `signalfx.aws.TokenIntegration` to define this
         * 
         * @return builder
         * 
         */
        public Builder authMethod(@Nullable Output<String> authMethod) {
            $.authMethod = authMethod;
            return this;
        }

        /**
         * @param authMethod The mechanism used to authenticate with AWS. Use one of `signalfx.aws.ExternalIntegration` or
         * `signalfx.aws.TokenIntegration` to define this
         * 
         * @return builder
         * 
         */
        public Builder authMethod(String authMethod) {
            return authMethod(Output.of(authMethod));
        }

        /**
         * @param customCloudwatchNamespaces List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS;
         * Splunk Observability imports the metrics so you can monitor them.
         * 
         * @return builder
         * 
         */
        public Builder customCloudwatchNamespaces(@Nullable Output<List<String>> customCloudwatchNamespaces) {
            $.customCloudwatchNamespaces = customCloudwatchNamespaces;
            return this;
        }

        /**
         * @param customCloudwatchNamespaces List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS;
         * Splunk Observability imports the metrics so you can monitor them.
         * 
         * @return builder
         * 
         */
        public Builder customCloudwatchNamespaces(List<String> customCloudwatchNamespaces) {
            return customCloudwatchNamespaces(Output.of(customCloudwatchNamespaces));
        }

        /**
         * @param customCloudwatchNamespaces List of custom AWS CloudWatch namespaces to monitor. Custom namespaces contain custom metrics that you define in AWS;
         * Splunk Observability imports the metrics so you can monitor them.
         * 
         * @return builder
         * 
         */
        public Builder customCloudwatchNamespaces(String... customCloudwatchNamespaces) {
            return customCloudwatchNamespaces(List.of(customCloudwatchNamespaces));
        }

        /**
         * @param customNamespaceSyncRules Each element controls the data collected by Splunk Observability for the specified namespace. If you specify this
         * property, Splunk Observability ignores values in the &#34;custom_cloudwatch_namespaces&#34; property.
         * 
         * @return builder
         * 
         */
        public Builder customNamespaceSyncRules(@Nullable Output<List<IntegrationCustomNamespaceSyncRuleArgs>> customNamespaceSyncRules) {
            $.customNamespaceSyncRules = customNamespaceSyncRules;
            return this;
        }

        /**
         * @param customNamespaceSyncRules Each element controls the data collected by Splunk Observability for the specified namespace. If you specify this
         * property, Splunk Observability ignores values in the &#34;custom_cloudwatch_namespaces&#34; property.
         * 
         * @return builder
         * 
         */
        public Builder customNamespaceSyncRules(List<IntegrationCustomNamespaceSyncRuleArgs> customNamespaceSyncRules) {
            return customNamespaceSyncRules(Output.of(customNamespaceSyncRules));
        }

        /**
         * @param customNamespaceSyncRules Each element controls the data collected by Splunk Observability for the specified namespace. If you specify this
         * property, Splunk Observability ignores values in the &#34;custom_cloudwatch_namespaces&#34; property.
         * 
         * @return builder
         * 
         */
        public Builder customNamespaceSyncRules(IntegrationCustomNamespaceSyncRuleArgs... customNamespaceSyncRules) {
            return customNamespaceSyncRules(List.of(customNamespaceSyncRules));
        }

        /**
         * @param enableAwsUsage Flag that controls how Splunk Observability imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`,
         * Splunk Observability imports the metrics.
         * 
         * @return builder
         * 
         */
        public Builder enableAwsUsage(@Nullable Output<Boolean> enableAwsUsage) {
            $.enableAwsUsage = enableAwsUsage;
            return this;
        }

        /**
         * @param enableAwsUsage Flag that controls how Splunk Observability imports usage metrics from AWS to use with AWS Cost Optimizer. If `true`,
         * Splunk Observability imports the metrics.
         * 
         * @return builder
         * 
         */
        public Builder enableAwsUsage(Boolean enableAwsUsage) {
            return enableAwsUsage(Output.of(enableAwsUsage));
        }

        /**
         * @param enableCheckLargeVolume Controls how Splunk Observability checks for large amounts of data for this AWS integration. If true, Splunk
         * Observability monitors the amount of data coming in from the integration.
         * 
         * @return builder
         * 
         */
        public Builder enableCheckLargeVolume(@Nullable Output<Boolean> enableCheckLargeVolume) {
            $.enableCheckLargeVolume = enableCheckLargeVolume;
            return this;
        }

        /**
         * @param enableCheckLargeVolume Controls how Splunk Observability checks for large amounts of data for this AWS integration. If true, Splunk
         * Observability monitors the amount of data coming in from the integration.
         * 
         * @return builder
         * 
         */
        public Builder enableCheckLargeVolume(Boolean enableCheckLargeVolume) {
            return enableCheckLargeVolume(Output.of(enableCheckLargeVolume));
        }

        /**
         * @param enableLogsSync Enables AWS logs synchronization.
         * 
         * @return builder
         * 
         */
        public Builder enableLogsSync(@Nullable Output<Boolean> enableLogsSync) {
            $.enableLogsSync = enableLogsSync;
            return this;
        }

        /**
         * @param enableLogsSync Enables AWS logs synchronization.
         * 
         * @return builder
         * 
         */
        public Builder enableLogsSync(Boolean enableLogsSync) {
            return enableLogsSync(Output.of(enableLogsSync));
        }

        /**
         * @param enabled Whether the integration is enabled or not
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether the integration is enabled or not
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param externalId Used with `signalfx.aws.ExternalIntegration`. Use this property to specify the external id.
         * 
         * @return builder
         * 
         */
        public Builder externalId(@Nullable Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId Used with `signalfx.aws.ExternalIntegration`. Use this property to specify the external id.
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
        }

        /**
         * @param importCloudWatch Flag that controls how Splunk Observability imports Cloud Watch metrics. If true, Splunk Observability imports Cloud
         * Watch metrics from AWS.
         * 
         * @return builder
         * 
         */
        public Builder importCloudWatch(@Nullable Output<Boolean> importCloudWatch) {
            $.importCloudWatch = importCloudWatch;
            return this;
        }

        /**
         * @param importCloudWatch Flag that controls how Splunk Observability imports Cloud Watch metrics. If true, Splunk Observability imports Cloud
         * Watch metrics from AWS.
         * 
         * @return builder
         * 
         */
        public Builder importCloudWatch(Boolean importCloudWatch) {
            return importCloudWatch(Output.of(importCloudWatch));
        }

        /**
         * @param integrationId The ID of this integration
         * 
         * @return builder
         * 
         */
        public Builder integrationId(@Nullable Output<String> integrationId) {
            $.integrationId = integrationId;
            return this;
        }

        /**
         * @param integrationId The ID of this integration
         * 
         * @return builder
         * 
         */
        public Builder integrationId(String integrationId) {
            return integrationId(Output.of(integrationId));
        }

        /**
         * @param key Used with `signalfx.aws.TokenIntegration`. Use this property to specify the token.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Used with `signalfx.aws.TokenIntegration`. Use this property to specify the token.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param metricStatsToSyncs Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics
         * that Splunk Observability collects for this metric. If you specify this property, Splunk Observability retrieves only
         * specified AWS statistics. If you don&#39;t specify this property, Splunk Observability retrieves the AWS standard set of
         * statistics.
         * 
         * @return builder
         * 
         */
        public Builder metricStatsToSyncs(@Nullable Output<List<IntegrationMetricStatsToSyncArgs>> metricStatsToSyncs) {
            $.metricStatsToSyncs = metricStatsToSyncs;
            return this;
        }

        /**
         * @param metricStatsToSyncs Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics
         * that Splunk Observability collects for this metric. If you specify this property, Splunk Observability retrieves only
         * specified AWS statistics. If you don&#39;t specify this property, Splunk Observability retrieves the AWS standard set of
         * statistics.
         * 
         * @return builder
         * 
         */
        public Builder metricStatsToSyncs(List<IntegrationMetricStatsToSyncArgs> metricStatsToSyncs) {
            return metricStatsToSyncs(Output.of(metricStatsToSyncs));
        }

        /**
         * @param metricStatsToSyncs Each element in the array is an object that contains an AWS namespace name, AWS metric name and a list of statistics
         * that Splunk Observability collects for this metric. If you specify this property, Splunk Observability retrieves only
         * specified AWS statistics. If you don&#39;t specify this property, Splunk Observability retrieves the AWS standard set of
         * statistics.
         * 
         * @return builder
         * 
         */
        public Builder metricStatsToSyncs(IntegrationMetricStatsToSyncArgs... metricStatsToSyncs) {
            return metricStatsToSyncs(List.of(metricStatsToSyncs));
        }

        /**
         * @param name Name of the integration. Please specify the name in `signalfx.aws.ExternalIntegration` or
         * `signalfx_aws_integration_token`
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the integration. Please specify the name in `signalfx.aws.ExternalIntegration` or
         * `signalfx_aws_integration_token`
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namedToken A named token to use for ingest
         * 
         * @return builder
         * 
         */
        public Builder namedToken(@Nullable Output<String> namedToken) {
            $.namedToken = namedToken;
            return this;
        }

        /**
         * @param namedToken A named token to use for ingest
         * 
         * @return builder
         * 
         */
        public Builder namedToken(String namedToken) {
            return namedToken(Output.of(namedToken));
        }

        /**
         * @param namespaceSyncRules Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that
         * Splunk Observability collects for the namespace. If you specify this property, Splunk Observability ignores the values
         * in the AWS CloudWatch Integration Model &#34;services&#34; property. If you don&#39;t specify either property, Splunk Observability
         * syncs all data in all AWS namespaces.
         * 
         * @return builder
         * 
         */
        public Builder namespaceSyncRules(@Nullable Output<List<IntegrationNamespaceSyncRuleArgs>> namespaceSyncRules) {
            $.namespaceSyncRules = namespaceSyncRules;
            return this;
        }

        /**
         * @param namespaceSyncRules Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that
         * Splunk Observability collects for the namespace. If you specify this property, Splunk Observability ignores the values
         * in the AWS CloudWatch Integration Model &#34;services&#34; property. If you don&#39;t specify either property, Splunk Observability
         * syncs all data in all AWS namespaces.
         * 
         * @return builder
         * 
         */
        public Builder namespaceSyncRules(List<IntegrationNamespaceSyncRuleArgs> namespaceSyncRules) {
            return namespaceSyncRules(Output.of(namespaceSyncRules));
        }

        /**
         * @param namespaceSyncRules Each element in the array is an object that contains an AWS namespace name and a filter that controls the data that
         * Splunk Observability collects for the namespace. If you specify this property, Splunk Observability ignores the values
         * in the AWS CloudWatch Integration Model &#34;services&#34; property. If you don&#39;t specify either property, Splunk Observability
         * syncs all data in all AWS namespaces.
         * 
         * @return builder
         * 
         */
        public Builder namespaceSyncRules(IntegrationNamespaceSyncRuleArgs... namespaceSyncRules) {
            return namespaceSyncRules(List.of(namespaceSyncRules));
        }

        /**
         * @param pollRate AWS poll rate (in seconds). Between `60` and `600`.
         * 
         * @return builder
         * 
         */
        public Builder pollRate(@Nullable Output<Integer> pollRate) {
            $.pollRate = pollRate;
            return this;
        }

        /**
         * @param pollRate AWS poll rate (in seconds). Between `60` and `600`.
         * 
         * @return builder
         * 
         */
        public Builder pollRate(Integer pollRate) {
            return pollRate(Output.of(pollRate));
        }

        /**
         * @param regions List of AWS regions that Splunk Observability should monitor.
         * 
         * @return builder
         * 
         */
        public Builder regions(@Nullable Output<List<String>> regions) {
            $.regions = regions;
            return this;
        }

        /**
         * @param regions List of AWS regions that Splunk Observability should monitor.
         * 
         * @return builder
         * 
         */
        public Builder regions(List<String> regions) {
            return regions(Output.of(regions));
        }

        /**
         * @param regions List of AWS regions that Splunk Observability should monitor.
         * 
         * @return builder
         * 
         */
        public Builder regions(String... regions) {
            return regions(List.of(regions));
        }

        /**
         * @param roleArn Used with `signalfx.aws.ExternalIntegration`. Use this property to specify the AIM role ARN.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn Used with `signalfx.aws.ExternalIntegration`. Use this property to specify the AIM role ARN.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param services List of AWS services that you want Splunk Observability to monitor. Each element is a string designating an AWS service.
         * 
         * @return builder
         * 
         */
        public Builder services(@Nullable Output<List<String>> services) {
            $.services = services;
            return this;
        }

        /**
         * @param services List of AWS services that you want Splunk Observability to monitor. Each element is a string designating an AWS service.
         * 
         * @return builder
         * 
         */
        public Builder services(List<String> services) {
            return services(Output.of(services));
        }

        /**
         * @param services List of AWS services that you want Splunk Observability to monitor. Each element is a string designating an AWS service.
         * 
         * @return builder
         * 
         */
        public Builder services(String... services) {
            return services(List.of(services));
        }

        /**
         * @param syncCustomNamespacesOnly Indicates that Splunk Observability should sync metrics and metadata from custom AWS namespaces only (see the
         * `custom_namespace_sync_rule` field for details). Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder syncCustomNamespacesOnly(@Nullable Output<Boolean> syncCustomNamespacesOnly) {
            $.syncCustomNamespacesOnly = syncCustomNamespacesOnly;
            return this;
        }

        /**
         * @param syncCustomNamespacesOnly Indicates that Splunk Observability should sync metrics and metadata from custom AWS namespaces only (see the
         * `custom_namespace_sync_rule` field for details). Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder syncCustomNamespacesOnly(Boolean syncCustomNamespacesOnly) {
            return syncCustomNamespacesOnly(Output.of(syncCustomNamespacesOnly));
        }

        /**
         * @param token Used with `signalfx.aws.TokenIntegration`. Use this property to specify the token.
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token Used with `signalfx.aws.TokenIntegration`. Use this property to specify the token.
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        /**
         * @param useMetricStreamsSync Enables the use of Cloudwatch Metric Streams for metrics synchronization.
         * 
         * @return builder
         * 
         */
        public Builder useMetricStreamsSync(@Nullable Output<Boolean> useMetricStreamsSync) {
            $.useMetricStreamsSync = useMetricStreamsSync;
            return this;
        }

        /**
         * @param useMetricStreamsSync Enables the use of Cloudwatch Metric Streams for metrics synchronization.
         * 
         * @return builder
         * 
         */
        public Builder useMetricStreamsSync(Boolean useMetricStreamsSync) {
            return useMetricStreamsSync(Output.of(useMetricStreamsSync));
        }

        public IntegrationState build() {
            return $;
        }
    }

}
