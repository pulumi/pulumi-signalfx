// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DataLinkTargetExternalUrl {
    /**
     * @return The [minimum time window](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) for a search sent to an external site. Defaults to `6000`
     * 
     */
    private @Nullable String minimumTimeWindow;
    /**
     * @return User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     * 
     */
    private String name;
    /**
     * @return Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.
     * 
     */
    private @Nullable Map<String,String> propertyKeyMapping;
    /**
     * @return [Designates the format](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) of `minimum_time_window` in the same data link target object. Must be one of `&#34;ISO8601&#34;`, `&#34;EpochSeconds&#34;` or `&#34;Epoch&#34;` (which is milliseconds). Defaults to `&#34;ISO8601&#34;`.
     * 
     */
    private @Nullable String timeFormat;
    /**
     * @return URL string for a Splunk instance or external system data link target. [See the supported template variables](https://developers.signalfx.com/administration/data_links_overview.html#_external_link_targets).
     * 
     */
    private String url;

    private DataLinkTargetExternalUrl() {}
    /**
     * @return The [minimum time window](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) for a search sent to an external site. Defaults to `6000`
     * 
     */
    public Optional<String> minimumTimeWindow() {
        return Optional.ofNullable(this.minimumTimeWindow);
    }
    /**
     * @return User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.
     * 
     */
    public Map<String,String> propertyKeyMapping() {
        return this.propertyKeyMapping == null ? Map.of() : this.propertyKeyMapping;
    }
    /**
     * @return [Designates the format](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) of `minimum_time_window` in the same data link target object. Must be one of `&#34;ISO8601&#34;`, `&#34;EpochSeconds&#34;` or `&#34;Epoch&#34;` (which is milliseconds). Defaults to `&#34;ISO8601&#34;`.
     * 
     */
    public Optional<String> timeFormat() {
        return Optional.ofNullable(this.timeFormat);
    }
    /**
     * @return URL string for a Splunk instance or external system data link target. [See the supported template variables](https://developers.signalfx.com/administration/data_links_overview.html#_external_link_targets).
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DataLinkTargetExternalUrl defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String minimumTimeWindow;
        private String name;
        private @Nullable Map<String,String> propertyKeyMapping;
        private @Nullable String timeFormat;
        private String url;
        public Builder() {}
        public Builder(DataLinkTargetExternalUrl defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.minimumTimeWindow = defaults.minimumTimeWindow;
    	      this.name = defaults.name;
    	      this.propertyKeyMapping = defaults.propertyKeyMapping;
    	      this.timeFormat = defaults.timeFormat;
    	      this.url = defaults.url;
        }

        @CustomType.Setter
        public Builder minimumTimeWindow(@Nullable String minimumTimeWindow) {
            this.minimumTimeWindow = minimumTimeWindow;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder propertyKeyMapping(@Nullable Map<String,String> propertyKeyMapping) {
            this.propertyKeyMapping = propertyKeyMapping;
            return this;
        }
        @CustomType.Setter
        public Builder timeFormat(@Nullable String timeFormat) {
            this.timeFormat = timeFormat;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        public DataLinkTargetExternalUrl build() {
            final var _resultValue = new DataLinkTargetExternalUrl();
            _resultValue.minimumTimeWindow = minimumTimeWindow;
            _resultValue.name = name;
            _resultValue.propertyKeyMapping = propertyKeyMapping;
            _resultValue.timeFormat = timeFormat;
            _resultValue.url = url;
            return _resultValue;
        }
    }
}
