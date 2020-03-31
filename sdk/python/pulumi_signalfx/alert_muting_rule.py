# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class AlertMutingRule(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description for this muting rule
    """
    detectors: pulumi.Output[list]
    """
    A convenience attribute that associated this muting rule with specific detector ids.
    """
    effective_start_time: pulumi.Output[float]
    filters: pulumi.Output[list]
    """
    Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.

      * `negated` (`bool`) - Determines if this is a "not" filter. Defaults to `false`.
      * `property` (`str`) - The property to filter.
      * `property_value` (`str`) - The property value to filter.
    """
    start_time: pulumi.Output[float]
    """
    Starting time of an alert muting rule as a Unit time stamp in seconds.
    """
    stop_time: pulumi.Output[float]
    """
    Starting time of an alert muting rule as a Unix time stamp in seconds.
    """
    def __init__(__self__, resource_name, opts=None, description=None, detectors=None, filters=None, start_time=None, stop_time=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a SignalFx resource for managing alert muting rules. See [Mute Notifications](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html) for more information.

        > **WARNING** SignalFx does not allow the start time of a **currently active** muting rule to be modified. As such, attempting to modify a currently active rule will destroy the existing rule and create a new rule. This may result in the emission of notifications.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/alert_muting_rule.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description for this muting rule
        :param pulumi.Input[list] detectors: A convenience attribute that associated this muting rule with specific detector ids.
        :param pulumi.Input[list] filters: Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.
        :param pulumi.Input[float] start_time: Starting time of an alert muting rule as a Unit time stamp in seconds.
        :param pulumi.Input[float] stop_time: Starting time of an alert muting rule as a Unix time stamp in seconds.

        The **filters** object supports the following:

          * `negated` (`pulumi.Input[bool]`) - Determines if this is a "not" filter. Defaults to `false`.
          * `property` (`pulumi.Input[str]`) - The property to filter.
          * `property_value` (`pulumi.Input[str]`) - The property value to filter.
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

            if description is None:
                raise TypeError("Missing required property 'description'")
            __props__['description'] = description
            __props__['detectors'] = detectors
            if filters is None:
                raise TypeError("Missing required property 'filters'")
            __props__['filters'] = filters
            if start_time is None:
                raise TypeError("Missing required property 'start_time'")
            __props__['start_time'] = start_time
            __props__['stop_time'] = stop_time
            __props__['effective_start_time'] = None
        super(AlertMutingRule, __self__).__init__(
            'signalfx:index/alertMutingRule:AlertMutingRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, detectors=None, effective_start_time=None, filters=None, start_time=None, stop_time=None):
        """
        Get an existing AlertMutingRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description for this muting rule
        :param pulumi.Input[list] detectors: A convenience attribute that associated this muting rule with specific detector ids.
        :param pulumi.Input[list] filters: Filters for this rule. See [Creating muting rules from scratch](https://docs.signalfx.com/en/latest/detect-alert/mute-notifications.html#rule-from-scratch) for more information.
        :param pulumi.Input[float] start_time: Starting time of an alert muting rule as a Unit time stamp in seconds.
        :param pulumi.Input[float] stop_time: Starting time of an alert muting rule as a Unix time stamp in seconds.

        The **filters** object supports the following:

          * `negated` (`pulumi.Input[bool]`) - Determines if this is a "not" filter. Defaults to `false`.
          * `property` (`pulumi.Input[str]`) - The property to filter.
          * `property_value` (`pulumi.Input[str]`) - The property value to filter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["detectors"] = detectors
        __props__["effective_start_time"] = effective_start_time
        __props__["filters"] = filters
        __props__["start_time"] = start_time
        __props__["stop_time"] = stop_time
        return AlertMutingRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
