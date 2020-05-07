// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Outputs
{

    [OutputType]
    public sealed class ListChartLegendOptionsField
    {
        /// <summary>
        /// True or False depending on if you want the property to be shown or hidden.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The name of the property to display. Note the special values of `sf_metric` (corresponding with the API's `Plot Name`) which shows the label of the time series `publish()` and `sf_originatingMetric` (corresponding with the API's `metric (sf metric)`) that shows the [name of the metric](https://developers.signalfx.com/signalflow_analytics/functions/data_function.html#table-1-parameter-definitions) for the time series being displayed.
        /// </summary>
        public readonly string Property;

        [OutputConstructor]
        private ListChartLegendOptionsField(
            bool? enabled,

            string property)
        {
            Enabled = enabled;
            Property = property;
        }
    }
}