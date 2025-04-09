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
    'IntegrationCustomNamespacesPerService',
    'IntegrationResourceFilterRule',
]

@pulumi.output_type
class IntegrationCustomNamespacesPerService(dict):
    def __init__(__self__, *,
                 namespaces: Sequence[builtins.str],
                 service: builtins.str):
        """
        :param Sequence[builtins.str] namespaces: The additional namespaces.
        :param builtins.str service: The name of the service.
        """
        pulumi.set(__self__, "namespaces", namespaces)
        pulumi.set(__self__, "service", service)

    @property
    @pulumi.getter
    def namespaces(self) -> Sequence[builtins.str]:
        """
        The additional namespaces.
        """
        return pulumi.get(self, "namespaces")

    @property
    @pulumi.getter
    def service(self) -> builtins.str:
        """
        The name of the service.
        """
        return pulumi.get(self, "service")


@pulumi.output_type
class IntegrationResourceFilterRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "filterSource":
            suggest = "filter_source"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IntegrationResourceFilterRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IntegrationResourceFilterRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IntegrationResourceFilterRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 filter_source: builtins.str):
        """
        :param builtins.str filter_source: Expression that selects the data that Splunk Observability Cloud should sync for the resource associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function. The source of each filter rule must be in the form filter('key', 'value'). You can join multiple filter statements using the and and or operators. Referenced keys are limited to tags and must start with the azure_tag_ prefix.
        """
        pulumi.set(__self__, "filter_source", filter_source)

    @property
    @pulumi.getter(name="filterSource")
    def filter_source(self) -> builtins.str:
        """
        Expression that selects the data that Splunk Observability Cloud should sync for the resource associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function. The source of each filter rule must be in the form filter('key', 'value'). You can join multiple filter statements using the and and or operators. Referenced keys are limited to tags and must start with the azure_tag_ prefix.
        """
        return pulumi.get(self, "filter_source")


