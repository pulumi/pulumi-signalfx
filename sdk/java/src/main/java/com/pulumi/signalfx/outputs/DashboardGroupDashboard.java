// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
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
     * @return The dashboard id to mirror
     * 
     */
    private String dashboardId;
    /**
     * @return The description that will override the original dashboards&#39;s description.
     * 
     */
    private @Nullable String descriptionOverride;
    /**
     * @return The description that will override the original dashboards&#39;s description.
     * 
     */
    private @Nullable List<DashboardGroupDashboardFilterOverride> filterOverrides;
    /**
     * @return The name that will override the original dashboards&#39;s name.
     * 
     */
    private @Nullable String nameOverride;
    private @Nullable List<DashboardGroupDashboardVariableOverride> variableOverrides;

    private DashboardGroupDashboard() {}
    /**
     * @return The dashboard id to mirror
     * 
     */
    public String dashboardId() {
        return this.dashboardId;
    }
    /**
     * @return The description that will override the original dashboards&#39;s description.
     * 
     */
    public Optional<String> descriptionOverride() {
        return Optional.ofNullable(this.descriptionOverride);
    }
    /**
     * @return The description that will override the original dashboards&#39;s description.
     * 
     */
    public List<DashboardGroupDashboardFilterOverride> filterOverrides() {
        return this.filterOverrides == null ? List.of() : this.filterOverrides;
    }
    /**
     * @return The name that will override the original dashboards&#39;s name.
     * 
     */
    public Optional<String> nameOverride() {
        return Optional.ofNullable(this.nameOverride);
    }
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
        private String dashboardId;
        private @Nullable String descriptionOverride;
        private @Nullable List<DashboardGroupDashboardFilterOverride> filterOverrides;
        private @Nullable String nameOverride;
        private @Nullable List<DashboardGroupDashboardVariableOverride> variableOverrides;
        public Builder() {}
        public Builder(DashboardGroupDashboard defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dashboardId = defaults.dashboardId;
    	      this.descriptionOverride = defaults.descriptionOverride;
    	      this.filterOverrides = defaults.filterOverrides;
    	      this.nameOverride = defaults.nameOverride;
    	      this.variableOverrides = defaults.variableOverrides;
        }

        @CustomType.Setter
        public Builder dashboardId(String dashboardId) {
            this.dashboardId = Objects.requireNonNull(dashboardId);
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
            final var o = new DashboardGroupDashboard();
            o.dashboardId = dashboardId;
            o.descriptionOverride = descriptionOverride;
            o.filterOverrides = filterOverrides;
            o.nameOverride = nameOverride;
            o.variableOverrides = variableOverrides;
            return o;
        }
    }
}
