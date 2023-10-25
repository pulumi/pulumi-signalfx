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
    'ViewColumn',
    'ViewSortOption',
]

@pulumi.output_type
class ViewColumn(dict):
    def __init__(__self__, *,
                 name: str):
        """
        :param str name: Name of the log view.
        """
        ViewColumn._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if name is None:
            raise TypeError("Missing 'name' argument")

        _setter("name", name)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the log view.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class ViewSortOption(dict):
    def __init__(__self__, *,
                 descending: bool,
                 field: str):
        ViewSortOption._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            descending=descending,
            field=field,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             descending: Optional[bool] = None,
             field: Optional[str] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if descending is None:
            raise TypeError("Missing 'descending' argument")
        if field is None:
            raise TypeError("Missing 'field' argument")

        _setter("descending", descending)
        _setter("field", field)

    @property
    @pulumi.getter
    def descending(self) -> bool:
        return pulumi.get(self, "descending")

    @property
    @pulumi.getter
    def field(self) -> str:
        return pulumi.get(self, "field")


