// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx
{
    /// <summary>
    /// This chart type displays an overview of your SLO and can give more specific insights into your SLI performance using different filter and customized time ranges.
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
    ///     var myslochart0 = new SignalFx.SloChart("myslochart0", new()
    ///     {
    ///         SloId = "GbOHXbSAEAA",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:index/sloChart:SloChart")]
    public partial class SloChart : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of SLO object.
        /// </summary>
        [Output("sloId")]
        public Output<string> SloId { get; private set; } = null!;

        /// <summary>
        /// The URL of the chart.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a SloChart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SloChart(string name, SloChartArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/sloChart:SloChart", name, args ?? new SloChartArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SloChart(string name, Input<string> id, SloChartState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/sloChart:SloChart", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SloChart resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SloChart Get(string name, Input<string> id, SloChartState? state = null, CustomResourceOptions? options = null)
        {
            return new SloChart(name, id, state, options);
        }
    }

    public sealed class SloChartArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of SLO object.
        /// </summary>
        [Input("sloId", required: true)]
        public Input<string> SloId { get; set; } = null!;

        public SloChartArgs()
        {
        }
        public static new SloChartArgs Empty => new SloChartArgs();
    }

    public sealed class SloChartState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of SLO object.
        /// </summary>
        [Input("sloId")]
        public Input<string>? SloId { get; set; }

        /// <summary>
        /// The URL of the chart.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public SloChartState()
        {
        }
        public static new SloChartState Empty => new SloChartState();
    }
}
