// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class WebhookIntegrationHeaderArgs : Pulumi.ResourceArgs
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

        public WebhookIntegrationHeaderArgs()
        {
        }
    }
}
