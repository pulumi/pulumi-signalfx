// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TableChartVizOption {
    private @Nullable String color;
    private @Nullable String displayName;
    private String label;
    private @Nullable String valuePrefix;
    private @Nullable String valueSuffix;
    private @Nullable String valueUnit;

    private TableChartVizOption() {}
    public Optional<String> color() {
        return Optional.ofNullable(this.color);
    }
    public Optional<String> displayName() {
        return Optional.ofNullable(this.displayName);
    }
    public String label() {
        return this.label;
    }
    public Optional<String> valuePrefix() {
        return Optional.ofNullable(this.valuePrefix);
    }
    public Optional<String> valueSuffix() {
        return Optional.ofNullable(this.valueSuffix);
    }
    public Optional<String> valueUnit() {
        return Optional.ofNullable(this.valueUnit);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TableChartVizOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String color;
        private @Nullable String displayName;
        private String label;
        private @Nullable String valuePrefix;
        private @Nullable String valueSuffix;
        private @Nullable String valueUnit;
        public Builder() {}
        public Builder(TableChartVizOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.color = defaults.color;
    	      this.displayName = defaults.displayName;
    	      this.label = defaults.label;
    	      this.valuePrefix = defaults.valuePrefix;
    	      this.valueSuffix = defaults.valueSuffix;
    	      this.valueUnit = defaults.valueUnit;
        }

        @CustomType.Setter
        public Builder color(@Nullable String color) {
            this.color = color;
            return this;
        }
        @CustomType.Setter
        public Builder displayName(@Nullable String displayName) {
            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder label(String label) {
            this.label = Objects.requireNonNull(label);
            return this;
        }
        @CustomType.Setter
        public Builder valuePrefix(@Nullable String valuePrefix) {
            this.valuePrefix = valuePrefix;
            return this;
        }
        @CustomType.Setter
        public Builder valueSuffix(@Nullable String valueSuffix) {
            this.valueSuffix = valueSuffix;
            return this;
        }
        @CustomType.Setter
        public Builder valueUnit(@Nullable String valueUnit) {
            this.valueUnit = valueUnit;
            return this;
        }
        public TableChartVizOption build() {
            final var o = new TableChartVizOption();
            o.color = color;
            o.displayName = displayName;
            o.label = label;
            o.valuePrefix = valuePrefix;
            o.valueSuffix = valueSuffix;
            o.valueUnit = valueUnit;
            return o;
        }
    }
}
