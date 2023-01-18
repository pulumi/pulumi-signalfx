# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TableChartArgs', 'TableChart']

@pulumi.input_type
class TableChartArgs:
    def __init__(__self__, *,
                 program_text: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 disable_sampling: Optional[pulumi.Input[bool]] = None,
                 group_bies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hide_timestamp: Optional[pulumi.Input[bool]] = None,
                 max_delay: Optional[pulumi.Input[int]] = None,
                 minimum_resolution: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 refresh_interval: Optional[pulumi.Input[int]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 unit_prefix: Optional[pulumi.Input[str]] = None,
                 viz_options: Optional[pulumi.Input[Sequence[pulumi.Input['TableChartVizOptionArgs']]]] = None):
        """
        The set of arguments for constructing a TableChart resource.
        :param pulumi.Input[str] program_text: The SignalFlow for your Data Table Chart
        :param pulumi.Input[str] description: Description of the table chart.
        :param pulumi.Input[bool] disable_sampling: (false by default) If false, samples a subset of the output MTS, which improves UI performance
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_bies: Dimension to group by
        :param pulumi.Input[bool] hide_timestamp: (false by default) Whether to show the timestamp in the chart
        :param pulumi.Input[int] max_delay: How long (in seconds) to wait for late datapoints
        :param pulumi.Input[int] minimum_resolution: The minimum resolution (in seconds) to use for computing the underlying program
        :param pulumi.Input[str] name: Name of the table chart.
        :param pulumi.Input[int] refresh_interval: How often (in seconds) to refresh the values of the Table
        :param pulumi.Input[str] timezone: The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        :param pulumi.Input[str] unit_prefix: (Metric by default) Must be "Metric" or "Binary"
        :param pulumi.Input[Sequence[pulumi.Input['TableChartVizOptionArgs']]] viz_options: Plot-level customization options, associated with a publish statement
        """
        pulumi.set(__self__, "program_text", program_text)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable_sampling is not None:
            pulumi.set(__self__, "disable_sampling", disable_sampling)
        if group_bies is not None:
            pulumi.set(__self__, "group_bies", group_bies)
        if hide_timestamp is not None:
            pulumi.set(__self__, "hide_timestamp", hide_timestamp)
        if max_delay is not None:
            pulumi.set(__self__, "max_delay", max_delay)
        if minimum_resolution is not None:
            pulumi.set(__self__, "minimum_resolution", minimum_resolution)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if refresh_interval is not None:
            pulumi.set(__self__, "refresh_interval", refresh_interval)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)
        if unit_prefix is not None:
            pulumi.set(__self__, "unit_prefix", unit_prefix)
        if viz_options is not None:
            pulumi.set(__self__, "viz_options", viz_options)

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> pulumi.Input[str]:
        """
        The SignalFlow for your Data Table Chart
        """
        return pulumi.get(self, "program_text")

    @program_text.setter
    def program_text(self, value: pulumi.Input[str]):
        pulumi.set(self, "program_text", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the table chart.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disableSampling")
    def disable_sampling(self) -> Optional[pulumi.Input[bool]]:
        """
        (false by default) If false, samples a subset of the output MTS, which improves UI performance
        """
        return pulumi.get(self, "disable_sampling")

    @disable_sampling.setter
    def disable_sampling(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_sampling", value)

    @property
    @pulumi.getter(name="groupBies")
    def group_bies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Dimension to group by
        """
        return pulumi.get(self, "group_bies")

    @group_bies.setter
    def group_bies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "group_bies", value)

    @property
    @pulumi.getter(name="hideTimestamp")
    def hide_timestamp(self) -> Optional[pulumi.Input[bool]]:
        """
        (false by default) Whether to show the timestamp in the chart
        """
        return pulumi.get(self, "hide_timestamp")

    @hide_timestamp.setter
    def hide_timestamp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hide_timestamp", value)

    @property
    @pulumi.getter(name="maxDelay")
    def max_delay(self) -> Optional[pulumi.Input[int]]:
        """
        How long (in seconds) to wait for late datapoints
        """
        return pulumi.get(self, "max_delay")

    @max_delay.setter
    def max_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_delay", value)

    @property
    @pulumi.getter(name="minimumResolution")
    def minimum_resolution(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum resolution (in seconds) to use for computing the underlying program
        """
        return pulumi.get(self, "minimum_resolution")

    @minimum_resolution.setter
    def minimum_resolution(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minimum_resolution", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the table chart.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="refreshInterval")
    def refresh_interval(self) -> Optional[pulumi.Input[int]]:
        """
        How often (in seconds) to refresh the values of the Table
        """
        return pulumi.get(self, "refresh_interval")

    @refresh_interval.setter
    def refresh_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "refresh_interval", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        """
        The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        """
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)

    @property
    @pulumi.getter(name="unitPrefix")
    def unit_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        (Metric by default) Must be "Metric" or "Binary"
        """
        return pulumi.get(self, "unit_prefix")

    @unit_prefix.setter
    def unit_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unit_prefix", value)

    @property
    @pulumi.getter(name="vizOptions")
    def viz_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TableChartVizOptionArgs']]]]:
        """
        Plot-level customization options, associated with a publish statement
        """
        return pulumi.get(self, "viz_options")

    @viz_options.setter
    def viz_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TableChartVizOptionArgs']]]]):
        pulumi.set(self, "viz_options", value)


@pulumi.input_type
class _TableChartState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_sampling: Optional[pulumi.Input[bool]] = None,
                 group_bies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hide_timestamp: Optional[pulumi.Input[bool]] = None,
                 max_delay: Optional[pulumi.Input[int]] = None,
                 minimum_resolution: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 program_text: Optional[pulumi.Input[str]] = None,
                 refresh_interval: Optional[pulumi.Input[int]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 unit_prefix: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 viz_options: Optional[pulumi.Input[Sequence[pulumi.Input['TableChartVizOptionArgs']]]] = None):
        """
        Input properties used for looking up and filtering TableChart resources.
        :param pulumi.Input[str] description: Description of the table chart.
        :param pulumi.Input[bool] disable_sampling: (false by default) If false, samples a subset of the output MTS, which improves UI performance
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_bies: Dimension to group by
        :param pulumi.Input[bool] hide_timestamp: (false by default) Whether to show the timestamp in the chart
        :param pulumi.Input[int] max_delay: How long (in seconds) to wait for late datapoints
        :param pulumi.Input[int] minimum_resolution: The minimum resolution (in seconds) to use for computing the underlying program
        :param pulumi.Input[str] name: Name of the table chart.
        :param pulumi.Input[str] program_text: The SignalFlow for your Data Table Chart
        :param pulumi.Input[int] refresh_interval: How often (in seconds) to refresh the values of the Table
        :param pulumi.Input[str] timezone: The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        :param pulumi.Input[str] unit_prefix: (Metric by default) Must be "Metric" or "Binary"
        :param pulumi.Input[str] url: The URL of the chart.
        :param pulumi.Input[Sequence[pulumi.Input['TableChartVizOptionArgs']]] viz_options: Plot-level customization options, associated with a publish statement
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable_sampling is not None:
            pulumi.set(__self__, "disable_sampling", disable_sampling)
        if group_bies is not None:
            pulumi.set(__self__, "group_bies", group_bies)
        if hide_timestamp is not None:
            pulumi.set(__self__, "hide_timestamp", hide_timestamp)
        if max_delay is not None:
            pulumi.set(__self__, "max_delay", max_delay)
        if minimum_resolution is not None:
            pulumi.set(__self__, "minimum_resolution", minimum_resolution)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if program_text is not None:
            pulumi.set(__self__, "program_text", program_text)
        if refresh_interval is not None:
            pulumi.set(__self__, "refresh_interval", refresh_interval)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)
        if unit_prefix is not None:
            pulumi.set(__self__, "unit_prefix", unit_prefix)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if viz_options is not None:
            pulumi.set(__self__, "viz_options", viz_options)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the table chart.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disableSampling")
    def disable_sampling(self) -> Optional[pulumi.Input[bool]]:
        """
        (false by default) If false, samples a subset of the output MTS, which improves UI performance
        """
        return pulumi.get(self, "disable_sampling")

    @disable_sampling.setter
    def disable_sampling(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_sampling", value)

    @property
    @pulumi.getter(name="groupBies")
    def group_bies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Dimension to group by
        """
        return pulumi.get(self, "group_bies")

    @group_bies.setter
    def group_bies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "group_bies", value)

    @property
    @pulumi.getter(name="hideTimestamp")
    def hide_timestamp(self) -> Optional[pulumi.Input[bool]]:
        """
        (false by default) Whether to show the timestamp in the chart
        """
        return pulumi.get(self, "hide_timestamp")

    @hide_timestamp.setter
    def hide_timestamp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hide_timestamp", value)

    @property
    @pulumi.getter(name="maxDelay")
    def max_delay(self) -> Optional[pulumi.Input[int]]:
        """
        How long (in seconds) to wait for late datapoints
        """
        return pulumi.get(self, "max_delay")

    @max_delay.setter
    def max_delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_delay", value)

    @property
    @pulumi.getter(name="minimumResolution")
    def minimum_resolution(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum resolution (in seconds) to use for computing the underlying program
        """
        return pulumi.get(self, "minimum_resolution")

    @minimum_resolution.setter
    def minimum_resolution(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "minimum_resolution", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the table chart.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> Optional[pulumi.Input[str]]:
        """
        The SignalFlow for your Data Table Chart
        """
        return pulumi.get(self, "program_text")

    @program_text.setter
    def program_text(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "program_text", value)

    @property
    @pulumi.getter(name="refreshInterval")
    def refresh_interval(self) -> Optional[pulumi.Input[int]]:
        """
        How often (in seconds) to refresh the values of the Table
        """
        return pulumi.get(self, "refresh_interval")

    @refresh_interval.setter
    def refresh_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "refresh_interval", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        """
        The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        """
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)

    @property
    @pulumi.getter(name="unitPrefix")
    def unit_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        (Metric by default) Must be "Metric" or "Binary"
        """
        return pulumi.get(self, "unit_prefix")

    @unit_prefix.setter
    def unit_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "unit_prefix", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the chart.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="vizOptions")
    def viz_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TableChartVizOptionArgs']]]]:
        """
        Plot-level customization options, associated with a publish statement
        """
        return pulumi.get(self, "viz_options")

    @viz_options.setter
    def viz_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TableChartVizOptionArgs']]]]):
        pulumi.set(self, "viz_options", value)


class TableChart(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_sampling: Optional[pulumi.Input[bool]] = None,
                 group_bies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hide_timestamp: Optional[pulumi.Input[bool]] = None,
                 max_delay: Optional[pulumi.Input[int]] = None,
                 minimum_resolution: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 program_text: Optional[pulumi.Input[str]] = None,
                 refresh_interval: Optional[pulumi.Input[int]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 unit_prefix: Optional[pulumi.Input[str]] = None,
                 viz_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableChartVizOptionArgs']]]]] = None,
                 __props__=None):
        """
        This special type of chart displays a Data Table. This Table can be grouped by a Dimension.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        # signalfx_list_chart.Logs-Exec_0:
        table0 = signalfx.TableChart("table0",
            description="beep",
            disable_sampling=False,
            group_bies=["ClusterName"],
            max_delay=0,
            program_text="A = data('cpu.usage.total').publish(label='CPU Total')")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the table chart.
        :param pulumi.Input[bool] disable_sampling: (false by default) If false, samples a subset of the output MTS, which improves UI performance
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_bies: Dimension to group by
        :param pulumi.Input[bool] hide_timestamp: (false by default) Whether to show the timestamp in the chart
        :param pulumi.Input[int] max_delay: How long (in seconds) to wait for late datapoints
        :param pulumi.Input[int] minimum_resolution: The minimum resolution (in seconds) to use for computing the underlying program
        :param pulumi.Input[str] name: Name of the table chart.
        :param pulumi.Input[str] program_text: The SignalFlow for your Data Table Chart
        :param pulumi.Input[int] refresh_interval: How often (in seconds) to refresh the values of the Table
        :param pulumi.Input[str] timezone: The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        :param pulumi.Input[str] unit_prefix: (Metric by default) Must be "Metric" or "Binary"
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableChartVizOptionArgs']]]] viz_options: Plot-level customization options, associated with a publish statement
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TableChartArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This special type of chart displays a Data Table. This Table can be grouped by a Dimension.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        # signalfx_list_chart.Logs-Exec_0:
        table0 = signalfx.TableChart("table0",
            description="beep",
            disable_sampling=False,
            group_bies=["ClusterName"],
            max_delay=0,
            program_text="A = data('cpu.usage.total').publish(label='CPU Total')")
        ```

        :param str resource_name: The name of the resource.
        :param TableChartArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TableChartArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_sampling: Optional[pulumi.Input[bool]] = None,
                 group_bies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 hide_timestamp: Optional[pulumi.Input[bool]] = None,
                 max_delay: Optional[pulumi.Input[int]] = None,
                 minimum_resolution: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 program_text: Optional[pulumi.Input[str]] = None,
                 refresh_interval: Optional[pulumi.Input[int]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 unit_prefix: Optional[pulumi.Input[str]] = None,
                 viz_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableChartVizOptionArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TableChartArgs.__new__(TableChartArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["disable_sampling"] = disable_sampling
            __props__.__dict__["group_bies"] = group_bies
            __props__.__dict__["hide_timestamp"] = hide_timestamp
            __props__.__dict__["max_delay"] = max_delay
            __props__.__dict__["minimum_resolution"] = minimum_resolution
            __props__.__dict__["name"] = name
            if program_text is None and not opts.urn:
                raise TypeError("Missing required property 'program_text'")
            __props__.__dict__["program_text"] = program_text
            __props__.__dict__["refresh_interval"] = refresh_interval
            __props__.__dict__["timezone"] = timezone
            __props__.__dict__["unit_prefix"] = unit_prefix
            __props__.__dict__["viz_options"] = viz_options
            __props__.__dict__["url"] = None
        super(TableChart, __self__).__init__(
            'signalfx:index/tableChart:TableChart',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            disable_sampling: Optional[pulumi.Input[bool]] = None,
            group_bies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            hide_timestamp: Optional[pulumi.Input[bool]] = None,
            max_delay: Optional[pulumi.Input[int]] = None,
            minimum_resolution: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            program_text: Optional[pulumi.Input[str]] = None,
            refresh_interval: Optional[pulumi.Input[int]] = None,
            timezone: Optional[pulumi.Input[str]] = None,
            unit_prefix: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            viz_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableChartVizOptionArgs']]]]] = None) -> 'TableChart':
        """
        Get an existing TableChart resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the table chart.
        :param pulumi.Input[bool] disable_sampling: (false by default) If false, samples a subset of the output MTS, which improves UI performance
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_bies: Dimension to group by
        :param pulumi.Input[bool] hide_timestamp: (false by default) Whether to show the timestamp in the chart
        :param pulumi.Input[int] max_delay: How long (in seconds) to wait for late datapoints
        :param pulumi.Input[int] minimum_resolution: The minimum resolution (in seconds) to use for computing the underlying program
        :param pulumi.Input[str] name: Name of the table chart.
        :param pulumi.Input[str] program_text: The SignalFlow for your Data Table Chart
        :param pulumi.Input[int] refresh_interval: How often (in seconds) to refresh the values of the Table
        :param pulumi.Input[str] timezone: The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        :param pulumi.Input[str] unit_prefix: (Metric by default) Must be "Metric" or "Binary"
        :param pulumi.Input[str] url: The URL of the chart.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TableChartVizOptionArgs']]]] viz_options: Plot-level customization options, associated with a publish statement
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TableChartState.__new__(_TableChartState)

        __props__.__dict__["description"] = description
        __props__.__dict__["disable_sampling"] = disable_sampling
        __props__.__dict__["group_bies"] = group_bies
        __props__.__dict__["hide_timestamp"] = hide_timestamp
        __props__.__dict__["max_delay"] = max_delay
        __props__.__dict__["minimum_resolution"] = minimum_resolution
        __props__.__dict__["name"] = name
        __props__.__dict__["program_text"] = program_text
        __props__.__dict__["refresh_interval"] = refresh_interval
        __props__.__dict__["timezone"] = timezone
        __props__.__dict__["unit_prefix"] = unit_prefix
        __props__.__dict__["url"] = url
        __props__.__dict__["viz_options"] = viz_options
        return TableChart(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the table chart.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableSampling")
    def disable_sampling(self) -> pulumi.Output[Optional[bool]]:
        """
        (false by default) If false, samples a subset of the output MTS, which improves UI performance
        """
        return pulumi.get(self, "disable_sampling")

    @property
    @pulumi.getter(name="groupBies")
    def group_bies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Dimension to group by
        """
        return pulumi.get(self, "group_bies")

    @property
    @pulumi.getter(name="hideTimestamp")
    def hide_timestamp(self) -> pulumi.Output[Optional[bool]]:
        """
        (false by default) Whether to show the timestamp in the chart
        """
        return pulumi.get(self, "hide_timestamp")

    @property
    @pulumi.getter(name="maxDelay")
    def max_delay(self) -> pulumi.Output[Optional[int]]:
        """
        How long (in seconds) to wait for late datapoints
        """
        return pulumi.get(self, "max_delay")

    @property
    @pulumi.getter(name="minimumResolution")
    def minimum_resolution(self) -> pulumi.Output[Optional[int]]:
        """
        The minimum resolution (in seconds) to use for computing the underlying program
        """
        return pulumi.get(self, "minimum_resolution")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the table chart.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> pulumi.Output[str]:
        """
        The SignalFlow for your Data Table Chart
        """
        return pulumi.get(self, "program_text")

    @property
    @pulumi.getter(name="refreshInterval")
    def refresh_interval(self) -> pulumi.Output[Optional[int]]:
        """
        How often (in seconds) to refresh the values of the Table
        """
        return pulumi.get(self, "refresh_interval")

    @property
    @pulumi.getter
    def timezone(self) -> pulumi.Output[Optional[str]]:
        """
        The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        """
        return pulumi.get(self, "timezone")

    @property
    @pulumi.getter(name="unitPrefix")
    def unit_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        (Metric by default) Must be "Metric" or "Binary"
        """
        return pulumi.get(self, "unit_prefix")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the chart.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="vizOptions")
    def viz_options(self) -> pulumi.Output[Optional[Sequence['outputs.TableChartVizOption']]]:
        """
        Plot-level customization options, associated with a publish statement
        """
        return pulumi.get(self, "viz_options")

