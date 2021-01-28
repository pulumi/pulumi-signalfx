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
    /// This chart type displays the specified plot in a heatmap fashion. This format is similar to the [Infrastructure Navigator](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/built-in-content/infra-nav.html#infra), with squares representing each source for the selected metric, and the color of each square representing the value range of the metric.
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
    ///         var myheatmapchart0 = new SignalFx.HeatmapChart("myheatmapchart0", new SignalFx.HeatmapChartArgs
    ///         {
    ///             ColorRange = new SignalFx.Inputs.HeatmapChartColorRangeArgs
    ///             {
    ///                 Color = "#ff0000",
    ///                 MaxValue = 100,
    ///                 MinValue = 0,
    ///             },
    ///             ColorScales = 
    ///             {
    ///                 new SignalFx.Inputs.HeatmapChartColorScaleArgs
    ///                 {
    ///                     Color = "green",
    ///                     Gte = 99,
    ///                 },
    ///                 new SignalFx.Inputs.HeatmapChartColorScaleArgs
    ///                 {
    ///                     Color = "yellow",
    ///                     Gte = 95,
    ///                     Lt = 99,
    ///                 },
    ///                 new SignalFx.Inputs.HeatmapChartColorScaleArgs
    ///                 {
    ///                     Color = "red",
    ///                     Lt = 95,
    ///                 },
    ///             },
    ///             Description = "Very cool Heatmap",
    ///             DisableSampling = true,
    ///             GroupBies = 
    ///             {
    ///                 "hostname",
    ///                 "host",
    ///             },
    ///             HideTimestamp = true,
    ///             ProgramText = @"myfilters = filter(""cluster_name"", ""prod"") and filter(""role"", ""search"")
    /// data(""cpu.total.idle"", filter=myfilters).publish()
    /// 
    /// ",
    ///             SortBy = "+host",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:index/heatmapChart:HeatmapChart")]
    public partial class HeatmapChart : Pulumi.CustomResource
    {
        /// <summary>
        /// Values and color for the color range. Example: `color_range : { min : 0, max : 100, color : "#0000ff" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        /// </summary>
        [Output("colorRange")]
        public Output<Outputs.HeatmapChartColorRange?> ColorRange { get; private set; } = null!;

        /// <summary>
        /// One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = "blue" } color_scale { lte = 60, color = "yellow" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        /// </summary>
        [Output("colorScales")]
        public Output<ImmutableArray<Outputs.HeatmapChartColorScale>> ColorScales { get; private set; } = null!;

        /// <summary>
        /// Description of the chart.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
        /// </summary>
        [Output("disableSampling")]
        public Output<bool?> DisableSampling { get; private set; } = null!;

        /// <summary>
        /// Properties to group by in the heatmap (in nesting order).
        /// </summary>
        [Output("groupBies")]
        public Output<ImmutableArray<string>> GroupBies { get; private set; } = null!;

        /// <summary>
        /// Whether to show the timestamp in the chart. `false` by default.
        /// </summary>
        [Output("hideTimestamp")]
        public Output<bool?> HideTimestamp { get; private set; } = null!;

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
        /// Signalflow program text for the chart. More info at &lt;https://developers.signalfx.com/docs/signalflow-overview&gt;.
        /// </summary>
        [Output("programText")]
        public Output<string> ProgramText { get; private set; } = null!;

        /// <summary>
        /// How often (in seconds) to refresh the values of the heatmap.
        /// </summary>
        [Output("refreshInterval")]
        public Output<int?> RefreshInterval { get; private set; } = null!;

        /// <summary>
        /// The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
        /// </summary>
        [Output("sortBy")]
        public Output<string?> SortBy { get; private set; } = null!;

        /// <summary>
        /// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        /// </summary>
        [Output("unitPrefix")]
        public Output<string?> UnitPrefix { get; private set; } = null!;

        /// <summary>
        /// The URL of the chart.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a HeatmapChart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HeatmapChart(string name, HeatmapChartArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/heatmapChart:HeatmapChart", name, args ?? new HeatmapChartArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HeatmapChart(string name, Input<string> id, HeatmapChartState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/heatmapChart:HeatmapChart", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HeatmapChart resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HeatmapChart Get(string name, Input<string> id, HeatmapChartState? state = null, CustomResourceOptions? options = null)
        {
            return new HeatmapChart(name, id, state, options);
        }
    }

    public sealed class HeatmapChartArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Values and color for the color range. Example: `color_range : { min : 0, max : 100, color : "#0000ff" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        /// </summary>
        [Input("colorRange")]
        public Input<Inputs.HeatmapChartColorRangeArgs>? ColorRange { get; set; }

        [Input("colorScales")]
        private InputList<Inputs.HeatmapChartColorScaleArgs>? _colorScales;

        /// <summary>
        /// One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = "blue" } color_scale { lte = 60, color = "yellow" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        /// </summary>
        public InputList<Inputs.HeatmapChartColorScaleArgs> ColorScales
        {
            get => _colorScales ?? (_colorScales = new InputList<Inputs.HeatmapChartColorScaleArgs>());
            set => _colorScales = value;
        }

        /// <summary>
        /// Description of the chart.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
        /// </summary>
        [Input("disableSampling")]
        public Input<bool>? DisableSampling { get; set; }

        [Input("groupBies")]
        private InputList<string>? _groupBies;

        /// <summary>
        /// Properties to group by in the heatmap (in nesting order).
        /// </summary>
        public InputList<string> GroupBies
        {
            get => _groupBies ?? (_groupBies = new InputList<string>());
            set => _groupBies = value;
        }

        /// <summary>
        /// Whether to show the timestamp in the chart. `false` by default.
        /// </summary>
        [Input("hideTimestamp")]
        public Input<bool>? HideTimestamp { get; set; }

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
        /// Signalflow program text for the chart. More info at &lt;https://developers.signalfx.com/docs/signalflow-overview&gt;.
        /// </summary>
        [Input("programText", required: true)]
        public Input<string> ProgramText { get; set; } = null!;

        /// <summary>
        /// How often (in seconds) to refresh the values of the heatmap.
        /// </summary>
        [Input("refreshInterval")]
        public Input<int>? RefreshInterval { get; set; }

        /// <summary>
        /// The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
        /// </summary>
        [Input("sortBy")]
        public Input<string>? SortBy { get; set; }

        /// <summary>
        /// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        /// </summary>
        [Input("unitPrefix")]
        public Input<string>? UnitPrefix { get; set; }

        public HeatmapChartArgs()
        {
        }
    }

    public sealed class HeatmapChartState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Values and color for the color range. Example: `color_range : { min : 0, max : 100, color : "#0000ff" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        /// </summary>
        [Input("colorRange")]
        public Input<Inputs.HeatmapChartColorRangeGetArgs>? ColorRange { get; set; }

        [Input("colorScales")]
        private InputList<Inputs.HeatmapChartColorScaleGetArgs>? _colorScales;

        /// <summary>
        /// One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = "blue" } color_scale { lte = 60, color = "yellow" }`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        /// </summary>
        public InputList<Inputs.HeatmapChartColorScaleGetArgs> ColorScales
        {
            get => _colorScales ?? (_colorScales = new InputList<Inputs.HeatmapChartColorScaleGetArgs>());
            set => _colorScales = value;
        }

        /// <summary>
        /// Description of the chart.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
        /// </summary>
        [Input("disableSampling")]
        public Input<bool>? DisableSampling { get; set; }

        [Input("groupBies")]
        private InputList<string>? _groupBies;

        /// <summary>
        /// Properties to group by in the heatmap (in nesting order).
        /// </summary>
        public InputList<string> GroupBies
        {
            get => _groupBies ?? (_groupBies = new InputList<string>());
            set => _groupBies = value;
        }

        /// <summary>
        /// Whether to show the timestamp in the chart. `false` by default.
        /// </summary>
        [Input("hideTimestamp")]
        public Input<bool>? HideTimestamp { get; set; }

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
        /// Signalflow program text for the chart. More info at &lt;https://developers.signalfx.com/docs/signalflow-overview&gt;.
        /// </summary>
        [Input("programText")]
        public Input<string>? ProgramText { get; set; }

        /// <summary>
        /// How often (in seconds) to refresh the values of the heatmap.
        /// </summary>
        [Input("refreshInterval")]
        public Input<int>? RefreshInterval { get; set; }

        /// <summary>
        /// The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
        /// </summary>
        [Input("sortBy")]
        public Input<string>? SortBy { get; set; }

        /// <summary>
        /// Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        /// </summary>
        [Input("unitPrefix")]
        public Input<string>? UnitPrefix { get; set; }

        /// <summary>
        /// The URL of the chart.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public HeatmapChartState()
        {
        }
    }
}
