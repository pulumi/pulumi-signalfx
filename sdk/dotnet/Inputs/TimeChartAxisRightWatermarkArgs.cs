// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class TimeChartAxisRightWatermarkArgs : global::Pulumi.ResourceArgs
    {
        [Input("label")]
        public Input<string>? Label { get; set; }

        [Input("value", required: true)]
        public Input<double> Value { get; set; } = null!;

        public TimeChartAxisRightWatermarkArgs()
        {
        }
        public static new TimeChartAxisRightWatermarkArgs Empty => new TimeChartAxisRightWatermarkArgs();
    }
}
