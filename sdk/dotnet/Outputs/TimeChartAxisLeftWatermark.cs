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
    public sealed class TimeChartAxisLeftWatermark
    {
        /// <summary>
        /// Label of the left axis.
        /// </summary>
        public readonly string? Label;
        public readonly double Value;

        [OutputConstructor]
        private TimeChartAxisLeftWatermark(
            string? label,

            double value)
        {
            Label = label;
            Value = value;
        }
    }
}
