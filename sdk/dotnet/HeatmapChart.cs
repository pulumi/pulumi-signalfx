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
    /// This chart type shows the specified plot in a heat map fashion. This format is similar to the [Infrastructure Navigator](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/built-in-content/infra-nav.html#infra), with squares representing each source for the selected metric, and the color of each square representing the value range of the metric.
    /// 
    /// ## Example
    /// 
    /// ## Arguments
    /// 
    /// The following arguments are supported in the resource block:
    /// 
    /// * `name` - (Required) Name of the chart.
    /// * `program_text` - (Required) Signalflow program text for the chart. More info at &lt;https://dev.splunk.com/observability/docs/signalflow/&gt;.
    /// * `description` - (Optional) Description of the chart.
    /// * `unit_prefix` - (Optional) Must be `"Metric"` or `"Binary`". `"Metric"` by default.
    /// * `minimum_resolution` - (Optional) The minimum resolution (in seconds) to use for computing the underlying program.
    /// * `max_delay` - (Optional) How long (in seconds) to wait for late datapoints.
    /// * `timezone` - (Optional) The property value is a string that denotes the geographic region associated with the time zone, (default UTC).
    /// * `refresh_interval` - (Optional) How often (in seconds) to refresh the values of the heatmap.
    /// * `disable_sampling` - (Optional) If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default.
    /// * `group_by` - (Optional) Properties to group by in the heatmap (in nesting order).
    /// * `sort_by` - (Optional) The property to use when sorting the elements. Must be prepended with `+` for ascending or `-` for descending (e.g. `-foo`).
    /// * `hide_timestamp` - (Optional) Whether to show the timestamp in the chart. `false` by default.
    /// * `color_range` - (Optional, Default) Values and color for the color range. Example: `color_range : { min : 0, max : 100, color : "#0000ff" }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
    ///     * `min_value` - (Optional) The minimum value within the coloring range.
    ///     * `max_value` - (Optional) The maximum value within the coloring range.
    ///     * `color` - (Required) The color range to use. The starting hex color value for data values in a heatmap chart. Specify the value as a 6-character hexadecimal value preceded by the '#' character, for example "#ea1849" (grass green).
    /// * `color_scale` - (Optional.  Conflicts with `color_range`) One to N blocks, each defining a single color range including both the color to display for that range and the borders of the range. Example: `color_scale { gt = 60, color = "blue" } color_scale { lte = 60, color = "yellow" }`. Look at this [link](https://docs.splunk.com/observability/en/data-visualization/charts/chart-options.html).
    ///     * `gt` - (Optional) Indicates the lower threshold non-inclusive value for this range.
    ///     * `gte` - (Optional) Indicates the lower threshold inclusive value for this range.
    ///     * `lt` - (Optional) Indicates the upper threshold non-inclusive value for this range.
    ///     * `lte` - (Optional) Indicates the upper threshold inclusive value for this range.
    ///     * `color` - (Required) The color range to use. Hex values are not supported here. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
    /// 
    /// ## Attributes
    /// 
    /// In a addition to all arguments above, the following attributes are exported:
    /// 
    /// * `id` - The ID of the chart.
    /// * `url` - The URL of the chart.
    /// </summary>
    [SignalFxResourceType("signalfx:index/heatmapChart:HeatmapChart")]
    public partial class HeatmapChart : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Values and color for the color range. Example: colorRange : { min : 0, max : 100, color : "#0000ff" }
        /// </summary>
        [Output("colorRange")]
        public Output<Outputs.HeatmapChartColorRange?> ColorRange { get; private set; } = null!;

        /// <summary>
        /// Single color range including both the color to display for that range and the borders of the range
        /// </summary>
        [Output("colorScales")]
        public Output<ImmutableArray<Outputs.HeatmapChartColorScale>> ColorScales { get; private set; } = null!;

        /// <summary>
        /// Description of the chart (Optional)
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// (false by default) If false, samples a subset of the output MTS, which improves UI performance
        /// </summary>
        [Output("disableSampling")]
        public Output<bool?> DisableSampling { get; private set; } = null!;

        /// <summary>
        /// Properties to group by in the heatmap (in nesting order)
        /// </summary>
        [Output("groupBies")]
        public Output<ImmutableArray<string>> GroupBies { get; private set; } = null!;

        /// <summary>
        /// (false by default) Whether to show the timestamp in the chart
        /// </summary>
        [Output("hideTimestamp")]
        public Output<bool?> HideTimestamp { get; private set; } = null!;

        /// <summary>
        /// How long (in seconds) to wait for late datapoints
        /// </summary>
        [Output("maxDelay")]
        public Output<int?> MaxDelay { get; private set; } = null!;

        /// <summary>
        /// The minimum resolution (in seconds) to use for computing the underlying program
        /// </summary>
        [Output("minimumResolution")]
        public Output<int?> MinimumResolution { get; private set; } = null!;

        /// <summary>
        /// Name of the chart
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
        /// </summary>
        [Output("programText")]
        public Output<string> ProgramText { get; private set; } = null!;

        /// <summary>
        /// How often (in seconds) to refresh the values of the heatmap
        /// </summary>
        [Output("refreshInterval")]
        public Output<int?> RefreshInterval { get; private set; } = null!;

        /// <summary>
        /// The property to use when sorting the elements. Must be prepended with + for ascending or - for descending (e.g. -foo)
        /// </summary>
        [Output("sortBy")]
        public Output<string?> SortBy { get; private set; } = null!;

        /// <summary>
        /// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;

        /// <summary>
        /// (Metric by default) Must be "Metric" or "Binary"
        /// </summary>
        [Output("unitPrefix")]
        public Output<string?> UnitPrefix { get; private set; } = null!;

        /// <summary>
        /// URL of the chart
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

    public sealed class HeatmapChartArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Values and color for the color range. Example: colorRange : { min : 0, max : 100, color : "#0000ff" }
        /// </summary>
        [Input("colorRange")]
        public Input<Inputs.HeatmapChartColorRangeArgs>? ColorRange { get; set; }

        [Input("colorScales")]
        private InputList<Inputs.HeatmapChartColorScaleArgs>? _colorScales;

        /// <summary>
        /// Single color range including both the color to display for that range and the borders of the range
        /// </summary>
        public InputList<Inputs.HeatmapChartColorScaleArgs> ColorScales
        {
            get => _colorScales ?? (_colorScales = new InputList<Inputs.HeatmapChartColorScaleArgs>());
            set => _colorScales = value;
        }

        /// <summary>
        /// Description of the chart (Optional)
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (false by default) If false, samples a subset of the output MTS, which improves UI performance
        /// </summary>
        [Input("disableSampling")]
        public Input<bool>? DisableSampling { get; set; }

        [Input("groupBies")]
        private InputList<string>? _groupBies;

        /// <summary>
        /// Properties to group by in the heatmap (in nesting order)
        /// </summary>
        public InputList<string> GroupBies
        {
            get => _groupBies ?? (_groupBies = new InputList<string>());
            set => _groupBies = value;
        }

        /// <summary>
        /// (false by default) Whether to show the timestamp in the chart
        /// </summary>
        [Input("hideTimestamp")]
        public Input<bool>? HideTimestamp { get; set; }

        /// <summary>
        /// How long (in seconds) to wait for late datapoints
        /// </summary>
        [Input("maxDelay")]
        public Input<int>? MaxDelay { get; set; }

        /// <summary>
        /// The minimum resolution (in seconds) to use for computing the underlying program
        /// </summary>
        [Input("minimumResolution")]
        public Input<int>? MinimumResolution { get; set; }

        /// <summary>
        /// Name of the chart
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
        /// </summary>
        [Input("programText", required: true)]
        public Input<string> ProgramText { get; set; } = null!;

        /// <summary>
        /// How often (in seconds) to refresh the values of the heatmap
        /// </summary>
        [Input("refreshInterval")]
        public Input<int>? RefreshInterval { get; set; }

        /// <summary>
        /// The property to use when sorting the elements. Must be prepended with + for ascending or - for descending (e.g. -foo)
        /// </summary>
        [Input("sortBy")]
        public Input<string>? SortBy { get; set; }

        /// <summary>
        /// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// (Metric by default) Must be "Metric" or "Binary"
        /// </summary>
        [Input("unitPrefix")]
        public Input<string>? UnitPrefix { get; set; }

        public HeatmapChartArgs()
        {
        }
        public static new HeatmapChartArgs Empty => new HeatmapChartArgs();
    }

    public sealed class HeatmapChartState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Values and color for the color range. Example: colorRange : { min : 0, max : 100, color : "#0000ff" }
        /// </summary>
        [Input("colorRange")]
        public Input<Inputs.HeatmapChartColorRangeGetArgs>? ColorRange { get; set; }

        [Input("colorScales")]
        private InputList<Inputs.HeatmapChartColorScaleGetArgs>? _colorScales;

        /// <summary>
        /// Single color range including both the color to display for that range and the borders of the range
        /// </summary>
        public InputList<Inputs.HeatmapChartColorScaleGetArgs> ColorScales
        {
            get => _colorScales ?? (_colorScales = new InputList<Inputs.HeatmapChartColorScaleGetArgs>());
            set => _colorScales = value;
        }

        /// <summary>
        /// Description of the chart (Optional)
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (false by default) If false, samples a subset of the output MTS, which improves UI performance
        /// </summary>
        [Input("disableSampling")]
        public Input<bool>? DisableSampling { get; set; }

        [Input("groupBies")]
        private InputList<string>? _groupBies;

        /// <summary>
        /// Properties to group by in the heatmap (in nesting order)
        /// </summary>
        public InputList<string> GroupBies
        {
            get => _groupBies ?? (_groupBies = new InputList<string>());
            set => _groupBies = value;
        }

        /// <summary>
        /// (false by default) Whether to show the timestamp in the chart
        /// </summary>
        [Input("hideTimestamp")]
        public Input<bool>? HideTimestamp { get; set; }

        /// <summary>
        /// How long (in seconds) to wait for late datapoints
        /// </summary>
        [Input("maxDelay")]
        public Input<int>? MaxDelay { get; set; }

        /// <summary>
        /// The minimum resolution (in seconds) to use for computing the underlying program
        /// </summary>
        [Input("minimumResolution")]
        public Input<int>? MinimumResolution { get; set; }

        /// <summary>
        /// Name of the chart
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Signalflow program text for the chart. More info at "https://developers.signalfx.com/docs/signalflow-overview"
        /// </summary>
        [Input("programText")]
        public Input<string>? ProgramText { get; set; }

        /// <summary>
        /// How often (in seconds) to refresh the values of the heatmap
        /// </summary>
        [Input("refreshInterval")]
        public Input<int>? RefreshInterval { get; set; }

        /// <summary>
        /// The property to use when sorting the elements. Must be prepended with + for ascending or - for descending (e.g. -foo)
        /// </summary>
        [Input("sortBy")]
        public Input<string>? SortBy { get; set; }

        /// <summary>
        /// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// (Metric by default) Must be "Metric" or "Binary"
        /// </summary>
        [Input("unitPrefix")]
        public Input<string>? UnitPrefix { get; set; }

        /// <summary>
        /// URL of the chart
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public HeatmapChartState()
        {
        }
        public static new HeatmapChartState Empty => new HeatmapChartState();
    }
}
