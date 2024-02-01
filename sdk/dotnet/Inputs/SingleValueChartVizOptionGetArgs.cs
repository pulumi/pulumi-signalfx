// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class SingleValueChartVizOptionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color to use
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The label used in the publish statement that displays the plot (metric time series data) you want to customize
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// An arbitrary prefix to display with the value of this plot
        /// </summary>
        [Input("valuePrefix")]
        public Input<string>? ValuePrefix { get; set; }

        /// <summary>
        /// An arbitrary suffix to display with the value of this plot
        /// </summary>
        [Input("valueSuffix")]
        public Input<string>? ValueSuffix { get; set; }

        /// <summary>
        /// A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
        /// </summary>
        [Input("valueUnit")]
        public Input<string>? ValueUnit { get; set; }

        public SingleValueChartVizOptionGetArgs()
        {
        }
        public static new SingleValueChartVizOptionGetArgs Empty => new SingleValueChartVizOptionGetArgs();
    }
}
