# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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
from . import outputs
from ._inputs import *

__all__ = ['AlertMutingRuleArgs', 'AlertMutingRule']

@pulumi.input_type
class AlertMutingRuleArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[_builtins.str],
                 start_time: pulumi.Input[_builtins.int],
                 detectors: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input['AlertMutingRuleFilterArgs']]]] = None,
                 recurrence: Optional[pulumi.Input['AlertMutingRuleRecurrenceArgs']] = None,
                 stop_time: Optional[pulumi.Input[_builtins.int]] = None):
        """
        The set of arguments for constructing a AlertMutingRule resource.
        :param pulumi.Input[_builtins.str] description: The description for this muting rule
        :param pulumi.Input[_builtins.int] start_time: Starting time of an alert muting rule as a Unit time stamp in seconds.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] detectors: A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
        :param pulumi.Input[Sequence[pulumi.Input['AlertMutingRuleFilterArgs']]] filters: Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
        :param pulumi.Input['AlertMutingRuleRecurrenceArgs'] recurrence: Defines the recurrence of the muting rule. Allows setting a recurring muting rule based on specified days or weeks.
        :param pulumi.Input[_builtins.int] stop_time: Stop time of an alert muting rule as a Unix time stamp in seconds.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "start_time", start_time)
        if detectors is not None:
            pulumi.set(__self__, "detectors", detectors)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if recurrence is not None:
            pulumi.set(__self__, "recurrence", recurrence)
        if stop_time is not None:
            pulumi.set(__self__, "stop_time", stop_time)

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Input[_builtins.str]:
        """
        The description for this muting rule
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Input[_builtins.int]:
        """
        Starting time of an alert muting rule as a Unit time stamp in seconds.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: pulumi.Input[_builtins.int]):
        pulumi.set(self, "start_time", value)

    @_builtins.property
    @pulumi.getter
    def detectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
        """
        return pulumi.get(self, "detectors")

    @detectors.setter
    def detectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "detectors", value)

    @_builtins.property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlertMutingRuleFilterArgs']]]]:
        """
        Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlertMutingRuleFilterArgs']]]]):
        pulumi.set(self, "filters", value)

    @_builtins.property
    @pulumi.getter
    def recurrence(self) -> Optional[pulumi.Input['AlertMutingRuleRecurrenceArgs']]:
        """
        Defines the recurrence of the muting rule. Allows setting a recurring muting rule based on specified days or weeks.
        """
        return pulumi.get(self, "recurrence")

    @recurrence.setter
    def recurrence(self, value: Optional[pulumi.Input['AlertMutingRuleRecurrenceArgs']]):
        pulumi.set(self, "recurrence", value)

    @_builtins.property
    @pulumi.getter(name="stopTime")
    def stop_time(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Stop time of an alert muting rule as a Unix time stamp in seconds.
        """
        return pulumi.get(self, "stop_time")

    @stop_time.setter
    def stop_time(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "stop_time", value)


