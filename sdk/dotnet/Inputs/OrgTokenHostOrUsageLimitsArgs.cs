// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class OrgTokenHostOrUsageLimitsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Max number of Docker containers that can use this token
        /// </summary>
        [Input("containerLimit")]
        public Input<int>? ContainerLimit { get; set; }

        /// <summary>
        /// Notification threshold for Docker containers
        /// </summary>
        [Input("containerNotificationThreshold")]
        public Input<int>? ContainerNotificationThreshold { get; set; }

        /// <summary>
        /// Max number of custom metrics that can be sent with this token
        /// </summary>
        [Input("customMetricsLimit")]
        public Input<int>? CustomMetricsLimit { get; set; }

        /// <summary>
        /// Notification threshold for custom metrics
        /// </summary>
        [Input("customMetricsNotificationThreshold")]
        public Input<int>? CustomMetricsNotificationThreshold { get; set; }

        /// <summary>
        /// Max number of hi-res metrics that can be sent with this toke
        /// </summary>
        [Input("highResMetricsLimit")]
        public Input<int>? HighResMetricsLimit { get; set; }

        /// <summary>
        /// Notification threshold for hi-res metrics
        /// </summary>
        [Input("highResMetricsNotificationThreshold")]
        public Input<int>? HighResMetricsNotificationThreshold { get; set; }

        /// <summary>
        /// Max number of hosts that can use this token
        /// </summary>
        [Input("hostLimit")]
        public Input<int>? HostLimit { get; set; }

        /// <summary>
        /// Notification threshold for hosts
        /// </summary>
        [Input("hostNotificationThreshold")]
        public Input<int>? HostNotificationThreshold { get; set; }

        public OrgTokenHostOrUsageLimitsArgs()
        {
        }
        public static new OrgTokenHostOrUsageLimitsArgs Empty => new OrgTokenHostOrUsageLimitsArgs();
    }
}
