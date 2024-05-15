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
    public sealed class HeatmapChartColorRange
    {
        /// <summary>
        /// The color range to use. The starting hex color value for data values in a heatmap chart. Specify the value as a 6-character hexadecimal value preceded by the '#' character, for example "#ea1849" (grass green).
        /// </summary>
        public readonly string Color;
        /// <summary>
        /// The maximum value within the coloring range.
        /// </summary>
        public readonly double? MaxValue;
        /// <summary>
        /// The minimum value within the coloring range.
        /// </summary>
        public readonly double? MinValue;

        [OutputConstructor]
        private HeatmapChartColorRange(
            string color,

            double? maxValue,

            double? minValue)
        {
            Color = color;
            MaxValue = maxValue;
            MinValue = minValue;
        }
    }
}
