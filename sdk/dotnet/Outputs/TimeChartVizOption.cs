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
    public sealed class TimeChartVizOption
    {
        /// <summary>
        /// Y-axis associated with values for this plot. Must be either `right` or `left`.
        /// </summary>
        public readonly string? Axis;
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
        /// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
        /// </summary>
        public readonly string? PlotType;
        /// <summary>
        /// , `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
        /// </summary>
        public readonly string? ValuePrefix;
        /// <summary>
        /// An arbitrary suffix to display with the value of this plot
        /// </summary>
        public readonly string? ValueSuffix;
        /// <summary>
        /// A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
        /// </summary>
        public readonly string? ValueUnit;

        [OutputConstructor]
        private TimeChartVizOption(
            string? axis,

            string? color,

            string? displayName,

            string label,

            string? plotType,

            string? valuePrefix,

            string? valueSuffix,

            string? valueUnit)
        {
            Axis = axis;
            Color = color;
            DisplayName = displayName;
            Label = label;
            PlotType = plotType;
            ValuePrefix = valuePrefix;
            ValueSuffix = valueSuffix;
            ValueUnit = valueUnit;
        }
    }
}
