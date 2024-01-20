// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class HeatmapChartColorScaleArgs : global::Pulumi.ResourceArgs
    {
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        [Input("gt")]
        public Input<double>? Gt { get; set; }

        [Input("gte")]
        public Input<double>? Gte { get; set; }

        [Input("lt")]
        public Input<double>? Lt { get; set; }

        [Input("lte")]
        public Input<double>? Lte { get; set; }

        public HeatmapChartColorScaleArgs()
        {
        }
        public static new HeatmapChartColorScaleArgs Empty => new HeatmapChartColorScaleArgs();
    }
}
