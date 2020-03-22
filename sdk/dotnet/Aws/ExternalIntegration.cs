// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Aws
{
    /// <summary>
    /// SignalFx AWS CloudWatch integrations using Role ARNs. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
    /// 
    /// &gt; **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.
    /// 
    /// &gt; **WARNING** This resource implements a part of a workflow. You must use it with `signalfx.aws.Integration`. Check with SignalFx support for your realm's AWS account id.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/aws_external_integration.html.markdown.
    /// </summary>
    public partial class ExternalIntegration : Pulumi.CustomResource
    {
        /// <summary>
        /// The external ID to use with your IAM role and with `signalfx.aws.Integration`.
        /// </summary>
        [Output("externalId")]
        public Output<string> ExternalId { get; private set; } = null!;

        /// <summary>
        /// The name of this integration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
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
            : base("signalfx:aws/externalIntegration:ExternalIntegration", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class ExternalIntegrationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of this integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ExternalIntegrationArgs()
        {
        }
    }

    public sealed class ExternalIntegrationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The external ID to use with your IAM role and with `signalfx.aws.Integration`.
        /// </summary>
        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        /// <summary>
        /// The name of this integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
        /// </summary>
        [Input("signalfxAwsAccount")]
        public Input<string>? SignalfxAwsAccount { get; set; }

        public ExternalIntegrationState()
        {
        }
    }
}
