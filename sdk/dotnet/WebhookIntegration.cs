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
    /// SignalFx Webhook integration.
    /// 
    /// &gt; **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/webhook_integration.html.markdown.
    /// </summary>
    public partial class WebhookIntegration : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// A header to send with the request
        /// </summary>
        [Output("headers")]
        public Output<ImmutableArray<Outputs.WebhookIntegrationHeaders>> Headers { get; private set; } = null!;

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("sharedSecret")]
        public Output<string?> SharedSecret { get; private set; } = null!;

        /// <summary>
        /// The URL to request
        /// </summary>
        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;


        /// <summary>
        /// Create a WebhookIntegration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebhookIntegration(string name, WebhookIntegrationArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/webhookIntegration:WebhookIntegration", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private WebhookIntegration(string name, Input<string> id, WebhookIntegrationState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/webhookIntegration:WebhookIntegration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WebhookIntegration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebhookIntegration Get(string name, Input<string> id, WebhookIntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new WebhookIntegration(name, id, state, options);
        }
    }

    public sealed class WebhookIntegrationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("headers")]
        private InputList<Inputs.WebhookIntegrationHeadersArgs>? _headers;

        /// <summary>
        /// A header to send with the request
        /// </summary>
        public InputList<Inputs.WebhookIntegrationHeadersArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.WebhookIntegrationHeadersArgs>());
            set => _headers = value;
        }

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("sharedSecret")]
        public Input<string>? SharedSecret { get; set; }

        /// <summary>
        /// The URL to request
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public WebhookIntegrationArgs()
        {
        }
    }

    public sealed class WebhookIntegrationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("headers")]
        private InputList<Inputs.WebhookIntegrationHeadersGetArgs>? _headers;

        /// <summary>
        /// A header to send with the request
        /// </summary>
        public InputList<Inputs.WebhookIntegrationHeadersGetArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.WebhookIntegrationHeadersGetArgs>());
            set => _headers = value;
        }

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("sharedSecret")]
        public Input<string>? SharedSecret { get; set; }

        /// <summary>
        /// The URL to request
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public WebhookIntegrationState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class WebhookIntegrationHeadersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of the header to send
        /// </summary>
        [Input("headerKey", required: true)]
        public Input<string> HeaderKey { get; set; } = null!;

        /// <summary>
        /// The value of the header to send
        /// </summary>
        [Input("headerValue", required: true)]
        public Input<string> HeaderValue { get; set; } = null!;

        public WebhookIntegrationHeadersArgs()
        {
        }
    }

    public sealed class WebhookIntegrationHeadersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of the header to send
        /// </summary>
        [Input("headerKey", required: true)]
        public Input<string> HeaderKey { get; set; } = null!;

        /// <summary>
        /// The value of the header to send
        /// </summary>
        [Input("headerValue", required: true)]
        public Input<string> HeaderValue { get; set; } = null!;

        public WebhookIntegrationHeadersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class WebhookIntegrationHeaders
    {
        /// <summary>
        /// The key of the header to send
        /// </summary>
        public readonly string HeaderKey;
        /// <summary>
        /// The value of the header to send
        /// </summary>
        public readonly string HeaderValue;

        [OutputConstructor]
        private WebhookIntegrationHeaders(
            string headerKey,
            string headerValue)
        {
            HeaderKey = headerKey;
            HeaderValue = headerValue;
        }
    }
    }
}
