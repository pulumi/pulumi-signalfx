# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, api_url=None, auth_token=None, custom_app_url=None, timeout_seconds=None, __props__=None, __name__=None, __opts__=None):
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
        :param pulumi.Input[float] timeout_seconds: Timeout duration for a single HTTP call in seconds. Defaults to 120
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

            __props__['api_url'] = api_url
            if auth_token is None:
                auth_token = utilities.get_env('SFX_AUTH_TOKEN')
            __props__['auth_token'] = auth_token
            __props__['custom_app_url'] = custom_app_url
            __props__['timeout_seconds'] = pulumi.Output.from_input(timeout_seconds).apply(json.dumps) if timeout_seconds is not None else None
        super(Provider, __self__).__init__(
            'signalfx',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

