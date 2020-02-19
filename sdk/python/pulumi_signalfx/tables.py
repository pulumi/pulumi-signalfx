# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

_SNAKE_TO_CAMEL_CASE_TABLE = {
    "api_key": "apiKey",
    "api_token": "apiToken",
    "api_url": "apiUrl",
    "app_id": "appId",
    "assignee_display_name": "assigneeDisplayName",
    "assignee_name": "assigneeName",
    "auth_method": "authMethod",
    "auth_token": "authToken",
    "authorized_writer_teams": "authorizedWriterTeams",
    "authorized_writer_users": "authorizedWriterUsers",
    "axes_include_zero": "axesIncludeZero",
    "axes_precision": "axesPrecision",
    "axis_left": "axisLeft",
    "axis_right": "axisRight",
    "base_url": "baseUrl",
    "charts_resolution": "chartsResolution",
    "color_by": "colorBy",
    "color_range": "colorRange",
    "color_scales": "colorScales",
    "context_dashboard_id": "contextDashboardId",
    "custom_app_url": "customAppUrl",
    "custom_cloudwatch_namespaces": "customCloudwatchNamespaces",
    "custom_namespace_sync_rules": "customNamespaceSyncRules",
    "dashboard_group": "dashboardGroup",
    "disable_sampling": "disableSampling",
    "dpm_limits": "dpmLimits",
    "effective_start_time": "effectiveStartTime",
    "enable_aws_usage": "enableAwsUsage",
    "end_time": "endTime",
    "event_options": "eventOptions",
    "event_overlays": "eventOverlays",
    "external_id": "externalId",
    "group_bies": "groupBies",
    "hide_timestamp": "hideTimestamp",
    "histogram_options": "histogramOptions",
    "host_or_usage_limits": "hostOrUsageLimits",
    "import_cloud_watch": "importCloudWatch",
    "integration_id": "integrationId",
    "is_timestamp_hidden": "isTimestampHidden",
    "issue_type": "issueType",
    "legend_fields_to_hides": "legendFieldsToHides",
    "legend_options_fields": "legendOptionsFields",
    "max_delay": "maxDelay",
    "max_precision": "maxPrecision",
    "minimum_resolution": "minimumResolution",
    "namespace_sync_rules": "namespaceSyncRules",
    "notifications_criticals": "notificationsCriticals",
    "notifications_defaults": "notificationsDefaults",
    "notifications_infos": "notificationsInfos",
    "notifications_majors": "notificationsMajors",
    "notifications_minors": "notificationsMinors",
    "notifications_warnings": "notificationsWarnings",
    "on_chart_legend_dimension": "onChartLegendDimension",
    "plot_type": "plotType",
    "poll_rate": "pollRate",
    "post_url": "postUrl",
    "program_text": "programText",
    "project_key": "projectKey",
    "project_service_keys": "projectServiceKeys",
    "property_name": "propertyName",
    "property_value": "propertyValue",
    "refresh_interval": "refreshInterval",
    "role_arn": "roleArn",
    "secondary_visualization": "secondaryVisualization",
    "secret_key": "secretKey",
    "selected_event_overlays": "selectedEventOverlays",
    "show_data_markers": "showDataMarkers",
    "show_event_lines": "showEventLines",
    "show_spark_line": "showSparkLine",
    "signalfx_aws_account": "signalfxAwsAccount",
    "sort_by": "sortBy",
    "start_time": "startTime",
    "stop_time": "stopTime",
    "target_external_urls": "targetExternalUrls",
    "target_signalfx_dashboards": "targetSignalfxDashboards",
    "target_splunks": "targetSplunks",
    "tenant_id": "tenantId",
    "time_range": "timeRange",
    "token_id": "tokenId",
    "unit_prefix": "unitPrefix",
    "use_get_metric_data_method": "useGetMetricDataMethod",
    "user_email": "userEmail",
    "viz_options": "vizOptions",
    "webhook_url": "webhookUrl",
}

