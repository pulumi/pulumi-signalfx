// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Dashboard struct {
	pulumi.CustomResourceState

	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
	//
	// Deprecated: Please use permissions_* fields now
	AuthorizedWriterTeams pulumi.StringArrayOutput `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
	//
	// Deprecated: Please use permissions fields now
	AuthorizedWriterUsers pulumi.StringArrayOutput `pulumi:"authorizedWriterUsers"`
	// Chart ID and layout information for the charts in the dashboard.
	Charts DashboardChartArrayOutput `pulumi:"charts"`
	// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`, `"low"`, `"high"`, or `"highest"`.
	ChartsResolution pulumi.StringPtrOutput `pulumi:"chartsResolution"`
	// Column layout. Charts listed will be placed in a single column with the same width and height.
	Columns DashboardColumnArrayOutput `pulumi:"columns"`
	// The ID of the dashboard group that contains the dashboard.
	DashboardGroup pulumi.StringOutput `pulumi:"dashboardGroup"`
	// Description of the dashboard.
	Description               pulumi.StringPtrOutput   `pulumi:"description"`
	DiscoveryOptionsQuery     pulumi.StringPtrOutput   `pulumi:"discoveryOptionsQuery"`
	DiscoveryOptionsSelectors pulumi.StringArrayOutput `pulumi:"discoveryOptionsSelectors"`
	// Seconds since epoch. Used for visualization.
	EndTime pulumi.IntPtrOutput `pulumi:"endTime"`
	// Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selectedEventOverlay` property instead.
	EventOverlays DashboardEventOverlayArrayOutput `pulumi:"eventOverlays"`
	// Filter to apply to the charts when displaying the dashboard.
	Filters DashboardFilterArrayOutput `pulumi:"filters"`
	// Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
	Grids DashboardGridArrayOutput `pulumi:"grids"`
	// Name of the dashboard.
	Name pulumi.StringOutput `pulumi:"name"`
	// [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) Controls who can view and/or edit your dashboard. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
	Permissions DashboardPermissionsOutput `pulumi:"permissions"`
	// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
	SelectedEventOverlays DashboardSelectedEventOverlayArrayOutput `pulumi:"selectedEventOverlays"`
	// Seconds since epoch. Used for visualization.
	StartTime pulumi.IntPtrOutput `pulumi:"startTime"`
	// Tags of the dashboard.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The time range prior to now to visualize. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`).
	TimeRange pulumi.StringPtrOutput `pulumi:"timeRange"`
	// The URL of the dashboard.
	Url pulumi.StringOutput `pulumi:"url"`
	// Dashboard variable to apply to each chart in the dashboard.
	Variables DashboardVariableArrayOutput `pulumi:"variables"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DashboardGroup == nil {
		return nil, errors.New("invalid value for required argument 'DashboardGroup'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Dashboard
	err := ctx.RegisterResource("signalfx:index/dashboard:Dashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboard gets an existing Dashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardState, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	var resource Dashboard
	err := ctx.ReadResource("signalfx:index/dashboard:Dashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dashboard resources.
type dashboardState struct {
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
	//
	// Deprecated: Please use permissions_* fields now
	AuthorizedWriterTeams []string `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
	//
	// Deprecated: Please use permissions fields now
	AuthorizedWriterUsers []string `pulumi:"authorizedWriterUsers"`
	// Chart ID and layout information for the charts in the dashboard.
	Charts []DashboardChart `pulumi:"charts"`
	// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`, `"low"`, `"high"`, or `"highest"`.
	ChartsResolution *string `pulumi:"chartsResolution"`
	// Column layout. Charts listed will be placed in a single column with the same width and height.
	Columns []DashboardColumn `pulumi:"columns"`
	// The ID of the dashboard group that contains the dashboard.
	DashboardGroup *string `pulumi:"dashboardGroup"`
	// Description of the dashboard.
	Description               *string  `pulumi:"description"`
	DiscoveryOptionsQuery     *string  `pulumi:"discoveryOptionsQuery"`
	DiscoveryOptionsSelectors []string `pulumi:"discoveryOptionsSelectors"`
	// Seconds since epoch. Used for visualization.
	EndTime *int `pulumi:"endTime"`
	// Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selectedEventOverlay` property instead.
	EventOverlays []DashboardEventOverlay `pulumi:"eventOverlays"`
	// Filter to apply to the charts when displaying the dashboard.
	Filters []DashboardFilter `pulumi:"filters"`
	// Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
	Grids []DashboardGrid `pulumi:"grids"`
	// Name of the dashboard.
	Name *string `pulumi:"name"`
	// [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) Controls who can view and/or edit your dashboard. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
	Permissions *DashboardPermissions `pulumi:"permissions"`
	// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
	SelectedEventOverlays []DashboardSelectedEventOverlay `pulumi:"selectedEventOverlays"`
	// Seconds since epoch. Used for visualization.
	StartTime *int `pulumi:"startTime"`
	// Tags of the dashboard.
	Tags []string `pulumi:"tags"`
	// The time range prior to now to visualize. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`).
	TimeRange *string `pulumi:"timeRange"`
	// The URL of the dashboard.
	Url *string `pulumi:"url"`
	// Dashboard variable to apply to each chart in the dashboard.
	Variables []DashboardVariable `pulumi:"variables"`
}

