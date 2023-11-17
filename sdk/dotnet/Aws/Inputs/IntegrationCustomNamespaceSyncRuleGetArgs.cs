// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Aws.Inputs
{

    public sealed class IntegrationCustomNamespaceSyncRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        [Input("defaultAction")]
        public Input<string>? DefaultAction { get; set; }

        /// <summary>
        /// Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        [Input("filterAction")]
        public Input<string>? FilterAction { get; set; }

        /// <summary>
        /// Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        /// </summary>
        [Input("filterSource")]
        public Input<string>? FilterSource { get; set; }

        /// <summary>
        /// An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
        /// </summary>
        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        public IntegrationCustomNamespaceSyncRuleGetArgs()
        {
        }
        public static new IntegrationCustomNamespaceSyncRuleGetArgs Empty => new IntegrationCustomNamespaceSyncRuleGetArgs();
    }
}
