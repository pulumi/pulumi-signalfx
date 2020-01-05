// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Gcp
{
    /// <summary>
    /// SignalFx GCP Integration
    /// 
    /// **Note:** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/gcp_integration.html.markdown.
    /// </summary>
    public partial class Integration : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// GCP integration poll rate in milliseconds. Can be set to either 60000 or 300000 (1 minute or 5 minutes).
        /// </summary>
        [Output("pollRate")]
        public Output<int?> PollRate { get; private set; } = null!;

        /// <summary>
        /// GCP projects to add.
        /// </summary>
        [Output("projectServiceKeys")]
        public Output<ImmutableArray<Outputs.IntegrationProjectServiceKeys>> ProjectServiceKeys { get; private set; } = null!;

        /// <summary>
        /// GCP service metrics to import. Can be an empty list, or not included, to import 'All services'.
        /// </summary>
        [Output("services")]
        public Output<ImmutableArray<string>> Services { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("signalfx:gcp/integration:Integration", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class IntegrationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// GCP integration poll rate in milliseconds. Can be set to either 60000 or 300000 (1 minute or 5 minutes).
        /// </summary>
        [Input("pollRate")]
        public Input<int>? PollRate { get; set; }

        [Input("projectServiceKeys")]
        private InputList<Inputs.IntegrationProjectServiceKeysArgs>? _projectServiceKeys;

        /// <summary>
        /// GCP projects to add.
        /// </summary>
        public InputList<Inputs.IntegrationProjectServiceKeysArgs> ProjectServiceKeys
        {
            get => _projectServiceKeys ?? (_projectServiceKeys = new InputList<Inputs.IntegrationProjectServiceKeysArgs>());
            set => _projectServiceKeys = value;
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// GCP service metrics to import. Can be an empty list, or not included, to import 'All services'.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        public IntegrationArgs()
        {
        }
    }

    public sealed class IntegrationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// GCP integration poll rate in milliseconds. Can be set to either 60000 or 300000 (1 minute or 5 minutes).
        /// </summary>
        [Input("pollRate")]
        public Input<int>? PollRate { get; set; }

        [Input("projectServiceKeys")]
        private InputList<Inputs.IntegrationProjectServiceKeysGetArgs>? _projectServiceKeys;

        /// <summary>
        /// GCP projects to add.
        /// </summary>
        public InputList<Inputs.IntegrationProjectServiceKeysGetArgs> ProjectServiceKeys
        {
            get => _projectServiceKeys ?? (_projectServiceKeys = new InputList<Inputs.IntegrationProjectServiceKeysGetArgs>());
            set => _projectServiceKeys = value;
        }

        [Input("services")]
        private InputList<string>? _services;

        /// <summary>
        /// GCP service metrics to import. Can be an empty list, or not included, to import 'All services'.
        /// </summary>
        public InputList<string> Services
        {
            get => _services ?? (_services = new InputList<string>());
            set => _services = value;
        }

        public IntegrationState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class IntegrationProjectServiceKeysArgs : Pulumi.ResourceArgs
    {
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("projectKey", required: true)]
        public Input<string> ProjectKey { get; set; } = null!;

        public IntegrationProjectServiceKeysArgs()
        {
        }
    }

    public sealed class IntegrationProjectServiceKeysGetArgs : Pulumi.ResourceArgs
    {
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("projectKey", required: true)]
        public Input<string> ProjectKey { get; set; } = null!;

        public IntegrationProjectServiceKeysGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class IntegrationProjectServiceKeys
    {
        public readonly string ProjectId;
        public readonly string ProjectKey;

        [OutputConstructor]
        private IntegrationProjectServiceKeys(
            string projectId,
            string projectKey)
        {
            ProjectId = projectId;
            ProjectKey = projectKey;
        }
    }
    }
}
