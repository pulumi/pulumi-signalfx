// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class AlertMutingRuleFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if this is a "not" filter. Defaults to `false`.
        /// </summary>
        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        /// <summary>
        /// The property to filter.
        /// </summary>
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        /// <summary>
        /// The property value to filter.
        /// </summary>
        [Input("propertyValue", required: true)]
        public Input<string> PropertyValue { get; set; } = null!;

        public AlertMutingRuleFilterArgs()
        {
        }
    }
}