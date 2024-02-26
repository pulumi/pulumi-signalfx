// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface AlertMutingRuleFilter {
    /**
     * (false by default) whether this filter should be a "not" filter
     */
    negated?: boolean;
    /**
     * the property to filter by
     */
    property: string;
    /**
     * the value of the property to filter by
     */
    propertyValue: string;
}

export interface DashboardChart {
    /**
     * ID of the chart to display
     */
    chartId: string;
    /**
     * The column to show the chart in (zero-based); this value always represents the leftmost column of the chart. (between 0 and 11)
     */
    column?: number;
    /**
     * How many rows the chart should take up. (greater than or equal to 1)
     */
    height?: number;
    /**
     * The row to show the chart in (zero-based); if height > 1, this value represents the topmost row of the chart. (greater than or equal to 0)
     */
    row?: number;
    /**
     * How many columns (out of a total of 12, one-based) the chart should take up. (between 1 and 12)
     */
    width?: number;
}

export interface DashboardColumn {
    /**
     * Charts to use for the column
     */
    chartIds: string[];
    /**
     * The column to show the chart in (zero-based); this value always represents the leftmost column of the chart. (between 0 and 11)
     */
    column?: number;
    /**
     * How many rows each chart should take up. (greater than or equal to 1)
     */
    height?: number;
    /**
     * Number of columns (out of a total of 12) each chart should take up. (between 1 and 12)
     */
    width?: number;
}

export interface DashboardEventOverlay {
    /**
     * Color to use
     */
    color?: string;
    /**
     * The text displaying in the dropdown menu used to select this event overlay as an active overlay for the dashboard.
     */
    label?: string;
    /**
     * (false by default) Whether a vertical line should be displayed in the plot at the time the event occurs
     */
    line?: boolean;
    /**
     * Search term used to define events
     */
    signal: string;
    sources?: outputs.DashboardEventOverlaySource[];
    /**
     * Source for this event's data. Can be "eventTimeSeries" (default) or "detectorEvents".
     */
    type?: string;
}

export interface DashboardEventOverlaySource {
    /**
     * (false by default) Whether this filter should be a "not" filter
     */
    negated?: boolean;
    /**
     * A metric time series dimension or property name
     */
    property: string;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values: string[];
}

export interface DashboardFilter {
    /**
     * If true, this filter will also match data that does not have the specified property
     */
    applyIfExist?: boolean;
    /**
     * (false by default) Whether this filter should be a "not" filter
     */
    negated?: boolean;
    /**
     * A metric time series dimension or property name
     */
    property: string;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values: string[];
}

export interface DashboardGrid {
    /**
     * Charts to use for the grid
     */
    chartIds: string[];
    /**
     * How many rows each chart should take up. (greater than or equal to 1)
     */
    height?: number;
    /**
     * Number of columns (out of a total of 12, one-based) each chart should take up. (between 1 and 12)
     */
    width?: number;
}

export interface DashboardGroupDashboard {
    /**
     * Unique identifier of an association between a dashboard group and a dashboard
     */
    configId: string;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    dashboardId: string;
    /**
     * String that provides a description override for a mirrored dashboard
     */
    descriptionOverride?: string;
    /**
     * Filter to apply to each chart in the dashboard
     */
    filterOverrides?: outputs.DashboardGroupDashboardFilterOverride[];
    /**
     * String that provides a name override for a mirrored dashboard
     */
    nameOverride?: string;
    /**
     * Dashboard variable to apply to each chart in the dashboard
     */
    variableOverrides?: outputs.DashboardGroupDashboardVariableOverride[];
}

export interface DashboardGroupDashboardFilterOverride {
    /**
     * (false by default) Whether this filter should be a "not" filter
     */
    negated?: boolean;
    /**
     * A metric time series dimension or property name
     */
    property: string;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values: string[];
}

