// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class SloTargetGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("alertRules", required: true)]
        private InputList<Inputs.SloTargetAlertRuleGetArgs>? _alertRules;

        /// <summary>
        /// List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
        /// </summary>
        public InputList<Inputs.SloTargetAlertRuleGetArgs> AlertRules
        {
            get => _alertRules ?? (_alertRules = new InputList<Inputs.SloTargetAlertRuleGetArgs>());
            set => _alertRules = value;
        }

        /// <summary>
        /// Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
        /// </summary>
        [Input("compliancePeriod")]
        public Input<string>? CompliancePeriod { get; set; }

        /// <summary>
        /// Target value in the form of a percentage
        /// </summary>
        [Input("slo", required: true)]
        public Input<double> Slo { get; set; } = null!;

        /// <summary>
        /// SLO alert rule can be one of the following types: BREACH, ERROR_BUDGET_LEFT, BURN_RATE. Within an SLO object, you can only specify one SLO alert_rule per type. For example, you can't specify two alert_rule of type BREACH. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SloTargetGetArgs()
        {
        }
        public static new SloTargetGetArgs Empty => new SloTargetGetArgs();
    }
}
