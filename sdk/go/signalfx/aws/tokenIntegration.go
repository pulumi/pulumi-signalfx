// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 	"github.com/pulumi/pulumi-signalfx/sdk/v5/go/signalfx/aws"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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

type TokenIntegrationInput interface {
	pulumi.Input

	ToTokenIntegrationOutput() TokenIntegrationOutput
	ToTokenIntegrationOutputWithContext(ctx context.Context) TokenIntegrationOutput
}

func (*TokenIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenIntegration)(nil))
}

func (i *TokenIntegration) ToTokenIntegrationOutput() TokenIntegrationOutput {
	return i.ToTokenIntegrationOutputWithContext(context.Background())
}

func (i *TokenIntegration) ToTokenIntegrationOutputWithContext(ctx context.Context) TokenIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenIntegrationOutput)
}

func (i *TokenIntegration) ToTokenIntegrationPtrOutput() TokenIntegrationPtrOutput {
	return i.ToTokenIntegrationPtrOutputWithContext(context.Background())
}

func (i *TokenIntegration) ToTokenIntegrationPtrOutputWithContext(ctx context.Context) TokenIntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenIntegrationPtrOutput)
}

type TokenIntegrationPtrInput interface {
	pulumi.Input

	ToTokenIntegrationPtrOutput() TokenIntegrationPtrOutput
	ToTokenIntegrationPtrOutputWithContext(ctx context.Context) TokenIntegrationPtrOutput
}

type tokenIntegrationPtrType TokenIntegrationArgs

func (*tokenIntegrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenIntegration)(nil))
}

func (i *tokenIntegrationPtrType) ToTokenIntegrationPtrOutput() TokenIntegrationPtrOutput {
	return i.ToTokenIntegrationPtrOutputWithContext(context.Background())
}

func (i *tokenIntegrationPtrType) ToTokenIntegrationPtrOutputWithContext(ctx context.Context) TokenIntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenIntegrationPtrOutput)
}

// TokenIntegrationArrayInput is an input type that accepts TokenIntegrationArray and TokenIntegrationArrayOutput values.
// You can construct a concrete instance of `TokenIntegrationArrayInput` via:
//
//          TokenIntegrationArray{ TokenIntegrationArgs{...} }
type TokenIntegrationArrayInput interface {
	pulumi.Input

	ToTokenIntegrationArrayOutput() TokenIntegrationArrayOutput
	ToTokenIntegrationArrayOutputWithContext(context.Context) TokenIntegrationArrayOutput
}

type TokenIntegrationArray []TokenIntegrationInput

func (TokenIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TokenIntegration)(nil))
}

func (i TokenIntegrationArray) ToTokenIntegrationArrayOutput() TokenIntegrationArrayOutput {
	return i.ToTokenIntegrationArrayOutputWithContext(context.Background())
}

func (i TokenIntegrationArray) ToTokenIntegrationArrayOutputWithContext(ctx context.Context) TokenIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenIntegrationArrayOutput)
}

// TokenIntegrationMapInput is an input type that accepts TokenIntegrationMap and TokenIntegrationMapOutput values.
// You can construct a concrete instance of `TokenIntegrationMapInput` via:
//
//          TokenIntegrationMap{ "key": TokenIntegrationArgs{...} }
type TokenIntegrationMapInput interface {
	pulumi.Input

	ToTokenIntegrationMapOutput() TokenIntegrationMapOutput
	ToTokenIntegrationMapOutputWithContext(context.Context) TokenIntegrationMapOutput
}

type TokenIntegrationMap map[string]TokenIntegrationInput

func (TokenIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TokenIntegration)(nil))
}

func (i TokenIntegrationMap) ToTokenIntegrationMapOutput() TokenIntegrationMapOutput {
	return i.ToTokenIntegrationMapOutputWithContext(context.Background())
}

func (i TokenIntegrationMap) ToTokenIntegrationMapOutputWithContext(ctx context.Context) TokenIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenIntegrationMapOutput)
}

type TokenIntegrationOutput struct {
	*pulumi.OutputState
}

func (TokenIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenIntegration)(nil))
}

func (o TokenIntegrationOutput) ToTokenIntegrationOutput() TokenIntegrationOutput {
	return o
}

func (o TokenIntegrationOutput) ToTokenIntegrationOutputWithContext(ctx context.Context) TokenIntegrationOutput {
	return o
}

func (o TokenIntegrationOutput) ToTokenIntegrationPtrOutput() TokenIntegrationPtrOutput {
	return o.ToTokenIntegrationPtrOutputWithContext(context.Background())
}

func (o TokenIntegrationOutput) ToTokenIntegrationPtrOutputWithContext(ctx context.Context) TokenIntegrationPtrOutput {
	return o.ApplyT(func(v TokenIntegration) *TokenIntegration {
		return &v
	}).(TokenIntegrationPtrOutput)
}

type TokenIntegrationPtrOutput struct {
	*pulumi.OutputState
}

func (TokenIntegrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenIntegration)(nil))
}

func (o TokenIntegrationPtrOutput) ToTokenIntegrationPtrOutput() TokenIntegrationPtrOutput {
	return o
}

func (o TokenIntegrationPtrOutput) ToTokenIntegrationPtrOutputWithContext(ctx context.Context) TokenIntegrationPtrOutput {
	return o
}

type TokenIntegrationArrayOutput struct{ *pulumi.OutputState }

func (TokenIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenIntegration)(nil))
}

func (o TokenIntegrationArrayOutput) ToTokenIntegrationArrayOutput() TokenIntegrationArrayOutput {
	return o
}

func (o TokenIntegrationArrayOutput) ToTokenIntegrationArrayOutputWithContext(ctx context.Context) TokenIntegrationArrayOutput {
	return o
}

func (o TokenIntegrationArrayOutput) Index(i pulumi.IntInput) TokenIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenIntegration {
		return vs[0].([]TokenIntegration)[vs[1].(int)]
	}).(TokenIntegrationOutput)
}

type TokenIntegrationMapOutput struct{ *pulumi.OutputState }

func (TokenIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TokenIntegration)(nil))
}

func (o TokenIntegrationMapOutput) ToTokenIntegrationMapOutput() TokenIntegrationMapOutput {
	return o
}

func (o TokenIntegrationMapOutput) ToTokenIntegrationMapOutputWithContext(ctx context.Context) TokenIntegrationMapOutput {
	return o
}

func (o TokenIntegrationMapOutput) MapIndex(k pulumi.StringInput) TokenIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TokenIntegration {
		return vs[0].(map[string]TokenIntegration)[vs[1].(string)]
	}).(TokenIntegrationOutput)
}

func init() {
	pulumi.RegisterOutputType(TokenIntegrationOutput{})
	pulumi.RegisterOutputType(TokenIntegrationPtrOutput{})
	pulumi.RegisterOutputType(TokenIntegrationArrayOutput{})
	pulumi.RegisterOutputType(TokenIntegrationMapOutput{})
}
