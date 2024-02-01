// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Outputs
{

    [OutputType]
    public sealed class TimeChartHistogramOption
    {
        /// <summary>
        /// Base color theme to use for the graph.
        /// </summary>
        public readonly string? ColorTheme;

        [OutputConstructor]
        private TimeChartHistogramOption(string? colorTheme)
        {
            ColorTheme = colorTheme;
        }
    }
}
