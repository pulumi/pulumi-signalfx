// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Splunk Observability Cloud GCP Integration.
//
// > **NOTE** When managing integrations, use a session token of an administrator to authenticate the Splunk  Observability Cloud provider. See [Operations that require a session token for an administrator](https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator). Otherwise you'll receive a 4xx error.
//
// ## Example
//
// ## Arguments
//
// * `customMetricTypeDomains` - (Optional) List of additional GCP service domain names that Splunk Observability Cloud will monitor. See [Custom Metric Type Domains documentation](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/#Custom-metric-type-domains)
// * `enabled` - (Required) Whether the integration is enabled.
// * `importGcpMetrics` - (Optional) If enabled, Splunk Observability Cloud will sync also Google Cloud Monitoring data. If disabled, Splunk Observability Cloud will import only metadata. Defaults to true.
// * `includeList` - (Optional) [Compute Metadata Include List](https://dev.splunk.com/observability/docs/integrations/gcp_integration_overview/).
// * `name` - (Required) Name of the integration.
// * `namedToken` - (Optional) Name of the org token to be used for data ingestion. If not specified then default access token is used.
// * `pollRate` - (Optional) GCP integration poll rate (in seconds). Value between `60` and `600`. Default: `300`.
// * `projectServiceKeys` - (Required) GCP projects to add.
// * `services` - (Optional) GCP service metrics to import. Can be an empty list, or not included, to import 'All services'. See [Google Cloud Platform services](https://docs.splunk.com/Observability/gdi/get-data-in/integrations.html#google-cloud-platform-services) for a list of valid values.
// * `useMetricSourceProjectForQuota` - (Optional) When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are stored. For this to work the service account provided for the project needs to be provided with serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota settings are used.
//
// ## Attributes
//
// In addition to all arguments above, the following attributes are exported:
//
// * `id` - The ID of the integration.
type Integration struct {
	pulumi.CustomResourceState

	// List of additional GCP service domain names that you want to monitor
	CustomMetricTypeDomains pulumi.StringArrayOutput `pulumi:"customMetricTypeDomains"`
	// Whether the integration is enabled or not
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// If enabled, Splunk Observability Cloud will sync also Google Cloud Metrics data. If disabled, Splunk Observability Cloud
	// will import only metadata. Defaults to true.
	ImportGcpMetrics pulumi.BoolPtrOutput `pulumi:"importGcpMetrics"`
	// List of custom metadata keys that you want Observability Cloud to collect for Compute Engine instances.
	IncludeLists pulumi.StringArrayOutput `pulumi:"includeLists"`
	// Name of the integration
	Name pulumi.StringOutput `pulumi:"name"`
	// A named token to use for ingest
	NamedToken pulumi.StringPtrOutput `pulumi:"namedToken"`
	// GCP poll rate (in seconds). Between `60` and `600`.
	PollRate pulumi.IntPtrOutput `pulumi:"pollRate"`
	// GCP project service keys
	ProjectServiceKeys IntegrationProjectServiceKeyArrayOutput `pulumi:"projectServiceKeys"`
	// GCP enabled services
	Services pulumi.StringArrayOutput `pulumi:"services"`
	// When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are
	// stored. For this to work the service account provided for the project needs to be provided with
	// serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota
	// settings are used.
	UseMetricSourceProjectForQuota pulumi.BoolPtrOutput `pulumi:"useMetricSourceProjectForQuota"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.ProjectServiceKeys != nil {
		args.ProjectServiceKeys = pulumi.ToSecret(args.ProjectServiceKeys).(IntegrationProjectServiceKeyArrayInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"projectServiceKeys",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Integration
	err := ctx.RegisterResource("signalfx:gcp/integration:Integration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("signalfx:gcp/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// List of additional GCP service domain names that you want to monitor
	CustomMetricTypeDomains []string `pulumi:"customMetricTypeDomains"`
	// Whether the integration is enabled or not
	Enabled *bool `pulumi:"enabled"`
	// If enabled, Splunk Observability Cloud will sync also Google Cloud Metrics data. If disabled, Splunk Observability Cloud
	// will import only metadata. Defaults to true.
	ImportGcpMetrics *bool `pulumi:"importGcpMetrics"`
	// List of custom metadata keys that you want Observability Cloud to collect for Compute Engine instances.
	IncludeLists []string `pulumi:"includeLists"`
	// Name of the integration
	Name *string `pulumi:"name"`
	// A named token to use for ingest
	NamedToken *string `pulumi:"namedToken"`
	// GCP poll rate (in seconds). Between `60` and `600`.
	PollRate *int `pulumi:"pollRate"`
	// GCP project service keys
	ProjectServiceKeys []IntegrationProjectServiceKey `pulumi:"projectServiceKeys"`
	// GCP enabled services
	Services []string `pulumi:"services"`
	// When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are
	// stored. For this to work the service account provided for the project needs to be provided with
	// serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota
	// settings are used.
	UseMetricSourceProjectForQuota *bool `pulumi:"useMetricSourceProjectForQuota"`
}

type IntegrationState struct {
	// List of additional GCP service domain names that you want to monitor
	CustomMetricTypeDomains pulumi.StringArrayInput
	// Whether the integration is enabled or not
	Enabled pulumi.BoolPtrInput
	// If enabled, Splunk Observability Cloud will sync also Google Cloud Metrics data. If disabled, Splunk Observability Cloud
	// will import only metadata. Defaults to true.
	ImportGcpMetrics pulumi.BoolPtrInput
	// List of custom metadata keys that you want Observability Cloud to collect for Compute Engine instances.
	IncludeLists pulumi.StringArrayInput
	// Name of the integration
	Name pulumi.StringPtrInput
	// A named token to use for ingest
	NamedToken pulumi.StringPtrInput
	// GCP poll rate (in seconds). Between `60` and `600`.
	PollRate pulumi.IntPtrInput
	// GCP project service keys
	ProjectServiceKeys IntegrationProjectServiceKeyArrayInput
	// GCP enabled services
	Services pulumi.StringArrayInput
	// When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are
	// stored. For this to work the service account provided for the project needs to be provided with
	// serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota
	// settings are used.
	UseMetricSourceProjectForQuota pulumi.BoolPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// List of additional GCP service domain names that you want to monitor
	CustomMetricTypeDomains []string `pulumi:"customMetricTypeDomains"`
	// Whether the integration is enabled or not
	Enabled bool `pulumi:"enabled"`
	// If enabled, Splunk Observability Cloud will sync also Google Cloud Metrics data. If disabled, Splunk Observability Cloud
	// will import only metadata. Defaults to true.
	ImportGcpMetrics *bool `pulumi:"importGcpMetrics"`
	// List of custom metadata keys that you want Observability Cloud to collect for Compute Engine instances.
	IncludeLists []string `pulumi:"includeLists"`
	// Name of the integration
	Name *string `pulumi:"name"`
	// A named token to use for ingest
	NamedToken *string `pulumi:"namedToken"`
	// GCP poll rate (in seconds). Between `60` and `600`.
	PollRate *int `pulumi:"pollRate"`
	// GCP project service keys
	ProjectServiceKeys []IntegrationProjectServiceKey `pulumi:"projectServiceKeys"`
	// GCP enabled services
	Services []string `pulumi:"services"`
	// When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are
	// stored. For this to work the service account provided for the project needs to be provided with
	// serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota
	// settings are used.
	UseMetricSourceProjectForQuota *bool `pulumi:"useMetricSourceProjectForQuota"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// List of additional GCP service domain names that you want to monitor
	CustomMetricTypeDomains pulumi.StringArrayInput
	// Whether the integration is enabled or not
	Enabled pulumi.BoolInput
	// If enabled, Splunk Observability Cloud will sync also Google Cloud Metrics data. If disabled, Splunk Observability Cloud
	// will import only metadata. Defaults to true.
	ImportGcpMetrics pulumi.BoolPtrInput
	// List of custom metadata keys that you want Observability Cloud to collect for Compute Engine instances.
	IncludeLists pulumi.StringArrayInput
	// Name of the integration
	Name pulumi.StringPtrInput
	// A named token to use for ingest
	NamedToken pulumi.StringPtrInput
	// GCP poll rate (in seconds). Between `60` and `600`.
	PollRate pulumi.IntPtrInput
	// GCP project service keys
	ProjectServiceKeys IntegrationProjectServiceKeyArrayInput
	// GCP enabled services
	Services pulumi.StringArrayInput
	// When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are
	// stored. For this to work the service account provided for the project needs to be provided with
	// serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota
	// settings are used.
	UseMetricSourceProjectForQuota pulumi.BoolPtrInput
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
//	IntegrationArray{ IntegrationArgs{...} }
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
//	IntegrationMap{ "key": IntegrationArgs{...} }
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

// List of additional GCP service domain names that you want to monitor
func (o IntegrationOutput) CustomMetricTypeDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringArrayOutput { return v.CustomMetricTypeDomains }).(pulumi.StringArrayOutput)
}

