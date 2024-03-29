// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardGroupDashboardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier of an association between a dashboard group and a dashboard
        /// </summary>
        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        /// <summary>
        /// The label used in the publish statement that displays the plot (metric time series data) you want to customize
        /// </summary>
        [Input("dashboardId", required: true)]
        public Input<string> DashboardId { get; set; } = null!;

        /// <summary>
        /// String that provides a description override for a mirrored dashboard
        /// </summary>
        [Input("descriptionOverride")]
        public Input<string>? DescriptionOverride { get; set; }

        [Input("filterOverrides")]
        private InputList<Inputs.DashboardGroupDashboardFilterOverrideArgs>? _filterOverrides;

        /// <summary>
        /// Filter to apply to each chart in the dashboard
        /// </summary>
        public InputList<Inputs.DashboardGroupDashboardFilterOverrideArgs> FilterOverrides
        {
            get => _filterOverrides ?? (_filterOverrides = new InputList<Inputs.DashboardGroupDashboardFilterOverrideArgs>());
            set => _filterOverrides = value;
        }

        /// <summary>
        /// String that provides a name override for a mirrored dashboard
        /// </summary>
        [Input("nameOverride")]
        public Input<string>? NameOverride { get; set; }

        [Input("variableOverrides")]
        private InputList<Inputs.DashboardGroupDashboardVariableOverrideArgs>? _variableOverrides;

        /// <summary>
        /// Dashboard variable to apply to each chart in the dashboard
        /// </summary>
        public InputList<Inputs.DashboardGroupDashboardVariableOverrideArgs> VariableOverrides
        {
            get => _variableOverrides ?? (_variableOverrides = new InputList<Inputs.DashboardGroupDashboardVariableOverrideArgs>());
            set => _variableOverrides = value;
        }

        public DashboardGroupDashboardArgs()
        {
        }
        public static new DashboardGroupDashboardArgs Empty => new DashboardGroupDashboardArgs();
    }
}
