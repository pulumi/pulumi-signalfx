module github.com/pulumi/pulumi-signalfx/provider/v5

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.20.0
	github.com/pulumi/pulumi/sdk/v3 v3.27.0
	github.com/splunk-terraform/terraform-provider-signalfx v1.9.2-0.20200914150050-df1863718382
)

replace github.com/splunk-terraform/terraform-provider-signalfx => github.com/pulumi/terraform-provider-signalfx v1.9.2-0.20220329183112-559fc8c45047

replace github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0

replace github.com/hashicorp/terraform-plugin-test => github.com/hashicorp/terraform-plugin-test v1.3.0
