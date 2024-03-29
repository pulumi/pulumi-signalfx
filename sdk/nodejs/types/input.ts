// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface AlertMutingRuleFilter {
    /**
     * (false by default) whether this filter should be a "not" filter
     */
    negated?: pulumi.Input<boolean>;
    /**
     * the property to filter by
     */
    property: pulumi.Input<string>;
    /**
     * the value of the property to filter by
     */
    propertyValue: pulumi.Input<string>;
}

export interface DashboardChart {
    /**
     * ID of the chart to display
     */
    chartId: pulumi.Input<string>;
    /**
     * The column to show the chart in (zero-based); this value always represents the leftmost column of the chart. (between 0 and 11)
     */
    column?: pulumi.Input<number>;
    /**
     * How many rows the chart should take up. (greater than or equal to 1)
     */
    height?: pulumi.Input<number>;
    /**
     * The row to show the chart in (zero-based); if height > 1, this value represents the topmost row of the chart. (greater than or equal to 0)
     */
    row?: pulumi.Input<number>;
    /**
     * How many columns (out of a total of 12, one-based) the chart should take up. (between 1 and 12)
     */
    width?: pulumi.Input<number>;
}

export interface DashboardColumn {
    /**
     * Charts to use for the column
     */
    chartIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The column to show the chart in (zero-based); this value always represents the leftmost column of the chart. (between 0 and 11)
     */
    column?: pulumi.Input<number>;
    /**
     * How many rows each chart should take up. (greater than or equal to 1)
     */
    height?: pulumi.Input<number>;
    /**
     * Number of columns (out of a total of 12) each chart should take up. (between 1 and 12)
     */
    width?: pulumi.Input<number>;
}

export interface DashboardEventOverlay {
    /**
     * Color to use
     */
    color?: pulumi.Input<string>;
    /**
     * The text displaying in the dropdown menu used to select this event overlay as an active overlay for the dashboard.
     */
    label?: pulumi.Input<string>;
    /**
     * (false by default) Whether a vertical line should be displayed in the plot at the time the event occurs
     */
    line?: pulumi.Input<boolean>;
    /**
     * Search term used to define events
     */
    signal: pulumi.Input<string>;
    sources?: pulumi.Input<pulumi.Input<inputs.DashboardEventOverlaySource>[]>;
    /**
     * Source for this event's data. Can be "eventTimeSeries" (default) or "detectorEvents".
     */
    type?: pulumi.Input<string>;
}

