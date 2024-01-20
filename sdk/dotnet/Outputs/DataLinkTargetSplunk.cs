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
    public sealed class DataLinkTargetSplunk
    {
        public readonly string Name;
        public readonly ImmutableDictionary<string, string>? PropertyKeyMapping;

        [OutputConstructor]
        private DataLinkTargetSplunk(
            string name,

            ImmutableDictionary<string, string>? propertyKeyMapping)
        {
            Name = name;
            PropertyKeyMapping = propertyKeyMapping;
        }
    }
}
