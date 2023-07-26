// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package signalfx

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/pulumi/pulumi-signalfx/provider/v6/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/splunk-terraform/terraform-provider-signalfx/signalfx"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "signalfx"

	// modules:
	mainMod       = "index"
	awsMod        = "Aws"
	azureMod      = "Azure"
	gcpMod        = "Gcp"
	jiraMod       = "Jira"
	logsMod       = "Logs"
	opsgenieMod   = "Opsgenie"
	pagerdutyMod  = "PagerDuty"
	slackMod      = "Slack"
	victoropsMod  = "VictorOps"
	serviceNowMod = "ServiceNow"
)

var namespaceMap = map[string]string{
	mainPkg: "SignalFx",
}

func makeMember(moduleTitle string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	fn := string(unicode.ToLower(rune(mem[0]))) + mem[1:]
	token := moduleName + "/" + fn
	return tokens.ModuleMember(mainPkg + ":" + token + ":" + mem)
}

func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

func makeResource(mod string, res string) tokens.Type {
	return makeType(mod, res)
}

func makeDataSource(mod string, res string) tokens.ModuleMember {
	return makeMember(mod, res)
}

func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(signalfx.Provider())
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "signalfx",
		Description:          "A Pulumi package for creating and managing SignalFx resources.",
		Keywords:             []string{"pulumi", "signalfx"},
		License:              "Apache-2.0",
		Homepage:             "https://pulumi.io",
		Repository:           "https://github.com/pulumi/pulumi-signalfx",
		GitHubOrg:            "splunk-terraform",
		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"signalfx_dashboard":           {Tok: makeResource(mainMod, "Dashboard")},
			"signalfx_dashboard_group":     {Tok: makeResource(mainMod, "DashboardGroup")},
			"signalfx_detector":            {Tok: makeResource(mainMod, "Detector")},
			"signalfx_event_feed_chart":    {Tok: makeResource(mainMod, "EventFeedChart")},
			"signalfx_heatmap_chart":       {Tok: makeResource(mainMod, "HeatmapChart")},
			"signalfx_list_chart":          {Tok: makeResource(mainMod, "ListChart")},
			"signalfx_org_token":           {Tok: makeResource(mainMod, "OrgToken")},
			"signalfx_single_value_chart":  {Tok: makeResource(mainMod, "SingleValueChart")},
			"signalfx_table_chart":         {Tok: makeResource(mainMod, "TableChart")},
			"signalfx_team":                {Tok: makeResource(mainMod, "Team")},
			"signalfx_text_chart":          {Tok: makeResource(mainMod, "TextChart")},
			"signalfx_time_chart":          {Tok: makeResource(mainMod, "TimeChart")},
			"signalfx_alert_muting_rule":   {Tok: makeResource(mainMod, "AlertMutingRule")},
			"signalfx_data_link":           {Tok: makeResource(mainMod, "DataLink")},
			"signalfx_webhook_integration": {Tok: makeResource(mainMod, "WebhookIntegration")},
			"signalfx_metric_ruleset":      {Tok: makeResource(mainMod, "MetricRuleset")},

			"signalfx_log_view": {Tok: makeResource(logsMod, "View")},

			"signalfx_aws_external_integration": {Tok: makeResource(awsMod, "ExternalIntegration")},
			"signalfx_aws_integration":          {Tok: makeResource(awsMod, "Integration")},
			"signalfx_aws_token_integration":    {Tok: makeResource(awsMod, "TokenIntegration")},

			"signalfx_azure_integration": {Tok: makeResource(azureMod, "Integration")},

			"signalfx_gcp_integration": {Tok: makeResource(gcpMod, "Integration")},

			"signalfx_opsgenie_integration": {Tok: makeResource(opsgenieMod, "Integration")},

			"signalfx_pagerduty_integration": {Tok: makeResource(pagerdutyMod, "Integration")},

			"signalfx_slack_integration": {Tok: makeResource(slackMod, "Integration")},

			"signalfx_victor_ops_integration": {Tok: makeResource(victoropsMod, "Integration")},

			"signalfx_jira_integration": {Tok: makeResource(jiraMod, "Integration")},

			"signalfx_service_now_integration": {Tok: makeResource(serviceNowMod, "Integration")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"signalfx_dimension_values": {Tok: makeDataSource(mainMod, "getDimensionValues")},

			"signalfx_gcp_services": {Tok: makeDataSource(gcpMod, "getServices")},

			"signalfx_pagerduty_integration": {Tok: makeDataSource(pagerdutyMod, "getIntegration")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: namespaceMap,
		},
	}

	mappedMods := map[string]string{
		"aws":         "Aws",
		"azure":       "Azure",
		"gcp":         "Gcp",
		"jira":        "Jira",
		"log":         "Logs",
		"opsgenie":    "Opsgenie",
		"pagerduty":   "PagerDuty",
		"service_now": "ServiceNow",
		"slack":       "Slack",
		"victor_ops":  "VictorOps",
	}

	mappedModKeys := make([]string, 0, len(mappedMods))
	for k := range mappedMods {
		mappedModKeys = append(mappedModKeys, k)
	}

	moduleNameMap := make(map[string]string, len(mappedMods))
	for _, v := range mappedMods {
		moduleNameMap[strings.ToLower(v)] = v
	}

	err := prov.ComputeTokens(tfbridgetokens.KnownModules("signalfx_", mainMod, mappedModKeys,
		func(mod, name string) (string, error) {
			m, ok := moduleNameMap[strings.ToLower(mod)]
			contract.Assertf(ok, "all mods must be mapped: '%s'", strings.ToLower(mod))
			return makeResource(m, name).String(), nil
		}))
	contract.AssertNoErrorf(err, "failed to compute defaults")

	prov.SetAutonaming(255, "-")

	return prov
}
