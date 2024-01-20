# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('signalfx')


class _ExportableConfig(types.ModuleType):
    @property
    def api_url(self) -> Optional[str]:
        """
        API URL for your Splunk Observability Cloud org, may include a realm
        """
        return __config__.get('apiUrl')

    @property
    def auth_token(self) -> Optional[str]:
        """
        Splunk Observability Cloud auth token
        """
        return __config__.get('authToken')

    @property
    def custom_app_url(self) -> Optional[str]:
        """
        Application URL for your Splunk Observability Cloud org, often customized for organizations using SSO
        """
        return __config__.get('customAppUrl')

    @property
    def retry_max_attempts(self) -> Optional[int]:
        """
        Max retries for a single HTTP call. Defaults to 4
        """
        return __config__.get_int('retryMaxAttempts')

    @property
    def retry_wait_max_seconds(self) -> Optional[int]:
        """
        Maximum retry wait for a single HTTP call in seconds. Defaults to 30
        """
        return __config__.get_int('retryWaitMaxSeconds')

    @property
    def retry_wait_min_seconds(self) -> Optional[int]:
        """
        Minimum retry wait for a single HTTP call in seconds. Defaults to 1
        """
        return __config__.get_int('retryWaitMinSeconds')

    @property
    def timeout_seconds(self) -> Optional[int]:
        """
        Timeout duration for a single HTTP call in seconds. Defaults to 120
        """
        return __config__.get_int('timeoutSeconds')

