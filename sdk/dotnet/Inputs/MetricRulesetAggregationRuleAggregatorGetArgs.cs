// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// List of dimensions to either be kept or dropped in the new aggregated MTSs
        /// </summary>
        public InputList<string> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<string>());
            set => _dimensions = value;
        }

        /// <summary>
        /// when true, the specified dimensions will be dropped from the aggregated MTSs
        /// </summary>
        [Input("dropDimensions", required: true)]
        public Input<bool> DropDimensions { get; set; } = null!;

        /// <summary>
        /// name of the new aggregated metric
        /// </summary>
        [Input("outputName", required: true)]
        public Input<string> OutputName { get; set; } = null!;

        /// <summary>
        /// Type of aggregator. Must always be "rollup"
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public MetricRulesetAggregationRuleAggregatorGetArgs()
        {
        }
        public static new MetricRulesetAggregationRuleAggregatorGetArgs Empty => new MetricRulesetAggregationRuleAggregatorGetArgs();
    }
}
