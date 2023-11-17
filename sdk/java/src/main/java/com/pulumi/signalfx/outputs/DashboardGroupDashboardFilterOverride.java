// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DashboardGroupDashboardFilterOverride {
    /**
     * @return If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
     * 
     */
    private @Nullable Boolean negated;
    /**
     * @return A metric time series dimension or property name.
     * 
     */
    private String property;
    /**
     * @return List of of strings (which will be treated as an OR filter on the property).
     * 
     */
    private List<String> values;

    private DashboardGroupDashboardFilterOverride() {}
    /**
     * @return If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
     * 
     */
    public Optional<Boolean> negated() {
        return Optional.ofNullable(this.negated);
    }
    /**
     * @return A metric time series dimension or property name.
     * 
     */
    public String property() {
        return this.property;
    }
    /**
     * @return List of of strings (which will be treated as an OR filter on the property).
     * 
     */
    public List<String> values() {
        return this.values;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardGroupDashboardFilterOverride defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean negated;
        private String property;
        private List<String> values;
        public Builder() {}
        public Builder(DashboardGroupDashboardFilterOverride defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.negated = defaults.negated;
    	      this.property = defaults.property;
    	      this.values = defaults.values;
        }

        @CustomType.Setter
        public Builder negated(@Nullable Boolean negated) {
            this.negated = negated;
            return this;
        }
        @CustomType.Setter
        public Builder property(String property) {
            this.property = Objects.requireNonNull(property);
            return this;
        }
        @CustomType.Setter
        public Builder values(List<String> values) {
            this.values = Objects.requireNonNull(values);
            return this;
        }
        public Builder values(String... values) {
            return values(List.of(values));
        }
        public DashboardGroupDashboardFilterOverride build() {
            final var o = new DashboardGroupDashboardFilterOverride();
            o.negated = negated;
            o.property = property;
            o.values = values;
            return o;
        }
    }
}
