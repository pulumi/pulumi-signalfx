# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class TextChart(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    Description of the text note.
    """
    markdown: pulumi.Output[str]
    """
    Markdown text to display.
    """
    name: pulumi.Output[str]
    """
    Name of the text note.
    """
    url: pulumi.Output[str]
    """
    URL of the chart
    """
    def __init__(__self__, resource_name, opts=None, description=None, markdown=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        This special type of chart doesn’t display any metric data. Rather, it lets you place a text note on the dashboard.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/text_chart.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the text note.
        :param pulumi.Input[str] markdown: Markdown text to display.
        :param pulumi.Input[str] name: Name of the text note.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            if markdown is None:
                raise TypeError("Missing required property 'markdown'")
            __props__['markdown'] = markdown
            __props__['name'] = name
            __props__['url'] = None
        super(TextChart, __self__).__init__(
            'signalfx:index/textChart:TextChart',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, markdown=None, name=None, url=None):
        """
        Get an existing TextChart resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the text note.
        :param pulumi.Input[str] markdown: Markdown text to display.
        :param pulumi.Input[str] name: Name of the text note.
        :param pulumi.Input[str] url: URL of the chart
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["markdown"] = markdown
        __props__["name"] = name
        __props__["url"] = url
        return TextChart(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