export interface DashboardGroupDashboardVariableOverride {
    /**
     * A metric time series dimension or property name
     */
    property: string;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values?: string[];
    /**
     * A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable
     */
    valuesSuggesteds?: string[];
}

export interface DashboardGroupImportQualifier {
    /**
     * Filter to apply to each chart in the dashboard
     */
    filters?: outputs.DashboardGroupImportQualifierFilter[];
    metric?: string;
}

export interface DashboardGroupImportQualifierFilter {
    /**
     * (false by default) Whether this filter should be a "not" filter
     */
    negated?: boolean;
    /**
     * A metric time series dimension or property name
     */
    property: string;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values: string[];
}

export interface DashboardGroupPermission {
    /**
     * Actions level, possible values: READ, WRITE
     */
    actions?: string[];
    /**
     * ID of the principal with access
     */
    principalId: string;
    /**
     * Type of principal, possible values: ORG, TEAM, USER
     */
    principalType: string;
}

export interface DashboardPermissions {
    /**
     * The custom access control list for this dashboard
     */
    acls?: outputs.DashboardPermissionsAcl[];
    /**
     * The ID of the dashboard group that this dashboard inherits permissions from
     */
    parent?: string;
}

export interface DashboardPermissionsAcl {
    /**
     * Actions level, possible values: READ, WRITE
     */
    actions?: string[];
    /**
     * ID of the principal with access
     */
    principalId: string;
    /**
     * Type of principal, possible values: ORG, TEAM, USER
     */
    principalType: string;
}

export interface DashboardSelectedEventOverlay {
    /**
     * Search term used to define events
     */
    signal: string;
    sources?: outputs.DashboardSelectedEventOverlaySource[];
    /**
     * Source for this event's data. Can be "eventTimeSeries" (default) or "detectorEvents".
     */
    type?: string;
}

export interface DashboardSelectedEventOverlaySource {
    /**
     * (false by default) Whether this filter should be a "not" filter
     */
    negated?: boolean;
    /**
     * A metric time series dimension or property name
     */
    property: string;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values: string[];
}

export interface DashboardVariable {
    /**
     * An alias for the dashboard variable. This text will appear as the label for the dropdown field on the dashboard
     */
    alias: string;
    /**
     * If true, this variable will also match data that does not have the specified property
     */
    applyIfExist?: boolean;
    /**
     * Variable description
     */
    description?: string;
    /**
     * A metric time series dimension or property name
     */
    property: string;
    /**
     * If true, this variable will only apply to charts with a filter on the named property.
     */
    replaceOnly?: boolean;
    /**
     * If true, this variable may only be set to the values listed in preferredSuggestions. and only these values will appear in autosuggestion menus. false by default
     */
    restrictedSuggestions?: boolean;
    /**
     * Determines whether a value is required for this variable (and therefore whether it will be possible to view this dashboard without this filter applied). false by default
     */
    valueRequired?: boolean;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values?: string[];
    /**
     * A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable
     */
    valuesSuggesteds?: string[];
}

export interface DataLinkTargetExternalUrl {
    /**
     * The minimum time window for a search sent to an external site. Depends on the value set for `timeFormat`.
     */
    minimumTimeWindow?: string;
    /**
     * User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     */
    name: string;
    /**
     * Describes the relationship between Splunk Observability Cloud metadata keys and external system properties when the key names are different
     */
    propertyKeyMapping?: {[key: string]: string};
    /**
     * Designates the format of minimumTimeWindow in the same data link target object.
     */
    timeFormat?: string;
    /**
     * URL string for a Splunk instance or external system data link target.
     */
    url: string;
}

export interface DataLinkTargetSignalfxDashboard {
    /**
     * SignalFx-assigned ID of the dashboard link target's dashboard group
     */
    dashboardGroupId: string;
    /**
     * SignalFx-assigned ID of the dashboard link target
     */
    dashboardId: string;
    /**
     * Flag that designates a target as the default for a data link object.
     */
    isDefault?: boolean;
    /**
     * User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     */
    name: string;
}

