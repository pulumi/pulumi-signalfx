// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SignalFx Azure integrations. For help with this integration see [Monitoring Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure).
//
// > **NOTE** When managing integrations use a session token for an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator].(https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-signalfx/sdk/v5/go/signalfx/azure"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azure.NewIntegration(ctx, "azureMyteam", &azure.IntegrationArgs{
// 			AppId: pulumi.String("YYY"),
// 			CustomNamespacesPerServices: azure.IntegrationCustomNamespacesPerServiceArray{
// 				&azure.IntegrationCustomNamespacesPerServiceArgs{
// 					Namespaces: pulumi.StringArray{
// 						pulumi.String("monitoringAgent"),
// 						pulumi.String("customNamespace"),
// 					},
// 					Service: pulumi.String("Microsoft.Compute/virtualMachines"),
// 				},
// 			},
// 			Enabled:     pulumi.Bool(true),
// 			Environment: pulumi.String("azure"),
// 			PollRate:    pulumi.Int(300),
// 			SecretKey:   pulumi.String("XXX"),
// 			Services: pulumi.StringArray{
// 				pulumi.String("microsoft.sql/servers/elasticpools"),
// 			},
// 			Subscriptions: pulumi.StringArray{
// 				pulumi.String("sub-guid-here"),
// 			},
// 			TenantId: pulumi.String("ZZZ"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Service Names
//
// > **NOTE** You can use the data source "azure.getServices" to specify all services.
type Integration struct {
	pulumi.CustomResourceState

	// Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `syncGuestOsNamespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
	CustomNamespacesPerServices IntegrationCustomNamespacesPerServiceArrayOutput `pulumi:"customNamespacesPerServices"`
	// Whether the integration is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// Name of the integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// A named token to use for ingest
	NamedToken pulumi.StringPtrOutput `pulumi:"namedToken"`
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate pulumi.IntPtrOutput `pulumi:"pollRate"`
	// Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services pulumi.StringArrayOutput `pulumi:"services"`
	// List of Azure subscriptions that SignalFx should monitor.
	Subscriptions pulumi.StringArrayOutput `pulumi:"subscriptions"`
	// If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
	SyncGuestOsNamespaces pulumi.BoolPtrOutput `pulumi:"syncGuestOsNamespaces"`
	// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	if args.Services == nil {
		return nil, errors.New("invalid value for required argument 'Services'")
	}
	if args.Subscriptions == nil {
		return nil, errors.New("invalid value for required argument 'Subscriptions'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource Integration
	err := ctx.RegisterResource("signalfx:azure/integration:Integration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegration gets an existing Integration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationState, opts ...pulumi.ResourceOption) (*Integration, error) {
	var resource Integration
	err := ctx.ReadResource("signalfx:azure/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
	AppId *string `pulumi:"appId"`
	// Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `syncGuestOsNamespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
	CustomNamespacesPerServices []IntegrationCustomNamespacesPerService `pulumi:"customNamespacesPerServices"`
	// Whether the integration is enabled.
	Enabled *bool `pulumi:"enabled"`
	// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
	Environment *string `pulumi:"environment"`
	// Name of the integration.
	Name *string `pulumi:"name"`
	// A named token to use for ingest
	NamedToken *string `pulumi:"namedToken"`
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate *int `pulumi:"pollRate"`
	// Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	SecretKey *string `pulumi:"secretKey"`
	// List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services []string `pulumi:"services"`
	// List of Azure subscriptions that SignalFx should monitor.
	Subscriptions []string `pulumi:"subscriptions"`
	// If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
	SyncGuestOsNamespaces *bool `pulumi:"syncGuestOsNamespaces"`
	// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	TenantId *string `pulumi:"tenantId"`
}

type IntegrationState struct {
	// Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
	AppId pulumi.StringPtrInput
	// Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `syncGuestOsNamespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
	CustomNamespacesPerServices IntegrationCustomNamespacesPerServiceArrayInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolPtrInput
	// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
	Environment pulumi.StringPtrInput
	// Name of the integration.
	Name pulumi.StringPtrInput
	// A named token to use for ingest
	NamedToken pulumi.StringPtrInput
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate pulumi.IntPtrInput
	// Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	SecretKey pulumi.StringPtrInput
	// List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services pulumi.StringArrayInput
	// List of Azure subscriptions that SignalFx should monitor.
	Subscriptions pulumi.StringArrayInput
	// If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
	SyncGuestOsNamespaces pulumi.BoolPtrInput
	// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	TenantId pulumi.StringPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
	AppId string `pulumi:"appId"`
	// Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `syncGuestOsNamespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
	CustomNamespacesPerServices []IntegrationCustomNamespacesPerService `pulumi:"customNamespacesPerServices"`
	// Whether the integration is enabled.
	Enabled bool `pulumi:"enabled"`
	// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
	Environment *string `pulumi:"environment"`
	// Name of the integration.
	Name *string `pulumi:"name"`
	// A named token to use for ingest
	NamedToken *string `pulumi:"namedToken"`
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate *int `pulumi:"pollRate"`
	// Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	SecretKey string `pulumi:"secretKey"`
	// List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services []string `pulumi:"services"`
	// List of Azure subscriptions that SignalFx should monitor.
	Subscriptions []string `pulumi:"subscriptions"`
	// If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
	SyncGuestOsNamespaces *bool `pulumi:"syncGuestOsNamespaces"`
	// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
	AppId pulumi.StringInput
	// Allows for more fine-grained control of syncing of custom namespaces, should the boolean convenience parameter `syncGuestOsNamespaces` be not enough. The customer may specify a map of services to custom namespaces. If they do so, for each service which is a key in this map, we will attempt to sync metrics from namespaces in the value list in addition to the default namespaces.
	CustomNamespacesPerServices IntegrationCustomNamespacesPerServiceArrayInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolInput
	// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
	Environment pulumi.StringPtrInput
	// Name of the integration.
	Name pulumi.StringPtrInput
	// A named token to use for ingest
	NamedToken pulumi.StringPtrInput
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate pulumi.IntPtrInput
	// Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	SecretKey pulumi.StringInput
	// List of Microsoft Azure service names for the Azure services you want SignalFx to monitor. See the documentation for [Creating Integrations](https://developers.signalfx.com/integrations_reference.html#operation/Create%20Integration) for valida values.
	Services pulumi.StringArrayInput
	// List of Azure subscriptions that SignalFx should monitor.
	Subscriptions pulumi.StringArrayInput
	// If enabled, SignalFx will try to sync additional namespaces for VMs (including VMs in scale sets): telegraf/mem, telegraf/cpu, azure.vm.windows.guest (these are namespaces recommended by Azure when enabling their Diagnostic Extension). If there are no metrics there, no new datapoints will be ingested. Defaults to false.
	SyncGuestOsNamespaces pulumi.BoolPtrInput
	// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	TenantId pulumi.StringInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}

type IntegrationInput interface {
	pulumi.Input

	ToIntegrationOutput() IntegrationOutput
	ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput
}

func (*Integration) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil)).Elem()
}

func (i *Integration) ToIntegrationOutput() IntegrationOutput {
	return i.ToIntegrationOutputWithContext(context.Background())
}

func (i *Integration) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationOutput)
}

