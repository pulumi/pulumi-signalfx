# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetServicesResult',
    'AwaitableGetServicesResult',
    'get_services',
    'get_services_output',
]

@pulumi.output_type
class GetServicesResult:
    """
    A collection of values returned by getServices.
    """
    def __init__(__self__, id=None, services=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def services(self) -> Optional[Sequence['outputs.GetServicesServiceResult']]:
        return pulumi.get(self, "services")


class AwaitableGetServicesResult(GetServicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServicesResult(
            id=self.id,
            services=self.services)


def get_services(services: Optional[Sequence[pulumi.InputType['GetServicesServiceArgs']]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServicesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['services'] = services
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('signalfx:azure/getServices:getServices', __args__, opts=opts, typ=GetServicesResult).value

    return AwaitableGetServicesResult(
        id=__ret__.id,
        services=__ret__.services)


@_utilities.lift_output_func(get_services)
def get_services_output(services: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetServicesServiceArgs']]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServicesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
