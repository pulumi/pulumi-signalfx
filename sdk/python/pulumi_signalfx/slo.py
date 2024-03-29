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

__all__ = ['SloArgs', 'Slo']

@pulumi.input_type
class SloArgs:
    def __init__(__self__, *,
                 input: pulumi.Input['SloInputArgs'],
                 target: pulumi.Input['SloTargetArgs'],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Slo resource.
        :param pulumi.Input['SloInputArgs'] input: SignalFlow program and arguments text strings that define the streams used as successful event count and total event
               count
        :param pulumi.Input['SloTargetArgs'] target: Define target value of the service level indicator in the appropriate time period.
        :param pulumi.Input[str] type: Type of the SLO. Currently only RequestBased SLO is supported
        :param pulumi.Input[str] description: Description of the SLO
        :param pulumi.Input[str] name: Name of the SLO
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
        SignalFlow program and arguments text strings that define the streams used as successful event count and total event
        count
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
    def type(self) -> pulumi.Input[str]:
        """
        Type of the SLO. Currently only RequestBased SLO is supported
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the SLO
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the SLO
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _SloState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 input: Optional[pulumi.Input['SloInputArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input['SloTargetArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Slo resources.
        :param pulumi.Input[str] description: Description of the SLO
        :param pulumi.Input['SloInputArgs'] input: SignalFlow program and arguments text strings that define the streams used as successful event count and total event
               count
        :param pulumi.Input[str] name: Name of the SLO
        :param pulumi.Input['SloTargetArgs'] target: Define target value of the service level indicator in the appropriate time period.
        :param pulumi.Input[str] type: Type of the SLO. Currently only RequestBased SLO is supported
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
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the SLO
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def input(self) -> Optional[pulumi.Input['SloInputArgs']]:
        """
        SignalFlow program and arguments text strings that define the streams used as successful event count and total event
        count
        """
        return pulumi.get(self, "input")

    @input.setter
    def input(self, value: Optional[pulumi.Input['SloInputArgs']]):
        pulumi.set(self, "input", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the SLO
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
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
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the SLO. Currently only RequestBased SLO is supported
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Slo(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 input: Optional[pulumi.Input[pulumi.InputType['SloInputArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[pulumi.InputType['SloTargetArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Splunk Observability Cloud slo resource. This can be used to create and manage SLOs.

        To learn more about this feature take a look on [documentation for SLO](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/slo-intro.html).

        ## Example

        ## Notification format

        As Splunk Observability Cloud supports different notification mechanisms, use a comma-delimited string to provide inputs. If you want to specify multiple notifications, each must be a member in the list, like so:

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        See [Splunk Observability Cloud Docs](https://dev.splunk.com/observability/reference/api/detectors/latest) for more information.

        Here are some example of how to configure each notification type:

        ### Email

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### Jira

        Note that the `credentialId` is the Splunk-provided ID shown after setting up your Jira integration. See also `jira.Integration`.

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### OpsGenie

        Note that the `credentialId` is the Splunk-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### PagerDuty

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### Slack

        Exclude the `#` on the channel name:

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### Team

        Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### TeamEmail

        Sends an email to every member of a team.

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### Splunk On-Call (formerly VictorOps)

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### Webhooks

        You need to include all the commas even if you only use a credential id.

        You can either configure a Webhook to use an existing integration's credential id:
        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        Or configure one inline:

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ## Arguments

        * `name` - (Required) Name of the SLO. Each SLO name must be unique within an organization.
        * `description` - (Optional) Description of the SLO.
        * `type` - (Required) Type of the SLO. Currently just: `"RequestBased"` is supported.
        * `input` - (Required) Properties to configure an SLO object inputs
          * `program_text` - (Required) SignalFlow program and arguments text strings that define the streams used as successful event count and total event count
          * `good_events_label` - (Required) Label used in `"program_text"` that refers to the data block which contains the stream of successful events
          * `total_events_label` - (Required) Label used in `"program_text"` that refers to the data block which contains the stream of total events
        * `target` - (Required) Define target value of the service level indicator in the appropriate time period.
          * `type` - (Required) SLO target type can be the following type: `"RollingWindow"`
          * `compliance_period` - (Required for `"RollingWindow"` type) Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
          * `slo` - (Required) Target value in the form of a percentage
          * `alert_rule` - (Required) List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
            * `type` - (Required) SLO alert rule can be one of the following types: BREACH, ERROR_BUDGET_LEFT, BURN_RATE. Within an SLO object, you can only specify one SLO alert_rule per type. For example, you can't specify two alert_rule of type BREACH. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
            * `rule` - (Required) Set of rules used for alerting.
                * `severity` - (Required) The severity of the rule, must be one of: `"Critical"`, `"Major"`, `"Minor"`, `"Warning"`, `"Info"`.
                * `description` - (Optional) Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
                * `disabled` - (Optional) When true, notifications and events will not be generated for the detect label. `false` by default.
                * `notifications` - (Optional) List of strings specifying where notifications will be sent when an incident occurs. See [Create SLO](https://dev.splunk.com/observability/reference/api/slo/latest#endpoint-create-new-slo) for more info.
                * `parameterized_body` - (Optional) Custom notification message body when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
                * `parameterized_subject` - (Optional) Custom notification message subject when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
                * `runbook_url` - (Optional) URL of page to consult when an alert is triggered. This can be used with custom notification messages.
                * `tip` - (Optional) Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
                * `parameters` - (Optional) Parameters for the SLO alert rule. Each SLO alert rule type accepts different parameters. If not specified, default parameters are used.
                  * `fire_lasting` - (Optional) Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: `"BREACH"` and `"ERROR_BUDGET_LEFT"` alert rules use the fireLasting parameter. Default: `"5m"`
                  * `percent_of_lasting` - (Optional) Percentage of the `"fire_lasting"` duration that the alert condition is met before the alert is triggered. Note: `"BREACH"` and `"ERROR_BUDGET_LEFT"` alert rules use the `"percent_of_lasting"` parameter. Default: `100`
                  * `percent_error_budget_left` - (Optional) Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: `"ERROR_BUDGET_LEFT"` alert rules use the `"percent_error_budget_left"` parameter. Default: `100`
                  * `short_window_1` - (Optional) Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of `"long_window_1"`. Note: `"BURN_RATE"` alert rules use the `"short_window_1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
                  * `short_window_2` - (Optional) Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of `"long_window_2"`. Note: `"BURN_RATE"` alert rules use the `"short_window_2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
                  * `long_window_1` - (Optional) Long window 1 used in burn rate alert calculation. This value must be longer than `"short_window_1"` and shorter than 90 days. Note: `"BURN_RATE"` alert rules use the `"long_window_1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
                  * `long_window_2` - (Optional) Long window 2 used in burn rate alert calculation. This value must be longer than `"short_window_2"` and shorter than 90 days. Note: `"BURN_RATE"` alert rules use the `"long_window_2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
                  * `burn_rate_threshold_1` - (Optional) Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: `"BURN_RATE"` alert rules use the `"burn_rate_threshold_1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
                  * `burn_rate_threshold_2` - (Optional) Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: `"BURN_RATE"` alert rules use the `"burn_rate_threshold_2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the SLO
        :param pulumi.Input[pulumi.InputType['SloInputArgs']] input: SignalFlow program and arguments text strings that define the streams used as successful event count and total event
               count
        :param pulumi.Input[str] name: Name of the SLO
        :param pulumi.Input[pulumi.InputType['SloTargetArgs']] target: Define target value of the service level indicator in the appropriate time period.
        :param pulumi.Input[str] type: Type of the SLO. Currently only RequestBased SLO is supported
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

        ## Notification format

        As Splunk Observability Cloud supports different notification mechanisms, use a comma-delimited string to provide inputs. If you want to specify multiple notifications, each must be a member in the list, like so:

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        See [Splunk Observability Cloud Docs](https://dev.splunk.com/observability/reference/api/detectors/latest) for more information.

        Here are some example of how to configure each notification type:

        ### Email

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### Jira

        Note that the `credentialId` is the Splunk-provided ID shown after setting up your Jira integration. See also `jira.Integration`.

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### OpsGenie

        Note that the `credentialId` is the Splunk-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### PagerDuty

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### Slack

        Exclude the `#` on the channel name:

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### Team

        Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### TeamEmail

        Sends an email to every member of a team.

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### Splunk On-Call (formerly VictorOps)

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ### Webhooks

        You need to include all the commas even if you only use a credential id.

        You can either configure a Webhook to use an existing integration's credential id:
        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        Or configure one inline:

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        ```
        <!--End PulumiCodeChooser -->

        ## Arguments

        * `name` - (Required) Name of the SLO. Each SLO name must be unique within an organization.
        * `description` - (Optional) Description of the SLO.
        * `type` - (Required) Type of the SLO. Currently just: `"RequestBased"` is supported.
        * `input` - (Required) Properties to configure an SLO object inputs
          * `program_text` - (Required) SignalFlow program and arguments text strings that define the streams used as successful event count and total event count
          * `good_events_label` - (Required) Label used in `"program_text"` that refers to the data block which contains the stream of successful events
          * `total_events_label` - (Required) Label used in `"program_text"` that refers to the data block which contains the stream of total events
        * `target` - (Required) Define target value of the service level indicator in the appropriate time period.
          * `type` - (Required) SLO target type can be the following type: `"RollingWindow"`
          * `compliance_period` - (Required for `"RollingWindow"` type) Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
          * `slo` - (Required) Target value in the form of a percentage
          * `alert_rule` - (Required) List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
            * `type` - (Required) SLO alert rule can be one of the following types: BREACH, ERROR_BUDGET_LEFT, BURN_RATE. Within an SLO object, you can only specify one SLO alert_rule per type. For example, you can't specify two alert_rule of type BREACH. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
            * `rule` - (Required) Set of rules used for alerting.
                * `severity` - (Required) The severity of the rule, must be one of: `"Critical"`, `"Major"`, `"Minor"`, `"Warning"`, `"Info"`.
                * `description` - (Optional) Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
                * `disabled` - (Optional) When true, notifications and events will not be generated for the detect label. `false` by default.
                * `notifications` - (Optional) List of strings specifying where notifications will be sent when an incident occurs. See [Create SLO](https://dev.splunk.com/observability/reference/api/slo/latest#endpoint-create-new-slo) for more info.
                * `parameterized_body` - (Optional) Custom notification message body when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
                * `parameterized_subject` - (Optional) Custom notification message subject when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
                * `runbook_url` - (Optional) URL of page to consult when an alert is triggered. This can be used with custom notification messages.
                * `tip` - (Optional) Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
                * `parameters` - (Optional) Parameters for the SLO alert rule. Each SLO alert rule type accepts different parameters. If not specified, default parameters are used.
                  * `fire_lasting` - (Optional) Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: `"BREACH"` and `"ERROR_BUDGET_LEFT"` alert rules use the fireLasting parameter. Default: `"5m"`
                  * `percent_of_lasting` - (Optional) Percentage of the `"fire_lasting"` duration that the alert condition is met before the alert is triggered. Note: `"BREACH"` and `"ERROR_BUDGET_LEFT"` alert rules use the `"percent_of_lasting"` parameter. Default: `100`
                  * `percent_error_budget_left` - (Optional) Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: `"ERROR_BUDGET_LEFT"` alert rules use the `"percent_error_budget_left"` parameter. Default: `100`
                  * `short_window_1` - (Optional) Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of `"long_window_1"`. Note: `"BURN_RATE"` alert rules use the `"short_window_1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
                  * `short_window_2` - (Optional) Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of `"long_window_2"`. Note: `"BURN_RATE"` alert rules use the `"short_window_2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
                  * `long_window_1` - (Optional) Long window 1 used in burn rate alert calculation. This value must be longer than `"short_window_1"` and shorter than 90 days. Note: `"BURN_RATE"` alert rules use the `"long_window_1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
                  * `long_window_2` - (Optional) Long window 2 used in burn rate alert calculation. This value must be longer than `"short_window_2"` and shorter than 90 days. Note: `"BURN_RATE"` alert rules use the `"long_window_2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
                  * `burn_rate_threshold_1` - (Optional) Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: `"BURN_RATE"` alert rules use the `"burn_rate_threshold_1"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
                  * `burn_rate_threshold_2` - (Optional) Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: `"BURN_RATE"` alert rules use the `"burn_rate_threshold_2"` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.

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
                 description: Optional[pulumi.Input[str]] = None,
                 input: Optional[pulumi.Input[pulumi.InputType['SloInputArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[pulumi.InputType['SloTargetArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
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
            description: Optional[pulumi.Input[str]] = None,
            input: Optional[pulumi.Input[pulumi.InputType['SloInputArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            target: Optional[pulumi.Input[pulumi.InputType['SloTargetArgs']]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Slo':
        """
        Get an existing Slo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the SLO
        :param pulumi.Input[pulumi.InputType['SloInputArgs']] input: SignalFlow program and arguments text strings that define the streams used as successful event count and total event
               count
        :param pulumi.Input[str] name: Name of the SLO
        :param pulumi.Input[pulumi.InputType['SloTargetArgs']] target: Define target value of the service level indicator in the appropriate time period.
        :param pulumi.Input[str] type: Type of the SLO. Currently only RequestBased SLO is supported
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
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the SLO
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def input(self) -> pulumi.Output['outputs.SloInput']:
        """
        SignalFlow program and arguments text strings that define the streams used as successful event count and total event
        count
        """
        return pulumi.get(self, "input")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the SLO
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
    def type(self) -> pulumi.Output[str]:
        """
        Type of the SLO. Currently only RequestBased SLO is supported
        """
        return pulumi.get(self, "type")

