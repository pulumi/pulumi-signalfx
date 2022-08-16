// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.inputs.GetAzureServicesServiceArgs;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAzureServicesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAzureServicesArgs Empty = new GetAzureServicesArgs();

    @Import(name="services")
    private @Nullable Output<List<GetAzureServicesServiceArgs>> services;

    public Optional<Output<List<GetAzureServicesServiceArgs>>> services() {
        return Optional.ofNullable(this.services);
    }

    private GetAzureServicesArgs() {}

    private GetAzureServicesArgs(GetAzureServicesArgs $) {
        this.services = $.services;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAzureServicesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAzureServicesArgs $;

        public Builder() {
            $ = new GetAzureServicesArgs();
        }

        public Builder(GetAzureServicesArgs defaults) {
            $ = new GetAzureServicesArgs(Objects.requireNonNull(defaults));
        }

        public Builder services(@Nullable Output<List<GetAzureServicesServiceArgs>> services) {
            $.services = services;
            return this;
        }

        public Builder services(List<GetAzureServicesServiceArgs> services) {
            return services(Output.of(services));
        }

        public Builder services(GetAzureServicesServiceArgs... services) {
            return services(List.of(services));
        }

        public GetAzureServicesArgs build() {
            return $;
        }
    }

}
