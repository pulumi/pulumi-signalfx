// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface AlertMutingRuleFilter {
    /**
     * Determines if this is a "not" filter. Defaults to `false`.
     */
    negated?: boolean;
    /**
     * The property to filter.
     */
    property: string;
    /**
     * The property value to filter.
     */
    propertyValue: string;
}

export interface DashboardChart {
    /**
     * ID of the chart to display.
     */
    chartId: string;
    /**
     * Column number for the layout.
     */
    column?: number;
    /**
     * How many rows every chart should take up (greater than or equal to 1). 1 by default.
     */
    height?: number;
    /**
     * The row to show the chart in (zero-based); if `height > 1`, this value represents the topmost row of the chart (greater than or equal to `0`).
     */
    row?: number;
    /**
     * How many columns (out of a total of `12`) every chart should take up (between `1` and `12`). `12` by default.
     */
    width?: number;
}

export interface DashboardColumn {
    /**
     * List of IDs of the charts to display.
     */
    chartIds: string[];
    /**
     * Column number for the layout.
     */
    column?: number;
    /**
     * How many rows every chart should take up (greater than or equal to 1). 1 by default.
     */
    height?: number;
    /**
     * How many columns (out of a total of `12`) every chart should take up (between `1` and `12`). `12` by default.
     */
    width?: number;
}

export interface DashboardEventOverlay {
    /**
     * Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
     */
    color?: string;
    /**
     * Text shown in the dropdown when selecting this overlay from the menu.
     */
    label?: string;
    /**
     * Show a vertical line for the event. `false` by default.
     */
    line?: boolean;
    /**
     * Search term used to choose the events shown in the overlay.
     */
    signal: string;
    /**
     * Each element specifies a filter to use against the signal specified in the `signal`.
     */
    sources?: outputs.DashboardEventOverlaySource[];
    /**
     * Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
     */
    type?: string;
}

export interface DashboardEventOverlaySource {
    /**
     * If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
     */
    negated?: boolean;
    /**
     * The name of a dimension to filter against.
     */
    property: string;
    /**
     * A list of values to be used with the `property`, they will be combined via `OR`.
     */
    values: string[];
}

export interface DashboardFilter {
    /**
     * If true, this variable will also match data that doesn't have this property at all.
     */
    applyIfExist?: boolean;
    /**
     * If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
     */
    negated?: boolean;
    /**
     * The name of a dimension to filter against.
     */
    property: string;
    /**
     * A list of values to be used with the `property`, they will be combined via `OR`.
     */
    values: string[];
}

export interface DashboardGrid {
    /**
     * List of IDs of the charts to display.
     */
    chartIds: string[];
    /**
     * How many rows every chart should take up (greater than or equal to 1). 1 by default.
     */
    height?: number;
    /**
     * How many columns (out of a total of `12`) every chart should take up (between `1` and `12`). `12` by default.
     */
    width?: number;
}

export interface DashboardGroupDashboard {
    /**
     * The dashboard id to mirror
     */
    dashboardId: string;
    /**
     * The description that will override the original dashboards's description.
     */
    descriptionOverride?: string;
    /**
     * The description that will override the original dashboards's description.
     */
    filterOverrides?: outputs.DashboardGroupDashboardFilterOverride[];
    /**
     * The name that will override the original dashboards's name.
     */
    nameOverride?: string;
    variableOverrides?: outputs.DashboardGroupDashboardVariableOverride[];
}

export interface DashboardGroupDashboardFilterOverride {
    /**
     * If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
     */
    negated?: boolean;
    /**
     * A metric time series dimension or property name.
     */
    property: string;
    /**
     * (Optional) List of of strings (which will be treated as an OR filter on the property).
     */
    values: string[];
}

export interface DashboardGroupDashboardVariableOverride {
    /**
     * A metric time series dimension or property name.
     */
    property: string;
    /**
     * (Optional) List of of strings (which will be treated as an OR filter on the property).
     */
    values?: string[];
    /**
     * A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
     */
    valuesSuggesteds?: string[];
}

