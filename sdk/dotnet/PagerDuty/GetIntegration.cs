// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.PagerDuty
{
    public static class GetIntegration
    {
        /// <summary>
        /// Use this data source to get information on an existing PagerDuty integration.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using SignalFx = Pulumi.SignalFx;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var pdIntegration = Output.Create(SignalFx.PagerDuty.GetIntegration.InvokeAsync(new SignalFx.PagerDuty.GetIntegrationArgs
        ///         {
        ///             Name = "PD-Integration",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIntegrationResult> InvokeAsync(GetIntegrationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationResult>("signalfx:pagerduty/getIntegration:getIntegration", args ?? new GetIntegrationArgs(), options.WithVersion());
    }


    public sealed class GetIntegrationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the exact name of the desired PagerDuty integration
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetIntegrationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIntegrationResult
    {
        /// <summary>
        /// Whether the integration is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the integration.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetIntegrationResult(
            bool enabled,

            string id,

            string name)
        {
            Enabled = enabled;
            Id = id;
            Name = name;
        }
    }
}
