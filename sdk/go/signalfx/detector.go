// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a SignalFx detector resource. This can be used to create and manage detectors.
//
// > **NOTE** If you're interested in using SignalFx detector features such as Historical Anomaly, Resource Running Out, or others then consider building them in the UI first then using the "Show SignalFlow" feature to extract the value for `programText`. You may also consult the [documentation for detector functions in signalflow-library](https://github.com/signalfx/signalflow-library/tree/master/library/signalfx/detectors).
//
// ## Notification Format
//
// As SignalFx supports different notification mechanisms a comma-delimited string is used to provide inputs. If you'd like to specify multiple notifications, then each should be a member in the list, like so:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// This will likely be changed in a future iteration of the provider. See [SignalFx Docs](https://developers.signalfx.com/detectors_reference.html#operation/Create%20Single%20Detector) for more information. For now, here are some example of how to configure each notification type:
//
// ### Email
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ### Jira
//
// Note that the `credentialId` is the SignalFx-provided ID shown after setting up your Jira integration. (See also `jira.Integration`.)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ### Opsgenie
//
// Note that the `credentialId` is the SignalFx-provided ID shown after setting up your Opsgenie integration. `Team` here is hardcoded as the `responderType` as that is the only acceptable type as per the API docs.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ### PagerDuty
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ### Slack
//
// Exclude the `#` on the channel name!
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ### Team
//
// Sends [notifications to a team](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html).
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ### TeamEmail
//
// Sends an email to every member of a team.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ### VictorOps
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ### Webhook
//
// > **NOTE** You need to include all the commas even if you only use a credential id below.
//
// You can either configure a Webhook to use an existing integration's credential id:
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// or configure one inline:
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Detectors can be imported using their string ID (recoverable from URL`/#/detector/v2/abc123/edit`, e.g.
//
// ```sh
//  $ pulumi import signalfx:index/detector:Detector application_delay abc123
// ```
type Detector struct {
	pulumi.CustomResourceState

	// Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorizedWriterUsers`).
	AuthorizedWriterTeams pulumi.StringArrayOutput `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers pulumi.StringArrayOutput `pulumi:"authorizedWriterUsers"`
	// Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
	DisableSampling pulumi.BoolPtrOutput `pulumi:"disableSampling"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrOutput `pulumi:"endTime"`
	// The resolutions of the detector alerts in milliseconds that indicate how often data is analyzed to determine if an alert should be triggered.
	LabelResolutions pulumi.IntMapOutput `pulumi:"labelResolutions"`
	// How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
	MaxDelay pulumi.IntPtrOutput `pulumi:"maxDelay"`
	// How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
	MinDelay pulumi.IntPtrOutput `pulumi:"minDelay"`
	// Name of the detector.
	Name pulumi.StringOutput `pulumi:"name"`
	// Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringOutput `pulumi:"programText"`
	// Set of rules used for alerting.
	Rules DetectorRuleArrayOutput `pulumi:"rules"`
	// When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
	ShowDataMarkers pulumi.BoolPtrOutput `pulumi:"showDataMarkers"`
	// When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
	ShowEventLines pulumi.BoolPtrOutput `pulumi:"showEventLines"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrOutput `pulumi:"startTime"`
	// Tags associated with the detector.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Team IDs to associate the detector to.
	Teams pulumi.StringArrayOutput `pulumi:"teams"`
	// Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
	TimeRange pulumi.IntPtrOutput `pulumi:"timeRange"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// The URL of the detector.
	Url pulumi.StringOutput `pulumi:"url"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions DetectorVizOptionArrayOutput `pulumi:"vizOptions"`
}

// NewDetector registers a new resource with the given unique name, arguments, and options.
func NewDetector(ctx *pulumi.Context,
	name string, args *DetectorArgs, opts ...pulumi.ResourceOption) (*Detector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProgramText == nil {
		return nil, errors.New("invalid value for required argument 'ProgramText'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	var resource Detector
	err := ctx.RegisterResource("signalfx:index/detector:Detector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDetector gets an existing Detector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDetector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DetectorState, opts ...pulumi.ResourceOption) (*Detector, error) {
	var resource Detector
	err := ctx.ReadResource("signalfx:index/detector:Detector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Detector resources.
type detectorState struct {
	// Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorizedWriterUsers`).
	AuthorizedWriterTeams []string `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers []string `pulumi:"authorizedWriterUsers"`
	// Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
	Description *string `pulumi:"description"`
	// When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
	DisableSampling *bool `pulumi:"disableSampling"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime *int `pulumi:"endTime"`
	// The resolutions of the detector alerts in milliseconds that indicate how often data is analyzed to determine if an alert should be triggered.
	LabelResolutions map[string]int `pulumi:"labelResolutions"`
	// How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
	MaxDelay *int `pulumi:"maxDelay"`
	// How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
	MinDelay *int `pulumi:"minDelay"`
	// Name of the detector.
	Name *string `pulumi:"name"`
	// Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText *string `pulumi:"programText"`
	// Set of rules used for alerting.
	Rules []DetectorRule `pulumi:"rules"`
	// When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
	ShowDataMarkers *bool `pulumi:"showDataMarkers"`
	// When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
	ShowEventLines *bool `pulumi:"showEventLines"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime *int `pulumi:"startTime"`
	// Tags associated with the detector.
	Tags []string `pulumi:"tags"`
	// Team IDs to associate the detector to.
	Teams []string `pulumi:"teams"`
	// Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
	TimeRange *int `pulumi:"timeRange"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone *string `pulumi:"timezone"`
	// The URL of the detector.
	Url *string `pulumi:"url"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions []DetectorVizOption `pulumi:"vizOptions"`
}

type DetectorState struct {
	// Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorizedWriterUsers`).
	AuthorizedWriterTeams pulumi.StringArrayInput
	// User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers pulumi.StringArrayInput
	// Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
	Description pulumi.StringPtrInput
	// When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
	DisableSampling pulumi.BoolPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrInput
	// The resolutions of the detector alerts in milliseconds that indicate how often data is analyzed to determine if an alert should be triggered.
	LabelResolutions pulumi.IntMapInput
	// How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
	MaxDelay pulumi.IntPtrInput
	// How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
	MinDelay pulumi.IntPtrInput
	// Name of the detector.
	Name pulumi.StringPtrInput
	// Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringPtrInput
	// Set of rules used for alerting.
	Rules DetectorRuleArrayInput
	// When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
	ShowDataMarkers pulumi.BoolPtrInput
	// When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
	ShowEventLines pulumi.BoolPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrInput
	// Tags associated with the detector.
	Tags pulumi.StringArrayInput
	// Team IDs to associate the detector to.
	Teams pulumi.StringArrayInput
	// Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
	TimeRange pulumi.IntPtrInput
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrInput
	// The URL of the detector.
	Url pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement.
	VizOptions DetectorVizOptionArrayInput
}

