// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardPermissionsArgs : Pulumi.ResourceArgs
    {
        [Input("acls")]
        private InputList<Inputs.DashboardPermissionsAclArgs>? _acls;

        /// <summary>
        /// List of read and write permission configurations to specify which user, team, and organization can view and/or edit your dashboard. Use the `permissions.parent` instead if you want to inherit permissions.
        /// </summary>
        public InputList<Inputs.DashboardPermissionsAclArgs> Acls
        {
            get => _acls ?? (_acls = new InputList<Inputs.DashboardPermissionsAclArgs>());
            set => _acls = value;
        }

        /// <summary>
        /// ID of the dashboard group you want your dashboard to inherit permissions from. Use the `permissions.acl` instead if you want to specify various read and write permission configurations.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        public DashboardPermissionsArgs()
        {
        }
    }
}