type DashboardState struct {
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
	//
	// Deprecated: Please use permissions_* fields now
	AuthorizedWriterTeams pulumi.StringArrayInput
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
	//
	// Deprecated: Please use permissions fields now
	AuthorizedWriterUsers pulumi.StringArrayInput
	// Chart ID and layout information for the charts in the dashboard.
	Charts DashboardChartArrayInput
	// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`, `"low"`, `"high"`, or `"highest"`.
	ChartsResolution pulumi.StringPtrInput
	// Column layout. Charts listed will be placed in a single column with the same width and height.
	Columns DashboardColumnArrayInput
	// The ID of the dashboard group that contains the dashboard.
	DashboardGroup pulumi.StringPtrInput
	// Description of the dashboard.
	Description               pulumi.StringPtrInput
	DiscoveryOptionsQuery     pulumi.StringPtrInput
	DiscoveryOptionsSelectors pulumi.StringArrayInput
	// Seconds since epoch. Used for visualization.
	EndTime pulumi.IntPtrInput
	// Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selectedEventOverlay` property instead.
	EventOverlays DashboardEventOverlayArrayInput
	// Filter to apply to the charts when displaying the dashboard.
	Filters DashboardFilterArrayInput
	// Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
	Grids DashboardGridArrayInput
	// Name of the dashboard.
	Name pulumi.StringPtrInput
	// [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) Controls who can view and/or edit your dashboard. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
	Permissions DashboardPermissionsPtrInput
	// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
	SelectedEventOverlays DashboardSelectedEventOverlayArrayInput
	// Seconds since epoch. Used for visualization.
	StartTime pulumi.IntPtrInput
	// Tags of the dashboard.
	Tags pulumi.StringArrayInput
	// The time range prior to now to visualize. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`).
	TimeRange pulumi.StringPtrInput
	// The URL of the dashboard.
	Url pulumi.StringPtrInput
	// Dashboard variable to apply to each chart in the dashboard.
	Variables DashboardVariableArrayInput
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
	//
	// Deprecated: Please use permissions_* fields now
	AuthorizedWriterTeams []string `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
	//
	// Deprecated: Please use permissions fields now
	AuthorizedWriterUsers []string `pulumi:"authorizedWriterUsers"`
	// Chart ID and layout information for the charts in the dashboard.
	Charts []DashboardChart `pulumi:"charts"`
	// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`, `"low"`, `"high"`, or `"highest"`.
	ChartsResolution *string `pulumi:"chartsResolution"`
	// Column layout. Charts listed will be placed in a single column with the same width and height.
	Columns []DashboardColumn `pulumi:"columns"`
	// The ID of the dashboard group that contains the dashboard.
	DashboardGroup string `pulumi:"dashboardGroup"`
	// Description of the dashboard.
	Description               *string  `pulumi:"description"`
	DiscoveryOptionsQuery     *string  `pulumi:"discoveryOptionsQuery"`
	DiscoveryOptionsSelectors []string `pulumi:"discoveryOptionsSelectors"`
	// Seconds since epoch. Used for visualization.
	EndTime *int `pulumi:"endTime"`
	// Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selectedEventOverlay` property instead.
	EventOverlays []DashboardEventOverlay `pulumi:"eventOverlays"`
	// Filter to apply to the charts when displaying the dashboard.
	Filters []DashboardFilter `pulumi:"filters"`
	// Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
	Grids []DashboardGrid `pulumi:"grids"`
	// Name of the dashboard.
	Name *string `pulumi:"name"`
	// [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) Controls who can view and/or edit your dashboard. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
	Permissions *DashboardPermissions `pulumi:"permissions"`
	// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
	SelectedEventOverlays []DashboardSelectedEventOverlay `pulumi:"selectedEventOverlays"`
	// Seconds since epoch. Used for visualization.
	StartTime *int `pulumi:"startTime"`
	// Tags of the dashboard.
	Tags []string `pulumi:"tags"`
	// The time range prior to now to visualize. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`).
	TimeRange *string `pulumi:"timeRange"`
	// Dashboard variable to apply to each chart in the dashboard.
	Variables []DashboardVariable `pulumi:"variables"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
	//
	// Deprecated: Please use permissions_* fields now
	AuthorizedWriterTeams pulumi.StringArrayInput
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
	//
	// Deprecated: Please use permissions fields now
	AuthorizedWriterUsers pulumi.StringArrayInput
	// Chart ID and layout information for the charts in the dashboard.
	Charts DashboardChartArrayInput
	// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`, `"low"`, `"high"`, or `"highest"`.
	ChartsResolution pulumi.StringPtrInput
	// Column layout. Charts listed will be placed in a single column with the same width and height.
	Columns DashboardColumnArrayInput
	// The ID of the dashboard group that contains the dashboard.
	DashboardGroup pulumi.StringInput
	// Description of the dashboard.
	Description               pulumi.StringPtrInput
	DiscoveryOptionsQuery     pulumi.StringPtrInput
	DiscoveryOptionsSelectors pulumi.StringArrayInput
	// Seconds since epoch. Used for visualization.
	EndTime pulumi.IntPtrInput
	// Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selectedEventOverlay` property instead.
	EventOverlays DashboardEventOverlayArrayInput
	// Filter to apply to the charts when displaying the dashboard.
	Filters DashboardFilterArrayInput
	// Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
	Grids DashboardGridArrayInput
	// Name of the dashboard.
	Name pulumi.StringPtrInput
	// [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) Controls who can view and/or edit your dashboard. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
	Permissions DashboardPermissionsPtrInput
	// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
	SelectedEventOverlays DashboardSelectedEventOverlayArrayInput
	// Seconds since epoch. Used for visualization.
	StartTime pulumi.IntPtrInput
	// Tags of the dashboard.
	Tags pulumi.StringArrayInput
	// The time range prior to now to visualize. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`).
	TimeRange pulumi.StringPtrInput
	// Dashboard variable to apply to each chart in the dashboard.
	Variables DashboardVariableArrayInput
}

