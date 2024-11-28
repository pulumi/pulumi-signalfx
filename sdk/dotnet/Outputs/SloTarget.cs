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
    public sealed class SloTarget
    {
        /// <summary>
        /// List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
        /// </summary>
        public readonly ImmutableArray<Outputs.SloTargetAlertRule> AlertRules;
        /// <summary>
        /// Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
        /// </summary>
        public readonly string? CompliancePeriod;
        /// <summary>
        /// It can be used to change the cycle start time. For example, you can specify sunday as the start of the week (instead of the default monday)
        /// </summary>
        public readonly string? CycleStart;
        /// <summary>
        /// The cycle type of the calendar window, e.g. week, month.
        /// </summary>
        public readonly string? CycleType;
        /// <summary>
        /// Target value in the form of a percentage
        /// </summary>
        public readonly double Slo;
        /// <summary>
        /// SLO target type can be the following type: `"RollingWindow"`, `"CalendarWindow"`
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SloTarget(
            ImmutableArray<Outputs.SloTargetAlertRule> alertRules,

            string? compliancePeriod,

            string? cycleStart,

            string? cycleType,

            double slo,

            string type)
        {
            AlertRules = alertRules;
            CompliancePeriod = compliancePeriod;
            CycleStart = cycleStart;
            CycleType = cycleType;
            Slo = slo;
            Type = type;
        }
    }
}