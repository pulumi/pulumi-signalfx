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
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.AlertMutingRule;
 * import com.pulumi.signalfx.AlertMutingRuleArgs;
 * import com.pulumi.signalfx.inputs.AlertMutingRuleFilterArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var roolMooterOne = new AlertMutingRule(&#34;roolMooterOne&#34;, AlertMutingRuleArgs.builder()        
 *             .description(&#34;mooted it NEW&#34;)
 *             .startTime(1573063243)
 *             .stopTime(0)
 *             .detectors(signalfx_detector.some_detector_id())
 *             .filters(AlertMutingRuleFilterArgs.builder()
 *                 .property(&#34;foo&#34;)
 *                 .propertyValue(&#34;bar&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Arguments
 * 
 * * `description` - (Required) The description for this muting rule
 * * `start_time` - (Required) Starting time of an alert muting rule as a Unit time stamp in seconds.
 * * `stop_time` - (Optional) Stop time of an alert muting rule as a Unix time stamp in seconds.
 * * `detectors` - (Optional) A convenience attribute that associated this muting rule with specific detector IDs. Currently, only one ID is supported.
 * * `filter` - (Optional) Filters for this rule. See [Creating muting rules from scratch](https://docs.splunk.com/Observability/alerts-detectors-notifications/mute-notifications.html#rule-from-scratch) for more information.
 *   * `property` - (Required) The property to filter.
 *   * `property_value` - (Required) The property value to filter.
 *   * `negated` - (Optional) Determines if this is a &#34;not&#34; filter. Defaults to `false`.
 * 
 * ## Attributes
 * 
 * In a addition to all arguments above, the following attributes are exported:
 * 
 * * `id` - The ID of the alert muting rule.
 * * `effective_start_time`
 * 
 */
@ResourceType(type="signalfx:index/alertMutingRule:AlertMutingRule")
public class AlertMutingRule extends com.pulumi.resources.CustomResource {
    /**
     * description of the rule
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return description of the rule
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * detectors to which this muting rule applies
     * 
     */
    @Export(name="detectors", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> detectors;

    /**
     * @return detectors to which this muting rule applies
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
     * list of alert muting filters for this rule
     * 
     */
    @Export(name="filters", refs={List.class,AlertMutingRuleFilter.class}, tree="[0,1]")
    private Output</* @Nullable */ List<AlertMutingRuleFilter>> filters;

    /**
     * @return list of alert muting filters for this rule
     * 
     */
    public Output<Optional<List<AlertMutingRuleFilter>>> filters() {
        return Codegen.optional(this.filters);
    }
    /**
     * starting time of an alert muting rule as a Unix timestamp, in seconds
     * 
     */
    @Export(name="startTime", refs={Integer.class}, tree="[0]")
    private Output<Integer> startTime;

    /**
     * @return starting time of an alert muting rule as a Unix timestamp, in seconds
     * 
     */
    public Output<Integer> startTime() {
        return this.startTime;
    }
    /**
     * stop time of an alert muting rule as a Unix timestamp, in seconds
     * 
     */
    @Export(name="stopTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> stopTime;

    /**
     * @return stop time of an alert muting rule as a Unix timestamp, in seconds
     * 
     */
    public Output<Optional<Integer>> stopTime() {
        return Codegen.optional(this.stopTime);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AlertMutingRule(String name) {
        this(name, AlertMutingRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AlertMutingRule(String name, AlertMutingRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AlertMutingRule(String name, AlertMutingRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/alertMutingRule:AlertMutingRule", name, args == null ? AlertMutingRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AlertMutingRule(String name, Output<String> id, @Nullable AlertMutingRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/alertMutingRule:AlertMutingRule", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static AlertMutingRule get(String name, Output<String> id, @Nullable AlertMutingRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AlertMutingRule(name, id, state, options);
    }
}
