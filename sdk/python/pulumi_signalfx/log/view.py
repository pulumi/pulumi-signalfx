# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ViewArgs', 'View']

@pulumi.input_type
class ViewArgs:
    def __init__(__self__, *,
                 program_text: pulumi.Input[str],
                 columns: Optional[pulumi.Input[Sequence[pulumi.Input['ViewColumnArgs']]]] = None,
                 default_connection: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sort_options: Optional[pulumi.Input[Sequence[pulumi.Input['ViewSortOptionArgs']]]] = None,
                 start_time: Optional[pulumi.Input[int]] = None,
                 time_range: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a View resource.
        :param pulumi.Input[str] program_text: Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
        :param pulumi.Input[Sequence[pulumi.Input['ViewColumnArgs']]] columns: The column headers to show on the log view.
        :param pulumi.Input[str] default_connection: The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
        :param pulumi.Input[str] description: Description of the log view.
        :param pulumi.Input[int] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[str] name: Name of the log view.
        :param pulumi.Input[Sequence[pulumi.Input['ViewSortOptionArgs']]] sort_options: The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
        :param pulumi.Input[int] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[int] time_range: From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        ViewArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            program_text=program_text,
            columns=columns,
            default_connection=default_connection,
            description=description,
            end_time=end_time,
            name=name,
            sort_options=sort_options,
            start_time=start_time,
            time_range=time_range,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             program_text: Optional[pulumi.Input[str]] = None,
             columns: Optional[pulumi.Input[Sequence[pulumi.Input['ViewColumnArgs']]]] = None,
             default_connection: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             end_time: Optional[pulumi.Input[int]] = None,
             name: Optional[pulumi.Input[str]] = None,
             sort_options: Optional[pulumi.Input[Sequence[pulumi.Input['ViewSortOptionArgs']]]] = None,
             start_time: Optional[pulumi.Input[int]] = None,
             time_range: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if program_text is None and 'programText' in kwargs:
            program_text = kwargs['programText']
        if program_text is None:
            raise TypeError("Missing 'program_text' argument")
        if default_connection is None and 'defaultConnection' in kwargs:
            default_connection = kwargs['defaultConnection']
        if end_time is None and 'endTime' in kwargs:
            end_time = kwargs['endTime']
        if sort_options is None and 'sortOptions' in kwargs:
            sort_options = kwargs['sortOptions']
        if start_time is None and 'startTime' in kwargs:
            start_time = kwargs['startTime']
        if time_range is None and 'timeRange' in kwargs:
            time_range = kwargs['timeRange']

        _setter("program_text", program_text)
        if columns is not None:
            _setter("columns", columns)
        if default_connection is not None:
            _setter("default_connection", default_connection)
        if description is not None:
            _setter("description", description)
        if end_time is not None:
            _setter("end_time", end_time)
        if name is not None:
            _setter("name", name)
        if sort_options is not None:
            _setter("sort_options", sort_options)
        if start_time is not None:
            _setter("start_time", start_time)
        if time_range is not None:
            _setter("time_range", time_range)

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> pulumi.Input[str]:
        """
        Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
        """
        return pulumi.get(self, "program_text")

    @program_text.setter
    def program_text(self, value: pulumi.Input[str]):
        pulumi.set(self, "program_text", value)

    @property
    @pulumi.getter
    def columns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ViewColumnArgs']]]]:
        """
        The column headers to show on the log view.
        """
        return pulumi.get(self, "columns")

    @columns.setter
    def columns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ViewColumnArgs']]]]):
        pulumi.set(self, "columns", value)

    @property
    @pulumi.getter(name="defaultConnection")
    def default_connection(self) -> Optional[pulumi.Input[str]]:
        """
        The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
        """
        return pulumi.get(self, "default_connection")

    @default_connection.setter
    def default_connection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_connection", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the log view.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the log view.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sortOptions")
    def sort_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ViewSortOptionArgs']]]]:
        """
        The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
        """
        return pulumi.get(self, "sort_options")

    @sort_options.setter
    def sort_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ViewSortOptionArgs']]]]):
        pulumi.set(self, "sort_options", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> Optional[pulumi.Input[int]]:
        """
        From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        return pulumi.get(self, "time_range")

    @time_range.setter
    def time_range(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_range", value)


@pulumi.input_type
class _ViewState:
    def __init__(__self__, *,
                 columns: Optional[pulumi.Input[Sequence[pulumi.Input['ViewColumnArgs']]]] = None,
                 default_connection: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 program_text: Optional[pulumi.Input[str]] = None,
                 sort_options: Optional[pulumi.Input[Sequence[pulumi.Input['ViewSortOptionArgs']]]] = None,
                 start_time: Optional[pulumi.Input[int]] = None,
                 time_range: Optional[pulumi.Input[int]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering View resources.
        :param pulumi.Input[Sequence[pulumi.Input['ViewColumnArgs']]] columns: The column headers to show on the log view.
        :param pulumi.Input[str] default_connection: The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
        :param pulumi.Input[str] description: Description of the log view.
        :param pulumi.Input[int] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[str] name: Name of the log view.
        :param pulumi.Input[str] program_text: Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
        :param pulumi.Input[Sequence[pulumi.Input['ViewSortOptionArgs']]] sort_options: The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
        :param pulumi.Input[int] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[int] time_range: From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        :param pulumi.Input[str] url: The URL of the log view.
        """
        _ViewState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            columns=columns,
            default_connection=default_connection,
            description=description,
            end_time=end_time,
            name=name,
            program_text=program_text,
            sort_options=sort_options,
            start_time=start_time,
            time_range=time_range,
            url=url,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             columns: Optional[pulumi.Input[Sequence[pulumi.Input['ViewColumnArgs']]]] = None,
             default_connection: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             end_time: Optional[pulumi.Input[int]] = None,
             name: Optional[pulumi.Input[str]] = None,
             program_text: Optional[pulumi.Input[str]] = None,
             sort_options: Optional[pulumi.Input[Sequence[pulumi.Input['ViewSortOptionArgs']]]] = None,
             start_time: Optional[pulumi.Input[int]] = None,
             time_range: Optional[pulumi.Input[int]] = None,
             url: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if default_connection is None and 'defaultConnection' in kwargs:
            default_connection = kwargs['defaultConnection']
        if end_time is None and 'endTime' in kwargs:
            end_time = kwargs['endTime']
        if program_text is None and 'programText' in kwargs:
            program_text = kwargs['programText']
        if sort_options is None and 'sortOptions' in kwargs:
            sort_options = kwargs['sortOptions']
        if start_time is None and 'startTime' in kwargs:
            start_time = kwargs['startTime']
        if time_range is None and 'timeRange' in kwargs:
            time_range = kwargs['timeRange']

        if columns is not None:
            _setter("columns", columns)
        if default_connection is not None:
            _setter("default_connection", default_connection)
        if description is not None:
            _setter("description", description)
        if end_time is not None:
            _setter("end_time", end_time)
        if name is not None:
            _setter("name", name)
        if program_text is not None:
            _setter("program_text", program_text)
        if sort_options is not None:
            _setter("sort_options", sort_options)
        if start_time is not None:
            _setter("start_time", start_time)
        if time_range is not None:
            _setter("time_range", time_range)
        if url is not None:
            _setter("url", url)

    @property
    @pulumi.getter
    def columns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ViewColumnArgs']]]]:
        """
        The column headers to show on the log view.
        """
        return pulumi.get(self, "columns")

    @columns.setter
    def columns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ViewColumnArgs']]]]):
        pulumi.set(self, "columns", value)

    @property
    @pulumi.getter(name="defaultConnection")
    def default_connection(self) -> Optional[pulumi.Input[str]]:
        """
        The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
        """
        return pulumi.get(self, "default_connection")

    @default_connection.setter
    def default_connection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_connection", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the log view.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the log view.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> Optional[pulumi.Input[str]]:
        """
        Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
        """
        return pulumi.get(self, "program_text")

    @program_text.setter
    def program_text(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "program_text", value)

    @property
    @pulumi.getter(name="sortOptions")
    def sort_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ViewSortOptionArgs']]]]:
        """
        The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
        """
        return pulumi.get(self, "sort_options")

    @sort_options.setter
    def sort_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ViewSortOptionArgs']]]]):
        pulumi.set(self, "sort_options", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> Optional[pulumi.Input[int]]:
        """
        From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        return pulumi.get(self, "time_range")

    @time_range.setter
    def time_range(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_range", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the log view.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class View(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewColumnArgs']]]]] = None,
                 default_connection: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 program_text: Optional[pulumi.Input[str]] = None,
                 sort_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewSortOptionArgs']]]]] = None,
                 start_time: Optional[pulumi.Input[int]] = None,
                 time_range: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        You can add logs data to your Observability Cloud dashboards without turning your logs into metrics first. A log view displays log lines in a table form in a dashboard and shows you in detail what is happening and why.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewColumnArgs']]]] columns: The column headers to show on the log view.
        :param pulumi.Input[str] default_connection: The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
        :param pulumi.Input[str] description: Description of the log view.
        :param pulumi.Input[int] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[str] name: Name of the log view.
        :param pulumi.Input[str] program_text: Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewSortOptionArgs']]]] sort_options: The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
        :param pulumi.Input[int] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[int] time_range: From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ViewArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        You can add logs data to your Observability Cloud dashboards without turning your logs into metrics first. A log view displays log lines in a table form in a dashboard and shows you in detail what is happening and why.

        :param str resource_name: The name of the resource.
        :param ViewArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ViewArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ViewArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewColumnArgs']]]]] = None,
                 default_connection: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 program_text: Optional[pulumi.Input[str]] = None,
                 sort_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewSortOptionArgs']]]]] = None,
                 start_time: Optional[pulumi.Input[int]] = None,
                 time_range: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ViewArgs.__new__(ViewArgs)

            __props__.__dict__["columns"] = columns
            __props__.__dict__["default_connection"] = default_connection
            __props__.__dict__["description"] = description
            __props__.__dict__["end_time"] = end_time
            __props__.__dict__["name"] = name
            if program_text is None and not opts.urn:
                raise TypeError("Missing required property 'program_text'")
            __props__.__dict__["program_text"] = program_text
            __props__.__dict__["sort_options"] = sort_options
            __props__.__dict__["start_time"] = start_time
            __props__.__dict__["time_range"] = time_range
            __props__.__dict__["url"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="signalfx:logs/view:View")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(View, __self__).__init__(
            'signalfx:log/view:View',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            columns: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewColumnArgs']]]]] = None,
            default_connection: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            end_time: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            program_text: Optional[pulumi.Input[str]] = None,
            sort_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewSortOptionArgs']]]]] = None,
            start_time: Optional[pulumi.Input[int]] = None,
            time_range: Optional[pulumi.Input[int]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'View':
        """
        Get an existing View resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewColumnArgs']]]] columns: The column headers to show on the log view.
        :param pulumi.Input[str] default_connection: The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
        :param pulumi.Input[str] description: Description of the log view.
        :param pulumi.Input[int] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[str] name: Name of the log view.
        :param pulumi.Input[str] program_text: Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ViewSortOptionArgs']]]] sort_options: The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
        :param pulumi.Input[int] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[int] time_range: From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        :param pulumi.Input[str] url: The URL of the log view.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ViewState.__new__(_ViewState)

        __props__.__dict__["columns"] = columns
        __props__.__dict__["default_connection"] = default_connection
        __props__.__dict__["description"] = description
        __props__.__dict__["end_time"] = end_time
        __props__.__dict__["name"] = name
        __props__.__dict__["program_text"] = program_text
        __props__.__dict__["sort_options"] = sort_options
        __props__.__dict__["start_time"] = start_time
        __props__.__dict__["time_range"] = time_range
        __props__.__dict__["url"] = url
        return View(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def columns(self) -> pulumi.Output[Optional[Sequence['outputs.ViewColumn']]]:
        """
        The column headers to show on the log view.
        """
        return pulumi.get(self, "columns")

    @property
    @pulumi.getter(name="defaultConnection")
    def default_connection(self) -> pulumi.Output[Optional[str]]:
        """
        The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
        """
        return pulumi.get(self, "default_connection")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the log view.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the log view.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> pulumi.Output[str]:
        """
        Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
        """
        return pulumi.get(self, "program_text")

    @property
    @pulumi.getter(name="sortOptions")
    def sort_options(self) -> pulumi.Output[Optional[Sequence['outputs.ViewSortOption']]]:
        """
        The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
        """
        return pulumi.get(self, "sort_options")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> pulumi.Output[Optional[int]]:
        """
        From when to display data. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`). Conflicts with `start_time` and `end_time`.
        """
        return pulumi.get(self, "time_range")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the log view.
        """
        return pulumi.get(self, "url")

