// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.signalfx.inputs.DashboardSelectedEventOverlaySourceArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardSelectedEventOverlayArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardSelectedEventOverlayArgs Empty = new DashboardSelectedEventOverlayArgs();

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
    private @Nullable Output<List<DashboardSelectedEventOverlaySourceArgs>> sources;

    public Optional<Output<List<DashboardSelectedEventOverlaySourceArgs>>> sources() {
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

        public Builder sources(@Nullable Output<List<DashboardSelectedEventOverlaySourceArgs>> sources) {
            $.sources = sources;
            return this;
        }

        public Builder sources(List<DashboardSelectedEventOverlaySourceArgs> sources) {
            return sources(Output.of(sources));
        }

        public Builder sources(DashboardSelectedEventOverlaySourceArgs... sources) {
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

        public DashboardSelectedEventOverlayArgs build() {
            if ($.signal == null) {
                throw new MissingRequiredPropertyException("DashboardSelectedEventOverlayArgs", "signal");
            }
            return $;
        }
    }

}
