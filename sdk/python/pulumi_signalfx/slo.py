# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
from . import outputs
from ._inputs import *

__all__ = ['SloArgs', 'Slo']

@pulumi.input_type
class SloArgs:
    def __init__(__self__, *,
                 input: pulumi.Input['SloInputArgs'],
                 target: pulumi.Input['SloTargetArgs'],
                 type: pulumi.Input[builtins.str],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Slo resource.
        :param pulumi.Input['SloInputArgs'] input: Properties to configure an SLO object inputs
        :param pulumi.Input['SloTargetArgs'] target: Define target value of the service level indicator in the appropriate time period.
        :param pulumi.Input[builtins.str] type: Type of the SLO. Currently just: `"RequestBased"` is supported.
        :param pulumi.Input[builtins.str] description: Description of the SLO.
        :param pulumi.Input[builtins.str] name: Name of the SLO. Each SLO name must be unique within an organization.
        """
        pulumi.set(__self__, "input", input)
        pulumi.set(__self__, "target", target)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def input(self) -> pulumi.Input['SloInputArgs']:
        """
        Properties to configure an SLO object inputs
        """
        return pulumi.get(self, "input")

    @input.setter
    def input(self, value: pulumi.Input['SloInputArgs']):
        pulumi.set(self, "input", value)

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input['SloTargetArgs']:
        """
        Define target value of the service level indicator in the appropriate time period.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: pulumi.Input['SloTargetArgs']):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[builtins.str]:
        """
        Type of the SLO. Currently just: `"RequestBased"` is supported.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the SLO.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the SLO. Each SLO name must be unique within an organization.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _SloState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 input: Optional[pulumi.Input['SloInputArgs']] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 target: Optional[pulumi.Input['SloTargetArgs']] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Slo resources.
        :param pulumi.Input[builtins.str] description: Description of the SLO.
        :param pulumi.Input['SloInputArgs'] input: Properties to configure an SLO object inputs
        :param pulumi.Input[builtins.str] name: Name of the SLO. Each SLO name must be unique within an organization.
        :param pulumi.Input['SloTargetArgs'] target: Define target value of the service level indicator in the appropriate time period.
        :param pulumi.Input[builtins.str] type: Type of the SLO. Currently just: `"RequestBased"` is supported.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if input is not None:
            pulumi.set(__self__, "input", input)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if target is not None:
            pulumi.set(__self__, "target", target)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description of the SLO.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def input(self) -> Optional[pulumi.Input['SloInputArgs']]:
        """
        Properties to configure an SLO object inputs
        """
        return pulumi.get(self, "input")

    @input.setter
    def input(self, value: Optional[pulumi.Input['SloInputArgs']]):
        pulumi.set(self, "input", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the SLO. Each SLO name must be unique within an organization.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input['SloTargetArgs']]:
        """
        Define target value of the service level indicator in the appropriate time period.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input['SloTargetArgs']]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Type of the SLO. Currently just: `"RequestBased"` is supported.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


