// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Dashboard struct {
	pulumi.CustomResourceState

	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
	AuthorizedWriterTeams pulumi.StringArrayOutput `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers pulumi.StringArrayOutput `pulumi:"authorizedWriterUsers"`
	// Chart ID and layout information for the charts in the dashboard.
	Charts DashboardChartArrayOutput `pulumi:"charts"`
	// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
	ChartsResolution pulumi.StringPtrOutput `pulumi:"chartsResolution"`
	// Column number for the layout.
	Columns DashboardColumnArrayOutput `pulumi:"columns"`
	// The ID of the dashboard group that contains the dashboard.
	DashboardGroup pulumi.StringOutput `pulumi:"dashboardGroup"`
	// Variable description.
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
	// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
	SelectedEventOverlays DashboardSelectedEventOverlayArrayOutput `pulumi:"selectedEventOverlays"`
	// Seconds since epoch. Used for visualization.
	StartTime pulumi.IntPtrOutput `pulumi:"startTime"`
	// The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
	TimeRange pulumi.StringPtrOutput `pulumi:"timeRange"`
	// The URL of the dashboard.
	Url pulumi.StringOutput `pulumi:"url"`
	// Dashboard variable to apply to each chart in the dashboard.
	Variables DashboardVariableArrayOutput `pulumi:"variables"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil || args.DashboardGroup == nil {
		return nil, errors.New("missing required argument 'DashboardGroup'")
	}
	if args == nil {
		args = &DashboardArgs{}
	}
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
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
	AuthorizedWriterTeams []string `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers []string `pulumi:"authorizedWriterUsers"`
	// Chart ID and layout information for the charts in the dashboard.
	Charts []DashboardChart `pulumi:"charts"`
	// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
	ChartsResolution *string `pulumi:"chartsResolution"`
	// Column number for the layout.
	Columns []DashboardColumn `pulumi:"columns"`
	// The ID of the dashboard group that contains the dashboard.
	DashboardGroup *string `pulumi:"dashboardGroup"`
	// Variable description.
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
	// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
	SelectedEventOverlays []DashboardSelectedEventOverlay `pulumi:"selectedEventOverlays"`
	// Seconds since epoch. Used for visualization.
	StartTime *int `pulumi:"startTime"`
	// The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
	TimeRange *string `pulumi:"timeRange"`
	// The URL of the dashboard.
	Url *string `pulumi:"url"`
	// Dashboard variable to apply to each chart in the dashboard.
	Variables []DashboardVariable `pulumi:"variables"`
}

type DashboardState struct {
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
	AuthorizedWriterTeams pulumi.StringArrayInput
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers pulumi.StringArrayInput
	// Chart ID and layout information for the charts in the dashboard.
	Charts DashboardChartArrayInput
	// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
	ChartsResolution pulumi.StringPtrInput
	// Column number for the layout.
	Columns DashboardColumnArrayInput
	// The ID of the dashboard group that contains the dashboard.
	DashboardGroup pulumi.StringPtrInput
	// Variable description.
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
	// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
	SelectedEventOverlays DashboardSelectedEventOverlayArrayInput
	// Seconds since epoch. Used for visualization.
	StartTime pulumi.IntPtrInput
	// The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
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
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
	AuthorizedWriterTeams []string `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers []string `pulumi:"authorizedWriterUsers"`
	// Chart ID and layout information for the charts in the dashboard.
	Charts []DashboardChart `pulumi:"charts"`
	// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
	ChartsResolution *string `pulumi:"chartsResolution"`
	// Column number for the layout.
	Columns []DashboardColumn `pulumi:"columns"`
	// The ID of the dashboard group that contains the dashboard.
	DashboardGroup string `pulumi:"dashboardGroup"`
	// Variable description.
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
	// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
	SelectedEventOverlays []DashboardSelectedEventOverlay `pulumi:"selectedEventOverlays"`
	// Seconds since epoch. Used for visualization.
	StartTime *int `pulumi:"startTime"`
	// The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
	TimeRange *string `pulumi:"timeRange"`
	// Dashboard variable to apply to each chart in the dashboard.
	Variables []DashboardVariable `pulumi:"variables"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// Team IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's team (or user id in `authorizedWriterTeams`).
	AuthorizedWriterTeams pulumi.StringArrayInput
	// User IDs that have write access to this dashboard group. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers pulumi.StringArrayInput
	// Chart ID and layout information for the charts in the dashboard.
	Charts DashboardChartArrayInput
	// Specifies the chart data display resolution for charts in this dashboard. Value can be one of `"default"`,  `"low"`, `"high"`, or  `"highest"`.
	ChartsResolution pulumi.StringPtrInput
	// Column number for the layout.
	Columns DashboardColumnArrayInput
	// The ID of the dashboard group that contains the dashboard.
	DashboardGroup pulumi.StringInput
	// Variable description.
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
	// Defines event overlays which are enabled by **default**. Any overlay specified here should have an accompanying entry in `eventOverlay`, which are similar to the properties here.
	SelectedEventOverlays DashboardSelectedEventOverlayArrayInput
	// Seconds since epoch. Used for visualization.
	StartTime pulumi.IntPtrInput
	// The time range prior to now to visualize. SignalFx time syntax (e.g. `"-5m"`, `"-1h"`).
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

func (Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((*Dashboard)(nil)).Elem()
}

func (i Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

type DashboardOutput struct {
	*pulumi.OutputState
}

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardOutput)(nil)).Elem()
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DashboardOutput{})
}
