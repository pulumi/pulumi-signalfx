# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class DashboardGroup(pulumi.CustomResource):
    authorized_writer_teams: pulumi.Output[list]
    """
    Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
    """
    authorized_writer_users: pulumi.Output[list]
    """
    User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
    """
    dashboards: pulumi.Output[list]
    """
    [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.

      * `dashboardId` (`str`) - The dashboard id to mirror
      * `descriptionOverride` (`str`) - The description that will override the original dashboards's description.
      * `filterOverrides` (`list`) - The description that will override the original dashboards's description.
        * `negated` (`bool`) - If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
        * `property` (`str`) - A metric time series dimension or property name.
        * `values` (`list`) - (Optional) List of of strings (which will be treated as an OR filter on the property).

      * `nameOverride` (`str`) - The name that will override the original dashboards's name.
      * `variableOverrides` (`list`)
        * `property` (`str`) - A metric time series dimension or property name.
        * `values` (`list`) - (Optional) List of of strings (which will be treated as an OR filter on the property).
        * `valuesSuggesteds` (`list`) - A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
    """
    description: pulumi.Output[str]
    """
    Description of the dashboard group.
    """
    import_qualifiers: pulumi.Output[list]
    name: pulumi.Output[str]
    """
    Name of the dashboard group.
    """
    teams: pulumi.Output[list]
    """
    Team IDs to associate the dashboard group to.
    """
    def __init__(__self__, resource_name, opts=None, authorized_writer_teams=None, authorized_writer_users=None, dashboards=None, description=None, import_qualifiers=None, name=None, teams=None, __props__=None, __name__=None, __opts__=None):
        """
        In the SignalFx web UI, a [dashboard group](https://developers.signalfx.com/dashboard_groups_reference.html) is a collection of dashboards.

        > **NOTE** Dashboard groups cannot be accessed directly, but just via a dashboard contained in them. This is the reason why make show won't show any of yours dashboard groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        mydashboardgroup0 = signalfx.DashboardGroup("mydashboardgroup0",
            description="Cool dashboard group",
            authorized_writer_teams=[signalfx_team["mycoolteam"]["id"]],
            authorized_writer_users=["abc123"])
        ```
        ### With Mirrored Dashboards

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        mydashboardgroup_withmirrors = signalfx.DashboardGroup("mydashboardgroupWithmirrors",
            description="Cool dashboard group",
            dashboards=[{
                "dashboardId": signalfx_dashboard["gc_dashboard"]["id"],
                "nameOverride": "GC For My Service",
                "descriptionOverride": "Garbage Collection dashboard maintained by JVM team",
                "filterOverrides": [{
                    "property": "service",
                    "values": ["myservice"],
                    "negated": False,
                }],
                "variableOverrides": [{
                    "property": "region",
                    "values": ["us-west1"],
                    "valuesSuggesteds": [
                        "us-west-1",
                        "us-east-1",
                    ],
                }],
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] authorized_writer_teams: Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        :param pulumi.Input[list] authorized_writer_users: User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[list] dashboards: [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        :param pulumi.Input[str] description: Description of the dashboard group.
        :param pulumi.Input[str] name: Name of the dashboard group.
        :param pulumi.Input[list] teams: Team IDs to associate the dashboard group to.

        The **dashboards** object supports the following:

          * `dashboardId` (`pulumi.Input[str]`) - The dashboard id to mirror
          * `descriptionOverride` (`pulumi.Input[str]`) - The description that will override the original dashboards's description.
          * `filterOverrides` (`pulumi.Input[list]`) - The description that will override the original dashboards's description.
            * `negated` (`pulumi.Input[bool]`) - If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
            * `property` (`pulumi.Input[str]`) - A metric time series dimension or property name.
            * `values` (`pulumi.Input[list]`) - (Optional) List of of strings (which will be treated as an OR filter on the property).

          * `nameOverride` (`pulumi.Input[str]`) - The name that will override the original dashboards's name.
          * `variableOverrides` (`pulumi.Input[list]`)
            * `property` (`pulumi.Input[str]`) - A metric time series dimension or property name.
            * `values` (`pulumi.Input[list]`) - (Optional) List of of strings (which will be treated as an OR filter on the property).
            * `valuesSuggesteds` (`pulumi.Input[list]`) - A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.

        The **import_qualifiers** object supports the following:

          * `filters` (`pulumi.Input[list]`)
            * `negated` (`pulumi.Input[bool]`) - If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
            * `property` (`pulumi.Input[str]`) - A metric time series dimension or property name.
            * `values` (`pulumi.Input[list]`) - (Optional) List of of strings (which will be treated as an OR filter on the property).

          * `metric` (`pulumi.Input[str]`)
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

            __props__['authorized_writer_teams'] = authorized_writer_teams
            __props__['authorized_writer_users'] = authorized_writer_users
            __props__['dashboards'] = dashboards
            __props__['description'] = description
            __props__['import_qualifiers'] = import_qualifiers
            __props__['name'] = name
            __props__['teams'] = teams
        super(DashboardGroup, __self__).__init__(
            'signalfx:index/dashboardGroup:DashboardGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, authorized_writer_teams=None, authorized_writer_users=None, dashboards=None, description=None, import_qualifiers=None, name=None, teams=None):
        """
        Get an existing DashboardGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] authorized_writer_teams: Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        :param pulumi.Input[list] authorized_writer_users: User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[list] dashboards: [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        :param pulumi.Input[str] description: Description of the dashboard group.
        :param pulumi.Input[str] name: Name of the dashboard group.
        :param pulumi.Input[list] teams: Team IDs to associate the dashboard group to.

        The **dashboards** object supports the following:

          * `dashboardId` (`pulumi.Input[str]`) - The dashboard id to mirror
          * `descriptionOverride` (`pulumi.Input[str]`) - The description that will override the original dashboards's description.
          * `filterOverrides` (`pulumi.Input[list]`) - The description that will override the original dashboards's description.
            * `negated` (`pulumi.Input[bool]`) - If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
            * `property` (`pulumi.Input[str]`) - A metric time series dimension or property name.
            * `values` (`pulumi.Input[list]`) - (Optional) List of of strings (which will be treated as an OR filter on the property).

          * `nameOverride` (`pulumi.Input[str]`) - The name that will override the original dashboards's name.
          * `variableOverrides` (`pulumi.Input[list]`)
            * `property` (`pulumi.Input[str]`) - A metric time series dimension or property name.
            * `values` (`pulumi.Input[list]`) - (Optional) List of of strings (which will be treated as an OR filter on the property).
            * `valuesSuggesteds` (`pulumi.Input[list]`) - A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.

        The **import_qualifiers** object supports the following:

          * `filters` (`pulumi.Input[list]`)
            * `negated` (`pulumi.Input[bool]`) - If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
            * `property` (`pulumi.Input[str]`) - A metric time series dimension or property name.
            * `values` (`pulumi.Input[list]`) - (Optional) List of of strings (which will be treated as an OR filter on the property).

          * `metric` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authorized_writer_teams"] = authorized_writer_teams
        __props__["authorized_writer_users"] = authorized_writer_users
        __props__["dashboards"] = dashboards
        __props__["description"] = description
        __props__["import_qualifiers"] = import_qualifiers
        __props__["name"] = name
        __props__["teams"] = teams
        return DashboardGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
