// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx
{
    /// <summary>
    /// Provides a Splunk Observability Cloud slo resource. This can be used to create and manage SLOs.
    /// 
    /// To learn more about this feature take a look on [documentation for SLO](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/slo-intro.html).
    /// 
    /// ## Example
    /// 
    /// ## Notification format
    /// 
    /// As Splunk Observability Cloud supports different notification mechanisms, use a comma-delimited string to provide inputs. If you want to specify multiple notifications, each must be a member in the list, like so:
    /// 
    /// See [Splunk Observability Cloud Docs](https://dev.splunk.com/observability/reference/api/detectors/latest) for more information.
    /// 
    /// Here are some example of how to configure each notification type:
    /// 
    /// ### Email
    /// 
    /// ### Jira
    /// 
    /// Note that the `credentialId` is the Splunk-provided ID shown after setting up your Jira integration. See also `signalfx.jira.Integration`.
    /// 
    /// ### OpsGenie
    /// 
    /// Note that the `credentialId` is the Splunk-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.
    /// 
    /// ### PagerDuty
    /// 
    /// ### Slack
    /// 
    /// Exclude the `#` on the channel name:
    /// 
    /// ### Team
    /// 
    /// Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).
    /// 
    /// ### TeamEmail
    /// 
    /// Sends an email to every member of a team.
    /// 
    /// ### Splunk On-Call (formerly VictorOps)
    /// 
    /// ### Webhooks
    /// 
    /// You need to include all the commas even if you only use a credential id.
    /// 
    /// You can either configure a Webhook to use an existing integration's credential id:
    /// 
    /// Or configure one inline:
    /// 
    /// ## Arguments
    /// 
    /// * `name` - (Required) Name of the SLO. Each SLO name must be unique within an organization.
    /// * `description` - (Optional) Description of the SLO.
    /// * `type` - (Required) Type of the SLO. Currently just: `"RequestBased"` is supported.
    /// * `input` - (Required) Properties to configure an SLO object inputs
    ///   * `program_text` - (Required) SignalFlow program and arguments text strings that define the streams used as successful event count and total event count
    ///   * `good_events_label` - (Required) Label used in `"program_text"` that refers to the data block which contains the stream of successful events
    ///   * `total_events_label` - (Required) Label used in `"program_text"` that refers to the data block which contains the stream of total events
    /// * `target` - (Required) Define target value of the service level indicator in the appropriate time period.
    ///   * `type` - (Required) SLO target type can be the following type: `"RollingWindow"`
    ///   * `compliance_period` - (Required for `"RollingWindow"` type) Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
    ///   * `slo` - (Required) Target value in the form of a percentage
    ///   * `alert_rule` - (Required) List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
    ///     * `type` - (Required) SLO alert rule can be one of the following types: BREACH, ERROR_BUDGET_LEFT, BURN_RATE. Within an SLO object, you can only specify one SLO alert_rule per type. For example, you can't specify two alert_rule of type BREACH. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
    ///     * `rule` - (Required) Set of rules used for alerting.
    ///         * `severity` - (Required) The severity of the rule, must be one of: `"Critical"`, `"Major"`, `"Minor"`, `"Warning"`, `"Info"`.
    ///         * `description` - (Optional) Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
    ///         * `disabled` - (Optional) When true, notifications and events will not be generated for the detect label. `false` by default.
    ///         * `notifications` - (Optional) List of strings specifying where notifications will be sent when an incident occurs. See [Create SLO](https://dev.splunk.com/observability/reference/api/slo/latest#endpoint-create-new-slo) for more info.
    ///         * `parameterized_body` - (Optional) Custom notification message body when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
    ///         * `parameterized_subject` - (Optional) Custom notification message subject when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
    ///         * `runbook_url` - (Optional) URL of page to consult when an alert is triggered. This can be used with custom notification messages.
    ///         * `tip` - (Optional) Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
    ///         * `parameters` - (Optional) Parameters for the SLO alert rule. Each SLO alert rule type accepts different parameters. If not specified, default parameters are used.
    ///           * `fire_lasting` - (Optional) Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: `"BREACH"` and `"ERROR_BUDGET_LEFT"` alert rules use the fireLasting parameter. Default: `"5m"`
    ///           * `percent_of_lasting` - (Optional) Percentage of the `"fire_lasting"` duration that the alert condition is met before the alert is triggered. Note: `"BREACH"` and `"ERROR_BUDGET_LEFT"` alert rules use the `"percent_of_lasting"` parameter. Default: `100`
    ///           * `percent_error_budget_left` - (Optional) Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: `"ERROR_BUDGET_LEFT"` alert rules use the `"percent_error_budget_left"` parameter. Default: `100`
    ///           * `short_window_1` - (Optional) Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of `"long_window_1"`. Note: `"BURN_RATE"` alert rules use the `"short_window_1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
    ///           * `short_window_2` - (Optional) Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of `"long_window_2"`. Note: `"BURN_RATE"` alert rules use the `"short_window_2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
    ///           * `long_window_1` - (Optional) Long window 1 used in burn rate alert calculation. This value must be longer than `"short_window_1"` and shorter than 90 days. Note: `"BURN_RATE"` alert rules use the `"long_window_1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
    ///           * `long_window_2` - (Optional) Long window 2 used in burn rate alert calculation. This value must be longer than `"short_window_2"` and shorter than 90 days. Note: `"BURN_RATE"` alert rules use the `"long_window_2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
    ///           * `burn_rate_threshold_1` - (Optional) Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: `"BURN_RATE"` alert rules use the `"burn_rate_threshold_1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
    ///           * `burn_rate_threshold_2` - (Optional) Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: `"BURN_RATE"` alert rules use the `"burn_rate_threshold_2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
    /// </summary>
    [SignalFxResourceType("signalfx:index/slo:Slo")]
    public partial class Slo : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the SLO
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// SignalFlow program and arguments text strings that define the streams used as successful event count and total event
        /// count
        /// </summary>
        [Output("input")]
        public Output<Outputs.SloInput> Input { get; private set; } = null!;

        /// <summary>
        /// Name of the SLO
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Define target value of the service level indicator in the appropriate time period.
        /// </summary>
        [Output("target")]
        public Output<Outputs.SloTarget> Target { get; private set; } = null!;

        /// <summary>
        /// Type of the SLO. Currently only RequestBased SLO is supported
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Slo resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Slo(string name, SloArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/slo:Slo", name, args ?? new SloArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Slo(string name, Input<string> id, SloState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/slo:Slo", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Slo resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Slo Get(string name, Input<string> id, SloState? state = null, CustomResourceOptions? options = null)
        {
            return new Slo(name, id, state, options);
        }
    }

    public sealed class SloArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the SLO
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// SignalFlow program and arguments text strings that define the streams used as successful event count and total event
        /// count
        /// </summary>
        [Input("input", required: true)]
        public Input<Inputs.SloInputArgs> Input { get; set; } = null!;

        /// <summary>
        /// Name of the SLO
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Define target value of the service level indicator in the appropriate time period.
        /// </summary>
        [Input("target", required: true)]
        public Input<Inputs.SloTargetArgs> Target { get; set; } = null!;

        /// <summary>
        /// Type of the SLO. Currently only RequestBased SLO is supported
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SloArgs()
        {
        }
        public static new SloArgs Empty => new SloArgs();
    }

    public sealed class SloState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the SLO
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// SignalFlow program and arguments text strings that define the streams used as successful event count and total event
        /// count
        /// </summary>
        [Input("input")]
        public Input<Inputs.SloInputGetArgs>? Input { get; set; }

        /// <summary>
        /// Name of the SLO
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Define target value of the service level indicator in the appropriate time period.
        /// </summary>
        [Input("target")]
        public Input<Inputs.SloTargetGetArgs>? Target { get; set; }

        /// <summary>
        /// Type of the SLO. Currently only RequestBased SLO is supported
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SloState()
        {
        }
        public static new SloState Empty => new SloState();
    }
}
