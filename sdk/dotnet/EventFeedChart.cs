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
    /// Displays a listing of events as a widget in a dashboard.
    /// </summary>
    [SignalFxResourceType("signalfx:index/eventFeedChart:EventFeedChart")]
    public partial class EventFeedChart : Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the text note.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Output("endTime")]
        public Output<int?> EndTime { get; private set; } = null!;

        /// <summary>
        /// Name of the text note.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        /// </summary>
        [Output("programText")]
        public Output<string> ProgramText { get; private set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Output("startTime")]
        public Output<int?> StartTime { get; private set; } = null!;

        /// <summary>
        /// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        /// </summary>
        [Output("timeRange")]
        public Output<int?> TimeRange { get; private set; } = null!;

        /// <summary>
        /// The URL of the chart.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a EventFeedChart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventFeedChart(string name, EventFeedChartArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/eventFeedChart:EventFeedChart", name, args ?? new EventFeedChartArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventFeedChart(string name, Input<string> id, EventFeedChartState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/eventFeedChart:EventFeedChart", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventFeedChart resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventFeedChart Get(string name, Input<string> id, EventFeedChartState? state = null, CustomResourceOptions? options = null)
        {
            return new EventFeedChart(name, id, state, options);
        }
    }

    public sealed class EventFeedChartArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the text note.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        /// <summary>
        /// Name of the text note.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        /// </summary>
        [Input("programText", required: true)]
        public Input<string> ProgramText { get; set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        /// </summary>
        [Input("timeRange")]
        public Input<int>? TimeRange { get; set; }

        public EventFeedChartArgs()
        {
        }
    }

    public sealed class EventFeedChartState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the text note.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        /// <summary>
        /// Name of the text note.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        /// </summary>
        [Input("programText")]
        public Input<string>? ProgramText { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        /// </summary>
        [Input("timeRange")]
        public Input<int>? TimeRange { get; set; }

        /// <summary>
        /// The URL of the chart.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public EventFeedChartState()
        {
        }
    }
}
