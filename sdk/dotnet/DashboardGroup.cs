// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx
{
    /// <summary>
    /// In the SignalFx web UI, a [dashboard group](https://developers.signalfx.com/v2/docs/dashboard-group-model) is a collection of dashboards.
    /// 
    /// **NOTE:** Dashboard groups cannot be accessed directly, but just via a dashboard contained in them. This is the reason why make show won't show any of yours dashboard groups.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/dashboard_group.html.markdown.
    /// </summary>
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
        public Output<ImmutableArray<Outputs.DashboardGroupDashboards>> Dashboards { get; private set; } = null!;

        /// <summary>
        /// Description of the dashboard group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

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
            : base("signalfx:index/dashboardGroup:DashboardGroup", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
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
        private InputList<Inputs.DashboardGroupDashboardsArgs>? _dashboards;

        /// <summary>
        /// [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        /// </summary>
        public InputList<Inputs.DashboardGroupDashboardsArgs> Dashboards
        {
            get => _dashboards ?? (_dashboards = new InputList<Inputs.DashboardGroupDashboardsArgs>());
            set => _dashboards = value;
        }

        /// <summary>
        /// Description of the dashboard group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        private InputList<Inputs.DashboardGroupDashboardsGetArgs>? _dashboards;

        /// <summary>
        /// [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        /// </summary>
        public InputList<Inputs.DashboardGroupDashboardsGetArgs> Dashboards
        {
            get => _dashboards ?? (_dashboards = new InputList<Inputs.DashboardGroupDashboardsGetArgs>());
            set => _dashboards = value;
        }

        /// <summary>
        /// Description of the dashboard group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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

    namespace Inputs
    {

    public sealed class DashboardGroupDashboardsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dashboard id to mirror
        /// </summary>
        [Input("dashboardId", required: true)]
        public Input<string> DashboardId { get; set; } = null!;

        /// <summary>
        /// The description that will override the original dashboards's description.
        /// </summary>
        [Input("descriptionOverride")]
        public Input<string>? DescriptionOverride { get; set; }

        [Input("filterOverrides")]
        private InputList<DashboardGroupDashboardsFilterOverridesArgs>? _filterOverrides;

        /// <summary>
        /// The description that will override the original dashboards's description.
        /// </summary>
        public InputList<DashboardGroupDashboardsFilterOverridesArgs> FilterOverrides
        {
            get => _filterOverrides ?? (_filterOverrides = new InputList<DashboardGroupDashboardsFilterOverridesArgs>());
            set => _filterOverrides = value;
        }

        /// <summary>
        /// The name that will override the original dashboards's name.
        /// </summary>
        [Input("nameOverride")]
        public Input<string>? NameOverride { get; set; }

        [Input("variableOverrides")]
        private InputList<DashboardGroupDashboardsVariableOverridesArgs>? _variableOverrides;
        public InputList<DashboardGroupDashboardsVariableOverridesArgs> VariableOverrides
        {
            get => _variableOverrides ?? (_variableOverrides = new InputList<DashboardGroupDashboardsVariableOverridesArgs>());
            set => _variableOverrides = value;
        }

        public DashboardGroupDashboardsArgs()
        {
        }
    }

    public sealed class DashboardGroupDashboardsFilterOverridesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
        /// </summary>
        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        /// <summary>
        /// A metric time series dimension or property name.
        /// </summary>
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// (Optional) List of of strings (which will be treated as an OR filter on the property).
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public DashboardGroupDashboardsFilterOverridesArgs()
        {
        }
    }

    public sealed class DashboardGroupDashboardsFilterOverridesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
        /// </summary>
        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        /// <summary>
        /// A metric time series dimension or property name.
        /// </summary>
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// (Optional) List of of strings (which will be treated as an OR filter on the property).
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public DashboardGroupDashboardsFilterOverridesGetArgs()
        {
        }
    }

    public sealed class DashboardGroupDashboardsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dashboard id to mirror
        /// </summary>
        [Input("dashboardId", required: true)]
        public Input<string> DashboardId { get; set; } = null!;

        /// <summary>
        /// The description that will override the original dashboards's description.
        /// </summary>
        [Input("descriptionOverride")]
        public Input<string>? DescriptionOverride { get; set; }

        [Input("filterOverrides")]
        private InputList<DashboardGroupDashboardsFilterOverridesGetArgs>? _filterOverrides;

        /// <summary>
        /// The description that will override the original dashboards's description.
        /// </summary>
        public InputList<DashboardGroupDashboardsFilterOverridesGetArgs> FilterOverrides
        {
            get => _filterOverrides ?? (_filterOverrides = new InputList<DashboardGroupDashboardsFilterOverridesGetArgs>());
            set => _filterOverrides = value;
        }

        /// <summary>
        /// The name that will override the original dashboards's name.
        /// </summary>
        [Input("nameOverride")]
        public Input<string>? NameOverride { get; set; }

        [Input("variableOverrides")]
        private InputList<DashboardGroupDashboardsVariableOverridesGetArgs>? _variableOverrides;
        public InputList<DashboardGroupDashboardsVariableOverridesGetArgs> VariableOverrides
        {
            get => _variableOverrides ?? (_variableOverrides = new InputList<DashboardGroupDashboardsVariableOverridesGetArgs>());
            set => _variableOverrides = value;
        }

        public DashboardGroupDashboardsGetArgs()
        {
        }
    }

    public sealed class DashboardGroupDashboardsVariableOverridesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A metric time series dimension or property name.
        /// </summary>
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// (Optional) List of of strings (which will be treated as an OR filter on the property).
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        [Input("valuesSuggesteds")]
        private InputList<string>? _valuesSuggesteds;

        /// <summary>
        /// A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
        /// </summary>
        public InputList<string> ValuesSuggesteds
        {
            get => _valuesSuggesteds ?? (_valuesSuggesteds = new InputList<string>());
            set => _valuesSuggesteds = value;
        }

        public DashboardGroupDashboardsVariableOverridesArgs()
        {
        }
    }

    public sealed class DashboardGroupDashboardsVariableOverridesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A metric time series dimension or property name.
        /// </summary>
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// (Optional) List of of strings (which will be treated as an OR filter on the property).
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        [Input("valuesSuggesteds")]
        private InputList<string>? _valuesSuggesteds;

        /// <summary>
        /// A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
        /// </summary>
        public InputList<string> ValuesSuggesteds
        {
            get => _valuesSuggesteds ?? (_valuesSuggesteds = new InputList<string>());
            set => _valuesSuggesteds = value;
        }

        public DashboardGroupDashboardsVariableOverridesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DashboardGroupDashboards
    {
        /// <summary>
        /// The dashboard id to mirror
        /// </summary>
        public readonly string DashboardId;
        /// <summary>
        /// The description that will override the original dashboards's description.
        /// </summary>
        public readonly string? DescriptionOverride;
        /// <summary>
        /// The description that will override the original dashboards's description.
        /// </summary>
        public readonly ImmutableArray<DashboardGroupDashboardsFilterOverrides> FilterOverrides;
        /// <summary>
        /// The name that will override the original dashboards's name.
        /// </summary>
        public readonly string? NameOverride;
        public readonly ImmutableArray<DashboardGroupDashboardsVariableOverrides> VariableOverrides;

        [OutputConstructor]
        private DashboardGroupDashboards(
            string dashboardId,
            string? descriptionOverride,
            ImmutableArray<DashboardGroupDashboardsFilterOverrides> filterOverrides,
            string? nameOverride,
            ImmutableArray<DashboardGroupDashboardsVariableOverrides> variableOverrides)
        {
            DashboardId = dashboardId;
            DescriptionOverride = descriptionOverride;
            FilterOverrides = filterOverrides;
            NameOverride = nameOverride;
            VariableOverrides = variableOverrides;
        }
    }

    [OutputType]
    public sealed class DashboardGroupDashboardsFilterOverrides
    {
        /// <summary>
        /// If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
        /// </summary>
        public readonly bool? Negated;
        /// <summary>
        /// A metric time series dimension or property name.
        /// </summary>
        public readonly string Property;
        /// <summary>
        /// (Optional) List of of strings (which will be treated as an OR filter on the property).
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private DashboardGroupDashboardsFilterOverrides(
            bool? negated,
            string property,
            ImmutableArray<string> values)
        {
            Negated = negated;
            Property = property;
            Values = values;
        }
    }

    [OutputType]
    public sealed class DashboardGroupDashboardsVariableOverrides
    {
        /// <summary>
        /// A metric time series dimension or property name.
        /// </summary>
        public readonly string Property;
        /// <summary>
        /// (Optional) List of of strings (which will be treated as an OR filter on the property).
        /// </summary>
        public readonly ImmutableArray<string> Values;
        /// <summary>
        /// A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
        /// </summary>
        public readonly ImmutableArray<string> ValuesSuggesteds;

        [OutputConstructor]
        private DashboardGroupDashboardsVariableOverrides(
            string property,
            ImmutableArray<string> values,
            ImmutableArray<string> valuesSuggesteds)
        {
            Property = property;
            Values = values;
            ValuesSuggesteds = valuesSuggesteds;
        }
    }
    }
}
