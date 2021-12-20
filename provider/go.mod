module github.com/pulumi/pulumi-signalfx/provider/v5

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.11.0
	github.com/pulumi/pulumi/sdk/v3 v3.17.0
	github.com/splunk-terraform/terraform-provider-signalfx v1.9.2-0.20200914150050-df1863718382
)

replace github.com/splunk-terraform/terraform-provider-signalfx => github.com/pulumi/terraform-provider-signalfx v1.9.2-0.20211220170238-19f8b573964a
