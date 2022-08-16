// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.signalfx.inputs.OrgTokenDpmLimitsArgs;
import com.pulumi.signalfx.inputs.OrgTokenHostOrUsageLimitsArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrgTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final OrgTokenArgs Empty = new OrgTokenArgs();

    /**
     * Authentication scope, ex: INGEST, API, RUM ... (Optional)
     * 
     */
    @Import(name="authScopes")
    private @Nullable Output<List<String>> authScopes;

    /**
     * @return Authentication scope, ex: INGEST, API, RUM ... (Optional)
     * 
     */
    public Optional<Output<List<String>>> authScopes() {
        return Optional.ofNullable(this.authScopes);
    }

    /**
     * Description of the token.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the token.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Flag that controls enabling the token. If set to `true`, the token is disabled, and you can&#39;t use it for authentication. Defaults to `false`.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Flag that controls enabling the token. If set to `true`, the token is disabled, and you can&#39;t use it for authentication. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * Specify DPM-based limits for this token.
     * 
     */
    @Import(name="dpmLimits")
    private @Nullable Output<OrgTokenDpmLimitsArgs> dpmLimits;

    /**
     * @return Specify DPM-based limits for this token.
     * 
     */
    public Optional<Output<OrgTokenDpmLimitsArgs>> dpmLimits() {
        return Optional.ofNullable(this.dpmLimits);
    }

    /**
     * Specify Usage-based limits for this token.
     * 
     */
    @Import(name="hostOrUsageLimits")
    private @Nullable Output<OrgTokenHostOrUsageLimitsArgs> hostOrUsageLimits;

    /**
     * @return Specify Usage-based limits for this token.
     * 
     */
    public Optional<Output<OrgTokenHostOrUsageLimitsArgs>> hostOrUsageLimits() {
        return Optional.ofNullable(this.hostOrUsageLimits);
    }

    /**
     * Name of the token.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the token.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See
     * https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     * 
     */
    @Import(name="notifications")
    private @Nullable Output<List<String>> notifications;

    /**
     * @return List of strings specifying where notifications will be sent when an incident occurs. See
     * https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     * 
     */
    public Optional<Output<List<String>>> notifications() {
        return Optional.ofNullable(this.notifications);
    }

    private OrgTokenArgs() {}

    private OrgTokenArgs(OrgTokenArgs $) {
        this.authScopes = $.authScopes;
        this.description = $.description;
        this.disabled = $.disabled;
        this.dpmLimits = $.dpmLimits;
        this.hostOrUsageLimits = $.hostOrUsageLimits;
        this.name = $.name;
        this.notifications = $.notifications;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrgTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrgTokenArgs $;

        public Builder() {
            $ = new OrgTokenArgs();
        }

        public Builder(OrgTokenArgs defaults) {
            $ = new OrgTokenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authScopes Authentication scope, ex: INGEST, API, RUM ... (Optional)
         * 
         * @return builder
         * 
         */
        public Builder authScopes(@Nullable Output<List<String>> authScopes) {
            $.authScopes = authScopes;
            return this;
        }

        /**
         * @param authScopes Authentication scope, ex: INGEST, API, RUM ... (Optional)
         * 
         * @return builder
         * 
         */
        public Builder authScopes(List<String> authScopes) {
            return authScopes(Output.of(authScopes));
        }

        /**
         * @param authScopes Authentication scope, ex: INGEST, API, RUM ... (Optional)
         * 
         * @return builder
         * 
         */
        public Builder authScopes(String... authScopes) {
            return authScopes(List.of(authScopes));
        }

        /**
         * @param description Description of the token.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the token.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disabled Flag that controls enabling the token. If set to `true`, the token is disabled, and you can&#39;t use it for authentication. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Flag that controls enabling the token. If set to `true`, the token is disabled, and you can&#39;t use it for authentication. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param dpmLimits Specify DPM-based limits for this token.
         * 
         * @return builder
         * 
         */
        public Builder dpmLimits(@Nullable Output<OrgTokenDpmLimitsArgs> dpmLimits) {
            $.dpmLimits = dpmLimits;
            return this;
        }

        /**
         * @param dpmLimits Specify DPM-based limits for this token.
         * 
         * @return builder
         * 
         */
        public Builder dpmLimits(OrgTokenDpmLimitsArgs dpmLimits) {
            return dpmLimits(Output.of(dpmLimits));
        }

        /**
         * @param hostOrUsageLimits Specify Usage-based limits for this token.
         * 
         * @return builder
         * 
         */
        public Builder hostOrUsageLimits(@Nullable Output<OrgTokenHostOrUsageLimitsArgs> hostOrUsageLimits) {
            $.hostOrUsageLimits = hostOrUsageLimits;
            return this;
        }

        /**
         * @param hostOrUsageLimits Specify Usage-based limits for this token.
         * 
         * @return builder
         * 
         */
        public Builder hostOrUsageLimits(OrgTokenHostOrUsageLimitsArgs hostOrUsageLimits) {
            return hostOrUsageLimits(Output.of(hostOrUsageLimits));
        }

        /**
         * @param name Name of the token.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the token.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param notifications List of strings specifying where notifications will be sent when an incident occurs. See
         * https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
         * 
         * @return builder
         * 
         */
        public Builder notifications(@Nullable Output<List<String>> notifications) {
            $.notifications = notifications;
            return this;
        }

        /**
         * @param notifications List of strings specifying where notifications will be sent when an incident occurs. See
         * https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
         * 
         * @return builder
         * 
         */
        public Builder notifications(List<String> notifications) {
            return notifications(Output.of(notifications));
        }

        /**
         * @param notifications List of strings specifying where notifications will be sent when an incident occurs. See
         * https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
         * 
         * @return builder
         * 
         */
        public Builder notifications(String... notifications) {
            return notifications(List.of(notifications));
        }

        public OrgTokenArgs build() {
            return $;
        }
    }

}
