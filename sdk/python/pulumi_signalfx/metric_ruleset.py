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

__all__ = ['MetricRulesetArgs', 'MetricRuleset']

@pulumi.input_type
class MetricRulesetArgs:
    def __init__(__self__, *,
                 metric_name: pulumi.Input[str],
                 routing_rules: pulumi.Input[Sequence[pulumi.Input['MetricRulesetRoutingRuleArgs']]],
                 aggregation_rules: Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesetAggregationRuleArgs']]]] = None):
        """
        The set of arguments for constructing a MetricRuleset resource.
        :param pulumi.Input[str] metric_name: Name of the metric
        :param pulumi.Input[Sequence[pulumi.Input['MetricRulesetRoutingRuleArgs']]] routing_rules: Location to send the input metric
        :param pulumi.Input[Sequence[pulumi.Input['MetricRulesetAggregationRuleArgs']]] aggregation_rules: Aggregation rules in the ruleset
        """
        pulumi.set(__self__, "metric_name", metric_name)
        pulumi.set(__self__, "routing_rules", routing_rules)
        if aggregation_rules is not None:
            pulumi.set(__self__, "aggregation_rules", aggregation_rules)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Input[str]:
        """
        Name of the metric
        """
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "metric_name", value)

    @property
    @pulumi.getter(name="routingRules")
    def routing_rules(self) -> pulumi.Input[Sequence[pulumi.Input['MetricRulesetRoutingRuleArgs']]]:
        """
        Location to send the input metric
        """
        return pulumi.get(self, "routing_rules")

    @routing_rules.setter
    def routing_rules(self, value: pulumi.Input[Sequence[pulumi.Input['MetricRulesetRoutingRuleArgs']]]):
        pulumi.set(self, "routing_rules", value)

    @property
    @pulumi.getter(name="aggregationRules")
    def aggregation_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesetAggregationRuleArgs']]]]:
        """
        Aggregation rules in the ruleset
        """
        return pulumi.get(self, "aggregation_rules")

    @aggregation_rules.setter
    def aggregation_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesetAggregationRuleArgs']]]]):
        pulumi.set(self, "aggregation_rules", value)


