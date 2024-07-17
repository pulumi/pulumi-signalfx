// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Outputs
{

    [OutputType]
    public sealed class DashboardSelectedEventOverlaySource
    {
        /// <summary>
        /// If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
        /// </summary>
        public readonly bool? Negated;
        /// <summary>
        /// The name of a dimension to filter against.
        /// </summary>
        public readonly string Property;
        /// <summary>
        /// A list of values to be used with the `property`, they will be combined via `OR`.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private DashboardSelectedEventOverlaySource(
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