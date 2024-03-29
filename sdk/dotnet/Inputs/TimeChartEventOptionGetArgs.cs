// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class TimeChartEventOptionGetArgs : global::Pulumi.ResourceArgs
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
        /// The label used in the publish statement that displays the events you want to customize
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        public TimeChartEventOptionGetArgs()
        {
        }
        public static new TimeChartEventOptionGetArgs Empty => new TimeChartEventOptionGetArgs();
    }
}
