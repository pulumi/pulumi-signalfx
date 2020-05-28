# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Integration(pulumi.CustomResource):
    app_id: pulumi.Output[str]
    """
    Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
    """
    enabled: pulumi.Output[bool]
    """
    Whether the integration is enabled.
    """
    environment: pulumi.Output[str]
    """
    What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
    """
    name: pulumi.Output[str]
    """
    Name of the integration.
    """
    poll_rate: pulumi.Output[float]
    """
    AWS poll rate (in seconds). One of `60` or `300`.
    """
    secret_key: pulumi.Output[str]
    """
    Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
    """
    services: pulumi.Output[list]
    """
    List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
    """
    subscriptions: pulumi.Output[list]
    """
    List of Azure subscriptions that SignalFx should monitor.
    """
    tenant_id: pulumi.Output[str]
    """
    Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
    """
    def __init__(__self__, resource_name, opts=None, app_id=None, enabled=None, environment=None, name=None, poll_rate=None, secret_key=None, services=None, subscriptions=None, tenant_id=None, __props__=None, __name__=None, __opts__=None):
        """
        SignalFx Azure integrations. For help with this integration see [Monitoring Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure).

        > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        azure_myteam = signalfx.azure.Integration("azureMyteam",
            enabled=True,
            resource=[{
                "signalfxAzureIntegration": [{
                    "azureMyteamXX": [{
                        "app_id": "YYY",
                        "enabled": False,
                        "environment": "azure",
                        "name": "AzureFoo",
                        "poll_rate": 300,
                        "secret_key": "XXX",
                        "services": ["microsoft.sql/servers/elasticpools"],
                        "subscriptions": ["sub-guid-here"],
                        "tenant_id": "ZZZ",
                    }],
                }],
            }])
        ```

        ## Service Names

        > **NOTE** You can use the data source "azure.getServices" to specify all services.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] environment: What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[float] poll_rate: AWS poll rate (in seconds). One of `60` or `300`.
        :param pulumi.Input[str] secret_key: Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
        :param pulumi.Input[list] services: List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
        :param pulumi.Input[list] subscriptions: List of Azure subscriptions that SignalFx should monitor.
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, app_id=None, enabled=None, environment=None, name=None, poll_rate=None, secret_key=None, services=None, subscriptions=None, tenant_id=None):
        """
        Get an existing Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] environment: What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[float] poll_rate: AWS poll rate (in seconds). One of `60` or `300`.
        :param pulumi.Input[str] secret_key: Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
        :param pulumi.Input[list] services: List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
        :param pulumi.Input[list] subscriptions: List of Azure subscriptions that SignalFx should monitor.
        :param pulumi.Input[str] tenant_id: Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_id"] = app_id
        __props__["enabled"] = enabled
        __props__["environment"] = environment
        __props__["name"] = name
        __props__["poll_rate"] = poll_rate
        __props__["secret_key"] = secret_key
        __props__["services"] = services
        __props__["subscriptions"] = subscriptions
        __props__["tenant_id"] = tenant_id
        return Integration(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

