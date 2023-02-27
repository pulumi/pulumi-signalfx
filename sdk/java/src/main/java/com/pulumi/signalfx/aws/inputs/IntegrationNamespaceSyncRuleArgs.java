// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.aws.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IntegrationNamespaceSyncRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationNamespaceSyncRuleArgs Empty = new IntegrationNamespaceSyncRuleArgs();

    /**
     * Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn&#39;t match the filter. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
     * 
     */
    @Import(name="defaultAction")
    private @Nullable Output<String> defaultAction;

    /**
     * @return Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn&#39;t match the filter. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
     * 
     */
    public Optional<Output<String>> defaultAction() {
        return Optional.ofNullable(this.defaultAction);
    }

    /**
     * Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
     * 
     */
    @Import(name="filterAction")
    private @Nullable Output<String> filterAction;

    /**
     * @return Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
     * 
     */
    public Optional<Output<String>> filterAction() {
        return Optional.ofNullable(this.filterAction);
    }

    /**
     * Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
     * 
     */
    @Import(name="filterSource")
    private @Nullable Output<String> filterSource;

    /**
     * @return Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
     * 
     */
    public Optional<Output<String>> filterSource() {
        return Optional.ofNullable(this.filterSource);
    }

    /**
     * An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
     * 
     */
    @Import(name="namespace", required=true)
    private Output<String> namespace;

    /**
     * @return An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
     * 
     */
    public Output<String> namespace() {
        return this.namespace;
    }

    private IntegrationNamespaceSyncRuleArgs() {}

    private IntegrationNamespaceSyncRuleArgs(IntegrationNamespaceSyncRuleArgs $) {
        this.defaultAction = $.defaultAction;
        this.filterAction = $.filterAction;
        this.filterSource = $.filterSource;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationNamespaceSyncRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationNamespaceSyncRuleArgs $;

        public Builder() {
            $ = new IntegrationNamespaceSyncRuleArgs();
        }

        public Builder(IntegrationNamespaceSyncRuleArgs defaults) {
            $ = new IntegrationNamespaceSyncRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultAction Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn&#39;t match the filter. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder defaultAction(@Nullable Output<String> defaultAction) {
            $.defaultAction = defaultAction;
            return this;
        }

        /**
         * @param defaultAction Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn&#39;t match the filter. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder defaultAction(String defaultAction) {
            return defaultAction(Output.of(defaultAction));
        }

        /**
         * @param filterAction Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder filterAction(@Nullable Output<String> filterAction) {
            $.filterAction = filterAction;
            return this;
        }

        /**
         * @param filterAction Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `&#34;Include&#34;` or `&#34;Exclude&#34;`.
         * 
         * @return builder
         * 
         */
        public Builder filterAction(String filterAction) {
            return filterAction(Output.of(filterAction));
        }

        /**
         * @param filterSource Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
         * 
         * @return builder
         * 
         */
        public Builder filterSource(@Nullable Output<String> filterSource) {
            $.filterSource = filterSource;
            return this;
        }

        /**
         * @param filterSource Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
         * 
         * @return builder
         * 
         */
        public Builder filterSource(String filterSource) {
            return filterSource(Output.of(filterSource));
        }

        /**
         * @param namespace An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
         * 
         * @return builder
         * 
         */
        public Builder namespace(Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public IntegrationNamespaceSyncRuleArgs build() {
            $.namespace = Objects.requireNonNull($.namespace, "expected parameter 'namespace' to be non-null");
            return $;
        }
    }

}