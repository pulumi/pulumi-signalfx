// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalfx

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-signalfx/sdk/v7/go/signalfx/internal"
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
	case "signalfx:index/alertMutingRule:AlertMutingRule":
		r = &AlertMutingRule{}
	case "signalfx:index/dashboard:Dashboard":
		r = &Dashboard{}
	case "signalfx:index/dashboardGroup:DashboardGroup":
		r = &DashboardGroup{}
	case "signalfx:index/dataLink:DataLink":
		r = &DataLink{}
	case "signalfx:index/detector:Detector":
		r = &Detector{}
	case "signalfx:index/eventFeedChart:EventFeedChart":
		r = &EventFeedChart{}
	case "signalfx:index/heatmapChart:HeatmapChart":
		r = &HeatmapChart{}
	case "signalfx:index/listChart:ListChart":
		r = &ListChart{}
	case "signalfx:index/metricRuleset:MetricRuleset":
		r = &MetricRuleset{}
	case "signalfx:index/orgToken:OrgToken":
		r = &OrgToken{}
	case "signalfx:index/singleValueChart:SingleValueChart":
		r = &SingleValueChart{}
	case "signalfx:index/slo:Slo":
		r = &Slo{}
	case "signalfx:index/tableChart:TableChart":
		r = &TableChart{}
	case "signalfx:index/team:Team":
		r = &Team{}
	case "signalfx:index/textChart:TextChart":
		r = &TextChart{}
	case "signalfx:index/timeChart:TimeChart":
		r = &TimeChart{}
	case "signalfx:index/webhookIntegration:WebhookIntegration":
		r = &WebhookIntegration{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:signalfx" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/alertMutingRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/dashboard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/dashboardGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/dataLink",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/detector",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/eventFeedChart",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/heatmapChart",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/listChart",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/metricRuleset",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/orgToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/singleValueChart",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/slo",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/tableChart",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/team",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/textChart",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/timeChart",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"signalfx",
		"index/webhookIntegration",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"signalfx",
		&pkg{version},
	)
}
