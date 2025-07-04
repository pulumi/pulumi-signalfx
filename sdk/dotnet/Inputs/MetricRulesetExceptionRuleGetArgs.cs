// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class MetricRulesetExceptionRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information about an exception rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When false, this rule will not route matched data to real-time
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("matchers", required: true)]
        private InputList<Inputs.MetricRulesetExceptionRuleMatcherGetArgs>? _matchers;

        /// <summary>
        /// Matcher object
        /// </summary>
        public InputList<Inputs.MetricRulesetExceptionRuleMatcherGetArgs> Matchers
        {
            get => _matchers ?? (_matchers = new InputList<Inputs.MetricRulesetExceptionRuleMatcherGetArgs>());
            set => _matchers = value;
        }

        /// <summary>
        /// name of the exception rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("restorations")]
        private InputList<Inputs.MetricRulesetExceptionRuleRestorationGetArgs>? _restorations;

        /// <summary>
        /// Properties of a restoration job
        /// </summary>
        public InputList<Inputs.MetricRulesetExceptionRuleRestorationGetArgs> Restorations
        {
            get => _restorations ?? (_restorations = new InputList<Inputs.MetricRulesetExceptionRuleRestorationGetArgs>());
            set => _restorations = value;
        }

        public MetricRulesetExceptionRuleGetArgs()
        {
        }
        public static new MetricRulesetExceptionRuleGetArgs Empty => new MetricRulesetExceptionRuleGetArgs();
    }
}
