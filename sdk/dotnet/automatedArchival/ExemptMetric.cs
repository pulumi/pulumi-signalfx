// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.automatedArchival
{
    [SignalFxResourceType("signalfx:automatedarchival/exemptMetric:ExemptMetric")]
    public partial class ExemptMetric : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of metrics to be exempted from automated archival
        /// </summary>
        [Output("exemptMetrics")]
        public Output<ImmutableArray<Outputs.ExemptMetricExemptMetric>> ExemptMetrics { get; private set; } = null!;


        /// <summary>
        /// Create a ExemptMetric resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExemptMetric(string name, ExemptMetricArgs args, CustomResourceOptions? options = null)
            : base("signalfx:automatedarchival/exemptMetric:ExemptMetric", name, args ?? new ExemptMetricArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExemptMetric(string name, Input<string> id, ExemptMetricState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:automatedarchival/exemptMetric:ExemptMetric", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ExemptMetric resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExemptMetric Get(string name, Input<string> id, ExemptMetricState? state = null, CustomResourceOptions? options = null)
        {
            return new ExemptMetric(name, id, state, options);
        }
    }

    public sealed class ExemptMetricArgs : global::Pulumi.ResourceArgs
    {
        [Input("exemptMetrics", required: true)]
        private InputList<Inputs.ExemptMetricExemptMetricArgs>? _exemptMetrics;

        /// <summary>
        /// List of metrics to be exempted from automated archival
        /// </summary>
        public InputList<Inputs.ExemptMetricExemptMetricArgs> ExemptMetrics
        {
            get => _exemptMetrics ?? (_exemptMetrics = new InputList<Inputs.ExemptMetricExemptMetricArgs>());
            set => _exemptMetrics = value;
        }

        public ExemptMetricArgs()
        {
        }
        public static new ExemptMetricArgs Empty => new ExemptMetricArgs();
    }

    public sealed class ExemptMetricState : global::Pulumi.ResourceArgs
    {
        [Input("exemptMetrics")]
        private InputList<Inputs.ExemptMetricExemptMetricGetArgs>? _exemptMetrics;

        /// <summary>
        /// List of metrics to be exempted from automated archival
        /// </summary>
        public InputList<Inputs.ExemptMetricExemptMetricGetArgs> ExemptMetrics
        {
            get => _exemptMetrics ?? (_exemptMetrics = new InputList<Inputs.ExemptMetricExemptMetricGetArgs>());
            set => _exemptMetrics = value;
        }

        public ExemptMetricState()
        {
        }
        public static new ExemptMetricState Empty => new ExemptMetricState();
    }
}
