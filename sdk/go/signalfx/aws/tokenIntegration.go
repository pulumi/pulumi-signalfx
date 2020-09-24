// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// SignalFx AWS CloudWatch integrations using security tokens. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
//
// > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.
//
// > **WARNING** This resource implements a part of a workflow. You must use it with `aws.Integration`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi-signalfx/sdk/v3/go/signalfx/aws"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		awsMyteamToken, err := aws.NewTokenIntegration(ctx, "awsMyteamToken", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRole(ctx, "awsSfxRole", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = aws.NewIntegration(ctx, "awsMyteam", &aws.IntegrationArgs{
// 			Enabled:       pulumi.Bool(true),
// 			IntegrationId: awsMyteamToken.ID(),
// 			Token:         pulumi.String("put_your_token_here"),
// 			Key:           pulumi.String("put_your_key_here"),
// 			Regions: pulumi.StringArray{
// 				pulumi.String("us-east-1"),
// 			},
// 			PollRate:         pulumi.Int(300),
// 			ImportCloudWatch: pulumi.Bool(true),
// 			EnableAwsUsage:   pulumi.Bool(true),
// 			CustomNamespaceSyncRules: aws.IntegrationCustomNamespaceSyncRuleArray{
// 				&aws.IntegrationCustomNamespaceSyncRuleArgs{
// 					DefaultAction: pulumi.String("Exclude"),
// 					FilterAction:  pulumi.String("Include"),
// 					FilterSource:  pulumi.String("filter('code', '200')"),
// 					Namespace:     pulumi.String("fart"),
// 				},
// 			},
// 			NamespaceSyncRules: aws.IntegrationNamespaceSyncRuleArray{
// 				&aws.IntegrationNamespaceSyncRuleArgs{
// 					DefaultAction: pulumi.String("Exclude"),
// 					FilterAction:  pulumi.String("Include"),
// 					FilterSource:  pulumi.String("filter('code', '200')"),
// 					Namespace:     pulumi.String("AWS/EC2"),
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
type TokenIntegration struct {
	pulumi.CustomResourceState

	// The name of this integration
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
	SignalfxAwsAccount pulumi.StringOutput `pulumi:"signalfxAwsAccount"`
	// The SignalFx-generated AWS token to use with an AWS integration.
	TokenId pulumi.StringOutput `pulumi:"tokenId"`
}

// NewTokenIntegration registers a new resource with the given unique name, arguments, and options.
func NewTokenIntegration(ctx *pulumi.Context,
	name string, args *TokenIntegrationArgs, opts ...pulumi.ResourceOption) (*TokenIntegration, error) {
	if args == nil {
		args = &TokenIntegrationArgs{}
	}
	var resource TokenIntegration
	err := ctx.RegisterResource("signalfx:aws/tokenIntegration:TokenIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTokenIntegration gets an existing TokenIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTokenIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenIntegrationState, opts ...pulumi.ResourceOption) (*TokenIntegration, error) {
	var resource TokenIntegration
	err := ctx.ReadResource("signalfx:aws/tokenIntegration:TokenIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TokenIntegration resources.
type tokenIntegrationState struct {
	// The name of this integration
	Name *string `pulumi:"name"`
	// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
	SignalfxAwsAccount *string `pulumi:"signalfxAwsAccount"`
	// The SignalFx-generated AWS token to use with an AWS integration.
	TokenId *string `pulumi:"tokenId"`
}

type TokenIntegrationState struct {
	// The name of this integration
	Name pulumi.StringPtrInput
	// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
	SignalfxAwsAccount pulumi.StringPtrInput
	// The SignalFx-generated AWS token to use with an AWS integration.
	TokenId pulumi.StringPtrInput
}

func (TokenIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenIntegrationState)(nil)).Elem()
}

type tokenIntegrationArgs struct {
	// The name of this integration
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a TokenIntegration resource.
type TokenIntegrationArgs struct {
	// The name of this integration
	Name pulumi.StringPtrInput
}

func (TokenIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenIntegrationArgs)(nil)).Elem()
}
