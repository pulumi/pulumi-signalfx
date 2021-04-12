# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DashboardArgs', 'Dashboard']

@pulumi.input_type
class DashboardArgs:
    def __init__(__self__, *,
                 dashboard_group: pulumi.Input[str],
                 authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 charts: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardChartArgs']]]] = None,
                 charts_resolution: Optional[pulumi.Input[str]] = None,
                 columns: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardColumnArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 discovery_options_query: Optional[pulumi.Input[str]] = None,
                 discovery_options_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 end_time: Optional[pulumi.Input[int]] = None,
                 event_overlays: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardEventOverlayArgs']]]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardFilterArgs']]]] = None,
                 grids: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGridArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 selected_event_overlays: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardSelectedEventOverlayArgs']]]] = None,
                 start_time: Optional[pulumi.Input[int]] = None,
                 time_range: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardVariableArgs']]]] = None):
        """
        The set of arguments for constructing a Dashboard resource.
        :param pulumi.Input[str] dashboard_group: The ID of the dashboard group that contains the dashboard.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_teams: Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_users: User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input['DashboardChartArgs']]] charts: Chart ID and layout information for the charts in the dashboard.
        :param pulumi.Input[str] charts_resolution: Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
        :param pulumi.Input[Sequence[pulumi.Input['DashboardColumnArgs']]] columns: Column number for the layout.
        :param pulumi.Input[str] description: Variable description.
        :param pulumi.Input[int] end_time: Seconds since epoch. Used for visualization.
        :param pulumi.Input[Sequence[pulumi.Input['DashboardEventOverlayArgs']]] event_overlays: Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
        :param pulumi.Input[Sequence[pulumi.Input['DashboardFilterArgs']]] filters: Filter to apply to the charts when displaying the dashboard.
        :param pulumi.Input[Sequence[pulumi.Input['DashboardGridArgs']]] grids: Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
        :param pulumi.Input[str] name: Name of the dashboard.
        :param pulumi.Input[Sequence[pulumi.Input['DashboardSelectedEventOverlayArgs']]] selected_event_overlays: Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
        :param pulumi.Input[int] start_time: Seconds since epoch. Used for visualization.
        :param pulumi.Input[str] time_range: The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
        :param pulumi.Input[Sequence[pulumi.Input['DashboardVariableArgs']]] variables: Dashboard variable to apply to each chart in the dashboard.
        """
        pulumi.set(__self__, "dashboard_group", dashboard_group)
        if authorized_writer_teams is not None:
            pulumi.set(__self__, "authorized_writer_teams", authorized_writer_teams)
        if authorized_writer_users is not None:
            pulumi.set(__self__, "authorized_writer_users", authorized_writer_users)
        if charts is not None:
            pulumi.set(__self__, "charts", charts)
        if charts_resolution is not None:
            pulumi.set(__self__, "charts_resolution", charts_resolution)
        if columns is not None:
            pulumi.set(__self__, "columns", columns)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if discovery_options_query is not None:
            pulumi.set(__self__, "discovery_options_query", discovery_options_query)
        if discovery_options_selectors is not None:
            pulumi.set(__self__, "discovery_options_selectors", discovery_options_selectors)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if event_overlays is not None:
            pulumi.set(__self__, "event_overlays", event_overlays)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if grids is not None:
            pulumi.set(__self__, "grids", grids)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if selected_event_overlays is not None:
            pulumi.set(__self__, "selected_event_overlays", selected_event_overlays)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if time_range is not None:
            pulumi.set(__self__, "time_range", time_range)
        if variables is not None:
            pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter(name="dashboardGroup")
    def dashboard_group(self) -> pulumi.Input[str]:
        """
        The ID of the dashboard group that contains the dashboard.
        """
        return pulumi.get(self, "dashboard_group")

    @dashboard_group.setter
    def dashboard_group(self, value: pulumi.Input[str]):
        pulumi.set(self, "dashboard_group", value)

    @property
    @pulumi.getter(name="authorizedWriterTeams")
    def authorized_writer_teams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_teams")

    @authorized_writer_teams.setter
    def authorized_writer_teams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_writer_teams", value)

    @property
    @pulumi.getter(name="authorizedWriterUsers")
    def authorized_writer_users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_users")

    @authorized_writer_users.setter
    def authorized_writer_users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_writer_users", value)

    @property
    @pulumi.getter
    def charts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardChartArgs']]]]:
        """
        Chart ID and layout information for the charts in the dashboard.
        """
        return pulumi.get(self, "charts")

    @charts.setter
    def charts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardChartArgs']]]]):
        pulumi.set(self, "charts", value)

    @property
    @pulumi.getter(name="chartsResolution")
    def charts_resolution(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
        """
        return pulumi.get(self, "charts_resolution")

    @charts_resolution.setter
    def charts_resolution(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charts_resolution", value)

    @property
    @pulumi.getter
    def columns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardColumnArgs']]]]:
        """
        Column number for the layout.
        """
        return pulumi.get(self, "columns")

    @columns.setter
    def columns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardColumnArgs']]]]):
        pulumi.set(self, "columns", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Variable description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="discoveryOptionsQuery")
    def discovery_options_query(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "discovery_options_query")

    @discovery_options_query.setter
    def discovery_options_query(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "discovery_options_query", value)

    @property
    @pulumi.getter(name="discoveryOptionsSelectors")
    def discovery_options_selectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "discovery_options_selectors")

    @discovery_options_selectors.setter
    def discovery_options_selectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "discovery_options_selectors", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[int]]:
        """
        Seconds since epoch. Used for visualization.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter(name="eventOverlays")
    def event_overlays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardEventOverlayArgs']]]]:
        """
        Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
        """
        return pulumi.get(self, "event_overlays")

    @event_overlays.setter
    def event_overlays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardEventOverlayArgs']]]]):
        pulumi.set(self, "event_overlays", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardFilterArgs']]]]:
        """
        Filter to apply to the charts when displaying the dashboard.
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardFilterArgs']]]]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter
    def grids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGridArgs']]]]:
        """
        Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
        """
        return pulumi.get(self, "grids")

    @grids.setter
    def grids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGridArgs']]]]):
        pulumi.set(self, "grids", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the dashboard.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="selectedEventOverlays")
    def selected_event_overlays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardSelectedEventOverlayArgs']]]]:
        """
        Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
        """
        return pulumi.get(self, "selected_event_overlays")

    @selected_event_overlays.setter
    def selected_event_overlays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardSelectedEventOverlayArgs']]]]):
        pulumi.set(self, "selected_event_overlays", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[int]]:
        """
        Seconds since epoch. Used for visualization.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> Optional[pulumi.Input[str]]:
        """
        The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
        """
        return pulumi.get(self, "time_range")

    @time_range.setter
    def time_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_range", value)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardVariableArgs']]]]:
        """
        Dashboard variable to apply to each chart in the dashboard.
        """
        return pulumi.get(self, "variables")

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardVariableArgs']]]]):
        pulumi.set(self, "variables", value)


class Dashboard(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 charts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardChartArgs']]]]] = None,
                 charts_resolution: Optional[pulumi.Input[str]] = None,
                 columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardColumnArgs']]]]] = None,
                 dashboard_group: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 discovery_options_query: Optional[pulumi.Input[str]] = None,
                 discovery_options_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 end_time: Optional[pulumi.Input[int]] = None,
                 event_overlays: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardEventOverlayArgs']]]]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardFilterArgs']]]]] = None,
                 grids: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGridArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 selected_event_overlays: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardSelectedEventOverlayArgs']]]]] = None,
                 start_time: Optional[pulumi.Input[int]] = None,
                 time_range: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardVariableArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Dashboard resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_teams: Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_users: User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardChartArgs']]]] charts: Chart ID and layout information for the charts in the dashboard.
        :param pulumi.Input[str] charts_resolution: Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardColumnArgs']]]] columns: Column number for the layout.
        :param pulumi.Input[str] dashboard_group: The ID of the dashboard group that contains the dashboard.
        :param pulumi.Input[str] description: Variable description.
        :param pulumi.Input[int] end_time: Seconds since epoch. Used for visualization.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardEventOverlayArgs']]]] event_overlays: Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardFilterArgs']]]] filters: Filter to apply to the charts when displaying the dashboard.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGridArgs']]]] grids: Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
        :param pulumi.Input[str] name: Name of the dashboard.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardSelectedEventOverlayArgs']]]] selected_event_overlays: Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
        :param pulumi.Input[int] start_time: Seconds since epoch. Used for visualization.
        :param pulumi.Input[str] time_range: The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardVariableArgs']]]] variables: Dashboard variable to apply to each chart in the dashboard.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DashboardArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Dashboard resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DashboardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DashboardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 charts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardChartArgs']]]]] = None,
                 charts_resolution: Optional[pulumi.Input[str]] = None,
                 columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardColumnArgs']]]]] = None,
                 dashboard_group: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 discovery_options_query: Optional[pulumi.Input[str]] = None,
                 discovery_options_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 end_time: Optional[pulumi.Input[int]] = None,
                 event_overlays: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardEventOverlayArgs']]]]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardFilterArgs']]]]] = None,
                 grids: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGridArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 selected_event_overlays: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardSelectedEventOverlayArgs']]]]] = None,
                 start_time: Optional[pulumi.Input[int]] = None,
                 time_range: Optional[pulumi.Input[str]] = None,
                 variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardVariableArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['authorized_writer_teams'] = authorized_writer_teams
            __props__['authorized_writer_users'] = authorized_writer_users
            __props__['charts'] = charts
            __props__['charts_resolution'] = charts_resolution
            __props__['columns'] = columns
            if dashboard_group is None and not opts.urn:
                raise TypeError("Missing required property 'dashboard_group'")
            __props__['dashboard_group'] = dashboard_group
            __props__['description'] = description
            __props__['discovery_options_query'] = discovery_options_query
            __props__['discovery_options_selectors'] = discovery_options_selectors
            __props__['end_time'] = end_time
            __props__['event_overlays'] = event_overlays
            __props__['filters'] = filters
            __props__['grids'] = grids
            __props__['name'] = name
            __props__['selected_event_overlays'] = selected_event_overlays
            __props__['start_time'] = start_time
            __props__['time_range'] = time_range
            __props__['variables'] = variables
            __props__['url'] = None
        super(Dashboard, __self__).__init__(
            'signalfx:index/dashboard:Dashboard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            charts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardChartArgs']]]]] = None,
            charts_resolution: Optional[pulumi.Input[str]] = None,
            columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardColumnArgs']]]]] = None,
            dashboard_group: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            discovery_options_query: Optional[pulumi.Input[str]] = None,
            discovery_options_selectors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            end_time: Optional[pulumi.Input[int]] = None,
            event_overlays: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardEventOverlayArgs']]]]] = None,
            filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardFilterArgs']]]]] = None,
            grids: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGridArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            selected_event_overlays: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardSelectedEventOverlayArgs']]]]] = None,
            start_time: Optional[pulumi.Input[int]] = None,
            time_range: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            variables: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardVariableArgs']]]]] = None) -> 'Dashboard':
        """
        Get an existing Dashboard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_teams: Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_users: User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardChartArgs']]]] charts: Chart ID and layout information for the charts in the dashboard.
        :param pulumi.Input[str] charts_resolution: Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardColumnArgs']]]] columns: Column number for the layout.
        :param pulumi.Input[str] dashboard_group: The ID of the dashboard group that contains the dashboard.
        :param pulumi.Input[str] description: Variable description.
        :param pulumi.Input[int] end_time: Seconds since epoch. Used for visualization.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardEventOverlayArgs']]]] event_overlays: Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardFilterArgs']]]] filters: Filter to apply to the charts when displaying the dashboard.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGridArgs']]]] grids: Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
        :param pulumi.Input[str] name: Name of the dashboard.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardSelectedEventOverlayArgs']]]] selected_event_overlays: Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
        :param pulumi.Input[int] start_time: Seconds since epoch. Used for visualization.
        :param pulumi.Input[str] time_range: The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
        :param pulumi.Input[str] url: The URL of the dashboard.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardVariableArgs']]]] variables: Dashboard variable to apply to each chart in the dashboard.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authorized_writer_teams"] = authorized_writer_teams
        __props__["authorized_writer_users"] = authorized_writer_users
        __props__["charts"] = charts
        __props__["charts_resolution"] = charts_resolution
        __props__["columns"] = columns
        __props__["dashboard_group"] = dashboard_group
        __props__["description"] = description
        __props__["discovery_options_query"] = discovery_options_query
        __props__["discovery_options_selectors"] = discovery_options_selectors
        __props__["end_time"] = end_time
        __props__["event_overlays"] = event_overlays
        __props__["filters"] = filters
        __props__["grids"] = grids
        __props__["name"] = name
        __props__["selected_event_overlays"] = selected_event_overlays
        __props__["start_time"] = start_time
        __props__["time_range"] = time_range
        __props__["url"] = url
        __props__["variables"] = variables
        return Dashboard(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizedWriterTeams")
    def authorized_writer_teams(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_teams")

    @property
    @pulumi.getter(name="authorizedWriterUsers")
    def authorized_writer_users(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_users")

    @property
    @pulumi.getter
    def charts(self) -> pulumi.Output[Optional[Sequence['outputs.DashboardChart']]]:
        """
        Chart ID and layout information for the charts in the dashboard.
        """
        return pulumi.get(self, "charts")

    @property
    @pulumi.getter(name="chartsResolution")
    def charts_resolution(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
        """
        return pulumi.get(self, "charts_resolution")

    @property
    @pulumi.getter
    def columns(self) -> pulumi.Output[Optional[Sequence['outputs.DashboardColumn']]]:
        """
        Column number for the layout.
        """
        return pulumi.get(self, "columns")

    @property
    @pulumi.getter(name="dashboardGroup")
    def dashboard_group(self) -> pulumi.Output[str]:
        """
        The ID of the dashboard group that contains the dashboard.
        """
        return pulumi.get(self, "dashboard_group")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Variable description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="discoveryOptionsQuery")
    def discovery_options_query(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "discovery_options_query")

    @property
    @pulumi.getter(name="discoveryOptionsSelectors")
    def discovery_options_selectors(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "discovery_options_selectors")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[int]]:
        """
        Seconds since epoch. Used for visualization.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="eventOverlays")
    def event_overlays(self) -> pulumi.Output[Optional[Sequence['outputs.DashboardEventOverlay']]]:
        """
        Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selected_event_overlay` property instead.
        """
        return pulumi.get(self, "event_overlays")

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Optional[Sequence['outputs.DashboardFilter']]]:
        """
        Filter to apply to the charts when displaying the dashboard.
        """
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def grids(self) -> pulumi.Output[Optional[Sequence['outputs.DashboardGrid']]]:
        """
        Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
        """
        return pulumi.get(self, "grids")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the dashboard.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="selectedEventOverlays")
    def selected_event_overlays(self) -> pulumi.Output[Optional[Sequence['outputs.DashboardSelectedEventOverlay']]]:
        """
        Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `event_overlay`, which are similar to the properties here.
        """
        return pulumi.get(self, "selected_event_overlays")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[int]]:
        """
        Seconds since epoch. Used for visualization.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> pulumi.Output[Optional[str]]:
        """
        The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
        """
        return pulumi.get(self, "time_range")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the dashboard.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Optional[Sequence['outputs.DashboardVariable']]]:
        """
        Dashboard variable to apply to each chart in the dashboard.
        """
        return pulumi.get(self, "variables")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

