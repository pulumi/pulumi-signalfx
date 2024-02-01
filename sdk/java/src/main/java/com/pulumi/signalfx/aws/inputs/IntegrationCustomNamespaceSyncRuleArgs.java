// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.aws.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IntegrationCustomNamespaceSyncRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final IntegrationCustomNamespaceSyncRuleArgs Empty = new IntegrationCustomNamespaceSyncRuleArgs();

    /**
     * Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn&#39;t match the filter. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
     * 
     */
    @Import(name="defaultAction")
    private @Nullable Output<String> defaultAction;

    /**
     * @return Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn&#39;t match the filter. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
     * 
     */
    public Optional<Output<String>> defaultAction() {
        return Optional.ofNullable(this.defaultAction);
    }

    /**
     * Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
     * 
     */
    @Import(name="filterAction")
    private @Nullable Output<String> filterAction;

    /**
     * @return Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
     * 
     */
    public Optional<Output<String>> filterAction() {
        return Optional.ofNullable(this.filterAction);
    }

    /**
     * Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
     * 
     */
    @Import(name="filterSource")
    private @Nullable Output<String> filterSource;

    /**
     * @return Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
     * 
     */
    public Optional<Output<String>> filterSource() {
        return Optional.ofNullable(this.filterSource);
    }

    /**
     * An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
     * 
     */
    @Import(name="namespace", required=true)
    private Output<String> namespace;

    /**
     * @return An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
     * 
     */
    public Output<String> namespace() {
        return this.namespace;
    }

    private IntegrationCustomNamespaceSyncRuleArgs() {}

    private IntegrationCustomNamespaceSyncRuleArgs(IntegrationCustomNamespaceSyncRuleArgs $) {
        this.defaultAction = $.defaultAction;
        this.filterAction = $.filterAction;
        this.filterSource = $.filterSource;
        this.namespace = $.namespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IntegrationCustomNamespaceSyncRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IntegrationCustomNamespaceSyncRuleArgs $;

        public Builder() {
            $ = new IntegrationCustomNamespaceSyncRuleArgs();
        }

        public Builder(IntegrationCustomNamespaceSyncRuleArgs defaults) {
            $ = new IntegrationCustomNamespaceSyncRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultAction Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn&#39;t match the filter. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
         * 
         * @return builder
         * 
         */
        public Builder defaultAction(@Nullable Output<String> defaultAction) {
            $.defaultAction = defaultAction;
            return this;
        }

        /**
         * @param defaultAction Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filter_action` and `filter_source` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn&#39;t match the filter. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
         * 
         * @return builder
         * 
         */
        public Builder defaultAction(String defaultAction) {
            return defaultAction(Output.of(defaultAction));
        }

        /**
         * @param filterAction Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
         * 
         * @return builder
         * 
         */
        public Builder filterAction(@Nullable Output<String> filterAction) {
            $.filterAction = filterAction;
            return this;
        }

        /**
         * @param filterAction Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of &#34;Include&#34; or &#34;Exclude&#34;.
         * 
         * @return builder
         * 
         */
        public Builder filterAction(String filterAction) {
            return filterAction(Output.of(filterAction));
        }

        /**
         * @param filterSource Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
         * 
         * @return builder
         * 
         */
        public Builder filterSource(@Nullable Output<String> filterSource) {
            $.filterSource = filterSource;
            return this;
        }

        /**
         * @param filterSource Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
         * 
         * @return builder
         * 
         */
        public Builder filterSource(String filterSource) {
            return filterSource(Output.of(filterSource));
        }

        /**
         * @param namespace An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
         * 
         * @return builder
         * 
         */
        public Builder namespace(Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public IntegrationCustomNamespaceSyncRuleArgs build() {
            if ($.namespace == null) {
                throw new MissingRequiredPropertyException("IntegrationCustomNamespaceSyncRuleArgs", "namespace");
            }
            return $;
        }
    }

}
