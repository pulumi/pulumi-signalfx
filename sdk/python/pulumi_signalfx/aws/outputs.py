# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'IntegrationCustomNamespaceSyncRule',
    'IntegrationNamespaceSyncRule',
    'GetServicesServiceResult',
]

@pulumi.output_type
class IntegrationCustomNamespaceSyncRule(dict):
    def __init__(__self__, *,
                 namespace: str,
                 default_action: Optional[str] = None,
                 filter_action: Optional[str] = None,
                 filter_source: Optional[str] = None):
        """
        :param str namespace: An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
        :param str default_action: Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        :param str filter_action: Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        :param str filter_source: Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
        pulumi.set(__self__, "namespace", namespace)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if filter_action is not None:
            pulumi.set(__self__, "filter_action", filter_action)
        if filter_source is not None:
            pulumi.set(__self__, "filter_source", filter_source)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[str]:
        """
        Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="filterAction")
    def filter_action(self) -> Optional[str]:
        """
        Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "filter_action")

    @property
    @pulumi.getter(name="filterSource")
    def filter_source(self) -> Optional[str]:
        """
        Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
        return pulumi.get(self, "filter_source")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IntegrationNamespaceSyncRule(dict):
    def __init__(__self__, *,
                 namespace: str,
                 default_action: Optional[str] = None,
                 filter_action: Optional[str] = None,
                 filter_source: Optional[str] = None):
        """
        :param str namespace: An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
        :param str default_action: Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        :param str filter_action: Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        :param str filter_source: Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
        pulumi.set(__self__, "namespace", namespace)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if filter_action is not None:
            pulumi.set(__self__, "filter_action", filter_action)
        if filter_source is not None:
            pulumi.set(__self__, "filter_source", filter_source)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[str]:
        """
        Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="filterAction")
    def filter_action(self) -> Optional[str]:
        """
        Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "filter_action")

    @property
    @pulumi.getter(name="filterSource")
    def filter_source(self) -> Optional[str]:
        """
        Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
        return pulumi.get(self, "filter_source")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetServicesServiceResult(dict):
    def __init__(__self__, *,
                 name: str):
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


