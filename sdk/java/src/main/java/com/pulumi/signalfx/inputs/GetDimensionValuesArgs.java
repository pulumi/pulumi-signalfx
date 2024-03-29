// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetDimensionValuesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDimensionValuesArgs Empty = new GetDimensionValuesArgs();

    @Import(name="query", required=true)
    private Output<String> query;

    public Output<String> query() {
        return this.query;
    }

    private GetDimensionValuesArgs() {}

    private GetDimensionValuesArgs(GetDimensionValuesArgs $) {
        this.query = $.query;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDimensionValuesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDimensionValuesArgs $;

        public Builder() {
            $ = new GetDimensionValuesArgs();
        }

        public Builder(GetDimensionValuesArgs defaults) {
            $ = new GetDimensionValuesArgs(Objects.requireNonNull(defaults));
        }

        public Builder query(Output<String> query) {
            $.query = query;
            return this;
        }

        public Builder query(String query) {
            return query(Output.of(query));
        }

        public GetDimensionValuesArgs build() {
            if ($.query == null) {
                throw new MissingRequiredPropertyException("GetDimensionValuesArgs", "query");
            }
            return $;
        }
    }

}
