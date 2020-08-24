[![Actions Status](https://github.com/pulumi/pulumi-signalfx/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-signalfx/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fsignalfx.svg)](https://www.npmjs.com/package/@pulumi/signalfx)
[![Python version](https://badge.fury.io/py/pulumi-signalfx.svg)](https://pypi.org/project/pulumi-signalfx)
[![NuGet version](https://badge.fury.io/nu/pulumi.signalfx.svg)](https://badge.fury.io/nu/pulumi.signalfx)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-signalfx/sdk/v2/go)](https://pkg.go.dev/github.com/pulumi/pulumi-signalfx/sdk/v2/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-signalfx/blob/master/LICENSE)

# SignalFx Resource Provider

The SignalFx resource provider for Pulumi lets you manage SignalFx resources in your cloud programs. To use
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

    $ go get github.com/pulumi/pulumi-signalfx/sdk/v2/go/...

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Signalfx

## Configuration

The following configuration points are available:

- `signalfx:authToken` - (Required) The auth token for authentication. This can also be set via the `SFX_AUTH_TOKEN` 
  environment variable.
- `signalfx:apiUrl` - (Optional) The API URL to use for communicating with SignalFx. This is helpful for organizations 
  who need to set their Realm or use a proxy. Note: You likely want to change `customAppUrl` too!
- `signalfx:customAppUrl` - (Optional) The application URL that users should use to interact with assets in the browser.
  This is used by organizations using specific realms or those with a custom SSO domain.

## Reference

For further information, please visit [the SignalFx provider docs](https://www.pulumi.com/docs/intro/cloud-providers/signalfx) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/signalfx).
