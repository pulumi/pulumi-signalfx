// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class MetricRulesetAggregationRuleMatcherArgs : global::Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.MetricRulesetAggregationRuleMatcherFilterArgs>? _filters;

        /// <summary>
        /// List of filters to filter the set of input MTSs
        /// </summary>
        public InputList<Inputs.MetricRulesetAggregationRuleMatcherFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.MetricRulesetAggregationRuleMatcherFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Type of matcher. Must always be "dimension"
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public MetricRulesetAggregationRuleMatcherArgs()
        {
        }
        public static new MetricRulesetAggregationRuleMatcherArgs Empty => new MetricRulesetAggregationRuleMatcherArgs();
    }
}
