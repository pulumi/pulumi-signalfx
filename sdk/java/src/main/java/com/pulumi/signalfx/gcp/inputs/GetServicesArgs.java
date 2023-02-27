// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.gcp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.gcp.inputs.GetServicesServiceArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServicesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServicesArgs Empty = new GetServicesArgs();

    @Import(name="services")
    private @Nullable Output<List<GetServicesServiceArgs>> services;

    public Optional<Output<List<GetServicesServiceArgs>>> services() {
        return Optional.ofNullable(this.services);
    }

    private GetServicesArgs() {}

    private GetServicesArgs(GetServicesArgs $) {
        this.services = $.services;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServicesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServicesArgs $;

        public Builder() {
            $ = new GetServicesArgs();
        }

        public Builder(GetServicesArgs defaults) {
            $ = new GetServicesArgs(Objects.requireNonNull(defaults));
        }

        public Builder services(@Nullable Output<List<GetServicesServiceArgs>> services) {
            $.services = services;
            return this;
        }

        public Builder services(List<GetServicesServiceArgs> services) {
            return services(Output.of(services));
        }

        public Builder services(GetServicesServiceArgs... services) {
            return services(List.of(services));
        }

        public GetServicesArgs build() {
            return $;
        }
    }

}