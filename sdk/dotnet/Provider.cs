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
    /// The provider type for the signalfx package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [SignalFxResourceType("pulumi:providers:signalfx")]
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// API URL for your SignalFx org, may include a realm
        /// </summary>
        [Output("apiUrl")]
        public Output<string?> ApiUrl { get; private set; } = null!;

        /// <summary>
        /// SignalFx auth token
        /// </summary>
        [Output("authToken")]
        public Output<string?> AuthToken { get; private set; } = null!;

        /// <summary>
        /// Application URL for your SignalFx org, often customized for organizations using SSO
        /// </summary>
        [Output("customAppUrl")]
        public Output<string?> CustomAppUrl { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("signalfx", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API URL for your SignalFx org, may include a realm
        /// </summary>
        [Input("apiUrl")]
        public Input<string>? ApiUrl { get; set; }

        /// <summary>
        /// SignalFx auth token
        /// </summary>
        [Input("authToken")]
        public Input<string>? AuthToken { get; set; }

        /// <summary>
        /// Application URL for your SignalFx org, often customized for organizations using SSO
        /// </summary>
        [Input("customAppUrl")]
        public Input<string>? CustomAppUrl { get; set; }

        /// <summary>
        /// Timeout duration for a single HTTP call in seconds. Defaults to 120
        /// </summary>
        [Input("timeoutSeconds", json: true)]
        public Input<int>? TimeoutSeconds { get; set; }

        public ProviderArgs()
        {
        }
    }
}
