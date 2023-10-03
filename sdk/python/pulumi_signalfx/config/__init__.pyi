# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
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
Application URL for your SignalFx org, often customized for organizations using SSO
"""

retryMaxAttempts: Optional[int]
"""
Max retries for a single HTTP call. Defaults to 4
"""

retryWaitMaxSeconds: Optional[int]
"""
Maximum retry wait for a single HTTP call in seconds. Defaults to 30
"""

retryWaitMinSeconds: Optional[int]
"""
Minimum retry wait for a single HTTP call in seconds. Defaults to 1
"""

timeoutSeconds: Optional[int]
"""
Timeout duration for a single HTTP call in seconds. Defaults to 120
"""

