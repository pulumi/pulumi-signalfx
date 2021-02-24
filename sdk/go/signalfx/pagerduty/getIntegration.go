// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pagerduty

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information on an existing PagerDuty integration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-signalfx/sdk/v4/go/signalfx/pagerduty"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := pagerduty.LookupIntegration(ctx, &pagerduty.LookupIntegrationArgs{
// 			Name: "PD-Integration",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupIntegration(ctx *pulumi.Context, args *LookupIntegrationArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationResult, error) {
	var rv LookupIntegrationResult
	err := ctx.Invoke("signalfx:pagerduty/getIntegration:getIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIntegration.
type LookupIntegrationArgs struct {
	// Specify the exact name of the desired PagerDuty integration
	Name string `pulumi:"name"`
}

// A collection of values returned by getIntegration.
type LookupIntegrationResult struct {
	// Whether the integration is enabled.
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the integration.
	Name string `pulumi:"name"`
}