export interface DataLinkTargetSplunk {
    /**
     * User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     */
    name: string;
    /**
     * Describes the relationship between Splunk Observability Cloud metadata keys and external system properties when the key names are different
     */
    propertyKeyMapping?: {[key: string]: string};
}

export interface DetectorRule {
    /**
     * Description of the rule
     */
    description?: string;
    /**
     * A detect label which matches a detect label within the program text
     */
    detectLabel: string;
    /**
     * (default: false) When true, notifications and events will not be generated for the detect label
     */
    disabled?: boolean;
    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     */
    notifications?: string[];
    /**
     * Custom notification message body when an alert is triggered. See https://developers.signalfx.com/v2/reference#detector-model for more info
     */
    parameterizedBody?: string;
    /**
     * Custom notification message subject when an alert is triggered. See https://developers.signalfx.com/v2/reference#detector-model for more info
     */
    parameterizedSubject?: string;
    /**
     * URL of page to consult when an alert is triggered
     */
    runbookUrl?: string;
    /**
     * The severity of the rule, must be one of: Critical, Warning, Major, Minor, Info
     */
    severity: string;
    /**
     * Plain text suggested first course of action, such as a command to execute.
     */
    tip?: string;
}

export interface DetectorVizOption {
    /**
     * Color to use
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    label: string;
    /**
     * An arbitrary prefix to display with the value of this plot
     */
    valuePrefix?: string;
    /**
     * An arbitrary suffix to display with the value of this plot
     */
    valueSuffix?: string;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     */
    valueUnit?: string;
}

export interface HeatmapChartColorRange {
    /**
     * The color range to use. The starting hex color value for data values in a heatmap chart. Specify the value as a 6-character hexadecimal value preceded by the '#' character, for example "#ea1849" (grass green).
     */
    color: string;
    /**
     * The maximum value within the coloring range
     */
    maxValue?: number;
    /**
     * The minimum value within the coloring range
     */
    minValue?: number;
}

export interface HeatmapChartColorScale {
    /**
     * The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color: string;
    /**
     * Indicates the lower threshold non-inclusive value for this range
     */
    gt?: number;
    /**
     * Indicates the lower threshold inclusive value for this range
     */
    gte?: number;
    /**
     * Indicates the upper threshold non-inculsive value for this range
     */
    lt?: number;
    /**
     * Indicates the upper threshold inclusive value for this range
     */
    lte?: number;
}

export interface ListChartColorScale {
    /**
     * The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color: string;
    /**
     * Indicates the lower threshold non-inclusive value for this range
     */
    gt?: number;
    /**
     * Indicates the lower threshold inclusive value for this range
     */
    gte?: number;
    /**
     * Indicates the upper threshold non-inculsive value for this range
     */
    lt?: number;
    /**
     * Indicates the upper threshold inclusive value for this range
     */
    lte?: number;
}

export interface ListChartLegendOptionsField {
    /**
     * (true by default) Determines if this property is displayed in the data table.
     */
    enabled?: boolean;
    /**
     * The name of a property to hide or show in the data table.
     */
    property: string;
}

export interface ListChartVizOption {
    /**
     * Color to use
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    label: string;
    /**
     * An arbitrary prefix to display with the value of this plot
     */
    valuePrefix?: string;
    /**
     * An arbitrary suffix to display with the value of this plot
     */
    valueSuffix?: string;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     */
    valueUnit?: string;
}

export interface MetricRulesetAggregationRule {
    /**
     * The aggregator for this rule
     */
    aggregators: outputs.MetricRulesetAggregationRuleAggregator[];
    /**
     * Status of this aggregation rule
     */
    enabled: boolean;
    /**
     * The matcher for this rule
     */
    matchers: outputs.MetricRulesetAggregationRuleMatcher[];
    /**
     * Name of this aggregation rule
     */
    name?: string;
}

