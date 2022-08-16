// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DataLinkTargetSignalfxDashboard {
    /**
     * @return SignalFx-assigned ID of the dashboard link target&#39;s dashboard group
     * 
     */
    private final String dashboardGroupId;
    /**
     * @return SignalFx-assigned ID of the dashboard link target
     * 
     */
    private final String dashboardId;
    /**
     * @return Flag that designates a target as the default for a data link object. `true` by default
     * 
     */
    private final @Nullable Boolean isDefault;
    /**
     * @return User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     * 
     */
    private final String name;

    @CustomType.Constructor
    private DataLinkTargetSignalfxDashboard(
        @CustomType.Parameter("dashboardGroupId") String dashboardGroupId,
        @CustomType.Parameter("dashboardId") String dashboardId,
        @CustomType.Parameter("isDefault") @Nullable Boolean isDefault,
        @CustomType.Parameter("name") String name) {
        this.dashboardGroupId = dashboardGroupId;
        this.dashboardId = dashboardId;
        this.isDefault = isDefault;
        this.name = name;
    }

    /**
     * @return SignalFx-assigned ID of the dashboard link target&#39;s dashboard group
     * 
     */
    public String dashboardGroupId() {
        return this.dashboardGroupId;
    }
    /**
     * @return SignalFx-assigned ID of the dashboard link target
     * 
     */
    public String dashboardId() {
        return this.dashboardId;
    }
    /**
     * @return Flag that designates a target as the default for a data link object. `true` by default
     * 
     */
    public Optional<Boolean> isDefault() {
        return Optional.ofNullable(this.isDefault);
    }
    /**
     * @return User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DataLinkTargetSignalfxDashboard defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String dashboardGroupId;
        private String dashboardId;
        private @Nullable Boolean isDefault;
        private String name;

        public Builder() {
    	      // Empty
        }

        public Builder(DataLinkTargetSignalfxDashboard defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dashboardGroupId = defaults.dashboardGroupId;
    	      this.dashboardId = defaults.dashboardId;
    	      this.isDefault = defaults.isDefault;
    	      this.name = defaults.name;
        }

        public Builder dashboardGroupId(String dashboardGroupId) {
            this.dashboardGroupId = Objects.requireNonNull(dashboardGroupId);
            return this;
        }
        public Builder dashboardId(String dashboardId) {
            this.dashboardId = Objects.requireNonNull(dashboardId);
            return this;
        }
        public Builder isDefault(@Nullable Boolean isDefault) {
            this.isDefault = isDefault;
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }        public DataLinkTargetSignalfxDashboard build() {
            return new DataLinkTargetSignalfxDashboard(dashboardGroupId, dashboardId, isDefault, name);
        }
    }
}
