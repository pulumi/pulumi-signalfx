// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.aws.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IntegrationNamespaceSyncRule {
    /**
     * @return Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn&#39;t match the filter. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
     * 
     */
    private final @Nullable String defaultAction;
    /**
     * @return Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
     * 
     */
    private final @Nullable String filterAction;
    /**
     * @return Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
     * 
     */
    private final @Nullable String filterSource;
    /**
     * @return An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
     * 
     */
    private final String namespace;

    @CustomType.Constructor
    private IntegrationNamespaceSyncRule(
        @CustomType.Parameter("defaultAction") @Nullable String defaultAction,
        @CustomType.Parameter("filterAction") @Nullable String filterAction,
        @CustomType.Parameter("filterSource") @Nullable String filterSource,
        @CustomType.Parameter("namespace") String namespace) {
        this.defaultAction = defaultAction;
        this.filterAction = filterAction;
        this.filterSource = filterSource;
        this.namespace = namespace;
    }

    /**
     * @return Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn&#39;t match the filter. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
     * 
     */
    public Optional<String> defaultAction() {
        return Optional.ofNullable(this.defaultAction);
    }
    /**
     * @return Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
     * 
     */
    public Optional<String> filterAction() {
        return Optional.ofNullable(this.filterAction);
    }
    /**
     * @return Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
     * 
     */
    public Optional<String> filterSource() {
        return Optional.ofNullable(this.filterSource);
    }
    /**
     * @return An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
     * 
     */
    public String namespace() {
        return this.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IntegrationNamespaceSyncRule defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String defaultAction;
        private @Nullable String filterAction;
        private @Nullable String filterSource;
        private String namespace;

        public Builder() {
    	      // Empty
        }

        public Builder(IntegrationNamespaceSyncRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.defaultAction = defaults.defaultAction;
    	      this.filterAction = defaults.filterAction;
    	      this.filterSource = defaults.filterSource;
    	      this.namespace = defaults.namespace;
        }

        public Builder defaultAction(@Nullable String defaultAction) {
            this.defaultAction = defaultAction;
            return this;
        }
        public Builder filterAction(@Nullable String filterAction) {
            this.filterAction = filterAction;
            return this;
        }
        public Builder filterSource(@Nullable String filterSource) {
            this.filterSource = filterSource;
            return this;
        }
        public Builder namespace(String namespace) {
            this.namespace = Objects.requireNonNull(namespace);
            return this;
        }        public IntegrationNamespaceSyncRule build() {
            return new IntegrationNamespaceSyncRule(defaultAction, filterAction, filterSource, namespace);
        }
    }
}
