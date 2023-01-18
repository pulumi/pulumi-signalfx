// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx
{
    /// <summary>
    /// Manage SignalFx org tokens.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using SignalFx = Pulumi.SignalFx;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myteamkey0 = new SignalFx.OrgToken("myteamkey0", new()
    ///     {
    ///         Description = "My team's rad key",
    ///         HostOrUsageLimits = new SignalFx.Inputs.OrgTokenHostOrUsageLimitsArgs
    ///         {
    ///             ContainerLimit = 200,
    ///             ContainerNotificationThreshold = 180,
    ///             CustomMetricsLimit = 1000,
    ///             CustomMetricsNotificationThreshold = 900,
    ///             HighResMetricsLimit = 1000,
    ///             HighResMetricsNotificationThreshold = 900,
    ///             HostLimit = 100,
    ///             HostNotificationThreshold = 90,
    ///         },
    ///         Notifications = new[]
    ///         {
    ///             "Email,foo-alerts@bar.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:index/orgToken:OrgToken")]
    public partial class OrgToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Authentication scope, ex: INGEST, API, RUM ... (Optional)
        /// </summary>
        [Output("authScopes")]
        public Output<ImmutableArray<string>> AuthScopes { get; private set; } = null!;

        /// <summary>
        /// Description of the token.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// Specify DPM-based limits for this token.
        /// </summary>
        [Output("dpmLimits")]
        public Output<Outputs.OrgTokenDpmLimits?> DpmLimits { get; private set; } = null!;

        /// <summary>
        /// Specify Usage-based limits for this token.
        /// </summary>
        [Output("hostOrUsageLimits")]
        public Output<Outputs.OrgTokenHostOrUsageLimits?> HostOrUsageLimits { get; private set; } = null!;

        /// <summary>
        /// Name of the token.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of strings specifying where notifications will be sent when an incident occurs. See
        /// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        /// </summary>
        [Output("notifications")]
        public Output<ImmutableArray<string>> Notifications { get; private set; } = null!;

        /// <summary>
        /// The secret token created by the API. You cannot set this value.
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;


        /// <summary>
        /// Create a OrgToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrgToken(string name, OrgTokenArgs? args = null, CustomResourceOptions? options = null)
            : base("signalfx:index/orgToken:OrgToken", name, args ?? new OrgTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrgToken(string name, Input<string> id, OrgTokenState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/orgToken:OrgToken", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "secret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrgToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrgToken Get(string name, Input<string> id, OrgTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new OrgToken(name, id, state, options);
        }
    }

    public sealed class OrgTokenArgs : global::Pulumi.ResourceArgs
    {
        [Input("authScopes")]
        private InputList<string>? _authScopes;

        /// <summary>
        /// Authentication scope, ex: INGEST, API, RUM ... (Optional)
        /// </summary>
        public InputList<string> AuthScopes
        {
            get => _authScopes ?? (_authScopes = new InputList<string>());
            set => _authScopes = value;
        }

        /// <summary>
        /// Description of the token.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Specify DPM-based limits for this token.
        /// </summary>
        [Input("dpmLimits")]
        public Input<Inputs.OrgTokenDpmLimitsArgs>? DpmLimits { get; set; }

        /// <summary>
        /// Specify Usage-based limits for this token.
        /// </summary>
        [Input("hostOrUsageLimits")]
        public Input<Inputs.OrgTokenHostOrUsageLimitsArgs>? HostOrUsageLimits { get; set; }

        /// <summary>
        /// Name of the token.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notifications")]
        private InputList<string>? _notifications;

        /// <summary>
        /// List of strings specifying where notifications will be sent when an incident occurs. See
        /// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        /// </summary>
        public InputList<string> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<string>());
            set => _notifications = value;
        }

        public OrgTokenArgs()
        {
        }
        public static new OrgTokenArgs Empty => new OrgTokenArgs();
    }

    public sealed class OrgTokenState : global::Pulumi.ResourceArgs
    {
        [Input("authScopes")]
        private InputList<string>? _authScopes;

        /// <summary>
        /// Authentication scope, ex: INGEST, API, RUM ... (Optional)
        /// </summary>
        public InputList<string> AuthScopes
        {
            get => _authScopes ?? (_authScopes = new InputList<string>());
            set => _authScopes = value;
        }

        /// <summary>
        /// Description of the token.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// Specify DPM-based limits for this token.
        /// </summary>
        [Input("dpmLimits")]
        public Input<Inputs.OrgTokenDpmLimitsGetArgs>? DpmLimits { get; set; }

        /// <summary>
        /// Specify Usage-based limits for this token.
        /// </summary>
        [Input("hostOrUsageLimits")]
        public Input<Inputs.OrgTokenHostOrUsageLimitsGetArgs>? HostOrUsageLimits { get; set; }

        /// <summary>
        /// Name of the token.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notifications")]
        private InputList<string>? _notifications;

        /// <summary>
        /// List of strings specifying where notifications will be sent when an incident occurs. See
        /// https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        /// </summary>
        public InputList<string> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<string>());
            set => _notifications = value;
        }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// The secret token created by the API. You cannot set this value.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public OrgTokenState()
        {
        }
        public static new OrgTokenState Empty => new OrgTokenState();
    }
}
