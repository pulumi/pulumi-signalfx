// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.SloArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.SloState;
import com.pulumi.signalfx.outputs.SloInput;
import com.pulumi.signalfx.outputs.SloTarget;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Splunk Observability Cloud slo resource. This can be used to create and manage SLOs.
 * 
 * To learn more about this feature take a look on [documentation for SLO](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/slo-intro.html).
 * 
 * ## Example
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.Slo;
 * import com.pulumi.signalfx.SloArgs;
 * import com.pulumi.signalfx.inputs.SloInputArgs;
 * import com.pulumi.signalfx.inputs.SloTargetArgs;
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
 *         var fooServiceSlo = new Slo(&#34;fooServiceSlo&#34;, SloArgs.builder()        
 *             .description(&#34;SLO monitoring for foo service&#34;)
 *             .input(SloInputArgs.builder()
 *                 .goodEventsLabel(&#34;G&#34;)
 *                 .programText(&#34;&#34;&#34;
 * G = data(&#39;spans.count&#39;, filter=filter(&#39;sf_error&#39;, &#39;false&#39;) and filter(&#39;sf_service&#39;, &#39;foo-service&#39;))
 * T = data(&#39;spans.count&#39;, filter=filter(&#39;sf_service&#39;, &#39;foo-service&#39;))
 *                 &#34;&#34;&#34;)
 *                 .totalEventsLabel(&#34;T&#34;)
 *                 .build())
 *             .target(SloTargetArgs.builder()
 *                 .alertRules(SloTargetAlertRuleArgs.builder()
 *                     .rule(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *                     .type(&#34;BREACH&#34;)
 *                     .build())
 *                 .compliancePeriod(&#34;30d&#34;)
 *                 .slo(95)
 *                 .type(&#34;RollingWindow&#34;)
 *                 .build())
 *             .type(&#34;RequestBased&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Notification format
 * 
 * As Splunk Observability Cloud supports different notification mechanisms, use a comma-delimited string to provide inputs. If you want to specify multiple notifications, each must be a member in the list, like so:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * See [Splunk Observability Cloud Docs](https://dev.splunk.com/observability/reference/api/detectors/latest) for more information.
 * 
 * Here are some example of how to configure each notification type:
 * 
 * ### Email
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * ### Jira
 * 
 * Note that the `credentialId` is the Splunk-provided ID shown after setting up your Jira integration. See also `signalfx.jira.Integration`.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * ### OpsGenie
 * 
 * Note that the `credentialId` is the Splunk-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * ### PagerDuty
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * ### Slack
 * 
 * Exclude the `#` on the channel name:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * ### Team
 * 
 * Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * ### TeamEmail
 * 
 * Sends an email to every member of a team.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * ### Splunk On-Call (formerly VictorOps)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * ### Webhooks
 * 
 * You need to include all the commas even if you only use a credential id.
 * 
 * You can either configure a Webhook to use an existing integration&#39;s credential id:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * Or configure one inline:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
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
 *     }
 * }
 * ```
 * 
 * ## Arguments
 * 
 * * `name` - (Required) Name of the SLO. Each SLO name must be unique within an organization.
 * * `description` - (Optional) Description of the SLO.
 * * `type` - (Required) Type of the SLO. Currently just: `&#34;RequestBased&#34;` is supported.
 * * `input` - (Required) Properties to configure an SLO object inputs
 *   * `program_text` - (Required) SignalFlow program and arguments text strings that define the streams used as successful event count and total event count
 *   * `good_events_label` - (Required) Label used in `&#34;program_text&#34;` that refers to the data block which contains the stream of successful events
 *   * `total_events_label` - (Required) Label used in `&#34;program_text&#34;` that refers to the data block which contains the stream of total events
 * * `target` - (Required) Define target value of the service level indicator in the appropriate time period.
 *   * `type` - (Required) SLO target type can be the following type: `&#34;RollingWindow&#34;`
 *   * `compliance_period` - (Required for `&#34;RollingWindow&#34;` type) Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
 *   * `slo` - (Required) Target value in the form of a percentage
 *   * `alert_rule` - (Required) List of alert rules you want to set for this SLO target. An SLO alert rule of type BREACH is always required.
 *     * `type` - (Required) SLO alert rule can be one of the following types: BREACH, ERROR_BUDGET_LEFT, BURN_RATE. Within an SLO object, you can only specify one SLO alert_rule per type. For example, you can&#39;t specify two alert_rule of type BREACH. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
 *     * `rule` - (Required) Set of rules used for alerting.
 *         * `severity` - (Required) The severity of the rule, must be one of: `&#34;Critical&#34;`, `&#34;Major&#34;`, `&#34;Minor&#34;`, `&#34;Warning&#34;`, `&#34;Info&#34;`.
 *         * `description` - (Optional) Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
 *         * `disabled` - (Optional) When true, notifications and events will not be generated for the detect label. `false` by default.
 *         * `notifications` - (Optional) List of strings specifying where notifications will be sent when an incident occurs. See [Create SLO](https://dev.splunk.com/observability/reference/api/slo/latest#endpoint-create-new-slo) for more info.
 *         * `parameterized_body` - (Optional) Custom notification message body when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
 *         * `parameterized_subject` - (Optional) Custom notification message subject when an alert is triggered. See [Alert message](https://docs.splunk.com/observability/en/alerts-detectors-notifications/create-detectors-for-alerts.html#alert-messages) for more info.
 *         * `runbook_url` - (Optional) URL of page to consult when an alert is triggered. This can be used with custom notification messages.
 *         * `tip` - (Optional) Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
 *         * `parameters` - (Optional) Parameters for the SLO alert rule. Each SLO alert rule type accepts different parameters. If not specified, default parameters are used.
 *           * `fire_lasting` - (Optional) Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: `&#34;BREACH&#34;` and `&#34;ERROR_BUDGET_LEFT&#34;` alert rules use the fireLasting parameter. Default: `&#34;5m&#34;`
 *           * `percent_of_lasting` - (Optional) Percentage of the `&#34;fire_lasting&#34;` duration that the alert condition is met before the alert is triggered. Note: `&#34;BREACH&#34;` and `&#34;ERROR_BUDGET_LEFT&#34;` alert rules use the `&#34;percent_of_lasting&#34;` parameter. Default: `100`
 *           * `percent_error_budget_left` - (Optional) Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: `&#34;ERROR_BUDGET_LEFT&#34;` alert rules use the `&#34;percent_error_budget_left&#34;` parameter. Default: `100`
 *           * `short_window_1` - (Optional) Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of `&#34;long_window_1&#34;`. Note: `&#34;BURN_RATE&#34;` alert rules use the `&#34;short_window_1&#34;` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
 *           * `short_window_2` - (Optional) Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of `&#34;long_window_2&#34;`. Note: `&#34;BURN_RATE&#34;` alert rules use the `&#34;short_window_2&#34;` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
 *           * `long_window_1` - (Optional) Long window 1 used in burn rate alert calculation. This value must be longer than `&#34;short_window_1&#34;` and shorter than 90 days. Note: `&#34;BURN_RATE&#34;` alert rules use the `&#34;long_window_1&#34;` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
 *           * `long_window_2` - (Optional) Long window 2 used in burn rate alert calculation. This value must be longer than `&#34;short_window_2&#34;` and shorter than 90 days. Note: `&#34;BURN_RATE&#34;` alert rules use the `&#34;long_window_2&#34;` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
 *           * `burn_rate_threshold_1` - (Optional) Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: `&#34;BURN_RATE&#34;` alert rules use the `&#34;burn_rate_threshold_1&#34;` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
 *           * `burn_rate_threshold_2` - (Optional) Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: `&#34;BURN_RATE&#34;` alert rules use the `&#34;burn_rate_threshold_2&#34;` parameter. See [SLO alerts](https://docs.splunk.com/observability/en/alerts-detectors-notifications/slo/burn-rate-alerts.html) for more info.
 * 
 */
@ResourceType(type="signalfx:index/slo:Slo")
public class Slo extends com.pulumi.resources.CustomResource {
    /**
     * Description of the SLO
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the SLO
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * SignalFlow program and arguments text strings that define the streams used as successful event count and total event
     * count
     * 
     */
    @Export(name="input", refs={SloInput.class}, tree="[0]")
    private Output<SloInput> input;

    /**
     * @return SignalFlow program and arguments text strings that define the streams used as successful event count and total event
     * count
     * 
     */
    public Output<SloInput> input() {
        return this.input;
    }
    /**
     * Name of the SLO
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the SLO
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Define target value of the service level indicator in the appropriate time period.
     * 
     */
    @Export(name="target", refs={SloTarget.class}, tree="[0]")
    private Output<SloTarget> target;

    /**
     * @return Define target value of the service level indicator in the appropriate time period.
     * 
     */
    public Output<SloTarget> target() {
        return this.target;
    }
    /**
     * Type of the SLO. Currently only RequestBased SLO is supported
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of the SLO. Currently only RequestBased SLO is supported
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Slo(String name) {
        this(name, SloArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Slo(String name, SloArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Slo(String name, SloArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/slo:Slo", name, args == null ? SloArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Slo(String name, Output<String> id, @Nullable SloState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/slo:Slo", name, state, makeResourceOptions(options, id));
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
    public static Slo get(String name, Output<String> id, @Nullable SloState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Slo(name, id, state, options);
    }
}
