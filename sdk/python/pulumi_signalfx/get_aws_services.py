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
    'GetAwsServicesResult',
    'AwaitableGetAwsServicesResult',
    'get_aws_services',
    'get_aws_services_output',
]

warnings.warn("""signalfx.getAwsServices has been deprecated in favor of signalfx.aws.getServices""", DeprecationWarning)

@pulumi.output_type
class GetAwsServicesResult:
    """
    A collection of values returned by getAwsServices.
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
    def services(self) -> Optional[Sequence['outputs.GetAwsServicesServiceResult']]:
        return pulumi.get(self, "services")


class AwaitableGetAwsServicesResult(GetAwsServicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAwsServicesResult(
            id=self.id,
            services=self.services)


def get_aws_services(services: Optional[Sequence[pulumi.InputType['GetAwsServicesServiceArgs']]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAwsServicesResult:
    """
    Use this data source to access information about an existing resource.
    """
    pulumi.log.warn("""get_aws_services is deprecated: signalfx.getAwsServices has been deprecated in favor of signalfx.aws.getServices""")
    __args__ = dict()
    __args__['services'] = services
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('signalfx:index/getAwsServices:getAwsServices', __args__, opts=opts, typ=GetAwsServicesResult).value

    return AwaitableGetAwsServicesResult(
        id=__ret__.id,
        services=__ret__.services)


@_utilities.lift_output_func(get_aws_services)
def get_aws_services_output(services: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetAwsServicesServiceArgs']]]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAwsServicesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    pulumi.log.warn("""get_aws_services is deprecated: signalfx.getAwsServices has been deprecated in favor of signalfx.aws.getServices""")
    ...
