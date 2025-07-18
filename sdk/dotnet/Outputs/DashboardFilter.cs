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
    public sealed class DashboardFilter
    {
        /// <summary>
        /// If true, this filter will also match data that doesn't have this property at all.
        /// </summary>
        public readonly bool? ApplyIfExist;
        /// <summary>
        /// Whether this filter should be a not filter. `false` by default.
        /// </summary>
        public readonly bool? Negated;
        /// <summary>
        /// A metric time series dimension or property name.
        /// </summary>
        public readonly string Property;
        /// <summary>
        /// List of of strings (which will be treated as an OR filter on the property).
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private DashboardFilter(
            bool? applyIfExist,

            bool? negated,

            string property,

            ImmutableArray<string> values)
        {
            ApplyIfExist = applyIfExist;
            Negated = negated;
            Property = property;
            Values = values;
        }
    }
}
