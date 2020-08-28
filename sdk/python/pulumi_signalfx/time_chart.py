# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['TimeChart']


class TimeChart(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 axes_include_zero: Optional[pulumi.Input[bool]] = None,
                 axes_precision: Optional[pulumi.Input[float]] = None,
                 axis_left: Optional[pulumi.Input[pulumi.InputType['TimeChartAxisLeftArgs']]] = None,
                 axis_right: Optional[pulumi.Input[pulumi.InputType['TimeChartAxisRightArgs']]] = None,
                 color_by: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_sampling: Optional[pulumi.Input[bool]] = None,
                 end_time: Optional[pulumi.Input[float]] = None,
                 event_options: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartEventOptionArgs']]]]] = None,
                 histogram_options: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartHistogramOptionArgs']]]]] = None,
                 legend_fields_to_hides: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 legend_options_fields: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartLegendOptionsFieldArgs']]]]] = None,
                 max_delay: Optional[pulumi.Input[float]] = None,
                 minimum_resolution: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 on_chart_legend_dimension: Optional[pulumi.Input[str]] = None,
                 plot_type: Optional[pulumi.Input[str]] = None,
                 program_text: Optional[pulumi.Input[str]] = None,
                 show_data_markers: Optional[pulumi.Input[bool]] = None,
                 show_event_lines: Optional[pulumi.Input[bool]] = None,
                 stacked: Optional[pulumi.Input[bool]] = None,
                 start_time: Optional[pulumi.Input[float]] = None,
                 tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 time_range: Optional[pulumi.Input[float]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 unit_prefix: Optional[pulumi.Input[str]] = None,
                 viz_options: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartVizOptionArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a SignalFx time chart resource. This can be used to create and manage the different types of time charts.

        Time charts display data points over a period of time.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        mychart0 = signalfx.TimeChart("mychart0",
            axis_left=signalfx.TimeChartAxisLeftArgs(
                label="CPU Total Idle",
                low_watermark=1000,
            ),
            legend_options_fields=[
                signalfx.TimeChartLegendOptionsFieldArgs(
                    enabled=False,
                    property="collector",
                ),
                signalfx.TimeChartLegendOptionsFieldArgs(
                    enabled=False,
                    property="hostname",
                ),
            ],
            plot_type="LineChart",
            program_text=\"\"\"data("cpu.total.idle").publish(label="CPU Idle")

        \"\"\",
            show_data_markers=True,
            time_range=3600,
            viz_options=[signalfx.TimeChartVizOptionArgs(
                axis="left",
                color="orange",
                label="CPU Idle",
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] axes_include_zero: Force the chart to display zero on the y-axes, even if none of the data is near zero.
        :param pulumi.Input[float] axes_precision: Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
        :param pulumi.Input[pulumi.InputType['TimeChartAxisLeftArgs']] axis_left: Set of axis options.
        :param pulumi.Input[pulumi.InputType['TimeChartAxisRightArgs']] axis_right: Set of axis options.
        :param pulumi.Input[str] color_by: Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
        :param pulumi.Input[str] description: Description of the chart.
        :param pulumi.Input[bool] disable_sampling: If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
        :param pulumi.Input[float] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartEventOptionArgs']]]] event_options: Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartHistogramOptionArgs']]]] histogram_options: Only used when `plot_type` is `"Histogram"`. Histogram specific options.
        :param pulumi.Input[List[pulumi.Input[str]]] legend_fields_to_hides: List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartLegendOptionsFieldArgs']]]] legend_options_fields: List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
        :param pulumi.Input[float] max_delay: How long (in seconds) to wait for late datapoints.
        :param pulumi.Input[float] minimum_resolution: The minimum resolution (in seconds) to use for computing the underlying program.
        :param pulumi.Input[str] name: Name of the chart.
        :param pulumi.Input[str] on_chart_legend_dimension: Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plot_label"` and any dimension.
        :param pulumi.Input[str] plot_type: The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
        :param pulumi.Input[str] program_text: Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        :param pulumi.Input[bool] show_data_markers: Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
        :param pulumi.Input[bool] show_event_lines: Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
        :param pulumi.Input[bool] stacked: Whether area and bar charts in the visualization should be stacked. `false` by default.
        :param pulumi.Input[float] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: Tags associated with the chart
        :param pulumi.Input[float] time_range: How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
        :param pulumi.Input[str] timezone: Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
        :param pulumi.Input[str] unit_prefix: Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartVizOptionArgs']]]] viz_options: Plot-level customization options, associated with a publish statement.
        """
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

            __props__['axes_include_zero'] = axes_include_zero
            __props__['axes_precision'] = axes_precision
            __props__['axis_left'] = axis_left
            __props__['axis_right'] = axis_right
            __props__['color_by'] = color_by
            __props__['description'] = description
            __props__['disable_sampling'] = disable_sampling
            __props__['end_time'] = end_time
            __props__['event_options'] = event_options
            __props__['histogram_options'] = histogram_options
            if legend_fields_to_hides is not None:
                warnings.warn("Please use legend_options_fields", DeprecationWarning)
                pulumi.log.warn("legend_fields_to_hides is deprecated: Please use legend_options_fields")
            __props__['legend_fields_to_hides'] = legend_fields_to_hides
            __props__['legend_options_fields'] = legend_options_fields
            __props__['max_delay'] = max_delay
            __props__['minimum_resolution'] = minimum_resolution
            __props__['name'] = name
            __props__['on_chart_legend_dimension'] = on_chart_legend_dimension
            __props__['plot_type'] = plot_type
            if program_text is None:
                raise TypeError("Missing required property 'program_text'")
            __props__['program_text'] = program_text
            __props__['show_data_markers'] = show_data_markers
            __props__['show_event_lines'] = show_event_lines
            __props__['stacked'] = stacked
            __props__['start_time'] = start_time
            if tags is not None:
                warnings.warn("signalfx_time_chart.tags is being removed in the next release", DeprecationWarning)
                pulumi.log.warn("tags is deprecated: signalfx_time_chart.tags is being removed in the next release")
            __props__['tags'] = tags
            __props__['time_range'] = time_range
            __props__['timezone'] = timezone
            __props__['unit_prefix'] = unit_prefix
            __props__['viz_options'] = viz_options
            __props__['url'] = None
        super(TimeChart, __self__).__init__(
            'signalfx:index/timeChart:TimeChart',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            axes_include_zero: Optional[pulumi.Input[bool]] = None,
            axes_precision: Optional[pulumi.Input[float]] = None,
            axis_left: Optional[pulumi.Input[pulumi.InputType['TimeChartAxisLeftArgs']]] = None,
            axis_right: Optional[pulumi.Input[pulumi.InputType['TimeChartAxisRightArgs']]] = None,
            color_by: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disable_sampling: Optional[pulumi.Input[bool]] = None,
            end_time: Optional[pulumi.Input[float]] = None,
            event_options: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartEventOptionArgs']]]]] = None,
            histogram_options: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartHistogramOptionArgs']]]]] = None,
            legend_fields_to_hides: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            legend_options_fields: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartLegendOptionsFieldArgs']]]]] = None,
            max_delay: Optional[pulumi.Input[float]] = None,
            minimum_resolution: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            on_chart_legend_dimension: Optional[pulumi.Input[str]] = None,
            plot_type: Optional[pulumi.Input[str]] = None,
            program_text: Optional[pulumi.Input[str]] = None,
            show_data_markers: Optional[pulumi.Input[bool]] = None,
            show_event_lines: Optional[pulumi.Input[bool]] = None,
            stacked: Optional[pulumi.Input[bool]] = None,
            start_time: Optional[pulumi.Input[float]] = None,
            tags: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            time_range: Optional[pulumi.Input[float]] = None,
            timezone: Optional[pulumi.Input[str]] = None,
            unit_prefix: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            viz_options: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartVizOptionArgs']]]]] = None) -> 'TimeChart':
        """
        Get an existing TimeChart resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] axes_include_zero: Force the chart to display zero on the y-axes, even if none of the data is near zero.
        :param pulumi.Input[float] axes_precision: Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
        :param pulumi.Input[pulumi.InputType['TimeChartAxisLeftArgs']] axis_left: Set of axis options.
        :param pulumi.Input[pulumi.InputType['TimeChartAxisRightArgs']] axis_right: Set of axis options.
        :param pulumi.Input[str] color_by: Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
        :param pulumi.Input[str] description: Description of the chart.
        :param pulumi.Input[bool] disable_sampling: If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
        :param pulumi.Input[float] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartEventOptionArgs']]]] event_options: Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartHistogramOptionArgs']]]] histogram_options: Only used when `plot_type` is `"Histogram"`. Histogram specific options.
        :param pulumi.Input[List[pulumi.Input[str]]] legend_fields_to_hides: List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartLegendOptionsFieldArgs']]]] legend_options_fields: List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
        :param pulumi.Input[float] max_delay: How long (in seconds) to wait for late datapoints.
        :param pulumi.Input[float] minimum_resolution: The minimum resolution (in seconds) to use for computing the underlying program.
        :param pulumi.Input[str] name: Name of the chart.
        :param pulumi.Input[str] on_chart_legend_dimension: Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plot_label"` and any dimension.
        :param pulumi.Input[str] plot_type: The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
        :param pulumi.Input[str] program_text: Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        :param pulumi.Input[bool] show_data_markers: Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
        :param pulumi.Input[bool] show_event_lines: Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
        :param pulumi.Input[bool] stacked: Whether area and bar charts in the visualization should be stacked. `false` by default.
        :param pulumi.Input[float] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[List[pulumi.Input[str]]] tags: Tags associated with the chart
        :param pulumi.Input[float] time_range: How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
        :param pulumi.Input[str] timezone: Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
        :param pulumi.Input[str] unit_prefix: Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        :param pulumi.Input[str] url: URL of the chart
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TimeChartVizOptionArgs']]]] viz_options: Plot-level customization options, associated with a publish statement.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["axes_include_zero"] = axes_include_zero
        __props__["axes_precision"] = axes_precision
        __props__["axis_left"] = axis_left
        __props__["axis_right"] = axis_right
        __props__["color_by"] = color_by
        __props__["description"] = description
        __props__["disable_sampling"] = disable_sampling
        __props__["end_time"] = end_time
        __props__["event_options"] = event_options
        __props__["histogram_options"] = histogram_options
        __props__["legend_fields_to_hides"] = legend_fields_to_hides
        __props__["legend_options_fields"] = legend_options_fields
        __props__["max_delay"] = max_delay
        __props__["minimum_resolution"] = minimum_resolution
        __props__["name"] = name
        __props__["on_chart_legend_dimension"] = on_chart_legend_dimension
        __props__["plot_type"] = plot_type
        __props__["program_text"] = program_text
        __props__["show_data_markers"] = show_data_markers
        __props__["show_event_lines"] = show_event_lines
        __props__["stacked"] = stacked
        __props__["start_time"] = start_time
        __props__["tags"] = tags
        __props__["time_range"] = time_range
        __props__["timezone"] = timezone
        __props__["unit_prefix"] = unit_prefix
        __props__["url"] = url
        __props__["viz_options"] = viz_options
        return TimeChart(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="axesIncludeZero")
    def axes_include_zero(self) -> pulumi.Output[Optional[bool]]:
        """
        Force the chart to display zero on the y-axes, even if none of the data is near zero.
        """
        return pulumi.get(self, "axes_include_zero")

    @property
    @pulumi.getter(name="axesPrecision")
    def axes_precision(self) -> pulumi.Output[Optional[float]]:
        """
        Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
        """
        return pulumi.get(self, "axes_precision")

    @property
    @pulumi.getter(name="axisLeft")
    def axis_left(self) -> pulumi.Output[Optional['outputs.TimeChartAxisLeft']]:
        """
        Set of axis options.
        """
        return pulumi.get(self, "axis_left")

    @property
    @pulumi.getter(name="axisRight")
    def axis_right(self) -> pulumi.Output[Optional['outputs.TimeChartAxisRight']]:
        """
        Set of axis options.
        """
        return pulumi.get(self, "axis_right")

    @property
    @pulumi.getter(name="colorBy")
    def color_by(self) -> pulumi.Output[Optional[str]]:
        """
        Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
        """
        return pulumi.get(self, "color_by")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the chart.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableSampling")
    def disable_sampling(self) -> pulumi.Output[Optional[bool]]:
        """
        If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
        """
        return pulumi.get(self, "disable_sampling")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[float]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="eventOptions")
    def event_options(self) -> pulumi.Output[Optional[List['outputs.TimeChartEventOption']]]:
        """
        Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
        """
        return pulumi.get(self, "event_options")

    @property
    @pulumi.getter(name="histogramOptions")
    def histogram_options(self) -> pulumi.Output[Optional[List['outputs.TimeChartHistogramOption']]]:
        """
        Only used when `plot_type` is `"Histogram"`. Histogram specific options.
        """
        return pulumi.get(self, "histogram_options")

    @property
    @pulumi.getter(name="legendFieldsToHides")
    def legend_fields_to_hides(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
        """
        return pulumi.get(self, "legend_fields_to_hides")

    @property
    @pulumi.getter(name="legendOptionsFields")
    def legend_options_fields(self) -> pulumi.Output[Optional[List['outputs.TimeChartLegendOptionsField']]]:
        """
        List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
        """
        return pulumi.get(self, "legend_options_fields")

    @property
    @pulumi.getter(name="maxDelay")
    def max_delay(self) -> pulumi.Output[Optional[float]]:
        """
        How long (in seconds) to wait for late datapoints.
        """
        return pulumi.get(self, "max_delay")

    @property
    @pulumi.getter(name="minimumResolution")
    def minimum_resolution(self) -> pulumi.Output[Optional[float]]:
        """
        The minimum resolution (in seconds) to use for computing the underlying program.
        """
        return pulumi.get(self, "minimum_resolution")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the chart.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="onChartLegendDimension")
    def on_chart_legend_dimension(self) -> pulumi.Output[Optional[str]]:
        """
        Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plot_label"` and any dimension.
        """
        return pulumi.get(self, "on_chart_legend_dimension")

    @property
    @pulumi.getter(name="plotType")
    def plot_type(self) -> pulumi.Output[Optional[str]]:
        """
        The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
        """
        return pulumi.get(self, "plot_type")

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> pulumi.Output[str]:
        """
        Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        """
        return pulumi.get(self, "program_text")

    @property
    @pulumi.getter(name="showDataMarkers")
    def show_data_markers(self) -> pulumi.Output[Optional[bool]]:
        """
        Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
        """
        return pulumi.get(self, "show_data_markers")

    @property
    @pulumi.getter(name="showEventLines")
    def show_event_lines(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
        """
        return pulumi.get(self, "show_event_lines")

    @property
    @pulumi.getter
    def stacked(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether area and bar charts in the visualization should be stacked. `false` by default.
        """
        return pulumi.get(self, "stacked")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[float]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[List[str]]]:
        """
        Tags associated with the chart
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> pulumi.Output[Optional[float]]:
        """
        How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
        """
        return pulumi.get(self, "time_range")

    @property
    @pulumi.getter
    def timezone(self) -> pulumi.Output[Optional[str]]:
        """
        Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
        """
        return pulumi.get(self, "timezone")

    @property
    @pulumi.getter(name="unitPrefix")
    def unit_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        """
        return pulumi.get(self, "unit_prefix")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL of the chart
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="vizOptions")
    def viz_options(self) -> pulumi.Output[Optional[List['outputs.TimeChartVizOption']]]:
        """
        Plot-level customization options, associated with a publish statement.
        """
        return pulumi.get(self, "viz_options")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

