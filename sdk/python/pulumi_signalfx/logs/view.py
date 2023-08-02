# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
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
        pulumi.set(__self__, "program_text", program_text)
        if columns is not None:
            pulumi.set(__self__, "columns", columns)
        if default_connection is not None:
            pulumi.set(__self__, "default_connection", default_connection)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sort_options is not None:
            pulumi.set(__self__, "sort_options", sort_options)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if time_range is not None:
            pulumi.set(__self__, "time_range", time_range)

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
        if columns is not None:
            pulumi.set(__self__, "columns", columns)
        if default_connection is not None:
            pulumi.set(__self__, "default_connection", default_connection)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if program_text is not None:
            pulumi.set(__self__, "program_text", program_text)
        if sort_options is not None:
            pulumi.set(__self__, "sort_options", sort_options)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if time_range is not None:
            pulumi.set(__self__, "time_range", time_range)
        if url is not None:
            pulumi.set(__self__, "url", url)

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


warnings.warn("""signalfx.logs/view.View has been deprecated in favor of signalfx.log/view.View""", DeprecationWarning)


class View(pulumi.CustomResource):
    warnings.warn("""signalfx.logs/view.View has been deprecated in favor of signalfx.log/view.View""", DeprecationWarning)

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

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        my_log_view = signalfx.log.View("myLogView",
            columns=[
                signalfx.log.ViewColumnArgs(
                    name="severity",
                ),
                signalfx.log.ViewColumnArgs(
                    name="time",
                ),
                signalfx.log.ViewColumnArgs(
                    name="amount.currency_code",
                ),
                signalfx.log.ViewColumnArgs(
                    name="amount.nanos",
                ),
                signalfx.log.ViewColumnArgs(
                    name="amount.units",
                ),
                signalfx.log.ViewColumnArgs(
                    name="message",
                ),
            ],
            description="Lorem ipsum dolor sit amet, laudem tibique iracundia at mea. Nam posse dolores ex, nec cu adhuc putent honestatis",
            program_text=\"\"\"logs(filter=field('message') == 'Transaction processed' and field('service.name') == 'paymentservice').publish()

        \"\"\",
            sort_options=[signalfx.log.ViewSortOptionArgs(
                descending=False,
                field="severity",
            )],
            time_range=900)
        ```

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

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        my_log_view = signalfx.log.View("myLogView",
            columns=[
                signalfx.log.ViewColumnArgs(
                    name="severity",
                ),
                signalfx.log.ViewColumnArgs(
                    name="time",
                ),
                signalfx.log.ViewColumnArgs(
                    name="amount.currency_code",
                ),
                signalfx.log.ViewColumnArgs(
                    name="amount.nanos",
                ),
                signalfx.log.ViewColumnArgs(
                    name="amount.units",
                ),
                signalfx.log.ViewColumnArgs(
                    name="message",
                ),
            ],
            description="Lorem ipsum dolor sit amet, laudem tibique iracundia at mea. Nam posse dolores ex, nec cu adhuc putent honestatis",
            program_text=\"\"\"logs(filter=field('message') == 'Transaction processed' and field('service.name') == 'paymentservice').publish()

        \"\"\",
            sort_options=[signalfx.log.ViewSortOptionArgs(
                descending=False,
                field="severity",
            )],
            time_range=900)
        ```

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
        pulumi.log.warn("""View is deprecated: signalfx.logs/view.View has been deprecated in favor of signalfx.log/view.View""")
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
        super(View, __self__).__init__(
            'signalfx:logs/view:View',
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

