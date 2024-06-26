// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TimeChartHistogramOptionArgs extends com.pulumi.resources.ResourceArgs {

    public static final TimeChartHistogramOptionArgs Empty = new TimeChartHistogramOptionArgs();

    /**
     * Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine, red, gold, greenyellow, chartreuse, jade
     * 
     */
    @Import(name="colorTheme")
    private @Nullable Output<String> colorTheme;

    /**
     * @return Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine, red, gold, greenyellow, chartreuse, jade
     * 
     */
    public Optional<Output<String>> colorTheme() {
        return Optional.ofNullable(this.colorTheme);
    }

    private TimeChartHistogramOptionArgs() {}

    private TimeChartHistogramOptionArgs(TimeChartHistogramOptionArgs $) {
        this.colorTheme = $.colorTheme;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TimeChartHistogramOptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TimeChartHistogramOptionArgs $;

        public Builder() {
            $ = new TimeChartHistogramOptionArgs();
        }

        public Builder(TimeChartHistogramOptionArgs defaults) {
            $ = new TimeChartHistogramOptionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param colorTheme Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine, red, gold, greenyellow, chartreuse, jade
         * 
         * @return builder
         * 
         */
        public Builder colorTheme(@Nullable Output<String> colorTheme) {
            $.colorTheme = colorTheme;
            return this;
        }

        /**
         * @param colorTheme Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine, red, gold, greenyellow, chartreuse, jade
         * 
         * @return builder
         * 
         */
        public Builder colorTheme(String colorTheme) {
            return colorTheme(Output.of(colorTheme));
        }

        public TimeChartHistogramOptionArgs build() {
            return $;
        }
    }

}
