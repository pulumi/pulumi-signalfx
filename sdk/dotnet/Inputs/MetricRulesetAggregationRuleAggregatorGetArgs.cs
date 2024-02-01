// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class MetricRulesetAggregationRuleAggregatorGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dimensions", required: true)]
        private InputList<string>? _dimensions;

        /// <summary>
        /// List of dimensions to keep or drop in aggregated metric
        /// </summary>
        public InputList<string> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<string>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Flag specifying to keep or drop given dimensions
        /// </summary>
        [Input("dropDimensions", required: true)]
        public Input<bool> DropDimensions { get; set; } = null!;

        /// <summary>
        /// The aggregated metric name
        /// </summary>
        [Input("outputName", required: true)]
        public Input<string> OutputName { get; set; } = null!;

        /// <summary>
        /// The type of the aggregator
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public MetricRulesetAggregationRuleAggregatorGetArgs()
        {
        }
        public static new MetricRulesetAggregationRuleAggregatorGetArgs Empty => new MetricRulesetAggregationRuleAggregatorGetArgs();
    }
}
