// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class MetricRulesetRoutingRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// end destination of the input metric. Must be `RealTime` or `Drop`
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        public MetricRulesetRoutingRuleGetArgs()
        {
        }
        public static new MetricRulesetRoutingRuleGetArgs Empty => new MetricRulesetRoutingRuleGetArgs();
    }
}
