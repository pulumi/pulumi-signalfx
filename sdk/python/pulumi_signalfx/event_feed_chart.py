# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['EventFeedChartArgs', 'EventFeedChart']

@pulumi.input_type
class EventFeedChartArgs:
    def __init__(__self__, *,
                 program_text: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 end_time: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 start_time: Optional[pulumi.Input[builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 time_range: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a EventFeedChart resource.
        :param pulumi.Input[builtins.str] program_text: Signalflow program text for the chart. More info[in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
        :param pulumi.Input[builtins.str] description: Description of the text note.
        :param pulumi.Input[builtins.int] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[builtins.str] name: Name of the text note.
        :param pulumi.Input[builtins.int] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: Tags associated with the resource
        :param pulumi.Input[builtins.int] time_range: From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        pulumi.set(__self__, "program_text", program_text)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if time_range is not None:
            pulumi.set(__self__, "time_range", time_range)

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> pulumi.Input[builtins.str]:
        """
        Signalflow program text for the chart. More info[in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
        """
        return pulumi.get(self, "program_text")

    @program_text.setter
    def program_text(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "program_text", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the text note.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the text note.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Tags associated with the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        return pulumi.get(self, "time_range")

    @time_range.setter
    def time_range(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "time_range", value)


@pulumi.input_type
class _EventFeedChartState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 end_time: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 program_text: Optional[pulumi.Input[builtins.str]] = None,
                 start_time: Optional[pulumi.Input[builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 time_range: Optional[pulumi.Input[builtins.int]] = None,
                 url: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering EventFeedChart resources.
        :param pulumi.Input[builtins.str] description: Description of the text note.
        :param pulumi.Input[builtins.int] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[builtins.str] name: Name of the text note.
        :param pulumi.Input[builtins.str] program_text: Signalflow program text for the chart. More info[in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
        :param pulumi.Input[builtins.int] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: Tags associated with the resource
        :param pulumi.Input[builtins.int] time_range: From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        :param pulumi.Input[builtins.str] url: The URL of the chart.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if program_text is not None:
            pulumi.set(__self__, "program_text", program_text)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if time_range is not None:
            pulumi.set(__self__, "time_range", time_range)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the text note.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the text note.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Signalflow program text for the chart. More info[in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
        """
        return pulumi.get(self, "program_text")

    @program_text.setter
    def program_text(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "program_text", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Tags associated with the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        return pulumi.get(self, "time_range")

    @time_range.setter
    def time_range(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "time_range", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The URL of the chart.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "url", value)


class EventFeedChart(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 end_time: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 program_text: Optional[pulumi.Input[builtins.str]] = None,
                 start_time: Optional[pulumi.Input[builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 time_range: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        Displays a listing of events as a widget in a dashboard.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Description of the text note.
        :param pulumi.Input[builtins.int] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[builtins.str] name: Name of the text note.
        :param pulumi.Input[builtins.str] program_text: Signalflow program text for the chart. More info[in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
        :param pulumi.Input[builtins.int] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: Tags associated with the resource
        :param pulumi.Input[builtins.int] time_range: From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventFeedChartArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Displays a listing of events as a widget in a dashboard.

        :param str resource_name: The name of the resource.
        :param EventFeedChartArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventFeedChartArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 end_time: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 program_text: Optional[pulumi.Input[builtins.str]] = None,
                 start_time: Optional[pulumi.Input[builtins.int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 time_range: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventFeedChartArgs.__new__(EventFeedChartArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["end_time"] = end_time
            __props__.__dict__["name"] = name
            if program_text is None and not opts.urn:
                raise TypeError("Missing required property 'program_text'")
            __props__.__dict__["program_text"] = program_text
            __props__.__dict__["start_time"] = start_time
            __props__.__dict__["tags"] = tags
            __props__.__dict__["time_range"] = time_range
            __props__.__dict__["url"] = None
        super(EventFeedChart, __self__).__init__(
            'signalfx:index/eventFeedChart:EventFeedChart',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            end_time: Optional[pulumi.Input[builtins.int]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            program_text: Optional[pulumi.Input[builtins.str]] = None,
            start_time: Optional[pulumi.Input[builtins.int]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            time_range: Optional[pulumi.Input[builtins.int]] = None,
            url: Optional[pulumi.Input[builtins.str]] = None) -> 'EventFeedChart':
        """
        Get an existing EventFeedChart resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Description of the text note.
        :param pulumi.Input[builtins.int] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[builtins.str] name: Name of the text note.
        :param pulumi.Input[builtins.str] program_text: Signalflow program text for the chart. More info[in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
        :param pulumi.Input[builtins.int] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: Tags associated with the resource
        :param pulumi.Input[builtins.int] time_range: From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        :param pulumi.Input[builtins.str] url: The URL of the chart.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EventFeedChartState.__new__(_EventFeedChartState)

        __props__.__dict__["description"] = description
        __props__.__dict__["end_time"] = end_time
        __props__.__dict__["name"] = name
        __props__.__dict__["program_text"] = program_text
        __props__.__dict__["start_time"] = start_time
        __props__.__dict__["tags"] = tags
        __props__.__dict__["time_range"] = time_range
        __props__.__dict__["url"] = url
        return EventFeedChart(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description of the text note.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the text note.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> pulumi.Output[builtins.str]:
        """
        Signalflow program text for the chart. More info[in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
        """
        return pulumi.get(self, "program_text")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        Tags associated with the resource
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        From when to display data. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        return pulumi.get(self, "time_range")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[builtins.str]:
        """
        The URL of the chart.
        """
        return pulumi.get(self, "url")

