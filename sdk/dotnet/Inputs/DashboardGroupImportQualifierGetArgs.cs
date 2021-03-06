// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardGroupImportQualifierGetArgs : Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.DashboardGroupImportQualifierFilterGetArgs>? _filters;
        public InputList<Inputs.DashboardGroupImportQualifierFilterGetArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.DashboardGroupImportQualifierFilterGetArgs>());
            set => _filters = value;
        }

        [Input("metric")]
        public Input<string>? Metric { get; set; }

        public DashboardGroupImportQualifierGetArgs()
        {
        }
    }
}
