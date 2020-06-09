# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class TimeChart(pulumi.CustomResource):
    axes_include_zero: pulumi.Output[bool]
    """
    Force the chart to display zero on the y-axes, even if none of the data is near zero.
    """
    axes_precision: pulumi.Output[float]
    """
    Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
    """
    axis_left: pulumi.Output[dict]
    """
    Set of axis options.

      * `highWatermark` (`float`) - A line to draw as a high watermark.
      * `highWatermarkLabel` (`str`) - A label to attach to the high watermark line.
      * `label` (`str`) - Label used in the publish statement that displays the event query you want to customize.
      * `lowWatermark` (`float`) - A line to draw as a low watermark.
      * `lowWatermarkLabel` (`str`) - A label to attach to the low watermark line.
      * `maxValue` (`float`) - The maximum value for the right axis.
      * `minValue` (`float`) - The minimum value for the right axis.
      * `watermarks` (`list`)
        * `label` (`str`) - Label used in the publish statement that displays the event query you want to customize.
        * `value` (`float`)
    """
    axis_right: pulumi.Output[dict]
    """
    Set of axis options.

      * `highWatermark` (`float`) - A line to draw as a high watermark.
      * `highWatermarkLabel` (`str`) - A label to attach to the high watermark line.
      * `label` (`str`) - Label used in the publish statement that displays the event query you want to customize.
      * `lowWatermark` (`float`) - A line to draw as a low watermark.
      * `lowWatermarkLabel` (`str`) - A label to attach to the low watermark line.
      * `maxValue` (`float`) - The maximum value for the right axis.
      * `minValue` (`float`) - The minimum value for the right axis.
      * `watermarks` (`list`)
        * `label` (`str`) - Label used in the publish statement that displays the event query you want to customize.
        * `value` (`float`)
    """
    color_by: pulumi.Output[str]
    """
    Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
    """
    description: pulumi.Output[str]
    """
    Description of the chart.
    """
    disable_sampling: pulumi.Output[bool]
    """
    If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
    """
    end_time: pulumi.Output[float]
    """
    Seconds since epoch. Used for visualization. Conflicts with `time_range`.
    """
    event_options: pulumi.Output[list]
    """
    Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.

      * `color` (`str`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
      * `displayName` (`str`) - Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
      * `label` (`str`) - Label used in the publish statement that displays the event query you want to customize.
    """
    histogram_options: pulumi.Output[list]
    """
    Only used when `plot_type` is `"Histogram"`. Histogram specific options.

      * `colorTheme` (`str`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine, red, gold, greenyellow, chartreuse, jade
    """
    legend_fields_to_hides: pulumi.Output[list]
    """
    List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
    """
    legend_options_fields: pulumi.Output[list]
    """
    List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.

      * `enabled` (`bool`) - True or False depending on if you want the property to be shown or hidden.
      * `property` (`str`) - The name of the property to display. Note the special values of `plot_label` (corresponding with the API's `sf_metric`) which shows the label of the time series `publish()` and `metric` (corresponding with the API's `sf_originatingMetric`) that shows the name of the metric for the time series being displayed.
    """
    max_delay: pulumi.Output[float]
    """
    How long (in seconds) to wait for late datapoints.
    """
    minimum_resolution: pulumi.Output[float]
    """
    The minimum resolution (in seconds) to use for computing the underlying program.
    """
    name: pulumi.Output[str]
    """
    Name of the chart.
    """
    on_chart_legend_dimension: pulumi.Output[str]
    """
    Dimensions to show in the on-chart legend. On-chart legend is off unless a dimension is specified. Allowed: `"metric"`, `"plot_label"` and any dimension.
    """
    plot_type: pulumi.Output[str]
    """
    The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
    """
    program_text: pulumi.Output[str]
    """
    Signalflow program text for the chart. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
    """
    show_data_markers: pulumi.Output[bool]
    """
    Show markers (circles) for each datapoint used to draw line or area charts. `false` by default.
    """
    show_event_lines: pulumi.Output[bool]
    """
    Whether vertical highlight lines should be drawn in the visualizations at times when events occurred. `false` by default.
    """
    stacked: pulumi.Output[bool]
    """
    Whether area and bar charts in the visualization should be stacked. `false` by default.
    """
    start_time: pulumi.Output[float]
    """
    Seconds since epoch. Used for visualization. Conflicts with `time_range`.
    """
    tags: pulumi.Output[list]
    """
    Tags associated with the chart
    """
    time_range: pulumi.Output[float]
    """
    How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
    """
    timezone: pulumi.Output[str]
    """
    Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
    """
    unit_prefix: pulumi.Output[str]
    """
    Must be `"Metric"` or `"Binary`". `"Metric"` by default.
    """
    url: pulumi.Output[str]
    """
    URL of the chart
    """
    viz_options: pulumi.Output[list]
    """
    Plot-level customization options, associated with a publish statement.

      * `axis` (`str`) - Y-axis associated with values for this plot. Must be either `right` or `left`.
      * `color` (`str`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
      * `displayName` (`str`) - Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
      * `label` (`str`) - Label used in the publish statement that displays the event query you want to customize.
      * `plot_type` (`str`) - The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
      * `valuePrefix` (`str`)
      * `valueSuffix` (`str`)
      * `valueUnit` (`str`) - A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
        * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
    """
    def __init__(__self__, resource_name, opts=None, axes_include_zero=None, axes_precision=None, axis_left=None, axis_right=None, color_by=None, description=None, disable_sampling=None, end_time=None, event_options=None, histogram_options=None, legend_fields_to_hides=None, legend_options_fields=None, max_delay=None, minimum_resolution=None, name=None, on_chart_legend_dimension=None, plot_type=None, program_text=None, show_data_markers=None, show_event_lines=None, stacked=None, start_time=None, tags=None, time_range=None, timezone=None, unit_prefix=None, viz_options=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a SignalFx time chart resource. This can be used to create and manage the different types of time charts.

        Time charts display data points over a period of time.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        mychart0 = signalfx.TimeChart("mychart0",
            axis_left={
                "label": "CPU Total Idle",
                "lowWatermark": 1000,
            },
            legend_options_fields=[
                {
                    "enabled": False,
                    "property": "collector",
                },
                {
                    "enabled": False,
                    "property": "hostname",
                },
            ],
            plot_type="LineChart",
            program_text=\"\"\"data("cpu.total.idle").publish(label="CPU Idle")

        \"\"\",
            show_data_markers=True,
            time_range=3600,
            viz_options=[{
                "axis": "left",
                "color": "orange",
                "label": "CPU Idle",
            }])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] axes_include_zero: Force the chart to display zero on the y-axes, even if none of the data is near zero.
        :param pulumi.Input[float] axes_precision: Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
        :param pulumi.Input[dict] axis_left: Set of axis options.
        :param pulumi.Input[dict] axis_right: Set of axis options.
        :param pulumi.Input[str] color_by: Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
        :param pulumi.Input[str] description: Description of the chart.
        :param pulumi.Input[bool] disable_sampling: If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
        :param pulumi.Input[float] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[list] event_options: Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
        :param pulumi.Input[list] histogram_options: Only used when `plot_type` is `"Histogram"`. Histogram specific options.
        :param pulumi.Input[list] legend_fields_to_hides: List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
        :param pulumi.Input[list] legend_options_fields: List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
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
        :param pulumi.Input[list] tags: Tags associated with the chart
        :param pulumi.Input[float] time_range: How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
        :param pulumi.Input[str] timezone: Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
        :param pulumi.Input[str] unit_prefix: Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        :param pulumi.Input[list] viz_options: Plot-level customization options, associated with a publish statement.

        The **axis_left** object supports the following:

          * `highWatermark` (`pulumi.Input[float]`) - A line to draw as a high watermark.
          * `highWatermarkLabel` (`pulumi.Input[str]`) - A label to attach to the high watermark line.
          * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.
          * `lowWatermark` (`pulumi.Input[float]`) - A line to draw as a low watermark.
          * `lowWatermarkLabel` (`pulumi.Input[str]`) - A label to attach to the low watermark line.
          * `maxValue` (`pulumi.Input[float]`) - The maximum value for the right axis.
          * `minValue` (`pulumi.Input[float]`) - The minimum value for the right axis.
          * `watermarks` (`pulumi.Input[list]`)
            * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.
            * `value` (`pulumi.Input[float]`)

        The **axis_right** object supports the following:

          * `highWatermark` (`pulumi.Input[float]`) - A line to draw as a high watermark.
          * `highWatermarkLabel` (`pulumi.Input[str]`) - A label to attach to the high watermark line.
          * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.
          * `lowWatermark` (`pulumi.Input[float]`) - A line to draw as a low watermark.
          * `lowWatermarkLabel` (`pulumi.Input[str]`) - A label to attach to the low watermark line.
          * `maxValue` (`pulumi.Input[float]`) - The maximum value for the right axis.
          * `minValue` (`pulumi.Input[float]`) - The minimum value for the right axis.
          * `watermarks` (`pulumi.Input[list]`)
            * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.
            * `value` (`pulumi.Input[float]`)

        The **event_options** object supports the following:

          * `color` (`pulumi.Input[str]`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
          * `displayName` (`pulumi.Input[str]`) - Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
          * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.

        The **histogram_options** object supports the following:

          * `colorTheme` (`pulumi.Input[str]`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine, red, gold, greenyellow, chartreuse, jade

        The **legend_options_fields** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - True or False depending on if you want the property to be shown or hidden.
          * `property` (`pulumi.Input[str]`) - The name of the property to display. Note the special values of `plot_label` (corresponding with the API's `sf_metric`) which shows the label of the time series `publish()` and `metric` (corresponding with the API's `sf_originatingMetric`) that shows the name of the metric for the time series being displayed.

        The **viz_options** object supports the following:

          * `axis` (`pulumi.Input[str]`) - Y-axis associated with values for this plot. Must be either `right` or `left`.
          * `color` (`pulumi.Input[str]`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
          * `displayName` (`pulumi.Input[str]`) - Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
          * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.
          * `plot_type` (`pulumi.Input[str]`) - The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
          * `valuePrefix` (`pulumi.Input[str]`)
          * `valueSuffix` (`pulumi.Input[str]`)
          * `valueUnit` (`pulumi.Input[str]`) - A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
            * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, axes_include_zero=None, axes_precision=None, axis_left=None, axis_right=None, color_by=None, description=None, disable_sampling=None, end_time=None, event_options=None, histogram_options=None, legend_fields_to_hides=None, legend_options_fields=None, max_delay=None, minimum_resolution=None, name=None, on_chart_legend_dimension=None, plot_type=None, program_text=None, show_data_markers=None, show_event_lines=None, stacked=None, start_time=None, tags=None, time_range=None, timezone=None, unit_prefix=None, url=None, viz_options=None):
        """
        Get an existing TimeChart resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] axes_include_zero: Force the chart to display zero on the y-axes, even if none of the data is near zero.
        :param pulumi.Input[float] axes_precision: Specifies the digits SignalFx displays for values plotted on the chart. Defaults to `3`.
        :param pulumi.Input[dict] axis_left: Set of axis options.
        :param pulumi.Input[dict] axis_right: Set of axis options.
        :param pulumi.Input[str] color_by: Must be `"Dimension"` or `"Metric"`. `"Dimension"` by default.
        :param pulumi.Input[str] description: Description of the chart.
        :param pulumi.Input[bool] disable_sampling: If `false`, samples a subset of the output MTS, which improves UI performance. `false` by default
        :param pulumi.Input[float] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[list] event_options: Event customization options, associated with a publish statement. You will need to use this to change settings for any `events(…)` statements you use.
        :param pulumi.Input[list] histogram_options: Only used when `plot_type` is `"Histogram"`. Histogram specific options.
        :param pulumi.Input[list] legend_fields_to_hides: List of properties that should not be displayed in the chart legend (i.e. dimension names). All the properties are visible by default. Deprecated, please use `legend_options_fields`.
        :param pulumi.Input[list] legend_options_fields: List of property names and enabled flags that should be displayed in the data table for the chart, in the order provided. This option cannot be used with `legend_fields_to_hide`.
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
        :param pulumi.Input[list] tags: Tags associated with the chart
        :param pulumi.Input[float] time_range: How many seconds ago from which to display data. For example, the last hour would be `3600`, etc. Conflicts with `start_time` and `end_time`.
        :param pulumi.Input[str] timezone: Time zone that SignalFlow uses as the basis of calendar window transformation methods. For example, if you set "timezone": "Europe/Paris" and then use the transformation sum(cycle="week", cycle_start="Monday") in your chart's SignalFlow program, the calendar window starts on Monday, Paris time. See the [full list of timezones for more](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_supported_signalflow_time_zones). `"UTC"` by default.
        :param pulumi.Input[str] unit_prefix: Must be `"Metric"` or `"Binary`". `"Metric"` by default.
        :param pulumi.Input[str] url: URL of the chart
        :param pulumi.Input[list] viz_options: Plot-level customization options, associated with a publish statement.

        The **axis_left** object supports the following:

          * `highWatermark` (`pulumi.Input[float]`) - A line to draw as a high watermark.
          * `highWatermarkLabel` (`pulumi.Input[str]`) - A label to attach to the high watermark line.
          * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.
          * `lowWatermark` (`pulumi.Input[float]`) - A line to draw as a low watermark.
          * `lowWatermarkLabel` (`pulumi.Input[str]`) - A label to attach to the low watermark line.
          * `maxValue` (`pulumi.Input[float]`) - The maximum value for the right axis.
          * `minValue` (`pulumi.Input[float]`) - The minimum value for the right axis.
          * `watermarks` (`pulumi.Input[list]`)
            * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.
            * `value` (`pulumi.Input[float]`)

        The **axis_right** object supports the following:

          * `highWatermark` (`pulumi.Input[float]`) - A line to draw as a high watermark.
          * `highWatermarkLabel` (`pulumi.Input[str]`) - A label to attach to the high watermark line.
          * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.
          * `lowWatermark` (`pulumi.Input[float]`) - A line to draw as a low watermark.
          * `lowWatermarkLabel` (`pulumi.Input[str]`) - A label to attach to the low watermark line.
          * `maxValue` (`pulumi.Input[float]`) - The maximum value for the right axis.
          * `minValue` (`pulumi.Input[float]`) - The minimum value for the right axis.
          * `watermarks` (`pulumi.Input[list]`)
            * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.
            * `value` (`pulumi.Input[float]`)

        The **event_options** object supports the following:

          * `color` (`pulumi.Input[str]`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
          * `displayName` (`pulumi.Input[str]`) - Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
          * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.

        The **histogram_options** object supports the following:

          * `colorTheme` (`pulumi.Input[str]`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine, red, gold, greenyellow, chartreuse, jade

        The **legend_options_fields** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - True or False depending on if you want the property to be shown or hidden.
          * `property` (`pulumi.Input[str]`) - The name of the property to display. Note the special values of `plot_label` (corresponding with the API's `sf_metric`) which shows the label of the time series `publish()` and `metric` (corresponding with the API's `sf_originatingMetric`) that shows the name of the metric for the time series being displayed.

        The **viz_options** object supports the following:

          * `axis` (`pulumi.Input[str]`) - Y-axis associated with values for this plot. Must be either `right` or `left`.
          * `color` (`pulumi.Input[str]`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
          * `displayName` (`pulumi.Input[str]`) - Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
          * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the event query you want to customize.
          * `plot_type` (`pulumi.Input[str]`) - The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plot_type` by default.
          * `valuePrefix` (`pulumi.Input[str]`)
          * `valueSuffix` (`pulumi.Input[str]`)
          * `valueUnit` (`pulumi.Input[str]`) - A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
            * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

