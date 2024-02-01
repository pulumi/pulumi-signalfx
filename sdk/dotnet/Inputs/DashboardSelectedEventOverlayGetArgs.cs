// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardSelectedEventOverlayGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search term used to define events
        /// </summary>
        [Input("signal", required: true)]
        public Input<string> Signal { get; set; } = null!;

        [Input("sources")]
        private InputList<Inputs.DashboardSelectedEventOverlaySourceGetArgs>? _sources;
        public InputList<Inputs.DashboardSelectedEventOverlaySourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.DashboardSelectedEventOverlaySourceGetArgs>());
            set => _sources = value;
        }

        /// <summary>
        /// Source for this event's data. Can be "eventTimeSeries" (default) or "detectorEvents".
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DashboardSelectedEventOverlayGetArgs()
        {
        }
        public static new DashboardSelectedEventOverlayGetArgs Empty => new DashboardSelectedEventOverlayGetArgs();
    }
}
