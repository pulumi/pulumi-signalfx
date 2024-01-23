// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Aws.Outputs
{

    [OutputType]
    public sealed class IntegrationCustomNamespaceSyncRule
    {
        public readonly string? DefaultAction;
        public readonly string? FilterAction;
        public readonly string? FilterSource;
        public readonly string Namespace;

        [OutputConstructor]
        private IntegrationCustomNamespaceSyncRule(
            string? defaultAction,

            string? filterAction,

            string? filterSource,

            string @namespace)
        {
            DefaultAction = defaultAction;
            FilterAction = filterAction;
            FilterSource = filterSource;
            Namespace = @namespace;
        }
    }
}
