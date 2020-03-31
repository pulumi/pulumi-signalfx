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
    api_token: pulumi.Output[str]
    """
    The API token for the user email
    """
    assignee_display_name: pulumi.Output[str]
    """
    Jira display name for the assignee.
    """
    assignee_name: pulumi.Output[str]
    """
    Jira user name for the assignee.
    """
    auth_method: pulumi.Output[str]
    """
    Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
    """
    base_url: pulumi.Output[str]
    """
    Base URL of the Jira instance that's integrated with SignalFx.
    """
    enabled: pulumi.Output[bool]
    """
    Whether the integration is enabled.
    """
    issue_type: pulumi.Output[str]
    """
    Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
    """
    name: pulumi.Output[str]
    """
    Name of the integration.
    """
    password: pulumi.Output[str]
    """
    Password used to authenticate the Jira integration.
    """
    project_key: pulumi.Output[str]
    """
    Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
    """
    user_email: pulumi.Output[str]
    """
    Email address used to authenticate the Jira integration.
    """
    username: pulumi.Output[str]
    """
    User name used to authenticate the Jira integration.
    """
    def __init__(__self__, resource_name, opts=None, api_token=None, assignee_display_name=None, assignee_name=None, auth_method=None, base_url=None, enabled=None, issue_type=None, name=None, password=None, project_key=None, user_email=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        SignalFx Jira integrations. For help with this integration see [Integration with Jira](https://docs.signalfx.com/en/latest/admin-guide/integrate-notifications.html#integrate-with-jira).

        > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/jira_integration.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_token: The API token for the user email
        :param pulumi.Input[str] assignee_display_name: Jira display name for the assignee.
        :param pulumi.Input[str] assignee_name: Jira user name for the assignee.
        :param pulumi.Input[str] auth_method: Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        :param pulumi.Input[str] base_url: Base URL of the Jira instance that's integrated with SignalFx.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] issue_type: Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] password: Password used to authenticate the Jira integration.
        :param pulumi.Input[str] project_key: Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        :param pulumi.Input[str] user_email: Email address used to authenticate the Jira integration.
        :param pulumi.Input[str] username: User name used to authenticate the Jira integration.
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

            __props__['api_token'] = api_token
            __props__['assignee_display_name'] = assignee_display_name
            if assignee_name is None:
                raise TypeError("Missing required property 'assignee_name'")
            __props__['assignee_name'] = assignee_name
            if auth_method is None:
                raise TypeError("Missing required property 'auth_method'")
            __props__['auth_method'] = auth_method
            if base_url is None:
                raise TypeError("Missing required property 'base_url'")
            __props__['base_url'] = base_url
            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            if issue_type is None:
                raise TypeError("Missing required property 'issue_type'")
            __props__['issue_type'] = issue_type
            __props__['name'] = name
            __props__['password'] = password
            if project_key is None:
                raise TypeError("Missing required property 'project_key'")
            __props__['project_key'] = project_key
            __props__['user_email'] = user_email
            __props__['username'] = username
        super(Integration, __self__).__init__(
            'signalfx:jira/integration:Integration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_token=None, assignee_display_name=None, assignee_name=None, auth_method=None, base_url=None, enabled=None, issue_type=None, name=None, password=None, project_key=None, user_email=None, username=None):
        """
        Get an existing Integration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_token: The API token for the user email
        :param pulumi.Input[str] assignee_display_name: Jira display name for the assignee.
        :param pulumi.Input[str] assignee_name: Jira user name for the assignee.
        :param pulumi.Input[str] auth_method: Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        :param pulumi.Input[str] base_url: Base URL of the Jira instance that's integrated with SignalFx.
        :param pulumi.Input[bool] enabled: Whether the integration is enabled.
        :param pulumi.Input[str] issue_type: Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        :param pulumi.Input[str] name: Name of the integration.
        :param pulumi.Input[str] password: Password used to authenticate the Jira integration.
        :param pulumi.Input[str] project_key: Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        :param pulumi.Input[str] user_email: Email address used to authenticate the Jira integration.
        :param pulumi.Input[str] username: User name used to authenticate the Jira integration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_token"] = api_token
        __props__["assignee_display_name"] = assignee_display_name
        __props__["assignee_name"] = assignee_name
        __props__["auth_method"] = auth_method
        __props__["base_url"] = base_url
        __props__["enabled"] = enabled
        __props__["issue_type"] = issue_type
        __props__["name"] = name
        __props__["password"] = password
        __props__["project_key"] = project_key
        __props__["user_email"] = user_email
        __props__["username"] = username
        return Integration(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
