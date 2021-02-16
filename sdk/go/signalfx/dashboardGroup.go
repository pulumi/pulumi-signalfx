// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// In the SignalFx web UI, a [dashboard group](https://developers.signalfx.com/dashboard_groups_reference.html) is a collection of dashboards.
//
// > **NOTE** Dashboard groups cannot be accessed directly, but just via a dashboard contained in them. This is the reason why make show won't show any of yours dashboard groups.
//
// ## Example Usage
// ### With Mirrored Dashboards
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
// 		_, err := signalfx.NewDashboardGroup(ctx, "mydashboardgroupWithmirrors", &signalfx.DashboardGroupArgs{
// 			Description: pulumi.String("Cool dashboard group"),
// 			Dashboards: signalfx.DashboardGroupDashboardArray{
// 				&signalfx.DashboardGroupDashboardArgs{
// 					DashboardId:         pulumi.Any(signalfx_dashboard.Gc_dashboard.Id),
// 					NameOverride:        pulumi.String("GC For My Service"),
// 					DescriptionOverride: pulumi.String("Garbage Collection dashboard maintained by JVM team"),
// 					FilterOverrides: signalfx.DashboardGroupDashboardFilterOverrideArray{
// 						&signalfx.DashboardGroupDashboardFilterOverrideArgs{
// 							Property: pulumi.String("service"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("myservice"),
// 							},
// 							Negated: pulumi.Bool(false),
// 						},
// 					},
// 					VariableOverrides: signalfx.DashboardGroupDashboardVariableOverrideArray{
// 						&signalfx.DashboardGroupDashboardVariableOverrideArgs{
// 							Property: pulumi.String("region"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("us-west1"),
// 							},
// 							ValuesSuggesteds: pulumi.StringArray{
// 								pulumi.String("us-west-1"),
// 								pulumi.String("us-east-1"),
// 							},
// 						},
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
type DashboardGroup struct {
	pulumi.CustomResourceState

	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
	AuthorizedWriterTeams pulumi.StringArrayOutput `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers pulumi.StringArrayOutput `pulumi:"authorizedWriterUsers"`
	// [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
	Dashboards DashboardGroupDashboardArrayOutput `pulumi:"dashboards"`
	// Description of the dashboard group.
	Description      pulumi.StringPtrOutput                   `pulumi:"description"`
	ImportQualifiers DashboardGroupImportQualifierArrayOutput `pulumi:"importQualifiers"`
	// Name of the dashboard group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Team IDs to associate the dashboard group to.
	Teams pulumi.StringArrayOutput `pulumi:"teams"`
}

// NewDashboardGroup registers a new resource with the given unique name, arguments, and options.
func NewDashboardGroup(ctx *pulumi.Context,
	name string, args *DashboardGroupArgs, opts ...pulumi.ResourceOption) (*DashboardGroup, error) {
	if args == nil {
		args = &DashboardGroupArgs{}
	}

	var resource DashboardGroup
	err := ctx.RegisterResource("signalfx:index/dashboardGroup:DashboardGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardGroup gets an existing DashboardGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardGroupState, opts ...pulumi.ResourceOption) (*DashboardGroup, error) {
	var resource DashboardGroup
	err := ctx.ReadResource("signalfx:index/dashboardGroup:DashboardGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardGroup resources.
type dashboardGroupState struct {
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
	AuthorizedWriterTeams []string `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers []string `pulumi:"authorizedWriterUsers"`
	// [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
	Dashboards []DashboardGroupDashboard `pulumi:"dashboards"`
	// Description of the dashboard group.
	Description      *string                         `pulumi:"description"`
	ImportQualifiers []DashboardGroupImportQualifier `pulumi:"importQualifiers"`
	// Name of the dashboard group.
	Name *string `pulumi:"name"`
	// Team IDs to associate the dashboard group to.
	Teams []string `pulumi:"teams"`
}

