// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DashboardColumn {
    private List<String> chartIds;
    private @Nullable Integer column;
    private @Nullable Integer height;
    private @Nullable Integer width;

    private DashboardColumn() {}
    public List<String> chartIds() {
        return this.chartIds;
    }
    public Optional<Integer> column() {
        return Optional.ofNullable(this.column);
    }
    public Optional<Integer> height() {
        return Optional.ofNullable(this.height);
    }
    public Optional<Integer> width() {
        return Optional.ofNullable(this.width);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardColumn defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> chartIds;
        private @Nullable Integer column;
        private @Nullable Integer height;
        private @Nullable Integer width;
        public Builder() {}
        public Builder(DashboardColumn defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.chartIds = defaults.chartIds;
    	      this.column = defaults.column;
    	      this.height = defaults.height;
    	      this.width = defaults.width;
        }

        @CustomType.Setter
        public Builder chartIds(List<String> chartIds) {
            if (chartIds == null) {
              throw new MissingRequiredPropertyException("DashboardColumn", "chartIds");
            }
            this.chartIds = chartIds;
            return this;
        }
        public Builder chartIds(String... chartIds) {
            return chartIds(List.of(chartIds));
        }
        @CustomType.Setter
        public Builder column(@Nullable Integer column) {

            this.column = column;
            return this;
        }
        @CustomType.Setter
        public Builder height(@Nullable Integer height) {

            this.height = height;
            return this;
        }
        @CustomType.Setter
        public Builder width(@Nullable Integer width) {

            this.width = width;
            return this;
        }
        public DashboardColumn build() {
            final var _resultValue = new DashboardColumn();
            _resultValue.chartIds = chartIds;
            _resultValue.column = column;
            _resultValue.height = height;
            _resultValue.width = width;
            return _resultValue;
        }
    }
}
