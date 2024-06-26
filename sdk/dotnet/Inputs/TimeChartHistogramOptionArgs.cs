// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class TimeChartHistogramOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine, red, gold, greenyellow, chartreuse, jade
        /// </summary>
        [Input("colorTheme")]
        public Input<string>? ColorTheme { get; set; }

        public TimeChartHistogramOptionArgs()
        {
        }
        public static new TimeChartHistogramOptionArgs Empty => new TimeChartHistogramOptionArgs();
    }
}
