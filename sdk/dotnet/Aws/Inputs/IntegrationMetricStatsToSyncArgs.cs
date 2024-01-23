// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Aws.Inputs
{

    public sealed class IntegrationMetricStatsToSyncArgs : global::Pulumi.ResourceArgs
    {
        [Input("metric", required: true)]
        public Input<string> Metric { get; set; } = null!;

        [Input("namespace", required: true)]
        public Input<string> Namespace { get; set; } = null!;

        [Input("stats", required: true)]
        private InputList<string>? _stats;
        public InputList<string> Stats
        {
            get => _stats ?? (_stats = new InputList<string>());
            set => _stats = value;
        }

        public IntegrationMetricStatsToSyncArgs()
        {
        }
        public static new IntegrationMetricStatsToSyncArgs Empty => new IntegrationMetricStatsToSyncArgs();
    }
}
