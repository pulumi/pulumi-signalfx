// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.logs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ListChartSortOption {
    private final Boolean descending;
    private final String field;

    @CustomType.Constructor
    private ListChartSortOption(
        @CustomType.Parameter("descending") Boolean descending,
        @CustomType.Parameter("field") String field) {
        this.descending = descending;
        this.field = field;
    }

    public Boolean descending() {
        return this.descending;
    }
    public String field() {
        return this.field;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ListChartSortOption defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Boolean descending;
        private String field;

        public Builder() {
    	      // Empty
        }

        public Builder(ListChartSortOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.descending = defaults.descending;
    	      this.field = defaults.field;
        }

        public Builder descending(Boolean descending) {
            this.descending = Objects.requireNonNull(descending);
            return this;
        }
        public Builder field(String field) {
            this.field = Objects.requireNonNull(field);
            return this;
        }        public ListChartSortOption build() {
            return new ListChartSortOption(descending, field);
        }
    }
}
