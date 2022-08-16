// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AlertMutingRuleFilter {
    /**
     * @return Determines if this is a &#34;not&#34; filter. Defaults to `false`.
     * 
     */
    private final @Nullable Boolean negated;
    /**
     * @return The property to filter.
     * 
     */
    private final String property;
    /**
     * @return The property value to filter.
     * 
     */
    private final String propertyValue;

    @CustomType.Constructor
    private AlertMutingRuleFilter(
        @CustomType.Parameter("negated") @Nullable Boolean negated,
        @CustomType.Parameter("property") String property,
        @CustomType.Parameter("propertyValue") String propertyValue) {
        this.negated = negated;
        this.property = property;
        this.propertyValue = propertyValue;
    }

    /**
     * @return Determines if this is a &#34;not&#34; filter. Defaults to `false`.
     * 
     */
    public Optional<Boolean> negated() {
        return Optional.ofNullable(this.negated);
    }
    /**
     * @return The property to filter.
     * 
     */
    public String property() {
        return this.property;
    }
    /**
     * @return The property value to filter.
     * 
     */
    public String propertyValue() {
        return this.propertyValue;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AlertMutingRuleFilter defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Boolean negated;
        private String property;
        private String propertyValue;

        public Builder() {
    	      // Empty
        }

        public Builder(AlertMutingRuleFilter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.negated = defaults.negated;
    	      this.property = defaults.property;
    	      this.propertyValue = defaults.propertyValue;
        }

        public Builder negated(@Nullable Boolean negated) {
            this.negated = negated;
            return this;
        }
        public Builder property(String property) {
            this.property = Objects.requireNonNull(property);
            return this;
        }
        public Builder propertyValue(String propertyValue) {
            this.propertyValue = Objects.requireNonNull(propertyValue);
            return this;
        }        public AlertMutingRuleFilter build() {
            return new AlertMutingRuleFilter(negated, property, propertyValue);
        }
    }
}