export interface MetricRulesetAggregationRuleAggregator {
    /**
     * List of dimensions to keep or drop in aggregated metric
     */
    dimensions: string[];
    /**
     * Flag specifying to keep or drop given dimensions
     */
    dropDimensions: boolean;
    /**
     * The aggregated metric name
     */
    outputName: string;
    /**
     * The type of the aggregator
     */
    type: string;
}

export interface MetricRulesetAggregationRuleMatcher {
    /**
     * List of filters to match on
     */
    filters?: outputs.MetricRulesetAggregationRuleMatcherFilter[];
    /**
     * The type of the matcher
     */
    type: string;
}

export interface MetricRulesetAggregationRuleMatcherFilter {
    /**
     * Flag specifying equals or not equals
     */
    not: boolean;
    /**
     * Name of dimension to match
     */
    property: string;
    /**
     * List of property values to match
     */
    propertyValues: string[];
}

export interface MetricRulesetRoutingRule {
    /**
     * Destination to send the input metric
     */
    destination: string;
}

export interface OrgTokenDpmLimits {
    /**
     * The datapoints per minute (dpm) limit for this token. If you exceed this limit, Splunk Observability Cloud sends out an alert.
     */
    dpmLimit: number;
    /**
     * DPM level at which Splunk Observability Cloud sends the notification for this token. If you don't specify a notification, Splunk Observability Cloud sends the generic notification.
     */
    dpmNotificationThreshold?: number;
}

export interface OrgTokenHostOrUsageLimits {
    /**
     * Max number of containers that can use this token
     */
    containerLimit?: number;
    /**
     * Notification threshold for containers
     */
    containerNotificationThreshold?: number;
    /**
     * Max number of custom metrics that can be sent with this token
     */
    customMetricsLimit?: number;
    /**
     * Notification threshold for custom metrics
     */
    customMetricsNotificationThreshold?: number;
    /**
     * Max number of high-res metrics that can be sent with this token
     */
    highResMetricsLimit?: number;
    /**
     * Notification threshold for high-res metrics
     */
    highResMetricsNotificationThreshold?: number;
    /**
     * Max number of hosts that can use this token
     */
    hostLimit?: number;
    /**
     * Notification threshold for hosts
     */
    hostNotificationThreshold?: number;
}

export interface SingleValueChartColorScale {
    /**
     * The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color: string;
    /**
     * Indicates the lower threshold non-inclusive value for this range
     */
    gt?: number;
    /**
     * Indicates the lower threshold inclusive value for this range
     */
    gte?: number;
    /**
     * Indicates the upper threshold non-inculsive value for this range
     */
    lt?: number;
    /**
     * Indicates the upper threshold inclusive value for this range
     */
    lte?: number;
}

export interface SingleValueChartVizOption {
    /**
     * Color to use
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    label: string;
    /**
     * An arbitrary prefix to display with the value of this plot
     */
    valuePrefix?: string;
    /**
     * An arbitrary suffix to display with the value of this plot
     */
    valueSuffix?: string;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     */
    valueUnit?: string;
}

export interface SloInput {
    /**
     * Label used in `programText` that refers to the data block which contains the stream of successful events
     */
    goodEventsLabel?: string;
    /**
     * Signalflow program text for the SLO. More info at "https://dev.splunk.com/observability/docs/signalflow". We require this Signalflow program text to contain at least 2 data blocks - one for the total stream and one for the good stream, whose labels are specified by goodEventsLabel and totalEventsLabel
     */
    programText: string;
    /**
     * Label used in `programText` that refers to the data block which contains the stream of total events
     */
    totalEventsLabel?: string;
}

export interface SloTarget {
    /**
     * SLO alert rules
     */
    alertRules: outputs.SloTargetAlertRule[];
    /**
     * (Required for `RollingWindow` type) Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
     */
    compliancePeriod?: string;
    /**
     * Target value in the form of a percentage
     */
    slo: number;
    /**
     * SLO target type can be the following type: `RollingWindow`
     */
    type: string;
}

