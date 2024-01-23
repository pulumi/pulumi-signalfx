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
    public sealed class TimeChartAxisLeft
    {
        public readonly double? HighWatermark;
        public readonly string? HighWatermarkLabel;
        public readonly string? Label;
        public readonly double? LowWatermark;
        public readonly string? LowWatermarkLabel;
        public readonly double? MaxValue;
        public readonly double? MinValue;
        public readonly ImmutableArray<Outputs.TimeChartAxisLeftWatermark> Watermarks;

        [OutputConstructor]
        private TimeChartAxisLeft(
            double? highWatermark,

            string? highWatermarkLabel,

            string? label,

            double? lowWatermark,

            string? lowWatermarkLabel,

            double? maxValue,

            double? minValue,

            ImmutableArray<Outputs.TimeChartAxisLeftWatermark> watermarks)
        {
            HighWatermark = highWatermark;
            HighWatermarkLabel = highWatermarkLabel;
            Label = label;
            LowWatermark = lowWatermark;
            LowWatermarkLabel = lowWatermarkLabel;
            MaxValue = maxValue;
            MinValue = minValue;
            Watermarks = watermarks;
        }
    }
}
