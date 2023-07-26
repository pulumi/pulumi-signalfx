// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.azure;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.azure.IntegrationArgs;
import com.pulumi.signalfx.azure.inputs.IntegrationState;
import com.pulumi.signalfx.azure.outputs.IntegrationCustomNamespacesPerService;
import com.pulumi.signalfx.azure.outputs.IntegrationResourceFilterRule;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * SignalFx Azure integrations. For help with this integration see [Monitoring Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure).
 * 
 * &gt; **NOTE** When managing integrations, use a session token of an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you&#39;ll receive a 4xx error.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.azure.Integration;
 * import com.pulumi.signalfx.azure.IntegrationArgs;
 * import com.pulumi.signalfx.azure.inputs.IntegrationCustomNamespacesPerServiceArgs;
 * import com.pulumi.signalfx.azure.inputs.IntegrationResourceFilterRuleArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var azureMyteam = new Integration(&#34;azureMyteam&#34;, IntegrationArgs.builder()        
 *             .additionalServices(            
 *                 &#34;some/service&#34;,
 *                 &#34;another/service&#34;)
 *             .appId(&#34;YYY&#34;)
 *             .customNamespacesPerServices(IntegrationCustomNamespacesPerServiceArgs.builder()
 *                 .namespaces(                
 *                     &#34;monitoringAgent&#34;,
 *                     &#34;customNamespace&#34;)
 *                 .service(&#34;Microsoft.Compute/virtualMachines&#34;)
 *                 .build())
 *             .enabled(true)
 *             .environment(&#34;azure&#34;)
 *             .pollRate(300)
 *             .resourceFilterRules(            
 *                 IntegrationResourceFilterRuleArgs.builder()
 *                     .filterSource(&#34;filter(&#39;azure_tag_service&#39;, &#39;payment&#39;) and (filter(&#39;azure_tag_env&#39;, &#39;prod-us&#39;) or filter(&#39;azure_tag_env&#39;, &#39;prod-eu&#39;))&#34;)
 *                     .build(),
 *                 IntegrationResourceFilterRuleArgs.builder()
 *                     .filterSource(&#34;filter(&#39;azure_tag_service&#39;, &#39;notification&#39;) and (filter(&#39;azure_tag_env&#39;, &#39;prod-us&#39;) or filter(&#39;azure_tag_env&#39;, &#39;prod-eu&#39;))&#34;)
 *                     .build())
 *             .secretKey(&#34;XXX&#34;)
 *             .services(&#34;microsoft.sql/servers/elasticpools&#34;)
 *             .subscriptions(&#34;sub-guid-here&#34;)
 *             .tenantId(&#34;ZZZ&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="signalfx:azure/integration:Integration")
public class Integration extends com.pulumi.resources.CustomResource {
    /**
     * Additional Azure resource types that you want to sync with Observability Cloud.
     * 
     */
    @Export(name="additionalServices", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> additionalServices;

    /**
     * @return Additional Azure resource types that you want to sync with Observability Cloud.
     * 
     */
    public Output<Optional<List<String>>> additionalServices() {
        return Codegen.optional(this.additionalServices);
    }
    /**
     * Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
     * 
     */
    @Export(name="appId", type=String.class, parameters={})
    private Output<String> appId;

    /**
     * @return Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
     * 
     */
    public Output<String> appId() {
        return this.appId;
    }
    /**
     * Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `sync_guest_os_namespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
     * 
     */
    @Export(name="customNamespacesPerServices", type=List.class, parameters={IntegrationCustomNamespacesPerService.class})
    private Output</* @Nullable */ List<IntegrationCustomNamespacesPerService>> customNamespacesPerServices;

    /**
     * @return Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `sync_guest_os_namespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
     * 
     */
    public Output<Optional<List<IntegrationCustomNamespacesPerService>>> customNamespacesPerServices() {
        return Codegen.optional(this.customNamespacesPerServices);
    }
    /**
     * Whether the integration is enabled.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output<Boolean> enabled;

    /**
     * @return Whether the integration is enabled.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * What type of Azure integration this is. The allowed values are `\&#34;azure_us_government\&#34;` and `\&#34;azure\&#34;`. Defaults to `\&#34;azure\&#34;`.
     * 
     */
    @Export(name="environment", type=String.class, parameters={})
    private Output</* @Nullable */ String> environment;

    /**
     * @return What type of Azure integration this is. The allowed values are `\&#34;azure_us_government\&#34;` and `\&#34;azure\&#34;`. Defaults to `\&#34;azure\&#34;`.
     * 
     */
    public Output<Optional<String>> environment() {
        return Codegen.optional(this.environment);
    }
    /**
     * If enabled, SignalFx will sync also Azure Monitor data. If disabled, SignalFx will import only metadata. Defaults to true.
     * 
     */
    @Export(name="importAzureMonitor", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> importAzureMonitor;

    /**
     * @return If enabled, SignalFx will sync also Azure Monitor data. If disabled, SignalFx will import only metadata. Defaults to true.
     * 
     */
    public Output<Optional<Boolean>> importAzureMonitor() {
        return Codegen.optional(this.importAzureMonitor);
    }
    /**
     * Name of the integration.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the integration.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     * 
     */
    @Export(name="namedToken", type=String.class, parameters={})
    private Output</* @Nullable */ String> namedToken;

    /**
     * @return Name of the org token to be used for data ingestion. If not specified then default access token is used.
     * 
     */
    public Output<Optional<String>> namedToken() {
        return Codegen.optional(this.namedToken);
    }
    /**
     * Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     * 
     */
    @Export(name="pollRate", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> pollRate;

    /**
     * @return Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     * 
     */
    public Output<Optional<Integer>> pollRate() {
        return Codegen.optional(this.pollRate);
    }
    /**
     * List of rules for filtering Azure resources by their tags.
     * 
     */
    @Export(name="resourceFilterRules", type=List.class, parameters={IntegrationResourceFilterRule.class})
    private Output</* @Nullable */ List<IntegrationResourceFilterRule>> resourceFilterRules;

    /**
     * @return List of rules for filtering Azure resources by their tags.
     * 
     */
    public Output<Optional<List<IntegrationResourceFilterRule>>> resourceFilterRules() {
        return Codegen.optional(this.resourceFilterRules);
    }
    /**
     * Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     * 
     */
    @Export(name="secretKey", type=String.class, parameters={})
    private Output<String> secretKey;

    /**
     * @return Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     * 
     */
    public Output<String> secretKey() {
        return this.secretKey;
    }
    /**
     * List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
     * 
     */
    @Export(name="services", type=List.class, parameters={String.class})
    private Output<List<String>> services;

    /**
     * @return List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
     * 
     */
    public Output<List<String>> services() {
        return this.services;
    }
    /**
     * List of Azure subscriptions that SignalFx should monitor.
     * 
     */
    @Export(name="subscriptions", type=List.class, parameters={String.class})
    private Output<List<String>> subscriptions;

    /**
     * @return List of Azure subscriptions that SignalFx should monitor.
     * 
     */
    public Output<List<String>> subscriptions() {
        return this.subscriptions;
    }
    /**
     * If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
     * 
     */
    @Export(name="syncGuestOsNamespaces", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> syncGuestOsNamespaces;

    /**
     * @return If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> syncGuestOsNamespaces() {
        return Codegen.optional(this.syncGuestOsNamespaces);
    }
    /**
     * Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     * 
     */
    @Export(name="tenantId", type=String.class, parameters={})
    private Output<String> tenantId;

    /**
     * @return Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     * 
     */
    public Output<String> tenantId() {
        return this.tenantId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Integration(String name) {
        this(name, IntegrationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Integration(String name, IntegrationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Integration(String name, IntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:azure/integration:Integration", name, args == null ? IntegrationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Integration(String name, Output<String> id, @Nullable IntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:azure/integration:Integration", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "appId",
                "environment",
                "secretKey"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Integration get(String name, Output<String> id, @Nullable IntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Integration(name, id, state, options);
    }
}
