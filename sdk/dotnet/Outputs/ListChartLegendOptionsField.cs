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
    public sealed class ListChartLegendOptionsField
    {
        /// <summary>
        /// (true by default) Determines if this property is displayed in the data table.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The name of a property to hide or show in the data table.
        /// </summary>
        public readonly string Property;

        [OutputConstructor]
        private ListChartLegendOptionsField(
            bool? enabled,

            string property)
        {
            Enabled = enabled;
            Property = property;
        }
    }
}
