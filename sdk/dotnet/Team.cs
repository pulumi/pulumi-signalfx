// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Signalfx
{
    /// <summary>
    /// Handles management of SignalFx teams.
    /// 
    /// You can configure [team notification policies](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html) using this resource and the various `notifications_*` properties.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/team.html.markdown.
    /// </summary>
    public partial class Team : Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the team.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of user IDs to include in the team.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// Name of the team.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Where to send notifications for critical alerts
        /// </summary>
        [Output("notificationsCriticals")]
        public Output<ImmutableArray<string>> NotificationsCriticals { get; private set; } = null!;

        /// <summary>
        /// Where to send notifications for default alerts
        /// </summary>
        [Output("notificationsDefaults")]
        public Output<ImmutableArray<string>> NotificationsDefaults { get; private set; } = null!;

        /// <summary>
        /// Where to send notifications for info alerts
        /// </summary>
        [Output("notificationsInfos")]
        public Output<ImmutableArray<string>> NotificationsInfos { get; private set; } = null!;

        /// <summary>
        /// Where to send notifications for major alerts
        /// </summary>
        [Output("notificationsMajors")]
        public Output<ImmutableArray<string>> NotificationsMajors { get; private set; } = null!;

        /// <summary>
        /// Where to send notifications for minor alerts
        /// </summary>
        [Output("notificationsMinors")]
        public Output<ImmutableArray<string>> NotificationsMinors { get; private set; } = null!;

        /// <summary>
        /// Where to send notifications for warning alerts
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
            : base("signalfx:index/team:Team", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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

    public sealed class TeamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the team.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// List of user IDs to include in the team.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Name of the team.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationsCriticals")]
        private InputList<string>? _notificationsCriticals;

        /// <summary>
        /// Where to send notifications for critical alerts
        /// </summary>
        public InputList<string> NotificationsCriticals
        {
            get => _notificationsCriticals ?? (_notificationsCriticals = new InputList<string>());
            set => _notificationsCriticals = value;
        }

        [Input("notificationsDefaults")]
        private InputList<string>? _notificationsDefaults;

        /// <summary>
        /// Where to send notifications for default alerts
        /// </summary>
        public InputList<string> NotificationsDefaults
        {
            get => _notificationsDefaults ?? (_notificationsDefaults = new InputList<string>());
            set => _notificationsDefaults = value;
        }

        [Input("notificationsInfos")]
        private InputList<string>? _notificationsInfos;

        /// <summary>
        /// Where to send notifications for info alerts
        /// </summary>
        public InputList<string> NotificationsInfos
        {
            get => _notificationsInfos ?? (_notificationsInfos = new InputList<string>());
            set => _notificationsInfos = value;
        }

        [Input("notificationsMajors")]
        private InputList<string>? _notificationsMajors;

        /// <summary>
        /// Where to send notifications for major alerts
        /// </summary>
        public InputList<string> NotificationsMajors
        {
            get => _notificationsMajors ?? (_notificationsMajors = new InputList<string>());
            set => _notificationsMajors = value;
        }

        [Input("notificationsMinors")]
        private InputList<string>? _notificationsMinors;

        /// <summary>
        /// Where to send notifications for minor alerts
        /// </summary>
        public InputList<string> NotificationsMinors
        {
            get => _notificationsMinors ?? (_notificationsMinors = new InputList<string>());
            set => _notificationsMinors = value;
        }

        [Input("notificationsWarnings")]
        private InputList<string>? _notificationsWarnings;

        /// <summary>
        /// Where to send notifications for warning alerts
        /// </summary>
        public InputList<string> NotificationsWarnings
        {
            get => _notificationsWarnings ?? (_notificationsWarnings = new InputList<string>());
            set => _notificationsWarnings = value;
        }

        public TeamArgs()
        {
        }
    }

    public sealed class TeamState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the team.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// List of user IDs to include in the team.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Name of the team.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationsCriticals")]
        private InputList<string>? _notificationsCriticals;

        /// <summary>
        /// Where to send notifications for critical alerts
        /// </summary>
        public InputList<string> NotificationsCriticals
        {
            get => _notificationsCriticals ?? (_notificationsCriticals = new InputList<string>());
            set => _notificationsCriticals = value;
        }

        [Input("notificationsDefaults")]
        private InputList<string>? _notificationsDefaults;

        /// <summary>
        /// Where to send notifications for default alerts
        /// </summary>
        public InputList<string> NotificationsDefaults
        {
            get => _notificationsDefaults ?? (_notificationsDefaults = new InputList<string>());
            set => _notificationsDefaults = value;
        }

        [Input("notificationsInfos")]
        private InputList<string>? _notificationsInfos;

        /// <summary>
        /// Where to send notifications for info alerts
        /// </summary>
        public InputList<string> NotificationsInfos
        {
            get => _notificationsInfos ?? (_notificationsInfos = new InputList<string>());
            set => _notificationsInfos = value;
        }

        [Input("notificationsMajors")]
        private InputList<string>? _notificationsMajors;

        /// <summary>
        /// Where to send notifications for major alerts
        /// </summary>
        public InputList<string> NotificationsMajors
        {
            get => _notificationsMajors ?? (_notificationsMajors = new InputList<string>());
            set => _notificationsMajors = value;
        }

        [Input("notificationsMinors")]
        private InputList<string>? _notificationsMinors;

        /// <summary>
        /// Where to send notifications for minor alerts
        /// </summary>
        public InputList<string> NotificationsMinors
        {
            get => _notificationsMinors ?? (_notificationsMinors = new InputList<string>());
            set => _notificationsMinors = value;
        }

        [Input("notificationsWarnings")]
        private InputList<string>? _notificationsWarnings;

        /// <summary>
        /// Where to send notifications for warning alerts
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
    }
}
