# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class SingleValueChart(pulumi.CustomResource):
    color_by: pulumi.Output[str]
    """
    Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
    """
    color_scales: pulumi.Output[list]
    """
    Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
    
      * `color` (`str`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
      * `gt` (`float`) - Indicates the lower threshold non-inclusive value for this range.
      * `gte` (`float`) - Indicates the lower threshold inclusive value for this range.
      * `lt` (`float`) - Indicates the upper threshold non-inculsive value for this range.
      * `lte` (`float`) - Indicates the upper threshold inclusive value for this range.
    """
    description: pulumi.Output[str]
    """
    Description of the chart.
    """
    is_timestamp_hidden: pulumi.Output[bool]
    """
    Whether to hide the timestamp in the chart. `false` by default.
    """
    max_delay: pulumi.Output[float]
    """
    How long (in seconds) to wait for late datapoints
    """
    max_precision: pulumi.Output[float]
    """
    The maximum precision to for value displayed.
    """
    name: pulumi.Output[str]
    """
    Name of the chart.
    """
    program_text: pulumi.Output[str]
    """
    Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
    """
    refresh_interval: pulumi.Output[float]
    """
    How often (in seconds) to refresh the value.
    """
    secondary_visualization: pulumi.Output[str]
    """
    The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
    """
    show_spark_line: pulumi.Output[bool]
    """
    Whether to show a trend line below the current value. `false` by default.
    """
    unit_prefix: pulumi.Output[str]
    """
    Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
    """
    url: pulumi.Output[str]
    viz_options: pulumi.Output[list]
    """
    Plot-level customization options, associated with a publish statement.
    
      * `color` (`str`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
      * `displayName` (`str`) - Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
      * `label` (`str`) - Label used in the publish statement that displays the plot (metric time series data) you want to customize.
      * `valuePrefix` (`str`)
      * `valueSuffix` (`str`)
      * `valueUnit` (`str`) - A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
        * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
    """
    def __init__(__self__, resource_name, opts=None, color_by=None, color_scales=None, description=None, is_timestamp_hidden=None, max_delay=None, max_precision=None, name=None, program_text=None, refresh_interval=None, secondary_visualization=None, show_spark_line=None, unit_prefix=None, viz_options=None, __props__=None, __name__=None, __opts__=None):
        """
        This chart type displays a single number in a large font, representing the current value of a single metric on a plot line.
        
        If the time period is in the past, the number represents the value of the metric near the end of the time period.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color_by: Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
        :param pulumi.Input[list] color_scales: Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        :param pulumi.Input[str] description: Description of the chart.
        :param pulumi.Input[bool] is_timestamp_hidden: Whether to hide the timestamp in the chart. `false` by default.
        :param pulumi.Input[float] max_delay: How long (in seconds) to wait for late datapoints
        :param pulumi.Input[float] max_precision: The maximum precision to for value displayed.
        :param pulumi.Input[str] name: Name of the chart.
        :param pulumi.Input[str] program_text: Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
        :param pulumi.Input[float] refresh_interval: How often (in seconds) to refresh the value.
        :param pulumi.Input[str] secondary_visualization: The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
        :param pulumi.Input[bool] show_spark_line: Whether to show a trend line below the current value. `false` by default.
        :param pulumi.Input[str] unit_prefix: Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
        :param pulumi.Input[list] viz_options: Plot-level customization options, associated with a publish statement.
        
        The **color_scales** object supports the following:
        
          * `color` (`pulumi.Input[str]`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
          * `gt` (`pulumi.Input[float]`) - Indicates the lower threshold non-inclusive value for this range.
          * `gte` (`pulumi.Input[float]`) - Indicates the lower threshold inclusive value for this range.
          * `lt` (`pulumi.Input[float]`) - Indicates the upper threshold non-inculsive value for this range.
          * `lte` (`pulumi.Input[float]`) - Indicates the upper threshold inclusive value for this range.
        
        The **viz_options** object supports the following:
        
          * `color` (`pulumi.Input[str]`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
          * `displayName` (`pulumi.Input[str]`) - Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
          * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the plot (metric time series data) you want to customize.
          * `valuePrefix` (`pulumi.Input[str]`)
          * `valueSuffix` (`pulumi.Input[str]`)
          * `valueUnit` (`pulumi.Input[str]`) - A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
            * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/single_value_chart.html.markdown.
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

            __props__['color_by'] = color_by
            __props__['color_scales'] = color_scales
            __props__['description'] = description
            __props__['is_timestamp_hidden'] = is_timestamp_hidden
            __props__['max_delay'] = max_delay
            __props__['max_precision'] = max_precision
            __props__['name'] = name
            if program_text is None:
                raise TypeError("Missing required property 'program_text'")
            __props__['program_text'] = program_text
            __props__['refresh_interval'] = refresh_interval
            __props__['secondary_visualization'] = secondary_visualization
            __props__['show_spark_line'] = show_spark_line
            __props__['unit_prefix'] = unit_prefix
            __props__['viz_options'] = viz_options
            __props__['url'] = None
        super(SingleValueChart, __self__).__init__(
            'signalfx:index/singleValueChart:SingleValueChart',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, color_by=None, color_scales=None, description=None, is_timestamp_hidden=None, max_delay=None, max_precision=None, name=None, program_text=None, refresh_interval=None, secondary_visualization=None, show_spark_line=None, unit_prefix=None, url=None, viz_options=None):
        """
        Get an existing SingleValueChart resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color_by: Must be `"Dimension"`, `"Scale"` or `"Metric"`. `"Dimension"` by default.
        :param pulumi.Input[list] color_scales: Single color range including both the color to display for that range and the borders of the range. Example: `[{ gt = 60, color = "blue" }, { lte = 60, color = "yellow" }]`. Look at this [link](https://docs.signalfx.com/en/latest/charts/chart-options-tab.html).
        :param pulumi.Input[str] description: Description of the chart.
        :param pulumi.Input[bool] is_timestamp_hidden: Whether to hide the timestamp in the chart. `false` by default.
        :param pulumi.Input[float] max_delay: How long (in seconds) to wait for late datapoints
        :param pulumi.Input[float] max_precision: The maximum precision to for value displayed.
        :param pulumi.Input[str] name: Name of the chart.
        :param pulumi.Input[str] program_text: Signalflow program text for the chart. More info at <https://developers.signalfx.com/docs/signalflow-overview>.
        :param pulumi.Input[float] refresh_interval: How often (in seconds) to refresh the value.
        :param pulumi.Input[str] secondary_visualization: The type of secondary visualization. Can be `None`, `Radial`, `Linear`, or `Sparkline`. If unset, the SignalFx default is used (`None`).
        :param pulumi.Input[bool] show_spark_line: Whether to show a trend line below the current value. `false` by default.
        :param pulumi.Input[str] unit_prefix: Must be `"Metric"` or `"Binary"`. `"Metric"` by default.
        :param pulumi.Input[list] viz_options: Plot-level customization options, associated with a publish statement.
        
        The **color_scales** object supports the following:
        
          * `color` (`pulumi.Input[str]`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
          * `gt` (`pulumi.Input[float]`) - Indicates the lower threshold non-inclusive value for this range.
          * `gte` (`pulumi.Input[float]`) - Indicates the lower threshold inclusive value for this range.
          * `lt` (`pulumi.Input[float]`) - Indicates the upper threshold non-inculsive value for this range.
          * `lte` (`pulumi.Input[float]`) - Indicates the upper threshold inclusive value for this range.
        
        The **viz_options** object supports the following:
        
          * `color` (`pulumi.Input[str]`) - Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
          * `displayName` (`pulumi.Input[str]`) - Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
          * `label` (`pulumi.Input[str]`) - Label used in the publish statement that displays the plot (metric time series data) you want to customize.
          * `valuePrefix` (`pulumi.Input[str]`)
          * `valueSuffix` (`pulumi.Input[str]`)
          * `valueUnit` (`pulumi.Input[str]`) - A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
            * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/single_value_chart.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["color_by"] = color_by
        __props__["color_scales"] = color_scales
        __props__["description"] = description
        __props__["is_timestamp_hidden"] = is_timestamp_hidden
        __props__["max_delay"] = max_delay
        __props__["max_precision"] = max_precision
        __props__["name"] = name
        __props__["program_text"] = program_text
        __props__["refresh_interval"] = refresh_interval
        __props__["secondary_visualization"] = secondary_visualization
        __props__["show_spark_line"] = show_spark_line
        __props__["unit_prefix"] = unit_prefix
        __props__["url"] = url
        __props__["viz_options"] = viz_options
        return SingleValueChart(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

