// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardVariableArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardVariableArgs Empty = new DashboardVariableArgs();

    /**
     * An alias for the dashboard variable. This text will appear as the label for the dropdown field on the dashboard.
     * 
     */
    @Import(name="alias", required=true)
    private Output<String> alias;

    /**
     * @return An alias for the dashboard variable. This text will appear as the label for the dropdown field on the dashboard.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }

    /**
     * If true, this variable will also match data that doesn&#39;t have this property at all.
     * 
     */
    @Import(name="applyIfExist")
    private @Nullable Output<Boolean> applyIfExist;

    /**
     * @return If true, this variable will also match data that doesn&#39;t have this property at all.
     * 
     */
    public Optional<Output<Boolean>> applyIfExist() {
        return Optional.ofNullable(this.applyIfExist);
    }

    /**
     * Variable description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Variable description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of a dimension to filter against.
     * 
     */
    @Import(name="property", required=true)
    private Output<String> property;

    /**
     * @return The name of a dimension to filter against.
     * 
     */
    public Output<String> property() {
        return this.property;
    }

    /**
     * If `true`, this variable will only apply to charts that have a filter for the property.
     * 
     */
    @Import(name="replaceOnly")
    private @Nullable Output<Boolean> replaceOnly;

    /**
     * @return If `true`, this variable will only apply to charts that have a filter for the property.
     * 
     */
    public Optional<Output<Boolean>> replaceOnly() {
        return Optional.ofNullable(this.replaceOnly);
    }

    /**
     * If `true`, this variable may only be set to the values listed in `values_suggested` and only these values will appear in autosuggestion menus. `false` by default.
     * 
     */
    @Import(name="restrictedSuggestions")
    private @Nullable Output<Boolean> restrictedSuggestions;

    /**
     * @return If `true`, this variable may only be set to the values listed in `values_suggested` and only these values will appear in autosuggestion menus. `false` by default.
     * 
     */
    public Optional<Output<Boolean>> restrictedSuggestions() {
        return Optional.ofNullable(this.restrictedSuggestions);
    }

    /**
     * Determines whether a value is required for this variable (and therefore whether it will be possible to view this dashboard without this filter applied). `false` by default.
     * 
     */
    @Import(name="valueRequired")
    private @Nullable Output<Boolean> valueRequired;

    /**
     * @return Determines whether a value is required for this variable (and therefore whether it will be possible to view this dashboard without this filter applied). `false` by default.
     * 
     */
    public Optional<Output<Boolean>> valueRequired() {
        return Optional.ofNullable(this.valueRequired);
    }

    /**
     * A list of values to be used with the `property`, they will be combined via `OR`.
     * 
     */
    @Import(name="values")
    private @Nullable Output<List<String>> values;

    /**
     * @return A list of values to be used with the `property`, they will be combined via `OR`.
     * 
     */
    public Optional<Output<List<String>>> values() {
        return Optional.ofNullable(this.values);
    }

    /**
     * A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
     * 
     */
    @Import(name="valuesSuggesteds")
    private @Nullable Output<List<String>> valuesSuggesteds;

    /**
     * @return A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
     * 
     */
    public Optional<Output<List<String>>> valuesSuggesteds() {
        return Optional.ofNullable(this.valuesSuggesteds);
    }

    private DashboardVariableArgs() {}

    private DashboardVariableArgs(DashboardVariableArgs $) {
        this.alias = $.alias;
        this.applyIfExist = $.applyIfExist;
        this.description = $.description;
        this.property = $.property;
        this.replaceOnly = $.replaceOnly;
        this.restrictedSuggestions = $.restrictedSuggestions;
        this.valueRequired = $.valueRequired;
        this.values = $.values;
        this.valuesSuggesteds = $.valuesSuggesteds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardVariableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardVariableArgs $;

        public Builder() {
            $ = new DashboardVariableArgs();
        }

        public Builder(DashboardVariableArgs defaults) {
            $ = new DashboardVariableArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alias An alias for the dashboard variable. This text will appear as the label for the dropdown field on the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder alias(Output<String> alias) {
            $.alias = alias;
            return this;
        }

        /**
         * @param alias An alias for the dashboard variable. This text will appear as the label for the dropdown field on the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        /**
         * @param applyIfExist If true, this variable will also match data that doesn&#39;t have this property at all.
         * 
         * @return builder
         * 
         */
        public Builder applyIfExist(@Nullable Output<Boolean> applyIfExist) {
            $.applyIfExist = applyIfExist;
            return this;
        }

        /**
         * @param applyIfExist If true, this variable will also match data that doesn&#39;t have this property at all.
         * 
         * @return builder
         * 
         */
        public Builder applyIfExist(Boolean applyIfExist) {
            return applyIfExist(Output.of(applyIfExist));
        }

        /**
         * @param description Variable description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Variable description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param property The name of a dimension to filter against.
         * 
         * @return builder
         * 
         */
        public Builder property(Output<String> property) {
            $.property = property;
            return this;
        }

        /**
         * @param property The name of a dimension to filter against.
         * 
         * @return builder
         * 
         */
        public Builder property(String property) {
            return property(Output.of(property));
        }

        /**
         * @param replaceOnly If `true`, this variable will only apply to charts that have a filter for the property.
         * 
         * @return builder
         * 
         */
        public Builder replaceOnly(@Nullable Output<Boolean> replaceOnly) {
            $.replaceOnly = replaceOnly;
            return this;
        }

        /**
         * @param replaceOnly If `true`, this variable will only apply to charts that have a filter for the property.
         * 
         * @return builder
         * 
         */
        public Builder replaceOnly(Boolean replaceOnly) {
            return replaceOnly(Output.of(replaceOnly));
        }

        /**
         * @param restrictedSuggestions If `true`, this variable may only be set to the values listed in `values_suggested` and only these values will appear in autosuggestion menus. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder restrictedSuggestions(@Nullable Output<Boolean> restrictedSuggestions) {
            $.restrictedSuggestions = restrictedSuggestions;
            return this;
        }

        /**
         * @param restrictedSuggestions If `true`, this variable may only be set to the values listed in `values_suggested` and only these values will appear in autosuggestion menus. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder restrictedSuggestions(Boolean restrictedSuggestions) {
            return restrictedSuggestions(Output.of(restrictedSuggestions));
        }

        /**
         * @param valueRequired Determines whether a value is required for this variable (and therefore whether it will be possible to view this dashboard without this filter applied). `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder valueRequired(@Nullable Output<Boolean> valueRequired) {
            $.valueRequired = valueRequired;
            return this;
        }

        /**
         * @param valueRequired Determines whether a value is required for this variable (and therefore whether it will be possible to view this dashboard without this filter applied). `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder valueRequired(Boolean valueRequired) {
            return valueRequired(Output.of(valueRequired));
        }

        /**
         * @param values A list of values to be used with the `property`, they will be combined via `OR`.
         * 
         * @return builder
         * 
         */
        public Builder values(@Nullable Output<List<String>> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values A list of values to be used with the `property`, they will be combined via `OR`.
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        /**
         * @param values A list of values to be used with the `property`, they will be combined via `OR`.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        /**
         * @param valuesSuggesteds A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
         * 
         * @return builder
         * 
         */
        public Builder valuesSuggesteds(@Nullable Output<List<String>> valuesSuggesteds) {
            $.valuesSuggesteds = valuesSuggesteds;
            return this;
        }

        /**
         * @param valuesSuggesteds A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
         * 
         * @return builder
         * 
         */
        public Builder valuesSuggesteds(List<String> valuesSuggesteds) {
            return valuesSuggesteds(Output.of(valuesSuggesteds));
        }

        /**
         * @param valuesSuggesteds A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
         * 
         * @return builder
         * 
         */
        public Builder valuesSuggesteds(String... valuesSuggesteds) {
            return valuesSuggesteds(List.of(valuesSuggesteds));
        }

        public DashboardVariableArgs build() {
            $.alias = Objects.requireNonNull($.alias, "expected parameter 'alias' to be non-null");
            $.property = Objects.requireNonNull($.property, "expected parameter 'property' to be non-null");
            return $;
        }
    }

}
