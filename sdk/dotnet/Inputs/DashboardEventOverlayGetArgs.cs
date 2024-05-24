// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardEventOverlayGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// Text shown in the dropdown when selecting this overlay from the menu.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Show a vertical line for the event. `false` by default.
        /// </summary>
        [Input("line")]
        public Input<bool>? Line { get; set; }

        /// <summary>
        /// Search term used to choose the events shown in the overlay.
        /// </summary>
        [Input("signal", required: true)]
        public Input<string> Signal { get; set; } = null!;

        [Input("sources")]
        private InputList<Inputs.DashboardEventOverlaySourceGetArgs>? _sources;

        /// <summary>
        /// Each element specifies a filter to use against the signal specified in the `signal`.
        /// </summary>
        public InputList<Inputs.DashboardEventOverlaySourceGetArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.DashboardEventOverlaySourceGetArgs>());
            set => _sources = value;
        }

        /// <summary>
        /// Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DashboardEventOverlayGetArgs()
        {
        }
        public static new DashboardEventOverlayGetArgs Empty => new DashboardEventOverlayGetArgs();
    }
}
