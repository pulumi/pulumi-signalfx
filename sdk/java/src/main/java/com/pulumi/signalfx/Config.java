// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.signalfx;

import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("signalfx");
/**
 * API URL for your SignalFx org, may include a realm
 * 
 */
    public Optional<String> apiUrl() {
        return Codegen.stringProp("apiUrl").config(config).get();
    }
/**
 * SignalFx auth token
 * 
 */
    public Optional<String> authToken() {
        return Codegen.stringProp("authToken").config(config).get();
    }
/**
 * Application URL for your SignalFx org, often customized for organizations using SSO
 * 
 */
    public Optional<String> customAppUrl() {
        return Codegen.stringProp("customAppUrl").config(config).get();
    }
/**
 * Max retries for a single HTTP call. Defaults to 4
 * 
 */
    public Optional<Integer> retryMaxAttempts() {
        return Codegen.integerProp("retryMaxAttempts").config(config).get();
    }
/**
 * Maximum retry wait for a single HTTP call in seconds. Defaults to 30
 * 
 */
    public Optional<Integer> retryWaitMaxSeconds() {
        return Codegen.integerProp("retryWaitMaxSeconds").config(config).get();
    }
/**
 * Minimum retry wait for a single HTTP call in seconds. Defaults to 1
 * 
 */
    public Optional<Integer> retryWaitMinSeconds() {
        return Codegen.integerProp("retryWaitMinSeconds").config(config).get();
    }
/**
 * Timeout duration for a single HTTP call in seconds. Defaults to 120
 * 
 */
    public Optional<Integer> timeoutSeconds() {
        return Codegen.integerProp("timeoutSeconds").config(config).get();
    }
}
