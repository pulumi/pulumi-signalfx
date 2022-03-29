# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .alert_muting_rule import *
from .dashboard import *
from .dashboard_group import *
from .data_link import *
from .detector import *
from .event_feed_chart import *
from .get_aws_services import *
from .get_azure_services import *
from .get_dimension_values import *
from .heatmap_chart import *
from .list_chart import *
from .org_token import *
from .provider import *
from .single_value_chart import *
from .team import *
from .text_chart import *
from .time_chart import *
from .webhook_integration import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_signalfx.aws as __aws
    aws = __aws
    import pulumi_signalfx.azure as __azure
    azure = __azure
    import pulumi_signalfx.config as __config
    config = __config
    import pulumi_signalfx.gcp as __gcp
    gcp = __gcp
    import pulumi_signalfx.jira as __jira
    jira = __jira
    import pulumi_signalfx.opsgenie as __opsgenie
    opsgenie = __opsgenie
    import pulumi_signalfx.pagerduty as __pagerduty
    pagerduty = __pagerduty
    import pulumi_signalfx.servicenow as __servicenow
    servicenow = __servicenow
    import pulumi_signalfx.slack as __slack
    slack = __slack
    import pulumi_signalfx.victorops as __victorops
    victorops = __victorops
else:
    aws = _utilities.lazy_import('pulumi_signalfx.aws')
    azure = _utilities.lazy_import('pulumi_signalfx.azure')
    config = _utilities.lazy_import('pulumi_signalfx.config')
    gcp = _utilities.lazy_import('pulumi_signalfx.gcp')
    jira = _utilities.lazy_import('pulumi_signalfx.jira')
    opsgenie = _utilities.lazy_import('pulumi_signalfx.opsgenie')
    pagerduty = _utilities.lazy_import('pulumi_signalfx.pagerduty')
    servicenow = _utilities.lazy_import('pulumi_signalfx.servicenow')
    slack = _utilities.lazy_import('pulumi_signalfx.slack')
    victorops = _utilities.lazy_import('pulumi_signalfx.victorops')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "signalfx",
  "mod": "aws/externalIntegration",
  "fqn": "pulumi_signalfx.aws",
  "classes": {
   "signalfx:aws/externalIntegration:ExternalIntegration": "ExternalIntegration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "aws/integration",
  "fqn": "pulumi_signalfx.aws",
  "classes": {
   "signalfx:aws/integration:Integration": "Integration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "aws/tokenIntegration",
  "fqn": "pulumi_signalfx.aws",
  "classes": {
   "signalfx:aws/tokenIntegration:TokenIntegration": "TokenIntegration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "azure/integration",
  "fqn": "pulumi_signalfx.azure",
  "classes": {
   "signalfx:azure/integration:Integration": "Integration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "gcp/integration",
  "fqn": "pulumi_signalfx.gcp",
  "classes": {
   "signalfx:gcp/integration:Integration": "Integration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/alertMutingRule",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/alertMutingRule:AlertMutingRule": "AlertMutingRule"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/dashboard",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/dashboard:Dashboard": "Dashboard"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/dashboardGroup",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/dashboardGroup:DashboardGroup": "DashboardGroup"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/dataLink",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/dataLink:DataLink": "DataLink"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/detector",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/detector:Detector": "Detector"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/eventFeedChart",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/eventFeedChart:EventFeedChart": "EventFeedChart"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/heatmapChart",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/heatmapChart:HeatmapChart": "HeatmapChart"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/listChart",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/listChart:ListChart": "ListChart"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/orgToken",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/orgToken:OrgToken": "OrgToken"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/singleValueChart",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/singleValueChart:SingleValueChart": "SingleValueChart"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/team",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/team:Team": "Team"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/textChart",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/textChart:TextChart": "TextChart"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/timeChart",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/timeChart:TimeChart": "TimeChart"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "index/webhookIntegration",
  "fqn": "pulumi_signalfx",
  "classes": {
   "signalfx:index/webhookIntegration:WebhookIntegration": "WebhookIntegration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "jira/integration",
  "fqn": "pulumi_signalfx.jira",
  "classes": {
   "signalfx:jira/integration:Integration": "Integration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "opsgenie/integration",
  "fqn": "pulumi_signalfx.opsgenie",
  "classes": {
   "signalfx:opsgenie/integration:Integration": "Integration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "pagerduty/integration",
  "fqn": "pulumi_signalfx.pagerduty",
  "classes": {
   "signalfx:pagerduty/integration:Integration": "Integration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "servicenow/integration",
  "fqn": "pulumi_signalfx.servicenow",
  "classes": {
   "signalfx:servicenow/integration:Integration": "Integration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "slack/integration",
  "fqn": "pulumi_signalfx.slack",
  "classes": {
   "signalfx:slack/integration:Integration": "Integration"
  }
 },
 {
  "pkg": "signalfx",
  "mod": "victorops/integration",
  "fqn": "pulumi_signalfx.victorops",
  "classes": {
   "signalfx:victorops/integration:Integration": "Integration"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "signalfx",
  "token": "pulumi:providers:signalfx",
  "fqn": "pulumi_signalfx",
  "class": "Provider"
 }
]
"""
)
