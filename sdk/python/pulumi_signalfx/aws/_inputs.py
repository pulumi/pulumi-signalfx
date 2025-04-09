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
from .. import _utilities

__all__ = [
    'IntegrationCustomNamespaceSyncRuleArgs',
    'IntegrationCustomNamespaceSyncRuleArgsDict',
    'IntegrationMetricStatsToSyncArgs',
    'IntegrationMetricStatsToSyncArgsDict',
    'IntegrationNamespaceSyncRuleArgs',
    'IntegrationNamespaceSyncRuleArgsDict',
]

MYPY = False

if not MYPY:
    class IntegrationCustomNamespaceSyncRuleArgsDict(TypedDict):
        namespace: pulumi.Input[builtins.str]
        """
        An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability Cloud. See the AWS documentation on publishing metrics for more information.
        """
        default_action: NotRequired[pulumi.Input[builtins.str]]
        """
        Controls the Splunk Observability Cloud default behavior for processing data from an AWS namespace. Splunk Observability Cloud ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability Cloud treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        """
        filter_action: NotRequired[pulumi.Input[builtins.str]]
        """
        Controls how Splunk Observability Cloud processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        """
        filter_source: NotRequired[pulumi.Input[builtins.str]]
        """
        Expression that selects the data that Splunk Observability Cloud should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
elif False:
    IntegrationCustomNamespaceSyncRuleArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IntegrationCustomNamespaceSyncRuleArgs:
    def __init__(__self__, *,
                 namespace: pulumi.Input[builtins.str],
                 default_action: Optional[pulumi.Input[builtins.str]] = None,
                 filter_action: Optional[pulumi.Input[builtins.str]] = None,
                 filter_source: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] namespace: An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability Cloud. See the AWS documentation on publishing metrics for more information.
        :param pulumi.Input[builtins.str] default_action: Controls the Splunk Observability Cloud default behavior for processing data from an AWS namespace. Splunk Observability Cloud ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability Cloud treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        :param pulumi.Input[builtins.str] filter_action: Controls how Splunk Observability Cloud processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        :param pulumi.Input[builtins.str] filter_source: Expression that selects the data that Splunk Observability Cloud should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
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
    def namespace(self) -> pulumi.Input[builtins.str]:
        """
        An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability Cloud. See the AWS documentation on publishing metrics for more information.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Controls the Splunk Observability Cloud default behavior for processing data from an AWS namespace. Splunk Observability Cloud ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability Cloud treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter(name="filterAction")
    def filter_action(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Controls how Splunk Observability Cloud processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "filter_action")

    @filter_action.setter
    def filter_action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "filter_action", value)

    @property
    @pulumi.getter(name="filterSource")
    def filter_source(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Expression that selects the data that Splunk Observability Cloud should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
        return pulumi.get(self, "filter_source")

    @filter_source.setter
    def filter_source(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "filter_source", value)


if not MYPY:
    class IntegrationMetricStatsToSyncArgsDict(TypedDict):
        metric: pulumi.Input[builtins.str]
        """
        AWS metric that you want to pick statistics for
        """
        namespace: pulumi.Input[builtins.str]
        """
        An AWS namespace having AWS metric that you want to pick statistics for
        """
        stats: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]
        """
        AWS statistics you want to collect
        """
elif False:
    IntegrationMetricStatsToSyncArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IntegrationMetricStatsToSyncArgs:
    def __init__(__self__, *,
                 metric: pulumi.Input[builtins.str],
                 namespace: pulumi.Input[builtins.str],
                 stats: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        """
        :param pulumi.Input[builtins.str] metric: AWS metric that you want to pick statistics for
        :param pulumi.Input[builtins.str] namespace: An AWS namespace having AWS metric that you want to pick statistics for
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] stats: AWS statistics you want to collect
        """
        pulumi.set(__self__, "metric", metric)
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "stats", stats)

    @property
    @pulumi.getter
    def metric(self) -> pulumi.Input[builtins.str]:
        """
        AWS metric that you want to pick statistics for
        """
        return pulumi.get(self, "metric")

    @metric.setter
    def metric(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "metric", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[builtins.str]:
        """
        An AWS namespace having AWS metric that you want to pick statistics for
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def stats(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        AWS statistics you want to collect
        """
        return pulumi.get(self, "stats")

    @stats.setter
    def stats(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "stats", value)


if not MYPY:
    class IntegrationNamespaceSyncRuleArgsDict(TypedDict):
        namespace: pulumi.Input[builtins.str]
        """
        An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability Cloud. See `services` field description below for additional information.
        """
        default_action: NotRequired[pulumi.Input[builtins.str]]
        """
        Controls the Splunk Observability Cloud default behavior for processing data from an AWS namespace. Splunk Observability Cloud ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability Cloud treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        """
        filter_action: NotRequired[pulumi.Input[builtins.str]]
        """
        Controls how Splunk Observability Cloud processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        """
        filter_source: NotRequired[pulumi.Input[builtins.str]]
        """
        Expression that selects the data that Splunk Observability Cloud should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
elif False:
    IntegrationNamespaceSyncRuleArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IntegrationNamespaceSyncRuleArgs:
    def __init__(__self__, *,
                 namespace: pulumi.Input[builtins.str],
                 default_action: Optional[pulumi.Input[builtins.str]] = None,
                 filter_action: Optional[pulumi.Input[builtins.str]] = None,
                 filter_source: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] namespace: An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability Cloud. See `services` field description below for additional information.
        :param pulumi.Input[builtins.str] default_action: Controls the Splunk Observability Cloud default behavior for processing data from an AWS namespace. Splunk Observability Cloud ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability Cloud treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        :param pulumi.Input[builtins.str] filter_action: Controls how Splunk Observability Cloud processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        :param pulumi.Input[builtins.str] filter_source: Expression that selects the data that Splunk Observability Cloud should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
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
    def namespace(self) -> pulumi.Input[builtins.str]:
        """
        An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability Cloud. See `services` field description below for additional information.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Controls the Splunk Observability Cloud default behavior for processing data from an AWS namespace. Splunk Observability Cloud ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability Cloud treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter(name="filterAction")
    def filter_action(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Controls how Splunk Observability Cloud processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "filter_action")

    @filter_action.setter
    def filter_action(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "filter_action", value)

    @property
    @pulumi.getter(name="filterSource")
    def filter_source(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Expression that selects the data that Splunk Observability Cloud should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
        return pulumi.get(self, "filter_source")

    @filter_source.setter
    def filter_source(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "filter_source", value)


