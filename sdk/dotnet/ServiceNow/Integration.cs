// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.ServiceNow
{
    /// <summary>
    /// Observability Cloud ServiceNow integrations. For help with this integration see [Integration with ServiceNow](https://docs.splunk.com/Observability/admin/notif-services/servicenow.html).
    /// 
    /// &gt; **NOTE** When managing integrations, use a session token of an administrator to authenticate the Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
    /// </summary>
    [SignalFxResourceType("signalfx:servicenow/integration:Integration")]
    public partial class Integration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
        /// </summary>
        [Output("alertResolvedPayloadTemplate")]
        public Output<string?> AlertResolvedPayloadTemplate { get; private set; } = null!;

        /// <summary>
        /// A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
        /// </summary>
        [Output("alertTriggeredPayloadTemplate")]
        public Output<string?> AlertTriggeredPayloadTemplate { get; private set; } = null!;

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Name of the ServiceNow instance, for example `myinst.service-now.com`.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        /// </summary>
        [Output("issueType")]
        public Output<string> IssueType { get; private set; } = null!;

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password used to authenticate the ServiceNow integration.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// User name used to authenticate the ServiceNow integration.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("signalfx:servicenow/integration:Integration", name, args ?? new IntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Integration(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:servicenow/integration:Integration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Integration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Integration Get(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new Integration(name, id, state, options);
        }
    }

    public sealed class IntegrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
        /// </summary>
        [Input("alertResolvedPayloadTemplate")]
        public Input<string>? AlertResolvedPayloadTemplate { get; set; }

        /// <summary>
        /// A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
        /// </summary>
        [Input("alertTriggeredPayloadTemplate")]
        public Input<string>? AlertTriggeredPayloadTemplate { get; set; }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Name of the ServiceNow instance, for example `myinst.service-now.com`.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        /// </summary>
        [Input("issueType", required: true)]
        public Input<string> IssueType { get; set; } = null!;

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Password used to authenticate the ServiceNow integration.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// User name used to authenticate the ServiceNow integration.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public IntegrationArgs()
        {
        }
        public static new IntegrationArgs Empty => new IntegrationArgs();
    }

    public sealed class IntegrationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
        /// </summary>
        [Input("alertResolvedPayloadTemplate")]
        public Input<string>? AlertResolvedPayloadTemplate { get; set; }

        /// <summary>
        /// A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
        /// </summary>
        [Input("alertTriggeredPayloadTemplate")]
        public Input<string>? AlertTriggeredPayloadTemplate { get; set; }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Name of the ServiceNow instance, for example `myinst.service-now.com`.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        /// </summary>
        [Input("issueType")]
        public Input<string>? IssueType { get; set; }

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password used to authenticate the ServiceNow integration.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// User name used to authenticate the ServiceNow integration.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public IntegrationState()
        {
        }
        public static new IntegrationState Empty => new IntegrationState();
    }
}
