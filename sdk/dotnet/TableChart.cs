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
    /// This special type of chart displays a data table. This table can be grouped by a dimension.
    /// 
    /// ## Example
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using SignalFx = Pulumi.SignalFx;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // signalfx_list_chart.Logs-Exec_0:
    ///     var table0 = new SignalFx.TableChart("table0", new()
    ///     {
    ///         Description = "beep",
    ///         DisableSampling = false,
    ///         GroupBies = new[]
    ///         {
    ///             "ClusterName",
    ///         },
    ///         MaxDelay = 0,
    ///         ProgramText = "A = data('cpu.usage.total').publish(label='CPU Total')",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Arguments
    /// 
    /// The following arguments are supported in the resource block:
    /// 
    /// * `name` - (Required) Name of the table chart.
    /// * `program_text` - (Required) The SignalFlow for your Data Table Chart
    /// * `description` - (Optional) Description of the table chart.
    /// * `group_by` - (Optional) Dimension to group by
    /// 
    /// ## Attributes
    /// 
    /// In a addition to all arguments above, the following attributes are exported:
    /// 
    /// * `id` - The ID of the chart.
    /// * `url` - The URL of the chart.
    /// </summary>
    [SignalFxResourceType("signalfx:index/tableChart:TableChart")]
    public partial class TableChart : global::Pulumi.CustomResource
    {
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
        /// Properties to group by in the Table (in nesting order)
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
        /// How often (in seconds) to refresh the values of the Table
        /// </summary>
        [Output("refreshInterval")]
        public Output<int?> RefreshInterval { get; private set; } = null!;

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
        /// Plot-level customization options, associated with a publish statement
        /// </summary>
        [Output("vizOptions")]
        public Output<ImmutableArray<Outputs.TableChartVizOption>> VizOptions { get; private set; } = null!;


        /// <summary>
        /// Create a TableChart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TableChart(string name, TableChartArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/tableChart:TableChart", name, args ?? new TableChartArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TableChart(string name, Input<string> id, TableChartState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/tableChart:TableChart", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TableChart resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TableChart Get(string name, Input<string> id, TableChartState? state = null, CustomResourceOptions? options = null)
        {
            return new TableChart(name, id, state, options);
        }
    }

    public sealed class TableChartArgs : global::Pulumi.ResourceArgs
    {
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
        /// Properties to group by in the Table (in nesting order)
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
        /// How often (in seconds) to refresh the values of the Table
        /// </summary>
        [Input("refreshInterval")]
        public Input<int>? RefreshInterval { get; set; }

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

        [Input("vizOptions")]
        private InputList<Inputs.TableChartVizOptionArgs>? _vizOptions;

        /// <summary>
        /// Plot-level customization options, associated with a publish statement
        /// </summary>
        public InputList<Inputs.TableChartVizOptionArgs> VizOptions
        {
            get => _vizOptions ?? (_vizOptions = new InputList<Inputs.TableChartVizOptionArgs>());
            set => _vizOptions = value;
        }

        public TableChartArgs()
        {
        }
        public static new TableChartArgs Empty => new TableChartArgs();
    }

    public sealed class TableChartState : global::Pulumi.ResourceArgs
    {
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
        /// Properties to group by in the Table (in nesting order)
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
        /// How often (in seconds) to refresh the values of the Table
        /// </summary>
        [Input("refreshInterval")]
        public Input<int>? RefreshInterval { get; set; }

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

        [Input("vizOptions")]
        private InputList<Inputs.TableChartVizOptionGetArgs>? _vizOptions;

        /// <summary>
        /// Plot-level customization options, associated with a publish statement
        /// </summary>
        public InputList<Inputs.TableChartVizOptionGetArgs> VizOptions
        {
            get => _vizOptions ?? (_vizOptions = new InputList<Inputs.TableChartVizOptionGetArgs>());
            set => _vizOptions = value;
        }

        public TableChartState()
        {
        }
        public static new TableChartState Empty => new TableChartState();
    }
}
