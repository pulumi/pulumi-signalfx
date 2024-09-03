// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage Splunk Observability Cloud [Data Links](https://docs.signalfx.com/en/latest/managing/data-links.html).
//
// ## Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// A global link to Splunk Observability Cloud dashboard.
//			_, err := signalfx.NewDataLink(ctx, "my_data_link", &signalfx.DataLinkArgs{
//				PropertyName:  pulumi.String("pname"),
//				PropertyValue: pulumi.String("pvalue"),
//				TargetSignalfxDashboards: signalfx.DataLinkTargetSignalfxDashboardArray{
//					&signalfx.DataLinkTargetSignalfxDashboardArgs{
//						IsDefault:        pulumi.Bool(true),
//						Name:             pulumi.String("sfx_dash"),
//						DashboardGroupId: pulumi.Any(mydashboardgroup0.Id),
//						DashboardId:      pulumi.Any(mydashboard0.Id),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// A dashboard-specific link to an external URL
//			_, err = signalfx.NewDataLink(ctx, "my_data_link_dash", &signalfx.DataLinkArgs{
//				ContextDashboardId: pulumi.Any(mydashboard0.Id),
//				PropertyName:       pulumi.String("pname2"),
//				PropertyValue:      pulumi.String("pvalue"),
//				TargetExternalUrls: signalfx.DataLinkTargetExternalUrlArray{
//					&signalfx.DataLinkTargetExternalUrlArgs{
//						Name:       pulumi.String("ex_url"),
//						TimeFormat: pulumi.String("ISO8601"),
//						Url:        pulumi.String("https://www.example.com"),
//						PropertyKeyMapping: pulumi.StringMap{
//							"foo": pulumi.String("bar"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DataLink struct {
	pulumi.CustomResourceState

	// If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
	ContextDashboardId pulumi.StringPtrOutput `pulumi:"contextDashboardId"`
	// Name (key) of the metadata that's the trigger of a data link. If you specify `propertyValue`, you must specify `propertyName`.
	PropertyName pulumi.StringPtrOutput `pulumi:"propertyName"`
	// Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `propertyName`.
	PropertyValue pulumi.StringPtrOutput `pulumi:"propertyValue"`
	// Link to an external URL
	TargetExternalUrls DataLinkTargetExternalUrlArrayOutput `pulumi:"targetExternalUrls"`
	// Link to a Splunk Observability Cloud dashboard
	TargetSignalfxDashboards DataLinkTargetSignalfxDashboardArrayOutput `pulumi:"targetSignalfxDashboards"`
	// Link to an external URL
	TargetSplunks DataLinkTargetSplunkArrayOutput `pulumi:"targetSplunks"`
}

