// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.DashboardGroupArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.DashboardGroupState;
import com.pulumi.signalfx.outputs.DashboardGroupDashboard;
import com.pulumi.signalfx.outputs.DashboardGroupImportQualifier;
import com.pulumi.signalfx.outputs.DashboardGroupPermission;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * In the Splunk Observability Cloud web UI, a [dashboard group](https://developers.signalfx.com/dashboard_groups_reference.html) is a collection of dashboards.
 * 
 * Dashboard groups cannot be accessed directly. You can access them through a dashboard within a group.
 * 
 * &gt; **NOTE** When you want to change or remove write permissions for a user other than yourself regarding dashboard groups, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
 * 
 * ## Example
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.DashboardGroup;
 * import com.pulumi.signalfx.DashboardGroupArgs;
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
 *         var mydashboardgroup0 = new DashboardGroup("mydashboardgroup0", DashboardGroupArgs.builder()
 *             .name("My team dashboard group")
 *             .description("Cool dashboard group")
 *             .authorizedWriterTeams(mycoolteam.id())
 *             .authorizedWriterUsers("abc123")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Example with permissions
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.DashboardGroup;
 * import com.pulumi.signalfx.DashboardGroupArgs;
 * import com.pulumi.signalfx.inputs.DashboardGroupPermissionArgs;
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
 *         var mydashboardgroupWithpermissions = new DashboardGroup("mydashboardgroupWithpermissions", DashboardGroupArgs.builder()
 *             .name("My team dashboard group")
 *             .description("Cool dashboard group")
 *             .permissions(            
 *                 DashboardGroupPermissionArgs.builder()
 *                     .principalId("abc123")
 *                     .principalType("ORG")
 *                     .actions("READ")
 *                     .build(),
 *                 DashboardGroupPermissionArgs.builder()
 *                     .principalId("abc456")
 *                     .principalType("USER")
 *                     .actions(                    
 *                         "READ",
 *                         "WRITE")
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Example With mirrored dashboards
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.DashboardGroup;
 * import com.pulumi.signalfx.DashboardGroupArgs;
 * import com.pulumi.signalfx.inputs.DashboardGroupDashboardArgs;
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
 *         var mydashboardgroupWithmirrors = new DashboardGroup("mydashboardgroupWithmirrors", DashboardGroupArgs.builder()
 *             .name("My team dashboard group")
 *             .description("Cool dashboard group")
 *             .dashboards(DashboardGroupDashboardArgs.builder()
 *                 .dashboardId(gcDashboard.id())
 *                 .nameOverride("GC For My Service")
 *                 .descriptionOverride("Garbage Collection dashboard maintained by JVM team")
 *                 .filterOverrides(DashboardGroupDashboardFilterOverrideArgs.builder()
 *                     .property("service")
 *                     .values("myservice")
 *                     .negated(false)
 *                     .build())
 *                 .variableOverrides(DashboardGroupDashboardVariableOverrideArgs.builder()
 *                     .property("region")
 *                     .values("us-west1")
 *                     .valuesSuggesteds(                    
 *                         "us-west-1",
 *                         "us-east-1")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="signalfx:index/dashboardGroup:DashboardGroup")
public class DashboardGroup extends com.pulumi.resources.CustomResource {
    /**
     * Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     * @deprecated
     * Please use permissions field now
     * 
     */
    @Deprecated /* Please use permissions field now */
    @Export(name="authorizedWriterTeams", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> authorizedWriterTeams;

    /**
     * @return Team IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s team (or user id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     */
    public Output<Optional<List<String>>> authorizedWriterTeams() {
        return Codegen.optional(this.authorizedWriterTeams);
    }
    /**
     * User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     * @deprecated
     * Please use permissions field now
     * 
     */
    @Deprecated /* Please use permissions field now */
    @Export(name="authorizedWriterUsers", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> authorizedWriterUsers;

    /**
     * @return User IDs that have write access to this dashboard group. Remember to use an admin&#39;s token if using this feature and to include that admin&#39;s user id (or team id in `authorized_writer_teams`). **Note:** Deprecated use `permissions` instead.
     * 
     */
    public Output<Optional<List<String>>> authorizedWriterUsers() {
        return Codegen.optional(this.authorizedWriterUsers);
    }
    /**
     * [Mirrored dashboards](https://docs.splunk.com/observability/en/data-visualization/dashboards/dashboard-share-clone-mirror.html#mirror-dashboard) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    @Export(name="dashboards", refs={List.class,DashboardGroupDashboard.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DashboardGroupDashboard>> dashboards;

    /**
     * @return [Mirrored dashboards](https://docs.splunk.com/observability/en/data-visualization/dashboards/dashboard-share-clone-mirror.html#mirror-dashboard) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    public Output<Optional<List<DashboardGroupDashboard>>> dashboards() {
        return Codegen.optional(this.dashboards);
    }
    /**
     * Description of the dashboard group.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the dashboard group.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="importQualifiers", refs={List.class,DashboardGroupImportQualifier.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DashboardGroupImportQualifier>> importQualifiers;

    public Output<Optional<List<DashboardGroupImportQualifier>>> importQualifiers() {
        return Codegen.optional(this.importQualifiers);
    }
    /**
     * Name of the dashboard group.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the dashboard group.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) List of read and write permission configuration to specify which user, team, and organization can view and/or edit your dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    @Export(name="permissions", refs={List.class,DashboardGroupPermission.class}, tree="[0,1]")
    private Output<List<DashboardGroupPermission>> permissions;

    /**
     * @return [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) List of read and write permission configuration to specify which user, team, and organization can view and/or edit your dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    public Output<List<DashboardGroupPermission>> permissions() {
        return this.permissions;
    }
    /**
     * Team IDs to associate the dashboard group to.
     * 
     */
    @Export(name="teams", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> teams;

    /**
     * @return Team IDs to associate the dashboard group to.
     * 
     */
    public Output<Optional<List<String>>> teams() {
        return Codegen.optional(this.teams);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DashboardGroup(java.lang.String name) {
        this(name, DashboardGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DashboardGroup(java.lang.String name, @Nullable DashboardGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DashboardGroup(java.lang.String name, @Nullable DashboardGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/dashboardGroup:DashboardGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DashboardGroup(java.lang.String name, Output<java.lang.String> id, @Nullable DashboardGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/dashboardGroup:DashboardGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static DashboardGroupArgs makeArgs(@Nullable DashboardGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DashboardGroupArgs.Empty : args;
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
    public static DashboardGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable DashboardGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DashboardGroup(name, id, state, options);
    }
}
