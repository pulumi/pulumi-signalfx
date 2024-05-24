// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.inputs.DashboardChartArgs;
import com.pulumi.signalfx.inputs.DashboardColumnArgs;
import com.pulumi.signalfx.inputs.DashboardEventOverlayArgs;
import com.pulumi.signalfx.inputs.DashboardFilterArgs;
import com.pulumi.signalfx.inputs.DashboardGridArgs;
import com.pulumi.signalfx.inputs.DashboardPermissionsArgs;
import com.pulumi.signalfx.inputs.DashboardSelectedEventOverlayArgs;
import com.pulumi.signalfx.inputs.DashboardVariableArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardState extends com.pulumi.resources.ResourceArgs {

    public static final DashboardState Empty = new DashboardState();

    /**
     * Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     * @deprecated
     * Please use permissions_* fields now
     * 
     */
    @Deprecated /* Please use permissions_* fields now */
    @Import(name="authorizedWriterTeams")
    private @Nullable Output<List<String>> authorizedWriterTeams;

    /**
     * @return Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     * @deprecated
     * Please use permissions_* fields now
     * 
     */
    @Deprecated /* Please use permissions_* fields now */
    public Optional<Output<List<String>>> authorizedWriterTeams() {
        return Optional.ofNullable(this.authorizedWriterTeams);
    }

    /**
     * User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     * @deprecated
     * Please use permissions fields now
     * 
     */
    @Deprecated /* Please use permissions fields now */
    @Import(name="authorizedWriterUsers")
    private @Nullable Output<List<String>> authorizedWriterUsers;

    /**
     * @return User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     * @deprecated
     * Please use permissions fields now
     * 
     */
    @Deprecated /* Please use permissions fields now */
    public Optional<Output<List<String>>> authorizedWriterUsers() {
        return Optional.ofNullable(this.authorizedWriterUsers);
    }

    /**
     * Chart ID and layout information for the charts in the dashboard.
     * 
     */
    @Import(name="charts")
    private @Nullable Output<List<DashboardChartArgs>> charts;

    /**
     * @return Chart ID and layout information for the charts in the dashboard.
     * 
     */
    public Optional<Output<List<DashboardChartArgs>>> charts() {
        return Optional.ofNullable(this.charts);
    }

    /**
     * Specifies the chart data display resolution for charts in this dashboard. Value can be one of `&#34;default&#34;`,  `&#34;low&#34;`, `&#34;high&#34;`, or  `&#34;highest&#34;`.
     * 
     */
    @Import(name="chartsResolution")
    private @Nullable Output<String> chartsResolution;

    /**
     * @return Specifies the chart data display resolution for charts in this dashboard. Value can be one of `&#34;default&#34;`,  `&#34;low&#34;`, `&#34;high&#34;`, or  `&#34;highest&#34;`.
     * 
     */
    public Optional<Output<String>> chartsResolution() {
        return Optional.ofNullable(this.chartsResolution);
    }

    /**
     * Column layout. Charts listed will be placed in a single column with the same width and height.
     * 
     */
    @Import(name="columns")
    private @Nullable Output<List<DashboardColumnArgs>> columns;

    /**
     * @return Column layout. Charts listed will be placed in a single column with the same width and height.
     * 
     */
    public Optional<Output<List<DashboardColumnArgs>>> columns() {
        return Optional.ofNullable(this.columns);
    }

    /**
     * The ID of the dashboard group that contains the dashboard.
     * 
     */
    @Import(name="dashboardGroup")
    private @Nullable Output<String> dashboardGroup;

    /**
     * @return The ID of the dashboard group that contains the dashboard.
     * 
     */
    public Optional<Output<String>> dashboardGroup() {
        return Optional.ofNullable(this.dashboardGroup);
    }

    /**
     * Description of the dashboard.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the dashboard.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="discoveryOptionsQuery")
    private @Nullable Output<String> discoveryOptionsQuery;

    public Optional<Output<String>> discoveryOptionsQuery() {
        return Optional.ofNullable(this.discoveryOptionsQuery);
    }

    @Import(name="discoveryOptionsSelectors")
    private @Nullable Output<List<String>> discoveryOptionsSelectors;

    public Optional<Output<List<String>>> discoveryOptionsSelectors() {
        return Optional.ofNullable(this.discoveryOptionsSelectors);
    }

    /**
     * Seconds since epoch. Used for visualization.
     * 
     */
    @Import(name="endTime")
    private @Nullable Output<Integer> endTime;

    /**
     * @return Seconds since epoch. Used for visualization.
     * 
     */
    public Optional<Output<Integer>> endTime() {
        return Optional.ofNullable(this.endTime);
    }

    /**
     * Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they&#39;re not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
     * 
     */
    @Import(name="eventOverlays")
    private @Nullable Output<List<DashboardEventOverlayArgs>> eventOverlays;

    /**
     * @return Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they&#39;re not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
     * 
     */
    public Optional<Output<List<DashboardEventOverlayArgs>>> eventOverlays() {
        return Optional.ofNullable(this.eventOverlays);
    }

    /**
     * Filter to apply to the charts when displaying the dashboard.
     * 
     */
    @Import(name="filters")
    private @Nullable Output<List<DashboardFilterArgs>> filters;

    /**
     * @return Filter to apply to the charts when displaying the dashboard.
     * 
     */
    public Optional<Output<List<DashboardFilterArgs>>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
     * 
     */
    @Import(name="grids")
    private @Nullable Output<List<DashboardGridArgs>> grids;

    /**
     * @return Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
     * 
     */
    public Optional<Output<List<DashboardGridArgs>>> grids() {
        return Optional.ofNullable(this.grids);
    }

    /**
     * Name of the dashboard.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the dashboard.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) Controls who can view and/or edit your dashboard. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<DashboardPermissionsArgs> permissions;

    /**
     * @return [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) Controls who can view and/or edit your dashboard. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    public Optional<Output<DashboardPermissionsArgs>> permissions() {
        return Optional.ofNullable(this.permissions);
    }

    /**
     * Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
     * 
     */
    @Import(name="selectedEventOverlays")
    private @Nullable Output<List<DashboardSelectedEventOverlayArgs>> selectedEventOverlays;

    /**
     * @return Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
     * 
     */
    public Optional<Output<List<DashboardSelectedEventOverlayArgs>>> selectedEventOverlays() {
        return Optional.ofNullable(this.selectedEventOverlays);
    }

    /**
     * Seconds since epoch. Used for visualization.
     * 
     */
    @Import(name="startTime")
    private @Nullable Output<Integer> startTime;

    /**
     * @return Seconds since epoch. Used for visualization.
     * 
     */
    public Optional<Output<Integer>> startTime() {
        return Optional.ofNullable(this.startTime);
    }

    /**
     * Tags of the dashboard.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return Tags of the dashboard.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The time range prior to now to visualize. Splunk Observability Cloud time syntax (e.g. `&#34;-5m&#34;`, `&#34;-1h&#34;`).
     * 
     */
    @Import(name="timeRange")
    private @Nullable Output<String> timeRange;

    /**
     * @return The time range prior to now to visualize. Splunk Observability Cloud time syntax (e.g. `&#34;-5m&#34;`, `&#34;-1h&#34;`).
     * 
     */
    public Optional<Output<String>> timeRange() {
        return Optional.ofNullable(this.timeRange);
    }

    /**
     * The URL of the dashboard.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The URL of the dashboard.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * Dashboard variable to apply to each chart in the dashboard.
     * 
     */
    @Import(name="variables")
    private @Nullable Output<List<DashboardVariableArgs>> variables;

    /**
     * @return Dashboard variable to apply to each chart in the dashboard.
     * 
     */
    public Optional<Output<List<DashboardVariableArgs>>> variables() {
        return Optional.ofNullable(this.variables);
    }

    private DashboardState() {}

    private DashboardState(DashboardState $) {
        this.authorizedWriterTeams = $.authorizedWriterTeams;
        this.authorizedWriterUsers = $.authorizedWriterUsers;
        this.charts = $.charts;
        this.chartsResolution = $.chartsResolution;
        this.columns = $.columns;
        this.dashboardGroup = $.dashboardGroup;
        this.description = $.description;
        this.discoveryOptionsQuery = $.discoveryOptionsQuery;
        this.discoveryOptionsSelectors = $.discoveryOptionsSelectors;
        this.endTime = $.endTime;
        this.eventOverlays = $.eventOverlays;
        this.filters = $.filters;
        this.grids = $.grids;
        this.name = $.name;
        this.permissions = $.permissions;
        this.selectedEventOverlays = $.selectedEventOverlays;
        this.startTime = $.startTime;
        this.tags = $.tags;
        this.timeRange = $.timeRange;
        this.url = $.url;
        this.variables = $.variables;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardState $;

        public Builder() {
            $ = new DashboardState();
        }

        public Builder(DashboardState defaults) {
            $ = new DashboardState(Objects.requireNonNull(defaults));
        }

        /**
         * @param authorizedWriterTeams Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use permissions_* fields now
         * 
         */
        @Deprecated /* Please use permissions_* fields now */
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
         * Please use permissions_* fields now
         * 
         */
        @Deprecated /* Please use permissions_* fields now */
        public Builder authorizedWriterTeams(List<String> authorizedWriterTeams) {
            return authorizedWriterTeams(Output.of(authorizedWriterTeams));
        }

        /**
         * @param authorizedWriterTeams Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use permissions_* fields now
         * 
         */
        @Deprecated /* Please use permissions_* fields now */
        public Builder authorizedWriterTeams(String... authorizedWriterTeams) {
            return authorizedWriterTeams(List.of(authorizedWriterTeams));
        }

        /**
         * @param authorizedWriterUsers User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use permissions fields now
         * 
         */
        @Deprecated /* Please use permissions fields now */
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
         * Please use permissions fields now
         * 
         */
        @Deprecated /* Please use permissions fields now */
        public Builder authorizedWriterUsers(List<String> authorizedWriterUsers) {
            return authorizedWriterUsers(Output.of(authorizedWriterUsers));
        }

        /**
         * @param authorizedWriterUsers User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use permissions fields now
         * 
         */
        @Deprecated /* Please use permissions fields now */
        public Builder authorizedWriterUsers(String... authorizedWriterUsers) {
            return authorizedWriterUsers(List.of(authorizedWriterUsers));
        }

        /**
         * @param charts Chart ID and layout information for the charts in the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder charts(@Nullable Output<List<DashboardChartArgs>> charts) {
            $.charts = charts;
            return this;
        }

        /**
         * @param charts Chart ID and layout information for the charts in the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder charts(List<DashboardChartArgs> charts) {
            return charts(Output.of(charts));
        }

        /**
         * @param charts Chart ID and layout information for the charts in the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder charts(DashboardChartArgs... charts) {
            return charts(List.of(charts));
        }

        /**
         * @param chartsResolution Specifies the chart data display resolution for charts in this dashboard. Value can be one of `&#34;default&#34;`,  `&#34;low&#34;`, `&#34;high&#34;`, or  `&#34;highest&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder chartsResolution(@Nullable Output<String> chartsResolution) {
            $.chartsResolution = chartsResolution;
            return this;
        }

        /**
         * @param chartsResolution Specifies the chart data display resolution for charts in this dashboard. Value can be one of `&#34;default&#34;`,  `&#34;low&#34;`, `&#34;high&#34;`, or  `&#34;highest&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder chartsResolution(String chartsResolution) {
            return chartsResolution(Output.of(chartsResolution));
        }

        /**
         * @param columns Column layout. Charts listed will be placed in a single column with the same width and height.
         * 
         * @return builder
         * 
         */
        public Builder columns(@Nullable Output<List<DashboardColumnArgs>> columns) {
            $.columns = columns;
            return this;
        }

        /**
         * @param columns Column layout. Charts listed will be placed in a single column with the same width and height.
         * 
         * @return builder
         * 
         */
        public Builder columns(List<DashboardColumnArgs> columns) {
            return columns(Output.of(columns));
        }

        /**
         * @param columns Column layout. Charts listed will be placed in a single column with the same width and height.
         * 
         * @return builder
         * 
         */
        public Builder columns(DashboardColumnArgs... columns) {
            return columns(List.of(columns));
        }

        /**
         * @param dashboardGroup The ID of the dashboard group that contains the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder dashboardGroup(@Nullable Output<String> dashboardGroup) {
            $.dashboardGroup = dashboardGroup;
            return this;
        }

        /**
         * @param dashboardGroup The ID of the dashboard group that contains the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder dashboardGroup(String dashboardGroup) {
            return dashboardGroup(Output.of(dashboardGroup));
        }

        /**
         * @param description Description of the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder discoveryOptionsQuery(@Nullable Output<String> discoveryOptionsQuery) {
            $.discoveryOptionsQuery = discoveryOptionsQuery;
            return this;
        }

        public Builder discoveryOptionsQuery(String discoveryOptionsQuery) {
            return discoveryOptionsQuery(Output.of(discoveryOptionsQuery));
        }

        public Builder discoveryOptionsSelectors(@Nullable Output<List<String>> discoveryOptionsSelectors) {
            $.discoveryOptionsSelectors = discoveryOptionsSelectors;
            return this;
        }

        public Builder discoveryOptionsSelectors(List<String> discoveryOptionsSelectors) {
            return discoveryOptionsSelectors(Output.of(discoveryOptionsSelectors));
        }

        public Builder discoveryOptionsSelectors(String... discoveryOptionsSelectors) {
            return discoveryOptionsSelectors(List.of(discoveryOptionsSelectors));
        }

        /**
         * @param endTime Seconds since epoch. Used for visualization.
         * 
         * @return builder
         * 
         */
        public Builder endTime(@Nullable Output<Integer> endTime) {
            $.endTime = endTime;
            return this;
        }

        /**
         * @param endTime Seconds since epoch. Used for visualization.
         * 
         * @return builder
         * 
         */
        public Builder endTime(Integer endTime) {
            return endTime(Output.of(endTime));
        }

        /**
         * @param eventOverlays Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they&#39;re not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
         * 
         * @return builder
         * 
         */
        public Builder eventOverlays(@Nullable Output<List<DashboardEventOverlayArgs>> eventOverlays) {
            $.eventOverlays = eventOverlays;
            return this;
        }

        /**
         * @param eventOverlays Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they&#39;re not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
         * 
         * @return builder
         * 
         */
        public Builder eventOverlays(List<DashboardEventOverlayArgs> eventOverlays) {
            return eventOverlays(Output.of(eventOverlays));
        }

        /**
         * @param eventOverlays Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they&#39;re not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
         * 
         * @return builder
         * 
         */
        public Builder eventOverlays(DashboardEventOverlayArgs... eventOverlays) {
            return eventOverlays(List.of(eventOverlays));
        }

        /**
         * @param filters Filter to apply to the charts when displaying the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable Output<List<DashboardFilterArgs>> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Filter to apply to the charts when displaying the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder filters(List<DashboardFilterArgs> filters) {
            return filters(Output.of(filters));
        }

        /**
         * @param filters Filter to apply to the charts when displaying the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder filters(DashboardFilterArgs... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param grids Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
         * 
         * @return builder
         * 
         */
        public Builder grids(@Nullable Output<List<DashboardGridArgs>> grids) {
            $.grids = grids;
            return this;
        }

        /**
         * @param grids Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
         * 
         * @return builder
         * 
         */
        public Builder grids(List<DashboardGridArgs> grids) {
            return grids(Output.of(grids));
        }

        /**
         * @param grids Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
         * 
         * @return builder
         * 
         */
        public Builder grids(DashboardGridArgs... grids) {
            return grids(List.of(grids));
        }

        /**
         * @param name Name of the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param permissions [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) Controls who can view and/or edit your dashboard. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
         * 
         * @return builder
         * 
         */
        public Builder permissions(@Nullable Output<DashboardPermissionsArgs> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) Controls who can view and/or edit your dashboard. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
         * 
         * @return builder
         * 
         */
        public Builder permissions(DashboardPermissionsArgs permissions) {
            return permissions(Output.of(permissions));
        }

        /**
         * @param selectedEventOverlays Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
         * 
         * @return builder
         * 
         */
        public Builder selectedEventOverlays(@Nullable Output<List<DashboardSelectedEventOverlayArgs>> selectedEventOverlays) {
            $.selectedEventOverlays = selectedEventOverlays;
            return this;
        }

        /**
         * @param selectedEventOverlays Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
         * 
         * @return builder
         * 
         */
        public Builder selectedEventOverlays(List<DashboardSelectedEventOverlayArgs> selectedEventOverlays) {
            return selectedEventOverlays(Output.of(selectedEventOverlays));
        }

        /**
         * @param selectedEventOverlays Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
         * 
         * @return builder
         * 
         */
        public Builder selectedEventOverlays(DashboardSelectedEventOverlayArgs... selectedEventOverlays) {
            return selectedEventOverlays(List.of(selectedEventOverlays));
        }

        /**
         * @param startTime Seconds since epoch. Used for visualization.
         * 
         * @return builder
         * 
         */
        public Builder startTime(@Nullable Output<Integer> startTime) {
            $.startTime = startTime;
            return this;
        }

        /**
         * @param startTime Seconds since epoch. Used for visualization.
         * 
         * @return builder
         * 
         */
        public Builder startTime(Integer startTime) {
            return startTime(Output.of(startTime));
        }

        /**
         * @param tags Tags of the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags of the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags Tags of the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param timeRange The time range prior to now to visualize. Splunk Observability Cloud time syntax (e.g. `&#34;-5m&#34;`, `&#34;-1h&#34;`).
         * 
         * @return builder
         * 
         */
        public Builder timeRange(@Nullable Output<String> timeRange) {
            $.timeRange = timeRange;
            return this;
        }

        /**
         * @param timeRange The time range prior to now to visualize. Splunk Observability Cloud time syntax (e.g. `&#34;-5m&#34;`, `&#34;-1h&#34;`).
         * 
         * @return builder
         * 
         */
        public Builder timeRange(String timeRange) {
            return timeRange(Output.of(timeRange));
        }

        /**
         * @param url The URL of the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL of the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param variables Dashboard variable to apply to each chart in the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder variables(@Nullable Output<List<DashboardVariableArgs>> variables) {
            $.variables = variables;
            return this;
        }

        /**
         * @param variables Dashboard variable to apply to each chart in the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder variables(List<DashboardVariableArgs> variables) {
            return variables(Output.of(variables));
        }

        /**
         * @param variables Dashboard variable to apply to each chart in the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder variables(DashboardVariableArgs... variables) {
            return variables(List.of(variables));
        }

        public DashboardState build() {
            return $;
        }
    }

}