func (DetectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*detectorState)(nil)).Elem()
}

type detectorArgs struct {
	// Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorizedWriterUsers`).
	AuthorizedWriterTeams []string `pulumi:"authorizedWriterTeams"`
	// User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers []string `pulumi:"authorizedWriterUsers"`
	// Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
	Description *string `pulumi:"description"`
	// When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
	DisableSampling *bool `pulumi:"disableSampling"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime *int `pulumi:"endTime"`
	// How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
	MaxDelay *int `pulumi:"maxDelay"`
	// How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
	MinDelay *int `pulumi:"minDelay"`
	// Name of the detector.
	Name *string `pulumi:"name"`
	// Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText string `pulumi:"programText"`
	// Set of rules used for alerting.
	Rules []DetectorRule `pulumi:"rules"`
	// When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
	ShowDataMarkers *bool `pulumi:"showDataMarkers"`
	// When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
	ShowEventLines *bool `pulumi:"showEventLines"`
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime *int `pulumi:"startTime"`
	// Tags associated with the detector.
	Tags []string `pulumi:"tags"`
	// Team IDs to associate the detector to.
	Teams []string `pulumi:"teams"`
	// Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
	TimeRange *int `pulumi:"timeRange"`
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone *string `pulumi:"timezone"`
	// Plot-level customization options, associated with a publish statement.
	VizOptions []DetectorVizOption `pulumi:"vizOptions"`
}

// The set of arguments for constructing a Detector resource.
type DetectorArgs struct {
	// Team IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's team id (or user id in `authorizedWriterUsers`).
	AuthorizedWriterTeams pulumi.StringArrayInput
	// User IDs that have write access to this detector. Remember to use an admin's token if using this feature and to include that admin's user id (or team id in `authorizedWriterTeams`).
	AuthorizedWriterUsers pulumi.StringArrayInput
	// Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
	Description pulumi.StringPtrInput
	// When `false`, the visualization may sample the output timeseries rather than displaying them all. `false` by default.
	DisableSampling pulumi.BoolPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	EndTime pulumi.IntPtrInput
	// How long (in seconds) to wait for late datapoints. See [Delayed Datapoints](https://signalfx-product-docs.readthedocs-hosted.com/en/latest/charts/chart-builder.html#delayed-datapoints) for more info. Max value is `900` seconds (15 minutes). `Auto` (as little as possible) by default.
	MaxDelay pulumi.IntPtrInput
	// How long (in seconds) to wait even if the datapoints are arriving in a timely fashion. Max value is 900 (15m).
	MinDelay pulumi.IntPtrInput
	// Name of the detector.
	Name pulumi.StringPtrInput
	// Signalflow program text for the detector. More info [in the SignalFx docs](https://developers.signalfx.com/signalflow_analytics/signalflow_overview.html#_signalflow_programming_language).
	ProgramText pulumi.StringInput
	// Set of rules used for alerting.
	Rules DetectorRuleArrayInput
	// When `true`, markers will be drawn for each datapoint within the visualization. `true` by default.
	ShowDataMarkers pulumi.BoolPtrInput
	// When `true`, the visualization will display a vertical line for each event trigger. `false` by default.
	ShowEventLines pulumi.BoolPtrInput
	// Seconds since epoch. Used for visualization. Conflicts with `timeRange`.
	StartTime pulumi.IntPtrInput
	// Tags associated with the detector.
	Tags pulumi.StringArrayInput
	// Team IDs to associate the detector to.
	Teams pulumi.StringArrayInput
	// Seconds to display in the visualization. This is a rolling range from the current time. Example: `3600` corresponds to `-1h` in web UI. `3600` by default.
	TimeRange pulumi.IntPtrInput
	// The property value is a string that denotes the geographic region associated with the time zone, (e.g. Australia/Sydney)
	Timezone pulumi.StringPtrInput
	// Plot-level customization options, associated with a publish statement.
	VizOptions DetectorVizOptionArrayInput
}

func (DetectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*detectorArgs)(nil)).Elem()
}

type DetectorInput interface {
	pulumi.Input

	ToDetectorOutput() DetectorOutput
	ToDetectorOutputWithContext(ctx context.Context) DetectorOutput
}

func (*Detector) ElementType() reflect.Type {
	return reflect.TypeOf((**Detector)(nil)).Elem()
}

func (i *Detector) ToDetectorOutput() DetectorOutput {
	return i.ToDetectorOutputWithContext(context.Background())
}

func (i *Detector) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorOutput)
}

// DetectorArrayInput is an input type that accepts DetectorArray and DetectorArrayOutput values.
// You can construct a concrete instance of `DetectorArrayInput` via:
//
//          DetectorArray{ DetectorArgs{...} }
type DetectorArrayInput interface {
	pulumi.Input

	ToDetectorArrayOutput() DetectorArrayOutput
	ToDetectorArrayOutputWithContext(context.Context) DetectorArrayOutput
}

type DetectorArray []DetectorInput

func (DetectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Detector)(nil)).Elem()
}

func (i DetectorArray) ToDetectorArrayOutput() DetectorArrayOutput {
	return i.ToDetectorArrayOutputWithContext(context.Background())
}

func (i DetectorArray) ToDetectorArrayOutputWithContext(ctx context.Context) DetectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorArrayOutput)
}

// DetectorMapInput is an input type that accepts DetectorMap and DetectorMapOutput values.
// You can construct a concrete instance of `DetectorMapInput` via:
//
//          DetectorMap{ "key": DetectorArgs{...} }
type DetectorMapInput interface {
	pulumi.Input

	ToDetectorMapOutput() DetectorMapOutput
	ToDetectorMapOutputWithContext(context.Context) DetectorMapOutput
}

type DetectorMap map[string]DetectorInput

func (DetectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Detector)(nil)).Elem()
}

func (i DetectorMap) ToDetectorMapOutput() DetectorMapOutput {
	return i.ToDetectorMapOutputWithContext(context.Background())
}

func (i DetectorMap) ToDetectorMapOutputWithContext(ctx context.Context) DetectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorMapOutput)
}

type DetectorOutput struct{ *pulumi.OutputState }

func (DetectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Detector)(nil)).Elem()
}

func (o DetectorOutput) ToDetectorOutput() DetectorOutput {
	return o
}

func (o DetectorOutput) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return o
}

type DetectorArrayOutput struct{ *pulumi.OutputState }

func (DetectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Detector)(nil)).Elem()
}

func (o DetectorArrayOutput) ToDetectorArrayOutput() DetectorArrayOutput {
	return o
}

func (o DetectorArrayOutput) ToDetectorArrayOutputWithContext(ctx context.Context) DetectorArrayOutput {
	return o
}

func (o DetectorArrayOutput) Index(i pulumi.IntInput) DetectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Detector {
		return vs[0].([]*Detector)[vs[1].(int)]
	}).(DetectorOutput)
}

type DetectorMapOutput struct{ *pulumi.OutputState }

func (DetectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Detector)(nil)).Elem()
}

func (o DetectorMapOutput) ToDetectorMapOutput() DetectorMapOutput {
	return o
}

func (o DetectorMapOutput) ToDetectorMapOutputWithContext(ctx context.Context) DetectorMapOutput {
	return o
}

func (o DetectorMapOutput) MapIndex(k pulumi.StringInput) DetectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Detector {
		return vs[0].(map[string]*Detector)[vs[1].(string)]
	}).(DetectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorInput)(nil)).Elem(), &Detector{})
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorArrayInput)(nil)).Elem(), DetectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorMapInput)(nil)).Elem(), DetectorMap{})
	pulumi.RegisterOutputType(DetectorOutput{})
	pulumi.RegisterOutputType(DetectorArrayOutput{})
	pulumi.RegisterOutputType(DetectorMapOutput{})
}
