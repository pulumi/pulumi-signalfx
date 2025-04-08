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
from .. import _utilities

__all__ = ['IntegrationArgs', 'Integration']

@pulumi.input_type
class IntegrationArgs:
    def __init__(__self__, *,
                 assignee_name: pulumi.Input[builtins.str],
                 auth_method: pulumi.Input[builtins.str],
                 base_url: pulumi.Input[builtins.str],
                 enabled: pulumi.Input[builtins.bool],
                 issue_type: pulumi.Input[builtins.str],
                 project_key: pulumi.Input[builtins.str],
                 api_token: Optional[pulumi.Input[builtins.str]] = None,
                 assignee_display_name: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 user_email: Optional[pulumi.Input[builtins.str]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Integration resource.
        :param pulumi.Input[builtins.str] assignee_name: Jira user name for the assignee.
        :param pulumi.Input[builtins.str] auth_method: Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        :param pulumi.Input[builtins.str] base_url: Base URL of the Jira instance that's integrated with SignalFx.
        :param pulumi.Input[builtins.bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[builtins.str] issue_type: Issue type (for example, Story) for tickets that Jira creates for detector notifications. Splunk Observability Cloud validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        :param pulumi.Input[builtins.str] project_key: Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        :param pulumi.Input[builtins.str] api_token: The API token for the user email
        :param pulumi.Input[builtins.str] assignee_display_name: Jira display name for the assignee.
        :param pulumi.Input[builtins.str] name: Name of the integration.
        :param pulumi.Input[builtins.str] password: Password used to authenticate the Jira integration.
        :param pulumi.Input[builtins.str] user_email: Email address used to authenticate the Jira integration.
        :param pulumi.Input[builtins.str] username: User name used to authenticate the Jira integration.
        """
        pulumi.set(__self__, "assignee_name", assignee_name)
        pulumi.set(__self__, "auth_method", auth_method)
        pulumi.set(__self__, "base_url", base_url)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "issue_type", issue_type)
        pulumi.set(__self__, "project_key", project_key)
        if api_token is not None:
            pulumi.set(__self__, "api_token", api_token)
        if assignee_display_name is not None:
            pulumi.set(__self__, "assignee_display_name", assignee_display_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if user_email is not None:
            pulumi.set(__self__, "user_email", user_email)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="assigneeName")
    def assignee_name(self) -> pulumi.Input[builtins.str]:
        """
        Jira user name for the assignee.
        """
        return pulumi.get(self, "assignee_name")

    @assignee_name.setter
    def assignee_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "assignee_name", value)

    @property
    @pulumi.getter(name="authMethod")
    def auth_method(self) -> pulumi.Input[builtins.str]:
        """
        Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        """
        return pulumi.get(self, "auth_method")

    @auth_method.setter
    def auth_method(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "auth_method", value)

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> pulumi.Input[builtins.str]:
        """
        Base URL of the Jira instance that's integrated with SignalFx.
        """
        return pulumi.get(self, "base_url")

    @base_url.setter
    def base_url(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "base_url", value)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[builtins.bool]:
        """
        Whether the integration is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[builtins.bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="issueType")
    def issue_type(self) -> pulumi.Input[builtins.str]:
        """
        Issue type (for example, Story) for tickets that Jira creates for detector notifications. Splunk Observability Cloud validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        """
        return pulumi.get(self, "issue_type")

    @issue_type.setter
    def issue_type(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "issue_type", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> pulumi.Input[builtins.str]:
        """
        Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The API token for the user email
        """
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "api_token", value)

    @property
    @pulumi.getter(name="assigneeDisplayName")
    def assignee_display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Jira display name for the assignee.
        """
        return pulumi.get(self, "assignee_display_name")

    @assignee_display_name.setter
    def assignee_display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "assignee_display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the integration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Password used to authenticate the Jira integration.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="userEmail")
    def user_email(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Email address used to authenticate the Jira integration.
        """
        return pulumi.get(self, "user_email")

    @user_email.setter
    def user_email(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "user_email", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        User name used to authenticate the Jira integration.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _IntegrationState:
    def __init__(__self__, *,
                 api_token: Optional[pulumi.Input[builtins.str]] = None,
                 assignee_display_name: Optional[pulumi.Input[builtins.str]] = None,
                 assignee_name: Optional[pulumi.Input[builtins.str]] = None,
                 auth_method: Optional[pulumi.Input[builtins.str]] = None,
                 base_url: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 issue_type: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 project_key: Optional[pulumi.Input[builtins.str]] = None,
                 user_email: Optional[pulumi.Input[builtins.str]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Integration resources.
        :param pulumi.Input[builtins.str] api_token: The API token for the user email
        :param pulumi.Input[builtins.str] assignee_display_name: Jira display name for the assignee.
        :param pulumi.Input[builtins.str] assignee_name: Jira user name for the assignee.
        :param pulumi.Input[builtins.str] auth_method: Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        :param pulumi.Input[builtins.str] base_url: Base URL of the Jira instance that's integrated with SignalFx.
        :param pulumi.Input[builtins.bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[builtins.str] issue_type: Issue type (for example, Story) for tickets that Jira creates for detector notifications. Splunk Observability Cloud validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        :param pulumi.Input[builtins.str] name: Name of the integration.
        :param pulumi.Input[builtins.str] password: Password used to authenticate the Jira integration.
        :param pulumi.Input[builtins.str] project_key: Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        :param pulumi.Input[builtins.str] user_email: Email address used to authenticate the Jira integration.
        :param pulumi.Input[builtins.str] username: User name used to authenticate the Jira integration.
        """
        if api_token is not None:
            pulumi.set(__self__, "api_token", api_token)
        if assignee_display_name is not None:
            pulumi.set(__self__, "assignee_display_name", assignee_display_name)
        if assignee_name is not None:
            pulumi.set(__self__, "assignee_name", assignee_name)
        if auth_method is not None:
            pulumi.set(__self__, "auth_method", auth_method)
        if base_url is not None:
            pulumi.set(__self__, "base_url", base_url)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if issue_type is not None:
            pulumi.set(__self__, "issue_type", issue_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if project_key is not None:
            pulumi.set(__self__, "project_key", project_key)
        if user_email is not None:
            pulumi.set(__self__, "user_email", user_email)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The API token for the user email
        """
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "api_token", value)

    @property
    @pulumi.getter(name="assigneeDisplayName")
    def assignee_display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Jira display name for the assignee.
        """
        return pulumi.get(self, "assignee_display_name")

    @assignee_display_name.setter
    def assignee_display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "assignee_display_name", value)

    @property
    @pulumi.getter(name="assigneeName")
    def assignee_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Jira user name for the assignee.
        """
        return pulumi.get(self, "assignee_name")

    @assignee_name.setter
    def assignee_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "assignee_name", value)

    @property
    @pulumi.getter(name="authMethod")
    def auth_method(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        """
        return pulumi.get(self, "auth_method")

    @auth_method.setter
    def auth_method(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "auth_method", value)

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Base URL of the Jira instance that's integrated with SignalFx.
        """
        return pulumi.get(self, "base_url")

    @base_url.setter
    def base_url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "base_url", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether the integration is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="issueType")
    def issue_type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Issue type (for example, Story) for tickets that Jira creates for detector notifications. Splunk Observability Cloud validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        """
        return pulumi.get(self, "issue_type")

    @issue_type.setter
    def issue_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "issue_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Name of the integration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Password used to authenticate the Jira integration.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        """
        return pulumi.get(self, "project_key")

    @project_key.setter
    def project_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_key", value)

    @property
    @pulumi.getter(name="userEmail")
    def user_email(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Email address used to authenticate the Jira integration.
        """
        return pulumi.get(self, "user_email")

    @user_email.setter
    def user_email(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "user_email", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        User name used to authenticate the Jira integration.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "username", value)


class Integration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_token: Optional[pulumi.Input[builtins.str]] = None,
                 assignee_display_name: Optional[pulumi.Input[builtins.str]] = None,
                 assignee_name: Optional[pulumi.Input[builtins.str]] = None,
                 auth_method: Optional[pulumi.Input[builtins.str]] = None,
                 base_url: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 issue_type: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 project_key: Optional[pulumi.Input[builtins.str]] = None,
                 user_email: Optional[pulumi.Input[builtins.str]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Splunk Observability Cloud Jira integrations. For help with this integration see [Integration with Jira](https://docs.splunk.com/observability/en/admin/notif-services/jira.html).

        > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

        ## Example

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        jira_myteam_xx = signalfx.jira.Integration("jira_myteamXX",
            name="JiraFoo",
            enabled=False,
            auth_method="UsernameAndPassword",
            username="yoosername",
            password="paasword",
            assignee_name="testytesterson",
            assignee_display_name="Testy Testerson",
            base_url="https://www.example.com",
            issue_type="Story",
            project_key="TEST")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] api_token: The API token for the user email
        :param pulumi.Input[builtins.str] assignee_display_name: Jira display name for the assignee.
        :param pulumi.Input[builtins.str] assignee_name: Jira user name for the assignee.
        :param pulumi.Input[builtins.str] auth_method: Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        :param pulumi.Input[builtins.str] base_url: Base URL of the Jira instance that's integrated with SignalFx.
        :param pulumi.Input[builtins.bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[builtins.str] issue_type: Issue type (for example, Story) for tickets that Jira creates for detector notifications. Splunk Observability Cloud validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        :param pulumi.Input[builtins.str] name: Name of the integration.
        :param pulumi.Input[builtins.str] password: Password used to authenticate the Jira integration.
        :param pulumi.Input[builtins.str] project_key: Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        :param pulumi.Input[builtins.str] user_email: Email address used to authenticate the Jira integration.
        :param pulumi.Input[builtins.str] username: User name used to authenticate the Jira integration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Splunk Observability Cloud Jira integrations. For help with this integration see [Integration with Jira](https://docs.splunk.com/observability/en/admin/notif-services/jira.html).

        > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.

        ## Example

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        jira_myteam_xx = signalfx.jira.Integration("jira_myteamXX",
            name="JiraFoo",
            enabled=False,
            auth_method="UsernameAndPassword",
            username="yoosername",
            password="paasword",
            assignee_name="testytesterson",
            assignee_display_name="Testy Testerson",
            base_url="https://www.example.com",
            issue_type="Story",
            project_key="TEST")
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
                 api_token: Optional[pulumi.Input[builtins.str]] = None,
                 assignee_display_name: Optional[pulumi.Input[builtins.str]] = None,
                 assignee_name: Optional[pulumi.Input[builtins.str]] = None,
                 auth_method: Optional[pulumi.Input[builtins.str]] = None,
                 base_url: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 issue_type: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 password: Optional[pulumi.Input[builtins.str]] = None,
                 project_key: Optional[pulumi.Input[builtins.str]] = None,
                 user_email: Optional[pulumi.Input[builtins.str]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IntegrationArgs.__new__(IntegrationArgs)

            __props__.__dict__["api_token"] = None if api_token is None else pulumi.Output.secret(api_token)
            __props__.__dict__["assignee_display_name"] = assignee_display_name
            if assignee_name is None and not opts.urn:
                raise TypeError("Missing required property 'assignee_name'")
            __props__.__dict__["assignee_name"] = assignee_name
            if auth_method is None and not opts.urn:
                raise TypeError("Missing required property 'auth_method'")
            __props__.__dict__["auth_method"] = auth_method
            if base_url is None and not opts.urn:
                raise TypeError("Missing required property 'base_url'")
            __props__.__dict__["base_url"] = base_url
            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            if issue_type is None and not opts.urn:
                raise TypeError("Missing required property 'issue_type'")
            __props__.__dict__["issue_type"] = issue_type
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            if project_key is None and not opts.urn:
                raise TypeError("Missing required property 'project_key'")
            __props__.__dict__["project_key"] = project_key
            __props__.__dict__["user_email"] = user_email
            __props__.__dict__["username"] = username
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["apiToken", "password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Integration, __self__).__init__(
            'signalfx:jira/integration:Integration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_token: Optional[pulumi.Input[builtins.str]] = None,
            assignee_display_name: Optional[pulumi.Input[builtins.str]] = None,
            assignee_name: Optional[pulumi.Input[builtins.str]] = None,
            auth_method: Optional[pulumi.Input[builtins.str]] = None,
            base_url: Optional[pulumi.Input[builtins.str]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            issue_type: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            password: Optional[pulumi.Input[builtins.str]] = None,
            project_key: Optional[pulumi.Input[builtins.str]] = None,
            user_email: Optional[pulumi.Input[builtins.str]] = None,
            username: Optional[pulumi.Input[builtins.str]] = None) -> 'Integration':
        """
        Get an existing Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] api_token: The API token for the user email
        :param pulumi.Input[builtins.str] assignee_display_name: Jira display name for the assignee.
        :param pulumi.Input[builtins.str] assignee_name: Jira user name for the assignee.
        :param pulumi.Input[builtins.str] auth_method: Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        :param pulumi.Input[builtins.str] base_url: Base URL of the Jira instance that's integrated with SignalFx.
        :param pulumi.Input[builtins.bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[builtins.str] issue_type: Issue type (for example, Story) for tickets that Jira creates for detector notifications. Splunk Observability Cloud validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        :param pulumi.Input[builtins.str] name: Name of the integration.
        :param pulumi.Input[builtins.str] password: Password used to authenticate the Jira integration.
        :param pulumi.Input[builtins.str] project_key: Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        :param pulumi.Input[builtins.str] user_email: Email address used to authenticate the Jira integration.
        :param pulumi.Input[builtins.str] username: User name used to authenticate the Jira integration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IntegrationState.__new__(_IntegrationState)

        __props__.__dict__["api_token"] = api_token
        __props__.__dict__["assignee_display_name"] = assignee_display_name
        __props__.__dict__["assignee_name"] = assignee_name
        __props__.__dict__["auth_method"] = auth_method
        __props__.__dict__["base_url"] = base_url
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["issue_type"] = issue_type
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["project_key"] = project_key
        __props__.__dict__["user_email"] = user_email
        __props__.__dict__["username"] = username
        return Integration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The API token for the user email
        """
        return pulumi.get(self, "api_token")

    @property
    @pulumi.getter(name="assigneeDisplayName")
    def assignee_display_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Jira display name for the assignee.
        """
        return pulumi.get(self, "assignee_display_name")

    @property
    @pulumi.getter(name="assigneeName")
    def assignee_name(self) -> pulumi.Output[builtins.str]:
        """
        Jira user name for the assignee.
        """
        return pulumi.get(self, "assignee_name")

    @property
    @pulumi.getter(name="authMethod")
    def auth_method(self) -> pulumi.Output[builtins.str]:
        """
        Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        """
        return pulumi.get(self, "auth_method")

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> pulumi.Output[builtins.str]:
        """
        Base URL of the Jira instance that's integrated with SignalFx.
        """
        return pulumi.get(self, "base_url")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[builtins.bool]:
        """
        Whether the integration is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="issueType")
    def issue_type(self) -> pulumi.Output[builtins.str]:
        """
        Issue type (for example, Story) for tickets that Jira creates for detector notifications. Splunk Observability Cloud validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        """
        return pulumi.get(self, "issue_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Name of the integration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Password used to authenticate the Jira integration.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="projectKey")
    def project_key(self) -> pulumi.Output[builtins.str]:
        """
        Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        """
        return pulumi.get(self, "project_key")

    @property
    @pulumi.getter(name="userEmail")
    def user_email(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Email address used to authenticate the Jira integration.
        """
        return pulumi.get(self, "user_email")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        User name used to authenticate the Jira integration.
        """
        return pulumi.get(self, "username")

