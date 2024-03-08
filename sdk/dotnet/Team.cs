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
    /// Handles management of Splunk Observability Cloud teams.
    /// 
    /// You can configure [team notification policies](https://docs.splunk.com/observability/en/admin/user-management/teams/team-notifications.html) using this resource and the various `notifications_*` properties.
    /// 
    /// &gt; **NOTE** When managing teams, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
    /// 
    /// ## Example
    /// 
    /// ## Arguments
    /// 
    /// The following arguments are supported in the resource block:
    /// 
    /// * `name` - (Required) Name of the team.
    /// * `description` - (Optional) Description of the team.
    /// * `members` - (Optional) List of user IDs to include in the team.
    /// * `notifications_critical` - (Optional) Where to send notifications for critical alerts
    /// * `notifications_default` - (Optional) Where to send notifications for default alerts
    /// * `notifications_info` - (Optional) Where to send notifications for info alerts
    /// * `notifications_major` - (Optional) Where to send notifications for major alerts
    /// * `notifications_minor` - (Optional) Where to send notifications for minor alerts
    /// * `notifications_warning` - (Optional) Where to send notifications for warning alerts
    /// 
    /// ## Attributes
    /// 
    /// In a addition to all arguments above, the following attributes are exported:
    /// 
    /// * `id` - The ID of the team.
    /// * `url` - The URL of the team.
    /// </summary>
    [SignalFxResourceType("signalfx:index/team:Team")]
    public partial class Team : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the team (Optional)
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Members of team
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// Name of the team
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of notification destinations to use for the critical alerts category.
        /// </summary>
        [Output("notificationsCriticals")]
        public Output<ImmutableArray<string>> NotificationsCriticals { get; private set; } = null!;

        /// <summary>
        /// List of notification destinations to use for the default alerts category.
        /// </summary>
        [Output("notificationsDefaults")]
        public Output<ImmutableArray<string>> NotificationsDefaults { get; private set; } = null!;

        /// <summary>
        /// List of notification destinations to use for the info alerts category.
        /// </summary>
        [Output("notificationsInfos")]
        public Output<ImmutableArray<string>> NotificationsInfos { get; private set; } = null!;

        /// <summary>
        /// List of notification destinations to use for the major alerts category.
        /// </summary>
        [Output("notificationsMajors")]
        public Output<ImmutableArray<string>> NotificationsMajors { get; private set; } = null!;

        /// <summary>
        /// List of notification destinations to use for the minor alerts category.
        /// </summary>
        [Output("notificationsMinors")]
        public Output<ImmutableArray<string>> NotificationsMinors { get; private set; } = null!;

        /// <summary>
        /// List of notification destinations to use for the warning alerts category.
        /// </summary>
        [Output("notificationsWarnings")]
        public Output<ImmutableArray<string>> NotificationsWarnings { get; private set; } = null!;

        /// <summary>
        /// URL of the team
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Team resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Team(string name, TeamArgs? args = null, CustomResourceOptions? options = null)
            : base("signalfx:index/team:Team", name, args ?? new TeamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Team(string name, Input<string> id, TeamState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/team:Team", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Team resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Team Get(string name, Input<string> id, TeamState? state = null, CustomResourceOptions? options = null)
        {
            return new Team(name, id, state, options);
        }
    }

    public sealed class TeamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the team (Optional)
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// Members of team
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Name of the team
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationsCriticals")]
        private InputList<string>? _notificationsCriticals;

        /// <summary>
        /// List of notification destinations to use for the critical alerts category.
        /// </summary>
        public InputList<string> NotificationsCriticals
        {
            get => _notificationsCriticals ?? (_notificationsCriticals = new InputList<string>());
            set => _notificationsCriticals = value;
        }

        [Input("notificationsDefaults")]
        private InputList<string>? _notificationsDefaults;

        /// <summary>
        /// List of notification destinations to use for the default alerts category.
        /// </summary>
        public InputList<string> NotificationsDefaults
        {
            get => _notificationsDefaults ?? (_notificationsDefaults = new InputList<string>());
            set => _notificationsDefaults = value;
        }

        [Input("notificationsInfos")]
        private InputList<string>? _notificationsInfos;

        /// <summary>
        /// List of notification destinations to use for the info alerts category.
        /// </summary>
        public InputList<string> NotificationsInfos
        {
            get => _notificationsInfos ?? (_notificationsInfos = new InputList<string>());
            set => _notificationsInfos = value;
        }

        [Input("notificationsMajors")]
        private InputList<string>? _notificationsMajors;

        /// <summary>
        /// List of notification destinations to use for the major alerts category.
        /// </summary>
        public InputList<string> NotificationsMajors
        {
            get => _notificationsMajors ?? (_notificationsMajors = new InputList<string>());
            set => _notificationsMajors = value;
        }

        [Input("notificationsMinors")]
        private InputList<string>? _notificationsMinors;

        /// <summary>
        /// List of notification destinations to use for the minor alerts category.
        /// </summary>
        public InputList<string> NotificationsMinors
        {
            get => _notificationsMinors ?? (_notificationsMinors = new InputList<string>());
            set => _notificationsMinors = value;
        }

        [Input("notificationsWarnings")]
        private InputList<string>? _notificationsWarnings;

        /// <summary>
        /// List of notification destinations to use for the warning alerts category.
        /// </summary>
        public InputList<string> NotificationsWarnings
        {
            get => _notificationsWarnings ?? (_notificationsWarnings = new InputList<string>());
            set => _notificationsWarnings = value;
        }

        public TeamArgs()
        {
        }
        public static new TeamArgs Empty => new TeamArgs();
    }

    public sealed class TeamState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the team (Optional)
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// Members of team
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Name of the team
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationsCriticals")]
        private InputList<string>? _notificationsCriticals;

        /// <summary>
        /// List of notification destinations to use for the critical alerts category.
        /// </summary>
        public InputList<string> NotificationsCriticals
        {
            get => _notificationsCriticals ?? (_notificationsCriticals = new InputList<string>());
            set => _notificationsCriticals = value;
        }

        [Input("notificationsDefaults")]
        private InputList<string>? _notificationsDefaults;

        /// <summary>
        /// List of notification destinations to use for the default alerts category.
        /// </summary>
        public InputList<string> NotificationsDefaults
        {
            get => _notificationsDefaults ?? (_notificationsDefaults = new InputList<string>());
            set => _notificationsDefaults = value;
        }

        [Input("notificationsInfos")]
        private InputList<string>? _notificationsInfos;

        /// <summary>
        /// List of notification destinations to use for the info alerts category.
        /// </summary>
        public InputList<string> NotificationsInfos
        {
            get => _notificationsInfos ?? (_notificationsInfos = new InputList<string>());
            set => _notificationsInfos = value;
        }

        [Input("notificationsMajors")]
        private InputList<string>? _notificationsMajors;

        /// <summary>
        /// List of notification destinations to use for the major alerts category.
        /// </summary>
        public InputList<string> NotificationsMajors
        {
            get => _notificationsMajors ?? (_notificationsMajors = new InputList<string>());
            set => _notificationsMajors = value;
        }

        [Input("notificationsMinors")]
        private InputList<string>? _notificationsMinors;

        /// <summary>
        /// List of notification destinations to use for the minor alerts category.
        /// </summary>
        public InputList<string> NotificationsMinors
        {
            get => _notificationsMinors ?? (_notificationsMinors = new InputList<string>());
            set => _notificationsMinors = value;
        }

        [Input("notificationsWarnings")]
        private InputList<string>? _notificationsWarnings;

        /// <summary>
        /// List of notification destinations to use for the warning alerts category.
        /// </summary>
        public InputList<string> NotificationsWarnings
        {
            get => _notificationsWarnings ?? (_notificationsWarnings = new InputList<string>());
            set => _notificationsWarnings = value;
        }

        /// <summary>
        /// URL of the team
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public TeamState()
        {
        }
        public static new TeamState Empty => new TeamState();
    }
}
