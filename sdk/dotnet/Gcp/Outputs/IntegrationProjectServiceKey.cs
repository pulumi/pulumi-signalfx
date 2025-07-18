// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Gcp.Outputs
{

    [OutputType]
    public sealed class IntegrationProjectServiceKey
    {
        public readonly string ProjectId;
        public readonly string ProjectKey;

        [OutputConstructor]
        private IntegrationProjectServiceKey(
            string projectId,

            string projectKey)
        {
            ProjectId = projectId;
            ProjectKey = projectKey;
        }
    }
}
