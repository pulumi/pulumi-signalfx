// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.gcp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.gcp.inputs.IntegrationProjectServiceKeyArgs;
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
     * List of additional GCP service domain names that you want to monitor
     * 
     */
    @Import(name="customMetricTypeDomains")
    private @Nullable Output<List<String>> customMetricTypeDomains;

    /**
     * @return List of additional GCP service domain names that you want to monitor
     * 
     */
    public Optional<Output<List<String>>> customMetricTypeDomains() {
        return Optional.ofNullable(this.customMetricTypeDomains);
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
     * If enabled, SignalFx will sync also Google Cloud Metrics data. If disabled, SignalFx will import only metadata. Defaults
     * to true.
     * 
     */
    @Import(name="importGcpMetrics")
    private @Nullable Output<Boolean> importGcpMetrics;

    /**
     * @return If enabled, SignalFx will sync also Google Cloud Metrics data. If disabled, SignalFx will import only metadata. Defaults
     * to true.
     * 
     */
    public Optional<Output<Boolean>> importGcpMetrics() {
        return Optional.ofNullable(this.importGcpMetrics);
    }

    /**
     * [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
     * 
     */
    @Import(name="includeLists")
    private @Nullable Output<List<String>> includeLists;

    /**
     * @return [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
     * 
     */
    public Optional<Output<List<String>>> includeLists() {
        return Optional.ofNullable(this.includeLists);
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
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     * 
     */
    @Import(name="namedToken")
    private @Nullable Output<String> namedToken;

    /**
     * @return Name of the org token to be used for data ingestion. If not specified then default access token is used.
     * 
     */
    public Optional<Output<String>> namedToken() {
        return Optional.ofNullable(this.namedToken);
    }

    /**
     * GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     * 
     */
    @Import(name="pollRate")
    private @Nullable Output<Integer> pollRate;

    /**
     * @return GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     * 
     */
    public Optional<Output<Integer>> pollRate() {
        return Optional.ofNullable(this.pollRate);
    }

    /**
     * GCP projects to add.
     * 
     */
    @Import(name="projectServiceKeys")
    private @Nullable Output<List<IntegrationProjectServiceKeyArgs>> projectServiceKeys;

    /**
     * @return GCP projects to add.
     * 
     */
    public Optional<Output<List<IntegrationProjectServiceKeyArgs>>> projectServiceKeys() {
        return Optional.ofNullable(this.projectServiceKeys);
    }

    /**
     * GCP service metrics to import. Can be an empty list, or not included, to import &#39;All services&#39;. See the documentation for [Creating Integrations](https://dev.splunk.com/observability/reference/api/integrations/latest#endpoint-create-integration) for valid values.
     * 
     */
    @Import(name="services")
    private @Nullable Output<List<String>> services;

    /**
     * @return GCP service metrics to import. Can be an empty list, or not included, to import &#39;All services&#39;. See the documentation for [Creating Integrations](https://dev.splunk.com/observability/reference/api/integrations/latest#endpoint-create-integration) for valid values.
     * 
     */
    public Optional<Output<List<String>>> services() {
        return Optional.ofNullable(this.services);
    }

    /**
     * When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
     * 
     */
    @Import(name="useMetricSourceProjectForQuota")
    private @Nullable Output<Boolean> useMetricSourceProjectForQuota;

    /**
     * @return When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
     * 
     */
    public Optional<Output<Boolean>> useMetricSourceProjectForQuota() {
        return Optional.ofNullable(this.useMetricSourceProjectForQuota);
    }

    /**
     * [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
     * 
     * @deprecated
     * Please use include_list instead
     * 
     */
    @Deprecated /* Please use include_list instead */
    @Import(name="whitelists")
    private @Nullable Output<List<String>> whitelists;

    /**
     * @return [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
     * 
     * @deprecated
     * Please use include_list instead
     * 
     */
    @Deprecated /* Please use include_list instead */
    public Optional<Output<List<String>>> whitelists() {
        return Optional.ofNullable(this.whitelists);
    }

    private IntegrationState() {}

    private IntegrationState(IntegrationState $) {
        this.customMetricTypeDomains = $.customMetricTypeDomains;
        this.enabled = $.enabled;
        this.importGcpMetrics = $.importGcpMetrics;
        this.includeLists = $.includeLists;
        this.name = $.name;
        this.namedToken = $.namedToken;
        this.pollRate = $.pollRate;
        this.projectServiceKeys = $.projectServiceKeys;
        this.services = $.services;
        this.useMetricSourceProjectForQuota = $.useMetricSourceProjectForQuota;
        this.whitelists = $.whitelists;
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
         * @param customMetricTypeDomains List of additional GCP service domain names that you want to monitor
         * 
         * @return builder
         * 
         */
        public Builder customMetricTypeDomains(@Nullable Output<List<String>> customMetricTypeDomains) {
            $.customMetricTypeDomains = customMetricTypeDomains;
            return this;
        }

        /**
         * @param customMetricTypeDomains List of additional GCP service domain names that you want to monitor
         * 
         * @return builder
         * 
         */
        public Builder customMetricTypeDomains(List<String> customMetricTypeDomains) {
            return customMetricTypeDomains(Output.of(customMetricTypeDomains));
        }

        /**
         * @param customMetricTypeDomains List of additional GCP service domain names that you want to monitor
         * 
         * @return builder
         * 
         */
        public Builder customMetricTypeDomains(String... customMetricTypeDomains) {
            return customMetricTypeDomains(List.of(customMetricTypeDomains));
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
         * @param importGcpMetrics If enabled, SignalFx will sync also Google Cloud Metrics data. If disabled, SignalFx will import only metadata. Defaults
         * to true.
         * 
         * @return builder
         * 
         */
        public Builder importGcpMetrics(@Nullable Output<Boolean> importGcpMetrics) {
            $.importGcpMetrics = importGcpMetrics;
            return this;
        }

        /**
         * @param importGcpMetrics If enabled, SignalFx will sync also Google Cloud Metrics data. If disabled, SignalFx will import only metadata. Defaults
         * to true.
         * 
         * @return builder
         * 
         */
        public Builder importGcpMetrics(Boolean importGcpMetrics) {
            return importGcpMetrics(Output.of(importGcpMetrics));
        }

        /**
         * @param includeLists [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
         * 
         * @return builder
         * 
         */
        public Builder includeLists(@Nullable Output<List<String>> includeLists) {
            $.includeLists = includeLists;
            return this;
        }

        /**
         * @param includeLists [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
         * 
         * @return builder
         * 
         */
        public Builder includeLists(List<String> includeLists) {
            return includeLists(Output.of(includeLists));
        }

        /**
         * @param includeLists [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
         * 
         * @return builder
         * 
         */
        public Builder includeLists(String... includeLists) {
            return includeLists(List.of(includeLists));
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
         * @param namedToken Name of the org token to be used for data ingestion. If not specified then default access token is used.
         * 
         * @return builder
         * 
         */
        public Builder namedToken(@Nullable Output<String> namedToken) {
            $.namedToken = namedToken;
            return this;
        }

        /**
         * @param namedToken Name of the org token to be used for data ingestion. If not specified then default access token is used.
         * 
         * @return builder
         * 
         */
        public Builder namedToken(String namedToken) {
            return namedToken(Output.of(namedToken));
        }

        /**
         * @param pollRate GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
         * 
         * @return builder
         * 
         */
        public Builder pollRate(@Nullable Output<Integer> pollRate) {
            $.pollRate = pollRate;
            return this;
        }

        /**
         * @param pollRate GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
         * 
         * @return builder
         * 
         */
        public Builder pollRate(Integer pollRate) {
            return pollRate(Output.of(pollRate));
        }

        /**
         * @param projectServiceKeys GCP projects to add.
         * 
         * @return builder
         * 
         */
        public Builder projectServiceKeys(@Nullable Output<List<IntegrationProjectServiceKeyArgs>> projectServiceKeys) {
            $.projectServiceKeys = projectServiceKeys;
            return this;
        }

        /**
         * @param projectServiceKeys GCP projects to add.
         * 
         * @return builder
         * 
         */
        public Builder projectServiceKeys(List<IntegrationProjectServiceKeyArgs> projectServiceKeys) {
            return projectServiceKeys(Output.of(projectServiceKeys));
        }

        /**
         * @param projectServiceKeys GCP projects to add.
         * 
         * @return builder
         * 
         */
        public Builder projectServiceKeys(IntegrationProjectServiceKeyArgs... projectServiceKeys) {
            return projectServiceKeys(List.of(projectServiceKeys));
        }

        /**
         * @param services GCP service metrics to import. Can be an empty list, or not included, to import &#39;All services&#39;. See the documentation for [Creating Integrations](https://dev.splunk.com/observability/reference/api/integrations/latest#endpoint-create-integration) for valid values.
         * 
         * @return builder
         * 
         */
        public Builder services(@Nullable Output<List<String>> services) {
            $.services = services;
            return this;
        }

        /**
         * @param services GCP service metrics to import. Can be an empty list, or not included, to import &#39;All services&#39;. See the documentation for [Creating Integrations](https://dev.splunk.com/observability/reference/api/integrations/latest#endpoint-create-integration) for valid values.
         * 
         * @return builder
         * 
         */
        public Builder services(List<String> services) {
            return services(Output.of(services));
        }

        /**
         * @param services GCP service metrics to import. Can be an empty list, or not included, to import &#39;All services&#39;. See the documentation for [Creating Integrations](https://dev.splunk.com/observability/reference/api/integrations/latest#endpoint-create-integration) for valid values.
         * 
         * @return builder
         * 
         */
        public Builder services(String... services) {
            return services(List.of(services));
        }

        /**
         * @param useMetricSourceProjectForQuota When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
         * 
         * @return builder
         * 
         */
        public Builder useMetricSourceProjectForQuota(@Nullable Output<Boolean> useMetricSourceProjectForQuota) {
            $.useMetricSourceProjectForQuota = useMetricSourceProjectForQuota;
            return this;
        }

        /**
         * @param useMetricSourceProjectForQuota When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
         * 
         * @return builder
         * 
         */
        public Builder useMetricSourceProjectForQuota(Boolean useMetricSourceProjectForQuota) {
            return useMetricSourceProjectForQuota(Output.of(useMetricSourceProjectForQuota));
        }

        /**
         * @param whitelists [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
         * 
         * @return builder
         * 
         * @deprecated
         * Please use include_list instead
         * 
         */
        @Deprecated /* Please use include_list instead */
        public Builder whitelists(@Nullable Output<List<String>> whitelists) {
            $.whitelists = whitelists;
            return this;
        }

        /**
         * @param whitelists [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
         * 
         * @return builder
         * 
         * @deprecated
         * Please use include_list instead
         * 
         */
        @Deprecated /* Please use include_list instead */
        public Builder whitelists(List<String> whitelists) {
            return whitelists(Output.of(whitelists));
        }

        /**
         * @param whitelists [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
         * 
         * @return builder
         * 
         * @deprecated
         * Please use include_list instead
         * 
         */
        @Deprecated /* Please use include_list instead */
        public Builder whitelists(String... whitelists) {
            return whitelists(List.of(whitelists));
        }

        public IntegrationState build() {
            return $;
        }
    }

}
