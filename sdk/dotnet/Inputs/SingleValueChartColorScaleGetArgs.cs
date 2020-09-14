// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class SingleValueChartColorScaleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
        /// </summary>
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        /// <summary>
        /// Indicates the lower threshold non-inclusive value for this range.
        /// </summary>
        [Input("gt")]
        public Input<double>? Gt { get; set; }

        /// <summary>
        /// Indicates the lower threshold inclusive value for this range.
        /// </summary>
        [Input("gte")]
        public Input<double>? Gte { get; set; }

        /// <summary>
        /// Indicates the upper threshold non-inculsive value for this range.
        /// </summary>
        [Input("lt")]
        public Input<double>? Lt { get; set; }

        /// <summary>
        /// Indicates the upper threshold inclusive value for this range.
        /// </summary>
        [Input("lte")]
        public Input<double>? Lte { get; set; }

        public SingleValueChartColorScaleGetArgs()
        {
        }
    }
}
