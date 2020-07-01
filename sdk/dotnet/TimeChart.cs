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
    /// Provides a SignalFx time chart resource. This can be used to create and manage the different types of time charts.
    /// 
    /// Time charts display data points over a period of time.
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
    ///         var mychart0 = new SignalFx.TimeChart("mychart0", new SignalFx.TimeChartArgs
    ///         {
    ///             AxisLeft = new SignalFx.Inputs.TimeChartAxisLeftArgs
    ///             {
    ///                 Label = "CPU Total Idle",
    ///                 LowWatermark = 1000,
    ///             },
    ///             LegendOptionsFields = 
    ///             {
    ///                 new SignalFx.Inputs.TimeChartLegendOptionsFieldArgs
    ///                 {
    ///                     Enabled = false,
    ///                     Property = "collector",
    ///                 },
    ///                 new SignalFx.Inputs.TimeChartLegendOptionsFieldArgs
    ///                 {
    ///                     Enabled = false,
    ///                     Property = "hostname",
    ///                 },
    ///             },
    ///             PlotType = "LineChart",
    ///             ProgramText = @"data(""cpu.total.idle"").publish(label=""CPU Idle"")
    /// 
    /// ",
    ///             ShowDataMarkers = true,
    ///             TimeRange = 3600,
    ///             VizOptions = 
    ///             {
    ///                 new SignalFx.Inputs.TimeChartVizOptionArgs
    ///                 {
    ///                     Axis = "left",
    ///                     Color = "orange",
    ///                     Label = "CPU Idle",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class TimeChart : Pulumi.CustomResource
    {
        /// <summary>
        /// Force the chart to display zero on the y-axes, even if none of the data is near zero.
        /// </summary>
        [Output("axesIncludeZero")]
        public Output<bool?> AxesIncludeZero { get; private set; } = null!;

        /// <summary>
        /// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
        /// </summary>
        [Output("axesPrecision")]
        public Output<int?> AxesPrecision { get; private set; } = null!;

        /// <summary>
        /// Set of axis options.
        /// </summary>
        [Output("axisLeft")]
        public Output<Outputs.TimeChartAxisLeft?> AxisLeft { get; private set; } = null!;

        /// <summary>
        /// Set of axis options.
        /// </summary>
        [Output("axisRight")]
        public Output<Outputs.TimeChartAxisRight?> AxisRight { get; private set; } = null!;

        /// <summary>
        /// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
        /// </summary>
        [Output("colorBy")]
        public Output<string?> ColorBy { get; private set; } = null!;

        /// <summary>
        /// Description of the chart.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
        /// </summary>
        [Output("disableSampling")]
        public Output<bool?> DisableSampling { get; private set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Output("endTime")]
        public Output<int?> EndTime { get; private set; } = null!;

        /// <summary>
        /// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
        /// </summary>
        [Output("eventOptions")]
        public Output<ImmutableArray<Outputs.TimeChartEventOption>> EventOptions { get; private set; } = null!;

        /// <summary>
        /// Only used when `plot_type` is `"Histogram"`. Histogram specific options.
        /// </summary>
        [Output("histogramOptions")]
        public Output<ImmutableArray<Outputs.TimeChartHistogramOption>> HistogramOptions { get; private set; } = null!;

        /// <summary>
        /// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
        /// </summary>
        [Output("legendFieldsToHides")]
        public Output<ImmutableArray<string>> LegendFieldsToHides { get; private set; } = null!;

        /// <summary>
        /// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
        /// </summary>
        [Output("legendOptionsFields")]
        public Output<ImmutableArray<Outputs.TimeChartLegendOptionsField>> LegendOptionsFields { get; private set; } = null!;

        /// <summary>
        /// How long (in seconds) to wait for late datapoints.
        /// </summary>
        [Output("maxDelay")]
        public Output<int?> MaxDelay { get; private set; } = null!;

        /// <summary>
        /// The minimum resolution (in seconds) to use for computing the underlying program.
        /// </summary>
        [Output("minimumResolution")]
        public Output<int?> MinimumResolution { get; private set; } = null!;

        /// <summary>
        /// Name of the chart.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plot_label"` and any dimension.
        /// </summary>
        [Output("onChartLegendDimension")]
        public Output<string?> OnChartLegendDimension { get; private set; } = null!;

        /// <summary>
        /// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
        /// </summary>
        [Output("plotType")]
        public Output<string?> PlotType { get; private set; } = null!;

        /// <summary>
        /// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        /// </summary>
        [Output("programText")]
        public Output<string> ProgramText { get; private set; } = null!;

        /// <summary>
        /// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
        /// </summary>
        [Output("showDataMarkers")]
        public Output<bool?> ShowDataMarkers { get; private set; } = null!;

        /// <summary>
        /// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
        /// </summary>
        [Output("showEventLines")]
        public Output<bool?> ShowEventLines { get; private set; } = null!;

        /// <summary>
        /// Whether area and bar charts in the visualization should be stacked. `false` by default.
        /// </summary>
        [Output("stacked")]
        public Output<bool?> Stacked { get; private set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Output("startTime")]
        public Output<int?> StartTime { get; private set; } = null!;

        /// <summary>
        /// Tags associated with the chart
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
        /// </summary>
        [Output("timeRange")]
        public Output<int?> TimeRange { get; private set; } = null!;

        /// <summary>
        /// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;

        /// <summary>
        /// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        /// </summary>
        [Output("unitPrefix")]
        public Output<string?> UnitPrefix { get; private set; } = null!;

        /// <summary>
        /// URL of the chart
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Plot-level customization options, associated with a publish statement.
        /// </summary>
        [Output("vizOptions")]
        public Output<ImmutableArray<Outputs.TimeChartVizOption>> VizOptions { get; private set; } = null!;


        /// <summary>
        /// Create a TimeChart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TimeChart(string name, TimeChartArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/timeChart:TimeChart", name, args ?? new TimeChartArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TimeChart(string name, Input<string> id, TimeChartState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/timeChart:TimeChart", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TimeChart resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TimeChart Get(string name, Input<string> id, TimeChartState? state = null, CustomResourceOptions? options = null)
        {
            return new TimeChart(name, id, state, options);
        }
    }

    public sealed class TimeChartArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Force the chart to display zero on the y-axes, even if none of the data is near zero.
        /// </summary>
        [Input("axesIncludeZero")]
        public Input<bool>? AxesIncludeZero { get; set; }

        /// <summary>
        /// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
        /// </summary>
        [Input("axesPrecision")]
        public Input<int>? AxesPrecision { get; set; }

        /// <summary>
        /// Set of axis options.
        /// </summary>
        [Input("axisLeft")]
        public Input<Inputs.TimeChartAxisLeftArgs>? AxisLeft { get; set; }

        /// <summary>
        /// Set of axis options.
        /// </summary>
        [Input("axisRight")]
        public Input<Inputs.TimeChartAxisRightArgs>? AxisRight { get; set; }

        /// <summary>
        /// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
        /// </summary>
        [Input("colorBy")]
        public Input<string>? ColorBy { get; set; }

        /// <summary>
        /// Description of the chart.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
        /// </summary>
        [Input("disableSampling")]
        public Input<bool>? DisableSampling { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        [Input("eventOptions")]
        private InputList<Inputs.TimeChartEventOptionArgs>? _eventOptions;

        /// <summary>
        /// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
        /// </summary>
        public InputList<Inputs.TimeChartEventOptionArgs> EventOptions
        {
            get => _eventOptions ?? (_eventOptions = new InputList<Inputs.TimeChartEventOptionArgs>());
            set => _eventOptions = value;
        }

        [Input("histogramOptions")]
        private InputList<Inputs.TimeChartHistogramOptionArgs>? _histogramOptions;

        /// <summary>
        /// Only used when `plot_type` is `"Histogram"`. Histogram specific options.
        /// </summary>
        public InputList<Inputs.TimeChartHistogramOptionArgs> HistogramOptions
        {
            get => _histogramOptions ?? (_histogramOptions = new InputList<Inputs.TimeChartHistogramOptionArgs>());
            set => _histogramOptions = value;
        }

        [Input("legendFieldsToHides")]
        private InputList<string>? _legendFieldsToHides;

        /// <summary>
        /// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
        /// </summary>
        [Obsolete(@"Please use legend_options_fields")]
        public InputList<string> LegendFieldsToHides
        {
            get => _legendFieldsToHides ?? (_legendFieldsToHides = new InputList<string>());
            set => _legendFieldsToHides = value;
        }

        [Input("legendOptionsFields")]
        private InputList<Inputs.TimeChartLegendOptionsFieldArgs>? _legendOptionsFields;

        /// <summary>
        /// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
        /// </summary>
        public InputList<Inputs.TimeChartLegendOptionsFieldArgs> LegendOptionsFields
        {
            get => _legendOptionsFields ?? (_legendOptionsFields = new InputList<Inputs.TimeChartLegendOptionsFieldArgs>());
            set => _legendOptionsFields = value;
        }

        /// <summary>
        /// How long (in seconds) to wait for late datapoints.
        /// </summary>
        [Input("maxDelay")]
        public Input<int>? MaxDelay { get; set; }

        /// <summary>
        /// The minimum resolution (in seconds) to use for computing the underlying program.
        /// </summary>
        [Input("minimumResolution")]
        public Input<int>? MinimumResolution { get; set; }

        /// <summary>
        /// Name of the chart.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plot_label"` and any dimension.
        /// </summary>
        [Input("onChartLegendDimension")]
        public Input<string>? OnChartLegendDimension { get; set; }

        /// <summary>
        /// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
        /// </summary>
        [Input("plotType")]
        public Input<string>? PlotType { get; set; }

        /// <summary>
        /// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        /// </summary>
        [Input("programText", required: true)]
        public Input<string> ProgramText { get; set; } = null!;

        /// <summary>
        /// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
        /// </summary>
        [Input("showDataMarkers")]
        public Input<bool>? ShowDataMarkers { get; set; }

        /// <summary>
        /// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
        /// </summary>
        [Input("showEventLines")]
        public Input<bool>? ShowEventLines { get; set; }

        /// <summary>
        /// Whether area and bar charts in the visualization should be stacked. `false` by default.
        /// </summary>
        [Input("stacked")]
        public Input<bool>? Stacked { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags associated with the chart
        /// </summary>
        [Obsolete(@"signalfx_time_chart.tags is being removed in the next release")]
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
        /// </summary>
        [Input("timeRange")]
        public Input<int>? TimeRange { get; set; }

        /// <summary>
        /// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        /// </summary>
        [Input("unitPrefix")]
        public Input<string>? UnitPrefix { get; set; }

        [Input("vizOptions")]
        private InputList<Inputs.TimeChartVizOptionArgs>? _vizOptions;

        /// <summary>
        /// Plot-level customization options, associated with a publish statement.
        /// </summary>
        public InputList<Inputs.TimeChartVizOptionArgs> VizOptions
        {
            get => _vizOptions ?? (_vizOptions = new InputList<Inputs.TimeChartVizOptionArgs>());
            set => _vizOptions = value;
        }

        public TimeChartArgs()
        {
        }
    }

    public sealed class TimeChartState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Force the chart to display zero on the y-axes, even if none of the data is near zero.
        /// </summary>
        [Input("axesIncludeZero")]
        public Input<bool>? AxesIncludeZero { get; set; }

        /// <summary>
        /// Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
        /// </summary>
        [Input("axesPrecision")]
        public Input<int>? AxesPrecision { get; set; }

        /// <summary>
        /// Set of axis options.
        /// </summary>
        [Input("axisLeft")]
        public Input<Inputs.TimeChartAxisLeftGetArgs>? AxisLeft { get; set; }

        /// <summary>
        /// Set of axis options.
        /// </summary>
        [Input("axisRight")]
        public Input<Inputs.TimeChartAxisRightGetArgs>? AxisRight { get; set; }

        /// <summary>
        /// Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
        /// </summary>
        [Input("colorBy")]
        public Input<string>? ColorBy { get; set; }

        /// <summary>
        /// Description of the chart.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
        /// </summary>
        [Input("disableSampling")]
        public Input<bool>? DisableSampling { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        [Input("eventOptions")]
        private InputList<Inputs.TimeChartEventOptionGetArgs>? _eventOptions;

        /// <summary>
        /// Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
        /// </summary>
        public InputList<Inputs.TimeChartEventOptionGetArgs> EventOptions
        {
            get => _eventOptions ?? (_eventOptions = new InputList<Inputs.TimeChartEventOptionGetArgs>());
            set => _eventOptions = value;
        }

        [Input("histogramOptions")]
        private InputList<Inputs.TimeChartHistogramOptionGetArgs>? _histogramOptions;

        /// <summary>
        /// Only used when `plot_type` is `"Histogram"`. Histogram specific options.
        /// </summary>
        public InputList<Inputs.TimeChartHistogramOptionGetArgs> HistogramOptions
        {
            get => _histogramOptions ?? (_histogramOptions = new InputList<Inputs.TimeChartHistogramOptionGetArgs>());
            set => _histogramOptions = value;
        }

        [Input("legendFieldsToHides")]
        private InputList<string>? _legendFieldsToHides;

        /// <summary>
        /// List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
        /// </summary>
        [Obsolete(@"Please use legend_options_fields")]
        public InputList<string> LegendFieldsToHides
        {
            get => _legendFieldsToHides ?? (_legendFieldsToHides = new InputList<string>());
            set => _legendFieldsToHides = value;
        }

        [Input("legendOptionsFields")]
        private InputList<Inputs.TimeChartLegendOptionsFieldGetArgs>? _legendOptionsFields;

        /// <summary>
        /// List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
        /// </summary>
        public InputList<Inputs.TimeChartLegendOptionsFieldGetArgs> LegendOptionsFields
        {
            get => _legendOptionsFields ?? (_legendOptionsFields = new InputList<Inputs.TimeChartLegendOptionsFieldGetArgs>());
            set => _legendOptionsFields = value;
        }

        /// <summary>
        /// How long (in seconds) to wait for late datapoints.
        /// </summary>
        [Input("maxDelay")]
        public Input<int>? MaxDelay { get; set; }

        /// <summary>
        /// The minimum resolution (in seconds) to use for computing the underlying program.
        /// </summary>
        [Input("minimumResolution")]
        public Input<int>? MinimumResolution { get; set; }

        /// <summary>
        /// Name of the chart.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plot_label"` and any dimension.
        /// </summary>
        [Input("onChartLegendDimension")]
        public Input<string>? OnChartLegendDimension { get; set; }

        /// <summary>
        /// The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
        /// </summary>
        [Input("plotType")]
        public Input<string>? PlotType { get; set; }

        /// <summary>
        /// Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        /// </summary>
        [Input("programText")]
        public Input<string>? ProgramText { get; set; }

        /// <summary>
        /// Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
        /// </summary>
        [Input("showDataMarkers")]
        public Input<bool>? ShowDataMarkers { get; set; }

        /// <summary>
        /// Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
        /// </summary>
        [Input("showEventLines")]
        public Input<bool>? ShowEventLines { get; set; }

        /// <summary>
        /// Whether area and bar charts in the visualization should be stacked. `false` by default.
        /// </summary>
        [Input("stacked")]
        public Input<bool>? Stacked { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags associated with the chart
        /// </summary>
        [Obsolete(@"signalfx_time_chart.tags is being removed in the next release")]
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
        /// </summary>
        [Input("timeRange")]
        public Input<int>? TimeRange { get; set; }

        /// <summary>
        /// Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        /// </summary>
        [Input("unitPrefix")]
        public Input<string>? UnitPrefix { get; set; }

        /// <summary>
        /// URL of the chart
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("vizOptions")]
        private InputList<Inputs.TimeChartVizOptionGetArgs>? _vizOptions;

        /// <summary>
        /// Plot-level customization options, associated with a publish statement.
        /// </summary>
        public InputList<Inputs.TimeChartVizOptionGetArgs> VizOptions
        {
            get => _vizOptions ?? (_vizOptions = new InputList<Inputs.TimeChartVizOptionGetArgs>());
            set => _vizOptions = value;
        }

        public TimeChartState()
        {
        }
    }
}