// NewDataLink registers a new resource with the given unique name, arguments, and options.
func NewDataLink(ctx *pulumi.Context,
	name string, args *DataLinkArgs, opts ...pulumi.ResourceOption) (*DataLink, error) {
	if args == nil {
		args = &DataLinkArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataLink
	err := ctx.RegisterResource("signalfx:index/dataLink:DataLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataLink gets an existing DataLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataLinkState, opts ...pulumi.ResourceOption) (*DataLink, error) {
	var resource DataLink
	err := ctx.ReadResource("signalfx:index/dataLink:DataLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataLink resources.
type dataLinkState struct {
	// If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
	ContextDashboardId *string `pulumi:"contextDashboardId"`
	// Name (key) of the metadata that's the trigger of a data link. If you specify `propertyValue`, you must specify `propertyName`.
	PropertyName *string `pulumi:"propertyName"`
	// Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `propertyName`.
	PropertyValue *string `pulumi:"propertyValue"`
	// Link to an external URL
	TargetExternalUrls []DataLinkTargetExternalUrl `pulumi:"targetExternalUrls"`
	// Link to a Splunk Observability Cloud dashboard
	TargetSignalfxDashboards []DataLinkTargetSignalfxDashboard `pulumi:"targetSignalfxDashboards"`
	// Link to an external URL
	TargetSplunks []DataLinkTargetSplunk `pulumi:"targetSplunks"`
}

type DataLinkState struct {
	// If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
	ContextDashboardId pulumi.StringPtrInput
	// Name (key) of the metadata that's the trigger of a data link. If you specify `propertyValue`, you must specify `propertyName`.
	PropertyName pulumi.StringPtrInput
	// Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `propertyName`.
	PropertyValue pulumi.StringPtrInput
	// Link to an external URL
	TargetExternalUrls DataLinkTargetExternalUrlArrayInput
	// Link to a Splunk Observability Cloud dashboard
	TargetSignalfxDashboards DataLinkTargetSignalfxDashboardArrayInput
	// Link to an external URL
	TargetSplunks DataLinkTargetSplunkArrayInput
}

func (DataLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLinkState)(nil)).Elem()
}

type dataLinkArgs struct {
	// If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
	ContextDashboardId *string `pulumi:"contextDashboardId"`
	// Name (key) of the metadata that's the trigger of a data link. If you specify `propertyValue`, you must specify `propertyName`.
	PropertyName *string `pulumi:"propertyName"`
	// Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `propertyName`.
	PropertyValue *string `pulumi:"propertyValue"`
	// Link to an external URL
	TargetExternalUrls []DataLinkTargetExternalUrl `pulumi:"targetExternalUrls"`
	// Link to a Splunk Observability Cloud dashboard
	TargetSignalfxDashboards []DataLinkTargetSignalfxDashboard `pulumi:"targetSignalfxDashboards"`
	// Link to an external URL
	TargetSplunks []DataLinkTargetSplunk `pulumi:"targetSplunks"`
}

// The set of arguments for constructing a DataLink resource.
type DataLinkArgs struct {
	// If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
	ContextDashboardId pulumi.StringPtrInput
	// Name (key) of the metadata that's the trigger of a data link. If you specify `propertyValue`, you must specify `propertyName`.
	PropertyName pulumi.StringPtrInput
	// Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `propertyName`.
	PropertyValue pulumi.StringPtrInput
	// Link to an external URL
	TargetExternalUrls DataLinkTargetExternalUrlArrayInput
	// Link to a Splunk Observability Cloud dashboard
	TargetSignalfxDashboards DataLinkTargetSignalfxDashboardArrayInput
	// Link to an external URL
	TargetSplunks DataLinkTargetSplunkArrayInput
}

func (DataLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLinkArgs)(nil)).Elem()
}

type DataLinkInput interface {
	pulumi.Input

	ToDataLinkOutput() DataLinkOutput
	ToDataLinkOutputWithContext(ctx context.Context) DataLinkOutput
}

func (*DataLink) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLink)(nil)).Elem()
}

func (i *DataLink) ToDataLinkOutput() DataLinkOutput {
	return i.ToDataLinkOutputWithContext(context.Background())
}

func (i *DataLink) ToDataLinkOutputWithContext(ctx context.Context) DataLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLinkOutput)
}

// DataLinkArrayInput is an input type that accepts DataLinkArray and DataLinkArrayOutput values.
// You can construct a concrete instance of `DataLinkArrayInput` via:
//
//	DataLinkArray{ DataLinkArgs{...} }
type DataLinkArrayInput interface {
	pulumi.Input

	ToDataLinkArrayOutput() DataLinkArrayOutput
	ToDataLinkArrayOutputWithContext(context.Context) DataLinkArrayOutput
}

type DataLinkArray []DataLinkInput

func (DataLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataLink)(nil)).Elem()
}

func (i DataLinkArray) ToDataLinkArrayOutput() DataLinkArrayOutput {
	return i.ToDataLinkArrayOutputWithContext(context.Background())
}

func (i DataLinkArray) ToDataLinkArrayOutputWithContext(ctx context.Context) DataLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLinkArrayOutput)
}

// DataLinkMapInput is an input type that accepts DataLinkMap and DataLinkMapOutput values.
// You can construct a concrete instance of `DataLinkMapInput` via:
//
//	DataLinkMap{ "key": DataLinkArgs{...} }
type DataLinkMapInput interface {
	pulumi.Input

	ToDataLinkMapOutput() DataLinkMapOutput
	ToDataLinkMapOutputWithContext(context.Context) DataLinkMapOutput
}

