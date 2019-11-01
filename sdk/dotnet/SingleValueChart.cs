// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Signalfx
{
    /// <summary>
    /// This chart type displays a single number in a large font, representing the current value of a single metric on a plot line.
    /// 
    /// If the time period is in the past, the number represents the value of the metric near the end of the time period.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/single_value_chart.html.markdown.
    /// </summary>
    public partial class SingleValueChart : Pulumi.CustomResource
    {
        /// <summary>
        /// Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
        /// </summary>
        [Output("colorBy")]
        public Output<string?> ColorBy { get; private set; } = null!;

        /// <summary>
        /// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        /// </summary>
        [Output("colorScales")]
        public Output<ImmutableArray<Outputs.SingleValueChartColorScales>> ColorScales { get; private set; } = null!;

        /// <summary>
        /// Description of the chart.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to hide the timestamp in the chart. `false` by default.
        /// </summary>
        [Output("isTimestampHidden")]
        public Output<bool?> IsTimestampHidden { get; private set; } = null!;

        /// <summary>
        /// How long (in seconds) to wait for late datapoints
        /// </summary>
        [Output("maxDelay")]
        public Output<int?> MaxDelay { get; private set; } = null!;

        /// <summary>
        /// The maximum precision to for value displayed.
        /// </summary>
        [Output("maxPrecision")]
        public Output<int?> MaxPrecision { get; private set; } = null!;

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
        /// How often (in seconds) to refresh the value.
        /// </summary>
        [Output("refreshInterval")]
        public Output<int?> RefreshInterval { get; private set; } = null!;

        /// <summary>
        /// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
        /// </summary>
        [Output("secondaryVisualization")]
        public Output<string?> SecondaryVisualization { get; private set; } = null!;

        /// <summary>
        /// Whether to show a trend line below the current value. `false` by default.
        /// </summary>
        [Output("showSparkLine")]
        public Output<bool?> ShowSparkLine { get; private set; } = null!;

        /// <summary>
        /// Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
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
        public Output<ImmutableArray<Outputs.SingleValueChartVizOptions>> VizOptions { get; private set; } = null!;


        /// <summary>
        /// Create a SingleValueChart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SingleValueChart(string name, SingleValueChartArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/singleValueChart:SingleValueChart", name, args, MakeResourceOptions(options, ""))
        {
        }

        private SingleValueChart(string name, Input<string> id, SingleValueChartState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/singleValueChart:SingleValueChart", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SingleValueChart resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SingleValueChart Get(string name, Input<string> id, SingleValueChartState? state = null, CustomResourceOptions? options = null)
        {
            return new SingleValueChart(name, id, state, options);
        }
    }

    public sealed class SingleValueChartArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
        /// </summary>
        [Input("colorBy")]
        public Input<string>? ColorBy { get; set; }

        [Input("colorScales")]
        private InputList<Inputs.SingleValueChartColorScalesArgs>? _colorScales;

        /// <summary>
        /// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        /// </summary>
        public InputList<Inputs.SingleValueChartColorScalesArgs> ColorScales
        {
            get => _colorScales ?? (_colorScales = new InputList<Inputs.SingleValueChartColorScalesArgs>());
            set => _colorScales = value;
        }

        /// <summary>
        /// Description of the chart.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to hide the timestamp in the chart. `false` by default.
        /// </summary>
        [Input("isTimestampHidden")]
        public Input<bool>? IsTimestampHidden { get; set; }

        /// <summary>
        /// How long (in seconds) to wait for late datapoints
        /// </summary>
        [Input("maxDelay")]
        public Input<int>? MaxDelay { get; set; }

        /// <summary>
        /// The maximum precision to for value displayed.
        /// </summary>
        [Input("maxPrecision")]
        public Input<int>? MaxPrecision { get; set; }

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
        /// How often (in seconds) to refresh the value.
        /// </summary>
        [Input("refreshInterval")]
        public Input<int>? RefreshInterval { get; set; }

        /// <summary>
        /// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
        /// </summary>
        [Input("secondaryVisualization")]
        public Input<string>? SecondaryVisualization { get; set; }

        /// <summary>
        /// Whether to show a trend line below the current value. `false` by default.
        /// </summary>
        [Input("showSparkLine")]
        public Input<bool>? ShowSparkLine { get; set; }

        /// <summary>
        /// Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
        /// </summary>
        [Input("unitPrefix")]
        public Input<string>? UnitPrefix { get; set; }

        [Input("vizOptions")]
        private InputList<Inputs.SingleValueChartVizOptionsArgs>? _vizOptions;

        /// <summary>
        /// Plot-level customization options, associated with a publish statement.
        /// </summary>
        public InputList<Inputs.SingleValueChartVizOptionsArgs> VizOptions
        {
            get => _vizOptions ?? (_vizOptions = new InputList<Inputs.SingleValueChartVizOptionsArgs>());
            set => _vizOptions = value;
        }

        public SingleValueChartArgs()
        {
        }
    }

    public sealed class SingleValueChartState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
        /// </summary>
        [Input("colorBy")]
        public Input<string>? ColorBy { get; set; }

        [Input("colorScales")]
        private InputList<Inputs.SingleValueChartColorScalesGetArgs>? _colorScales;

        /// <summary>
        /// Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        /// </summary>
        public InputList<Inputs.SingleValueChartColorScalesGetArgs> ColorScales
        {
            get => _colorScales ?? (_colorScales = new InputList<Inputs.SingleValueChartColorScalesGetArgs>());
            set => _colorScales = value;
        }

        /// <summary>
        /// Description of the chart.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to hide the timestamp in the chart. `false` by default.
        /// </summary>
        [Input("isTimestampHidden")]
        public Input<bool>? IsTimestampHidden { get; set; }

        /// <summary>
        /// How long (in seconds) to wait for late datapoints
        /// </summary>
        [Input("maxDelay")]
        public Input<int>? MaxDelay { get; set; }

        /// <summary>
        /// The maximum precision to for value displayed.
        /// </summary>
        [Input("maxPrecision")]
        public Input<int>? MaxPrecision { get; set; }

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
        /// How often (in seconds) to refresh the value.
        /// </summary>
        [Input("refreshInterval")]
        public Input<int>? RefreshInterval { get; set; }

        /// <summary>
        /// The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
        /// </summary>
        [Input("secondaryVisualization")]
        public Input<string>? SecondaryVisualization { get; set; }

        /// <summary>
        /// Whether to show a trend line below the current value. `false` by default.
        /// </summary>
        [Input("showSparkLine")]
        public Input<bool>? ShowSparkLine { get; set; }

        /// <summary>
        /// Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
        /// </summary>
        [Input("unitPrefix")]
        public Input<string>? UnitPrefix { get; set; }

        /// <summary>
        /// URL of the chart
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("vizOptions")]
        private InputList<Inputs.SingleValueChartVizOptionsGetArgs>? _vizOptions;

        /// <summary>
        /// Plot-level customization options, associated with a publish statement.
        /// </summary>
        public InputList<Inputs.SingleValueChartVizOptionsGetArgs> VizOptions
        {
            get => _vizOptions ?? (_vizOptions = new InputList<Inputs.SingleValueChartVizOptionsGetArgs>());
            set => _vizOptions = value;
        }

        public SingleValueChartState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SingleValueChartColorScalesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
        /// </summary>
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        /// <summary>
        /// Indicates the lower threshold non-inclusive value for this range.
        /// </summary>
        [Input("gt")]
        public Input<double>? Gt { get; set; }

        /// <summary>
        /// Indicates the lower threshold inclusive value for this range.
        /// </summary>
        [Input("gte")]
        public Input<double>? Gte { get; set; }

        /// <summary>
        /// Indicates the upper threshold non-inculsive value for this range.
        /// </summary>
        [Input("lt")]
        public Input<double>? Lt { get; set; }

        /// <summary>
        /// Indicates the upper threshold inclusive value for this range.
        /// </summary>
        [Input("lte")]
        public Input<double>? Lte { get; set; }

        public SingleValueChartColorScalesArgs()
        {
        }
    }

    public sealed class SingleValueChartColorScalesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
        /// </summary>
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        /// <summary>
        /// Indicates the lower threshold non-inclusive value for this range.
        /// </summary>
        [Input("gt")]
        public Input<double>? Gt { get; set; }

        /// <summary>
        /// Indicates the lower threshold inclusive value for this range.
        /// </summary>
        [Input("gte")]
        public Input<double>? Gte { get; set; }

        /// <summary>
        /// Indicates the upper threshold non-inculsive value for this range.
        /// </summary>
        [Input("lt")]
        public Input<double>? Lt { get; set; }

        /// <summary>
        /// Indicates the upper threshold inclusive value for this range.
        /// </summary>
        [Input("lte")]
        public Input<double>? Lte { get; set; }

        public SingleValueChartColorScalesGetArgs()
        {
        }
    }

    public sealed class SingleValueChartVizOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Label used in the publish statement that displays the plot (metric time series data) you want to customize.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        [Input("valuePrefix")]
        public Input<string>? ValuePrefix { get; set; }

        [Input("valueSuffix")]
        public Input<string>? ValueSuffix { get; set; }

        /// <summary>
        /// A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
        /// * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
        /// </summary>
        [Input("valueUnit")]
        public Input<string>? ValueUnit { get; set; }

        public SingleValueChartVizOptionsArgs()
        {
        }
    }

    public sealed class SingleValueChartVizOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
        /// </summary>
        [Input("color")]
        public Input<string>? Color { get; set; }

        /// <summary>
        /// Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Label used in the publish statement that displays the plot (metric time series data) you want to customize.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        [Input("valuePrefix")]
        public Input<string>? ValuePrefix { get; set; }

        [Input("valueSuffix")]
        public Input<string>? ValueSuffix { get; set; }

        /// <summary>
        /// A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
        /// * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
        /// </summary>
        [Input("valueUnit")]
        public Input<string>? ValueUnit { get; set; }

        public SingleValueChartVizOptionsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SingleValueChartColorScales
    {
        /// <summary>
        /// Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
        /// </summary>
        public readonly string Color;
        /// <summary>
        /// Indicates the lower threshold non-inclusive value for this range.
        /// </summary>
        public readonly double? Gt;
        /// <summary>
        /// Indicates the lower threshold inclusive value for this range.
        /// </summary>
        public readonly double? Gte;
        /// <summary>
        /// Indicates the upper threshold non-inculsive value for this range.
        /// </summary>
        public readonly double? Lt;
        /// <summary>
        /// Indicates the upper threshold inclusive value for this range.
        /// </summary>
        public readonly double? Lte;

        [OutputConstructor]
        private SingleValueChartColorScales(
            string color,
            double? gt,
            double? gte,
            double? lt,
            double? lte)
        {
            Color = color;
            Gt = gt;
            Gte = gte;
            Lt = lt;
            Lte = lte;
        }
    }

    [OutputType]
    public sealed class SingleValueChartVizOptions
    {
        /// <summary>
        /// Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
        /// </summary>
        public readonly string? Color;
        /// <summary>
        /// Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Label used in the publish statement that displays the plot (metric time series data) you want to customize.
        /// </summary>
        public readonly string Label;
        public readonly string? ValuePrefix;
        public readonly string? ValueSuffix;
        /// <summary>
        /// A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
        /// * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
        /// </summary>
        public readonly string? ValueUnit;

        [OutputConstructor]
        private SingleValueChartVizOptions(
            string? color,
            string? displayName,
            string label,
            string? valuePrefix,
            string? valueSuffix,
            string? valueUnit)
        {
            Color = color;
            DisplayName = displayName;
            Label = label;
            ValuePrefix = valuePrefix;
            ValueSuffix = valueSuffix;
            ValueUnit = valueUnit;
        }
    }
    }
}
