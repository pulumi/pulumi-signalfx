// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.outputs.DashboardGroupDashboardFilterOverride;
import com.pulumi.signalfx.outputs.DashboardGroupDashboardVariableOverride;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DashboardGroupDashboard {
    /**
     * @return Unique identifier of an association between a dashboard group and a dashboard
     * 
     */
    private @Nullable String configId;
    /**
     * @return The label used in the publish statement that displays the plot (metric time series data) you want to customize
     * 
     */
    private String dashboardId;
    /**
     * @return String that provides a description override for a mirrored dashboard
     * 
     */
    private @Nullable String descriptionOverride;
    /**
     * @return Filter to apply to each chart in the dashboard
     * 
     */
    private @Nullable List<DashboardGroupDashboardFilterOverride> filterOverrides;
    /**
     * @return String that provides a name override for a mirrored dashboard
     * 
     */
    private @Nullable String nameOverride;
    /**
     * @return Dashboard variable to apply to each chart in the dashboard
     * 
     */
    private @Nullable List<DashboardGroupDashboardVariableOverride> variableOverrides;

    private DashboardGroupDashboard() {}
    /**
     * @return Unique identifier of an association between a dashboard group and a dashboard
     * 
     */
    public Optional<String> configId() {
        return Optional.ofNullable(this.configId);
    }
    /**
     * @return The label used in the publish statement that displays the plot (metric time series data) you want to customize
     * 
     */
    public String dashboardId() {
        return this.dashboardId;
    }
    /**
     * @return String that provides a description override for a mirrored dashboard
     * 
     */
    public Optional<String> descriptionOverride() {
        return Optional.ofNullable(this.descriptionOverride);
    }
    /**
     * @return Filter to apply to each chart in the dashboard
     * 
     */
    public List<DashboardGroupDashboardFilterOverride> filterOverrides() {
        return this.filterOverrides == null ? List.of() : this.filterOverrides;
    }
    /**
     * @return String that provides a name override for a mirrored dashboard
     * 
     */
    public Optional<String> nameOverride() {
        return Optional.ofNullable(this.nameOverride);
    }
    /**
     * @return Dashboard variable to apply to each chart in the dashboard
     * 
     */
    public List<DashboardGroupDashboardVariableOverride> variableOverrides() {
        return this.variableOverrides == null ? List.of() : this.variableOverrides;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardGroupDashboard defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String configId;
        private String dashboardId;
        private @Nullable String descriptionOverride;
        private @Nullable List<DashboardGroupDashboardFilterOverride> filterOverrides;
        private @Nullable String nameOverride;
        private @Nullable List<DashboardGroupDashboardVariableOverride> variableOverrides;
        public Builder() {}
        public Builder(DashboardGroupDashboard defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configId = defaults.configId;
    	      this.dashboardId = defaults.dashboardId;
    	      this.descriptionOverride = defaults.descriptionOverride;
    	      this.filterOverrides = defaults.filterOverrides;
    	      this.nameOverride = defaults.nameOverride;
    	      this.variableOverrides = defaults.variableOverrides;
        }

        @CustomType.Setter
        public Builder configId(@Nullable String configId) {

            this.configId = configId;
            return this;
        }
        @CustomType.Setter
        public Builder dashboardId(String dashboardId) {
            if (dashboardId == null) {
              throw new MissingRequiredPropertyException("DashboardGroupDashboard", "dashboardId");
            }
            this.dashboardId = dashboardId;
            return this;
        }
        @CustomType.Setter
        public Builder descriptionOverride(@Nullable String descriptionOverride) {

            this.descriptionOverride = descriptionOverride;
            return this;
        }
        @CustomType.Setter
        public Builder filterOverrides(@Nullable List<DashboardGroupDashboardFilterOverride> filterOverrides) {

            this.filterOverrides = filterOverrides;
            return this;
        }
        public Builder filterOverrides(DashboardGroupDashboardFilterOverride... filterOverrides) {
            return filterOverrides(List.of(filterOverrides));
        }
        @CustomType.Setter
        public Builder nameOverride(@Nullable String nameOverride) {

            this.nameOverride = nameOverride;
            return this;
        }
        @CustomType.Setter
        public Builder variableOverrides(@Nullable List<DashboardGroupDashboardVariableOverride> variableOverrides) {

            this.variableOverrides = variableOverrides;
            return this;
        }
        public Builder variableOverrides(DashboardGroupDashboardVariableOverride... variableOverrides) {
            return variableOverrides(List.of(variableOverrides));
        }
        public DashboardGroupDashboard build() {
            final var _resultValue = new DashboardGroupDashboard();
            _resultValue.configId = configId;
            _resultValue.dashboardId = dashboardId;
            _resultValue.descriptionOverride = descriptionOverride;
            _resultValue.filterOverrides = filterOverrides;
            _resultValue.nameOverride = nameOverride;
            _resultValue.variableOverrides = variableOverrides;
            return _resultValue;
        }
    }
}
