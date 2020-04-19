// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// SignalFx Azure integrations. For help with this integration see [Monitoring Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure).
//
// > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.
//
//
// ## Service Names
//
// Fields that expect an Azure service will work with one of: "microsoft.sql/servers/elasticpools" "microsoft.storage/storageaccounts" "microsoft.storage/storageaccountsservices/tableservices" "microsoft.storage/storageaccountsservices/blobservices" "microsoft.storage/storageaccounts/queueservices" "microsoft.storage/storageaccounts/fileservices" "microsoft.compute/virtualmachinescalesets" "microsoft.compute/virtualmachinescalesets/virtualmachines" "microsoft.compute/virtualmachines" "microsoft.devices/iothubs" "microsoft.eventHub/namespaces" "microsoft.batch/batchaccounts" "microsoft.sql/servers/databases" "microsoft.cache/redis" "microsoft.logic/workflows".
type Integration struct {
	pulumi.CustomResourceState

	// Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// Whether the integration is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// Name of the integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate pulumi.IntPtrOutput `pulumi:"pollRate"`
	// Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// List of Microsoft Azure service names for the Azure services you want SignalFx to monitor.
	Services pulumi.StringArrayOutput `pulumi:"services"`
	// List of Azure subscriptions that SignalFx should monitor.
	Subscriptions pulumi.StringArrayOutput `pulumi:"subscriptions"`
	// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil || args.AppId == nil {
		return nil, errors.New("missing required argument 'AppId'")
	}
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.SecretKey == nil {
		return nil, errors.New("missing required argument 'SecretKey'")
	}
	if args == nil || args.Services == nil {
		return nil, errors.New("missing required argument 'Services'")
	}
	if args == nil || args.Subscriptions == nil {
		return nil, errors.New("missing required argument 'Subscriptions'")
	}
	if args == nil || args.TenantId == nil {
		return nil, errors.New("missing required argument 'TenantId'")
	}
	if args == nil {
		args = &IntegrationArgs{}
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
	// Whether the integration is enabled.
	Enabled *bool `pulumi:"enabled"`
	// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
	Environment *string `pulumi:"environment"`
	// Name of the integration.
	Name *string `pulumi:"name"`
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate *int `pulumi:"pollRate"`
	// Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	SecretKey *string `pulumi:"secretKey"`
	// List of Microsoft Azure service names for the Azure services you want SignalFx to monitor.
	Services []string `pulumi:"services"`
	// List of Azure subscriptions that SignalFx should monitor.
	Subscriptions []string `pulumi:"subscriptions"`
	// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	TenantId *string `pulumi:"tenantId"`
}

type IntegrationState struct {
	// Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
	AppId pulumi.StringPtrInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolPtrInput
	// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
	Environment pulumi.StringPtrInput
	// Name of the integration.
	Name pulumi.StringPtrInput
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate pulumi.IntPtrInput
	// Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	SecretKey pulumi.StringPtrInput
	// List of Microsoft Azure service names for the Azure services you want SignalFx to monitor.
	Services pulumi.StringArrayInput
	// List of Azure subscriptions that SignalFx should monitor.
	Subscriptions pulumi.StringArrayInput
	// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	TenantId pulumi.StringPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
	AppId string `pulumi:"appId"`
	// Whether the integration is enabled.
	Enabled bool `pulumi:"enabled"`
	// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
	Environment *string `pulumi:"environment"`
	// Name of the integration.
	Name *string `pulumi:"name"`
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate *int `pulumi:"pollRate"`
	// Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	SecretKey string `pulumi:"secretKey"`
	// List of Microsoft Azure service names for the Azure services you want SignalFx to monitor.
	Services []string `pulumi:"services"`
	// List of Azure subscriptions that SignalFx should monitor.
	Subscriptions []string `pulumi:"subscriptions"`
	// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// Azure application ID for the SignalFx app. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/getting-started/send-data.html#connect-to-microsoft-azure) in the product documentation.
	AppId pulumi.StringInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolInput
	// What type of Azure integration this is. The allowed values are `\"azure_us_government\"` and `\"azure\"`. Defaults to `\"azure\"`.
	Environment pulumi.StringPtrInput
	// Name of the integration.
	Name pulumi.StringPtrInput
	// AWS poll rate (in seconds). One of `60` or `300`.
	PollRate pulumi.IntPtrInput
	// Azure secret key that associates the SignalFx app in Azure with the Azure tenant ID. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	SecretKey pulumi.StringInput
	// List of Microsoft Azure service names for the Azure services you want SignalFx to monitor.
	Services pulumi.StringArrayInput
	// List of Azure subscriptions that SignalFx should monitor.
	Subscriptions pulumi.StringArrayInput
	// Azure ID of the Azure tenant. To learn how to get this ID, see the topic [Connect to Microsoft Azure](https://docs.signalfx.com/en/latest/integrations/azure-info.html#connect-to-azure) in the product documentation.
	TenantId pulumi.StringInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}
