// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx
{
    /// <summary>
    /// Provides a Splunk Observability Cloud detector resource. This can be used to create and manage detectors.
    /// 
    /// If you're interested in using Splunk Observability Cloud detector features such as Historical Anomaly, Resource Running Out, or others, consider building them in the UI first and then use the "Show SignalFlow" feature to extract the value for `program_text`. You can also see the [documentation for detector functions in signalflow-library](https://github.com/signalfx/signalflow-library/tree/master/library/signalfx/detectors).
    /// 
    /// &gt; **NOTE** When you want to change or remove write permissions for a user other than yourself regarding detectors, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
    /// 
    /// ## Example
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using SignalFx = Pulumi.SignalFx;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var clusters = config.GetObject&lt;dynamic&gt;("clusters") ?? new[]
    ///     {
    ///         "clusterA",
    ///         "clusterB",
    ///     };
    ///     var applicationDelay = new List&lt;SignalFx.Detector&gt;();
    ///     for (var rangeIndex = 0; rangeIndex &lt; clusters.Length; rangeIndex++)
    ///     {
    ///         var range = new { Value = rangeIndex };
    ///         applicationDelay.Add(new SignalFx.Detector($"application_delay-{range.Value}", new()
    ///         {
    ///             Name = $" max average delay - {clusters[range.Value]}",
    ///             Description = $"your application is slow - {clusters[range.Value]}",
    ///             MaxDelay = 30,
    ///             Tags = new[]
    ///             {
    ///                 "app-backend",
    ///                 "staging",
    ///             },
    ///             AuthorizedWriterTeams = new[]
    ///             {
    ///                 mycoolteam.Id,
    ///             },
    ///             AuthorizedWriterUsers = new[]
    ///             {
    ///                 "abc123",
    ///             },
    ///             ProgramText = @$"signal = data('app.delay', filter('cluster','{clusters[range.Value]}'), extrapolation='last_value', maxExtrapolations=5).max()
    /// detect(when(signal &gt; 60, '5m')).publish('Processing old messages 5m')
    /// detect(when(signal &gt; 60, '30m')).publish('Processing old messages 30m')
    /// ",
    ///             Rules = new[]
    ///             {
    ///                 new SignalFx.Inputs.DetectorRuleArgs
    ///                 {
    ///                     Description = "maximum &gt; 60 for 5m",
    ///                     Severity = "Warning",
    ///                     DetectLabel = "Processing old messages 5m",
    ///                     Notifications = new[]
    ///                     {
    ///                         "Email,foo-alerts@bar.com",
    ///                     },
    ///                 },
    ///                 new SignalFx.Inputs.DetectorRuleArgs
    ///                 {
    ///                     Description = "maximum &gt; 60 for 30m",
    ///                     Severity = "Critical",
    ///                     DetectLabel = "Processing old messages 30m",
    ///                     Notifications = new[]
    ///                     {
    ///                         "Email,foo-alerts@bar.com",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///     }
    /// });
    /// ```
    /// 
    /// ## Notification format
    /// 
    /// As Splunk Observability Cloud supports different notification mechanisms, use a comma-delimited string to provide inputs. If you want to specify multiple notifications, each must be a member in the list, like so:
    /// 
    /// See [Splunk Observability Cloud Docs](https://dev.splunk.com/observability/reference/api/detectors/latest) for more information.
    /// 
    /// Here are some example of how to configure each notification type:
    /// 
    /// ### Email
    /// 
    /// ### Jira
    /// 
    /// Note that the `credentialId` is the Splunk-provided ID shown after setting up your Jira integration. See also `signalfx.jira.Integration`.
    /// 
    /// ### OpsGenie
    /// 
    /// Note that the `credentialId` is the Splunk-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.
    /// 
    /// ### PagerDuty
    /// 
    /// ### Slack
    /// 
    /// Exclude the `#` on the channel name:
    /// 
    /// ### Team
    /// 
    /// Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).
    /// 
    /// ### TeamEmail
    /// 
    /// Sends an email to every member of a team.
    /// 
    /// ### Splunk On-Call (formerly VictorOps)
    /// 
    /// ### Webhooks
    /// 
    /// You need to include all the commas even if you only use a credential id.
    /// 
    /// You can either configure a Webhook to use an existing integration's credential id:
    /// 
    /// Or configure one inline:
    /// 
    /// ## Import
    /// 
    /// Detectors can be imported using their string ID (recoverable from URL: `/#/detector/v2/abc123/edit`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import signalfx:index/detector:Detector application_delay abc123
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:index/detector:Detector")]
    public partial class Detector : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorized_writer_users`).
        /// </summary>
        [Output("authorizedWriterTeams")]
        public Output<ImmutableArray<string>> AuthorizedWriterTeams { get; private set; } = null!;

        /// <summary>
        /// User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        /// </summary>
        [Output("authorizedWriterUsers")]
        public Output<ImmutableArray<string>> AuthorizedWriterUsers { get; private set; } = null!;

        /// <summary>
        /// Description of the detector.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates how a detector was created. The possible values are: Standard and AutoDetectCustomization. The value can only be set when creating the detector and cannot be modified later.
        /// </summary>
        [Output("detectorOrigin")]
        public Output<string?> DetectorOrigin { get; private set; } = null!;

        /// <summary>
        /// When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
        /// </summary>
        [Output("disableSampling")]
        public Output<bool?> DisableSampling { get; private set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Output("endTime")]
        public Output<int?> EndTime { get; private set; } = null!;

        /// <summary>
        /// The resolutions of the detector alerts in milliseconds that indicate how often data is analyzed to determine if an alert should be triggered.
        /// </summary>
        [Output("labelResolutions")]
        public Output<ImmutableDictionary<string, int>> LabelResolutions { get; private set; } = null!;

        /// <summary>
        /// allows Splunk Observability Cloud to continue with computation if there is a lag in receiving data points.
        /// </summary>
        [Output("maxDelay")]
        public Output<int?> MaxDelay { get; private set; } = null!;

        /// <summary>
        /// How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
        /// </summary>
        [Output("minDelay")]
        public Output<int?> MinDelay { get; private set; } = null!;

        /// <summary>
        /// Name of the detector.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the AutoDetect parent detector from which this detector is customized and created. This property is required for detectors with detectorOrigin of type AutoDetectCustomization. The value can only be set when creating the detector and cannot be modified later.
        /// </summary>
        [Output("parentDetectorId")]
        public Output<string?> ParentDetectorId { get; private set; } = null!;

        /// <summary>
        /// Signalflow program text for the detector. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
        /// </summary>
        [Output("programText")]
        public Output<string> ProgramText { get; private set; } = null!;

        /// <summary>
        /// Set of rules used for alerting.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.DetectorRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
        /// </summary>
        [Output("showDataMarkers")]
        public Output<bool?> ShowDataMarkers { get; private set; } = null!;

        /// <summary>
        /// When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
        /// </summary>
        [Output("showEventLines")]
        public Output<bool?> ShowEventLines { get; private set; } = null!;

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Output("startTime")]
        public Output<int?> StartTime { get; private set; } = null!;

        /// <summary>
        /// Tags associated with the detector.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Team IDs to associate the detector to.
        /// </summary>
        [Output("teams")]
        public Output<ImmutableArray<string>> Teams { get; private set; } = null!;

        /// <summary>
        /// Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
        /// </summary>
        [Output("timeRange")]
        public Output<int?> TimeRange { get; private set; } = null!;

        /// <summary>
        /// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;

        /// <summary>
        /// The URL of the detector.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// Plot-level customization options, associated with a publish statement.
        /// </summary>
        [Output("vizOptions")]
        public Output<ImmutableArray<Outputs.DetectorVizOption>> VizOptions { get; private set; } = null!;


        /// <summary>
        /// Create a Detector resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Detector(string name, DetectorArgs args, CustomResourceOptions? options = null)
            : base("signalfx:index/detector:Detector", name, args ?? new DetectorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Detector(string name, Input<string> id, DetectorState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:index/detector:Detector", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Detector resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Detector Get(string name, Input<string> id, DetectorState? state = null, CustomResourceOptions? options = null)
        {
            return new Detector(name, id, state, options);
        }
    }

    public sealed class DetectorArgs : global::Pulumi.ResourceArgs
    {
        [Input("authorizedWriterTeams")]
        private InputList<string>? _authorizedWriterTeams;

        /// <summary>
        /// Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorized_writer_users`).
        /// </summary>
        public InputList<string> AuthorizedWriterTeams
        {
            get => _authorizedWriterTeams ?? (_authorizedWriterTeams = new InputList<string>());
            set => _authorizedWriterTeams = value;
        }

        [Input("authorizedWriterUsers")]
        private InputList<string>? _authorizedWriterUsers;

        /// <summary>
        /// User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        /// </summary>
        public InputList<string> AuthorizedWriterUsers
        {
            get => _authorizedWriterUsers ?? (_authorizedWriterUsers = new InputList<string>());
            set => _authorizedWriterUsers = value;
        }

        /// <summary>
        /// Description of the detector.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates how a detector was created. The possible values are: Standard and AutoDetectCustomization. The value can only be set when creating the detector and cannot be modified later.
        /// </summary>
        [Input("detectorOrigin")]
        public Input<string>? DetectorOrigin { get; set; }

        /// <summary>
        /// When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
        /// </summary>
        [Input("disableSampling")]
        public Input<bool>? DisableSampling { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        /// <summary>
        /// allows Splunk Observability Cloud to continue with computation if there is a lag in receiving data points.
        /// </summary>
        [Input("maxDelay")]
        public Input<int>? MaxDelay { get; set; }

        /// <summary>
        /// How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
        /// </summary>
        [Input("minDelay")]
        public Input<int>? MinDelay { get; set; }

        /// <summary>
        /// Name of the detector.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the AutoDetect parent detector from which this detector is customized and created. This property is required for detectors with detectorOrigin of type AutoDetectCustomization. The value can only be set when creating the detector and cannot be modified later.
        /// </summary>
        [Input("parentDetectorId")]
        public Input<string>? ParentDetectorId { get; set; }

        /// <summary>
        /// Signalflow program text for the detector. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
        /// </summary>
        [Input("programText", required: true)]
        public Input<string> ProgramText { get; set; } = null!;

        [Input("rules", required: true)]
        private InputList<Inputs.DetectorRuleArgs>? _rules;

        /// <summary>
        /// Set of rules used for alerting.
        /// </summary>
        public InputList<Inputs.DetectorRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.DetectorRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
        /// </summary>
        [Input("showDataMarkers")]
        public Input<bool>? ShowDataMarkers { get; set; }

        /// <summary>
        /// When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
        /// </summary>
        [Input("showEventLines")]
        public Input<bool>? ShowEventLines { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags associated with the detector.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("teams")]
        private InputList<string>? _teams;

        /// <summary>
        /// Team IDs to associate the detector to.
        /// </summary>
        public InputList<string> Teams
        {
            get => _teams ?? (_teams = new InputList<string>());
            set => _teams = value;
        }

        /// <summary>
        /// Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
        /// </summary>
        [Input("timeRange")]
        public Input<int>? TimeRange { get; set; }

        /// <summary>
        /// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        [Input("vizOptions")]
        private InputList<Inputs.DetectorVizOptionArgs>? _vizOptions;

        /// <summary>
        /// Plot-level customization options, associated with a publish statement.
        /// </summary>
        public InputList<Inputs.DetectorVizOptionArgs> VizOptions
        {
            get => _vizOptions ?? (_vizOptions = new InputList<Inputs.DetectorVizOptionArgs>());
            set => _vizOptions = value;
        }

        public DetectorArgs()
        {
        }
        public static new DetectorArgs Empty => new DetectorArgs();
    }

    public sealed class DetectorState : global::Pulumi.ResourceArgs
    {
        [Input("authorizedWriterTeams")]
        private InputList<string>? _authorizedWriterTeams;

        /// <summary>
        /// Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorized_writer_users`).
        /// </summary>
        public InputList<string> AuthorizedWriterTeams
        {
            get => _authorizedWriterTeams ?? (_authorizedWriterTeams = new InputList<string>());
            set => _authorizedWriterTeams = value;
        }

        [Input("authorizedWriterUsers")]
        private InputList<string>? _authorizedWriterUsers;

        /// <summary>
        /// User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorized_writer_teams`).
        /// </summary>
        public InputList<string> AuthorizedWriterUsers
        {
            get => _authorizedWriterUsers ?? (_authorizedWriterUsers = new InputList<string>());
            set => _authorizedWriterUsers = value;
        }

        /// <summary>
        /// Description of the detector.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates how a detector was created. The possible values are: Standard and AutoDetectCustomization. The value can only be set when creating the detector and cannot be modified later.
        /// </summary>
        [Input("detectorOrigin")]
        public Input<string>? DetectorOrigin { get; set; }

        /// <summary>
        /// When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
        /// </summary>
        [Input("disableSampling")]
        public Input<bool>? DisableSampling { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("endTime")]
        public Input<int>? EndTime { get; set; }

        [Input("labelResolutions")]
        private InputMap<int>? _labelResolutions;

        /// <summary>
        /// The resolutions of the detector alerts in milliseconds that indicate how often data is analyzed to determine if an alert should be triggered.
        /// </summary>
        public InputMap<int> LabelResolutions
        {
            get => _labelResolutions ?? (_labelResolutions = new InputMap<int>());
            set => _labelResolutions = value;
        }

        /// <summary>
        /// allows Splunk Observability Cloud to continue with computation if there is a lag in receiving data points.
        /// </summary>
        [Input("maxDelay")]
        public Input<int>? MaxDelay { get; set; }

        /// <summary>
        /// How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
        /// </summary>
        [Input("minDelay")]
        public Input<int>? MinDelay { get; set; }

        /// <summary>
        /// Name of the detector.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the AutoDetect parent detector from which this detector is customized and created. This property is required for detectors with detectorOrigin of type AutoDetectCustomization. The value can only be set when creating the detector and cannot be modified later.
        /// </summary>
        [Input("parentDetectorId")]
        public Input<string>? ParentDetectorId { get; set; }

        /// <summary>
        /// Signalflow program text for the detector. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
        /// </summary>
        [Input("programText")]
        public Input<string>? ProgramText { get; set; }

        [Input("rules")]
        private InputList<Inputs.DetectorRuleGetArgs>? _rules;

        /// <summary>
        /// Set of rules used for alerting.
        /// </summary>
        public InputList<Inputs.DetectorRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.DetectorRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
        /// </summary>
        [Input("showDataMarkers")]
        public Input<bool>? ShowDataMarkers { get; set; }

        /// <summary>
        /// When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
        /// </summary>
        [Input("showEventLines")]
        public Input<bool>? ShowEventLines { get; set; }

        /// <summary>
        /// Seconds since epoch. Used for visualization. Conflicts with `time_range`.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// Tags associated with the detector.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("teams")]
        private InputList<string>? _teams;

        /// <summary>
        /// Team IDs to associate the detector to.
        /// </summary>
        public InputList<string> Teams
        {
            get => _teams ?? (_teams = new InputList<string>());
            set => _teams = value;
        }

        /// <summary>
        /// Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
        /// </summary>
        [Input("timeRange")]
        public Input<int>? TimeRange { get; set; }

        /// <summary>
        /// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// The URL of the detector.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("vizOptions")]
        private InputList<Inputs.DetectorVizOptionGetArgs>? _vizOptions;

        /// <summary>
        /// Plot-level customization options, associated with a publish statement.
        /// </summary>
        public InputList<Inputs.DetectorVizOptionGetArgs> VizOptions
        {
            get => _vizOptions ?? (_vizOptions = new InputList<Inputs.DetectorVizOptionGetArgs>());
            set => _vizOptions = value;
        }

        public DetectorState()
        {
        }
        public static new DetectorState Empty => new DetectorState();
    }
}
