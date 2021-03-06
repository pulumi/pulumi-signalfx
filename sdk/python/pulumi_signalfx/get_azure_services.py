# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetAzureServicesResult',
    'AwaitableGetAzureServicesResult',
    'get_azure_services',
]

warnings.warn("""signalfx.getAzureServices has been deprecated in favor of signalfx.azure.getServices""", DeprecationWarning)

@pulumi.output_type
class GetAzureServicesResult:
    """
    A collection of values returned by getAzureServices.
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
    def services(self) -> Optional[Sequence['outputs.GetAzureServicesServiceResult']]:
        return pulumi.get(self, "services")


class AwaitableGetAzureServicesResult(GetAzureServicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAzureServicesResult(
            id=self.id,
            services=self.services)


def get_azure_services(services: Optional[Sequence[pulumi.InputType['GetAzureServicesServiceArgs']]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAzureServicesResult:
    """
    Use this data source to access information about an existing resource.
    """
    pulumi.log.warn("""get_azure_services is deprecated: signalfx.getAzureServices has been deprecated in favor of signalfx.azure.getServices""")
    __args__ = dict()
    __args__['services'] = services
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('signalfx:index/getAzureServices:getAzureServices', __args__, opts=opts, typ=GetAzureServicesResult).value

    return AwaitableGetAzureServicesResult(
        id=__ret__.id,
        services=__ret__.services)
