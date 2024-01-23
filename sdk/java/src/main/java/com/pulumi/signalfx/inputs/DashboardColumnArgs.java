// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardColumnArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardColumnArgs Empty = new DashboardColumnArgs();

    @Import(name="chartIds", required=true)
    private Output<List<String>> chartIds;

    public Output<List<String>> chartIds() {
        return this.chartIds;
    }

    @Import(name="column")
    private @Nullable Output<Integer> column;

    public Optional<Output<Integer>> column() {
        return Optional.ofNullable(this.column);
    }

    @Import(name="height")
    private @Nullable Output<Integer> height;

    public Optional<Output<Integer>> height() {
        return Optional.ofNullable(this.height);
    }

    @Import(name="width")
    private @Nullable Output<Integer> width;

    public Optional<Output<Integer>> width() {
        return Optional.ofNullable(this.width);
    }

    private DashboardColumnArgs() {}

    private DashboardColumnArgs(DashboardColumnArgs $) {
        this.chartIds = $.chartIds;
        this.column = $.column;
        this.height = $.height;
        this.width = $.width;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardColumnArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardColumnArgs $;

        public Builder() {
            $ = new DashboardColumnArgs();
        }

        public Builder(DashboardColumnArgs defaults) {
            $ = new DashboardColumnArgs(Objects.requireNonNull(defaults));
        }

        public Builder chartIds(Output<List<String>> chartIds) {
            $.chartIds = chartIds;
            return this;
        }

        public Builder chartIds(List<String> chartIds) {
            return chartIds(Output.of(chartIds));
        }

        public Builder chartIds(String... chartIds) {
            return chartIds(List.of(chartIds));
        }

        public Builder column(@Nullable Output<Integer> column) {
            $.column = column;
            return this;
        }

        public Builder column(Integer column) {
            return column(Output.of(column));
        }

        public Builder height(@Nullable Output<Integer> height) {
            $.height = height;
            return this;
        }

        public Builder height(Integer height) {
            return height(Output.of(height));
        }

        public Builder width(@Nullable Output<Integer> width) {
            $.width = width;
            return this;
        }

        public Builder width(Integer width) {
            return width(Output.of(width));
        }

        public DashboardColumnArgs build() {
            if ($.chartIds == null) {
                throw new MissingRequiredPropertyException("DashboardColumnArgs", "chartIds");
            }
            return $;
        }
    }

}
