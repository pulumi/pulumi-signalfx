// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SignalFx AWS CloudWatch integrations using Role ARNs. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
//
// > **NOTE** When managing integrations use a session token for an administrator to authenticate the SignalFx provider. See [Operations that require a session token for an administrator].(https://dev.splunk.com/observability/docs/administration/authtokens#Operations-that-require-a-session-token-for-an-administrator).
//
// > **WARNING** This resource implements a part of a workflow. You must use it with `aws.Integration`. Check with SignalFx support for your realm's AWS account id.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi-signalfx/sdk/v5/go/signalfx/aws"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		awsMyteamExtern, err := aws.NewExternalIntegration(ctx, "awsMyteamExtern", nil)
// 		if err != nil {
// 			return err
// 		}
// 		signalfxAssumePolicy := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
// 			Statements: iam.GetPolicyDocumentStatementArray{
// 				&iam.GetPolicyDocumentStatementArgs{
// 					Actions: pulumi.StringArray{
// 						pulumi.String("sts:AssumeRole"),
// 					},
// 					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
// 						&iam.GetPolicyDocumentStatementPrincipalArgs{
// 							Type: pulumi.String("AWS"),
// 							Identifiers: pulumi.StringArray{
// 								awsMyteamExtern.SignalfxAwsAccount,
// 							},
// 						},
// 					},
// 					Conditions: iam.GetPolicyDocumentStatementConditionArray{
// 						&iam.GetPolicyDocumentStatementConditionArgs{
// 							Test:     pulumi.String("StringEquals"),
// 							Variable: pulumi.String("sts:ExternalId"),
// 							Values: pulumi.StringArray{
// 								awsMyteamExtern.ExternalId,
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		awsSplunkRole, err := iam.NewRole(ctx, "awsSplunkRole", &iam.RoleArgs{
// 			Description: pulumi.String("signalfx integration to read out data and send it to signalfxs aws account"),
// 			AssumeRolePolicy: signalfxAssumePolicy.ApplyT(func(signalfxAssumePolicy iam.GetPolicyDocumentResult) (string, error) {
// 				return signalfxAssumePolicy.Json, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		awsSplunkPolicy, err := iam.NewPolicy(ctx, "awsSplunkPolicy", &iam.PolicyArgs{
// 			Description: pulumi.String("AWS permissions required by the Splunk Observability Cloud"),
// 			Policy:      pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Action\": [\n", "        \"apigateway:GET\",\n", "        \"autoscaling:DescribeAutoScalingGroups\",\n", "        \"cloudfront:GetDistributionConfig\",\n", "        \"cloudfront:ListDistributions\",\n", "        \"cloudfront:ListTagsForResource\",\n", "        \"cloudwatch:DescribeAlarms\",\n", "        \"cloudwatch:GetMetricData\",\n", "        \"cloudwatch:GetMetricStatistics\",\n", "        \"cloudwatch:ListMetrics\",\n", "        \"directconnect:DescribeConnections\",\n", "        \"dynamodb:DescribeTable\",\n", "        \"dynamodb:ListTables\",\n", "        \"dynamodb:ListTagsOfResource\",\n", "        \"ec2:DescribeInstances\",\n", "        \"ec2:DescribeInstanceStatus\",\n", "        \"ec2:DescribeRegions\",\n", "        \"ec2:DescribeReservedInstances\",\n", "        \"ec2:DescribeReservedInstancesModifications\",\n", "        \"ec2:DescribeTags\",\n", "        \"ec2:DescribeVolumes\",\n", "        \"ecs:DescribeClusters\",\n", "        \"ecs:DescribeServices\",\n", "        \"ecs:DescribeTasks\",\n", "        \"ecs:ListClusters\",\n", "        \"ecs:ListServices\",\n", "        \"ecs:ListTagsForResource\",\n", "        \"ecs:ListTaskDefinitions\",\n", "        \"ecs:ListTasks\",\n", "        \"elasticache:DescribeCacheClusters\",\n", "        \"elasticloadbalancing:DescribeLoadBalancerAttributes\",\n", "        \"elasticloadbalancing:DescribeLoadBalancers\",\n", "        \"elasticloadbalancing:DescribeTags\",\n", "        \"elasticloadbalancing:DescribeTargetGroups\",\n", "        \"elasticmapreduce:DescribeCluster\",\n", "        \"elasticmapreduce:ListClusters\",\n", "        \"es:DescribeElasticsearchDomain\",\n", "        \"es:ListDomainNames\",\n", "        \"kinesis:DescribeStream\",\n", "        \"kinesis:ListShards\",\n", "        \"kinesis:ListStreams\",\n", "        \"kinesis:ListTagsForStream\",\n", "        \"lambda:GetAlias\",\n", "        \"lambda:ListFunctions\",\n", "        \"lambda:ListTags\",\n", "        \"logs:DeleteSubscriptionFilter\",\n", "        \"logs:DescribeLogGroups\",\n", "        \"logs:DescribeSubscriptionFilters\",\n", "        \"logs:PutSubscriptionFilter\",\n", "        \"organizations:DescribeOrganization\",\n", "        \"rds:DescribeDBClusters\",\n", "        \"rds:DescribeDBInstances\",\n", "        \"rds:ListTagsForResource\",\n", "        \"redshift:DescribeClusters\",\n", "        \"redshift:DescribeLoggingStatus\",\n", "        \"s3:GetBucketLocation\",\n", "        \"s3:GetBucketLogging\",\n", "        \"s3:GetBucketNotification\",\n", "        \"s3:GetBucketTagging\",\n", "        \"s3:ListAllMyBuckets\",\n", "        \"s3:ListBucket\",\n", "        \"s3:PutBucketNotification\",\n", "        \"sqs:GetQueueAttributes\",\n", "        \"sqs:ListQueues\",\n", "        \"sqs:ListQueueTags\",\n", "        \"states:ListStateMachines\",\n", "        \"tag:GetResources\",\n", "        \"workspaces:DescribeWorkspaces\"\n", "      ],\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "splunkRolePolicyAttach", &iam.RolePolicyAttachmentArgs{
// 			Role:      awsSplunkRole.Name,
// 			PolicyArn: awsSplunkPolicy.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = aws.NewIntegration(ctx, "awsMyteam", &aws.IntegrationArgs{
// 			Enabled:       pulumi.Bool(true),
// 			IntegrationId: awsMyteamExtern.ID(),
// 			ExternalId:    awsMyteamExtern.ExternalId,
// 			RoleArn:       awsSplunkRole.Arn,
// 			Regions: pulumi.StringArray{
// 				pulumi.String("us-east-1"),
// 			},
// 			PollRate:         pulumi.Int(300),
// 			ImportCloudWatch: pulumi.Bool(true),
// 			EnableAwsUsage:   pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ExternalIntegration struct {
	pulumi.CustomResourceState

	// The external ID to use with your IAM role and with `aws.Integration`.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// The name of this integration
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
	SignalfxAwsAccount pulumi.StringOutput `pulumi:"signalfxAwsAccount"`
}

// NewExternalIntegration registers a new resource with the given unique name, arguments, and options.
func NewExternalIntegration(ctx *pulumi.Context,
	name string, args *ExternalIntegrationArgs, opts ...pulumi.ResourceOption) (*ExternalIntegration, error) {
	if args == nil {
		args = &ExternalIntegrationArgs{}
	}

	var resource ExternalIntegration
	err := ctx.RegisterResource("signalfx:aws/externalIntegration:ExternalIntegration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalIntegration gets an existing ExternalIntegration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalIntegrationState, opts ...pulumi.ResourceOption) (*ExternalIntegration, error) {
	var resource ExternalIntegration
	err := ctx.ReadResource("signalfx:aws/externalIntegration:ExternalIntegration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalIntegration resources.
type externalIntegrationState struct {
	// The external ID to use with your IAM role and with `aws.Integration`.
	ExternalId *string `pulumi:"externalId"`
	// The name of this integration
	Name *string `pulumi:"name"`
	// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
	SignalfxAwsAccount *string `pulumi:"signalfxAwsAccount"`
}

type ExternalIntegrationState struct {
	// The external ID to use with your IAM role and with `aws.Integration`.
	ExternalId pulumi.StringPtrInput
	// The name of this integration
	Name pulumi.StringPtrInput
	// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
	SignalfxAwsAccount pulumi.StringPtrInput
}

func (ExternalIntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalIntegrationState)(nil)).Elem()
}

type externalIntegrationArgs struct {
	// The name of this integration
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ExternalIntegration resource.
type ExternalIntegrationArgs struct {
	// The name of this integration
	Name pulumi.StringPtrInput
}

func (ExternalIntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalIntegrationArgs)(nil)).Elem()
}

type ExternalIntegrationInput interface {
	pulumi.Input

	ToExternalIntegrationOutput() ExternalIntegrationOutput
	ToExternalIntegrationOutputWithContext(ctx context.Context) ExternalIntegrationOutput
}

func (*ExternalIntegration) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalIntegration)(nil)).Elem()
}

