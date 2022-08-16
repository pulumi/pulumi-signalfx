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
    'GetIntegrationResult',
    'AwaitableGetIntegrationResult',
    'get_integration',
    'get_integration_output',
]

@pulumi.output_type
class GetIntegrationResult:
    """
    A collection of values returned by getIntegration.
    """
    def __init__(__self__, enabled=None, id=None, name=None):
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether the integration is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the integration.
        """
        return pulumi.get(self, "name")


class AwaitableGetIntegrationResult(GetIntegrationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIntegrationResult(
            enabled=self.enabled,
            id=self.id,
            name=self.name)


def get_integration(name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntegrationResult:
    """
    Use this data source to get information on an existing PagerDuty integration.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_signalfx as signalfx

    pd_integration = signalfx.pagerduty.get_integration(name="PD-Integration")
    ```


    :param str name: Specify the exact name of the desired PagerDuty integration
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('signalfx:pagerduty/getIntegration:getIntegration', __args__, opts=opts, typ=GetIntegrationResult).value

    return AwaitableGetIntegrationResult(
        enabled=__ret__.enabled,
        id=__ret__.id,
        name=__ret__.name)


@_utilities.lift_output_func(get_integration)
def get_integration_output(name: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIntegrationResult]:
    """
    Use this data source to get information on an existing PagerDuty integration.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_signalfx as signalfx

    pd_integration = signalfx.pagerduty.get_integration(name="PD-Integration")
    ```


    :param str name: Specify the exact name of the desired PagerDuty integration
    """
    ...
