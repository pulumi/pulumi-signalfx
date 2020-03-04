module github.com/pulumi/pulumi-signalfx

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.9.1
	github.com/pulumi/pulumi-terraform-bridge v1.6.4
	github.com/terraform-providers/terraform-provider-signalfx v1.9.2-0.20200303222832-3119e91a5162
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
