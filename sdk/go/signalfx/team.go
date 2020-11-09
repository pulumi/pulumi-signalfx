// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Handles management of SignalFx teams.
//
// You can configure [team notification policies](https://docs.signalfx.com/en/latest/managing/teams/team-notifications.html) using this resource and the various `notifications_*` properties.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-signalfx/sdk/v3/go/signalfx"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := signalfx.NewTeam(ctx, "myteam0", &signalfx.TeamArgs{
// 			Description: pulumi.String("Super great team no jerks definitely"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("userid1"),
// 				pulumi.String("userid2"),
// 			},
// 			NotificationsCriticals: pulumi.StringArray{
// 				pulumi.String("PagerDuty,credentialId"),
// 			},
// 			NotificationsInfos: pulumi.StringArray{
// 				pulumi.String("Email,notify@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Team struct {
	pulumi.CustomResourceState

	// Description of the team.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of user IDs to include in the team.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Name of the team.
	Name pulumi.StringOutput `pulumi:"name"`
	// Where to send notifications for critical alerts
	NotificationsCriticals pulumi.StringArrayOutput `pulumi:"notificationsCriticals"`
	// Where to send notifications for default alerts
	NotificationsDefaults pulumi.StringArrayOutput `pulumi:"notificationsDefaults"`
	// Where to send notifications for info alerts
	NotificationsInfos pulumi.StringArrayOutput `pulumi:"notificationsInfos"`
	// Where to send notifications for major alerts
	NotificationsMajors pulumi.StringArrayOutput `pulumi:"notificationsMajors"`
	// Where to send notifications for minor alerts
	NotificationsMinors pulumi.StringArrayOutput `pulumi:"notificationsMinors"`
	// Where to send notifications for warning alerts
	NotificationsWarnings pulumi.StringArrayOutput `pulumi:"notificationsWarnings"`
	// The URL of the team.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewTeam registers a new resource with the given unique name, arguments, and options.
func NewTeam(ctx *pulumi.Context,
	name string, args *TeamArgs, opts ...pulumi.ResourceOption) (*Team, error) {
	if args == nil {
		args = &TeamArgs{}
	}
	var resource Team
	err := ctx.RegisterResource("signalfx:index/team:Team", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeam gets an existing Team resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamState, opts ...pulumi.ResourceOption) (*Team, error) {
	var resource Team
	err := ctx.ReadResource("signalfx:index/team:Team", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Team resources.
type teamState struct {
	// Description of the team.
	Description *string `pulumi:"description"`
	// List of user IDs to include in the team.
	Members []string `pulumi:"members"`
	// Name of the team.
	Name *string `pulumi:"name"`
	// Where to send notifications for critical alerts
	NotificationsCriticals []string `pulumi:"notificationsCriticals"`
	// Where to send notifications for default alerts
	NotificationsDefaults []string `pulumi:"notificationsDefaults"`
	// Where to send notifications for info alerts
	NotificationsInfos []string `pulumi:"notificationsInfos"`
	// Where to send notifications for major alerts
	NotificationsMajors []string `pulumi:"notificationsMajors"`
	// Where to send notifications for minor alerts
	NotificationsMinors []string `pulumi:"notificationsMinors"`
	// Where to send notifications for warning alerts
	NotificationsWarnings []string `pulumi:"notificationsWarnings"`
	// The URL of the team.
	Url *string `pulumi:"url"`
}

type TeamState struct {
	// Description of the team.
	Description pulumi.StringPtrInput
	// List of user IDs to include in the team.
	Members pulumi.StringArrayInput
	// Name of the team.
	Name pulumi.StringPtrInput
	// Where to send notifications for critical alerts
	NotificationsCriticals pulumi.StringArrayInput
	// Where to send notifications for default alerts
	NotificationsDefaults pulumi.StringArrayInput
	// Where to send notifications for info alerts
	NotificationsInfos pulumi.StringArrayInput
	// Where to send notifications for major alerts
	NotificationsMajors pulumi.StringArrayInput
	// Where to send notifications for minor alerts
	NotificationsMinors pulumi.StringArrayInput
	// Where to send notifications for warning alerts
	NotificationsWarnings pulumi.StringArrayInput
	// The URL of the team.
	Url pulumi.StringPtrInput
}

func (TeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamState)(nil)).Elem()
}

type teamArgs struct {
	// Description of the team.
	Description *string `pulumi:"description"`
	// List of user IDs to include in the team.
	Members []string `pulumi:"members"`
	// Name of the team.
	Name *string `pulumi:"name"`
	// Where to send notifications for critical alerts
	NotificationsCriticals []string `pulumi:"notificationsCriticals"`
	// Where to send notifications for default alerts
	NotificationsDefaults []string `pulumi:"notificationsDefaults"`
	// Where to send notifications for info alerts
	NotificationsInfos []string `pulumi:"notificationsInfos"`
	// Where to send notifications for major alerts
	NotificationsMajors []string `pulumi:"notificationsMajors"`
	// Where to send notifications for minor alerts
	NotificationsMinors []string `pulumi:"notificationsMinors"`
	// Where to send notifications for warning alerts
	NotificationsWarnings []string `pulumi:"notificationsWarnings"`
}

// The set of arguments for constructing a Team resource.
type TeamArgs struct {
	// Description of the team.
	Description pulumi.StringPtrInput
	// List of user IDs to include in the team.
	Members pulumi.StringArrayInput
	// Name of the team.
	Name pulumi.StringPtrInput
	// Where to send notifications for critical alerts
	NotificationsCriticals pulumi.StringArrayInput
	// Where to send notifications for default alerts
	NotificationsDefaults pulumi.StringArrayInput
	// Where to send notifications for info alerts
	NotificationsInfos pulumi.StringArrayInput
	// Where to send notifications for major alerts
	NotificationsMajors pulumi.StringArrayInput
	// Where to send notifications for minor alerts
	NotificationsMinors pulumi.StringArrayInput
	// Where to send notifications for warning alerts
	NotificationsWarnings pulumi.StringArrayInput
}

func (TeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamArgs)(nil)).Elem()
}
