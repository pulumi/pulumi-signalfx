// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class SloTargetAlertRuleRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the SLO.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (default: false) When true, notifications and events will not be generated for the detect label
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("notifications")]
        private InputList<string>? _notifications;

        /// <summary>
        /// List of strings specifying where notifications will be sent when an incident occurs. See https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        /// </summary>
        public InputList<string> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<string>());
            set => _notifications = value;
        }

        /// <summary>
        /// Custom notification message body when an alert is triggered. See https://developers.signalfx.com/v2/reference#detector-model for more info
        /// </summary>
        [Input("parameterizedBody")]
        public Input<string>? ParameterizedBody { get; set; }

        /// <summary>
        /// Custom notification message subject when an alert is triggered. See https://developers.signalfx.com/v2/reference#detector-model for more info
        /// </summary>
        [Input("parameterizedSubject")]
        public Input<string>? ParameterizedSubject { get; set; }

        /// <summary>
        /// Parameters for the SLO alert rule. Each SLO alert rule type accepts different parameters. If not specified, default parameters are used.
        /// </summary>
        [Input("parameters")]
        public Input<Inputs.SloTargetAlertRuleRuleParametersGetArgs>? Parameters { get; set; }

        /// <summary>
        /// URL of page to consult when an alert is triggered
        /// </summary>
        [Input("runbookUrl")]
        public Input<string>? RunbookUrl { get; set; }

        /// <summary>
        /// The severity of the rule, must be one of: Critical, Warning, Major, Minor, Info
        /// </summary>
        [Input("severity", required: true)]
        public Input<string> Severity { get; set; } = null!;

        /// <summary>
        /// Plain text suggested first course of action, such as a command to execute.
        /// </summary>
        [Input("tip")]
        public Input<string>? Tip { get; set; }

        public SloTargetAlertRuleRuleGetArgs()
        {
        }
        public static new SloTargetAlertRuleRuleGetArgs Empty => new SloTargetAlertRuleRuleGetArgs();
    }
}