@pulumi.input_type
class _AlertMutingRuleState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 detectors: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 effective_start_time: Optional[pulumi.Input[_builtins.int]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input['AlertMutingRuleFilterArgs']]]] = None,
                 recurrence: Optional[pulumi.Input['AlertMutingRuleRecurrenceArgs']] = None,
                 start_time: Optional[pulumi.Input[_builtins.int]] = None,
                 stop_time: Optional[pulumi.Input[_builtins.int]] = None):
        """
        Input properties used for looking up and filtering AlertMutingRule resources.
        :param pulumi.Input[_builtins.str] description: The description for this muting rule
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] detectors: A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
        :param pulumi.Input[Sequence[pulumi.Input['AlertMutingRuleFilterArgs']]] filters: Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
        :param pulumi.Input['AlertMutingRuleRecurrenceArgs'] recurrence: Defines the recurrence of the muting rule. Allows setting a recurring muting rule based on specified days or weeks.
        :param pulumi.Input[_builtins.int] start_time: Starting time of an alert muting rule as a Unit time stamp in seconds.
        :param pulumi.Input[_builtins.int] stop_time: Stop time of an alert muting rule as a Unix time stamp in seconds.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if detectors is not None:
            pulumi.set(__self__, "detectors", detectors)
        if effective_start_time is not None:
            pulumi.set(__self__, "effective_start_time", effective_start_time)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if recurrence is not None:
            pulumi.set(__self__, "recurrence", recurrence)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if stop_time is not None:
            pulumi.set(__self__, "stop_time", stop_time)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The description for this muting rule
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter
    def detectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
        """
        return pulumi.get(self, "detectors")

    @detectors.setter
    def detectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "detectors", value)

    @_builtins.property
    @pulumi.getter(name="effectiveStartTime")
    def effective_start_time(self) -> Optional[pulumi.Input[_builtins.int]]:
        return pulumi.get(self, "effective_start_time")

    @effective_start_time.setter
    def effective_start_time(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "effective_start_time", value)

    @_builtins.property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlertMutingRuleFilterArgs']]]]:
        """
        Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlertMutingRuleFilterArgs']]]]):
        pulumi.set(self, "filters", value)

    @_builtins.property
    @pulumi.getter
    def recurrence(self) -> Optional[pulumi.Input['AlertMutingRuleRecurrenceArgs']]:
        """
        Defines the recurrence of the muting rule. Allows setting a recurring muting rule based on specified days or weeks.
        """
        return pulumi.get(self, "recurrence")

    @recurrence.setter
    def recurrence(self, value: Optional[pulumi.Input['AlertMutingRuleRecurrenceArgs']]):
        pulumi.set(self, "recurrence", value)

    @_builtins.property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Starting time of an alert muting rule as a Unit time stamp in seconds.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "start_time", value)

    @_builtins.property
    @pulumi.getter(name="stopTime")
    def stop_time(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Stop time of an alert muting rule as a Unix time stamp in seconds.
        """
        return pulumi.get(self, "stop_time")

    @stop_time.setter
    def stop_time(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "stop_time", value)


