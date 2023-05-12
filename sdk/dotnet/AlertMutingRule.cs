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
    /// Provides an Observability Cloud resource for managing alert muting rules. See [Mute Notifications](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html) for more information.
    /// 
    /// &gt; **WARNING** Observability Cloud does not allow the start time of a **currently active** muting rule to be modified. As such, attempting to modify a currently active rule will destroy the existing rule and create a new rule. This may result in the emission of notifications.
    /// 
    /// &gt; **WARNING** Observability Cloud currently allows linking alert muting rule with only one detector ID. Specifying multiple detector IDs will make the muting rule obsolete.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using SignalFx = Pulumi.SignalFx;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var roolMooterOne = new SignalFx.AlertMutingRule("roolMooterOne", new()
    ///     {
    ///         Description = "mooted it NEW",
    ///         StartTime = 1573063243,
    ///         StopTime = 0,
    ///         Detectors = new[]
    ///         {
    ///             signalfx_detector.Some_detector_id,
    ///         },
    ///         Filters = new[]
    ///         {
    ///             new SignalFx.Inputs.AlertMutingRuleFilterArgs
    ///             {
    ///                 Property = "foo",
    ///                 PropertyValue = "bar",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:index/alertMutingRule:AlertMutingRule")]
    public partial class AlertMutingRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description for this muting rule
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
        /// </summary>
        [Output("detectors")]
        public Output<ImmutableArray<string>> Detectors { get; private set; } = null!;

        [Output("effectiveStartTime")]
        public Output<int> EffectiveStartTime { get; private set; } = null!;

        /// <summary>
        /// Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
        /// </summary>
        [Output("filters")]
        public Output<ImmutableArray<Outputs.AlertMutingRuleFilter>> Filters { get; private set; } = null!;

        /// <summary>
        /// Starting time of an alert muting rule as a Unit time stamp in seconds.
        /// </summary>
        [Output("startTime")]
        public Output<int> StartTime { get; private set; } = null!;

        /// <summary>
        /// Stop time of an alert muting rule as a Unix time stamp in seconds.
        /// </summary>
        [Output("stopTime")]
        public Output<int?> StopTime { get; private set; } = null!;


        /// <summary>
        /// Create a AlertMutingRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertMutingRule(string name, AlertMutingRuleArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/alertMutingRule:AlertMutingRule", name, args ?? new AlertMutingRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlertMutingRule(string name, Input<string> id, AlertMutingRuleState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/alertMutingRule:AlertMutingRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlertMutingRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertMutingRule Get(string name, Input<string> id, AlertMutingRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new AlertMutingRule(name, id, state, options);
        }
    }

    public sealed class AlertMutingRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description for this muting rule
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("detectors")]
        private InputList<string>? _detectors;

        /// <summary>
        /// A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
        /// </summary>
        public InputList<string> Detectors
        {
            get => _detectors ?? (_detectors = new InputList<string>());
            set => _detectors = value;
        }

        [Input("filters")]
        private InputList<Inputs.AlertMutingRuleFilterArgs>? _filters;

        /// <summary>
        /// Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
        /// </summary>
        public InputList<Inputs.AlertMutingRuleFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.AlertMutingRuleFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Starting time of an alert muting rule as a Unit time stamp in seconds.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<int> StartTime { get; set; } = null!;

        /// <summary>
        /// Stop time of an alert muting rule as a Unix time stamp in seconds.
        /// </summary>
        [Input("stopTime")]
        public Input<int>? StopTime { get; set; }

        public AlertMutingRuleArgs()
        {
        }
        public static new AlertMutingRuleArgs Empty => new AlertMutingRuleArgs();
    }

    public sealed class AlertMutingRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description for this muting rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("detectors")]
        private InputList<string>? _detectors;

        /// <summary>
        /// A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
        /// </summary>
        public InputList<string> Detectors
        {
            get => _detectors ?? (_detectors = new InputList<string>());
            set => _detectors = value;
        }

        [Input("effectiveStartTime")]
        public Input<int>? EffectiveStartTime { get; set; }

        [Input("filters")]
        private InputList<Inputs.AlertMutingRuleFilterGetArgs>? _filters;

        /// <summary>
        /// Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
        /// </summary>
        public InputList<Inputs.AlertMutingRuleFilterGetArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.AlertMutingRuleFilterGetArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// Starting time of an alert muting rule as a Unit time stamp in seconds.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// Stop time of an alert muting rule as a Unix time stamp in seconds.
        /// </summary>
        [Input("stopTime")]
        public Input<int>? StopTime { get; set; }

        public AlertMutingRuleState()
        {
        }
        public static new AlertMutingRuleState Empty => new AlertMutingRuleState();
    }
}
