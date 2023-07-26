// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Azure.Inputs
{

    public sealed class IntegrationResourceFilterRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Expression that selects the data that SignalFx should sync for the resource associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function. The source of each filter rule must be in the form filter('key', 'value'). You can join multiple filter statements using the and and or operators. Referenced keys are limited to tags and must start with the azure_tag_ prefix.
        /// </summary>
        [Input("filterSource", required: true)]
        public Input<string> FilterSource { get; set; } = null!;

        public IntegrationResourceFilterRuleGetArgs()
        {
        }
        public static new IntegrationResourceFilterRuleGetArgs Empty => new IntegrationResourceFilterRuleGetArgs();
    }
}
