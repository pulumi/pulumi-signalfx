// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DashboardChart {
    /**
     * @return ID of the chart to display.
     * 
     */
    private String chartId;
    /**
     * @return The column to show the chart in (zero-based); this value always represents the leftmost column of the chart (between `0` and `11`).
     * 
     */
    private @Nullable Integer column;
    /**
     * @return How many rows the chart should take up (greater than or equal to `1`). `1` by default.
     * 
     */
    private @Nullable Integer height;
    /**
     * @return The row to show the chart in (zero-based); if `height &gt; 1`, this value represents the topmost row of the chart (greater than or equal to `0`).
     * 
     */
    private @Nullable Integer row;
    /**
     * @return How many columns (out of a total of 12) the chart should take up (between `1` and `12`). `12` by default.
     * 
     */
    private @Nullable Integer width;

    private DashboardChart() {}
    /**
     * @return ID of the chart to display.
     * 
     */
    public String chartId() {
        return this.chartId;
    }
    /**
     * @return The column to show the chart in (zero-based); this value always represents the leftmost column of the chart (between `0` and `11`).
     * 
     */
    public Optional<Integer> column() {
        return Optional.ofNullable(this.column);
    }
    /**
     * @return How many rows the chart should take up (greater than or equal to `1`). `1` by default.
     * 
     */
    public Optional<Integer> height() {
        return Optional.ofNullable(this.height);
    }
    /**
     * @return The row to show the chart in (zero-based); if `height &gt; 1`, this value represents the topmost row of the chart (greater than or equal to `0`).
     * 
     */
    public Optional<Integer> row() {
        return Optional.ofNullable(this.row);
    }
    /**
     * @return How many columns (out of a total of 12) the chart should take up (between `1` and `12`). `12` by default.
     * 
     */
    public Optional<Integer> width() {
        return Optional.ofNullable(this.width);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardChart defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String chartId;
        private @Nullable Integer column;
        private @Nullable Integer height;
        private @Nullable Integer row;
        private @Nullable Integer width;
        public Builder() {}
        public Builder(DashboardChart defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.chartId = defaults.chartId;
    	      this.column = defaults.column;
    	      this.height = defaults.height;
    	      this.row = defaults.row;
    	      this.width = defaults.width;
        }

        @CustomType.Setter
        public Builder chartId(String chartId) {
            if (chartId == null) {
              throw new MissingRequiredPropertyException("DashboardChart", "chartId");
            }
            this.chartId = chartId;
            return this;
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
        public Builder row(@Nullable Integer row) {

            this.row = row;
            return this;
        }
        @CustomType.Setter
        public Builder width(@Nullable Integer width) {

            this.width = width;
            return this;
        }
        public DashboardChart build() {
            final var _resultValue = new DashboardChart();
            _resultValue.chartId = chartId;
            _resultValue.column = column;
            _resultValue.height = height;
            _resultValue.row = row;
            _resultValue.width = width;
            return _resultValue;
        }
    }
}
