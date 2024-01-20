// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class DashboardGroupDashboardVariableOverride {
    private String property;
    private @Nullable List<String> values;
    private @Nullable List<String> valuesSuggesteds;

    private DashboardGroupDashboardVariableOverride() {}
    public String property() {
        return this.property;
    }
    public List<String> values() {
        return this.values == null ? List.of() : this.values;
    }
    public List<String> valuesSuggesteds() {
        return this.valuesSuggesteds == null ? List.of() : this.valuesSuggesteds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardGroupDashboardVariableOverride defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String property;
        private @Nullable List<String> values;
        private @Nullable List<String> valuesSuggesteds;
        public Builder() {}
        public Builder(DashboardGroupDashboardVariableOverride defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.property = defaults.property;
    	      this.values = defaults.values;
    	      this.valuesSuggesteds = defaults.valuesSuggesteds;
        }

        @CustomType.Setter
        public Builder property(String property) {
            if (property == null) {
              throw new MissingRequiredPropertyException("DashboardGroupDashboardVariableOverride", "property");
            }
            this.property = property;
            return this;
        }
        @CustomType.Setter
        public Builder values(@Nullable List<String> values) {

            this.values = values;
            return this;
        }
        public Builder values(String... values) {
            return values(List.of(values));
        }
        @CustomType.Setter
        public Builder valuesSuggesteds(@Nullable List<String> valuesSuggesteds) {

            this.valuesSuggesteds = valuesSuggesteds;
            return this;
        }
        public Builder valuesSuggesteds(String... valuesSuggesteds) {
            return valuesSuggesteds(List.of(valuesSuggesteds));
        }
        public DashboardGroupDashboardVariableOverride build() {
            final var _resultValue = new DashboardGroupDashboardVariableOverride();
            _resultValue.property = property;
            _resultValue.values = values;
            _resultValue.valuesSuggesteds = valuesSuggesteds;
            return _resultValue;
        }
    }
}
