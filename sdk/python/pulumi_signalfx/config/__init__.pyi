# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

apiUrl: Optional[str]
"""
API URL for your SignalFx org, may include a realm
"""

authToken: Optional[str]
"""
SignalFx auth token
"""

customAppUrl: Optional[str]
"""
Application URL for your SignalFx org, often customzied for organizations using SSO
"""

timeoutSeconds: Optional[int]
"""
Timeout duration for a single HTTP call in seconds. Defaults to 120
"""
