// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Logs.Inputs
{

    public sealed class ViewColumnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the log view.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ViewColumnArgs()
        {
        }
        public static new ViewColumnArgs Empty => new ViewColumnArgs();
    }
}
