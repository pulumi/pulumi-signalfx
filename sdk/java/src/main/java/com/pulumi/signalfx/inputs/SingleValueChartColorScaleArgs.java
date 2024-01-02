// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SingleValueChartColorScaleArgs extends com.pulumi.resources.ResourceArgs {

    public static final SingleValueChartColorScaleArgs Empty = new SingleValueChartColorScaleArgs();

    /**
     * The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     * 
     */
    @Import(name="color", required=true)
    private Output<String> color;

    /**
     * @return The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     * 
     */
    public Output<String> color() {
        return this.color;
    }

    /**
     * Indicates the lower threshold non-inclusive value for this range.
     * 
     */
    @Import(name="gt")
    private @Nullable Output<Double> gt;

    /**
     * @return Indicates the lower threshold non-inclusive value for this range.
     * 
     */
    public Optional<Output<Double>> gt() {
        return Optional.ofNullable(this.gt);
    }

    /**
     * Indicates the lower threshold inclusive value for this range.
     * 
     */
    @Import(name="gte")
    private @Nullable Output<Double> gte;

    /**
     * @return Indicates the lower threshold inclusive value for this range.
     * 
     */
    public Optional<Output<Double>> gte() {
        return Optional.ofNullable(this.gte);
    }

    /**
     * Indicates the upper threshold non-inculsive value for this range.
     * 
     */
    @Import(name="lt")
    private @Nullable Output<Double> lt;

    /**
     * @return Indicates the upper threshold non-inculsive value for this range.
     * 
     */
    public Optional<Output<Double>> lt() {
        return Optional.ofNullable(this.lt);
    }

    /**
     * Indicates the upper threshold inclusive value for this range.
     * 
     */
    @Import(name="lte")
    private @Nullable Output<Double> lte;

    /**
     * @return Indicates the upper threshold inclusive value for this range.
     * 
     */
    public Optional<Output<Double>> lte() {
        return Optional.ofNullable(this.lte);
    }

    private SingleValueChartColorScaleArgs() {}

    private SingleValueChartColorScaleArgs(SingleValueChartColorScaleArgs $) {
        this.color = $.color;
        this.gt = $.gt;
        this.gte = $.gte;
        this.lt = $.lt;
        this.lte = $.lte;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SingleValueChartColorScaleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SingleValueChartColorScaleArgs $;

        public Builder() {
            $ = new SingleValueChartColorScaleArgs();
        }

        public Builder(SingleValueChartColorScaleArgs defaults) {
            $ = new SingleValueChartColorScaleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param color The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
         * 
         * @return builder
         * 
         */
        public Builder color(Output<String> color) {
            $.color = color;
            return this;
        }

        /**
         * @param color The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
         * 
         * @return builder
         * 
         */
        public Builder color(String color) {
            return color(Output.of(color));
        }

        /**
         * @param gt Indicates the lower threshold non-inclusive value for this range.
         * 
         * @return builder
         * 
         */
        public Builder gt(@Nullable Output<Double> gt) {
            $.gt = gt;
            return this;
        }

        /**
         * @param gt Indicates the lower threshold non-inclusive value for this range.
         * 
         * @return builder
         * 
         */
        public Builder gt(Double gt) {
            return gt(Output.of(gt));
        }

        /**
         * @param gte Indicates the lower threshold inclusive value for this range.
         * 
         * @return builder
         * 
         */
        public Builder gte(@Nullable Output<Double> gte) {
            $.gte = gte;
            return this;
        }

        /**
         * @param gte Indicates the lower threshold inclusive value for this range.
         * 
         * @return builder
         * 
         */
        public Builder gte(Double gte) {
            return gte(Output.of(gte));
        }

        /**
         * @param lt Indicates the upper threshold non-inculsive value for this range.
         * 
         * @return builder
         * 
         */
        public Builder lt(@Nullable Output<Double> lt) {
            $.lt = lt;
            return this;
        }

        /**
         * @param lt Indicates the upper threshold non-inculsive value for this range.
         * 
         * @return builder
         * 
         */
        public Builder lt(Double lt) {
            return lt(Output.of(lt));
        }

        /**
         * @param lte Indicates the upper threshold inclusive value for this range.
         * 
         * @return builder
         * 
         */
        public Builder lte(@Nullable Output<Double> lte) {
            $.lte = lte;
            return this;
        }

        /**
         * @param lte Indicates the upper threshold inclusive value for this range.
         * 
         * @return builder
         * 
         */
        public Builder lte(Double lte) {
            return lte(Output.of(lte));
        }

        public SingleValueChartColorScaleArgs build() {
            if ($.color == null) {
                throw new MissingRequiredPropertyException("SingleValueChartColorScaleArgs", "color");
            }
            return $;
        }
    }

}
