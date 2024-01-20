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
    'ViewColumn',
    'ViewSortOption',
]

@pulumi.output_type
class ViewColumn(dict):
    def __init__(__self__, *,
                 name: str):
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


@pulumi.output_type
class ViewSortOption(dict):
    def __init__(__self__, *,
                 descending: bool,
                 field: str):
        pulumi.set(__self__, "descending", descending)
        pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def descending(self) -> bool:
        return pulumi.get(self, "descending")

    @property
    @pulumi.getter
    def field(self) -> str:
        return pulumi.get(self, "field")


