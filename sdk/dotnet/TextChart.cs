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
    /// This special type of chart doesn’t display any metric data. Rather, it lets you place a text note on the dashboard.
    /// 
    /// ## Example
    /// </summary>
    [SignalFxResourceType("signalfx:index/textChart:TextChart")]
    public partial class TextChart : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of the text note.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Markdown text to display.
        /// </summary>
        [Output("markdown")]
        public Output<string> Markdown { get; private set; } = null!;

        /// <summary>
        /// Name of the text note.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The URL of the chart.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a TextChart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TextChart(string name, TextChartArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/textChart:TextChart", name, args ?? new TextChartArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TextChart(string name, Input<string> id, TextChartState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/textChart:TextChart", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TextChart resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TextChart Get(string name, Input<string> id, TextChartState? state = null, CustomResourceOptions? options = null)
        {
            return new TextChart(name, id, state, options);
        }
    }

    public sealed class TextChartArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the text note.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Markdown text to display.
        /// </summary>
        [Input("markdown", required: true)]
        public Input<string> Markdown { get; set; } = null!;

        /// <summary>
        /// Name of the text note.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TextChartArgs()
        {
        }
        public static new TextChartArgs Empty => new TextChartArgs();
    }

    public sealed class TextChartState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the text note.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Markdown text to display.
        /// </summary>
        [Input("markdown")]
        public Input<string>? Markdown { get; set; }

        /// <summary>
        /// Name of the text note.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The URL of the chart.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public TextChartState()
        {
        }
        public static new TextChartState Empty => new TextChartState();
    }
}
