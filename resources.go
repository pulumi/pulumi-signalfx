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
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pulumi/pulumi-terraform-bridge/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-signalfx/signalfx"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "signalfx"

	// modules:
	mainMod      = "index"
	awsMod       = "Aws"
	azureMod     = "Azure"
	gcpMod       = "Gcp"
	jiraMod      = "Jira"
	opsgenieMod  = "Opsgenie"
	pagerdutyMod = "PagerDuty"
	slackMod     = "Slack"
	victoropsMod = "VictorOps"
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

func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	p := signalfx.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "signalfx",
		Description: "A Pulumi package for creating and managing SignalFx resources.",
		Keywords:    []string{"pulumi", "signalfx"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-signalfx",
		Config: map[string]*tfbridge.SchemaInfo{
			"auth_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{
						"SFX_AUTH_TOKEN",
					},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"signalfx_dashboard":          {Tok: makeResource(mainMod, "Dashboard")},
			"signalfx_dashboard_group":    {Tok: makeResource(mainMod, "DashboardGroup")},
			"signalfx_detector":           {Tok: makeResource(mainMod, "Detector")},
			"signalfx_event_feed_chart":   {Tok: makeResource(mainMod, "EventFeedChart")},
			"signalfx_heatmap_chart":      {Tok: makeResource(mainMod, "HeatmapChart")},
			"signalfx_list_chart":         {Tok: makeResource(mainMod, "ListChart")},
			"signalfx_org_token":          {Tok: makeResource(mainMod, "OrgToken")},
			"signalfx_single_value_chart": {Tok: makeResource(mainMod, "SingleValueChart")},
			"signalfx_team":               {Tok: makeResource(mainMod, "Team")},
			"signalfx_text_chart":         {Tok: makeResource(mainMod, "TextChart")},
			"signalfx_time_chart":         {Tok: makeResource(mainMod, "TimeChart")},
			"signalfx_alert_muting_rule":  {Tok: makeResource(mainMod, "AlertMutingRule")},

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
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "latest",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=1.0.0,<2.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "1.5.0-*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: namespaceMap,
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if resourceSchema := p.ResourcesMap[resname]; resourceSchema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := resourceSchema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
