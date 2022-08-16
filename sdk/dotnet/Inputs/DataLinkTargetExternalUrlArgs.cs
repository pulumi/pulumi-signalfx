// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Inputs
{

    public sealed class DataLinkTargetExternalUrlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [minimum time window](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) for a search sent to an external site. Defaults to `6000`
        /// </summary>
        [Input("minimumTimeWindow")]
        public Input<string>? MinimumTimeWindow { get; set; }

        /// <summary>
        /// User-assigned target name. Use this value to differentiate between the link targets for a data link object.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("propertyKeyMapping")]
        private InputMap<string>? _propertyKeyMapping;

        /// <summary>
        /// Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.
        /// </summary>
        public InputMap<string> PropertyKeyMapping
        {
            get => _propertyKeyMapping ?? (_propertyKeyMapping = new InputMap<string>());
            set => _propertyKeyMapping = value;
        }

        /// <summary>
        /// [Designates the format](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) of `minimum_time_window` in the same data link target object. Must be one of `"ISO8601"`, `"EpochSeconds"` or `"Epoch"` (which is milliseconds). Defaults to `"ISO8601"`.
        /// </summary>
        [Input("timeFormat")]
        public Input<string>? TimeFormat { get; set; }

        /// <summary>
        /// URL string for a Splunk instance or external system data link target. [See the supported template variables](https://developers.signalfx.com/administration/data_links_overview.html#_external_link_targets).
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public DataLinkTargetExternalUrlArgs()
        {
        }
        public static new DataLinkTargetExternalUrlArgs Empty => new DataLinkTargetExternalUrlArgs();
    }
}
