// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Azure
{
    /// <summary>
    /// Splunk Observability Cloud Azure integrations. For help with this integration see [Monitoring Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html).
    /// 
    /// &gt; **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
    /// 
    /// ## Example
    /// </summary>
    [SignalFxResourceType("signalfx:azure/integration:Integration")]
    public partial class Integration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Additional Azure resource types that you want to sync with Observability Cloud.
        /// </summary>
        [Output("additionalServices")]
        public Output<ImmutableArray<string>> AdditionalServices { get; private set; } = null!;

        /// <summary>
        /// Azure application ID for the Splunk Observability Cloud app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `sync_guest_os_namespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
        /// </summary>
        [Output("customNamespacesPerServices")]
        public Output<ImmutableArray<Outputs.IntegrationCustomNamespacesPerService>> CustomNamespacesPerServices { get; private set; } = null!;

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
        /// </summary>
        [Output("environment")]
        public Output<string?> Environment { get; private set; } = null!;

        /// <summary>
        /// If enabled, Splunk Observability Cloud will sync also Azure Monitor data. If disabled, Splunk Observability Cloud will import only metadata. Defaults to true.
        /// </summary>
        [Output("importAzureMonitor")]
        public Output<bool?> ImportAzureMonitor { get; private set; } = null!;

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name of the org token to be used for data ingestion. If not specified then default access token is used.
        /// </summary>
        [Output("namedToken")]
        public Output<string?> NamedToken { get; private set; } = null!;

        /// <summary>
        /// Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
        /// </summary>
        [Output("pollRate")]
        public Output<int?> PollRate { get; private set; } = null!;

        /// <summary>
        /// List of rules for filtering Azure resources by their tags.
        /// </summary>
        [Output("resourceFilterRules")]
        public Output<ImmutableArray<Outputs.IntegrationResourceFilterRule>> ResourceFilterRules { get; private set; } = null!;

        /// <summary>
        /// Azure secret key that associates the Splunk Observability Cloud app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
        /// </summary>
        [Output("secretKey")]
        public Output<string> SecretKey { get; private set; } = null!;

        /// <summary>
        /// List of Microsoft Azure service names for the Azure services you want Splunk Observability Cloud to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<string>> Services { get; private set; } = null!;

        /// <summary>
        /// List of Azure subscriptions that Splunk Observability Cloud should monitor.
        /// </summary>
        [Output("subscriptions")]
        public Output<ImmutableArray<string>> Subscriptions { get; private set; } = null!;

        /// <summary>
        /// If enabled, Splunk Observability Cloud will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
        /// </summary>
        [Output("syncGuestOsNamespaces")]
        public Output<bool?> SyncGuestOsNamespaces { get; private set; } = null!;

        /// <summary>
        /// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("signalfx:azure/integration:Integration", name, args ?? new IntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Integration(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:azure/integration:Integration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "appId",
                    "environment",
                    "secretKey",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Integration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Integration Get(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new Integration(name, id, state, options);
        }
    }

    public sealed class IntegrationArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalServices")]
        private InputList<string>? _additionalServices;

        /// <summary>
        /// Additional Azure resource types that you want to sync with Observability Cloud.
        /// </summary>
        public InputList<string> AdditionalServices
        {
            get => _additionalServices ?? (_additionalServices = new InputList<string>());
            set => _additionalServices = value;
        }

        [Input("appId", required: true)]
        private Input<string>? _appId;

        /// <summary>
        /// Azure application ID for the Splunk Observability Cloud app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
        /// </summary>
        public Input<string>? AppId
        {
            get => _appId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _appId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("customNamespacesPerServices")]
        private InputList<Inputs.IntegrationCustomNamespacesPerServiceArgs>? _customNamespacesPerServices;

        /// <summary>
        /// Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `sync_guest_os_namespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
        /// </summary>
        public InputList<Inputs.IntegrationCustomNamespacesPerServiceArgs> CustomNamespacesPerServices
        {
            get => _customNamespacesPerServices ?? (_customNamespacesPerServices = new InputList<Inputs.IntegrationCustomNamespacesPerServiceArgs>());
            set => _customNamespacesPerServices = value;
        }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("environment")]
        private Input<string>? _environment;

        /// <summary>
        /// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
        /// </summary>
        public Input<string>? Environment
        {
            get => _environment;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _environment = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// If enabled, Splunk Observability Cloud will sync also Azure Monitor data. If disabled, Splunk Observability Cloud will import only metadata. Defaults to true.
        /// </summary>
        [Input("importAzureMonitor")]
        public Input<bool>? ImportAzureMonitor { get; set; }

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the org token to be used for data ingestion. If not specified then default access token is used.
        /// </summary>
        [Input("namedToken")]
        public Input<string>? NamedToken { get; set; }

        /// <summary>
        /// Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
        /// </summary>
        [Input("pollRate")]
        public Input<int>? PollRate { get; set; }

        [Input("resourceFilterRules")]
        private InputList<Inputs.IntegrationResourceFilterRuleArgs>? _resourceFilterRules;

        /// <summary>
        /// List of rules for filtering Azure resources by their tags.
        /// </summary>
        public InputList<Inputs.IntegrationResourceFilterRuleArgs> ResourceFilterRules
        {
            get => _resourceFilterRules ?? (_resourceFilterRules = new InputList<Inputs.IntegrationResourceFilterRuleArgs>());
            set => _resourceFilterRules = value;
        }

        [Input("secretKey", required: true)]
        private Input<string>? _secretKey;

        /// <summary>
        /// Azure secret key that associates the Splunk Observability Cloud app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("services", required: true)]
        private InputList<string>? _services;

        /// <summary>
        /// List of Microsoft Azure service names for the Azure services you want Splunk Observability Cloud to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        [Input("subscriptions", required: true)]
        private InputList<string>? _subscriptions;

        /// <summary>
        /// List of Azure subscriptions that Splunk Observability Cloud should monitor.
        /// </summary>
        public InputList<string> Subscriptions
        {
            get => _subscriptions ?? (_subscriptions = new InputList<string>());
            set => _subscriptions = value;
        }

        /// <summary>
        /// If enabled, Splunk Observability Cloud will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
        /// </summary>
        [Input("syncGuestOsNamespaces")]
        public Input<bool>? SyncGuestOsNamespaces { get; set; }

        /// <summary>
        /// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public IntegrationArgs()
        {
        }
        public static new IntegrationArgs Empty => new IntegrationArgs();
    }

    public sealed class IntegrationState : global::Pulumi.ResourceArgs
    {
        [Input("additionalServices")]
        private InputList<string>? _additionalServices;

        /// <summary>
        /// Additional Azure resource types that you want to sync with Observability Cloud.
        /// </summary>
        public InputList<string> AdditionalServices
        {
            get => _additionalServices ?? (_additionalServices = new InputList<string>());
            set => _additionalServices = value;
        }

        [Input("appId")]
        private Input<string>? _appId;

        /// <summary>
        /// Azure application ID for the Splunk Observability Cloud app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
        /// </summary>
        public Input<string>? AppId
        {
            get => _appId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _appId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("customNamespacesPerServices")]
        private InputList<Inputs.IntegrationCustomNamespacesPerServiceGetArgs>? _customNamespacesPerServices;

        /// <summary>
        /// Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `sync_guest_os_namespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
        /// </summary>
        public InputList<Inputs.IntegrationCustomNamespacesPerServiceGetArgs> CustomNamespacesPerServices
        {
            get => _customNamespacesPerServices ?? (_customNamespacesPerServices = new InputList<Inputs.IntegrationCustomNamespacesPerServiceGetArgs>());
            set => _customNamespacesPerServices = value;
        }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("environment")]
        private Input<string>? _environment;

        /// <summary>
        /// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
        /// </summary>
        public Input<string>? Environment
        {
            get => _environment;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _environment = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// If enabled, Splunk Observability Cloud will sync also Azure Monitor data. If disabled, Splunk Observability Cloud will import only metadata. Defaults to true.
        /// </summary>
        [Input("importAzureMonitor")]
        public Input<bool>? ImportAzureMonitor { get; set; }

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Name of the org token to be used for data ingestion. If not specified then default access token is used.
        /// </summary>
        [Input("namedToken")]
        public Input<string>? NamedToken { get; set; }

        /// <summary>
        /// Azure poll rate (in seconds). Value between `60` and `600`. Default: `300`.
        /// </summary>
        [Input("pollRate")]
        public Input<int>? PollRate { get; set; }

        [Input("resourceFilterRules")]
        private InputList<Inputs.IntegrationResourceFilterRuleGetArgs>? _resourceFilterRules;

        /// <summary>
        /// List of rules for filtering Azure resources by their tags.
        /// </summary>
        public InputList<Inputs.IntegrationResourceFilterRuleGetArgs> ResourceFilterRules
        {
            get => _resourceFilterRules ?? (_resourceFilterRules = new InputList<Inputs.IntegrationResourceFilterRuleGetArgs>());
            set => _resourceFilterRules = value;
        }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// Azure secret key that associates the Splunk Observability Cloud app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// List of Microsoft Azure service names for the Azure services you want Splunk Observability Cloud to monitor. Can be an empty list to import data for all supported services. See [Microsoft Azure services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#azure-integrations) for a list of valid values.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        [Input("subscriptions")]
        private InputList<string>? _subscriptions;

        /// <summary>
        /// List of Azure subscriptions that Splunk Observability Cloud should monitor.
        /// </summary>
        public InputList<string> Subscriptions
        {
            get => _subscriptions ?? (_subscriptions = new InputList<string>());
            set => _subscriptions = value;
        }

        /// <summary>
        /// If enabled, Splunk Observability Cloud will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
        /// </summary>
        [Input("syncGuestOsNamespaces")]
        public Input<bool>? SyncGuestOsNamespaces { get; set; }

        /// <summary>
        /// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/azure/azure.html) in the product documentation.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public IntegrationState()
        {
        }
        public static new IntegrationState Empty => new IntegrationState();
    }
}
