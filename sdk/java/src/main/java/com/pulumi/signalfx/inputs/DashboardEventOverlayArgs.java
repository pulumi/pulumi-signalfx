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
     * Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
     * 
     */
    @Import(name="color")
    private @Nullable Output<String> color;

    /**
     * @return Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
     * 
     */
    public Optional<Output<String>> color() {
        return Optional.ofNullable(this.color);
    }

    /**
     * Text shown in the dropdown when selecting this overlay from the menu.
     * 
     */
    @Import(name="label")
    private @Nullable Output<String> label;

    /**
     * @return Text shown in the dropdown when selecting this overlay from the menu.
     * 
     */
    public Optional<Output<String>> label() {
        return Optional.ofNullable(this.label);
    }

    /**
     * Show a vertical line for the event. `false` by default.
     * 
     */
    @Import(name="line")
    private @Nullable Output<Boolean> line;

    /**
     * @return Show a vertical line for the event. `false` by default.
     * 
     */
    public Optional<Output<Boolean>> line() {
        return Optional.ofNullable(this.line);
    }

    /**
     * Search term used to choose the events shown in the overlay.
     * 
     */
    @Import(name="signal", required=true)
    private Output<String> signal;

    /**
     * @return Search term used to choose the events shown in the overlay.
     * 
     */
    public Output<String> signal() {
        return this.signal;
    }

    /**
     * Each element specifies a filter to use against the signal specified in the `signal`.
     * 
     */
    @Import(name="sources")
    private @Nullable Output<List<DashboardEventOverlaySourceArgs>> sources;

    /**
     * @return Each element specifies a filter to use against the signal specified in the `signal`.
     * 
     */
    public Optional<Output<List<DashboardEventOverlaySourceArgs>>> sources() {
        return Optional.ofNullable(this.sources);
    }

    /**
     * Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
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
         * @param color Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
         * 
         * @return builder
         * 
         */
        public Builder color(@Nullable Output<String> color) {
            $.color = color;
            return this;
        }

        /**
         * @param color Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
         * 
         * @return builder
         * 
         */
        public Builder color(String color) {
            return color(Output.of(color));
        }

        /**
         * @param label Text shown in the dropdown when selecting this overlay from the menu.
         * 
         * @return builder
         * 
         */
        public Builder label(@Nullable Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label Text shown in the dropdown when selecting this overlay from the menu.
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param line Show a vertical line for the event. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder line(@Nullable Output<Boolean> line) {
            $.line = line;
            return this;
        }

        /**
         * @param line Show a vertical line for the event. `false` by default.
         * 
         * @return builder
         * 
         */
        public Builder line(Boolean line) {
            return line(Output.of(line));
        }

        /**
         * @param signal Search term used to choose the events shown in the overlay.
         * 
         * @return builder
         * 
         */
        public Builder signal(Output<String> signal) {
            $.signal = signal;
            return this;
        }

        /**
         * @param signal Search term used to choose the events shown in the overlay.
         * 
         * @return builder
         * 
         */
        public Builder signal(String signal) {
            return signal(Output.of(signal));
        }

        /**
         * @param sources Each element specifies a filter to use against the signal specified in the `signal`.
         * 
         * @return builder
         * 
         */
        public Builder sources(@Nullable Output<List<DashboardEventOverlaySourceArgs>> sources) {
            $.sources = sources;
            return this;
        }

        /**
         * @param sources Each element specifies a filter to use against the signal specified in the `signal`.
         * 
         * @return builder
         * 
         */
        public Builder sources(List<DashboardEventOverlaySourceArgs> sources) {
            return sources(Output.of(sources));
        }

        /**
         * @param sources Each element specifies a filter to use against the signal specified in the `signal`.
         * 
         * @return builder
         * 
         */
        public Builder sources(DashboardEventOverlaySourceArgs... sources) {
            return sources(List.of(sources));
        }

        /**
         * @param type Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
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
