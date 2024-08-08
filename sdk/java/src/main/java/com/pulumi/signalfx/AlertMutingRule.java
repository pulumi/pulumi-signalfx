// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.AlertMutingRuleArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.AlertMutingRuleState;
import com.pulumi.signalfx.outputs.AlertMutingRuleFilter;
import com.pulumi.signalfx.outputs.AlertMutingRuleRecurrence;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Splunk Observability Cloud resource for managing alert muting rules. See [Mute Notifications](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html) for more information.
 * 
 * Splunk Observability Cloud currently allows linking an alert muting rule with only one detector ID. Specifying multiple detector IDs makes the muting rule obsolete.
 * 
 * &gt; **WARNING** Splunk Observability Cloud does not allow the start time of a **currently active** muting rule to be modified. Attempting to modify a currently active rule destroys the existing rule and creates a new rule. This might result in the emission of notifications.
 * 
 * ## Example
 * 
 */
@ResourceType(type="signalfx:index/alertMutingRule:AlertMutingRule")
public class AlertMutingRule extends com.pulumi.resources.CustomResource {
    /**
     * The description for this muting rule
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description for this muting rule
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
     * 
     */
    @Export(name="detectors", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> detectors;

    /**
     * @return A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
     * 
     */
    public Output<Optional<List<String>>> detectors() {
        return Codegen.optional(this.detectors);
    }
    @Export(name="effectiveStartTime", refs={Integer.class}, tree="[0]")
    private Output<Integer> effectiveStartTime;

    public Output<Integer> effectiveStartTime() {
        return this.effectiveStartTime;
    }
    /**
     * Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
     * 
     */
    @Export(name="filters", refs={List.class,AlertMutingRuleFilter.class}, tree="[0,1]")
    private Output</* @Nullable */ List<AlertMutingRuleFilter>> filters;

    /**
     * @return Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
     * 
     */
    public Output<Optional<List<AlertMutingRuleFilter>>> filters() {
        return Codegen.optional(this.filters);
    }
    /**
     * Defines the recurrence of the muting rule. Allows setting a recurring muting rule based on specified days or weeks.
     * 
     */
    @Export(name="recurrence", refs={AlertMutingRuleRecurrence.class}, tree="[0]")
    private Output</* @Nullable */ AlertMutingRuleRecurrence> recurrence;

    /**
     * @return Defines the recurrence of the muting rule. Allows setting a recurring muting rule based on specified days or weeks.
     * 
     */
    public Output<Optional<AlertMutingRuleRecurrence>> recurrence() {
        return Codegen.optional(this.recurrence);
    }
    /**
     * Starting time of an alert muting rule as a Unit time stamp in seconds.
     * 
     */
    @Export(name="startTime", refs={Integer.class}, tree="[0]")
    private Output<Integer> startTime;

    /**
     * @return Starting time of an alert muting rule as a Unit time stamp in seconds.
     * 
     */
    public Output<Integer> startTime() {
        return this.startTime;
    }
    /**
     * Stop time of an alert muting rule as a Unix time stamp in seconds.
     * 
     */
    @Export(name="stopTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> stopTime;

    /**
     * @return Stop time of an alert muting rule as a Unix time stamp in seconds.
     * 
     */
    public Output<Optional<Integer>> stopTime() {
        return Codegen.optional(this.stopTime);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AlertMutingRule(java.lang.String name) {
        this(name, AlertMutingRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AlertMutingRule(java.lang.String name, AlertMutingRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AlertMutingRule(java.lang.String name, AlertMutingRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/alertMutingRule:AlertMutingRule", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AlertMutingRule(java.lang.String name, Output<java.lang.String> id, @Nullable AlertMutingRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/alertMutingRule:AlertMutingRule", name, state, makeResourceOptions(options, id), false);
    }

    private static AlertMutingRuleArgs makeArgs(AlertMutingRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AlertMutingRuleArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AlertMutingRule get(java.lang.String name, Output<java.lang.String> id, @Nullable AlertMutingRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AlertMutingRule(name, id, state, options);
    }
}
