# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['OrgTokenArgs', 'OrgToken']

@pulumi.input_type
class OrgTokenArgs:
    def __init__(__self__, *,
                 auth_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 dpm_limits: Optional[pulumi.Input['OrgTokenDpmLimitsArgs']] = None,
                 host_or_usage_limits: Optional[pulumi.Input['OrgTokenHostOrUsageLimitsArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a OrgToken resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] auth_scopes: Authentication scope, ex: INGEST, API, RUM ... (Optional)
        :param pulumi.Input[str] description: Description of the token.
        :param pulumi.Input[bool] disabled: Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
        :param pulumi.Input['OrgTokenDpmLimitsArgs'] dpm_limits: Specify DPM-based limits for this token.
        :param pulumi.Input['OrgTokenHostOrUsageLimitsArgs'] host_or_usage_limits: Specify Usage-based limits for this token.
        :param pulumi.Input[str] name: Name of the token.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications: List of strings specifying where notifications will be sent when an incident occurs. See
               https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        """
        if auth_scopes is not None:
            pulumi.set(__self__, "auth_scopes", auth_scopes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if dpm_limits is not None:
            pulumi.set(__self__, "dpm_limits", dpm_limits)
        if host_or_usage_limits is not None:
            pulumi.set(__self__, "host_or_usage_limits", host_or_usage_limits)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notifications is not None:
            pulumi.set(__self__, "notifications", notifications)

    @property
    @pulumi.getter(name="authScopes")
    def auth_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Authentication scope, ex: INGEST, API, RUM ... (Optional)
        """
        return pulumi.get(self, "auth_scopes")

    @auth_scopes.setter
    def auth_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "auth_scopes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the token.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="dpmLimits")
    def dpm_limits(self) -> Optional[pulumi.Input['OrgTokenDpmLimitsArgs']]:
        """
        Specify DPM-based limits for this token.
        """
        return pulumi.get(self, "dpm_limits")

    @dpm_limits.setter
    def dpm_limits(self, value: Optional[pulumi.Input['OrgTokenDpmLimitsArgs']]):
        pulumi.set(self, "dpm_limits", value)

    @property
    @pulumi.getter(name="hostOrUsageLimits")
    def host_or_usage_limits(self) -> Optional[pulumi.Input['OrgTokenHostOrUsageLimitsArgs']]:
        """
        Specify Usage-based limits for this token.
        """
        return pulumi.get(self, "host_or_usage_limits")

    @host_or_usage_limits.setter
    def host_or_usage_limits(self, value: Optional[pulumi.Input['OrgTokenHostOrUsageLimitsArgs']]):
        pulumi.set(self, "host_or_usage_limits", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def notifications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of strings specifying where notifications will be sent when an incident occurs. See
        https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        """
        return pulumi.get(self, "notifications")

    @notifications.setter
    def notifications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications", value)


@pulumi.input_type
class _OrgTokenState:
    def __init__(__self__, *,
                 auth_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 dpm_limits: Optional[pulumi.Input['OrgTokenDpmLimitsArgs']] = None,
                 host_or_usage_limits: Optional[pulumi.Input['OrgTokenHostOrUsageLimitsArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrgToken resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] auth_scopes: Authentication scope, ex: INGEST, API, RUM ... (Optional)
        :param pulumi.Input[str] description: Description of the token.
        :param pulumi.Input[bool] disabled: Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
        :param pulumi.Input['OrgTokenDpmLimitsArgs'] dpm_limits: Specify DPM-based limits for this token.
        :param pulumi.Input['OrgTokenHostOrUsageLimitsArgs'] host_or_usage_limits: Specify Usage-based limits for this token.
        :param pulumi.Input[str] name: Name of the token.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications: List of strings specifying where notifications will be sent when an incident occurs. See
               https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        :param pulumi.Input[str] secret: The secret token created by the API. You cannot set this value.
        """
        if auth_scopes is not None:
            pulumi.set(__self__, "auth_scopes", auth_scopes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if dpm_limits is not None:
            pulumi.set(__self__, "dpm_limits", dpm_limits)
        if host_or_usage_limits is not None:
            pulumi.set(__self__, "host_or_usage_limits", host_or_usage_limits)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notifications is not None:
            pulumi.set(__self__, "notifications", notifications)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="authScopes")
    def auth_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Authentication scope, ex: INGEST, API, RUM ... (Optional)
        """
        return pulumi.get(self, "auth_scopes")

    @auth_scopes.setter
    def auth_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "auth_scopes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the token.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="dpmLimits")
    def dpm_limits(self) -> Optional[pulumi.Input['OrgTokenDpmLimitsArgs']]:
        """
        Specify DPM-based limits for this token.
        """
        return pulumi.get(self, "dpm_limits")

    @dpm_limits.setter
    def dpm_limits(self, value: Optional[pulumi.Input['OrgTokenDpmLimitsArgs']]):
        pulumi.set(self, "dpm_limits", value)

    @property
    @pulumi.getter(name="hostOrUsageLimits")
    def host_or_usage_limits(self) -> Optional[pulumi.Input['OrgTokenHostOrUsageLimitsArgs']]:
        """
        Specify Usage-based limits for this token.
        """
        return pulumi.get(self, "host_or_usage_limits")

    @host_or_usage_limits.setter
    def host_or_usage_limits(self, value: Optional[pulumi.Input['OrgTokenHostOrUsageLimitsArgs']]):
        pulumi.set(self, "host_or_usage_limits", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def notifications(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of strings specifying where notifications will be sent when an incident occurs. See
        https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        """
        return pulumi.get(self, "notifications")

    @notifications.setter
    def notifications(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        The secret token created by the API. You cannot set this value.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


class OrgToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 dpm_limits: Optional[pulumi.Input[pulumi.InputType['OrgTokenDpmLimitsArgs']]] = None,
                 host_or_usage_limits: Optional[pulumi.Input[pulumi.InputType['OrgTokenHostOrUsageLimitsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Manage SignalFx org tokens.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        myteamkey0 = signalfx.OrgToken("myteamkey0",
            description="My team's rad key",
            host_or_usage_limits=signalfx.OrgTokenHostOrUsageLimitsArgs(
                container_limit=200,
                container_notification_threshold=180,
                custom_metrics_limit=1000,
                custom_metrics_notification_threshold=900,
                high_res_metrics_limit=1000,
                high_res_metrics_notification_threshold=900,
                host_limit=100,
                host_notification_threshold=90,
            ),
            notifications=["Email,foo-alerts@bar.com"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] auth_scopes: Authentication scope, ex: INGEST, API, RUM ... (Optional)
        :param pulumi.Input[str] description: Description of the token.
        :param pulumi.Input[bool] disabled: Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['OrgTokenDpmLimitsArgs']] dpm_limits: Specify DPM-based limits for this token.
        :param pulumi.Input[pulumi.InputType['OrgTokenHostOrUsageLimitsArgs']] host_or_usage_limits: Specify Usage-based limits for this token.
        :param pulumi.Input[str] name: Name of the token.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications: List of strings specifying where notifications will be sent when an incident occurs. See
               https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OrgTokenArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage SignalFx org tokens.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        myteamkey0 = signalfx.OrgToken("myteamkey0",
            description="My team's rad key",
            host_or_usage_limits=signalfx.OrgTokenHostOrUsageLimitsArgs(
                container_limit=200,
                container_notification_threshold=180,
                custom_metrics_limit=1000,
                custom_metrics_notification_threshold=900,
                high_res_metrics_limit=1000,
                high_res_metrics_notification_threshold=900,
                host_limit=100,
                host_notification_threshold=90,
            ),
            notifications=["Email,foo-alerts@bar.com"])
        ```

        :param str resource_name: The name of the resource.
        :param OrgTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrgTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 dpm_limits: Optional[pulumi.Input[pulumi.InputType['OrgTokenDpmLimitsArgs']]] = None,
                 host_or_usage_limits: Optional[pulumi.Input[pulumi.InputType['OrgTokenHostOrUsageLimitsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrgTokenArgs.__new__(OrgTokenArgs)

            __props__.__dict__["auth_scopes"] = auth_scopes
            __props__.__dict__["description"] = description
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["dpm_limits"] = dpm_limits
            __props__.__dict__["host_or_usage_limits"] = host_or_usage_limits
            __props__.__dict__["name"] = name
            __props__.__dict__["notifications"] = notifications
            __props__.__dict__["secret"] = None
        super(OrgToken, __self__).__init__(
            'signalfx:index/orgToken:OrgToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            dpm_limits: Optional[pulumi.Input[pulumi.InputType['OrgTokenDpmLimitsArgs']]] = None,
            host_or_usage_limits: Optional[pulumi.Input[pulumi.InputType['OrgTokenHostOrUsageLimitsArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notifications: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            secret: Optional[pulumi.Input[str]] = None) -> 'OrgToken':
        """
        Get an existing OrgToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] auth_scopes: Authentication scope, ex: INGEST, API, RUM ... (Optional)
        :param pulumi.Input[str] description: Description of the token.
        :param pulumi.Input[bool] disabled: Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['OrgTokenDpmLimitsArgs']] dpm_limits: Specify DPM-based limits for this token.
        :param pulumi.Input[pulumi.InputType['OrgTokenHostOrUsageLimitsArgs']] host_or_usage_limits: Specify Usage-based limits for this token.
        :param pulumi.Input[str] name: Name of the token.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications: List of strings specifying where notifications will be sent when an incident occurs. See
               https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        :param pulumi.Input[str] secret: The secret token created by the API. You cannot set this value.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrgTokenState.__new__(_OrgTokenState)

        __props__.__dict__["auth_scopes"] = auth_scopes
        __props__.__dict__["description"] = description
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["dpm_limits"] = dpm_limits
        __props__.__dict__["host_or_usage_limits"] = host_or_usage_limits
        __props__.__dict__["name"] = name
        __props__.__dict__["notifications"] = notifications
        __props__.__dict__["secret"] = secret
        return OrgToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authScopes")
    def auth_scopes(self) -> pulumi.Output[Sequence[str]]:
        """
        Authentication scope, ex: INGEST, API, RUM ... (Optional)
        """
        return pulumi.get(self, "auth_scopes")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the token.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag that controls enabling the token. If set to `true`, the token is disabled, and you can't use it for authentication. Defaults to `false`.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="dpmLimits")
    def dpm_limits(self) -> pulumi.Output[Optional['outputs.OrgTokenDpmLimits']]:
        """
        Specify DPM-based limits for this token.
        """
        return pulumi.get(self, "dpm_limits")

    @property
    @pulumi.getter(name="hostOrUsageLimits")
    def host_or_usage_limits(self) -> pulumi.Output[Optional['outputs.OrgTokenHostOrUsageLimits']]:
        """
        Specify Usage-based limits for this token.
        """
        return pulumi.get(self, "host_or_usage_limits")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the token.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notifications(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of strings specifying where notifications will be sent when an incident occurs. See
        https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        """
        The secret token created by the API. You cannot set this value.
        """
        return pulumi.get(self, "secret")

