// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.aws.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IntegrationCustomNamespaceSyncRule {
    /**
     * @return Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn&#39;t match the filter. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
     * 
     */
    private @Nullable String defaultAction;
    /**
     * @return Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
     * 
     */
    private @Nullable String filterAction;
    /**
     * @return Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
     * 
     */
    private @Nullable String filterSource;
    /**
     * @return An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
     * 
     */
    private String namespace;

    private IntegrationCustomNamespaceSyncRule() {}
    /**
     * @return Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn&#39;t match the filter. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
     * 
     */
    public Optional<String> defaultAction() {
        return Optional.ofNullable(this.defaultAction);
    }
    /**
     * @return Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
     * 
     */
    public Optional<String> filterAction() {
        return Optional.ofNullable(this.filterAction);
    }
    /**
     * @return Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
     * 
     */
    public Optional<String> filterSource() {
        return Optional.ofNullable(this.filterSource);
    }
    /**
     * @return An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
     * 
     */
    public String namespace() {
        return this.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IntegrationCustomNamespaceSyncRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String defaultAction;
        private @Nullable String filterAction;
        private @Nullable String filterSource;
        private String namespace;
        public Builder() {}
        public Builder(IntegrationCustomNamespaceSyncRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.defaultAction = defaults.defaultAction;
    	      this.filterAction = defaults.filterAction;
    	      this.filterSource = defaults.filterSource;
    	      this.namespace = defaults.namespace;
        }

        @CustomType.Setter
        public Builder defaultAction(@Nullable String defaultAction) {

            this.defaultAction = defaultAction;
            return this;
        }
        @CustomType.Setter
        public Builder filterAction(@Nullable String filterAction) {

            this.filterAction = filterAction;
            return this;
        }
        @CustomType.Setter
        public Builder filterSource(@Nullable String filterSource) {

            this.filterSource = filterSource;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(String namespace) {
            if (namespace == null) {
              throw new MissingRequiredPropertyException("IntegrationCustomNamespaceSyncRule", "namespace");
            }
            this.namespace = namespace;
            return this;
        }
        public IntegrationCustomNamespaceSyncRule build() {
            final var _resultValue = new IntegrationCustomNamespaceSyncRule();
            _resultValue.defaultAction = defaultAction;
            _resultValue.filterAction = filterAction;
            _resultValue.filterSource = filterSource;
            _resultValue.namespace = namespace;
            return _resultValue;
        }
    }
}
