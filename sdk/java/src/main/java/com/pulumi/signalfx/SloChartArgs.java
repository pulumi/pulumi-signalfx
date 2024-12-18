// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class SloChartArgs extends com.pulumi.resources.ResourceArgs {

    public static final SloChartArgs Empty = new SloChartArgs();

    /**
     * ID of SLO object.
     * 
     */
    @Import(name="sloId", required=true)
    private Output<String> sloId;

    /**
     * @return ID of SLO object.
     * 
     */
    public Output<String> sloId() {
        return this.sloId;
    }

    private SloChartArgs() {}

    private SloChartArgs(SloChartArgs $) {
        this.sloId = $.sloId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SloChartArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SloChartArgs $;

        public Builder() {
            $ = new SloChartArgs();
        }

        public Builder(SloChartArgs defaults) {
            $ = new SloChartArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param sloId ID of SLO object.
         * 
         * @return builder
         * 
         */
        public Builder sloId(Output<String> sloId) {
            $.sloId = sloId;
            return this;
        }

        /**
         * @param sloId ID of SLO object.
         * 
         * @return builder
         * 
         */
        public Builder sloId(String sloId) {
            return sloId(Output.of(sloId));
        }

        public SloChartArgs build() {
            if ($.sloId == null) {
                throw new MissingRequiredPropertyException("SloChartArgs", "sloId");
            }
            return $;
        }
    }

}
