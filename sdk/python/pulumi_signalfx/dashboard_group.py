# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DashboardGroupArgs', 'DashboardGroup']

@pulumi.input_type
class DashboardGroupArgs:
    def __init__(__self__, *,
                 authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dashboards: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupDashboardArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 import_qualifiers: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupImportQualifierArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a DashboardGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_teams: Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_users: User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input['DashboardGroupDashboardArgs']]] dashboards: [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        :param pulumi.Input[str] description: Description of the dashboard group.
        :param pulumi.Input[str] name: Name of the dashboard group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] teams: Team IDs to associate the dashboard group to.
        """
        if authorized_writer_teams is not None:
            pulumi.set(__self__, "authorized_writer_teams", authorized_writer_teams)
        if authorized_writer_users is not None:
            pulumi.set(__self__, "authorized_writer_users", authorized_writer_users)
        if dashboards is not None:
            pulumi.set(__self__, "dashboards", dashboards)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if import_qualifiers is not None:
            pulumi.set(__self__, "import_qualifiers", import_qualifiers)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if teams is not None:
            pulumi.set(__self__, "teams", teams)

    @property
    @pulumi.getter(name="authorizedWriterTeams")
    def authorized_writer_teams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_teams")

    @authorized_writer_teams.setter
    def authorized_writer_teams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_writer_teams", value)

    @property
    @pulumi.getter(name="authorizedWriterUsers")
    def authorized_writer_users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_users")

    @authorized_writer_users.setter
    def authorized_writer_users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_writer_users", value)

    @property
    @pulumi.getter
    def dashboards(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupDashboardArgs']]]]:
        """
        [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        """
        return pulumi.get(self, "dashboards")

    @dashboards.setter
    def dashboards(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupDashboardArgs']]]]):
        pulumi.set(self, "dashboards", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the dashboard group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="importQualifiers")
    def import_qualifiers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupImportQualifierArgs']]]]:
        return pulumi.get(self, "import_qualifiers")

    @import_qualifiers.setter
    def import_qualifiers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupImportQualifierArgs']]]]):
        pulumi.set(self, "import_qualifiers", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the dashboard group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def teams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Team IDs to associate the dashboard group to.
        """
        return pulumi.get(self, "teams")

    @teams.setter
    def teams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "teams", value)


@pulumi.input_type
class _DashboardGroupState:
    def __init__(__self__, *,
                 authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dashboards: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupDashboardArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 import_qualifiers: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupImportQualifierArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering DashboardGroup resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_teams: Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_users: User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input['DashboardGroupDashboardArgs']]] dashboards: [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        :param pulumi.Input[str] description: Description of the dashboard group.
        :param pulumi.Input[str] name: Name of the dashboard group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] teams: Team IDs to associate the dashboard group to.
        """
        if authorized_writer_teams is not None:
            pulumi.set(__self__, "authorized_writer_teams", authorized_writer_teams)
        if authorized_writer_users is not None:
            pulumi.set(__self__, "authorized_writer_users", authorized_writer_users)
        if dashboards is not None:
            pulumi.set(__self__, "dashboards", dashboards)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if import_qualifiers is not None:
            pulumi.set(__self__, "import_qualifiers", import_qualifiers)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if teams is not None:
            pulumi.set(__self__, "teams", teams)

    @property
    @pulumi.getter(name="authorizedWriterTeams")
    def authorized_writer_teams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_teams")

    @authorized_writer_teams.setter
    def authorized_writer_teams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_writer_teams", value)

    @property
    @pulumi.getter(name="authorizedWriterUsers")
    def authorized_writer_users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_users")

    @authorized_writer_users.setter
    def authorized_writer_users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_writer_users", value)

    @property
    @pulumi.getter
    def dashboards(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupDashboardArgs']]]]:
        """
        [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        """
        return pulumi.get(self, "dashboards")

    @dashboards.setter
    def dashboards(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupDashboardArgs']]]]):
        pulumi.set(self, "dashboards", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the dashboard group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="importQualifiers")
    def import_qualifiers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupImportQualifierArgs']]]]:
        return pulumi.get(self, "import_qualifiers")

    @import_qualifiers.setter
    def import_qualifiers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardGroupImportQualifierArgs']]]]):
        pulumi.set(self, "import_qualifiers", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the dashboard group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def teams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Team IDs to associate the dashboard group to.
        """
        return pulumi.get(self, "teams")

    @teams.setter
    def teams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "teams", value)


class DashboardGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dashboards: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGroupDashboardArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 import_qualifiers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGroupImportQualifierArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            dashboards=[signalfx.DashboardGroupDashboardArgs(
                dashboard_id=signalfx_dashboard["gc_dashboard"]["id"],
                name_override="GC For My Service",
                description_override="Garbage Collection dashboard maintained by JVM team",
                filter_overrides=[signalfx.DashboardGroupDashboardFilterOverrideArgs(
                    property="service",
                    values=["myservice"],
                    negated=False,
                )],
                variable_overrides=[signalfx.DashboardGroupDashboardVariableOverrideArgs(
                    property="region",
                    values=["us-west1"],
                    values_suggesteds=[
                        "us-west-1",
                        "us-east-1",
                    ],
                )],
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_teams: Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_users: User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGroupDashboardArgs']]]] dashboards: [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        :param pulumi.Input[str] description: Description of the dashboard group.
        :param pulumi.Input[str] name: Name of the dashboard group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] teams: Team IDs to associate the dashboard group to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DashboardGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
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
            dashboards=[signalfx.DashboardGroupDashboardArgs(
                dashboard_id=signalfx_dashboard["gc_dashboard"]["id"],
                name_override="GC For My Service",
                description_override="Garbage Collection dashboard maintained by JVM team",
                filter_overrides=[signalfx.DashboardGroupDashboardFilterOverrideArgs(
                    property="service",
                    values=["myservice"],
                    negated=False,
                )],
                variable_overrides=[signalfx.DashboardGroupDashboardVariableOverrideArgs(
                    property="region",
                    values=["us-west1"],
                    values_suggesteds=[
                        "us-west-1",
                        "us-east-1",
                    ],
                )],
            )])
        ```

        :param str resource_name: The name of the resource.
        :param DashboardGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DashboardGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dashboards: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGroupDashboardArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 import_qualifiers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGroupImportQualifierArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            __props__ = DashboardGroupArgs.__new__(DashboardGroupArgs)

            __props__.__dict__["authorized_writer_teams"] = authorized_writer_teams
            __props__.__dict__["authorized_writer_users"] = authorized_writer_users
            __props__.__dict__["dashboards"] = dashboards
            __props__.__dict__["description"] = description
            __props__.__dict__["import_qualifiers"] = import_qualifiers
            __props__.__dict__["name"] = name
            __props__.__dict__["teams"] = teams
        super(DashboardGroup, __self__).__init__(
            'signalfx:index/dashboardGroup:DashboardGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            dashboards: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGroupDashboardArgs']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            import_qualifiers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGroupImportQualifierArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'DashboardGroup':
        """
        Get an existing DashboardGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_teams: Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_users: User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardGroupDashboardArgs']]]] dashboards: [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        :param pulumi.Input[str] description: Description of the dashboard group.
        :param pulumi.Input[str] name: Name of the dashboard group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] teams: Team IDs to associate the dashboard group to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DashboardGroupState.__new__(_DashboardGroupState)

        __props__.__dict__["authorized_writer_teams"] = authorized_writer_teams
        __props__.__dict__["authorized_writer_users"] = authorized_writer_users
        __props__.__dict__["dashboards"] = dashboards
        __props__.__dict__["description"] = description
        __props__.__dict__["import_qualifiers"] = import_qualifiers
        __props__.__dict__["name"] = name
        __props__.__dict__["teams"] = teams
        return DashboardGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizedWriterTeams")
    def authorized_writer_teams(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_teams")

    @property
    @pulumi.getter(name="authorizedWriterUsers")
    def authorized_writer_users(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_users")

    @property
    @pulumi.getter
    def dashboards(self) -> pulumi.Output[Optional[Sequence['outputs.DashboardGroupDashboard']]]:
        """
        [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
        """
        return pulumi.get(self, "dashboards")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the dashboard group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="importQualifiers")
    def import_qualifiers(self) -> pulumi.Output[Optional[Sequence['outputs.DashboardGroupImportQualifier']]]:
        return pulumi.get(self, "import_qualifiers")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the dashboard group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def teams(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Team IDs to associate the dashboard group to.
        """
        return pulumi.get(self, "teams")

