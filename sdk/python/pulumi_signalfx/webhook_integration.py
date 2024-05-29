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

__all__ = ['WebhookIntegrationArgs', 'WebhookIntegration']

@pulumi.input_type
class WebhookIntegrationArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input['WebhookIntegrationHeaderArgs']]]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload_template: Optional[pulumi.Input[str]] = None,
                 shared_secret: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebhookIntegration resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[Sequence[pulumi.Input['WebhookIntegrationHeaderArgs']]] headers: A header to send with the request
        :param pulumi.Input[str] method: HTTP method used for the webhook request, such as 'GET', 'POST' and 'PUT'
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] payload_template: Template for the payload to be sent with the webhook request in JSON format
        :param pulumi.Input[str] url: The URL to request
        """
        pulumi.set(__self__, "enabled", enabled)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if method is not None:
            pulumi.set(__self__, "method", method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if payload_template is not None:
            pulumi.set(__self__, "payload_template", payload_template)
        if shared_secret is not None:
            pulumi.set(__self__, "shared_secret", shared_secret)
        if url is not None:
            pulumi.set(__self__, "url", url)

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
    def headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WebhookIntegrationHeaderArgs']]]]:
        """
        A header to send with the request
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WebhookIntegrationHeaderArgs']]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP method used for the webhook request, such as 'GET', 'POST' and 'PUT'
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)

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
    @pulumi.getter(name="payloadTemplate")
    def payload_template(self) -> Optional[pulumi.Input[str]]:
        """
        Template for the payload to be sent with the webhook request in JSON format
        """
        return pulumi.get(self, "payload_template")

    @payload_template.setter
    def payload_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload_template", value)

    @property
    @pulumi.getter(name="sharedSecret")
    def shared_secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "shared_secret")

    @shared_secret.setter
    def shared_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "shared_secret", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL to request
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


@pulumi.input_type
class _WebhookIntegrationState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input['WebhookIntegrationHeaderArgs']]]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload_template: Optional[pulumi.Input[str]] = None,
                 shared_secret: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebhookIntegration resources.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[Sequence[pulumi.Input['WebhookIntegrationHeaderArgs']]] headers: A header to send with the request
        :param pulumi.Input[str] method: HTTP method used for the webhook request, such as 'GET', 'POST' and 'PUT'
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] payload_template: Template for the payload to be sent with the webhook request in JSON format
        :param pulumi.Input[str] url: The URL to request
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if headers is not None:
            pulumi.set(__self__, "headers", headers)
        if method is not None:
            pulumi.set(__self__, "method", method)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if payload_template is not None:
            pulumi.set(__self__, "payload_template", payload_template)
        if shared_secret is not None:
            pulumi.set(__self__, "shared_secret", shared_secret)
        if url is not None:
            pulumi.set(__self__, "url", url)

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
    def headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WebhookIntegrationHeaderArgs']]]]:
        """
        A header to send with the request
        """
        return pulumi.get(self, "headers")

    @headers.setter
    def headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WebhookIntegrationHeaderArgs']]]]):
        pulumi.set(self, "headers", value)

    @property
    @pulumi.getter
    def method(self) -> Optional[pulumi.Input[str]]:
        """
        HTTP method used for the webhook request, such as 'GET', 'POST' and 'PUT'
        """
        return pulumi.get(self, "method")

    @method.setter
    def method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method", value)

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
    @pulumi.getter(name="payloadTemplate")
    def payload_template(self) -> Optional[pulumi.Input[str]]:
        """
        Template for the payload to be sent with the webhook request in JSON format
        """
        return pulumi.get(self, "payload_template")

    @payload_template.setter
    def payload_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "payload_template", value)

    @property
    @pulumi.getter(name="sharedSecret")
    def shared_secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "shared_secret")

    @shared_secret.setter
    def shared_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "shared_secret", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL to request
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class WebhookIntegration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookIntegrationHeaderArgs']]]]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload_template: Optional[pulumi.Input[str]] = None,
                 shared_secret: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Splunk Observability Cloud webhook integration.

        > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

        ## Example

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookIntegrationHeaderArgs']]]] headers: A header to send with the request
        :param pulumi.Input[str] method: HTTP method used for the webhook request, such as 'GET', 'POST' and 'PUT'
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] payload_template: Template for the payload to be sent with the webhook request in JSON format
        :param pulumi.Input[str] url: The URL to request
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WebhookIntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Splunk Observability Cloud webhook integration.

        > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

        ## Example

        :param str resource_name: The name of the resource.
        :param WebhookIntegrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebhookIntegrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 headers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookIntegrationHeaderArgs']]]]] = None,
                 method: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 payload_template: Optional[pulumi.Input[str]] = None,
                 shared_secret: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebhookIntegrationArgs.__new__(WebhookIntegrationArgs)

            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["headers"] = None if headers is None else pulumi.Output.secret(headers)
            __props__.__dict__["method"] = method
            __props__.__dict__["name"] = name
            __props__.__dict__["payload_template"] = payload_template
            __props__.__dict__["shared_secret"] = None if shared_secret is None else pulumi.Output.secret(shared_secret)
            __props__.__dict__["url"] = url
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["headers", "sharedSecret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(WebhookIntegration, __self__).__init__(
            'signalfx:index/webhookIntegration:WebhookIntegration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            headers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookIntegrationHeaderArgs']]]]] = None,
            method: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            payload_template: Optional[pulumi.Input[str]] = None,
            shared_secret: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'WebhookIntegration':
        """
        Get an existing WebhookIntegration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookIntegrationHeaderArgs']]]] headers: A header to send with the request
        :param pulumi.Input[str] method: HTTP method used for the webhook request, such as 'GET', 'POST' and 'PUT'
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] payload_template: Template for the payload to be sent with the webhook request in JSON format
        :param pulumi.Input[str] url: The URL to request
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WebhookIntegrationState.__new__(_WebhookIntegrationState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["headers"] = headers
        __props__.__dict__["method"] = method
        __props__.__dict__["name"] = name
        __props__.__dict__["payload_template"] = payload_template
        __props__.__dict__["shared_secret"] = shared_secret
        __props__.__dict__["url"] = url
        return WebhookIntegration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether the integration is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def headers(self) -> pulumi.Output[Optional[Sequence['outputs.WebhookIntegrationHeader']]]:
        """
        A header to send with the request
        """
        return pulumi.get(self, "headers")

    @property
    @pulumi.getter
    def method(self) -> pulumi.Output[Optional[str]]:
        """
        HTTP method used for the webhook request, such as 'GET', 'POST' and 'PUT'
        """
        return pulumi.get(self, "method")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the integration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="payloadTemplate")
    def payload_template(self) -> pulumi.Output[Optional[str]]:
        """
        Template for the payload to be sent with the webhook request in JSON format
        """
        return pulumi.get(self, "payload_template")

    @property
    @pulumi.getter(name="sharedSecret")
    def shared_secret(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "shared_secret")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        """
        The URL to request
        """
        return pulumi.get(self, "url")

