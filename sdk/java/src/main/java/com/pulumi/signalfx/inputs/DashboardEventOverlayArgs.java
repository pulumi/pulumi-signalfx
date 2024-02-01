// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.inputs.DashboardEventOverlaySourceArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardEventOverlayArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardEventOverlayArgs Empty = new DashboardEventOverlayArgs();

    /**
     * Color to use
     * 
     */
    @Import(name="color")
    private @Nullable Output<String> color;

    /**
     * @return Color to use
     * 
     */
    public Optional<Output<String>> color() {
        return Optional.ofNullable(this.color);
    }

    /**
     * The text displaying in the dropdown menu used to select this event overlay as an active overlay for the dashboard.
     * 
     */
    @Import(name="label")
    private @Nullable Output<String> label;

    /**
     * @return The text displaying in the dropdown menu used to select this event overlay as an active overlay for the dashboard.
     * 
     */
    public Optional<Output<String>> label() {
        return Optional.ofNullable(this.label);
    }

    /**
     * (false by default) Whether a vertical line should be displayed in the plot at the time the event occurs
     * 
     */
    @Import(name="line")
    private @Nullable Output<Boolean> line;

    /**
     * @return (false by default) Whether a vertical line should be displayed in the plot at the time the event occurs
     * 
     */
    public Optional<Output<Boolean>> line() {
        return Optional.ofNullable(this.line);
    }

    /**
     * Search term used to define events
     * 
     */
    @Import(name="signal", required=true)
    private Output<String> signal;

    /**
     * @return Search term used to define events
     * 
     */
    public Output<String> signal() {
        return this.signal;
    }

    @Import(name="sources")
    private @Nullable Output<List<DashboardEventOverlaySourceArgs>> sources;

    public Optional<Output<List<DashboardEventOverlaySourceArgs>>> sources() {
        return Optional.ofNullable(this.sources);
    }

    /**
     * Source for this event&#39;s data. Can be &#34;eventTimeSeries&#34; (default) or &#34;detectorEvents&#34;.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Source for this event&#39;s data. Can be &#34;eventTimeSeries&#34; (default) or &#34;detectorEvents&#34;.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private DashboardEventOverlayArgs() {}

    private DashboardEventOverlayArgs(DashboardEventOverlayArgs $) {
        this.color = $.color;
        this.label = $.label;
        this.line = $.line;
        this.signal = $.signal;
        this.sources = $.sources;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardEventOverlayArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardEventOverlayArgs $;

        public Builder() {
            $ = new DashboardEventOverlayArgs();
        }

        public Builder(DashboardEventOverlayArgs defaults) {
            $ = new DashboardEventOverlayArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param color Color to use
         * 
         * @return builder
         * 
         */
        public Builder color(@Nullable Output<String> color) {
            $.color = color;
            return this;
        }

        /**
         * @param color Color to use
         * 
         * @return builder
         * 
         */
        public Builder color(String color) {
            return color(Output.of(color));
        }

        /**
         * @param label The text displaying in the dropdown menu used to select this event overlay as an active overlay for the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder label(@Nullable Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label The text displaying in the dropdown menu used to select this event overlay as an active overlay for the dashboard.
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param line (false by default) Whether a vertical line should be displayed in the plot at the time the event occurs
         * 
         * @return builder
         * 
         */
        public Builder line(@Nullable Output<Boolean> line) {
            $.line = line;
            return this;
        }

        /**
         * @param line (false by default) Whether a vertical line should be displayed in the plot at the time the event occurs
         * 
         * @return builder
         * 
         */
        public Builder line(Boolean line) {
            return line(Output.of(line));
        }

        /**
         * @param signal Search term used to define events
         * 
         * @return builder
         * 
         */
        public Builder signal(Output<String> signal) {
            $.signal = signal;
            return this;
        }

        /**
         * @param signal Search term used to define events
         * 
         * @return builder
         * 
         */
        public Builder signal(String signal) {
            return signal(Output.of(signal));
        }

        public Builder sources(@Nullable Output<List<DashboardEventOverlaySourceArgs>> sources) {
            $.sources = sources;
            return this;
        }

        public Builder sources(List<DashboardEventOverlaySourceArgs> sources) {
            return sources(Output.of(sources));
        }

        public Builder sources(DashboardEventOverlaySourceArgs... sources) {
            return sources(List.of(sources));
        }

        /**
         * @param type Source for this event&#39;s data. Can be &#34;eventTimeSeries&#34; (default) or &#34;detectorEvents&#34;.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Source for this event&#39;s data. Can be &#34;eventTimeSeries&#34; (default) or &#34;detectorEvents&#34;.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public DashboardEventOverlayArgs build() {
            if ($.signal == null) {
                throw new MissingRequiredPropertyException("DashboardEventOverlayArgs", "signal");
            }
            return $;
        }
    }

}