export interface SloTargetAlertRule {
    /**
     * Set of rules used for alerting
     */
    rules: outputs.SloTargetAlertRuleRule[];
    /**
     * SLO alert rule type
     */
    type: string;
}

export interface SloTargetAlertRuleRule {
    /**
     * Description of the rule
     */
    description?: string;
    /**
     * (default: false) When true, notifications and events will not be generated for the detect label
     */
    disabled?: boolean;
    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     */
    notifications?: string[];
    /**
     * Custom notification message body when an alert is triggered. See https://developers.signalfx.com/v2/reference#detector-model for more info
     */
    parameterizedBody?: string;
    /**
     * Custom notification message subject when an alert is triggered. See https://developers.signalfx.com/v2/reference#detector-model for more info
     */
    parameterizedSubject?: string;
    /**
     * Parameters for the SLO alert rule. Each SLO alert rule type accepts different parameters. If not specified, default parameters are used.
     */
    parameters?: outputs.SloTargetAlertRuleRuleParameters;
    /**
     * URL of page to consult when an alert is triggered
     */
    runbookUrl?: string;
    /**
     * The severity of the rule, must be one of: Critical, Warning, Major, Minor, Info
     */
    severity: string;
    /**
     * Plain text suggested first course of action, such as a command to execute.
     */
    tip?: string;
}

export interface SloTargetAlertRuleRuleParameters {
    /**
     * Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burnRateThreshold1 parameter.
     */
    burnRateThreshold1: number;
    /**
     * Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burnRateThreshold2 parameter.
     */
    burnRateThreshold2: number;
    /**
     * Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the fireLasting parameter
     */
    fireLasting: string;
    /**
     * Long window 1 used in burn rate alert calculation. This value must be longer than shortWindow1` and shorter than 90 days. Note: BURN_RATE alert rules use the longWindow1 parameter.
     */
    longWindow1: string;
    /**
     * Long window 2 used in burn rate alert calculation. This value must be longer than shortWindow2` and shorter than 90 days. Note: BURN_RATE alert rules use the longWindow2 parameter.
     */
    longWindow2: string;
    /**
     * Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: ERROR_BUDGET_LEFT alert rules use the percentErrorBudgetLeft parameter.
     */
    percentErrorBudgetLeft: number;
    /**
     * Percentage of the fireLasting duration that the alert condition is met before the alert is triggered. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the percentOfLasting parameter
     */
    percentOfLasting: number;
    /**
     * Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_1. Note: BURN_RATE alert rules use the shortWindow1 parameter.
     */
    shortWindow1: string;
    /**
     * Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_2. Note: BURN_RATE alert rules use the shortWindow2 parameter.
     */
    shortWindow2: string;
}

export interface TableChartVizOption {
    /**
     * Color to use
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    label: string;
    /**
     * An arbitrary prefix to display with the value of this plot
     */
    valuePrefix?: string;
    /**
     * An arbitrary suffix to display with the value of this plot
     */
    valueSuffix?: string;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     */
    valueUnit?: string;
}

export interface TimeChartAxisLeft {
    /**
     * A line to draw as a high watermark
     */
    highWatermark?: number;
    /**
     * A label to attach to the high watermark line
     */
    highWatermarkLabel?: string;
    /**
     * Label of the left axis
     */
    label?: string;
    /**
     * A line to draw as a low watermark
     */
    lowWatermark?: number;
    /**
     * A label to attach to the low watermark line
     */
    lowWatermarkLabel?: string;
    /**
     * The maximum value for the left axis
     */
    maxValue?: number;
    /**
     * The minimum value for the left axis
     */
    minValue?: number;
    watermarks?: outputs.TimeChartAxisLeftWatermark[];
}

export interface TimeChartAxisLeftWatermark {
    /**
     * Label to display associated with the watermark line
     */
    label?: string;
    /**
     * Axis value where the watermark line will be displayed
     */
    value: number;
}

