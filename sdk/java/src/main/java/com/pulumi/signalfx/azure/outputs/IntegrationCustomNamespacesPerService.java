// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.azure.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class IntegrationCustomNamespacesPerService {
    /**
     * @return The additional namespaces.
     * 
     */
    private List<String> namespaces;
    /**
     * @return The name of the service.
     * 
     */
    private String service;

    private IntegrationCustomNamespacesPerService() {}
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
    @CustomType.Builder
    public static final class Builder {
        private List<String> namespaces;
        private String service;
        public Builder() {}
        public Builder(IntegrationCustomNamespacesPerService defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.namespaces = defaults.namespaces;
    	      this.service = defaults.service;
        }

        @CustomType.Setter
        public Builder namespaces(List<String> namespaces) {
            if (namespaces == null) {
              throw new MissingRequiredPropertyException("IntegrationCustomNamespacesPerService", "namespaces");
            }
            this.namespaces = namespaces;
            return this;
        }
        public Builder namespaces(String... namespaces) {
            return namespaces(List.of(namespaces));
        }
        @CustomType.Setter
        public Builder service(String service) {
            if (service == null) {
              throw new MissingRequiredPropertyException("IntegrationCustomNamespacesPerService", "service");
            }
            this.service = service;
            return this;
        }
        public IntegrationCustomNamespacesPerService build() {
            final var _resultValue = new IntegrationCustomNamespacesPerService();
            _resultValue.namespaces = namespaces;
            _resultValue.service = service;
            return _resultValue;
        }
    }
}
