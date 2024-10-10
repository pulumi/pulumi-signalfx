# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'ViewColumnArgs',
    'ViewColumnArgsDict',
    'ViewSortOptionArgs',
    'ViewSortOptionArgsDict',
]

MYPY = False

if not MYPY:
    class ViewColumnArgsDict(TypedDict):
        name: pulumi.Input[str]
        """
        Name of the log view.
        """
elif False:
    ViewColumnArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ViewColumnArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: Name of the log view.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the log view.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


if not MYPY:
    class ViewSortOptionArgsDict(TypedDict):
        descending: pulumi.Input[bool]
        """
        Name of the column
        """
        field: pulumi.Input[str]
        """
        Name of the column
        """
elif False:
    ViewSortOptionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ViewSortOptionArgs:
    def __init__(__self__, *,
                 descending: pulumi.Input[bool],
                 field: pulumi.Input[str]):
        """
        :param pulumi.Input[bool] descending: Name of the column
        :param pulumi.Input[str] field: Name of the column
        """
        pulumi.set(__self__, "descending", descending)
        pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def descending(self) -> pulumi.Input[bool]:
        """
        Name of the column
        """
        return pulumi.get(self, "descending")

    @descending.setter
    def descending(self, value: pulumi.Input[bool]):
        pulumi.set(self, "descending", value)

    @property
    @pulumi.getter
    def field(self) -> pulumi.Input[str]:
        """
        Name of the column
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: pulumi.Input[str]):
        pulumi.set(self, "field", value)


