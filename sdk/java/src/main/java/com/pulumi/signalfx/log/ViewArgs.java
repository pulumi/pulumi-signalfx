// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.log;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.log.inputs.ViewColumnArgs;
import com.pulumi.signalfx.log.inputs.ViewSortOptionArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ViewArgs extends com.pulumi.resources.ResourceArgs {

    public static final ViewArgs Empty = new ViewArgs();

    /**
     * The column headers to show on the log view.
     * 
     */
    @Import(name="columns")
    private @Nullable Output<List<ViewColumnArgs>> columns;

    /**
     * @return The column headers to show on the log view.
     * 
     */
    public Optional<Output<List<ViewColumnArgs>>> columns() {
        return Optional.ofNullable(this.columns);
    }

    /**
     * The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
     * 
     */
    @Import(name="defaultConnection")
    private @Nullable Output<String> defaultConnection;

    /**
     * @return The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
     * 
     */
    public Optional<Output<String>> defaultConnection() {
        return Optional.ofNullable(this.defaultConnection);
    }

    /**
     * Description of the log view.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the log view.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
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
     * Name of the log view.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the log view.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
     * 
     */
    @Import(name="programText", required=true)
    private Output<String> programText;

    /**
     * @return Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
     * 
     */
    public Output<String> programText() {
        return this.programText;
    }

    /**
     * The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
     * 
     */
    @Import(name="sortOptions")
    private @Nullable Output<List<ViewSortOptionArgs>> sortOptions;

    /**
     * @return The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
     * 
     */
    public Optional<Output<List<ViewSortOptionArgs>>> sortOptions() {
        return Optional.ofNullable(this.sortOptions);
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
     * Tags associated with the resource
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return Tags associated with the resource
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * From when to display data. Splunk Observability Cloud time syntax (e.g. `&#34;-5m&#34;`, `&#34;-1h&#34;`). Conflicts with `start_time` and `end_time`.
     * 
     */
    @Import(name="timeRange")
    private @Nullable Output<Integer> timeRange;

    /**
     * @return From when to display data. Splunk Observability Cloud time syntax (e.g. `&#34;-5m&#34;`, `&#34;-1h&#34;`). Conflicts with `start_time` and `end_time`.
     * 
     */
    public Optional<Output<Integer>> timeRange() {
        return Optional.ofNullable(this.timeRange);
    }

    private ViewArgs() {}

    private ViewArgs(ViewArgs $) {
        this.columns = $.columns;
        this.defaultConnection = $.defaultConnection;
        this.description = $.description;
        this.endTime = $.endTime;
        this.name = $.name;
        this.programText = $.programText;
        this.sortOptions = $.sortOptions;
        this.startTime = $.startTime;
        this.tags = $.tags;
        this.timeRange = $.timeRange;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ViewArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ViewArgs $;

        public Builder() {
            $ = new ViewArgs();
        }

        public Builder(ViewArgs defaults) {
            $ = new ViewArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param columns The column headers to show on the log view.
         * 
         * @return builder
         * 
         */
        public Builder columns(@Nullable Output<List<ViewColumnArgs>> columns) {
            $.columns = columns;
            return this;
        }

        /**
         * @param columns The column headers to show on the log view.
         * 
         * @return builder
         * 
         */
        public Builder columns(List<ViewColumnArgs> columns) {
            return columns(Output.of(columns));
        }

        /**
         * @param columns The column headers to show on the log view.
         * 
         * @return builder
         * 
         */
        public Builder columns(ViewColumnArgs... columns) {
            return columns(List.of(columns));
        }

        /**
         * @param defaultConnection The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
         * 
         * @return builder
         * 
         */
        public Builder defaultConnection(@Nullable Output<String> defaultConnection) {
            $.defaultConnection = defaultConnection;
            return this;
        }

        /**
         * @param defaultConnection The connection that the log view uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
         * 
         * @return builder
         * 
         */
        public Builder defaultConnection(String defaultConnection) {
            return defaultConnection(Output.of(defaultConnection));
        }

        /**
         * @param description Description of the log view.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the log view.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
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
         * @param name Name of the log view.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the log view.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param programText Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
         * 
         * @return builder
         * 
         */
        public Builder programText(Output<String> programText) {
            $.programText = programText;
            return this;
        }

        /**
         * @param programText Signalflow program text for the log view. More info at https://developers.signalfx.com/docs/signalflow-overview.
         * 
         * @return builder
         * 
         */
        public Builder programText(String programText) {
            return programText(Output.of(programText));
        }

        /**
         * @param sortOptions The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
         * 
         * @return builder
         * 
         */
        public Builder sortOptions(@Nullable Output<List<ViewSortOptionArgs>> sortOptions) {
            $.sortOptions = sortOptions;
            return this;
        }

        /**
         * @param sortOptions The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
         * 
         * @return builder
         * 
         */
        public Builder sortOptions(List<ViewSortOptionArgs> sortOptions) {
            return sortOptions(Output.of(sortOptions));
        }

        /**
         * @param sortOptions The sorting options configuration to specify if the log view table needs to be sorted in a particular field.
         * 
         * @return builder
         * 
         */
        public Builder sortOptions(ViewSortOptionArgs... sortOptions) {
            return sortOptions(List.of(sortOptions));
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
         * @param tags Tags associated with the resource
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags associated with the resource
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags Tags associated with the resource
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param timeRange From when to display data. Splunk Observability Cloud time syntax (e.g. `&#34;-5m&#34;`, `&#34;-1h&#34;`). Conflicts with `start_time` and `end_time`.
         * 
         * @return builder
         * 
         */
        public Builder timeRange(@Nullable Output<Integer> timeRange) {
            $.timeRange = timeRange;
            return this;
        }

        /**
         * @param timeRange From when to display data. Splunk Observability Cloud time syntax (e.g. `&#34;-5m&#34;`, `&#34;-1h&#34;`). Conflicts with `start_time` and `end_time`.
         * 
         * @return builder
         * 
         */
        public Builder timeRange(Integer timeRange) {
            return timeRange(Output.of(timeRange));
        }

        public ViewArgs build() {
            if ($.programText == null) {
                throw new MissingRequiredPropertyException("ViewArgs", "programText");
            }
            return $;
        }
    }

}