export interface DashboardGroupImportQualifier {
    filters?: outputs.DashboardGroupImportQualifierFilter[];
    metric?: string;
}

export interface DashboardGroupImportQualifierFilter {
    /**
     * If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
     */
    negated?: boolean;
    /**
     * A metric time series dimension or property name.
     */
    property: string;
    /**
     * (Optional) List of of strings (which will be treated as an OR filter on the property).
     */
    values: string[];
}

export interface DashboardSelectedEventOverlay {
    /**
     * Search term used to choose the events shown in the overlay.
     */
    signal: string;
    /**
     * Each element specifies a filter to use against the signal specified in the `signal`.
     */
    sources?: outputs.DashboardSelectedEventOverlaySource[];
    /**
     * Can be set to `eventTimeSeries` (the default) to refer to externally reported events, or `detectorEvents` to refer to events from detector triggers.
     */
    type?: string;
}

export interface DashboardSelectedEventOverlaySource {
    /**
     * If true,  only data that does not match the specified value of the specified property appear in the event overlay. Defaults to `false`.
     */
    negated?: boolean;
    /**
     * The name of a dimension to filter against.
     */
    property: string;
    /**
     * A list of values to be used with the `property`, they will be combined via `OR`.
     */
    values: string[];
}

export interface DashboardVariable {
    /**
     * An alias for the dashboard variable. This text will appear as the label for the dropdown field on the dashboard.
     */
    alias: string;
    /**
     * If true, this variable will also match data that doesn't have this property at all.
     */
    applyIfExist?: boolean;
    /**
     * Variable description.
     */
    description?: string;
    /**
     * The name of a dimension to filter against.
     */
    property: string;
    /**
     * If `true`, this variable will only apply to charts that have a filter for the property.
     */
    replaceOnly?: boolean;
    /**
     * If `true`, this variable may only be set to the values listed in `valuesSuggested` and only these values will appear in autosuggestion menus. `false` by default.
     */
    restrictedSuggestions?: boolean;
    /**
     * Determines whether a value is required for this variable (and therefore whether it will be possible to view this dashboard without this filter applied). `false` by default.
     */
    valueRequired?: boolean;
    /**
     * A list of values to be used with the `property`, they will be combined via `OR`.
     */
    values?: string[];
    /**
     * A list of strings of suggested values for this variable; these suggestions will receive priority when values are autosuggested for this variable.
     */
    valuesSuggesteds?: string[];
}

export interface DataLinkTargetExternalUrl {
    /**
     * Flag that designates a target as the default for a data link object. `true` by default
     */
    isDefault?: boolean;
    /**
     * The [minimum time window](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) for a search sent to an external site. Defaults to `6000`
     */
    minimumTimeWindow?: string;
    /**
     * User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     */
    name: string;
    /**
     * Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.
     */
    propertyKeyMapping?: {[key: string]: string};
    /**
     * [Designates the format](https://developers.signalfx.com/administration/data_links_overview.html#_minimum_time_window) of `minimumTimeWindow` in the same data link target object. Must be one of `"ISO8601"`, `"EpochSeconds"` or `"Epoch"` (which is milliseconds). Defaults to `"ISO8601"`.
     */
    timeFormat?: string;
    /**
     * URL string for a Splunk instance or external system data link target. [See the supported template variables](https://developers.signalfx.com/administration/data_links_overview.html#_external_link_targets).
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
     * Flag that designates a target as the default for a data link object. `true` by default
     */
    isDefault?: boolean;
    /**
     * User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     */
    name: string;
}

export interface DataLinkTargetSplunk {
    /**
     * Flag that designates a target as the default for a data link object. `true` by default
     */
    isDefault?: boolean;
    /**
     * User-assigned target name. Use this value to differentiate between the link targets for a data link object.
     */
    name: string;
    /**
     * Describes the relationship between SignalFx metadata keys and external system properties when the key names are different.
     */
    propertyKeyMapping?: {[key: string]: string};
}

