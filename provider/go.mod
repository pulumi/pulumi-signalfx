module github.com/pulumi/pulumi-signalfx/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.0.0-20200414185723-c9aee63e6d4c
	github.com/pulumi/pulumi/sdk/v2 v2.0.0-beta.3
	github.com/terraform-providers/terraform-provider-signalfx v1.9.2-0.20200327151148-93279286044a
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
