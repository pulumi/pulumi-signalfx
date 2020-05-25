// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.SignalFx.Aws
{
    public static class GetServices
    {
        /// <summary>
        /// Use this data source to get a list of AWS service names.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServicesResult> InvokeAsync(GetServicesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServicesResult>("signalfx:aws/getServices:getServices", args ?? new GetServicesArgs(), options.WithVersion());
    }


    public sealed class GetServicesArgs : Pulumi.InvokeArgs
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