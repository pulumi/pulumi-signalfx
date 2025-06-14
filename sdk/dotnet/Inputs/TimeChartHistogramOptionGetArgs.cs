// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class TimeChartHistogramOptionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color to use. Must be one of red, gold, iris, green, jade, gray, blue, azure, navy, brown, orange, yellow, magenta, cerise, pink, violet, purple, lilac, emerald, chartreuse, yellowgreen, aquamarine.
        /// </summary>
        [Input("colorTheme")]
        public Input<string>? ColorTheme { get; set; }

        public TimeChartHistogramOptionGetArgs()
        {
        }
        public static new TimeChartHistogramOptionGetArgs Empty => new TimeChartHistogramOptionGetArgs();
    }
}
