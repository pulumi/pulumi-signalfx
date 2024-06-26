// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class DashboardGroupPermission {
    /**
     * @return Action the user, team, or organization can take with the dashboard group. List of values (value can be &#34;READ&#34; or &#34;WRITE&#34;).
     * 
     */
    private @Nullable List<String> actions;
    /**
     * @return ID of the user, team, or organization for which you&#39;re granting permissions.
     * 
     */
    private String principalId;
    /**
     * @return Clarify whether this permission configuration is for a user, a team, or an organization. Value can be one of &#34;USER&#34;, &#34;TEAM&#34;, or &#34;ORG&#34;.
     * 
     */
    private String principalType;

    private DashboardGroupPermission() {}
    /**
     * @return Action the user, team, or organization can take with the dashboard group. List of values (value can be &#34;READ&#34; or &#34;WRITE&#34;).
     * 
     */
    public List<String> actions() {
        return this.actions == null ? List.of() : this.actions;
    }
    /**
     * @return ID of the user, team, or organization for which you&#39;re granting permissions.
     * 
     */
    public String principalId() {
        return this.principalId;
    }
    /**
     * @return Clarify whether this permission configuration is for a user, a team, or an organization. Value can be one of &#34;USER&#34;, &#34;TEAM&#34;, or &#34;ORG&#34;.
     * 
     */
    public String principalType() {
        return this.principalType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DashboardGroupPermission defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> actions;
        private String principalId;
        private String principalType;
        public Builder() {}
        public Builder(DashboardGroupPermission defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actions = defaults.actions;
    	      this.principalId = defaults.principalId;
    	      this.principalType = defaults.principalType;
        }

        @CustomType.Setter
        public Builder actions(@Nullable List<String> actions) {

            this.actions = actions;
            return this;
        }
        public Builder actions(String... actions) {
            return actions(List.of(actions));
        }
        @CustomType.Setter
        public Builder principalId(String principalId) {
            if (principalId == null) {
              throw new MissingRequiredPropertyException("DashboardGroupPermission", "principalId");
            }
            this.principalId = principalId;
            return this;
        }
        @CustomType.Setter
        public Builder principalType(String principalType) {
            if (principalType == null) {
              throw new MissingRequiredPropertyException("DashboardGroupPermission", "principalType");
            }
            this.principalType = principalType;
            return this;
        }
        public DashboardGroupPermission build() {
            final var _resultValue = new DashboardGroupPermission();
            _resultValue.actions = actions;
            _resultValue.principalId = principalId;
            _resultValue.principalType = principalType;
            return _resultValue;
        }
    }
}
