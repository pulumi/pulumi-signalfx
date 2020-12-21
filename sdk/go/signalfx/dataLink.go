// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manage SignalFx [Data Links](https://docs.signalfx.com/en/latest/managing/data-links.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-signalfx/sdk/v4/go/signalfx"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := signalfx.NewDataLink(ctx, "myDataLink", &signalfx.DataLinkArgs{
// 			PropertyName:  pulumi.String("pname"),
// 			PropertyValue: pulumi.String("pvalue"),
// 			TargetSignalfxDashboards: signalfx.DataLinkTargetSignalfxDashboardArray{
// 				&signalfx.DataLinkTargetSignalfxDashboardArgs{
// 					IsDefault:        pulumi.Bool(true),
// 					Name:             pulumi.String("sfx_dash"),
// 					DashboardGroupId: pulumi.Any(signalfx_dashboard_group.Mydashboardgroup0.Id),
// 					DashboardId:      pulumi.Any(signalfx_dashboard.Mydashboard0.Id),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = signalfx.NewDataLink(ctx, "myDataLinkDash", &signalfx.DataLinkArgs{
// 			ContextDashboardId: pulumi.Any(signalfx_dashboard.Mydashboard0.Id),
// 			PropertyName:       pulumi.String("pname2"),
// 			PropertyValue:      pulumi.String("pvalue"),
// 			TargetExternalUrls: signalfx.DataLinkTargetExternalUrlArray{
// 				&signalfx.DataLinkTargetExternalUrlArgs{
// 					IsDefault:  pulumi.Bool(false),
// 					Name:       pulumi.String("ex_url"),
// 					TimeFormat: pulumi.String("ISO8601"),
// 					Url:        pulumi.String("https://www.example.com"),
// 					PropertyKeyMapping: pulumi.StringMap{
// 						"foo": pulumi.String("bar"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
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
	// Link to a SignalFx dashboard
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
	// Link to a SignalFx dashboard
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
	// Link to a SignalFx dashboard
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
	// Link to a SignalFx dashboard
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
	// Link to a SignalFx dashboard
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

func (DataLink) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLink)(nil)).Elem()
}

func (i DataLink) ToDataLinkOutput() DataLinkOutput {
	return i.ToDataLinkOutputWithContext(context.Background())
}

func (i DataLink) ToDataLinkOutputWithContext(ctx context.Context) DataLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLinkOutput)
}

type DataLinkOutput struct {
	*pulumi.OutputState
}

func (DataLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataLinkOutput)(nil)).Elem()
}

func (o DataLinkOutput) ToDataLinkOutput() DataLinkOutput {
	return o
}

func (o DataLinkOutput) ToDataLinkOutputWithContext(ctx context.Context) DataLinkOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataLinkOutput{})
}
