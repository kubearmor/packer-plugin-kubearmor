# KubeArmor Plugins

A plugin for Packer which provides [KubeArmor Hardening Host Security policies](https://docs.kubearmor.io/kubearmor/use-cases/hardening_guide) for the build workloads.

> [KubeArmor](https://docs.kubearmor.io/kubearmor/) is a security solution for the Kubernetes and cloud native platforms that helps protect your workloads from attacks and threats. It does this by providing a set of hardening policies that are based on industry-leading compliance and attack frameworks such as CIS, MITRE, NIST-800-53, and STIGs. These policies are designed to help you secure your workloads in a way that is compliant with these frameworks and recommended best practices.

The KubeArmor provisioner will configure the build and provide the [KubeArmor Hardening Host Security Policies](https://docs.kubearmor.io/kubearmor/use-cases/hardening_guide). These hardening policies will be in the context of your workload, so you can see how they will be applied and what impact they will have on your system. This allows you to make informed decisions about which policies to apply, and helps you understand the trade-offs between security and functionality.

<!--
  Include a short overview about the plugin.

  This document is a great location for creating a table of contents for each
  of the components the plugin may provide. This document should load automatically
  when navigating to the docs directory for a plugin.

-->

## Installation

### Using pre-built releases

#### Using the `packer init` command

Starting from version 1.7, Packer supports a new `packer init` command allowing
automatic installation of Packer plugins. Read the
[Packer documentation](https://www.packer.io/docs/commands/init) for more information.

To install this plugin, copy and paste this code into your Packer configuration .
Then, run [`packer init`](https://www.packer.io/docs/commands/init).

```hcl
packer {
  required_plugins {
    kubearmor = {
      version = ">= 0.0.1"
      source  = "github.com/hashicorp/kubearmor"
    }
  }
}
```

#### Manual installation

You can find pre-built binary releases of the plugin [here](https://github.com/hashicorp/packer-plugin-name/releases).
Once you have downloaded the latest archive corresponding to your target OS,
uncompress it to retrieve the plugin binary file corresponding to your platform.
To install the plugin, please follow the Packer documentation on
[installing a plugin](https://www.packer.io/docs/extending/plugins/#installing-plugins).


#### From Source

If you prefer to build the plugin from its source code, clone the GitHub
repository locally and run the command `go build` from the root
directory. Upon successful compilation, a `packer-plugin-name` plugin
binary file can be found in the root directory.
To install the compiled plugin, please follow the official Packer documentation
on [installing a plugin](https://developer.hashicorp.com/packer/docs/plugins/install-plugins).


## Plugin Contents

### Provisioners

- [kubearmor](/packer/integrations/KubeArmor/latest/components/provisioner/kubearmor) - The kubearmor provisioner is used to provisioner
  Packer builds and provide the packer builds KubeArmor Host Security Policies.
