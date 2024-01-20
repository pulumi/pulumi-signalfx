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
    public sealed class DashboardGroupDashboardFilterOverride
    {
        public readonly bool? Negated;
        public readonly string Property;
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
