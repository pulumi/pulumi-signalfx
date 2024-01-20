// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DashboardSelectedEventOverlaySourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        [Input("property", required: true)]
        public Input<string> Property { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public DashboardSelectedEventOverlaySourceArgs()
        {
        }
        public static new DashboardSelectedEventOverlaySourceArgs Empty => new DashboardSelectedEventOverlaySourceArgs();
    }
}
