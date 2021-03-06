# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 api_url: Optional[pulumi.Input[str]] = None,
                 auth_token: Optional[pulumi.Input[str]] = None,
                 custom_app_url: Optional[pulumi.Input[str]] = None,
                 timeout_seconds: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] api_url: API URL for your SignalFx org, may include a realm
        :param pulumi.Input[str] auth_token: SignalFx auth token
        :param pulumi.Input[str] custom_app_url: Application URL for your SignalFx org, often customzied for organizations using SSO
        :param pulumi.Input[int] timeout_seconds: Timeout duration for a single HTTP call in seconds. Defaults to 120
        """
        if api_url is not None:
            pulumi.set(__self__, "api_url", api_url)
        if auth_token is not None:
            pulumi.set(__self__, "auth_token", auth_token)
        if custom_app_url is not None:
            pulumi.set(__self__, "custom_app_url", custom_app_url)
        if timeout_seconds is not None:
            pulumi.set(__self__, "timeout_seconds", timeout_seconds)

    @property
    @pulumi.getter(name="apiUrl")
    def api_url(self) -> Optional[pulumi.Input[str]]:
        """
        API URL for your SignalFx org, may include a realm
        """
        return pulumi.get(self, "api_url")

    @api_url.setter
    def api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_url", value)

    @property
    @pulumi.getter(name="authToken")
    def auth_token(self) -> Optional[pulumi.Input[str]]:
        """
        SignalFx auth token
        """
        return pulumi.get(self, "auth_token")

    @auth_token.setter
    def auth_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_token", value)

    @property
    @pulumi.getter(name="customAppUrl")
    def custom_app_url(self) -> Optional[pulumi.Input[str]]:
        """
        Application URL for your SignalFx org, often customzied for organizations using SSO
        """
        return pulumi.get(self, "custom_app_url")

    @custom_app_url.setter
    def custom_app_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_app_url", value)

    @property
    @pulumi.getter(name="timeoutSeconds")
    def timeout_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout duration for a single HTTP call in seconds. Defaults to 120
        """
        return pulumi.get(self, "timeout_seconds")

    @timeout_seconds.setter
    def timeout_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_seconds", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_url: Optional[pulumi.Input[str]] = None,
                 auth_token: Optional[pulumi.Input[str]] = None,
                 custom_app_url: Optional[pulumi.Input[str]] = None,
                 timeout_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The provider type for the signalfx package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_url: API URL for your SignalFx org, may include a realm
        :param pulumi.Input[str] auth_token: SignalFx auth token
        :param pulumi.Input[str] custom_app_url: Application URL for your SignalFx org, often customzied for organizations using SSO
        :param pulumi.Input[int] timeout_seconds: Timeout duration for a single HTTP call in seconds. Defaults to 120
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the signalfx package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_url: Optional[pulumi.Input[str]] = None,
                 auth_token: Optional[pulumi.Input[str]] = None,
                 custom_app_url: Optional[pulumi.Input[str]] = None,
                 timeout_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["api_url"] = api_url
            __props__.__dict__["auth_token"] = auth_token
            __props__.__dict__["custom_app_url"] = custom_app_url
            __props__.__dict__["timeout_seconds"] = pulumi.Output.from_input(timeout_seconds).apply(pulumi.runtime.to_json) if timeout_seconds is not None else None
        super(Provider, __self__).__init__(
            'signalfx',
            resource_name,
            __props__,
            opts)

