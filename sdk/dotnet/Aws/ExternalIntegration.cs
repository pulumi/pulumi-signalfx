// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Aws
{
    /// <summary>
    /// Splunk Observability AWS CloudWatch integrations using Role ARNs. For help with this integration see [Connect to AWS CloudWatch](https://docs.splunk.com/observability/en/gdi/get-data-in/connect/aws/aws-apiconfig.html).
    /// 
    /// &gt; **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
    /// 
    /// &gt; **WARNING** This resource implements part of a workflow. Use it with `signalfx.aws.Integration`. Check with Splunk support for your realm's AWS account id.
    /// 
    /// ## Example
    /// 
    /// ## Arguments
    /// 
    /// * `name` - (Required) The name of this integration
    /// 
    /// ## Attributes
    /// 
    /// In addition to all arguments above, the following attributes are exported:
    /// 
    /// * `id` - The ID of this integration, used with `signalfx.aws.Integration`
    /// * `external_id` - The external ID to use with your IAM role and with `signalfx.aws.Integration`.
    /// * `signalfx_aws_account` - The AWS Account ARN to use with your policies/roles, provided by Splunk Observability Cloud.
    /// </summary>
    [SignalFxResourceType("signalfx:aws/externalIntegration:ExternalIntegration")]
    public partial class ExternalIntegration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS external ID generated by Splunk Observability to use with an AWS integration.
        /// </summary>
        [Output("externalId")]
        public Output<string> ExternalId { get; private set; } = null!;

        /// <summary>
        /// Name of the integration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Splunk Observability AWS account ID to use with an AWS role.
        /// </summary>
        [Output("signalfxAwsAccount")]
        public Output<string> SignalfxAwsAccount { get; private set; } = null!;


        /// <summary>
        /// Create a ExternalIntegration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExternalIntegration(string name, ExternalIntegrationArgs? args = null, CustomResourceOptions? options = null)
            : base("signalfx:aws/externalIntegration:ExternalIntegration", name, args ?? new ExternalIntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExternalIntegration(string name, Input<string> id, ExternalIntegrationState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:aws/externalIntegration:ExternalIntegration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "externalId",
                    "signalfxAwsAccount",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExternalIntegration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExternalIntegration Get(string name, Input<string> id, ExternalIntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new ExternalIntegration(name, id, state, options);
        }
    }

    public sealed class ExternalIntegrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ExternalIntegrationArgs()
        {
        }
        public static new ExternalIntegrationArgs Empty => new ExternalIntegrationArgs();
    }

    public sealed class ExternalIntegrationState : global::Pulumi.ResourceArgs
    {
        [Input("externalId")]
        private Input<string>? _externalId;

        /// <summary>
        /// The AWS external ID generated by Splunk Observability to use with an AWS integration.
        /// </summary>
        public Input<string>? ExternalId
        {
            get => _externalId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _externalId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Name of the integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("signalfxAwsAccount")]
        private Input<string>? _signalfxAwsAccount;

        /// <summary>
        /// The Splunk Observability AWS account ID to use with an AWS role.
        /// </summary>
        public Input<string>? SignalfxAwsAccount
        {
            get => _signalfxAwsAccount;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _signalfxAwsAccount = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ExternalIntegrationState()
        {
        }
        public static new ExternalIntegrationState Empty => new ExternalIntegrationState();
    }
}
