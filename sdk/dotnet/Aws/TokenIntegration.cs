// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Aws
{
    /// <summary>
    /// SignalFx AWS CloudWatch integrations using security tokens. For help with this integration see [Connect to AWS CloudWatch](https://docs.signalfx.com/en/latest/integrations/amazon-web-services.html#connect-to-aws).
    /// 
    /// &gt; **NOTE** When managing integrations you'll need to use an admin token to authenticate the SignalFx provider.
    /// 
    /// &gt; **WARNING** This resource implements a part of a workflow. You must use it with `signalfx.aws.Integration`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// using SignalFx = Pulumi.SignalFx;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var awsMyteamToken = new SignalFx.Aws.TokenIntegration("awsMyteamToken", new SignalFx.Aws.TokenIntegrationArgs
    ///         {
    ///         });
    ///         // Make yourself an AWS IAM role here
    ///         var awsSfxRole = new Aws.Iam.Role("awsSfxRole", new Aws.Iam.RoleArgs
    ///         {
    ///         });
    ///         // Stuff here that uses the external and account ID
    ///         var awsMyteam = new SignalFx.Aws.Integration("awsMyteam", new SignalFx.Aws.IntegrationArgs
    ///         {
    ///             Enabled = true,
    ///             IntegrationId = awsMyteamToken.Id,
    ///             Token = "put_your_token_here",
    ///             Key = "put_your_key_here",
    ///             Regions = 
    ///             {
    ///                 "us-east-1",
    ///             },
    ///             PollRate = 300,
    ///             ImportCloudWatch = true,
    ///             EnableAwsUsage = true,
    ///             CustomNamespaceSyncRules = 
    ///             {
    ///                 new SignalFx.Aws.Inputs.IntegrationCustomNamespaceSyncRuleArgs
    ///                 {
    ///                     DefaultAction = "Exclude",
    ///                     FilterAction = "Include",
    ///                     FilterSource = "filter('code', '200')",
    ///                     Namespace = "fart",
    ///                 },
    ///             },
    ///             NamespaceSyncRules = 
    ///             {
    ///                 new SignalFx.Aws.Inputs.IntegrationNamespaceSyncRuleArgs
    ///                 {
    ///                     DefaultAction = "Exclude",
    ///                     FilterAction = "Include",
    ///                     FilterSource = "filter('code', '200')",
    ///                     Namespace = "AWS/EC2",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [SignalFxResourceType("signalfx:aws/tokenIntegration:TokenIntegration")]
    public partial class TokenIntegration : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of this integration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
        /// </summary>
        [Output("signalfxAwsAccount")]
        public Output<string> SignalfxAwsAccount { get; private set; } = null!;

        /// <summary>
        /// The SignalFx-generated AWS token to use with an AWS integration.
        /// </summary>
        [Output("tokenId")]
        public Output<string> TokenId { get; private set; } = null!;


        /// <summary>
        /// Create a TokenIntegration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TokenIntegration(string name, TokenIntegrationArgs? args = null, CustomResourceOptions? options = null)
            : base("signalfx:aws/tokenIntegration:TokenIntegration", name, args ?? new TokenIntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TokenIntegration(string name, Input<string> id, TokenIntegrationState? state = null, CustomResourceOptions? options = null)
            : base("signalfx:aws/tokenIntegration:TokenIntegration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TokenIntegration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TokenIntegration Get(string name, Input<string> id, TokenIntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new TokenIntegration(name, id, state, options);
        }
    }

    public sealed class TokenIntegrationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of this integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TokenIntegrationArgs()
        {
        }
    }

    public sealed class TokenIntegrationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of this integration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AWS Account ARN to use with your policies/roles, provided by SignalFx.
        /// </summary>
        [Input("signalfxAwsAccount")]
        public Input<string>? SignalfxAwsAccount { get; set; }

        /// <summary>
        /// The SignalFx-generated AWS token to use with an AWS integration.
        /// </summary>
        [Input("tokenId")]
        public Input<string>? TokenId { get; set; }

        public TokenIntegrationState()
        {
        }
    }
}
