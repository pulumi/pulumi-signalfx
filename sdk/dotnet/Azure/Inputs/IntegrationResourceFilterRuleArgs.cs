// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Azure.Inputs
{

    public sealed class IntegrationResourceFilterRuleArgs : Pulumi.ResourceArgs
    {
        [Input("filter", required: true)]
        public Input<Inputs.IntegrationResourceFilterRuleFilterArgs> Filter { get; set; } = null!;

        public IntegrationResourceFilterRuleArgs()
        {
        }
    }
}
