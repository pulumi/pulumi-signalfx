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
    public sealed class DataLinkTargetSignalfxDashboard
    {
        /// <summary>
        /// SignalFx-assigned ID of the dashboard link target's dashboard group
        /// </summary>
        public readonly string DashboardGroupId;
        /// <summary>
        /// SignalFx-assigned ID of the dashboard link target
        /// </summary>
        public readonly string DashboardId;
        /// <summary>
        /// Flag that designates a target as the default for a data link object.
        /// </summary>
        public readonly bool? IsDefault;
        /// <summary>
        /// User-assigned target name. Use this value to differentiate between the link targets for a data link object.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private DataLinkTargetSignalfxDashboard(
            string dashboardGroupId,

            string dashboardId,

            bool? isDefault,

            string name)
        {
            DashboardGroupId = dashboardGroupId;
            DashboardId = dashboardId;
            IsDefault = isDefault;
            Name = name;
        }
    }
}
