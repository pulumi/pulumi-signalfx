// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.log;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.log.TimelineArgs;
import com.pulumi.signalfx.log.inputs.TimelineState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * You can add logs data to your Observability Cloud dashboards without turning your logs into metrics first.
 * A log timeline chart displays timeline visualization in a dashboard and shows you in detail what is happening and why.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.log.Timeline;
 * import com.pulumi.signalfx.log.TimelineArgs;
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
 *         var myLogTimeline = new Timeline(&#34;myLogTimeline&#34;, TimelineArgs.builder()        
 *             .description(&#34;Lorem ipsum dolor sit amet, laudem tibique iracundia at mea. Nam posse dolores ex, nec cu adhuc putent honestatis&#34;)
 *             .programText(&#34;&#34;&#34;
 * logs(filter=field(&#39;message&#39;) == &#39;Transaction processed&#39; and field(&#39;service.name&#39;) == &#39;paymentservice&#39;).publish()
 * 
 *             &#34;&#34;&#34;)
 *             .timeRange(900)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="signalfx:log/timeline:Timeline")
public class Timeline extends com.pulumi.resources.CustomResource {
    /**
     * The connection that the log timeline uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
     * 
     */
    @Export(name="defaultConnection", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultConnection;

    /**
     * @return The connection that the log timeline uses to fetch data. This could be Splunk Enterprise, Splunk Enterprise Cloud or Observability Cloud.
     * 
     */
    public Output<Optional<String>> defaultConnection() {
        return Codegen.optional(this.defaultConnection);
    }
    /**
     * Description of the log timeline.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the log timeline.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    @Export(name="endTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> endTime;

    /**
     * @return Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    public Output<Optional<Integer>> endTime() {
        return Codegen.optional(this.endTime);
    }
    /**
     * Name of the log timeline.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the log timeline.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Signalflow program text for the log timeline. More info at https://dev.splunk.com/observability/docs/.
     * 
     */
    @Export(name="programText", refs={String.class}, tree="[0]")
    private Output<String> programText;

    /**
     * @return Signalflow program text for the log timeline. More info at https://dev.splunk.com/observability/docs/.
     * 
     */
    public Output<String> programText() {
        return this.programText;
    }
    /**
     * Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    @Export(name="startTime", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> startTime;

    /**
     * @return Seconds since epoch. Used for visualization. Conflicts with `time_range`.
     * 
     */
    public Output<Optional<Integer>> startTime() {
        return Codegen.optional(this.startTime);
    }
    /**
     * From when to display data. SignalFx time syntax (e.g. `&#34;-5m&#34;`, `&#34;-1h&#34;`). Conflicts with `start_time` and `end_time`.
     * 
     */
    @Export(name="timeRange", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> timeRange;

    /**
     * @return From when to display data. SignalFx time syntax (e.g. `&#34;-5m&#34;`, `&#34;-1h&#34;`). Conflicts with `start_time` and `end_time`.
     * 
     */
    public Output<Optional<Integer>> timeRange() {
        return Codegen.optional(this.timeRange);
    }
    /**
     * The URL of the log timeline.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL of the log timeline.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Timeline(String name) {
        this(name, TimelineArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Timeline(String name, TimelineArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Timeline(String name, TimelineArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:log/timeline:Timeline", name, args == null ? TimelineArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Timeline(String name, Output<String> id, @Nullable TimelineState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:log/timeline:Timeline", name, state, makeResourceOptions(options, id));
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
    public static Timeline get(String name, Output<String> id, @Nullable TimelineState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Timeline(name, id, state, options);
    }
}