_CAMEL_TO_SNAKE_CASE_TABLE = {
    "apiKey": "api_key",
    "apiToken": "api_token",
    "apiUrl": "api_url",
    "appId": "app_id",
    "assigneeDisplayName": "assignee_display_name",
    "assigneeName": "assignee_name",
    "authMethod": "auth_method",
    "authToken": "auth_token",
    "authorizedWriterTeams": "authorized_writer_teams",
    "authorizedWriterUsers": "authorized_writer_users",
    "axesIncludeZero": "axes_include_zero",
    "axesPrecision": "axes_precision",
    "axisLeft": "axis_left",
    "axisRight": "axis_right",
    "baseUrl": "base_url",
    "chartsResolution": "charts_resolution",
    "colorBy": "color_by",
    "colorRange": "color_range",
    "colorScales": "color_scales",
    "contextDashboardId": "context_dashboard_id",
    "customAppUrl": "custom_app_url",
    "customCloudwatchNamespaces": "custom_cloudwatch_namespaces",
    "customNamespaceSyncRules": "custom_namespace_sync_rules",
    "dashboardGroup": "dashboard_group",
    "disableSampling": "disable_sampling",
    "dpmLimits": "dpm_limits",
    "effectiveStartTime": "effective_start_time",
    "enableAwsUsage": "enable_aws_usage",
    "endTime": "end_time",
    "eventOptions": "event_options",
    "eventOverlays": "event_overlays",
    "externalId": "external_id",
    "groupBies": "group_bies",
    "hideTimestamp": "hide_timestamp",
    "histogramOptions": "histogram_options",
    "hostOrUsageLimits": "host_or_usage_limits",
    "importCloudWatch": "import_cloud_watch",
    "integrationId": "integration_id",
    "isTimestampHidden": "is_timestamp_hidden",
    "issueType": "issue_type",
    "legendFieldsToHides": "legend_fields_to_hides",
    "legendOptionsFields": "legend_options_fields",
    "maxDelay": "max_delay",
    "maxPrecision": "max_precision",
    "minimumResolution": "minimum_resolution",
    "namespaceSyncRules": "namespace_sync_rules",
    "notificationsCriticals": "notifications_criticals",
    "notificationsDefaults": "notifications_defaults",
    "notificationsInfos": "notifications_infos",
    "notificationsMajors": "notifications_majors",
    "notificationsMinors": "notifications_minors",
    "notificationsWarnings": "notifications_warnings",
    "onChartLegendDimension": "on_chart_legend_dimension",
    "plotType": "plot_type",
    "pollRate": "poll_rate",
    "postUrl": "post_url",
    "programText": "program_text",
    "projectKey": "project_key",
    "projectServiceKeys": "project_service_keys",
    "propertyName": "property_name",
    "propertyValue": "property_value",
    "refreshInterval": "refresh_interval",
    "roleArn": "role_arn",
    "secondaryVisualization": "secondary_visualization",
    "secretKey": "secret_key",
    "selectedEventOverlays": "selected_event_overlays",
    "showDataMarkers": "show_data_markers",
    "showEventLines": "show_event_lines",
    "showSparkLine": "show_spark_line",
    "signalfxAwsAccount": "signalfx_aws_account",
    "sortBy": "sort_by",
    "startTime": "start_time",
    "stopTime": "stop_time",
    "targetExternalUrls": "target_external_urls",
    "targetSignalfxDashboards": "target_signalfx_dashboards",
    "targetSplunks": "target_splunks",
    "tenantId": "tenant_id",
    "timeRange": "time_range",
    "tokenId": "token_id",
    "unitPrefix": "unit_prefix",
    "useGetMetricDataMethod": "use_get_metric_data_method",
    "userEmail": "user_email",
    "vizOptions": "viz_options",
    "webhookUrl": "webhook_url",
}
