// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Outputs
{

    [OutputType]
    public sealed class MetricRulesetAggregationRuleMatcher
    {
        /// <summary>
        /// List of filters to filter the set of input MTSs
        /// </summary>
        public readonly ImmutableArray<Outputs.MetricRulesetAggregationRuleMatcherFilter> Filters;
        /// <summary>
        /// Type of matcher. Must always be "dimension"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private MetricRulesetAggregationRuleMatcher(
            ImmutableArray<Outputs.MetricRulesetAggregationRuleMatcherFilter> filters,

            string type)
        {
            Filters = filters;
            Type = type;
        }
    }
}
