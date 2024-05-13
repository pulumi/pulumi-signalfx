// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Log
{
    /// <summary>
    /// You can add logs data to your Observability Cloud dashboards without turning your logs into metrics first.
    /// 
    /// A log view displays log lines in a table form in a dashboard and shows you in detail what is happening and why.
    /// 
    /// ## Example
    /// </summary>
    [SignalFxResourceType("signalfx:log/view:View")]
    public partial class View : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The column headers to show on the log view.
        /// </summary>
        [Output("columns")]
        public Output<ImmutableArray<Outputs.ViewColumn>> Columns { get; private set; } = null!;

        /// <summary>
        /// The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
        /// </summary>
        [Output("defaultConnection")]
        public Output<string?> DefaultConnection { get; private set; } = null!;

        /// <summary>
        /// Description of the log view.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Output("endTime")]
        public Output<int?> EndTime { get; private set; } = null!;

        /// <summary>
        /// Name of the log view.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
        /// </summary>
        [Output("programText")]
        public Output<string> ProgramText { get; private set; } = null!;

        /// <summary>
        /// The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
        /// </summary>
        [Output("sortOptions")]
        public Output<ImmutableArray<Outputs.ViewSortOption>> SortOptions { get; private set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Output("startTime")]
        public Output<int?> StartTime { get; private set; } = null!;

        /// <summary>
        /// From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        /// </summary>
        [Output("timeRange")]
        public Output<int?> TimeRange { get; private set; } = null!;

        /// <summary>
        /// The URL of the log view.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a View resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public View(string name, ViewArgs args, CustomResourceOptions? options = null)
            : base("signalfx:log/view:View", name, args ?? new ViewArgs(), MakeResourceOptions(options, ""))
        {
        }

        private View(string name, Input<string> id, ViewState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:log/view:View", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "signalfx:logs/view:View" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing View resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static View Get(string name, Input<string> id, ViewState? state = null, CustomResourceOptions? options = null)
        {
            return new View(name, id, state, options);
        }
    }

    public sealed class ViewArgs : global::Pulumi.ResourceArgs
    {
        [Input("columns")]
        private InputList<Inputs.ViewColumnArgs>? _columns;

        /// <summary>
        /// The column headers to show on the log view.
        /// </summary>
        public InputList<Inputs.ViewColumnArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.ViewColumnArgs>());
            set => _columns = value;
        }

        /// <summary>
        /// The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
        /// </summary>
        [Input("defaultConnection")]
        public Input<string>? DefaultConnection { get; set; }

        /// <summary>
        /// Description of the log view.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        /// <summary>
        /// Name of the log view.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
        /// </summary>
        [Input("programText", required: true)]
        public Input<string> ProgramText { get; set; } = null!;

        [Input("sortOptions")]
        private InputList<Inputs.ViewSortOptionArgs>? _sortOptions;

        /// <summary>
        /// The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
        /// </summary>
        public InputList<Inputs.ViewSortOptionArgs> SortOptions
        {
            get => _sortOptions ?? (_sortOptions = new InputList<Inputs.ViewSortOptionArgs>());
            set => _sortOptions = value;
        }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        /// </summary>
        [Input("timeRange")]
        public Input<int>? TimeRange { get; set; }

        public ViewArgs()
        {
        }
        public static new ViewArgs Empty => new ViewArgs();
    }

    public sealed class ViewState : global::Pulumi.ResourceArgs
    {
        [Input("columns")]
        private InputList<Inputs.ViewColumnGetArgs>? _columns;

        /// <summary>
        /// The column headers to show on the log view.
        /// </summary>
        public InputList<Inputs.ViewColumnGetArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.ViewColumnGetArgs>());
            set => _columns = value;
        }

        /// <summary>
        /// The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
        /// </summary>
        [Input("defaultConnection")]
        public Input<string>? DefaultConnection { get; set; }

        /// <summary>
        /// Description of the log view.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        /// <summary>
        /// Name of the log view.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
        /// </summary>
        [Input("programText")]
        public Input<string>? ProgramText { get; set; }

        [Input("sortOptions")]
        private InputList<Inputs.ViewSortOptionGetArgs>? _sortOptions;

        /// <summary>
        /// The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
        /// </summary>
        public InputList<Inputs.ViewSortOptionGetArgs> SortOptions
        {
            get => _sortOptions ?? (_sortOptions = new InputList<Inputs.ViewSortOptionGetArgs>());
            set => _sortOptions = value;
        }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        /// </summary>
        [Input("timeRange")]
        public Input<int>? TimeRange { get; set; }

        /// <summary>
        /// The URL of the log view.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ViewState()
        {
        }
        public static new ViewState Empty => new ViewState();
    }
}
