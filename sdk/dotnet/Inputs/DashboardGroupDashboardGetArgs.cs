// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardGroupDashboardGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        [Input("dashboardId", required: true)]
        public Input<string> DashboardId { get; set; } = null!;

        [Input("descriptionOverride")]
        public Input<string>? DescriptionOverride { get; set; }

        [Input("filterOverrides")]
        private InputList<Inputs.DashboardGroupDashboardFilterOverrideGetArgs>? _filterOverrides;
        public InputList<Inputs.DashboardGroupDashboardFilterOverrideGetArgs> FilterOverrides
        {
            get => _filterOverrides ?? (_filterOverrides = new InputList<Inputs.DashboardGroupDashboardFilterOverrideGetArgs>());
            set => _filterOverrides = value;
        }

        [Input("nameOverride")]
        public Input<string>? NameOverride { get; set; }

        [Input("variableOverrides")]
        private InputList<Inputs.DashboardGroupDashboardVariableOverrideGetArgs>? _variableOverrides;
        public InputList<Inputs.DashboardGroupDashboardVariableOverrideGetArgs> VariableOverrides
        {
            get => _variableOverrides ?? (_variableOverrides = new InputList<Inputs.DashboardGroupDashboardVariableOverrideGetArgs>());
            set => _variableOverrides = value;
        }

        public DashboardGroupDashboardGetArgs()
        {
        }
        public static new DashboardGroupDashboardGetArgs Empty => new DashboardGroupDashboardGetArgs();
    }
}
