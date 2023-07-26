// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.azure.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.azure.inputs.IntegrationCustomNamespacesPerServiceArgs;
import com.pulumi.signalfx.azure.inputs.IntegrationResourceFilterRuleArgs;
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
     * Additional Azure resource types that you want to sync with Observability Cloud.
     * 
     */
    @Import(name="additionalServices")
    private @Nullable Output<List<String>> additionalServices;

    /**
     * @return Additional Azure resource types that you want to sync with Observability Cloud.
     * 
     */
    public Optional<Output<List<String>>> additionalServices() {
        return Optional.ofNullable(this.additionalServices);
    }

    /**
     * Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
     * 
     */
    @Import(name="appId")
    private @Nullable Output<String> appId;

    /**
     * @return Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
     * 
     */
    public Optional<Output<String>> appId() {
        return Optional.ofNullable(this.appId);
    }

    /**
     * Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `sync_guest_os_namespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
     * 
     */
    @Import(name="customNamespacesPerServices")
    private @Nullable Output<List<IntegrationCustomNamespacesPerServiceArgs>> customNamespacesPerServices;

    /**
     * @return Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `sync_guest_os_namespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
     * 
     */
    public Optional<Output<List<IntegrationCustomNamespacesPerServiceArgs>>> customNamespacesPerServices() {
        return Optional.ofNullable(this.customNamespacesPerServices);
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
     * What type of Azure integration this is. The allowed values are `\&#34;azure_us_government\&#34;` and `\&#34;azure\&#34;`. Defaults to `\&#34;azure\&#34;`.
     * 
     */
    @Import(name="environment")
    private @Nullable Output<String> environment;

    /**
     * @return What type of Azure integration this is. The allowed values are `\&#34;azure_us_government\&#34;` and `\&#34;azure\&#34;`. Defaults to `\&#34;azure\&#34;`.
     * 
     */
    public Optional<Output<String>> environment() {
        return Optional.ofNullable(this.environment);
    }

    /**
     * If enabled, SignalFx will sync also Azure Monitor data. If disabled, SignalFx will import only metadata. Defaults to true.
     * 
     */
    @Import(name="importAzureMonitor")
    private @Nullable Output<Boolean> importAzureMonitor;

    /**
     * @return If enabled, SignalFx will sync also Azure Monitor data. If disabled, SignalFx will import only metadata. Defaults to true.
     * 
     */
    public Optional<Output<Boolean>> importAzureMonitor() {
        return Optional.ofNullable(this.importAzureMonitor);
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
     * Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     * 
     */
    @Import(name="pollRate")
    private @Nullable Output<Integer> pollRate;

    /**
     * @return Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     * 
     */
    public Optional<Output<Integer>> pollRate() {
        return Optional.ofNullable(this.pollRate);
    }

    /**
     * List of rules for filtering Azure resources by their tags.
     * 
     */
    @Import(name="resourceFilterRules")
    private @Nullable Output<List<IntegrationResourceFilterRuleArgs>> resourceFilterRules;

    /**
     * @return List of rules for filtering Azure resources by their tags.
     * 
     */
    public Optional<Output<List<IntegrationResourceFilterRuleArgs>>> resourceFilterRules() {
        return Optional.ofNullable(this.resourceFilterRules);
    }

    /**
     * Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     * 
     */
    @Import(name="secretKey")
    private @Nullable Output<String> secretKey;

    /**
     * @return Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     * 
     */
    public Optional<Output<String>> secretKey() {
        return Optional.ofNullable(this.secretKey);
    }

    /**
     * List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
     * 
     */
    @Import(name="services")
    private @Nullable Output<List<String>> services;

    /**
     * @return List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
     * 
     */
    public Optional<Output<List<String>>> services() {
        return Optional.ofNullable(this.services);
    }

    /**
     * List of Azure subscriptions that SignalFx should monitor.
     * 
     */
    @Import(name="subscriptions")
    private @Nullable Output<List<String>> subscriptions;

    /**
     * @return List of Azure subscriptions that SignalFx should monitor.
     * 
     */
    public Optional<Output<List<String>>> subscriptions() {
        return Optional.ofNullable(this.subscriptions);
    }

    /**
     * If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
     * 
     */
    @Import(name="syncGuestOsNamespaces")
    private @Nullable Output<Boolean> syncGuestOsNamespaces;

    /**
     * @return If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> syncGuestOsNamespaces() {
        return Optional.ofNullable(this.syncGuestOsNamespaces);
    }

    /**
     * Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     * 
     */
    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    /**
     * @return Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     * 
     */
    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    private IntegrationState() {}

    private IntegrationState(IntegrationState $) {
        this.additionalServices = $.additionalServices;
        this.appId = $.appId;
        this.customNamespacesPerServices = $.customNamespacesPerServices;
        this.enabled = $.enabled;
        this.environment = $.environment;
        this.importAzureMonitor = $.importAzureMonitor;
        this.name = $.name;
        this.namedToken = $.namedToken;
        this.pollRate = $.pollRate;
        this.resourceFilterRules = $.resourceFilterRules;
        this.secretKey = $.secretKey;
        this.services = $.services;
        this.subscriptions = $.subscriptions;
        this.syncGuestOsNamespaces = $.syncGuestOsNamespaces;
        this.tenantId = $.tenantId;
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
         * @param additionalServices Additional Azure resource types that you want to sync with Observability Cloud.
         * 
         * @return builder
         * 
         */
        public Builder additionalServices(@Nullable Output<List<String>> additionalServices) {
            $.additionalServices = additionalServices;
            return this;
        }

        /**
         * @param additionalServices Additional Azure resource types that you want to sync with Observability Cloud.
         * 
         * @return builder
         * 
         */
        public Builder additionalServices(List<String> additionalServices) {
            return additionalServices(Output.of(additionalServices));
        }

        /**
         * @param additionalServices Additional Azure resource types that you want to sync with Observability Cloud.
         * 
         * @return builder
         * 
         */
        public Builder additionalServices(String... additionalServices) {
            return additionalServices(List.of(additionalServices));
        }

        /**
         * @param appId Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
         * 
         * @return builder
         * 
         */
        public Builder appId(@Nullable Output<String> appId) {
            $.appId = appId;
            return this;
        }

        /**
         * @param appId Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
         * 
         * @return builder
         * 
         */
        public Builder appId(String appId) {
            return appId(Output.of(appId));
        }

        /**
         * @param customNamespacesPerServices Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `sync_guest_os_namespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
         * 
         * @return builder
         * 
         */
        public Builder customNamespacesPerServices(@Nullable Output<List<IntegrationCustomNamespacesPerServiceArgs>> customNamespacesPerServices) {
            $.customNamespacesPerServices = customNamespacesPerServices;
            return this;
        }

        /**
         * @param customNamespacesPerServices Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `sync_guest_os_namespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
         * 
         * @return builder
         * 
         */
        public Builder customNamespacesPerServices(List<IntegrationCustomNamespacesPerServiceArgs> customNamespacesPerServices) {
            return customNamespacesPerServices(Output.of(customNamespacesPerServices));
        }

        /**
         * @param customNamespacesPerServices Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `sync_guest_os_namespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
         * 
         * @return builder
         * 
         */
        public Builder customNamespacesPerServices(IntegrationCustomNamespacesPerServiceArgs... customNamespacesPerServices) {
            return customNamespacesPerServices(List.of(customNamespacesPerServices));
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
         * @param environment What type of Azure integration this is. The allowed values are `\&#34;azure_us_government\&#34;` and `\&#34;azure\&#34;`. Defaults to `\&#34;azure\&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder environment(@Nullable Output<String> environment) {
            $.environment = environment;
            return this;
        }

        /**
         * @param environment What type of Azure integration this is. The allowed values are `\&#34;azure_us_government\&#34;` and `\&#34;azure\&#34;`. Defaults to `\&#34;azure\&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder environment(String environment) {
            return environment(Output.of(environment));
        }

        /**
         * @param importAzureMonitor If enabled, SignalFx will sync also Azure Monitor data. If disabled, SignalFx will import only metadata. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder importAzureMonitor(@Nullable Output<Boolean> importAzureMonitor) {
            $.importAzureMonitor = importAzureMonitor;
            return this;
        }

        /**
         * @param importAzureMonitor If enabled, SignalFx will sync also Azure Monitor data. If disabled, SignalFx will import only metadata. Defaults to true.
         * 
         * @return builder
         * 
         */
        public Builder importAzureMonitor(Boolean importAzureMonitor) {
            return importAzureMonitor(Output.of(importAzureMonitor));
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
         * @param pollRate Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
         * 
         * @return builder
         * 
         */
        public Builder pollRate(@Nullable Output<Integer> pollRate) {
            $.pollRate = pollRate;
            return this;
        }

        /**
         * @param pollRate Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
         * 
         * @return builder
         * 
         */
        public Builder pollRate(Integer pollRate) {
            return pollRate(Output.of(pollRate));
        }

        /**
         * @param resourceFilterRules List of rules for filtering Azure resources by their tags.
         * 
         * @return builder
         * 
         */
        public Builder resourceFilterRules(@Nullable Output<List<IntegrationResourceFilterRuleArgs>> resourceFilterRules) {
            $.resourceFilterRules = resourceFilterRules;
            return this;
        }

        /**
         * @param resourceFilterRules List of rules for filtering Azure resources by their tags.
         * 
         * @return builder
         * 
         */
        public Builder resourceFilterRules(List<IntegrationResourceFilterRuleArgs> resourceFilterRules) {
            return resourceFilterRules(Output.of(resourceFilterRules));
        }

        /**
         * @param resourceFilterRules List of rules for filtering Azure resources by their tags.
         * 
         * @return builder
         * 
         */
        public Builder resourceFilterRules(IntegrationResourceFilterRuleArgs... resourceFilterRules) {
            return resourceFilterRules(List.of(resourceFilterRules));
        }

        /**
         * @param secretKey Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
         * 
         * @return builder
         * 
         */
        public Builder secretKey(@Nullable Output<String> secretKey) {
            $.secretKey = secretKey;
            return this;
        }

        /**
         * @param secretKey Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
         * 
         * @return builder
         * 
         */
        public Builder secretKey(String secretKey) {
            return secretKey(Output.of(secretKey));
        }

        /**
         * @param services List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
         * 
         * @return builder
         * 
         */
        public Builder services(@Nullable Output<List<String>> services) {
            $.services = services;
            return this;
        }

        /**
         * @param services List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
         * 
         * @return builder
         * 
         */
        public Builder services(List<String> services) {
            return services(Output.of(services));
        }

        /**
         * @param services List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
         * 
         * @return builder
         * 
         */
        public Builder services(String... services) {
            return services(List.of(services));
        }

        /**
         * @param subscriptions List of Azure subscriptions that SignalFx should monitor.
         * 
         * @return builder
         * 
         */
        public Builder subscriptions(@Nullable Output<List<String>> subscriptions) {
            $.subscriptions = subscriptions;
            return this;
        }

        /**
         * @param subscriptions List of Azure subscriptions that SignalFx should monitor.
         * 
         * @return builder
         * 
         */
        public Builder subscriptions(List<String> subscriptions) {
            return subscriptions(Output.of(subscriptions));
        }

        /**
         * @param subscriptions List of Azure subscriptions that SignalFx should monitor.
         * 
         * @return builder
         * 
         */
        public Builder subscriptions(String... subscriptions) {
            return subscriptions(List.of(subscriptions));
        }

        /**
         * @param syncGuestOsNamespaces If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder syncGuestOsNamespaces(@Nullable Output<Boolean> syncGuestOsNamespaces) {
            $.syncGuestOsNamespaces = syncGuestOsNamespaces;
            return this;
        }

        /**
         * @param syncGuestOsNamespaces If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder syncGuestOsNamespaces(Boolean syncGuestOsNamespaces) {
            return syncGuestOsNamespaces(Output.of(syncGuestOsNamespaces));
        }

        /**
         * @param tenantId Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        /**
         * @param tenantId Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public IntegrationState build() {
            return $;
        }
    }

}
