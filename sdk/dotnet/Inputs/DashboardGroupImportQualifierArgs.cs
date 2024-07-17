// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardGroupImportQualifierArgs : global::Pulumi.ResourceArgs
    {
        [Input("filters")]
        private InputList<Inputs.DashboardGroupImportQualifierFilterArgs>? _filters;

        /// <summary>
        /// Filter to apply to each chart in the dashboard
        /// </summary>
        public InputList<Inputs.DashboardGroupImportQualifierFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.DashboardGroupImportQualifierFilterArgs>());
            set => _filters = value;
        }

        [Input("metric")]
        public Input<string>? Metric { get; set; }

        public DashboardGroupImportQualifierArgs()
        {
        }
        public static new DashboardGroupImportQualifierArgs Empty => new DashboardGroupImportQualifierArgs();
    }
}