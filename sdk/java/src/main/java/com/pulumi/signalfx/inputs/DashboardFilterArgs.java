// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DashboardFilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final DashboardFilterArgs Empty = new DashboardFilterArgs();

    /**
     * If true, this filter will also match data that does not have the specified property
     * 
     */
    @Import(name="applyIfExist")
    private @Nullable Output<Boolean> applyIfExist;

    /**
     * @return If true, this filter will also match data that does not have the specified property
     * 
     */
    public Optional<Output<Boolean>> applyIfExist() {
        return Optional.ofNullable(this.applyIfExist);
    }

    /**
     * (false by default) Whether this filter should be a &#34;not&#34; filter
     * 
     */
    @Import(name="negated")
    private @Nullable Output<Boolean> negated;

    /**
     * @return (false by default) Whether this filter should be a &#34;not&#34; filter
     * 
     */
    public Optional<Output<Boolean>> negated() {
        return Optional.ofNullable(this.negated);
    }

    /**
     * A metric time series dimension or property name
     * 
     */
    @Import(name="property", required=true)
    private Output<String> property;

    /**
     * @return A metric time series dimension or property name
     * 
     */
    public Output<String> property() {
        return this.property;
    }

    /**
     * List of strings (which will be treated as an OR filter on the property)
     * 
     */
    @Import(name="values", required=true)
    private Output<List<String>> values;

    /**
     * @return List of strings (which will be treated as an OR filter on the property)
     * 
     */
    public Output<List<String>> values() {
        return this.values;
    }

    private DashboardFilterArgs() {}

    private DashboardFilterArgs(DashboardFilterArgs $) {
        this.applyIfExist = $.applyIfExist;
        this.negated = $.negated;
        this.property = $.property;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DashboardFilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DashboardFilterArgs $;

        public Builder() {
            $ = new DashboardFilterArgs();
        }

        public Builder(DashboardFilterArgs defaults) {
            $ = new DashboardFilterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applyIfExist If true, this filter will also match data that does not have the specified property
         * 
         * @return builder
         * 
         */
        public Builder applyIfExist(@Nullable Output<Boolean> applyIfExist) {
            $.applyIfExist = applyIfExist;
            return this;
        }

        /**
         * @param applyIfExist If true, this filter will also match data that does not have the specified property
         * 
         * @return builder
         * 
         */
        public Builder applyIfExist(Boolean applyIfExist) {
            return applyIfExist(Output.of(applyIfExist));
        }

        /**
         * @param negated (false by default) Whether this filter should be a &#34;not&#34; filter
         * 
         * @return builder
         * 
         */
        public Builder negated(@Nullable Output<Boolean> negated) {
            $.negated = negated;
            return this;
        }

        /**
         * @param negated (false by default) Whether this filter should be a &#34;not&#34; filter
         * 
         * @return builder
         * 
         */
        public Builder negated(Boolean negated) {
            return negated(Output.of(negated));
        }

        /**
         * @param property A metric time series dimension or property name
         * 
         * @return builder
         * 
         */
        public Builder property(Output<String> property) {
            $.property = property;
            return this;
        }

        /**
         * @param property A metric time series dimension or property name
         * 
         * @return builder
         * 
         */
        public Builder property(String property) {
            return property(Output.of(property));
        }

        /**
         * @param values List of strings (which will be treated as an OR filter on the property)
         * 
         * @return builder
         * 
         */
        public Builder values(Output<List<String>> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values List of strings (which will be treated as an OR filter on the property)
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        /**
         * @param values List of strings (which will be treated as an OR filter on the property)
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public DashboardFilterArgs build() {
            if ($.property == null) {
                throw new MissingRequiredPropertyException("DashboardFilterArgs", "property");
            }
            if ($.values == null) {
                throw new MissingRequiredPropertyException("DashboardFilterArgs", "values");
            }
            return $;
        }
    }

}
