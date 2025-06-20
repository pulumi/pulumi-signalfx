// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Outputs
{

    [OutputType]
    public sealed class SloTargetAlertRule
    {
        /// <summary>
        /// Set of rules used for alerting.
        /// </summary>
        public readonly ImmutableArray<Outputs.SloTargetAlertRuleRule> Rules;
        /// <summary>
        /// SLO alert rule can be one of the following types: BREACH, ERROR_BUDGET_LEFT, BURN_RATE. Within an SLO object, you can only specify one SLO alert_rule per type. For example, you can't specify two alert_rule of type BREACH. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private SloTargetAlertRule(
            ImmutableArray<Outputs.SloTargetAlertRuleRule> rules,

            string type)
        {
            Rules = rules;
            Type = type;
        }
    }
}
