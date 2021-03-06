// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SignalFx AWS CloudWatch integrations using Role ARNs. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
//
// > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.
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
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
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
// 		awsSfxRole, err := iam.NewRole(ctx, "awsSfxRole", &iam.RoleArgs{
// 			Description: pulumi.String("signalfx integration to read out data and send it to signalfxs aws account"),
// 			AssumeRolePolicy: signalfxAssumePolicy.ApplyT(func(signalfxAssumePolicy iam.GetPolicyDocumentResult) (string, error) {
// 				return signalfxAssumePolicy.Json, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		awsReadPermissions, err := iam.NewPolicy(ctx, "awsReadPermissions", &iam.PolicyArgs{
// 			Description: pulumi.String("farts"),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "	\"Version\": \"2012-10-17\",\n", "	\"Statement\": [\n", "		{\n", "			\"Action\": [\n", "				\"dynamodb:ListTables\",\n", "		    \"dynamodb:DescribeTable\",\n", "		    \"dynamodb:ListTagsOfResource\",\n", "		    \"ec2:DescribeInstances\",\n", "		    \"ec2:DescribeInstanceStatus\",\n", "		    \"ec2:DescribeVolumes\",\n", "		    \"ec2:DescribeReservedInstances\",\n", "		    \"ec2:DescribeReservedInstancesModifications\",\n", "		    \"ec2:DescribeTags\",\n", "		    \"organizations:DescribeOrganization\",\n", "		    \"cloudwatch:ListMetrics\",\n", "		    \"cloudwatch:GetMetricData\",\n", "		    \"cloudwatch:GetMetricStatistics\",\n", "		    \"cloudwatch:DescribeAlarms\",\n", "		    \"sqs:ListQueues\",\n", "		    \"sqs:GetQueueAttributes\",\n", "		    \"sqs:ListQueueTags\",\n", "		    \"elasticmapreduce:ListClusters\",\n", "		    \"elasticmapreduce:DescribeCluster\",\n", "		    \"kinesis:ListShards\",\n", "		    \"kinesis:ListStreams\",\n", "		    \"kinesis:DescribeStream\",\n", "		    \"kinesis:ListTagsForStream\",\n", "		    \"rds:DescribeDBInstances\",\n", "		    \"rds:ListTagsForResource\",\n", "		    \"elasticloadbalancing:DescribeLoadBalancers\",\n", "		    \"elasticloadbalancing:DescribeTags\",\n", "		    \"elasticache:describeCacheClusters\",\n", "		    \"redshift:DescribeClusters\",\n", "		    \"lambda:GetAlias\",\n", "		    \"lambda:ListFunctions\",\n", "		    \"lambda:ListTags\",\n", "		    \"autoscaling:DescribeAutoScalingGroups\",\n", "		    \"s3:ListAllMyBuckets\",\n", "		    \"s3:ListBucket\",\n", "		    \"s3:GetBucketLocation\",\n", "		    \"s3:GetBucketTagging\",\n", "		    \"ecs:ListServices\",\n", "		    \"ecs:ListTasks\",\n", "		    \"ecs:DescribeTasks\",\n", "		    \"ecs:DescribeServices\",\n", "		    \"ecs:ListClusters\",\n", "		    \"ecs:DescribeClusters\",\n", "		    \"ecs:ListTaskDefinitions\",\n", "		    \"ecs:ListTagsForResource\",\n", "		    \"apigateway:GET\",\n", "		    \"cloudfront:ListDistributions\",\n", "		    \"cloudfront:ListTagsForResource\",\n", "		    \"tag:GetResources\",\n", "		    \"es:ListDomainNames\",\n", "		    \"es:DescribeElasticsearchDomain\"\n", "			],\n", "			\"Effect\": \"Allow\",\n", "			\"Resource\": \"*\"\n", "		}\n", "	]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicyAttachment(ctx, "sfx_read_attach", &iam.RolePolicyAttachmentArgs{
// 			Role:      awsSfxRole.Name,
// 			PolicyArn: awsReadPermissions.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = aws.NewIntegration(ctx, "awsMyteam", &aws.IntegrationArgs{
// 			Enabled:       pulumi.Bool(true),
// 			IntegrationId: awsMyteamExtern.ID(),
// 			ExternalId:    awsMyteamExtern.ExternalId,
// 			RoleArn:       awsSfxRole.Arn,
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
	return reflect.TypeOf((*ExternalIntegration)(nil))
}

func (i *ExternalIntegration) ToExternalIntegrationOutput() ExternalIntegrationOutput {
	return i.ToExternalIntegrationOutputWithContext(context.Background())
}

func (i *ExternalIntegration) ToExternalIntegrationOutputWithContext(ctx context.Context) ExternalIntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalIntegrationOutput)
}

