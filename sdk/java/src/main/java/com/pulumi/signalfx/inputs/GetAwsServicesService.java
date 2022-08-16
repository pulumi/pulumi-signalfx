// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetAwsServicesService extends com.pulumi.resources.InvokeArgs {

    public static final GetAwsServicesService Empty = new GetAwsServicesService();

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    private GetAwsServicesService() {}

    private GetAwsServicesService(GetAwsServicesService $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAwsServicesService defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAwsServicesService $;

        public Builder() {
            $ = new GetAwsServicesService();
        }

        public Builder(GetAwsServicesService defaults) {
            $ = new GetAwsServicesService(Objects.requireNonNull(defaults));
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetAwsServicesService build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
