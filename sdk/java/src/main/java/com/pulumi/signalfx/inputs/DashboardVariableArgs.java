// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardVariableArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardVariableArgs Empty = new DashboardVariableArgs();

    @Import(name="alias", required=true)
    private Output<String> alias;

    public Output<String> alias() {
        return this.alias;
    }

    @Import(name="applyIfExist")
    private @Nullable Output<Boolean> applyIfExist;

    public Optional<Output<Boolean>> applyIfExist() {
        return Optional.ofNullable(this.applyIfExist);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="property", required=true)
    private Output<String> property;

    public Output<String> property() {
        return this.property;
    }

    @Import(name="replaceOnly")
    private @Nullable Output<Boolean> replaceOnly;

    public Optional<Output<Boolean>> replaceOnly() {
        return Optional.ofNullable(this.replaceOnly);
    }

    @Import(name="restrictedSuggestions")
    private @Nullable Output<Boolean> restrictedSuggestions;

    public Optional<Output<Boolean>> restrictedSuggestions() {
        return Optional.ofNullable(this.restrictedSuggestions);
    }

    @Import(name="valueRequired")
    private @Nullable Output<Boolean> valueRequired;

    public Optional<Output<Boolean>> valueRequired() {
        return Optional.ofNullable(this.valueRequired);
    }

    @Import(name="values")
    private @Nullable Output<List<String>> values;

    public Optional<Output<List<String>>> values() {
        return Optional.ofNullable(this.values);
    }

    @Import(name="valuesSuggesteds")
    private @Nullable Output<List<String>> valuesSuggesteds;

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

        public Builder alias(Output<String> alias) {
            $.alias = alias;
            return this;
        }

        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        public Builder applyIfExist(@Nullable Output<Boolean> applyIfExist) {
            $.applyIfExist = applyIfExist;
            return this;
        }

        public Builder applyIfExist(Boolean applyIfExist) {
            return applyIfExist(Output.of(applyIfExist));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder property(Output<String> property) {
            $.property = property;
            return this;
        }

        public Builder property(String property) {
            return property(Output.of(property));
        }

        public Builder replaceOnly(@Nullable Output<Boolean> replaceOnly) {
            $.replaceOnly = replaceOnly;
            return this;
        }

        public Builder replaceOnly(Boolean replaceOnly) {
            return replaceOnly(Output.of(replaceOnly));
        }

        public Builder restrictedSuggestions(@Nullable Output<Boolean> restrictedSuggestions) {
            $.restrictedSuggestions = restrictedSuggestions;
            return this;
        }

        public Builder restrictedSuggestions(Boolean restrictedSuggestions) {
            return restrictedSuggestions(Output.of(restrictedSuggestions));
        }

        public Builder valueRequired(@Nullable Output<Boolean> valueRequired) {
            $.valueRequired = valueRequired;
            return this;
        }

        public Builder valueRequired(Boolean valueRequired) {
            return valueRequired(Output.of(valueRequired));
        }

        public Builder values(@Nullable Output<List<String>> values) {
            $.values = values;
            return this;
        }

        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        public Builder values(String... values) {
            return values(List.of(values));
        }

        public Builder valuesSuggesteds(@Nullable Output<List<String>> valuesSuggesteds) {
            $.valuesSuggesteds = valuesSuggesteds;
            return this;
        }

        public Builder valuesSuggesteds(List<String> valuesSuggesteds) {
            return valuesSuggesteds(Output.of(valuesSuggesteds));
        }

        public Builder valuesSuggesteds(String... valuesSuggesteds) {
            return valuesSuggesteds(List.of(valuesSuggesteds));
        }

        public DashboardVariableArgs build() {
            if ($.alias == null) {
                throw new MissingRequiredPropertyException("DashboardVariableArgs", "alias");
            }
            if ($.property == null) {
                throw new MissingRequiredPropertyException("DashboardVariableArgs", "property");
            }
            return $;
        }
    }

}
