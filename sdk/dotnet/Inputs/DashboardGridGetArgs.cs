// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardGridGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("chartIds", required: true)]
        private InputList<string>? _chartIds;

        /// <summary>
        /// List of IDs of the charts to display.
        /// </summary>
        public InputList<string> ChartIds
        {
            get => _chartIds ?? (_chartIds = new InputList<string>());
            set => _chartIds = value;
        }

        /// <summary>
        /// How many rows every chart should take up (greater than or equal to 1). 1 by default.
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// How many columns (out of a total of `12`) every chart should take up (between `1` and `12`). `12` by default.
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public DashboardGridGetArgs()
        {
        }
        public static new DashboardGridGetArgs Empty => new DashboardGridGetArgs();
    }
}
