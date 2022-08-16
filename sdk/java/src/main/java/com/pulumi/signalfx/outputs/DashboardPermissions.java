// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.signalfx.outputs.DashboardPermissionsAcl;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DashboardPermissions {
    /**
     * @return List of read and write permission configurations to specify which user, team, and organization can view and/or edit your dashboard. Use the `permissions.parent` instead if you want to inherit permissions.
     * 
     */
    private final @Nullable List<DashboardPermissionsAcl> acls;
    /**
     * @return ID of the dashboard group you want your dashboard to inherit permissions from. Use the `permissions.acl` instead if you want to specify various read and write permission configurations.
     * 
     */
    private final @Nullable String parent;

    @CustomType.Constructor
    private DashboardPermissions(
        @CustomType.Parameter("acls") @Nullable List<DashboardPermissionsAcl> acls,
        @CustomType.Parameter("parent") @Nullable String parent) {
        this.acls = acls;
        this.parent = parent;
    }

    /**
     * @return List of read and write permission configurations to specify which user, team, and organization can view and/or edit your dashboard. Use the `permissions.parent` instead if you want to inherit permissions.
     * 
     */
    public List<DashboardPermissionsAcl> acls() {
        return this.acls == null ? List.of() : this.acls;
    }
    /**
     * @return ID of the dashboard group you want your dashboard to inherit permissions from. Use the `permissions.acl` instead if you want to specify various read and write permission configurations.
     * 
     */
    public Optional<String> parent() {
        return Optional.ofNullable(this.parent);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardPermissions defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<DashboardPermissionsAcl> acls;
        private @Nullable String parent;

        public Builder() {
    	      // Empty
        }

        public Builder(DashboardPermissions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.acls = defaults.acls;
    	      this.parent = defaults.parent;
        }

        public Builder acls(@Nullable List<DashboardPermissionsAcl> acls) {
            this.acls = acls;
            return this;
        }
        public Builder acls(DashboardPermissionsAcl... acls) {
            return acls(List.of(acls));
        }
        public Builder parent(@Nullable String parent) {
            this.parent = parent;
            return this;
        }        public DashboardPermissions build() {
            return new DashboardPermissions(acls, parent);
        }
    }
}
