CHANGELOG
=========

## Notice (2022-01-06)

*As of this notice, using CHANGELOG.md is DEPRECATED. We will be using [GitHub Releases](https://github.com/pulumi/pulumi-signalfx/releases) for this repository*

## HEAD (Unreleased)
_(none)_

---

## 5.1.1 (2021-12-20)
* Upgrade to v6.7.10 of the SignalFx Terraform Provider

## 5.1.0 (2021-11-11)
* Upgrade to terraform-bridge 3.11.0
* Upgrade to pulumi 3.17.0

## 5.0.5 (2021-10-19)
* Upgrade to v6.7.8 of the SignalFx Terraform Provider

## 5.0.4 (2021-09-27)
* Upgrade to v6.7.7 of the SignalFx Terraform Provider

## 5.0.3 (2021-09-09)
* Upgrade to v6.7.6 of the SignalFx Terraform Provider

## 5.0.2 (2021-07-20)
* Upgrade to v6.7.5 of the SignalFx Terraform Provider

## 5.0.1 (2021-05-21)
* Upgrade to v6.7.4 of the SignalFx Terraform Provider

## 5.0.0 (2021-04-19)
* Depend on Pulumi 3.0, which includes improvements to Python resource arguments and key translation, Go SDK performance,
  Node SDK performance, general availability of Automation API, and more.

## 4.6.0 (2021-04-12)
* Upgrade to pulumi-terraform-bridge v2.23.0

## 4.5.2 (2021-04-05)
* Upgrade to v6.7.3 of the SignalFx Terraform Provider

## 4.5.1 (2021-03-23)
* Upgrade to pulumi-terraform-bridge v2.22.1  
  **Please Note:** This includes a bug fix to the refresh operation regarding arrays

## 4.5.0 (2021-03-16)
* Upgrade to pulumi-terraform-bridge v2.21.0
* Release macOS arm64 binary

## 4.4.1 (2021-02-24)
* Upgrade to v6.7.1 of the SignalFx Terraform Provider

## 4.4.0 (2021-02-19)
* Upgrade to v6.7.0 of the SignalFx Terraform Provider

## 4.3.1 (2021-02-16)
* Upgrade to pulumi-terraform-bridge v2.19.0  
  **Please Note:** This includes a bug fix that stops mutating resources options in the nodejs provider
* Avoid storing config from the environment into the state

## 4.3.0 (2021-02-15)
* Upgrade to v6.6.0 of the SignalFx Terraform Provider

## 4.2.0 (2021-02-08)
* Upgrade to v6.5.0 of the SignalFx Terraform Provider

## 4.1.0 (2021-01-29)
* Upgrading pulumi-terraform-bridge to v2.18.0

## 4.0.1 (2021-01-13)
* Upgrade to pulumi-terraform-bridge v2.17.0
* Upgrade to Pulumi v2.17.0

## 4.0.0 (2020-12-21)
* Upgrade to v6.2.0 of the SignalFx Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.13.2  
  * This adds support for import specific examples in documentation

## 3.1.1 (2020-10-28)
* Upgrade to v5.0.2 of the SignalFx Terraform Provider

## 3.1.0 (2020-10-26)
* Upgrade to Pulumi v2.12.0 and pulumi-terraform-bridge v2.11.0
* Improving the accuracy of previews leading to a more accurate understanding of what will actually change rather than assuming all output properties will change.  
  ** PLEASE NOTE:**  
  This new preview functionality can be disabled by setting `PULUMI_DISABLE_PROVIDER_PREVIEW` to `1` or `false`.

## 3.0.1 (2020-09-25)
* Upgrade to v5.0.1 of the SignalFx Terraform Provider

## 3.0.0 (2020-09-14)
* Upgrade to v5.0.0 of the SignalFx Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.8.0
* Upgrade to Pulumi v2.10.0

## 2.6.0 (2020-08-31)
* Upgrade to pulumi-terraform-bridge v2.7.3
* Upgrade to Pulumi v2.9.0, which adds type annotations and input/output classes to Python

## 2.5.2 (2020-06-30)
* Upgrade to v4.23.2 of the SignalFx Terraform Provider

## 2.5.1 (2020-06-11)
* Upgrade to v4.23.1 of the SignalFx Terraform Provider
* Switch to GitHub actions for build

## 2.5.0 (2020-06-04)
* Upgrade to v4.23.0 of the SignalFx Terraform Provider

## 2.4.1 (2020-05-28)
* Upgrade to Pulumi v2.3.0
* Upgrade to pulumi-terraform-bridge v2.4.0

## 2.4.0 (2020-05-18)
* Upgrade to v4.21.0 of the SignalFx Terraform Provider
* Deprecate `signalFx.getAwsServices` in favor of `signalFx.aws.getServices`
* Deprecate `signalFx.getAzureServices` in favor of `signalFx.azure.getServices`

## 2.3.0 (2020-05-13)
* Upgrade to v4.20.1 of the SignalFx Terraform Provider
* Upgrade to pulumi-terraform-bridge v2.3.1

## 2.2.1 (2020-04-29)
* Upgrade to v4.19.5 of the SignalFx Terraform Provider

## 2.2.0 (2020-04-28)
* Upgrade to pulumi-terraform-bridge v2.1.0

## 2.1.2 (2020-04-24)
* Upgrade to v4.19.4 of the SignalFx Terraform Provider

## 2.1.1 (2020-04-23)
* Upgrade to v4.19.3 of the SignalFx Terraform Provider

## 2.1.0 (2020-04-20)
* Upgrade to v4.19.1 of the SignalFx Terraform Provider

## 2.0.0 (2020-04-18)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0

## 1.10.0 (2020-04-14)
* Upgrade to Pulumi v1.13.1
* Upgrade to pulumi-terraform-bridge v1.8.4
* Refactor layout to support Go modules

## 1.9.0 (2020-03-31)
* Upgrade to v4.18.6 of the SignalFx Terraform Provider

## 1.8.0 (2020-03-23)
* Upgrade to v4.18.4 of the SignalFx Terraform Provider
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.7.0 (2020-03-06)
* Upgrade to v4.18.0 of the SignalFx Terraform Provider

## 1.6.0 (2020-03-04)
* Upgrade to v4.17.0 of the SignalFx Terraform Provider

## 1.5.0 (2020-02-27)
* Upgrade to v4.15.0 of the SignalFx Terraform Provider

## 1.4.0 (2020-02-19)
* Upgrade to v4.13.0 of the SignalFx Terraform Provider

## 1.3.1 (2020-01-31)
* Upgrade to v4.12.2 of the SignalFx Terraform Provider

## 1.3.0 (2020-01-29)
* Upgrade to v4.12.0 of the SignalFx Terraform Provider
* Upgrade to pulumi-terraform-bridge v1.6.4

## 1.2.0 (2020-01-06)
* Upgrade to support go 1.13.x
* Upgrade to v4.11.0 of the SignalFx Terraform Provider
* Upgrade to pulumi-terraform-bridge v1.5.2
* Namespace names in .NET SDK are adjusted to PascalCase
([#6](https://github.com/pulumi/pulumi-signalfx/pull/6)).

## 1.1.0 (2019-11-15)
* Add support for DotNet SDK Generation

## 1.0.0 (2019-10-30)
* Initial release of the provider
