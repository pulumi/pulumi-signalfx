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
    public sealed class DashboardGroupDashboardFilterOverride
    {
        /// <summary>
        /// If true, only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
        /// </summary>
        public readonly bool? Negated;
        /// <summary>
        /// A metric time series dimension or property name.
        /// </summary>
        public readonly string Property;
        /// <summary>
        /// (Optional) List of of strings (which will be treated as an OR filter on the property).
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private DashboardGroupDashboardFilterOverride(
            bool? negated,

            string property,

            ImmutableArray<string> values)
        {
            Negated = negated;
            Property = property;
            Values = values;
        }
    }
}
