// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-signalfx/sdk/v5/go/signalfx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "signalfx:aws/externalIntegration:ExternalIntegration":
		r = &ExternalIntegration{}
	case "signalfx:aws/integration:Integration":
		r = &Integration{}
	case "signalfx:aws/tokenIntegration:TokenIntegration":
		r = &TokenIntegration{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := signalfx.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"signalfx",
		"aws/externalIntegration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"aws/integration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"aws/tokenIntegration",
		&module{version},
	)
}
