// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.inputs.DashboardSelectedEventOverlaySourceArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardSelectedEventOverlayArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardSelectedEventOverlayArgs Empty = new DashboardSelectedEventOverlayArgs();

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
    private @Nullable Output<List<DashboardSelectedEventOverlaySourceArgs>> sources;

    /**
     * @return Each element specifies a filter to use against the signal specified in the `signal`.
     * 
     */
    public Optional<Output<List<DashboardSelectedEventOverlaySourceArgs>>> sources() {
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

    private DashboardSelectedEventOverlayArgs() {}

    private DashboardSelectedEventOverlayArgs(DashboardSelectedEventOverlayArgs $) {
        this.signal = $.signal;
        this.sources = $.sources;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardSelectedEventOverlayArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardSelectedEventOverlayArgs $;

        public Builder() {
            $ = new DashboardSelectedEventOverlayArgs();
        }

        public Builder(DashboardSelectedEventOverlayArgs defaults) {
            $ = new DashboardSelectedEventOverlayArgs(Objects.requireNonNull(defaults));
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
        public Builder sources(@Nullable Output<List<DashboardSelectedEventOverlaySourceArgs>> sources) {
            $.sources = sources;
            return this;
        }

        /**
         * @param sources Each element specifies a filter to use against the signal specified in the `signal`.
         * 
         * @return builder
         * 
         */
        public Builder sources(List<DashboardSelectedEventOverlaySourceArgs> sources) {
            return sources(Output.of(sources));
        }

        /**
         * @param sources Each element specifies a filter to use against the signal specified in the `signal`.
         * 
         * @return builder
         * 
         */
        public Builder sources(DashboardSelectedEventOverlaySourceArgs... sources) {
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

        public DashboardSelectedEventOverlayArgs build() {
            $.signal = Objects.requireNonNull($.signal, "expected parameter 'signal' to be non-null");
            return $;
        }
    }

}