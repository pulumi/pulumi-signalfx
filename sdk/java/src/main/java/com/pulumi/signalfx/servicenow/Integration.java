// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx.servicenow;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.servicenow.IntegrationArgs;
import com.pulumi.signalfx.servicenow.inputs.IntegrationState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Observability Cloud ServiceNow integrations. For help with this integration see [Integration with ServiceNow](https://docs.splunk.com/Observability/admin/notif-services/servicenow.html).
 * 
 * &gt; **NOTE** When managing integrations, use a session token of an administrator to authenticate the Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you&#39;ll receive a 4xx error.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.signalfx.servicenow.Integration;
 * import com.pulumi.signalfx.servicenow.IntegrationArgs;
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
 *         var serviceNowMyteam = new Integration(&#34;serviceNowMyteam&#34;, IntegrationArgs.builder()        
 *             .alertResolvedPayloadTemplate(&#34;{\&#34;close_notes\&#34;: \&#34;{{{messageTitle}}} (customized close msg)\&#34;}&#34;)
 *             .alertTriggeredPayloadTemplate(&#34;{\&#34;short_description\&#34;: \&#34;{{{messageTitle}}} (customized)\&#34;}&#34;)
 *             .enabled(false)
 *             .instanceName(&#34;myinst.service-now.com&#34;)
 *             .issueType(&#34;Incident&#34;)
 *             .password(&#34;youd0ntsee1t&#34;)
 *             .username(&#34;thisis_me&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="signalfx:servicenow/integration:Integration")
public class Integration extends com.pulumi.resources.CustomResource {
    /**
     * A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     * 
     */
    @Export(name="alertResolvedPayloadTemplate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> alertResolvedPayloadTemplate;

    /**
     * @return A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     * 
     */
    public Output<Optional<String>> alertResolvedPayloadTemplate() {
        return Codegen.optional(this.alertResolvedPayloadTemplate);
    }
    /**
     * A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     * 
     */
    @Export(name="alertTriggeredPayloadTemplate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> alertTriggeredPayloadTemplate;

    /**
     * @return A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     * 
     */
    public Output<Optional<String>> alertTriggeredPayloadTemplate() {
        return Codegen.optional(this.alertTriggeredPayloadTemplate);
    }
    /**
     * Whether the integration is enabled.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return Whether the integration is enabled.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * Name of the ServiceNow instance, for example `myinst.service-now.com`.
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output<String> instanceName;

    /**
     * @return Name of the ServiceNow instance, for example `myinst.service-now.com`.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }
    /**
     * The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
     * 
     */
    @Export(name="issueType", refs={String.class}, tree="[0]")
    private Output<String> issueType;

    /**
     * @return The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
     * 
     */
    public Output<String> issueType() {
        return this.issueType;
    }
    /**
     * Name of the integration.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the integration.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Password used to authenticate the ServiceNow integration.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return Password used to authenticate the ServiceNow integration.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * User name used to authenticate the ServiceNow integration.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return User name used to authenticate the ServiceNow integration.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Integration(String name) {
        this(name, IntegrationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Integration(String name, IntegrationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Integration(String name, IntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:servicenow/integration:Integration", name, args == null ? IntegrationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Integration(String name, Output<String> id, @Nullable IntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:servicenow/integration:Integration", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static Integration get(String name, Output<String> id, @Nullable IntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Integration(name, id, state, options);
    }
}
