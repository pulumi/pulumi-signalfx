// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DetectorRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A detect label which matches a detect label within `program_text`.
        /// </summary>
        [Input("detectLabel", required: true)]
        public Input<string> DetectLabel { get; set; } = null!;

        /// <summary>
        /// When true, notifications and events will not be generated for the detect label. `false` by default.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("notifications")]
        private InputList<string>? _notifications;

        /// <summary>
        /// List of strings specifying where notifications will be sent when an incident occurs. See [Create A Single Detector](https://developers.signalfx.com/detectors_reference.html#operation/Create%20Single%20Detector) for more info.
        /// </summary>
        public InputList<string> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<string>());
            set => _notifications = value;
        }

        /// <summary>
        /// Custom notification message body when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.signalfx.com/en/latest/detect-alert/set-up-detectors.html#about-detectors#alert-settings) for more info.
        /// </summary>
        [Input("parameterizedBody")]
        public Input<string>? ParameterizedBody { get; set; }

        /// <summary>
        /// Custom notification message subject when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.signalfx.com/en/latest/detect-alert/set-up-detectors.html#about-detectors#alert-settings) for more info.
        /// </summary>
        [Input("parameterizedSubject")]
        public Input<string>? ParameterizedSubject { get; set; }

        /// <summary>
        /// URL of page to consult when an alert is triggered. This can be used with custom notification messages.
        /// </summary>
        [Input("runbookUrl")]
        public Input<string>? RunbookUrl { get; set; }

        /// <summary>
        /// The severity of the rule, must be one of: `"Critical"`, `"Major"`, `"Minor"`, `"Warning"`, `"Info"`.
        /// </summary>
        [Input("severity", required: true)]
        public Input<string> Severity { get; set; } = null!;

        /// <summary>
        /// Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
        /// </summary>
        [Input("tip")]
        public Input<string>? Tip { get; set; }

        public DetectorRuleArgs()
        {
        }
    }
}