// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// The ID of the association between the dashboard group and the dashboard
        /// </summary>
        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        /// <summary>
        /// The dashboard id to mirror
        /// </summary>
        [Input("dashboardId", required: true)]
        public Input<string> DashboardId { get; set; } = null!;

        /// <summary>
        /// The description that will override the original dashboards's description.
        /// </summary>
        [Input("descriptionOverride")]
        public Input<string>? DescriptionOverride { get; set; }

        [Input("filterOverrides")]
        private InputList<Inputs.DashboardGroupDashboardFilterOverrideArgs>? _filterOverrides;

        /// <summary>
        /// The description that will override the original dashboards's description.
        /// </summary>
        public InputList<Inputs.DashboardGroupDashboardFilterOverrideArgs> FilterOverrides
        {
            get => _filterOverrides ?? (_filterOverrides = new InputList<Inputs.DashboardGroupDashboardFilterOverrideArgs>());
            set => _filterOverrides = value;
        }

        /// <summary>
        /// The name that will override the original dashboards's name.
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
