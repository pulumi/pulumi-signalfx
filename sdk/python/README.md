[![Build Status](https://travis-ci.com/pulumi/pulumi-signalfx.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-signalfx)

# SignalFX Resource Provider

The SignalFX resource provider for Pulumi lets you manage SignalFX resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/signalfx

or `yarn`:

    $ yarn add @pulumi/signalfx

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_signalfx

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-signalfx/sdk/go/...

## Configuration

The following configuration points are available:

- `signalfx:authToken` - (Required) The auth token for authentication. This can also be set via the `SFX_AUTH_TOKEN` 
  environment variable.
- `signalfx:apiUrl` - (Optional) The API URL to use for communicating with SignalFx. This is helpful for organizations 
  who need to set their Realm or use a proxy. Note: You likely want to change `customAppUrl` too!
- `signalfx:customAppUrl` - (Optional) The application URL that users should use to interact with assets in the browser.
  This is used by organizations using specific realms or those with a custom SSO domain.

## Reference

For detailed reference documentation, please visit [the API docs](https://pulumi.io/reference/pkg/nodejs/@pulumi/signalfx/index.html).
