# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetServicesResult:
    """
    A collection of values returned by getServices.
    """
    def __init__(__self__, id=None, services=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        __self__.services = services
class AwaitableGetServicesResult(GetServicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServicesResult(
            id=self.id,
            services=self.services)

def get_services(services=None,opts=None):
    """
    Use this data source to get a list of AWS service names.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_signalfx as signalfx

    aws_services = signalfx.aws.get_services()
    # Leaves out most of the integration bits, see the docs
    # for aws.Integration for more
    aws_myteam = signalfx.aws.Integration("awsMyteam", services=[__item["name"] for __item in [aws_services.services]])
    ```




    The **services** object supports the following:

      * `name` (`str`)
    """
    __args__ = dict()


    __args__['services'] = services
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('signalfx:aws/getServices:getServices', __args__, opts=opts).value

    return AwaitableGetServicesResult(
        id=__ret__.get('id'),
        services=__ret__.get('services'))