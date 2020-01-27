module github.com/pulumi/pulumi-signalfx

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.5.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform-bridge v1.5.2
	github.com/terraform-providers/terraform-provider-signalfx v1.9.2-0.20200127180001-2c64c78ee3ec
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
