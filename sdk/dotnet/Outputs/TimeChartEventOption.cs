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
    public sealed class TimeChartEventOption
    {
        /// <summary>
        /// Color to use
        /// </summary>
        public readonly string? Color;
        /// <summary>
        /// Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The label used in the publish statement that displays the events you want to customize
        /// </summary>
        public readonly string Label;

        [OutputConstructor]
        private TimeChartEventOption(
            string? color,

            string? displayName,

            string label)
        {
            Color = color;
            DisplayName = displayName;
            Label = label;
        }
    }
}
