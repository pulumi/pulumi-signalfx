module github.com/pulumi/pulumi-signalfx/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.11.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.3.3
	github.com/pulumi/pulumi/pkg/v2 v2.2.1 // indirect
	github.com/pulumi/pulumi/sdk/v2 v2.2.1
	github.com/terraform-providers/terraform-provider-signalfx v1.9.2-0.20200507164037-8e3b51013117
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/terraform-providers/terraform-provider-signalfx => github.com/pulumi/terraform-provider-signalfx v1.9.2-0.20200518165554-501754b024c8
)
