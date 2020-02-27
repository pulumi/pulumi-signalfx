# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class DataLink(pulumi.CustomResource):
    context_dashboard_id: pulumi.Output[str]
    """
    If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
    """
    property_name: pulumi.Output[str]
    """
    Name (key) of the metadata that's the trigger of a data link. If you specify `property_value`, you must specify `property_name`.
    """
    property_value: pulumi.Output[str]
    """
    Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `property_name`.
    """
    target_external_urls: pulumi.Output[list]
    """
    Link to an external URL
    
      * `isDefault` (`bool`) - Flag that designates a target as the default for a data link object. `true` by default
      * `minimumTimeWindow` (`str`) - The [minimum time window](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) for a search sent to an external site. Defaults to `6000`
      * `name` (`str`) - User-assigned target name. Use this value to differentiate between the link targets for a data link object.
      * `propertyKeyMapping` (`dict`) - Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.
      * `timeFormat` (`str`) - [Designates the format](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) of `minimum_time_window` in the same data link target object. Must be one of `"ISO8601"`, `"EpochSeconds"` or `"Epoch"` (which is milliseconds). Defaults to `"ISO8601"`.
      * `url` (`str`)
    """
    target_signalfx_dashboards: pulumi.Output[list]
    """
    Link to a SignalFx dashboard
    
      * `dashboardGroupId` (`str`) - SignalFx-assigned ID of the dashboard link target's dashboard group
      * `dashboardId` (`str`) - SignalFx-assigned ID of the dashboard link target
      * `isDefault` (`bool`) - Flag that designates a target as the default for a data link object. `true` by default
      * `name` (`str`) - User-assigned target name. Use this value to differentiate between the link targets for a data link object.
    """
    target_splunks: pulumi.Output[list]
    """
    Link to an external URL
    
      * `isDefault` (`bool`) - Flag that designates a target as the default for a data link object. `true` by default
      * `name` (`str`) - User-assigned target name. Use this value to differentiate between the link targets for a data link object.
      * `propertyKeyMapping` (`dict`) - Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.
    """
    def __init__(__self__, resource_name, opts=None, context_dashboard_id=None, property_name=None, property_value=None, target_external_urls=None, target_signalfx_dashboards=None, target_splunks=None, __props__=None, __name__=None, __opts__=None):
        """
        Manage SignalFx [Data Links](https://docs.signalfx.com/en/latest/managing/data-links.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] context_dashboard_id: If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
        :param pulumi.Input[str] property_name: Name (key) of the metadata that's the trigger of a data link. If you specify `property_value`, you must specify `property_name`.
        :param pulumi.Input[str] property_value: Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `property_name`.
        :param pulumi.Input[list] target_external_urls: Link to an external URL
        :param pulumi.Input[list] target_signalfx_dashboards: Link to a SignalFx dashboard
        :param pulumi.Input[list] target_splunks: Link to an external URL
        
        The **target_external_urls** object supports the following:
        
          * `isDefault` (`pulumi.Input[bool]`) - Flag that designates a target as the default for a data link object. `true` by default
          * `minimumTimeWindow` (`pulumi.Input[str]`) - The [minimum time window](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) for a search sent to an external site. Defaults to `6000`
          * `name` (`pulumi.Input[str]`) - User-assigned target name. Use this value to differentiate between the link targets for a data link object.
          * `propertyKeyMapping` (`pulumi.Input[dict]`) - Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.
          * `timeFormat` (`pulumi.Input[str]`) - [Designates the format](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) of `minimum_time_window` in the same data link target object. Must be one of `"ISO8601"`, `"EpochSeconds"` or `"Epoch"` (which is milliseconds). Defaults to `"ISO8601"`.
          * `url` (`pulumi.Input[str]`)
        
        The **target_signalfx_dashboards** object supports the following:
        
          * `dashboardGroupId` (`pulumi.Input[str]`) - SignalFx-assigned ID of the dashboard link target's dashboard group
          * `dashboardId` (`pulumi.Input[str]`) - SignalFx-assigned ID of the dashboard link target
          * `isDefault` (`pulumi.Input[bool]`) - Flag that designates a target as the default for a data link object. `true` by default
          * `name` (`pulumi.Input[str]`) - User-assigned target name. Use this value to differentiate between the link targets for a data link object.
        
        The **target_splunks** object supports the following:
        
          * `isDefault` (`pulumi.Input[bool]`) - Flag that designates a target as the default for a data link object. `true` by default
          * `name` (`pulumi.Input[str]`) - User-assigned target name. Use this value to differentiate between the link targets for a data link object.
          * `propertyKeyMapping` (`pulumi.Input[dict]`) - Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/data_link.html.markdown.
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

            __props__['context_dashboard_id'] = context_dashboard_id
            __props__['property_name'] = property_name
            __props__['property_value'] = property_value
            __props__['target_external_urls'] = target_external_urls
            __props__['target_signalfx_dashboards'] = target_signalfx_dashboards
            __props__['target_splunks'] = target_splunks
        super(DataLink, __self__).__init__(
            'signalfx:index/dataLink:DataLink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, context_dashboard_id=None, property_name=None, property_value=None, target_external_urls=None, target_signalfx_dashboards=None, target_splunks=None):
        """
        Get an existing DataLink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] context_dashboard_id: If provided, scopes this data link to the supplied dashboard id. If omitted then the link will be global.
        :param pulumi.Input[str] property_name: Name (key) of the metadata that's the trigger of a data link. If you specify `property_value`, you must specify `property_name`.
        :param pulumi.Input[str] property_value: Value of the metadata that's the trigger of a data link. If you specify this property, you must also specify `property_name`.
        :param pulumi.Input[list] target_external_urls: Link to an external URL
        :param pulumi.Input[list] target_signalfx_dashboards: Link to a SignalFx dashboard
        :param pulumi.Input[list] target_splunks: Link to an external URL
        
        The **target_external_urls** object supports the following:
        
          * `isDefault` (`pulumi.Input[bool]`) - Flag that designates a target as the default for a data link object. `true` by default
          * `minimumTimeWindow` (`pulumi.Input[str]`) - The [minimum time window](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) for a search sent to an external site. Defaults to `6000`
          * `name` (`pulumi.Input[str]`) - User-assigned target name. Use this value to differentiate between the link targets for a data link object.
          * `propertyKeyMapping` (`pulumi.Input[dict]`) - Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.
          * `timeFormat` (`pulumi.Input[str]`) - [Designates the format](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) of `minimum_time_window` in the same data link target object. Must be one of `"ISO8601"`, `"EpochSeconds"` or `"Epoch"` (which is milliseconds). Defaults to `"ISO8601"`.
          * `url` (`pulumi.Input[str]`)
        
        The **target_signalfx_dashboards** object supports the following:
        
          * `dashboardGroupId` (`pulumi.Input[str]`) - SignalFx-assigned ID of the dashboard link target's dashboard group
          * `dashboardId` (`pulumi.Input[str]`) - SignalFx-assigned ID of the dashboard link target
          * `isDefault` (`pulumi.Input[bool]`) - Flag that designates a target as the default for a data link object. `true` by default
          * `name` (`pulumi.Input[str]`) - User-assigned target name. Use this value to differentiate between the link targets for a data link object.
        
        The **target_splunks** object supports the following:
        
          * `isDefault` (`pulumi.Input[bool]`) - Flag that designates a target as the default for a data link object. `true` by default
          * `name` (`pulumi.Input[str]`) - User-assigned target name. Use this value to differentiate between the link targets for a data link object.
          * `propertyKeyMapping` (`pulumi.Input[dict]`) - Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/data_link.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["context_dashboard_id"] = context_dashboard_id
        __props__["property_name"] = property_name
        __props__["property_value"] = property_value
        __props__["target_external_urls"] = target_external_urls
        __props__["target_signalfx_dashboards"] = target_signalfx_dashboards
        __props__["target_splunks"] = target_splunks
        return DataLink(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

