// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardChartArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardChartArgs Empty = new DashboardChartArgs();

    /**
     * ID of the chart to display
     * 
     */
    @Import(name="chartId", required=true)
    private Output<String> chartId;

    /**
     * @return ID of the chart to display
     * 
     */
    public Output<String> chartId() {
        return this.chartId;
    }

    /**
     * The column to show the chart in (zero-based); this value always represents the leftmost column of the chart. (between 0 and 11)
     * 
     */
    @Import(name="column")
    private @Nullable Output<Integer> column;

    /**
     * @return The column to show the chart in (zero-based); this value always represents the leftmost column of the chart. (between 0 and 11)
     * 
     */
    public Optional<Output<Integer>> column() {
        return Optional.ofNullable(this.column);
    }

    /**
     * How many rows the chart should take up. (greater than or equal to 1)
     * 
     */
    @Import(name="height")
    private @Nullable Output<Integer> height;

    /**
     * @return How many rows the chart should take up. (greater than or equal to 1)
     * 
     */
    public Optional<Output<Integer>> height() {
        return Optional.ofNullable(this.height);
    }

    /**
     * The row to show the chart in (zero-based); if height &gt; 1, this value represents the topmost row of the chart. (greater than or equal to 0)
     * 
     */
    @Import(name="row")
    private @Nullable Output<Integer> row;

    /**
     * @return The row to show the chart in (zero-based); if height &gt; 1, this value represents the topmost row of the chart. (greater than or equal to 0)
     * 
     */
    public Optional<Output<Integer>> row() {
        return Optional.ofNullable(this.row);
    }

    /**
     * How many columns (out of a total of 12, one-based) the chart should take up. (between 1 and 12)
     * 
     */
    @Import(name="width")
    private @Nullable Output<Integer> width;

    /**
     * @return How many columns (out of a total of 12, one-based) the chart should take up. (between 1 and 12)
     * 
     */
    public Optional<Output<Integer>> width() {
        return Optional.ofNullable(this.width);
    }

    private DashboardChartArgs() {}

    private DashboardChartArgs(DashboardChartArgs $) {
        this.chartId = $.chartId;
        this.column = $.column;
        this.height = $.height;
        this.row = $.row;
        this.width = $.width;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardChartArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardChartArgs $;

        public Builder() {
            $ = new DashboardChartArgs();
        }

        public Builder(DashboardChartArgs defaults) {
            $ = new DashboardChartArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param chartId ID of the chart to display
         * 
         * @return builder
         * 
         */
        public Builder chartId(Output<String> chartId) {
            $.chartId = chartId;
            return this;
        }

        /**
         * @param chartId ID of the chart to display
         * 
         * @return builder
         * 
         */
        public Builder chartId(String chartId) {
            return chartId(Output.of(chartId));
        }

        /**
         * @param column The column to show the chart in (zero-based); this value always represents the leftmost column of the chart. (between 0 and 11)
         * 
         * @return builder
         * 
         */
        public Builder column(@Nullable Output<Integer> column) {
            $.column = column;
            return this;
        }

        /**
         * @param column The column to show the chart in (zero-based); this value always represents the leftmost column of the chart. (between 0 and 11)
         * 
         * @return builder
         * 
         */
        public Builder column(Integer column) {
            return column(Output.of(column));
        }

        /**
         * @param height How many rows the chart should take up. (greater than or equal to 1)
         * 
         * @return builder
         * 
         */
        public Builder height(@Nullable Output<Integer> height) {
            $.height = height;
            return this;
        }

        /**
         * @param height How many rows the chart should take up. (greater than or equal to 1)
         * 
         * @return builder
         * 
         */
        public Builder height(Integer height) {
            return height(Output.of(height));
        }

        /**
         * @param row The row to show the chart in (zero-based); if height &gt; 1, this value represents the topmost row of the chart. (greater than or equal to 0)
         * 
         * @return builder
         * 
         */
        public Builder row(@Nullable Output<Integer> row) {
            $.row = row;
            return this;
        }

        /**
         * @param row The row to show the chart in (zero-based); if height &gt; 1, this value represents the topmost row of the chart. (greater than or equal to 0)
         * 
         * @return builder
         * 
         */
        public Builder row(Integer row) {
            return row(Output.of(row));
        }

        /**
         * @param width How many columns (out of a total of 12, one-based) the chart should take up. (between 1 and 12)
         * 
         * @return builder
         * 
         */
        public Builder width(@Nullable Output<Integer> width) {
            $.width = width;
            return this;
        }

        /**
         * @param width How many columns (out of a total of 12, one-based) the chart should take up. (between 1 and 12)
         * 
         * @return builder
         * 
         */
        public Builder width(Integer width) {
            return width(Output.of(width));
        }

        public DashboardChartArgs build() {
            if ($.chartId == null) {
                throw new MissingRequiredPropertyException("DashboardChartArgs", "chartId");
            }
            return $;
        }
    }

}
