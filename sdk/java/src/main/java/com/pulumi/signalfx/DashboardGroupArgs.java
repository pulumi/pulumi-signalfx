// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.inputs.DashboardGroupDashboardArgs;
import com.pulumi.signalfx.inputs.DashboardGroupImportQualifierArgs;
import com.pulumi.signalfx.inputs.DashboardGroupPermissionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardGroupArgs Empty = new DashboardGroupArgs();

    /**
     * Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     * @deprecated
     * Please use permissions field now
     * 
     */
    @Deprecated /* Please use permissions field now */
    @Import(name="authorizedWriterTeams")
    private @Nullable Output<List<String>> authorizedWriterTeams;

    /**
     * @return Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     * @deprecated
     * Please use permissions field now
     * 
     */
    @Deprecated /* Please use permissions field now */
    public Optional<Output<List<String>>> authorizedWriterTeams() {
        return Optional.ofNullable(this.authorizedWriterTeams);
    }

    /**
     * User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     * @deprecated
     * Please use permissions field now
     * 
     */
    @Deprecated /* Please use permissions field now */
    @Import(name="authorizedWriterUsers")
    private @Nullable Output<List<String>> authorizedWriterUsers;

    /**
     * @return User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     * @deprecated
     * Please use permissions field now
     * 
     */
    @Deprecated /* Please use permissions field now */
    public Optional<Output<List<String>>> authorizedWriterUsers() {
        return Optional.ofNullable(this.authorizedWriterUsers);
    }

    /**
     * [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    @Import(name="dashboards")
    private @Nullable Output<List<DashboardGroupDashboardArgs>> dashboards;

    /**
     * @return [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    public Optional<Output<List<DashboardGroupDashboardArgs>>> dashboards() {
        return Optional.ofNullable(this.dashboards);
    }

    /**
     * Description of the dashboard group.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the dashboard group.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="importQualifiers")
    private @Nullable Output<List<DashboardGroupImportQualifierArgs>> importQualifiers;

    public Optional<Output<List<DashboardGroupImportQualifierArgs>>> importQualifiers() {
        return Optional.ofNullable(this.importQualifiers);
    }

    /**
     * Name of the dashboard group.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the dashboard group.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) List of read and write permission configuration to specify which user, team, and organization can view and/or edit your dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<List<DashboardGroupPermissionArgs>> permissions;

    /**
     * @return [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) List of read and write permission configuration to specify which user, team, and organization can view and/or edit your dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    public Optional<Output<List<DashboardGroupPermissionArgs>>> permissions() {
        return Optional.ofNullable(this.permissions);
    }

    /**
     * Team IDs to associate the dashboard group to.
     * 
     */
    @Import(name="teams")
    private @Nullable Output<List<String>> teams;

    /**
     * @return Team IDs to associate the dashboard group to.
     * 
     */
    public Optional<Output<List<String>>> teams() {
        return Optional.ofNullable(this.teams);
    }

    private DashboardGroupArgs() {}

    private DashboardGroupArgs(DashboardGroupArgs $) {
        this.authorizedWriterTeams = $.authorizedWriterTeams;
        this.authorizedWriterUsers = $.authorizedWriterUsers;
        this.dashboards = $.dashboards;
        this.description = $.description;
        this.importQualifiers = $.importQualifiers;
        this.name = $.name;
        this.permissions = $.permissions;
        this.teams = $.teams;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardGroupArgs $;

        public Builder() {
            $ = new DashboardGroupArgs();
        }

        public Builder(DashboardGroupArgs defaults) {
            $ = new DashboardGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authorizedWriterTeams Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use permissions field now
         * 
         */
        @Deprecated /* Please use permissions field now */
        public Builder authorizedWriterTeams(@Nullable Output<List<String>> authorizedWriterTeams) {
            $.authorizedWriterTeams = authorizedWriterTeams;
            return this;
        }

        /**
         * @param authorizedWriterTeams Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use permissions field now
         * 
         */
        @Deprecated /* Please use permissions field now */
        public Builder authorizedWriterTeams(List<String> authorizedWriterTeams) {
            return authorizedWriterTeams(Output.of(authorizedWriterTeams));
        }

        /**
         * @param authorizedWriterTeams Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use permissions field now
         * 
         */
        @Deprecated /* Please use permissions field now */
        public Builder authorizedWriterTeams(String... authorizedWriterTeams) {
            return authorizedWriterTeams(List.of(authorizedWriterTeams));
        }

        /**
         * @param authorizedWriterUsers User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use permissions field now
         * 
         */
        @Deprecated /* Please use permissions field now */
        public Builder authorizedWriterUsers(@Nullable Output<List<String>> authorizedWriterUsers) {
            $.authorizedWriterUsers = authorizedWriterUsers;
            return this;
        }

        /**
         * @param authorizedWriterUsers User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use permissions field now
         * 
         */
        @Deprecated /* Please use permissions field now */
        public Builder authorizedWriterUsers(List<String> authorizedWriterUsers) {
            return authorizedWriterUsers(Output.of(authorizedWriterUsers));
        }

        /**
         * @param authorizedWriterUsers User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use permissions field now
         * 
         */
        @Deprecated /* Please use permissions field now */
        public Builder authorizedWriterUsers(String... authorizedWriterUsers) {
            return authorizedWriterUsers(List.of(authorizedWriterUsers));
        }

        /**
         * @param dashboards [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
         * 
         * @return builder
         * 
         */
        public Builder dashboards(@Nullable Output<List<DashboardGroupDashboardArgs>> dashboards) {
            $.dashboards = dashboards;
            return this;
        }

        /**
         * @param dashboards [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
         * 
         * @return builder
         * 
         */
        public Builder dashboards(List<DashboardGroupDashboardArgs> dashboards) {
            return dashboards(Output.of(dashboards));
        }

        /**
         * @param dashboards [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
         * 
         * @return builder
         * 
         */
        public Builder dashboards(DashboardGroupDashboardArgs... dashboards) {
            return dashboards(List.of(dashboards));
        }

        /**
         * @param description Description of the dashboard group.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the dashboard group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder importQualifiers(@Nullable Output<List<DashboardGroupImportQualifierArgs>> importQualifiers) {
            $.importQualifiers = importQualifiers;
            return this;
        }

        public Builder importQualifiers(List<DashboardGroupImportQualifierArgs> importQualifiers) {
            return importQualifiers(Output.of(importQualifiers));
        }

        public Builder importQualifiers(DashboardGroupImportQualifierArgs... importQualifiers) {
            return importQualifiers(List.of(importQualifiers));
        }

        /**
         * @param name Name of the dashboard group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the dashboard group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param permissions [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) List of read and write permission configuration to specify which user, team, and organization can view and/or edit your dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
         * 
         * @return builder
         * 
         */
        public Builder permissions(@Nullable Output<List<DashboardGroupPermissionArgs>> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) List of read and write permission configuration to specify which user, team, and organization can view and/or edit your dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
         * 
         * @return builder
         * 
         */
        public Builder permissions(List<DashboardGroupPermissionArgs> permissions) {
            return permissions(Output.of(permissions));
        }

        /**
         * @param permissions [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) List of read and write permission configuration to specify which user, team, and organization can view and/or edit your dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
         * 
         * @return builder
         * 
         */
        public Builder permissions(DashboardGroupPermissionArgs... permissions) {
            return permissions(List.of(permissions));
        }

        /**
         * @param teams Team IDs to associate the dashboard group to.
         * 
         * @return builder
         * 
         */
        public Builder teams(@Nullable Output<List<String>> teams) {
            $.teams = teams;
            return this;
        }

        /**
         * @param teams Team IDs to associate the dashboard group to.
         * 
         * @return builder
         * 
         */
        public Builder teams(List<String> teams) {
            return teams(Output.of(teams));
        }

        /**
         * @param teams Team IDs to associate the dashboard group to.
         * 
         * @return builder
         * 
         */
        public Builder teams(String... teams) {
            return teams(List.of(teams));
        }

        public DashboardGroupArgs build() {
            return $;
        }
    }

}