type DashboardGroupState struct {
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
	AuthorizedWriterTeams pulumi.StringArrayInput
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers pulumi.StringArrayInput
	// [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
	Dashboards DashboardGroupDashboardArrayInput
	// Description of the dashboard group.
	Description      pulumi.StringPtrInput
	ImportQualifiers DashboardGroupImportQualifierArrayInput
	// Name of the dashboard group.
	Name pulumi.StringPtrInput
	// Team IDs to associate the dashboard group to.
	Teams pulumi.StringArrayInput
}

func (DashboardGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardGroupState)(nil)).Elem()
}

type dashboardGroupArgs struct {
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
	AuthorizedWriterTeams []string `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers []string `pulumi:"authorizedWriterUsers"`
	// [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
	Dashboards []DashboardGroupDashboard `pulumi:"dashboards"`
	// Description of the dashboard group.
	Description      *string                         `pulumi:"description"`
	ImportQualifiers []DashboardGroupImportQualifier `pulumi:"importQualifiers"`
	// Name of the dashboard group.
	Name *string `pulumi:"name"`
	// Team IDs to associate the dashboard group to.
	Teams []string `pulumi:"teams"`
}

// The set of arguments for constructing a DashboardGroup resource.
type DashboardGroupArgs struct {
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
	AuthorizedWriterTeams pulumi.StringArrayInput
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers pulumi.StringArrayInput
	// [Mirrored dashboards](https://docs.signalfx.com/en/latest/dashboards/dashboard-mirrors.html) in this dashboard group. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
	Dashboards DashboardGroupDashboardArrayInput
	// Description of the dashboard group.
	Description      pulumi.StringPtrInput
	ImportQualifiers DashboardGroupImportQualifierArrayInput
	// Name of the dashboard group.
	Name pulumi.StringPtrInput
	// Team IDs to associate the dashboard group to.
	Teams pulumi.StringArrayInput
}

func (DashboardGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardGroupArgs)(nil)).Elem()
}

type DashboardGroupInput interface {
	pulumi.Input

	ToDashboardGroupOutput() DashboardGroupOutput
	ToDashboardGroupOutputWithContext(ctx context.Context) DashboardGroupOutput
}

func (*DashboardGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardGroup)(nil))
}

func (i *DashboardGroup) ToDashboardGroupOutput() DashboardGroupOutput {
	return i.ToDashboardGroupOutputWithContext(context.Background())
}

func (i *DashboardGroup) ToDashboardGroupOutputWithContext(ctx context.Context) DashboardGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardGroupOutput)
}

func (i *DashboardGroup) ToDashboardGroupPtrOutput() DashboardGroupPtrOutput {
	return i.ToDashboardGroupPtrOutputWithContext(context.Background())
}

func (i *DashboardGroup) ToDashboardGroupPtrOutputWithContext(ctx context.Context) DashboardGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardGroupPtrOutput)
}

type DashboardGroupPtrInput interface {
	pulumi.Input

	ToDashboardGroupPtrOutput() DashboardGroupPtrOutput
	ToDashboardGroupPtrOutputWithContext(ctx context.Context) DashboardGroupPtrOutput
}

type dashboardGroupPtrType DashboardGroupArgs

func (*dashboardGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardGroup)(nil))
}

func (i *dashboardGroupPtrType) ToDashboardGroupPtrOutput() DashboardGroupPtrOutput {
	return i.ToDashboardGroupPtrOutputWithContext(context.Background())
}

func (i *dashboardGroupPtrType) ToDashboardGroupPtrOutputWithContext(ctx context.Context) DashboardGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardGroupPtrOutput)
}

// DashboardGroupArrayInput is an input type that accepts DashboardGroupArray and DashboardGroupArrayOutput values.
// You can construct a concrete instance of `DashboardGroupArrayInput` via:
//
//          DashboardGroupArray{ DashboardGroupArgs{...} }
type DashboardGroupArrayInput interface {
	pulumi.Input

	ToDashboardGroupArrayOutput() DashboardGroupArrayOutput
	ToDashboardGroupArrayOutputWithContext(context.Context) DashboardGroupArrayOutput
}

type DashboardGroupArray []DashboardGroupInput

func (DashboardGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DashboardGroup)(nil))
}

func (i DashboardGroupArray) ToDashboardGroupArrayOutput() DashboardGroupArrayOutput {
	return i.ToDashboardGroupArrayOutputWithContext(context.Background())
}

func (i DashboardGroupArray) ToDashboardGroupArrayOutputWithContext(ctx context.Context) DashboardGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardGroupArrayOutput)
}

// DashboardGroupMapInput is an input type that accepts DashboardGroupMap and DashboardGroupMapOutput values.
// You can construct a concrete instance of `DashboardGroupMapInput` via:
//
//          DashboardGroupMap{ "key": DashboardGroupArgs{...} }
type DashboardGroupMapInput interface {
	pulumi.Input

	ToDashboardGroupMapOutput() DashboardGroupMapOutput
	ToDashboardGroupMapOutputWithContext(context.Context) DashboardGroupMapOutput
}

type DashboardGroupMap map[string]DashboardGroupInput

func (DashboardGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DashboardGroup)(nil))
}

func (i DashboardGroupMap) ToDashboardGroupMapOutput() DashboardGroupMapOutput {
	return i.ToDashboardGroupMapOutputWithContext(context.Background())
}

func (i DashboardGroupMap) ToDashboardGroupMapOutputWithContext(ctx context.Context) DashboardGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardGroupMapOutput)
}

type DashboardGroupOutput struct {
	*pulumi.OutputState
}

func (DashboardGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardGroup)(nil))
}

func (o DashboardGroupOutput) ToDashboardGroupOutput() DashboardGroupOutput {
	return o
}

func (o DashboardGroupOutput) ToDashboardGroupOutputWithContext(ctx context.Context) DashboardGroupOutput {
	return o
}

func (o DashboardGroupOutput) ToDashboardGroupPtrOutput() DashboardGroupPtrOutput {
	return o.ToDashboardGroupPtrOutputWithContext(context.Background())
}

func (o DashboardGroupOutput) ToDashboardGroupPtrOutputWithContext(ctx context.Context) DashboardGroupPtrOutput {
	return o.ApplyT(func(v DashboardGroup) *DashboardGroup {
		return &v
	}).(DashboardGroupPtrOutput)
}

type DashboardGroupPtrOutput struct {
	*pulumi.OutputState
}

func (DashboardGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardGroup)(nil))
}

func (o DashboardGroupPtrOutput) ToDashboardGroupPtrOutput() DashboardGroupPtrOutput {
	return o
}

func (o DashboardGroupPtrOutput) ToDashboardGroupPtrOutputWithContext(ctx context.Context) DashboardGroupPtrOutput {
	return o
}

type DashboardGroupArrayOutput struct{ *pulumi.OutputState }

func (DashboardGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DashboardGroup)(nil))
}

func (o DashboardGroupArrayOutput) ToDashboardGroupArrayOutput() DashboardGroupArrayOutput {
	return o
}

func (o DashboardGroupArrayOutput) ToDashboardGroupArrayOutputWithContext(ctx context.Context) DashboardGroupArrayOutput {
	return o
}

func (o DashboardGroupArrayOutput) Index(i pulumi.IntInput) DashboardGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DashboardGroup {
		return vs[0].([]DashboardGroup)[vs[1].(int)]
	}).(DashboardGroupOutput)
}

type DashboardGroupMapOutput struct{ *pulumi.OutputState }

func (DashboardGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardGroup)(nil))
}

func (o DashboardGroupMapOutput) ToDashboardGroupMapOutput() DashboardGroupMapOutput {
	return o
}

func (o DashboardGroupMapOutput) ToDashboardGroupMapOutputWithContext(ctx context.Context) DashboardGroupMapOutput {
	return o
}

func (o DashboardGroupMapOutput) MapIndex(k pulumi.StringInput) DashboardGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DashboardGroup {
		return vs[0].(map[string]DashboardGroup)[vs[1].(string)]
	}).(DashboardGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(DashboardGroupOutput{})
	pulumi.RegisterOutputType(DashboardGroupPtrOutput{})
	pulumi.RegisterOutputType(DashboardGroupArrayOutput{})
	pulumi.RegisterOutputType(DashboardGroupMapOutput{})
}
