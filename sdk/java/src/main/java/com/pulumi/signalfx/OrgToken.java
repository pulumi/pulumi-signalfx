// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.signalfx.OrgTokenArgs;
import com.pulumi.signalfx.Utilities;
import com.pulumi.signalfx.inputs.OrgTokenState;
import com.pulumi.signalfx.outputs.OrgTokenDpmLimits;
import com.pulumi.signalfx.outputs.OrgTokenHostOrUsageLimits;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manage Splunk Observability Cloud org tokens.
 * 
 * &gt; **NOTE** When managing Org tokens, use a session token of an administrator to authenticate the Splunk Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
 * 
 * ## Example
 * 
 * ## Arguments
 * 
 * The following arguments are supported in the resource block:
 * 
 * * `name` - (Required) Name of the token.
 * * `description` - (Optional) Description of the token.
 * * `disabled` - (Optional) Flag that controls enabling the token. If set to `true`, the token is disabled, and you can&#39;t use it for authentication. Defaults to `false`.
 * * `secret` - The secret token created by the API. You cannot set this value.
 * * `notifications` - (Optional) Where to send notifications about this token&#39;s limits. See the Notification Format laid out in detectors.
 * * `host_or_usage_limits` - (Optional) Specify Usage-based limits for this token.
 *   * `host_limit` - (Optional) Max number of hosts that can use this token
 *   * `host_notification_threshold` - (Optional) Notification threshold for hosts
 *   * `container_limit` - (Optional) Max number of Docker containers that can use this token
 *   * `container_notification_threshold` - (Optional) Notification threshold for Docker containers
 *   * `custom_metrics_limit` - (Optional) Max number of custom metrics that can be sent with this token
 *   * `custom_metrics_notification_threshold` - (Optional) Notification threshold for custom metrics
 *   * `high_res_metrics_limit` - (Optional) Max number of hi-res metrics that can be sent with this toke
 *   * `high_res_metrics_notification_threshold` - (Optional) Notification threshold for hi-res metrics
 * * `dpm_limits` (Optional) Specify DPM-based limits for this token.
 *   * `dpm_notification_threshold` - (Optional) DPM level at which Splunk Observability Cloud sends the notification for this token. If you don&#39;t specify a notification, Splunk Observability Cloud sends the generic notification.
 *   * `dpm_limit` - (Required) The datapoints per minute (dpm) limit for this token. If you exceed this limit, Splunk Observability Cloud sends out an alert.
 * 
 * ## Attributes
 * 
 * In a addition to all arguments above, the following attributes are exported:
 * 
 * * `id` - The ID of the token.
 * * `secret` - The assigned token.
 * 
 */
@ResourceType(type="signalfx:index/orgToken:OrgToken")
public class OrgToken extends com.pulumi.resources.CustomResource {
    /**
     * Authentication scope, ex: INGEST, API, RUM ... (Optional)
     * 
     */
    @Export(name="authScopes", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> authScopes;

    /**
     * @return Authentication scope, ex: INGEST, API, RUM ... (Optional)
     * 
     */
    public Output<List<String>> authScopes() {
        return this.authScopes;
    }
    /**
     * Description of the token (Optional)
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the token (Optional)
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Flag that controls enabling the token. If set to `true`, the token is disabled, and you can&#39;t use it for authentication.
     * Defaults to `false`
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disabled;

    /**
     * @return Flag that controls enabling the token. If set to `true`, the token is disabled, and you can&#39;t use it for authentication.
     * Defaults to `false`
     * 
     */
    public Output<Optional<Boolean>> disabled() {
        return Codegen.optional(this.disabled);
    }
    @Export(name="dpmLimits", refs={OrgTokenDpmLimits.class}, tree="[0]")
    private Output</* @Nullable */ OrgTokenDpmLimits> dpmLimits;

    public Output<Optional<OrgTokenDpmLimits>> dpmLimits() {
        return Codegen.optional(this.dpmLimits);
    }
    @Export(name="hostOrUsageLimits", refs={OrgTokenHostOrUsageLimits.class}, tree="[0]")
    private Output</* @Nullable */ OrgTokenHostOrUsageLimits> hostOrUsageLimits;

    public Output<Optional<OrgTokenHostOrUsageLimits>> hostOrUsageLimits() {
        return Codegen.optional(this.hostOrUsageLimits);
    }
    /**
     * Name of the token
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the token
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See
     * https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     * 
     */
    @Export(name="notifications", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> notifications;

    /**
     * @return List of strings specifying where notifications will be sent when an incident occurs. See
     * https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     * 
     */
    public Output<Optional<List<String>>> notifications() {
        return Codegen.optional(this.notifications);
    }
    @Export(name="secret", refs={String.class}, tree="[0]")
    private Output<String> secret;

    public Output<String> secret() {
        return this.secret;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OrgToken(String name) {
        this(name, OrgTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OrgToken(String name, @Nullable OrgTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OrgToken(String name, @Nullable OrgTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/orgToken:OrgToken", name, args == null ? OrgTokenArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OrgToken(String name, Output<String> id, @Nullable OrgTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("signalfx:index/orgToken:OrgToken", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secret"
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
    public static OrgToken get(String name, Output<String> id, @Nullable OrgTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OrgToken(name, id, state, options);
    }
}
