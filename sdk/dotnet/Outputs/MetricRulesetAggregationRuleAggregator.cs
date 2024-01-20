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
    public sealed class MetricRulesetAggregationRuleAggregator
    {
        public readonly ImmutableArray<string> Dimensions;
        public readonly bool DropDimensions;
        public readonly string OutputName;
        public readonly string Type;

        [OutputConstructor]
        private MetricRulesetAggregationRuleAggregator(
            ImmutableArray<string> dimensions,

            bool dropDimensions,

            string outputName,

            string type)
        {
            Dimensions = dimensions;
            DropDimensions = dropDimensions;
            OutputName = outputName;
            Type = type;
        }
    }
}
