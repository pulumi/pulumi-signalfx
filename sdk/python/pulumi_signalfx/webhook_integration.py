# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class WebhookIntegration(pulumi.CustomResource):
    enabled: pulumi.Output[bool]
    """
    Whether the integration is enabled.
    """
    headers: pulumi.Output[list]
    """
    A header to send with the request

      * `headerKey` (`str`) - The key of the header to send
      * `headerValue` (`str`) - The value of the header to send
    """
    name: pulumi.Output[str]
    """
    Name of the integration.
    """
    shared_secret: pulumi.Output[str]
    url: pulumi.Output[str]
    """
    The URL to request
    """
    def __init__(__self__, resource_name, opts=None, enabled=None, headers=None, name=None, shared_secret=None, url=None, __props__=None, __name__=None, __opts__=None):
        """
        SignalFx Webhook integration.

        > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[list] headers: A header to send with the request
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] url: The URL to request

        The **headers** object supports the following:

          * `headerKey` (`pulumi.Input[str]`) - The key of the header to send
          * `headerValue` (`pulumi.Input[str]`) - The value of the header to send
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            __props__['headers'] = headers
            __props__['name'] = name
            __props__['shared_secret'] = shared_secret
            __props__['url'] = url
        super(WebhookIntegration, __self__).__init__(
            'signalfx:index/webhookIntegration:WebhookIntegration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, enabled=None, headers=None, name=None, shared_secret=None, url=None):
        """
        Get an existing WebhookIntegration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[list] headers: A header to send with the request
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] url: The URL to request

        The **headers** object supports the following:

          * `headerKey` (`pulumi.Input[str]`) - The key of the header to send
          * `headerValue` (`pulumi.Input[str]`) - The value of the header to send
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["enabled"] = enabled
        __props__["headers"] = headers
        __props__["name"] = name
        __props__["shared_secret"] = shared_secret
        __props__["url"] = url
        return WebhookIntegration(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

