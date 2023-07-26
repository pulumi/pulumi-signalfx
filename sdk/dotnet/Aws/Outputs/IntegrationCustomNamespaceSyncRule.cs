// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Aws.Outputs
{

    [OutputType]
    public sealed class IntegrationCustomNamespaceSyncRule
    {
        /// <summary>
        /// Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        public readonly string? DefaultAction;
        /// <summary>
        /// Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        /// </summary>
        public readonly string? FilterAction;
        /// <summary>
        /// Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        /// </summary>
        public readonly string? FilterSource;
        /// <summary>
        /// An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See `services` field description below for additional information.
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private IntegrationCustomNamespaceSyncRule(
            string? defaultAction,

            string? filterAction,

            string? filterSource,

            string @namespace)
        {
            DefaultAction = defaultAction;
            FilterAction = filterAction;
            FilterSource = filterSource;
            Namespace = @namespace;
        }
    }
}
