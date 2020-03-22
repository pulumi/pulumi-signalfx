// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Jira
{
    /// <summary>
    /// SignalFx Jira integrations. For help with this integration see [Integration with Jira](https://docs.signalfx.com/en/latest/admin-guide/integrate-notifications.html#integrate-with-jira).
    /// 
    /// &gt; **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider. Otherwise you'll receive a 4xx error.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-signalfx/blob/master/website/docs/r/jira_integration.html.markdown.
    /// </summary>
    public partial class Integration : Pulumi.CustomResource
    {
        /// <summary>
        /// The API token for the user email
        /// </summary>
        [Output("apiToken")]
        public Output<string?> ApiToken { get; private set; } = null!;

        /// <summary>
        /// Jira display name for the assignee.
        /// </summary>
        [Output("assigneeDisplayName")]
        public Output<string?> AssigneeDisplayName { get; private set; } = null!;

        /// <summary>
        /// Jira user name for the assignee.
        /// </summary>
        [Output("assigneeName")]
        public Output<string> AssigneeName { get; private set; } = null!;

        /// <summary>
        /// Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        /// </summary>
        [Output("authMethod")]
        public Output<string> AuthMethod { get; private set; } = null!;

        /// <summary>
        /// Base URL of the Jira instance that's integrated with SignalFx.
        /// </summary>
        [Output("baseUrl")]
        public Output<string> BaseUrl { get; private set; } = null!;

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        /// </summary>
        [Output("issueType")]
        public Output<string> IssueType { get; private set; } = null!;

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password used to authenticate the Jira integration.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        /// </summary>
        [Output("projectKey")]
        public Output<string> ProjectKey { get; private set; } = null!;

        /// <summary>
        /// Email address used to authenticate the Jira integration.
        /// </summary>
        [Output("userEmail")]
        public Output<string?> UserEmail { get; private set; } = null!;

        /// <summary>
        /// User name used to authenticate the Jira integration.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Integration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Integration(string name, IntegrationArgs args, CustomResourceOptions? options = null)
            : base("signalfx:jira/integration:Integration", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Integration(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:jira/integration:Integration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Integration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Integration Get(string name, Input<string> id, IntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new Integration(name, id, state, options);
        }
    }

    public sealed class IntegrationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API token for the user email
        /// </summary>
        [Input("apiToken")]
        public Input<string>? ApiToken { get; set; }

        /// <summary>
        /// Jira display name for the assignee.
        /// </summary>
        [Input("assigneeDisplayName")]
        public Input<string>? AssigneeDisplayName { get; set; }

        /// <summary>
        /// Jira user name for the assignee.
        /// </summary>
        [Input("assigneeName", required: true)]
        public Input<string> AssigneeName { get; set; } = null!;

        /// <summary>
        /// Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        /// </summary>
        [Input("authMethod", required: true)]
        public Input<string> AuthMethod { get; set; } = null!;

        /// <summary>
        /// Base URL of the Jira instance that's integrated with SignalFx.
        /// </summary>
        [Input("baseUrl", required: true)]
        public Input<string> BaseUrl { get; set; } = null!;

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        /// </summary>
        [Input("issueType", required: true)]
        public Input<string> IssueType { get; set; } = null!;

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password used to authenticate the Jira integration.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        /// </summary>
        [Input("projectKey", required: true)]
        public Input<string> ProjectKey { get; set; } = null!;

        /// <summary>
        /// Email address used to authenticate the Jira integration.
        /// </summary>
        [Input("userEmail")]
        public Input<string>? UserEmail { get; set; }

        /// <summary>
        /// User name used to authenticate the Jira integration.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public IntegrationArgs()
        {
        }
    }

    public sealed class IntegrationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API token for the user email
        /// </summary>
        [Input("apiToken")]
        public Input<string>? ApiToken { get; set; }

        /// <summary>
        /// Jira display name for the assignee.
        /// </summary>
        [Input("assigneeDisplayName")]
        public Input<string>? AssigneeDisplayName { get; set; }

        /// <summary>
        /// Jira user name for the assignee.
        /// </summary>
        [Input("assigneeName")]
        public Input<string>? AssigneeName { get; set; }

        /// <summary>
        /// Authentication method used when creating the Jira integration. One of `EmailAndToken` (using `user_email` and `api_token`) or `UsernameAndPassword` (using `username` and `password`).
        /// </summary>
        [Input("authMethod")]
        public Input<string>? AuthMethod { get; set; }

        /// <summary>
        /// Base URL of the Jira instance that's integrated with SignalFx.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Issue type (for example, Story) for tickets that Jira creates for detector notifications. SignalFx validates issue types, so you must specify a type that's valid for the Jira project specified in `projectKey`.
        /// </summary>
        [Input("issueType")]
        public Input<string>? IssueType { get; set; }

        /// <summary>
        /// Name of the integration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password used to authenticate the Jira integration.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Jira key of an existing project. When Jira creates a new ticket for a detector notification, the ticket is assigned to this project.
        /// </summary>
        [Input("projectKey")]
        public Input<string>? ProjectKey { get; set; }

        /// <summary>
        /// Email address used to authenticate the Jira integration.
        /// </summary>
        [Input("userEmail")]
        public Input<string>? UserEmail { get; set; }

        /// <summary>
        /// User name used to authenticate the Jira integration.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public IntegrationState()
        {
        }
    }
}
