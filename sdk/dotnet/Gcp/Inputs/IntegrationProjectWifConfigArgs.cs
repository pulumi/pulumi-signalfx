// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Gcp.Inputs
{

    public sealed class IntegrationProjectWifConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("wifConfig", required: true)]
        private Input<string>? _wifConfig;
        public Input<string>? WifConfig
        {
            get => _wifConfig;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _wifConfig = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public IntegrationProjectWifConfigArgs()
        {
        }
        public static new IntegrationProjectWifConfigArgs Empty => new IntegrationProjectWifConfigArgs();
    }
}