func (i *ExternalIntegration) ToExternalIntegrationOutput() ExternalIntegrationOutput {
	return i.ToExternalIntegrationOutputWithContext(context.Background())
}

func (i *ExternalIntegration) ToExternalIntegrationOutputWithContext(ctx context.Context) ExternalIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalIntegrationOutput)
}

// ExternalIntegrationArrayInput is an input type that accepts ExternalIntegrationArray and ExternalIntegrationArrayOutput values.
// You can construct a concrete instance of `ExternalIntegrationArrayInput` via:
//
//          ExternalIntegrationArray{ ExternalIntegrationArgs{...} }
type ExternalIntegrationArrayInput interface {
	pulumi.Input

	ToExternalIntegrationArrayOutput() ExternalIntegrationArrayOutput
	ToExternalIntegrationArrayOutputWithContext(context.Context) ExternalIntegrationArrayOutput
}

type ExternalIntegrationArray []ExternalIntegrationInput

func (ExternalIntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalIntegration)(nil)).Elem()
}

func (i ExternalIntegrationArray) ToExternalIntegrationArrayOutput() ExternalIntegrationArrayOutput {
	return i.ToExternalIntegrationArrayOutputWithContext(context.Background())
}

func (i ExternalIntegrationArray) ToExternalIntegrationArrayOutputWithContext(ctx context.Context) ExternalIntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalIntegrationArrayOutput)
}

