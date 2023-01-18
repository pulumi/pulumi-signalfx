// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Gcp
{
    public static class GetServices
    {
        /// <summary>
        /// Use this data source to get a list of GCP service names.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using SignalFx = Pulumi.SignalFx;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var gcpServices = SignalFx.Gcp.GetServices.Invoke();
        /// 
        ///     // Leaves out most of the integration bits, see the docs
        ///     // for signalfx_gcp_integration for more
        ///     // …
        ///     var gcpMyteam = new SignalFx.Gcp.Integration("gcpMyteam", new()
        ///     {
        ///         Services = new[]
        ///         {
        ///             gcpServices.Apply(getServicesResult =&gt; getServicesResult.Services),
        ///         }.Select(__item =&gt; __item?.Name).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServicesResult> InvokeAsync(GetServicesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServicesResult>("signalfx:gcp/getServices:getServices", args ?? new GetServicesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a list of GCP service names.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using SignalFx = Pulumi.SignalFx;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var gcpServices = SignalFx.Gcp.GetServices.Invoke();
        /// 
        ///     // Leaves out most of the integration bits, see the docs
        ///     // for signalfx_gcp_integration for more
        ///     // …
        ///     var gcpMyteam = new SignalFx.Gcp.Integration("gcpMyteam", new()
        ///     {
        ///         Services = new[]
        ///         {
        ///             gcpServices.Apply(getServicesResult =&gt; getServicesResult.Services),
        ///         }.Select(__item =&gt; __item?.Name).ToList(),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetServicesResult> Invoke(GetServicesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServicesResult>("signalfx:gcp/getServices:getServices", args ?? new GetServicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServicesArgs : global::Pulumi.InvokeArgs
    {
        [Input("services")]
        private List<Inputs.GetServicesServiceArgs>? _services;
        public List<Inputs.GetServicesServiceArgs> Services
        {
            get => _services ?? (_services = new List<Inputs.GetServicesServiceArgs>());
            set => _services = value;
        }

        public GetServicesArgs()
        {
        }
        public static new GetServicesArgs Empty => new GetServicesArgs();
    }

    public sealed class GetServicesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("services")]
        private InputList<Inputs.GetServicesServiceInputArgs>? _services;
        public InputList<Inputs.GetServicesServiceInputArgs> Services
        {
            get => _services ?? (_services = new InputList<Inputs.GetServicesServiceInputArgs>());
            set => _services = value;
        }

        public GetServicesInvokeArgs()
        {
        }
        public static new GetServicesInvokeArgs Empty => new GetServicesInvokeArgs();
    }


    [OutputType]
    public sealed class GetServicesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetServicesServiceResult> Services;

        [OutputConstructor]
        private GetServicesResult(
            string id,

            ImmutableArray<Outputs.GetServicesServiceResult> services)
        {
            Id = id;
            Services = services;
        }
    }
}