@pulumi.input_type
class _MetricRulesetState:
    def __init__(__self__, *,
                 aggregation_rules: Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesetAggregationRuleArgs']]]] = None,
                 created: Optional[pulumi.Input[str]] = None,
                 creator: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[str]] = None,
                 last_updated_by: Optional[pulumi.Input[str]] = None,
                 last_updated_by_name: Optional[pulumi.Input[str]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 routing_rules: Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesetRoutingRuleArgs']]]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MetricRuleset resources.
        :param pulumi.Input[Sequence[pulumi.Input['MetricRulesetAggregationRuleArgs']]] aggregation_rules: Aggregation rules in the ruleset
        :param pulumi.Input[str] created: Timestamp of when the metric ruleset was created
        :param pulumi.Input[str] creator: ID of the creator of the metric ruleset
        :param pulumi.Input[str] last_updated: Timestamp of when the metric ruleset was last updated
        :param pulumi.Input[str] last_updated_by: ID of user who last updated the metric ruleset
        :param pulumi.Input[str] last_updated_by_name: Name of user who last updated this metric ruleset
        :param pulumi.Input[str] metric_name: Name of the metric
        :param pulumi.Input[Sequence[pulumi.Input['MetricRulesetRoutingRuleArgs']]] routing_rules: Location to send the input metric
        :param pulumi.Input[str] version: Version of the ruleset
        """
        if aggregation_rules is not None:
            pulumi.set(__self__, "aggregation_rules", aggregation_rules)
        if created is not None:
            pulumi.set(__self__, "created", created)
        if creator is not None:
            pulumi.set(__self__, "creator", creator)
        if last_updated is not None:
            pulumi.set(__self__, "last_updated", last_updated)
        if last_updated_by is not None:
            pulumi.set(__self__, "last_updated_by", last_updated_by)
        if last_updated_by_name is not None:
            pulumi.set(__self__, "last_updated_by_name", last_updated_by_name)
        if metric_name is not None:
            pulumi.set(__self__, "metric_name", metric_name)
        if routing_rules is not None:
            pulumi.set(__self__, "routing_rules", routing_rules)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="aggregationRules")
    def aggregation_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesetAggregationRuleArgs']]]]:
        """
        Aggregation rules in the ruleset
        """
        return pulumi.get(self, "aggregation_rules")

    @aggregation_rules.setter
    def aggregation_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesetAggregationRuleArgs']]]]):
        pulumi.set(self, "aggregation_rules", value)

    @property
    @pulumi.getter
    def created(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp of when the metric ruleset was created
        """
        return pulumi.get(self, "created")

    @created.setter
    def created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created", value)

    @property
    @pulumi.getter
    def creator(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the creator of the metric ruleset
        """
        return pulumi.get(self, "creator")

    @creator.setter
    def creator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creator", value)

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp of when the metric ruleset was last updated
        """
        return pulumi.get(self, "last_updated")

    @last_updated.setter
    def last_updated(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_updated", value)

    @property
    @pulumi.getter(name="lastUpdatedBy")
    def last_updated_by(self) -> Optional[pulumi.Input[str]]:
        """
        ID of user who last updated the metric ruleset
        """
        return pulumi.get(self, "last_updated_by")

    @last_updated_by.setter
    def last_updated_by(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_updated_by", value)

    @property
    @pulumi.getter(name="lastUpdatedByName")
    def last_updated_by_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of user who last updated this metric ruleset
        """
        return pulumi.get(self, "last_updated_by_name")

    @last_updated_by_name.setter
    def last_updated_by_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_updated_by_name", value)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the metric
        """
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metric_name", value)

    @property
    @pulumi.getter(name="routingRules")
    def routing_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesetRoutingRuleArgs']]]]:
        """
        Location to send the input metric
        """
        return pulumi.get(self, "routing_rules")

    @routing_rules.setter
    def routing_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['MetricRulesetRoutingRuleArgs']]]]):
        pulumi.set(self, "routing_rules", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Version of the ruleset
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class MetricRuleset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aggregation_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricRulesetAggregationRuleArgs']]]]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 routing_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricRulesetRoutingRuleArgs']]]]] = None,
                 __props__=None):
        """
        Provides an Observability Cloud resource for managing metric rulesets.

        > **NOTE** When managing metric rulesets to drop data use a session token for an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

        ## Arguments

        The following arguments are supported in the resource block:

        * `metric_name` - (Required) Name of the input metric
        * `aggregation_rules` - (Optional) List of aggregation rules for the metric
          * `enabled` - (Required) When false, this rule will not generate aggregated MTSs
          * `name` - (Optional) name of the aggregation rule
          * `matcher` - (Required) Matcher object
            * `type` - (Required) Type of matcher. Must always be "dimension"
            * `filters` - (Optional) List of filters to filter the set of input MTSs
              * `property` - (Required) - Name of the dimension
              * `property_value` - (Required) - Value of the dimension
              * `not` - When true, this filter will match all values not matching the property_values
          * `aggregator` - (Required) - Aggregator object
            * `type` - (Required) Type of aggregator. Must always be "rollup"
            * `dimensions` - (Required) List of dimensions to either be kept or dropped in the new aggregated MTSs
            * `drop_dimensions` - (Required) when true, the specified dimensions will be dropped from the aggregated MTSs
            * `output_name` - (Required) name of the new aggregated metric
        * `routing_rule` - (Required) Routing Rule object
          * `destination` - (Required) - end destination of the input metric. Must be `RealTime` or `Drop`

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricRulesetAggregationRuleArgs']]]] aggregation_rules: Aggregation rules in the ruleset
        :param pulumi.Input[str] metric_name: Name of the metric
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricRulesetRoutingRuleArgs']]]] routing_rules: Location to send the input metric
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MetricRulesetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Observability Cloud resource for managing metric rulesets.

        > **NOTE** When managing metric rulesets to drop data use a session token for an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

        ## Arguments

        The following arguments are supported in the resource block:

        * `metric_name` - (Required) Name of the input metric
        * `aggregation_rules` - (Optional) List of aggregation rules for the metric
          * `enabled` - (Required) When false, this rule will not generate aggregated MTSs
          * `name` - (Optional) name of the aggregation rule
          * `matcher` - (Required) Matcher object
            * `type` - (Required) Type of matcher. Must always be "dimension"
            * `filters` - (Optional) List of filters to filter the set of input MTSs
              * `property` - (Required) - Name of the dimension
              * `property_value` - (Required) - Value of the dimension
              * `not` - When true, this filter will match all values not matching the property_values
          * `aggregator` - (Required) - Aggregator object
            * `type` - (Required) Type of aggregator. Must always be "rollup"
            * `dimensions` - (Required) List of dimensions to either be kept or dropped in the new aggregated MTSs
            * `drop_dimensions` - (Required) when true, the specified dimensions will be dropped from the aggregated MTSs
            * `output_name` - (Required) name of the new aggregated metric
        * `routing_rule` - (Required) Routing Rule object
          * `destination` - (Required) - end destination of the input metric. Must be `RealTime` or `Drop`

        :param str resource_name: The name of the resource.
        :param MetricRulesetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MetricRulesetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aggregation_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricRulesetAggregationRuleArgs']]]]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 routing_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricRulesetRoutingRuleArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MetricRulesetArgs.__new__(MetricRulesetArgs)

            __props__.__dict__["aggregation_rules"] = aggregation_rules
            if metric_name is None and not opts.urn:
                raise TypeError("Missing required property 'metric_name'")
            __props__.__dict__["metric_name"] = metric_name
            if routing_rules is None and not opts.urn:
                raise TypeError("Missing required property 'routing_rules'")
            __props__.__dict__["routing_rules"] = routing_rules
            __props__.__dict__["created"] = None
            __props__.__dict__["creator"] = None
            __props__.__dict__["last_updated"] = None
            __props__.__dict__["last_updated_by"] = None
            __props__.__dict__["last_updated_by_name"] = None
            __props__.__dict__["version"] = None
        super(MetricRuleset, __self__).__init__(
            'signalfx:index/metricRuleset:MetricRuleset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aggregation_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricRulesetAggregationRuleArgs']]]]] = None,
            created: Optional[pulumi.Input[str]] = None,
            creator: Optional[pulumi.Input[str]] = None,
            last_updated: Optional[pulumi.Input[str]] = None,
            last_updated_by: Optional[pulumi.Input[str]] = None,
            last_updated_by_name: Optional[pulumi.Input[str]] = None,
            metric_name: Optional[pulumi.Input[str]] = None,
            routing_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricRulesetRoutingRuleArgs']]]]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'MetricRuleset':
        """
        Get an existing MetricRuleset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricRulesetAggregationRuleArgs']]]] aggregation_rules: Aggregation rules in the ruleset
        :param pulumi.Input[str] created: Timestamp of when the metric ruleset was created
        :param pulumi.Input[str] creator: ID of the creator of the metric ruleset
        :param pulumi.Input[str] last_updated: Timestamp of when the metric ruleset was last updated
        :param pulumi.Input[str] last_updated_by: ID of user who last updated the metric ruleset
        :param pulumi.Input[str] last_updated_by_name: Name of user who last updated this metric ruleset
        :param pulumi.Input[str] metric_name: Name of the metric
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MetricRulesetRoutingRuleArgs']]]] routing_rules: Location to send the input metric
        :param pulumi.Input[str] version: Version of the ruleset
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MetricRulesetState.__new__(_MetricRulesetState)

        __props__.__dict__["aggregation_rules"] = aggregation_rules
        __props__.__dict__["created"] = created
        __props__.__dict__["creator"] = creator
        __props__.__dict__["last_updated"] = last_updated
        __props__.__dict__["last_updated_by"] = last_updated_by
        __props__.__dict__["last_updated_by_name"] = last_updated_by_name
        __props__.__dict__["metric_name"] = metric_name
        __props__.__dict__["routing_rules"] = routing_rules
        __props__.__dict__["version"] = version
        return MetricRuleset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aggregationRules")
    def aggregation_rules(self) -> pulumi.Output[Optional[Sequence['outputs.MetricRulesetAggregationRule']]]:
        """
        Aggregation rules in the ruleset
        """
        return pulumi.get(self, "aggregation_rules")

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        Timestamp of when the metric ruleset was created
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def creator(self) -> pulumi.Output[str]:
        """
        ID of the creator of the metric ruleset
        """
        return pulumi.get(self, "creator")

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> pulumi.Output[str]:
        """
        Timestamp of when the metric ruleset was last updated
        """
        return pulumi.get(self, "last_updated")

    @property
    @pulumi.getter(name="lastUpdatedBy")
    def last_updated_by(self) -> pulumi.Output[str]:
        """
        ID of user who last updated the metric ruleset
        """
        return pulumi.get(self, "last_updated_by")

    @property
    @pulumi.getter(name="lastUpdatedByName")
    def last_updated_by_name(self) -> pulumi.Output[str]:
        """
        Name of user who last updated this metric ruleset
        """
        return pulumi.get(self, "last_updated_by_name")

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Output[str]:
        """
        Name of the metric
        """
        return pulumi.get(self, "metric_name")

    @property
    @pulumi.getter(name="routingRules")
    def routing_rules(self) -> pulumi.Output[Sequence['outputs.MetricRulesetRoutingRule']]:
        """
        Location to send the input metric
        """
        return pulumi.get(self, "routing_rules")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Version of the ruleset
        """
        return pulumi.get(self, "version")

