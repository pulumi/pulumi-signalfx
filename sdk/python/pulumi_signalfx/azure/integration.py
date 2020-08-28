# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Integration']


class Integration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 named_token: Optional[pulumi.Input[str]] = None,
                 poll_rate: Optional[pulumi.Input[float]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 subscriptions: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        SignalFx Azure integrations. For help with this integration see [Monitoring Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure).

        > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

        ## Service Names

        > **NOTE** You can use the data source "azure.getServices" to specify all services.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] environment: What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] named_token: A named token to use for ingest
        :param pulumi.Input[float] poll_rate: AWS poll rate (in seconds). One of `60` or `300`.
        :param pulumi.Input[str] secret_key: Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
        :param pulumi.Input[List[pulumi.Input[str]]] services: List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
        :param pulumi.Input[List[pulumi.Input[str]]] subscriptions: List of Azure subscriptions that SignalFx should monitor.
        :param pulumi.Input[str] tenant_id: Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if app_id is None:
                raise TypeError("Missing required property 'app_id'")
            __props__['app_id'] = app_id
            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            __props__['environment'] = environment
            __props__['name'] = name
            __props__['named_token'] = named_token
            __props__['poll_rate'] = poll_rate
            if secret_key is None:
                raise TypeError("Missing required property 'secret_key'")
            __props__['secret_key'] = secret_key
            if services is None:
                raise TypeError("Missing required property 'services'")
            __props__['services'] = services
            if subscriptions is None:
                raise TypeError("Missing required property 'subscriptions'")
            __props__['subscriptions'] = subscriptions
            if tenant_id is None:
                raise TypeError("Missing required property 'tenant_id'")
            __props__['tenant_id'] = tenant_id
        super(Integration, __self__).__init__(
            'signalfx:azure/integration:Integration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            named_token: Optional[pulumi.Input[str]] = None,
            poll_rate: Optional[pulumi.Input[float]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            services: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            subscriptions: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'Integration':
        """
        Get an existing Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] environment: What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] named_token: A named token to use for ingest
        :param pulumi.Input[float] poll_rate: AWS poll rate (in seconds). One of `60` or `300`.
        :param pulumi.Input[str] secret_key: Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
        :param pulumi.Input[List[pulumi.Input[str]]] services: List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
        :param pulumi.Input[List[pulumi.Input[str]]] subscriptions: List of Azure subscriptions that SignalFx should monitor.
        :param pulumi.Input[str] tenant_id: Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["enabled"] = enabled
        __props__["environment"] = environment
        __props__["name"] = name
        __props__["named_token"] = named_token
        __props__["poll_rate"] = poll_rate
        __props__["secret_key"] = secret_key
        __props__["services"] = services
        __props__["subscriptions"] = subscriptions
        __props__["tenant_id"] = tenant_id
        return Integration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether the integration is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Optional[str]]:
        """
        What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the integration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namedToken")
    def named_token(self) -> pulumi.Output[Optional[str]]:
        """
        A named token to use for ingest
        """
        return pulumi.get(self, "named_token")

    @property
    @pulumi.getter(name="pollRate")
    def poll_rate(self) -> pulumi.Output[Optional[float]]:
        """
        AWS poll rate (in seconds). One of `60` or `300`.
        """
        return pulumi.get(self, "poll_rate")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter
    def services(self) -> pulumi.Output[List[str]]:
        """
        List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def subscriptions(self) -> pulumi.Output[List[str]]:
        """
        List of Azure subscriptions that SignalFx should monitor.
        """
        return pulumi.get(self, "subscriptions")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
        """
        return pulumi.get(self, "tenant_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