// IntegrationArrayInput is an input type that accepts IntegrationArray and IntegrationArrayOutput values.
// You can construct a concrete instance of `IntegrationArrayInput` via:
//
//          IntegrationArray{ IntegrationArgs{...} }
type IntegrationArrayInput interface {
	pulumi.Input

	ToIntegrationArrayOutput() IntegrationArrayOutput
	ToIntegrationArrayOutputWithContext(context.Context) IntegrationArrayOutput
}

type IntegrationArray []IntegrationInput

func (IntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Integration)(nil)).Elem()
}

func (i IntegrationArray) ToIntegrationArrayOutput() IntegrationArrayOutput {
	return i.ToIntegrationArrayOutputWithContext(context.Background())
}

func (i IntegrationArray) ToIntegrationArrayOutputWithContext(ctx context.Context) IntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationArrayOutput)
}

// IntegrationMapInput is an input type that accepts IntegrationMap and IntegrationMapOutput values.
// You can construct a concrete instance of `IntegrationMapInput` via:
//
//          IntegrationMap{ "key": IntegrationArgs{...} }
type IntegrationMapInput interface {
	pulumi.Input

	ToIntegrationMapOutput() IntegrationMapOutput
	ToIntegrationMapOutputWithContext(context.Context) IntegrationMapOutput
}

type IntegrationMap map[string]IntegrationInput

func (IntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Integration)(nil)).Elem()
}

func (i IntegrationMap) ToIntegrationMapOutput() IntegrationMapOutput {
	return i.ToIntegrationMapOutputWithContext(context.Background())
}

func (i IntegrationMap) ToIntegrationMapOutputWithContext(ctx context.Context) IntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMapOutput)
}

type IntegrationOutput struct{ *pulumi.OutputState }

func (IntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil)).Elem()
}

func (o IntegrationOutput) ToIntegrationOutput() IntegrationOutput {
	return o
}

func (o IntegrationOutput) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return o
}

type IntegrationArrayOutput struct{ *pulumi.OutputState }

func (IntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Integration)(nil)).Elem()
}

func (o IntegrationArrayOutput) ToIntegrationArrayOutput() IntegrationArrayOutput {
	return o
}

func (o IntegrationArrayOutput) ToIntegrationArrayOutputWithContext(ctx context.Context) IntegrationArrayOutput {
	return o
}

func (o IntegrationArrayOutput) Index(i pulumi.IntInput) IntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Integration {
		return vs[0].([]*Integration)[vs[1].(int)]
	}).(IntegrationOutput)
}

type IntegrationMapOutput struct{ *pulumi.OutputState }

func (IntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Integration)(nil)).Elem()
}

func (o IntegrationMapOutput) ToIntegrationMapOutput() IntegrationMapOutput {
	return o
}

func (o IntegrationMapOutput) ToIntegrationMapOutputWithContext(ctx context.Context) IntegrationMapOutput {
	return o
}

func (o IntegrationMapOutput) MapIndex(k pulumi.StringInput) IntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Integration {
		return vs[0].(map[string]*Integration)[vs[1].(string)]
	}).(IntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationInput)(nil)).Elem(), &Integration{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationArrayInput)(nil)).Elem(), IntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMapInput)(nil)).Elem(), IntegrationMap{})
	pulumi.RegisterOutputType(IntegrationOutput{})
	pulumi.RegisterOutputType(IntegrationArrayOutput{})
	pulumi.RegisterOutputType(IntegrationMapOutput{})
}
