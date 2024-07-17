// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetDimensionValuesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDimensionValuesPlainArgs Empty = new GetDimensionValuesPlainArgs();

    @Import(name="query", required=true)
    private String query;

    public String query() {
        return this.query;
    }

    private GetDimensionValuesPlainArgs() {}

    private GetDimensionValuesPlainArgs(GetDimensionValuesPlainArgs $) {
        this.query = $.query;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDimensionValuesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDimensionValuesPlainArgs $;

        public Builder() {
            $ = new GetDimensionValuesPlainArgs();
        }

        public Builder(GetDimensionValuesPlainArgs defaults) {
            $ = new GetDimensionValuesPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder query(String query) {
            $.query = query;
            return this;
        }

        public GetDimensionValuesPlainArgs build() {
            if ($.query == null) {
                throw new MissingRequiredPropertyException("GetDimensionValuesPlainArgs", "query");
            }
            return $;
        }
    }

}