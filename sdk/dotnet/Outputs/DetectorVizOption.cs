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
    public sealed class DetectorVizOption
    {
        /// <summary>
        /// Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
        /// </summary>
        public readonly string? Color;
        /// <summary>
        /// Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Label used in the publish statement that displays the plot (metric time series data) you want to customize.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// , `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
        /// 
        /// **Notes**
        /// 
        /// It is highly recommended that you use both `max_delay` in your detector configuration and an `extrapolation` policy in your program text to reduce false positives/negatives.
        /// 
        /// `max_delay` allows SignalFx to continue with computation if there is a lag in receiving data points.
        /// 
        /// `extrapolation` allows you to specify how to handle missing data. An extrapolation policy can be added to individual signals by updating the data block in your `program_text`.
        /// 
        /// See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info.
        /// </summary>
        public readonly string? ValuePrefix;
        public readonly string? ValueSuffix;
        /// <summary>
        /// A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
        /// </summary>
        public readonly string? ValueUnit;

        [OutputConstructor]
        private DetectorVizOption(
            string? color,

            string? displayName,

            string label,

            string? valuePrefix,

            string? valueSuffix,

            string? valueUnit)
        {
            Color = color;
            DisplayName = displayName;
            Label = label;
            ValuePrefix = valuePrefix;
            ValueSuffix = valueSuffix;
            ValueUnit = valueUnit;
        }
    }
}