export interface TimeChartAxisRight {
    /**
     * A line to draw as a high watermark
     */
    highWatermark?: number;
    /**
     * A label to attach to the high watermark line
     */
    highWatermarkLabel?: string;
    /**
     * Label of the right axis
     */
    label?: string;
    /**
     * A line to draw as a low watermark
     */
    lowWatermark?: number;
    /**
     * A label to attach to the low watermark line
     */
    lowWatermarkLabel?: string;
    /**
     * The maximum value for the right axis
     */
    maxValue?: number;
    /**
     * The minimum value for the right axis
     */
    minValue?: number;
    watermarks?: outputs.TimeChartAxisRightWatermark[];
}

export interface TimeChartAxisRightWatermark {
    /**
     * Label to display associated with the watermark line
     */
    label?: string;
    /**
     * Axis value where the watermark line will be displayed
     */
    value: number;
}

export interface TimeChartEventOption {
    /**
     * Color to use
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * The label used in the publish statement that displays the events you want to customize
     */
    label: string;
}

export interface TimeChartHistogramOption {
    /**
     * Base color theme to use for the graph.
     */
    colorTheme?: string;
}

export interface TimeChartLegendOptionsField {
    /**
     * (true by default) Determines if this property is displayed in the data table.
     */
    enabled?: boolean;
    /**
     * The name of a property to hide or show in the data table.
     */
    property: string;
}

export interface TimeChartVizOption {
    /**
     * The Y-axis associated with values for this plot. Must be either "right" or "left". Defaults to "left".
     */
    axis?: string;
    /**
     * Color to use
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    label: string;
    /**
     * (Chart plotType by default) The visualization style to use. Must be "LineChart", "AreaChart", "ColumnChart", or "Histogram"
     */
    plotType?: string;
    /**
     * An arbitrary prefix to display with the value of this plot
     */
    valuePrefix?: string;
    /**
     * An arbitrary suffix to display with the value of this plot
     */
    valueSuffix?: string;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     */
    valueUnit?: string;
}

export interface WebhookIntegrationHeader {
    headerKey: string;
    headerValue: string;
}

export namespace aws {
    export interface IntegrationCustomNamespaceSyncRule {
        /**
         * Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filterAction` and `filterSource` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of "Include" or "Exclude".
         */
        defaultAction?: string;
        /**
         * Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of "Include" or "Exclude".
         */
        filterAction?: string;
        /**
         * Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
         */
        filterSource?: string;
        /**
         * An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
         */
        namespace: string;
    }

    export interface IntegrationMetricStatsToSync {
        /**
         * AWS metric that you want to pick statistics for
         */
        metric: string;
        /**
         * An AWS namespace having AWS metric that you want to pick statistics for
         */
        namespace: string;
        /**
         * AWS statistics you want to collect
         */
        stats: string[];
    }

    export interface IntegrationNamespaceSyncRule {
        /**
         * Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filterAction` and `filterSource` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of "Include" or "Exclude".
         */
        defaultAction?: string;
        /**
         * Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of "Include" or "Exclude".
         */
        filterAction?: string;
        /**
         * Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
         */
        filterSource?: string;
        /**
         * An AWS namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
         */
        namespace: string;
    }

}

export namespace azure {
    export interface IntegrationCustomNamespacesPerService {
        /**
         * The namespaces to sync
         */
        namespaces: string[];
        /**
         * The name of the service
         */
        service: string;
    }

    export interface IntegrationResourceFilterRule {
        filterSource: string;
    }

}

export namespace gcp {
    export interface IntegrationProjectServiceKey {
        projectId: string;
        projectKey: string;
    }

}

export namespace log {
    export interface ViewColumn {
        /**
         * Name of the column
         */
        name: string;
    }

    export interface ViewSortOption {
        /**
         * Name of the column
         */
        descending: boolean;
        /**
         * Name of the column
         */
        field: string;
    }

}
