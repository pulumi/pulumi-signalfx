// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class SloTargetAlertRuleRuleParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burn_rate_threshold_1 parameter.
        /// </summary>
        [Input("burnRateThreshold1")]
        public Input<double>? BurnRateThreshold1 { get; set; }

        /// <summary>
        /// Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burn_rate_threshold_2 parameter.
        /// </summary>
        [Input("burnRateThreshold2")]
        public Input<double>? BurnRateThreshold2 { get; set; }

        /// <summary>
        /// Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the fire_lasting parameter
        /// </summary>
        [Input("fireLasting")]
        public Input<string>? FireLasting { get; set; }

        /// <summary>
        /// Long window 1 used in burn rate alert calculation. This value must be longer than short_window_1` and shorter than 90 days. Note: BURN_RATE alert rules use the long_window_1 parameter.
        /// </summary>
        [Input("longWindow1")]
        public Input<string>? LongWindow1 { get; set; }

        /// <summary>
        /// Long window 2 used in burn rate alert calculation. This value must be longer than short_window_2` and shorter than 90 days. Note: BURN_RATE alert rules use the long_window_2 parameter.
        /// </summary>
        [Input("longWindow2")]
        public Input<string>? LongWindow2 { get; set; }

        /// <summary>
        /// Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: ERROR_BUDGET_LEFT alert rules use the percent_error_budget_left parameter.
        /// </summary>
        [Input("percentErrorBudgetLeft")]
        public Input<double>? PercentErrorBudgetLeft { get; set; }

        /// <summary>
        /// Percentage of the fire_lasting duration that the alert condition is met before the alert is triggered. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the percent_of_lasting parameter
        /// </summary>
        [Input("percentOfLasting")]
        public Input<double>? PercentOfLasting { get; set; }

        /// <summary>
        /// Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_1. Note: BURN_RATE alert rules use the short_window_1 parameter.
        /// </summary>
        [Input("shortWindow1")]
        public Input<string>? ShortWindow1 { get; set; }

        /// <summary>
        /// Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_2. Note: BURN_RATE alert rules use the short_window_2 parameter.
        /// </summary>
        [Input("shortWindow2")]
        public Input<string>? ShortWindow2 { get; set; }

        public SloTargetAlertRuleRuleParametersArgs()
        {
        }
        public static new SloTargetAlertRuleRuleParametersArgs Empty => new SloTargetAlertRuleRuleParametersArgs();
    }
}
