// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Azure.Inputs
{

    public sealed class IntegrationCustomNamespacesPerServiceGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("namespaces", required: true)]
        private InputList<string>? _namespaces;

        /// <summary>
        /// The namespaces to sync
        /// </summary>
        public InputList<string> Namespaces
        {
            get => _namespaces ?? (_namespaces = new InputList<string>());
            set => _namespaces = value;
        }

        /// <summary>
        /// The name of the service
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public IntegrationCustomNamespacesPerServiceGetArgs()
        {
        }
        public static new IntegrationCustomNamespacesPerServiceGetArgs Empty => new IntegrationCustomNamespacesPerServiceGetArgs();
    }
}
