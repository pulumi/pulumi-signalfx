// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Logs.Inputs
{

    public sealed class ViewSortOptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("descending", required: true)]
        public Input<bool> Descending { get; set; } = null!;

        [Input("field", required: true)]
        public Input<string> Field { get; set; } = null!;

        public ViewSortOptionArgs()
        {
        }
        public static new ViewSortOptionArgs Empty => new ViewSortOptionArgs();
    }
}
