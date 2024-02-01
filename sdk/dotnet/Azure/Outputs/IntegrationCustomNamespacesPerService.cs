// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Azure.Outputs
{

    [OutputType]
    public sealed class IntegrationCustomNamespacesPerService
    {
        /// <summary>
        /// The namespaces to sync
        /// </summary>
        public readonly ImmutableArray<string> Namespaces;
        /// <summary>
        /// The name of the service
        /// </summary>
        public readonly string Service;

        [OutputConstructor]
        private IntegrationCustomNamespacesPerService(
            ImmutableArray<string> namespaces,

            string service)
        {
            Namespaces = namespaces;
            Service = service;
        }
    }
}