class Slo(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 input: Optional[pulumi.Input[Union['SloInputArgs', 'SloInputArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 target: Optional[pulumi.Input[Union['SloTargetArgs', 'SloTargetArgsDict']]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Provides a Splunk Observability Cloud slo resource. This can be used to create and manage SLOs.

        To learn more about this feature take a look on [documentation for SLO](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/slo-intro.html).

        ## Example

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        foo_service_slo = signalfx.Slo("foo_service_slo",
            name="foo service SLO",
            type="RequestBased",
            description="SLO monitoring for foo service",
            input={
                "program_text": \"\"\"G = data('spans.count', filter=filter('sf_error', 'false') and filter('sf_service', 'foo-service'))
        T = data('spans.count', filter=filter('sf_service', 'foo-service'))\"\"\",
                "good_events_label": "G",
                "total_events_label": "T",
            },
            target={
                "type": "RollingWindow",
                "slo": 95,
                "compliance_period": "30d",
                "alert_rules": [{
                    "type": "BREACH",
                    "rules": [{
                        "severity": "Warning",
                        "notifications": ["Email,foo-alerts@bar.com"],
                    }],
                }],
            })
        ```

        ## Notification format

        As Splunk Observability Cloud supports different notification mechanisms, use a comma-delimited string to provide inputs. If you want to specify multiple notifications, each must be a member in the list, like so:

        See [Splunk Observability Cloud Docs](https://dev.splunk.com/observability/reference/api/detectors/latest) for more information.

        Here are some example of how to configure each notification type:

        ### Email

        ### Jira

        Note that the `credentialId` is the Splunk-provided ID shown after setting up your Jira integration. See also `jira.Integration`.

        ### OpsGenie

        Note that the `credentialId` is the Splunk-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.

        ### PagerDuty

        ### Slack

        Exclude the `#` on the channel name:

        ### Team

        Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).

        ### TeamEmail

        Sends an email to every member of a team.

        ### Splunk On-Call (formerly VictorOps)

        ### Webhooks

        You need to include all the commas even if you only use a credential id.

        You can either configure a Webhook to use an existing integration's credential id:

        Or configure one inline:

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Description of the SLO.
        :param pulumi.Input[Union['SloInputArgs', 'SloInputArgsDict']] input: Properties to configure an SLO object inputs
        :param pulumi.Input[builtins.str] name: Name of the SLO. Each SLO name must be unique within an organization.
        :param pulumi.Input[Union['SloTargetArgs', 'SloTargetArgsDict']] target: Define target value of the service level indicator in the appropriate time period.
        :param pulumi.Input[builtins.str] type: Type of the SLO. Currently just: `"RequestBased"` is supported.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SloArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Splunk Observability Cloud slo resource. This can be used to create and manage SLOs.

        To learn more about this feature take a look on [documentation for SLO](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/slo-intro.html).

        ## Example

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        foo_service_slo = signalfx.Slo("foo_service_slo",
            name="foo service SLO",
            type="RequestBased",
            description="SLO monitoring for foo service",
            input={
                "program_text": \"\"\"G = data('spans.count', filter=filter('sf_error', 'false') and filter('sf_service', 'foo-service'))
        T = data('spans.count', filter=filter('sf_service', 'foo-service'))\"\"\",
                "good_events_label": "G",
                "total_events_label": "T",
            },
            target={
                "type": "RollingWindow",
                "slo": 95,
                "compliance_period": "30d",
                "alert_rules": [{
                    "type": "BREACH",
                    "rules": [{
                        "severity": "Warning",
                        "notifications": ["Email,foo-alerts@bar.com"],
                    }],
                }],
            })
        ```

        ## Notification format

        As Splunk Observability Cloud supports different notification mechanisms, use a comma-delimited string to provide inputs. If you want to specify multiple notifications, each must be a member in the list, like so:

        See [Splunk Observability Cloud Docs](https://dev.splunk.com/observability/reference/api/detectors/latest) for more information.

        Here are some example of how to configure each notification type:

        ### Email

        ### Jira

        Note that the `credentialId` is the Splunk-provided ID shown after setting up your Jira integration. See also `jira.Integration`.

        ### OpsGenie

        Note that the `credentialId` is the Splunk-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.

        ### PagerDuty

        ### Slack

        Exclude the `#` on the channel name:

        ### Team

        Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).

        ### TeamEmail

        Sends an email to every member of a team.

        ### Splunk On-Call (formerly VictorOps)

        ### Webhooks

        You need to include all the commas even if you only use a credential id.

        You can either configure a Webhook to use an existing integration's credential id:

        Or configure one inline:

        :param str resource_name: The name of the resource.
        :param SloArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SloArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 input: Optional[pulumi.Input[Union['SloInputArgs', 'SloInputArgsDict']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 target: Optional[pulumi.Input[Union['SloTargetArgs', 'SloTargetArgsDict']]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SloArgs.__new__(SloArgs)

            __props__.__dict__["description"] = description
            if input is None and not opts.urn:
                raise TypeError("Missing required property 'input'")
            __props__.__dict__["input"] = input
            __props__.__dict__["name"] = name
            if target is None and not opts.urn:
                raise TypeError("Missing required property 'target'")
            __props__.__dict__["target"] = target
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(Slo, __self__).__init__(
            'signalfx:index/slo:Slo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            input: Optional[pulumi.Input[Union['SloInputArgs', 'SloInputArgsDict']]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            target: Optional[pulumi.Input[Union['SloTargetArgs', 'SloTargetArgsDict']]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None) -> 'Slo':
        """
        Get an existing Slo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Description of the SLO.
        :param pulumi.Input[Union['SloInputArgs', 'SloInputArgsDict']] input: Properties to configure an SLO object inputs
        :param pulumi.Input[builtins.str] name: Name of the SLO. Each SLO name must be unique within an organization.
        :param pulumi.Input[Union['SloTargetArgs', 'SloTargetArgsDict']] target: Define target value of the service level indicator in the appropriate time period.
        :param pulumi.Input[builtins.str] type: Type of the SLO. Currently just: `"RequestBased"` is supported.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SloState.__new__(_SloState)

        __props__.__dict__["description"] = description
        __props__.__dict__["input"] = input
        __props__.__dict__["name"] = name
        __props__.__dict__["target"] = target
        __props__.__dict__["type"] = type
        return Slo(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description of the SLO.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def input(self) -> pulumi.Output['outputs.SloInput']:
        """
        Properties to configure an SLO object inputs
        """
        return pulumi.get(self, "input")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the SLO. Each SLO name must be unique within an organization.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output['outputs.SloTarget']:
        """
        Define target value of the service level indicator in the appropriate time period.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        Type of the SLO. Currently just: `"RequestBased"` is supported.
        """
        return pulumi.get(self, "type")

