// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package jira

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// SignalFx Jira integrations. For help with this integration see [Integration with Jira](https://docs.signalfx.com/en/latest/admin-guide/integrate-notifications.html#integrate-with-jira).
//
// > **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/jira_integration.html.markdown.
type Integration struct {
	pulumi.CustomResourceState

	// The API token for the user email
	ApiToken pulumi.StringPtrOutput `pulumi:"apiToken"`
	// Jira display name for the assignee.
	AssigneeDisplayName pulumi.StringPtrOutput `pulumi:"assigneeDisplayName"`
	// Jira user name for the assignee.
	AssigneeName pulumi.StringOutput `pulumi:"assigneeName"`
	// Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `userEmail` and `apiToken`) or `UsernameAndPassword` (using `username` and `password`).
	AuthMethod pulumi.StringOutput `pulumi:"authMethod"`
	// Base URL of the Jira instance that's integrated with SignalFx.
	BaseUrl pulumi.StringOutput `pulumi:"baseUrl"`
	// Whether the integration is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
	IssueType pulumi.StringOutput `pulumi:"issueType"`
	// Name of the integration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password used to authenticate the Jira integration.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
	ProjectKey pulumi.StringOutput `pulumi:"projectKey"`
	// Email address used to authenticate the Jira integration.
	UserEmail pulumi.StringPtrOutput `pulumi:"userEmail"`
	// User name used to authenticate the Jira integration.
	Username pulumi.StringPtrOutput `pulumi:"username"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil || args.AssigneeName == nil {
		return nil, errors.New("missing required argument 'AssigneeName'")
	}
	if args == nil || args.AuthMethod == nil {
		return nil, errors.New("missing required argument 'AuthMethod'")
	}
	if args == nil || args.BaseUrl == nil {
		return nil, errors.New("missing required argument 'BaseUrl'")
	}
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.IssueType == nil {
		return nil, errors.New("missing required argument 'IssueType'")
	}
	if args == nil || args.ProjectKey == nil {
		return nil, errors.New("missing required argument 'ProjectKey'")
	}
	if args == nil {
		args = &IntegrationArgs{}
	}
	var resource Integration
	err := ctx.RegisterResource("signalfx:jira/integration:Integration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegration gets an existing Integration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationState, opts ...pulumi.ResourceOption) (*Integration, error) {
	var resource Integration
	err := ctx.ReadResource("signalfx:jira/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// The API token for the user email
	ApiToken *string `pulumi:"apiToken"`
	// Jira display name for the assignee.
	AssigneeDisplayName *string `pulumi:"assigneeDisplayName"`
	// Jira user name for the assignee.
	AssigneeName *string `pulumi:"assigneeName"`
	// Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `userEmail` and `apiToken`) or `UsernameAndPassword` (using `username` and `password`).
	AuthMethod *string `pulumi:"authMethod"`
	// Base URL of the Jira instance that's integrated with SignalFx.
	BaseUrl *string `pulumi:"baseUrl"`
	// Whether the integration is enabled.
	Enabled *bool `pulumi:"enabled"`
	// Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
	IssueType *string `pulumi:"issueType"`
	// Name of the integration.
	Name *string `pulumi:"name"`
	// Password used to authenticate the Jira integration.
	Password *string `pulumi:"password"`
	// Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
	ProjectKey *string `pulumi:"projectKey"`
	// Email address used to authenticate the Jira integration.
	UserEmail *string `pulumi:"userEmail"`
	// User name used to authenticate the Jira integration.
	Username *string `pulumi:"username"`
}

type IntegrationState struct {
	// The API token for the user email
	ApiToken pulumi.StringPtrInput
	// Jira display name for the assignee.
	AssigneeDisplayName pulumi.StringPtrInput
	// Jira user name for the assignee.
	AssigneeName pulumi.StringPtrInput
	// Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `userEmail` and `apiToken`) or `UsernameAndPassword` (using `username` and `password`).
	AuthMethod pulumi.StringPtrInput
	// Base URL of the Jira instance that's integrated with SignalFx.
	BaseUrl pulumi.StringPtrInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolPtrInput
	// Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
	IssueType pulumi.StringPtrInput
	// Name of the integration.
	Name pulumi.StringPtrInput
	// Password used to authenticate the Jira integration.
	Password pulumi.StringPtrInput
	// Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
	ProjectKey pulumi.StringPtrInput
	// Email address used to authenticate the Jira integration.
	UserEmail pulumi.StringPtrInput
	// User name used to authenticate the Jira integration.
	Username pulumi.StringPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// The API token for the user email
	ApiToken *string `pulumi:"apiToken"`
	// Jira display name for the assignee.
	AssigneeDisplayName *string `pulumi:"assigneeDisplayName"`
	// Jira user name for the assignee.
	AssigneeName string `pulumi:"assigneeName"`
	// Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `userEmail` and `apiToken`) or `UsernameAndPassword` (using `username` and `password`).
	AuthMethod string `pulumi:"authMethod"`
	// Base URL of the Jira instance that's integrated with SignalFx.
	BaseUrl string `pulumi:"baseUrl"`
	// Whether the integration is enabled.
	Enabled bool `pulumi:"enabled"`
	// Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
	IssueType string `pulumi:"issueType"`
	// Name of the integration.
	Name *string `pulumi:"name"`
	// Password used to authenticate the Jira integration.
	Password *string `pulumi:"password"`
	// Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
	ProjectKey string `pulumi:"projectKey"`
	// Email address used to authenticate the Jira integration.
	UserEmail *string `pulumi:"userEmail"`
	// User name used to authenticate the Jira integration.
	Username *string `pulumi:"username"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// The API token for the user email
	ApiToken pulumi.StringPtrInput
	// Jira display name for the assignee.
	AssigneeDisplayName pulumi.StringPtrInput
	// Jira user name for the assignee.
	AssigneeName pulumi.StringInput
	// Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `userEmail` and `apiToken`) or `UsernameAndPassword` (using `username` and `password`).
	AuthMethod pulumi.StringInput
	// Base URL of the Jira instance that's integrated with SignalFx.
	BaseUrl pulumi.StringInput
	// Whether the integration is enabled.
	Enabled pulumi.BoolInput
	// Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
	IssueType pulumi.StringInput
	// Name of the integration.
	Name pulumi.StringPtrInput
	// Password used to authenticate the Jira integration.
	Password pulumi.StringPtrInput
	// Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
	ProjectKey pulumi.StringInput
	// Email address used to authenticate the Jira integration.
	UserEmail pulumi.StringPtrInput
	// User name used to authenticate the Jira integration.
	Username pulumi.StringPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}

