// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardGroupDashboardVariableOverrideGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A metric time series dimension or property name.
        /// </summary>
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// (Optional) List of of strings (which will be treated as an OR filter on the property).
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        [Input("valuesSuggesteds")]
        private InputList<string>? _valuesSuggesteds;

        /// <summary>
        /// A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
        /// </summary>
        public InputList<string> ValuesSuggesteds
        {
            get => _valuesSuggesteds ?? (_valuesSuggesteds = new InputList<string>());
            set => _valuesSuggesteds = value;
        }

        public DashboardGroupDashboardVariableOverrideGetArgs()
        {
        }
    }
}
