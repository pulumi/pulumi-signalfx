// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.SignalFx
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("signalfx");

        private static readonly __Value<string?> _apiUrl = new __Value<string?>(() => __config.Get("apiUrl"));
        /// <summary>
        /// API URL for your SignalFx org, may include a realm
        /// </summary>
        public static string? ApiUrl
        {
            get => _apiUrl.Get();
            set => _apiUrl.Set(value);
        }

        private static readonly __Value<string?> _authToken = new __Value<string?>(() => __config.Get("authToken"));
        /// <summary>
        /// SignalFx auth token
        /// </summary>
        public static string? AuthToken
        {
            get => _authToken.Get();
            set => _authToken.Set(value);
        }

        private static readonly __Value<string?> _customAppUrl = new __Value<string?>(() => __config.Get("customAppUrl"));
        /// <summary>
        /// Application URL for your SignalFx org, often customized for organizations using SSO
        /// </summary>
        public static string? CustomAppUrl
        {
            get => _customAppUrl.Get();
            set => _customAppUrl.Set(value);
        }

        private static readonly __Value<int?> _retryMaxAttempts = new __Value<int?>(() => __config.GetInt32("retryMaxAttempts"));
        /// <summary>
        /// Max retries for a single HTTP call. Defaults to 4
        /// </summary>
        public static int? RetryMaxAttempts
        {
            get => _retryMaxAttempts.Get();
            set => _retryMaxAttempts.Set(value);
        }

        private static readonly __Value<int?> _retryWaitMaxSeconds = new __Value<int?>(() => __config.GetInt32("retryWaitMaxSeconds"));
        /// <summary>
        /// Maximum retry wait for a single HTTP call in seconds. Defaults to 30
        /// </summary>
        public static int? RetryWaitMaxSeconds
        {
            get => _retryWaitMaxSeconds.Get();
            set => _retryWaitMaxSeconds.Set(value);
        }

        private static readonly __Value<int?> _retryWaitMinSeconds = new __Value<int?>(() => __config.GetInt32("retryWaitMinSeconds"));
        /// <summary>
        /// Minimum retry wait for a single HTTP call in seconds. Defaults to 1
        /// </summary>
        public static int? RetryWaitMinSeconds
        {
            get => _retryWaitMinSeconds.Get();
            set => _retryWaitMinSeconds.Set(value);
        }

        private static readonly __Value<int?> _timeoutSeconds = new __Value<int?>(() => __config.GetInt32("timeoutSeconds"));
        /// <summary>
        /// Timeout duration for a single HTTP call in seconds. Defaults to 120
        /// </summary>
        public static int? TimeoutSeconds
        {
            get => _timeoutSeconds.Get();
            set => _timeoutSeconds.Set(value);
        }

    }
}