type DataLinkMap map[string]DataLinkInput

func (DataLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataLink)(nil)).Elem()
}

func (i DataLinkMap) ToDataLinkMapOutput() DataLinkMapOutput {
	return i.ToDataLinkMapOutputWithContext(context.Background())
}

func (i DataLinkMap) ToDataLinkMapOutputWithContext(ctx context.Context) DataLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLinkMapOutput)
}

type DataLinkOutput struct{ *pulumi.OutputState }

func (DataLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLink)(nil)).Elem()
}

func (o DataLinkOutput) ToDataLinkOutput() DataLinkOutput {
	return o
}

func (o DataLinkOutput) ToDataLinkOutputWithContext(ctx context.Context) DataLinkOutput {
	return o
}

// If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
func (o DataLinkOutput) ContextDashboardId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLink) pulumi.StringPtrOutput { return v.ContextDashboardId }).(pulumi.StringPtrOutput)
}

// Name (key) of the metadata that's the trigger of a data link. If you specify `propertyValue`, you must specify `propertyName`.
func (o DataLinkOutput) PropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLink) pulumi.StringPtrOutput { return v.PropertyName }).(pulumi.StringPtrOutput)
}

// Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `propertyName`.
func (o DataLinkOutput) PropertyValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLink) pulumi.StringPtrOutput { return v.PropertyValue }).(pulumi.StringPtrOutput)
}

// Link to an external URL
func (o DataLinkOutput) TargetExternalUrls() DataLinkTargetExternalUrlArrayOutput {
	return o.ApplyT(func(v *DataLink) DataLinkTargetExternalUrlArrayOutput { return v.TargetExternalUrls }).(DataLinkTargetExternalUrlArrayOutput)
}

// Link to a Splunk Observability Cloud dashboard
func (o DataLinkOutput) TargetSignalfxDashboards() DataLinkTargetSignalfxDashboardArrayOutput {
	return o.ApplyT(func(v *DataLink) DataLinkTargetSignalfxDashboardArrayOutput { return v.TargetSignalfxDashboards }).(DataLinkTargetSignalfxDashboardArrayOutput)
}

// Link to an external URL
func (o DataLinkOutput) TargetSplunks() DataLinkTargetSplunkArrayOutput {
	return o.ApplyT(func(v *DataLink) DataLinkTargetSplunkArrayOutput { return v.TargetSplunks }).(DataLinkTargetSplunkArrayOutput)
}

type DataLinkArrayOutput struct{ *pulumi.OutputState }

func (DataLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataLink)(nil)).Elem()
}

func (o DataLinkArrayOutput) ToDataLinkArrayOutput() DataLinkArrayOutput {
	return o
}

func (o DataLinkArrayOutput) ToDataLinkArrayOutputWithContext(ctx context.Context) DataLinkArrayOutput {
	return o
}

func (o DataLinkArrayOutput) Index(i pulumi.IntInput) DataLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataLink {
		return vs[0].([]*DataLink)[vs[1].(int)]
	}).(DataLinkOutput)
}

type DataLinkMapOutput struct{ *pulumi.OutputState }

func (DataLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataLink)(nil)).Elem()
}

func (o DataLinkMapOutput) ToDataLinkMapOutput() DataLinkMapOutput {
	return o
}

func (o DataLinkMapOutput) ToDataLinkMapOutputWithContext(ctx context.Context) DataLinkMapOutput {
	return o
}

func (o DataLinkMapOutput) MapIndex(k pulumi.StringInput) DataLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataLink {
		return vs[0].(map[string]*DataLink)[vs[1].(string)]
	}).(DataLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataLinkInput)(nil)).Elem(), &DataLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLinkArrayInput)(nil)).Elem(), DataLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLinkMapInput)(nil)).Elem(), DataLinkMap{})
	pulumi.RegisterOutputType(DataLinkOutput{})
	pulumi.RegisterOutputType(DataLinkArrayOutput{})
	pulumi.RegisterOutputType(DataLinkMapOutput{})
}
