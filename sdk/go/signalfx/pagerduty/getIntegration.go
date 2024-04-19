// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pagerduty

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an existing PagerDuty integration.
//
// ## Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/pagerduty"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pagerduty.LookupIntegration(ctx, &pagerduty.LookupIntegrationArgs{
//				Name: "PD-Integration",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Arguments
//
// * `name` - Specify the exact name of the desired PagerDuty integration
//
// ## Attributes
//
// * `id` - The ID of the integration.
// * `name` - The name of the integration.
// * `enabled` - Whether the integration is enabled.
func LookupIntegration(ctx *pulumi.Context, args *LookupIntegrationArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIntegrationResult
	err := ctx.Invoke("signalfx:pagerduty/getIntegration:getIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIntegration.
type LookupIntegrationArgs struct {
	Name string `pulumi:"name"`
}

// A collection of values returned by getIntegration.
type LookupIntegrationResult struct {
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupIntegrationOutput(ctx *pulumi.Context, args LookupIntegrationOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationResult, error) {
			args := v.(LookupIntegrationArgs)
			r, err := LookupIntegration(ctx, &args, opts...)
			var s LookupIntegrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationResultOutput)
}

// A collection of arguments for invoking getIntegration.
type LookupIntegrationOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationArgs)(nil)).Elem()
}

// A collection of values returned by getIntegration.
type LookupIntegrationResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationResult)(nil)).Elem()
}

func (o LookupIntegrationResultOutput) ToLookupIntegrationResultOutput() LookupIntegrationResultOutput {
	return o
}

func (o LookupIntegrationResultOutput) ToLookupIntegrationResultOutputWithContext(ctx context.Context) LookupIntegrationResultOutput {
	return o
}

func (o LookupIntegrationResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIntegrationResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIntegrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationResultOutput{})
}
