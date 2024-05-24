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
    public sealed class OrgTokenHostOrUsageLimits
    {
        /// <summary>
        /// Max number of Docker containers that can use this token
        /// </summary>
        public readonly int? ContainerLimit;
        /// <summary>
        /// Notification threshold for Docker containers
        /// </summary>
        public readonly int? ContainerNotificationThreshold;
        /// <summary>
        /// Max number of custom metrics that can be sent with this token
        /// </summary>
        public readonly int? CustomMetricsLimit;
        /// <summary>
        /// Notification threshold for custom metrics
        /// </summary>
        public readonly int? CustomMetricsNotificationThreshold;
        /// <summary>
        /// Max number of hi-res metrics that can be sent with this toke
        /// </summary>
        public readonly int? HighResMetricsLimit;
        /// <summary>
        /// Notification threshold for hi-res metrics
        /// </summary>
        public readonly int? HighResMetricsNotificationThreshold;
        /// <summary>
        /// Max number of hosts that can use this token
        /// </summary>
        public readonly int? HostLimit;
        /// <summary>
        /// Notification threshold for hosts
        /// </summary>
        public readonly int? HostNotificationThreshold;

        [OutputConstructor]
        private OrgTokenHostOrUsageLimits(
            int? containerLimit,

            int? containerNotificationThreshold,

            int? customMetricsLimit,

            int? customMetricsNotificationThreshold,

            int? highResMetricsLimit,

            int? highResMetricsNotificationThreshold,

            int? hostLimit,

            int? hostNotificationThreshold)
        {
            ContainerLimit = containerLimit;
            ContainerNotificationThreshold = containerNotificationThreshold;
            CustomMetricsLimit = customMetricsLimit;
            CustomMetricsNotificationThreshold = customMetricsNotificationThreshold;
            HighResMetricsLimit = highResMetricsLimit;
            HighResMetricsNotificationThreshold = highResMetricsNotificationThreshold;
            HostLimit = hostLimit;
            HostNotificationThreshold = hostNotificationThreshold;
        }
    }
}
