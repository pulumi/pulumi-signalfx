// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardGroupDashboardVariableOverrideGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        [Input("valuesSuggesteds")]
        private InputList<string>? _valuesSuggesteds;
        public InputList<string> ValuesSuggesteds
        {
            get => _valuesSuggesteds ?? (_valuesSuggesteds = new InputList<string>());
            set => _valuesSuggesteds = value;
        }

        public DashboardGroupDashboardVariableOverrideGetArgs()
        {
        }
        public static new DashboardGroupDashboardVariableOverrideGetArgs Empty => new DashboardGroupDashboardVariableOverrideGetArgs();
    }
}
