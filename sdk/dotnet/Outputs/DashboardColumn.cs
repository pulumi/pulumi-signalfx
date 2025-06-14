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
    public sealed class DashboardColumn
    {
        /// <summary>
        /// List of IDs of the charts to display.
        /// </summary>
        public readonly ImmutableArray<string> ChartIds;
        /// <summary>
        /// Column number for the layout.
        /// </summary>
        public readonly int? Column;
        /// <summary>
        /// How many rows every chart should take up (greater than or equal to 1). 1 by default.
        /// </summary>
        public readonly int? Height;
        /// <summary>
        /// How many columns (out of a total of `12`) every chart should take up (between `1` and `12`). `12` by default.
        /// </summary>
        public readonly int? Width;

        [OutputConstructor]
        private DashboardColumn(
            ImmutableArray<string> chartIds,

            int? column,

            int? height,

            int? width)
        {
            ChartIds = chartIds;
            Column = column;
            Height = height;
            Width = width;
        }
    }
}
