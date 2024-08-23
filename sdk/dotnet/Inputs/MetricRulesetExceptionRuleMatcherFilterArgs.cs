// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class MetricRulesetExceptionRuleMatcherFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, this filter will match all values not matching the property_values
        /// </summary>
        [Input("not", required: true)]
        public Input<bool> Not { get; set; } = null!;

        /// <summary>
        /// Name of the dimension
        /// </summary>
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        [Input("propertyValues", required: true)]
        private InputList<string>? _propertyValues;

        /// <summary>
        /// Value of the dimension
        /// </summary>
        public InputList<string> PropertyValues
        {
            get => _propertyValues ?? (_propertyValues = new InputList<string>());
            set => _propertyValues = value;
        }

        public MetricRulesetExceptionRuleMatcherFilterArgs()
        {
        }
        public static new MetricRulesetExceptionRuleMatcherFilterArgs Empty => new MetricRulesetExceptionRuleMatcherFilterArgs();
    }
}
