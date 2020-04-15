// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class TimeChartAxisLeftArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A line to draw as a high watermark.
        /// </summary>
        [Input("highWatermark")]
        public Input<double>? HighWatermark { get; set; }

        /// <summary>
        /// A label to attach to the high watermark line.
        /// </summary>
        [Input("highWatermarkLabel")]
        public Input<string>? HighWatermarkLabel { get; set; }

        /// <summary>
        /// Label used in the publish statement that displays the event query you want to customize.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// A line to draw as a low watermark.
        /// </summary>
        [Input("lowWatermark")]
        public Input<double>? LowWatermark { get; set; }

        /// <summary>
        /// A label to attach to the low watermark line.
        /// </summary>
        [Input("lowWatermarkLabel")]
        public Input<string>? LowWatermarkLabel { get; set; }

        /// <summary>
        /// The maximum value for the right axis.
        /// </summary>
        [Input("maxValue")]
        public Input<double>? MaxValue { get; set; }

        /// <summary>
        /// The minimum value for the right axis.
        /// </summary>
        [Input("minValue")]
        public Input<double>? MinValue { get; set; }

        [Input("watermarks")]
        private InputList<Inputs.TimeChartAxisLeftWatermarkArgs>? _watermarks;
        public InputList<Inputs.TimeChartAxisLeftWatermarkArgs> Watermarks
        {
            get => _watermarks ?? (_watermarks = new InputList<Inputs.TimeChartAxisLeftWatermarkArgs>());
            set => _watermarks = value;
        }

        public TimeChartAxisLeftArgs()
        {
        }
    }
}
