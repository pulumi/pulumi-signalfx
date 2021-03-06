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
    /// In the SignalFx web UI, a [dashboard group](https://developers.signalfx.com/dashboard_groups_reference.html) is a collection of dashboards.
    /// 
    /// &gt; **NOTE** Dashboard groups cannot be accessed directly, but just via a dashboard contained in them. This is the reason why make show won't show any of yours dashboard groups.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using SignalFx = Pulumi.SignalFx;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var mydashboardgroup0 = new SignalFx.DashboardGroup("mydashboardgroup0", new SignalFx.DashboardGroupArgs
    ///         {
    ///             Description = "Cool dashboard group",
    ///             AuthorizedWriterTeams = 
    ///             {
    ///                 signalfx_team.Mycoolteam.Id,
    ///             },
    ///             AuthorizedWriterUsers = 
    ///             {
    ///                 "abc123",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### With Mirrored Dashboards
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using SignalFx = Pulumi.SignalFx;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var mydashboardgroupWithmirrors = new SignalFx.DashboardGroup("mydashboardgroupWithmirrors", new SignalFx.DashboardGroupArgs
    ///         {
    ///             Description = "Cool dashboard group",
    ///             Dashboards = 
    ///             {
    ///                 new SignalFx.Inputs.DashboardGroupDashboardArgs
    ///                 {
    ///                     DashboardId = signalfx_dashboard.Gc_dashboard.Id,
    ///                     NameOverride = "GC For My Service",
    ///                     DescriptionOverride = "Garbage Collection dashboard maintained by JVM team",
    ///                     FilterOverrides = 
    ///                     {
    ///                         new SignalFx.Inputs.DashboardGroupDashboardFilterOverrideArgs
    ///                         {
    ///                             Property = "service",
    ///                             Values = 
    ///                             {
    ///                                 "myservice",
    ///                             },
    ///                             Negated = false,
    ///                         },
    ///                     },
    ///                     VariableOverrides = 
    ///                     {
    ///                         new SignalFx.Inputs.DashboardGroupDashboardVariableOverrideArgs
    ///                         {
    ///                             Property = "region",
    ///                             Values = 
    ///                             {
    ///                                 "us-west1",
    ///                             },
    ///                             ValuesSuggesteds = 
    ///                             {
    ///                                 "us-west-1",
    ///                                 "us-east-1",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:index/dashboardGroup:DashboardGroup")]
    public partial class DashboardGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        /// </summary>
        [Output("authorizedWriterTeams")]
        public Output<ImmutableArray<string>> AuthorizedWriterTeams { get; private set; } = null!;

        /// <summary>
        /// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        /// </summary>
        [Output("authorizedWriterUsers")]
        public Output<ImmutableArray<string>> AuthorizedWriterUsers { get; private set; } = null!;

        /// <summary>
        /// [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        /// </summary>
        [Output("dashboards")]
        public Output<ImmutableArray<Outputs.DashboardGroupDashboard>> Dashboards { get; private set; } = null!;

        /// <summary>
        /// Description of the dashboard group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("importQualifiers")]
        public Output<ImmutableArray<Outputs.DashboardGroupImportQualifier>> ImportQualifiers { get; private set; } = null!;

        /// <summary>
        /// Name of the dashboard group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Team IDs to associate the dashboard group to.
        /// </summary>
        [Output("teams")]
        public Output<ImmutableArray<string>> Teams { get; private set; } = null!;


        /// <summary>
        /// Create a DashboardGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DashboardGroup(string name, DashboardGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("signalfx:index/dashboardGroup:DashboardGroup", name, args ?? new DashboardGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DashboardGroup(string name, Input<string> id, DashboardGroupState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/dashboardGroup:DashboardGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DashboardGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DashboardGroup Get(string name, Input<string> id, DashboardGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new DashboardGroup(name, id, state, options);
        }
    }

    public sealed class DashboardGroupArgs : Pulumi.ResourceArgs
    {
        [Input("authorizedWriterTeams")]
        private InputList<string>? _authorizedWriterTeams;

        /// <summary>
        /// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        /// </summary>
        public InputList<string> AuthorizedWriterTeams
        {
            get => _authorizedWriterTeams ?? (_authorizedWriterTeams = new InputList<string>());
            set => _authorizedWriterTeams = value;
        }

        [Input("authorizedWriterUsers")]
        private InputList<string>? _authorizedWriterUsers;

        /// <summary>
        /// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        /// </summary>
        public InputList<string> AuthorizedWriterUsers
        {
            get => _authorizedWriterUsers ?? (_authorizedWriterUsers = new InputList<string>());
            set => _authorizedWriterUsers = value;
        }

        [Input("dashboards")]
        private InputList<Inputs.DashboardGroupDashboardArgs>? _dashboards;

        /// <summary>
        /// [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        /// </summary>
        public InputList<Inputs.DashboardGroupDashboardArgs> Dashboards
        {
            get => _dashboards ?? (_dashboards = new InputList<Inputs.DashboardGroupDashboardArgs>());
            set => _dashboards = value;
        }

        /// <summary>
        /// Description of the dashboard group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("importQualifiers")]
        private InputList<Inputs.DashboardGroupImportQualifierArgs>? _importQualifiers;
        public InputList<Inputs.DashboardGroupImportQualifierArgs> ImportQualifiers
        {
            get => _importQualifiers ?? (_importQualifiers = new InputList<Inputs.DashboardGroupImportQualifierArgs>());
            set => _importQualifiers = value;
        }

        /// <summary>
        /// Name of the dashboard group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("teams")]
        private InputList<string>? _teams;

        /// <summary>
        /// Team IDs to associate the dashboard group to.
        /// </summary>
        public InputList<string> Teams
        {
            get => _teams ?? (_teams = new InputList<string>());
            set => _teams = value;
        }

        public DashboardGroupArgs()
        {
        }
    }

    public sealed class DashboardGroupState : Pulumi.ResourceArgs
    {
        [Input("authorizedWriterTeams")]
        private InputList<string>? _authorizedWriterTeams;

        /// <summary>
        /// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        /// </summary>
        public InputList<string> AuthorizedWriterTeams
        {
            get => _authorizedWriterTeams ?? (_authorizedWriterTeams = new InputList<string>());
            set => _authorizedWriterTeams = value;
        }

        [Input("authorizedWriterUsers")]
        private InputList<string>? _authorizedWriterUsers;

        /// <summary>
        /// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        /// </summary>
        public InputList<string> AuthorizedWriterUsers
        {
            get => _authorizedWriterUsers ?? (_authorizedWriterUsers = new InputList<string>());
            set => _authorizedWriterUsers = value;
        }

        [Input("dashboards")]
        private InputList<Inputs.DashboardGroupDashboardGetArgs>? _dashboards;

        /// <summary>
        /// [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        /// </summary>
        public InputList<Inputs.DashboardGroupDashboardGetArgs> Dashboards
        {
            get => _dashboards ?? (_dashboards = new InputList<Inputs.DashboardGroupDashboardGetArgs>());
            set => _dashboards = value;
        }

        /// <summary>
        /// Description of the dashboard group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("importQualifiers")]
        private InputList<Inputs.DashboardGroupImportQualifierGetArgs>? _importQualifiers;
        public InputList<Inputs.DashboardGroupImportQualifierGetArgs> ImportQualifiers
        {
            get => _importQualifiers ?? (_importQualifiers = new InputList<Inputs.DashboardGroupImportQualifierGetArgs>());
            set => _importQualifiers = value;
        }

        /// <summary>
        /// Name of the dashboard group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("teams")]
        private InputList<string>? _teams;

        /// <summary>
        /// Team IDs to associate the dashboard group to.
        /// </summary>
        public InputList<string> Teams
        {
            get => _teams ?? (_teams = new InputList<string>());
            set => _teams = value;
        }

        public DashboardGroupState()
        {
        }
    }
}
