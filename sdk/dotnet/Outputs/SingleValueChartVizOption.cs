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
    public sealed class SingleValueChartVizOption
    {
        /// <summary>
        /// The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
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
        public readonly string? ValuePrefix;
        public readonly string? ValueSuffix;
        /// <summary>
        /// A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
        /// * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
        /// </summary>
        public readonly string? ValueUnit;

        [OutputConstructor]
        private SingleValueChartVizOption(
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
