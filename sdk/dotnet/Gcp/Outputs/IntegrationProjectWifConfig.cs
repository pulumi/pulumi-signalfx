// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Gcp.Outputs
{

    [OutputType]
    public sealed class IntegrationProjectWifConfig
    {
        public readonly string ProjectId;
        public readonly string WifConfig;

        [OutputConstructor]
        private IntegrationProjectWifConfig(
            string projectId,

            string wifConfig)
        {
            ProjectId = projectId;
            WifConfig = wifConfig;
        }
    }
}