// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class TimeChartEventOptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("color")]
        public Input<string>? Color { get; set; }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        public TimeChartEventOptionArgs()
        {
        }
        public static new TimeChartEventOptionArgs Empty => new TimeChartEventOptionArgs();
    }
}