// ExternalIntegrationMapInput is an input type that accepts ExternalIntegrationMap and ExternalIntegrationMapOutput values.
// You can construct a concrete instance of `ExternalIntegrationMapInput` via:
//
//          ExternalIntegrationMap{ "key": ExternalIntegrationArgs{...} }
type ExternalIntegrationMapInput interface {
	pulumi.Input

	ToExternalIntegrationMapOutput() ExternalIntegrationMapOutput
	ToExternalIntegrationMapOutputWithContext(context.Context) ExternalIntegrationMapOutput
}

type ExternalIntegrationMap map[string]ExternalIntegrationInput

func (ExternalIntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalIntegration)(nil)).Elem()
}

func (i ExternalIntegrationMap) ToExternalIntegrationMapOutput() ExternalIntegrationMapOutput {
	return i.ToExternalIntegrationMapOutputWithContext(context.Background())
}

func (i ExternalIntegrationMap) ToExternalIntegrationMapOutputWithContext(ctx context.Context) ExternalIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalIntegrationMapOutput)
}

type ExternalIntegrationOutput struct{ *pulumi.OutputState }

func (ExternalIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalIntegration)(nil)).Elem()
}

func (o ExternalIntegrationOutput) ToExternalIntegrationOutput() ExternalIntegrationOutput {
	return o
}

func (o ExternalIntegrationOutput) ToExternalIntegrationOutputWithContext(ctx context.Context) ExternalIntegrationOutput {
	return o
}

type ExternalIntegrationArrayOutput struct{ *pulumi.OutputState }

func (ExternalIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalIntegration)(nil)).Elem()
}

func (o ExternalIntegrationArrayOutput) ToExternalIntegrationArrayOutput() ExternalIntegrationArrayOutput {
	return o
}

func (o ExternalIntegrationArrayOutput) ToExternalIntegrationArrayOutputWithContext(ctx context.Context) ExternalIntegrationArrayOutput {
	return o
}

func (o ExternalIntegrationArrayOutput) Index(i pulumi.IntInput) ExternalIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalIntegration {
		return vs[0].([]*ExternalIntegration)[vs[1].(int)]
	}).(ExternalIntegrationOutput)
}

type ExternalIntegrationMapOutput struct{ *pulumi.OutputState }

func (ExternalIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalIntegration)(nil)).Elem()
}

func (o ExternalIntegrationMapOutput) ToExternalIntegrationMapOutput() ExternalIntegrationMapOutput {
	return o
}

func (o ExternalIntegrationMapOutput) ToExternalIntegrationMapOutputWithContext(ctx context.Context) ExternalIntegrationMapOutput {
	return o
}

func (o ExternalIntegrationMapOutput) MapIndex(k pulumi.StringInput) ExternalIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalIntegration {
		return vs[0].(map[string]*ExternalIntegration)[vs[1].(string)]
	}).(ExternalIntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalIntegrationInput)(nil)).Elem(), &ExternalIntegration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalIntegrationArrayInput)(nil)).Elem(), ExternalIntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalIntegrationMapInput)(nil)).Elem(), ExternalIntegrationMap{})
	pulumi.RegisterOutputType(ExternalIntegrationOutput{})
	pulumi.RegisterOutputType(ExternalIntegrationArrayOutput{})
	pulumi.RegisterOutputType(ExternalIntegrationMapOutput{})
}
