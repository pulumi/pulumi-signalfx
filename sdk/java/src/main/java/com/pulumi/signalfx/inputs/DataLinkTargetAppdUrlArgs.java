// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class DataLinkTargetAppdUrlArgs extends com.pulumi.resources.ResourceArgs {

    public static final DataLinkTargetAppdUrlArgs Empty = new DataLinkTargetAppdUrlArgs();

    /**
     * User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * URL string for an AppDynamics instance.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return URL string for an AppDynamics instance.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private DataLinkTargetAppdUrlArgs() {}

    private DataLinkTargetAppdUrlArgs(DataLinkTargetAppdUrlArgs $) {
        this.name = $.name;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DataLinkTargetAppdUrlArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DataLinkTargetAppdUrlArgs $;

        public Builder() {
            $ = new DataLinkTargetAppdUrlArgs();
        }

        public Builder(DataLinkTargetAppdUrlArgs defaults) {
            $ = new DataLinkTargetAppdUrlArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name User-assigned target name. Use this value to differentiate between the link targets for a data link object.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name User-assigned target name. Use this value to differentiate between the link targets for a data link object.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param url URL string for an AppDynamics instance.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url URL string for an AppDynamics instance.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public DataLinkTargetAppdUrlArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("DataLinkTargetAppdUrlArgs", "name");
            }
            if ($.url == null) {
                throw new MissingRequiredPropertyException("DataLinkTargetAppdUrlArgs", "url");
            }
            return $;
        }
    }

}
