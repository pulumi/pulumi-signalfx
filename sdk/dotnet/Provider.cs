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
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// API URL for your Splunk Observability Cloud org, may include a realm
        /// </summary>
        [Output("apiUrl")]
        public Output<string?> ApiUrl { get; private set; } = null!;

        /// <summary>
        /// Splunk Observability Cloud auth token
        /// </summary>
        [Output("authToken")]
        public Output<string?> AuthToken { get; private set; } = null!;

        /// <summary>
        /// Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
        /// </summary>
        [Output("customAppUrl")]
        public Output<string?> CustomAppUrl { get; private set; } = null!;

        /// <summary>
        /// Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
        /// Password
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// Required if the user is configured to be part of multiple organizations
        /// </summary>
        [Output("organizationId")]
        public Output<string?> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
        /// Password
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;


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
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API URL for your Splunk Observability Cloud org, may include a realm
        /// </summary>
        [Input("apiUrl")]
        public Input<string>? ApiUrl { get; set; }

        /// <summary>
        /// Splunk Observability Cloud auth token
        /// </summary>
        [Input("authToken")]
        public Input<string>? AuthToken { get; set; }

        /// <summary>
        /// Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
        /// </summary>
        [Input("customAppUrl")]
        public Input<string>? CustomAppUrl { get; set; }

        /// <summary>
        /// Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
        /// Password
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Required if the user is configured to be part of multiple organizations
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
        /// Password
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Max retries for a single HTTP call. Defaults to 4
        /// </summary>
        [Input("retryMaxAttempts", json: true)]
        public Input<int>? RetryMaxAttempts { get; set; }

        /// <summary>
        /// Maximum retry wait for a single HTTP call in seconds. Defaults to 30
        /// </summary>
        [Input("retryWaitMaxSeconds", json: true)]
        public Input<int>? RetryWaitMaxSeconds { get; set; }

        /// <summary>
        /// Minimum retry wait for a single HTTP call in seconds. Defaults to 1
        /// </summary>
        [Input("retryWaitMinSeconds", json: true)]
        public Input<int>? RetryWaitMinSeconds { get; set; }

        /// <summary>
        /// Timeout duration for a single HTTP call in seconds. Defaults to 120
        /// </summary>
        [Input("timeoutSeconds", json: true)]
        public Input<int>? TimeoutSeconds { get; set; }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
