// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// API URL for your Splunk Observability Cloud org, may include a realm
func GetApiUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "signalfx:apiUrl")
}

// Splunk Observability Cloud auth token
func GetAuthToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "signalfx:authToken")
}

// Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
//
// Deprecated: Remove the definition, the provider will automatically populate the custom app URL as needed
func GetCustomAppUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "signalfx:customAppUrl")
}

// Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
// Password
func GetEmail(ctx *pulumi.Context) string {
	return config.Get(ctx, "signalfx:email")
}

// Allows for users to opt-in to new features that are considered experimental or not ready for general availability yet.
func GetFeaturePreview(ctx *pulumi.Context) string {
	return config.Get(ctx, "signalfx:featurePreview")
}

// Required if the user is configured to be part of multiple organizations
func GetOrganizationId(ctx *pulumi.Context) string {
	return config.Get(ctx, "signalfx:organizationId")
}

// Used to create a session token instead of an API token, it requires the account to be configured to login with Email and
// Password
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "signalfx:password")
}

// Max retries for a single HTTP call. Defaults to 4
func GetRetryMaxAttempts(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "signalfx:retryMaxAttempts")
}

// Maximum retry wait for a single HTTP call in seconds. Defaults to 30
func GetRetryWaitMaxSeconds(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "signalfx:retryWaitMaxSeconds")
}

// Minimum retry wait for a single HTTP call in seconds. Defaults to 1
func GetRetryWaitMinSeconds(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "signalfx:retryWaitMinSeconds")
}

// Allows for Tags to be added by default to resources that allow for tags to be included. If there is already tags
// configured, the global tags are added in prefix.
func GetTags(ctx *pulumi.Context) string {
	return config.Get(ctx, "signalfx:tags")
}

// Allows for teams to be defined at a provider level, and apply to all applicable resources created.
func GetTeams(ctx *pulumi.Context) string {
	return config.Get(ctx, "signalfx:teams")
}

// Timeout duration for a single HTTP call in seconds. Defaults to 120
func GetTimeoutSeconds(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "signalfx:timeoutSeconds")
}
