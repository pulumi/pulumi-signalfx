// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Gcp
{
    /// <summary>
    /// SignalFx GCP Integration
    /// 
    /// &gt; **NOTE** When managing integrations, use a session token of an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using SignalFx = Pulumi.SignalFx;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var gcpMyteam = new SignalFx.Gcp.Integration("gcpMyteam", new()
    ///     {
    ///         CustomMetricTypeDomains = new[]
    ///         {
    ///             "istio.io",
    ///         },
    ///         Enabled = true,
    ///         ImportGcpMetrics = true,
    ///         PollRate = 300,
    ///         ProjectServiceKeys = new[]
    ///         {
    ///             new SignalFx.Gcp.Inputs.IntegrationProjectServiceKeyArgs
    ///             {
    ///                 ProjectId = "gcp_project_id_1",
    ///                 ProjectKey = File.ReadAllText("/path/to/gcp_credentials_1.json"),
    ///             },
    ///             new SignalFx.Gcp.Inputs.IntegrationProjectServiceKeyArgs
    ///             {
    ///                 ProjectId = "gcp_project_id_2",
    ///                 ProjectKey = File.ReadAllText("/path/to/gcp_credentials_2.json"),
    ///             },
    ///         },
    ///         Services = new[]
    ///         {
    ///             "compute",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:gcp/integration:Integration")]
    public partial class Integration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of additional GCP service domain names that Splunk Observability Cloud will monitor. See [Custom Metric Type Domains documentation](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/#Custom-metric-type-domains)
        /// </summary>
        [Output("customMetricTypeDomains")]
        public Output<ImmutableArray<string>> CustomMetricTypeDomains { get; private set; } = null!;

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// If enabled, Splunk Observability Cloud will sync also Google Cloud Monitoring data. If disabled, Splunk Observability Cloud will import only metadata. Defaults to true.
        /// </summary>
        [Output("importGcpMetrics")]
        public Output<bool?> ImportGcpMetrics { get; private set; } = null!;

        /// <summary>
        /// [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
        /// </summary>
        [Output("includeLists")]
        public Output<ImmutableArray<string>> IncludeLists { get; private set; } = null!;

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
        /// GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
        /// </summary>
        [Output("pollRate")]
        public Output<int?> PollRate { get; private set; } = null!;

        /// <summary>
        /// GCP projects to add.
        /// </summary>
        [Output("projectServiceKeys")]
        public Output<ImmutableArray<Outputs.IntegrationProjectServiceKey>> ProjectServiceKeys { get; private set; } = null!;

        /// <summary>
        /// GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See [Google Cloud Platform services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#google-cloud-platform-services) for a list of valid values.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<string>> Services { get; private set; } = null!;

        /// <summary>
        /// When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
        /// </summary>
        [Output("useMetricSourceProjectForQuota")]
        public Output<bool?> UseMetricSourceProjectForQuota { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("signalfx:gcp/integration:Integration", name, args ?? new IntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Integration(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:gcp/integration:Integration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "projectServiceKeys",
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
        [Input("customMetricTypeDomains")]
        private InputList<string>? _customMetricTypeDomains;

        /// <summary>
        /// List of additional GCP service domain names that Splunk Observability Cloud will monitor. See [Custom Metric Type Domains documentation](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/#Custom-metric-type-domains)
        /// </summary>
        public InputList<string> CustomMetricTypeDomains
        {
            get => _customMetricTypeDomains ?? (_customMetricTypeDomains = new InputList<string>());
            set => _customMetricTypeDomains = value;
        }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// If enabled, Splunk Observability Cloud will sync also Google Cloud Monitoring data. If disabled, Splunk Observability Cloud will import only metadata. Defaults to true.
        /// </summary>
        [Input("importGcpMetrics")]
        public Input<bool>? ImportGcpMetrics { get; set; }

        [Input("includeLists")]
        private InputList<string>? _includeLists;

        /// <summary>
        /// [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
        /// </summary>
        public InputList<string> IncludeLists
        {
            get => _includeLists ?? (_includeLists = new InputList<string>());
            set => _includeLists = value;
        }

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
        /// GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
        /// </summary>
        [Input("pollRate")]
        public Input<int>? PollRate { get; set; }

        [Input("projectServiceKeys")]
        private InputList<Inputs.IntegrationProjectServiceKeyArgs>? _projectServiceKeys;

        /// <summary>
        /// GCP projects to add.
        /// </summary>
        public InputList<Inputs.IntegrationProjectServiceKeyArgs> ProjectServiceKeys
        {
            get => _projectServiceKeys ?? (_projectServiceKeys = new InputList<Inputs.IntegrationProjectServiceKeyArgs>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableArray.Create<Inputs.IntegrationProjectServiceKeyArgs>());
                _projectServiceKeys = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See [Google Cloud Platform services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#google-cloud-platform-services) for a list of valid values.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        /// <summary>
        /// When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
        /// </summary>
        [Input("useMetricSourceProjectForQuota")]
        public Input<bool>? UseMetricSourceProjectForQuota { get; set; }

        public IntegrationArgs()
        {
        }
        public static new IntegrationArgs Empty => new IntegrationArgs();
    }

    public sealed class IntegrationState : global::Pulumi.ResourceArgs
    {
        [Input("customMetricTypeDomains")]
        private InputList<string>? _customMetricTypeDomains;

        /// <summary>
        /// List of additional GCP service domain names that Splunk Observability Cloud will monitor. See [Custom Metric Type Domains documentation](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/#Custom-metric-type-domains)
        /// </summary>
        public InputList<string> CustomMetricTypeDomains
        {
            get => _customMetricTypeDomains ?? (_customMetricTypeDomains = new InputList<string>());
            set => _customMetricTypeDomains = value;
        }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// If enabled, Splunk Observability Cloud will sync also Google Cloud Monitoring data. If disabled, Splunk Observability Cloud will import only metadata. Defaults to true.
        /// </summary>
        [Input("importGcpMetrics")]
        public Input<bool>? ImportGcpMetrics { get; set; }

        [Input("includeLists")]
        private InputList<string>? _includeLists;

        /// <summary>
        /// [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
        /// </summary>
        public InputList<string> IncludeLists
        {
            get => _includeLists ?? (_includeLists = new InputList<string>());
            set => _includeLists = value;
        }

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
        /// GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
        /// </summary>
        [Input("pollRate")]
        public Input<int>? PollRate { get; set; }

        [Input("projectServiceKeys")]
        private InputList<Inputs.IntegrationProjectServiceKeyGetArgs>? _projectServiceKeys;

        /// <summary>
        /// GCP projects to add.
        /// </summary>
        public InputList<Inputs.IntegrationProjectServiceKeyGetArgs> ProjectServiceKeys
        {
            get => _projectServiceKeys ?? (_projectServiceKeys = new InputList<Inputs.IntegrationProjectServiceKeyGetArgs>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableArray.Create<Inputs.IntegrationProjectServiceKeyGetArgs>());
                _projectServiceKeys = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See [Google Cloud Platform services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#google-cloud-platform-services) for a list of valid values.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        /// <summary>
        /// When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
        /// </summary>
        [Input("useMetricSourceProjectForQuota")]
        public Input<bool>? UseMetricSourceProjectForQuota { get; set; }

        public IntegrationState()
        {
        }
        public static new IntegrationState Empty => new IntegrationState();
    }
}
