# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'IntegrationCustomNamespaceSyncRule',
    'IntegrationMetricStatsToSync',
    'IntegrationNamespaceSyncRule',
]

@pulumi.output_type
class IntegrationCustomNamespaceSyncRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "defaultAction":
            suggest = "default_action"
        elif key == "filterAction":
            suggest = "filter_action"
        elif key == "filterSource":
            suggest = "filter_source"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IntegrationCustomNamespaceSyncRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IntegrationCustomNamespaceSyncRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IntegrationCustomNamespaceSyncRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 namespace: str,
                 default_action: Optional[str] = None,
                 filter_action: Optional[str] = None,
                 filter_source: Optional[str] = None):
        """
        :param str namespace: An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See `services` field description below for additional information.
        :param str default_action: Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        :param str filter_action: Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        :param str filter_source: Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
        IntegrationCustomNamespaceSyncRule._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            namespace=namespace,
            default_action=default_action,
            filter_action=filter_action,
            filter_source=filter_source,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             namespace: str,
             default_action: Optional[str] = None,
             filter_action: Optional[str] = None,
             filter_source: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("namespace", namespace)
        if default_action is not None:
            _setter("default_action", default_action)
        if filter_action is not None:
            _setter("filter_action", filter_action)
        if filter_source is not None:
            _setter("filter_source", filter_source)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See `services` field description below for additional information.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[str]:
        """
        Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="filterAction")
    def filter_action(self) -> Optional[str]:
        """
        Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "filter_action")

    @property
    @pulumi.getter(name="filterSource")
    def filter_source(self) -> Optional[str]:
        """
        Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
        return pulumi.get(self, "filter_source")


@pulumi.output_type
class IntegrationMetricStatsToSync(dict):
    def __init__(__self__, *,
                 metric: str,
                 namespace: str,
                 stats: Sequence[str]):
        """
        :param str metric: AWS metric that you want to pick statistics for
        :param str namespace: An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See `services` field description below for additional information.
        :param Sequence[str] stats: AWS statistics you want to collect
        """
        IntegrationMetricStatsToSync._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            metric=metric,
            namespace=namespace,
            stats=stats,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             metric: str,
             namespace: str,
             stats: Sequence[str],
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("metric", metric)
        _setter("namespace", namespace)
        _setter("stats", stats)

    @property
    @pulumi.getter
    def metric(self) -> str:
        """
        AWS metric that you want to pick statistics for
        """
        return pulumi.get(self, "metric")

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See `services` field description below for additional information.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def stats(self) -> Sequence[str]:
        """
        AWS statistics you want to collect
        """
        return pulumi.get(self, "stats")


@pulumi.output_type
class IntegrationNamespaceSyncRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "defaultAction":
            suggest = "default_action"
        elif key == "filterAction":
            suggest = "filter_action"
        elif key == "filterSource":
            suggest = "filter_source"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IntegrationNamespaceSyncRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IntegrationNamespaceSyncRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IntegrationNamespaceSyncRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 namespace: str,
                 default_action: Optional[str] = None,
                 filter_action: Optional[str] = None,
                 filter_source: Optional[str] = None):
        """
        :param str namespace: An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See `services` field description below for additional information.
        :param str default_action: Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        :param str filter_action: Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        :param str filter_source: Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
        IntegrationNamespaceSyncRule._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            namespace=namespace,
            default_action=default_action,
            filter_action=filter_action,
            filter_source=filter_source,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             namespace: str,
             default_action: Optional[str] = None,
             filter_action: Optional[str] = None,
             filter_source: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("namespace", namespace)
        if default_action is not None:
            _setter("default_action", default_action)
        if filter_action is not None:
            _setter("filter_action", filter_action)
        if filter_source is not None:
            _setter("filter_source", filter_source)

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See `services` field description below for additional information.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[str]:
        """
        Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter(name="filterAction")
    def filter_action(self) -> Optional[str]:
        """
        Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
        """
        return pulumi.get(self, "filter_action")

    @property
    @pulumi.getter(name="filterSource")
    def filter_source(self) -> Optional[str]:
        """
        Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
        """
        return pulumi.get(self, "filter_source")


