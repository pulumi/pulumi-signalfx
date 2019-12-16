// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestAccDashboardGroup(t *testing.T) {
	t.Skip("Skipping as no credit in signalfx account")
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "dashboardGroup"),
		})

	integration.ProgramTest(t, &test)
}

func getApiUrl(t *testing.T) string {
	url := os.Getenv("SFX_API_URL")
	if url == "" {
		t.Skipf("Skipping test due to missing SFX_API_URL environment variable")
	}

	return url
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
	}
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
