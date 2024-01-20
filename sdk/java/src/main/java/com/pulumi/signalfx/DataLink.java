// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.DataLinkArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.DataLinkState;
import com.pulumi.signalfx.outputs.DataLinkTargetExternalUrl;
import com.pulumi.signalfx.outputs.DataLinkTargetSignalfxDashboard;
import com.pulumi.signalfx.outputs.DataLinkTargetSplunk;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manage Splunk Observability Cloud [Data Links](https://docs.signalfx.com/en/latest/managing/data-links.html).
 * 
 * ## Example
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.DataLink;
 * import com.pulumi.signalfx.DataLinkArgs;
 * import com.pulumi.signalfx.inputs.DataLinkTargetSignalfxDashboardArgs;
 * import com.pulumi.signalfx.inputs.DataLinkTargetExternalUrlArgs;
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
 *         var myDataLink = new DataLink(&#34;myDataLink&#34;, DataLinkArgs.builder()        
 *             .propertyName(&#34;pname&#34;)
 *             .propertyValue(&#34;pvalue&#34;)
 *             .targetSignalfxDashboards(DataLinkTargetSignalfxDashboardArgs.builder()
 *                 .isDefault(true)
 *                 .name(&#34;sfx_dash&#34;)
 *                 .dashboardGroupId(signalfx_dashboard_group.mydashboardgroup0().id())
 *                 .dashboardId(signalfx_dashboard.mydashboard0().id())
 *                 .build())
 *             .build());
 * 
 *         var myDataLinkDash = new DataLink(&#34;myDataLinkDash&#34;, DataLinkArgs.builder()        
 *             .contextDashboardId(signalfx_dashboard.mydashboard0().id())
 *             .propertyName(&#34;pname2&#34;)
 *             .propertyValue(&#34;pvalue&#34;)
 *             .targetExternalUrls(DataLinkTargetExternalUrlArgs.builder()
 *                 .name(&#34;ex_url&#34;)
 *                 .timeFormat(&#34;ISO8601&#34;)
 *                 .url(&#34;https://www.example.com&#34;)
 *                 .propertyKeyMapping(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Arguments
 * 
 * The following arguments are supported in the resource block:
 * 
 * * `property_name` - (Optional) Name (key) of the metadata that&#39;s the trigger of a data link. If you specify `property_value`, you must specify `property_name`.
 * * `property_value` - (Optional) Value of the metadata that&#39;s the trigger of a data link. If you specify this property, you must also specify `property_name`.
 * * `context_dashboard_id` - (Optional) If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
 * * `target_external_url` - (Optional) Link to an external URL
 *   * `name` (Required) User-assigned target name. Use this value to differentiate between the link targets for a data link object.
 *   * `url`- (Required) URL string for a Splunk instance or external system data link target. [See the supported template variables](https://dev.splunk.com/observability/docs/administration/datalinks/).
 *   * `time_format` - (Optional) [Designates the format](https://dev.splunk.com/observability/docs/administration/datalinks/) of `minimum_time_window` in the same data link target object. Must be one of `&#34;ISO8601&#34;`, `&#34;EpochSeconds&#34;` or `&#34;Epoch&#34;` (which is milliseconds). Defaults to `&#34;ISO8601&#34;`.
 *   * `minimum_time_window` - (Optional) The [minimum time window](https://dev.splunk.com/observability/docs/administration/datalinks/) for a search sent to an external site. Defaults to `6000`
 *   * `property_key_mapping` - Describes the relationship between Splunk Observability Cloud metadata keys and external system properties when the key names are different.
 * * `target_signalfx_dashboard` - (Optional) Link to a Splunk Observability Cloud dashboard
 *   * `name` (Required) User-assigned target name. Use this value to differentiate between the link targets for a data link object.
 *   * `is_default` - (Optional) Flag that designates a target as the default for a data link object. `true` by default
 *   * `dashboard_id` - (Required) SignalFx-assigned ID of the dashboard link target
 *   * `dashboard_group_id` - (Required) SignalFx-assigned ID of the dashboard link target&#39;s dashboard group
 * * `target_splunk` - (Optional) Link to an external URL
 *   * `name` (Required) User-assigned target name. Use this value to differentiate between the link targets for a data link object.
 *   * `property_key_mapping` - Describes the relationship between Splunk Observability Cloud metadata keys and external system properties when the key names are different.
 * 
 * ## Attributes
 * 
 * In a addition to all arguments above, the following attributes are exported:
 * 
 * * `id` - The ID of the link.
 * 
 */
@ResourceType(type="signalfx:index/dataLink:DataLink")
public class DataLink extends com.pulumi.resources.CustomResource {
    /**
     * The dashobard ID to which this data link will be applied
     * 
     */
    @Export(name="contextDashboardId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> contextDashboardId;

    /**
     * @return The dashobard ID to which this data link will be applied
     * 
     */
    public Output<Optional<String>> contextDashboardId() {
        return Codegen.optional(this.contextDashboardId);
    }
    /**
     * Name (key) of the metadata that&#39;s the trigger of a data link. If you specify `property_value`, you must specify
     * `property_name`.
     * 
     */
    @Export(name="propertyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> propertyName;

    /**
     * @return Name (key) of the metadata that&#39;s the trigger of a data link. If you specify `property_value`, you must specify
     * `property_name`.
     * 
     */
    public Output<Optional<String>> propertyName() {
        return Codegen.optional(this.propertyName);
    }
    /**
     * Value of the metadata that&#39;s the trigger of a data link. If you specify this property, you must also specify
     * `property_name`.
     * 
     */
    @Export(name="propertyValue", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> propertyValue;

    /**
     * @return Value of the metadata that&#39;s the trigger of a data link. If you specify this property, you must also specify
     * `property_name`.
     * 
     */
    public Output<Optional<String>> propertyValue() {
        return Codegen.optional(this.propertyValue);
    }
    /**
     * Link to an external URL
     * 
     */
    @Export(name="targetExternalUrls", refs={List.class,DataLinkTargetExternalUrl.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DataLinkTargetExternalUrl>> targetExternalUrls;

    /**
     * @return Link to an external URL
     * 
     */
    public Output<Optional<List<DataLinkTargetExternalUrl>>> targetExternalUrls() {
        return Codegen.optional(this.targetExternalUrls);
    }
    /**
     * Link to a Splunk Observability Cloud dashboard
     * 
     */
    @Export(name="targetSignalfxDashboards", refs={List.class,DataLinkTargetSignalfxDashboard.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DataLinkTargetSignalfxDashboard>> targetSignalfxDashboards;

    /**
     * @return Link to a Splunk Observability Cloud dashboard
     * 
     */
    public Output<Optional<List<DataLinkTargetSignalfxDashboard>>> targetSignalfxDashboards() {
        return Codegen.optional(this.targetSignalfxDashboards);
    }
    /**
     * Link to a Splunk instance
     * 
     */
    @Export(name="targetSplunks", refs={List.class,DataLinkTargetSplunk.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DataLinkTargetSplunk>> targetSplunks;

    /**
     * @return Link to a Splunk instance
     * 
     */
    public Output<Optional<List<DataLinkTargetSplunk>>> targetSplunks() {
        return Codegen.optional(this.targetSplunks);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DataLink(String name) {
        this(name, DataLinkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DataLink(String name, @Nullable DataLinkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DataLink(String name, @Nullable DataLinkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/dataLink:DataLink", name, args == null ? DataLinkArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DataLink(String name, Output<String> id, @Nullable DataLinkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/dataLink:DataLink", name, state, makeResourceOptions(options, id));
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
    public static DataLink get(String name, Output<String> id, @Nullable DataLinkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DataLink(name, id, state, options);
    }
}
