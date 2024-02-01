// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class MetricRulesetAggregationRuleMatcherGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.MetricRulesetAggregationRuleMatcherFilterGetArgs>? _filters;

        /// <summary>
        /// List of filters to match on
        /// </summary>
        public InputList<Inputs.MetricRulesetAggregationRuleMatcherFilterGetArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.MetricRulesetAggregationRuleMatcherFilterGetArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The type of the matcher
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public MetricRulesetAggregationRuleMatcherGetArgs()
        {
        }
        public static new MetricRulesetAggregationRuleMatcherGetArgs Empty => new MetricRulesetAggregationRuleMatcherGetArgs();
    }
}
