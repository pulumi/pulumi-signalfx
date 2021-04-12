module github.com/pulumi/pulumi-signalfx/provider/v4

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.23.0
	github.com/pulumi/pulumi/sdk/v2 v2.24.1
	github.com/splunk-terraform/terraform-provider-signalfx v1.9.2-0.20200914150050-df1863718382
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/splunk-terraform/terraform-provider-signalfx => github.com/pulumi/terraform-provider-signalfx v1.9.2-0.20210330145014-71bc922955b7
)
