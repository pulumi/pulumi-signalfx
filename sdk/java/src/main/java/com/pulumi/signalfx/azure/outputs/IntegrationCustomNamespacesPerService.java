// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.azure.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class IntegrationCustomNamespacesPerService {
    /**
     * @return The additional namespaces.
     * 
     */
    private final List<String> namespaces;
    /**
     * @return The name of the service.
     * 
     */
    private final String service;

    @CustomType.Constructor
    private IntegrationCustomNamespacesPerService(
        @CustomType.Parameter("namespaces") List<String> namespaces,
        @CustomType.Parameter("service") String service) {
        this.namespaces = namespaces;
        this.service = service;
    }

    /**
     * @return The additional namespaces.
     * 
     */
    public List<String> namespaces() {
        return this.namespaces;
    }
    /**
     * @return The name of the service.
     * 
     */
    public String service() {
        return this.service;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IntegrationCustomNamespacesPerService defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private List<String> namespaces;
        private String service;

        public Builder() {
    	      // Empty
        }

        public Builder(IntegrationCustomNamespacesPerService defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.namespaces = defaults.namespaces;
    	      this.service = defaults.service;
        }

        public Builder namespaces(List<String> namespaces) {
            this.namespaces = Objects.requireNonNull(namespaces);
            return this;
        }
        public Builder namespaces(String... namespaces) {
            return namespaces(List.of(namespaces));
        }
        public Builder service(String service) {
            this.service = Objects.requireNonNull(service);
            return this;
        }        public IntegrationCustomNamespacesPerService build() {
            return new IntegrationCustomNamespacesPerService(namespaces, service);
        }
    }
}