func (DashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardArgs)(nil)).Elem()
}

type DashboardInput interface {
	pulumi.Input

	ToDashboardOutput() DashboardOutput
	ToDashboardOutputWithContext(ctx context.Context) DashboardOutput
}

func (*Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (i *Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

// DashboardArrayInput is an input type that accepts DashboardArray and DashboardArrayOutput values.
// You can construct a concrete instance of `DashboardArrayInput` via:
//
//	DashboardArray{ DashboardArgs{...} }
type DashboardArrayInput interface {
	pulumi.Input

	ToDashboardArrayOutput() DashboardArrayOutput
	ToDashboardArrayOutputWithContext(context.Context) DashboardArrayOutput
}

type DashboardArray []DashboardInput

func (DashboardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dashboard)(nil)).Elem()
}

func (i DashboardArray) ToDashboardArrayOutput() DashboardArrayOutput {
	return i.ToDashboardArrayOutputWithContext(context.Background())
}

func (i DashboardArray) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardArrayOutput)
}

// DashboardMapInput is an input type that accepts DashboardMap and DashboardMapOutput values.
// You can construct a concrete instance of `DashboardMapInput` via:
//
//	DashboardMap{ "key": DashboardArgs{...} }
type DashboardMapInput interface {
	pulumi.Input

	ToDashboardMapOutput() DashboardMapOutput
	ToDashboardMapOutputWithContext(context.Context) DashboardMapOutput
}

type DashboardMap map[string]DashboardInput

func (DashboardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dashboard)(nil)).Elem()
}

func (i DashboardMap) ToDashboardMapOutput() DashboardMapOutput {
	return i.ToDashboardMapOutputWithContext(context.Background())
}

func (i DashboardMap) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardMapOutput)
}

type DashboardOutput struct{ *pulumi.OutputState }

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
//
// Deprecated: Please use permissions_* fields now
func (o DashboardOutput) AuthorizedWriterTeams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringArrayOutput { return v.AuthorizedWriterTeams }).(pulumi.StringArrayOutput)
}

// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`). **Note:** Deprecated use `permissions` instead.
//
// Deprecated: Please use permissions fields now
func (o DashboardOutput) AuthorizedWriterUsers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringArrayOutput { return v.AuthorizedWriterUsers }).(pulumi.StringArrayOutput)
}

// Chart ID and layout information for the charts in the dashboard.
func (o DashboardOutput) Charts() DashboardChartArrayOutput {
	return o.ApplyT(func(v *Dashboard) DashboardChartArrayOutput { return v.Charts }).(DashboardChartArrayOutput)
}

// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`, `"low"`, `"high"`, or `"highest"`.
func (o DashboardOutput) ChartsResolution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringPtrOutput { return v.ChartsResolution }).(pulumi.StringPtrOutput)
}

// Column layout. Charts listed will be placed in a single column with the same width and height.
func (o DashboardOutput) Columns() DashboardColumnArrayOutput {
	return o.ApplyT(func(v *Dashboard) DashboardColumnArrayOutput { return v.Columns }).(DashboardColumnArrayOutput)
}

