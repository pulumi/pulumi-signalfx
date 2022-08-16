// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardChartGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the chart to display.
        /// </summary>
        [Input("chartId", required: true)]
        public Input<string> ChartId { get; set; } = null!;

        /// <summary>
        /// Column number for the layout.
        /// </summary>
        [Input("column")]
        public Input<int>? Column { get; set; }

        /// <summary>
        /// How many rows every chart should take up (greater than or equal to 1). 1 by default.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// The row to show the chart in (zero-based); if `height &gt; 1`, this value represents the topmost row of the chart (greater than or equal to `0`).
        /// </summary>
        [Input("row")]
        public Input<int>? Row { get; set; }

        /// <summary>
        /// How many columns (out of a total of `12`) every chart should take up (between `1` and `12`). `12` by default.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public DashboardChartGetArgs()
        {
        }
        public static new DashboardChartGetArgs Empty => new DashboardChartGetArgs();
    }
}
