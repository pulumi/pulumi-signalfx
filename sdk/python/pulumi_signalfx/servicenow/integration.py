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
                 instance_name: pulumi.Input[str],
                 issue_type: pulumi.Input[str],
                 password: pulumi.Input[str],
                 username: pulumi.Input[str],
                 alert_resolved_payload_template: Optional[pulumi.Input[str]] = None,
                 alert_triggered_payload_template: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Integration resource.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled
        :param pulumi.Input[str] instance_name: Name of the ServiceNow instance, for example `myInstances.service-now.com`.
        :param pulumi.Input[str] issue_type: The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        :param pulumi.Input[str] password: Password used to authenticate the ServiceNow integration.
        :param pulumi.Input[str] username: User name used to authenticate the ServiceNow integration.
        :param pulumi.Input[str] alert_resolved_payload_template: A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in
               ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
               ServiceNow. See API reference for details.
        :param pulumi.Input[str] alert_triggered_payload_template: A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification
               to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
               ServiceNow. See API reference for details.
        :param pulumi.Input[str] name: Name of the integration
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "instance_name", instance_name)
        pulumi.set(__self__, "issue_type", issue_type)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)
        if alert_resolved_payload_template is not None:
            pulumi.set(__self__, "alert_resolved_payload_template", alert_resolved_payload_template)
        if alert_triggered_payload_template is not None:
            pulumi.set(__self__, "alert_triggered_payload_template", alert_triggered_payload_template)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Whether the integration is enabled
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Input[str]:
        """
        Name of the ServiceNow instance, for example `myInstances.service-now.com`.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="issueType")
    def issue_type(self) -> pulumi.Input[str]:
        """
        The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        """
        return pulumi.get(self, "issue_type")

    @issue_type.setter
    def issue_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "issue_type", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Password used to authenticate the ServiceNow integration.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        User name used to authenticate the ServiceNow integration.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="alertResolvedPayloadTemplate")
    def alert_resolved_payload_template(self) -> Optional[pulumi.Input[str]]:
        """
        A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in
        ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
        ServiceNow. See API reference for details.
        """
        return pulumi.get(self, "alert_resolved_payload_template")

    @alert_resolved_payload_template.setter
    def alert_resolved_payload_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_resolved_payload_template", value)

    @property
    @pulumi.getter(name="alertTriggeredPayloadTemplate")
    def alert_triggered_payload_template(self) -> Optional[pulumi.Input[str]]:
        """
        A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification
        to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
        ServiceNow. See API reference for details.
        """
        return pulumi.get(self, "alert_triggered_payload_template")

    @alert_triggered_payload_template.setter
    def alert_triggered_payload_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_triggered_payload_template", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the integration
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _IntegrationState:
    def __init__(__self__, *,
                 alert_resolved_payload_template: Optional[pulumi.Input[str]] = None,
                 alert_triggered_payload_template: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 issue_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Integration resources.
        :param pulumi.Input[str] alert_resolved_payload_template: A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in
               ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
               ServiceNow. See API reference for details.
        :param pulumi.Input[str] alert_triggered_payload_template: A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification
               to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
               ServiceNow. See API reference for details.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled
        :param pulumi.Input[str] instance_name: Name of the ServiceNow instance, for example `myInstances.service-now.com`.
        :param pulumi.Input[str] issue_type: The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        :param pulumi.Input[str] name: Name of the integration
        :param pulumi.Input[str] password: Password used to authenticate the ServiceNow integration.
        :param pulumi.Input[str] username: User name used to authenticate the ServiceNow integration.
        """
        if alert_resolved_payload_template is not None:
            pulumi.set(__self__, "alert_resolved_payload_template", alert_resolved_payload_template)
        if alert_triggered_payload_template is not None:
            pulumi.set(__self__, "alert_triggered_payload_template", alert_triggered_payload_template)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if issue_type is not None:
            pulumi.set(__self__, "issue_type", issue_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="alertResolvedPayloadTemplate")
    def alert_resolved_payload_template(self) -> Optional[pulumi.Input[str]]:
        """
        A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in
        ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
        ServiceNow. See API reference for details.
        """
        return pulumi.get(self, "alert_resolved_payload_template")

    @alert_resolved_payload_template.setter
    def alert_resolved_payload_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_resolved_payload_template", value)

    @property
    @pulumi.getter(name="alertTriggeredPayloadTemplate")
    def alert_triggered_payload_template(self) -> Optional[pulumi.Input[str]]:
        """
        A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification
        to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
        ServiceNow. See API reference for details.
        """
        return pulumi.get(self, "alert_triggered_payload_template")

    @alert_triggered_payload_template.setter
    def alert_triggered_payload_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_triggered_payload_template", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the integration is enabled
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the ServiceNow instance, for example `myInstances.service-now.com`.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter(name="issueType")
    def issue_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        """
        return pulumi.get(self, "issue_type")

    @issue_type.setter
    def issue_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issue_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the integration
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Password used to authenticate the ServiceNow integration.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        User name used to authenticate the ServiceNow integration.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Integration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_resolved_payload_template: Optional[pulumi.Input[str]] = None,
                 alert_triggered_payload_template: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 issue_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ServiceNow integrations. For help with this integration see [Integration with ServiceNow](https://docs.splunk.com/observability/en/admin/notif-services/servicenow.html).

        > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

        ## Example

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        service_now_myteam = signalfx.servicenow.Integration("serviceNowMyteam",
            alert_resolved_payload_template="{\\"close_notes\\": \\"{{{messageTitle}}} (customized close msg)\\"}",
            alert_triggered_payload_template="{\\"short_description\\": \\"{{{messageTitle}}} (customized)\\"}",
            enabled=False,
            instance_name="myinst.service-now.com",
            issue_type="Incident",
            password="youd0ntsee1t",
            username="thisis_me")
        ```

        ## Arguments

        * `name` - (Required) Name of the integration.
        * `enabled` - (Required) Whether the integration is enabled.
        * `username` - (Required) User name used to authenticate the ServiceNow integration.
        * `password` - (Required) Password used to authenticate the ServiceNow integration.
        * `instance_name` - (Required) Name of the ServiceNow instance, for example `myinst.service-now.com`.
        * `issue_type` - (Required) The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        * `alert_triggered_payload_template` - (Optional) A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
        * `alert_resolved_payload_template` - (Optional) A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.

        ## Attributes

        In a addition to all arguments above, the following attributes are exported:

        * `id` - The ID of the integration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_resolved_payload_template: A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in
               ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
               ServiceNow. See API reference for details.
        :param pulumi.Input[str] alert_triggered_payload_template: A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification
               to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
               ServiceNow. See API reference for details.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled
        :param pulumi.Input[str] instance_name: Name of the ServiceNow instance, for example `myInstances.service-now.com`.
        :param pulumi.Input[str] issue_type: The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        :param pulumi.Input[str] name: Name of the integration
        :param pulumi.Input[str] password: Password used to authenticate the ServiceNow integration.
        :param pulumi.Input[str] username: User name used to authenticate the ServiceNow integration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ServiceNow integrations. For help with this integration see [Integration with ServiceNow](https://docs.splunk.com/observability/en/admin/notif-services/servicenow.html).

        > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

        ## Example

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        service_now_myteam = signalfx.servicenow.Integration("serviceNowMyteam",
            alert_resolved_payload_template="{\\"close_notes\\": \\"{{{messageTitle}}} (customized close msg)\\"}",
            alert_triggered_payload_template="{\\"short_description\\": \\"{{{messageTitle}}} (customized)\\"}",
            enabled=False,
            instance_name="myinst.service-now.com",
            issue_type="Incident",
            password="youd0ntsee1t",
            username="thisis_me")
        ```

        ## Arguments

        * `name` - (Required) Name of the integration.
        * `enabled` - (Required) Whether the integration is enabled.
        * `username` - (Required) User name used to authenticate the ServiceNow integration.
        * `password` - (Required) Password used to authenticate the ServiceNow integration.
        * `instance_name` - (Required) Name of the ServiceNow instance, for example `myinst.service-now.com`.
        * `issue_type` - (Required) The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        * `alert_triggered_payload_template` - (Optional) A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
        * `alert_resolved_payload_template` - (Optional) A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.

        ## Attributes

        In a addition to all arguments above, the following attributes are exported:

        * `id` - The ID of the integration.

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
                 alert_resolved_payload_template: Optional[pulumi.Input[str]] = None,
                 alert_triggered_payload_template: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 issue_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntegrationArgs.__new__(IntegrationArgs)

            __props__.__dict__["alert_resolved_payload_template"] = alert_resolved_payload_template
            __props__.__dict__["alert_triggered_payload_template"] = alert_triggered_payload_template
            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if instance_name is None and not opts.urn:
                raise TypeError("Missing required property 'instance_name'")
            __props__.__dict__["instance_name"] = instance_name
            if issue_type is None and not opts.urn:
                raise TypeError("Missing required property 'issue_type'")
            __props__.__dict__["issue_type"] = issue_type
            __props__.__dict__["name"] = name
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Integration, __self__).__init__(
            'signalfx:servicenow/integration:Integration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_resolved_payload_template: Optional[pulumi.Input[str]] = None,
            alert_triggered_payload_template: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            issue_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'Integration':
        """
        Get an existing Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_resolved_payload_template: A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in
               ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
               ServiceNow. See API reference for details.
        :param pulumi.Input[str] alert_triggered_payload_template: A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification
               to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
               ServiceNow. See API reference for details.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled
        :param pulumi.Input[str] instance_name: Name of the ServiceNow instance, for example `myInstances.service-now.com`.
        :param pulumi.Input[str] issue_type: The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        :param pulumi.Input[str] name: Name of the integration
        :param pulumi.Input[str] password: Password used to authenticate the ServiceNow integration.
        :param pulumi.Input[str] username: User name used to authenticate the ServiceNow integration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IntegrationState.__new__(_IntegrationState)

        __props__.__dict__["alert_resolved_payload_template"] = alert_resolved_payload_template
        __props__.__dict__["alert_triggered_payload_template"] = alert_triggered_payload_template
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["issue_type"] = issue_type
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["username"] = username
        return Integration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertResolvedPayloadTemplate")
    def alert_resolved_payload_template(self) -> pulumi.Output[Optional[str]]:
        """
        A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in
        ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
        ServiceNow. See API reference for details.
        """
        return pulumi.get(self, "alert_resolved_payload_template")

    @property
    @pulumi.getter(name="alertTriggeredPayloadTemplate")
    def alert_triggered_payload_template(self) -> pulumi.Output[Optional[str]]:
        """
        A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification
        to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in
        ServiceNow. See API reference for details.
        """
        return pulumi.get(self, "alert_triggered_payload_template")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Whether the integration is enabled
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        Name of the ServiceNow instance, for example `myInstances.service-now.com`.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="issueType")
    def issue_type(self) -> pulumi.Output[str]:
        """
        The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
        """
        return pulumi.get(self, "issue_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the integration
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password used to authenticate the ServiceNow integration.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        User name used to authenticate the ServiceNow integration.
        """
        return pulumi.get(self, "username")

