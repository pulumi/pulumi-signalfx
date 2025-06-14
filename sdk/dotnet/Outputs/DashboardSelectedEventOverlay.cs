// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Outputs
{

    [OutputType]
    public sealed class DashboardSelectedEventOverlay
    {
        /// <summary>
        /// Search term used to choose the events shown in the overlay.
        /// </summary>
        public readonly string Signal;
        /// <summary>
        /// Each element specifies a filter to use against the signal specified in the `signal`.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardSelectedEventOverlaySource> Sources;
        /// <summary>
        /// Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private DashboardSelectedEventOverlay(
            string signal,

            ImmutableArray<Outputs.DashboardSelectedEventOverlaySource> sources,

            string? type)
        {
            Signal = signal;
            Sources = sources;
            Type = type;
        }
    }
}
