// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.DetectorArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.DetectorState;
import com.pulumi.signalfx.outputs.DetectorRule;
import com.pulumi.signalfx.outputs.DetectorVizOption;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Splunk Observability Cloud detector resource. This can be used to create and manage detectors.
 * 
 * If you&#39;re interested in using Splunk Observability Cloud detector features such as Historical Anomaly, Resource Running Out, or others, consider building them in the UI first and then use the &#34;Show SignalFlow&#34; feature to extract the value for `program_text`. You can also see the [documentation for detector functions in signalflow-library](https://github.com/signalfx/signalflow-library/tree/master/library/signalfx/detectors).
 * 
 * &gt; **NOTE** When you want to change or remove write permissions for a user other than yourself regarding detectors, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
 * 
 * ## Example
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.Detector;
 * import com.pulumi.signalfx.DetectorArgs;
 * import com.pulumi.signalfx.inputs.DetectorRuleArgs;
 * import com.pulumi.codegen.internal.KeyedValue;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var clusters = config.get(&#34;clusters&#34;).orElse(        
 *             &#34;clusterA&#34;,
 *             &#34;clusterB&#34;);
 *         for (var i = 0; i &lt; clusters.length(); i++) {
 *             new Detector(&#34;applicationDelay-&#34; + i, DetectorArgs.builder()            
 *                 .description(String.format(&#34;your application is slow - %s&#34;, clusters[range.value()]))
 *                 .maxDelay(30)
 *                 .tags(                
 *                     &#34;app-backend&#34;,
 *                     &#34;staging&#34;)
 *                 .authorizedWriterTeams(signalfx_team.mycoolteam().id())
 *                 .authorizedWriterUsers(&#34;abc123&#34;)
 *                 .programText(&#34;&#34;&#34;
 * signal = data(&#39;app.delay&#39;, filter(&#39;cluster&#39;,&#39;%s&#39;), extrapolation=&#39;last_value&#39;, maxExtrapolations=5).max()
 * detect(when(signal &gt; 60, &#39;5m&#39;)).publish(&#39;Processing old messages 5m&#39;)
 * detect(when(signal &gt; 60, &#39;30m&#39;)).publish(&#39;Processing old messages 30m&#39;)
 * &#34;, clusters[range.value()]))
 *                 .rules(                
 *                     DetectorRuleArgs.builder()
 *                         .description(&#34;maximum &gt; 60 for 5m&#34;)
 *                         .severity(&#34;Warning&#34;)
 *                         .detectLabel(&#34;Processing old messages 5m&#34;)
 *                         .notifications(&#34;Email,foo-alerts@bar.com&#34;)
 *                         .build(),
 *                     DetectorRuleArgs.builder()
 *                         .description(&#34;maximum &gt; 60 for 30m&#34;)
 *                         .severity(&#34;Critical&#34;)
 *                         .detectLabel(&#34;Processing old messages 30m&#34;)
 *                         .notifications(&#34;Email,foo-alerts@bar.com&#34;)
 *                         .build())
 *                 .build());
 * 
 *         
 * }
 *     }
 * }
 * ```
 * 
 * ## Notification format
 * 
 * As Splunk Observability Cloud supports different notification mechanisms, use a comma-delimited string to provide inputs. If you want to specify multiple notifications, each must be a member in the list, like so:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * See [Splunk Observability Cloud Docs](https://dev.splunk.com/observability/reference/api/detectors/latest) for more information.
 * 
 * Here are some example of how to configure each notification type:
 * 
 * ### Email
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * ### Jira
 * 
 * Note that the `credentialId` is the Splunk-provided ID shown after setting up your Jira integration. See also `signalfx.jira.Integration`.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * ### OpsGenie
 * 
 * Note that the `credentialId` is the Splunk-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * ### PagerDuty
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * ### Slack
 * 
 * Exclude the `#` on the channel name:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * ### Team
 * 
 * Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * ### TeamEmail
 * 
 * Sends an email to every member of a team.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * ### Splunk On-Call (formerly VictorOps)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * ### Webhooks
 * 
 * You need to include all the commas even if you only use a credential id.
 * 
 * You can either configure a Webhook to use an existing integration&#39;s credential id:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * Or configure one inline:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *     }
 * }
 * ```
 * 
 * ## Arguments
 * 
 * * `name` - (Required) Name of the detector.
 * * `program_text` - (Required) Signalflow program text for the detector. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
 * * `description` - (Optional) Description of the detector.
 * * `authorized_writer_teams` - (Optional) Team IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team id (or user id in `authorized_writer_users`).
 * * `authorized_writer_users` - (Optional) User IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`).
 * * `max_delay` - (Optional) How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://docs.splunk.com/observability/en/data-visualization/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
 * * `min_delay` - (Optional) How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
 * * `show_data_markers` - (Optional) When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
 * * `show_event_lines` - (Optional) When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
 * * `disable_sampling` - (Optional) When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
 * * `time_range` - (Optional) Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
 * * `start_time` - (Optional) Seconds since epoch. Used for visualization. Conflicts with `time_range`.
 * * `end_time` - (Optional) Seconds since epoch. Used for visualization. Conflicts with `time_range`.
 * * `tags` - (Optional) Tags associated with the detector.
 * * `teams` - (Optional) Team IDs to associate the detector to.
 * * `rule` - (Required) Set of rules used for alerting.
 *     * `detect_label` - (Required) A detect label which matches a detect label within `program_text`.
 *     * `severity` - (Required) The severity of the rule, must be one of: `&#34;Critical&#34;`, `&#34;Major&#34;`, `&#34;Minor&#34;`, `&#34;Warning&#34;`, `&#34;Info&#34;`.
 *     * `description` - (Optional) Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
 *     * `disabled` - (Optional) When true, notifications and events will not be generated for the detect label. `false` by default.
 *     * `notifications` - (Optional) List of strings specifying where notifications will be sent when an incident occurs. See [Create A Single Detector](https://dev.splunk.com/observability/reference/api/detectors/latest) for more info.
 *     * `parameterized_body` - (Optional) Custom notification message body when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html) for more info.
 *     * `parameterized_subject` - (Optional) Custom notification message subject when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html) for more info.
 *     * `runbook_url` - (Optional) URL of page to consult when an alert is triggered. This can be used with custom notification messages.
 *     * `tip` - (Optional) Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
 * * `viz_options` - (Optional) Plot-level customization options, associated with a publish statement.
 *     * `label` - (Required) Label used in the publish statement that displays the plot (metric time series data) you want to customize.
 *     * `display_name` - (Optional) Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
 *     * `color` - (Optional) Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
 *     * `value_unit` - (Optional) A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gibibyte (note: this was previously typoed as Gigibyte), Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
 *     * `value_prefix`, `value_suffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
 * 
 * **Notes**
 * 
 * Use both `max_delay` in your detector configuration and an `extrapolation` policy in your program text to reduce false positives and false negatives.
 * 
 * - `max_delay` allows Splunk Observability Cloud to continue with computation if there is a lag in receiving data points.
 * - `extrapolation` allows you to specify how to handle missing data. An extrapolation policy can be added to individual signals by updating the data block in your `program_text`.
 * 
 * See [Delayed Datapoints](https://docs.splunk.com/observability/en/data-visualization/charts/chart-builder.html#delayed-datapoints) for more info.
 * 
 * ## Attributes
 * 
 * In a addition to all arguments above, the following attributes are exported:
 * 
 * * `id` - The ID of the detector.
 * * `label_resolutions` - The resolutions of the detector alerts in milliseconds that indicate how often data is analyzed to determine if an alert should be triggered.
 * * `url` - The URL of the detector.
 * 
 * ## Import
 * 
 * Detectors can be imported using their string ID (recoverable from URL: `/#/detector/v2/abc123/edit`, e.g.
 * 
 * ```sh
 *  $ pulumi import signalfx:index/detector:Detector application_delay abc123
 * ```
 * 
 */
@ResourceType(type="signalfx:index/detector:Detector")
public class Detector extends com.pulumi.resources.CustomResource {
    /**
     * Team IDs that have write access to this dashboard
     * 
     */
    @Export(name="authorizedWriterTeams", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> authorizedWriterTeams;

    /**
     * @return Team IDs that have write access to this dashboard
     * 
     */
    public Output<Optional<List<String>>> authorizedWriterTeams() {
        return Codegen.optional(this.authorizedWriterTeams);
    }
    /**
     * User IDs that have write access to this dashboard
     * 
     */
    @Export(name="authorizedWriterUsers", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> authorizedWriterUsers;

    /**
     * @return User IDs that have write access to this dashboard
     * 
     */
    public Output<Optional<List<String>>> authorizedWriterUsers() {
        return Codegen.optional(this.authorizedWriterUsers);
    }
    /**
     * Description of the detector
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the detector
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * (false by default) When false, samples a subset of the output MTS in the visualization.
     * 
     */
    @Export(name="disableSampling", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableSampling;

    /**
     * @return (false by default) When false, samples a subset of the output MTS in the visualization.
     * 
     */
    public Output<Optional<Boolean>> disableSampling() {
        return Codegen.optional(this.disableSampling);
    }
    /**
     * Seconds since epoch. Used for visualization
     * 
     */
    @Export(name="endTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> endTime;

    /**
     * @return Seconds since epoch. Used for visualization
     * 
     */
    public Output<Optional<Integer>> endTime() {
        return Codegen.optional(this.endTime);
    }
    /**
     * Resolutions of the detector alerts in milliseconds that indicate how often data is analyzed to determine if an alert
     * should be triggered
     * 
     */
    @Export(name="labelResolutions", refs={Map.class,String.class,Integer.class}, tree="[0,1,2]")
    private Output<Map<String,Integer>> labelResolutions;

    /**
     * @return Resolutions of the detector alerts in milliseconds that indicate how often data is analyzed to determine if an alert
     * should be triggered
     * 
     */
    public Output<Map<String,Integer>> labelResolutions() {
        return this.labelResolutions;
    }
    /**
     * Maximum time (in seconds) to wait for late datapoints. Max value is 900 (15m)
     * 
     */
    @Export(name="maxDelay", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxDelay;

    /**
     * @return Maximum time (in seconds) to wait for late datapoints. Max value is 900 (15m)
     * 
     */
    public Output<Optional<Integer>> maxDelay() {
        return Codegen.optional(this.maxDelay);
    }
    /**
     * Minimum time (in seconds) for the computation to wait even if the datapoints are arriving in a timely fashion. Max value
     * is 900 (15m)
     * 
     */
    @Export(name="minDelay", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minDelay;

    /**
     * @return Minimum time (in seconds) for the computation to wait even if the datapoints are arriving in a timely fashion. Max value
     * is 900 (15m)
     * 
     */
    public Output<Optional<Integer>> minDelay() {
        return Codegen.optional(this.minDelay);
    }
    /**
     * Name of the detector
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the detector
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Signalflow program text for the detector. More info at &#34;https://developers.signalfx.com/docs/signalflow-overview&#34;
     * 
     */
    @Export(name="programText", refs={String.class}, tree="[0]")
    private Output<String> programText;

    /**
     * @return Signalflow program text for the detector. More info at &#34;https://developers.signalfx.com/docs/signalflow-overview&#34;
     * 
     */
    public Output<String> programText() {
        return this.programText;
    }
    /**
     * Set of rules used for alerting
     * 
     */
    @Export(name="rules", refs={List.class,DetectorRule.class}, tree="[0,1]")
    private Output<List<DetectorRule>> rules;

    /**
     * @return Set of rules used for alerting
     * 
     */
    public Output<List<DetectorRule>> rules() {
        return this.rules;
    }
    /**
     * (true by default) When true, markers will be drawn for each datapoint within the visualization.
     * 
     */
    @Export(name="showDataMarkers", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> showDataMarkers;

    /**
     * @return (true by default) When true, markers will be drawn for each datapoint within the visualization.
     * 
     */
    public Output<Optional<Boolean>> showDataMarkers() {
        return Codegen.optional(this.showDataMarkers);
    }
    /**
     * (false by default) When true, vertical lines will be drawn for each triggered event within the visualization.
     * 
     */
    @Export(name="showEventLines", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> showEventLines;

    /**
     * @return (false by default) When true, vertical lines will be drawn for each triggered event within the visualization.
     * 
     */
    public Output<Optional<Boolean>> showEventLines() {
        return Codegen.optional(this.showEventLines);
    }
    /**
     * Seconds since epoch. Used for visualization
     * 
     */
    @Export(name="startTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> startTime;

    /**
     * @return Seconds since epoch. Used for visualization
     * 
     */
    public Output<Optional<Integer>> startTime() {
        return Codegen.optional(this.startTime);
    }
    /**
     * Tags associated with the detector
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return Tags associated with the detector
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Team IDs to associate the detector to
     * 
     */
    @Export(name="teams", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> teams;

    /**
     * @return Team IDs to associate the detector to
     * 
     */
    public Output<Optional<List<String>>> teams() {
        return Codegen.optional(this.teams);
    }
    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`. Defaults
     * to 3600
     * 
     */
    @Export(name="timeRange", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeRange;

    /**
     * @return Seconds to display in the visualization. This is a rolling range from the current time. Example: 3600 = `-1h`. Defaults
     * to 3600
     * 
     */
    public Output<Optional<Integer>> timeRange() {
        return Codegen.optional(this.timeRange);
    }
    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     * 
     */
    @Export(name="timezone", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> timezone;

    /**
     * @return The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     * 
     */
    public Output<Optional<String>> timezone() {
        return Codegen.optional(this.timezone);
    }
    /**
     * URL of the detector
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return URL of the detector
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * Plot-level customization options, associated with a publish statement
     * 
     */
    @Export(name="vizOptions", refs={List.class,DetectorVizOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DetectorVizOption>> vizOptions;

    /**
     * @return Plot-level customization options, associated with a publish statement
     * 
     */
    public Output<Optional<List<DetectorVizOption>>> vizOptions() {
        return Codegen.optional(this.vizOptions);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Detector(String name) {
        this(name, DetectorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Detector(String name, DetectorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Detector(String name, DetectorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/detector:Detector", name, args == null ? DetectorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Detector(String name, Output<String> id, @Nullable DetectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/detector:Detector", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Detector get(String name, Output<String> id, @Nullable DetectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Detector(name, id, state, options);
    }
}