func (i *ExternalIntegration) ToExternalIntegrationPtrOutput() ExternalIntegrationPtrOutput {
	return i.ToExternalIntegrationPtrOutputWithContext(context.Background())
}

func (i *ExternalIntegration) ToExternalIntegrationPtrOutputWithContext(ctx context.Context) ExternalIntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalIntegrationPtrOutput)
}

type ExternalIntegrationPtrInput interface {
	pulumi.Input

	ToExternalIntegrationPtrOutput() ExternalIntegrationPtrOutput
	ToExternalIntegrationPtrOutputWithContext(ctx context.Context) ExternalIntegrationPtrOutput
}

type externalIntegrationPtrType ExternalIntegrationArgs

func (*externalIntegrationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalIntegration)(nil))
}

func (i *externalIntegrationPtrType) ToExternalIntegrationPtrOutput() ExternalIntegrationPtrOutput {
	return i.ToExternalIntegrationPtrOutputWithContext(context.Background())
}

func (i *externalIntegrationPtrType) ToExternalIntegrationPtrOutputWithContext(ctx context.Context) ExternalIntegrationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalIntegrationPtrOutput)
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
	return reflect.TypeOf(([]*ExternalIntegration)(nil))
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
	return reflect.TypeOf((map[string]*ExternalIntegration)(nil))
}

func (i ExternalIntegrationMap) ToExternalIntegrationMapOutput() ExternalIntegrationMapOutput {
	return i.ToExternalIntegrationMapOutputWithContext(context.Background())
}

func (i ExternalIntegrationMap) ToExternalIntegrationMapOutputWithContext(ctx context.Context) ExternalIntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalIntegrationMapOutput)
}

type ExternalIntegrationOutput struct {
	*pulumi.OutputState
}

func (ExternalIntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalIntegration)(nil))
}

func (o ExternalIntegrationOutput) ToExternalIntegrationOutput() ExternalIntegrationOutput {
	return o
}

func (o ExternalIntegrationOutput) ToExternalIntegrationOutputWithContext(ctx context.Context) ExternalIntegrationOutput {
	return o
}

func (o ExternalIntegrationOutput) ToExternalIntegrationPtrOutput() ExternalIntegrationPtrOutput {
	return o.ToExternalIntegrationPtrOutputWithContext(context.Background())
}

func (o ExternalIntegrationOutput) ToExternalIntegrationPtrOutputWithContext(ctx context.Context) ExternalIntegrationPtrOutput {
	return o.ApplyT(func(v ExternalIntegration) *ExternalIntegration {
		return &v
	}).(ExternalIntegrationPtrOutput)
}

type ExternalIntegrationPtrOutput struct {
	*pulumi.OutputState
}

func (ExternalIntegrationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalIntegration)(nil))
}

func (o ExternalIntegrationPtrOutput) ToExternalIntegrationPtrOutput() ExternalIntegrationPtrOutput {
	return o
}

func (o ExternalIntegrationPtrOutput) ToExternalIntegrationPtrOutputWithContext(ctx context.Context) ExternalIntegrationPtrOutput {
	return o
}

type ExternalIntegrationArrayOutput struct{ *pulumi.OutputState }

func (ExternalIntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExternalIntegration)(nil))
}

func (o ExternalIntegrationArrayOutput) ToExternalIntegrationArrayOutput() ExternalIntegrationArrayOutput {
	return o
}

func (o ExternalIntegrationArrayOutput) ToExternalIntegrationArrayOutputWithContext(ctx context.Context) ExternalIntegrationArrayOutput {
	return o
}

func (o ExternalIntegrationArrayOutput) Index(i pulumi.IntInput) ExternalIntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExternalIntegration {
		return vs[0].([]ExternalIntegration)[vs[1].(int)]
	}).(ExternalIntegrationOutput)
}

type ExternalIntegrationMapOutput struct{ *pulumi.OutputState }

func (ExternalIntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ExternalIntegration)(nil))
}

func (o ExternalIntegrationMapOutput) ToExternalIntegrationMapOutput() ExternalIntegrationMapOutput {
	return o
}

func (o ExternalIntegrationMapOutput) ToExternalIntegrationMapOutputWithContext(ctx context.Context) ExternalIntegrationMapOutput {
	return o
}

func (o ExternalIntegrationMapOutput) MapIndex(k pulumi.StringInput) ExternalIntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ExternalIntegration {
		return vs[0].(map[string]ExternalIntegration)[vs[1].(string)]
	}).(ExternalIntegrationOutput)
}

func init() {
	pulumi.RegisterOutputType(ExternalIntegrationOutput{})
	pulumi.RegisterOutputType(ExternalIntegrationPtrOutput{})
	pulumi.RegisterOutputType(ExternalIntegrationArrayOutput{})
	pulumi.RegisterOutputType(ExternalIntegrationMapOutput{})
}