export interface DetectorRule {
    /**
     * Description for the rule. Displays as the alert condition in the Alert Rules tab of the detector editor in the web UI.
     */
    description?: string;
    /**
     * A detect label which matches a detect label within `programText`.
     */
    detectLabel: string;
    /**
     * When true, notifications and events will not be generated for the detect label. `false` by default.
     */
    disabled?: boolean;
    /**
     * List of strings specifying where notifications will be sent when an incident occurs. See [Create A Single Detector](https://developers.signalfx.com/detectors_reference.html#operation/Create%20Single%20Detector) for more info.
     */
    notifications?: string[];
    /**
     * Custom notification message body when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.signalfx.com/en/latest/detect-alert/set-up-detectors.html#about-detectors#alert-settings) for more info.
     */
    parameterizedBody?: string;
    /**
     * Custom notification message subject when an alert is triggered. See [Set Up Detectors to Trigger Alerts](https://docs.signalfx.com/en/latest/detect-alert/set-up-detectors.html#about-detectors#alert-settings) for more info.
     */
    parameterizedSubject?: string;
    /**
     * URL of page to consult when an alert is triggered. This can be used with custom notification messages.
     */
    runbookUrl?: string;
    /**
     * The severity of the rule, must be one of: `"Critical"`, `"Major"`, `"Minor"`, `"Warning"`, `"Info"`.
     */
    severity: string;
    /**
     * Plain text suggested first course of action, such as a command line to execute. This can be used with custom notification messages.
     */
    tip?: string;
}

export interface DetectorVizOption {
    /**
     * Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * Label used in the publish statement that displays the plot (metric time series data) you want to customize.
     */
    label: string;
    valuePrefix?: string;
    valueSuffix?: string;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
     * * `valuePrefix`, `valueSuffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
     */
    valueUnit?: string;
}

export interface GetAwsServicesService {
    name: string;
}

export interface GetAzureServicesService {
    name: string;
}

export interface HeatmapChartColorRange {
    /**
     * The color range to use. Hex values are not supported here. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color: string;
    /**
     * The maximum value within the coloring range.
     */
    maxValue?: number;
    /**
     * The minimum value within the coloring range.
     */
    minValue?: number;
}

export interface HeatmapChartColorScale {
    /**
     * The color range to use. Hex values are not supported here. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color: string;
    /**
     * Indicates the lower threshold non-inclusive value for this range.
     */
    gt?: number;
    /**
     * Indicates the lower threshold inclusive value for this range.
     */
    gte?: number;
    /**
     * Indicates the upper threshold non-inclusive value for this range.
     */
    lt?: number;
    /**
     * Indicates the upper threshold inclusive value for this range.
     */
    lte?: number;
}

export interface ListChartColorScale {
    /**
     * The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color: string;
    /**
     * Indicates the lower threshold non-inclusive value for this range.
     */
    gt?: number;
    /**
     * Indicates the lower threshold inclusive value for this range.
     */
    gte?: number;
    /**
     * Indicates the upper threshold non-inculsive value for this range.
     */
    lt?: number;
    /**
     * Indicates the upper threshold inclusive value for this range.
     */
    lte?: number;
}

export interface ListChartLegendOptionsField {
    /**
     * True or False depending on if you want the property to be shown or hidden.
     */
    enabled?: boolean;
    /**
     * The name of the property to display. Note the special values of `sfMetric` (corresponding with the API's `Plot Name`) which shows the label of the time series `publish()` and `sf_originatingMetric` (corresponding with the API's `metric (sf metric)`) that shows the [name of the metric](https://developers.signalfx.com/signalflow_analytics/functions/data_function.html#table-1-parameter-definitions) for the time series being displayed.
     */
    property: string;
}

export interface ListChartVizOption {
    /**
     * The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * Label used in the publish statement that displays the plot (metric time series data) you want to customize.
     */
    label: string;
    valuePrefix?: string;
    valueSuffix?: string;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
     * * `valuePrefix`, `valueSuffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
     */
    valueUnit?: string;
}

