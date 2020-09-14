# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['TokenIntegration']


class TokenIntegration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        SignalFx AWS CloudWatch integrations using security tokens. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).

        > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.

        > **WARNING** This resource implements a part of a workflow. You must use it with `aws.Integration`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_signalfx as signalfx

        aws_myteam_token = signalfx.aws.TokenIntegration("awsMyteamToken")
        # Make yourself an AWS IAM role here
        aws_sfx_role = aws.iam.Role("awsSfxRole")
        # Stuff here that uses the external and account ID
        aws_myteam = signalfx.aws.Integration("awsMyteam",
            enabled=True,
            integration_id=aws_myteam_token.id,
            token="put_your_token_here",
            key="put_your_key_here",
            regions=["us-east-1"],
            poll_rate=300,
            import_cloud_watch=True,
            enable_aws_usage=True,
            custom_namespace_sync_rules=[signalfx.aws.IntegrationCustomNamespaceSyncRuleArgs(
                default_action="Exclude",
                filter_action="Include",
                filter_source="filter('code', '200')",
                namespace="fart",
            )],
            namespace_sync_rules=[signalfx.aws.IntegrationNamespaceSyncRuleArgs(
                default_action="Exclude",
                filter_action="Include",
                filter_source="filter('code', '200')",
                namespace="AWS/EC2",
            )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of this integration
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['name'] = name
            __props__['signalfx_aws_account'] = None
            __props__['token_id'] = None
        super(TokenIntegration, __self__).__init__(
            'signalfx:aws/tokenIntegration:TokenIntegration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            signalfx_aws_account: Optional[pulumi.Input[str]] = None,
            token_id: Optional[pulumi.Input[str]] = None) -> 'TokenIntegration':
        """
        Get an existing TokenIntegration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of this integration
        :param pulumi.Input[str] signalfx_aws_account: The AWS Account ARN to use with your policies/roles, provided by SignalFx.
        :param pulumi.Input[str] token_id: The SignalFx-generated AWS token to use with an AWS integration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["signalfx_aws_account"] = signalfx_aws_account
        __props__["token_id"] = token_id
        return TokenIntegration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this integration
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="signalfxAwsAccount")
    def signalfx_aws_account(self) -> pulumi.Output[str]:
        """
        The AWS Account ARN to use with your policies/roles, provided by SignalFx.
        """
        return pulumi.get(self, "signalfx_aws_account")

    @property
    @pulumi.getter(name="tokenId")
    def token_id(self) -> pulumi.Output[str]:
        """
        The SignalFx-generated AWS token to use with an AWS integration.
        """
        return pulumi.get(self, "token_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

