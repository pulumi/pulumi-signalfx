// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class ListChartLegendOptionsFieldGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (true by default) Determines if this property is displayed in the data table.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The name of a property to hide or show in the data table.
        /// </summary>
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        public ListChartLegendOptionsFieldGetArgs()
        {
        }
        public static new ListChartLegendOptionsFieldGetArgs Empty => new ListChartLegendOptionsFieldGetArgs();
    }
}
