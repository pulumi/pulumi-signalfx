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
    /// Provides an Observability Cloud resource for managing metric rulesets.
    /// 
    /// &gt; **NOTE** When managing metric rulesets to drop data use a session token for an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
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
    ///     var cpuUtilizationMetricRuleset = new SignalFx.MetricRuleset("cpu_utilization_metric_ruleset", new()
    ///     {
    ///         MetricName = "cpu.utilization",
    ///         Description = "Routing ruleset for cpu.utilization",
    ///         AggregationRules = new[]
    ///         {
    ///             new SignalFx.Inputs.MetricRulesetAggregationRuleArgs
    ///             {
    ///                 Name = "cpu.utilization by service rule",
    ///                 Description = "Aggregates cpu.utilization data by service",
    ///                 Enabled = true,
    ///                 Matchers = new[]
    ///                 {
    ///                     new SignalFx.Inputs.MetricRulesetAggregationRuleMatcherArgs
    ///                     {
    ///                         Type = "dimension",
    ///                         Filters = new[]
    ///                         {
    ///                             new SignalFx.Inputs.MetricRulesetAggregationRuleMatcherFilterArgs
    ///                             {
    ///                                 Property = "realm",
    ///                                 PropertyValues = new[]
    ///                                 {
    ///                                     "us-east-1",
    ///                                 },
    ///                                 Not = false,
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///                 Aggregators = new[]
    ///                 {
    ///                     new SignalFx.Inputs.MetricRulesetAggregationRuleAggregatorArgs
    ///                     {
    ///                         Type = "rollup",
    ///                         Dimensions = new[]
    ///                         {
    ///                             "service",
    ///                         },
    ///                         DropDimensions = false,
    ///                         OutputName = "cpu.utilization.by.service.agg",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         ExceptionRules = new[]
    ///         {
    ///             new SignalFx.Inputs.MetricRulesetExceptionRuleArgs
    ///             {
    ///                 Name = "Exception rule us-east-2",
    ///                 Description = "Routes us-east-2 data to real-time",
    ///                 Enabled = true,
    ///                 Matchers = new[]
    ///                 {
    ///                     new SignalFx.Inputs.MetricRulesetExceptionRuleMatcherArgs
    ///                     {
    ///                         Type = "dimension",
    ///                         Filters = new[]
    ///                         {
    ///                             new SignalFx.Inputs.MetricRulesetExceptionRuleMatcherFilterArgs
    ///                             {
    ///                                 Property = "realm",
    ///                                 PropertyValues = new[]
    ///                                 {
    ///                                     "us-east-2",
    ///                                 },
    ///                                 Not = false,
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         RoutingRules = new[]
    ///         {
    ///             new SignalFx.Inputs.MetricRulesetRoutingRuleArgs
    ///             {
    ///                 Destination = "Archived",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:index/metricRuleset:MetricRuleset")]
    public partial class MetricRuleset : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of aggregation rules for the metric
        /// </summary>
        [Output("aggregationRules")]
        public Output<ImmutableArray<Outputs.MetricRulesetAggregationRule>> AggregationRules { get; private set; } = null!;

        /// <summary>
        /// Timestamp of when the metric ruleset was created
        /// </summary>
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        /// <summary>
        /// ID of the creator of the metric ruleset
        /// </summary>
        [Output("creator")]
        public Output<string> Creator { get; private set; } = null!;

        /// <summary>
        /// Information about the metric ruleset
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of exception rules for the metric
        /// </summary>
        [Output("exceptionRules")]
        public Output<ImmutableArray<Outputs.MetricRulesetExceptionRule>> ExceptionRules { get; private set; } = null!;

        /// <summary>
        /// Timestamp of when the metric ruleset was last updated
        /// </summary>
        [Output("lastUpdated")]
        public Output<string> LastUpdated { get; private set; } = null!;

        /// <summary>
        /// ID of user who last updated the metric ruleset
        /// </summary>
        [Output("lastUpdatedBy")]
        public Output<string> LastUpdatedBy { get; private set; } = null!;

        /// <summary>
        /// Name of user who last updated this metric ruleset
        /// </summary>
        [Output("lastUpdatedByName")]
        public Output<string> LastUpdatedByName { get; private set; } = null!;

        /// <summary>
        /// Name of the input metric
        /// </summary>
        [Output("metricName")]
        public Output<string> MetricName { get; private set; } = null!;

        /// <summary>
        /// Routing Rule object
        /// </summary>
        [Output("routingRules")]
        public Output<ImmutableArray<Outputs.MetricRulesetRoutingRule>> RoutingRules { get; private set; } = null!;

        /// <summary>
        /// Version of the ruleset
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a MetricRuleset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MetricRuleset(string name, MetricRulesetArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/metricRuleset:MetricRuleset", name, args ?? new MetricRulesetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MetricRuleset(string name, Input<string> id, MetricRulesetState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/metricRuleset:MetricRuleset", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MetricRuleset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MetricRuleset Get(string name, Input<string> id, MetricRulesetState? state = null, CustomResourceOptions? options = null)
        {
            return new MetricRuleset(name, id, state, options);
        }
    }

    public sealed class MetricRulesetArgs : global::Pulumi.ResourceArgs
    {
        [Input("aggregationRules")]
        private InputList<Inputs.MetricRulesetAggregationRuleArgs>? _aggregationRules;

        /// <summary>
        /// List of aggregation rules for the metric
        /// </summary>
        public InputList<Inputs.MetricRulesetAggregationRuleArgs> AggregationRules
        {
            get => _aggregationRules ?? (_aggregationRules = new InputList<Inputs.MetricRulesetAggregationRuleArgs>());
            set => _aggregationRules = value;
        }

        /// <summary>
        /// Information about the metric ruleset
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("exceptionRules")]
        private InputList<Inputs.MetricRulesetExceptionRuleArgs>? _exceptionRules;

        /// <summary>
        /// List of exception rules for the metric
        /// </summary>
        public InputList<Inputs.MetricRulesetExceptionRuleArgs> ExceptionRules
        {
            get => _exceptionRules ?? (_exceptionRules = new InputList<Inputs.MetricRulesetExceptionRuleArgs>());
            set => _exceptionRules = value;
        }

        /// <summary>
        /// Name of the input metric
        /// </summary>
        [Input("metricName", required: true)]
        public Input<string> MetricName { get; set; } = null!;

        [Input("routingRules", required: true)]
        private InputList<Inputs.MetricRulesetRoutingRuleArgs>? _routingRules;

        /// <summary>
        /// Routing Rule object
        /// </summary>
        public InputList<Inputs.MetricRulesetRoutingRuleArgs> RoutingRules
        {
            get => _routingRules ?? (_routingRules = new InputList<Inputs.MetricRulesetRoutingRuleArgs>());
            set => _routingRules = value;
        }

        public MetricRulesetArgs()
        {
        }
        public static new MetricRulesetArgs Empty => new MetricRulesetArgs();
    }

    public sealed class MetricRulesetState : global::Pulumi.ResourceArgs
    {
        [Input("aggregationRules")]
        private InputList<Inputs.MetricRulesetAggregationRuleGetArgs>? _aggregationRules;

        /// <summary>
        /// List of aggregation rules for the metric
        /// </summary>
        public InputList<Inputs.MetricRulesetAggregationRuleGetArgs> AggregationRules
        {
            get => _aggregationRules ?? (_aggregationRules = new InputList<Inputs.MetricRulesetAggregationRuleGetArgs>());
            set => _aggregationRules = value;
        }

        /// <summary>
        /// Timestamp of when the metric ruleset was created
        /// </summary>
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// ID of the creator of the metric ruleset
        /// </summary>
        [Input("creator")]
        public Input<string>? Creator { get; set; }

        /// <summary>
        /// Information about the metric ruleset
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("exceptionRules")]
        private InputList<Inputs.MetricRulesetExceptionRuleGetArgs>? _exceptionRules;

        /// <summary>
        /// List of exception rules for the metric
        /// </summary>
        public InputList<Inputs.MetricRulesetExceptionRuleGetArgs> ExceptionRules
        {
            get => _exceptionRules ?? (_exceptionRules = new InputList<Inputs.MetricRulesetExceptionRuleGetArgs>());
            set => _exceptionRules = value;
        }

        /// <summary>
        /// Timestamp of when the metric ruleset was last updated
        /// </summary>
        [Input("lastUpdated")]
        public Input<string>? LastUpdated { get; set; }

        /// <summary>
        /// ID of user who last updated the metric ruleset
        /// </summary>
        [Input("lastUpdatedBy")]
        public Input<string>? LastUpdatedBy { get; set; }

        /// <summary>
        /// Name of user who last updated this metric ruleset
        /// </summary>
        [Input("lastUpdatedByName")]
        public Input<string>? LastUpdatedByName { get; set; }

        /// <summary>
        /// Name of the input metric
        /// </summary>
        [Input("metricName")]
        public Input<string>? MetricName { get; set; }

        [Input("routingRules")]
        private InputList<Inputs.MetricRulesetRoutingRuleGetArgs>? _routingRules;

        /// <summary>
        /// Routing Rule object
        /// </summary>
        public InputList<Inputs.MetricRulesetRoutingRuleGetArgs> RoutingRules
        {
            get => _routingRules ?? (_routingRules = new InputList<Inputs.MetricRulesetRoutingRuleGetArgs>());
            set => _routingRules = value;
        }

        /// <summary>
        /// Version of the ruleset
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public MetricRulesetState()
        {
        }
        public static new MetricRulesetState Empty => new MetricRulesetState();
    }
}