export interface OrgTokenDpmLimits {
    /**
     * The datapoints per minute (dpm) limit for this token. If you exceed this limit, SignalFx sends out an alert.
     */
    dpmLimit: number;
    /**
     * DPM level at which SignalFx sends the notification for this token. If you don't specify a notification, SignalFx sends the generic notification.
     */
    dpmNotificationThreshold?: number;
}

export interface OrgTokenHostOrUsageLimits {
    /**
     * Max number of Docker containers that can use this token
     */
    containerLimit?: number;
    /**
     * Notification threshold for Docker containers
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
     * Max number of hi-res metrics that can be sent with this toke
     */
    highResMetricsLimit?: number;
    /**
     * Notification threshold for hi-res metrics
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
     * Indicates the lower threshold non-inclusive value for this range.
     */
    gt?: number;
    /**
     * Indicates the lower threshold inclusive value for this range.
     */
    gte?: number;
    /**
     * Indicates the upper threshold non-inculsive value for this range.
     */
    lt?: number;
    /**
     * Indicates the upper threshold inclusive value for this range.
     */
    lte?: number;
}

export interface SingleValueChartVizOption {
    /**
     * The color to use. Must be one of gray, blue, light_blue, navy, dark_orange, orange, dark_yellow, magenta, cerise, pink, violet, purple, gray_blue, dark_green, green, aquamarine, red, yellow, vivid_yellow, light_green, or lime_green.
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * Label used in the publish statement that displays the plot (metric time series data) you want to customize.
     */
    label: string;
    valuePrefix?: string;
    valueSuffix?: string;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
     * * `valuePrefix`, `valueSuffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
     */
    valueUnit?: string;
}

export interface TimeChartAxisLeft {
    /**
     * A line to draw as a high watermark.
     */
    highWatermark?: number;
    /**
     * A label to attach to the high watermark line.
     */
    highWatermarkLabel?: string;
    /**
     * Label used in the publish statement that displays the event query you want to customize.
     */
    label?: string;
    /**
     * A line to draw as a low watermark.
     */
    lowWatermark?: number;
    /**
     * A label to attach to the low watermark line.
     */
    lowWatermarkLabel?: string;
    /**
     * The maximum value for the right axis.
     */
    maxValue?: number;
    /**
     * The minimum value for the right axis.
     */
    minValue?: number;
    watermarks?: outputs.TimeChartAxisLeftWatermark[];
}

export interface TimeChartAxisLeftWatermark {
    /**
     * Label used in the publish statement that displays the event query you want to customize.
     */
    label?: string;
    value: number;
}

export interface TimeChartAxisRight {
    /**
     * A line to draw as a high watermark.
     */
    highWatermark?: number;
    /**
     * A label to attach to the high watermark line.
     */
    highWatermarkLabel?: string;
    /**
     * Label used in the publish statement that displays the event query you want to customize.
     */
    label?: string;
    /**
     * A line to draw as a low watermark.
     */
    lowWatermark?: number;
    /**
     * A label to attach to the low watermark line.
     */
    lowWatermarkLabel?: string;
    /**
     * The maximum value for the right axis.
     */
    maxValue?: number;
    /**
     * The minimum value for the right axis.
     */
    minValue?: number;
    watermarks?: outputs.TimeChartAxisRightWatermark[];
}

export interface TimeChartAxisRightWatermark {
    /**
     * Label used in the publish statement that displays the event query you want to customize.
     */
    label?: string;
    value: number;
}

export interface TimeChartEventOption {
    /**
     * Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * Label used in the publish statement that displays the event query you want to customize.
     */
    label: string;
}

export interface TimeChartHistogramOption {
    /**
     * Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine, red, gold, greenyellow, chartreuse, jade
     */
    colorTheme?: string;
}

export interface TimeChartLegendOptionsField {
    /**
     * True or False depending on if you want the property to be shown or hidden.
     */
    enabled?: boolean;
    /**
     * The name of the property to display. Note the special values of `plotLabel` (corresponding with the API's `sfMetric`) which shows the label of the time series `publish()` and `metric` (corresponding with the API's `sf_originatingMetric`) that shows the name of the metric for the time series being displayed.
     */
    property: string;
}

export interface TimeChartVizOption {
    /**
     * Y-axis associated with values for this plot. Must be either `right` or `left`.
     */
    axis?: string;
    /**
     * Color to use : gray, blue, azure, navy, brown, orange, yellow, iris, magenta, pink, purple, violet, lilac, emerald, green, aquamarine.
     */
    color?: string;
    /**
     * Specifies an alternate value for the Plot Name column of the Data Table associated with the chart.
     */
    displayName?: string;
    /**
     * Label used in the publish statement that displays the event query you want to customize.
     */
    label: string;
    /**
     * The visualization style to use. Must be `"LineChart"`, `"AreaChart"`, `"ColumnChart"`, or `"Histogram"`. Chart level `plotType` by default.
     */
    plotType?: string;
    valuePrefix?: string;
    valueSuffix?: string;
    /**
     * A unit to attach to this plot. Units support automatic scaling (eg thousands of bytes will be displayed as kilobytes). Values values are `Bit, Kilobit, Megabit, Gigabit, Terabit, Petabit, Exabit, Zettabit, Yottabit, Byte, Kibibyte, Mebibyte, Gigibyte, Tebibyte, Pebibyte, Exbibyte, Zebibyte, Yobibyte, Nanosecond, Microsecond, Millisecond, Second, Minute, Hour, Day, Week`.
     * * `valuePrefix`, `valueSuffix` - (Optional) Arbitrary prefix/suffix to display with the value of this plot.
     */
    valueUnit?: string;
}

export interface WebhookIntegrationHeader {
    /**
     * The key of the header to send
     */
    headerKey: string;
    /**
     * The value of the header to send
     */
    headerValue: string;
}
export namespace aws {
    export interface GetServicesService {
        name: string;
    }

    export interface IntegrationCustomNamespaceSyncRule {
        /**
         * Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
         */
        defaultAction?: string;
        /**
         * Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
         */
        filterAction?: string;
        /**
         * Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
         */
        filterSource?: string;
        /**
         * An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
         */
        namespace: string;
    }

    export interface IntegrationNamespaceSyncRule {
        /**
         * Controls the SignalFx default behavior for processing data from an AWS namespace. If you do specify a filter, use this property to control how SignalFx treats data that doesn't match the filter. The available actions are one of `"Include"` or `"Exclude"`.
         */
        defaultAction?: string;
        /**
         * Controls how SignalFx processes data from a custom AWS namespace. The available actions are one of `"Include"` or `"Exclude"`.
         */
        filterAction?: string;
        /**
         * Expression that selects the data that SignalFx should sync for the custom namespace associated with this sync rule. The expression uses the syntax defined for the SignalFlow `filter()` function; it can be any valid SignalFlow filter expression.
         */
        filterSource?: string;
        /**
         * An AWS custom namespace having custom AWS metrics that you want to sync with SignalFx. See the AWS documentation on publishing metrics for more information.
         */
        namespace: string;
    }
}

export namespace azure {
    export interface GetServicesService {
        name: string;
    }

    export interface IntegrationCustomNamespacesPerService {
        /**
         * The additional namespaces.
         */
        namespaces: string[];
        /**
         * The name of the service.
         */
        service: string;
    }
}

export namespace gcp {
    export interface IntegrationProjectServiceKey {
        projectId: string;
        projectKey: string;
    }
}
