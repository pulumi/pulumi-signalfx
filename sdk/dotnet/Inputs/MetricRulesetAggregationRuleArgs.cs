// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class MetricRulesetAggregationRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("aggregators", required: true)]
        private InputList<Inputs.MetricRulesetAggregationRuleAggregatorArgs>? _aggregators;

        /// <summary>
        /// Aggregator object
        /// </summary>
        public InputList<Inputs.MetricRulesetAggregationRuleAggregatorArgs> Aggregators
        {
            get => _aggregators ?? (_aggregators = new InputList<Inputs.MetricRulesetAggregationRuleAggregatorArgs>());
            set => _aggregators = value;
        }

        /// <summary>
        /// Information about an aggregation rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When false, this rule will not generate aggregated MTSs
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("matchers", required: true)]
        private InputList<Inputs.MetricRulesetAggregationRuleMatcherArgs>? _matchers;

        /// <summary>
        /// Matcher object
        /// </summary>
        public InputList<Inputs.MetricRulesetAggregationRuleMatcherArgs> Matchers
        {
            get => _matchers ?? (_matchers = new InputList<Inputs.MetricRulesetAggregationRuleMatcherArgs>());
            set => _matchers = value;
        }

        /// <summary>
        /// name of the aggregation rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public MetricRulesetAggregationRuleArgs()
        {
        }
        public static new MetricRulesetAggregationRuleArgs Empty => new MetricRulesetAggregationRuleArgs();
    }
}
