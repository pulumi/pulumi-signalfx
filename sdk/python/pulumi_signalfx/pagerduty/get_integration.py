# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

    @_builtins.property
    @pulumi.getter
    def enabled(self) -> _builtins.bool:
        """
        Whether the integration is currently enabled.
        """
        return pulumi.get(self, "enabled")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def name(self) -> _builtins.str:
        """
        This is the configured name of the PagerDuty integration.
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


def get_integration(name: Optional[_builtins.str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIntegrationResult:
    """
    Use this data source to fetch the PagerDuty integration details.

    ```python
    import pulumi
    import pulumi_signalfx as signalfx

    pd_integration = signalfx.pagerduty.get_integration(name="PD-Integration")
    ```


    :param _builtins.str name: This is the configured name of the PagerDuty integration.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('signalfx:pagerduty/getIntegration:getIntegration', __args__, opts=opts, typ=GetIntegrationResult).value

    return AwaitableGetIntegrationResult(
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))
def get_integration_output(name: Optional[pulumi.Input[_builtins.str]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIntegrationResult]:
    """
    Use this data source to fetch the PagerDuty integration details.

    ```python
    import pulumi
    import pulumi_signalfx as signalfx

    pd_integration = signalfx.pagerduty.get_integration(name="PD-Integration")
    ```


    :param _builtins.str name: This is the configured name of the PagerDuty integration.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('signalfx:pagerduty/getIntegration:getIntegration', __args__, opts=opts, typ=GetIntegrationResult)
    return __ret__.apply(lambda __response__: GetIntegrationResult(
        enabled=pulumi.get(__response__, 'enabled'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name')))