@pulumi.type_token("signalfx:index/alertMutingRule:AlertMutingRule")
class AlertMutingRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 detectors: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AlertMutingRuleFilterArgs', 'AlertMutingRuleFilterArgsDict']]]]] = None,
                 recurrence: Optional[pulumi.Input[Union['AlertMutingRuleRecurrenceArgs', 'AlertMutingRuleRecurrenceArgsDict']]] = None,
                 start_time: Optional[pulumi.Input[_builtins.int]] = None,
                 stop_time: Optional[pulumi.Input[_builtins.int]] = None,
                 __props__=None):
        """
        Provides a Splunk Observability Cloud resource for managing alert muting rules. See [Mute Notifications](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html) for more information.

        Splunk Observability Cloud currently allows linking an alert muting rule with only one detector ID. Specifying multiple detector IDs makes the muting rule obsolete.

        > **WARNING** Splunk Observability Cloud does not allow the start time of a **currently active** muting rule to be modified. Attempting to modify a currently active rule destroys the existing rule and creates a new rule. This might result in the emission of notifications.

        ## Example

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        rool_mooter_one = signalfx.AlertMutingRule("rool_mooter_one",
            description="mooted it NEW",
            start_time=1573063243,
            stop_time=0,
            detectors=[some_detector_id],
            filters=[{
                "property": "foo",
                "property_value": "bar",
            }],
            recurrence={
                "unit": "d",
                "value": 2,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] description: The description for this muting rule
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] detectors: A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AlertMutingRuleFilterArgs', 'AlertMutingRuleFilterArgsDict']]]] filters: Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
        :param pulumi.Input[Union['AlertMutingRuleRecurrenceArgs', 'AlertMutingRuleRecurrenceArgsDict']] recurrence: Defines the recurrence of the muting rule. Allows setting a recurring muting rule based on specified days or weeks.
        :param pulumi.Input[_builtins.int] start_time: Starting time of an alert muting rule as a Unit time stamp in seconds.
        :param pulumi.Input[_builtins.int] stop_time: Stop time of an alert muting rule as a Unix time stamp in seconds.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AlertMutingRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Splunk Observability Cloud resource for managing alert muting rules. See [Mute Notifications](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html) for more information.

        Splunk Observability Cloud currently allows linking an alert muting rule with only one detector ID. Specifying multiple detector IDs makes the muting rule obsolete.

        > **WARNING** Splunk Observability Cloud does not allow the start time of a **currently active** muting rule to be modified. Attempting to modify a currently active rule destroys the existing rule and creates a new rule. This might result in the emission of notifications.

        ## Example

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        rool_mooter_one = signalfx.AlertMutingRule("rool_mooter_one",
            description="mooted it NEW",
            start_time=1573063243,
            stop_time=0,
            detectors=[some_detector_id],
            filters=[{
                "property": "foo",
                "property_value": "bar",
            }],
            recurrence={
                "unit": "d",
                "value": 2,
            })
        ```

        :param str resource_name: The name of the resource.
        :param AlertMutingRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlertMutingRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 detectors: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AlertMutingRuleFilterArgs', 'AlertMutingRuleFilterArgsDict']]]]] = None,
                 recurrence: Optional[pulumi.Input[Union['AlertMutingRuleRecurrenceArgs', 'AlertMutingRuleRecurrenceArgsDict']]] = None,
                 start_time: Optional[pulumi.Input[_builtins.int]] = None,
                 stop_time: Optional[pulumi.Input[_builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AlertMutingRuleArgs.__new__(AlertMutingRuleArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["detectors"] = detectors
            __props__.__dict__["filters"] = filters
            __props__.__dict__["recurrence"] = recurrence
            if start_time is None and not opts.urn:
                raise TypeError("Missing required property 'start_time'")
            __props__.__dict__["start_time"] = start_time
            __props__.__dict__["stop_time"] = stop_time
            __props__.__dict__["effective_start_time"] = None
        super(AlertMutingRule, __self__).__init__(
            'signalfx:index/alertMutingRule:AlertMutingRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[_builtins.str]] = None,
            detectors: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            effective_start_time: Optional[pulumi.Input[_builtins.int]] = None,
            filters: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AlertMutingRuleFilterArgs', 'AlertMutingRuleFilterArgsDict']]]]] = None,
            recurrence: Optional[pulumi.Input[Union['AlertMutingRuleRecurrenceArgs', 'AlertMutingRuleRecurrenceArgsDict']]] = None,
            start_time: Optional[pulumi.Input[_builtins.int]] = None,
            stop_time: Optional[pulumi.Input[_builtins.int]] = None) -> 'AlertMutingRule':
        """
        Get an existing AlertMutingRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] description: The description for this muting rule
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] detectors: A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
        :param pulumi.Input[Sequence[pulumi.Input[Union['AlertMutingRuleFilterArgs', 'AlertMutingRuleFilterArgsDict']]]] filters: Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
        :param pulumi.Input[Union['AlertMutingRuleRecurrenceArgs', 'AlertMutingRuleRecurrenceArgsDict']] recurrence: Defines the recurrence of the muting rule. Allows setting a recurring muting rule based on specified days or weeks.
        :param pulumi.Input[_builtins.int] start_time: Starting time of an alert muting rule as a Unit time stamp in seconds.
        :param pulumi.Input[_builtins.int] stop_time: Stop time of an alert muting rule as a Unix time stamp in seconds.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlertMutingRuleState.__new__(_AlertMutingRuleState)

        __props__.__dict__["description"] = description
        __props__.__dict__["detectors"] = detectors
        __props__.__dict__["effective_start_time"] = effective_start_time
        __props__.__dict__["filters"] = filters
        __props__.__dict__["recurrence"] = recurrence
        __props__.__dict__["start_time"] = start_time
        __props__.__dict__["stop_time"] = stop_time
        return AlertMutingRule(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Output[_builtins.str]:
        """
        The description for this muting rule
        """
        return pulumi.get(self, "description")

    @_builtins.property
    @pulumi.getter
    def detectors(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
        """
        return pulumi.get(self, "detectors")

    @_builtins.property
    @pulumi.getter(name="effectiveStartTime")
    def effective_start_time(self) -> pulumi.Output[_builtins.int]:
        return pulumi.get(self, "effective_start_time")

    @_builtins.property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Optional[Sequence['outputs.AlertMutingRuleFilter']]]:
        """
        Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
        """
        return pulumi.get(self, "filters")

    @_builtins.property
    @pulumi.getter
    def recurrence(self) -> pulumi.Output[Optional['outputs.AlertMutingRuleRecurrence']]:
        """
        Defines the recurrence of the muting rule. Allows setting a recurring muting rule based on specified days or weeks.
        """
        return pulumi.get(self, "recurrence")

    @_builtins.property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[_builtins.int]:
        """
        Starting time of an alert muting rule as a Unit time stamp in seconds.
        """
        return pulumi.get(self, "start_time")

    @_builtins.property
    @pulumi.getter(name="stopTime")
    def stop_time(self) -> pulumi.Output[Optional[_builtins.int]]:
        """
        Stop time of an alert muting rule as a Unix time stamp in seconds.
        """
        return pulumi.get(self, "stop_time")

