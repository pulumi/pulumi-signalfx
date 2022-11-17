// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.signalfx.outputs.DashboardSelectedEventOverlaySource;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DashboardSelectedEventOverlay {
    /**
     * @return Search term used to choose the events shown in the overlay.
     * 
     */
    private String signal;
    /**
     * @return Each element specifies a filter to use against the signal specified in the `signal`.
     * 
     */
    private @Nullable List<DashboardSelectedEventOverlaySource> sources;
    /**
     * @return Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
     * 
     */
    private @Nullable String type;

    private DashboardSelectedEventOverlay() {}
    /**
     * @return Search term used to choose the events shown in the overlay.
     * 
     */
    public String signal() {
        return this.signal;
    }
    /**
     * @return Each element specifies a filter to use against the signal specified in the `signal`.
     * 
     */
    public List<DashboardSelectedEventOverlaySource> sources() {
        return this.sources == null ? List.of() : this.sources;
    }
    /**
     * @return Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardSelectedEventOverlay defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String signal;
        private @Nullable List<DashboardSelectedEventOverlaySource> sources;
        private @Nullable String type;
        public Builder() {}
        public Builder(DashboardSelectedEventOverlay defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.signal = defaults.signal;
    	      this.sources = defaults.sources;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder signal(String signal) {
            this.signal = Objects.requireNonNull(signal);
            return this;
        }
        @CustomType.Setter
        public Builder sources(@Nullable List<DashboardSelectedEventOverlaySource> sources) {
            this.sources = sources;
            return this;
        }
        public Builder sources(DashboardSelectedEventOverlaySource... sources) {
            return sources(List.of(sources));
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        public DashboardSelectedEventOverlay build() {
            final var o = new DashboardSelectedEventOverlay();
            o.signal = signal;
            o.sources = sources;
            o.type = type;
            return o;
        }
    }
}
