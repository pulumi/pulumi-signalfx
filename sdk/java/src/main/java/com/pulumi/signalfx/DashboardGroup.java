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
 * In the SignalFx web UI, a [dashboard group](https://developers.signalfx.com/dashboard_groups_reference.html) is a collection of dashboards.
 * 
 * &gt; **NOTE** Dashboard groups cannot be accessed directly, but just via a dashboard contained in them. This is the reason why make show won&#39;t show any of yours dashboard groups.
 * 
 * &gt; **NOTE** When you want to &#34;Change or remove write permissions for a user other than yourself&#34; regarding dashboard groups, use a session token of an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
 * 
 * ## Example Usage
 * ```java
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
 *         var mydashboardgroup0 = new DashboardGroup(&#34;mydashboardgroup0&#34;, DashboardGroupArgs.builder()        
 *             .description(&#34;Cool dashboard group&#34;)
 *             .authorizedWriterTeams(signalfx_team.mycoolteam().id())
 *             .authorizedWriterUsers(&#34;abc123&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With Permissions
 * ```java
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
 *         var mydashboardgroupWithpermissions = new DashboardGroup(&#34;mydashboardgroupWithpermissions&#34;, DashboardGroupArgs.builder()        
 *             .description(&#34;Cool dashboard group&#34;)
 *             .permissions(            
 *                 DashboardGroupPermissionArgs.builder()
 *                     .actions(&#34;READ&#34;)
 *                     .principalId(&#34;abc123&#34;)
 *                     .principalType(&#34;ORG&#34;)
 *                     .build(),
 *                 DashboardGroupPermissionArgs.builder()
 *                     .actions(                    
 *                         &#34;READ&#34;,
 *                         &#34;WRITE&#34;)
 *                     .principalId(&#34;abc456&#34;)
 *                     .principalType(&#34;USER&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### With Mirrored Dashboards
 * ```java
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
 *         var mydashboardgroupWithmirrors = new DashboardGroup(&#34;mydashboardgroupWithmirrors&#34;, DashboardGroupArgs.builder()        
 *             .description(&#34;Cool dashboard group&#34;)
 *             .dashboards(DashboardGroupDashboardArgs.builder()
 *                 .dashboardId(signalfx_dashboard.gc_dashboard().id())
 *                 .nameOverride(&#34;GC For My Service&#34;)
 *                 .descriptionOverride(&#34;Garbage Collection dashboard maintained by JVM team&#34;)
 *                 .filterOverrides(DashboardGroupDashboardFilterOverrideArgs.builder()
 *                     .property(&#34;service&#34;)
 *                     .values(&#34;myservice&#34;)
 *                     .negated(false)
 *                     .build())
 *                 .variableOverrides(DashboardGroupDashboardVariableOverrideArgs.builder()
 *                     .property(&#34;region&#34;)
 *                     .values(&#34;us-west1&#34;)
 *                     .valuesSuggesteds(                    
 *                         &#34;us-west-1&#34;,
 *                         &#34;us-east-1&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
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
     * [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
     * 
     */
    @Export(name="dashboards", refs={List.class,DashboardGroupDashboard.class}, tree="[0,1]")
    private Output</* @Nullable */ List<DashboardGroupDashboard>> dashboards;

    /**
     * @return [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
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
    public DashboardGroup(String name) {
        this(name, DashboardGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DashboardGroup(String name, @Nullable DashboardGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DashboardGroup(String name, @Nullable DashboardGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/dashboardGroup:DashboardGroup", name, args == null ? DashboardGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DashboardGroup(String name, Output<String> id, @Nullable DashboardGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/dashboardGroup:DashboardGroup", name, state, makeResourceOptions(options, id));
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
    public static DashboardGroup get(String name, Output<String> id, @Nullable DashboardGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DashboardGroup(name, id, state, options);
    }
}