export interface DashboardEventOverlaySource {
    /**
     * (false by default) Whether this filter should be a "not" filter
     */
    negated?: pulumi.Input<boolean>;
    /**
     * A metric time series dimension or property name
     */
    property: pulumi.Input<string>;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DashboardFilter {
    /**
     * If true, this filter will also match data that does not have the specified property
     */
    applyIfExist?: pulumi.Input<boolean>;
    /**
     * (false by default) Whether this filter should be a "not" filter
     */
    negated?: pulumi.Input<boolean>;
    /**
     * A metric time series dimension or property name
     */
    property: pulumi.Input<string>;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DashboardGrid {
    /**
     * Charts to use for the grid
     */
    chartIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * How many rows each chart should take up. (greater than or equal to 1)
     */
    height?: pulumi.Input<number>;
    /**
     * Number of columns (out of a total of 12, one-based) each chart should take up. (between 1 and 12)
     */
    width?: pulumi.Input<number>;
}

export interface DashboardGroupDashboard {
    /**
     * Unique identifier of an association between a dashboard group and a dashboard
     */
    configId?: pulumi.Input<string>;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    dashboardId: pulumi.Input<string>;
    /**
     * String that provides a description override for a mirrored dashboard
     */
    descriptionOverride?: pulumi.Input<string>;
    /**
     * Filter to apply to each chart in the dashboard
     */
    filterOverrides?: pulumi.Input<pulumi.Input<inputs.DashboardGroupDashboardFilterOverride>[]>;
    /**
     * String that provides a name override for a mirrored dashboard
     */
    nameOverride?: pulumi.Input<string>;
    /**
     * Dashboard variable to apply to each chart in the dashboard
     */
    variableOverrides?: pulumi.Input<pulumi.Input<inputs.DashboardGroupDashboardVariableOverride>[]>;
}

export interface DashboardGroupDashboardFilterOverride {
    /**
     * (false by default) Whether this filter should be a "not" filter
     */
    negated?: pulumi.Input<boolean>;
    /**
     * A metric time series dimension or property name
     */
    property: pulumi.Input<string>;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DashboardGroupDashboardVariableOverride {
    /**
     * A metric time series dimension or property name
     */
    property: pulumi.Input<string>;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable
     */
    valuesSuggesteds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DashboardGroupImportQualifier {
    /**
     * Filter to apply to each chart in the dashboard
     */
    filters?: pulumi.Input<pulumi.Input<inputs.DashboardGroupImportQualifierFilter>[]>;
    metric?: pulumi.Input<string>;
}

export interface DashboardGroupImportQualifierFilter {
    /**
     * (false by default) Whether this filter should be a "not" filter
     */
    negated?: pulumi.Input<boolean>;
    /**
     * A metric time series dimension or property name
     */
    property: pulumi.Input<string>;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DashboardGroupPermission {
    /**
     * Actions level, possible values: READ, WRITE
     */
    actions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the principal with access
     */
    principalId: pulumi.Input<string>;
    /**
     * Type of principal, possible values: ORG, TEAM, USER
     */
    principalType: pulumi.Input<string>;
}

export interface DashboardPermissions {
    /**
     * The custom access control list for this dashboard
     */
    acls?: pulumi.Input<pulumi.Input<inputs.DashboardPermissionsAcl>[]>;
    /**
     * The ID of the dashboard group that this dashboard inherits permissions from
     */
    parent?: pulumi.Input<string>;
}

export interface DashboardPermissionsAcl {
    /**
     * Actions level, possible values: READ, WRITE
     */
    actions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the principal with access
     */
    principalId: pulumi.Input<string>;
    /**
     * Type of principal, possible values: ORG, TEAM, USER
     */
    principalType: pulumi.Input<string>;
}

export interface DashboardSelectedEventOverlay {
    /**
     * Search term used to define events
     */
    signal: pulumi.Input<string>;
    sources?: pulumi.Input<pulumi.Input<inputs.DashboardSelectedEventOverlaySource>[]>;
    /**
     * Source for this event's data. Can be "eventTimeSeries" (default) or "detectorEvents".
     */
    type?: pulumi.Input<string>;
}

export interface DashboardSelectedEventOverlaySource {
    /**
     * (false by default) Whether this filter should be a "not" filter
     */
    negated?: pulumi.Input<boolean>;
    /**
     * A metric time series dimension or property name
     */
    property: pulumi.Input<string>;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DashboardVariable {
    /**
     * An alias for the dashboard variable. This text will appear as the label for the dropdown field on the dashboard
     */
    alias: pulumi.Input<string>;
    /**
     * If true, this variable will also match data that does not have the specified property
     */
    applyIfExist?: pulumi.Input<boolean>;
    /**
     * Variable description
     */
    description?: pulumi.Input<string>;
    /**
     * A metric time series dimension or property name
     */
    property: pulumi.Input<string>;
    /**
     * If true, this variable will only apply to charts with a filter on the named property.
     */
    replaceOnly?: pulumi.Input<boolean>;
    /**
     * If true, this variable may only be set to the values listed in preferredSuggestions. and only these values will appear in autosuggestion menus. false by default
     */
    restrictedSuggestions?: pulumi.Input<boolean>;
    /**
     * Determines whether a value is required for this variable (and therefore whether it will be possible to view this dashboard without this filter applied). false by default
     */
    valueRequired?: pulumi.Input<boolean>;
    /**
     * List of strings (which will be treated as an OR filter on the property)
     */
    values?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable
     */
    valuesSuggesteds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface DataLinkTargetExternalUrl {
    /**
     * The minimum time window for a search sent to an external site. Depends on the value set for `timeFormat`.
     */
    minimumTimeWindow?: pulumi.Input<string>;
    /**
     * User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     */
    name: pulumi.Input<string>;
    /**
     * Describes the relationship between Splunk Observability Cloud metadata keys and external system properties when the key names are different
     */
    propertyKeyMapping?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Designates the format of minimumTimeWindow in the same data link target object.
     */
    timeFormat?: pulumi.Input<string>;
    /**
     * URL string for a Splunk instance or external system data link target.
     */
    url: pulumi.Input<string>;
}

export interface DataLinkTargetSignalfxDashboard {
    /**
     * SignalFx-assigned ID of the dashboard link target's dashboard group
     */
    dashboardGroupId: pulumi.Input<string>;
    /**
     * SignalFx-assigned ID of the dashboard link target
     */
    dashboardId: pulumi.Input<string>;
    /**
     * Flag that designates a target as the default for a data link object.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     */
    name: pulumi.Input<string>;
}

export interface DataLinkTargetSplunk {
    /**
     * User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     */
    name: pulumi.Input<string>;
    /**
     * Describes the relationship between Splunk Observability Cloud metadata keys and external system properties when the key names are different
     */
    propertyKeyMapping?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

export interface DetectorRule {
    /**
     * Description of the rule
     */
    description?: pulumi.Input<string>;
    /**
     * A detect label which matches a detect label within the program text
     */
    detectLabel: pulumi.Input<string>;
    /**
     * (default: false) When true, notifications and events will not be generated for the detect label
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     */
    notifications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom notification message body when an alert is triggered. See https://developers.signalfx.com/v2/reference#detector-model for more info
     */
    parameterizedBody?: pulumi.Input<string>;
    /**
     * Custom notification message subject when an alert is triggered. See https://developers.signalfx.com/v2/reference#detector-model for more info
     */
    parameterizedSubject?: pulumi.Input<string>;
    /**
     * URL of page to consult when an alert is triggered
     */
    runbookUrl?: pulumi.Input<string>;
    /**
     * The severity of the rule, must be one of: Critical, Warning, Major, Minor, Info
     */
    severity: pulumi.Input<string>;
    /**
     * Plain text suggested first course of action, such as a command to execute.
     */
    tip?: pulumi.Input<string>;
}

export interface DetectorVizOption {
    /**
     * Color to use
     */
    color?: pulumi.Input<string>;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    label: pulumi.Input<string>;
    /**
     * An arbitrary prefix to display with the value of this plot
     */
    valuePrefix?: pulumi.Input<string>;
    /**
     * An arbitrary suffix to display with the value of this plot
     */
    valueSuffix?: pulumi.Input<string>;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     */
    valueUnit?: pulumi.Input<string>;
}

export interface HeatmapChartColorRange {
    /**
     * The color range to use. The starting hex color value for data values in a heatmap chart. Specify the value as a 6-character hexadecimal value preceded by the '#' character, for example "#ea1849" (grass green).
     */
    color: pulumi.Input<string>;
    /**
     * The maximum value within the coloring range
     */
    maxValue?: pulumi.Input<number>;
    /**
     * The minimum value within the coloring range
     */
    minValue?: pulumi.Input<number>;
}

export interface HeatmapChartColorScale {
    /**
     * The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color: pulumi.Input<string>;
    /**
     * Indicates the lower threshold non-inclusive value for this range
     */
    gt?: pulumi.Input<number>;
    /**
     * Indicates the lower threshold inclusive value for this range
     */
    gte?: pulumi.Input<number>;
    /**
     * Indicates the upper threshold non-inculsive value for this range
     */
    lt?: pulumi.Input<number>;
    /**
     * Indicates the upper threshold inclusive value for this range
     */
    lte?: pulumi.Input<number>;
}

export interface ListChartColorScale {
    /**
     * The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color: pulumi.Input<string>;
    /**
     * Indicates the lower threshold non-inclusive value for this range
     */
    gt?: pulumi.Input<number>;
    /**
     * Indicates the lower threshold inclusive value for this range
     */
    gte?: pulumi.Input<number>;
    /**
     * Indicates the upper threshold non-inculsive value for this range
     */
    lt?: pulumi.Input<number>;
    /**
     * Indicates the upper threshold inclusive value for this range
     */
    lte?: pulumi.Input<number>;
}

export interface ListChartLegendOptionsField {
    /**
     * (true by default) Determines if this property is displayed in the data table.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of a property to hide or show in the data table.
     */
    property: pulumi.Input<string>;
}

export interface ListChartVizOption {
    /**
     * Color to use
     */
    color?: pulumi.Input<string>;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    label: pulumi.Input<string>;
    /**
     * An arbitrary prefix to display with the value of this plot
     */
    valuePrefix?: pulumi.Input<string>;
    /**
     * An arbitrary suffix to display with the value of this plot
     */
    valueSuffix?: pulumi.Input<string>;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     */
    valueUnit?: pulumi.Input<string>;
}

export interface MetricRulesetAggregationRule {
    /**
     * The aggregator for this rule
     */
    aggregators: pulumi.Input<pulumi.Input<inputs.MetricRulesetAggregationRuleAggregator>[]>;
    /**
     * Status of this aggregation rule
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The matcher for this rule
     */
    matchers: pulumi.Input<pulumi.Input<inputs.MetricRulesetAggregationRuleMatcher>[]>;
    /**
     * Name of this aggregation rule
     */
    name?: pulumi.Input<string>;
}

export interface MetricRulesetAggregationRuleAggregator {
    /**
     * List of dimensions to keep or drop in aggregated metric
     */
    dimensions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Flag specifying to keep or drop given dimensions
     */
    dropDimensions: pulumi.Input<boolean>;
    /**
     * The aggregated metric name
     */
    outputName: pulumi.Input<string>;
    /**
     * The type of the aggregator
     */
    type: pulumi.Input<string>;
}

export interface MetricRulesetAggregationRuleMatcher {
    /**
     * List of filters to match on
     */
    filters?: pulumi.Input<pulumi.Input<inputs.MetricRulesetAggregationRuleMatcherFilter>[]>;
    /**
     * The type of the matcher
     */
    type: pulumi.Input<string>;
}

export interface MetricRulesetAggregationRuleMatcherFilter {
    /**
     * Flag specifying equals or not equals
     */
    not: pulumi.Input<boolean>;
    /**
     * Name of dimension to match
     */
    property: pulumi.Input<string>;
    /**
     * List of property values to match
     */
    propertyValues: pulumi.Input<pulumi.Input<string>[]>;
}

export interface MetricRulesetRoutingRule {
    /**
     * Destination to send the input metric
     */
    destination: pulumi.Input<string>;
}

export interface OrgTokenDpmLimits {
    /**
     * The datapoints per minute (dpm) limit for this token. If you exceed this limit, Splunk Observability Cloud sends out an alert.
     */
    dpmLimit: pulumi.Input<number>;
    /**
     * DPM level at which Splunk Observability Cloud sends the notification for this token. If you don't specify a notification, Splunk Observability Cloud sends the generic notification.
     */
    dpmNotificationThreshold?: pulumi.Input<number>;
}

export interface OrgTokenHostOrUsageLimits {
    /**
     * Max number of containers that can use this token
     */
    containerLimit?: pulumi.Input<number>;
    /**
     * Notification threshold for containers
     */
    containerNotificationThreshold?: pulumi.Input<number>;
    /**
     * Max number of custom metrics that can be sent with this token
     */
    customMetricsLimit?: pulumi.Input<number>;
    /**
     * Notification threshold for custom metrics
     */
    customMetricsNotificationThreshold?: pulumi.Input<number>;
    /**
     * Max number of high-res metrics that can be sent with this token
     */
    highResMetricsLimit?: pulumi.Input<number>;
    /**
     * Notification threshold for high-res metrics
     */
    highResMetricsNotificationThreshold?: pulumi.Input<number>;
    /**
     * Max number of hosts that can use this token
     */
    hostLimit?: pulumi.Input<number>;
    /**
     * Notification threshold for hosts
     */
    hostNotificationThreshold?: pulumi.Input<number>;
}

export interface SingleValueChartColorScale {
    /**
     * The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color: pulumi.Input<string>;
    /**
     * Indicates the lower threshold non-inclusive value for this range
     */
    gt?: pulumi.Input<number>;
    /**
     * Indicates the lower threshold inclusive value for this range
     */
    gte?: pulumi.Input<number>;
    /**
     * Indicates the upper threshold non-inculsive value for this range
     */
    lt?: pulumi.Input<number>;
    /**
     * Indicates the upper threshold inclusive value for this range
     */
    lte?: pulumi.Input<number>;
}

export interface SingleValueChartVizOption {
    /**
     * Color to use
     */
    color?: pulumi.Input<string>;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    label: pulumi.Input<string>;
    /**
     * An arbitrary prefix to display with the value of this plot
     */
    valuePrefix?: pulumi.Input<string>;
    /**
     * An arbitrary suffix to display with the value of this plot
     */
    valueSuffix?: pulumi.Input<string>;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     */
    valueUnit?: pulumi.Input<string>;
}

export interface SloInput {
    /**
     * Label used in `programText` that refers to the data block which contains the stream of successful events
     */
    goodEventsLabel?: pulumi.Input<string>;
    /**
     * Signalflow program text for the SLO. More info at "https://dev.splunk.com/observability/docs/signalflow". We require this Signalflow program text to contain at least 2 data blocks - one for the total stream and one for the good stream, whose labels are specified by goodEventsLabel and totalEventsLabel
     */
    programText: pulumi.Input<string>;
    /**
     * Label used in `programText` that refers to the data block which contains the stream of total events
     */
    totalEventsLabel?: pulumi.Input<string>;
}

export interface SloTarget {
    /**
     * SLO alert rules
     */
    alertRules: pulumi.Input<pulumi.Input<inputs.SloTargetAlertRule>[]>;
    /**
     * (Required for `RollingWindow` type) Compliance period of this SLO. This value must be within the range of 1d (1 days) to 30d (30 days), inclusive.
     */
    compliancePeriod?: pulumi.Input<string>;
    /**
     * Target value in the form of a percentage
     */
    slo: pulumi.Input<number>;
    /**
     * SLO target type can be the following type: `RollingWindow`
     */
    type: pulumi.Input<string>;
}

export interface SloTargetAlertRule {
    /**
     * Set of rules used for alerting
     */
    rules: pulumi.Input<pulumi.Input<inputs.SloTargetAlertRuleRule>[]>;
    /**
     * SLO alert rule type
     */
    type: pulumi.Input<string>;
}

export interface SloTargetAlertRuleRule {
    /**
     * Description of the rule
     */
    description?: pulumi.Input<string>;
    /**
     * (default: false) When true, notifications and events will not be generated for the detect label
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See https://developers.signalfx.com/v2/docs/detector-model#notifications-models for more info
     */
    notifications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom notification message body when an alert is triggered. See https://developers.signalfx.com/v2/reference#detector-model for more info
     */
    parameterizedBody?: pulumi.Input<string>;
    /**
     * Custom notification message subject when an alert is triggered. See https://developers.signalfx.com/v2/reference#detector-model for more info
     */
    parameterizedSubject?: pulumi.Input<string>;
    /**
     * Parameters for the SLO alert rule. Each SLO alert rule type accepts different parameters. If not specified, default parameters are used.
     */
    parameters?: pulumi.Input<inputs.SloTargetAlertRuleRuleParameters>;
    /**
     * URL of page to consult when an alert is triggered
     */
    runbookUrl?: pulumi.Input<string>;
    /**
     * The severity of the rule, must be one of: Critical, Warning, Major, Minor, Info
     */
    severity: pulumi.Input<string>;
    /**
     * Plain text suggested first course of action, such as a command to execute.
     */
    tip?: pulumi.Input<string>;
}

export interface SloTargetAlertRuleRuleParameters {
    /**
     * Burn rate threshold 1 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burnRateThreshold1 parameter.
     */
    burnRateThreshold1?: pulumi.Input<number>;
    /**
     * Burn rate threshold 2 used in burn rate alert calculation. This value must be between 0 and 100/(100-SLO target). Note: BURN_RATE alert rules use the burnRateThreshold2 parameter.
     */
    burnRateThreshold2?: pulumi.Input<number>;
    /**
     * Duration that indicates how long the alert condition is met before the alert is triggered. The value must be positive and smaller than the compliance period of the SLO target. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the fireLasting parameter
     */
    fireLasting?: pulumi.Input<string>;
    /**
     * Long window 1 used in burn rate alert calculation. This value must be longer than shortWindow1` and shorter than 90 days. Note: BURN_RATE alert rules use the longWindow1 parameter.
     */
    longWindow1?: pulumi.Input<string>;
    /**
     * Long window 2 used in burn rate alert calculation. This value must be longer than shortWindow2` and shorter than 90 days. Note: BURN_RATE alert rules use the longWindow2 parameter.
     */
    longWindow2?: pulumi.Input<string>;
    /**
     * Error budget must be equal to or smaller than this percentage for the alert to be triggered. Note: ERROR_BUDGET_LEFT alert rules use the percentErrorBudgetLeft parameter.
     */
    percentErrorBudgetLeft?: pulumi.Input<number>;
    /**
     * Percentage of the fireLasting duration that the alert condition is met before the alert is triggered. Note: BREACH and ERROR_BUDGET_LEFT alert rules use the percentOfLasting parameter
     */
    percentOfLasting?: pulumi.Input<number>;
    /**
     * Short window 1 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_1. Note: BURN_RATE alert rules use the shortWindow1 parameter.
     */
    shortWindow1?: pulumi.Input<string>;
    /**
     * Short window 2 used in burn rate alert calculation. This value must be longer than 1/30 of long_window_2. Note: BURN_RATE alert rules use the shortWindow2 parameter.
     */
    shortWindow2?: pulumi.Input<string>;
}

export interface TableChartVizOption {
    /**
     * Color to use
     */
    color?: pulumi.Input<string>;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    label: pulumi.Input<string>;
    /**
     * An arbitrary prefix to display with the value of this plot
     */
    valuePrefix?: pulumi.Input<string>;
    /**
     * An arbitrary suffix to display with the value of this plot
     */
    valueSuffix?: pulumi.Input<string>;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     */
    valueUnit?: pulumi.Input<string>;
}

export interface TimeChartAxisLeft {
    /**
     * A line to draw as a high watermark
     */
    highWatermark?: pulumi.Input<number>;
    /**
     * A label to attach to the high watermark line
     */
    highWatermarkLabel?: pulumi.Input<string>;
    /**
     * Label of the left axis
     */
    label?: pulumi.Input<string>;
    /**
     * A line to draw as a low watermark
     */
    lowWatermark?: pulumi.Input<number>;
    /**
     * A label to attach to the low watermark line
     */
    lowWatermarkLabel?: pulumi.Input<string>;
    /**
     * The maximum value for the left axis
     */
    maxValue?: pulumi.Input<number>;
    /**
     * The minimum value for the left axis
     */
    minValue?: pulumi.Input<number>;
    watermarks?: pulumi.Input<pulumi.Input<inputs.TimeChartAxisLeftWatermark>[]>;
}

export interface TimeChartAxisLeftWatermark {
    /**
     * Label to display associated with the watermark line
     */
    label?: pulumi.Input<string>;
    /**
     * Axis value where the watermark line will be displayed
     */
    value: pulumi.Input<number>;
}

export interface TimeChartAxisRight {
    /**
     * A line to draw as a high watermark
     */
    highWatermark?: pulumi.Input<number>;
    /**
     * A label to attach to the high watermark line
     */
    highWatermarkLabel?: pulumi.Input<string>;
    /**
     * Label of the right axis
     */
    label?: pulumi.Input<string>;
    /**
     * A line to draw as a low watermark
     */
    lowWatermark?: pulumi.Input<number>;
    /**
     * A label to attach to the low watermark line
     */
    lowWatermarkLabel?: pulumi.Input<string>;
    /**
     * The maximum value for the right axis
     */
    maxValue?: pulumi.Input<number>;
    /**
     * The minimum value for the right axis
     */
    minValue?: pulumi.Input<number>;
    watermarks?: pulumi.Input<pulumi.Input<inputs.TimeChartAxisRightWatermark>[]>;
}

export interface TimeChartAxisRightWatermark {
    /**
     * Label to display associated with the watermark line
     */
    label?: pulumi.Input<string>;
    /**
     * Axis value where the watermark line will be displayed
     */
    value: pulumi.Input<number>;
}

export interface TimeChartEventOption {
    /**
     * Color to use
     */
    color?: pulumi.Input<string>;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The label used in the publish statement that displays the events you want to customize
     */
    label: pulumi.Input<string>;
}

export interface TimeChartHistogramOption {
    /**
     * Base color theme to use for the graph.
     */
    colorTheme?: pulumi.Input<string>;
}

export interface TimeChartLegendOptionsField {
    /**
     * (true by default) Determines if this property is displayed in the data table.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The name of a property to hide or show in the data table.
     */
    property: pulumi.Input<string>;
}

export interface TimeChartVizOption {
    /**
     * The Y-axis associated with values for this plot. Must be either "right" or "left". Defaults to "left".
     */
    axis?: pulumi.Input<string>;
    /**
     * Color to use
     */
    color?: pulumi.Input<string>;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The label used in the publish statement that displays the plot (metric time series data) you want to customize
     */
    label: pulumi.Input<string>;
    /**
     * (Chart plotType by default) The visualization style to use. Must be "LineChart", "AreaChart", "ColumnChart", or "Histogram"
     */
    plotType?: pulumi.Input<string>;
    /**
     * An arbitrary prefix to display with the value of this plot
     */
    valuePrefix?: pulumi.Input<string>;
    /**
     * An arbitrary suffix to display with the value of this plot
     */
    valueSuffix?: pulumi.Input<string>;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes)
     */
    valueUnit?: pulumi.Input<string>;
}

export interface WebhookIntegrationHeader {
    headerKey: pulumi.Input<string>;
    headerValue: pulumi.Input<string>;
}
export namespace aws {
    export interface IntegrationCustomNamespaceSyncRule {
        /**
         * Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filterAction` and `filterSource` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of "Include" or "Exclude".
         */
        defaultAction?: pulumi.Input<string>;
        /**
         * Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of "Include" or "Exclude".
         */
        filterAction?: pulumi.Input<string>;
        /**
         * Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
         */
        filterSource?: pulumi.Input<string>;
        /**
         * An AWS custom namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
         */
        namespace: pulumi.Input<string>;
    }

    export interface IntegrationMetricStatsToSync {
        /**
         * AWS metric that you want to pick statistics for
         */
        metric: pulumi.Input<string>;
        /**
         * An AWS namespace having AWS metric that you want to pick statistics for
         */
        namespace: pulumi.Input<string>;
        /**
         * AWS statistics you want to collect
         */
        stats: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface IntegrationNamespaceSyncRule {
        /**
         * Controls the Splunk Observability default behavior for processing data from an AWS namespace. Splunk Observability ignores this property unless you specify the `filterAction` and `filterSource` properties. If you do specify them, use this property to control how Splunk Observability treats data that doesn't match the filter. The available actions are one of "Include" or "Exclude".
         */
        defaultAction?: pulumi.Input<string>;
        /**
         * Controls how Splunk Observability processes data from a custom AWS namespace. The available actions are one of "Include" or "Exclude".
         */
        filterAction?: pulumi.Input<string>;
        /**
         * Expression that selects the data that Splunk Observability should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
         */
        filterSource?: pulumi.Input<string>;
        /**
         * An AWS namespace having custom AWS metrics that you want to sync with Splunk Observability. See the AWS documentation on publishing metrics for more information.
         */
        namespace: pulumi.Input<string>;
    }
}

export namespace azure {
    export interface IntegrationCustomNamespacesPerService {
        /**
         * The namespaces to sync
         */
        namespaces: pulumi.Input<pulumi.Input<string>[]>;
        /**
         * The name of the service
         */
        service: pulumi.Input<string>;
    }

    export interface IntegrationResourceFilterRule {
        filterSource: pulumi.Input<string>;
    }
}

export namespace gcp {
    export interface IntegrationProjectServiceKey {
        projectId: pulumi.Input<string>;
        projectKey: pulumi.Input<string>;
    }
}

export namespace log {
    export interface ViewColumn {
        /**
         * Name of the column
         */
        name: pulumi.Input<string>;
    }

    export interface ViewSortOption {
        /**
         * Name of the column
         */
        descending: pulumi.Input<boolean>;
        /**
         * Name of the column
         */
        field: pulumi.Input<string>;
    }
}
