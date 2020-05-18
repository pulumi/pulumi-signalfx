// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * SignalFx Azure integrations. For help with this integration see [Monitoring Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure).
 * 
 * > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as signalfx from "@pulumi/signalfx";
 * 
 * const azureMyteam = new signalfx.azure.Integration("azureMyteam", {
 *     enabled: true,
 *     resource: [{
 *         signalfxAzureIntegration: [{
 *             azureMyteamXX: [{
 *                 appId: "YYY",
 *                 enabled: false,
 *                 environment: "azure",
 *                 name: "AzureFoo",
 *                 pollRate: 300,
 *                 secretKey: "XXX",
 *                 services: ["microsoft.sql/servers/elasticpools"],
 *                 subscriptions: ["sub-guid-here"],
 *                 tenantId: "ZZZ",
 *             }],
 *         }],
 *     }],
 * });
 * ```
 * 
 * ## Service Names
 * 
 * > **NOTE** You can use the data source "signalfx.azure.getServices" to specify all services.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/azure_integration.html.markdown.
 */
export class Integration extends pulumi.CustomResource {
    /**
     * Get an existing Integration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationState, opts?: pulumi.CustomResourceOptions): Integration {
        return new Integration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'signalfx:azure/integration:Integration';

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
     * Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * Whether the integration is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
     */
    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * Name of the integration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * AWS poll rate (in seconds). One of `60` or `300`.
     */
    public readonly pollRate!: pulumi.Output<number | undefined>;
    /**
     * Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     */
    public readonly secretKey!: pulumi.Output<string>;
    /**
     * List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
     */
    public readonly services!: pulumi.Output<string[]>;
    /**
     * List of Azure subscriptions that SignalFx should monitor.
     */
    public readonly subscriptions!: pulumi.Output<string[]>;
    /**
     * Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a Integration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationArgs | IntegrationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IntegrationState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["environment"] = state ? state.environment : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pollRate"] = state ? state.pollRate : undefined;
            inputs["secretKey"] = state ? state.secretKey : undefined;
            inputs["services"] = state ? state.services : undefined;
            inputs["subscriptions"] = state ? state.subscriptions : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as IntegrationArgs | undefined;
            if (!args || args.appId === undefined) {
                throw new Error("Missing required property 'appId'");
            }
            if (!args || args.enabled === undefined) {
                throw new Error("Missing required property 'enabled'");
            }
            if (!args || args.secretKey === undefined) {
                throw new Error("Missing required property 'secretKey'");
            }
            if (!args || args.services === undefined) {
                throw new Error("Missing required property 'services'");
            }
            if (!args || args.subscriptions === undefined) {
                throw new Error("Missing required property 'subscriptions'");
            }
            if (!args || args.tenantId === undefined) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["appId"] = args ? args.appId : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pollRate"] = args ? args.pollRate : undefined;
            inputs["secretKey"] = args ? args.secretKey : undefined;
            inputs["services"] = args ? args.services : undefined;
            inputs["subscriptions"] = args ? args.subscriptions : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Integration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Integration resources.
 */
export interface IntegrationState {
    /**
     * Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * Whether the integration is enabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
     */
    readonly environment?: pulumi.Input<string>;
    /**
     * Name of the integration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * AWS poll rate (in seconds). One of `60` or `300`.
     */
    readonly pollRate?: pulumi.Input<number>;
    /**
     * Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     */
    readonly secretKey?: pulumi.Input<string>;
    /**
     * List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
     */
    readonly services?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Azure subscriptions that SignalFx should monitor.
     */
    readonly subscriptions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     */
    readonly tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Integration resource.
 */
export interface IntegrationArgs {
    /**
     * Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
     */
    readonly appId: pulumi.Input<string>;
    /**
     * Whether the integration is enabled.
     */
    readonly enabled: pulumi.Input<boolean>;
    /**
     * What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
     */
    readonly environment?: pulumi.Input<string>;
    /**
     * Name of the integration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * AWS poll rate (in seconds). One of `60` or `300`.
     */
    readonly pollRate?: pulumi.Input<number>;
    /**
     * Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     */
    readonly secretKey: pulumi.Input<string>;
    /**
     * List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
     */
    readonly services: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of Azure subscriptions that SignalFx should monitor.
     */
    readonly subscriptions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
     */
    readonly tenantId: pulumi.Input<string>;
}
