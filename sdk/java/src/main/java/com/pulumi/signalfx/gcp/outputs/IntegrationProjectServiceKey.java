// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.gcp.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class IntegrationProjectServiceKey {
    private String projectId;
    private String projectKey;

    private IntegrationProjectServiceKey() {}
    public String projectId() {
        return this.projectId;
    }
    public String projectKey() {
        return this.projectKey;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IntegrationProjectServiceKey defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String projectId;
        private String projectKey;
        public Builder() {}
        public Builder(IntegrationProjectServiceKey defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.projectId = defaults.projectId;
    	      this.projectKey = defaults.projectKey;
        }

        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("IntegrationProjectServiceKey", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder projectKey(String projectKey) {
            if (projectKey == null) {
              throw new MissingRequiredPropertyException("IntegrationProjectServiceKey", "projectKey");
            }
            this.projectKey = projectKey;
            return this;
        }
        public IntegrationProjectServiceKey build() {
            final var _resultValue = new IntegrationProjectServiceKey();
            _resultValue.projectId = projectId;
            _resultValue.projectKey = projectKey;
            return _resultValue;
        }
    }
}