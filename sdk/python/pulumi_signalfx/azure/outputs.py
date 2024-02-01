# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'IntegrationCustomNamespacesPerService',
    'IntegrationResourceFilterRule',
]

@pulumi.output_type
class IntegrationCustomNamespacesPerService(dict):
    def __init__(__self__, *,
                 namespaces: Sequence[str],
                 service: str):
        """
        :param Sequence[str] namespaces: The namespaces to sync
        :param str service: The name of the service
        """
        pulumi.set(__self__, "namespaces", namespaces)
        pulumi.set(__self__, "service", service)

    @property
    @pulumi.getter
    def namespaces(self) -> Sequence[str]:
        """
        The namespaces to sync
        """
        return pulumi.get(self, "namespaces")

    @property
    @pulumi.getter
    def service(self) -> str:
        """
        The name of the service
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
                 filter_source: str):
        pulumi.set(__self__, "filter_source", filter_source)

    @property
    @pulumi.getter(name="filterSource")
    def filter_source(self) -> str:
        return pulumi.get(self, "filter_source")


