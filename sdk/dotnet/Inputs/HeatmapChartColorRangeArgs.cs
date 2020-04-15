// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class HeatmapChartColorRangeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The color range to use. Must be either gray, blue, navy, orange, yellow, magenta, purple, violet, lilac, green, aquamarine.
        /// </summary>
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        /// <summary>
        /// The maximum value within the coloring range.
        /// </summary>
        [Input("maxValue")]
        public Input<double>? MaxValue { get; set; }

        /// <summary>
        /// The minimum value within the coloring range.
        /// </summary>
        [Input("minValue")]
        public Input<double>? MinValue { get; set; }

        public HeatmapChartColorRangeArgs()
        {
        }
    }
}
