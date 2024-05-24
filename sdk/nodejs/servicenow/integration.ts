// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ServiceNow integrations. For help with this integration see [Integration with ServiceNow](https://docs.splunk.com/observability/en/admin/notif-services/servicenow.html).
 *
 * > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
 *
 * ## Example
 */
export class Integration extends pulumi.CustomResource {
    /**
     * Get an existing Integration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationState, opts?: pulumi.CustomResourceOptions): Integration {
        return new Integration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:servicenow/integration:Integration';

    /**
     * Returns true if the given object is an instance of Integration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Integration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Integration.__pulumiType;
    }

    /**
     * A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     */
    public readonly alertResolvedPayloadTemplate!: pulumi.Output<string | undefined>;
    /**
     * A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     */
    public readonly alertTriggeredPayloadTemplate!: pulumi.Output<string | undefined>;
    /**
     * Whether the integration is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Name of the ServiceNow instance, for example `myinst.service-now.com`.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
     */
    public readonly issueType!: pulumi.Output<string>;
    /**
     * Name of the integration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password used to authenticate the ServiceNow integration.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * User name used to authenticate the ServiceNow integration.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a Integration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationArgs | IntegrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationState | undefined;
            resourceInputs["alertResolvedPayloadTemplate"] = state ? state.alertResolvedPayloadTemplate : undefined;
            resourceInputs["alertTriggeredPayloadTemplate"] = state ? state.alertTriggeredPayloadTemplate : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["issueType"] = state ? state.issueType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as IntegrationArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.issueType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issueType'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["alertResolvedPayloadTemplate"] = args ? args.alertResolvedPayloadTemplate : undefined;
            resourceInputs["alertTriggeredPayloadTemplate"] = args ? args.alertTriggeredPayloadTemplate : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["issueType"] = args ? args.issueType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Integration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Integration resources.
 */
export interface IntegrationState {
    /**
     * A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     */
    alertResolvedPayloadTemplate?: pulumi.Input<string>;
    /**
     * A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     */
    alertTriggeredPayloadTemplate?: pulumi.Input<string>;
    /**
     * Whether the integration is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Name of the ServiceNow instance, for example `myinst.service-now.com`.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
     */
    issueType?: pulumi.Input<string>;
    /**
     * Name of the integration.
     */
    name?: pulumi.Input<string>;
    /**
     * Password used to authenticate the ServiceNow integration.
     */
    password?: pulumi.Input<string>;
    /**
     * User name used to authenticate the ServiceNow integration.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Integration resource.
 */
export interface IntegrationArgs {
    /**
     * A template that Observability Cloud uses to create the ServiceNow PUT JSON payloads when an alert is cleared in ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     */
    alertResolvedPayloadTemplate?: pulumi.Input<string>;
    /**
     * A template that Observability Cloud uses to create the ServiceNow POST JSON payloads when an alert sends a notification to ServiceNow. Use this optional field to send the values of Observability Cloud alert properties to specific fields in ServiceNow. See [API reference](https://dev.splunk.com/observability/reference/api/integrations/latest) for details.
     */
    alertTriggeredPayloadTemplate?: pulumi.Input<string>;
    /**
     * Whether the integration is enabled.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Name of the ServiceNow instance, for example `myinst.service-now.com`.
     */
    instanceName: pulumi.Input<string>;
    /**
     * The type of issue in standard ITIL terminology. The allowed values are `Incident` and `Problem`.
     */
    issueType: pulumi.Input<string>;
    /**
     * Name of the integration.
     */
    name?: pulumi.Input<string>;
    /**
     * Password used to authenticate the ServiceNow integration.
     */
    password: pulumi.Input<string>;
    /**
     * User name used to authenticate the ServiceNow integration.
     */
    username: pulumi.Input<string>;
}
