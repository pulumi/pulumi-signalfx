module github.com/pulumi/pulumi-signalfx/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.0.0
	github.com/pulumi/pulumi/sdk/v2 v2.0.0
	github.com/terraform-providers/terraform-provider-signalfx v1.9.2-0.20200424202336-777b2c494eef
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