// Whether the integration is enabled or not
func (o IntegrationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Integration) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// If enabled, Splunk Observability Cloud will sync also Google Cloud Metrics data. If disabled, Splunk Observability Cloud
// will import only metadata. Defaults to true.
func (o IntegrationOutput) ImportGcpMetrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Integration) pulumi.BoolPtrOutput { return v.ImportGcpMetrics }).(pulumi.BoolPtrOutput)
}

// List of custom metadata keys that you want Observability Cloud to collect for Compute Engine instances.
func (o IntegrationOutput) IncludeLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringArrayOutput { return v.IncludeLists }).(pulumi.StringArrayOutput)
}

// Name of the integration
func (o IntegrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A named token to use for ingest
func (o IntegrationOutput) NamedToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringPtrOutput { return v.NamedToken }).(pulumi.StringPtrOutput)
}

// GCP poll rate (in seconds). Between `60` and `600`.
func (o IntegrationOutput) PollRate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Integration) pulumi.IntPtrOutput { return v.PollRate }).(pulumi.IntPtrOutput)
}

// GCP project service keys
func (o IntegrationOutput) ProjectServiceKeys() IntegrationProjectServiceKeyArrayOutput {
	return o.ApplyT(func(v *Integration) IntegrationProjectServiceKeyArrayOutput { return v.ProjectServiceKeys }).(IntegrationProjectServiceKeyArrayOutput)
}

// GCP enabled services
func (o IntegrationOutput) Services() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringArrayOutput { return v.Services }).(pulumi.StringArrayOutput)
}

// When this value is set to true Observability Cloud will force usage of a quota from the project where metrics are
// stored. For this to work the service account provided for the project needs to be provided with
// serviceusage.services.use permission or Service Usage Consumer role in this project. When set to false default quota
// settings are used.
func (o IntegrationOutput) UseMetricSourceProjectForQuota() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Integration) pulumi.BoolPtrOutput { return v.UseMetricSourceProjectForQuota }).(pulumi.BoolPtrOutput)
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
