// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.inputs.DetectorRuleArgs;
import com.pulumi.signalfx.inputs.DetectorVizOptionArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DetectorArgs extends com.pulumi.resources.ResourceArgs {

    public static final DetectorArgs Empty = new DetectorArgs();

    /**
     * Team IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team id (or user id in `authorized_writer_users`).
     * 
     */
    @Import(name="authorizedWriterTeams")
    private @Nullable Output<List<String>> authorizedWriterTeams;

    /**
     * @return Team IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team id (or user id in `authorized_writer_users`).
     * 
     */
    public Optional<Output<List<String>>> authorizedWriterTeams() {
        return Optional.ofNullable(this.authorizedWriterTeams);
    }

    /**
     * User IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`).
     * 
     */
    @Import(name="authorizedWriterUsers")
    private @Nullable Output<List<String>> authorizedWriterUsers;

    /**
     * @return User IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`).
     * 
     */
    public Optional<Output<List<String>>> authorizedWriterUsers() {
        return Optional.ofNullable(this.authorizedWriterUsers);
    }

    /**
     * Description of the detector.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the detector.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
     * 
     */
    @Import(name="disableSampling")
    private @Nullable Output<Boolean> disableSampling;

    /**
     * @return When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
     * 
     */
    public Optional<Output<Boolean>> disableSampling() {
        return Optional.ofNullable(this.disableSampling);
    }

    /**
     * Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    @Import(name="endTime")
    private @Nullable Output<Integer> endTime;

    /**
     * @return Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    public Optional<Output<Integer>> endTime() {
        return Optional.ofNullable(this.endTime);
    }

    /**
     * allows Splunk Observability Cloud to continue with computation if there is a lag in receiving data points.
     * 
     */
    @Import(name="maxDelay")
    private @Nullable Output<Integer> maxDelay;

    /**
     * @return allows Splunk Observability Cloud to continue with computation if there is a lag in receiving data points.
     * 
     */
    public Optional<Output<Integer>> maxDelay() {
        return Optional.ofNullable(this.maxDelay);
    }

    /**
     * How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
     * 
     */
    @Import(name="minDelay")
    private @Nullable Output<Integer> minDelay;

    /**
     * @return How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
     * 
     */
    public Optional<Output<Integer>> minDelay() {
        return Optional.ofNullable(this.minDelay);
    }

    /**
     * Name of the detector.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the detector.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Signalflow program text for the detector. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
     * 
     */
    @Import(name="programText", required=true)
    private Output<String> programText;

    /**
     * @return Signalflow program text for the detector. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
     * 
     */
    public Output<String> programText() {
        return this.programText;
    }

    /**
     * Set of rules used for alerting.
     * 
     */
    @Import(name="rules", required=true)
    private Output<List<DetectorRuleArgs>> rules;

    /**
     * @return Set of rules used for alerting.
     * 
     */
    public Output<List<DetectorRuleArgs>> rules() {
        return this.rules;
    }

    /**
     * When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
     * 
     */
    @Import(name="showDataMarkers")
    private @Nullable Output<Boolean> showDataMarkers;

    /**
     * @return When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
     * 
     */
    public Optional<Output<Boolean>> showDataMarkers() {
        return Optional.ofNullable(this.showDataMarkers);
    }

    /**
     * When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
     * 
     */
    @Import(name="showEventLines")
    private @Nullable Output<Boolean> showEventLines;

    /**
     * @return When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
     * 
     */
    public Optional<Output<Boolean>> showEventLines() {
        return Optional.ofNullable(this.showEventLines);
    }

    /**
     * Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    @Import(name="startTime")
    private @Nullable Output<Integer> startTime;

    /**
     * @return Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    public Optional<Output<Integer>> startTime() {
        return Optional.ofNullable(this.startTime);
    }

    /**
     * Tags associated with the detector.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return Tags associated with the detector.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Team IDs to associate the detector to.
     * 
     */
    @Import(name="teams")
    private @Nullable Output<List<String>> teams;

    /**
     * @return Team IDs to associate the detector to.
     * 
     */
    public Optional<Output<List<String>>> teams() {
        return Optional.ofNullable(this.teams);
    }

    /**
     * Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
     * 
     */
    @Import(name="timeRange")
    private @Nullable Output<Integer> timeRange;

    /**
     * @return Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
     * 
     */
    public Optional<Output<Integer>> timeRange() {
        return Optional.ofNullable(this.timeRange);
    }

    /**
     * The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     * 
     */
    @Import(name="timezone")
    private @Nullable Output<String> timezone;

    /**
     * @return The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
     * 
     */
    public Optional<Output<String>> timezone() {
        return Optional.ofNullable(this.timezone);
    }

    /**
     * Plot-level customization options, associated with a publish statement.
     * 
     */
    @Import(name="vizOptions")
    private @Nullable Output<List<DetectorVizOptionArgs>> vizOptions;

    /**
     * @return Plot-level customization options, associated with a publish statement.
     * 
     */
    public Optional<Output<List<DetectorVizOptionArgs>>> vizOptions() {
        return Optional.ofNullable(this.vizOptions);
    }

    private DetectorArgs() {}

    private DetectorArgs(DetectorArgs $) {
        this.authorizedWriterTeams = $.authorizedWriterTeams;
        this.authorizedWriterUsers = $.authorizedWriterUsers;
        this.description = $.description;
        this.disableSampling = $.disableSampling;
        this.endTime = $.endTime;
        this.maxDelay = $.maxDelay;
        this.minDelay = $.minDelay;
        this.name = $.name;
        this.programText = $.programText;
        this.rules = $.rules;
        this.showDataMarkers = $.showDataMarkers;
        this.showEventLines = $.showEventLines;
        this.startTime = $.startTime;
        this.tags = $.tags;
        this.teams = $.teams;
        this.timeRange = $.timeRange;
        this.timezone = $.timezone;
        this.vizOptions = $.vizOptions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DetectorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DetectorArgs $;

        public Builder() {
            $ = new DetectorArgs();
        }

        public Builder(DetectorArgs defaults) {
            $ = new DetectorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authorizedWriterTeams Team IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team id (or user id in `authorized_writer_users`).
         * 
         * @return builder
         * 
         */
        public Builder authorizedWriterTeams(@Nullable Output<List<String>> authorizedWriterTeams) {
            $.authorizedWriterTeams = authorizedWriterTeams;
            return this;
        }

        /**
         * @param authorizedWriterTeams Team IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team id (or user id in `authorized_writer_users`).
         * 
         * @return builder
         * 
         */
        public Builder authorizedWriterTeams(List<String> authorizedWriterTeams) {
            return authorizedWriterTeams(Output.of(authorizedWriterTeams));
        }

        /**
         * @param authorizedWriterTeams Team IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team id (or user id in `authorized_writer_users`).
         * 
         * @return builder
         * 
         */
        public Builder authorizedWriterTeams(String... authorizedWriterTeams) {
            return authorizedWriterTeams(List.of(authorizedWriterTeams));
        }

        /**
         * @param authorizedWriterUsers User IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`).
         * 
         * @return builder
         * 
         */
        public Builder authorizedWriterUsers(@Nullable Output<List<String>> authorizedWriterUsers) {
            $.authorizedWriterUsers = authorizedWriterUsers;
            return this;
        }

        /**
         * @param authorizedWriterUsers User IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`).
         * 
         * @return builder
         * 
         */
        public Builder authorizedWriterUsers(List<String> authorizedWriterUsers) {
            return authorizedWriterUsers(Output.of(authorizedWriterUsers));
        }

        /**
         * @param authorizedWriterUsers User IDs that have write access to this detector. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`).
         * 
         * @return builder
         * 
         */
        public Builder authorizedWriterUsers(String... authorizedWriterUsers) {
            return authorizedWriterUsers(List.of(authorizedWriterUsers));
        }

        /**
         * @param description Description of the detector.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the detector.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disableSampling When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder disableSampling(@Nullable Output<Boolean> disableSampling) {
            $.disableSampling = disableSampling;
            return this;
        }

        /**
         * @param disableSampling When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder disableSampling(Boolean disableSampling) {
            return disableSampling(Output.of(disableSampling));
        }

        /**
         * @param endTime Seconds since epoch. Used for visualization. Conflicts with `time_range`.
         * 
         * @return builder
         * 
         */
        public Builder endTime(@Nullable Output<Integer> endTime) {
            $.endTime = endTime;
            return this;
        }

        /**
         * @param endTime Seconds since epoch. Used for visualization. Conflicts with `time_range`.
         * 
         * @return builder
         * 
         */
        public Builder endTime(Integer endTime) {
            return endTime(Output.of(endTime));
        }

        /**
         * @param maxDelay allows Splunk Observability Cloud to continue with computation if there is a lag in receiving data points.
         * 
         * @return builder
         * 
         */
        public Builder maxDelay(@Nullable Output<Integer> maxDelay) {
            $.maxDelay = maxDelay;
            return this;
        }

        /**
         * @param maxDelay allows Splunk Observability Cloud to continue with computation if there is a lag in receiving data points.
         * 
         * @return builder
         * 
         */
        public Builder maxDelay(Integer maxDelay) {
            return maxDelay(Output.of(maxDelay));
        }

        /**
         * @param minDelay How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
         * 
         * @return builder
         * 
         */
        public Builder minDelay(@Nullable Output<Integer> minDelay) {
            $.minDelay = minDelay;
            return this;
        }

        /**
         * @param minDelay How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
         * 
         * @return builder
         * 
         */
        public Builder minDelay(Integer minDelay) {
            return minDelay(Output.of(minDelay));
        }

        /**
         * @param name Name of the detector.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the detector.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param programText Signalflow program text for the detector. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
         * 
         * @return builder
         * 
         */
        public Builder programText(Output<String> programText) {
            $.programText = programText;
            return this;
        }

        /**
         * @param programText Signalflow program text for the detector. More info [in the Splunk Observability Cloud docs](https://dev.splunk.com/observability/docs/signalflow/).
         * 
         * @return builder
         * 
         */
        public Builder programText(String programText) {
            return programText(Output.of(programText));
        }

        /**
         * @param rules Set of rules used for alerting.
         * 
         * @return builder
         * 
         */
        public Builder rules(Output<List<DetectorRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules Set of rules used for alerting.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<DetectorRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules Set of rules used for alerting.
         * 
         * @return builder
         * 
         */
        public Builder rules(DetectorRuleArgs... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param showDataMarkers When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
         * 
         * @return builder
         * 
         */
        public Builder showDataMarkers(@Nullable Output<Boolean> showDataMarkers) {
            $.showDataMarkers = showDataMarkers;
            return this;
        }

        /**
         * @param showDataMarkers When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
         * 
         * @return builder
         * 
         */
        public Builder showDataMarkers(Boolean showDataMarkers) {
            return showDataMarkers(Output.of(showDataMarkers));
        }

        /**
         * @param showEventLines When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder showEventLines(@Nullable Output<Boolean> showEventLines) {
            $.showEventLines = showEventLines;
            return this;
        }

        /**
         * @param showEventLines When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder showEventLines(Boolean showEventLines) {
            return showEventLines(Output.of(showEventLines));
        }

        /**
         * @param startTime Seconds since epoch. Used for visualization. Conflicts with `time_range`.
         * 
         * @return builder
         * 
         */
        public Builder startTime(@Nullable Output<Integer> startTime) {
            $.startTime = startTime;
            return this;
        }

        /**
         * @param startTime Seconds since epoch. Used for visualization. Conflicts with `time_range`.
         * 
         * @return builder
         * 
         */
        public Builder startTime(Integer startTime) {
            return startTime(Output.of(startTime));
        }

        /**
         * @param tags Tags associated with the detector.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags associated with the detector.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags Tags associated with the detector.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param teams Team IDs to associate the detector to.
         * 
         * @return builder
         * 
         */
        public Builder teams(@Nullable Output<List<String>> teams) {
            $.teams = teams;
            return this;
        }

        /**
         * @param teams Team IDs to associate the detector to.
         * 
         * @return builder
         * 
         */
        public Builder teams(List<String> teams) {
            return teams(Output.of(teams));
        }

        /**
         * @param teams Team IDs to associate the detector to.
         * 
         * @return builder
         * 
         */
        public Builder teams(String... teams) {
            return teams(List.of(teams));
        }

        /**
         * @param timeRange Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
         * 
         * @return builder
         * 
         */
        public Builder timeRange(@Nullable Output<Integer> timeRange) {
            $.timeRange = timeRange;
            return this;
        }

        /**
         * @param timeRange Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
         * 
         * @return builder
         * 
         */
        public Builder timeRange(Integer timeRange) {
            return timeRange(Output.of(timeRange));
        }

        /**
         * @param timezone The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
         * 
         * @return builder
         * 
         */
        public Builder timezone(@Nullable Output<String> timezone) {
            $.timezone = timezone;
            return this;
        }

        /**
         * @param timezone The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
         * 
         * @return builder
         * 
         */
        public Builder timezone(String timezone) {
            return timezone(Output.of(timezone));
        }

        /**
         * @param vizOptions Plot-level customization options, associated with a publish statement.
         * 
         * @return builder
         * 
         */
        public Builder vizOptions(@Nullable Output<List<DetectorVizOptionArgs>> vizOptions) {
            $.vizOptions = vizOptions;
            return this;
        }

        /**
         * @param vizOptions Plot-level customization options, associated with a publish statement.
         * 
         * @return builder
         * 
         */
        public Builder vizOptions(List<DetectorVizOptionArgs> vizOptions) {
            return vizOptions(Output.of(vizOptions));
        }

        /**
         * @param vizOptions Plot-level customization options, associated with a publish statement.
         * 
         * @return builder
         * 
         */
        public Builder vizOptions(DetectorVizOptionArgs... vizOptions) {
            return vizOptions(List.of(vizOptions));
        }

        public DetectorArgs build() {
            if ($.programText == null) {
                throw new MissingRequiredPropertyException("DetectorArgs", "programText");
            }
            if ($.rules == null) {
                throw new MissingRequiredPropertyException("DetectorArgs", "rules");
            }
            return $;
        }
    }

}
