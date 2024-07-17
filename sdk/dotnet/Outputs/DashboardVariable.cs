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
    public sealed class DashboardVariable
    {
        /// <summary>
        /// An alias for the dashboard variable. This text will appear as the label for the dropdown field on the dashboard.
        /// </summary>
        public readonly string Alias;
        /// <summary>
        /// If true, this variable will also match data that doesn't have this property at all.
        /// </summary>
        public readonly bool? ApplyIfExist;
        /// <summary>
        /// Variable description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A metric time series dimension or property name.
        /// </summary>
        public readonly string Property;
        /// <summary>
        /// If `true`, this variable will only apply to charts that have a filter for the property.
        /// </summary>
        public readonly bool? ReplaceOnly;
        /// <summary>
        /// If `true`, this variable may only be set to the values listed in `values_suggested` and only these values will appear in autosuggestion menus. `false` by default.
        /// </summary>
        public readonly bool? RestrictedSuggestions;
        /// <summary>
        /// Determines whether a value is required for this variable (and therefore whether it will be possible to view this dashboard without this filter applied). `false` by default.
        /// </summary>
        public readonly bool? ValueRequired;
        /// <summary>
        /// List of of strings (which will be treated as an OR filter on the property).
        /// </summary>
        public readonly ImmutableArray<string> Values;
        /// <summary>
        /// A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
        /// </summary>
        public readonly ImmutableArray<string> ValuesSuggesteds;

        [OutputConstructor]
        private DashboardVariable(
            string alias,

            bool? applyIfExist,

            string? description,

            string property,

            bool? replaceOnly,

            bool? restrictedSuggestions,

            bool? valueRequired,

            ImmutableArray<string> values,

            ImmutableArray<string> valuesSuggesteds)
        {
            Alias = alias;
            ApplyIfExist = applyIfExist;
            Description = description;
            Property = property;
            ReplaceOnly = replaceOnly;
            RestrictedSuggestions = restrictedSuggestions;
            ValueRequired = valueRequired;
            Values = values;
            ValuesSuggesteds = valuesSuggesteds;
        }
    }
}