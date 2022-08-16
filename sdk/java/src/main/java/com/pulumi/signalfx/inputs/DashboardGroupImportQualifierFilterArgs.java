// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardGroupImportQualifierFilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardGroupImportQualifierFilterArgs Empty = new DashboardGroupImportQualifierFilterArgs();

    /**
     * If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
     * 
     */
    @Import(name="negated")
    private @Nullable Output<Boolean> negated;

    /**
     * @return If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> negated() {
        return Optional.ofNullable(this.negated);
    }

    /**
     * A metric time series dimension or property name.
     * 
     */
    @Import(name="property", required=true)
    private Output<String> property;

    /**
     * @return A metric time series dimension or property name.
     * 
     */
    public Output<String> property() {
        return this.property;
    }

    /**
     * (Optional) List of of strings (which will be treated as an OR filter on the property).
     * 
     */
    @Import(name="values", required=true)
    private Output<List<String>> values;

    /**
     * @return (Optional) List of of strings (which will be treated as an OR filter on the property).
     * 
     */
    public Output<List<String>> values() {
        return this.values;
    }

    private DashboardGroupImportQualifierFilterArgs() {}

    private DashboardGroupImportQualifierFilterArgs(DashboardGroupImportQualifierFilterArgs $) {
        this.negated = $.negated;
        this.property = $.property;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardGroupImportQualifierFilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardGroupImportQualifierFilterArgs $;

        public Builder() {
            $ = new DashboardGroupImportQualifierFilterArgs();
        }

        public Builder(DashboardGroupImportQualifierFilterArgs defaults) {
            $ = new DashboardGroupImportQualifierFilterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param negated If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder negated(@Nullable Output<Boolean> negated) {
            $.negated = negated;
            return this;
        }

        /**
         * @param negated If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder negated(Boolean negated) {
            return negated(Output.of(negated));
        }

        /**
         * @param property A metric time series dimension or property name.
         * 
         * @return builder
         * 
         */
        public Builder property(Output<String> property) {
            $.property = property;
            return this;
        }

        /**
         * @param property A metric time series dimension or property name.
         * 
         * @return builder
         * 
         */
        public Builder property(String property) {
            return property(Output.of(property));
        }

        /**
         * @param values (Optional) List of of strings (which will be treated as an OR filter on the property).
         * 
         * @return builder
         * 
         */
        public Builder values(Output<List<String>> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values (Optional) List of of strings (which will be treated as an OR filter on the property).
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        /**
         * @param values (Optional) List of of strings (which will be treated as an OR filter on the property).
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public DashboardGroupImportQualifierFilterArgs build() {
            $.property = Objects.requireNonNull($.property, "expected parameter 'property' to be non-null");
            $.values = Objects.requireNonNull($.values, "expected parameter 'values' to be non-null");
            return $;
        }
    }

}
