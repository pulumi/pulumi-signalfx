// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class MetricRulesetExceptionRuleRestorationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the restoration job.
        /// </summary>
        [Input("restorationId")]
        public Input<string>? RestorationId { get; set; }

        /// <summary>
        /// Time from which the restoration job will restore archived data, in the form of *nix time in milliseconds
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        /// <summary>
        /// Time to which the restoration job will restore archived data, in the form of *nix time in milliseconds
        /// </summary>
        [Input("stopTime")]
        public Input<string>? StopTime { get; set; }

        public MetricRulesetExceptionRuleRestorationArgs()
        {
        }
        public static new MetricRulesetExceptionRuleRestorationArgs Empty => new MetricRulesetExceptionRuleRestorationArgs();
    }
}