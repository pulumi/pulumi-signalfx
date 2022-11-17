// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ListChartColorScale {
    /**
     * @return The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     * 
     */
    private String color;
    /**
     * @return Indicates the lower threshold non-inclusive value for this range.
     * 
     */
    private @Nullable Double gt;
    /**
     * @return Indicates the lower threshold inclusive value for this range.
     * 
     */
    private @Nullable Double gte;
    /**
     * @return Indicates the upper threshold non-inculsive value for this range.
     * 
     */
    private @Nullable Double lt;
    /**
     * @return Indicates the upper threshold inclusive value for this range.
     * 
     */
    private @Nullable Double lte;

    private ListChartColorScale() {}
    /**
     * @return The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     * 
     */
    public String color() {
        return this.color;
    }
    /**
     * @return Indicates the lower threshold non-inclusive value for this range.
     * 
     */
    public Optional<Double> gt() {
        return Optional.ofNullable(this.gt);
    }
    /**
     * @return Indicates the lower threshold inclusive value for this range.
     * 
     */
    public Optional<Double> gte() {
        return Optional.ofNullable(this.gte);
    }
    /**
     * @return Indicates the upper threshold non-inculsive value for this range.
     * 
     */
    public Optional<Double> lt() {
        return Optional.ofNullable(this.lt);
    }
    /**
     * @return Indicates the upper threshold inclusive value for this range.
     * 
     */
    public Optional<Double> lte() {
        return Optional.ofNullable(this.lte);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ListChartColorScale defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String color;
        private @Nullable Double gt;
        private @Nullable Double gte;
        private @Nullable Double lt;
        private @Nullable Double lte;
        public Builder() {}
        public Builder(ListChartColorScale defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.color = defaults.color;
    	      this.gt = defaults.gt;
    	      this.gte = defaults.gte;
    	      this.lt = defaults.lt;
    	      this.lte = defaults.lte;
        }

        @CustomType.Setter
        public Builder color(String color) {
            this.color = Objects.requireNonNull(color);
            return this;
        }
        @CustomType.Setter
        public Builder gt(@Nullable Double gt) {
            this.gt = gt;
            return this;
        }
        @CustomType.Setter
        public Builder gte(@Nullable Double gte) {
            this.gte = gte;
            return this;
        }
        @CustomType.Setter
        public Builder lt(@Nullable Double lt) {
            this.lt = lt;
            return this;
        }
        @CustomType.Setter
        public Builder lte(@Nullable Double lte) {
            this.lte = lte;
            return this;
        }
        public ListChartColorScale build() {
            final var o = new ListChartColorScale();
            o.color = color;
            o.gt = gt;
            o.gte = gte;
            o.lt = lt;
            o.lte = lte;
            return o;
        }
    }
}
