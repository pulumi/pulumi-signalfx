# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'IntegrationProjectServiceKey',
    'GetServicesServiceResult',
]

@pulumi.output_type
class IntegrationProjectServiceKey(dict):
    def __init__(__self__, *,
                 project_id: str,
                 project_key: str):
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "project_key", project_key)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> str:
        return pulumi.get(self, "project_key")

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


