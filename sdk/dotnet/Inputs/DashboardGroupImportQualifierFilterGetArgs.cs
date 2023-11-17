// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardGroupImportQualifierFilterGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
        /// </summary>
        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        /// <summary>
        /// The name of a dimension to filter against.
        /// </summary>
        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// A list of values to be used with the `property`, they will be combined via `OR`.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public DashboardGroupImportQualifierFilterGetArgs()
        {
        }
        public static new DashboardGroupImportQualifierFilterGetArgs Empty => new DashboardGroupImportQualifierFilterGetArgs();
    }
}
