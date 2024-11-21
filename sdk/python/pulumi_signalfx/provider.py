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
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 api_url: Optional[pulumi.Input[str]] = None,
                 auth_token: Optional[pulumi.Input[str]] = None,
                 custom_app_url: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 retry_max_attempts: Optional[pulumi.Input[int]] = None,
                 retry_wait_max_seconds: Optional[pulumi.Input[int]] = None,
                 retry_wait_min_seconds: Optional[pulumi.Input[int]] = None,
                 timeout_seconds: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] api_url: API URL for your Splunk Observability Cloud org, may include a realm
        :param pulumi.Input[str] auth_token: Splunk Observability Cloud auth token
        :param pulumi.Input[str] custom_app_url: Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
        :param pulumi.Input[str] email: Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
               Password
        :param pulumi.Input[str] organization_id: Required if the user is configured to be part of multiple organizations
        :param pulumi.Input[str] password: Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
               Password
        :param pulumi.Input[int] retry_max_attempts: Max retries for a single HTTP call. Defaults to 4
        :param pulumi.Input[int] retry_wait_max_seconds: Maximum retry wait for a single HTTP call in seconds. Defaults to 30
        :param pulumi.Input[int] retry_wait_min_seconds: Minimum retry wait for a single HTTP call in seconds. Defaults to 1
        :param pulumi.Input[int] timeout_seconds: Timeout duration for a single HTTP call in seconds. Defaults to 120
        """
        if api_url is not None:
            pulumi.set(__self__, "api_url", api_url)
        if auth_token is not None:
            pulumi.set(__self__, "auth_token", auth_token)
        if custom_app_url is not None:
            pulumi.set(__self__, "custom_app_url", custom_app_url)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if retry_max_attempts is not None:
            pulumi.set(__self__, "retry_max_attempts", retry_max_attempts)
        if retry_wait_max_seconds is not None:
            pulumi.set(__self__, "retry_wait_max_seconds", retry_wait_max_seconds)
        if retry_wait_min_seconds is not None:
            pulumi.set(__self__, "retry_wait_min_seconds", retry_wait_min_seconds)
        if timeout_seconds is not None:
            pulumi.set(__self__, "timeout_seconds", timeout_seconds)

    @property
    @pulumi.getter(name="apiUrl")
    def api_url(self) -> Optional[pulumi.Input[str]]:
        """
        API URL for your Splunk Observability Cloud org, may include a realm
        """
        return pulumi.get(self, "api_url")

    @api_url.setter
    def api_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_url", value)

    @property
    @pulumi.getter(name="authToken")
    def auth_token(self) -> Optional[pulumi.Input[str]]:
        """
        Splunk Observability Cloud auth token
        """
        return pulumi.get(self, "auth_token")

    @auth_token.setter
    def auth_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_token", value)

    @property
    @pulumi.getter(name="customAppUrl")
    def custom_app_url(self) -> Optional[pulumi.Input[str]]:
        """
        Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
        """
        return pulumi.get(self, "custom_app_url")

    @custom_app_url.setter
    def custom_app_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_app_url", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
        Password
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required if the user is configured to be part of multiple organizations
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
        Password
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="retryMaxAttempts")
    def retry_max_attempts(self) -> Optional[pulumi.Input[int]]:
        """
        Max retries for a single HTTP call. Defaults to 4
        """
        return pulumi.get(self, "retry_max_attempts")

    @retry_max_attempts.setter
    def retry_max_attempts(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retry_max_attempts", value)

    @property
    @pulumi.getter(name="retryWaitMaxSeconds")
    def retry_wait_max_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum retry wait for a single HTTP call in seconds. Defaults to 30
        """
        return pulumi.get(self, "retry_wait_max_seconds")

    @retry_wait_max_seconds.setter
    def retry_wait_max_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retry_wait_max_seconds", value)

    @property
    @pulumi.getter(name="retryWaitMinSeconds")
    def retry_wait_min_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum retry wait for a single HTTP call in seconds. Defaults to 1
        """
        return pulumi.get(self, "retry_wait_min_seconds")

    @retry_wait_min_seconds.setter
    def retry_wait_min_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retry_wait_min_seconds", value)

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
                 email: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 retry_max_attempts: Optional[pulumi.Input[int]] = None,
                 retry_wait_max_seconds: Optional[pulumi.Input[int]] = None,
                 retry_wait_min_seconds: Optional[pulumi.Input[int]] = None,
                 timeout_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        The provider type for the signalfx package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_url: API URL for your Splunk Observability Cloud org, may include a realm
        :param pulumi.Input[str] auth_token: Splunk Observability Cloud auth token
        :param pulumi.Input[str] custom_app_url: Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
        :param pulumi.Input[str] email: Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
               Password
        :param pulumi.Input[str] organization_id: Required if the user is configured to be part of multiple organizations
        :param pulumi.Input[str] password: Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
               Password
        :param pulumi.Input[int] retry_max_attempts: Max retries for a single HTTP call. Defaults to 4
        :param pulumi.Input[int] retry_wait_max_seconds: Maximum retry wait for a single HTTP call in seconds. Defaults to 30
        :param pulumi.Input[int] retry_wait_min_seconds: Minimum retry wait for a single HTTP call in seconds. Defaults to 1
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
                 email: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 retry_max_attempts: Optional[pulumi.Input[int]] = None,
                 retry_wait_max_seconds: Optional[pulumi.Input[int]] = None,
                 retry_wait_min_seconds: Optional[pulumi.Input[int]] = None,
                 timeout_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["api_url"] = api_url
            __props__.__dict__["auth_token"] = auth_token
            __props__.__dict__["custom_app_url"] = custom_app_url
            __props__.__dict__["email"] = email
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["retry_max_attempts"] = pulumi.Output.from_input(retry_max_attempts).apply(pulumi.runtime.to_json) if retry_max_attempts is not None else None
            __props__.__dict__["retry_wait_max_seconds"] = pulumi.Output.from_input(retry_wait_max_seconds).apply(pulumi.runtime.to_json) if retry_wait_max_seconds is not None else None
            __props__.__dict__["retry_wait_min_seconds"] = pulumi.Output.from_input(retry_wait_min_seconds).apply(pulumi.runtime.to_json) if retry_wait_min_seconds is not None else None
            __props__.__dict__["timeout_seconds"] = pulumi.Output.from_input(timeout_seconds).apply(pulumi.runtime.to_json) if timeout_seconds is not None else None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Provider, __self__).__init__(
            'signalfx',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="apiUrl")
    def api_url(self) -> pulumi.Output[Optional[str]]:
        """
        API URL for your Splunk Observability Cloud org, may include a realm
        """
        return pulumi.get(self, "api_url")

    @property
    @pulumi.getter(name="authToken")
    def auth_token(self) -> pulumi.Output[Optional[str]]:
        """
        Splunk Observability Cloud auth token
        """
        return pulumi.get(self, "auth_token")

    @property
    @pulumi.getter(name="customAppUrl")
    def custom_app_url(self) -> pulumi.Output[Optional[str]]:
        """
        Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
        """
        return pulumi.get(self, "custom_app_url")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[str]]:
        """
        Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
        Password
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[Optional[str]]:
        """
        Required if the user is configured to be part of multiple organizations
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
        Password
        """
        return pulumi.get(self, "password")

