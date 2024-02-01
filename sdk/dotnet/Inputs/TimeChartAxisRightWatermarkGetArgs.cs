// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class TimeChartAxisRightWatermarkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Label to display associated with the watermark line
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Axis value where the watermark line will be displayed
        /// </summary>
        [Input("value", required: true)]
        public Input<double> Value { get; set; } = null!;

        public TimeChartAxisRightWatermarkGetArgs()
        {
        }
        public static new TimeChartAxisRightWatermarkGetArgs Empty => new TimeChartAxisRightWatermarkGetArgs();
    }
}
