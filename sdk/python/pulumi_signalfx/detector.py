# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Detector']


class Detector(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_sampling: Optional[pulumi.Input[bool]] = None,
                 end_time: Optional[pulumi.Input[int]] = None,
                 max_delay: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 program_text: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorRuleArgs']]]]] = None,
                 show_data_markers: Optional[pulumi.Input[bool]] = None,
                 show_event_lines: Optional[pulumi.Input[bool]] = None,
                 start_time: Optional[pulumi.Input[int]] = None,
                 teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 time_range: Optional[pulumi.Input[int]] = None,
                 viz_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorVizOptionArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a SignalFx detector resource. This can be used to create and manage detectors.

        > **NOTE** If you're interested in using SignalFx detector features such as Historical Anomaly, Resource Running Out, or others then consider building them in the UI first then using the "Show SignalFlow" feature to extract the value for `program_text`. You may also consult the [documentation for detector functions in signalflow-library](https://github.com/signalfx/signalflow-library/tree/master/library/signalfx/detectors).

        ## Notification Format

        As SignalFx supports different notification mechanisms a comma-delimited string is used to provide inputs. If you'd like to specify multiple notifications, then each should be a member in the list, like so:

        ```python
        import pulumi
        ```

        This will likely be changed in a future iteration of the provider. See [SignalFx Docs](https://developers.signalfx.com/detectors_reference.html#operation/Create%20Single%20Detector) for more information. For now, here are some example of how to configure each notification type:

        ### Email

        ```python
        import pulumi
        ```

        ### Jira

        Note that the `credentialId` is the SignalFx-provided ID shown after setting up your Jira integration. (See also `jira.Integration`.)

        ```python
        import pulumi
        ```

        ### Opsgenie

        Note that the `credentialId` is the SignalFx-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.

        ```python
        import pulumi
        ```

        ### PagerDuty

        ```python
        import pulumi
        ```

        ### Slack

        Exclude the `#` on the channel name!

        ```python
        import pulumi
        ```

        ### Team

        Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).

        ```python
        import pulumi
        ```

        ### TeamEmail

        Sends an email to every member of a team.

        ```python
        import pulumi
        ```

        ### VictorOps

        ```python
        import pulumi
        ```

        ### Webhook

        > **NOTE** You need to include all the commas even if you only use a credential id below.

        You can either configure a Webhook to use an existing integration's credential id:
        ```python
        import pulumi
        ```

        or configure one inline:
        ```python
        import pulumi
        ```

        ## Import

        Detectors can be imported using their string ID (recoverable from URL`/#/detector/v2/abc123/edit`, e.g.

        ```sh
         $ pulumi import signalfx:index/detector:Detector application_delay abc123
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_teams: Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorized_writer_users`).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_users: User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[str] description: Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
        :param pulumi.Input[bool] disable_sampling: When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
        :param pulumi.Input[int] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[int] max_delay: How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
        :param pulumi.Input[str] name: Name of the detector.
        :param pulumi.Input[str] program_text: Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorRuleArgs']]]] rules: Set of rules used for alerting.
        :param pulumi.Input[bool] show_data_markers: When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
        :param pulumi.Input[bool] show_event_lines: When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
        :param pulumi.Input[int] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] teams: Team IDs to associate the detector to.
        :param pulumi.Input[int] time_range: Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorVizOptionArgs']]]] viz_options: Plot-level customization options, associated with a publish statement.
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

            __props__['authorized_writer_teams'] = authorized_writer_teams
            __props__['authorized_writer_users'] = authorized_writer_users
            __props__['description'] = description
            __props__['disable_sampling'] = disable_sampling
            __props__['end_time'] = end_time
            __props__['max_delay'] = max_delay
            __props__['name'] = name
            if program_text is None:
                raise TypeError("Missing required property 'program_text'")
            __props__['program_text'] = program_text
            if rules is None:
                raise TypeError("Missing required property 'rules'")
            __props__['rules'] = rules
            __props__['show_data_markers'] = show_data_markers
            __props__['show_event_lines'] = show_event_lines
            __props__['start_time'] = start_time
            __props__['teams'] = teams
            __props__['time_range'] = time_range
            __props__['viz_options'] = viz_options
            __props__['url'] = None
        super(Detector, __self__).__init__(
            'signalfx:index/detector:Detector',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorized_writer_teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            authorized_writer_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disable_sampling: Optional[pulumi.Input[bool]] = None,
            end_time: Optional[pulumi.Input[int]] = None,
            max_delay: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            program_text: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorRuleArgs']]]]] = None,
            show_data_markers: Optional[pulumi.Input[bool]] = None,
            show_event_lines: Optional[pulumi.Input[bool]] = None,
            start_time: Optional[pulumi.Input[int]] = None,
            teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            time_range: Optional[pulumi.Input[int]] = None,
            url: Optional[pulumi.Input[str]] = None,
            viz_options: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorVizOptionArgs']]]]] = None) -> 'Detector':
        """
        Get an existing Detector resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_teams: Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorized_writer_users`).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_writer_users: User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        :param pulumi.Input[str] description: Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
        :param pulumi.Input[bool] disable_sampling: When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
        :param pulumi.Input[int] end_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[int] max_delay: How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
        :param pulumi.Input[str] name: Name of the detector.
        :param pulumi.Input[str] program_text: Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorRuleArgs']]]] rules: Set of rules used for alerting.
        :param pulumi.Input[bool] show_data_markers: When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
        :param pulumi.Input[bool] show_event_lines: When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
        :param pulumi.Input[int] start_time: Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] teams: Team IDs to associate the detector to.
        :param pulumi.Input[int] time_range: Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
        :param pulumi.Input[str] url: The URL of the detector.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DetectorVizOptionArgs']]]] viz_options: Plot-level customization options, associated with a publish statement.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authorized_writer_teams"] = authorized_writer_teams
        __props__["authorized_writer_users"] = authorized_writer_users
        __props__["description"] = description
        __props__["disable_sampling"] = disable_sampling
        __props__["end_time"] = end_time
        __props__["max_delay"] = max_delay
        __props__["name"] = name
        __props__["program_text"] = program_text
        __props__["rules"] = rules
        __props__["show_data_markers"] = show_data_markers
        __props__["show_event_lines"] = show_event_lines
        __props__["start_time"] = start_time
        __props__["teams"] = teams
        __props__["time_range"] = time_range
        __props__["url"] = url
        __props__["viz_options"] = viz_options
        return Detector(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizedWriterTeams")
    def authorized_writer_teams(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorized_writer_users`).
        """
        return pulumi.get(self, "authorized_writer_teams")

    @property
    @pulumi.getter(name="authorizedWriterUsers")
    def authorized_writer_users(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        """
        return pulumi.get(self, "authorized_writer_users")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableSampling")
    def disable_sampling(self) -> pulumi.Output[Optional[bool]]:
        """
        When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
        """
        return pulumi.get(self, "disable_sampling")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="maxDelay")
    def max_delay(self) -> pulumi.Output[Optional[int]]:
        """
        How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
        """
        return pulumi.get(self, "max_delay")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the detector.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="programText")
    def program_text(self) -> pulumi.Output[str]:
        """
        Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
        """
        return pulumi.get(self, "program_text")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.DetectorRule']]:
        """
        Set of rules used for alerting.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="showDataMarkers")
    def show_data_markers(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
        """
        return pulumi.get(self, "show_data_markers")

    @property
    @pulumi.getter(name="showEventLines")
    def show_event_lines(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
        """
        return pulumi.get(self, "show_event_lines")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[int]]:
        """
        Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def teams(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Team IDs to associate the detector to.
        """
        return pulumi.get(self, "teams")

    @property
    @pulumi.getter(name="timeRange")
    def time_range(self) -> pulumi.Output[Optional[int]]:
        """
        Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
        """
        return pulumi.get(self, "time_range")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the detector.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter(name="vizOptions")
    def viz_options(self) -> pulumi.Output[Optional[Sequence['outputs.DetectorVizOption']]]:
        """
        Plot-level customization options, associated with a publish statement.
        """
        return pulumi.get(self, "viz_options")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

