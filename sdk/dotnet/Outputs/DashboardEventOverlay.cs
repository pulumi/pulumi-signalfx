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
    public sealed class DashboardEventOverlay
    {
        /// <summary>
        /// Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
        /// </summary>
        public readonly string? Color;
        /// <summary>
        /// Text shown in the dropdown when selecting this overlay from the menu.
        /// </summary>
        public readonly string? Label;
        /// <summary>
        /// Show a vertical line for the event. `false` by default.
        /// </summary>
        public readonly bool? Line;
        /// <summary>
        /// Search term used to choose the events shown in the overlay.
        /// </summary>
        public readonly string Signal;
        /// <summary>
        /// Each element specifies a filter to use against the signal specified in the `signal`.
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardEventOverlaySource> Sources;
        /// <summary>
        /// Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private DashboardEventOverlay(
            string? color,

            string? label,

            bool? line,

            string signal,

            ImmutableArray<Outputs.DashboardEventOverlaySource> sources,

            string? type)
        {
            Color = color;
            Label = label;
            Line = line;
            Signal = signal;
            Sources = sources;
            Type = type;
        }
    }
}
