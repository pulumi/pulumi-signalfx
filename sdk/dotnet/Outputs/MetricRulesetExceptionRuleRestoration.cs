// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Outputs
{

    [OutputType]
    public sealed class MetricRulesetExceptionRuleRestoration
    {
        /// <summary>
        /// ID of the restoration job.
        /// </summary>
        public readonly string? RestorationId;
        /// <summary>
        /// Time from which the restoration job will restore archived data, in the form of *nix time in milliseconds
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// Time to which the restoration job will restore archived data, in the form of *nix time in milliseconds
        /// </summary>
        public readonly string? StopTime;

        [OutputConstructor]
        private MetricRulesetExceptionRuleRestoration(
            string? restorationId,

            string startTime,

            string? stopTime)
        {
            RestorationId = restorationId;
            StartTime = startTime;
            StopTime = stopTime;
        }
    }
}