// The ID of the dashboard group that contains the dashboard.
func (o DashboardOutput) DashboardGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.DashboardGroup }).(pulumi.StringOutput)
}

// Description of the dashboard.
func (o DashboardOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DashboardOutput) DiscoveryOptionsQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringPtrOutput { return v.DiscoveryOptionsQuery }).(pulumi.StringPtrOutput)
}

func (o DashboardOutput) DiscoveryOptionsSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringArrayOutput { return v.DiscoveryOptionsSelectors }).(pulumi.StringArrayOutput)
}

// Seconds since epoch. Used for visualization.
func (o DashboardOutput) EndTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.IntPtrOutput { return v.EndTime }).(pulumi.IntPtrOutput)
}

// Specify a list of event overlays to include in the dashboard. Note: These overlays correspond to the *suggested* event overlays specified in the web UI, and they're not automatically applied as active overlays. To set default active event overlays, use the `selectedEventOverlay` property instead.
func (o DashboardOutput) EventOverlays() DashboardEventOverlayArrayOutput {
	return o.ApplyT(func(v *Dashboard) DashboardEventOverlayArrayOutput { return v.EventOverlays }).(DashboardEventOverlayArrayOutput)
}

// Filter to apply to the charts when displaying the dashboard.
func (o DashboardOutput) Filters() DashboardFilterArrayOutput {
	return o.ApplyT(func(v *Dashboard) DashboardFilterArrayOutput { return v.Filters }).(DashboardFilterArrayOutput)
}

// Grid dashboard layout. Charts listed will be placed in a grid by row with the same width and height. If a chart cannot fit in a row, it will be placed automatically in the next row.
func (o DashboardOutput) Grids() DashboardGridArrayOutput {
	return o.ApplyT(func(v *Dashboard) DashboardGridArrayOutput { return v.Grids }).(DashboardGridArrayOutput)
}

// Name of the dashboard.
func (o DashboardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [Permissions](https://docs.splunk.com/Observability/infrastructure/terms-concepts/permissions.html) Controls who can view and/or edit your dashboard. **Note:** This feature is not present in all accounts. Please contact support if you are unsure.
func (o DashboardOutput) Permissions() DashboardPermissionsOutput {
	return o.ApplyT(func(v *Dashboard) DashboardPermissionsOutput { return v.Permissions }).(DashboardPermissionsOutput)
}

// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
func (o DashboardOutput) SelectedEventOverlays() DashboardSelectedEventOverlayArrayOutput {
	return o.ApplyT(func(v *Dashboard) DashboardSelectedEventOverlayArrayOutput { return v.SelectedEventOverlays }).(DashboardSelectedEventOverlayArrayOutput)
}

// Seconds since epoch. Used for visualization.
func (o DashboardOutput) StartTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.IntPtrOutput { return v.StartTime }).(pulumi.IntPtrOutput)
}

// Tags of the dashboard.
func (o DashboardOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The time range prior to now to visualize. Splunk Observability Cloud time syntax (e.g. `"-5m"`, `"-1h"`).
func (o DashboardOutput) TimeRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringPtrOutput { return v.TimeRange }).(pulumi.StringPtrOutput)
}

// The URL of the dashboard.
func (o DashboardOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Dashboard variable to apply to each chart in the dashboard.
func (o DashboardOutput) Variables() DashboardVariableArrayOutput {
	return o.ApplyT(func(v *Dashboard) DashboardVariableArrayOutput { return v.Variables }).(DashboardVariableArrayOutput)
}

type DashboardArrayOutput struct{ *pulumi.OutputState }

func (DashboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dashboard)(nil)).Elem()
}

func (o DashboardArrayOutput) ToDashboardArrayOutput() DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) Index(i pulumi.IntInput) DashboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dashboard {
		return vs[0].([]*Dashboard)[vs[1].(int)]
	}).(DashboardOutput)
}

type DashboardMapOutput struct{ *pulumi.OutputState }

func (DashboardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dashboard)(nil)).Elem()
}

func (o DashboardMapOutput) ToDashboardMapOutput() DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) MapIndex(k pulumi.StringInput) DashboardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dashboard {
		return vs[0].(map[string]*Dashboard)[vs[1].(string)]
	}).(DashboardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardInput)(nil)).Elem(), &Dashboard{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardArrayInput)(nil)).Elem(), DashboardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardMapInput)(nil)).Elem(), DashboardMap{})
	pulumi.RegisterOutputType(DashboardOutput{})
	pulumi.RegisterOutputType(DashboardArrayOutput{})
	pulumi.RegisterOutputType(DashboardMapOutput{})
}
