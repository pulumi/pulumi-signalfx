// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class ListChartVizOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
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
        /// , `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
        /// </summary>
        [Input("valuePrefix")]
        public Input<string>? ValuePrefix { get; set; }

        [Input("valueSuffix")]
        public Input<string>? ValueSuffix { get; set; }

        /// <summary>
        /// A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
        /// </summary>
        [Input("valueUnit")]
        public Input<string>? ValueUnit { get; set; }

        public ListChartVizOptionArgs()
        {
        }
        public static new ListChartVizOptionArgs Empty => new ListChartVizOptionArgs();
    }
}
