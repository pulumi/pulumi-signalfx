// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Splunk Observability Cloud Azure integrations. For help with this integration see [Monitoring Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html).
 *
 * > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
 *
 * ## Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 *
 * const azureMyteam = new signalfx.azure.Integration("azure_myteam", {
 *     name: "Azure Foo",
 *     enabled: true,
 *     environment: "azure",
 *     pollRate: 300,
 *     secretKey: "XXX",
 *     appId: "YYY",
 *     tenantId: "ZZZ",
 *     services: ["microsoft.sql/servers/elasticpools"],
 *     subscriptions: ["sub-guid-here"],
 *     additionalServices: [
 *         "some/service",
 *         "another/service",
 *     ],
 *     customNamespacesPerServices: [{
 *         service: "Microsoft.Compute/virtualMachines",
 *         namespaces: [
 *             "monitoringAgent",
 *             "customNamespace",
 *         ],
 *     }],
 *     resourceFilterRules: [
 *         {
 *             filterSource: "filter('azure_tag_service', 'payment') and (filter('azure_tag_env', 'prod-us') or filter('azure_tag_env', 'prod-eu'))",
 *         },
 *         {
 *             filterSource: "filter('azure_tag_service', 'notification') and (filter('azure_tag_env', 'prod-us') or filter('azure_tag_env', 'prod-eu'))",
 *         },
 *     ],
 * });
 * ```
 */
export class Integration extends pulumi.CustomResource {
    /**
     * Get an existing Integration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationState, opts?: pulumi.CustomResourceOptions): Integration {
        return new Integration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:azure/integration:Integration';

    /**
     * Returns true if the given object is an instance of Integration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Integration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Integration.__pulumiType;
    }

    /**
     * Additional Azure resource types that you want to sync with Observability Cloud.
     */
    public readonly additionalServices!: pulumi.Output<string[] | undefined>;
    /**
     * Azure application ID for the Splunk Observability Cloud app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `syncGuestOsNamespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
     */
    public readonly customNamespacesPerServices!: pulumi.Output<outputs.azure.IntegrationCustomNamespacesPerService[] | undefined>;
    /**
     * Whether the integration is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
     */
    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * If enabled, Splunk Observability Cloud will sync also Azure Monitor data. If disabled, Splunk Observability Cloud will import only metadata. Defaults to true.
     */
    public readonly importAzureMonitor!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the integration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     */
    public readonly namedToken!: pulumi.Output<string | undefined>;
    /**
     * Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     */
    public readonly pollRate!: pulumi.Output<number | undefined>;
    /**
     * List of rules for filtering Azure resources by their tags.
     */
    public readonly resourceFilterRules!: pulumi.Output<outputs.azure.IntegrationResourceFilterRule[] | undefined>;
    /**
     * Azure secret key that associates the Splunk Observability Cloud app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
     */
    public readonly secretKey!: pulumi.Output<string>;
    /**
     * List of Microsoft Azure service names for the Azure services you want Splunk Observability Cloud to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
     */
    public readonly services!: pulumi.Output<string[]>;
    /**
     * List of Azure subscriptions that Splunk Observability Cloud should monitor.
     */
    public readonly subscriptions!: pulumi.Output<string[]>;
    /**
     * If enabled, Splunk Observability Cloud will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
     */
    public readonly syncGuestOsNamespaces!: pulumi.Output<boolean | undefined>;
    /**
     * Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * If enabled, Splunk Observability Cloud will collect datapoints using Azure Metrics Batch API. Consider this option if you are synchronizing high loads of data and you want to avoid throttling issues. Contrary to the default Metrics List API, Metrics Batch API is paid. Refer to [Azure documentation](https://azure.microsoft.com/en-us/pricing/details/api-management/) for pricing info.
     */
    public readonly useBatchApi!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Integration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationArgs | IntegrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationState | undefined;
            resourceInputs["additionalServices"] = state ? state.additionalServices : undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["customNamespacesPerServices"] = state ? state.customNamespacesPerServices : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["importAzureMonitor"] = state ? state.importAzureMonitor : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namedToken"] = state ? state.namedToken : undefined;
            resourceInputs["pollRate"] = state ? state.pollRate : undefined;
            resourceInputs["resourceFilterRules"] = state ? state.resourceFilterRules : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["subscriptions"] = state ? state.subscriptions : undefined;
            resourceInputs["syncGuestOsNamespaces"] = state ? state.syncGuestOsNamespaces : undefined;
            resourceInputs["tenantId"] = state ? state.tenantId : undefined;
            resourceInputs["useBatchApi"] = state ? state.useBatchApi : undefined;
        } else {
            const args = argsOrState as IntegrationArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.secretKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretKey'");
            }
            if ((!args || args.services === undefined) && !opts.urn) {
                throw new Error("Missing required property 'services'");
            }
            if ((!args || args.subscriptions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscriptions'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["additionalServices"] = args ? args.additionalServices : undefined;
            resourceInputs["appId"] = args?.appId ? pulumi.secret(args.appId) : undefined;
            resourceInputs["customNamespacesPerServices"] = args ? args.customNamespacesPerServices : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["environment"] = args?.environment ? pulumi.secret(args.environment) : undefined;
            resourceInputs["importAzureMonitor"] = args ? args.importAzureMonitor : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namedToken"] = args ? args.namedToken : undefined;
            resourceInputs["pollRate"] = args ? args.pollRate : undefined;
            resourceInputs["resourceFilterRules"] = args ? args.resourceFilterRules : undefined;
            resourceInputs["secretKey"] = args?.secretKey ? pulumi.secret(args.secretKey) : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["subscriptions"] = args ? args.subscriptions : undefined;
            resourceInputs["syncGuestOsNamespaces"] = args ? args.syncGuestOsNamespaces : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["useBatchApi"] = args ? args.useBatchApi : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["appId", "environment", "secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Integration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Integration resources.
 */
export interface IntegrationState {
    /**
     * Additional Azure resource types that you want to sync with Observability Cloud.
     */
    additionalServices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Azure application ID for the Splunk Observability Cloud app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
     */
    appId?: pulumi.Input<string>;
    /**
     * Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `syncGuestOsNamespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
     */
    customNamespacesPerServices?: pulumi.Input<pulumi.Input<inputs.azure.IntegrationCustomNamespacesPerService>[]>;
    /**
     * Whether the integration is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
     */
    environment?: pulumi.Input<string>;
    /**
     * If enabled, Splunk Observability Cloud will sync also Azure Monitor data. If disabled, Splunk Observability Cloud will import only metadata. Defaults to true.
     */
    importAzureMonitor?: pulumi.Input<boolean>;
    /**
     * Name of the integration.
     */
    name?: pulumi.Input<string>;
    /**
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     */
    namedToken?: pulumi.Input<string>;
    /**
     * Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     */
    pollRate?: pulumi.Input<number>;
    /**
     * List of rules for filtering Azure resources by their tags.
     */
    resourceFilterRules?: pulumi.Input<pulumi.Input<inputs.azure.IntegrationResourceFilterRule>[]>;
    /**
     * Azure secret key that associates the Splunk Observability Cloud app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * List of Microsoft Azure service names for the Azure services you want Splunk Observability Cloud to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
     */
    services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Azure subscriptions that Splunk Observability Cloud should monitor.
     */
    subscriptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If enabled, Splunk Observability Cloud will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
     */
    syncGuestOsNamespaces?: pulumi.Input<boolean>;
    /**
     * Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * If enabled, Splunk Observability Cloud will collect datapoints using Azure Metrics Batch API. Consider this option if you are synchronizing high loads of data and you want to avoid throttling issues. Contrary to the default Metrics List API, Metrics Batch API is paid. Refer to [Azure documentation](https://azure.microsoft.com/en-us/pricing/details/api-management/) for pricing info.
     */
    useBatchApi?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Integration resource.
 */
export interface IntegrationArgs {
    /**
     * Additional Azure resource types that you want to sync with Observability Cloud.
     */
    additionalServices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Azure application ID for the Splunk Observability Cloud app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
     */
    appId: pulumi.Input<string>;
    /**
     * Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `syncGuestOsNamespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
     */
    customNamespacesPerServices?: pulumi.Input<pulumi.Input<inputs.azure.IntegrationCustomNamespacesPerService>[]>;
    /**
     * Whether the integration is enabled.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
     */
    environment?: pulumi.Input<string>;
    /**
     * If enabled, Splunk Observability Cloud will sync also Azure Monitor data. If disabled, Splunk Observability Cloud will import only metadata. Defaults to true.
     */
    importAzureMonitor?: pulumi.Input<boolean>;
    /**
     * Name of the integration.
     */
    name?: pulumi.Input<string>;
    /**
     * Name of the org token to be used for data ingestion. If not specified then default access token is used.
     */
    namedToken?: pulumi.Input<string>;
    /**
     * Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
     */
    pollRate?: pulumi.Input<number>;
    /**
     * List of rules for filtering Azure resources by their tags.
     */
    resourceFilterRules?: pulumi.Input<pulumi.Input<inputs.azure.IntegrationResourceFilterRule>[]>;
    /**
     * Azure secret key that associates the Splunk Observability Cloud app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
     */
    secretKey: pulumi.Input<string>;
    /**
     * List of Microsoft Azure service names for the Azure services you want Splunk Observability Cloud to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
     */
    services: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Azure subscriptions that Splunk Observability Cloud should monitor.
     */
    subscriptions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If enabled, Splunk Observability Cloud will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
     */
    syncGuestOsNamespaces?: pulumi.Input<boolean>;
    /**
     * Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
     */
    tenantId: pulumi.Input<string>;
    /**
     * If enabled, Splunk Observability Cloud will collect datapoints using Azure Metrics Batch API. Consider this option if you are synchronizing high loads of data and you want to avoid throttling issues. Contrary to the default Metrics List API, Metrics Batch API is paid. Refer to [Azure documentation](https://azure.microsoft.com/en-us/pricing/details/api-management/) for pricing info.
     */
    useBatchApi?: pulumi.Input<boolean>;
}
