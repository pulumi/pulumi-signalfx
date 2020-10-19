# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Integration']


class Integration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 named_token: Optional[pulumi.Input[str]] = None,
                 poll_rate: Optional[pulumi.Input[int]] = None,
                 project_service_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IntegrationProjectServiceKeyArgs']]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 whitelists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        SignalFx GCP Integration

        > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        gcp_myteam = signalfx.gcp.Integration("gcpMyteam",
            enabled=True,
            poll_rate=300000,
            project_service_keys=[
                signalfx.gcp.IntegrationProjectServiceKeyArgs(
                    project_id="gcp_project_id_1",
                    project_key=(lambda path: open(path).read())("/path/to/gcp_credentials_1.json"),
                ),
                signalfx.gcp.IntegrationProjectServiceKeyArgs(
                    project_id="gcp_project_id_2",
                    project_key=(lambda path: open(path).read())("/path/to/gcp_credentials_2.json"),
                ),
            ],
            services=["compute"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] named_token: A named token to use for ingest
        :param pulumi.Input[int] poll_rate: GCP integration poll rate in seconds. Can be set to either 60 or 300 (1 minute or 5 minutes).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IntegrationProjectServiceKeyArgs']]]] project_service_keys: GCP projects to add.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valid values.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] whitelists: Compute Metadata Whitelist
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

            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            __props__['name'] = name
            __props__['named_token'] = named_token
            __props__['poll_rate'] = poll_rate
            __props__['project_service_keys'] = project_service_keys
            __props__['services'] = services
            __props__['whitelists'] = whitelists
        super(Integration, __self__).__init__(
            'signalfx:gcp/integration:Integration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            named_token: Optional[pulumi.Input[str]] = None,
            poll_rate: Optional[pulumi.Input[int]] = None,
            project_service_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IntegrationProjectServiceKeyArgs']]]]] = None,
            services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            whitelists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Integration':
        """
        Get an existing Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] named_token: A named token to use for ingest
        :param pulumi.Input[int] poll_rate: GCP integration poll rate in seconds. Can be set to either 60 or 300 (1 minute or 5 minutes).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IntegrationProjectServiceKeyArgs']]]] project_service_keys: GCP projects to add.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valid values.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] whitelists: Compute Metadata Whitelist
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["enabled"] = enabled
        __props__["name"] = name
        __props__["named_token"] = named_token
        __props__["poll_rate"] = poll_rate
        __props__["project_service_keys"] = project_service_keys
        __props__["services"] = services
        __props__["whitelists"] = whitelists
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
    @pulumi.getter(name="namedToken")
    def named_token(self) -> pulumi.Output[Optional[str]]:
        """
        A named token to use for ingest
        """
        return pulumi.get(self, "named_token")

    @property
    @pulumi.getter(name="pollRate")
    def poll_rate(self) -> pulumi.Output[Optional[int]]:
        """
        GCP integration poll rate in seconds. Can be set to either 60 or 300 (1 minute or 5 minutes).
        """
        return pulumi.get(self, "poll_rate")

    @property
    @pulumi.getter(name="projectServiceKeys")
    def project_service_keys(self) -> pulumi.Output[Optional[Sequence['outputs.IntegrationProjectServiceKey']]]:
        """
        GCP projects to add.
        """
        return pulumi.get(self, "project_service_keys")

    @property
    @pulumi.getter
    def services(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valid values.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def whitelists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Compute Metadata Whitelist
        """
        return pulumi.get(self, "whitelists")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

