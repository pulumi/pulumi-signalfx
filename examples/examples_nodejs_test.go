// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccDashboardGroup(t *testing.T) {
	t.Skip("Skipping as no credit in signalfx account")
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "dashboardGroup"),
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	apiUrl := getApiUrl(t)
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"signalfx:apiUrl": apiUrl,
		},
		Dependencies: []string{
			"@pulumi/signalfx",
		},
	})

	return baseJS
}
