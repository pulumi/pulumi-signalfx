// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx
{
    public partial class Dashboard : Pulumi.CustomResource
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
        /// Chart ID and layout information for the charts in the dashboard.
        /// </summary>
        [Output("charts")]
        public Output<ImmutableArray<Outputs.DashboardChart>> Charts { get; private set; } = null!;

        /// <summary>
        /// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
        /// </summary>
        [Output("chartsResolution")]
        public Output<string?> ChartsResolution { get; private set; } = null!;

        /// <summary>
        /// Column number for the layout.
        /// </summary>
        [Output("columns")]
        public Output<ImmutableArray<Outputs.DashboardColumn>> Columns { get; private set; } = null!;

        /// <summary>
        /// The ID of the dashboard group that contains the dashboard.
        /// </summary>
        [Output("dashboardGroup")]
        public Output<string> DashboardGroup { get; private set; } = null!;

        /// <summary>
        /// Variable description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("discoveryOptionsQuery")]
        public Output<string?> DiscoveryOptionsQuery { get; private set; } = null!;

        [Output("discoveryOptionsSelectors")]
        public Output<ImmutableArray<string>> DiscoveryOptionsSelectors { get; private set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization.
        /// </summary>
        [Output("endTime")]
        public Output<int?> EndTime { get; private set; } = null!;

        /// <summary>
        /// Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
        /// </summary>
        [Output("eventOverlays")]
        public Output<ImmutableArray<Outputs.DashboardEventOverlay>> EventOverlays { get; private set; } = null!;

        /// <summary>
        /// Filter to apply to the charts when displaying the dashboard.
        /// </summary>
        [Output("filters")]
        public Output<ImmutableArray<Outputs.DashboardFilter>> Filters { get; private set; } = null!;

        /// <summary>
        /// Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
        /// </summary>
        [Output("grids")]
        public Output<ImmutableArray<Outputs.DashboardGrid>> Grids { get; private set; } = null!;

        /// <summary>
        /// Name of the dashboard.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
        /// </summary>
        [Output("selectedEventOverlays")]
        public Output<ImmutableArray<Outputs.DashboardSelectedEventOverlay>> SelectedEventOverlays { get; private set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization.
        /// </summary>
        [Output("startTime")]
        public Output<int?> StartTime { get; private set; } = null!;

        /// <summary>
        /// The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
        /// </summary>
        [Output("timeRange")]
        public Output<string?> TimeRange { get; private set; } = null!;

        /// <summary>
        /// The URL of the dashboard.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Dashboard variable to apply to each chart in the dashboard.
        /// </summary>
        [Output("variables")]
        public Output<ImmutableArray<Outputs.DashboardVariable>> Variables { get; private set; } = null!;


        /// <summary>
        /// Create a Dashboard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dashboard(string name, DashboardArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/dashboard:Dashboard", name, args ?? new DashboardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dashboard(string name, Input<string> id, DashboardState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/dashboard:Dashboard", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dashboard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dashboard Get(string name, Input<string> id, DashboardState? state = null, CustomResourceOptions? options = null)
        {
            return new Dashboard(name, id, state, options);
        }
    }

    public sealed class DashboardArgs : Pulumi.ResourceArgs
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

        [Input("charts")]
        private InputList<Inputs.DashboardChartArgs>? _charts;

        /// <summary>
        /// Chart ID and layout information for the charts in the dashboard.
        /// </summary>
        public InputList<Inputs.DashboardChartArgs> Charts
        {
            get => _charts ?? (_charts = new InputList<Inputs.DashboardChartArgs>());
            set => _charts = value;
        }

        /// <summary>
        /// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
        /// </summary>
        [Input("chartsResolution")]
        public Input<string>? ChartsResolution { get; set; }

        [Input("columns")]
        private InputList<Inputs.DashboardColumnArgs>? _columns;

        /// <summary>
        /// Column number for the layout.
        /// </summary>
        public InputList<Inputs.DashboardColumnArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.DashboardColumnArgs>());
            set => _columns = value;
        }

        /// <summary>
        /// The ID of the dashboard group that contains the dashboard.
        /// </summary>
        [Input("dashboardGroup", required: true)]
        public Input<string> DashboardGroup { get; set; } = null!;

        /// <summary>
        /// Variable description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("discoveryOptionsQuery")]
        public Input<string>? DiscoveryOptionsQuery { get; set; }

        [Input("discoveryOptionsSelectors")]
        private InputList<string>? _discoveryOptionsSelectors;
        public InputList<string> DiscoveryOptionsSelectors
        {
            get => _discoveryOptionsSelectors ?? (_discoveryOptionsSelectors = new InputList<string>());
            set => _discoveryOptionsSelectors = value;
        }

        /// <summary>
        /// Seconds since epoch. Used for visualization.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        [Input("eventOverlays")]
        private InputList<Inputs.DashboardEventOverlayArgs>? _eventOverlays;

        /// <summary>
        /// Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
        /// </summary>
        public InputList<Inputs.DashboardEventOverlayArgs> EventOverlays
        {
            get => _eventOverlays ?? (_eventOverlays = new InputList<Inputs.DashboardEventOverlayArgs>());
            set => _eventOverlays = value;
        }

        [Input("filters")]
        private InputList<Inputs.DashboardFilterArgs>? _filters;

        /// <summary>
        /// Filter to apply to the charts when displaying the dashboard.
        /// </summary>
        public InputList<Inputs.DashboardFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.DashboardFilterArgs>());
            set => _filters = value;
        }

        [Input("grids")]
        private InputList<Inputs.DashboardGridArgs>? _grids;

        /// <summary>
        /// Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
        /// </summary>
        public InputList<Inputs.DashboardGridArgs> Grids
        {
            get => _grids ?? (_grids = new InputList<Inputs.DashboardGridArgs>());
            set => _grids = value;
        }

        /// <summary>
        /// Name of the dashboard.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("selectedEventOverlays")]
        private InputList<Inputs.DashboardSelectedEventOverlayArgs>? _selectedEventOverlays;

        /// <summary>
        /// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
        /// </summary>
        public InputList<Inputs.DashboardSelectedEventOverlayArgs> SelectedEventOverlays
        {
            get => _selectedEventOverlays ?? (_selectedEventOverlays = new InputList<Inputs.DashboardSelectedEventOverlayArgs>());
            set => _selectedEventOverlays = value;
        }

        /// <summary>
        /// Seconds since epoch. Used for visualization.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
        /// </summary>
        [Input("timeRange")]
        public Input<string>? TimeRange { get; set; }

        [Input("variables")]
        private InputList<Inputs.DashboardVariableArgs>? _variables;

        /// <summary>
        /// Dashboard variable to apply to each chart in the dashboard.
        /// </summary>
        public InputList<Inputs.DashboardVariableArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.DashboardVariableArgs>());
            set => _variables = value;
        }

        public DashboardArgs()
        {
        }
    }

    public sealed class DashboardState : Pulumi.ResourceArgs
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

        [Input("charts")]
        private InputList<Inputs.DashboardChartGetArgs>? _charts;

        /// <summary>
        /// Chart ID and layout information for the charts in the dashboard.
        /// </summary>
        public InputList<Inputs.DashboardChartGetArgs> Charts
        {
            get => _charts ?? (_charts = new InputList<Inputs.DashboardChartGetArgs>());
            set => _charts = value;
        }

        /// <summary>
        /// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
        /// </summary>
        [Input("chartsResolution")]
        public Input<string>? ChartsResolution { get; set; }

        [Input("columns")]
        private InputList<Inputs.DashboardColumnGetArgs>? _columns;

        /// <summary>
        /// Column number for the layout.
        /// </summary>
        public InputList<Inputs.DashboardColumnGetArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.DashboardColumnGetArgs>());
            set => _columns = value;
        }

        /// <summary>
        /// The ID of the dashboard group that contains the dashboard.
        /// </summary>
        [Input("dashboardGroup")]
        public Input<string>? DashboardGroup { get; set; }

        /// <summary>
        /// Variable description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("discoveryOptionsQuery")]
        public Input<string>? DiscoveryOptionsQuery { get; set; }

        [Input("discoveryOptionsSelectors")]
        private InputList<string>? _discoveryOptionsSelectors;
        public InputList<string> DiscoveryOptionsSelectors
        {
            get => _discoveryOptionsSelectors ?? (_discoveryOptionsSelectors = new InputList<string>());
            set => _discoveryOptionsSelectors = value;
        }

        /// <summary>
        /// Seconds since epoch. Used for visualization.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        [Input("eventOverlays")]
        private InputList<Inputs.DashboardEventOverlayGetArgs>? _eventOverlays;

        /// <summary>
        /// Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
        /// </summary>
        public InputList<Inputs.DashboardEventOverlayGetArgs> EventOverlays
        {
            get => _eventOverlays ?? (_eventOverlays = new InputList<Inputs.DashboardEventOverlayGetArgs>());
            set => _eventOverlays = value;
        }

        [Input("filters")]
        private InputList<Inputs.DashboardFilterGetArgs>? _filters;

        /// <summary>
        /// Filter to apply to the charts when displaying the dashboard.
        /// </summary>
        public InputList<Inputs.DashboardFilterGetArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.DashboardFilterGetArgs>());
            set => _filters = value;
        }

        [Input("grids")]
        private InputList<Inputs.DashboardGridGetArgs>? _grids;

        /// <summary>
        /// Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
        /// </summary>
        public InputList<Inputs.DashboardGridGetArgs> Grids
        {
            get => _grids ?? (_grids = new InputList<Inputs.DashboardGridGetArgs>());
            set => _grids = value;
        }

        /// <summary>
        /// Name of the dashboard.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("selectedEventOverlays")]
        private InputList<Inputs.DashboardSelectedEventOverlayGetArgs>? _selectedEventOverlays;

        /// <summary>
        /// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
        /// </summary>
        public InputList<Inputs.DashboardSelectedEventOverlayGetArgs> SelectedEventOverlays
        {
            get => _selectedEventOverlays ?? (_selectedEventOverlays = new InputList<Inputs.DashboardSelectedEventOverlayGetArgs>());
            set => _selectedEventOverlays = value;
        }

        /// <summary>
        /// Seconds since epoch. Used for visualization.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
        /// </summary>
        [Input("timeRange")]
        public Input<string>? TimeRange { get; set; }

        /// <summary>
        /// The URL of the dashboard.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("variables")]
        private InputList<Inputs.DashboardVariableGetArgs>? _variables;

        /// <summary>
        /// Dashboard variable to apply to each chart in the dashboard.
        /// </summary>
        public InputList<Inputs.DashboardVariableGetArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.DashboardVariableGetArgs>());
            set => _variables = value;
        }

        public DashboardState()
        {
        }
    }
}
