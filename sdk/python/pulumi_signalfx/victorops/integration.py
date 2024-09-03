# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['IntegrationArgs', 'Integration']

@pulumi.input_type
class IntegrationArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 name: Optional[pulumi.Input[str]] = None,
                 post_url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Integration resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] post_url: Splunk On-Call REST API URL.
        """
        pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if post_url is not None:
            pulumi.set(__self__, "post_url", post_url)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Whether the integration is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the integration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="postUrl")
    def post_url(self) -> Optional[pulumi.Input[str]]:
        """
        Splunk On-Call REST API URL.
        """
        return pulumi.get(self, "post_url")

    @post_url.setter
    def post_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "post_url", value)


@pulumi.input_type
class _IntegrationState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 post_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Integration resources.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] post_url: Splunk On-Call REST API URL.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if post_url is not None:
            pulumi.set(__self__, "post_url", post_url)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the integration is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the integration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="postUrl")
    def post_url(self) -> Optional[pulumi.Input[str]]:
        """
        Splunk On-Call REST API URL.
        """
        return pulumi.get(self, "post_url")

    @post_url.setter
    def post_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "post_url", value)


class Integration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 post_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Splunk On-Call integrations.

        > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

        ## Example

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        vioctor_ops_myteam = signalfx.victorops.Integration("vioctor_ops_myteam",
            name="Splunk On-Call - My Team",
            enabled=True,
            post_url="https://alert.victorops.com/integrations/generic/1234/alert/$key/$routing_key")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] post_url: Splunk On-Call REST API URL.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Splunk On-Call integrations.

        > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

        ## Example

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        vioctor_ops_myteam = signalfx.victorops.Integration("vioctor_ops_myteam",
            name="Splunk On-Call - My Team",
            enabled=True,
            post_url="https://alert.victorops.com/integrations/generic/1234/alert/$key/$routing_key")
        ```

        :param str resource_name: The name of the resource.
        :param IntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 post_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntegrationArgs.__new__(IntegrationArgs)

            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["post_url"] = post_url
        super(Integration, __self__).__init__(
            'signalfx:victorops/integration:Integration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            post_url: Optional[pulumi.Input[str]] = None) -> 'Integration':
        """
        Get an existing Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] post_url: Splunk On-Call REST API URL.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IntegrationState.__new__(_IntegrationState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["post_url"] = post_url
        return Integration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether the integration is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the integration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="postUrl")
    def post_url(self) -> pulumi.Output[Optional[str]]:
        """
        Splunk On-Call REST API URL.
        """
        return pulumi.get(self, "post_url")

