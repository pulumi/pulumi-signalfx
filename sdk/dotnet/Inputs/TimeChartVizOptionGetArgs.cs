// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class TimeChartVizOptionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Y-axis associated with values for this plot. Must be either `right` or `left`.
        /// </summary>
        [Input("axis")]
        public Input<string>? Axis { get; set; }

        /// <summary>
        /// Color to use. Must be one of red, gold, iris, green, jade, gray, blue, azure, navy, brown, orange, yellow, magenta, cerise, pink, violet, purple, lilac, emerald, chartreuse, yellowgreen, aquamarine.
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Label used in the publish statement that displays the plot (metric time series data) you want to customize.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
        /// </summary>
        [Input("plotType")]
        public Input<string>? PlotType { get; set; }

        /// <summary>
        /// , `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
        /// </summary>
        [Input("valuePrefix")]
        public Input<string>? ValuePrefix { get; set; }

        /// <summary>
        /// An arbitrary suffix to display with the value of this plot
        /// </summary>
        [Input("valueSuffix")]
        public Input<string>? ValueSuffix { get; set; }

        /// <summary>
        /// A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
        /// </summary>
        [Input("valueUnit")]
        public Input<string>? ValueUnit { get; set; }

        public TimeChartVizOptionGetArgs()
        {
        }
        public static new TimeChartVizOptionGetArgs Empty => new TimeChartVizOptionGetArgs();
    }
}
