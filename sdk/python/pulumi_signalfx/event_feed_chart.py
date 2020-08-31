# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['EventFeedChart']


class EventFeedChart(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 program_text: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[float]] = None,
                 time_range: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Displays a listing of events as a widget in a dashboard.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the text note.
        :param pulumi.Input[float] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[str] name: Name of the text note.
        :param pulumi.Input[str] program_text: Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        :param pulumi.Input[float] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[float] time_range: From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
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

            __props__['description'] = description
            __props__['end_time'] = end_time
            __props__['name'] = name
            if program_text is None:
                raise TypeError("Missing required property 'program_text'")
            __props__['program_text'] = program_text
            __props__['start_time'] = start_time
            __props__['time_range'] = time_range
            __props__['url'] = None
        super(EventFeedChart, __self__).__init__(
            'signalfx:index/eventFeedChart:EventFeedChart',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            end_time: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            program_text: Optional[pulumi.Input[str]] = None,
            start_time: Optional[pulumi.Input[float]] = None,
            time_range: Optional[pulumi.Input[float]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'EventFeedChart':
        """
        Get an existing EventFeedChart resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the text note.
        :param pulumi.Input[float] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[str] name: Name of the text note.
        :param pulumi.Input[str] program_text: Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        :param pulumi.Input[float] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[float] time_range: From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        :param pulumi.Input[str] url: URL of the chart
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["end_time"] = end_time
        __props__["name"] = name
        __props__["program_text"] = program_text
        __props__["start_time"] = start_time
        __props__["time_range"] = time_range
        __props__["url"] = url
        return EventFeedChart(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the text note.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[float]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the text note.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> pulumi.Output[str]:
        """
        Signalflow program text for the chart. More info[in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        """
        return pulumi.get(self, "program_text")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[float]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> pulumi.Output[Optional[float]]:
        """
        From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        return pulumi.get(self, "time_range")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL of the chart
        """
        return pulumi.get(self, "url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

