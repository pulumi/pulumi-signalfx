# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TeamArgs', 'Team']

@pulumi.input_type
class TeamArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications_criticals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_defaults: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_infos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_majors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_minors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_warnings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Team resource.
        :param pulumi.Input[str] description: Description of the team.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: List of user IDs to include in the team.
        :param pulumi.Input[str] name: Name of the team.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_criticals: Where to send notifications for critical alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_defaults: Where to send notifications for default alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_infos: Where to send notifications for info alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_majors: Where to send notifications for major alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_minors: Where to send notifications for minor alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_warnings: Where to send notifications for warning alerts
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notifications_criticals is not None:
            pulumi.set(__self__, "notifications_criticals", notifications_criticals)
        if notifications_defaults is not None:
            pulumi.set(__self__, "notifications_defaults", notifications_defaults)
        if notifications_infos is not None:
            pulumi.set(__self__, "notifications_infos", notifications_infos)
        if notifications_majors is not None:
            pulumi.set(__self__, "notifications_majors", notifications_majors)
        if notifications_minors is not None:
            pulumi.set(__self__, "notifications_minors", notifications_minors)
        if notifications_warnings is not None:
            pulumi.set(__self__, "notifications_warnings", notifications_warnings)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the team.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of user IDs to include in the team.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the team.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationsCriticals")
    def notifications_criticals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for critical alerts
        """
        return pulumi.get(self, "notifications_criticals")

    @notifications_criticals.setter
    def notifications_criticals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_criticals", value)

    @property
    @pulumi.getter(name="notificationsDefaults")
    def notifications_defaults(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for default alerts
        """
        return pulumi.get(self, "notifications_defaults")

    @notifications_defaults.setter
    def notifications_defaults(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_defaults", value)

    @property
    @pulumi.getter(name="notificationsInfos")
    def notifications_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for info alerts
        """
        return pulumi.get(self, "notifications_infos")

    @notifications_infos.setter
    def notifications_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_infos", value)

    @property
    @pulumi.getter(name="notificationsMajors")
    def notifications_majors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for major alerts
        """
        return pulumi.get(self, "notifications_majors")

    @notifications_majors.setter
    def notifications_majors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_majors", value)

    @property
    @pulumi.getter(name="notificationsMinors")
    def notifications_minors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for minor alerts
        """
        return pulumi.get(self, "notifications_minors")

    @notifications_minors.setter
    def notifications_minors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_minors", value)

    @property
    @pulumi.getter(name="notificationsWarnings")
    def notifications_warnings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for warning alerts
        """
        return pulumi.get(self, "notifications_warnings")

    @notifications_warnings.setter
    def notifications_warnings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_warnings", value)


@pulumi.input_type
class _TeamState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications_criticals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_defaults: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_infos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_majors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_minors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_warnings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Team resources.
        :param pulumi.Input[str] description: Description of the team.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: List of user IDs to include in the team.
        :param pulumi.Input[str] name: Name of the team.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_criticals: Where to send notifications for critical alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_defaults: Where to send notifications for default alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_infos: Where to send notifications for info alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_majors: Where to send notifications for major alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_minors: Where to send notifications for minor alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_warnings: Where to send notifications for warning alerts
        :param pulumi.Input[str] url: The URL of the team.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notifications_criticals is not None:
            pulumi.set(__self__, "notifications_criticals", notifications_criticals)
        if notifications_defaults is not None:
            pulumi.set(__self__, "notifications_defaults", notifications_defaults)
        if notifications_infos is not None:
            pulumi.set(__self__, "notifications_infos", notifications_infos)
        if notifications_majors is not None:
            pulumi.set(__self__, "notifications_majors", notifications_majors)
        if notifications_minors is not None:
            pulumi.set(__self__, "notifications_minors", notifications_minors)
        if notifications_warnings is not None:
            pulumi.set(__self__, "notifications_warnings", notifications_warnings)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the team.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of user IDs to include in the team.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the team.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationsCriticals")
    def notifications_criticals(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for critical alerts
        """
        return pulumi.get(self, "notifications_criticals")

    @notifications_criticals.setter
    def notifications_criticals(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_criticals", value)

    @property
    @pulumi.getter(name="notificationsDefaults")
    def notifications_defaults(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for default alerts
        """
        return pulumi.get(self, "notifications_defaults")

    @notifications_defaults.setter
    def notifications_defaults(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_defaults", value)

    @property
    @pulumi.getter(name="notificationsInfos")
    def notifications_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for info alerts
        """
        return pulumi.get(self, "notifications_infos")

    @notifications_infos.setter
    def notifications_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_infos", value)

    @property
    @pulumi.getter(name="notificationsMajors")
    def notifications_majors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for major alerts
        """
        return pulumi.get(self, "notifications_majors")

    @notifications_majors.setter
    def notifications_majors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_majors", value)

    @property
    @pulumi.getter(name="notificationsMinors")
    def notifications_minors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for minor alerts
        """
        return pulumi.get(self, "notifications_minors")

    @notifications_minors.setter
    def notifications_minors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_minors", value)

    @property
    @pulumi.getter(name="notificationsWarnings")
    def notifications_warnings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Where to send notifications for warning alerts
        """
        return pulumi.get(self, "notifications_warnings")

    @notifications_warnings.setter
    def notifications_warnings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notifications_warnings", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the team.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class Team(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications_criticals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_defaults: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_infos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_majors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_minors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_warnings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Handles management of SignalFx teams.

        You can configure [team notification policies](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html) using this resource and the various `notifications_*` properties.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        myteam0 = signalfx.Team("myteam0",
            description="Super great team no jerks definitely",
            members=[
                "userid1",
                "userid2",
            ],
            notifications_criticals=["PagerDuty,credentialId"],
            notifications_infos=["Email,notify@example.com"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the team.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: List of user IDs to include in the team.
        :param pulumi.Input[str] name: Name of the team.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_criticals: Where to send notifications for critical alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_defaults: Where to send notifications for default alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_infos: Where to send notifications for info alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_majors: Where to send notifications for major alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_minors: Where to send notifications for minor alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_warnings: Where to send notifications for warning alerts
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TeamArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Handles management of SignalFx teams.

        You can configure [team notification policies](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html) using this resource and the various `notifications_*` properties.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_signalfx as signalfx

        myteam0 = signalfx.Team("myteam0",
            description="Super great team no jerks definitely",
            members=[
                "userid1",
                "userid2",
            ],
            notifications_criticals=["PagerDuty,credentialId"],
            notifications_infos=["Email,notify@example.com"])
        ```

        :param str resource_name: The name of the resource.
        :param TeamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications_criticals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_defaults: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_infos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_majors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_minors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications_warnings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeamArgs.__new__(TeamArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["notifications_criticals"] = notifications_criticals
            __props__.__dict__["notifications_defaults"] = notifications_defaults
            __props__.__dict__["notifications_infos"] = notifications_infos
            __props__.__dict__["notifications_majors"] = notifications_majors
            __props__.__dict__["notifications_minors"] = notifications_minors
            __props__.__dict__["notifications_warnings"] = notifications_warnings
            __props__.__dict__["url"] = None
        super(Team, __self__).__init__(
            'signalfx:index/team:Team',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notifications_criticals: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            notifications_defaults: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            notifications_infos: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            notifications_majors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            notifications_minors: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            notifications_warnings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'Team':
        """
        Get an existing Team resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the team.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: List of user IDs to include in the team.
        :param pulumi.Input[str] name: Name of the team.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_criticals: Where to send notifications for critical alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_defaults: Where to send notifications for default alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_infos: Where to send notifications for info alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_majors: Where to send notifications for major alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_minors: Where to send notifications for minor alerts
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications_warnings: Where to send notifications for warning alerts
        :param pulumi.Input[str] url: The URL of the team.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamState.__new__(_TeamState)

        __props__.__dict__["description"] = description
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["notifications_criticals"] = notifications_criticals
        __props__.__dict__["notifications_defaults"] = notifications_defaults
        __props__.__dict__["notifications_infos"] = notifications_infos
        __props__.__dict__["notifications_majors"] = notifications_majors
        __props__.__dict__["notifications_minors"] = notifications_minors
        __props__.__dict__["notifications_warnings"] = notifications_warnings
        __props__.__dict__["url"] = url
        return Team(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the team.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of user IDs to include in the team.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the team.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationsCriticals")
    def notifications_criticals(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Where to send notifications for critical alerts
        """
        return pulumi.get(self, "notifications_criticals")

    @property
    @pulumi.getter(name="notificationsDefaults")
    def notifications_defaults(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Where to send notifications for default alerts
        """
        return pulumi.get(self, "notifications_defaults")

    @property
    @pulumi.getter(name="notificationsInfos")
    def notifications_infos(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Where to send notifications for info alerts
        """
        return pulumi.get(self, "notifications_infos")

    @property
    @pulumi.getter(name="notificationsMajors")
    def notifications_majors(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Where to send notifications for major alerts
        """
        return pulumi.get(self, "notifications_majors")

    @property
    @pulumi.getter(name="notificationsMinors")
    def notifications_minors(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Where to send notifications for minor alerts
        """
        return pulumi.get(self, "notifications_minors")

    @property
    @pulumi.getter(name="notificationsWarnings")
    def notifications_warnings(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Where to send notifications for warning alerts
        """
        return pulumi.get(self, "notifications_warnings")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the team.
        """
        return pulumi.get(self, "url")

